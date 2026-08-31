package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/github/terraform-provider-tester/engine"
	"github.com/github/terraform-provider-tester/internal/redact"
	gogithub "github.com/google/go-github/v88/github"
	"gopkg.in/yaml.v3"
)

const issueLabelAcctestFailure = "acctest-failure"

type IssueClient struct {
	client   *gogithub.Client
	owner    string
	repo     string
	redactor *redact.Redactor
}

type IssueMatch struct {
	Number int
	State  string
	Title  string
	URL    string
}

type IssueResult struct {
	Number int
	URL    string
}

type IssueDraft struct {
	Title  string
	Body   string
	Labels []string
}

type IssueDraftOptions struct {
	IssuesRepo string
	LogLines   []string
	Redactor   *redact.Redactor
}

func NewIssueClient(apiBase, token, repo string) (*IssueClient, error) {
	owner, name, err := splitRepo(repo)
	if err != nil {
		return nil, err
	}
	return &IssueClient{
		client:   newClient(apiBase, token),
		owner:    owner,
		repo:     name,
		redactor: redact.New([]string{token}),
	}, nil
}

func NewIssueClientFromEnv(repo string) (*IssueClient, error) {
	return NewIssueClient(os.Getenv("GITHUB_BASE_URL"), os.Getenv("GITHUB_TOKEN"), repo)
}

func (c *IssueClient) ListKnownIssues(ctx context.Context) ([]engine.KnownIssueEntry, error) {
	opts := &gogithub.IssueListByRepoOptions{
		State:  "open",
		Labels: []string{issueLabelAcctestFailure},
		ListOptions: gogithub.ListOptions{
			PerPage: 100,
		},
	}
	var entries []engine.KnownIssueEntry
	for {
		issues, resp, err := c.client.Issues.ListByRepo(ctx, c.owner, c.repo, opts)
		if err != nil {
			return nil, c.redactedError(err)
		}
		for _, issue := range issues {
			entry, ok := parseKnownIssueBlock(issue.GetBody())
			if !ok {
				continue
			}
			entry.Issue = issue.GetNumber()
			entry.State = issue.GetState()
			entry.Title = issue.GetTitle()
			entry.IssueURL = issue.GetHTMLURL()
			entries = append(entries, entry)
		}
		if resp == nil || resp.NextPage == 0 {
			break
		}
		opts.ListOptions.Page = resp.NextPage
	}
	return entries, nil
}

func (c *IssueClient) FindIssueByFingerprint(ctx context.Context, fingerprint string) (*IssueMatch, error) {
	query := fmt.Sprintf(`repo:%s/%s is:issue "%s"`, c.owner, c.repo, fingerprint)
	result, _, err := c.client.Search.Issues(ctx, query, &gogithub.SearchOptions{
		ListOptions: gogithub.ListOptions{PerPage: 10},
	})
	if err != nil {
		return nil, c.redactedError(err)
	}
	var firstClosed *IssueMatch
	for _, issue := range result.Issues {
		match := &IssueMatch{
			Number: issue.GetNumber(),
			State:  issue.GetState(),
			Title:  issue.GetTitle(),
			URL:    issue.GetHTMLURL(),
		}
		if strings.EqualFold(match.State, "open") {
			return match, nil
		}
		if firstClosed == nil {
			firstClosed = match
		}
	}
	return firstClosed, nil
}

func (c *IssueClient) CreateFailureIssue(ctx context.Context, draft IssueDraft) (*IssueResult, error) {
	title := c.redactor.String(draft.Title)
	body := c.redactor.String(draft.Body)
	labels := append([]string{}, draft.Labels...)
	issue, err := c.createIssue(ctx, title, body, labels)
	if err != nil {
		if ctxErr := ctx.Err(); ctxErr != nil {
			return nil, ctxErr
		}
		if shouldRetryIssueCreateWithGuaranteedLabels(err, labels) {
			issue, err = c.createIssue(ctx, title, body, labels[:3])
		}
		if err != nil {
			if ctxErr := ctx.Err(); ctxErr != nil {
				return nil, ctxErr
			}
			return nil, c.redactedError(err)
		}
	}
	return &IssueResult{Number: issue.GetNumber(), URL: issue.GetHTMLURL()}, nil
}

func shouldRetryIssueCreateWithGuaranteedLabels(err error, labels []string) bool {
	if len(labels) <= 3 {
		return false
	}
	var respErr *gogithub.ErrorResponse
	if !errors.As(err, &respErr) || respErr.Response == nil || respErr.Response.StatusCode != http.StatusUnprocessableEntity {
		return false
	}
	if isMissingOrInvalidLabelText(respErr.Message) || isMissingOrInvalidLabelText(err.Error()) {
		return true
	}
	for _, detail := range respErr.Errors {
		field := strings.ToLower(detail.Field)
		code := strings.ToLower(detail.Code)
		if field != "labels" && field != "label" && !isMissingOrInvalidLabelText(detail.Message) {
			continue
		}
		if code == "missing" || code == "invalid" || code == "missing_field" || isMissingOrInvalidLabelText(detail.Message) {
			return true
		}
	}
	return false
}

func isMissingOrInvalidLabelText(text string) bool {
	lower := strings.ToLower(text)
	return strings.Contains(lower, "label") &&
		(strings.Contains(lower, "does not exist") ||
			strings.Contains(lower, "not found") ||
			strings.Contains(lower, "missing") ||
			strings.Contains(lower, "invalid"))
}

func (c *IssueClient) createIssue(ctx context.Context, title, body string, labels []string) (*gogithub.Issue, error) {
	reqLabels := append([]string{}, labels...)
	req := &gogithub.IssueRequest{Title: &title, Body: &body, Labels: &reqLabels}
	issue, _, err := c.client.Issues.Create(ctx, c.owner, c.repo, req)
	return issue, err
}

func BuildFailureIssueDraft(f engine.PersistFailure, opts IssueDraftOptions) IssueDraft {
	red := opts.Redactor
	if red == nil {
		red = redact.New(nil)
	}
	short := f.ShortFingerprint
	if short == "" && strings.HasPrefix(f.Fingerprint, "sha256:") && len(f.Fingerprint) >= len("sha256:")+16 {
		short = strings.TrimPrefix(f.Fingerprint, "sha256:")[:16]
	}
	testName := f.Test
	if testName == "" {
		testName = "suite"
	}
	title := red.String(fmt.Sprintf("[acctest] %s %s (%s)", testName, f.Class, short))
	labels := issueLabelsForClass(f.Class)
	body := red.String(renderIssueBody(f, short, redactedExcerpt(opts.LogLines, red)))
	return IssueDraft{Title: title, Body: body, Labels: labels}
}

func renderIssueBody(f engine.PersistFailure, short, excerpt string) string {
	testName := f.Test
	if testName == "" {
		testName = "suite"
	}
	retry := "never"
	if f.Retryable {
		retry = "retryable"
	}
	retries := max(f.Attempts-1, 0)
	return fmt.Sprintf(`## Summary

terraform-provider-tester saw an acceptance-test failure and classified it as a real unknown failure after applying the retry policy.

## Fingerprint

`+"```yaml"+`
version: 1
fingerprint: %s
short: %s
class: %s
canonical: %q
mode: %s
package: %s
test: %s
subtest: %q
status: %s
retryable: %t
classification: %s
`+"```"+`

<!-- pulsar-known-issue
version: 1
fingerprint: %s
short: %s
mode: known-real
test: %s
classes:
  - %s
modes:
  - %s
retry: %s
-->

## Retry decision

- Initial attempt: failed.
- Attempts made: %d.
- Retryable: %t.
- Result: every attempt produced the same fingerprint.
- Known issue lookup: no open issue found with this fingerprint.

## Redacted log excerpt

`+"```text"+`
%s
`+"```"+`

If this excerpt contains `+"`***REDACTED***`"+`, treat the hidden value as sensitive. Do not ask anyone to paste the original secret.

## Reproduce

`+"```sh"+`
terraform-provider-tester run --mode %s --run '^%s$' --retries %d --triage --repo-root <provider checkout>
`+"```"+`

Equivalent raw command shape:

`+"```sh"+`
TF_ACC=1 CGO_ENABLED=0 go test ./github/... -json -run '^%s$' -timeout 120m -count=1
`+"```"+`

## AI-assisted disclosure

This issue was drafted by terraform-provider-tester with AI-assisted classification. A human maintainer must review it, confirm the excerpt contains no secrets, and verify the failure before acting on it.
`, f.Fingerprint, short, f.Class, f.Canonical, f.Mode, f.Package, testName, f.Sub, f.Status, f.Retryable, f.Classification,
		f.Fingerprint, short, testName, f.Class, f.Mode, retry, f.Attempts, f.Retryable, excerpt, f.Mode, testName, retries, testName)
}

func redactedExcerpt(lines []string, red *redact.Redactor) string {
	if len(lines) > 200 {
		lines = lines[:200]
	}
	joined := red.String(strings.Join(lines, "\n"))
	if len(joined) > 16*1024 {
		joined = string([]byte(joined)[:16*1024])
	}
	return joined
}

func issueLabelsForClass(class string) []string {
	labels := []string{issueLabelAcctestFailure, "tester-filed", "ai-assisted"}
	switch class {
	case engine.ClassLeftoverState:
		labels = append(labels, "leftover-state")
	case engine.ClassRateLimit, engine.ClassSecondaryRateLimit:
		labels = append(labels, "rate-limit")
	case engine.ClassTransientNetwork:
		labels = append(labels, "transient-network")
	}
	return labels
}

func parseKnownIssueBlock(body string) (engine.KnownIssueEntry, bool) {
	start := strings.Index(body, "<!-- pulsar-known-issue")
	if start < 0 {
		return engine.KnownIssueEntry{}, false
	}
	contentStart := start + len("<!-- pulsar-known-issue")
	end := strings.Index(body[contentStart:], "-->")
	if end < 0 {
		return engine.KnownIssueEntry{}, false
	}
	block := strings.TrimSpace(body[contentStart : contentStart+end])
	var parsed struct {
		Version     int      `yaml:"version"`
		Fingerprint string   `yaml:"fingerprint"`
		Short       string   `yaml:"short"`
		Mode        string   `yaml:"mode"`
		Test        string   `yaml:"test"`
		Classes     []string `yaml:"classes"`
		Modes       []string `yaml:"modes"`
		Retry       string   `yaml:"retry"`
	}
	if err := yaml.Unmarshal([]byte(block), &parsed); err != nil || parsed.Fingerprint == "" {
		return engine.KnownIssueEntry{}, false
	}
	entry := engine.KnownIssueEntry{
		Fingerprint: parsed.Fingerprint,
		Short:       parsed.Short,
		Mode:        parsed.Mode,
		Classes:     parsed.Classes,
		Modes:       parsed.Modes,
		Retry:       parsed.Retry,
	}
	if parsed.Test != "" {
		entry.Tests = []string{parsed.Test}
	}
	return entry, true
}

func splitRepo(repo string) (string, string, error) {
	parts := strings.Split(repo, "/")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return "", "", fmt.Errorf("issues repo must be owner/repo, got %q", repo)
	}
	return parts[0], parts[1], nil
}

func SplitIssueRepoForCLI(repo string) (string, string, error) {
	return splitRepo(repo)
}

func (c *IssueClient) redactedError(err error) error {
	if err == nil {
		return nil
	}
	msg := c.redactor.String(err.Error())
	if msg == "" {
		return errors.New("github issue request failed")
	}
	return errors.New(msg)
}
