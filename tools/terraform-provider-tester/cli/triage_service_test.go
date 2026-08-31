package cli

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/github/terraform-provider-tester/engine"
	"github.com/github/terraform-provider-tester/internal/redact"
)

func TestTriageStateServiceRefreshUsesCacheWithoutLiveLookup(t *testing.T) {
	root := cliScratchDir(t)
	statePath := filepath.Join(root, ".pulsar-state.json")
	failure := triageFailureFixture()
	st := engine.State{
		Provider: "github",
		Mode:     "organization",
		Failures: []engine.PersistFailure{failure},
	}
	if err := st.Save(statePath); err != nil {
		t.Fatalf("save state: %v", err)
	}

	cachePath := filepath.Join(root, "known-issues.yaml")
	now := time.Date(2026, 7, 9, 12, 0, 0, 0, time.UTC)
	if err := engine.SaveKnownIssuesFile(cachePath, engine.KnownIssueFile{
		Version:    1,
		IssuesRepo: defaultIssuesRepo,
		SyncedAt:   &now,
		Entries: []engine.KnownIssueEntry{{
			Fingerprint: failure.Fingerprint,
			Short:       failure.ShortFingerprint,
			Issue:       42,
			Mode:        "known-real",
			Title:       "known issue from cache",
			State:       "open",
		}},
	}); err != nil {
		t.Fatalf("save known issues cache: %v", err)
	}

	svc := triageStateService{
		root:      root,
		statePath: statePath,
		red:       redact.New(nil),
		registryFactory: func(triageOptions) (issueRegistry, error) {
			t.Fatal("offline refresh must not create a live known-issues registry")
			return nil, nil
		},
		filerFactory: func(string) (issueFiler, error) {
			t.Fatal("offline refresh must not create an issue filer")
			return nil, nil
		},
		getenv: func(key string) string {
			if key == "GITHUB_TOKEN" {
				return "ghp_FAKEFAKEFAKEFAKEFAKEFAKE"
			}
			return ""
		},
	}

	result, err := svc.Refresh(context.Background(), triageOptions{
		IssuesRepo:         defaultIssuesRepo,
		KnownIssuesPath:    cachePath,
		KnownIssuesOffline: true,
	}, false)
	if err != nil {
		t.Fatalf("Refresh: %v", err)
	}
	if !result.CacheAvailable {
		t.Fatal("CacheAvailable = false, want true")
	}
	if len(result.Failures) != 1 {
		t.Fatalf("failures = %+v, want one", result.Failures)
	}
	if result.Failures[0].IssueAction != "known" || result.Failures[0].KnownIssue != 42 {
		t.Fatalf("offline cache match = %+v, want known issue #42", result.Failures[0])
	}

	loaded, err := engine.Load(statePath)
	if err != nil {
		t.Fatalf("load refreshed state: %v", err)
	}
	if len(loaded.Failures) != 1 || loaded.Failures[0].IssueAction != "known" || loaded.Failures[0].KnownIssue != 42 {
		t.Fatalf("saved failures = %+v, want known cached issue", loaded.Failures)
	}
}

func TestTriageStateServiceForceReconstructClearsPassingFailure(t *testing.T) {
	root := cliScratchDir(t)
	statePath := filepath.Join(root, ".pulsar-state.json")
	st := engine.State{
		Provider: "github",
		Mode:     "organization",
		Results: []engine.PersistResult{{
			Package: "./github",
			Test:    "TestAccThing",
			Status:  "pass",
		}},
		Failures: []engine.PersistFailure{triageFailureFixture()},
	}
	if err := st.Save(statePath); err != nil {
		t.Fatalf("save state: %v", err)
	}

	svc := triageStateService{
		root:      root,
		statePath: statePath,
		red:       redact.New(nil),
	}

	result, err := svc.Refresh(context.Background(), triageOptions{IssuesRepo: defaultIssuesRepo}, true)
	if err != nil {
		t.Fatalf("Refresh: %v", err)
	}
	if len(result.Failures) != 0 {
		t.Fatalf("failures = %+v, want cleared stale failure", result.Failures)
	}

	loaded, err := engine.Load(statePath)
	if err != nil {
		t.Fatalf("load refreshed state: %v", err)
	}
	if len(loaded.Failures) != 0 {
		t.Fatalf("saved failures = %+v, want none after force reconstruct", loaded.Failures)
	}
}

func TestTriageStateServiceLocksBeforeLoad(t *testing.T) {
	root := cliScratchDir(t)
	statePath := filepath.Join(root, ".pulsar-state.json")
	if err := os.WriteFile(statePath, []byte("{not-json"), 0o600); err != nil {
		t.Fatalf("write broken state: %v", err)
	}

	release, err := engine.AcquireLock(statePath + ".lock")
	if err != nil {
		t.Fatalf("AcquireLock: %v", err)
	}
	defer func() {
		if err := release(); err != nil {
			t.Fatalf("release lock: %v", err)
		}
	}()

	svc := triageStateService{root: root, statePath: statePath, red: redact.New(nil)}
	_, err = svc.Refresh(context.Background(), triageOptions{IssuesRepo: defaultIssuesRepo}, false)
	if err == nil {
		t.Fatal("Refresh: expected lock error, got nil")
	}
	if !strings.Contains(err.Error(), "lock already held") {
		t.Fatalf("Refresh error = %q, want lock-held error", err)
	}
	if strings.Contains(err.Error(), "invalid character") {
		t.Fatalf("Refresh loaded state before taking lock: %v", err)
	}
}

func TestTriageStateServiceRefreshReturnsStateMode(t *testing.T) {
	root := cliScratchDir(t)
	statePath := filepath.Join(root, ".pulsar-state.json")
	st := engine.State{
		Provider: "github",
		Mode:     "organization",
		Failures: []engine.PersistFailure{{
			Package:          "./github",
			Test:             "TestAccThing",
			Status:           "fail",
			Fingerprint:      "sha256:" + strings.Repeat("a", 64),
			ShortFingerprint: strings.Repeat("a", 16),
			Class:            engine.ClassLeftoverState,
			Canonical:        "422 leftover state ***REDACTED***",
			Retryable:        true,
			Classification:   engine.ClassificationReal,
			Attempts:         3,
		}},
	}
	if err := st.Save(statePath); err != nil {
		t.Fatalf("save state: %v", err)
	}

	svc := triageStateService{root: root, statePath: statePath, red: redact.New(nil)}
	result, err := svc.Refresh(context.Background(), triageOptions{IssuesRepo: defaultIssuesRepo}, false)
	if err != nil {
		t.Fatalf("Refresh: %v", err)
	}
	if result.Mode != "organization" {
		t.Fatalf("result.Mode = %q, want organization", result.Mode)
	}
}

func TestPersistSuiteFailureLogsRedactsAndClears(t *testing.T) {
	root := cliScratchDir(t)
	failDir := filepath.Join(root, ".pulsar-failures")
	if err := os.MkdirAll(failDir, 0o700); err != nil {
		t.Fatalf("make failure dir: %v", err)
	}

	st := &engine.State{}
	secret := "ghp_FAKEFAKEFAKEFAKEFAKEFAKE"
	if err := persistSuiteFailureLogs(root, st, engine.RunResult{
		BuildFailed:  true,
		BuildOutput:  []string{"build " + secret + "\n"},
		PreRunFailed: true,
		PreRunOutput: []string{"pre-run " + secret + "\n"},
	}, redact.New(nil)); err != nil {
		t.Fatalf("persistSuiteFailureLogs failure case: %v", err)
	}
	if !st.BuildFailed || !st.PreRunFailed {
		t.Fatalf("suite failure flags = build:%v pre-run:%v, want both true", st.BuildFailed, st.PreRunFailed)
	}
	if st.BuildLog != filepath.Join(failDir, "build.log") || st.PreRunLog != filepath.Join(failDir, "pre-run.log") {
		t.Fatalf("suite failure paths = %q / %q", st.BuildLog, st.PreRunLog)
	}
	for _, path := range []string{st.BuildLog, st.PreRunLog} {
		data, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("read %s: %v", path, err)
		}
		if strings.Contains(string(data), secret) {
			t.Fatalf("%s leaked secret: %s", path, string(data))
		}
		if !strings.Contains(string(data), "***REDACTED***") {
			t.Fatalf("%s missing redaction marker: %s", path, string(data))
		}
	}

	if err := persistSuiteFailureLogs(root, st, engine.RunResult{}, redact.New(nil)); err != nil {
		t.Fatalf("persistSuiteFailureLogs green case: %v", err)
	}
	if st.BuildFailed || st.PreRunFailed {
		t.Fatalf("suite failure flags after green = build:%v pre-run:%v, want both false", st.BuildFailed, st.PreRunFailed)
	}
	if st.BuildLog != "" || st.PreRunLog != "" {
		t.Fatalf("suite failure paths after green = %q / %q, want cleared", st.BuildLog, st.PreRunLog)
	}
	for _, stale := range []string{"build.log", "pre-run.log"} {
		if _, err := os.Stat(filepath.Join(failDir, stale)); !os.IsNotExist(err) {
			t.Fatalf("stale %s still exists or stat failed: %v", stale, err)
		}
	}
}

func triageFailureFixture() engine.PersistFailure {
	return engine.PersistFailure{
		Package:          "./github",
		Test:             "TestAccThing",
		Status:           "fail",
		Fingerprint:      "sha256:" + strings.Repeat("a", 64),
		ShortFingerprint: strings.Repeat("a", 16),
		Class:            engine.ClassLeftoverState,
		Canonical:        "422 leftover state ***REDACTED***",
		Retryable:        true,
		Classification:   engine.ClassificationReal,
		Attempts:         3,
		Mode:             "organization",
	}
}
