package github

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/github/terraform-provider-tester/engine"
	"github.com/github/terraform-provider-tester/internal/redact"
)

func TestListKnownIssuesParsesFingerprintBlocks(t *testing.T) {
	fp := "sha256:" + strings.Repeat("d", 64)
	mux := http.NewServeMux()
	mux.HandleFunc("/repos/integrations/terraform-provider-github/issues", func(w http.ResponseWriter, r *http.Request) {
		if got := r.URL.Query().Get("labels"); got != "acctest-failure" {
			t.Fatalf("labels query = %q, want acctest-failure", got)
		}
		if got := r.URL.Query().Get("state"); got != "open" {
			t.Fatalf("state query = %q, want open", got)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[{
			"number": 42,
			"state": "open",
			"title": "[acctest] TestAccThing api/422-leftover-state",
			"html_url": "https://github.com/integrations/terraform-provider-github/issues/42",
			"body": "text\n<!-- pulsar-known-issue\nversion: 1\nfingerprint: %s\nshort: %s\nmode: known-real\ntest: TestAccThing\nclasses:\n  - api/422-leftover-state\nmodes:\n  - organization\nretry: retryable\n-->\n"
		}]`, fp, strings.Repeat("d", 16))
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	client, err := NewIssueClient(srv.URL+"/", "ghp_FAKEFAKEFAKEFAKEFAKEFAKE", "integrations/terraform-provider-github")
	if err != nil {
		t.Fatalf("NewIssueClient error: %v", err)
	}
	entries, err := client.ListKnownIssues(context.Background())
	if err != nil {
		t.Fatalf("ListKnownIssues error: %v", err)
	}
	if len(entries) != 1 {
		t.Fatalf("entries = %d, want 1", len(entries))
	}
	entry := entries[0]
	if entry.Fingerprint != fp || entry.Issue != 42 || entry.Mode != "known-real" || entry.Retry != "retryable" {
		t.Fatalf("entry not parsed from block: %+v", entry)
	}
	if got := strings.Join(entry.Classes, ","); got != "api/422-leftover-state" {
		t.Fatalf("classes = %q", got)
	}
}

func TestFindIssueByFingerprintUsesExactSearch(t *testing.T) {
	fp := "sha256:" + strings.Repeat("e", 64)
	mux := http.NewServeMux()
	mux.HandleFunc("/search/issues", func(w http.ResponseWriter, r *http.Request) {
		want := `repo:integrations/terraform-provider-github is:issue "` + fp + `"`
		if got := r.URL.Query().Get("q"); got != want {
			t.Fatalf("query = %q, want %q", got, want)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"total_count":1,"items":[{"number":77,"state":"closed","title":"old","html_url":"https://github.com/integrations/terraform-provider-github/issues/77"}]}`)
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	client, err := NewIssueClient(srv.URL+"/", "ghp_FAKEFAKEFAKEFAKEFAKEFAKE", "integrations/terraform-provider-github")
	if err != nil {
		t.Fatalf("NewIssueClient error: %v", err)
	}
	match, err := client.FindIssueByFingerprint(context.Background(), fp)
	if err != nil {
		t.Fatalf("FindIssueByFingerprint error: %v", err)
	}
	if match == nil || match.Number != 77 || match.State != "closed" {
		t.Fatalf("match = %+v, want closed issue #77", match)
	}
}

func TestBuildFailureIssueDraftRedactsAndCapsBody(t *testing.T) {
	secret := "ghp_FAKEFAKEFAKEFAKEFAKEFAKE"
	failure := engine.PersistFailure{
		Package: "./github", Test: "TestAccThing", Status: "fail",
		Fingerprint: "sha256:" + strings.Repeat("f", 64), ShortFingerprint: strings.Repeat("f", 16),
		Class: engine.ClassLeftoverState, Canonical: "422 leftover state " + secret,
		Retryable: true, Classification: engine.ClassificationReal, Attempts: 3, Mode: "organization",
	}
	var lines []string
	for i := 0; i < 220; i++ {
		lines = append(lines, fmt.Sprintf("line %03d %s", i, secret))
	}
	draft := BuildFailureIssueDraft(failure, IssueDraftOptions{
		IssuesRepo: "integrations/terraform-provider-github",
		LogLines:   lines,
		Redactor:   redact.New([]string{secret}),
	})

	combined := draft.Title + "\n" + strings.Join(draft.Labels, "\n") + "\n" + draft.Body
	if strings.Contains(combined, secret) {
		t.Fatalf("draft leaked token-shaped value: %s", combined)
	}
	for _, want := range []string{
		"[acctest] TestAccThing api/422-leftover-state (ffffffffffffffff)",
		"acctest-failure",
		"tester-filed",
		"ai-assisted",
		"<!-- pulsar-known-issue",
		"AI-assisted disclosure",
		"human maintainer must review",
		"***REDACTED***",
	} {
		if !strings.Contains(combined, want) {
			t.Fatalf("draft missing %q:\n%s", want, combined)
		}
	}
	if got := strings.Count(draft.Body, "line "); got != 200 {
		t.Fatalf("body excerpt lines = %d, want 200", got)
	}
	if len(draft.Body) > 16*1024+4096 {
		t.Fatalf("body too large after excerpt cap: %d bytes", len(draft.Body))
	}
}

func TestBuildFailureIssueDraftUsesObservedRetryFacts(t *testing.T) {
	failure := engine.PersistFailure{
		Package: "./github", Test: "TestAccThing", Status: "fail",
		Fingerprint: "sha256:" + strings.Repeat("f", 64), ShortFingerprint: strings.Repeat("f", 16),
		Class: engine.ClassLeftoverState, Canonical: "422 leftover state",
		Retryable: true, Classification: engine.ClassificationReal, Attempts: 1, Mode: "organization",
	}
	draft := BuildFailureIssueDraft(failure, IssueDraftOptions{})
	for _, want := range []string{
		"<!-- pulsar-known-issue",
		"- Attempts made: 1.",
		"- Retryable: true.",
		"terraform-provider-tester run --mode organization --run '^TestAccThing$' --retries 0 --triage",
	} {
		if !strings.Contains(draft.Body, want) {
			t.Fatalf("body missing %q:\n%s", want, draft.Body)
		}
	}
	for _, notWant := range []string{"2 retries", "pulsar run"} {
		if strings.Contains(draft.Body, notWant) {
			t.Fatalf("body contains obsolete retry fact or command %q:\n%s", notWant, draft.Body)
		}
	}

	failure.Attempts = 3
	draft = BuildFailureIssueDraft(failure, IssueDraftOptions{})
	if !strings.Contains(draft.Body, "terraform-provider-tester run --mode organization --run '^TestAccThing$' --retries 2 --triage") {
		t.Fatalf("body missing observed two-retry reproduce command:\n%s", draft.Body)
	}
}

func TestCreateFailureIssueRedactsCreateErrors(t *testing.T) {
	secret := "ghp_FAKEFAKEFAKEFAKEFAKEFAKE"
	mux := http.NewServeMux()
	mux.HandleFunc("/repos/integrations/terraform-provider-github/issues", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("method = %s, want POST", r.Method)
		}
		http.Error(w, `{"message":"create failed `+secret+`"}`, http.StatusInternalServerError)
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	client, err := NewIssueClient(srv.URL+"/", secret, "integrations/terraform-provider-github")
	if err != nil {
		t.Fatalf("NewIssueClient error: %v", err)
	}
	_, err = client.CreateFailureIssue(context.Background(), IssueDraft{
		Title:  "safe title",
		Body:   "safe body",
		Labels: []string{"acctest-failure", "tester-filed", "ai-assisted"},
	})
	if err == nil {
		t.Fatal("expected create error")
	}
	if strings.Contains(err.Error(), secret) {
		t.Fatalf("create error leaked token: %v", err)
	}
	if !strings.Contains(err.Error(), "***REDACTED***") {
		t.Fatalf("create error should contain redaction marker, got: %v", err)
	}
}

func TestCreateFailureIssueDoesNotRetryNonValidationError(t *testing.T) {
	calls := 0
	mux := http.NewServeMux()
	mux.HandleFunc("/repos/integrations/terraform-provider-github/issues", func(w http.ResponseWriter, r *http.Request) {
		calls++
		if r.Method != http.MethodPost {
			t.Fatalf("method = %s, want POST", r.Method)
		}
		http.Error(w, `{"message":"server failed"}`, http.StatusInternalServerError)
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	client, err := NewIssueClient(srv.URL+"/", "ghp_FAKEFAKEFAKEFAKEFAKEFAKE", "integrations/terraform-provider-github")
	if err != nil {
		t.Fatalf("NewIssueClient error: %v", err)
	}
	_, err = client.CreateFailureIssue(context.Background(), IssueDraft{
		Title:  "safe title",
		Body:   "safe body",
		Labels: []string{"acctest-failure", "tester-filed", "ai-assisted", "leftover-state"},
	})
	if err == nil {
		t.Fatal("expected create error")
	}
	if calls != 1 {
		t.Fatalf("create calls = %d, want 1 for non-validation error", calls)
	}
}

func TestCreateFailureIssueReturnsContextCancellationWithoutRetry(t *testing.T) {
	calls := 0
	mux := http.NewServeMux()
	mux.HandleFunc("/repos/integrations/terraform-provider-github/issues", func(w http.ResponseWriter, _ *http.Request) {
		calls++
		http.Error(w, `{"message":"should not be called"}`, http.StatusInternalServerError)
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	client, err := NewIssueClient(srv.URL+"/", "ghp_FAKEFAKEFAKEFAKEFAKEFAKE", "integrations/terraform-provider-github")
	if err != nil {
		t.Fatalf("NewIssueClient error: %v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err = client.CreateFailureIssue(ctx, IssueDraft{
		Title:  "safe title",
		Body:   "safe body",
		Labels: []string{"acctest-failure", "tester-filed", "ai-assisted", "leftover-state"},
	})
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("error = %v, want context.Canceled", err)
	}
	if calls != 0 {
		t.Fatalf("create calls = %d, want 0 for canceled context", calls)
	}
}

func TestCreateFailureIssueRetriesWithoutOptionalClassLabel(t *testing.T) {
	calls := 0
	mux := http.NewServeMux()
	mux.HandleFunc("/repos/integrations/terraform-provider-github/issues", func(w http.ResponseWriter, r *http.Request) {
		calls++
		var req struct {
			Labels []string `json:"labels"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("decode create request: %v", err)
		}
		if calls == 1 {
			if got := strings.Join(req.Labels, ","); got != "acctest-failure,tester-filed,ai-assisted,leftover-state" {
				t.Fatalf("first labels = %q", got)
			}
			http.Error(w, `{"message":"label does not exist"}`, http.StatusUnprocessableEntity)
			return
		}
		if got := strings.Join(req.Labels, ","); got != "acctest-failure,tester-filed,ai-assisted" {
			t.Fatalf("retry labels = %q", got)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, `{"number":88,"html_url":"https://github.com/integrations/terraform-provider-github/issues/88"}`)
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	client, err := NewIssueClient(srv.URL+"/", "ghp_FAKEFAKEFAKEFAKEFAKEFAKE", "integrations/terraform-provider-github")
	if err != nil {
		t.Fatalf("NewIssueClient error: %v", err)
	}
	result, err := client.CreateFailureIssue(context.Background(), IssueDraft{
		Title:  "safe title",
		Body:   "safe body",
		Labels: []string{"acctest-failure", "tester-filed", "ai-assisted", "leftover-state"},
	})
	if err != nil {
		t.Fatalf("CreateFailureIssue should retry without optional label: %v", err)
	}
	if result.Number != 88 || calls != 2 {
		t.Fatalf("result=%+v calls=%d, want issue 88 after retry", result, calls)
	}
}

func TestNewIssueClientFromEnvUsesGHESBaseURL(t *testing.T) {
	calls := 0
	mux := http.NewServeMux()
	mux.HandleFunc("/repos/integrations/terraform-provider-github/issues", func(w http.ResponseWriter, _ *http.Request) {
		calls++
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]map[string]any{})
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()
	t.Setenv("GITHUB_BASE_URL", srv.URL+"/")
	t.Setenv("GITHUB_TOKEN", "ghp_FAKEFAKEFAKEFAKEFAKEFAKE")

	client, err := NewIssueClientFromEnv("integrations/terraform-provider-github")
	if err != nil {
		t.Fatalf("NewIssueClientFromEnv error: %v", err)
	}
	if _, err := client.ListKnownIssues(context.Background()); err != nil {
		t.Fatalf("ListKnownIssues error: %v", err)
	}
	if calls != 1 {
		t.Fatalf("GHES server calls = %d, want 1", calls)
	}
}
