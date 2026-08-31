package cli

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/github/terraform-provider-tester/engine"
	"github.com/github/terraform-provider-tester/fakeprovider"
	"github.com/github/terraform-provider-tester/provider"
	"github.com/github/terraform-provider-tester/tui"
)

// fakeRunner records the RunSpec it receives and returns a pre-configured RunResult.
type fakeRunner struct {
	calls  int
	spec   engine.RunSpec
	result engine.RunResult
	err    error
}

func (f *fakeRunner) Run(_ context.Context, spec engine.RunSpec, sink func(engine.TestResult)) (engine.RunResult, error) {
	f.calls++
	f.spec = spec
	for _, tr := range f.result.Tests {
		if sink != nil {
			sink(tr)
		}
	}
	return f.result, f.err
}

// stubList returns a lister that always returns the given names.
func stubList(names []string) lister {
	return func(_ context.Context, _ string, _ []string, _ string) ([]string, error) {
		return names, nil
	}
}

// failList returns a lister that always fails, simulating a build/discovery
// error in the provider package (e.g. `go test -list` aborting on a compile
// failure, which would otherwise yield zero tests).
func failList(err error) lister {
	return func(_ context.Context, _ string, _ []string, _ string) ([]string, error) {
		return nil, err
	}
}

// TestCLIGroupsAgainstFixture verifies that groups lists buckets from the stub name set.
func TestCLIGroupsAgainstFixture(t *testing.T) {
	root := t.TempDir()
	pf := &fakeprovider.Fake{
		NameVal:  "test",
		Packages: []string{"./..."},
		Pattern:  "^TestAcc",
		GroupFunc: func(name string) string {
			switch {
			case strings.HasPrefix(name, "TestAccRepo"):
				return "repos"
			case strings.HasPrefix(name, "TestAccTeam"):
				return "teams"
			default:
				return "misc"
			}
		},
	}
	d := deps{
		provider: pf,
		list:     stubList([]string{"TestAccRepoCreate", "TestAccTeamAdd", "TestAccUnknown"}),
		cwd:      func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"groups", "--repo-root", root}, &out, &errOut, d)

	if code != 0 {
		t.Errorf("expected exit 0, got %d; stderr=%s", code, errOut.String())
	}
	output := out.String()
	for _, want := range []string{"repos", "teams", "misc"} {
		if !strings.Contains(output, want) {
			t.Errorf("expected %q in output; got: %s", want, output)
		}
	}
}

// TestCLIGroupsUnmatchedExit verifies that groups --unmatched exits 1 when misc tests exist.
func TestCLIGroupsUnmatchedExit(t *testing.T) {
	root := t.TempDir()
	pf := &fakeprovider.Fake{
		Packages:  []string{"./..."},
		Pattern:   "^TestAcc",
		GroupFunc: func(_ string) string { return "misc" },
	}
	d := deps{
		provider: pf,
		list:     stubList([]string{"TestAccFoo"}),
		cwd:      func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"groups", "--unmatched", "--repo-root", root}, &out, &errOut, d)

	if code != 1 {
		t.Errorf("expected exit 1, got %d", code)
	}
}

// TestCLIRunWritesState verifies that run writes .pulsar-state.json with test results.
func TestCLIRunWritesState(t *testing.T) {
	root := t.TempDir()
	fr := &fakeRunner{
		result: engine.RunResult{
			Tests: []engine.TestResult{
				{Name: "TestAccA", Status: provider.StatusPass},
				{Name: "TestAccB", Status: provider.StatusFail},
			},
		},
	}
	pf := &fakeprovider.Fake{
		NameVal:        "test",
		Packages:       []string{"./..."},
		Pattern:        "^TestAcc",
		ModesVal:       testModes,
		RequirementsFn: allowAllRequirements,
		PreflightFn: func(_ context.Context, _ string, _ provider.TestRequirements) provider.PreflightReport {
			return provider.PreflightReport{Mode: "anonymous"}
		},
	}
	d := deps{
		provider:  pf,
		newRunner: func(_ io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccA", "TestAccB"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	_ = runWithDeps([]string{"run", "--repo-root", root}, &out, &errOut, d)

	statePath := filepath.Join(root, ".pulsar-state.json")
	loaded, err := engine.Load(statePath)
	if err != nil {
		t.Fatalf("loading state: %v", err)
	}
	if len(loaded.Results) != 2 {
		t.Errorf("expected 2 results in state, got %d", len(loaded.Results))
	}
}

func TestStandaloneRunClearsPriorGuidedOrphanAccounting(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")
	previous := engine.State{
		Provider: "test",
		Mode:     "organization",
		Plan:     planWithEligible("organization", []string{"TestAccPrior"}),
		Orphans: &engine.OrphanAccounting{
			Mode:             "organization",
			Baseline:         []provider.Resource{{Kind: "repository", Name: "tf-acc-test-existing"}},
			Final:            []provider.Resource{{Kind: "repository", Name: "tf-acc-test-new"}},
			New:              []provider.Resource{{Kind: "repository", Name: "tf-acc-test-new"}},
			BaselineCaptured: true,
			FinalCaptured:    true,
			CleanupStatus:    engine.CleanupComplete,
		},
	}
	if err := previous.Save(statePath); err != nil {
		t.Fatalf("seeding prior guided accounting: %v", err)
	}

	runner := &fakeRunner{result: engine.RunResult{Tests: []engine.TestResult{{
		Package: "./github", Name: "TestAccCurrent", Status: provider.StatusPass,
	}}}}
	prov := &fakeprovider.Fake{
		NameVal:        "test",
		Packages:       []string{"./github/..."},
		Pattern:        "^TestAcc",
		ModesVal:       testModes,
		RequirementsFn: allowAllRequirements,
		PreflightFn: func(_ context.Context, mode string, _ provider.TestRequirements) provider.PreflightReport {
			return provider.PreflightReport{Mode: mode}
		},
	}
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccCurrent"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps(
		[]string{"run", "--mode", "organization", "--no-tui", "--repo-root", root},
		&out, &errOut, d,
	)

	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s stdout=%s", code, errOut.String(), out.String())
	}
	persisted, err := engine.Load(statePath)
	if err != nil {
		t.Fatal(err)
	}
	if persisted.Orphans != nil {
		t.Fatalf("standalone run retained prior guided orphan accounting: %+v", persisted.Orphans)
	}
}

// TestRunWritesFailureLogs verifies that run writes redacted failed-test output
// to per-test failure logs and prints the artifact path.
func TestRunWritesFailureLogs(t *testing.T) {
	root := t.TempDir()
	fr := &fakeRunner{
		result: engine.RunResult{
			Tests: []engine.TestResult{
				{Name: "TestAccB", Status: provider.StatusFail, Output: []string{"boom ghp_SECRET\n"}},
			},
		},
	}
	pf := &fakeprovider.Fake{
		NameVal:        "test",
		Packages:       []string{"./..."},
		Pattern:        "^TestAcc",
		ModesVal:       testModes,
		RequirementsFn: allowAllRequirements,
		Secrets:        []string{"GITHUB_TOKEN"},
		PreflightFn: func(_ context.Context, _ string, _ provider.TestRequirements) provider.PreflightReport {
			return provider.PreflightReport{Mode: "anonymous"}
		},
	}
	d := deps{
		provider:  pf,
		newRunner: func(_ io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccB"}),
		cwd:       func() (string, error) { return root, nil },
		getenv:    getenvFromMap(map[string]string{"GITHUB_TOKEN": "ghp_SECRET"}),
	}

	var out, errOut bytes.Buffer
	_ = runWithDeps([]string{"run", "--repo-root", root}, &out, &errOut, d)

	failDir := filepath.Join(root, ".pulsar-failures")
	entries, err := os.ReadDir(failDir)
	if err != nil {
		t.Fatalf("reading failure dir: %v; stderr=%s", err, errOut.String())
	}
	var logPath string
	for _, entry := range entries {
		if filepath.Ext(entry.Name()) == ".log" {
			logPath = filepath.Join(failDir, entry.Name())
			break
		}
	}
	if logPath == "" {
		t.Fatalf("expected a .log under %s", failDir)
	}
	content, err := os.ReadFile(logPath)
	if err != nil {
		t.Fatalf("reading failure log: %v", err)
	}
	if !strings.Contains(string(content), "***REDACTED***") {
		t.Errorf("expected redacted marker in failure log, got: %s", string(content))
	}
	if strings.Contains(string(content), "ghp_SECRET") {
		t.Errorf("failure log leaked secret: %s", string(content))
	}
	if !strings.Contains(out.String(), "failure output written to") {
		t.Errorf("expected failure log notice in stdout, got: %s", out.String())
	}
}

// TestRunRedactsBuildFailureSummary verifies build/pre-run failure output is
// redacted before it reaches terminal output.
func TestRunRedactsBuildFailureSummary(t *testing.T) {
	root := t.TempDir()
	fr := &fakeRunner{
		result: engine.RunResult{
			BuildFailed: true,
			BuildOutput: []string{
				"# github\n",
				"compiler printed ghp_SECRET\n",
			},
		},
	}
	pf := &fakeprovider.Fake{
		NameVal:        "test",
		Packages:       []string{"./..."},
		Pattern:        "^TestAcc",
		ModesVal:       testModes,
		RequirementsFn: allowAllRequirements,
		Secrets:        []string{"GITHUB_TOKEN"},
		PreflightFn: func(_ context.Context, _ string, _ provider.TestRequirements) provider.PreflightReport {
			return provider.PreflightReport{Mode: "anonymous"}
		},
	}
	d := deps{
		provider:  pf,
		newRunner: func(_ io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccA"}),
		cwd:       func() (string, error) { return root, nil },
		getenv:    getenvFromMap(map[string]string{"GITHUB_TOKEN": "ghp_SECRET"}),
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"run", "--repo-root", root}, &out, &errOut, d)
	if code != 1 {
		t.Fatalf("exit code = %d, want 1; stderr=%s", code, errOut.String())
	}
	if strings.Contains(out.String(), "ghp_SECRET") {
		t.Fatalf("terminal summary leaked secret: %s", out.String())
	}
	if !strings.Contains(out.String(), "***REDACTED***") {
		t.Fatalf("terminal summary did not redact secret: %s", out.String())
	}
}

// TestRunGreenLeavesNoFailureLogs verifies that a green run removes stale
// failure logs and leaves no new failure artifacts.
func TestRunGreenLeavesNoFailureLogs(t *testing.T) {
	root := t.TempDir()
	failDir := filepath.Join(root, ".pulsar-failures")
	if err := os.MkdirAll(failDir, 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(failDir, "._TestAccA.log"), []byte("old failure\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	fr := &fakeRunner{
		result: engine.RunResult{
			Tests: []engine.TestResult{
				{Name: "TestAccA", Status: provider.StatusPass},
			},
		},
	}
	pf := &fakeprovider.Fake{
		NameVal:        "test",
		Packages:       []string{"./..."},
		Pattern:        "^TestAcc",
		ModesVal:       testModes,
		RequirementsFn: allowAllRequirements,
		PreflightFn: func(_ context.Context, _ string, _ provider.TestRequirements) provider.PreflightReport {
			return provider.PreflightReport{Mode: "anonymous"}
		},
	}
	d := deps{
		provider:  pf,
		newRunner: func(_ io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccA"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	_ = runWithDeps([]string{"run", "--repo-root", root}, &out, &errOut, d)

	matches, err := filepath.Glob(filepath.Join(failDir, "*.log"))
	if err != nil {
		t.Fatalf("globbing failure logs: %v", err)
	}
	if len(matches) != 0 {
		t.Fatalf("expected no failure logs after green run, got %v", matches)
	}
}

// TestRunPartialRunPreservesOtherFailureLogs verifies that a filtered run only
// refreshes logs for tests it re-runs and preserves other cumulative failures.
func TestRunPartialRunPreservesOtherFailureLogs(t *testing.T) {
	root := t.TempDir()
	failDir := filepath.Join(root, ".pulsar-failures")
	if err := os.MkdirAll(failDir, 0o700); err != nil {
		t.Fatal(err)
	}
	otherLog := filepath.Join(failDir, "._TestAccB.log")
	if err := os.WriteFile(otherLog, []byte("previous failure\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	fr := &fakeRunner{
		result: engine.RunResult{
			Tests: []engine.TestResult{
				{Name: "TestAccA", Status: provider.StatusFail, Output: []string{"boom\n"}},
			},
		},
	}
	pf := &fakeprovider.Fake{
		NameVal:        "test",
		Packages:       []string{"./..."},
		Pattern:        "^TestAcc",
		ModesVal:       testModes,
		RequirementsFn: allowAllRequirements,
		PreflightFn: func(_ context.Context, _ string, _ provider.TestRequirements) provider.PreflightReport {
			return provider.PreflightReport{Mode: "anonymous"}
		},
	}
	d := deps{
		provider:  pf,
		newRunner: func(_ io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccA", "TestAccB"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	_ = runWithDeps([]string{"run", "--repo-root", root, "--run", "^TestAccA$"}, &out, &errOut, d)

	if content, err := os.ReadFile(otherLog); err != nil {
		t.Fatalf("expected other failure log to be preserved: %v; stderr=%s", err, errOut.String())
	} else if string(content) != "previous failure\n" {
		t.Fatalf("other failure log = %q; want preserved previous failure", string(content))
	}
	failLog := filepath.Join(failDir, "._TestAccA.log")
	content, err := os.ReadFile(failLog)
	if err != nil {
		t.Fatalf("expected TestAccA failure log to be written: %v; stderr=%s", err, errOut.String())
	}
	if !strings.Contains(string(content), "boom\n") {
		t.Fatalf("TestAccA failure log missing output: %q", string(content))
	}
}

// TestReportListsFailureLogs verifies terminal reports list captured failure
// logs from the last run.
func TestReportListsFailureLogs(t *testing.T) {
	root := t.TempDir()
	failDir := filepath.Join(root, ".pulsar-failures")
	if err := os.MkdirAll(failDir, 0o700); err != nil {
		t.Fatal(err)
	}
	logPath := filepath.Join(failDir, "x.log")
	if err := os.WriteFile(logPath, []byte("failure\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	pf := &fakeprovider.Fake{
		Packages: []string{"./..."},
		Pattern:  "^TestAcc",
	}
	d := deps{
		provider: pf,
		list:     stubList([]string{"TestAccA"}),
		cwd:      func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"report", "--repo-root", root}, &out, &errOut, d)

	if code != 0 {
		t.Fatalf("expected exit 0, got %d; stderr=%s", code, errOut.String())
	}
	if !strings.Contains(out.String(), logPath) {
		t.Fatalf("expected report output to list %s, got: %s", logPath, out.String())
	}
}

func TestReportCommandWithoutPathsStillPrintsSummary(t *testing.T) {
	root := cliScratchDir(t)
	statePath := filepath.Join(root, ".pulsar-state.json")
	st := engine.State{
		Provider: "test",
		Mode:     "anonymous",
		Results: []engine.PersistResult{
			{Package: "./github", Test: "TestAccRepoPass", Status: "pass"},
			{Package: "./github", Test: "TestAccRepoFail", Status: "fail"},
		},
	}
	if err := st.Save(statePath); err != nil {
		t.Fatalf("save state: %v", err)
	}
	failDir := filepath.Join(root, ".pulsar-failures")
	if err := os.MkdirAll(failDir, 0o700); err != nil {
		t.Fatalf("make failure dir: %v", err)
	}
	logPath := filepath.Join(failDir, "repo.log")
	if err := os.WriteFile(logPath, []byte("failure\n"), 0o600); err != nil {
		t.Fatalf("write failure log: %v", err)
	}
	pf := &fakeprovider.Fake{
		Packages: []string{"./..."},
		Pattern:  "^TestAcc",
		GroupFunc: func(string) string {
			return "repos"
		},
	}
	d := deps{
		provider: pf,
		list:     stubList([]string{"TestAccRepoPass", "TestAccRepoFail"}),
		cwd:      func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"report", "--repo-root", root}, &out, &errOut, d)
	if code != 0 {
		t.Fatalf("expected exit 0, got %d; stderr=%s", code, errOut.String())
	}
	stdout := out.String()
	for _, want := range []string{
		"1 passed, 1 failed, 0 skipped",
		"repos: 1 passed, 1 failed, 0 skipped",
		"failure logs from the last run (1):",
		logPath,
	} {
		if !strings.Contains(stdout, want) {
			t.Fatalf("stdout missing %q:\n%s", want, stdout)
		}
	}
}

// TestCLIRetryUsesFailedOnly verifies retry --failed sends only the failed pattern to the runner.
func TestCLIRetryUsesFailedOnly(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")

	// Seed state with one passed and one failed.
	st := engine.State{Provider: "test", Mode: "anonymous"}
	st.Plan = planWithEligible("anonymous", []string{"TestAccA", "TestAccB"})
	st.Results = []engine.PersistResult{
		{Test: "TestAccA", Status: "pass"},
		{Test: "TestAccB", Status: "fail"},
	}
	if err := st.Save(statePath); err != nil {
		t.Fatalf("seeding state: %v", err)
	}

	fr := &fakeRunner{result: engine.RunResult{}}
	pf := &fakeprovider.Fake{
		Packages: []string{"./..."},
		Pattern:  "^TestAcc",
	}
	d := deps{
		provider:  pf,
		newRunner: func(_ io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccA", "TestAccB"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	_ = runWithDeps([]string{"retry", "--failed", "--repo-root", root}, &out, &errOut, d)

	wantPattern := engine.RunPattern([]string{"TestAccB"})
	if fr.calls != 1 {
		t.Errorf("expected 1 runner call, got %d", fr.calls)
	}
	if fr.spec.Pattern != wantPattern {
		t.Errorf("pattern = %q, want %q", fr.spec.Pattern, wantPattern)
	}
}

// TestCLIRetryNoFailuresNoOp verifies retry --failed with an all-pass state prints
// "nothing to run", returns exit 0, and invokes the runner zero times.
func TestCLIRetryNoFailuresNoOp(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")

	st := engine.State{Provider: "test", Mode: "anonymous"}
	st.Plan = planWithEligible("anonymous", []string{"TestAccA", "TestAccB"})
	st.Results = []engine.PersistResult{
		{Test: "TestAccA", Status: "pass"},
		{Test: "TestAccB", Status: "pass"},
	}
	if err := st.Save(statePath); err != nil {
		t.Fatalf("seeding state: %v", err)
	}

	fr := &fakeRunner{}
	pf := &fakeprovider.Fake{Packages: []string{"./..."}, Pattern: "^TestAcc"}
	d := deps{
		provider:  pf,
		newRunner: func(_ io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccA", "TestAccB"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"retry", "--failed", "--repo-root", root}, &out, &errOut, d)

	if code != 0 {
		t.Errorf("expected exit 0, got %d", code)
	}
	if fr.calls != 0 {
		t.Errorf("expected 0 runner calls (no-op rail), got %d", fr.calls)
	}
	if !strings.Contains(out.String(), "nothing to run") {
		t.Errorf("expected 'nothing to run' in output; got: %s", out.String())
	}
}

// TestCLIResumeAllPassedNoOp verifies resume with every test already passed
// exits 0 and invokes the runner zero times.
func TestCLIResumeAllPassedNoOp(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")

	st := engine.State{Provider: "test", Mode: "anonymous"}
	st.Plan = planWithEligible("anonymous", []string{"TestAccA", "TestAccB"})
	st.Results = []engine.PersistResult{
		{Test: "TestAccA", Status: "pass"},
		{Test: "TestAccB", Status: "pass"},
	}
	if err := st.Save(statePath); err != nil {
		t.Fatalf("seeding state: %v", err)
	}

	fr := &fakeRunner{}
	pf := &fakeprovider.Fake{Packages: []string{"./..."}, Pattern: "^TestAcc"}
	d := deps{
		provider:  pf,
		newRunner: func(_ io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccA", "TestAccB"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"resume", "--repo-root", root}, &out, &errOut, d)

	if code != 0 {
		t.Errorf("expected exit 0, got %d", code)
	}
	if fr.calls != 0 {
		t.Errorf("expected 0 runner calls (no-op rail), got %d", fr.calls)
	}
	if !strings.Contains(out.String(), "nothing to resume") {
		t.Errorf("expected 'nothing to resume' in output; got: %s", out.String())
	}
}

// TestCLIResumeUsesFailedAndNotRun verifies resume builds a pattern from
// failed + not-yet-run tests and invokes the runner exactly once.
func TestCLIResumeUsesFailedAndNotRun(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")

	// TestAccA passed, TestAccB failed; TestAccC is not in state (not run).
	st := engine.State{Provider: "test", Mode: "anonymous"}
	st.Plan = planWithEligible("anonymous", []string{"TestAccA", "TestAccB", "TestAccC"})
	st.Results = []engine.PersistResult{
		{Test: "TestAccA", Status: "pass"},
		{Test: "TestAccB", Status: "fail"},
	}
	if err := st.Save(statePath); err != nil {
		t.Fatalf("seeding state: %v", err)
	}

	fr := &fakeRunner{result: engine.RunResult{}}
	pf := &fakeprovider.Fake{Packages: []string{"./..."}, Pattern: "^TestAcc"}
	d := deps{
		provider:  pf,
		newRunner: func(_ io.Writer) testRunner { return fr },
		// TestAccC is in the full list but not in state → not run
		list: stubList([]string{"TestAccA", "TestAccB", "TestAccC"}),
		cwd:  func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	_ = runWithDeps([]string{"resume", "--repo-root", root}, &out, &errOut, d)

	if fr.calls != 1 {
		t.Errorf("expected 1 runner call, got %d", fr.calls)
	}
	// Pattern must include TestAccB (failed) and TestAccC (not run), but not TestAccA (passed).
	wantPattern := engine.RunPattern([]string{"TestAccB", "TestAccC"})
	if fr.spec.Pattern != wantPattern {
		t.Errorf("pattern = %q, want %q", fr.spec.Pattern, wantPattern)
	}
}

// TestCLIRepoRootResolved verifies FindRepoRoot auto-resolution and --repo-root override.
func TestCLIRepoRootResolved(t *testing.T) {
	tmp := t.TempDir()
	if err := os.Mkdir(filepath.Join(tmp, ".git"), 0o755); err != nil {
		t.Fatal(err)
	}
	sub := filepath.Join(tmp, "sub")
	if err := os.Mkdir(sub, 0o755); err != nil {
		t.Fatal(err)
	}

	pf := &fakeprovider.Fake{Packages: []string{"./..."}, Pattern: "^TestAcc"}

	// Test auto-resolve: cwd returns sub, root should be resolved to tmp (ancestor with .git).
	var capturedDir string
	d := deps{
		provider: pf,
		list: func(_ context.Context, dir string, _ []string, _ string) ([]string, error) {
			capturedDir = dir
			return nil, nil
		},
		cwd: func() (string, error) { return sub, nil },
	}
	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"groups"}, &out, &errOut, d)
	if code != 0 {
		t.Errorf("auto-resolve: exit code = %d, want 0; stderr=%s", code, errOut.String())
	}
	if capturedDir != tmp {
		t.Errorf("auto-resolve: root = %q, want %q", capturedDir, tmp)
	}

	// Test --repo-root override.
	var capturedDir2 string
	override := t.TempDir()
	d2 := deps{
		provider: pf,
		list: func(_ context.Context, dir string, _ []string, _ string) ([]string, error) {
			capturedDir2 = dir
			return nil, nil
		},
		cwd: func() (string, error) { return sub, nil },
	}
	var out2, errOut2 bytes.Buffer
	code2 := runWithDeps([]string{"groups", "--repo-root", override}, &out2, &errOut2, d2)
	if code2 != 0 {
		t.Errorf("override: exit code = %d, want 0; stderr=%s", code2, errOut2.String())
	}
	if capturedDir2 != override {
		t.Errorf("override: root = %q, want %q", capturedDir2, override)
	}
}

// TestCLIStripsTFLogByDefault verifies that RunSpec.AllowSensitiveLogs is false
// by default and true when --allow-sensitive-logs is given.
func TestCLIStripsTFLogByDefault(t *testing.T) {
	root := t.TempDir()

	pf := &fakeprovider.Fake{
		NameVal:        "test",
		Packages:       []string{"./..."},
		Pattern:        "^TestAcc",
		ModesVal:       testModes,
		RequirementsFn: allowAllRequirements,
		PreflightFn: func(_ context.Context, _ string, _ provider.TestRequirements) provider.PreflightReport {
			return provider.PreflightReport{Mode: "anonymous"}
		},
	}

	fr := &fakeRunner{result: engine.RunResult{}}
	d := deps{
		provider:  pf,
		newRunner: func(_ io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccFoo"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	_ = runWithDeps([]string{"run", "--repo-root", root}, &out, &errOut, d)

	if fr.spec.AllowSensitiveLogs {
		t.Error("AllowSensitiveLogs should be false by default")
	}

	// With --allow-sensitive-logs flag.
	fr2 := &fakeRunner{result: engine.RunResult{}}
	d.newRunner = func(_ io.Writer) testRunner { return fr2 }
	out.Reset()
	errOut.Reset()
	_ = runWithDeps([]string{"run", "--repo-root", root, "--allow-sensitive-logs"}, &out, &errOut, d)

	if !fr2.spec.AllowSensitiveLogs {
		t.Error("AllowSensitiveLogs should be true with --allow-sensitive-logs")
	}
}

// TestCLIModeDefaultsFromEnv verifies subcommands honor GH_TEST_AUTH_MODE when
// --mode is omitted.
func TestCLIModeDefaultsFromEnv(t *testing.T) {
	root := t.TempDir()
	var preflightMode string
	fr := &fakeRunner{result: engine.RunResult{}}
	pf := &fakeprovider.Fake{
		NameVal:        "test",
		Packages:       []string{"./..."},
		Pattern:        "^TestAcc",
		ModesVal:       testModes,
		RequirementsFn: allowAllRequirements,
		EnvByMode: map[string][]provider.EnvVar{
			"organization": {{Key: "GITHUB_OWNER"}},
		},
		PreflightFn: func(_ context.Context, mode string, _ provider.TestRequirements) provider.PreflightReport {
			preflightMode = mode
			return provider.PreflightReport{Mode: mode}
		},
	}
	d := deps{
		provider:  pf,
		newRunner: func(_ io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccA"}),
		cwd:       func() (string, error) { return root, nil },
		getenv:    getenvFromMap(map[string]string{"GH_TEST_AUTH_MODE": "organization"}),
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"run", "--repo-root", root}, &out, &errOut, d)
	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s", code, errOut.String())
	}
	if preflightMode != "organization" {
		t.Fatalf("preflight mode = %q, want organization", preflightMode)
	}
	if !extraEnvContains(fr.spec.ExtraEnv, "GH_TEST_AUTH_MODE=organization") {
		t.Fatalf("child ExtraEnv = %v, want GH_TEST_AUTH_MODE=organization", fr.spec.ExtraEnv)
	}
}

// TestCLISweepRequiresConfirm verifies that sweep without --confirm exits non-zero
// and never calls the provider Sweep method.
func TestCLISweepRequiresConfirm(t *testing.T) {
	sweepCalls := 0
	pf := &fakeprovider.Fake{
		SweepFn: func(_ context.Context, _ string, _ provider.SweepOpts) error {
			sweepCalls++
			return nil
		},
	}
	d := deps{
		provider: pf,
		list:     stubList(nil),
		cwd:      func() (string, error) { return t.TempDir(), nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"sweep"}, &out, &errOut, d)

	if code == 0 {
		t.Errorf("expected non-zero exit, got 0")
	}
	if sweepCalls != 0 {
		t.Errorf("expected 0 Sweep calls, got %d", sweepCalls)
	}
}

// TestCLINonTTYNoTUI verifies that bare invocation prints help and returns 0 without blocking.
func TestCLINonTTYNoTUI(t *testing.T) {
	pf := &fakeprovider.Fake{}
	d := deps{
		provider:  pf,
		newRunner: func(_ io.Writer) testRunner { return &fakeRunner{} },
		list:      stubList(nil),
		cwd:       func() (string, error) { return t.TempDir(), nil },
	}

	done := make(chan int, 1)
	go func() {
		var out, errOut bytes.Buffer
		done <- runWithDeps([]string{}, &out, &errOut, d)
	}()

	var code int
	select {
	case code = <-done:
	case <-time.After(3 * time.Second):
		t.Fatal("runWithDeps blocked; expected prompt return for non-TTY")
	}

	if code != 0 {
		t.Errorf("expected exit 0, got %d", code)
	}
}

// TestCLIExitCodes verifies the exit-code contract for run results.
func TestCLIExitCodes(t *testing.T) {
	root := t.TempDir()
	pf := &fakeprovider.Fake{
		Packages:       []string{"./..."},
		Pattern:        "^Test",
		ModesVal:       testModes,
		RequirementsFn: allowAllRequirements,
		PreflightFn: func(_ context.Context, _ string, _ provider.TestRequirements) provider.PreflightReport {
			return provider.PreflightReport{Mode: "anonymous"}
		},
	}

	cases := []struct {
		name     string
		result   engine.RunResult
		wantCode int
	}{
		{
			name: "failed test",
			result: engine.RunResult{
				Tests: []engine.TestResult{{Name: "TestFoo", Status: provider.StatusFail}},
			},
			wantCode: 1,
		},
		{
			name:     "build failed",
			result:   engine.RunResult{BuildFailed: true},
			wantCode: 1,
		},
		{
			name:     "pre-run failed",
			result:   engine.RunResult{PreRunFailed: true},
			wantCode: 1,
		},
		{
			name: "all pass",
			result: engine.RunResult{
				Tests: []engine.TestResult{{Name: "TestFoo", Status: provider.StatusPass}},
			},
			wantCode: 0,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			fr := &fakeRunner{result: tc.result}
			d := deps{
				provider:  pf,
				newRunner: func(_ io.Writer) testRunner { return fr },
				list:      stubList([]string{"TestFoo"}),
				cwd:       func() (string, error) { return root, nil },
			}
			var out, errOut bytes.Buffer
			code := runWithDeps([]string{"run", "--repo-root", root}, &out, &errOut, d)
			if code != tc.wantCode {
				t.Errorf("exit code = %d, want %d", code, tc.wantCode)
			}
		})
	}
}

// TestCLIRunBuildFailureStateFeedsReport verifies a build failure replaces a
// stale green state in later reports instead of preserving old passing results
// as the apparent latest run.
func TestCLIRunBuildFailureStateFeedsReport(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")
	st := engine.State{Provider: "test", Mode: "anonymous"}
	st.Results = []engine.PersistResult{{Test: "TestAccA", Status: "pass"}}
	if err := st.Save(statePath); err != nil {
		t.Fatalf("seeding state: %v", err)
	}

	fr := &fakeRunner{result: engine.RunResult{
		BuildFailed: true,
		BuildOutput: []string{"compile exploded\n"},
	}}
	pf := &fakeprovider.Fake{
		NameVal:        "test",
		Packages:       []string{"./..."},
		Pattern:        "^TestAcc",
		ModesVal:       testModes,
		RequirementsFn: allowAllRequirements,
		PreflightFn: func(_ context.Context, _ string, _ provider.TestRequirements) provider.PreflightReport {
			return provider.PreflightReport{Mode: "anonymous"}
		},
	}
	d := deps{
		provider:  pf,
		newRunner: func(_ io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccA"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	if code := runWithDeps([]string{"run", "--repo-root", root}, &out, &errOut, d); code != 1 {
		t.Fatalf("run exit code = %d, want 1; stderr=%s", code, errOut.String())
	}

	out.Reset()
	errOut.Reset()
	if code := runWithDeps([]string{"report", "--repo-root", root}, &out, &errOut, d); code != 0 {
		t.Fatalf("report exit code = %d, want 0; stderr=%s", code, errOut.String())
	}
	if !strings.Contains(out.String(), "BUILD FAILED") {
		t.Fatalf("report = %s, want BUILD FAILED from latest run", out.String())
	}
	if strings.Contains(out.String(), "1 passed, 0 failed") {
		t.Fatalf("report used stale passing state after build failure: %s", out.String())
	}
}

// TestCLIPreflightGateBlocksRun verifies that a failing preflight exits 1
// and the runner is invoked zero times.
func TestCLIPreflightGateBlocksRun(t *testing.T) {
	root := t.TempDir()
	fr := &fakeRunner{}
	pf := &fakeprovider.Fake{
		Packages:       []string{"./..."},
		Pattern:        "^TestAcc",
		ModesVal:       testModes,
		RequirementsFn: allowAllRequirements,
		PreflightFn: func(_ context.Context, _ string, _ provider.TestRequirements) provider.PreflightReport {
			return provider.PreflightReport{
				Mode: "anonymous",
				Checks: []provider.Check{
					{
						Name:   "token",
						Status: provider.CheckFail,
						Detail: "missing token",
						Fix:    "set GITHUB_TOKEN",
					},
				},
			}
		},
	}
	d := deps{
		provider:  pf,
		newRunner: func(_ io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccFoo"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"run", "--repo-root", root}, &out, &errOut, d)

	if code != 1 {
		t.Errorf("expected exit 1, got %d", code)
	}
	if fr.calls != 0 {
		t.Errorf("expected 0 runner calls, got %d", fr.calls)
	}
}

// ── planning integration (Task 6): ordering, redaction, and standalone
// preflight parity ─────────────────────────────────────────────────────────

// TestRunOrderDiscoversBeforePreflight verifies that test discovery happens
// before the provider preflight gate runs, locking in the brief's required
// resolve-root -> discover/plan -> preflight ordering.
func TestRunOrderDiscoversBeforePreflight(t *testing.T) {
	root := t.TempDir()
	var order []string
	fr := &fakeRunner{}
	pf := &fakeprovider.Fake{
		Packages:       []string{"./..."},
		Pattern:        "^TestAcc",
		ModesVal:       testModes,
		RequirementsFn: allowAllRequirements,
		PreflightFn: func(_ context.Context, _ string, _ provider.TestRequirements) provider.PreflightReport {
			order = append(order, "preflight")
			return provider.PreflightReport{Mode: "anonymous"}
		},
	}
	d := deps{
		provider:  pf,
		newRunner: func(_ io.Writer) testRunner { return fr },
		list: func(_ context.Context, _ string, _ []string, _ string) ([]string, error) {
			order = append(order, "list")
			return []string{"TestAccFoo"}, nil
		},
		cwd: func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"run", "--repo-root", root}, &out, &errOut, d)
	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s", code, errOut.String())
	}
	listIdx, preflightIdx := -1, -1
	for i, ev := range order {
		if ev == "list" && listIdx == -1 {
			listIdx = i
		}
		if ev == "preflight" && preflightIdx == -1 {
			preflightIdx = i
		}
	}
	if listIdx == -1 || preflightIdx == -1 || listIdx > preflightIdx {
		t.Fatalf("expected discovery (list) before preflight; order=%v", order)
	}
}

// TestRunPlanFailureSkipsPreflightAndRunner verifies that when planning fails
// (an unknown --group), neither Preflight nor the runner is invoked, and
// nothing is written to stdout - a planning failure must short-circuit
// before either later step runs.
func TestRunPlanFailureSkipsPreflightAndRunner(t *testing.T) {
	root := t.TempDir()
	fr := &fakeRunner{}
	preflightCalls := 0
	pf := &fakeprovider.Fake{
		Packages:       []string{"./..."},
		Pattern:        "^TestAcc",
		ModesVal:       testModes,
		RequirementsFn: allowAllRequirements,
		PreflightFn: func(_ context.Context, _ string, _ provider.TestRequirements) provider.PreflightReport {
			preflightCalls++
			return provider.PreflightReport{Mode: "anonymous"}
		},
	}
	d := deps{
		provider:  pf,
		newRunner: func(_ io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccFoo"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"run", "--group", "does-not-exist", "--repo-root", root}, &out, &errOut, d)

	if code != 1 {
		t.Errorf("exit code = %d, want 1 (plan/group-unknown)", code)
	}
	if preflightCalls != 0 {
		t.Errorf("expected 0 preflight calls, got %d", preflightCalls)
	}
	if fr.calls != 0 {
		t.Errorf("expected 0 runner calls, got %d", fr.calls)
	}
	if out.String() != "" {
		t.Errorf("expected no stdout on planning failure, got: %q", out.String())
	}
}

// TestRunPreflightReceivesPlanAggregatedRequirements verifies that Preflight
// is called with exactly the plan's aggregated Scopes/Capabilities/
// SideEffects across eligible tests - not an empty TestRequirements{} and not
// a per-test value.
func TestRunPreflightReceivesPlanAggregatedRequirements(t *testing.T) {
	root := t.TempDir()
	fr := &fakeRunner{}
	var got provider.TestRequirements
	pf := &fakeprovider.Fake{
		Packages: []string{"./..."},
		Pattern:  "^TestAcc",
		ModesVal: testModes,
		RequirementsFn: func(name string) (provider.TestRequirements, bool) {
			switch name {
			case "TestAccRepo":
				return provider.TestRequirements{Scopes: []string{"repo"}, SideEffects: []string{"repository"}}, true
			case "TestAccOrg":
				return provider.TestRequirements{Scopes: []string{"read:org"}, Capabilities: []string{"organization"}}, true
			}
			return provider.TestRequirements{}, true
		},
		PreflightFn: func(_ context.Context, _ string, requirements provider.TestRequirements) provider.PreflightReport {
			got = requirements
			return provider.PreflightReport{Mode: "anonymous"}
		},
	}
	d := deps{
		provider:  pf,
		newRunner: func(_ io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccRepo", "TestAccOrg"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"run", "--repo-root", root}, &out, &errOut, d)
	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s", code, errOut.String())
	}
	if want := []string{"read:org", "repo"}; !reflect.DeepEqual(got.Scopes, want) {
		t.Errorf("Preflight requirements.Scopes = %v, want %v", got.Scopes, want)
	}
	if want := []string{"organization"}; !reflect.DeepEqual(got.Capabilities, want) {
		t.Errorf("Preflight requirements.Capabilities = %v, want %v", got.Capabilities, want)
	}
	if want := []string{"repository"}; !reflect.DeepEqual(got.SideEffects, want) {
		t.Errorf("Preflight requirements.SideEffects = %v, want %v", got.SideEffects, want)
	}
}

// TestRunPatternMatchesPlanEligible verifies the runner receives exactly
// engine.RunPattern(plan.Eligible) - so a test excluded by planning (here:
// mode-incompatible) is never passed to the runner, even though it was
// discovered.
func TestRunPatternMatchesPlanEligible(t *testing.T) {
	root := t.TempDir()
	fr := &fakeRunner{}
	pf := &fakeprovider.Fake{
		Packages: []string{"./..."},
		Pattern:  "^TestAcc",
		ModesVal: testModes,
		RequirementsFn: func(name string) (provider.TestRequirements, bool) {
			if name == "TestAccOrgOnly" {
				return provider.TestRequirements{Modes: []string{"organization"}}, true
			}
			return provider.TestRequirements{}, true
		},
		PreflightFn: func(_ context.Context, _ string, _ provider.TestRequirements) provider.PreflightReport {
			return provider.PreflightReport{Mode: "anonymous"}
		},
	}
	d := deps{
		provider:  pf,
		newRunner: func(_ io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccKeep", "TestAccOrgOnly"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"run", "--mode", "anonymous", "--repo-root", root}, &out, &errOut, d)
	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s", code, errOut.String())
	}
	if fr.calls != 1 {
		t.Fatalf("expected 1 runner call, got %d", fr.calls)
	}
	want := engine.RunPattern([]string{"TestAccKeep"})
	if fr.spec.Pattern != want {
		t.Errorf("runner pattern = %q, want %q (plan.Eligible only)", fr.spec.Pattern, want)
	}
}

// TestRunEmptyPlanSkipsRunner verifies that when planning succeeds but no
// test is eligible (an empty plan), run still runs preflight and emits its
// plan summary, but never invokes the runner.
func TestRunEmptyPlanSkipsRunner(t *testing.T) {
	root := t.TempDir()
	fr := &fakeRunner{}
	preflightCalls := 0
	pf := &fakeprovider.Fake{
		Packages:       []string{"./..."},
		Pattern:        "^TestAcc",
		ModesVal:       testModes,
		RequirementsFn: allowAllRequirements,
		PreflightFn: func(_ context.Context, _ string, _ provider.TestRequirements) provider.PreflightReport {
			preflightCalls++
			return provider.PreflightReport{Mode: "anonymous"}
		},
	}
	d := deps{
		provider:  pf,
		newRunner: func(_ io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccFoo"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"run", "--run", "^NoSuchTest$", "--repo-root", root}, &out, &errOut, d)

	if code != 0 {
		t.Errorf("exit code = %d, want 0", code)
	}
	if preflightCalls != 1 {
		t.Errorf("expected preflight to still run once for an empty plan, got %d calls", preflightCalls)
	}
	if fr.calls != 0 {
		t.Errorf("expected 0 runner calls for an empty plan, got %d", fr.calls)
	}
	if !strings.Contains(out.String(), "nothing to run") {
		t.Errorf("expected 'nothing to run' in output; got: %s", out.String())
	}
	if !strings.Contains(out.String(), "mode anonymous: 0 selected, 0 eligible, 0 excluded, 0 unclassified") {
		t.Errorf("expected plan summary line in output; got: %s", out.String())
	}
}

// TestRunTextShowsModeSummaryLine verifies run's text output leads with the
// exact "mode <mode>: N selected, N eligible, N excluded, N unclassified"
// line naming the selected mode.
func TestRunTextShowsModeSummaryLine(t *testing.T) {
	root := t.TempDir()
	fr := &fakeRunner{result: engine.RunResult{Tests: []engine.TestResult{{Name: "TestAccFoo", Status: provider.StatusPass}}}}
	pf := &fakeprovider.Fake{
		Packages:       []string{"./..."},
		Pattern:        "^TestAcc",
		ModesVal:       testModes,
		RequirementsFn: allowAllRequirements,
		PreflightFn: func(_ context.Context, _ string, _ provider.TestRequirements) provider.PreflightReport {
			return provider.PreflightReport{Mode: "organization"}
		},
	}
	d := deps{
		provider:  pf,
		newRunner: func(_ io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccFoo"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"run", "--mode", "organization", "--repo-root", root}, &out, &errOut, d)
	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s", code, errOut.String())
	}
	want := "mode organization: 1 selected, 1 eligible, 0 excluded, 0 unclassified\n"
	if !strings.HasPrefix(out.String(), want) {
		t.Errorf("expected output to start with %q; got: %s", want, out.String())
	}
}

// TestRunTextPreflightFailureRedactsCheckSecret is a regression test for the
// Task 4 follow-up: printPreflight must route Check.Detail/Fix through
// redactorForProvider before printing, not merely rely on the provider's own
// unit tests. A configured secret embedded in a failing Check must never
// reach run's text output - only the redaction marker should.
func TestRunTextPreflightFailureRedactsCheckSecret(t *testing.T) {
	const secret = "ghp_SUPERSECRETVALUE"
	root := t.TempDir()
	fr := &fakeRunner{}
	pf := &fakeprovider.Fake{
		Packages:       []string{"./..."},
		Pattern:        "^TestAcc",
		ModesVal:       testModes,
		Secrets:        []string{"GITHUB_TOKEN"},
		RequirementsFn: allowAllRequirements,
		PreflightFn: func(_ context.Context, _ string, _ provider.TestRequirements) provider.PreflightReport {
			return provider.PreflightReport{
				Mode: "anonymous",
				Checks: []provider.Check{{
					Name:   "identity",
					Status: provider.CheckFail,
					Detail: "token " + secret + " rejected by GitHub",
					Fix:    "rotate " + secret + " and retry",
				}},
			}
		},
	}
	d := deps{
		provider:  pf,
		newRunner: func(_ io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccFoo"}),
		cwd:       func() (string, error) { return root, nil },
		getenv: func(key string) string {
			if key == "GITHUB_TOKEN" {
				return secret
			}
			return ""
		},
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"run", "--repo-root", root}, &out, &errOut, d)

	if code != 1 {
		t.Fatalf("exit code = %d, want 1", code)
	}
	if strings.Contains(errOut.String(), secret) {
		t.Errorf("preflight failure text leaked secret: %s", errOut.String())
	}
	if !strings.Contains(errOut.String(), "***REDACTED***") {
		t.Errorf("expected redaction marker in preflight failure text; got: %s", errOut.String())
	}
}

// TestPreflightStandaloneFlagParitySelectsSubset verifies standalone
// preflight accepts --group and --run with the same selection semantics as
// run: an unknown --group is a planning error (exit 1), and --run narrows
// which tests contribute to the aggregated requirements that get checked.
func TestPreflightStandaloneFlagParitySelectsSubset(t *testing.T) {
	root := t.TempDir()
	var got provider.TestRequirements
	pf := &fakeprovider.Fake{
		Packages: []string{"./..."},
		Pattern:  "^TestAcc",
		ModesVal: testModes,
		RequirementsFn: func(name string) (provider.TestRequirements, bool) {
			if name == "TestAccOrg" {
				return provider.TestRequirements{Scopes: []string{"read:org"}}, true
			}
			return provider.TestRequirements{Scopes: []string{"repo"}}, true
		},
		PreflightFn: func(_ context.Context, _ string, requirements provider.TestRequirements) provider.PreflightReport {
			got = requirements
			return provider.PreflightReport{Mode: "anonymous"}
		},
	}
	d := deps{
		provider: pf,
		list:     stubList([]string{"TestAccOrg", "TestAccRepo"}),
		cwd:      func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"preflight", "--run", "^TestAccOrg$", "--repo-root", root}, &out, &errOut, d)
	if code != 0 {
		t.Fatalf("--run: exit code = %d, want 0; stderr=%s", code, errOut.String())
	}
	if want := []string{"read:org"}; !reflect.DeepEqual(got.Scopes, want) {
		t.Errorf("--run: Preflight requirements.Scopes = %v, want %v (only TestAccOrg selected)", got.Scopes, want)
	}

	out.Reset()
	errOut.Reset()
	code = runWithDeps([]string{"preflight", "--group", "does-not-exist", "--repo-root", root}, &out, &errOut, d)
	if code != 1 {
		t.Errorf("unknown --group exit code = %d, want 1", code)
	}

	out.Reset()
	errOut.Reset()
	code = runWithDeps([]string{"preflight", "--repo-root", root}, &out, &errOut, d)
	if code != 0 {
		t.Errorf("default selection exit code = %d, want 0; stderr=%s", code, errOut.String())
	}
	if want := []string{"read:org", "repo"}; !reflect.DeepEqual(got.Scopes, want) {
		t.Errorf("default: Preflight requirements.Scopes = %v, want %v (both tests selected)", got.Scopes, want)
	}
}

// TestPreflightRunOverridesGroupSelection verifies at the CLI boundary that
// `--run` overrides `--group` exactly as the flag help and CLI reference
// promise: a regex pointing OUTSIDE the named group still selects its match
// (rather than intersecting to nothing), and an unknown `--group` is inert
// once `--run` is supplied instead of failing with plan/group-unknown.
func TestPreflightRunOverridesGroupSelection(t *testing.T) {
	root := t.TempDir()
	var got provider.TestRequirements
	pf := &fakeprovider.Fake{
		Packages: []string{"./..."},
		Pattern:  "^TestAcc",
		ModesVal: testModes,
		GroupFunc: func(name string) string {
			if name == "TestAccOrg" {
				return "organizations"
			}
			return "repositories"
		},
		RequirementsFn: func(name string) (provider.TestRequirements, bool) {
			if name == "TestAccOrg" {
				return provider.TestRequirements{Scopes: []string{"read:org"}}, true
			}
			return provider.TestRequirements{Scopes: []string{"repo"}}, true
		},
		PreflightFn: func(_ context.Context, _ string, requirements provider.TestRequirements) provider.PreflightReport {
			got = requirements
			return provider.PreflightReport{Mode: "anonymous"}
		},
	}
	d := deps{
		provider: pf,
		list:     stubList([]string{"TestAccOrg", "TestAccRepo"}),
		cwd:      func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"preflight", "--group", "repositories", "--run", "^TestAccOrg$", "--repo-root", root}, &out, &errOut, d)
	if code != 0 {
		t.Fatalf("--group + --run: exit code = %d, want 0; stderr=%s", code, errOut.String())
	}
	if want := []string{"read:org"}; !reflect.DeepEqual(got.Scopes, want) {
		t.Errorf("--group + --run: requirements.Scopes = %v, want %v (--run must override --group)", got.Scopes, want)
	}

	out.Reset()
	errOut.Reset()
	code = runWithDeps([]string{"preflight", "--group", "does-not-exist", "--run", "^TestAccOrg$", "--repo-root", root}, &out, &errOut, d)
	if code != 0 {
		t.Fatalf("unknown --group with --run: exit code = %d, want 0; stderr=%s", code, errOut.String())
	}
	if want := []string{"read:org"}; !reflect.DeepEqual(got.Scopes, want) {
		t.Errorf("unknown --group with --run: requirements.Scopes = %v, want %v", got.Scopes, want)
	}
}

// TestPreflightStandaloneAllowUnclassifiedFlag verifies that an unclassified
// test blocks standalone preflight by default (plan/unclassified) but is
// accepted, with a text warning, once --allow-unclassified is passed -
// matching run's hardcoded AllowUnclassified: true behavior on demand.
func TestPreflightStandaloneAllowUnclassifiedFlag(t *testing.T) {
	root := t.TempDir()
	pf := &fakeprovider.Fake{
		Packages: []string{"./..."},
		Pattern:  "^TestAcc",
		ModesVal: testModes,
		// No RequirementsFn: every test is unclassified (RequirementsFor returns ok=false).
		PreflightFn: func(_ context.Context, _ string, _ provider.TestRequirements) provider.PreflightReport {
			return provider.PreflightReport{Mode: "anonymous"}
		},
	}
	d := deps{
		provider: pf,
		list:     stubList([]string{"TestAccMystery"}),
		cwd:      func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"preflight", "--repo-root", root}, &out, &errOut, d)
	if code != 1 {
		t.Errorf("exit code without --allow-unclassified = %d, want 1", code)
	}

	out.Reset()
	errOut.Reset()
	code = runWithDeps([]string{"preflight", "--allow-unclassified", "--repo-root", root}, &out, &errOut, d)
	if code != 0 {
		t.Errorf("exit code with --allow-unclassified = %d, want 0; stderr=%s", code, errOut.String())
	}
	if !strings.Contains(out.String(), "warning: TestAccMystery is unclassified") {
		t.Errorf("expected unclassified warning in output; got: %s", out.String())
	}
}

// TestPreflightStandaloneTextRedactsCheckSecret mirrors
// TestRunTextPreflightFailureRedactsCheckSecret for the standalone preflight
// subcommand's own printPreflight call site (Task 4 follow-up).
func TestPreflightStandaloneTextRedactsCheckSecret(t *testing.T) {
	const secret = "ghp_SUPERSECRETVALUE"
	root := t.TempDir()
	pf := &fakeprovider.Fake{
		Packages:       []string{"./..."},
		Pattern:        "^TestAcc",
		ModesVal:       testModes,
		Secrets:        []string{"GITHUB_TOKEN"},
		RequirementsFn: allowAllRequirements,
		PreflightFn: func(_ context.Context, _ string, _ provider.TestRequirements) provider.PreflightReport {
			return provider.PreflightReport{
				Mode: "anonymous",
				Checks: []provider.Check{{
					Name:   "identity",
					Status: provider.CheckFail,
					Detail: "token " + secret + " rejected by GitHub",
					Fix:    "rotate " + secret + " and retry",
				}},
			}
		},
	}
	d := deps{
		provider: pf,
		list:     stubList([]string{"TestAccFoo"}),
		cwd:      func() (string, error) { return root, nil },
		getenv: func(key string) string {
			if key == "GITHUB_TOKEN" {
				return secret
			}
			return ""
		},
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"preflight", "--repo-root", root}, &out, &errOut, d)

	if code != 1 {
		t.Fatalf("exit code = %d, want 1", code)
	}
	if strings.Contains(out.String(), secret) {
		t.Errorf("preflight text leaked secret: %s", out.String())
	}
	if !strings.Contains(out.String(), "***REDACTED***") {
		t.Errorf("expected redaction marker in preflight text; got: %s", out.String())
	}
}

// findResult returns the first PersistResult with the given test name, or nil.
func findResult(results []engine.PersistResult, name string) *engine.PersistResult {
	for i := range results {
		if results[i].Test == name {
			return &results[i]
		}
	}
	return nil
}

// TestCLIRetryMergesStateKeepsPassed verifies that retry --failed preserves
// previously-passed results in the on-disk state (so a later resume does not
// re-run already-passed tests). Reproduces the multi-session resume requirement.
func TestCLIRetryMergesStateKeepsPassed(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")

	st := engine.State{Provider: "test", Mode: "anonymous"}
	st.Plan = planWithEligible("anonymous", []string{"TestAccA", "TestAccB"})
	st.Results = []engine.PersistResult{
		{Test: "TestAccA", Status: "pass"},
		{Test: "TestAccB", Status: "fail"},
	}
	if err := st.Save(statePath); err != nil {
		t.Fatalf("seeding state: %v", err)
	}

	// B now passes on retry.
	fr := &fakeRunner{result: engine.RunResult{
		Tests: []engine.TestResult{{Name: "TestAccB", Status: provider.StatusPass}},
	}}
	pf := &fakeprovider.Fake{Packages: []string{"./..."}, Pattern: "^TestAcc"}
	d := deps{
		provider:  pf,
		newRunner: func(_ io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccA", "TestAccB"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	_ = runWithDeps([]string{"retry", "--failed", "--repo-root", root}, &out, &errOut, d)

	loaded, err := engine.Load(statePath)
	if err != nil {
		t.Fatalf("loading state: %v", err)
	}
	a := findResult(loaded.Results, "TestAccA")
	if a == nil {
		t.Fatalf("previously-passed TestAccA dropped from state after retry; results=%+v", loaded.Results)
	}
	if a.Status != "pass" {
		t.Errorf("TestAccA status = %q, want pass", a.Status)
	}
	b := findResult(loaded.Results, "TestAccB")
	if b == nil || b.Status != "pass" {
		t.Errorf("TestAccB should be updated to pass; got %+v", b)
	}
}

// TestCLIResumeMergesStateKeepsPassed verifies resume preserves previously-passed
// results while recording the resumed (failed + not-run) tests.
func TestCLIResumeMergesStateKeepsPassed(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")

	st := engine.State{Provider: "test", Mode: "anonymous"}
	st.Plan = planWithEligible("anonymous", []string{"TestAccA", "TestAccB", "TestAccC"})
	st.Results = []engine.PersistResult{
		{Test: "TestAccA", Status: "pass"},
		{Test: "TestAccB", Status: "fail"},
	}
	if err := st.Save(statePath); err != nil {
		t.Fatalf("seeding state: %v", err)
	}

	fr := &fakeRunner{result: engine.RunResult{
		Tests: []engine.TestResult{
			{Name: "TestAccB", Status: provider.StatusPass},
			{Name: "TestAccC", Status: provider.StatusPass},
		},
	}}
	pf := &fakeprovider.Fake{Packages: []string{"./..."}, Pattern: "^TestAcc"}
	d := deps{
		provider:  pf,
		newRunner: func(_ io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccA", "TestAccB", "TestAccC"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	_ = runWithDeps([]string{"resume", "--repo-root", root}, &out, &errOut, d)

	loaded, err := engine.Load(statePath)
	if err != nil {
		t.Fatalf("loading state: %v", err)
	}
	for _, name := range []string{"TestAccA", "TestAccB", "TestAccC"} {
		r := findResult(loaded.Results, name)
		if r == nil {
			t.Errorf("%s missing from merged state; results=%+v", name, loaded.Results)
			continue
		}
		if r.Status != "pass" {
			t.Errorf("%s status = %q, want pass", name, r.Status)
		}
	}
}

// countResults returns how many PersistResults have the given top-level test
// name and an empty Sub (i.e. distinct top-level entries).
func countResults(results []engine.PersistResult, name string) int {
	n := 0
	for _, r := range results {
		if r.Test == name && r.Sub == "" {
			n++
		}
	}
	return n
}

// TestCLIRunInvalidRunRegexExits2 verifies that an invalid --run regex is
// rejected up front (exit 2), the runner is never invoked, and prior state is
// left untouched (not wiped).
func TestCLIRunInvalidRunRegexExits2(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")
	st := engine.State{Provider: "test", Mode: "anonymous"}
	st.Results = []engine.PersistResult{{Test: "TestAccA", Status: "pass"}}
	if err := st.Save(statePath); err != nil {
		t.Fatalf("seeding state: %v", err)
	}

	fr := &fakeRunner{}
	pf := &fakeprovider.Fake{
		Packages:       []string{"./..."},
		Pattern:        "^TestAcc",
		ModesVal:       testModes,
		RequirementsFn: allowAllRequirements,
		PreflightFn: func(_ context.Context, _ string, _ provider.TestRequirements) provider.PreflightReport {
			return provider.PreflightReport{Mode: "anonymous"}
		},
	}
	d := deps{
		provider:  pf,
		newRunner: func(_ io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccA"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"run", "--run", "[", "--repo-root", root}, &out, &errOut, d)

	if code != 2 {
		t.Errorf("expected exit 2 for invalid regex, got %d", code)
	}
	if fr.calls != 0 {
		t.Errorf("expected 0 runner calls, got %d", fr.calls)
	}
	loaded, err := engine.Load(statePath)
	if err != nil {
		t.Fatalf("loading state: %v", err)
	}
	if findResult(loaded.Results, "TestAccA") == nil {
		t.Errorf("prior state wiped by invalid-regex run; results=%+v", loaded.Results)
	}
}

// TestCLIRunSubtestRegexMergeNoDuplicate verifies that a run targeting a test
// (whose results include a subtest entry) replaces that test's prior entry
// exactly once - no stale or duplicate top-level entry - while preserving
// untouched tests. This guards the merge-by-results behavior for -run patterns.
func TestCLIRunSubtestRegexMergeNoDuplicate(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")
	st := engine.State{Provider: "test", Mode: "anonymous"}
	st.Results = []engine.PersistResult{
		{Test: "TestAccA", Status: "pass"},
		{Test: "TestAccB", Status: "fail"},
	}
	if err := st.Save(statePath); err != nil {
		t.Fatalf("seeding state: %v", err)
	}

	// The run reports a fresh top-level result for B plus a subtest entry.
	fr := &fakeRunner{result: engine.RunResult{
		Tests: []engine.TestResult{
			{Name: "TestAccB", Status: provider.StatusPass},
			{Name: "TestAccB", Sub: "Sub", Status: provider.StatusPass},
		},
	}}
	pf := &fakeprovider.Fake{
		Packages:       []string{"./..."},
		Pattern:        "^TestAcc",
		ModesVal:       testModes,
		RequirementsFn: allowAllRequirements,
		PreflightFn: func(_ context.Context, _ string, _ provider.TestRequirements) provider.PreflightReport {
			return provider.PreflightReport{Mode: "anonymous"}
		},
	}
	d := deps{
		provider:  pf,
		newRunner: func(_ io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccA", "TestAccB"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	_ = runWithDeps([]string{"run", "--run", "TestAccB", "--repo-root", root}, &out, &errOut, d)

	loaded, err := engine.Load(statePath)
	if err != nil {
		t.Fatalf("loading state: %v", err)
	}
	if got := countResults(loaded.Results, "TestAccB"); got != 1 {
		t.Errorf("expected exactly 1 top-level TestAccB entry, got %d; results=%+v", got, loaded.Results)
	}
	b := findResult(loaded.Results, "TestAccB")
	if b == nil || b.Status != "pass" {
		t.Errorf("TestAccB should be updated to pass; got %+v", b)
	}
	if a := findResult(loaded.Results, "TestAccA"); a == nil || a.Status != "pass" {
		t.Errorf("untouched TestAccA should be preserved as pass; got %+v", a)
	}
}

// TestRunForwardsAuthModeToChildEnv verifies that the selected --mode is propagated to
// the child go test process via GH_TEST_AUTH_MODE. The provider's TestMain reads that
// variable to decide which acceptance tests run; without it every credentialed mode
// silently runs as anonymous and skips its real tests (false green).
func TestRunForwardsAuthModeToChildEnv(t *testing.T) {
	root := t.TempDir()
	fr := &fakeRunner{
		result: engine.RunResult{
			Tests: []engine.TestResult{{Name: "TestAccA", Status: provider.StatusPass}},
		},
	}
	pf := &fakeprovider.Fake{
		NameVal:        "test",
		Packages:       []string{"./..."},
		Pattern:        "^TestAcc",
		ModesVal:       testModes,
		RequirementsFn: allowAllRequirements,
		PreflightFn: func(_ context.Context, m string, _ provider.TestRequirements) provider.PreflightReport {
			return provider.PreflightReport{Mode: m}
		},
	}
	d := deps{
		provider:  pf,
		newRunner: func(_ io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccA"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"run", "--mode", "organization", "--repo-root", root}, &out, &errOut, d)
	if code != 0 {
		t.Fatalf("expected exit 0, got %d (stderr: %s)", code, errOut.String())
	}
	if !extraEnvContains(fr.spec.ExtraEnv, "GH_TEST_AUTH_MODE=organization") {
		t.Errorf("child ExtraEnv missing GH_TEST_AUTH_MODE=organization, got: %v", fr.spec.ExtraEnv)
	}
}

func extraEnvContains(env []string, want string) bool {
	for _, e := range env {
		if e == want {
			return true
		}
	}
	return false
}

// ── discovery fail-open regression (Codex audit: cli.go:597, dashboard.go:261) ─

// TestDiscoverGroupsSurfacesListError proves discovery errors propagate instead
// of being swallowed into an empty group set.
func TestDiscoverGroupsSurfacesListError(t *testing.T) {
	d := deps{
		provider: &fakeprovider.Fake{NameVal: "test", Packages: []string{"./..."}, Pattern: "^TestAcc"},
		list:     failList(errors.New("build failed: ./github does not compile")),
	}
	_, _, err := discoverGroups(context.Background(), d, t.TempDir())
	if err == nil {
		t.Fatal("expected error when lister fails, got nil")
	}
	if !strings.Contains(err.Error(), "build failed") {
		t.Errorf("expected underlying list error wrapped, got: %v", err)
	}
}

// TestDiscoverGroupsReturnsGroups confirms the happy path still flattens names.
func TestDiscoverGroupsReturnsGroups(t *testing.T) {
	pf := &fakeprovider.Fake{
		NameVal:  "test",
		Packages: []string{"./..."},
		Pattern:  "^TestAcc",
		GroupFunc: func(name string) string {
			if strings.HasPrefix(name, "TestAccRepo") {
				return "repos"
			}
			return "misc"
		},
	}
	d := deps{provider: pf, list: stubList([]string{"TestAccRepoA", "TestAccOther"})}
	groups, allNames, err := discoverGroups(context.Background(), d, t.TempDir())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(groups) == 0 {
		t.Fatal("expected groups, got none")
	}
	if len(allNames) != 2 {
		t.Errorf("expected 2 flattened names, got %d: %v", len(allNames), allNames)
	}
}

// TestReportFailsWhenDiscoveryFails proves the report command exits non-zero and
// surfaces the error rather than printing an empty report on a build failure.
func TestReportFailsWhenDiscoveryFails(t *testing.T) {
	root := t.TempDir()
	d := deps{
		provider: &fakeprovider.Fake{NameVal: "test", Packages: []string{"./..."}, Pattern: "^TestAcc"},
		list:     failList(errors.New("build failed: ./github does not compile")),
		cwd:      func() (string, error) { return root, nil },
	}
	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"report", "--repo-root", root}, &out, &errOut, d)
	if code == 0 {
		t.Fatalf("expected non-zero exit when discovery fails; stdout=%q stderr=%q", out.String(), errOut.String())
	}
	if !strings.Contains(errOut.String(), "build failed") {
		t.Errorf("expected discovery error on stderr, got: %s", errOut.String())
	}
}

// TestDashboardFailsWhenDiscoveryFails proves the dashboard aborts before
// launching the TUI when discovery fails, instead of opening an empty board.
func TestDashboardFailsWhenDiscoveryFails(t *testing.T) {
	root := t.TempDir()
	d := deps{
		provider:  &fakeprovider.Fake{NameVal: "test", Packages: []string{"./..."}, Pattern: "^TestAcc"},
		list:      failList(errors.New("build failed: ./github does not compile")),
		newRunner: func(_ io.Writer) testRunner { return &fakeRunner{} },
	}
	var out, errOut bytes.Buffer
	code := runDashboard(d, root, "anonymous", "", true, false, &out, &errOut)
	if code == 0 {
		t.Fatalf("expected non-zero exit when discovery fails; stdout=%q stderr=%q", out.String(), errOut.String())
	}
	if !strings.Contains(errOut.String(), "build failed") {
		t.Errorf("expected discovery error on stderr, got: %s", errOut.String())
	}
}

// TestVersionCommand locks the `terraform-provider-tester version` output contract. The version
// string is a package var (not a const) so a release build can inject the real
// version via -ldflags "-X .../cli.version=...". This test pins the default and
// guards the output format against accidental change.
func TestVersionCommand(t *testing.T) {
	for _, arg := range []string{"version", "--version"} {
		var out, errOut bytes.Buffer
		code := runWithDeps([]string{arg}, &out, &errOut, deps{})
		if code != 0 {
			t.Fatalf("%s: expected exit 0, got %d (stderr=%q)", arg, code, errOut.String())
		}
		got := out.String()
		// The first line is the stable, machine-parseable identity line and must
		// not change shape.
		firstLine := tui.AppName + " " + version + " (terraform-provider-tester)"
		if !strings.HasPrefix(got, firstLine+"\n") {
			t.Errorf("%s: first line = %q, want prefix %q", arg, got, firstLine)
		}
		// The remaining lines carry the description, ownership, license, and a
		// pointer to the maintainers list.
		for _, want := range []string{tui.Description, tui.Vendor, tui.License, tui.MaintainersRef} {
			if !strings.Contains(got, want) {
				t.Errorf("%s: version output missing %q\n%s", arg, want, got)
			}
		}
	}
}

// TestCLIUnlockClearsStaleLock verifies `terraform-provider-tester unlock` removes a stuck lock file
// and is a no-op (exit 0) when none is present.
func TestCLIUnlockClearsStaleLock(t *testing.T) {
	root := t.TempDir()
	lockPath := filepath.Join(root, ".pulsar-state.json.lock")
	if err := os.WriteFile(lockPath, nil, 0o600); err != nil {
		t.Fatalf("seed lock: %v", err)
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"unlock", "--repo-root", root}, &out, &errOut, deps{})
	if code != 0 {
		t.Fatalf("unlock exit = %d, want 0; stderr=%q", code, errOut.String())
	}
	if !strings.Contains(out.String(), "removed lock") {
		t.Errorf("unlock stdout = %q, want it to mention 'removed lock'", out.String())
	}
	if _, err := os.Stat(lockPath); !os.IsNotExist(err) {
		t.Errorf("lock still present after unlock: %v", err)
	}

	out.Reset()
	errOut.Reset()
	code = runWithDeps([]string{"unlock", "--repo-root", root}, &out, &errOut, deps{})
	if code != 0 {
		t.Fatalf("second unlock exit = %d, want 0", code)
	}
	if !strings.Contains(out.String(), "no lock file present") {
		t.Errorf("second unlock stdout = %q, want 'no lock file present'", out.String())
	}
}

// TestCLIDiscoverBasic verifies that discover prints orgs and enterprises from the fake provider.
func TestCLIDiscoverBasic(t *testing.T) {
	pf := &fakeprovider.Fake{
		DiscoverFn: func(_ context.Context, opts provider.DiscoverOpts) (provider.DiscoveryResult, error) {
			return provider.DiscoveryResult{
				Orgs: []string{"acme", "widgets"},
				Enterprises: []provider.DiscoveredEnterprise{
					{Slug: "bigcorp", Name: "Big Corp"},
				},
			}, nil
		},
	}
	d := deps{provider: pf}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"discover"}, &out, &errOut, d)

	if code != 0 {
		t.Fatalf("expected exit 0, got %d; stderr=%s", code, errOut.String())
	}
	got := out.String()
	for _, want := range []string{"acme", "widgets", "bigcorp", "Big Corp"} {
		if !strings.Contains(got, want) {
			t.Errorf("expected %q in discover output; got:\n%s", want, got)
		}
	}
	if !strings.Contains(got, "Organizations") {
		t.Errorf("expected 'Organizations' header in output; got:\n%s", got)
	}
	if !strings.Contains(got, "Enterprises") {
		t.Errorf("expected 'Enterprises' header in output; got:\n%s", got)
	}
	if !strings.Contains(got, "GITHUB_OWNER") {
		t.Errorf("expected trailing hint with GITHUB_OWNER in output; got:\n%s", got)
	}
}

// TestCLIDiscoverWithOrg verifies that discover --org prints template repos.
func TestCLIDiscoverWithOrg(t *testing.T) {
	pf := &fakeprovider.Fake{
		DiscoverFn: func(_ context.Context, opts provider.DiscoverOpts) (provider.DiscoveryResult, error) {
			if opts.Org != "acme" {
				t.Errorf("expected org=acme, got %q", opts.Org)
			}
			return provider.DiscoveryResult{
				Orgs:          []string{"acme"},
				TemplateRepos: []string{"tmpl-base", "tmpl-extra"},
			}, nil
		},
	}
	d := deps{provider: pf}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"discover", "--org", "acme"}, &out, &errOut, d)

	if code != 0 {
		t.Fatalf("expected exit 0, got %d; stderr=%s", code, errOut.String())
	}
	got := out.String()
	for _, want := range []string{"tmpl-base", "tmpl-extra", "acme"} {
		if !strings.Contains(got, want) {
			t.Errorf("expected %q in discover --org output; got:\n%s", want, got)
		}
	}
	if !strings.Contains(got, "Template repos") {
		t.Errorf("expected 'Template repos' section in output; got:\n%s", got)
	}
}

// TestCLIDiscoverNotes verifies that Notes lines appear prefixed with "note: ".
func TestCLIDiscoverNotes(t *testing.T) {
	pf := &fakeprovider.Fake{
		DiscoverFn: func(_ context.Context, _ provider.DiscoverOpts) (provider.DiscoveryResult, error) {
			return provider.DiscoveryResult{
				Notes: []string{"orgs unverified: permission denied"},
			}, nil
		},
	}
	d := deps{provider: pf}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"discover"}, &out, &errOut, d)

	if code != 0 {
		t.Fatalf("expected exit 0 even with notes, got %d; stderr=%s", code, errOut.String())
	}
	if !strings.Contains(out.String(), "note: orgs unverified:") {
		t.Errorf("expected 'note: orgs unverified:' in output; got:\n%s", out.String())
	}
}

// ── Task 7: preserve plan eligibility through retry and resume ─────────────────

// planWithEligible builds a minimal *engine.ExecutionPlan for test state seeding.
func planWithEligible(mode string, eligible []string) *engine.ExecutionPlan {
	return &engine.ExecutionPlan{Mode: mode, Eligible: eligible}
}

// TestRetryRejectsLegacyPlanlessState verifies that retry --failed returns exit 2
// with a human-readable "state/plan-required" error when the persisted state has
// no execution plan (pre-Task-6 legacy run).
func TestRetryRejectsLegacyPlanlessState(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")
	// Seed legacy planless state (no Plan field).
	st := engine.State{Provider: "test", Mode: "anonymous"}
	st.Results = []engine.PersistResult{{Test: "TestAccA", Status: "fail"}}
	if err := st.Save(statePath); err != nil {
		t.Fatalf("seeding state: %v", err)
	}

	fr := &fakeRunner{}
	d := deps{
		provider:  &fakeprovider.Fake{Packages: []string{"./..."}, Pattern: "^TestAcc"},
		newRunner: func(_ io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccA"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"retry", "--failed", "--repo-root", root}, &out, &errOut, d)

	if code != 2 {
		t.Errorf("expected exit 2 for planless state, got %d; stderr=%s", code, errOut.String())
	}
	const wantFragment = planRequiredError
	if !strings.Contains(errOut.String(), wantFragment) {
		t.Errorf("expected %q in stderr; got: %s", wantFragment, errOut.String())
	}
	if fr.calls != 0 {
		t.Errorf("expected 0 runner calls for planless state, got %d", fr.calls)
	}
}

func TestRetryRejectsInternallyMismatchedPersistedModes(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, ".git"), []byte("gitdir: /tmp/fake\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	state := engine.State{
		Provider: "test",
		Mode:     "individual",
		Plan:     planWithEligible("organization", []string{"TestAccA"}),
		Results: []engine.PersistResult{{
			Test:    "TestAccA",
			Status:  "fail",
			Package: "./github",
		}},
	}
	if err := state.Save(filepath.Join(root, ".pulsar-state.json")); err != nil {
		t.Fatalf("seeding state: %v", err)
	}

	fr := &fakeRunner{result: engine.RunResult{}}
	d := deps{
		provider:  &fakeprovider.Fake{Packages: []string{"./..."}, Pattern: "^TestAcc", ModesVal: testModes},
		newRunner: func(_ io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccA"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"retry", "--failed"}, &out, &errOut, d)

	if code != 2 {
		t.Fatalf("exit code = %d, want 2; stderr=%s", code, errOut.String())
	}
	if fr.calls != 0 {
		t.Fatalf("runner calls = %d, want 0", fr.calls)
	}
	if !strings.Contains(errOut.String(), "state mode") || !strings.Contains(errOut.String(), "plan mode") {
		t.Fatalf("stderr = %q, want persisted mode mismatch guidance", errOut.String())
	}
}

// TestResumeRejectsLegacyPlanlessState verifies that resume returns exit 2 with a
// "state/plan-required" error when the persisted state pre-dates execution plans.
func TestResumeRejectsLegacyPlanlessState(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")
	st := engine.State{Provider: "test", Mode: "anonymous"}
	st.Results = []engine.PersistResult{{Test: "TestAccA", Status: "fail"}}
	if err := st.Save(statePath); err != nil {
		t.Fatalf("seeding state: %v", err)
	}

	fr := &fakeRunner{}
	d := deps{
		provider:  &fakeprovider.Fake{Packages: []string{"./..."}, Pattern: "^TestAcc"},
		newRunner: func(_ io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccA"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"resume", "--repo-root", root}, &out, &errOut, d)

	if code != 2 {
		t.Errorf("expected exit 2 for planless state, got %d; stderr=%s", code, errOut.String())
	}
	const wantFragment = planRequiredError
	if !strings.Contains(errOut.String(), wantFragment) {
		t.Errorf("expected %q in stderr; got: %s", wantFragment, errOut.String())
	}
	if fr.calls != 0 {
		t.Errorf("expected 0 runner calls for planless state, got %d", fr.calls)
	}
}

// TestRetryExcludesIneligibleFailure verifies that retry --failed returns only the
// intersection of FailedTopLevel and Plan.Eligible – a failed test excluded from
// the original plan must not be retried.
func TestRetryExcludesIneligibleFailure(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")
	// TestAccA is eligible and failed; TestAccB failed but is NOT in Plan.Eligible.
	st := engine.State{Provider: "test", Mode: "anonymous"}
	st.Plan = planWithEligible("anonymous", []string{"TestAccA"})
	st.Results = []engine.PersistResult{
		{Test: "TestAccA", Status: "fail"},
		{Test: "TestAccB", Status: "fail"},
	}
	if err := st.Save(statePath); err != nil {
		t.Fatalf("seeding state: %v", err)
	}

	fr := &fakeRunner{result: engine.RunResult{}}
	d := deps{
		provider:  &fakeprovider.Fake{Packages: []string{"./..."}, Pattern: "^TestAcc"},
		newRunner: func(_ io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccA", "TestAccB"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	_ = runWithDeps([]string{"retry", "--failed", "--repo-root", root}, &out, &errOut, d)

	wantPattern := engine.RunPattern([]string{"TestAccA"})
	if fr.calls != 1 {
		t.Errorf("expected 1 runner call, got %d", fr.calls)
	}
	if fr.spec.Pattern != wantPattern {
		t.Errorf("pattern = %q, want %q (only eligible failures)", fr.spec.Pattern, wantPattern)
	}
}

// TestResumeDoesNotReadmitExcludedTest verifies that a failed test that was excluded
// from the original execution plan (not in Plan.Eligible) is never re-admitted by
// resume, even though it is present in FailedTopLevel.
func TestResumeDoesNotReadmitExcludedTest(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")
	// TestAccA is eligible; TestAccB is a failure but NOT in Plan.Eligible.
	st := engine.State{Provider: "test", Mode: "anonymous"}
	st.Plan = planWithEligible("anonymous", []string{"TestAccA"})
	st.Results = []engine.PersistResult{
		{Test: "TestAccA", Status: "fail"},
		{Test: "TestAccB", Status: "fail"},
	}
	if err := st.Save(statePath); err != nil {
		t.Fatalf("seeding state: %v", err)
	}

	fr := &fakeRunner{result: engine.RunResult{}}
	d := deps{
		provider:  &fakeprovider.Fake{Packages: []string{"./..."}, Pattern: "^TestAcc"},
		newRunner: func(_ io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccA", "TestAccB"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	_ = runWithDeps([]string{"resume", "--repo-root", root}, &out, &errOut, d)

	wantPattern := engine.RunPattern([]string{"TestAccA"})
	if fr.calls != 1 {
		t.Errorf("expected 1 runner call, got %d; stderr=%s", fr.calls, errOut.String())
	}
	if fr.spec.Pattern != wantPattern {
		t.Errorf("pattern = %q, want %q (excluded failure TestAccB must not re-enter)", fr.spec.Pattern, wantPattern)
	}
}

// TestResumeUsesEligibleNotRunSet verifies that resume computes the not-run set from
// State.Plan.Eligible, not from a fresh unfiltered test discovery. A test that is
// discovered by the lister but absent from Plan.Eligible must not enter the selection.
func TestResumeUsesEligibleNotRunSet(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")
	// Plan.Eligible is {"TestAccA", "TestAccB"}; TestAccA already ran and passed.
	// Fresh discovery would return TestAccC too, but it is not in Plan.Eligible.
	st := engine.State{Provider: "test", Mode: "anonymous"}
	st.Plan = planWithEligible("anonymous", []string{"TestAccA", "TestAccB"})
	st.Results = []engine.PersistResult{
		{Test: "TestAccA", Status: "pass"},
	}
	if err := st.Save(statePath); err != nil {
		t.Fatalf("seeding state: %v", err)
	}

	fr := &fakeRunner{result: engine.RunResult{}}
	d := deps{
		provider:  &fakeprovider.Fake{Packages: []string{"./..."}, Pattern: "^TestAcc"},
		newRunner: func(_ io.Writer) testRunner { return fr },
		// Discovery returns TestAccC, but it must be ignored for selection.
		list: stubList([]string{"TestAccA", "TestAccB", "TestAccC"}),
		cwd:  func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	_ = runWithDeps([]string{"resume", "--repo-root", root}, &out, &errOut, d)

	wantPattern := engine.RunPattern([]string{"TestAccB"})
	if fr.calls != 1 {
		t.Errorf("expected 1 runner call, got %d; stderr=%s", fr.calls, errOut.String())
	}
	if fr.spec.Pattern != wantPattern {
		t.Errorf("pattern = %q, want %q (TestAccC outside Plan.Eligible must be excluded)", fr.spec.Pattern, wantPattern)
	}
}

// TestRetryResumeRejectModeMismatch verifies that passing an explicit --mode that
// disagrees with State.Plan.Mode is rejected before runner creation (exit 2).
func TestRetryResumeRejectModeMismatch(t *testing.T) {
	for _, subcmd := range []string{"retry --failed", "resume"} {
		t.Run(subcmd, func(t *testing.T) {
			root := t.TempDir()
			statePath := filepath.Join(root, ".pulsar-state.json")
			st := engine.State{Provider: "test", Mode: "anonymous"}
			st.Plan = planWithEligible("anonymous", []string{"TestAccA"})
			st.Results = []engine.PersistResult{{Test: "TestAccA", Status: "fail"}}
			if err := st.Save(statePath); err != nil {
				t.Fatalf("seeding state: %v", err)
			}

			fr := &fakeRunner{}
			d := deps{
				provider:  &fakeprovider.Fake{Packages: []string{"./..."}, Pattern: "^TestAcc", ModesVal: testModes},
				newRunner: func(_ io.Writer) testRunner { return fr },
				list:      stubList([]string{"TestAccA"}),
				cwd:       func() (string, error) { return root, nil },
			}

			args := strings.Fields(subcmd)
			args = append(args, "--mode", "organization", "--repo-root", root)
			var out, errOut bytes.Buffer
			code := runWithDeps(args, &out, &errOut, d)

			if code != 2 {
				t.Errorf("expected exit 2 for mode mismatch, got %d; stderr=%s", code, errOut.String())
			}
			if fr.calls != 0 {
				t.Errorf("expected 0 runner calls, got %d", fr.calls)
			}
		})
	}
}

// TestRetryResumeDefaultToPersistedMode verifies that omitting --mode causes retry
// and resume to default to State.Plan.Mode (not the environment default).
func TestRetryResumeDefaultToPersistedMode(t *testing.T) {
	for _, subcmd := range []string{"retry --failed", "resume"} {
		t.Run(subcmd, func(t *testing.T) {
			root := t.TempDir()
			statePath := filepath.Join(root, ".pulsar-state.json")
			st := engine.State{Provider: "test", Mode: "organization"}
			st.Plan = planWithEligible("organization", []string{"TestAccA"})
			st.Results = []engine.PersistResult{{Test: "TestAccA", Status: "fail"}}
			if err := st.Save(statePath); err != nil {
				t.Fatalf("seeding state: %v", err)
			}

			fr := &fakeRunner{result: engine.RunResult{}}
			d := deps{
				provider:  &fakeprovider.Fake{Packages: []string{"./..."}, Pattern: "^TestAcc", ModesVal: testModes},
				newRunner: func(_ io.Writer) testRunner { return fr },
				list:      stubList([]string{"TestAccA"}),
				cwd:       func() (string, error) { return root, nil },
				// Deliberately no GH_TEST_AUTH_MODE in env – must use plan's mode.
				getenv: getenvFromMap(map[string]string{}),
			}

			args := strings.Fields(subcmd)
			args = append(args, "--repo-root", root)
			var out, errOut bytes.Buffer
			_ = runWithDeps(args, &out, &errOut, d)

			if fr.calls != 1 {
				t.Errorf("expected 1 runner call, got %d; stderr=%s", fr.calls, errOut.String())
			}
			if !extraEnvContains(fr.spec.ExtraEnv, "GH_TEST_AUTH_MODE=organization") {
				t.Errorf("child ExtraEnv = %v, want GH_TEST_AUTH_MODE=organization (persisted mode)", fr.spec.ExtraEnv)
			}
		})
	}
}

// TestRetryResumeTUIRejectModeMismatchBeforeDashboardPlanning verifies that an
// explicit --mode mismatch is rejected before the dashboard ever plans or
// launches. The retry path regressed here because its TUI branch ran before the
// persisted-state mode check.
func TestRetryResumeTUIRejectModeMismatchBeforeDashboardPlanning(t *testing.T) {
	for _, subcmd := range []string{"retry --failed", "resume"} {
		t.Run(subcmd, func(t *testing.T) {
			root := t.TempDir()
			statePath := filepath.Join(root, ".pulsar-state.json")
			st := engine.State{Provider: "test", Mode: "anonymous"}
			st.Plan = planWithEligible("anonymous", []string{"TestAccA"})
			st.Results = []engine.PersistResult{{Test: "TestAccA", Status: "fail"}}
			if err := st.Save(statePath); err != nil {
				t.Fatalf("seeding state: %v", err)
			}

			listCalls := 0
			d := deps{
				provider: &fakeprovider.Fake{
					Packages: []string{"./..."},
					Pattern:  "^TestAcc",
					ModesVal: testModes,
				},
				list: func(_ context.Context, _ string, _ []string, _ string) ([]string, error) {
					listCalls++
					return nil, errors.New("dashboard planning must not run after a mode mismatch")
				},
				cwd:   func() (string, error) { return root, nil },
				isTTY: func() bool { return true },
			}

			args := append(strings.Fields(subcmd), "--mode", "organization", "--tui", "--repo-root", root)
			var out, errOut bytes.Buffer
			code := runWithDeps(args, &out, &errOut, d)

			if code != 2 {
				t.Fatalf("exit code = %d, want 2; stderr=%s", code, errOut.String())
			}
			if listCalls != 0 {
				t.Fatalf("dashboard planning ran %d times before the mismatch was rejected", listCalls)
			}
			if !strings.Contains(errOut.String(), "mode mismatch") {
				t.Fatalf("stderr = %q, want mode mismatch guidance", errOut.String())
			}
		})
	}
}

// TestRetryResumeTUIUsesPersistedPlanMode verifies that omitting --mode on the
// dashboard path still plans in the persisted plan mode rather than a shell
// value, an env-file GH_TEST_AUTH_MODE, or the anonymous fallback.
func TestRetryResumeTUIUsesPersistedPlanMode(t *testing.T) {
	type sourceCase struct {
		name         string
		baseEnv      map[string]string
		envFileValue string
		wrongMode    string
	}

	sources := []sourceCase{
		{name: "shell mode", baseEnv: map[string]string{"GH_TEST_AUTH_MODE": "shell-mode"}, wrongMode: "shell-mode"},
		{name: "env file mode", envFileValue: "env-file-mode", wrongMode: "env-file-mode"},
		{name: "anonymous fallback", wrongMode: "anonymous"},
	}

	for _, subcmd := range []string{"retry --failed", "resume"} {
		for _, source := range sources {
			t.Run(subcmd+" "+source.name, func(t *testing.T) {
				root := t.TempDir()
				statePath := filepath.Join(root, ".pulsar-state.json")
				st := engine.State{Provider: "test", Mode: "persisted-mode"}
				st.Plan = planWithEligible("persisted-mode", []string{"TestAccA"})
				st.Results = []engine.PersistResult{{Test: "TestAccA", Status: "fail"}}
				if err := st.Save(statePath); err != nil {
					t.Fatalf("seeding state: %v", err)
				}

				getenv, setenv, _ := mapEnv(source.baseEnv)
				args := append(strings.Fields(subcmd), "--tui", "--repo-root", root)
				if source.envFileValue != "" {
					envFile := filepath.Join(root, "retry-resume.env")
					if err := os.WriteFile(envFile, []byte("GH_TEST_AUTH_MODE="+source.envFileValue+"\n"), 0o600); err != nil {
						t.Fatalf("writing env file: %v", err)
					}
					args = append(args, "--env-file", envFile)
				}

				d := deps{
					provider: &fakeprovider.Fake{
						Packages: []string{"./..."},
						Pattern:  "^TestAcc",
						ModesVal: []provider.Mode{{Name: "supported-mode"}},
					},
					list:   stubList([]string{"TestAccA"}),
					cwd:    func() (string, error) { return root, nil },
					getenv: getenv,
					setenv: setenv,
					isTTY:  func() bool { return true },
				}

				var out, errOut bytes.Buffer
				code := runWithDeps(args, &out, &errOut, d)
				if code != 1 {
					t.Fatalf("exit code = %d, want 1 from unsupported persisted mode; stderr=%s", code, errOut.String())
				}
				if !strings.Contains(errOut.String(), `mode "persisted-mode" is not supported by this provider`) {
					t.Fatalf("stderr = %q, want persisted plan mode in dashboard planning failure", errOut.String())
				}
				if strings.Contains(errOut.String(), `mode "`+source.wrongMode+`" is not supported by this provider`) {
					t.Fatalf("stderr = %q, must not plan in %q when persisted mode is available", errOut.String(), source.wrongMode)
				}
			})
		}
	}
}

// TestRetryExplicitModeNilPlanE2EBypass verifies that when runRetryContext is
// called with a non-nil selected list (the internal test-only path that
// bypasses the normal retrySelection/planless-state rejection) and the
// persisted plan is nil (e.g. due to a nonfatal state-save failure), the
// explicit --mode flag value is still used by the runner — not discarded in
// favour of the environment/default mode.
func TestRetryExplicitModeNilPlanE2EBypass(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")
	// Seed planless state (simulates nonfatal save failure after initial run).
	st := engine.State{Provider: "test", Mode: "organization"}
	st.Results = []engine.PersistResult{{Test: "TestAccA", Status: "fail"}}
	if err := st.Save(statePath); err != nil {
		t.Fatalf("seeding state: %v", err)
	}

	fr := &fakeRunner{result: engine.RunResult{}}
	d := deps{
		provider:  &fakeprovider.Fake{Packages: []string{"./..."}, Pattern: "^TestAcc", ModesVal: testModes},
		newRunner: func(_ io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccA"}),
		cwd:       func() (string, error) { return root, nil },
		// No GH_TEST_AUTH_MODE in env – explicit flag must win.
		getenv: getenvFromMap(map[string]string{}),
	}

	// Use the internal selected-list bypass (selected non-nil); production paths
	// do not reach retry with a nil persisted plan.
	var out, errOut bytes.Buffer
	code := runRetryContext(context.Background(),
		[]string{"--failed", "--mode", "organization", "--repo-root", root},
		&out, &errOut, d, []string{"TestAccA"})

	if code != 0 {
		t.Errorf("expected exit 0, got %d; stderr=%s", code, errOut.String())
	}
	if fr.calls != 1 {
		t.Errorf("expected 1 runner call, got %d", fr.calls)
	}
	if !extraEnvContains(fr.spec.ExtraEnv, "GH_TEST_AUTH_MODE=organization") {
		t.Errorf("child ExtraEnv = %v, want GH_TEST_AUTH_MODE=organization (explicit mode must be preserved with nil plan)", fr.spec.ExtraEnv)
	}
}

// TestRetryIntersectionPreservesOrder verifies that retrySelection returns only
// the intersection of FailedTopLevel and Plan.Eligible, preserving the order
// defined by Plan.Eligible (not the order of failures in the result set).
func TestRetryIntersectionPreservesOrder(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")
	// Plan.Eligible order: C, A, B. Failures in result order: A, B, C.
	// Expected retry order follows Plan.Eligible: C, A, B.
	st := engine.State{Provider: "test", Mode: "anonymous"}
	st.Plan = planWithEligible("anonymous", []string{"TestAccC", "TestAccA", "TestAccB"})
	st.Results = []engine.PersistResult{
		{Test: "TestAccA", Status: "fail"},
		{Test: "TestAccB", Status: "fail"},
		{Test: "TestAccC", Status: "fail"},
	}
	if err := st.Save(statePath); err != nil {
		t.Fatalf("seeding state: %v", err)
	}

	fr := &fakeRunner{result: engine.RunResult{}}
	d := deps{
		provider:  &fakeprovider.Fake{Packages: []string{"./..."}, Pattern: "^TestAcc", ModesVal: testModes},
		newRunner: func(_ io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccA", "TestAccB", "TestAccC"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	_ = runWithDeps([]string{"retry", "--failed", "--repo-root", root}, &out, &errOut, d)

	wantPattern := engine.RunPattern([]string{"TestAccC", "TestAccA", "TestAccB"})
	if fr.calls != 1 {
		t.Errorf("expected 1 runner call, got %d; stderr=%s", fr.calls, errOut.String())
	}
	if fr.spec.Pattern != wantPattern {
		t.Errorf("pattern = %q, want %q (intersection must follow Plan.Eligible order)", fr.spec.Pattern, wantPattern)
	}
}

// TestResumeConcatDeduplicatesAndPreservesOrder verifies that resumeSelection
// produces: eligible-failed tests first (in Plan.Eligible order), then
// eligible-not-run tests (also in Plan.Eligible order), with no duplicates.
func TestResumeConcatDeduplicatesAndPreservesOrder(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")
	// Plan.Eligible order: A, B, C, D.
	// A=fail, B=pass, C=not-run, D=fail.
	// Expected resume order: eligible-failed=[A,D] ++ eligible-not-run=[C]. B excluded (passed).
	st := engine.State{Provider: "test", Mode: "anonymous"}
	st.Plan = planWithEligible("anonymous", []string{"TestAccA", "TestAccB", "TestAccC", "TestAccD"})
	st.Results = []engine.PersistResult{
		{Test: "TestAccA", Status: "fail"},
		{Test: "TestAccB", Status: "pass"},
		{Test: "TestAccD", Status: "fail"},
	}
	if err := st.Save(statePath); err != nil {
		t.Fatalf("seeding state: %v", err)
	}

	fr := &fakeRunner{result: engine.RunResult{}}
	d := deps{
		provider:  &fakeprovider.Fake{Packages: []string{"./..."}, Pattern: "^TestAcc", ModesVal: testModes},
		newRunner: func(_ io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccA", "TestAccB", "TestAccC", "TestAccD"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	_ = runWithDeps([]string{"resume", "--repo-root", root}, &out, &errOut, d)

	wantPattern := engine.RunPattern([]string{"TestAccA", "TestAccD", "TestAccC"})
	if fr.calls != 1 {
		t.Errorf("expected 1 runner call, got %d; stderr=%s", fr.calls, errOut.String())
	}
	if fr.spec.Pattern != wantPattern {
		t.Errorf("pattern = %q, want %q (failed first in eligible order, then not-run)", fr.spec.Pattern, wantPattern)
	}
}

// TestCLIDiscoverUnknownFlag verifies that an unknown flag causes exit 2.
func TestCLIDiscoverUnknownFlag(t *testing.T) {
	pf := &fakeprovider.Fake{}
	d := deps{provider: pf}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"discover", "--no-such-flag"}, &out, &errOut, d)

	if code != 2 {
		t.Errorf("expected exit 2 for unknown flag, got %d", code)
	}
}

func TestInterruptedRunSavesPartialResults(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")
	baseline := []provider.Resource{{
		Kind: "repository", Name: "tf-acc-test-existing", URL: "https://example.test/existing",
	}}
	plan := planWithEligible("organization", []string{"TestAccCompleted", "TestAccPending"})
	prev := engine.State{
		Provider: "test",
		Mode:     "organization",
		Plan:     plan,
		Orphans: &engine.OrphanAccounting{
			Mode:             "organization",
			Baseline:         baseline,
			BaselineCaptured: true,
			CleanupStatus:    engine.CleanupBaselineOnly,
		},
	}
	runner := &fakeRunner{
		result: engine.RunResult{
			BuildFailed:  true,
			BuildOutput:  []string{"build interrupted\n"},
			PreRunFailed: true,
			PreRunOutput: []string{"pre-run interrupted\n"},
			Tests: []engine.TestResult{{
				Package: "./github", Name: "TestAccCompleted", Status: provider.StatusPass,
			}},
		},
		err: context.Canceled,
	}

	var out, errOut bytes.Buffer
	code := runWithSpec(context.Background(), engine.RunSpec{Dir: root}, statePath, prev,
		&fakeprovider.Fake{NameVal: "test"}, "organization", runner, &out, &errOut,
		runOptions{Command: "run"})

	if code != 1 {
		t.Fatalf("exit code = %d, want 1; stderr=%s", code, errOut.String())
	}
	state, err := engine.Load(statePath)
	if err != nil {
		t.Fatal(err)
	}
	if len(state.Results) != 1 || state.Results[0].Test != "TestAccCompleted" || state.Results[0].Status != "pass" {
		t.Fatalf("results = %+v, want streamed passing result", state.Results)
	}
	if !state.BuildFailed || !state.PreRunFailed {
		t.Fatalf("suite failure flags = build:%v pre-run:%v, want both true", state.BuildFailed, state.PreRunFailed)
	}
	for label, path := range map[string]string{"build": state.BuildLog, "pre-run": state.PreRunLog} {
		if path == "" {
			t.Fatalf("%s log path is empty", label)
		}
		if _, err := os.Stat(path); err != nil {
			t.Fatalf("%s log %q: %v", label, path, err)
		}
		if !reflect.DeepEqual(state.Plan, plan) {
			t.Fatalf("plan = %+v, want %+v", state.Plan, plan)
		}
		if state.Orphans == nil || !reflect.DeepEqual(state.Orphans.Baseline, baseline) {
			t.Fatalf("orphans = %+v, want original baseline %+v", state.Orphans, baseline)
		}
	}
}

// unknownTestProvider returns a provider whose modes cover anonymous plus the
// authenticated set and whose RequirementsFor reports every name in unknown
// as unclassified, returning the conservative authenticated fallback the real
// provider catalog uses (see provider/github.conservativeRequirements).
func unknownTestProvider(unknown map[string]bool) *fakeprovider.Fake {
	return &fakeprovider.Fake{
		NameVal:  "test",
		Packages: []string{"./github/..."},
		Pattern:  "^TestAcc",
		ModesVal: []provider.Mode{
			{Name: "anonymous"},
			{Name: "individual"},
			{Name: "organization"},
		},
		RequirementsFn: func(name string) (provider.TestRequirements, bool) {
			if unknown[name] {
				return provider.TestRequirements{
					Modes:       []string{"individual", "organization", "team", "enterprise"},
					Scopes:      []string{"repo"},
					SideEffects: []string{"unknown"},
				}, false
			}
			return provider.TestRequirements{}, true
		},
		PreflightFn: func(_ context.Context, mode string, _ provider.TestRequirements) provider.PreflightReport {
			return provider.PreflightReport{Mode: mode}
		},
	}
}

// TestRunAllowUnclassifiedFlagGatesUnknownTests is the fail-closed regression
// for `run`, which hardcoded AllowUnclassified: true and so silently ran
// unknown tests with over-broad conservative requirements. Without the new
// --allow-unclassified flag the run must fail during planning, before
// preflight and before the runner; with it the unknown test runs.
func TestRunAllowUnclassifiedFlagGatesUnknownTests(t *testing.T) {
	root := t.TempDir()
	pf := unknownTestProvider(map[string]bool{"TestAccMystery": true})
	preflightCalls := 0
	pf.PreflightFn = func(_ context.Context, mode string, _ provider.TestRequirements) provider.PreflightReport {
		preflightCalls++
		return provider.PreflightReport{Mode: mode}
	}
	fr := &fakeRunner{result: engine.RunResult{Tests: []engine.TestResult{
		{Package: "./github", Name: "TestAccMystery", Status: provider.StatusPass},
	}}}
	d := deps{
		provider:  pf,
		newRunner: func(io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccMystery"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"run", "--mode", "organization", "--no-tui", "--repo-root", root}, &out, &errOut, d)
	if code != 1 {
		t.Fatalf("exit code without --allow-unclassified = %d, want 1; stderr=%s", code, errOut.String())
	}
	if preflightCalls != 0 || fr.calls != 0 {
		t.Fatalf("planning must fail before preflight and the runner: preflight=%d runner=%d", preflightCalls, fr.calls)
	}
	if !strings.Contains(errOut.String(), "has no requirements classification") {
		t.Fatalf("stderr = %q, want the fail-closed unclassified planning error", errOut.String())
	}

	out.Reset()
	errOut.Reset()
	code = runWithDeps([]string{"run", "--mode", "organization", "--allow-unclassified", "--no-tui", "--repo-root", root}, &out, &errOut, d)
	if code != 0 {
		t.Fatalf("exit code with --allow-unclassified = %d, want 0; stderr=%s", code, errOut.String())
	}
	if fr.calls != 1 {
		t.Fatalf("runner calls with --allow-unclassified = %d, want 1", fr.calls)
	}
	if fr.spec.Pattern != engine.RunPattern([]string{"TestAccMystery"}) {
		t.Fatalf("run pattern = %q, want the allowed unknown test", fr.spec.Pattern)
	}
}

// TestRunAllowUnclassifiedStillRespectsModeCompatibility verifies that an
// allowed unknown test is not blindly run: its conservative requirements are
// authenticated-only, so a default anonymous selection excludes it and never
// starts the runner, while an explicit --run naming it fails the plan closed.
func TestRunAllowUnclassifiedStillRespectsModeCompatibility(t *testing.T) {
	root := t.TempDir()
	pf := unknownTestProvider(map[string]bool{"TestAccMystery": true})
	fr := &fakeRunner{}
	d := deps{
		provider:  pf,
		newRunner: func(io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccMystery"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{
		"run", "--mode", "anonymous", "--allow-unclassified", "--no-tui", "--repo-root", root,
	}, &out, &errOut, d)
	if code != 0 {
		t.Fatalf("default anonymous exit code = %d, want 0 (nothing to run); stderr=%s", code, errOut.String())
	}
	if fr.calls != 0 {
		t.Fatalf("runner calls = %d, want 0: the unknown test is anonymous-incompatible", fr.calls)
	}
	if !strings.Contains(out.String(), "nothing to run") {
		t.Fatalf("stdout = %q, want the empty-selection message", out.String())
	}

	out.Reset()
	errOut.Reset()
	code = runWithDeps([]string{
		"run", "--mode", "anonymous", "--allow-unclassified", "--run", "^TestAccMystery$", "--no-tui", "--repo-root", root,
	}, &out, &errOut, d)
	if code != 1 {
		t.Fatalf("explicit anonymous exit code = %d, want 1; stderr=%s", code, errOut.String())
	}
	if fr.calls != 0 {
		t.Fatalf("runner calls = %d, want 0 after a failed plan", fr.calls)
	}
	if !strings.Contains(errOut.String(), `does not support mode "anonymous"`) {
		t.Fatalf("stderr = %q, want the mode-incompatible planning error", errOut.String())
	}
}

// TestRunAllowUnclassifiedRunsUnknownInCompatibleMode proves the allowed
// unknown test really executes under a compatible authenticated mode, that
// its conservative requirements reach preflight, and that the plan event
// reports it as both eligible and unclassified.
func TestRunAllowUnclassifiedRunsUnknownInCompatibleMode(t *testing.T) {
	root := t.TempDir()
	pf := unknownTestProvider(map[string]bool{"TestAccMystery": true})
	var gotReqs provider.TestRequirements
	pf.PreflightFn = func(_ context.Context, mode string, reqs provider.TestRequirements) provider.PreflightReport {
		gotReqs = reqs
		return provider.PreflightReport{Mode: mode}
	}
	fr := &fakeRunner{result: engine.RunResult{Tests: []engine.TestResult{
		{Package: "./github", Name: "TestAccMystery", Status: provider.StatusPass},
	}}}
	d := deps{
		provider:  pf,
		newRunner: func(io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccMystery"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{
		"run", "--mode", "individual", "--allow-unclassified", "--json", "--repo-root", root,
	}, &out, &errOut, d)
	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s", code, errOut.String())
	}
	if fr.calls != 1 {
		t.Fatalf("runner calls = %d, want 1", fr.calls)
	}
	if !reflect.DeepEqual(gotReqs.Scopes, []string{"repo"}) {
		t.Errorf("preflight scopes = %v, want the conservative [repo]", gotReqs.Scopes)
	}
	if !reflect.DeepEqual(gotReqs.SideEffects, []string{"unknown"}) {
		t.Errorf("preflight side effects = %v, want [unknown]", gotReqs.SideEffects)
	}
	events := decodeNDJSON(t, out.String())
	if events[0]["type"] != "plan" || events[0]["eligible"] != float64(1) || events[0]["unclassified"] != float64(1) {
		t.Fatalf("plan event = %v, want 1 eligible and 1 unclassified", events[0])
	}
}

// TestRunJSONPlanFailureStillEmitsFinalSummary is the regression for
// `run --json` returning an empty stdout stream when planning fails. An
// automation client that reads the NDJSON stream must still receive a
// terminal summary object naming the command, mode, and exit code instead of
// nothing at all. No plan event is fabricated: the plan was never built.
func TestRunJSONPlanFailureStillEmitsFinalSummary(t *testing.T) {
	root := t.TempDir()
	pf := unknownTestProvider(map[string]bool{"TestAccMystery": true})
	preflightCalls := 0
	pf.PreflightFn = func(_ context.Context, mode string, _ provider.TestRequirements) provider.PreflightReport {
		preflightCalls++
		return provider.PreflightReport{Mode: mode}
	}
	fr := &fakeRunner{}
	d := deps{
		provider:  pf,
		newRunner: func(io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccMystery"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"run", "--mode", "organization", "--json", "--repo-root", root}, &out, &errOut, d)

	if code != 1 {
		t.Fatalf("exit code = %d, want 1; stderr=%s", code, errOut.String())
	}
	if preflightCalls != 0 || fr.calls != 0 {
		t.Fatalf("planning must fail before preflight and the runner: preflight=%d runner=%d", preflightCalls, fr.calls)
	}
	events := decodeNDJSON(t, out.String())
	if len(events) != 1 {
		t.Fatalf("events = %d, want only the terminal summary: %v", len(events), events)
	}
	summary := events[0]
	if summary["type"] != "summary" || summary["command"] != "run" {
		t.Fatalf("terminal event = %v, want the run summary", summary)
	}
	if summary["mode"] != "organization" || summary["exit_code"] != float64(1) {
		t.Fatalf("bad terminal summary: %v", summary)
	}
}

// TestPreflightJSONPlanFailureStillEmitsFinalSummary is the same regression
// for `preflight --json`: a planning failure must still terminate the NDJSON
// stream with a preflight summary object rather than an empty stream.
func TestPreflightJSONPlanFailureStillEmitsFinalSummary(t *testing.T) {
	root := t.TempDir()
	pf := unknownTestProvider(map[string]bool{"TestAccMystery": true})
	preflightCalls := 0
	pf.PreflightFn = func(_ context.Context, mode string, _ provider.TestRequirements) provider.PreflightReport {
		preflightCalls++
		return provider.PreflightReport{Mode: mode}
	}
	d := deps{
		provider: pf,
		list:     stubList([]string{"TestAccMystery"}),
		cwd:      func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"preflight", "--mode", "individual", "--json", "--repo-root", root}, &out, &errOut, d)

	if code != 1 {
		t.Fatalf("exit code = %d, want 1; stderr=%s", code, errOut.String())
	}
	if preflightCalls != 0 {
		t.Fatalf("preflight calls = %d, want 0 after a failed plan", preflightCalls)
	}
	events := decodeNDJSON(t, out.String())
	if len(events) != 1 {
		t.Fatalf("events = %d, want only the terminal summary: %v", len(events), events)
	}
	summary := events[0]
	if summary["type"] != "summary" || summary["command"] != "preflight" {
		t.Fatalf("terminal event = %v, want the preflight summary", summary)
	}
	if summary["mode"] != "individual" || summary["ok"] != false || summary["exit_code"] != float64(1) {
		t.Fatalf("bad terminal summary: %v", summary)
	}
}

// TestPreflightJSONUsageErrorWritesNoJSON pins the boundary of the new
// planning-failure summary: a terminal summary is written for a RUNTIME
// planning failure (exit 1), not for a usage error (exit 2). An invalid --run
// regex is a usage error, so `preflight --json` must report it on stderr only
// and leave stdout empty - the same contract `run --json` already holds for
// the identical input (see TestRunJSONSummaryExitCodes).
func TestPreflightJSONUsageErrorWritesNoJSON(t *testing.T) {
	root := t.TempDir()
	pf := &fakeprovider.Fake{
		NameVal:        "test",
		Packages:       []string{"./github/..."},
		Pattern:        "^TestAcc",
		ModesVal:       testModes,
		RequirementsFn: allowAllRequirements,
		PreflightFn: func(_ context.Context, mode string, _ provider.TestRequirements) provider.PreflightReport {
			return provider.PreflightReport{Mode: mode}
		},
	}
	d := deps{
		provider: pf,
		list:     stubList([]string{"TestAccThing"}),
		cwd:      func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"preflight", "--run", "([", "--json", "--repo-root", root}, &out, &errOut, d)

	if code != 2 {
		t.Fatalf("exit code = %d, want 2 for an invalid --run regex; stderr=%s", code, errOut.String())
	}
	if strings.TrimSpace(out.String()) != "" {
		t.Fatalf("usage error wrote JSON stdout, want empty stdout: %s", out.String())
	}
	if errOut.String() == "" {
		t.Fatal("expected the regex error on stderr")
	}
}

// TestResumeFailsLoudlyOnGroupDiscoveryError is the regression for resume
// silently ignoring the error from its grouping discovery pass. A provider
// package that fails to build makes `go test -list` yield zero tests, and
// resume then ran with empty grouping metadata: every result reported as
// group "misc" and every group summary missing. retry already routes the
// same lookup through discoverGroups and fails loudly; resume must match.
func TestResumeFailsLoudlyOnGroupDiscoveryError(t *testing.T) {
	for _, format := range []string{"text", "json"} {
		t.Run(format, func(t *testing.T) {
			root := t.TempDir()
			statePath := filepath.Join(root, ".pulsar-state.json")
			st := engine.State{Provider: "test", Mode: "anonymous"}
			st.Plan = planWithEligible("anonymous", []string{"TestAccA"})
			st.Results = []engine.PersistResult{{Test: "TestAccA", Status: "fail"}}
			if err := st.Save(statePath); err != nil {
				t.Fatalf("seeding state: %v", err)
			}

			fr := &fakeRunner{}
			d := deps{
				provider:  &fakeprovider.Fake{Packages: []string{"./..."}, Pattern: "^TestAcc", ModesVal: testModes},
				newRunner: func(_ io.Writer) testRunner { return fr },
				list:      failList(errors.New("build failed: ./github does not compile")),
				cwd:       func() (string, error) { return root, nil },
			}

			args := []string{"resume", "--repo-root", root}
			if format == "json" {
				args = append(args, "--json")
			}
			var out, errOut bytes.Buffer
			code := runWithDeps(args, &out, &errOut, d)

			if code != 1 {
				t.Fatalf("exit code = %d, want 1 when grouping discovery fails; stdout=%q stderr=%q", code, out.String(), errOut.String())
			}
			if fr.calls != 0 {
				t.Fatalf("runner calls = %d, want 0: never run with empty grouping metadata", fr.calls)
			}
			if !strings.Contains(errOut.String(), "build failed") {
				t.Fatalf("stderr = %q, want the underlying discovery error", errOut.String())
			}
		})
	}
}
