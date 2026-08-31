package cli

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/github/terraform-provider-tester/engine"
)

type fakeKnownIssueLister struct {
	entries []engine.KnownIssueEntry
	err     error
	calls   int
}

func (f *fakeKnownIssueLister) ListKnownIssues(context.Context) ([]engine.KnownIssueEntry, error) {
	f.calls++
	if f.err != nil {
		return nil, f.err
	}
	return append([]engine.KnownIssueEntry(nil), f.entries...), nil
}

func TestKnownIssueSyncServiceWritesConfiguredPath(t *testing.T) {
	root := cliScratchDir(t)
	cachePath := filepath.Join(root, "config", "known-issues.yaml")
	syncedAt := time.Date(2026, 7, 9, 18, 33, 32, 0, time.FixedZone("EDT", -4*60*60))
	fingerprint := "sha256:" + strings.Repeat("1", 64)
	lister := &fakeKnownIssueLister{entries: []engine.KnownIssueEntry{{
		Fingerprint: fingerprint,
		Issue:       123,
		State:       "open",
		Mode:        "known-real",
		Tests:       []string{"TestAccThing"},
	}}}
	svc := knownIssueSyncService{
		issuesRepo: defaultIssuesRepo,
		cachePath:  cachePath,
		now: func() time.Time {
			return syncedAt
		},
		newLister: func(repo string) (knownIssueLister, error) {
			if repo != defaultIssuesRepo {
				t.Fatalf("repo = %q, want %q", repo, defaultIssuesRepo)
			}
			return lister, nil
		},
	}

	result, err := svc.Sync(context.Background())
	if err != nil {
		t.Fatalf("Sync: %v", err)
	}
	if result.Count != 1 {
		t.Fatalf("Count = %d, want 1", result.Count)
	}
	if result.CachePath != cachePath {
		t.Fatalf("CachePath = %q, want %q", result.CachePath, cachePath)
	}
	wantUTC := syncedAt.UTC()
	if !result.SyncedAt.Equal(wantUTC) {
		t.Fatalf("SyncedAt = %s, want %s", result.SyncedAt.Format(time.RFC3339), wantUTC.Format(time.RFC3339))
	}

	data, err := os.ReadFile(cachePath)
	if err != nil {
		t.Fatalf("read cache: %v", err)
	}
	got := string(data)
	for _, want := range []string{
		"version: 1",
		"source: github",
		"issues_repo: " + defaultIssuesRepo,
		"label: acctest-failure",
		wantUTC.Format(time.RFC3339),
		fingerprint,
		"origin: github",
	} {
		if !strings.Contains(got, want) {
			t.Fatalf("cache missing %q:\n%s", want, got)
		}
	}

	file, err := engine.LoadKnownIssuesFile(cachePath, wantUTC)
	if err != nil {
		t.Fatalf("LoadKnownIssuesFile: %v", err)
	}
	if file.Source != "github" || file.IssuesRepo != defaultIssuesRepo || file.Label != "acctest-failure" {
		t.Fatalf("metadata = %+v", file)
	}
	if file.SyncedAt == nil || !file.SyncedAt.Equal(wantUTC) {
		t.Fatalf("file.SyncedAt = %v, want %s", file.SyncedAt, wantUTC.Format(time.RFC3339))
	}
	if len(file.Entries) != 1 || file.Entries[0].Fingerprint != fingerprint {
		t.Fatalf("entries = %+v, want fingerprint %q", file.Entries, fingerprint)
	}
}

func TestKnownIssueSyncServicePreservesLocalEntriesAndReplacesGitHubSnapshot(t *testing.T) {
	root := cliScratchDir(t)
	cachePath := filepath.Join(root, "known-issues.yaml")
	localFP := "sha256:" + strings.Repeat("a", 64)
	staleGitHubFP := "sha256:" + strings.Repeat("b", 64)
	liveFP := "sha256:" + strings.Repeat("c", 64)
	localOverrideFP := "sha256:" + strings.Repeat("d", 64)
	if err := engine.SaveKnownIssuesFile(cachePath, engine.KnownIssueFile{
		Source: "github",
		Entries: []engine.KnownIssueEntry{
			{Fingerprint: localFP, Mode: "ignore", Origin: "local"},
			{Fingerprint: staleGitHubFP, Issue: 12, State: "open", Mode: "known-real", Origin: "github"},
			{Fingerprint: localOverrideFP, Mode: "ignore", Origin: "local"},
		},
	}); err != nil {
		t.Fatalf("seed cache: %v", err)
	}
	lister := &fakeKnownIssueLister{entries: []engine.KnownIssueEntry{{
		Fingerprint: liveFP,
		Issue:       13,
		State:       "open",
		Mode:        "known-real",
	}, {
		Fingerprint: localOverrideFP,
		Issue:       14,
		State:       "open",
		Mode:        "known-real",
	}}}
	svc := knownIssueSyncService{
		issuesRepo: defaultIssuesRepo,
		cachePath:  cachePath,
		now:        func() time.Time { return time.Date(2026, 7, 9, 22, 33, 32, 0, time.UTC) },
		newLister:  func(string) (knownIssueLister, error) { return lister, nil },
	}

	result, err := svc.Sync(context.Background())
	if err != nil {
		t.Fatalf("Sync: %v", err)
	}
	if result.Count != 2 {
		t.Fatalf("Count = %d, want live GitHub entry count", result.Count)
	}
	file, err := engine.LoadKnownIssuesFile(cachePath, result.SyncedAt)
	if err != nil {
		t.Fatalf("LoadKnownIssuesFile: %v", err)
	}
	got := map[string]string{}
	for _, entry := range file.Entries {
		got[entry.Fingerprint] = entry.Origin
	}
	want := map[string]string{liveFP: "github", localFP: "local", localOverrideFP: "local"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("synced entries = %+v, want %+v", got, want)
	}
}

func TestKnownIssueSyncServicePreservesExistingCacheOnFailure(t *testing.T) {
	if os.Getuid() == 0 {
		t.Skip("permission-based save failure is not meaningful as root")
	}

	root := cliScratchDir(t)
	cacheDir := filepath.Join(root, "config")
	cachePath := filepath.Join(cacheDir, "known-issues.yaml")
	original := []byte("version: 1\nentries:\n  - fingerprint: old-cache\n")
	if err := os.MkdirAll(cacheDir, 0o700); err != nil {
		t.Fatalf("mkdir cache dir: %v", err)
	}
	if err := os.WriteFile(cachePath, original, 0o600); err != nil {
		t.Fatalf("seed cache: %v", err)
	}
	if err := os.Chmod(cacheDir, 0o500); err != nil {
		t.Fatalf("chmod cache dir: %v", err)
	}
	defer func() {
		_ = os.Chmod(cacheDir, 0o700)
	}()

	lister := &fakeKnownIssueLister{entries: []engine.KnownIssueEntry{{
		Fingerprint: "sha256:" + strings.Repeat("2", 64),
		Issue:       456,
		State:       "open",
	}}}
	svc := knownIssueSyncService{
		issuesRepo: defaultIssuesRepo,
		cachePath:  cachePath,
		now: func() time.Time {
			return time.Date(2026, 7, 9, 22, 33, 32, 0, time.UTC)
		},
		newLister: func(string) (knownIssueLister, error) {
			return lister, nil
		},
	}

	if _, err := svc.Sync(context.Background()); err == nil {
		t.Fatal("Sync error = nil, want save failure")
	}

	data, err := os.ReadFile(cachePath)
	if err != nil {
		t.Fatalf("read cache after failure: %v", err)
	}
	if string(data) != string(original) {
		t.Fatalf("cache bytes changed on failed sync:\n%s", string(data))
	}
	matches, err := filepath.Glob(filepath.Join(cacheDir, ".known-issues.yaml.*.tmp"))
	if err != nil {
		t.Fatalf("glob temp files: %v", err)
	}
	if len(matches) != 0 {
		t.Fatalf("leftover temp files: %v", matches)
	}
}

type cancelingKnownIssueLister struct {
	cancel  context.CancelFunc
	entries []engine.KnownIssueEntry
}

func (l cancelingKnownIssueLister) ListKnownIssues(context.Context) ([]engine.KnownIssueEntry, error) {
	l.cancel()
	return append([]engine.KnownIssueEntry(nil), l.entries...), nil
}

func TestKnownIssueSyncServiceHonorsCancellationBeforePersist(t *testing.T) {
	root := cliScratchDir(t)
	cachePath := filepath.Join(root, "config", "known-issues.yaml")
	ctx, cancel := context.WithCancel(context.Background())
	svc := knownIssueSyncService{
		issuesRepo: defaultIssuesRepo,
		cachePath:  cachePath,
		now:        func() time.Time { return time.Date(2026, 7, 9, 22, 33, 32, 0, time.UTC) },
		newLister: func(string) (knownIssueLister, error) {
			return cancelingKnownIssueLister{cancel: cancel, entries: []engine.KnownIssueEntry{{Fingerprint: "sha256:" + strings.Repeat("4", 64), Issue: 444, State: "open"}}}, nil
		},
	}

	if _, err := svc.Sync(ctx); !errors.Is(err, context.Canceled) {
		t.Fatalf("Sync error = %v, want context.Canceled", err)
	}
	if _, err := os.Stat(cachePath); !os.IsNotExist(err) {
		t.Fatalf("cache path stat = %v, want not created", err)
	}
}
