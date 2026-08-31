package cli

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/github/terraform-provider-tester/engine"
	"github.com/github/terraform-provider-tester/fakeprovider"
	"github.com/github/terraform-provider-tester/provider"
)

type noisyRunner struct {
	w      io.Writer
	calls  int
	spec   engine.RunSpec
	result engine.RunResult
	err    error
}

func (n *noisyRunner) Run(_ context.Context, spec engine.RunSpec, sink func(engine.TestResult)) (engine.RunResult, error) {
	n.calls++
	n.spec = spec
	if n.w != nil {
		fmt.Fprintln(n.w, "human go test command")
	}
	for _, tr := range n.result.Tests {
		if sink != nil {
			sink(tr)
		}
	}
	return n.result, n.err
}

func jsonTestProvider() *fakeprovider.Fake {
	return &fakeprovider.Fake{
		NameVal:  "test",
		Packages: []string{"./github/..."},
		Pattern:  "^TestAcc",
		ModesVal: testModes,
		GroupFunc: func(name string) string {
			switch {
			case strings.Contains(name, "Issue"):
				return "issues"
			case strings.Contains(name, "Org"):
				return "organization"
			default:
				return "misc"
			}
		},
		RequirementsFn: allowAllRequirements,
		PreflightFn: func(_ context.Context, mode string, _ provider.TestRequirements) provider.PreflightReport {
			return provider.PreflightReport{Mode: mode}
		},
	}
}

func decodeNDJSON(t *testing.T, out string) []map[string]any {
	t.Helper()
	out = strings.TrimSpace(out)
	if out == "" {
		t.Fatal("expected NDJSON output, got empty stdout")
	}
	var events []map[string]any
	for i, line := range strings.Split(out, "\n") {
		var ev map[string]any
		if err := json.Unmarshal([]byte(line), &ev); err != nil {
			t.Fatalf("stdout line %d is not JSON: %q: %v", i+1, line, err)
		}
		events = append(events, ev)
	}
	return events
}

func lastEvent(t *testing.T, events []map[string]any) map[string]any {
	t.Helper()
	if len(events) == 0 {
		t.Fatal("expected at least one event")
	}
	return events[len(events)-1]
}

func TestInterruptedTriageRunSavesPartialResults(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")
	baseline := []provider.Resource{{
		Kind: "repository", Name: "tf-acc-test-existing", URL: "https://example.test/existing",
	}}
	plan := planWithEligible("organization", []string{"TestAccCompleted", "TestAccPending"})
	prev := engine.State{
		Provider: "test",
		Mode:     "organization",
		History:  map[string][]string{"TestAccCompleted": {"fail"}},
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
	code := runWithSpecTriage(context.Background(), engine.RunSpec{Dir: root}, statePath, prev,
		jsonTestProvider(), "organization", runner, &out, &errOut, runOptions{
			Format:  formatJSON,
			Command: "run",
			Triage: triageOptions{
				Enabled: true,
				Retries: 2,
			},
		})

	if code != 1 {
		t.Fatalf("exit code = %d, want 1; stderr=%s stdout=%s", code, errOut.String(), out.String())
	}
	if runner.calls != 1 {
		t.Fatalf("runner calls = %d, want 1 with no retry after cancellation", runner.calls)
	}
	state, err := engine.Load(statePath)
	if err != nil {
		t.Fatal(err)
	}
	if len(state.Results) != 1 || state.Results[0].Test != "TestAccCompleted" || state.Results[0].Status != "pass" {
		t.Fatalf("results = %+v, want retryResult.Final passing result", state.Results)
	}
	if got := state.History["TestAccCompleted"]; !reflect.DeepEqual(got, []string{"fail", "pass"}) {
		t.Fatalf("history = %v, want [fail pass]", got)
	}
	if !state.BuildFailed || !state.PreRunFailed || state.BuildLog == "" || state.PreRunLog == "" {
		t.Fatalf("suite failure state = %+v, want flags and log paths", state)
	}
	if !reflect.DeepEqual(state.Plan, plan) || state.Orphans == nil ||
		!reflect.DeepEqual(state.Orphans.Baseline, baseline) {
		t.Fatalf("plan/orphans = %+v/%+v, want persisted plan and baseline", state.Plan, state.Orphans)
	}
	events := decodeNDJSON(t, out.String())
	summary := lastEvent(t, events)
	if summary["type"] != "summary" || summary["exit_code"] != float64(1) {
		t.Fatalf("last event = %#v, want interruption summary with exit 1", summary)
	}
}

func TestRunJSONEmitsOnlyNDJSONAndFinalSummary(t *testing.T) {
	root := t.TempDir()
	runner := &noisyRunner{
		result: engine.RunResult{
			Tests: []engine.TestResult{
				{Package: "./github", Name: "TestAccOrgSettings", Status: provider.StatusPass, Elapsed: 1.23},
				{Package: "./github", Name: "TestAccIssueLabel", Status: provider.StatusFail, Elapsed: 2.5, Output: []string{"boom\n"}},
			},
		},
	}
	d := deps{
		provider: jsonTestProvider(),
		newRunner: func(w io.Writer) testRunner {
			runner.w = w
			return runner
		},
		list: stubList([]string{"TestAccOrgSettings", "TestAccIssueLabel"}),
		cwd:  func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"run", "--json", "--mode", "organization", "--repo-root", root}, &out, &errOut, d)

	if code != 1 {
		t.Fatalf("exit code = %d, want 1; stderr=%s", code, errOut.String())
	}
	events := decodeNDJSON(t, out.String())
	if got := lastEvent(t, events)["type"]; got != "summary" {
		t.Fatalf("last event type = %v, want summary; events=%v", got, events)
	}
	if len(events) < 4 {
		t.Fatalf("expected plan event plus two test events plus summary, got %d events: %v", len(events), events)
	}
	if events[0]["type"] != "plan" {
		t.Fatalf("first event type = %v, want plan", events[0]["type"])
	}
	failEvent := events[2]
	if failEvent["type"] != "test" {
		t.Fatalf("third event type = %v, want test", failEvent["type"])
	}
	if failEvent["group"] != "issues" {
		t.Fatalf("failure group = %v, want issues", failEvent["group"])
	}
	if failEvent["failure_log_path"] == nil {
		t.Fatalf("failed top-level test missing failure_log_path: %v", failEvent)
	}
	summary := lastEvent(t, events)
	if summary["command"] != "run" || summary["mode"] != "organization" {
		t.Fatalf("summary command/mode = %v/%v, want run/organization", summary["command"], summary["mode"])
	}
	if summary["exit_code"] != float64(1) {
		t.Fatalf("summary exit_code = %v, want 1", summary["exit_code"])
	}
	if !strings.Contains(fmt.Sprint(summary["go_test_command"]), "go test ./github/... -json") {
		t.Fatalf("summary missing go_test_command: %v", summary["go_test_command"])
	}
}

func TestRunJSONBypassesTUIOnTTY(t *testing.T) {
	root := t.TempDir()
	runner := &fakeRunner{
		result: engine.RunResult{
			Tests: []engine.TestResult{{Package: "./github", Name: "TestAccOrgSettings", Status: provider.StatusPass}},
		},
	}
	d := deps{
		provider:  jsonTestProvider(),
		newRunner: func(_ io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccOrgSettings"}),
		cwd:       func() (string, error) { return root, nil },
		isTTY:     func() bool { return true },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"run", "--json", "--repo-root", root}, &out, &errOut, d)

	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s", code, errOut.String())
	}
	if runner.calls != 1 {
		t.Fatalf("runner calls = %d, want 1", runner.calls)
	}
	events := decodeNDJSON(t, out.String())
	if lastEvent(t, events)["type"] != "summary" {
		t.Fatalf("last event = %v, want summary", lastEvent(t, events))
	}
}

func TestRunNoTUIForcesTextOnTTY(t *testing.T) {
	root := t.TempDir()
	runner := &fakeRunner{
		result: engine.RunResult{
			Tests: []engine.TestResult{{Package: "./github", Name: "TestAccOrgSettings", Status: provider.StatusPass, Elapsed: 1.23}},
		},
	}
	d := deps{
		provider:  jsonTestProvider(),
		newRunner: func(_ io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccOrgSettings"}),
		cwd:       func() (string, error) { return root, nil },
		isTTY:     func() bool { return true },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"run", "--no-tui", "--repo-root", root}, &out, &errOut, d)

	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s", code, errOut.String())
	}
	if runner.calls != 1 {
		t.Fatalf("runner calls = %d, want 1", runner.calls)
	}
	if !strings.Contains(out.String(), "pass organization TestAccOrgSettings 1.23s") {
		t.Fatalf("stdout missing streamed text result: %s", out.String())
	}
}

func TestRunTextStreamsTopLevelResults(t *testing.T) {
	root := t.TempDir()
	runner := &fakeRunner{
		result: engine.RunResult{
			Tests: []engine.TestResult{
				{Package: "./github", Name: "TestAccOrgSettings", Status: provider.StatusPass, Elapsed: 1.23},
				{Package: "./github", Name: "TestAccIssueLabel", Status: provider.StatusFail, Elapsed: 2.5, Output: []string{"boom\n"}},
				{Package: "./github", Name: "TestAccIssueLabel", Sub: "Sub", Status: provider.StatusFail, Elapsed: 0.1},
			},
		},
	}
	d := deps{
		provider:  jsonTestProvider(),
		newRunner: func(_ io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccOrgSettings", "TestAccIssueLabel"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"run", "--repo-root", root}, &out, &errOut, d)

	if code != 1 {
		t.Fatalf("exit code = %d, want 1; stderr=%s", code, errOut.String())
	}
	got := out.String()
	if !strings.Contains(got, "pass organization TestAccOrgSettings 1.23s") {
		t.Fatalf("stdout missing pass stream line: %s", got)
	}
	if !strings.Contains(got, "fail issues TestAccIssueLabel 2.50s log=") {
		t.Fatalf("stdout missing fail stream line with log path: %s", got)
	}
	if strings.Contains(got, "TestAccIssueLabel/Sub") {
		t.Fatalf("stdout streamed a subtest result: %s", got)
	}
}

func TestRunJSONExitCodes(t *testing.T) {
	root := t.TempDir()
	pf := jsonTestProvider()
	cases := []struct {
		name     string
		result   engine.RunResult
		wantCode int
	}{
		{
			name:     "all pass",
			result:   engine.RunResult{Tests: []engine.TestResult{{Package: "./github", Name: "TestAccOrgSettings", Status: provider.StatusPass}}},
			wantCode: 0,
		},
		{
			name:     "failed test",
			result:   engine.RunResult{Tests: []engine.TestResult{{Package: "./github", Name: "TestAccIssueLabel", Status: provider.StatusFail}}},
			wantCode: 1,
		},
		{
			name:     "build failed",
			result:   engine.RunResult{BuildFailed: true, BuildOutput: []string{"compile failed\n"}},
			wantCode: 1,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			runner := &fakeRunner{result: tc.result}
			d := deps{
				provider:  pf,
				newRunner: func(_ io.Writer) testRunner { return runner },
				list:      stubList([]string{"TestAccOrgSettings", "TestAccIssueLabel"}),
				cwd:       func() (string, error) { return root, nil },
			}
			var out, errOut bytes.Buffer
			code := runWithDeps([]string{"run", "--json", "--repo-root", root}, &out, &errOut, d)
			if code != tc.wantCode {
				t.Fatalf("exit code = %d, want %d; stderr=%s", code, tc.wantCode, errOut.String())
			}
			summary := lastEvent(t, decodeNDJSON(t, out.String()))
			if summary["exit_code"] != float64(tc.wantCode) {
				t.Fatalf("summary exit_code = %v, want %d", summary["exit_code"], tc.wantCode)
			}
		})
	}

	var out, errOut bytes.Buffer
	d := deps{
		provider:  pf,
		newRunner: func(_ io.Writer) testRunner { return &fakeRunner{} },
		list:      stubList([]string{"TestAccOrgSettings"}),
		cwd:       func() (string, error) { return root, nil },
	}
	code := runWithDeps([]string{"run", "--json", "--run", "[", "--repo-root", root}, &out, &errOut, d)
	if code != 2 {
		t.Fatalf("invalid regex exit code = %d, want 2", code)
	}
	if strings.TrimSpace(out.String()) != "" {
		t.Fatalf("usage error wrote JSON stdout, want empty stdout: %s", out.String())
	}
}

func TestGroupsJSONEmitsGroupEventsAndSummary(t *testing.T) {
	root := t.TempDir()
	d := deps{
		provider: jsonTestProvider(),
		list:     stubList([]string{"TestAccOrgSettings", "TestAccIssueLabel"}),
		cwd:      func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"groups", "--json", "--repo-root", root}, &out, &errOut, d)

	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s", code, errOut.String())
	}
	events := decodeNDJSON(t, out.String())
	if len(events) != 3 {
		t.Fatalf("expected two group events plus summary, got %d: %v", len(events), events)
	}
	if events[0]["type"] != "group" || events[1]["type"] != "group" {
		t.Fatalf("first two events should be groups: %v", events)
	}
	summary := lastEvent(t, events)
	if summary["type"] != "summary" || summary["command"] != "groups" {
		t.Fatalf("summary event = %v", summary)
	}
	if summary["groups"] != float64(2) || summary["tests"] != float64(2) || summary["exit_code"] != float64(0) {
		t.Fatalf("bad groups summary: %v", summary)
	}
}

func TestPreflightJSONEmitsChecksAndSummary(t *testing.T) {
	root := t.TempDir()
	pf := jsonTestProvider()
	pf.PreflightFn = func(_ context.Context, mode string, _ provider.TestRequirements) provider.PreflightReport {
		return provider.PreflightReport{
			Mode: mode,
			Checks: []provider.Check{
				{Name: "owner", Status: provider.CheckOK, Detail: "org accessible"},
				{Name: "rate-limit", Status: provider.CheckWarn, Detail: "low headroom", Fix: "wait"},
				{Name: "scopes", Status: provider.CheckFail, Detail: "missing scopes", Fix: "add scope"},
			},
		}
	}
	d := deps{
		provider: pf,
		list:     stubList([]string{"TestAccOrgSettings", "TestAccIssueLabel"}),
		cwd:      func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"preflight", "--json", "--mode", "organization", "--repo-root", root}, &out, &errOut, d)

	if code != 1 {
		t.Fatalf("exit code = %d, want 1; stderr=%s", code, errOut.String())
	}
	events := decodeNDJSON(t, out.String())
	if len(events) != 5 {
		t.Fatalf("expected plan event plus three check events plus summary, got %d: %v", len(events), events)
	}
	if events[0]["type"] != "plan" {
		t.Fatalf("first event = %v, want plan", events[0])
	}
	if events[1]["type"] != "check" || events[1]["status"] != "ok" {
		t.Fatalf("second event (first check) = %v", events[1])
	}
	summary := lastEvent(t, events)
	if summary["type"] != "summary" || summary["command"] != "preflight" {
		t.Fatalf("summary event = %v", summary)
	}
	if summary["ok"] != false || summary["exit_code"] != float64(1) {
		t.Fatalf("bad preflight summary: %v", summary)
	}
}

func TestRetryAndResumeJSONEmitGroups(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")
	st := engine.State{Provider: "test", Mode: "anonymous"}
	st.Plan = &engine.ExecutionPlan{Mode: "anonymous", Eligible: []string{"TestAccOrgSettings", "TestAccIssueLabel"}}
	st.Results = []engine.PersistResult{
		{Test: "TestAccIssueLabel", Status: "fail", Package: "./github"},
	}
	if err := st.Save(statePath); err != nil {
		t.Fatalf("seeding state: %v", err)
	}

	runner := &fakeRunner{
		result: engine.RunResult{
			Tests: []engine.TestResult{{Package: "./github", Name: "TestAccIssueLabel", Status: provider.StatusPass}},
		},
	}
	d := deps{
		provider:  jsonTestProvider(),
		newRunner: func(_ io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccOrgSettings", "TestAccIssueLabel"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"retry", "--failed", "--json", "--repo-root", root}, &out, &errOut, d)
	if code != 0 {
		t.Fatalf("retry exit code = %d, want 0; stderr=%s", code, errOut.String())
	}
	events := decodeNDJSON(t, out.String())
	if events[0]["group"] != "issues" {
		t.Fatalf("retry test event group = %v, want issues; events=%v", events[0]["group"], events)
	}

	st.Results = []engine.PersistResult{{Test: "TestAccIssueLabel", Status: "fail", Package: "./github"}}
	if err := st.Save(statePath); err != nil {
		t.Fatalf("reseeding state: %v", err)
	}
	runner.calls = 0
	out.Reset()
	errOut.Reset()
	code = runWithDeps([]string{"resume", "--json", "--repo-root", root}, &out, &errOut, d)
	if code != 0 {
		t.Fatalf("resume exit code = %d, want 0; stderr=%s", code, errOut.String())
	}
	events = decodeNDJSON(t, out.String())
	if events[0]["group"] != "issues" {
		t.Fatalf("resume test event group = %v, want issues; events=%v", events[0]["group"], events)
	}
}

func TestRunTUIFallsBackToTextWhenStdoutIsNotTTY(t *testing.T) {
	root := t.TempDir()
	runner := &fakeRunner{
		result: engine.RunResult{
			Tests: []engine.TestResult{{Package: "./github", Name: "TestAccOrgSettings", Status: provider.StatusPass}},
		},
	}
	d := deps{
		provider:  jsonTestProvider(),
		newRunner: func(_ io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccOrgSettings"}),
		cwd:       func() (string, error) { return root, nil },
		isTTY:     func() bool { return false },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"run", "--tui", "--repo-root", root}, &out, &errOut, d)

	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s", code, errOut.String())
	}
	if runner.calls != 1 {
		t.Fatalf("runner calls = %d, want 1", runner.calls)
	}
	if !strings.Contains(errOut.String(), "--tui requested but stdout is not a TTY") {
		t.Fatalf("stderr missing fallback note: %s", errOut.String())
	}
}

func TestGroupsJSONUnmatchedExitIncludesUnmatchedTests(t *testing.T) {
	root := t.TempDir()
	d := deps{
		provider: jsonTestProvider(),
		list:     stubList([]string{"TestAccOther"}),
		cwd:      func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"groups", "--json", "--unmatched", "--repo-root", root}, &out, &errOut, d)

	if code != 1 {
		t.Fatalf("exit code = %d, want 1; stderr=%s", code, errOut.String())
	}
	summary := lastEvent(t, decodeNDJSON(t, out.String()))
	if summary["exit_code"] != float64(1) {
		t.Fatalf("summary exit_code = %v, want 1", summary["exit_code"])
	}
	if _, ok := summary["unmatched_tests"]; !ok {
		t.Fatalf("summary missing unmatched_tests: %v", summary)
	}
}

func TestRunJSONWritesFailureLogsBeforeSummary(t *testing.T) {
	root := t.TempDir()
	runner := &fakeRunner{
		result: engine.RunResult{
			Tests: []engine.TestResult{{Package: "./github", Name: "TestAccIssueLabel", Status: provider.StatusFail, Output: []string{"boom\n"}}},
		},
	}
	d := deps{
		provider:  jsonTestProvider(),
		newRunner: func(_ io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccIssueLabel"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"run", "--json", "--repo-root", root}, &out, &errOut, d)
	if code != 1 {
		t.Fatalf("exit code = %d, want 1; stderr=%s", code, errOut.String())
	}
	summary := lastEvent(t, decodeNDJSON(t, out.String()))
	failures, ok := summary["failures"].([]any)
	if !ok || len(failures) != 1 {
		t.Fatalf("summary failures = %v, want one failure", summary["failures"])
	}
	failure, ok := failures[0].(map[string]any)
	if !ok {
		t.Fatalf("failure entry has type %T", failures[0])
	}
	logPath, ok := failure["failure_log_path"].(string)
	if !ok || logPath == "" {
		t.Fatalf("failure missing log path: %v", failure)
	}
	if _, err := os.Stat(logPath); err != nil {
		t.Fatalf("failure log was not written before summary: %v", err)
	}
}

// TestShouldLaunchDashboardPrecedence locks the TUI-decision precedence:
// --json > --no-tui > --tui > env (PULSAR_NO_TUI / PULSAR_FORCE_TTY) > TTY.
func TestShouldLaunchDashboardPrecedence(t *testing.T) {
	cases := []struct {
		name     string
		format   outputFormat
		tuiFlag  bool
		noTUI    bool
		isTTY    bool
		envNoTUI string
		envForce string
		want     bool
	}{
		{"json beats tui on tty", formatJSON, true, false, true, "", "", false},
		{"json beats no-tui", formatJSON, false, true, true, "", "", false},
		{"json beats force-tty env", formatJSON, false, false, false, "", "1", false},
		{"no-tui beats tui", formatText, true, true, true, "", "", false},
		{"no-tui beats force-tty env", formatText, false, true, false, "", "1", false},
		{"tui on tty launches", formatText, true, false, true, "", "", true},
		{"tui with force-tty on non-tty launches", formatText, true, false, false, "", "1", true},
		{"tui on non-tty falls back to text", formatText, true, false, false, "", "", false},
		{"env no-tui beats tty auto", formatText, false, false, true, "1", "", false},
		{"env no-tui beats force-tty env", formatText, false, false, false, "1", "1", false},
		{"tty auto launches", formatText, false, false, true, "", "", true},
		{"force-tty env auto launches on non-tty", formatText, false, false, false, "", "1", true},
		{"non-tty auto stays text", formatText, false, false, false, "", "", false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			d := deps{
				isTTY: func() bool { return tc.isTTY },
				getenv: getenvFromMap(map[string]string{
					"PULSAR_NO_TUI":    tc.envNoTUI,
					"PULSAR_FORCE_TTY": tc.envForce,
				}),
			}
			var errOut bytes.Buffer
			got := shouldLaunchDashboard(tc.format, tc.tuiFlag, tc.noTUI, d, &errOut)
			if got != tc.want {
				t.Fatalf("shouldLaunchDashboard = %v, want %v (stderr=%q)", got, tc.want, errOut.String())
			}
		})
	}
}

// TestRunJSONPreflightGateEmitsSummary verifies that when the provider
// preflight gate fails, `run --json` still emits a terminal summary object on
// stdout (pre_run_failed=true, exit_code 1) instead of an empty stream, and
// never invokes the runner.
func TestRunJSONPreflightGateEmitsSummary(t *testing.T) {
	root := t.TempDir()
	runner := &fakeRunner{result: engine.RunResult{}}
	prov := jsonTestProvider()
	prov.PreflightFn = func(_ context.Context, mode string, _ provider.TestRequirements) provider.PreflightReport {
		return provider.PreflightReport{
			Mode:   mode,
			Checks: []provider.Check{{Name: "identity", Status: provider.CheckFail, Detail: "no token"}},
		}
	}
	d := deps{
		provider:  prov,
		newRunner: func(_ io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccOrgSettings"}),
		cwd:       func() (string, error) { return root, nil },
		isTTY:     func() bool { return false },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"run", "--json", "--mode", "organization", "--repo-root", root}, &out, &errOut, d)

	if code != 1 {
		t.Fatalf("exit code = %d, want 1; stderr=%s", code, errOut.String())
	}
	if runner.calls != 0 {
		t.Fatalf("runner calls = %d, want 0 (must not run after gate)", runner.calls)
	}
	events := decodeNDJSON(t, out.String())
	if len(events) != 3 {
		t.Fatalf("stdout events = %d, want 3 (plan, capability, summary); got %v", len(events), events)
	}
	if events[0]["type"] != "plan" {
		t.Fatalf("event[0] type = %v, want plan", events[0]["type"])
	}
	if events[1]["type"] != "capability" || events[1]["name"] != "identity" || events[1]["status"] != "fail" {
		t.Fatalf("event[1] (capability) = %v", events[1])
	}
	summary := events[2]
	if summary["type"] != "summary" {
		t.Fatalf("event type = %v, want summary", summary["type"])
	}
	if summary["pre_run_failed"] != true {
		t.Fatalf("pre_run_failed = %v, want true", summary["pre_run_failed"])
	}
	if summary["exit_code"].(float64) != 1 {
		t.Fatalf("exit_code = %v, want 1", summary["exit_code"])
	}
	if cleanupStatus, ok := summary["cleanup_status"]; !ok || cleanupStatus != "" {
		t.Fatalf("cleanup_status = %v (present=%v), want present empty status before an accounting window opens",
			cleanupStatus, ok)
	}
	totals, ok := summary["totals"].(map[string]any)
	if !ok || totals["total"].(float64) != 0 {
		t.Fatalf("totals = %v, want zeroed totals", summary["totals"])
	}
}

// failOnceWriter fails the first Write call, then behaves normally. It models a
// transient stdout failure on the first streamed NDJSON event.
type failOnceWriter struct {
	failed bool
	buf    bytes.Buffer
}

func (w *failOnceWriter) Write(p []byte) (int, error) {
	if !w.failed {
		w.failed = true
		return 0, fmt.Errorf("simulated stdout failure")
	}
	return w.buf.Write(p)
}

// TestRunJSONStreamEncodeErrorForcesExit verifies that a failed streaming NDJSON
// event write forces a non-zero exit even when every test passed, so automation
// never sees a truncated stream reported as success. Once the underlying stdout
// write fails, encoding/json latches the error and the best-effort terminal
// summary cannot be written, so the exit code is the machine-observable signal.
func TestRunJSONStreamEncodeErrorForcesExit(t *testing.T) {
	root := t.TempDir()
	runner := &fakeRunner{
		result: engine.RunResult{
			Tests: []engine.TestResult{{Package: "./github", Name: "TestAccOrgSettings", Status: provider.StatusPass}},
		},
	}
	d := deps{
		provider:  jsonTestProvider(),
		newRunner: func(_ io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccOrgSettings"}),
		cwd:       func() (string, error) { return root, nil },
		isTTY:     func() bool { return false },
	}

	out := &failOnceWriter{}
	var errOut bytes.Buffer
	code := runWithDeps([]string{"run", "--json", "--repo-root", root}, out, &errOut, d)

	if code != 1 {
		t.Fatalf("exit code = %d, want 1 after stream write failure", code)
	}
	if !strings.Contains(errOut.String(), "writing JSON") {
		t.Fatalf("expected stderr to report the stream write failure, got: %s", errOut.String())
	}
}

// ── planning integration (Task 6): JSON ordering and redaction ─────────────

// TestRunJSONEventOrderPlanExcludedCapabilityTestSummary verifies the full
// NDJSON event order for `run --json`: plan, then one excluded event per
// plan.Excluded entry, then one capability event per preflight check, then
// one test event per executed test, then the final summary object.
func TestRunJSONEventOrderPlanExcludedCapabilityTestSummary(t *testing.T) {
	root := t.TempDir()
	runner := &fakeRunner{
		result: engine.RunResult{
			Tests: []engine.TestResult{{Package: "./github", Name: "TestAccKeep", Status: provider.StatusPass}},
		},
	}
	pf := &fakeprovider.Fake{
		NameVal:  "test",
		Packages: []string{"./..."},
		Pattern:  "^TestAcc",
		ModesVal: testModes,
		RequirementsFn: func(name string) (provider.TestRequirements, bool) {
			if name == "TestAccExcluded" {
				return provider.TestRequirements{Modes: []string{"organization"}}, true
			}
			return provider.TestRequirements{}, true
		},
		PreflightFn: func(_ context.Context, mode string, _ provider.TestRequirements) provider.PreflightReport {
			return provider.PreflightReport{
				Mode:   mode,
				Checks: []provider.Check{{Name: "rate-limit", Status: provider.CheckWarn, Detail: "low headroom", Fix: "wait"}},
			}
		},
	}
	d := deps{
		provider:  pf,
		newRunner: func(_ io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccKeep", "TestAccExcluded"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"run", "--json", "--mode", "anonymous", "--repo-root", root}, &out, &errOut, d)
	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s", code, errOut.String())
	}
	events := decodeNDJSON(t, out.String())
	if len(events) != 5 {
		t.Fatalf("expected plan+excluded+capability+test+summary (5 events), got %d: %v", len(events), events)
	}
	wantTypes := []string{"plan", "excluded", "capability", "test", "summary"}
	for i, want := range wantTypes {
		if got := events[i]["type"]; got != want {
			t.Errorf("events[%d][type] = %v, want %v; events=%v", i, got, want, events)
		}
	}
	if events[1]["test"] != "TestAccExcluded" {
		t.Errorf("excluded event test = %v, want TestAccExcluded", events[1]["test"])
	}
	if events[2]["name"] != "rate-limit" || events[2]["status"] != "warn" {
		t.Errorf("capability event = %v", events[2])
	}
	if events[3]["name"] != "TestAccKeep" {
		t.Errorf("test event name = %v, want TestAccKeep", events[3]["name"])
	}
}

func TestRunJSONSummaryTotalsExcludeIncompatibleTests(t *testing.T) {
	root := t.TempDir()
	runner := &fakeRunner{
		result: engine.RunResult{
			Tests: []engine.TestResult{{Package: "./github", Name: "TestAccKeep", Status: provider.StatusPass}},
		},
	}
	pf := jsonTestProvider()
	pf.RequirementsFn = func(name string) (provider.TestRequirements, bool) {
		if name == "TestAccIncompatible" {
			return provider.TestRequirements{Modes: []string{"organization"}}, true
		}
		return provider.TestRequirements{}, true
	}
	d := deps{
		provider:  pf,
		newRunner: func(_ io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccKeep", "TestAccIncompatible"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"run", "--json", "--mode", "anonymous", "--repo-root", root}, &out, &errOut, d)
	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s", code, errOut.String())
	}

	summary := lastEvent(t, decodeNDJSON(t, out.String()))
	totals := summary["totals"].(map[string]any)
	if totals["total"] != float64(1) {
		t.Errorf("summary total = %v, want 1", totals["total"])
	}
	groups := summary["groups"].([]any)
	if len(groups) != 1 {
		t.Fatalf("summary groups = %v, want one group", groups)
	}
	group := groups[0].(map[string]any)
	if group["total"] != float64(1) {
		t.Errorf("group total = %v, want 1", group["total"])
	}
}

// TestRunJSONPlanFailureEmitsOnlyTerminalSummary verifies that a planning
// failure (an unknown --group) writes exactly one object to stdout in JSON
// mode: the terminal run summary carrying the mapped exit code. No plan or
// excluded event is fabricated for a plan that was never built, and the
// stream is never left empty - automation must always have something to
// parse. Text mode still prints only the human error (see
// TestRunPlanFailureSkipsPreflightAndRunner).
func TestRunJSONPlanFailureEmitsOnlyTerminalSummary(t *testing.T) {
	root := t.TempDir()
	runner := &fakeRunner{}
	d := deps{
		provider:  jsonTestProvider(),
		newRunner: func(_ io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccOrgSettings"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"run", "--json", "--group", "does-not-exist", "--repo-root", root}, &out, &errOut, d)

	if code != 1 {
		t.Fatalf("exit code = %d, want 1", code)
	}
	if runner.calls != 0 {
		t.Fatalf("runner calls = %d, want 0", runner.calls)
	}
	events := decodeNDJSON(t, out.String())
	if len(events) != 1 {
		t.Fatalf("events = %d, want only the terminal summary: %v", len(events), events)
	}
	summary := events[0]
	if summary["type"] != "summary" || summary["command"] != "run" || summary["exit_code"] != float64(1) {
		t.Fatalf("terminal event = %v, want the run summary with exit_code 1", summary)
	}
}

// TestRunJSONCapabilityEventRedactsCheckSecret is the JSON-mode counterpart
// of TestRunTextPreflightFailureRedactsCheckSecret (Task 4 follow-up): a
// configured secret embedded in a failing Check's Detail/Fix must never
// appear in the emitted jsonCapabilityEvent.
func TestRunJSONCapabilityEventRedactsCheckSecret(t *testing.T) {
	const secret = "ghp_SUPERSECRETVALUE"
	root := t.TempDir()
	pf := &fakeprovider.Fake{
		NameVal:        "test",
		Packages:       []string{"./..."},
		Pattern:        "^TestAcc",
		ModesVal:       testModes,
		Secrets:        []string{"GITHUB_TOKEN"},
		RequirementsFn: allowAllRequirements,
		PreflightFn: func(_ context.Context, mode string, _ provider.TestRequirements) provider.PreflightReport {
			return provider.PreflightReport{
				Mode: mode,
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
		newRunner: func(_ io.Writer) testRunner { return &fakeRunner{} },
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
	code := runWithDeps([]string{"run", "--json", "--repo-root", root}, &out, &errOut, d)
	if code != 1 {
		t.Fatalf("exit code = %d, want 1", code)
	}
	if strings.Contains(out.String(), secret) {
		t.Fatalf("stdout leaked secret: %s", out.String())
	}
	events := decodeNDJSON(t, out.String())
	var capEvent map[string]any
	for _, ev := range events {
		if ev["type"] == "capability" {
			capEvent = ev
			break
		}
	}
	if capEvent == nil {
		t.Fatalf("expected a capability event, got: %v", events)
	}
	if !strings.Contains(fmt.Sprint(capEvent["detail"]), "***REDACTED***") {
		t.Errorf("capability detail not redacted: %v", capEvent["detail"])
	}
	if !strings.Contains(fmt.Sprint(capEvent["fix"]), "***REDACTED***") {
		t.Errorf("capability fix not redacted: %v", capEvent["fix"])
	}
}

// TestPreflightStandaloneJSONCheckEventRedactsSecret mirrors
// TestRunJSONCapabilityEventRedactsCheckSecret for standalone preflight's
// existing jsonCheckEvent emission path (Task 4 follow-up #1): the real
// printPreflight/JSON check emission must redact, not just provider-internal
// unit tests.
func TestPreflightStandaloneJSONCheckEventRedactsSecret(t *testing.T) {
	const secret = "ghp_SUPERSECRETVALUE"
	root := t.TempDir()
	pf := &fakeprovider.Fake{
		NameVal:        "test",
		Packages:       []string{"./..."},
		Pattern:        "^TestAcc",
		ModesVal:       testModes,
		Secrets:        []string{"GITHUB_TOKEN"},
		RequirementsFn: allowAllRequirements,
		PreflightFn: func(_ context.Context, mode string, _ provider.TestRequirements) provider.PreflightReport {
			return provider.PreflightReport{
				Mode: mode,
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
	code := runWithDeps([]string{"preflight", "--json", "--repo-root", root}, &out, &errOut, d)
	if code != 1 {
		t.Fatalf("exit code = %d, want 1", code)
	}
	if strings.Contains(out.String(), secret) {
		t.Fatalf("stdout leaked secret: %s", out.String())
	}
	events := decodeNDJSON(t, out.String())
	var checkEvent map[string]any
	for _, ev := range events {
		if ev["type"] == "check" {
			checkEvent = ev
			break
		}
	}
	if checkEvent == nil {
		t.Fatalf("expected a check event, got: %v", events)
	}
	if !strings.Contains(fmt.Sprint(checkEvent["detail"]), "***REDACTED***") {
		t.Errorf("check detail not redacted: %v", checkEvent["detail"])
	}
	if !strings.Contains(fmt.Sprint(checkEvent["fix"]), "***REDACTED***") {
		t.Errorf("check fix not redacted: %v", checkEvent["fix"])
	}
}

func TestE2EOrphanDeltaJSONShapeAndCleanupCommands(t *testing.T) {
	root := filepath.Join(t.TempDir(), "provider checkout")
	if err := os.MkdirAll(root, 0o755); err != nil {
		t.Fatal(err)
	}
	baseline := []provider.Resource{
		{Kind: "issue_label", Name: "tf-acc-test-existing-label", URL: "https://example.test/labels/existing"},
		{Kind: "repository", Name: "tf-acc-test-existing-repo", URL: "https://example.test/repos/existing"},
	}
	final := append(append([]provider.Resource{}, baseline...), provider.Resource{
		Kind: "repository", Name: "tf-acc-test-new-repo", URL: "https://example.test/repos/new",
	})
	runner := &e2eRunner{results: []engine.RunResult{{Tests: []engine.TestResult{{
		Package: "./github", Name: "TestAccThing", Status: provider.StatusPass,
	}}}}}
	prov := newE2EProvider(provider.PreflightReport{})
	prov.ModesVal = []provider.Mode{{Name: "anonymous"}, {Name: "individual"}, {Name: "organization"}}
	prov.orphanSeq = [][]provider.Resource{baseline, final}
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccThing"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"e2e", "--mode", "individual", "--json", "--repo-root", root}, &out, &errOut, d)
	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s stdout=%s", code, errOut.String(), out.String())
	}

	events := decodeNDJSON(t, out.String())
	if len(events) < 2 || events[len(events)-2]["type"] != "orphan_delta" {
		t.Fatalf("penultimate event = %v, want orphan_delta before final summary; events=%v", events[len(events)-2], events)
	}
	delta := events[len(events)-2]
	preview := "terraform-provider-tester orphans --repo-root '" + root + "' --mode individual --run-delta"
	cleanup := "terraform-provider-tester sweep --repo-root '" + root + "' --mode individual --run-delta --confirm"
	for field, want := range map[string]any{
		"schema_version":  float64(1),
		"type":            "orphan_delta",
		"mode":            "individual",
		"baseline":        float64(2),
		"final":           float64(3),
		"pre_existing":    float64(2),
		"new":             float64(1),
		"cleanup_status":  "complete",
		"preview_command": preview,
		"cleanup_command": cleanup,
	} {
		if got := delta[field]; got != want {
			t.Errorf("orphan_delta[%q] = %v, want %v; event=%v", field, got, want, delta)
		}
	}

	summary := lastEvent(t, events)
	if summary["type"] != "summary" || summary["command"] != "e2e" {
		t.Fatalf("final event = %v, want e2e summary", summary)
	}
	for field, want := range map[string]any{
		"baseline_orphans": float64(2),
		"new_orphans":      float64(1),
		"cleanup_status":   "complete",
		"preview_command":  preview,
		"cleanup_command":  cleanup,
		"exit_code":        float64(0),
	} {
		if got := summary[field]; got != want {
			t.Errorf("summary[%q] = %v, want %v; summary=%v", field, got, want, summary)
		}
	}
	if prov.sweepCalls != 0 {
		t.Fatalf("guided e2e sweep calls = %d, want 0", prov.sweepCalls)
	}
}

func TestE2EAnonymousJSONOmitsOrphanDeltaAndCleanupCommands(t *testing.T) {
	root := t.TempDir()
	runner := &e2eRunner{results: []engine.RunResult{{Tests: []engine.TestResult{{
		Package: "./github", Name: "TestAccPublic", Status: provider.StatusPass,
	}}}}}
	prov := newE2EProvider(provider.PreflightReport{})
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccPublic"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"e2e", "--mode", "anonymous", "--json", "--repo-root", root}, &out, &errOut, d)
	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s stdout=%s", code, errOut.String(), out.String())
	}
	events := decodeNDJSON(t, out.String())
	for _, event := range events {
		if event["type"] == "orphan_delta" {
			t.Fatalf("anonymous output emitted orphan_delta: %v", event)
		}
	}
	summary := lastEvent(t, events)
	if summary["type"] != "summary" || summary["command"] != "e2e" {
		t.Fatalf("final event = %v, want e2e summary", summary)
	}
	for _, field := range []string{"preview_command", "cleanup_command"} {
		if got := summary[field]; got != "" {
			t.Errorf("anonymous summary[%q] = %v, want empty", field, got)
		}
	}
	if got := summary["cleanup_status"]; got != "not-applicable" {
		t.Errorf("anonymous summary cleanup_status = %v, want not-applicable", got)
	}
}
