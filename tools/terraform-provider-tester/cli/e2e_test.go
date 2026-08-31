package cli

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
	"syscall"
	"testing"
	"time"

	"github.com/github/terraform-provider-tester/engine"
	"github.com/github/terraform-provider-tester/fakeprovider"
	"github.com/github/terraform-provider-tester/internal/redact"
	"github.com/github/terraform-provider-tester/provider"
)

type e2eRunner struct {
	results []engine.RunResult
	errs    []error
	specs   []engine.RunSpec
	onRun   func(engine.RunSpec)
}

type failE2ESummaryWriter struct {
	buf bytes.Buffer
}

func (w *failE2ESummaryWriter) Write(p []byte) (int, error) {
	if bytes.Contains(p, []byte(`"type":"summary"`)) &&
		bytes.Contains(p, []byte(`"command":"e2e"`)) {
		return 0, errors.New("simulated e2e summary write failure")
	}
	return w.buf.Write(p)
}

type signalBlockingE2ERunner struct{}

func (signalBlockingE2ERunner) Run(ctx context.Context, _ engine.RunSpec, _ func(engine.TestResult)) (engine.RunResult, error) {
	fmt.Fprintln(os.Stdout, "e2e-run-signal-ready")
	<-ctx.Done()
	return engine.RunResult{}, ctx.Err()
}

func runE2ESignalSequence(t *testing.T, testName, helperEnv string, markers ...string) (string, string, error) {
	t.Helper()
	cmd := exec.Command(os.Args[0], "-test.run=^"+testName+"$", "-test.v")
	cmd.Env = append(os.Environ(), helperEnv+"=1")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		t.Fatal(err)
	}
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Start(); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if cmd.ProcessState == nil {
			_ = cmd.Process.Kill()
		}
	})

	lines := make(chan string, 128)
	scanDone := make(chan error, 1)
	var childOutput bytes.Buffer
	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Fprintln(&childOutput, line)
			lines <- line
		}
		close(lines)
		scanDone <- scanner.Err()
	}()

	stopAndReport := func(message string) {
		t.Helper()
		_ = cmd.Process.Kill()
		waitErr := cmd.Wait()
		scanErr := <-scanDone
		t.Fatalf("%s; wait=%v scan=%v stderr=%s stdout=%s",
			message, waitErr, scanErr, stderr.String(), childOutput.String())
	}
	for _, marker := range markers {
		timer := time.NewTimer(5 * time.Second)
		found := false
		for !found {
			select {
			case line, ok := <-lines:
				if !ok {
					timer.Stop()
					waitErr := cmd.Wait()
					scanErr := <-scanDone
					t.Fatalf("helper exited before %q; wait=%v scan=%v stderr=%s stdout=%s",
						marker, waitErr, scanErr, stderr.String(), childOutput.String())
				}
				found = line == marker
			case <-timer.C:
				stopAndReport("timed out waiting for helper marker " + marker)
			}
		}
		timer.Stop()
		if err := cmd.Process.Signal(syscall.SIGTERM); err != nil {
			stopAndReport("sending SIGTERM after helper marker " + marker + ": " + err.Error())
		}
	}

	waitDone := make(chan error, 1)
	go func() { waitDone <- cmd.Wait() }()
	var waitErr error
	select {
	case waitErr = <-waitDone:
	case <-time.After(3 * time.Second):
		_ = cmd.Process.Kill()
		<-waitDone
		if err := <-scanDone; err != nil {
			t.Fatalf("reading helper stdout after timeout: %v", err)
		}
		t.Fatalf("runE2E helper did not exit promptly after SIGTERM; stderr=%s stdout=%s",
			stderr.String(), childOutput.String())
	}
	if err := <-scanDone; err != nil {
		t.Fatalf("reading helper stdout: %v", err)
	}
	return childOutput.String(), stderr.String(), waitErr
}

func (r *e2eRunner) Run(_ context.Context, spec engine.RunSpec, sink func(engine.TestResult)) (engine.RunResult, error) {
	r.specs = append(r.specs, spec)
	if r.onRun != nil {
		r.onRun(spec)
	}
	i := len(r.specs) - 1
	var result engine.RunResult
	if i < len(r.results) {
		result = r.results[i]
	}
	for _, test := range result.Tests {
		if sink != nil {
			sink(test)
		}
	}
	if i < len(r.errs) {
		return result, r.errs[i]
	}
	return result, nil
}

type e2eProvider struct {
	*fakeprovider.Fake
	orphanCalls int
	sweepCalls  int
	orphanModes []string
	sweepModes  []string
	orphans     []provider.Resource
	orphanSeq   [][]provider.Resource
	orphanErrs  []error
	orphanErr   error
	orphanFn    func(context.Context, string, int) ([]provider.Resource, error)
	callLog     *[]string
}

func (p *e2eProvider) Orphans(ctx context.Context, mode string) ([]provider.Resource, error) {
	p.orphanCalls++
	p.orphanModes = append(p.orphanModes, mode)
	if p.callLog != nil {
		label := "baseline"
		if p.orphanCalls > 1 {
			label = "final"
		}
		*p.callLog = append(*p.callLog, label)
	}
	if p.orphanFn != nil {
		return p.orphanFn(ctx, mode, p.orphanCalls)
	}
	if p.orphanCalls <= len(p.orphanErrs) && p.orphanErrs[p.orphanCalls-1] != nil {
		return nil, p.orphanErrs[p.orphanCalls-1]
	}
	if p.orphanErr != nil {
		return nil, p.orphanErr
	}
	if p.orphanCalls <= len(p.orphanSeq) {
		return p.orphanSeq[p.orphanCalls-1], nil
	}
	return p.orphans, nil
}

func (p *e2eProvider) Sweep(_ context.Context, mode string, _ provider.SweepOpts) error {
	p.sweepCalls++
	p.sweepModes = append(p.sweepModes, mode)
	return nil
}

func newE2EProvider(preflight provider.PreflightReport) *e2eProvider {
	return &e2eProvider{Fake: &fakeprovider.Fake{
		NameVal:        "test",
		Packages:       []string{"./github/..."},
		Pattern:        "^TestAcc",
		ModesVal:       testModes,
		RequirementsFn: allowAllRequirements,
		PreflightFn: func(_ context.Context, mode string, _ provider.TestRequirements) provider.PreflightReport {
			preflight.Mode = mode
			return preflight
		},
	}}
}

func saveInterruptedResumeState(t *testing.T, root string, eligible []string, results []engine.PersistResult, baseline []provider.Resource) {
	t.Helper()
	state := engine.State{
		Provider: "test",
		Mode:     "organization",
		Plan:     planWithEligible("organization", eligible),
		Results:  results,
		Orphans: &engine.OrphanAccounting{
			Mode:             "organization",
			Baseline:         baseline,
			BaselineCaptured: true,
			CleanupStatus:    engine.CleanupBaselineOnly,
		},
	}
	if err := state.Save(filepath.Join(root, ".pulsar-state.json")); err != nil {
		t.Fatalf("seeding interrupted state: %v", err)
	}
}

type forbiddenProvider struct {
	t *testing.T
}

func (p forbiddenProvider) called(method string) {
	p.t.Helper()
	p.t.Fatalf("provider call %s occurred before resume mode mismatch was rejected", method)
}

func (p forbiddenProvider) Name() string {
	p.called("Name")
	return ""
}
func (p forbiddenProvider) TestPackages() []string {
	p.called("TestPackages")
	return nil
}
func (p forbiddenProvider) TestPattern() string {
	p.called("TestPattern")
	return ""
}
func (p forbiddenProvider) Modes() []provider.Mode {
	p.called("Modes")
	return nil
}
func (p forbiddenProvider) EnvFor(string) []provider.EnvVar {
	p.called("EnvFor")
	return nil
}
func (p forbiddenProvider) GroupOf(string) string {
	p.called("GroupOf")
	return ""
}
func (p forbiddenProvider) RequirementsFor(string) (provider.TestRequirements, bool) {
	p.called("RequirementsFor")
	return provider.TestRequirements{}, false
}
func (p forbiddenProvider) Preflight(context.Context, string, provider.TestRequirements) provider.PreflightReport {
	p.called("Preflight")
	return provider.PreflightReport{}
}
func (p forbiddenProvider) Discover(context.Context, provider.DiscoverOpts) (provider.DiscoveryResult, error) {
	p.called("Discover")
	return provider.DiscoveryResult{}, nil
}
func (p forbiddenProvider) Orphans(context.Context, string) ([]provider.Resource, error) {
	p.called("Orphans")
	return nil, nil
}
func (p forbiddenProvider) Sweep(context.Context, string, provider.SweepOpts) error {
	p.called("Sweep")
	return nil
}
func (p forbiddenProvider) SecretEnvKeys() []string {
	p.called("SecretEnvKeys")
	return nil
}

func TestInterruptedRunPreservesPlanAndBaseline(t *testing.T) {
	root := t.TempDir()
	baseline := []provider.Resource{{
		Kind: "repository", Name: "tf-acc-test-existing", URL: "https://example.test/existing",
	}}
	runner := &e2eRunner{
		results: []engine.RunResult{{Tests: []engine.TestResult{{
			Package: "./github", Name: "TestAccCompleted", Status: provider.StatusPass,
		}}}},
		errs: []error{context.Canceled},
	}
	prov := newE2EProvider(provider.PreflightReport{})
	prov.orphanSeq = [][]provider.Resource{baseline, baseline}
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccCompleted", "TestAccPending"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"e2e", "--mode", "organization", "--json", "--repo-root", root}, &out, &errOut, d)

	if code != 1 {
		t.Fatalf("exit code = %d, want 1; stderr=%s stdout=%s", code, errOut.String(), out.String())
	}
	state, err := engine.Load(filepath.Join(root, ".pulsar-state.json"))
	if err != nil {
		t.Fatal(err)
	}
	if state.Plan == nil || !reflect.DeepEqual(state.Plan.Eligible, []string{"TestAccCompleted", "TestAccPending"}) {
		t.Fatalf("plan = %+v, want original eligible selection", state.Plan)
	}
	if state.Orphans == nil || !reflect.DeepEqual(state.Orphans.Baseline, baseline) {
		t.Fatalf("orphans = %+v, want original baseline %+v", state.Orphans, baseline)
	}
	if len(state.Results) != 1 || state.Results[0].Test != "TestAccCompleted" {
		t.Fatalf("results = %+v, want streamed partial result", state.Results)
	}
}

func TestResumeReusesOriginalBaseline(t *testing.T) {
	root := t.TempDir()
	baseline := []provider.Resource{{
		Kind: "repository", Name: "tf-acc-test-existing", URL: "https://example.test/existing",
	}}
	created := provider.Resource{Kind: "issue_label", Name: "tf-acc-test-created", URL: "https://example.test/created"}
	final := append(append([]provider.Resource{}, baseline...), created)
	saveInterruptedResumeState(t, root, []string{"TestAccPending"},
		[]engine.PersistResult{{Test: "TestAccPending", Status: "fail", Package: "./github"}}, baseline)
	runner := &e2eRunner{results: []engine.RunResult{{Tests: []engine.TestResult{{
		Package: "./github", Name: "TestAccPending", Status: provider.StatusPass,
	}}}}}
	prov := newE2EProvider(provider.PreflightReport{})
	prov.orphanSeq = [][]provider.Resource{final}
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccPending"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"resume", "--repo-root", root}, &out, &errOut, d)

	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s", code, errOut.String())
	}
	if prov.orphanCalls != 1 {
		t.Fatalf("orphan calls = %d, want exactly one final call and no new baseline", prov.orphanCalls)
	}
	state, err := engine.Load(filepath.Join(root, ".pulsar-state.json"))
	if err != nil {
		t.Fatal(err)
	}
	if state.Orphans == nil || !reflect.DeepEqual(state.Orphans.Baseline, baseline) ||
		!reflect.DeepEqual(state.Orphans.Final, final) ||
		!reflect.DeepEqual(state.Orphans.New, []provider.Resource{created}) {
		t.Fatalf("orphan accounting = %+v, want original baseline and final delta", state.Orphans)
	}
}

func TestResumeTextPrintsFinalizedOrphanDeltaAndSafeCommands(t *testing.T) {
	root := t.TempDir()
	baseline := []provider.Resource{{
		Kind: "repository", Name: "tf-acc-test-existing", URL: "https://example.test/existing",
	}}
	created := provider.Resource{
		Kind: "issue_label", Name: "tf-acc-test-created", URL: "https://example.test/created",
	}
	final := append(append([]provider.Resource{}, baseline...), created)
	saveInterruptedResumeState(t, root, []string{"TestAccPending"},
		[]engine.PersistResult{{Test: "TestAccPending", Status: "fail", Package: "./github"}}, baseline)
	runner := &e2eRunner{results: []engine.RunResult{{Tests: []engine.TestResult{{
		Package: "./github", Name: "TestAccPending", Status: provider.StatusPass,
	}}}}}
	prov := newE2EProvider(provider.PreflightReport{})
	prov.orphanSeq = [][]provider.Resource{final}
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccPending"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"resume", "--repo-root", root}, &out, &errOut, d)

	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s stdout=%s", code, errOut.String(), out.String())
	}
	text := out.String()
	for _, want := range []string{
		"Orphan accounting:",
		"Baseline: 1",
		"Pre-existing: 1",
		"Final: 2",
		"New (this run's cleanup obligation): 1",
		"Cleanup status: complete",
		"Preview command: terraform-provider-tester orphans --repo-root '" + root + "' --mode organization --run-delta",
		"Cleanup command: terraform-provider-tester sweep --repo-root '" + root + "' --mode organization --run-delta --confirm",
	} {
		if !strings.Contains(text, want) {
			t.Errorf("resume text missing %q:\n%s", want, text)
		}
	}
	if got := strings.Count(text, "this run's cleanup obligation"); got != 1 {
		t.Errorf("cleanup-obligation label count = %d, want 1 on New only:\n%s", got, text)
	}
}

func TestResumeTextUnknownAccountingSuppressesCommands(t *testing.T) {
	root := t.TempDir()
	baseline := []provider.Resource{{
		Kind: "repository", Name: "tf-acc-test-existing", URL: "https://example.test/existing",
	}}
	saveInterruptedResumeState(t, root, []string{"TestAccPending"},
		[]engine.PersistResult{{Test: "TestAccPending", Status: "fail", Package: "./github"}}, baseline)
	runner := &e2eRunner{results: []engine.RunResult{{Tests: []engine.TestResult{{
		Package: "./github", Name: "TestAccPending", Status: provider.StatusPass,
	}}}}}
	prov := newE2EProvider(provider.PreflightReport{})
	prov.orphanErr = errors.New("final orphan scan failed")
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccPending"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"resume", "--repo-root", root}, &out, &errOut, d)

	if code != 1 {
		t.Fatalf("exit code = %d, want 1; stderr=%s stdout=%s", code, errOut.String(), out.String())
	}
	text := out.String()
	for _, want := range []string{
		"Baseline: 1",
		"Pre-existing: unproven",
		"Final: unproven",
		"New (this run's cleanup obligation): unproven",
		"Cleanup status: unknown",
	} {
		if !strings.Contains(text, want) {
			t.Errorf("resume unknown text missing %q:\n%s", want, text)
		}
	}
	for _, forbidden := range []string{
		"Preview command:",
		"Cleanup command:",
		"terraform-provider-tester orphans",
		"terraform-provider-tester sweep",
	} {
		if strings.Contains(text, forbidden) {
			t.Errorf("resume unknown text advertises unavailable cleanup command %q:\n%s", forbidden, text)
		}
	}
}

func TestResumeTextEmptySelectionPrintsFinalizedOrphanDelta(t *testing.T) {
	root := t.TempDir()
	baseline := []provider.Resource{{
		Kind: "repository", Name: "tf-acc-test-existing", URL: "https://example.test/existing",
	}}
	saveInterruptedResumeState(t, root, []string{"TestAccComplete"},
		[]engine.PersistResult{{Test: "TestAccComplete", Status: "pass", Package: "./github"}}, baseline)
	prov := newE2EProvider(provider.PreflightReport{})
	prov.orphanSeq = [][]provider.Resource{baseline}
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return &e2eRunner{} },
		list:      stubList([]string{"TestAccComplete"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"resume", "--repo-root", root}, &out, &errOut, d)

	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s stdout=%s", code, errOut.String(), out.String())
	}
	for _, want := range []string{
		"Orphan accounting:",
		"Baseline: 1",
		"Pre-existing: 1",
		"Final: 1",
		"New (this run's cleanup obligation): 0",
		"Cleanup status: complete",
		"Preview command:",
		"Cleanup command:",
		"nothing to resume",
	} {
		if !strings.Contains(out.String(), want) {
			t.Errorf("empty resume text missing %q:\n%s", want, out.String())
		}
	}
}

func TestResumeJSONFinalizationSuccessIsTerminal(t *testing.T) {
	const secret = "ghp_RESUME_JSON_SECRET_1234567890"
	root := t.TempDir()
	baseline := []provider.Resource{{
		Kind: "repository", Name: "tf-acc-test-existing", URL: "https://example.test/existing",
	}}
	created := provider.Resource{
		Kind: "issue_label", Name: "tf-acc-test-created", URL: "https://example.test/created?token=" + secret,
	}
	final := append(append([]provider.Resource{}, baseline...), created)
	saveInterruptedResumeState(t, root, []string{"TestAccPending"},
		[]engine.PersistResult{{Test: "TestAccPending", Status: "fail", Package: "./github"}}, baseline)
	runner := &e2eRunner{results: []engine.RunResult{{Tests: []engine.TestResult{{
		Package: "./github", Name: "TestAccPending", Status: provider.StatusPass,
	}}}}}
	prov := newE2EProvider(provider.PreflightReport{})
	prov.Secrets = []string{"GITHUB_TOKEN"}
	prov.orphanSeq = [][]provider.Resource{final}
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccPending"}),
		cwd:       func() (string, error) { return root, nil },
		getenv: func(key string) string {
			if key == "GITHUB_TOKEN" {
				return secret
			}
			return ""
		},
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"resume", "--json", "--repo-root", root}, &out, &errOut, d)

	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s stdout=%s", code, errOut.String(), out.String())
	}
	if strings.Contains(out.String(), secret) {
		t.Fatalf("resume JSON leaked secret-shaped orphan metadata:\n%s", out.String())
	}
	events := decodeNDJSON(t, out.String())
	if len(events) < 2 {
		t.Fatalf("events = %v, want orphan_delta and terminal resume summary", events)
	}
	var summaryIndexes []int
	for i, event := range events {
		if event["type"] == "summary" {
			summaryIndexes = append(summaryIndexes, i)
		}
	}
	if !reflect.DeepEqual(summaryIndexes, []int{len(events) - 3, len(events) - 1}) {
		t.Fatalf("summary indexes = %v, want inner run summary then terminal resume summary; events=%v", summaryIndexes, events)
	}
	delta := events[len(events)-2]
	summary := lastEvent(t, events)
	for field, want := range map[string]any{
		"type":           "orphan_delta",
		"cleanup_status": engine.CleanupComplete,
		"baseline":       float64(1),
		"final":          float64(2),
		"pre_existing":   float64(1),
		"new":            float64(1),
	} {
		if got := delta[field]; got != want {
			t.Errorf("orphan_delta[%q] = %v, want %v; event=%v", field, got, want, delta)
		}
	}
	if summary["type"] != "summary" || summary["command"] != "resume" ||
		summary["cleanup_status"] != engine.CleanupComplete ||
		summary["exit_code"] != float64(code) {
		t.Fatalf("final event = %v, want complete resume summary matching process exit %d", summary, code)
	}
}

func TestResumeJSONFinalCheckFailureClearsStaleDeltaAndIsTerminal(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")
	baseline := []provider.Resource{{
		Kind: "repository", Name: "tf-acc-test-existing", URL: "https://example.test/existing",
	}}
	saveInterruptedResumeState(t, root, []string{"TestAccPending"},
		[]engine.PersistResult{{Test: "TestAccPending", Status: "fail", Package: "./github"}}, baseline)
	state, err := engine.Load(statePath)
	if err != nil {
		t.Fatal(err)
	}
	state.Orphans.Final = []provider.Resource{{Kind: "repository", Name: "tf-acc-test-stale-final"}}
	state.Orphans.PreExisting = []provider.Resource{{Kind: "repository", Name: "tf-acc-test-stale-existing"}}
	state.Orphans.New = []provider.Resource{{Kind: "repository", Name: "tf-acc-test-stale-new"}}
	state.Orphans.FinalCaptured = true
	state.Orphans.CleanupStatus = engine.CleanupComplete
	if err := state.Save(statePath); err != nil {
		t.Fatalf("seeding stale final accounting: %v", err)
	}

	runner := &e2eRunner{results: []engine.RunResult{{Tests: []engine.TestResult{{
		Package: "./github", Name: "TestAccPending", Status: provider.StatusPass,
	}}}}}
	prov := newE2EProvider(provider.PreflightReport{})
	prov.orphanErr = errors.New("final orphan scan failed")
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccPending"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"resume", "--json", "--repo-root", root}, &out, &errOut, d)

	if code != 1 {
		t.Fatalf("exit code = %d, want 1; stderr=%s stdout=%s", code, errOut.String(), out.String())
	}
	events := decodeNDJSON(t, out.String())
	if len(events) < 2 {
		t.Fatalf("events = %v, want unknown orphan_delta and terminal resume summary", events)
	}
	delta := events[len(events)-2]
	summary := lastEvent(t, events)
	if delta["type"] != "orphan_delta" || delta["cleanup_status"] != engine.CleanupUnknown {
		t.Fatalf("penultimate event = %v, want unknown orphan_delta", delta)
	}
	for _, field := range []string{"final", "pre_existing", "new"} {
		if delta[field] != float64(0) {
			t.Errorf("orphan_delta[%q] = %v, want cleared zero placeholder", field, delta[field])
		}
	}
	if summary["type"] != "summary" || summary["command"] != "resume" ||
		summary["cleanup_status"] != engine.CleanupUnknown ||
		summary["exit_code"] != float64(code) {
		t.Fatalf("final event = %v, want unknown resume summary matching process exit %d", summary, code)
	}
	for _, event := range []map[string]any{delta, summary} {
		for _, field := range []string{"preview_command", "cleanup_command"} {
			if value, ok := event[field]; ok && value != "" {
				t.Errorf("%s = %v, want empty for unknown cleanup; event=%v", field, value, event)
			}
		}
	}

	persisted, err := engine.Load(statePath)
	if err != nil {
		t.Fatal(err)
	}
	if persisted.Orphans == nil || persisted.Orphans.CleanupStatus != engine.CleanupUnknown ||
		persisted.Orphans.Final != nil || persisted.Orphans.PreExisting != nil ||
		persisted.Orphans.New != nil || persisted.Orphans.FinalCaptured {
		t.Fatalf("persisted orphan accounting = %+v, want cleared unknown final state", persisted.Orphans)
	}

	orphansCalls := 0
	sweepCalls := 0
	cleanupDeps := deps{
		provider: &fakeprovider.Fake{
			OrphansFn: func(_ context.Context, _ string) ([]provider.Resource, error) {
				orphansCalls++
				return nil, nil
			},
			SweepFn: func(_ context.Context, _ string, _ provider.SweepOpts) error {
				sweepCalls++
				return nil
			},
		},
		cwd: func() (string, error) { return root, nil },
	}
	for _, args := range [][]string{
		{"orphans", "--repo-root", root, "--run-delta"},
		{"sweep", "--repo-root", root, "--run-delta", "--confirm"},
	} {
		out.Reset()
		errOut.Reset()
		if cleanupCode := runWithDeps(args, &out, &errOut, cleanupDeps); cleanupCode != 2 {
			t.Fatalf("%v exit code = %d, want 2; stderr=%s", args, cleanupCode, errOut.String())
		}
	}
	if orphansCalls != 0 || sweepCalls != 0 {
		t.Fatalf("cleanup provider calls after unknown resume: Orphans=%d Sweep=%d, want zero", orphansCalls, sweepCalls)
	}
}

func TestResumeJSONEmptySelectionStillFinalizesBeforeSummary(t *testing.T) {
	root := t.TempDir()
	baseline := []provider.Resource{{
		Kind: "repository", Name: "tf-acc-test-existing", URL: "https://example.test/existing",
	}}
	saveInterruptedResumeState(t, root, []string{"TestAccComplete"},
		[]engine.PersistResult{{Test: "TestAccComplete", Status: "pass", Package: "./github"}}, baseline)
	runner := &e2eRunner{}
	prov := newE2EProvider(provider.PreflightReport{})
	prov.orphanSeq = [][]provider.Resource{baseline}
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccComplete"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"resume", "--json", "--repo-root", root}, &out, &errOut, d)

	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s stdout=%s", code, errOut.String(), out.String())
	}
	if len(runner.specs) != 0 {
		t.Fatalf("runner calls = %d, want 0 for empty selection", len(runner.specs))
	}
	events := decodeNDJSON(t, out.String())
	if len(events) != 2 {
		t.Fatalf("events = %v, want only orphan_delta then resume summary", events)
	}
	if events[0]["type"] != "orphan_delta" || events[0]["cleanup_status"] != engine.CleanupComplete {
		t.Fatalf("first event = %v, want complete orphan_delta", events[0])
	}
	summary := events[1]
	if summary["type"] != "summary" || summary["command"] != "resume" ||
		summary["cleanup_status"] != engine.CleanupComplete ||
		summary["exit_code"] != float64(code) {
		t.Fatalf("final event = %v, want complete empty resume summary matching exit %d", summary, code)
	}
	selection := summary["selection"].(map[string]any)
	if selection["tests"] != float64(0) {
		t.Fatalf("selection = %v, want zero tests", selection)
	}
}

func TestResumeJSONEmptySelectionDiscoveryFailureStillEndsWithSummary(t *testing.T) {
	root := t.TempDir()
	baseline := []provider.Resource{{
		Kind: "repository", Name: "tf-acc-test-existing", URL: "https://example.test/existing",
	}}
	saveInterruptedResumeState(t, root, []string{"TestAccComplete"},
		[]engine.PersistResult{{Test: "TestAccComplete", Status: "pass", Package: "./github"}}, baseline)
	runner := &e2eRunner{}
	prov := newE2EProvider(provider.PreflightReport{})
	prov.orphanSeq = [][]provider.Resource{baseline}
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      failList(errors.New("build failed during resume discovery")),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"resume", "--json", "--repo-root", root}, &out, &errOut, d)

	if code != 1 {
		t.Fatalf("exit code = %d, want 1; stderr=%s stdout=%s", code, errOut.String(), out.String())
	}
	if prov.orphanCalls != 1 {
		t.Fatalf("orphan calls = %d, want finalized accounting before discovery failure", prov.orphanCalls)
	}
	if len(runner.specs) != 0 {
		t.Fatalf("runner calls = %d, want 0 for empty selection", len(runner.specs))
	}
	if !strings.Contains(errOut.String(), "build failed during resume discovery") {
		t.Fatalf("stderr = %q, want discovery failure", errOut.String())
	}
	events := decodeNDJSON(t, out.String())
	if len(events) == 0 {
		t.Fatal("resume JSON stream is empty; want terminal failure summary")
	}
	summary := lastEvent(t, events)
	if summary["type"] != "summary" || summary["command"] != "resume" ||
		summary["cleanup_status"] != engine.CleanupComplete ||
		summary["exit_code"] != float64(code) {
		t.Fatalf("final event = %v, want nonzero resume summary after finalized accounting", summary)
	}
}

func TestResumePlainCredentialedRunWithoutAccountingKeepsLegacyBehavior(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")
	state := engine.State{
		Provider: "test",
		Mode:     "organization",
		Plan:     planWithEligible("organization", []string{"TestAccPending"}),
		Results:  []engine.PersistResult{{Test: "TestAccPending", Status: "fail", Package: "./github"}},
	}
	if err := state.Save(statePath); err != nil {
		t.Fatalf("seeding plain run state: %v", err)
	}
	runner := &e2eRunner{results: []engine.RunResult{{Tests: []engine.TestResult{{
		Package: "./github", Name: "TestAccPending", Status: provider.StatusPass,
	}}}}}
	prov := newE2EProvider(provider.PreflightReport{})
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccPending"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"resume", "--json", "--repo-root", root}, &out, &errOut, d)

	if code != 0 {
		t.Fatalf("exit code = %d, want 0 for plain run resume; stderr=%s stdout=%s", code, errOut.String(), out.String())
	}
	if prov.orphanCalls != 0 {
		t.Fatalf("orphan calls = %d, want 0 without a guided accounting record", prov.orphanCalls)
	}
	if prov.sweepCalls != 0 {
		t.Fatalf("sweep calls = %d, want 0", prov.sweepCalls)
	}
	persisted, err := engine.Load(statePath)
	if err != nil {
		t.Fatal(err)
	}
	if persisted.Orphans != nil {
		t.Fatalf("persisted orphan accounting = %+v, want nil legacy state", persisted.Orphans)
	}
	events := decodeNDJSON(t, out.String())
	summary := lastEvent(t, events)
	if summary["type"] != "summary" || summary["command"] != "resume" ||
		summary["exit_code"] != float64(0) {
		t.Fatalf("final event = %v, want successful resume summary", summary)
	}
	if cleanupStatus, ok := summary["cleanup_status"]; !ok || cleanupStatus != "" {
		t.Fatalf("legacy resume cleanup_status = %v (present=%v), want present empty status: %v",
			cleanupStatus, ok, summary)
	}
	for _, event := range events {
		if event["type"] == "orphan_delta" {
			t.Fatalf("legacy resume emitted orphan delta: %v", event)
		}
	}
}

func TestResumePlainCredentialedEmptySelectionWithoutAccountingKeepsLegacyBehavior(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")
	state := engine.State{
		Provider: "test",
		Mode:     "organization",
		Plan:     planWithEligible("organization", []string{"TestAccComplete"}),
		Results:  []engine.PersistResult{{Test: "TestAccComplete", Status: "pass", Package: "./github"}},
	}
	if err := state.Save(statePath); err != nil {
		t.Fatalf("seeding plain run state: %v", err)
	}
	prov := newE2EProvider(provider.PreflightReport{})
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return &e2eRunner{} },
		list:      stubList([]string{"TestAccComplete"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"resume", "--json", "--repo-root", root}, &out, &errOut, d)

	if code != 0 {
		t.Fatalf("exit code = %d, want 0 for empty plain run resume; stderr=%s stdout=%s", code, errOut.String(), out.String())
	}
	if prov.orphanCalls != 0 {
		t.Fatalf("orphan calls = %d, want 0 without a guided accounting record", prov.orphanCalls)
	}
	persisted, err := engine.Load(statePath)
	if err != nil {
		t.Fatal(err)
	}
	if persisted.Orphans != nil {
		t.Fatalf("persisted orphan accounting = %+v, want nil legacy state", persisted.Orphans)
	}
	events := decodeNDJSON(t, out.String())
	if len(events) != 1 {
		t.Fatalf("events = %v, want only terminal resume summary", events)
	}
	summary := lastEvent(t, events)
	if summary["type"] != "summary" || summary["command"] != "resume" ||
		summary["exit_code"] != float64(0) {
		t.Fatalf("final event = %v, want successful empty resume summary", summary)
	}
	if cleanupStatus, ok := summary["cleanup_status"]; !ok || cleanupStatus != "" {
		t.Fatalf("legacy resume cleanup_status = %v (present=%v), want present empty status: %v",
			cleanupStatus, ok, summary)
	}
}

func TestResumeJSONMissingBaselineFailsClosedWithoutProviderCall(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")
	state := engine.State{
		Provider: "test",
		Mode:     "organization",
		Plan:     planWithEligible("organization", []string{"TestAccPending"}),
		Results:  []engine.PersistResult{{Test: "TestAccPending", Status: "fail", Package: "./github"}},
		Orphans: &engine.OrphanAccounting{
			Mode: "organization",
			New:  []provider.Resource{{Kind: "repository", Name: "tf-acc-test-stale-new"}},
		},
	}
	if err := state.Save(statePath); err != nil {
		t.Fatalf("seeding state without captured baseline: %v", err)
	}
	runner := &e2eRunner{results: []engine.RunResult{{Tests: []engine.TestResult{{
		Package: "./github", Name: "TestAccPending", Status: provider.StatusPass,
	}}}}}
	prov := newE2EProvider(provider.PreflightReport{})
	prov.orphanFn = func(_ context.Context, _ string, _ int) ([]provider.Resource, error) {
		t.Fatal("final orphan provider call occurred without a captured baseline")
		return nil, nil
	}
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccPending"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"resume", "--json", "--repo-root", root}, &out, &errOut, d)

	if code != 1 {
		t.Fatalf("exit code = %d, want 1; stderr=%s stdout=%s", code, errOut.String(), out.String())
	}
	if prov.orphanCalls != 0 {
		t.Fatalf("orphan calls = %d, want 0 without a captured baseline", prov.orphanCalls)
	}
	events := decodeNDJSON(t, out.String())
	delta := events[len(events)-2]
	summary := lastEvent(t, events)
	if delta["type"] != "orphan_delta" || delta["cleanup_status"] != engine.CleanupUnknown {
		t.Fatalf("penultimate event = %v, want unknown orphan_delta", delta)
	}
	if summary["command"] != "resume" || summary["cleanup_status"] != engine.CleanupUnknown ||
		summary["exit_code"] != float64(code) {
		t.Fatalf("summary = %v, want unknown resume summary matching exit %d", summary, code)
	}
	persisted, err := engine.Load(statePath)
	if err != nil {
		t.Fatal(err)
	}
	if persisted.Orphans == nil || persisted.Orphans.CleanupStatus != engine.CleanupUnknown ||
		persisted.Orphans.New != nil || persisted.Orphans.FinalCaptured {
		t.Fatalf("persisted orphan accounting = %+v, want cleared unknown state", persisted.Orphans)
	}
}

func TestResumeFinalizationRedactsSecretShapedResourceURL(t *testing.T) {
	const secret = "ghp_RESUMEFINALIZATIONSECRET123456789"
	root := t.TempDir()
	final := []provider.Resource{{
		Kind: "repository",
		Name: "tf-acc-test-created",
		URL:  "https://example.test/repos/created?token=" + secret,
	}}
	saveInterruptedResumeState(t, root, []string{"TestAccPending"},
		[]engine.PersistResult{{Test: "TestAccPending", Status: "fail", Package: "./github"}}, nil)
	runner := &e2eRunner{results: []engine.RunResult{{Tests: []engine.TestResult{{
		Package: "./github", Name: "TestAccPending", Status: provider.StatusPass,
	}}}}}
	prov := newE2EProvider(provider.PreflightReport{})
	prov.orphanSeq = [][]provider.Resource{final}
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccPending"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"resume", "--repo-root", root}, &out, &errOut, d)
	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s", code, errOut.String())
	}
	rawState, err := os.ReadFile(filepath.Join(root, ".pulsar-state.json"))
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(rawState), secret) {
		t.Fatalf("persisted resume state leaked secret-shaped resource URL:\n%s", rawState)
	}
	if !strings.Contains(string(rawState), "***REDACTED***") {
		t.Fatalf("persisted resume state missing redaction marker:\n%s", rawState)
	}
}

func TestResumeAttributesPreInterruptionResourceToOriginalRun(t *testing.T) {
	root := t.TempDir()
	baseline := []provider.Resource{}
	createdBeforeInterruption := provider.Resource{
		Kind: "repository", Name: "tf-acc-test-before-interruption", URL: "https://example.test/before-interruption",
	}
	saveInterruptedResumeState(t, root, []string{"TestAccCompleted", "TestAccPending"},
		[]engine.PersistResult{{Test: "TestAccCompleted", Status: "pass", Package: "./github"}}, baseline)
	runner := &e2eRunner{results: []engine.RunResult{{Tests: []engine.TestResult{{
		Package: "./github", Name: "TestAccPending", Status: provider.StatusPass,
	}}}}}
	prov := newE2EProvider(provider.PreflightReport{})
	prov.orphanSeq = [][]provider.Resource{{createdBeforeInterruption}}
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccCompleted", "TestAccPending"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"resume", "--repo-root", root}, &out, &errOut, d)

	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s", code, errOut.String())
	}
	state, err := engine.Load(filepath.Join(root, ".pulsar-state.json"))
	if err != nil {
		t.Fatal(err)
	}
	if state.Orphans == nil || !reflect.DeepEqual(state.Orphans.Baseline, baseline) ||
		!reflect.DeepEqual(state.Orphans.New, []provider.Resource{createdBeforeInterruption}) {
		t.Fatalf("orphan accounting = %+v, want pre-interruption resource new relative to original baseline", state.Orphans)
	}
}

func TestRepeatedInterruptionKeepsOriginalBaseline(t *testing.T) {
	root := t.TempDir()
	baseline := []provider.Resource{{
		Kind: "repository", Name: "tf-acc-test-existing", URL: "https://example.test/existing",
	}}
	saveInterruptedResumeState(t, root, []string{"TestAccFirst", "TestAccSecond", "TestAccPending"}, nil, baseline)
	runner := &e2eRunner{
		results: []engine.RunResult{
			{Tests: []engine.TestResult{{Package: "./github", Name: "TestAccFirst", Status: provider.StatusPass}}},
			{Tests: []engine.TestResult{{Package: "./github", Name: "TestAccSecond", Status: provider.StatusPass}}},
		},
		errs: []error{context.Canceled, context.Canceled},
	}
	prov := newE2EProvider(provider.PreflightReport{})
	prov.orphanSeq = [][]provider.Resource{baseline, baseline}
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccFirst", "TestAccSecond", "TestAccPending"}),
		cwd:       func() (string, error) { return root, nil },
	}

	for attempt := 1; attempt <= 2; attempt++ {
		var out, errOut bytes.Buffer
		code := runWithDeps([]string{"resume", "--repo-root", root}, &out, &errOut, d)
		if code != 1 {
			t.Fatalf("attempt %d exit code = %d, want 1; stderr=%s", attempt, code, errOut.String())
		}
		state, err := engine.Load(filepath.Join(root, ".pulsar-state.json"))
		if err != nil {
			t.Fatal(err)
		}
		if state.Orphans == nil || !reflect.DeepEqual(state.Orphans.Baseline, baseline) {
			t.Fatalf("attempt %d orphan accounting = %+v, want original baseline %+v", attempt, state.Orphans, baseline)
		}
	}
	if prov.orphanCalls != 2 {
		t.Fatalf("orphan calls = %d, want one final check per interrupted resume", prov.orphanCalls)
	}
	state, err := engine.Load(filepath.Join(root, ".pulsar-state.json"))
	if err != nil {
		t.Fatal(err)
	}
	if len(state.Results) != 2 {
		t.Fatalf("results = %+v, want both streamed partial results after repeated interruption", state.Results)
	}
}

func TestResumeModeMismatchMakesNoProviderCalls(t *testing.T) {
	root := t.TempDir()
	saveInterruptedResumeState(t, root, []string{"TestAccPending"},
		[]engine.PersistResult{{Test: "TestAccPending", Status: "fail"}}, []provider.Resource{})
	runner := &fakeRunner{}
	d := deps{
		provider:  forbiddenProvider{t: t},
		newRunner: func(io.Writer) testRunner { return runner },
		list: func(context.Context, string, []string, string) ([]string, error) {
			t.Fatal("test discovery occurred before mode mismatch was rejected")
			return nil, nil
		},
		cwd: func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"resume", "--mode", "anonymous", "--repo-root", root}, &out, &errOut, d)

	if code != 2 {
		t.Fatalf("exit code = %d, want 2; stderr=%s", code, errOut.String())
	}
	if runner.calls != 0 {
		t.Fatalf("runner calls = %d, want 0", runner.calls)
	}
}

func TestResumeModeGuardIgnoresModeAfterPositionalArgument(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")
	baseline := []provider.Resource{{
		Kind: "repository", Name: "tf-acc-test-existing", URL: "https://example.test/existing",
	}}
	newResources := []provider.Resource{{
		Kind: "repository", Name: "tf-acc-test-new", URL: "https://example.test/new",
	}}
	final := append(append([]provider.Resource{}, baseline...), newResources...)
	accounting := &engine.OrphanAccounting{
		Mode:             "individual",
		Baseline:         baseline,
		Final:            final,
		PreExisting:      baseline,
		New:              newResources,
		BaselineCaptured: true,
		FinalCaptured:    true,
		CleanupStatus:    engine.CleanupComplete,
	}
	state := engine.State{
		Provider: "test",
		Mode:     "individual",
		Plan:     planWithEligible("individual", []string{"TestAccPending"}),
		Results:  []engine.PersistResult{{Test: "TestAccPending", Status: "fail", Package: "./github"}},
		Orphans:  accounting,
	}
	if err := state.Save(statePath); err != nil {
		t.Fatalf("seeding resume state: %v", err)
	}

	runner := &fakeRunner{}
	prov := newE2EProvider(provider.PreflightReport{})
	prov.orphans = final
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccPending"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps(
		[]string{"resume", "--repo-root", root, "bogus", "--mode", "organization"},
		&out,
		&errOut,
		d,
	)

	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s stdout=%s", code, errOut.String(), out.String())
	}
	if runner.calls != 1 {
		t.Fatalf("runner calls = %d, want 1", runner.calls)
	}
	if !extraEnvContains(runner.spec.ExtraEnv, "GH_TEST_AUTH_MODE=individual") {
		t.Fatalf("child ExtraEnv = %v, want persisted individual mode", runner.spec.ExtraEnv)
	}
	if !reflect.DeepEqual(prov.orphanModes, []string{"individual"}) {
		t.Fatalf("orphan modes = %v, want final scan in persisted individual mode", prov.orphanModes)
	}

	persisted, err := engine.Load(statePath)
	if err != nil {
		t.Fatalf("loading persisted state: %v", err)
	}
	if !reflect.DeepEqual(persisted.Orphans, accounting) {
		t.Fatalf("persisted orphan accounting = %+v, want unchanged %+v", persisted.Orphans, accounting)
	}
}

func TestResumePersistedOrphanModeMismatchFailsClosedInTextAndJSON(t *testing.T) {
	type terminalResult struct {
		exitCode      int
		cleanupStatus string
	}
	results := make(map[string]terminalResult)

	for _, tc := range []struct {
		name    string
		json    bool
		tty     bool
		envFile bool
		single  bool
	}{
		{name: "text"},
		{name: "JSON", json: true},
		{name: "TTY", tty: true},
		{name: "JSON with env file", json: true, envFile: true},
		{name: "single-dash JSON with env file", json: true, envFile: true, single: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := t.TempDir()
			cwdRoot := root
			if tc.single {
				cwdRoot = t.TempDir()
			}
			statePath := filepath.Join(root, ".pulsar-state.json")
			baseline := []provider.Resource{{
				Kind: "repository", Name: "tf-acc-test-existing", URL: "https://example.test/existing",
			}}
			state := engine.State{
				Provider: "test",
				Mode:     "organization",
				Plan:     planWithEligible("organization", []string{"TestAccPending"}),
				Results:  []engine.PersistResult{{Test: "TestAccPending", Status: "fail", Package: "./github"}},
				Orphans: &engine.OrphanAccounting{
					Mode:             "individual",
					Baseline:         baseline,
					Final:            []provider.Resource{{Kind: "repository", Name: "tf-acc-test-stale-final"}},
					PreExisting:      []provider.Resource{{Kind: "repository", Name: "tf-acc-test-stale-existing"}},
					New:              []provider.Resource{{Kind: "repository", Name: "tf-acc-test-stale-new"}},
					BaselineCaptured: true,
					FinalCaptured:    true,
					CleanupStatus:    engine.CleanupComplete,
				},
			}
			if err := state.Save(statePath); err != nil {
				t.Fatalf("seeding mismatched orphan mode: %v", err)
			}

			runner := &fakeRunner{}
			d := deps{
				provider:  forbiddenProvider{t: t},
				newRunner: func(io.Writer) testRunner { return runner },
				list: func(context.Context, string, []string, string) ([]string, error) {
					t.Fatal("test discovery occurred before persisted orphan mode mismatch was rejected")
					return nil, nil
				},
				cwd:   func() (string, error) { return cwdRoot, nil },
				isTTY: func() bool { return tc.tty },
			}

			repoRootFlag := "--repo-root"
			jsonFlag := "--json"
			if tc.single {
				repoRootFlag = "-repo-root"
				jsonFlag = "-json"
			}
			args := []string{"resume", repoRootFlag, root}
			if tc.json {
				args = append(args, jsonFlag)
			}
			if tc.envFile {
				envPath := filepath.Join(root, "resume.env")
				if err := os.WriteFile(envPath, []byte("GH_TEST_AUTH_MODE=organization\n"), 0o600); err != nil {
					t.Fatalf("writing env file: %v", err)
				}
				args = append(args, "--env-file", envPath)
			}
			var out, errOut bytes.Buffer
			code := runWithDeps(args, &out, &errOut, d)

			if code != 1 {
				t.Fatalf("exit code = %d, want 1; stderr=%s stdout=%s", code, errOut.String(), out.String())
			}
			if runner.calls != 0 {
				t.Fatalf("runner calls = %d, want 0", runner.calls)
			}
			for _, want := range []string{
				"state/orphan-mode-mismatch",
				`orphan accounting mode "individual"`,
				`resolved resume mode "organization"`,
				"start a new guided `e2e` run",
			} {
				if !strings.Contains(errOut.String(), want) {
					t.Errorf("stderr = %q, want actionable mode-mismatch text %q", errOut.String(), want)
				}
			}

			persisted, err := engine.Load(statePath)
			if err != nil {
				t.Fatal(err)
			}
			if persisted.Orphans == nil {
				t.Fatal("persisted orphan accounting = nil, want original accounting invalidated")
			}
			if persisted.Orphans.Mode != "individual" {
				t.Errorf("persisted orphan mode = %q, want original individual mode", persisted.Orphans.Mode)
			}
			if !persisted.Orphans.BaselineCaptured || !reflect.DeepEqual(persisted.Orphans.Baseline, baseline) {
				t.Errorf("persisted baseline = %+v (captured=%v), want original captured baseline %+v",
					persisted.Orphans.Baseline, persisted.Orphans.BaselineCaptured, baseline)
			}
			if persisted.Orphans.Final != nil || persisted.Orphans.PreExisting != nil ||
				persisted.Orphans.New != nil || persisted.Orphans.FinalCaptured {
				t.Errorf("persisted final accounting = %+v, want final/pre-existing/new cleared", persisted.Orphans)
			}
			if persisted.Orphans.CleanupStatus != engine.CleanupUnknown {
				t.Errorf("persisted cleanup status = %q, want %q",
					persisted.Orphans.CleanupStatus, engine.CleanupUnknown)
			}

			if !tc.json {
				if !strings.Contains(out.String(), "Cleanup status: unknown") ||
					!strings.Contains(out.String(), "resume stopped before execution") {
					t.Fatalf("text output = %q, want unknown accounting and terminal resume result", out.String())
				}
				for _, unsafe := range []string{"Preview command:", "Cleanup command:"} {
					if strings.Contains(out.String(), unsafe) {
						t.Errorf("text output contains %q for unknown accounting: %s", unsafe, out.String())
					}
				}
				results[tc.name] = terminalResult{exitCode: code, cleanupStatus: engine.CleanupUnknown}
				return
			}

			events := decodeNDJSON(t, out.String())
			if len(events) != 2 {
				t.Fatalf("events = %v, want orphan_delta then terminal resume summary", events)
			}
			delta := events[0]
			if delta["type"] != "orphan_delta" || delta["mode"] != "individual" ||
				delta["cleanup_status"] != engine.CleanupUnknown {
				t.Fatalf("orphan delta = %v, want original mode and unknown cleanup", delta)
			}
			for _, field := range []string{"final", "pre_existing", "new"} {
				if delta[field] != float64(0) {
					t.Errorf("orphan_delta[%q] = %v, want cleared zero placeholder", field, delta[field])
				}
			}
			summary := events[1]
			if summary["type"] != "summary" || summary["command"] != "resume" ||
				summary["cleanup_status"] != engine.CleanupUnknown ||
				summary["exit_code"] != float64(code) {
				t.Fatalf("terminal summary = %v, want unknown nonzero resume result", summary)
			}
			results[tc.name] = terminalResult{
				exitCode:      code,
				cleanupStatus: summary["cleanup_status"].(string),
			}
		})
	}

	if !reflect.DeepEqual(results["text"], results["JSON"]) {
		t.Fatalf("text/JSON terminal results differ: text=%+v JSON=%+v", results["text"], results["JSON"])
	}
	for _, name := range []string{"TTY", "JSON with env file", "single-dash JSON with env file"} {
		if !reflect.DeepEqual(results["text"], results[name]) {
			t.Fatalf("text/%s terminal results differ: text=%+v %s=%+v",
				name, results["text"], name, results[name])
		}
	}
}

func TestResumePerformsFinalOrphanCheckBeforeUnlock(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")
	lockPath := statePath + ".lock"
	baseline := []provider.Resource{{
		Kind: "repository", Name: "tf-acc-test-existing", URL: "https://example.test/existing",
	}}
	saveInterruptedResumeState(t, root, []string{"TestAccPending"},
		[]engine.PersistResult{{Test: "TestAccPending", Status: "fail", Package: "./github"}}, baseline)
	var calls []string
	runner := &e2eRunner{
		results: []engine.RunResult{{Tests: []engine.TestResult{{
			Package: "./github", Name: "TestAccPending", Status: provider.StatusPass,
		}}}},
		onRun: func(engine.RunSpec) {
			if _, err := os.Stat(lockPath); err != nil {
				t.Fatalf("runner did not observe logical lock: %v", err)
			}
			calls = append(calls, "runner")
		},
	}
	prov := newE2EProvider(provider.PreflightReport{})
	prov.orphanFn = func(_ context.Context, _ string, _ int) ([]provider.Resource, error) {
		if _, err := os.Stat(lockPath); err != nil {
			t.Fatalf("final orphan check occurred after unlock: %v", err)
		}
		calls = append(calls, "final")
		return baseline, nil
	}
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccPending"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"resume", "--repo-root", root}, &out, &errOut, d)

	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s", code, errOut.String())
	}
	if !reflect.DeepEqual(calls, []string{"runner", "final"}) {
		t.Fatalf("call order = %v, want runner then final under one lock", calls)
	}
	if prov.orphanCalls != 1 {
		t.Fatalf("orphan calls = %d, want exactly one final call", prov.orphanCalls)
	}
	if _, err := os.Stat(lockPath); !os.IsNotExist(err) {
		t.Fatalf("lock still exists after resume: %v", err)
	}
}

func TestE2EAnonymousDoesNotCaptureOrphanBaseline(t *testing.T) {
	root := t.TempDir()
	var calls []string
	runner := &e2eRunner{
		results: []engine.RunResult{{Tests: []engine.TestResult{{
			Package: "./github", Name: "TestAccPublic", Status: provider.StatusPass,
		}}}},
		onRun: func(engine.RunSpec) { calls = append(calls, "runner") },
	}
	prov := newE2EProvider(provider.PreflightReport{})
	prov.callLog = &calls
	prov.orphanErr = errors.New("anonymous must not list provider orphans")
	prov.PreflightFn = func(_ context.Context, mode string, _ provider.TestRequirements) provider.PreflightReport {
		calls = append(calls, "preflight")
		return provider.PreflightReport{Mode: mode}
	}
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
	if prov.orphanCalls != 0 {
		t.Fatalf("orphan calls = %d, want 0 for anonymous baseline", prov.orphanCalls)
	}
	if !reflect.DeepEqual(calls, []string{"preflight", "runner"}) {
		t.Fatalf("call order = %v, want preflight then runner without baseline", calls)
	}
	state, err := engine.Load(filepath.Join(root, ".pulsar-state.json"))
	if err != nil {
		t.Fatal(err)
	}
	if state.Orphans == nil {
		t.Fatal("state.Orphans = nil, want anonymous accounting")
	}
	if state.Orphans.Mode != "anonymous" || state.Orphans.CleanupStatus != engine.CleanupNotApplicable || state.Orphans.BaselineCaptured {
		t.Fatalf("anonymous orphan accounting = %+v, want mode anonymous cleanup not-applicable without captured baseline", state.Orphans)
	}
}

func TestE2EAnonymousTextOmitsOrphanDeltaAndCleanupCommands(t *testing.T) {
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
	code := runWithDeps([]string{"e2e", "--mode", "anonymous", "--repo-root", root}, &out, &errOut, d)
	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s stdout=%s", code, errOut.String(), out.String())
	}
	text := out.String()
	for _, unsafe := range []string{
		"Orphan accounting:",
		"terraform-provider-tester orphans",
		"terraform-provider-tester sweep",
	} {
		if strings.Contains(text, unsafe) {
			t.Errorf("anonymous text output contains %q:\n%s", unsafe, text)
		}
	}
	if !strings.HasSuffix(text, "destructive cleanup was not run\n") {
		t.Fatalf("anonymous text output does not end with the final e2e summary:\n%s", text)
	}
}

func TestE2ECapturesBaselineBeforeRunner(t *testing.T) {
	root := t.TempDir()
	baseline := []provider.Resource{{Kind: "repository", Name: "tf-acc-test-before", URL: "https://example.test/before"}}
	var calls []string
	runner := &e2eRunner{
		results: []engine.RunResult{{Tests: []engine.TestResult{{
			Package: "./github", Name: "TestAccThing", Status: provider.StatusPass,
		}}}},
		onRun: func(engine.RunSpec) {
			calls = append(calls, "runner")
			state, err := engine.Load(filepath.Join(root, ".pulsar-state.json"))
			if err != nil {
				t.Fatalf("loading state during runner: %v", err)
			}
			if state.Orphans == nil || !state.Orphans.BaselineCaptured {
				t.Fatalf("runner observed state.Orphans = %+v, want captured baseline before runner", state.Orphans)
			}
		},
	}
	prov := newE2EProvider(provider.PreflightReport{})
	prov.callLog = &calls
	prov.orphans = baseline
	prov.PreflightFn = func(_ context.Context, mode string, _ provider.TestRequirements) provider.PreflightReport {
		calls = append(calls, "preflight")
		return provider.PreflightReport{Mode: mode}
	}
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccThing"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"e2e", "--mode", "organization", "--json", "--repo-root", root}, &out, &errOut, d)

	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s stdout=%s", code, errOut.String(), out.String())
	}
	if !reflect.DeepEqual(calls, []string{"preflight", "baseline", "runner", "final"}) {
		t.Fatalf("call order = %v, want preflight, baseline, runner, final", calls)
	}
	state, err := engine.Load(filepath.Join(root, ".pulsar-state.json"))
	if err != nil {
		t.Fatal(err)
	}
	if state.Orphans == nil || state.Orphans.Mode != "organization" || !state.Orphans.BaselineCaptured ||
		!state.Orphans.FinalCaptured || state.Orphans.CleanupStatus != engine.CleanupComplete ||
		!reflect.DeepEqual(state.Orphans.Baseline, baseline) || !reflect.DeepEqual(state.Orphans.Final, baseline) {
		t.Fatalf("orphan accounting = %+v, want captured organization baseline/final %+v", state.Orphans, baseline)
	}
}

func TestE2EBaselineFailureBlocksRunner(t *testing.T) {
	root := t.TempDir()
	var calls []string
	runner := &e2eRunner{
		results: []engine.RunResult{{Tests: []engine.TestResult{{
			Package: "./github", Name: "TestAccThing", Status: provider.StatusPass,
		}}}},
		onRun: func(engine.RunSpec) { calls = append(calls, "runner") },
	}
	prov := newE2EProvider(provider.PreflightReport{})
	prov.callLog = &calls
	prov.orphanErr = errors.New("baseline failed")
	prov.PreflightFn = func(_ context.Context, mode string, _ provider.TestRequirements) provider.PreflightReport {
		calls = append(calls, "preflight")
		return provider.PreflightReport{Mode: mode}
	}
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccThing"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"e2e", "--mode", "organization", "--json", "--repo-root", root}, &out, &errOut, d)

	if code != 1 {
		t.Fatalf("exit code = %d, want 1; stderr=%s stdout=%s", code, errOut.String(), out.String())
	}
	if len(runner.specs) != 0 {
		t.Fatalf("runner calls = %d, want 0 when baseline capture fails", len(runner.specs))
	}
	if !reflect.DeepEqual(calls, []string{"preflight", "baseline"}) {
		t.Fatalf("call order = %v, want preflight, baseline and no runner", calls)
	}
	summary := lastEvent(t, decodeNDJSON(t, out.String()))
	if summary["cleanup_status"] != engine.CleanupUnknown || summary["exit_code"] != float64(1) {
		t.Fatalf("summary = %v, want cleanup_status unknown and exit_code 1", summary)
	}
}

func TestE2ESetupDeadlineDuringPreflightEmitsTerminalSummary(t *testing.T) {
	for _, tc := range []struct {
		name string
		args []string
		json bool
	}{
		{name: "text", args: []string{"--mode", "individual"}},
		{name: "JSON", args: []string{"--mode", "individual", "--json"}, json: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := t.TempDir()
			runner := &e2eRunner{}
			prov := newE2EProvider(provider.PreflightReport{})
			prov.ModesVal = []provider.Mode{{Name: "anonymous"}, {Name: "individual"}, {Name: "organization"}}
			prov.PreflightFn = func(ctx context.Context, mode string, _ provider.TestRequirements) provider.PreflightReport {
				<-ctx.Done()
				if !errors.Is(ctx.Err(), context.DeadlineExceeded) {
					t.Fatalf("preflight context error = %v, want deadline exceeded", ctx.Err())
				}
				return provider.PreflightReport{
					Mode: mode,
					Checks: []provider.Check{{
						Name:   "rate-limit",
						Status: provider.CheckWarn,
						Detail: "preflight probe reached its deadline",
					}},
				}
			}
			d := deps{
				provider:  prov,
				newRunner: func(io.Writer) testRunner { return runner },
				list:      stubList([]string{"TestAccThing"}),
				cwd:       func() (string, error) { return root, nil },
			}

			parent, cancel := context.WithTimeout(context.Background(), 250*time.Millisecond)
			defer cancel()
			var out, errOut bytes.Buffer
			args := append(append([]string{}, tc.args...), "--repo-root", root)
			code := runE2EContext(parent, args, &out, &errOut, d)

			if code != 1 {
				t.Fatalf("exit code = %d, want 1; stderr=%s stdout=%s", code, errOut.String(), out.String())
			}
			if len(runner.specs) != 0 || prov.orphanCalls != 0 {
				t.Fatalf("deadline reached downstream work: runner calls=%d orphan calls=%d", len(runner.specs), prov.orphanCalls)
			}
			if !strings.Contains(errOut.String(), "guided setup: context deadline exceeded") {
				t.Fatalf("stderr = %q, want guided setup deadline error", errOut.String())
			}

			if tc.json {
				events := decodeNDJSON(t, out.String())
				if len(events) != 3 {
					t.Fatalf("events = %v, want plan, warning check, and terminal summary", events)
				}
				if events[1]["type"] != "check" || events[1]["name"] != "rate-limit" || events[1]["status"] != "warn" {
					t.Fatalf("preflight classification = %v, want warning check preserved", events[1])
				}
				summary := lastEvent(t, events)
				if summary["type"] != "summary" || summary["command"] != "e2e" ||
					summary["mode"] != "individual" || summary["preflight_ok"] != true ||
					summary["run_attempted"] != false || summary["orphan_check_attempted"] != false ||
					summary["exit_code"] != float64(1) {
					t.Fatalf("terminal summary = %v, want deadline failure without collapsed preflight classification", summary)
				}
				return
			}

			text := out.String()
			if !strings.Contains(text, "[!] rate-limit: preflight probe reached its deadline") {
				t.Fatalf("text output lost warning classification:\n%s", text)
			}
			if !strings.HasSuffix(text, "individual acceptance tests were not started\ndestructive cleanup was not run\n") {
				t.Fatalf("text output missing terminal summary:\n%s", text)
			}
		})
	}
}

func TestE2ESetupDeadlineDuringBaselineEmitsTerminalSummary(t *testing.T) {
	for _, tc := range []struct {
		name string
		args []string
		json bool
	}{
		{name: "text", args: []string{"--mode", "individual"}},
		{name: "JSON", args: []string{"--mode", "individual", "--json"}, json: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := t.TempDir()
			runner := &e2eRunner{}
			prov := newE2EProvider(provider.PreflightReport{Checks: []provider.Check{{
				Name:   "identity",
				Status: provider.CheckOK,
				Detail: "credentials accepted",
			}}})
			prov.ModesVal = []provider.Mode{{Name: "anonymous"}, {Name: "individual"}, {Name: "organization"}}
			prov.orphanFn = func(ctx context.Context, _ string, call int) ([]provider.Resource, error) {
				if call != 1 {
					t.Fatalf("orphan call = %d, want only the baseline", call)
				}
				<-ctx.Done()
				if !errors.Is(ctx.Err(), context.DeadlineExceeded) {
					t.Fatalf("baseline context error = %v, want deadline exceeded", ctx.Err())
				}
				return nil, ctx.Err()
			}
			d := deps{
				provider:  prov,
				newRunner: func(io.Writer) testRunner { return runner },
				list:      stubList([]string{"TestAccThing"}),
				cwd:       func() (string, error) { return root, nil },
			}

			parent, cancel := context.WithTimeout(context.Background(), 250*time.Millisecond)
			defer cancel()
			var out, errOut bytes.Buffer
			args := append(append([]string{}, tc.args...), "--repo-root", root)
			code := runE2EContext(parent, args, &out, &errOut, d)

			if code != 1 {
				t.Fatalf("exit code = %d, want 1; stderr=%s stdout=%s", code, errOut.String(), out.String())
			}
			if len(runner.specs) != 0 || prov.orphanCalls != 1 {
				t.Fatalf("baseline deadline reached downstream work: runner calls=%d orphan calls=%d", len(runner.specs), prov.orphanCalls)
			}
			if !strings.Contains(errOut.String(), "capturing orphan baseline: context deadline exceeded") {
				t.Fatalf("stderr = %q, want baseline deadline error", errOut.String())
			}

			if tc.json {
				events := decodeNDJSON(t, out.String())
				summary := lastEvent(t, events)
				if summary["type"] != "summary" || summary["command"] != "e2e" ||
					summary["mode"] != "individual" || summary["preflight_ok"] != true ||
					summary["run_attempted"] != false || summary["orphan_check_attempted"] != true ||
					summary["orphan_check_exit_code"] != float64(1) ||
					summary["cleanup_status"] != engine.CleanupUnknown ||
					summary["exit_code"] != float64(1) {
					t.Fatalf("terminal summary = %v, want classified baseline deadline failure", summary)
				}
				return
			}

			text := out.String()
			if !strings.Contains(text, "[✓] identity: credentials accepted") {
				t.Fatalf("text output lost successful preflight classification:\n%s", text)
			}
			if !strings.HasSuffix(text, "individual acceptance tests were not started\ndestructive cleanup was not run\n") {
				t.Fatalf("text output missing terminal summary:\n%s", text)
			}
		})
	}
}

func TestE2ESetupUsesNamedSixtySecondDeadline(t *testing.T) {
	root := t.TempDir()
	var remaining time.Duration
	prov := newE2EProvider(provider.PreflightReport{})
	prov.ModesVal = []provider.Mode{{Name: "anonymous"}, {Name: "individual"}, {Name: "organization"}}
	prov.PreflightFn = func(ctx context.Context, mode string, _ provider.TestRequirements) provider.PreflightReport {
		deadline, ok := ctx.Deadline()
		if !ok {
			t.Fatal("guided setup context has no deadline")
		}
		remaining = time.Until(deadline)
		return provider.PreflightReport{
			Mode: mode,
			Checks: []provider.Check{{
				Name:   "stop",
				Status: provider.CheckFail,
				Detail: "stop after observing setup deadline",
			}},
		}
	}
	d := deps{
		provider: prov,
		list:     stubList([]string{"TestAccThing"}),
		cwd:      func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runE2EContext(context.Background(),
		[]string{"--mode", "individual", "--json", "--repo-root", root}, &out, &errOut, d)

	if code != 1 {
		t.Fatalf("exit code = %d, want failed preflight; stderr=%s stdout=%s", code, errOut.String(), out.String())
	}
	if guidedE2ESetupTimeout != 60*time.Second {
		t.Fatalf("guided setup timeout = %s, want 60s", guidedE2ESetupTimeout)
	}
	if remaining <= guidedE2ESetupTimeout-time.Second || remaining > guidedE2ESetupTimeout {
		t.Fatalf("guided setup deadline remaining = %s, want about %s", remaining, guidedE2ESetupTimeout)
	}
}

func TestE2EPlanningRespectsParentCancellation(t *testing.T) {
	root := t.TempDir()
	prov := newE2EProvider(provider.PreflightReport{})
	prov.PreflightFn = func(context.Context, string, provider.TestRequirements) provider.PreflightReport {
		t.Fatal("preflight called after planning cancellation")
		return provider.PreflightReport{}
	}
	var listErr error
	d := deps{
		provider: prov,
		list: func(ctx context.Context, _ string, _ []string, _ string) ([]string, error) {
			<-ctx.Done()
			listErr = ctx.Err()
			return nil, ctx.Err()
		},
		cwd: func() (string, error) { return root, nil },
	}
	parent, cancel := context.WithCancel(context.Background())
	cancel()

	var out, errOut bytes.Buffer
	code := runE2EContext(parent,
		[]string{"--mode", "individual", "--json", "--repo-root", root}, &out, &errOut, d)

	if code != 1 {
		t.Fatalf("exit code = %d, want 1; stderr=%s stdout=%s", code, errOut.String(), out.String())
	}
	if !errors.Is(listErr, context.Canceled) {
		t.Fatalf("list context error = %v, want parent cancellation", listErr)
	}
	events := decodeNDJSON(t, out.String())
	if got := len(events); got != 1 {
		t.Fatalf("events = %v, want only terminal summary after canceled planning", events)
	}
	summary := lastEvent(t, events)
	if summary["type"] != "summary" || summary["command"] != "e2e" ||
		summary["run_attempted"] != false || summary["exit_code"] != float64(1) {
		t.Fatalf("terminal summary = %v, want canceled planning failure", summary)
	}
}

func TestRunE2EHandlesSIGTERM(t *testing.T) {
	const helperEnv = "TPT_TEST_E2E_SIGTERM_HELPER"
	if os.Getenv(helperEnv) == "1" {
		root := t.TempDir()
		prov := newE2EProvider(provider.PreflightReport{})
		d := deps{
			provider: prov,
			list: func(ctx context.Context, _ string, _ []string, _ string) ([]string, error) {
				fmt.Fprintln(os.Stdout, "e2e-sigterm-ready")
				<-ctx.Done()
				return nil, ctx.Err()
			},
			cwd: func() (string, error) { return root, nil },
		}
		code := runE2E(
			[]string{"--mode", "anonymous", "--json", "--repo-root", root},
			os.Stdout,
			os.Stderr,
			d,
		)
		fmt.Fprintf(os.Stdout, "e2e-sigterm-exit=%d\n", code)
		return
	}

	childOutput, stderr, waitErr := runE2ESignalSequence(
		t,
		"TestRunE2EHandlesSIGTERM",
		helperEnv,
		"e2e-sigterm-ready",
	)
	if waitErr != nil {
		t.Fatalf("runE2E helper did not handle SIGTERM: %v; stderr=%s stdout=%s", waitErr, stderr, childOutput)
	}
	if !strings.Contains(childOutput, "e2e-sigterm-exit=1") {
		t.Fatalf("helper output = %q, want cancellation exit 1; stderr=%s", childOutput, stderr)
	}
	if !strings.Contains(stderr, "planning:") || !strings.Contains(stderr, "context canceled") {
		t.Fatalf("helper stderr = %q, want planning cancellation", stderr)
	}
}

func TestRunE2ESecondSIGTERMCancelsFinalAccounting(t *testing.T) {
	const helperEnv = "TPT_TEST_E2E_FINAL_SIGTERM_HELPER"
	if os.Getenv(helperEnv) == "1" {
		root := t.TempDir()
		prov := newE2EProvider(provider.PreflightReport{})
		prov.orphanFn = func(ctx context.Context, _ string, call int) ([]provider.Resource, error) {
			if call == 1 {
				return nil, nil
			}
			fmt.Fprintln(os.Stdout, "e2e-final-signal-ready")
			<-ctx.Done()
			return nil, ctx.Err()
		}
		d := deps{
			provider:  prov,
			newRunner: func(io.Writer) testRunner { return signalBlockingE2ERunner{} },
			list:      stubList([]string{"TestAccThing"}),
			cwd:       func() (string, error) { return root, nil },
		}
		code := runE2E(
			[]string{"--mode", "organization", "--json", "--repo-root", root},
			os.Stdout,
			os.Stderr,
			d,
		)
		fmt.Fprintf(os.Stdout, "e2e-final-sigterm-exit=%d\n", code)
		return
	}

	childOutput, stderr, waitErr := runE2ESignalSequence(
		t,
		"TestRunE2ESecondSIGTERMCancelsFinalAccounting",
		helperEnv,
		"e2e-run-signal-ready",
		"e2e-final-signal-ready",
	)
	if waitErr != nil {
		t.Fatalf("runE2E helper failed: %v; stderr=%s stdout=%s", waitErr, stderr, childOutput)
	}
	if !strings.Contains(childOutput, "e2e-final-sigterm-exit=1") {
		t.Fatalf("helper output = %q, want cancellation exit 1; stderr=%s", childOutput, stderr)
	}
	if !strings.Contains(stderr, "capturing final orphan state: context canceled") {
		t.Fatalf("helper stderr = %q, want prompt final-accounting cancellation", stderr)
	}
}

func TestE2EPersistsEmptyCapturedBaseline(t *testing.T) {
	root := t.TempDir()
	runner := &e2eRunner{results: []engine.RunResult{{Tests: []engine.TestResult{{
		Package: "./github", Name: "TestAccThing", Status: provider.StatusPass,
	}}}}}
	prov := newE2EProvider(provider.PreflightReport{})
	prov.orphans = []provider.Resource{}
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccThing"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"e2e", "--mode", "organization", "--json", "--repo-root", root}, &out, &errOut, d)

	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s stdout=%s", code, errOut.String(), out.String())
	}
	state, err := engine.Load(filepath.Join(root, ".pulsar-state.json"))
	if err != nil {
		t.Fatal(err)
	}
	if state.Orphans == nil || !state.Orphans.BaselineCaptured || !state.Orphans.FinalCaptured ||
		state.Orphans.CleanupStatus != engine.CleanupComplete ||
		state.Orphans.Baseline == nil || len(state.Orphans.Baseline) != 0 ||
		state.Orphans.Final == nil || len(state.Orphans.Final) != 0 {
		t.Fatalf("orphan accounting = %+v, want captured empty baseline and final", state.Orphans)
	}
}

func TestE2EPassingRunCapturesBaselineAndFinal(t *testing.T) {
	root := t.TempDir()
	baseline := []provider.Resource{{Kind: "repository", Name: "tf-acc-test-existing", URL: "https://example.test/old"}}
	final := []provider.Resource{
		{Kind: "repository", Name: "tf-acc-test-existing", URL: "https://example.test/current"},
		{Kind: "issue_label", Name: "tf-acc-test-new", URL: "https://example.test/new"},
	}
	runner := &e2eRunner{results: []engine.RunResult{{Tests: []engine.TestResult{{
		Package: "./github", Name: "TestAccThing", Status: provider.StatusPass,
	}}}}}
	prov := newE2EProvider(provider.PreflightReport{})
	prov.orphanSeq = [][]provider.Resource{baseline, final}
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccThing"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"e2e", "--mode", "organization", "--json", "--repo-root", root}, &out, &errOut, d)

	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s stdout=%s", code, errOut.String(), out.String())
	}
	if prov.orphanCalls != 2 {
		t.Fatalf("orphan calls = %d, want baseline plus final", prov.orphanCalls)
	}
	state, err := engine.Load(filepath.Join(root, ".pulsar-state.json"))
	if err != nil {
		t.Fatal(err)
	}
	wantPreExisting := []provider.Resource{{Kind: "repository", Name: "tf-acc-test-existing", URL: "https://example.test/current"}}
	wantNew := []provider.Resource{{Kind: "issue_label", Name: "tf-acc-test-new", URL: "https://example.test/new"}}
	if state.Orphans == nil || !state.Orphans.BaselineCaptured || !state.Orphans.FinalCaptured ||
		state.Orphans.CleanupStatus != engine.CleanupComplete ||
		!reflect.DeepEqual(state.Orphans.Baseline, baseline) ||
		!reflect.DeepEqual(state.Orphans.Final, final) ||
		!reflect.DeepEqual(state.Orphans.PreExisting, wantPreExisting) ||
		!reflect.DeepEqual(state.Orphans.New, wantNew) {
		t.Fatalf("orphan accounting = %+v, want baseline/final delta", state.Orphans)
	}
}

func TestE2EFailedRunCapturesBaselineAndFinal(t *testing.T) {
	root := t.TempDir()
	baseline := []provider.Resource{{Kind: "repository", Name: "tf-acc-test-existing", URL: "https://example.test/existing"}}
	final := append([]provider.Resource{}, baseline...)
	final = append(final, provider.Resource{Kind: "repository", Name: "tf-acc-test-leaked", URL: "https://example.test/leaked"})
	runner := &e2eRunner{results: []engine.RunResult{{Tests: []engine.TestResult{{
		Package: "./github", Name: "TestAccBroken", Status: provider.StatusFail, Output: []string{"assertion failed\n"},
	}}}}}
	prov := newE2EProvider(provider.PreflightReport{})
	prov.orphanSeq = [][]provider.Resource{baseline, final}
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccBroken"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"e2e", "--mode", "organization", "--json", "--repo-root", root}, &out, &errOut, d)

	if code != 1 {
		t.Fatalf("exit code = %d, want 1; stderr=%s stdout=%s", code, errOut.String(), out.String())
	}
	if prov.orphanCalls != 2 {
		t.Fatalf("orphan calls = %d, want baseline plus final after failed run", prov.orphanCalls)
	}
	state, err := engine.Load(filepath.Join(root, ".pulsar-state.json"))
	if err != nil {
		t.Fatal(err)
	}
	if state.Orphans == nil || !state.Orphans.FinalCaptured || state.Orphans.CleanupStatus != engine.CleanupComplete ||
		!reflect.DeepEqual(state.Orphans.Baseline, baseline) || !reflect.DeepEqual(state.Orphans.Final, final) ||
		!reflect.DeepEqual(state.Orphans.New, []provider.Resource{{Kind: "repository", Name: "tf-acc-test-leaked", URL: "https://example.test/leaked"}}) {
		t.Fatalf("orphan accounting = %+v, want final delta after failed run", state.Orphans)
	}
}

func TestE2EInternalRetryStillListsOrphansTwice(t *testing.T) {
	root := t.TempDir()
	runner := &e2eRunner{results: []engine.RunResult{
		{Tests: []engine.TestResult{{Package: "./github", Name: "TestAccRateLimited", Status: provider.StatusFail, Output: []string{"403 API rate limit exceeded\n"}}}},
		{Tests: []engine.TestResult{{Package: "./github", Name: "TestAccRateLimited", Status: provider.StatusPass}}},
	}}
	prov := newE2EProvider(provider.PreflightReport{})
	prov.orphanSeq = [][]provider.Resource{nil, {{Kind: "repository", Name: "tf-acc-test-new", URL: "https://example.test/new"}}}
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccRateLimited"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"e2e", "--mode", "organization", "--json", "--repo-root", root}, &out, &errOut, d)

	if code != 0 {
		t.Fatalf("exit code = %d, want 0 after retry; stderr=%s stdout=%s", code, errOut.String(), out.String())
	}
	if len(runner.specs) != 2 {
		t.Fatalf("runner calls = %d, want initial plus retry", len(runner.specs))
	}
	if prov.orphanCalls != 2 {
		t.Fatalf("orphan calls = %d, want exactly baseline and final regardless of retry", prov.orphanCalls)
	}
}

func TestE2EStartFailureStillCapturesFinal(t *testing.T) {
	root := t.TempDir()
	baseline := []provider.Resource{{Kind: "repository", Name: "tf-acc-test-existing", URL: "https://example.test/existing"}}
	final := append([]provider.Resource{}, baseline...)
	runner := &e2eRunner{
		results: []engine.RunResult{{}},
		errs:    []error{errors.New("fork/exec go: no such file or directory")},
	}
	prov := newE2EProvider(provider.PreflightReport{})
	prov.orphanSeq = [][]provider.Resource{baseline, final}
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccThing"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"e2e", "--mode", "organization", "--json", "--repo-root", root}, &out, &errOut, d)

	if code != 1 {
		t.Fatalf("exit code = %d, want 1 for runner start error; stderr=%s stdout=%s", code, errOut.String(), out.String())
	}
	if prov.orphanCalls != 2 {
		t.Fatalf("orphan calls = %d, want baseline plus final after start error", prov.orphanCalls)
	}
	state, err := engine.Load(filepath.Join(root, ".pulsar-state.json"))
	if err != nil {
		t.Fatal(err)
	}
	if state.Orphans == nil || !state.Orphans.FinalCaptured || state.Orphans.CleanupStatus != engine.CleanupComplete ||
		!reflect.DeepEqual(state.Orphans.Baseline, baseline) || !reflect.DeepEqual(state.Orphans.Final, final) {
		t.Fatalf("orphan accounting = %+v, want final after start error", state.Orphans)
	}
}

func TestE2EInterruptionStillCapturesFinal(t *testing.T) {
	root := t.TempDir()
	final := []provider.Resource{{Kind: "repository", Name: "tf-acc-test-interrupted", URL: "https://example.test/interrupted"}}
	runner := &e2eRunner{
		results: []engine.RunResult{{}},
		errs:    []error{context.Canceled},
	}
	prov := newE2EProvider(provider.PreflightReport{})
	prov.orphanSeq = [][]provider.Resource{nil, final}
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccThing"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"e2e", "--mode", "organization", "--json", "--repo-root", root}, &out, &errOut, d)

	if code != 1 {
		t.Fatalf("exit code = %d, want 1 for interruption; stderr=%s stdout=%s", code, errOut.String(), out.String())
	}
	if prov.orphanCalls != 2 {
		t.Fatalf("orphan calls = %d, want baseline plus final after interruption", prov.orphanCalls)
	}
	state, err := engine.Load(filepath.Join(root, ".pulsar-state.json"))
	if err != nil {
		t.Fatal(err)
	}
	if state.Orphans == nil || !state.Orphans.FinalCaptured || state.Orphans.CleanupStatus != engine.CleanupComplete ||
		!reflect.DeepEqual(state.Orphans.Final, final) || !reflect.DeepEqual(state.Orphans.New, final) {
		t.Fatalf("orphan accounting = %+v, want final after interruption", state.Orphans)
	}
}

func TestE2EFinalAccountingUsesFreshContextAfterRunCancellation(t *testing.T) {
	root := t.TempDir()
	parent, cancelRun := context.WithCancel(context.Background())
	runner := &e2eRunner{
		results: []engine.RunResult{{}},
		errs:    []error{context.Canceled},
		onRun: func(engine.RunSpec) {
			cancelRun()
		},
	}
	final := []provider.Resource{{
		Kind: "repository", Name: "tf-acc-test-after-cancel", URL: "https://example.test/after-cancel",
	}}
	prov := newE2EProvider(provider.PreflightReport{})
	prov.orphanFn = func(ctx context.Context, _ string, call int) ([]provider.Resource, error) {
		if call == 1 {
			return nil, nil
		}
		if err := ctx.Err(); err != nil {
			return nil, fmt.Errorf("final context inherited run cancellation: %w", err)
		}
		return final, nil
	}
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccThing"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runE2EContext(
		parent,
		[]string{"--mode", "organization", "--json", "--repo-root", root},
		&out,
		&errOut,
		d,
	)

	if code != 1 {
		t.Fatalf("exit code = %d, want interrupted run exit 1; stderr=%s stdout=%s", code, errOut.String(), out.String())
	}
	if prov.orphanCalls != 2 {
		t.Fatalf("orphan calls = %d, want baseline plus final after run cancellation", prov.orphanCalls)
	}
	state, err := engine.Load(filepath.Join(root, ".pulsar-state.json"))
	if err != nil {
		t.Fatal(err)
	}
	if state.Orphans == nil || !state.Orphans.FinalCaptured ||
		!reflect.DeepEqual(state.Orphans.Final, final) {
		t.Fatalf("orphan accounting = %+v, want final scan after run cancellation", state.Orphans)
	}
}

func TestE2EFinalFailureMarksCleanupUnknown(t *testing.T) {
	root := t.TempDir()
	baseline := []provider.Resource{{Kind: "repository", Name: "tf-acc-test-existing", URL: "https://example.test/existing"}}
	runner := &e2eRunner{results: []engine.RunResult{{Tests: []engine.TestResult{{
		Package: "./github", Name: "TestAccThing", Status: provider.StatusPass,
	}}}}}
	prov := newE2EProvider(provider.PreflightReport{})
	prov.orphanSeq = [][]provider.Resource{baseline}
	prov.orphanErrs = []error{nil, errors.New("final orphan scan failed")}
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccThing"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"e2e", "--mode", "organization", "--json", "--repo-root", root}, &out, &errOut, d)

	if code != 1 {
		t.Fatalf("exit code = %d, want 1 when final orphan scan fails; stderr=%s stdout=%s", code, errOut.String(), out.String())
	}
	if prov.orphanCalls != 2 {
		t.Fatalf("orphan calls = %d, want baseline plus failed final", prov.orphanCalls)
	}
	state, err := engine.Load(filepath.Join(root, ".pulsar-state.json"))
	if err != nil {
		t.Fatal(err)
	}
	if state.Orphans == nil || !state.Orphans.BaselineCaptured || state.Orphans.FinalCaptured ||
		state.Orphans.CleanupStatus != engine.CleanupUnknown ||
		!reflect.DeepEqual(state.Orphans.Baseline, baseline) {
		t.Fatalf("orphan accounting = %+v, want baseline intact and cleanup unknown", state.Orphans)
	}
}

func TestFinalOrphanCheckFailureInvalidatesStaleDelta(t *testing.T) {
	baseline := []provider.Resource{{
		Kind: "repository", Name: "tf-acc-test-existing", URL: "https://example.test/existing",
	}}
	state := &engine.State{Orphans: &engine.OrphanAccounting{
		Mode:             "organization",
		Baseline:         baseline,
		Final:            []provider.Resource{{Kind: "repository", Name: "tf-acc-test-stale-final"}},
		PreExisting:      []provider.Resource{{Kind: "repository", Name: "tf-acc-test-stale-existing"}},
		New:              []provider.Resource{{Kind: "repository", Name: "tf-acc-test-stale-new"}},
		BaselineCaptured: true,
		FinalCaptured:    true,
		CleanupStatus:    engine.CleanupComplete,
	}}
	prov := newE2EProvider(provider.PreflightReport{})
	prov.orphanErr = errors.New("final orphan scan failed")

	err := finalizeOrphanAccountingWithRedactor(
		context.Background(), prov, "organization", state, redact.New(nil),
	)

	if err == nil {
		t.Fatal("finalizeOrphanAccountingWithRedactor error = nil, want final scan failure")
	}
	if prov.orphanCalls != 1 {
		t.Fatalf("orphan calls = %d, want 1 failed final scan", prov.orphanCalls)
	}
	got := state.Orphans
	if got == nil || got.Mode != "organization" || !got.BaselineCaptured ||
		!reflect.DeepEqual(got.Baseline, baseline) {
		t.Fatalf("baseline accounting = %+v, want original mode and captured baseline %+v", got, baseline)
	}
	if got.Final != nil || got.PreExisting != nil || got.New != nil || got.FinalCaptured {
		t.Fatalf("stale final accounting survived failed scan: %+v", got)
	}
	if got.CleanupStatus != engine.CleanupUnknown {
		t.Fatalf("cleanup status = %q, want unknown", got.CleanupStatus)
	}
}

func TestFinalOrphanCheckRequiresCapturedBaseline(t *testing.T) {
	baseline := []provider.Resource{{
		Kind: "repository", Name: "tf-acc-test-existing", URL: "https://example.test/existing",
	}}
	for _, tc := range []struct {
		name       string
		accounting *engine.OrphanAccounting
	}{
		{name: "nil accounting"},
		{
			name: "uncaptured accounting",
			accounting: &engine.OrphanAccounting{
				Mode:        "organization",
				Baseline:    baseline,
				Final:       []provider.Resource{{Kind: "repository", Name: "tf-acc-test-stale-final"}},
				PreExisting: []provider.Resource{{Kind: "repository", Name: "tf-acc-test-stale-existing"}},
				New:         []provider.Resource{{Kind: "repository", Name: "tf-acc-test-stale-new"}},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			state := &engine.State{Orphans: tc.accounting}
			prov := newE2EProvider(provider.PreflightReport{})
			prov.orphans = []provider.Resource{{Kind: "repository", Name: "tf-acc-test-would-be-new"}}

			err := finalizeOrphanAccountingWithRedactor(
				context.Background(), prov, "organization", state, redact.New(nil),
			)

			if err == nil || !strings.Contains(err.Error(), "captured orphan baseline") {
				t.Fatalf("error = %v, want captured orphan baseline failure", err)
			}
			if prov.orphanCalls != 0 {
				t.Fatalf("orphan calls = %d, want 0 without a captured baseline", prov.orphanCalls)
			}
			if state.Orphans == nil || state.Orphans.CleanupStatus != engine.CleanupUnknown {
				t.Fatalf("orphan accounting = %+v, want cleanup status unknown", state.Orphans)
			}
			if state.Orphans.Final != nil || state.Orphans.PreExisting != nil ||
				state.Orphans.New != nil || state.Orphans.FinalCaptured {
				t.Fatalf("uncaptured baseline retained final accounting: %+v", state.Orphans)
			}
			if tc.accounting != nil &&
				(state.Orphans.Mode != "organization" ||
					!reflect.DeepEqual(state.Orphans.Baseline, baseline)) {
				t.Fatalf("orphan accounting = %+v, want original mode and baseline %+v", state.Orphans, baseline)
			}
		})
	}
}

func TestE2EFinalStateSaveFailureForcesUnknownNonzero(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")
	baseline := []provider.Resource{{
		Kind: "repository", Name: "tf-acc-test-existing", URL: "https://example.test/existing",
	}}
	runner := &e2eRunner{results: []engine.RunResult{{Tests: []engine.TestResult{{
		Package: "./github", Name: "TestAccThing", Status: provider.StatusPass,
	}}}}}
	prov := newE2EProvider(provider.PreflightReport{})
	prov.orphanFn = func(_ context.Context, _ string, call int) ([]provider.Resource, error) {
		if call == 1 {
			return baseline, nil
		}
		if err := os.Remove(statePath); err != nil {
			t.Fatalf("removing state before final save: %v", err)
		}
		if err := os.Mkdir(statePath, 0o700); err != nil {
			t.Fatalf("blocking final state save: %v", err)
		}
		return baseline, nil
	}
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccThing"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"e2e", "--mode", "organization", "--json", "--repo-root", root}, &out, &errOut, d)

	if code != 1 {
		t.Fatalf("exit code = %d, want 1 when final state save fails; stderr=%s stdout=%s", code, errOut.String(), out.String())
	}
	if !strings.Contains(errOut.String(), "saving state:") {
		t.Fatalf("stderr = %q, want final state save error", errOut.String())
	}
	events := decodeNDJSON(t, out.String())
	delta := events[len(events)-2]
	summary := lastEvent(t, events)
	if delta["type"] != "orphan_delta" || delta["cleanup_status"] != engine.CleanupUnknown {
		t.Fatalf("penultimate event = %v, want unknown orphan_delta", delta)
	}
	if summary["command"] != "e2e" || summary["cleanup_status"] != engine.CleanupUnknown ||
		summary["exit_code"] != float64(1) {
		t.Fatalf("summary = %v, want unknown cleanup and exit 1", summary)
	}
}

func TestE2EPreExistingResourcesAreNotNew(t *testing.T) {
	root := t.TempDir()
	baseline := []provider.Resource{{Kind: "repository", Name: "tf-acc-test-existing", URL: "https://example.test/baseline"}}
	final := []provider.Resource{
		{Kind: "repository", Name: "tf-acc-test-existing", URL: "https://example.test/final"},
		{Kind: "repository", Name: "tf-acc-test-new", URL: "https://example.test/new"},
	}
	runner := &e2eRunner{results: []engine.RunResult{{Tests: []engine.TestResult{{
		Package: "./github", Name: "TestAccThing", Status: provider.StatusPass,
	}}}}}
	prov := newE2EProvider(provider.PreflightReport{})
	prov.orphanSeq = [][]provider.Resource{baseline, final}
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccThing"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"e2e", "--mode", "organization", "--json", "--repo-root", root}, &out, &errOut, d)

	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s stdout=%s", code, errOut.String(), out.String())
	}
	state, err := engine.Load(filepath.Join(root, ".pulsar-state.json"))
	if err != nil {
		t.Fatal(err)
	}
	wantPreExisting := []provider.Resource{{Kind: "repository", Name: "tf-acc-test-existing", URL: "https://example.test/final"}}
	wantNew := []provider.Resource{{Kind: "repository", Name: "tf-acc-test-new", URL: "https://example.test/new"}}
	if state.Orphans == nil || !reflect.DeepEqual(state.Orphans.PreExisting, wantPreExisting) ||
		!reflect.DeepEqual(state.Orphans.New, wantNew) {
		t.Fatalf("delta = pre-existing %+v new %+v, want pre-existing %+v new %+v",
			state.Orphans.PreExisting, state.Orphans.New, wantPreExisting, wantNew)
	}
}

func TestFinalOrphanCheckUsesThirtySecondDeadline(t *testing.T) {
	parent := context.Background()
	state := &engine.State{Orphans: &engine.OrphanAccounting{
		Mode:             "organization",
		Baseline:         []provider.Resource{{Kind: "repository", Name: "tf-acc-test-existing", URL: "https://example.test/existing"}},
		BaselineCaptured: true,
		CleanupStatus:    engine.CleanupBaselineOnly,
	}}
	observed := false
	prov := newE2EProvider(provider.PreflightReport{})
	prov.orphanFn = func(ctx context.Context, mode string, call int) ([]provider.Resource, error) {
		if mode != "organization" {
			t.Fatalf("mode = %q, want organization", mode)
		}
		if call != 1 {
			t.Fatalf("orphan calls = %d, want exactly one final call", call)
		}
		if err := ctx.Err(); err != nil {
			t.Fatalf("final context error = %v, want uncanceled context despite canceled parent", err)
		}
		deadline, ok := ctx.Deadline()
		if !ok {
			t.Fatal("final context has no deadline")
		}
		remaining := time.Until(deadline)
		if remaining <= 29*time.Second || remaining > finalOrphanCheckTimeout {
			t.Fatalf("final deadline remaining = %s, want about 30s", remaining)
		}
		observed = true
		return state.Orphans.Baseline, nil
	}

	if err := finalizeOrphanAccountingWithRedactor(parent, prov, "organization", state, redact.New(nil)); err != nil {
		t.Fatalf("finalizeOrphanAccountingWithRedactor: %v", err)
	}
	if !observed {
		t.Fatal("final orphan check was not observed")
	}
	if !state.Orphans.FinalCaptured || state.Orphans.CleanupStatus != engine.CleanupComplete {
		t.Fatalf("orphan accounting = %+v, want successful finalization", state.Orphans)
	}
}

func TestFinalOrphanCheckRespectsFreshParentCancellation(t *testing.T) {
	parent, cancel := context.WithCancel(context.Background())
	state := &engine.State{Orphans: &engine.OrphanAccounting{
		Mode:             "organization",
		Baseline:         []provider.Resource{},
		BaselineCaptured: true,
		CleanupStatus:    engine.CleanupBaselineOnly,
	}}
	started := make(chan struct{})
	release := make(chan struct{})
	prov := newE2EProvider(provider.PreflightReport{})
	prov.orphanFn = func(ctx context.Context, _ string, _ int) ([]provider.Resource, error) {
		close(started)
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-release:
			return nil, errors.New("test released an uncanceled final orphan check")
		}
	}
	result := make(chan error, 1)
	go func() {
		result <- finalizeOrphanAccountingWithRedactor(parent, prov, "organization", state, redact.New(nil))
	}()

	<-started
	cancel()
	select {
	case err := <-result:
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("final orphan check error = %v, want context cancellation", err)
		}
	case <-time.After(500 * time.Millisecond):
		close(release)
		err := <-result
		t.Fatalf("fresh parent cancellation did not reach provider.Orphans promptly; eventual error=%v", err)
	}
}

func TestE2ENeverSweepsDuringFinalization(t *testing.T) {
	root := t.TempDir()
	runner := &e2eRunner{results: []engine.RunResult{{Tests: []engine.TestResult{{
		Package: "./github", Name: "TestAccThing", Status: provider.StatusPass,
	}}}}}
	prov := newE2EProvider(provider.PreflightReport{})
	prov.orphanSeq = [][]provider.Resource{
		{{Kind: "repository", Name: "tf-acc-test-existing", URL: "https://example.test/existing"}},
		{{Kind: "repository", Name: "tf-acc-test-existing", URL: "https://example.test/existing"}},
	}
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccThing"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"e2e", "--mode", "organization", "--json", "--repo-root", root}, &out, &errOut, d)

	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s stdout=%s", code, errOut.String(), out.String())
	}
	if prov.orphanCalls != 2 {
		t.Fatalf("orphan calls = %d, want baseline plus final", prov.orphanCalls)
	}
	if prov.sweepCalls != 0 {
		t.Fatalf("sweep calls = %d, want 0", prov.sweepCalls)
	}
}

func TestE2ERetryDoesNotRecaptureBaseline(t *testing.T) {
	root := t.TempDir()
	var calls []string
	runner := &e2eRunner{
		results: []engine.RunResult{
			{Tests: []engine.TestResult{{
				Package: "./github", Name: "TestAccRateLimited", Status: provider.StatusFail, Output: []string{"403 API rate limit exceeded\n"},
			}}},
			{Tests: []engine.TestResult{{
				Package: "./github", Name: "TestAccRateLimited", Status: provider.StatusPass,
			}}},
		},
		onRun: func(engine.RunSpec) { calls = append(calls, "runner") },
	}
	prov := newE2EProvider(provider.PreflightReport{})
	prov.callLog = &calls
	prov.orphans = []provider.Resource{{Kind: "repository", Name: "tf-acc-test-existing", URL: "https://example.test/existing"}}
	prov.PreflightFn = func(_ context.Context, mode string, _ provider.TestRequirements) provider.PreflightReport {
		calls = append(calls, "preflight")
		return provider.PreflightReport{Mode: mode}
	}
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccRateLimited"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"e2e", "--mode", "organization", "--json", "--repo-root", root}, &out, &errOut, d)

	if code != 0 {
		t.Fatalf("exit code = %d, want 0 after retry recovery; stderr=%s stdout=%s", code, errOut.String(), out.String())
	}
	if prov.orphanCalls != 2 {
		t.Fatalf("orphan calls = %d, want baseline plus final across initial run and retry", prov.orphanCalls)
	}
	if !reflect.DeepEqual(calls, []string{"preflight", "baseline", "runner", "runner", "final"}) {
		t.Fatalf("call order = %v, want baseline, initial run, retry, final", calls)
	}
}

func TestE2ETTYNestedRunSkipsDashboard(t *testing.T) {
	root := t.TempDir()
	runner := &e2eRunner{results: []engine.RunResult{{Tests: []engine.TestResult{{
		Package: "./github", Name: "TestAccThing", Status: provider.StatusPass,
	}}}}}
	prov := newE2EProvider(provider.PreflightReport{})
	listCalls := 0
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list: func(context.Context, string, []string, string) ([]string, error) {
			listCalls++
			if listCalls == 1 {
				return []string{"TestAccThing"}, nil
			}
			return nil, errors.New("nested guided run unexpectedly launched dashboard discovery")
		},
		cwd:   func() (string, error) { return root, nil },
		isTTY: func() bool { return true },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"e2e", "--mode", "organization", "--repo-root", root}, &out, &errOut, d)

	if code != 0 {
		t.Fatalf("exit code = %d, want 0 when nested run stays in batch mode; stderr=%s stdout=%s", code, errOut.String(), out.String())
	}
	if listCalls != 1 {
		t.Fatalf("list calls = %d, want 1 planning call with no dashboard rediscovery", listCalls)
	}
	if len(runner.specs) != 1 {
		t.Fatalf("runner calls = %d, want 1", len(runner.specs))
	}
}

func TestRunContextWithHeldLockDoesNotReacquire(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")
	release, err := engine.AcquireLock(statePath + ".lock")
	if err != nil {
		t.Fatalf("AcquireLock: %v", err)
	}
	defer func() {
		if err := release(); err != nil {
			t.Fatalf("release lock: %v", err)
		}
	}()
	runner := &e2eRunner{results: []engine.RunResult{{Tests: []engine.TestResult{{
		Package: "./github", Name: "TestAccThing", Status: provider.StatusPass,
	}}}}}
	prov := newE2EProvider(provider.PreflightReport{})
	plan := &engine.ExecutionPlan{
		Mode:     "organization",
		Selected: []string{"TestAccThing"},
		Eligible: []string{"TestAccThing"},
	}
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list: func(context.Context, string, []string, string) ([]string, error) {
			t.Fatal("runRunContext should reuse the supplied plan instead of rediscovering")
			return nil, nil
		},
		cwd: func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runRunContext(context.Background(), []string{"--mode", "organization", "--json", "--repo-root", root}, &out, &errOut, d, runContextOptions{
		Preflight:     false,
		PreserveState: false,
		LockHeld:      true,
		Plan:          plan,
	})

	if code != 0 {
		t.Fatalf("exit code = %d, want 0 with caller-held lock; stderr=%s stdout=%s", code, errOut.String(), out.String())
	}
	if len(runner.specs) != 1 {
		t.Fatalf("runner calls = %d, want 1", len(runner.specs))
	}
	state, err := engine.Load(statePath)
	if err != nil {
		t.Fatal(err)
	}
	if state.Plan == nil || !reflect.DeepEqual(state.Plan.Eligible, []string{"TestAccThing"}) {
		t.Fatalf("persisted plan = %+v, want supplied plan", state.Plan)
	}
}

func TestNestedGuidedRunPreservesCapturedOrphanBaseline(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")
	baseline := []provider.Resource{{
		Kind: "repository", Name: "tf-acc-test-existing", URL: "https://example.test/existing",
	}}
	plan := &engine.ExecutionPlan{
		Mode:     "organization",
		Selected: []string{"TestAccThing"},
		Eligible: []string{"TestAccThing"},
	}
	previous := engine.State{
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
	if err := previous.Save(statePath); err != nil {
		t.Fatalf("seeding guided baseline: %v", err)
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

	runner := &e2eRunner{results: []engine.RunResult{{Tests: []engine.TestResult{{
		Package: "./github", Name: "TestAccThing", Status: provider.StatusPass,
	}}}}}
	prov := newE2EProvider(provider.PreflightReport{})
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list: func(context.Context, string, []string, string) ([]string, error) {
			t.Fatal("nested guided run should reuse the supplied plan")
			return nil, nil
		},
		cwd: func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runRunContext(context.Background(),
		[]string{"--mode", "organization", "--json", "--repo-root", root},
		&out, &errOut, d, runContextOptions{
			Preflight:     false,
			PreserveState: false,
			LockHeld:      true,
			Plan:          plan,
		})

	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s stdout=%s", code, errOut.String(), out.String())
	}
	persisted, err := engine.Load(statePath)
	if err != nil {
		t.Fatal(err)
	}
	if persisted.Orphans == nil || persisted.Orphans.Mode != "organization" ||
		!persisted.Orphans.BaselineCaptured ||
		!reflect.DeepEqual(persisted.Orphans.Baseline, baseline) ||
		persisted.Orphans.CleanupStatus != engine.CleanupBaselineOnly {
		t.Fatalf("persisted orphan accounting = %+v, want captured guided baseline %+v",
			persisted.Orphans, baseline)
	}
}

func TestE2EStopsOnFailedOrganizationPreflight(t *testing.T) {
	root := t.TempDir()
	runner := &e2eRunner{}
	prov := newE2EProvider(provider.PreflightReport{Checks: []provider.Check{{
		Name:   "owner",
		Status: provider.CheckFail,
		Detail: "GITHUB_OWNER environment variable not set",
	}}})
	listCalls := 0
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list: func(context.Context, string, []string, string) ([]string, error) {
			listCalls++
			return []string{"TestAccThing"}, nil
		},
		cwd: func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"e2e", "--json", "--repo-root", root}, &out, &errOut, d)

	if code != 1 {
		t.Fatalf("exit code = %d, want 1; stderr=%s", code, errOut.String())
	}
	// e2e's own standalone preflight step now discovers tests before checking
	// (resolve root -> discover/plan -> preflight), so exactly one list call
	// comes from that plan; the run step itself is never reached, so the
	// runner performs no test setup.
	if listCalls != 1 || len(runner.specs) != 0 {
		t.Fatalf("expected discovery-before-preflight and no run attempt: list calls=%d runner calls=%d", listCalls, len(runner.specs))
	}
	if prov.orphanCalls != 0 {
		t.Fatalf("orphan calls = %d, want 0 before any attempted run", prov.orphanCalls)
	}
	if prov.sweepCalls != 0 {
		t.Fatalf("sweep calls = %d, want 0", prov.sweepCalls)
	}
	events := decodeNDJSON(t, out.String())
	if events[0]["type"] != "plan" {
		t.Fatalf("first event = %v, want plan", events[0])
	}
	if events[1]["type"] != "check" || events[1]["name"] != "owner" ||
		events[1]["detail"] != "GITHUB_OWNER environment variable not set" {
		t.Fatalf("failed owner check not reported: %v", events)
	}
	summary := lastEvent(t, events)
	if summary["command"] != "e2e" || summary["mode"] != "organization" ||
		summary["preflight_ok"] != false || summary["run_attempted"] != false ||
		summary["exit_code"] != float64(1) {
		t.Fatalf("bad final summary: %v", summary)
	}
}

// TestE2ETextModePrintsPreflightChecklistOnSuccess guards against a
// regression where e2e's text-mode preflight checklist was only printed on
// failure: printPreflight was called solely inside the `!report.OK()`
// branch, so a successful preflight's checks were silently dropped from
// text-mode output even though the standalone preflight subcommand (see
// runPreflightContext) always prints its checklist unconditionally in text
// mode, success or failure.
func TestE2ETextModePrintsPreflightChecklistOnSuccess(t *testing.T) {
	root := t.TempDir()
	runner := &e2eRunner{results: []engine.RunResult{{Tests: []engine.TestResult{{
		Package: "./github", Name: "TestAccThing", Status: provider.StatusPass,
	}}}}}
	prov := newE2EProvider(provider.PreflightReport{Checks: []provider.Check{{
		Name:   "owner",
		Status: provider.CheckOK,
		Detail: "GITHUB_OWNER is set",
	}}})
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccThing"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"e2e", "--repo-root", root}, &out, &errOut, d)

	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s", code, errOut.String())
	}
	if !strings.Contains(out.String(), "[✓] owner: GITHUB_OWNER is set") {
		t.Fatalf("text output missing the successful preflight checklist: %q", out.String())
	}
}

func TestE2EUsesOrganizationScopeForEnvFile(t *testing.T) {
	root := t.TempDir()
	envPath := filepath.Join(root, "tester.env")
	if err := os.WriteFile(envPath, []byte("PULSAR_ORGANIZATION_GITHUB_OWNER=sandbox-org\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	getenv, setenv, values := mapEnv(nil)
	runner := &e2eRunner{results: []engine.RunResult{{Tests: []engine.TestResult{{
		Package: "./github", Name: "TestAccThing", Status: provider.StatusPass,
	}}}}}
	prov := newE2EProvider(provider.PreflightReport{})
	prov.ModesVal = []provider.Mode{{Name: "anonymous"}, {Name: "organization"}}
	prov.PreflightFn = func(_ context.Context, mode string, _ provider.TestRequirements) provider.PreflightReport {
		if (*values)["GITHUB_OWNER"] == "" {
			return provider.PreflightReport{Mode: mode, Checks: []provider.Check{{
				Name: "owner", Status: provider.CheckFail, Detail: "GITHUB_OWNER environment variable not set",
			}}}
		}
		return provider.PreflightReport{Mode: mode}
	}
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccThing"}),
		cwd:       func() (string, error) { return root, nil },
		getenv:    getenv,
		setenv:    setenv,
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{
		"--env-file", envPath, "e2e", "--json", "--repo-root", root,
	}, &out, &errOut, d)

	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s stdout=%s", code, errOut.String(), out.String())
	}
	if (*values)["GITHUB_OWNER"] != "sandbox-org" {
		t.Fatalf("GITHUB_OWNER = %q, want organization-scoped env value", (*values)["GITHUB_OWNER"])
	}
}

func TestE2ERetriesOnlyFailedTestsOnceThenTriagesAndChecksOrphans(t *testing.T) {
	root := t.TempDir()
	runner := &e2eRunner{results: []engine.RunResult{
		{Tests: []engine.TestResult{
			{Package: "./github", Name: "TestAccPass", Status: provider.StatusPass},
			{Package: "./github", Name: "TestAccFail", Status: provider.StatusFail, Output: []string{"403 API rate limit exceeded\n"}},
		}},
		{Tests: []engine.TestResult{
			{Package: "./github", Name: "TestAccFail", Status: provider.StatusFail, Output: []string{"still broken\n"}},
		}},
	}}
	prov := newE2EProvider(provider.PreflightReport{})
	prov.orphans = []provider.Resource{{Kind: "repository", Name: "tf-acc-test-leftover", URL: "https://example.test/repo"}}
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccPass", "TestAccFail"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"e2e", "--json", "--repo-root", root}, &out, &errOut, d)

	if code != 1 {
		t.Fatalf("exit code = %d, want 1; stderr=%s", code, errOut.String())
	}
	if len(runner.specs) != 2 {
		t.Fatalf("runner calls = %d, want initial plus one retry", len(runner.specs))
	}
	if runner.specs[0].ExtraEnv[0] != "GH_TEST_AUTH_MODE=organization" {
		t.Fatalf("initial mode env = %v, want organization", runner.specs[0].ExtraEnv)
	}
	if runner.specs[1].Pattern != "^"+regexp.QuoteMeta("TestAccFail")+"$" {
		t.Fatalf("retry pattern = %q, want only failed test", runner.specs[1].Pattern)
	}
	if regexp.MustCompile(`TestAccPass`).MatchString(runner.specs[1].Pattern) {
		t.Fatalf("retry pattern included passing test: %q", runner.specs[1].Pattern)
	}
	if prov.orphanCalls != 2 {
		t.Fatalf("orphan calls = %d, want baseline plus final", prov.orphanCalls)
	}
	if !reflect.DeepEqual(prov.orphanModes, []string{"organization", "organization"}) {
		t.Fatalf("orphan modes = %v, want two organization scans", prov.orphanModes)
	}
	if prov.sweepCalls != 0 {
		t.Fatalf("sweep calls = %d, want 0", prov.sweepCalls)
	}
	events := decodeNDJSON(t, out.String())
	seenTriage := false
	for _, event := range events {
		seenTriage = seenTriage || event["type"] == "triage"
	}
	if !seenTriage {
		t.Fatalf("events missing triage result: %v", events)
	}
	state, err := engine.Load(filepath.Join(root, ".pulsar-state.json"))
	if err != nil {
		t.Fatal(err)
	}
	if state.Orphans == nil || !state.Orphans.BaselineCaptured {
		t.Fatalf("state.Orphans = %+v, want captured baseline", state.Orphans)
	}
	summary := lastEvent(t, events)
	if summary["command"] != "e2e" || summary["retry_attempted"] != true ||
		summary["triage_attempted"] != true || summary["orphan_check_attempted"] != true ||
		summary["exit_code"] != float64(1) {
		t.Fatalf("bad final summary: %v", summary)
	}
}

// TestE2EJSONSummaryDistinguishesInitialFromRetryExitCode guards against a
// regression where e2e's summary assigned the single final runWithSpecTriage
// exit code to both initial_exit_code and retry_exit_code. When the initial
// attempt fails a retryable test and the retry then passes, that collapsed
// the two fields to the same (successful) value and erased the
// initial-failure/recovered-retry distinction the schema exists to capture.
func TestE2EJSONSummaryDistinguishesInitialFromRetryExitCode(t *testing.T) {
	root := t.TempDir()
	runner := &e2eRunner{results: []engine.RunResult{
		{Tests: []engine.TestResult{
			{Package: "./github", Name: "TestAccRateLimited", Status: provider.StatusFail, Output: []string{"403 API rate limit exceeded\n"}},
		}},
		{Tests: []engine.TestResult{
			{Package: "./github", Name: "TestAccRateLimited", Status: provider.StatusPass},
		}},
	}}
	prov := newE2EProvider(provider.PreflightReport{})
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccRateLimited"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"e2e", "--json", "--repo-root", root}, &out, &errOut, d)

	if code != 0 {
		t.Fatalf("exit code = %d, want 0 after the retry recovers; stderr=%s", code, errOut.String())
	}
	if len(runner.specs) != 2 {
		t.Fatalf("runner calls = %d, want initial plus one retry", len(runner.specs))
	}

	summary := lastEvent(t, decodeNDJSON(t, out.String()))
	if summary["retry_attempted"] != true {
		t.Fatalf("retry_attempted = %v, want true", summary["retry_attempted"])
	}
	if summary["initial_exit_code"] != float64(1) {
		t.Fatalf("initial_exit_code = %v, want 1 (the first attempt failed)", summary["initial_exit_code"])
	}
	if summary["retry_exit_code"] != float64(0) {
		t.Fatalf("retry_exit_code = %v, want 0 (the retry recovered)", summary["retry_exit_code"])
	}
	if summary["exit_code"] != float64(0) {
		t.Fatalf("exit_code = %v, want 0", summary["exit_code"])
	}
}

func TestE2ERetryExcludesStaleFailuresFromEarlierState(t *testing.T) {
	root := t.TempDir()
	state := engine.State{Results: []engine.PersistResult{{
		Test: "TestAccStale", Package: "./github", Status: "fail",
	}}}
	if err := state.Save(filepath.Join(root, ".pulsar-state.json")); err != nil {
		t.Fatal(err)
	}
	runner := &e2eRunner{results: []engine.RunResult{
		{Tests: []engine.TestResult{
			{Package: "./github", Name: "TestAccCurrent", Status: provider.StatusFail, Output: []string{"403 API rate limit exceeded\n"}},
		}},
		{Tests: []engine.TestResult{
			{Package: "./github", Name: "TestAccCurrent", Status: provider.StatusPass},
		}},
	}}
	prov := newE2EProvider(provider.PreflightReport{})
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccCurrent"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	_ = runWithDeps([]string{"e2e", "--json", "--repo-root", root}, &out, &errOut, d)

	if len(runner.specs) != 2 {
		t.Fatalf("runner calls = %d, want initial plus one retry; stderr=%s", len(runner.specs), errOut.String())
	}
	if strings.Contains(runner.specs[1].Pattern, "TestAccStale") {
		t.Fatalf("retry pattern included stale failure: %q", runner.specs[1].Pattern)
	}
	if runner.specs[1].Pattern != "^"+regexp.QuoteMeta("TestAccCurrent")+"$" {
		t.Fatalf("retry pattern = %q, want current failed test only", runner.specs[1].Pattern)
	}
	summary := lastEvent(t, decodeNDJSON(t, out.String()))
	if summary["triage_attempted"] != false || summary["exit_code"] != float64(0) {
		t.Fatalf("stale failure affected workflow result: %v", summary)
	}
}

func TestE2EDoesNotRetryBuildFailureAndStillChecksOrphans(t *testing.T) {
	root := t.TempDir()
	runner := &e2eRunner{results: []engine.RunResult{{
		BuildFailed: true,
		BuildOutput: []string{"compile failed\n"},
	}}}
	prov := newE2EProvider(provider.PreflightReport{})
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccThing"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"e2e", "--json", "--repo-root", root}, &out, &errOut, d)

	if code != 1 {
		t.Fatalf("exit code = %d, want 1; stderr=%s", code, errOut.String())
	}
	if len(runner.specs) != 1 {
		t.Fatalf("runner calls = %d, want no build-failure retry", len(runner.specs))
	}
	if prov.orphanCalls != 2 {
		t.Fatalf("orphan calls = %d, want baseline plus final", prov.orphanCalls)
	}
	summary := lastEvent(t, decodeNDJSON(t, out.String()))
	if summary["retry_attempted"] != false || summary["triage_attempted"] != true {
		t.Fatalf("build failure treated incorrectly: %v", summary)
	}
}

func TestE2EChecksOrphansAfterSuccessfulRun(t *testing.T) {
	root := t.TempDir()
	runner := &e2eRunner{results: []engine.RunResult{{Tests: []engine.TestResult{{
		Package: "./github", Name: "TestAccThing", Status: provider.StatusPass,
	}}}}}
	prov := newE2EProvider(provider.PreflightReport{})
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccThing"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"e2e", "--json", "--repo-root", root}, &out, &errOut, d)

	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s", code, errOut.String())
	}
	if len(runner.specs) != 1 {
		t.Fatalf("runner calls = %d, want 1", len(runner.specs))
	}
	if prov.orphanCalls != 2 {
		t.Fatalf("orphan calls = %d, want baseline plus final", prov.orphanCalls)
	}
	summary := lastEvent(t, decodeNDJSON(t, out.String()))
	if summary["run_attempted"] != true || summary["retry_attempted"] != false ||
		summary["triage_attempted"] != false || summary["orphan_check_attempted"] != true ||
		summary["exit_code"] != float64(0) {
		t.Fatalf("bad success summary: %v", summary)
	}
}

func TestE2ESummaryWriteFailureForcesNonzeroExit(t *testing.T) {
	root := t.TempDir()
	runner := &e2eRunner{results: []engine.RunResult{{Tests: []engine.TestResult{{
		Package: "./github", Name: "TestAccThing", Status: provider.StatusPass,
	}}}}}
	prov := newE2EProvider(provider.PreflightReport{})
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccThing"}),
		cwd:       func() (string, error) { return root, nil },
	}

	out := &failE2ESummaryWriter{}
	var errOut bytes.Buffer
	code := runWithDeps([]string{"e2e", "--json", "--repo-root", root}, out, &errOut, d)

	if code != 1 {
		t.Fatalf("exit code = %d, want 1 when the terminal e2e summary cannot be written; stderr=%s stdout=%s",
			code, errOut.String(), out.buf.String())
	}
	if got := strings.Count(errOut.String(), "simulated e2e summary write failure"); got != 1 {
		t.Fatalf("summary write error count = %d, want exactly 1; stderr=%s", got, errOut.String())
	}
}

func TestE2EChecksOrphansAfterInterruptedRun(t *testing.T) {
	root := t.TempDir()
	state := engine.State{Results: []engine.PersistResult{{
		Test: "TestAccOldFailure", Package: "./github", Status: "fail",
	}}}
	if err := state.Save(filepath.Join(root, ".pulsar-state.json")); err != nil {
		t.Fatal(err)
	}
	runner := &e2eRunner{
		results: []engine.RunResult{{}},
		errs:    []error{context.Canceled},
	}
	prov := newE2EProvider(provider.PreflightReport{})
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccThing"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"e2e", "--json", "--repo-root", root}, &out, &errOut, d)

	if code != 1 {
		t.Fatalf("exit code = %d, want 1; stderr=%s", code, errOut.String())
	}
	if len(runner.specs) != 1 {
		t.Fatalf("runner calls = %d, want interrupted initial run only", len(runner.specs))
	}
	if prov.orphanCalls != 2 {
		t.Fatalf("orphan calls = %d, want baseline plus final after interruption", prov.orphanCalls)
	}
	if !strings.Contains(errOut.String(), errors.New("context canceled").Error()) {
		t.Fatalf("stderr missing interruption: %s", errOut.String())
	}
	summary := lastEvent(t, decodeNDJSON(t, out.String()))
	if summary["run_attempted"] != true || summary["retry_attempted"] != false ||
		summary["triage_attempted"] != false || summary["orphan_check_attempted"] != true {
		t.Fatalf("bad interruption summary: %v", summary)
	}
	// An interrupted run never reaches cur.Save, so state on disk is stale or
	// absent and lastRunCompleted(tracker) is false: initial_exit_code must
	// fall back to the run step's own exit code (1) rather than reporting a
	// success recovered from unrelated leftover state.
	if summary["initial_exit_code"] != float64(1) {
		t.Fatalf("initial_exit_code = %v, want 1 for an interrupted run", summary["initial_exit_code"])
	}
	if summary["retry_exit_code"] != nil {
		t.Fatalf("retry_exit_code = %v, want null when no retry was attempted", summary["retry_exit_code"])
	}
}

func TestE2EDefaultsToOrganization(t *testing.T) {
	root := t.TempDir()
	runner := &e2eRunner{results: []engine.RunResult{{Tests: []engine.TestResult{{
		Package: "./github", Name: "TestAccThing", Status: provider.StatusPass,
	}}}}}
	prov := newE2EProvider(provider.PreflightReport{})
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccThing"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"e2e", "--json", "--repo-root", root}, &out, &errOut, d)

	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s", code, errOut.String())
	}
	if len(runner.specs) != 1 || !extraEnvContains(runner.specs[0].ExtraEnv, "GH_TEST_AUTH_MODE=organization") {
		t.Fatalf("expected the default organization mode in child env, got specs=%v", runner.specs)
	}
	summary := lastEvent(t, decodeNDJSON(t, out.String()))
	if summary["mode"] != "organization" {
		t.Fatalf("summary mode = %v, want organization", summary["mode"])
	}
}

func TestE2EUsesIndividualEnvFileScope(t *testing.T) {
	root := t.TempDir()
	envPath := filepath.Join(root, "tester.env")
	if err := os.WriteFile(envPath, []byte("PULSAR_INDIVIDUAL_GITHUB_OWNER=solo-user\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	getenv, setenv, values := mapEnv(nil)
	runner := &e2eRunner{results: []engine.RunResult{{Tests: []engine.TestResult{{
		Package: "./github", Name: "TestAccThing", Status: provider.StatusPass,
	}}}}}
	prov := newE2EProvider(provider.PreflightReport{})
	prov.ModesVal = []provider.Mode{{Name: "anonymous"}, {Name: "individual"}, {Name: "organization"}}
	prov.PreflightFn = func(_ context.Context, mode string, _ provider.TestRequirements) provider.PreflightReport {
		if (*values)["GITHUB_OWNER"] == "" {
			return provider.PreflightReport{Mode: mode, Checks: []provider.Check{{
				Name: "owner", Status: provider.CheckFail, Detail: "GITHUB_OWNER environment variable not set",
			}}}
		}
		return provider.PreflightReport{Mode: mode}
	}
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccThing"}),
		cwd:       func() (string, error) { return root, nil },
		getenv:    getenv,
		setenv:    setenv,
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{
		"--env-file", envPath, "e2e", "--mode", "individual", "--json", "--repo-root", root,
	}, &out, &errOut, d)

	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s stdout=%s", code, errOut.String(), out.String())
	}
	if (*values)["GITHUB_OWNER"] != "solo-user" {
		t.Fatalf("GITHUB_OWNER = %q, want individual-scoped env value", (*values)["GITHUB_OWNER"])
	}
}

func TestE2EAnonymousSkipsCredentialedRequirements(t *testing.T) {
	root := t.TempDir()
	runner := &e2eRunner{results: []engine.RunResult{{Tests: []engine.TestResult{{
		Package: "./github", Name: "TestAccPublic", Status: provider.StatusPass,
	}}}}}
	prov := newE2EProvider(provider.PreflightReport{})
	prov.RequirementsFn = func(name string) (provider.TestRequirements, bool) {
		if name == "TestAccCredentialed" {
			return provider.TestRequirements{Modes: []string{"individual", "organization"}, Scopes: []string{"repo"}}, true
		}
		return provider.TestRequirements{}, true
	}
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccPublic", "TestAccCredentialed"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"e2e", "--mode", "anonymous", "--json", "--repo-root", root}, &out, &errOut, d)

	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s", code, errOut.String())
	}
	if len(runner.specs) != 1 {
		t.Fatalf("runner calls = %d, want 1", len(runner.specs))
	}
	if runner.specs[0].Pattern != engine.RunPattern([]string{"TestAccPublic"}) {
		t.Fatalf("run pattern = %q, want only the anonymous-eligible test", runner.specs[0].Pattern)
	}
	events := decodeNDJSON(t, out.String())
	if events[0]["type"] != "plan" || events[0]["eligible"] != float64(1) || events[0]["excluded"] != float64(1) {
		t.Fatalf("plan event missing the credentialed exclusion: %v", events[0])
	}
}

func TestE2ERejectsGuidedTeamAndEnterprise(t *testing.T) {
	for _, mode := range []string{"team", "enterprise"} {
		t.Run(mode, func(t *testing.T) {
			root := t.TempDir()
			runner := &e2eRunner{}
			listCalls := 0
			prov := newE2EProvider(provider.PreflightReport{})
			d := deps{
				provider:  prov,
				newRunner: func(io.Writer) testRunner { return runner },
				list: func(context.Context, string, []string, string) ([]string, error) {
					listCalls++
					return []string{"TestAccThing"}, nil
				},
				cwd: func() (string, error) { return root, nil },
			}

			var out, errOut bytes.Buffer
			code := runWithDeps([]string{"e2e", "--mode", mode, "--repo-root", root}, &out, &errOut, d)

			if code != 2 {
				t.Fatalf("exit code = %d, want 2; stderr=%s", code, errOut.String())
			}
			if listCalls != 0 || len(runner.specs) != 0 {
				t.Fatalf("expected no work before the mode is rejected: list calls=%d runner calls=%d", listCalls, len(runner.specs))
			}
			if errOut.String() == "" {
				t.Fatalf("expected a usage error on stderr")
			}
		})
	}
}

func TestE2ENeverRunsIncompatibleTests(t *testing.T) {
	root := t.TempDir()
	runner := &e2eRunner{results: []engine.RunResult{{Tests: []engine.TestResult{{
		Package: "./github", Name: "TestAccOrgCompatible", Status: provider.StatusPass,
	}}}}}
	prov := newE2EProvider(provider.PreflightReport{})
	prov.RequirementsFn = func(name string) (provider.TestRequirements, bool) {
		if name == "TestAccIndividualOnly" {
			return provider.TestRequirements{Modes: []string{"individual"}}, true
		}
		return provider.TestRequirements{}, true
	}
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccOrgCompatible", "TestAccIndividualOnly"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"e2e", "--json", "--repo-root", root}, &out, &errOut, d)

	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s", code, errOut.String())
	}
	if len(runner.specs) != 1 {
		t.Fatalf("runner calls = %d, want 1", len(runner.specs))
	}
	if runner.specs[0].Pattern != engine.RunPattern([]string{"TestAccOrgCompatible"}) {
		t.Fatalf("run pattern = %q, want only the organization-eligible test", runner.specs[0].Pattern)
	}
	if strings.Contains(runner.specs[0].Pattern, "TestAccIndividualOnly") {
		t.Fatalf("pattern included an incompatible test: %q", runner.specs[0].Pattern)
	}
}

func TestE2ERetriesOnlyRetryableFailuresOnce(t *testing.T) {
	root := t.TempDir()
	runner := &e2eRunner{results: []engine.RunResult{
		{Tests: []engine.TestResult{
			{Package: "./github", Name: "TestAccRateLimited", Status: provider.StatusFail, Output: []string{"403 API rate limit exceeded\n"}},
			{Package: "./github", Name: "TestAccBadInput", Status: provider.StatusFail, Output: []string{"boom\n"}},
		}},
		{Tests: []engine.TestResult{
			{Package: "./github", Name: "TestAccRateLimited", Status: provider.StatusPass},
		}},
	}}
	prov := newE2EProvider(provider.PreflightReport{})
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccRateLimited", "TestAccBadInput"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"e2e", "--json", "--repo-root", root}, &out, &errOut, d)

	if code != 1 {
		t.Fatalf("exit code = %d, want 1 (TestAccBadInput never retried-away); stderr=%s", code, errOut.String())
	}
	if len(runner.specs) != 2 {
		t.Fatalf("runner calls = %d, want initial plus exactly one retry", len(runner.specs))
	}
	if runner.specs[1].Pattern != "^"+regexp.QuoteMeta("TestAccRateLimited")+"$" {
		t.Fatalf("retry pattern = %q, want only the retryable failure", runner.specs[1].Pattern)
	}
	if strings.Contains(runner.specs[1].Pattern, "TestAccBadInput") {
		t.Fatalf("retry pattern included a non-retryable failure: %q", runner.specs[1].Pattern)
	}
}

func TestE2ESendsDeterministicFailuresDirectlyToTriage(t *testing.T) {
	root := t.TempDir()
	runner := &e2eRunner{results: []engine.RunResult{{Tests: []engine.TestResult{{
		Package: "./github", Name: "TestAccBroken", Status: provider.StatusFail, Output: []string{"boom\n"},
	}}}}}
	prov := newE2EProvider(provider.PreflightReport{})
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccBroken"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"e2e", "--json", "--repo-root", root}, &out, &errOut, d)

	if code != 1 {
		t.Fatalf("exit code = %d, want 1; stderr=%s", code, errOut.String())
	}
	if len(runner.specs) != 1 {
		t.Fatalf("runner calls = %d, want no retry attempt for a deterministic failure", len(runner.specs))
	}
	events := decodeNDJSON(t, out.String())
	seenTriage := false
	for _, event := range events {
		seenTriage = seenTriage || event["type"] == "triage"
	}
	if !seenTriage {
		t.Fatalf("expected a triage event for the deterministic failure: %v", events)
	}
	summary := lastEvent(t, events)
	if summary["retry_attempted"] != false || summary["triage_attempted"] != true || summary["exit_code"] != float64(1) {
		t.Fatalf("bad summary: %v", summary)
	}
}

func TestE2ETextSummaryNamesSelectedMode(t *testing.T) {
	root := t.TempDir()
	runner := &e2eRunner{results: []engine.RunResult{{Tests: []engine.TestResult{{
		Package: "./github", Name: "TestAccThing", Status: provider.StatusPass,
	}}}}}
	prov := newE2EProvider(provider.PreflightReport{})
	prov.ModesVal = []provider.Mode{{Name: "anonymous"}, {Name: "individual"}, {Name: "organization"}}
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccThing"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"e2e", "--mode", "individual", "--repo-root", root}, &out, &errOut, d)

	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s", code, errOut.String())
	}
	if !strings.Contains(out.String(), "individual acceptance-test workflow completed") {
		t.Fatalf("text summary missing the selected mode: %q", out.String())
	}
}

func TestE2EJSONFinalObjectRemainsSummary(t *testing.T) {
	root := t.TempDir()
	runner := &e2eRunner{results: []engine.RunResult{{Tests: []engine.TestResult{{
		Package: "./github", Name: "TestAccBroken", Status: provider.StatusFail, Output: []string{"boom\n"},
	}}}}}
	prov := newE2EProvider(provider.PreflightReport{})
	prov.orphans = []provider.Resource{{Kind: "repository", Name: "tf-acc-test-leftover", URL: "https://example.test/repo"}}
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccBroken"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"e2e", "--json", "--repo-root", root}, &out, &errOut, d)
	if code != 1 {
		t.Fatalf("exit code = %d, want 1; stderr=%s", code, errOut.String())
	}

	events := decodeNDJSON(t, out.String())
	if len(events) == 0 {
		t.Fatalf("expected at least one JSON event")
	}
	last := events[len(events)-1]
	if last["type"] != "summary" || last["command"] != "e2e" {
		t.Fatalf("final JSON object = %v, want the e2e summary", last)
	}
}

func TestE2EJSONPlanFailureStillEmitsFinalSummary(t *testing.T) {
	root := t.TempDir()
	prov := newE2EProvider(provider.PreflightReport{})
	prov.ModesVal = []provider.Mode{{Name: "anonymous"}, {Name: "individual"}, {Name: "organization"}}
	prov.RequirementsFn = func(name string) (provider.TestRequirements, bool) {
		if name == "TestAccUnclassified" {
			return provider.TestRequirements{}, false
		}
		return provider.TestRequirements{}, true
	}
	d := deps{
		provider: prov,
		list:     stubList([]string{"TestAccUnclassified"}),
		cwd:      func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"e2e", "--mode", "individual", "--json", "--repo-root", root}, &out, &errOut, d)

	if code != 1 {
		t.Fatalf("exit code = %d, want 1; stderr=%s", code, errOut.String())
	}
	events := decodeNDJSON(t, out.String())
	if got := len(events); got != 1 {
		t.Fatalf("events = %d, want only the terminal summary: %v", got, events)
	}
	summary := lastEvent(t, events)
	if summary["type"] != "summary" || summary["command"] != "e2e" {
		t.Fatalf("final JSON object = %v, want the e2e summary", summary)
	}
	if summary["mode"] != "individual" || summary["run_attempted"] != false || summary["exit_code"] != float64(1) {
		t.Fatalf("bad final summary: %v", summary)
	}
}

func TestE2EJSONLockFailureStillEmitsFinalSummary(t *testing.T) {
	root := t.TempDir()
	lockPath := filepath.Join(root, ".pulsar-state.json.lock")
	release, err := engine.AcquireLock(lockPath)
	if err != nil {
		t.Fatalf("AcquireLock: %v", err)
	}
	defer func() {
		if err := release(); err != nil {
			t.Fatalf("release lock: %v", err)
		}
	}()

	runner := &e2eRunner{}
	prov := newE2EProvider(provider.PreflightReport{Checks: []provider.Check{{
		Name:   "owner",
		Status: provider.CheckOK,
		Detail: "GITHUB_OWNER is set",
	}}})
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccThing"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"e2e", "--mode", "organization", "--json", "--repo-root", root}, &out, &errOut, d)

	if code != 1 {
		t.Fatalf("exit code = %d, want 1; stderr=%s", code, errOut.String())
	}
	if len(runner.specs) != 0 {
		t.Fatalf("runner calls = %d, want 0 when the run lock is already held", len(runner.specs))
	}
	events := decodeNDJSON(t, out.String())
	if got := len(events); got != 3 {
		t.Fatalf("events = %d, want plan + check + summary: %v", got, events)
	}
	if events[0]["type"] != "plan" || events[1]["type"] != "check" {
		t.Fatalf("unexpected event ordering before summary: %v", events)
	}
	summary := lastEvent(t, events)
	if summary["type"] != "summary" || summary["command"] != "e2e" {
		t.Fatalf("final JSON object = %v, want the e2e summary", summary)
	}
	if summary["mode"] != "organization" || summary["run_attempted"] != false || summary["exit_code"] != float64(1) {
		t.Fatalf("bad final summary: %v", summary)
	}
}

func TestE2ENeverCallsSweep(t *testing.T) {
	root := t.TempDir()
	runner := &e2eRunner{results: []engine.RunResult{{Tests: []engine.TestResult{{
		Package: "./github", Name: "TestAccBroken", Status: provider.StatusFail, Output: []string{"boom\n"},
	}}}}}
	prov := newE2EProvider(provider.PreflightReport{})
	prov.orphans = []provider.Resource{{Kind: "repository", Name: "tf-acc-test-leftover", URL: "https://example.test/repo"}}
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccBroken"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	_ = runWithDeps([]string{"e2e", "--json", "--repo-root", root}, &out, &errOut, d)

	if prov.sweepCalls != 0 {
		t.Fatalf("sweep calls = %d, want 0", prov.sweepCalls)
	}
}

// TestE2EAnonymousSkipsOrphanInspection is the regression for guided
// anonymous runs performing a credentialed orphan scan. The real provider's
// Orphans requires GITHUB_OWNER and issues authenticated queries, so an
// unauthenticated scan is both meaningless and guaranteed to fail - which
// turned a genuinely successful zero-credential anonymous run into a
// non-zero exit. The Orphans double here errors and counts calls: an
// eligible anonymous run must succeed with orphan_check_attempted false and
// zero Orphans calls.
func TestE2EAnonymousSkipsOrphanInspection(t *testing.T) {
	root := t.TempDir()
	runner := &e2eRunner{results: []engine.RunResult{{Tests: []engine.TestResult{
		{Package: "./github", Name: "TestAccPublic", Status: provider.StatusPass},
	}}}}
	prov := newE2EProvider(provider.PreflightReport{})
	prov.orphanErr = errors.New("GITHUB_OWNER is not set")
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
	if len(runner.specs) != 1 {
		t.Fatalf("runner calls = %d, want the anonymous run", len(runner.specs))
	}
	if prov.orphanCalls != 0 {
		t.Fatalf("orphan calls = %d, want 0 for an uncredentialed anonymous run", prov.orphanCalls)
	}
	summary := lastEvent(t, decodeNDJSON(t, out.String()))
	if summary["run_attempted"] != true {
		t.Fatalf("run_attempted = %v, want true", summary["run_attempted"])
	}
	if summary["orphan_check_attempted"] != false {
		t.Fatalf("orphan_check_attempted = %v, want false in anonymous mode", summary["orphan_check_attempted"])
	}
	if summary["orphan_check_exit_code"] != nil {
		t.Fatalf("orphan_check_exit_code = %v, want null when no orphan check ran", summary["orphan_check_exit_code"])
	}
	if summary["exit_code"] != float64(0) {
		t.Fatalf("exit_code = %v, want 0", summary["exit_code"])
	}
}

// TestE2ECredentialedModesStillInspectOrphans pins the other half of the
// credentialed gate: individual and organization runs, which do have
// credentials, must still perform the read-only orphan inspection.
func TestE2ECredentialedModesStillInspectOrphans(t *testing.T) {
	for _, mode := range []string{"individual", "organization"} {
		t.Run(mode, func(t *testing.T) {
			root := t.TempDir()
			runner := &e2eRunner{results: []engine.RunResult{{Tests: []engine.TestResult{
				{Package: "./github", Name: "TestAccThing", Status: provider.StatusPass},
			}}}}
			prov := newE2EProvider(provider.PreflightReport{})
			prov.ModesVal = []provider.Mode{{Name: "anonymous"}, {Name: "individual"}, {Name: "organization"}}
			d := deps{
				provider:  prov,
				newRunner: func(io.Writer) testRunner { return runner },
				list:      stubList([]string{"TestAccThing"}),
				cwd:       func() (string, error) { return root, nil },
			}

			var out, errOut bytes.Buffer
			code := runWithDeps([]string{"e2e", "--mode", mode, "--json", "--repo-root", root}, &out, &errOut, d)

			if code != 0 {
				t.Fatalf("exit code = %d, want 0; stderr=%s", code, errOut.String())
			}
			if prov.orphanCalls != 2 {
				t.Fatalf("orphan calls = %d, want baseline plus final in credentialed mode %s", prov.orphanCalls, mode)
			}
			if !reflect.DeepEqual(prov.orphanModes, []string{mode, mode}) {
				t.Fatalf("orphan modes = %v, want two %s scans", prov.orphanModes, mode)
			}
			summary := lastEvent(t, decodeNDJSON(t, out.String()))
			if summary["orphan_check_attempted"] != true {
				t.Fatalf("orphan_check_attempted = %v, want true in mode %s", summary["orphan_check_attempted"], mode)
			}
		})
	}
}

// TestE2EZeroEligibleSkipsRunnerAndOrphans is the permanent regression for
// the zero-eligible plan path: engine.RunPattern returns "" for an empty
// eligible set and an empty `go test -run` matches EVERY test, so e2e must
// skip the runner entirely rather than run the whole suite. The credentialed
// logical run still captures a baseline under the caller-owned lock, and the
// workflow still ends in a successful terminal summary.
func TestE2EZeroEligibleSkipsRunnerAndOrphans(t *testing.T) {
	root := t.TempDir()
	runner := &e2eRunner{}
	prov := newE2EProvider(provider.PreflightReport{})
	prov.ModesVal = []provider.Mode{{Name: "anonymous"}, {Name: "individual"}, {Name: "organization"}}
	prov.RequirementsFn = func(string) (provider.TestRequirements, bool) {
		return provider.TestRequirements{Modes: []string{"individual"}}, true
	}
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccIndividualOnly"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"e2e", "--mode", "organization", "--json", "--repo-root", root}, &out, &errOut, d)

	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s stdout=%s", code, errOut.String(), out.String())
	}
	if len(runner.specs) != 0 {
		t.Fatalf("runner calls = %d, want 0 for a zero-eligible plan: specs=%v", len(runner.specs), runner.specs)
	}
	if prov.orphanCalls != 1 {
		t.Fatalf("orphan baseline calls = %d, want 1 for the credentialed logical run", prov.orphanCalls)
	}
	events := decodeNDJSON(t, out.String())
	if events[0]["type"] != "plan" || events[0]["eligible"] != float64(0) {
		t.Fatalf("first event = %v, want a zero-eligible plan", events[0])
	}
	summary := lastEvent(t, events)
	if summary["command"] != "e2e" || summary["run_attempted"] != false ||
		summary["orphan_check_attempted"] != true || summary["exit_code"] != float64(0) {
		t.Fatalf("bad zero-eligible summary: %v", summary)
	}
	state, err := engine.Load(filepath.Join(root, ".pulsar-state.json"))
	if err != nil {
		t.Fatal(err)
	}
	if state.Orphans == nil || !state.Orphans.BaselineCaptured {
		t.Fatalf("state.Orphans = %+v, want captured credentialed baseline", state.Orphans)
	}
	if summary["initial_exit_code"] != nil {
		t.Fatalf("initial_exit_code = %v, want null when no run was attempted", summary["initial_exit_code"])
	}
}

func TestE2EZeroEligibleBaselineOnlyOrphanTextDoesNotProveUnmeasuredCounts(t *testing.T) {
	root := t.TempDir()
	baseline := []provider.Resource{{
		Kind: "repository", Name: "tf-acc-test-existing", URL: "https://example.test/existing",
	}}
	runner := &e2eRunner{}
	prov := newE2EProvider(provider.PreflightReport{})
	prov.ModesVal = []provider.Mode{{Name: "anonymous"}, {Name: "individual"}, {Name: "organization"}}
	prov.RequirementsFn = func(string) (provider.TestRequirements, bool) {
		return provider.TestRequirements{Modes: []string{"individual"}}, true
	}
	prov.orphanSeq = [][]provider.Resource{baseline}
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccIndividualOnly"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"e2e", "--mode", "organization", "--repo-root", root}, &out, &errOut, d)
	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s stdout=%s", code, errOut.String(), out.String())
	}
	if len(runner.specs) != 0 {
		t.Fatalf("runner calls = %d, want 0 for a zero-eligible plan: specs=%v", len(runner.specs), runner.specs)
	}
	if prov.orphanCalls != 1 {
		t.Fatalf("orphan calls = %d, want only the baseline capture", prov.orphanCalls)
	}

	text := out.String()
	for _, want := range []string{
		"Baseline: 1",
		"Pre-existing: not captured",
		"Final: not captured",
		"New: 0 (no run attempted; no attributed cleanup obligation)",
		"Cleanup status: baseline-only",
		"Preview command: terraform-provider-tester orphans --repo-root '" + root + "' --mode organization --run-delta",
		"Cleanup command: terraform-provider-tester sweep --repo-root '" + root + "' --mode organization --run-delta --confirm",
		"organization acceptance tests were not started",
	} {
		if !strings.Contains(text, want) {
			t.Errorf("baseline-only text missing %q:\n%s", want, text)
		}
	}
	for _, misleading := range []string{
		"Pre-existing: 0",
		"Final: 0",
		"New (this run's cleanup obligation): 0",
	} {
		if strings.Contains(text, misleading) {
			t.Errorf("baseline-only text presents an unmeasured count as proven zero (%q):\n%s", misleading, text)
		}
	}
}

// TestE2EAllowUnclassifiedRunsUnknownOnlyWhenModeCompatible proves the guided
// workflow honors the same conservative-requirements rule as the planner: an
// allowed unknown test runs under a credentialed mode whose requirements it
// supports, and is excluded (never run) under anonymous, where the
// conservative authenticated requirements are incompatible.
func TestE2EAllowUnclassifiedRunsUnknownOnlyWhenModeCompatible(t *testing.T) {
	newDeps := func(root string, runner *e2eRunner) (deps, *e2eProvider) {
		prov := newE2EProvider(provider.PreflightReport{})
		prov.ModesVal = []provider.Mode{{Name: "anonymous"}, {Name: "individual"}, {Name: "organization"}}
		prov.RequirementsFn = func(string) (provider.TestRequirements, bool) {
			return provider.TestRequirements{
				Modes:       []string{"individual", "organization", "team", "enterprise"},
				Scopes:      []string{"repo"},
				SideEffects: []string{"unknown"},
			}, false
		}
		return deps{
			provider:  prov,
			newRunner: func(io.Writer) testRunner { return runner },
			list:      stubList([]string{"TestAccMystery"}),
			cwd:       func() (string, error) { return root, nil },
		}, prov
	}

	t.Run("organization runs it", func(t *testing.T) {
		root := t.TempDir()
		runner := &e2eRunner{results: []engine.RunResult{{Tests: []engine.TestResult{
			{Package: "./github", Name: "TestAccMystery", Status: provider.StatusPass},
		}}}}
		d, _ := newDeps(root, runner)

		var out, errOut bytes.Buffer
		code := runWithDeps([]string{
			"e2e", "--mode", "organization", "--allow-unclassified", "--json", "--repo-root", root,
		}, &out, &errOut, d)

		if code != 0 {
			t.Fatalf("exit code = %d, want 0; stderr=%s", code, errOut.String())
		}
		if len(runner.specs) != 1 {
			t.Fatalf("runner calls = %d, want the allowed unknown test to run", len(runner.specs))
		}
		events := decodeNDJSON(t, out.String())
		if events[0]["eligible"] != float64(1) || events[0]["unclassified"] != float64(1) {
			t.Fatalf("plan event = %v, want 1 eligible and 1 unclassified", events[0])
		}
	})

	t.Run("anonymous excludes it", func(t *testing.T) {
		root := t.TempDir()
		runner := &e2eRunner{}
		d, prov := newDeps(root, runner)

		var out, errOut bytes.Buffer
		code := runWithDeps([]string{
			"e2e", "--mode", "anonymous", "--allow-unclassified", "--json", "--repo-root", root,
		}, &out, &errOut, d)

		if code != 0 {
			t.Fatalf("exit code = %d, want 0; stderr=%s", code, errOut.String())
		}
		if len(runner.specs) != 0 {
			t.Fatalf("runner calls = %d, want 0: the unknown test is anonymous-incompatible", len(runner.specs))
		}
		if prov.orphanCalls != 0 {
			t.Fatalf("orphan calls = %d, want 0", prov.orphanCalls)
		}
		events := decodeNDJSON(t, out.String())
		if events[0]["eligible"] != float64(0) || events[0]["excluded"] != float64(1) {
			t.Fatalf("plan event = %v, want 0 eligible and 1 excluded", events[0])
		}
		if events[1]["type"] != "excluded" || events[1]["code"] != "excluded/mode-incompatible" {
			t.Fatalf("second event = %v, want the mode-incompatible exclusion", events[1])
		}
	})
}

// TestE2EJSONPreservesProviderGroupAttribution verifies that guided e2e keeps
// the provider's group classification for every eligible test. The run step
// previously passed Groups: nil into runWithSpec, so every streamed `test`
// event fell back to "misc" (see groupForTest) and the embedded run summary
// carried an empty `groups` array, even though planForCommand had already
// computed the real buckets.
func TestE2EJSONPreservesProviderGroupAttribution(t *testing.T) {
	root := t.TempDir()
	runner := &e2eRunner{results: []engine.RunResult{{Tests: []engine.TestResult{
		{Package: "./github", Name: "TestAccGithubRepository", Status: provider.StatusPass, Elapsed: 1},
		{Package: "./github", Name: "TestAccGithubIssueLabel", Status: provider.StatusPass, Elapsed: 2},
	}}}}
	prov := newE2EProvider(provider.PreflightReport{})
	prov.GroupFunc = func(name string) string {
		if name == "TestAccGithubIssueLabel" {
			return "issues"
		}
		return "repositories"
	}
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccGithubRepository", "TestAccGithubIssueLabel"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"e2e", "--json", "--repo-root", root}, &out, &errOut, d)
	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s", code, errOut.String())
	}

	events := decodeNDJSON(t, out.String())
	gotGroups := map[string]string{}
	var runSummary map[string]any
	for _, ev := range events {
		switch ev["type"] {
		case "test":
			gotGroups[ev["name"].(string)] = ev["group"].(string)
		case "summary":
			if ev["command"] == "run" {
				runSummary = ev
			}
		}
	}
	want := map[string]string{
		"TestAccGithubRepository": "repositories",
		"TestAccGithubIssueLabel": "issues",
	}
	if !reflect.DeepEqual(gotGroups, want) {
		t.Errorf("streamed test event groups = %v, want %v", gotGroups, want)
	}
	if runSummary == nil {
		t.Fatalf("no embedded run summary in e2e output: %v", events)
	}
	groups, _ := runSummary["groups"].([]any)
	totals := map[string]float64{}
	for _, g := range groups {
		gm := g.(map[string]any)
		totals[gm["name"].(string)] = gm["total"].(float64)
	}
	wantTotals := map[string]float64{"repositories": 1, "issues": 1}
	if !reflect.DeepEqual(totals, wantTotals) {
		t.Errorf("run summary group totals = %v, want %v", totals, wantTotals)
	}
}

// TestE2EGroupSummariesExcludeIneligibleTests verifies that a test the plan
// excluded (mode-incompatible) never appears in any group summary or group
// total, so the guided summary counts only what actually ran.
func TestE2EGroupSummariesExcludeIneligibleTests(t *testing.T) {
	root := t.TempDir()
	runner := &e2eRunner{results: []engine.RunResult{{Tests: []engine.TestResult{
		{Package: "./github", Name: "TestAccOrgOnly", Status: provider.StatusPass, Elapsed: 1},
	}}}}
	prov := newE2EProvider(provider.PreflightReport{})
	prov.GroupFunc = func(name string) string {
		if name == "TestAccIndividualOnly" {
			return "users"
		}
		return "organizations"
	}
	prov.RequirementsFn = func(name string) (provider.TestRequirements, bool) {
		if name == "TestAccIndividualOnly" {
			return provider.TestRequirements{Modes: []string{"individual"}}, true
		}
		return provider.TestRequirements{}, true
	}
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccOrgOnly", "TestAccIndividualOnly"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"e2e", "--json", "--repo-root", root}, &out, &errOut, d)
	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s", code, errOut.String())
	}

	var runSummary map[string]any
	for _, ev := range decodeNDJSON(t, out.String()) {
		if ev["type"] == "summary" && ev["command"] == "run" {
			runSummary = ev
		}
	}
	if runSummary == nil {
		t.Fatal("no embedded run summary in e2e output")
	}
	groups, _ := runSummary["groups"].([]any)
	if len(groups) != 1 {
		t.Fatalf("group summaries = %v, want exactly the eligible organizations group", groups)
	}
	gm := groups[0].(map[string]any)
	if gm["name"] != "organizations" || gm["total"] != float64(1) || gm["pass"] != float64(1) {
		t.Errorf("group summary = %v, want organizations with total 1 and pass 1", gm)
	}
}

// TestE2ETextStreamsProviderGroupNames verifies the text surface keeps group
// attribution too: each streamed result line names the provider group rather
// than the "misc" fallback, and the terminal summary reports per-group totals.
func TestE2ETextStreamsProviderGroupNames(t *testing.T) {
	root := t.TempDir()
	runner := &e2eRunner{results: []engine.RunResult{{Tests: []engine.TestResult{
		{Package: "./github", Name: "TestAccGithubRepository", Status: provider.StatusPass, Elapsed: 1},
	}}}}
	prov := newE2EProvider(provider.PreflightReport{})
	prov.GroupFunc = func(string) string { return "repositories" }
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccGithubRepository"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"e2e", "--repo-root", root}, &out, &errOut, d)
	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s", code, errOut.String())
	}
	if !strings.Contains(out.String(), "pass repositories TestAccGithubRepository") {
		t.Errorf("text stream did not name the provider group:\n%s", out.String())
	}
	if strings.Contains(out.String(), "misc TestAccGithubRepository") {
		t.Errorf("text stream fell back to the misc group:\n%s", out.String())
	}
	if !strings.Contains(out.String(), "repositories: 1 passed, 0 failed, 0 skipped") {
		t.Errorf("terminal summary missing per-group totals:\n%s", out.String())
	}
}

// TestE2ERetryTriageKeepsGroupAttribution verifies group attribution survives
// the guided retry+triage pass: the retried test's streamed events and the
// triage failure entry both carry the provider group, not "misc".
func TestE2ERetryTriageKeepsGroupAttribution(t *testing.T) {
	root := t.TempDir()
	runner := &e2eRunner{results: []engine.RunResult{
		{Tests: []engine.TestResult{
			{Package: "./github", Name: "TestAccFlaky", Status: provider.StatusFail, Output: []string{"403 API rate limit exceeded\n"}},
			{Package: "./github", Name: "TestAccBroken", Status: provider.StatusFail, Output: []string{"boom\n"}},
		}},
		{Tests: []engine.TestResult{
			{Package: "./github", Name: "TestAccFlaky", Status: provider.StatusPass},
		}},
	}}
	prov := newE2EProvider(provider.PreflightReport{})
	prov.GroupFunc = func(name string) string {
		if name == "TestAccFlaky" {
			return "actions"
		}
		return "repositories"
	}
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccFlaky", "TestAccBroken"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"e2e", "--json", "--repo-root", root}, &out, &errOut, d)
	if code != 1 {
		t.Fatalf("exit code = %d, want 1; stderr=%s", code, errOut.String())
	}

	events := decodeNDJSON(t, out.String())
	gotGroups := map[string]string{}
	var runSummary map[string]any
	for _, ev := range events {
		switch ev["type"] {
		case "test":
			gotGroups[ev["name"].(string)] = ev["group"].(string)
		case "summary":
			if ev["command"] == "run" {
				runSummary = ev
			}
		}
	}
	want := map[string]string{"TestAccFlaky": "actions", "TestAccBroken": "repositories"}
	if !reflect.DeepEqual(gotGroups, want) {
		t.Errorf("streamed test event groups = %v, want %v", gotGroups, want)
	}
	if runSummary == nil {
		t.Fatal("no embedded run summary in e2e output")
	}
	failures, _ := runSummary["failures"].([]any)
	if len(failures) != 1 {
		t.Fatalf("summary failures = %v, want exactly TestAccBroken", failures)
	}
	fm := failures[0].(map[string]any)
	if fm["name"] != "TestAccBroken" || fm["group"] != "repositories" {
		t.Errorf("failure entry = %v, want TestAccBroken in the repositories group", fm)
	}
}

func TestE2ETextOrphanDeltaNamesOnlyNewCleanupObligationAndRedactsMetadata(t *testing.T) {
	const secret = "ghp_RESOURCE_METADATA_SECRET_1234567890"
	root := t.TempDir()
	baseline := []provider.Resource{{
		Kind: "repository", Name: "tf-acc-test-existing", URL: "https://example.test/existing?token=" + secret,
	}}
	final := append(append([]provider.Resource{}, baseline...), provider.Resource{
		Kind: "issue_label", Name: "tf-acc-test-new", URL: "https://example.test/new?token=" + secret,
	})
	runner := &e2eRunner{results: []engine.RunResult{{Tests: []engine.TestResult{{
		Package: "./github", Name: "TestAccThing", Status: provider.StatusPass,
	}}}}}
	prov := newE2EProvider(provider.PreflightReport{})
	prov.Secrets = []string{"GITHUB_TOKEN"}
	prov.orphanSeq = [][]provider.Resource{baseline, final}
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccThing"}),
		cwd:       func() (string, error) { return root, nil },
		getenv: func(key string) string {
			if key == "GITHUB_TOKEN" {
				return secret
			}
			return ""
		},
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"e2e", "--mode", "organization", "--repo-root", root}, &out, &errOut, d)
	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s stdout=%s", code, errOut.String(), out.String())
	}

	text := out.String()
	for _, want := range []string{
		"Baseline: 1",
		"Pre-existing: 1",
		"Final: 2",
		"New (this run's cleanup obligation): 1",
		"Cleanup status: complete",
		"Preview command: terraform-provider-tester orphans --repo-root '" + root + "' --mode organization --run-delta",
		"Cleanup command: terraform-provider-tester sweep --repo-root '" + root + "' --mode organization --run-delta --confirm",
	} {
		if !strings.Contains(text, want) {
			t.Errorf("text output missing %q:\n%s", want, text)
		}
	}
	if got := strings.Count(text, "this run's cleanup obligation"); got != 1 {
		t.Errorf("cleanup-obligation label count = %d, want 1 on New only:\n%s", got, text)
	}
	if strings.Contains(text, secret) {
		t.Fatalf("text output leaked resource metadata secret:\n%s", text)
	}
	rawState, err := os.ReadFile(filepath.Join(root, ".pulsar-state.json"))
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(rawState), secret) || !strings.Contains(string(rawState), "***REDACTED***") {
		t.Fatalf("persisted orphan metadata was not redacted:\n%s", rawState)
	}
	if prov.sweepCalls != 0 {
		t.Fatalf("guided e2e sweep calls = %d, want 0", prov.sweepCalls)
	}
}

func TestE2ETextOrphanDeltaUnknownLabelsCountsUnproven(t *testing.T) {
	root := t.TempDir()
	baseline := []provider.Resource{{
		Kind: "repository", Name: "tf-acc-test-existing", URL: "https://example.test/existing",
	}}
	runner := &e2eRunner{results: []engine.RunResult{{Tests: []engine.TestResult{{
		Package: "./github", Name: "TestAccThing", Status: provider.StatusPass,
	}}}}}
	prov := newE2EProvider(provider.PreflightReport{})
	prov.orphanSeq = [][]provider.Resource{baseline}
	prov.orphanErrs = []error{nil, errors.New("final orphan scan failed")}
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccThing"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"e2e", "--mode", "organization", "--repo-root", root}, &out, &errOut, d)
	if code != 1 {
		t.Fatalf("exit code = %d, want 1; stderr=%s stdout=%s", code, errOut.String(), out.String())
	}
	text := out.String()
	for _, want := range []string{
		"Baseline: 1",
		"Pre-existing: unproven",
		"Final: unproven",
		"New (this run's cleanup obligation): unproven",
		"Cleanup status: unknown",
	} {
		if !strings.Contains(text, want) {
			t.Errorf("unknown orphan text missing %q:\n%s", want, text)
		}
	}
	for _, misleading := range []string{"Pre-existing: 0", "Final: 0", "New (this run's cleanup obligation): 0"} {
		if strings.Contains(text, misleading) {
			t.Errorf("unknown orphan text presents an unproven count as zero (%q):\n%s", misleading, text)
		}
	}
	for _, commandText := range []string{
		"Preview command:",
		"Cleanup command:",
		"terraform-provider-tester orphans",
		"terraform-provider-tester sweep",
	} {
		if strings.Contains(text, commandText) {
			t.Errorf("unknown orphan text advertises unavailable cleanup command %q:\n%s", commandText, text)
		}
	}
}

func TestE2EOrphanDeltaUnknownCleanupForcesNonzero(t *testing.T) {
	root := t.TempDir()
	baseline := []provider.Resource{{
		Kind: "repository", Name: "tf-acc-test-existing", URL: "https://example.test/existing",
	}}
	runner := &e2eRunner{results: []engine.RunResult{{Tests: []engine.TestResult{{
		Package: "./github", Name: "TestAccThing", Status: provider.StatusPass,
	}}}}}
	prov := newE2EProvider(provider.PreflightReport{})
	prov.orphanSeq = [][]provider.Resource{baseline}
	prov.orphanErrs = []error{nil, errors.New("final orphan scan failed")}
	d := deps{
		provider:  prov,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccThing"}),
		cwd:       func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"e2e", "--mode", "organization", "--json", "--repo-root", root}, &out, &errOut, d)
	if code != 1 {
		t.Fatalf("exit code = %d, want 1 for unknown cleanup status; stderr=%s stdout=%s", code, errOut.String(), out.String())
	}
	events := decodeNDJSON(t, out.String())
	if len(events) < 2 || events[len(events)-2]["type"] != "orphan_delta" {
		t.Fatalf("penultimate event = %v, want orphan_delta; events=%v", events[len(events)-2], events)
	}
	delta := events[len(events)-2]
	if delta["cleanup_status"] != "unknown" {
		t.Fatalf("orphan_delta cleanup_status = %v, want unknown", delta["cleanup_status"])
	}
	for field, want := range map[string]any{
		"baseline":     float64(1),
		"final":        float64(0),
		"pre_existing": float64(0),
		"new":          float64(0),
	} {
		if got := delta[field]; got != want {
			t.Errorf("orphan_delta[%q] = %v, want required numeric field %v; event=%v", field, got, want, delta)
		}
	}
	summary := lastEvent(t, events)
	if summary["cleanup_status"] != "unknown" || summary["exit_code"] != float64(1) {
		t.Fatalf("summary = %v, want cleanup_status unknown and exit_code 1", summary)
	}
	if summary["type"] != "summary" || summary["command"] != "e2e" {
		t.Fatalf("final event = %v, want e2e summary", summary)
	}
	for _, fields := range []map[string]any{delta, summary} {
		for _, field := range []string{"preview_command", "cleanup_command"} {
			if got := fields[field]; got != "" {
				t.Errorf("%s = %v, want empty when cleanup status is unknown; object=%v", field, got, fields)
			}
		}
	}
	for _, command := range []string{"terraform-provider-tester orphans", "terraform-provider-tester sweep"} {
		if strings.Contains(out.String(), command) {
			t.Errorf("unknown orphan NDJSON advertises unavailable command %q:\n%s", command, out.String())
		}
	}
	if prov.sweepCalls != 0 {
		t.Fatalf("guided e2e sweep calls = %d, want 0", prov.sweepCalls)
	}
}

func TestOrphanCleanupCommandsQuoteSingleQuoteInRepoRoot(t *testing.T) {
	root := "it's a root"
	quotedRoot := `'it'"'"'s a root'`
	if got, want := orphanPreviewCommand(root, "individual"),
		"terraform-provider-tester orphans --repo-root "+quotedRoot+" --mode individual --run-delta"; got != want {
		t.Errorf("orphanPreviewCommand() = %q, want %q", got, want)
	}
	if got, want := orphanCleanupCommand(root, "individual"),
		"terraform-provider-tester sweep --repo-root "+quotedRoot+" --mode individual --run-delta --confirm"; got != want {
		t.Errorf("orphanCleanupCommand() = %q, want %q", got, want)
	}
}
