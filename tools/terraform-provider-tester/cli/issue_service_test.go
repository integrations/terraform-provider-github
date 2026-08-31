package cli

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/github/terraform-provider-tester/engine"
	"github.com/github/terraform-provider-tester/internal/redact"
	ghissues "github.com/github/terraform-provider-tester/provider/github"
)

func TestIssueServicePreviewOmitsBodyAndRedactsTitle(t *testing.T) {
	const secret = "ghp_FAKEPREVIEWTOKEN"
	root := cliScratchDir(t)
	statePath := filepath.Join(root, ".pulsar-state.json")
	failure := issueServiceFailure("a", func(f *engine.PersistFailure) {
		f.Test = "TestAcc" + secret
		f.LogPath = issueServiceWriteLog(t, root, f, "raw log "+secret+"\n")
	})
	st := engine.State{
		Provider: "github",
		Mode:     "organization",
		Failures: []engine.PersistFailure{failure},
	}
	if err := st.Save(statePath); err != nil {
		t.Fatalf("save state: %v", err)
	}
	rec := &issueServiceRecorder{}
	reg := &issueServiceFakeRegistry{rec: rec}
	filer := &issueServiceFakeFiler{rec: rec}
	svc := issueServiceFixture(root, statePath, reg, filer, redact.New([]string{secret}))

	got, err := svc.Preview(failure.Fingerprint)
	if err != nil {
		t.Fatalf("Preview: %v", err)
	}
	if got.Fingerprint != failure.Fingerprint || got.ShortFingerprint != failure.ShortFingerprint {
		t.Fatalf("fingerprints = %q/%q, want authoritative %q/%q", got.Fingerprint, got.ShortFingerprint, failure.Fingerprint, failure.ShortFingerprint)
	}
	if got.IssuesRepo != defaultIssuesRepo || got.Classification != engine.ClassificationReal {
		t.Fatalf("metadata = %+v, want repo %q classification real", got, defaultIssuesRepo)
	}
	if strings.Contains(got.Title, secret) || !strings.Contains(got.Title, "***REDACTED***") {
		t.Fatalf("title = %q, want redacted secret", got.Title)
	}
	if len(got.Labels) == 0 || got.Labels[0] != "acctest-failure" {
		t.Fatalf("labels = %+v, want copied issue labels", got.Labels)
	}
	typ := reflect.TypeOf(got)
	for _, name := range []string{"Body", "LogExcerpt", "LogLines"} {
		if _, ok := typ.FieldByName(name); ok {
			t.Fatalf("Preview result exposes %s; TUI must receive metadata only", name)
		}
	}
	if len(rec.calls) != 0 || filer.createCalls != 0 {
		t.Fatalf("preview calls = %+v createCalls=%d, want no network/create calls", rec.calls, filer.createCalls)
	}
}

func TestIssueServicePreviewReturnsReleaseError(t *testing.T) {
	root, statePath := issueServiceSeedState(t, engine.State{
		Provider: "github",
		Mode:     "organization",
		Failures: []engine.PersistFailure{issueServiceFailure("e")},
	})
	failure := issueServiceLoadFailures(t, statePath)[0]
	svc := issueServiceFixture(root, statePath, &issueServiceFakeRegistry{}, &issueServiceFakeFiler{}, redact.New(nil))
	svc.acquireLock = func(string) (func() error, error) {
		return func() error { return errors.New("release failed") }, nil
	}

	got, err := svc.Preview(failure.Fingerprint)
	if err == nil || !strings.Contains(err.Error(), "releasing lock: release failed") {
		t.Fatalf("Preview error = %v, want release failure", err)
	}
	if got.Fingerprint != failure.Fingerprint {
		t.Fatalf("Preview result fingerprint = %q, want %q", got.Fingerprint, failure.Fingerprint)
	}
}

func TestIssueServiceFileOneRejectsWrongPhrase(t *testing.T) {
	root, statePath := issueServiceSeedState(t, engine.State{
		Provider: "github",
		Mode:     "organization",
		Failures: []engine.PersistFailure{issueServiceFailure("b")},
	})
	before := issueServiceLoadFailures(t, statePath)
	rec := &issueServiceRecorder{}
	reg := &issueServiceFakeRegistry{rec: rec}
	filer := &issueServiceFakeFiler{rec: rec, created: &ghissues.IssueResult{Number: 99}}
	svc := issueServiceFixture(root, statePath, reg, filer, redact.New(nil))

	got, err := svc.FileOne(context.Background(), before[0].Fingerprint, "FILE "+strings.ToUpper(before[0].ShortFingerprint))
	if err == nil {
		t.Fatal("FileOne wrong phrase: expected error, got nil")
	}
	if !reflect.DeepEqual(got, issueFileResult{}) {
		t.Fatalf("result = %+v, want zero result on rejection", got)
	}
	if len(rec.calls) != 0 || filer.createCalls != 0 || filer.dedupCalls != 0 {
		t.Fatalf("calls = %+v dedup=%d create=%d, want zero side effects", rec.calls, filer.dedupCalls, filer.createCalls)
	}
	if after := issueServiceLoadFailures(t, statePath); !reflect.DeepEqual(after, before) {
		t.Fatalf("state changed on wrong phrase:\nafter=%+v\nbefore=%+v", after, before)
	}
}

func TestIssueServiceFileOneRejectsFlakeAndKnownFailure(t *testing.T) {
	cases := []struct {
		name   string
		mutate func(*engine.PersistFailure)
	}{
		{name: "flake", mutate: func(f *engine.PersistFailure) { f.Classification = engine.ClassificationFlakeConfirmed }},
		{name: "known", mutate: func(f *engine.PersistFailure) { f.KnownIssue = 42 }},
		{name: "actioned", mutate: func(f *engine.PersistFailure) { f.IssueAction = "dedup" }},
		{name: "malformed", mutate: func(f *engine.PersistFailure) { f.Fingerprint = "sha256:" + strings.Repeat("c", 63) }},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			failure := issueServiceFailure("c")
			tc.mutate(&failure)
			root, statePath := issueServiceSeedState(t, engine.State{Provider: "github", Mode: "organization", Failures: []engine.PersistFailure{failure}})
			before := issueServiceLoadFailures(t, statePath)
			rec := &issueServiceRecorder{}
			reg := &issueServiceFakeRegistry{rec: rec}
			filer := &issueServiceFakeFiler{rec: rec, created: &ghissues.IssueResult{Number: 99}}
			svc := issueServiceFixture(root, statePath, reg, filer, redact.New(nil))

			got, err := svc.FileOne(context.Background(), failure.Fingerprint, "FILE "+failure.ShortFingerprint)
			if err == nil {
				t.Fatal("FileOne ineligible failure: expected error, got nil")
			}
			if !reflect.DeepEqual(got, issueFileResult{}) {
				t.Fatalf("result = %+v, want zero result on rejection", got)
			}
			if len(rec.calls) != 0 || filer.createCalls != 0 || filer.dedupCalls != 0 {
				t.Fatalf("calls = %+v dedup=%d create=%d, want zero side effects", rec.calls, filer.dedupCalls, filer.createCalls)
			}
			if after := issueServiceLoadFailures(t, statePath); !reflect.DeepEqual(after, before) {
				t.Fatalf("state changed on ineligible failure:\nafter=%+v\nbefore=%+v", after, before)
			}
		})
	}
}

func TestIssueServiceFileOneUsesLiveRegistryBeforeDedup(t *testing.T) {
	root, statePath := issueServiceSeedState(t, engine.State{
		Provider: "github",
		Mode:     "organization",
		Failures: []engine.PersistFailure{issueServiceFailure("d", func(f *engine.PersistFailure) { f.Mode = "individual" })},
	})
	failure := issueServiceLoadFailures(t, statePath)[0]
	rec := &issueServiceRecorder{}
	reg := &issueServiceFakeRegistry{rec: rec, matches: map[string]engine.KnownIssueMatch{
		failure.Fingerprint: {Found: true, Issue: 44, Mode: "known-real", Suppress: true, Note: "registry suppressed"},
	}}
	filer := &issueServiceFakeFiler{rec: rec, created: &ghissues.IssueResult{Number: 99}}
	svc := issueServiceFixture(root, statePath, reg, filer, redact.New(nil))

	got, err := svc.FileOne(context.Background(), failure.Fingerprint, "FILE "+failure.ShortFingerprint)
	if err != nil {
		t.Fatalf("FileOne: %v", err)
	}
	wantCalls := []string{"registry:" + failure.Fingerprint + ":organization"}
	if !reflect.DeepEqual(rec.calls, wantCalls) {
		t.Fatalf("calls = %+v, want %+v", rec.calls, wantCalls)
	}
	if filer.dedupCalls != 0 || filer.createCalls != 0 {
		t.Fatalf("dedup/create calls = %d/%d, want 0/0 for registry suppression", filer.dedupCalls, filer.createCalls)
	}
	if got.Action != "known" || got.IssueNumber != 44 {
		t.Fatalf("result = %+v, want known issue #44", got)
	}
	after := issueServiceLoadFailures(t, statePath)
	if after[0].IssueAction != "known" || after[0].KnownIssue != 44 || !containsString(after[0].Reasons, "registry suppressed") {
		t.Fatalf("persisted failure = %+v, want known issue with note", after[0])
	}
}

func TestIssueServiceFileOneDedupsBeforeCreate(t *testing.T) {
	root, statePath := issueServiceSeedState(t, engine.State{
		Provider: "github",
		Mode:     "organization",
		Failures: []engine.PersistFailure{issueServiceFailure("e")},
	})
	failure := issueServiceLoadFailures(t, statePath)[0]
	rec := &issueServiceRecorder{}
	reg := &issueServiceFakeRegistry{rec: rec}
	filer := &issueServiceFakeFiler{rec: rec, dedup: map[string]*ghissues.IssueMatch{
		failure.Fingerprint: {Number: 77, State: "open", URL: "https://github.com/integrations/terraform-provider-github/issues/77"},
	}, created: &ghissues.IssueResult{Number: 99}}
	svc := issueServiceFixture(root, statePath, reg, filer, redact.New(nil))

	got, err := svc.FileOne(context.Background(), failure.Fingerprint, "FILE "+failure.ShortFingerprint)
	if err != nil {
		t.Fatalf("FileOne: %v", err)
	}
	wantCalls := []string{"registry:" + failure.Fingerprint + ":organization", "find:" + failure.Fingerprint}
	if !reflect.DeepEqual(rec.calls, wantCalls) {
		t.Fatalf("calls = %+v, want %+v", rec.calls, wantCalls)
	}
	if filer.createCalls != 0 {
		t.Fatalf("create calls = %d, want 0 after dedup", filer.createCalls)
	}
	if got.Action != "dedup" || got.IssueNumber != 77 {
		t.Fatalf("result = %+v, want dedup issue #77", got)
	}
	after := issueServiceLoadFailures(t, statePath)
	if after[0].IssueAction != "dedup" || after[0].KnownIssue != 77 {
		t.Fatalf("persisted failure = %+v, want dedup issue #77", after[0])
	}
}

func TestIssueServiceFileOneCreatesOnlySelectedFailure(t *testing.T) {
	selected := issueServiceFailure("f")
	other := issueServiceFailure("a", func(f *engine.PersistFailure) {
		f.Test = "TestAccOther"
		f.Canonical = "other canonical must be preserved"
		f.Reasons = []string{"keep me"}
	})
	root, statePath := issueServiceSeedState(t, engine.State{
		Provider: "github",
		Mode:     "organization",
		Failures: []engine.PersistFailure{other, selected},
	})
	before := issueServiceLoadFailures(t, statePath)
	rec := &issueServiceRecorder{}
	reg := &issueServiceFakeRegistry{rec: rec}
	filer := &issueServiceFakeFiler{rec: rec, created: &ghissues.IssueResult{Number: 101, URL: "https://github.com/integrations/terraform-provider-github/issues/101"}}
	svc := issueServiceFixture(root, statePath, reg, filer, redact.New(nil))

	got, err := svc.FileOne(context.Background(), selected.Fingerprint, "FILE "+selected.ShortFingerprint)
	if err != nil {
		t.Fatalf("FileOne: %v", err)
	}
	wantCalls := []string{"registry:" + selected.Fingerprint + ":organization", "find:" + selected.Fingerprint, "create:" + selected.Fingerprint}
	if !reflect.DeepEqual(rec.calls, wantCalls) {
		t.Fatalf("calls = %+v, want %+v", rec.calls, wantCalls)
	}
	if filer.createCalls != 1 || len(filer.drafts) != 1 || len(filer.createdFingerprints) != 1 {
		t.Fatalf("createCalls/drafts/fingerprints = %d/%d/%d, want 1/1/1", filer.createCalls, len(filer.drafts), len(filer.createdFingerprints))
	}
	if !strings.Contains(filer.drafts[0].Body, selected.Fingerprint) || strings.Contains(filer.drafts[0].Body, other.Fingerprint) {
		t.Fatalf("created draft body should contain only selected fingerprint:\n%s", filer.drafts[0].Body)
	}
	if got.Action != "filed" || got.IssueNumber != 101 || got.Fingerprint != selected.Fingerprint {
		t.Fatalf("result = %+v, want filed selected issue #101", got)
	}
	after := issueServiceLoadFailures(t, statePath)
	if !reflect.DeepEqual(after[0], before[0]) {
		t.Fatalf("nonselected failure changed:\nafter=%+v\nbefore=%+v", after[0], before[0])
	}
	if after[1].IssueAction != "filed" || after[1].KnownIssue != 101 {
		t.Fatalf("selected failure = %+v, want filed issue #101", after[1])
	}
}

func TestIssueServiceFileOneRevalidatesPersistedKnownIssue(t *testing.T) {
	failure := issueServiceFailure("9", func(f *engine.PersistFailure) {
		f.KnownIssue = 42
		f.IssueAction = "known"
	})
	root, statePath := issueServiceSeedState(t, engine.State{
		Provider: "github",
		Mode:     "organization",
		Failures: []engine.PersistFailure{failure},
	})
	rec := &issueServiceRecorder{}
	reg := &issueServiceFakeRegistry{rec: rec}
	filer := &issueServiceFakeFiler{
		rec:     rec,
		created: &ghissues.IssueResult{Number: 101, URL: "https://github.com/integrations/terraform-provider-github/issues/101"},
	}
	svc := issueServiceFixture(root, statePath, reg, filer, redact.New(nil))

	got, err := svc.FileOne(context.Background(), failure.Fingerprint, "FILE "+failure.ShortFingerprint)
	if err != nil {
		t.Fatalf("FileOne: %v", err)
	}
	wantCalls := []string{
		"registry:" + failure.Fingerprint + ":organization",
		"find:" + failure.Fingerprint,
		"create:" + failure.Fingerprint,
	}
	if !reflect.DeepEqual(rec.calls, wantCalls) {
		t.Fatalf("calls = %+v, want %+v", rec.calls, wantCalls)
	}
	if got.Action != "filed" || got.IssueNumber != 101 {
		t.Fatalf("result = %+v, want refreshed filing #101", got)
	}
	after := issueServiceLoadFailures(t, statePath)
	if after[0].IssueAction != "filed" || after[0].KnownIssue != 101 {
		t.Fatalf("persisted failure = %+v, want refreshed filed issue #101", after[0])
	}
}

func TestIssueServiceFileOneOmitsLogExcerptOutsideFailureDir(t *testing.T) {
	const outsideSecret = "outside-secret-token-10"
	root := cliScratchDir(t)
	outsidePath := filepath.Join(root, "unrelated-secret.log")
	if err := os.WriteFile(outsidePath, []byte("unrelated "+outsideSecret+"\n"), 0o600); err != nil {
		t.Fatalf("write outside log: %v", err)
	}
	failure := issueServiceFailure("a", func(f *engine.PersistFailure) {
		f.LogPath = outsidePath
	})
	statePath := filepath.Join(root, ".pulsar-state.json")
	st := engine.State{Provider: "github", Mode: "organization", Failures: []engine.PersistFailure{failure}}
	if err := st.Save(statePath); err != nil {
		t.Fatalf("save state: %v", err)
	}
	rec := &issueServiceRecorder{}
	reg := &issueServiceFakeRegistry{rec: rec}
	filer := &issueServiceFakeFiler{rec: rec, created: &ghissues.IssueResult{Number: 404, URL: "https://github.com/integrations/terraform-provider-github/issues/404"}}
	svc := issueServiceFixture(root, statePath, reg, filer, redact.New(nil))

	got, err := svc.FileOne(context.Background(), failure.Fingerprint, "FILE "+failure.ShortFingerprint)
	if err != nil {
		t.Fatalf("FileOne: %v", err)
	}
	if got.Action != "filed" || got.IssueNumber != 404 {
		t.Fatalf("result = %+v, want filing to continue with log excerpt omitted", got)
	}
	if filer.createCalls != 1 || len(filer.drafts) != 1 {
		t.Fatalf("create calls/drafts = %d/%d, want 1/1", filer.createCalls, len(filer.drafts))
	}
	draft := filer.drafts[0]
	if strings.Contains(draft.Title, outsideSecret) || strings.Contains(draft.Body, outsideSecret) || strings.Contains(strings.Join(draft.Labels, ","), outsideSecret) {
		t.Fatalf("outside log secret leaked into issue draft:\ntitle=%q\nlabels=%v\nbody=%s", draft.Title, draft.Labels, draft.Body)
	}
}

func TestIssueServiceFileOneOmitsSymlinkEscapedFailureLog(t *testing.T) {
	const outsideSecret = "symlink-outside-secret-token-10"
	root := cliScratchDir(t)
	failDir := filepath.Join(root, ".pulsar-failures")
	if err := os.MkdirAll(failDir, 0o700); err != nil {
		t.Fatalf("make failure dir: %v", err)
	}
	outsidePath := filepath.Join(root, "unrelated-symlink-secret.log")
	if err := os.WriteFile(outsidePath, []byte("unrelated "+outsideSecret+"\n"), 0o600); err != nil {
		t.Fatalf("write outside log: %v", err)
	}
	absOutsidePath, err := filepath.Abs(outsidePath)
	if err != nil {
		t.Fatalf("abs outside log: %v", err)
	}
	linkPath := filepath.Join(failDir, "escaped.log")
	if err := os.Symlink(absOutsidePath, linkPath); err != nil {
		t.Skipf("platform refused symlink creation: %v", err)
	}
	failure := issueServiceFailure("b", func(f *engine.PersistFailure) {
		f.LogPath = linkPath
	})
	draft := issueServiceFileDraft(t, root, failure, redact.New(nil), 405)

	if strings.Contains(draft.Title, outsideSecret) || strings.Contains(draft.Body, outsideSecret) || strings.Contains(strings.Join(draft.Labels, ","), outsideSecret) {
		t.Fatalf("symlinked outside log secret leaked into issue draft:\ntitle=%q\nlabels=%v\nbody=%s", draft.Title, draft.Labels, draft.Body)
	}
}

func TestIssueServiceFileOneOmitsLogWhenFailureDirSymlinkEscapes(t *testing.T) {
	const outsideSecret = "base-symlink-outside-secret-token-10"
	root := cliScratchDir(t)
	outsideDir := filepath.Join(root, "outside-failures")
	if err := os.MkdirAll(outsideDir, 0o700); err != nil {
		t.Fatalf("make outside failure dir: %v", err)
	}
	leakedLog := filepath.Join(outsideDir, "escaped.log")
	if err := os.WriteFile(leakedLog, []byte("unrelated "+outsideSecret+"\n"), 0o600); err != nil {
		t.Fatalf("write outside log: %v", err)
	}
	absOutsideDir, err := filepath.Abs(outsideDir)
	if err != nil {
		t.Fatalf("abs outside dir: %v", err)
	}
	failDir := filepath.Join(root, ".pulsar-failures")
	if err := os.Symlink(absOutsideDir, failDir); err != nil {
		t.Skipf("platform refused symlink creation: %v", err)
	}
	failure := issueServiceFailure("c", func(f *engine.PersistFailure) {
		f.LogPath = filepath.Join(failDir, "escaped.log")
	})
	draft := issueServiceFileDraft(t, root, failure, redact.New(nil), 406)

	if strings.Contains(draft.Title, outsideSecret) || strings.Contains(draft.Body, outsideSecret) || strings.Contains(strings.Join(draft.Labels, ","), outsideSecret) {
		t.Fatalf("base symlink outside log secret leaked into issue draft:\ntitle=%q\nlabels=%v\nbody=%s", draft.Title, draft.Labels, draft.Body)
	}
}

func TestIssueServiceFileOneIncludesRedactedInDirectoryLogExcerpt(t *testing.T) {
	const secret = "valid-log-secret-token-10"
	for _, tc := range []struct {
		name     string
		logPath  func(t *testing.T, root, path string) string
		issueNum int
	}{
		{
			name: "absolute",
			logPath: func(t *testing.T, _, path string) string {
				return issueServiceAbsPath(t, path)
			},
			issueNum: 407,
		},
		{
			name: "root-relative",
			logPath: func(t *testing.T, root, path string) string {
				return issueServiceRelPath(t, root, path)
			},
			issueNum: 408,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := cliScratchDir(t)
			failure := issueServiceFailure("d")
			logPath := issueServiceWriteLog(t, root, &failure, "raw "+secret+"\n")
			failure.LogPath = tc.logPath(t, root, logPath)

			draft := issueServiceFileDraft(t, root, failure, redact.New([]string{secret}), tc.issueNum)

			if !strings.Contains(draft.Body, "***REDACTED***") {
				t.Fatalf("draft body missing redacted in-directory log excerpt:\n%s", draft.Body)
			}
			if strings.Contains(draft.Title, secret) || strings.Contains(draft.Body, secret) || strings.Contains(strings.Join(draft.Labels, ","), secret) {
				t.Fatalf("raw in-directory log secret leaked into issue draft:\ntitle=%q\nlabels=%v\nbody=%s", draft.Title, draft.Labels, draft.Body)
			}
		})
	}
}

func TestIssueServiceFileOneReportsSaveFailureAfterCreate(t *testing.T) {
	root, statePath := issueServiceSeedState(t, engine.State{
		Provider: "github",
		Mode:     "organization",
		Failures: []engine.PersistFailure{issueServiceFailure("a")},
	})
	before := issueServiceLoadFailures(t, statePath)
	failure := before[0]
	rec := &issueServiceRecorder{}
	reg := &issueServiceFakeRegistry{rec: rec}
	filer := &issueServiceFakeFiler{rec: rec, created: &ghissues.IssueResult{Number: 202, URL: "https://github.com/integrations/terraform-provider-github/issues/202"}}
	saveErr := errors.New("save failed after create")
	svc := issueServiceFixture(root, statePath, reg, filer, redact.New(nil))
	svc.saveState = func(_ *engine.State, path string) error {
		if _, err := os.Stat(path + ".lock"); err != nil {
			t.Fatalf("state lock not held while saving: %v", err)
		}
		return saveErr
	}

	got, err := svc.FileOne(context.Background(), failure.Fingerprint, "FILE "+failure.ShortFingerprint)
	if !errors.Is(err, saveErr) {
		t.Fatalf("FileOne error = %v, want save error", err)
	}
	if !reflect.DeepEqual(got, issueFileResult{}) {
		t.Fatalf("result = %+v, want zero result when save fails", got)
	}
	if filer.createCalls != 1 {
		t.Fatalf("create calls = %d, want 1", filer.createCalls)
	}
	if after := issueServiceLoadFailures(t, statePath); !reflect.DeepEqual(after, before) {
		t.Fatalf("state persisted despite save failure:\nafter=%+v\nbefore=%+v", after, before)
	}
}

func TestIssueServiceFileOneRetryRecoversCreatedIssueByDedup(t *testing.T) {
	root, statePath := issueServiceSeedState(t, engine.State{
		Provider: "github",
		Mode:     "organization",
		Failures: []engine.PersistFailure{issueServiceFailure("b")},
	})
	failure := issueServiceLoadFailures(t, statePath)[0]
	rec := &issueServiceRecorder{}
	reg := &issueServiceFakeRegistry{rec: rec}
	filer := &issueServiceFakeFiler{rec: rec, created: &ghissues.IssueResult{Number: 303, URL: "https://github.com/integrations/terraform-provider-github/issues/303"}}
	saveErr := errors.New("first save failed")
	svc := issueServiceFixture(root, statePath, reg, filer, redact.New(nil))
	failSave := true
	svc.saveState = func(st *engine.State, path string) error {
		if failSave {
			failSave = false
			return saveErr
		}
		return st.Save(path)
	}
	_, err := svc.FileOne(context.Background(), failure.Fingerprint, "FILE "+failure.ShortFingerprint)
	if !errors.Is(err, saveErr) {
		t.Fatalf("first FileOne error = %v, want save error", err)
	}
	filer.dedup = map[string]*ghissues.IssueMatch{
		failure.Fingerprint: {Number: 303, State: "open", URL: "https://github.com/integrations/terraform-provider-github/issues/303"},
	}

	got, err := svc.FileOne(context.Background(), failure.Fingerprint, "FILE "+failure.ShortFingerprint)
	if err != nil {
		t.Fatalf("retry FileOne: %v", err)
	}
	if filer.createCalls != 1 {
		t.Fatalf("create calls = %d, want exactly 1 across save-failure retry", filer.createCalls)
	}
	wantCalls := []string{
		"registry:" + failure.Fingerprint + ":organization",
		"find:" + failure.Fingerprint,
		"create:" + failure.Fingerprint,
		"registry:" + failure.Fingerprint + ":organization",
		"find:" + failure.Fingerprint,
	}
	if !reflect.DeepEqual(rec.calls, wantCalls) {
		t.Fatalf("calls = %+v, want %+v", rec.calls, wantCalls)
	}
	if got.Action != "dedup" || got.IssueNumber != 303 {
		t.Fatalf("retry result = %+v, want dedup issue #303", got)
	}
	after := issueServiceLoadFailures(t, statePath)
	if after[0].IssueAction != "dedup" || after[0].KnownIssue != 303 {
		t.Fatalf("persisted failure = %+v, want recovered dedup issue #303", after[0])
	}
}

type issueServiceRecorder struct {
	calls []string
}

type issueServiceFakeRegistry struct {
	rec     *issueServiceRecorder
	matches map[string]engine.KnownIssueMatch
}

func (f *issueServiceFakeRegistry) LookupFingerprint(_ context.Context, fingerprint, mode string) (engine.KnownIssueMatch, error) {
	f.rec.calls = append(f.rec.calls, "registry:"+fingerprint+":"+mode)
	if f.matches != nil {
		return f.matches[fingerprint], nil
	}
	return engine.KnownIssueMatch{}, nil
}

type issueServiceFakeFiler struct {
	rec                 *issueServiceRecorder
	dedup               map[string]*ghissues.IssueMatch
	created             *ghissues.IssueResult
	dedupCalls          int
	createCalls         int
	findFingerprints    []string
	createdFingerprints []string
	drafts              []ghissues.IssueDraft
}

func (f *issueServiceFakeFiler) FindIssueByFingerprint(_ context.Context, fingerprint string) (*ghissues.IssueMatch, error) {
	f.dedupCalls++
	f.findFingerprints = append(f.findFingerprints, fingerprint)
	f.rec.calls = append(f.rec.calls, "find:"+fingerprint)
	if f.dedup != nil {
		return f.dedup[fingerprint], nil
	}
	return nil, nil
}

func (f *issueServiceFakeFiler) CreateFailureIssue(_ context.Context, draft ghissues.IssueDraft) (*ghissues.IssueResult, error) {
	f.createCalls++
	f.drafts = append(f.drafts, draft)
	fp := issueServiceDraftFingerprint(draft)
	f.createdFingerprints = append(f.createdFingerprints, fp)
	f.rec.calls = append(f.rec.calls, "create:"+fp)
	return f.created, nil
}

func issueServiceFixture(root, statePath string, reg issueRegistry, filer issueFiler, red *redact.Redactor) issueService {
	return issueService{
		root:       root,
		statePath:  statePath,
		issuesRepo: defaultIssuesRepo,
		red:        red,
		registryFactory: func(opts triageOptions) (issueRegistry, error) {
			if opts.IssuesRepo != defaultIssuesRepo {
				return nil, fmt.Errorf("issues repo = %q, want %q", opts.IssuesRepo, defaultIssuesRepo)
			}
			if opts.KnownIssuesOffline {
				return nil, errors.New("selected issue filing must use live registry lookup")
			}
			return reg, nil
		},
		filerFactory: func(repo string) (issueFiler, error) {
			if repo != defaultIssuesRepo {
				return nil, fmt.Errorf("filer repo = %q, want %q", repo, defaultIssuesRepo)
			}
			return filer, nil
		},
		getenv: func(string) string { return "" },
	}
}

func issueServiceSeedState(t *testing.T, st engine.State) (root, statePath string) {
	t.Helper()
	root = cliScratchDir(t)
	statePath = filepath.Join(root, ".pulsar-state.json")
	for i := range st.Failures {
		if st.Failures[i].LogPath == "" {
			st.Failures[i].LogPath = issueServiceWriteLog(t, root, &st.Failures[i], "failure log for "+st.Failures[i].Fingerprint+"\n")
		}
	}
	if err := st.Save(statePath); err != nil {
		t.Fatalf("save state: %v", err)
	}
	return root, statePath
}

func issueServiceFailure(hex string, mutate ...func(*engine.PersistFailure)) engine.PersistFailure {
	fp := "sha256:" + strings.Repeat(hex, 64)
	f := engine.PersistFailure{
		Package:          "./github",
		Test:             "TestAccIssue" + strings.ToUpper(hex),
		Status:           "fail",
		Fingerprint:      fp,
		ShortFingerprint: strings.Repeat(hex, 16),
		Class:            engine.ClassLeftoverState,
		Canonical:        "canonical " + hex,
		Retryable:        true,
		Classification:   engine.ClassificationReal,
		Attempts:         3,
		Mode:             "organization",
	}
	for _, fn := range mutate {
		fn(&f)
	}
	return f
}

func issueServiceWriteLog(t *testing.T, root string, f *engine.PersistFailure, body string) string {
	t.Helper()
	failDir := filepath.Join(root, ".pulsar-failures")
	if err := os.MkdirAll(failDir, 0o700); err != nil {
		t.Fatalf("make failure dir: %v", err)
	}
	path := engine.FailureLogPath(failDir, f.Package, f.Test)
	if err := os.WriteFile(path, []byte(body), 0o600); err != nil {
		t.Fatalf("write failure log: %v", err)
	}
	return path
}

func issueServiceFileDraft(t *testing.T, root string, failure engine.PersistFailure, red *redact.Redactor, issueNumber int) ghissues.IssueDraft {
	t.Helper()
	statePath := filepath.Join(root, ".pulsar-state.json")
	st := engine.State{Provider: "github", Mode: "organization", Failures: []engine.PersistFailure{failure}}
	if err := st.Save(statePath); err != nil {
		t.Fatalf("save state: %v", err)
	}
	rec := &issueServiceRecorder{}
	reg := &issueServiceFakeRegistry{rec: rec}
	filer := &issueServiceFakeFiler{rec: rec, created: &ghissues.IssueResult{Number: issueNumber, URL: fmt.Sprintf("https://github.com/integrations/terraform-provider-github/issues/%d", issueNumber)}}
	svc := issueServiceFixture(root, statePath, reg, filer, red)

	got, err := svc.FileOne(context.Background(), failure.Fingerprint, "FILE "+failure.ShortFingerprint)
	if err != nil {
		t.Fatalf("FileOne: %v", err)
	}
	if got.Action != "filed" || got.IssueNumber != issueNumber {
		t.Fatalf("result = %+v, want filed issue #%d", got, issueNumber)
	}
	if filer.createCalls != 1 || len(filer.drafts) != 1 {
		t.Fatalf("create calls/drafts = %d/%d, want 1/1", filer.createCalls, len(filer.drafts))
	}
	return filer.drafts[0]
}

func issueServiceAbsPath(t *testing.T, path string) string {
	t.Helper()
	abs, err := filepath.Abs(path)
	if err != nil {
		t.Fatalf("absolute path for %q: %v", path, err)
	}
	return abs
}

func issueServiceRelPath(t *testing.T, base, path string) string {
	t.Helper()
	rel, err := filepath.Rel(base, path)
	if err != nil {
		t.Fatalf("relative path from %q to %q: %v", base, path, err)
	}
	return rel
}

func issueServiceLoadFailures(t *testing.T, statePath string) []engine.PersistFailure {
	t.Helper()
	st, err := engine.Load(statePath)
	if err != nil {
		t.Fatalf("load state: %v", err)
	}
	return append([]engine.PersistFailure(nil), st.Failures...)
}

func issueServiceDraftFingerprint(draft ghissues.IssueDraft) string {
	re := regexp.MustCompile(`fingerprint:\s*(sha256:[a-f0-9]{64})`)
	if match := re.FindStringSubmatch(draft.Body); len(match) == 2 {
		return match[1]
	}
	return ""
}

func containsString(values []string, want string) bool {
	for _, value := range values {
		if value == want {
			return true
		}
	}
	return false
}

func TestIssueServiceFileOneHonorsCacheAddedAfterPreview(t *testing.T) {
	root, statePath := issueServiceSeedState(t, engine.State{
		Provider: "github",
		Mode:     "organization",
		Failures: []engine.PersistFailure{issueServiceFailure("1")},
	})
	failure := issueServiceLoadFailures(t, statePath)[0]
	cachePath := filepath.Join(root, "config", "known-issues.yaml")
	rec := &issueServiceRecorder{}
	filer := &issueServiceFakeFiler{rec: rec, created: &ghissues.IssueResult{Number: 909}}
	svc := issueService{
		root:            root,
		statePath:       statePath,
		issuesRepo:      defaultIssuesRepo,
		red:             redact.New(nil),
		knownIssuesPath: cachePath,
		registryFactory: func(opts triageOptions) (issueRegistry, error) {
			if opts.KnownIssuesPath != cachePath {
				t.Fatalf("KnownIssuesPath = %q, want %q", opts.KnownIssuesPath, cachePath)
			}
			return &engine.KnownIssueRegistry{CachePath: opts.KnownIssuesPath}, nil
		},
		filerFactory: func(string) (issueFiler, error) { return filer, nil },
	}

	if _, err := svc.Preview(failure.Fingerprint); err != nil {
		t.Fatalf("Preview before cache update: %v", err)
	}
	if err := engine.SaveKnownIssuesFile(cachePath, engine.KnownIssueFile{Entries: []engine.KnownIssueEntry{{
		Fingerprint: failure.Fingerprint,
		Issue:       808,
		State:       "open",
		Mode:        "known-real",
	}}}); err != nil {
		t.Fatalf("save cache: %v", err)
	}

	got, err := svc.FileOne(context.Background(), failure.Fingerprint, "FILE "+failure.ShortFingerprint)
	if err != nil {
		t.Fatalf("FileOne after cache update: %v", err)
	}
	if got.Action != "known" || got.IssueNumber != 808 {
		t.Fatalf("result = %+v, want known issue #808", got)
	}
	if filer.dedupCalls != 0 || filer.createCalls != 0 {
		t.Fatalf("filer calls = dedup:%d create:%d, want 0/0", filer.dedupCalls, filer.createCalls)
	}
}

func TestIssueServicePreviewRefusesConcurrentStateMutation(t *testing.T) {
	root, statePath := issueServiceSeedState(t, engine.State{Provider: "github", Mode: "organization", Failures: []engine.PersistFailure{issueServiceFailure("2")}})
	failure := issueServiceLoadFailures(t, statePath)[0]
	release, err := engine.AcquireLock(statePath + ".lock")
	if err != nil {
		t.Fatalf("acquire lock: %v", err)
	}
	defer func() { _ = release() }()

	svc := issueServiceFixture(root, statePath, &issueServiceFakeRegistry{rec: &issueServiceRecorder{}}, &issueServiceFakeFiler{rec: &issueServiceRecorder{}}, redact.New(nil))
	if _, err := svc.Preview(failure.Fingerprint); err == nil || !strings.Contains(err.Error(), "lock") {
		t.Fatalf("Preview error = %v, want lock error", err)
	}
}
