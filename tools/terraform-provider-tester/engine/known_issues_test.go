package engine

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestLoadKnownIssuesFileParsesSchemaAndMarksStale(t *testing.T) {
	now := time.Date(2026, 7, 8, 20, 0, 0, 0, time.UTC)
	path := writeKnownIssuesFixture(t, `version: 1
source: github
issues_repo: integrations/terraform-provider-github
label: acctest-failure
synced_at: "2026-06-30T18:00:00Z"
entries:
  - fingerprint: "sha256:`+strings.Repeat("a", 64)+`"
    short: "`+strings.Repeat("a", 16)+`"
    issue: 1234
    issue_url: "https://github.com/integrations/terraform-provider-github/issues/1234"
    title: "[acctest] TestAccGithubRepository api/422-leftover-state"
    state: open
    mode: known-real
    tests:
      - TestAccGithubRepository
    modes:
      - organization
    classes:
      - api/422-leftover-state
    retry: retryable
    expires_at: "2026-10-01T00:00:00Z"
    note: "Run sweep before retrying."
    updated_at: "2026-07-08T18:00:00Z"
`)

	file, err := LoadKnownIssuesFile(path, now)
	if err != nil {
		t.Fatalf("LoadKnownIssuesFile error: %v", err)
	}
	if file.Version != 1 || file.IssuesRepo != "integrations/terraform-provider-github" || file.Label != "acctest-failure" {
		t.Fatalf("bad file header: %+v", file)
	}
	if len(file.Entries) != 1 {
		t.Fatalf("entries = %d, want 1", len(file.Entries))
	}
	entry := file.Entries[0]
	if !entry.Stale {
		t.Fatalf("entry should be stale when synced_at is older than 7 days: %+v", entry)
	}
	if entry.Expired {
		t.Fatalf("entry should not be expired: %+v", entry)
	}
	if entry.Issue != 1234 || entry.Mode != "known-real" || entry.Retry != "retryable" {
		t.Fatalf("entry fields not parsed: %+v", entry)
	}
	if got := strings.Join(entry.Tests, ","); got != "TestAccGithubRepository" {
		t.Fatalf("tests = %q", got)
	}
}

func TestKnownIssueRegistryMergesLiveOverCacheAndHonorsOffline(t *testing.T) {
	fp := "sha256:" + strings.Repeat("b", 64)
	path := writeKnownIssuesFixture(t, `version: 1
source: github
synced_at: "2026-07-08T18:00:00Z"
entries:
  - fingerprint: "`+fp+`"
    origin: github
    issue: 100
    state: open
    mode: known-real
    modes:
      - organization
`)
	live := fakeKnownIssueSource{entries: []KnownIssueEntry{{
		Fingerprint: fp,
		Issue:       200,
		State:       "open",
		Mode:        "known-real",
		Modes:       []string{"organization"},
	}}}

	reg := &KnownIssueRegistry{CachePath: path, Live: &live, Now: func() time.Time {
		return time.Date(2026, 7, 8, 20, 0, 0, 0, time.UTC)
	}}
	got, err := reg.LookupFingerprint(context.Background(), fp, "organization")
	if err != nil {
		t.Fatalf("LookupFingerprint error: %v", err)
	}
	if got.Issue != 200 || !got.Suppress {
		t.Fatalf("live issue should win over cache, got %+v", got)
	}

	reg.Offline = true
	got, err = reg.LookupFingerprint(context.Background(), fp, "organization")
	if err != nil {
		t.Fatalf("offline LookupFingerprint error: %v", err)
	}
	if got.Issue != 100 || !got.Suppress {
		t.Fatalf("offline lookup should use cache, got %+v", got)
	}
	if live.calls != 1 {
		t.Fatalf("live source calls = %d, want only the online lookup", live.calls)
	}
}

func TestKnownIssueRegistryMissingCacheStillUsesLiveLookup(t *testing.T) {
	fp := "sha256:" + strings.Repeat("d", 64)
	live := fakeKnownIssueSource{entries: []KnownIssueEntry{{
		Fingerprint: fp,
		Issue:       900,
		State:       "open",
		Mode:        "known-real",
		Modes:       []string{"organization"},
	}}}

	reg := &KnownIssueRegistry{CachePath: missingKnownIssuesPath(t), Live: &live}
	got, err := reg.LookupFingerprint(context.Background(), fp, "organization")
	if err != nil {
		t.Fatalf("LookupFingerprint error: %v", err)
	}
	if got.Issue != 900 || !got.Suppress {
		t.Fatalf("missing cache should not disable live lookup, got %+v", got)
	}
	if live.calls != 1 {
		t.Fatalf("live source calls = %d, want 1", live.calls)
	}
}

func TestKnownIssueRegistryMissingCacheOfflineIsEmpty(t *testing.T) {
	fp := "sha256:" + strings.Repeat("e", 64)
	reg := &KnownIssueRegistry{CachePath: missingKnownIssuesPath(t), Offline: true}

	got, err := reg.LookupFingerprint(context.Background(), fp, "organization")
	if err != nil {
		t.Fatalf("LookupFingerprint error: %v", err)
	}
	if got.Found || got.Suppress || got.Issue != 0 {
		t.Fatalf("missing offline cache match = %+v, want empty match", got)
	}
}

func TestKnownIssueRegistryOnlineDropsMissingGitHubSnapshotButKeepsLocalEntries(t *testing.T) {
	remoteFP := "sha256:" + strings.Repeat("1", 64)
	localFP := "sha256:" + strings.Repeat("2", 64)
	legacyLocalFP := "sha256:" + strings.Repeat("3", 64)
	legacyRemoteFP := "sha256:" + strings.Repeat("4", 64)
	localOverrideFP := "sha256:" + strings.Repeat("5", 64)
	path := writeKnownIssuesFixture(t, `version: 1
source: github
synced_at: "2026-06-30T18:00:00Z"
entries:
  - fingerprint: "`+remoteFP+`"
    origin: github
    issue: 100
    issue_url: "https://github.com/integrations/terraform-provider-github/issues/100"
    state: open
    mode: known-real
  - fingerprint: "`+localFP+`"
    origin: local
    issue: 200
    issue_url: "https://github.com/integrations/terraform-provider-github/issues/200"
    state: open
    mode: known-real
  - fingerprint: "`+legacyLocalFP+`"
    mode: ignore
    note: "hand-authored local rule"
  - fingerprint: "`+legacyRemoteFP+`"
    issue: 400
    issue_url: "https://github.com/integrations/terraform-provider-github/issues/400"
    state: open
    mode: known-real
  - fingerprint: "`+localOverrideFP+`"
    origin: local
    mode: ignore
    note: "local override wins"
`)
	live := fakeKnownIssueSource{entries: []KnownIssueEntry{{
		Fingerprint: localOverrideFP,
		Issue:       500,
		State:       "open",
		Mode:        "known-real",
	}}}
	reg := &KnownIssueRegistry{
		CachePath: path,
		Live:      &live,
		Now: func() time.Time {
			return time.Date(2026, 7, 8, 20, 0, 0, 0, time.UTC)
		},
	}

	for name, fp := range map[string]string{
		"explicit GitHub origin": remoteFP,
		"legacy GitHub origin":   legacyRemoteFP,
	} {
		remote, err := reg.LookupFingerprint(context.Background(), fp, "organization")
		if err != nil {
			t.Fatalf("%s LookupFingerprint error: %v", name, err)
		}
		if remote.Found || remote.Suppress || remote.Issue != 0 {
			t.Fatalf("%s entry missing from live GitHub should not suppress, got %+v", name, remote)
		}
	}
	for name, fp := range map[string]string{
		"explicit local": localFP,
		"legacy local":   legacyLocalFP,
	} {
		match, err := reg.LookupFingerprint(context.Background(), fp, "organization")
		if err != nil {
			t.Fatalf("%s LookupFingerprint error: %v", name, err)
		}
		if !match.Found || !match.Suppress {
			t.Fatalf("%s entry should remain authoritative online, got %+v", name, match)
		}
	}
	override, err := reg.LookupFingerprint(context.Background(), localOverrideFP, "organization")
	if err != nil {
		t.Fatalf("local override LookupFingerprint error: %v", err)
	}
	if !override.Found || !override.Suppress || override.Issue != 0 || override.Mode != "ignore" || override.Note != "local override wins" {
		t.Fatalf("local override should win over same-fingerprint live issue, got %+v", override)
	}
	if live.calls != 1 {
		t.Fatalf("live source calls = %d, want one merged lookup", live.calls)
	}

	reg.Offline = true
	remote, err := reg.LookupFingerprint(context.Background(), remoteFP, "organization")
	if err != nil {
		t.Fatalf("offline remote LookupFingerprint error: %v", err)
	}
	if !remote.Found || !remote.Suppress || remote.Issue != 100 {
		t.Fatalf("explicit offline lookup should retain cached snapshot, got %+v", remote)
	}
}

func TestKnownIssueRegistryDoesNotSuppressExpiredEntries(t *testing.T) {
	fp := "sha256:" + strings.Repeat("c", 64)
	path := writeKnownIssuesFixture(t, `version: 1
synced_at: "2026-07-08T18:00:00Z"
entries:
  - fingerprint: "`+fp+`"
    issue: 300
    state: open
    mode: known-real
    modes:
      - organization
    expires_at: "2026-07-01T00:00:00Z"
`)

	reg := &KnownIssueRegistry{CachePath: path, Offline: true, Now: func() time.Time {
		return time.Date(2026, 7, 8, 20, 0, 0, 0, time.UTC)
	}}
	got, err := reg.LookupFingerprint(context.Background(), fp, "organization")
	if err != nil {
		t.Fatalf("LookupFingerprint error: %v", err)
	}
	if got.Found || got.Suppress || got.Issue != 0 {
		t.Fatalf("expired entry should not suppress filing, got %+v", got)
	}
}

func TestSaveKnownIssuesFileUsesPrivateAtomicWrite(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "known.yaml")

	file := KnownIssueFile{
		Source:     "github",
		IssuesRepo: "integrations/terraform-provider-github",
		Label:      "acctest-failure",
		Entries: []KnownIssueEntry{{
			Fingerprint: "sha256:" + strings.Repeat("g", 64),
			Issue:       42,
			State:       "open",
		}},
	}
	if err := SaveKnownIssuesFile(path, file); err != nil {
		t.Fatalf("SaveKnownIssuesFile: %v", err)
	}

	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("stat saved file: %v", err)
	}
	if got := info.Mode().Perm(); got != 0o600 {
		t.Fatalf("mode = %v, want 0600", got)
	}

	matches, err := filepath.Glob(filepath.Join(dir, ".*.tmp"))
	if err != nil {
		t.Fatalf("glob temp files: %v", err)
	}
	if len(matches) != 0 {
		t.Fatalf("leftover temp files: %v", matches)
	}
}

func TestKnownIssueRegistryMatchesLocalIgnoreWithoutIssue(t *testing.T) {
	fp := "sha256:" + strings.Repeat("f", 64)
	path := writeKnownIssuesFixture(t, `version: 1
synced_at: "2026-07-08T18:00:00Z"
entries:
  - fingerprint: "`+fp+`"
    mode: ignore
    modes:
      - organization
    note: "local suppress"
`)

	reg := &KnownIssueRegistry{CachePath: path, Offline: true}
	match, err := reg.LookupFingerprint(context.Background(), fp, "organization")
	if err != nil {
		t.Fatalf("LookupFingerprint error: %v", err)
	}
	if !match.Found || !match.Suppress || match.Issue != 0 || match.Mode != "ignore" || match.Note != "local suppress" {
		t.Fatalf("local ignore match = %+v, want suppressing match without issue", match)
	}
}

func TestKnownIssueRegistryExpiredLocalIgnoreDoesNotSuppress(t *testing.T) {
	fp := "sha256:" + strings.Repeat("0", 64)
	path := writeKnownIssuesFixture(t, `version: 1
synced_at: "2026-07-08T18:00:00Z"
entries:
  - fingerprint: "`+fp+`"
    mode: ignore
    modes:
      - organization
    expires_at: "2026-07-01T00:00:00Z"
`)

	reg := &KnownIssueRegistry{CachePath: path, Offline: true, Now: func() time.Time {
		return time.Date(2026, 7, 8, 20, 0, 0, 0, time.UTC)
	}}
	match, err := reg.LookupFingerprint(context.Background(), fp, "organization")
	if err != nil {
		t.Fatalf("LookupFingerprint error: %v", err)
	}
	if match.Found || match.Suppress || match.Issue != 0 {
		t.Fatalf("expired local ignore match = %+v, want no suppressing match", match)
	}
}

func writeKnownIssuesFixture(t *testing.T, body string) string {
	t.Helper()
	dir := filepath.Join(".", ".test-scratch", strings.NewReplacer("/", "_", " ", "_").Replace(t.Name()))
	if err := os.RemoveAll(dir); err != nil {
		t.Fatalf("clean scratch: %v", err)
	}
	if err := os.MkdirAll(dir, 0o700); err != nil {
		t.Fatalf("make scratch: %v", err)
	}
	t.Cleanup(func() { _ = os.RemoveAll(dir) })
	path := filepath.Join(dir, "known.yaml")
	if err := os.WriteFile(path, []byte(body), 0o600); err != nil {
		t.Fatalf("write fixture: %v", err)
	}
	return path
}

func missingKnownIssuesPath(t *testing.T) string {
	t.Helper()
	dir := filepath.Join(".", ".test-scratch", strings.NewReplacer("/", "_", " ", "_").Replace(t.Name()))
	if err := os.RemoveAll(dir); err != nil {
		t.Fatalf("clean scratch: %v", err)
	}
	if err := os.MkdirAll(dir, 0o700); err != nil {
		t.Fatalf("make scratch: %v", err)
	}
	t.Cleanup(func() { _ = os.RemoveAll(dir) })
	return filepath.Join(dir, "missing.yaml")
}

type fakeKnownIssueSource struct {
	entries []KnownIssueEntry
	calls   int
}

func (f *fakeKnownIssueSource) ListKnownIssues(context.Context) ([]KnownIssueEntry, error) {
	f.calls++
	return f.entries, nil
}
