package engine

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/github/terraform-provider-tester/provider"
)

// TestStateRecordStoresNoOutput confirms that no Output field, and no secret
// value from Output, reaches the serialized state.
func TestStateRecordStoresNoOutput(t *testing.T) {
	var s State
	s.Record(TestResult{
		Name:    "TestAccGithubRepo",
		Status:  provider.StatusPass,
		Package: "pkg/github",
		Output:  []string{"PASS", "ghp_SECRETTOKEN"},
	})

	data, err := json.Marshal(s)
	if err != nil {
		t.Fatalf("json.Marshal: %v", err)
	}
	js := string(data)

	if strings.Contains(js, `"output"`) {
		t.Errorf("JSON contains \"output\" field: %s", js)
	}
	if strings.Contains(js, "ghp_SECRETTOKEN") {
		t.Errorf("JSON contains secret token: %s", js)
	}
}

// TestStateSaveLoadRoundTrip saves a State with Results and History, then
// loads it back and asserts deep equality.
func TestStateSaveLoadRoundTrip(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "state.json")

	s := State{
		Provider: "github",
		Mode:     "organization",
		RunAt:    time.Unix(1700000000, 0).UTC(),
		Results: []PersistResult{
			{Test: "TestA", Status: "pass", Elapsed: 1.5, Package: "pkg/github"},
			{Test: "TestB", Sub: "scenario", Status: "fail", Elapsed: 0.3, Package: "pkg/github"},
		},
		History: map[string][]string{
			"TestA": {"pass", "fail"},
			"TestB": {"fail"},
		},
	}

	if err := s.Save(path); err != nil {
		t.Fatalf("Save: %v", err)
	}

	got, err := Load(path)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}

	if !got.RunAt.Equal(s.RunAt) {
		t.Errorf("RunAt: got %v, want %v", got.RunAt, s.RunAt)
	}
	// Normalize RunAt so reflect.DeepEqual is not tripped by internal clock representation.
	got.RunAt = s.RunAt

	if !reflect.DeepEqual(got, s) {
		t.Errorf("loaded state differs:\ngot  %+v\nwant %+v", got, s)
	}
}

func TestStateSaveEmitsVersionTwo(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "state.json")

	s := State{Provider: "github"}
	if err := s.Save(path); err != nil {
		t.Fatalf("Save: %v", err)
	}

	if s.Version != StateVersion {
		t.Fatalf("Save() left Version=%d, want %d", s.Version, StateVersion)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}

	var raw map[string]any
	if err := json.Unmarshal(data, &raw); err != nil {
		t.Fatalf("json.Unmarshal: %v", err)
	}

	got, ok := raw["version"].(float64)
	if !ok {
		t.Fatalf("state JSON missing numeric version field: %s", string(data))
	}
	if int(got) != StateVersion {
		t.Fatalf("state JSON version = %d, want %d", int(got), StateVersion)
	}
}

func TestStateLoadLegacyVersionZero(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "state.json")

	legacy := `{
  "provider": "github",
  "mode": "organization",
  "results": [],
  "history": {
    "TestAccGithubRepo": ["pass", "fail"]
  }
}`
	if err := os.WriteFile(path, []byte(legacy), 0o600); err != nil {
		t.Fatalf("WriteFile: %v", err)
	}

	got, err := Load(path)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}

	if got.Version != 0 {
		t.Fatalf("legacy version = %d, want 0", got.Version)
	}
	if got.Plan != nil {
		t.Fatalf("legacy state synthesized plan: %+v", got.Plan)
	}
	if got.Provider != "github" || got.Mode != "organization" {
		t.Fatalf("legacy fields not preserved: %+v", got)
	}
}

func TestStatePlanRoundTrip(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "state.json")

	s := State{
		Provider: "github",
		Mode:     "organization",
		RunAt:    time.Unix(1700000100, 0).UTC(),
		Plan: &ExecutionPlan{
			Mode:         "organization",
			Group:        "repositories",
			Run:          "^TestAccGithubRepository",
			Selected:     []string{"TestAccGithubRepository", "TestAccGithubUser"},
			Eligible:     []string{"TestAccGithubRepository"},
			Excluded:     []PlanExclusion{{Test: "TestAccGithubUser", Code: "excluded/mode-incompatible", Detail: "test does not support mode", Fix: "run in one of: [user]", RequiredModes: []string{"user"}}},
			Unclassified: []string{"TestAccGithubRepositoryUnknown"},
			Scopes:       []string{"org", "repo"},
			Capabilities: []string{"administration", "repository"},
			SideEffects:  []string{"unknown"},
		},
	}

	if err := s.Save(path); err != nil {
		t.Fatalf("Save: %v", err)
	}

	got, err := Load(path)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}

	if got.Version != StateVersion {
		t.Fatalf("Version = %d, want %d", got.Version, StateVersion)
	}
	if got.Plan == nil {
		t.Fatal("Plan = nil, want persisted execution plan")
	}
	if !reflect.DeepEqual(got.Plan, s.Plan) {
		t.Fatalf("Plan mismatch:\ngot  %+v\nwant %+v", got.Plan, s.Plan)
	}
}

func TestStateValidEmptyPlanDiffersFromLegacy(t *testing.T) {
	dir := t.TempDir()
	emptyPath := filepath.Join(dir, "empty-plan.json")
	legacyPath := filepath.Join(dir, "legacy-plan.json")

	empty := State{Plan: &ExecutionPlan{Eligible: []string{}}}
	legacy := State{Plan: nil}

	if err := empty.Save(emptyPath); err != nil {
		t.Fatalf("Save empty plan: %v", err)
	}
	if err := legacy.Save(legacyPath); err != nil {
		t.Fatalf("Save legacy plan: %v", err)
	}

	gotEmpty, err := Load(emptyPath)
	if err != nil {
		t.Fatalf("Load empty plan: %v", err)
	}
	gotLegacy, err := Load(legacyPath)
	if err != nil {
		t.Fatalf("Load legacy plan: %v", err)
	}

	if gotEmpty.Plan == nil {
		t.Fatal("empty plan round trip lost non-nil plan")
	}
	if !reflect.DeepEqual(gotEmpty.Plan, empty.Plan) {
		t.Fatalf("empty plan mismatch:\ngot  %+v\nwant %+v", gotEmpty.Plan, empty.Plan)
	}
	if gotLegacy.Plan != nil {
		t.Fatalf("legacy plan should remain nil after round trip, got %+v", gotLegacy.Plan)
	}
	if reflect.DeepEqual(gotEmpty, gotLegacy) {
		t.Fatalf("empty and legacy states should differ after round trip:\nempty  %+v\nlegacy %+v", gotEmpty, gotLegacy)
	}
}

func TestStatePlanStoresNoSecretMarkers(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "state.json")

	s := State{
		Plan: &ExecutionPlan{
			Mode:        "organization",
			Selected:    []string{"TestAccGithubRepository"},
			Eligible:    []string{"TestAccGithubRepository"},
			Excluded:    []PlanExclusion{{Test: "TestAccGithubRepository", Code: "excluded/mode-incompatible", Detail: "plan stays explicit", Fix: "pick a compatible mode"}},
			Scopes:      []string{"repo"},
			SideEffects: []string{"write"},
		},
	}
	if err := s.Save(path); err != nil {
		t.Fatalf("Save: %v", err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}

	js := string(data)
	if strings.Contains(js, "***REDACTED***") {
		t.Fatalf("state plan stored redaction marker: %s", js)
	}
	if !strings.Contains(js, `"plan"`) {
		t.Fatalf("state JSON missing plan: %s", js)
	}
}

func TestStateOrphanAccountingRoundTrip(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "state.json")

	s := State{
		Provider: "github",
		Mode:     "organization",
		RunAt:    time.Unix(1700000200, 0).UTC(),
		Orphans: &OrphanAccounting{
			Mode: "organization",
			Baseline: []provider.Resource{
				{Kind: "repository", Name: "tf-acc-test-a", URL: "https://example.test/repos/tf-acc-test-a-baseline"},
			},
			Final: []provider.Resource{
				{Kind: "repository", Name: "tf-acc-test-a", URL: "https://example.test/repos/tf-acc-test-a-final"},
				{Kind: "repository", Name: "tf-acc-test-new", URL: "https://example.test/repos/tf-acc-test-new"},
			},
			PreExisting: []provider.Resource{
				{Kind: "repository", Name: "tf-acc-test-a", URL: "https://example.test/repos/tf-acc-test-a-final"},
			},
			New: []provider.Resource{
				{Kind: "repository", Name: "tf-acc-test-new", URL: "https://example.test/repos/tf-acc-test-new"},
			},
			BaselineCaptured: true,
			FinalCaptured:    true,
			CleanupStatus:    CleanupComplete,
		},
	}

	if err := s.Save(path); err != nil {
		t.Fatalf("Save: %v", err)
	}

	got, err := Load(path)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}

	if !got.RunAt.Equal(s.RunAt) {
		t.Fatalf("RunAt: got %v, want %v", got.RunAt, s.RunAt)
	}
	got.RunAt = s.RunAt

	if !reflect.DeepEqual(got, s) {
		t.Fatalf("loaded state differs:\ngot  %+v\nwant %+v", got, s)
	}
}

func TestStateEmptyCapturedBaselineDiffersFromUncaptured(t *testing.T) {
	dir := t.TempDir()
	capturedPath := filepath.Join(dir, "captured.json")
	uncapturedPath := filepath.Join(dir, "uncaptured.json")

	captured := State{
		Provider: "github",
		Orphans: &OrphanAccounting{
			Mode:             "organization",
			Baseline:         []provider.Resource{},
			BaselineCaptured: true,
			CleanupStatus:    CleanupBaselineOnly,
		},
	}
	uncaptured := State{Provider: "github"}

	if err := captured.Save(capturedPath); err != nil {
		t.Fatalf("Save captured: %v", err)
	}
	if err := uncaptured.Save(uncapturedPath); err != nil {
		t.Fatalf("Save uncaptured: %v", err)
	}

	gotCaptured, err := Load(capturedPath)
	if err != nil {
		t.Fatalf("Load captured: %v", err)
	}
	gotUncaptured, err := Load(uncapturedPath)
	if err != nil {
		t.Fatalf("Load uncaptured: %v", err)
	}

	if gotCaptured.Orphans == nil {
		t.Fatal("captured baseline lost orphan accounting")
	}
	if !gotCaptured.Orphans.BaselineCaptured {
		t.Fatal("captured baseline lost BaselineCaptured marker")
	}
	if gotCaptured.Orphans.Baseline == nil {
		t.Fatal("captured baseline lost explicit empty baseline slice")
	}
	if gotUncaptured.Orphans != nil {
		t.Fatalf("uncaptured state synthesized orphan accounting: %+v", gotUncaptured.Orphans)
	}
	if reflect.DeepEqual(gotCaptured, gotUncaptured) {
		t.Fatalf("captured and uncaptured states should differ:\ncaptured  %+v\nuncaptured %+v", gotCaptured, gotUncaptured)
	}
}

func TestStateWithoutOrphansRemainsReadableAtVersionTwo(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "state.json")

	legacyV2 := `{
  "version": 2,
  "provider": "github",
  "mode": "organization",
  "results": [],
  "history": {}
}`
	if err := os.WriteFile(path, []byte(legacyV2), 0o600); err != nil {
		t.Fatalf("WriteFile: %v", err)
	}

	got, err := Load(path)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}

	if got.Version != StateVersion {
		t.Fatalf("Version = %d, want %d", got.Version, StateVersion)
	}
	if got.Orphans != nil {
		t.Fatalf("Orphans = %+v, want nil", got.Orphans)
	}
	if got.Provider != "github" || got.Mode != "organization" {
		t.Fatalf("state fields not preserved: %+v", got)
	}
}

// TestStateSaveAtomic confirms that no leftover temp file remains after a
// successful Save, and that the target file is readable.
func TestStateSaveAtomic(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "state.json")

	s := State{Provider: "github"}
	if err := s.Save(path); err != nil {
		t.Fatalf("Save: %v", err)
	}

	// No temp file should remain.
	matches, err := filepath.Glob(filepath.Join(dir, ".pulsar-state-*.tmp"))
	if err != nil {
		t.Fatalf("Glob: %v", err)
	}
	if len(matches) != 0 {
		t.Errorf("leftover temp files after Save: %v", matches)
	}

	// Target must exist and parse.
	if _, err := os.Stat(path); err != nil {
		t.Errorf("target file not found: %v", err)
	}
	if _, err := Load(path); err != nil {
		t.Errorf("Load after Save: %v", err)
	}
}

// TestStateLoadMissingFile confirms that loading a nonexistent path returns
// a zero State and nil error.
func TestStateLoadMissingFile(t *testing.T) {
	got, err := Load(filepath.Join(t.TempDir(), "does-not-exist.json"))
	if err != nil {
		t.Fatalf("expected nil error for missing file, got: %v", err)
	}
	if !reflect.DeepEqual(got, State{}) {
		t.Errorf("expected zero State, got %+v", got)
	}
}

// TestFailedTopLevelDedup records a passing parent, a failing subtest under
// the same parent, and a fully passing second parent. FailedTopLevel must
// return only the first parent (once), deduplicating and excluding B.
func TestFailedTopLevelDedup(t *testing.T) {
	var s State
	// Top-level A: pass.
	s.Record(TestResult{Name: "A", Sub: "", Status: provider.StatusPass, Package: "pkg"})
	// Subtest A/x: fail - parent name is still "A".
	s.Record(TestResult{Name: "A", Sub: "x", Status: provider.StatusFail, Package: "pkg"})
	// Top-level B: pass, no failing subtests.
	s.Record(TestResult{Name: "B", Sub: "", Status: provider.StatusPass, Package: "pkg"})

	got := s.FailedTopLevel()
	want := []string{"A"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("FailedTopLevel() = %v; want %v", got, want)
	}
}

// TestNotRunSelection records a result for "A" and confirms that
// NotRun(["A","B","C"]) returns ["B","C"] in input order.
func TestNotRunSelection(t *testing.T) {
	var s State
	s.Record(TestResult{Name: "A", Status: provider.StatusPass, Package: "pkg"})

	got := s.NotRun([]string{"A", "B", "C"})
	want := []string{"B", "C"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("NotRun() = %v; want %v", got, want)
	}
}

// TestFlakyDetection covers the four canonical Flaky cases from the brief.
func TestFlakyDetection(t *testing.T) {
	tests := []struct {
		name    string
		history []string
		failN   int
		lastM   int
		want    bool
	}{
		{
			name:    "two_of_five_is_flaky",
			history: []string{"pass", "fail", "pass", "pass", "fail"},
			failN:   2,
			lastM:   5,
			want:    true,
		},
		{
			name:    "all_fail_is_regression_not_flaky",
			history: []string{"fail", "fail", "fail"},
			failN:   2,
			lastM:   5,
			want:    false,
		},
		{
			name:    "all_pass_is_not_flaky",
			history: []string{"pass", "pass"},
			failN:   1,
			lastM:   5,
			want:    false,
		},
		{
			name:    "one_fail_in_last_five_window_is_flaky",
			history: []string{"pass", "pass", "pass", "pass", "pass", "fail"},
			failN:   1,
			lastM:   5,
			want:    true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Flaky(tc.history, tc.failN, tc.lastM)
			if got != tc.want {
				t.Errorf("Flaky(%v, %d, %d) = %v; want %v",
					tc.history, tc.failN, tc.lastM, got, tc.want)
			}
		})
	}
}

// TestResumeRegexFromState records two failing top-level results and confirms
// that RunPattern(FailedTopLevel()) produces the correct anchored alternation.
func TestResumeRegexFromState(t *testing.T) {
	var s State
	s.Record(TestResult{Name: "TestAccGithubA", Status: provider.StatusFail, Package: "pkg"})
	s.Record(TestResult{Name: "TestAccGithubB", Status: provider.StatusFail, Package: "pkg"})

	got := RunPattern(s.FailedTopLevel())
	want := "^(TestAccGithubA|TestAccGithubB)$"

	if got != want {
		t.Errorf("RunPattern(FailedTopLevel()) = %q; want %q", got, want)
	}
}

// TestAcquireLockRefusesSecond verifies mutual exclusion: a second acquire
// on a held lock returns an error, and after release a third acquire succeeds.
func TestAcquireLockRefusesSecond(t *testing.T) {
	lockPath := filepath.Join(t.TempDir(), "test.lock")

	release1, err := AcquireLock(lockPath)
	if err != nil {
		t.Fatalf("first AcquireLock: %v", err)
	}

	_, err = AcquireLock(lockPath)
	if err == nil {
		t.Fatal("second AcquireLock: expected error, got nil")
	}

	if err := release1(); err != nil {
		t.Fatalf("release1: %v", err)
	}

	release3, err := AcquireLock(lockPath)
	if err != nil {
		t.Fatalf("third AcquireLock after release: %v", err)
	}
	if err := release3(); err != nil {
		t.Fatalf("release3: %v", err)
	}
}

// reapedPID starts a trivial child process, waits for it to exit, and returns
// its now-dead PID. signal-0 liveness checks against this PID return "dead"
// (barring near-immediate PID reuse, which is vanishingly unlikely in a test).
func reapedPID(t *testing.T) int {
	t.Helper()
	cmd := exec.Command("sleep", "0")
	if err := cmd.Start(); err != nil {
		t.Fatalf("start helper process: %v", err)
	}
	pid := cmd.Process.Pid
	if err := cmd.Wait(); err != nil {
		t.Fatalf("wait helper process: %v", err)
	}
	return pid
}

func seedLock(t *testing.T, path string, pid int, host string) {
	t.Helper()
	body := fmt.Sprintf(`{"pid":%d,"host":%q,"acquired_at":%q}`,
		pid, host, time.Now().UTC().Format(time.RFC3339))
	if err := os.WriteFile(path, []byte(body), 0o600); err != nil {
		t.Fatalf("seed lock %s: %v", path, err)
	}
}

// TestAcquireLockReclaimsEmptyLegacyLock reproduces the real failure seen during
// validation: a crashed run leaves a 0-byte lock that has no owner metadata. A
// fresh acquire must reclaim it rather than bricking every later run. The lock
// is backdated past the write grace to stand in for a run that crashed earlier.
func TestAcquireLockReclaimsEmptyLegacyLock(t *testing.T) {
	lockPath := filepath.Join(t.TempDir(), "test.lock")
	if err := os.WriteFile(lockPath, nil, 0o600); err != nil {
		t.Fatalf("seed empty lock: %v", err)
	}
	old := time.Now().Add(-time.Hour)
	if err := os.Chtimes(lockPath, old, old); err != nil {
		t.Fatalf("backdate lock: %v", err)
	}

	release, err := AcquireLock(lockPath)
	if err != nil {
		t.Fatalf("AcquireLock over empty legacy lock: %v", err)
	}
	if release == nil {
		t.Fatal("release is nil")
	}
	if err := release(); err != nil {
		t.Fatalf("release: %v", err)
	}
}

// TestAcquireLockKeepsFreshEmptyLock proves the write-window guard: a just-created
// empty lock (as a concurrent acquirer would see mid-stamp) is left in place, not
// reclaimed, so two runs cannot both believe they hold the lock.
func TestAcquireLockKeepsFreshEmptyLock(t *testing.T) {
	lockPath := filepath.Join(t.TempDir(), "test.lock")
	if err := os.WriteFile(lockPath, nil, 0o600); err != nil {
		t.Fatalf("seed empty lock: %v", err)
	}

	if _, err := AcquireLock(lockPath); err == nil {
		t.Fatal("AcquireLock over a fresh empty lock: expected refusal, got nil")
	}
}

// TestAcquireLockDoesNotRemoveNewLockCreatedDuringStaleReclaim proves stale
// reclaim cannot unlink a replacement lock acquired by another run. This
// models two acquirers racing on the same abandoned file: once one removes the
// stale file and another creates a fresh lock, a later stale-removal attempt
// must not delete that fresh owner.
func TestAcquireLockDoesNotRemoveNewLockCreatedDuringStaleReclaim(t *testing.T) {
	lockPath := filepath.Join(t.TempDir(), "test.lock")
	host, _ := os.Hostname()
	stalePID := reapedPID(t)
	seedLock(t, lockPath, stalePID, host)

	data, err := os.ReadFile(lockPath)
	if err != nil {
		t.Fatalf("read stale lock: %v", err)
	}
	identity, err := lockIdentityForPath(lockPath)
	if err != nil {
		t.Fatalf("lock identity: %v", err)
	}
	if err := os.Remove(lockPath); err != nil {
		t.Fatalf("remove stale lock: %v", err)
	}
	release, err := AcquireLock(lockPath)
	if err != nil {
		t.Fatalf("acquire replacement lock: %v", err)
	}
	defer release() //nolint:errcheck

	reclaimed, err := reclaimStaleLockWithIdentity(lockPath, data, identity)
	if err != nil {
		t.Fatalf("reclaim with old identity: %v", err)
	}
	if reclaimed {
		t.Fatal("expected old stale-reclaim attempt not to remove a replacement lock")
	}
	if _, err := os.Stat(lockPath); err != nil {
		t.Fatalf("replacement lock was removed: %v", err)
	}
}

// TestLockIdentityChangesWhenContentReplacedInPlace pins the cross-platform
// regression behind the test above. A lock identity must change when the file's
// contents change even if its inode does not. Truncating and rewriting a file
// (what os.WriteFile does) keeps the same dev+ino, which models Linux reusing a
// freed inode for a replacement lock at the same path. If the identity ignored
// content, stale reclaim could unlink a live replacement lock.
func TestLockIdentityChangesWhenContentReplacedInPlace(t *testing.T) {
	lockPath := filepath.Join(t.TempDir(), "test.lock")
	if err := os.WriteFile(lockPath, []byte("original-stale-owner"), 0o600); err != nil {
		t.Fatalf("seed lock: %v", err)
	}
	before, err := lockIdentityForPath(lockPath)
	if err != nil {
		t.Fatalf("identity before: %v", err)
	}
	if err := os.WriteFile(lockPath, []byte("replacement-live-owner"), 0o600); err != nil {
		t.Fatalf("rewrite lock in place: %v", err)
	}
	after, err := lockIdentityForPath(lockPath)
	if err != nil {
		t.Fatalf("identity after: %v", err)
	}
	if after == before {
		t.Fatal("lock identity ignored a same-inode content change; stale reclaim could unlink a replacement lock")
	}
}
func TestAcquireLockReclaimsDeadOwner(t *testing.T) {
	lockPath := filepath.Join(t.TempDir(), "test.lock")
	host, _ := os.Hostname()
	seedLock(t, lockPath, reapedPID(t), host)

	release, err := AcquireLock(lockPath)
	if err != nil {
		t.Fatalf("AcquireLock over dead-owner lock: %v", err)
	}
	if err := release(); err != nil {
		t.Fatalf("release: %v", err)
	}
}

// TestAcquireLockRefusesForeignHost verifies that a lock written by a different
// host (whose process liveness cannot be checked) is left in place, not blindly
// reclaimed.
func TestAcquireLockRefusesForeignHost(t *testing.T) {
	lockPath := filepath.Join(t.TempDir(), "test.lock")
	seedLock(t, lockPath, os.Getpid(), "some-other-host-9f3a")

	if _, err := AcquireLock(lockPath); err == nil {
		t.Fatal("AcquireLock over a foreign-host lock: expected refusal, got nil")
	}
}

// TestRemoveLockReportsAndDeletes covers `terraform-provider-tester unlock`: it removes a stuck lock
// and is a no-op (no error) when none is present.
func TestRemoveLockReportsAndDeletes(t *testing.T) {
	lockPath := filepath.Join(t.TempDir(), "test.lock")
	if _, err := AcquireLock(lockPath); err != nil {
		t.Fatalf("AcquireLock: %v", err)
	}

	msg, err := RemoveLock(lockPath)
	if err != nil {
		t.Fatalf("RemoveLock: %v", err)
	}
	if !strings.Contains(msg, "removed lock") {
		t.Errorf("RemoveLock message = %q, want it to mention removal", msg)
	}
	if _, statErr := os.Stat(lockPath); !os.IsNotExist(statErr) {
		t.Errorf("lock file still present after RemoveLock: stat err = %v", statErr)
	}

	msg2, err := RemoveLock(lockPath)
	if err != nil {
		t.Fatalf("RemoveLock on absent lock: %v", err)
	}
	if !strings.Contains(msg2, "no lock file present") {
		t.Errorf("RemoveLock on absent lock = %q, want 'no lock file present'", msg2)
	}
}
