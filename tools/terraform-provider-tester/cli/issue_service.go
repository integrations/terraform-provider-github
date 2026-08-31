package cli

import (
	"context"
	"errors"
	"fmt"

	"github.com/github/terraform-provider-tester/engine"
	"github.com/github/terraform-provider-tester/internal/redact"
	ghissues "github.com/github/terraform-provider-tester/provider/github"
)

type issuePreviewResult struct {
	Fingerprint      string
	ShortFingerprint string
	IssuesRepo       string
	Title            string
	Labels           []string
	Classification   string
}

type issueFileResult struct {
	Fingerprint string
	IssueNumber int
	Action      string
	Failures    []engine.PersistFailure
}

type issueService struct {
	root            string
	statePath       string
	issuesRepo      string
	knownIssuesPath string
	red             *redact.Redactor
	registryFactory func(triageOptions) (issueRegistry, error)
	filerFactory    func(string) (issueFiler, error)
	getenv          func(string) string
	saveState       func(*engine.State, string) error
	acquireLock     func(string) (func() error, error)
}

func (s issueService) Preview(fingerprint string) (result issuePreviewResult, err error) {
	release, err := s.lockState()
	if err != nil {
		return issuePreviewResult{}, fmt.Errorf("acquiring lock: %w", err)
	}
	defer func() {
		if releaseErr := release(); releaseErr != nil {
			err = errors.Join(err, fmt.Errorf("releasing lock: %w", releaseErr))
		}
	}()

	st, err := engine.Load(s.statePath)
	if err != nil {
		return issuePreviewResult{}, fmt.Errorf("loading state: %w", err)
	}
	f, ok := findFailure(st.Failures, fingerprint)
	if !ok || !shouldFileIssue(f) {
		return issuePreviewResult{}, errors.New("selected failure is not eligible for issue filing")
	}
	red := s.redactor()
	draft := ghissues.BuildFailureIssueDraft(f, ghissues.IssueDraftOptions{
		IssuesRepo: s.repo(),
		LogLines:   readFailureLogLines(s.root, f),
		Redactor:   red,
	})
	return issuePreviewResult{
		Fingerprint:      f.Fingerprint,
		ShortFingerprint: f.ShortFingerprint,
		IssuesRepo:       s.repo(),
		Title:            red.String(draft.Title),
		Labels:           append([]string(nil), draft.Labels...),
		Classification:   f.Classification,
	}, nil
}

func (s issueService) FileOne(ctx context.Context, fingerprint, phrase string) (result issueFileResult, err error) {
	release, err := s.lockState()
	if err != nil {
		return issueFileResult{}, fmt.Errorf("acquiring lock: %w", err)
	}

	defer func() {
		if releaseErr := release(); releaseErr != nil {
			releaseErr = fmt.Errorf("releasing lock: %w", releaseErr)
			if err != nil {
				err = errors.Join(err, releaseErr)
				return
			}
			err = releaseErr
		}
	}()

	st, err := engine.Load(s.statePath)
	if err != nil {
		return issueFileResult{}, fmt.Errorf("loading state: %w", err)
	}
	f, ok := findFailure(st.Failures, fingerprint)
	if !ok {
		return issueFileResult{}, errors.New("selected failure is not eligible for issue filing")
	}
	if phrase != "FILE "+f.ShortFingerprint {
		return issueFileResult{}, errors.New("confirmation phrase does not match selected failure")
	}
	clearRefreshableKnownIssue(&f)
	if !shouldFileIssue(f) {
		return issueFileResult{}, errors.New("selected failure is not eligible for issue filing")
	}

	opts := triageOptions{IssuesRepo: s.repo(), KnownIssuesPath: s.knownIssuesPath}
	reg, err := s.liveRegistry(opts)
	if err != nil {
		return issueFileResult{}, err
	}
	match, err := reg.LookupFingerprint(ctx, f.Fingerprint, st.Mode)
	if err != nil {
		return issueFileResult{}, err
	}
	if match.Found && match.Suppress {
		if match.Issue != 0 {
			f.KnownIssue = match.Issue
		}
		f.IssueAction = "known"
		if match.Note != "" {
			f.Reasons = appendUniqueReason(f.Reasons, match.Note)
		}
		return s.saveSelected(&st, f)
	}

	filer, err := issueFilerForFactory(s.repo(), s.filerFactory)
	if err != nil {
		return issueFileResult{}, err
	}
	dedup, err := filer.FindIssueByFingerprint(ctx, f.Fingerprint)
	if err != nil {
		return issueFileResult{}, err
	}
	if dedup != nil {
		applyIssueDedupMatch(&f, dedup)
		return s.saveSelected(&st, f)
	}

	draft := ghissues.BuildFailureIssueDraft(f, ghissues.IssueDraftOptions{
		IssuesRepo: s.repo(),
		LogLines:   readFailureLogLines(s.root, f),
		Redactor:   s.redactor(),
	})
	created, err := filer.CreateFailureIssue(ctx, draft)
	if err != nil {
		return issueFileResult{}, err
	}
	if created == nil || created.Number <= 0 {
		return issueFileResult{}, errors.New("created issue did not return a valid issue number")
	}
	f.KnownIssue = created.Number
	f.IssueAction = "filed"
	return s.saveSelected(&st, f)
}

func (s issueService) lockState() (func() error, error) {
	if s.acquireLock != nil {
		return s.acquireLock(s.statePath + ".lock")
	}
	return engine.AcquireLock(s.statePath + ".lock")
}

func (s issueService) saveSelected(st *engine.State, updated engine.PersistFailure) (issueFileResult, error) {
	st.Failures = replaceFailure(st.Failures, updated)
	if err := s.save(st, s.statePath); err != nil {
		return issueFileResult{}, fmt.Errorf("saving state: %w", err)
	}
	return issueFileResult{
		Fingerprint: updated.Fingerprint,
		IssueNumber: updated.KnownIssue,
		Action:      updated.IssueAction,
		Failures:    append([]engine.PersistFailure(nil), st.Failures...),
	}, nil
}

func (s issueService) save(st *engine.State, path string) error {
	if s.saveState != nil {
		return s.saveState(st, path)
	}
	return st.Save(path)
}

func (s issueService) liveRegistry(opts triageOptions) (issueRegistry, error) {
	if s.registryFactory != nil {
		return s.registryFactory(opts)
	}
	return buildLiveIssueRegistry(opts)
}

func (s issueService) repo() string {
	if s.issuesRepo != "" {
		return s.issuesRepo
	}
	return defaultIssuesRepo
}

func (s issueService) redactor() *redact.Redactor {
	if s.red != nil {
		return s.red
	}
	return redact.New(nil)
}

func findFailure(all []engine.PersistFailure, fingerprint string) (engine.PersistFailure, bool) {
	for _, f := range all {
		if f.Fingerprint == fingerprint {
			return f, true
		}
	}
	return engine.PersistFailure{}, false
}

func replaceFailure(all []engine.PersistFailure, updated engine.PersistFailure) []engine.PersistFailure {
	out := append([]engine.PersistFailure(nil), all...)
	for i := range out {
		if out[i].Fingerprint == updated.Fingerprint {
			out[i] = updated
			return out
		}
	}
	return out
}
