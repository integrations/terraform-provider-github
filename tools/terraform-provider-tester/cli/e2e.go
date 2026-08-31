package cli

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/github/terraform-provider-tester/engine"
	"github.com/github/terraform-provider-tester/internal/redact"
	"github.com/github/terraform-provider-tester/provider"
)

// e2eDefaultMode is guided e2e's default --mode when the operator does not
// supply one. It also backs the env-file pre-scan fallback in cli.go, so the
// default mode used for PULSAR_<MODE>_* scoping always matches the default
// mode e2e itself will plan and preflight against.
const e2eDefaultMode = "organization"

const (
	guidedE2ESetupTimeout   = 60 * time.Second
	finalOrphanCheckTimeout = 30 * time.Second
)

var errCapturedOrphanBaselineUnavailable = errors.New(
	"captured orphan baseline is unavailable; cleanup status is unknown",
)

// e2eAllowedModes is the exact set of modes the guided e2e workflow
// supports. Guided e2e intentionally curates a safe subset of whatever modes
// the underlying provider technically supports: team and enterprise usage is
// rejected before any work happens, regardless of provider.Modes().
var e2eAllowedModes = map[string]bool{
	"anonymous":    true,
	"individual":   true,
	"organization": true,
}

// e2eCredentialedModes is the subset of e2eAllowedModes that actually has
// credentials. It gates guided-run baseline capture: the provider's Orphans
// implementation requires GITHUB_OWNER and issues authenticated queries scoped
// by GH_TEST_AUTH_MODE, so running it for anonymous mode is meaningless (there
// is no owner to scan and no credentialed test could have leaked a resource)
// and fails outright.
var e2eCredentialedModes = map[string]bool{
	"individual":   true,
	"organization": true,
}

type e2eRunTracker struct {
	calls   int
	results []engine.RunResult
	errs    []error
}

type trackedTestRunner struct {
	runner  testRunner
	tracker *e2eRunTracker
}

func (r *trackedTestRunner) Run(ctx context.Context, spec engine.RunSpec, sink func(engine.TestResult)) (engine.RunResult, error) {
	r.tracker.calls++
	result, err := r.runner.Run(ctx, spec, sink)
	r.tracker.results = append(r.tracker.results, result)
	r.tracker.errs = append(r.tracker.errs, err)
	return result, err
}

func trackE2ERuns(d deps, tracker *e2eRunTracker) deps {
	newRunner := d.newRunner
	d.newRunner = func(w io.Writer) testRunner {
		return &trackedTestRunner{runner: newRunner(w), tracker: tracker}
	}
	return d
}

type e2eSummary struct {
	SchemaVersion        int    `json:"schema_version"`
	Type                 string `json:"type"`
	Command              string `json:"command"`
	Mode                 string `json:"mode"`
	PreflightOK          bool   `json:"preflight_ok"`
	RunAttempted         bool   `json:"run_attempted"`
	InitialExitCode      *int   `json:"initial_exit_code"`
	RetryAttempted       bool   `json:"retry_attempted"`
	RetryExitCode        *int   `json:"retry_exit_code"`
	TriageAttempted      bool   `json:"triage_attempted"`
	OrphanCheckAttempted bool   `json:"orphan_check_attempted"`
	OrphanCheckExitCode  *int   `json:"orphan_check_exit_code"`
	DestructiveCleanup   bool   `json:"destructive_cleanup"`
	BaselineOrphans      int    `json:"baseline_orphans"`
	NewOrphans           int    `json:"new_orphans"`
	CleanupStatus        string `json:"cleanup_status"`
	PreviewCommand       string `json:"preview_command"`
	CleanupCommand       string `json:"cleanup_command"`
	ExitCode             int    `json:"exit_code"`
}

func intPtr(value int) *int {
	return &value
}

func shellQuoteArg(value string) string {
	return "'" + strings.ReplaceAll(value, "'", "'\"'\"'") + "'"
}

func orphanPreviewCommand(root, mode string) string {
	return fmt.Sprintf(
		"terraform-provider-tester orphans --repo-root %s --mode %s --run-delta",
		shellQuoteArg(root),
		mode,
	)
}

func orphanCleanupCommand(root, mode string) string {
	return fmt.Sprintf(
		"terraform-provider-tester sweep --repo-root %s --mode %s --run-delta --confirm",
		shellQuoteArg(root),
		mode,
	)
}

func redactOrphanResources(resources []provider.Resource, red *redact.Redactor) []provider.Resource {
	if resources == nil {
		return nil
	}
	if red == nil {
		red = redact.New(nil)
	}
	out := make([]provider.Resource, len(resources))
	for i, resource := range resources {
		out[i] = provider.Resource{
			Kind: red.String(resource.Kind),
			Name: red.String(resource.Name),
			URL:  red.String(resource.URL),
		}
	}
	return out
}

func orphanDeltaEvent(
	root string,
	accounting *engine.OrphanAccounting,
	red *redact.Redactor,
) (jsonOrphanDeltaEvent, bool) {
	if accounting == nil || accounting.CleanupStatus == engine.CleanupNotApplicable {
		return jsonOrphanDeltaEvent{}, false
	}
	if red == nil {
		red = redact.New(nil)
	}
	var preview, cleanup string
	if accounting.CleanupStatus != engine.CleanupUnknown {
		preview = red.String(orphanPreviewCommand(root, accounting.Mode))
		cleanup = red.String(orphanCleanupCommand(root, accounting.Mode))
	}
	return jsonOrphanDeltaEvent{
		SchemaVersion:  1,
		Type:           "orphan_delta",
		Mode:           accounting.Mode,
		Baseline:       len(accounting.Baseline),
		Final:          len(accounting.Final),
		PreExisting:    len(accounting.PreExisting),
		New:            len(accounting.New),
		CleanupStatus:  accounting.CleanupStatus,
		PreviewCommand: preview,
		CleanupCommand: cleanup,
	}, true
}

func emitE2EOrphanDelta(
	out io.Writer,
	jw *jsonWriter,
	format outputFormat,
	root string,
	accounting *engine.OrphanAccounting,
	red *redact.Redactor,
	summary *e2eSummary,
) error {
	if accounting == nil {
		return nil
	}
	summary.BaselineOrphans = len(accounting.Baseline)
	summary.NewOrphans = len(accounting.New)
	summary.CleanupStatus = accounting.CleanupStatus
	if accounting.CleanupStatus == engine.CleanupNotApplicable {
		return nil
	}
	event, _, err := emitOrphanDelta(out, jw, format, root, accounting, red)
	summary.PreviewCommand = event.PreviewCommand
	summary.CleanupCommand = event.CleanupCommand
	return err
}

func emitOrphanDelta(
	out io.Writer,
	jw *jsonWriter,
	format outputFormat,
	root string,
	accounting *engine.OrphanAccounting,
	red *redact.Redactor,
) (jsonOrphanDeltaEvent, bool, error) {
	event, ok := orphanDeltaEvent(root, accounting, red)
	if !ok {
		return jsonOrphanDeltaEvent{}, false, nil
	}
	if format == formatJSON {
		if jw == nil {
			jw = newJSONWriter(out)
		}
		return event, true, jw.Encode(event)
	}

	fmt.Fprintln(out, "Orphan accounting:")
	fmt.Fprintf(out, "  Baseline: %d\n", event.Baseline)
	switch event.CleanupStatus {
	case engine.CleanupUnknown:
		fmt.Fprintln(out, "  Pre-existing: unproven")
		fmt.Fprintln(out, "  Final: unproven")
		fmt.Fprintln(out, "  New (this run's cleanup obligation): unproven")
	case engine.CleanupBaselineOnly:
		fmt.Fprintln(out, "  Pre-existing: not captured")
		fmt.Fprintln(out, "  Final: not captured")
		fmt.Fprintln(out, "  New: 0 (no run attempted; no attributed cleanup obligation)")
	default:
		fmt.Fprintf(out, "  Pre-existing: %d\n", event.PreExisting)
		fmt.Fprintf(out, "  Final: %d\n", event.Final)
		fmt.Fprintf(out, "  New (this run's cleanup obligation): %d\n", event.New)
	}
	fmt.Fprintf(out, "  Cleanup status: %s\n", event.CleanupStatus)
	if event.PreviewCommand != "" {
		fmt.Fprintf(out, "  Preview command: %s\n", event.PreviewCommand)
	}
	if event.CleanupCommand != "" {
		fmt.Fprintf(out, "  Cleanup command: %s\n", event.CleanupCommand)
	}
	return event, true, nil
}

func emitResumeOrphanDelta(
	out io.Writer,
	jw *jsonWriter,
	format outputFormat,
	root string,
	accounting *engine.OrphanAccounting,
	red *redact.Redactor,
) error {
	_, _, err := emitOrphanDelta(out, jw, format, root, accounting, red)
	return err
}

func emitE2ESummary(out io.Writer, format outputFormat, summary e2eSummary) error {
	if format == formatJSON {
		return newJSONWriter(out).Encode(summary)
	}
	if !summary.RunAttempted {
		if _, err := fmt.Fprintf(out, "%s acceptance tests were not started\n", summary.Mode); err != nil {
			return err
		}
	} else if summary.ExitCode == 0 {
		if _, err := fmt.Fprintf(out, "%s acceptance-test workflow completed\n", summary.Mode); err != nil {
			return err
		}
	} else {
		if _, err := fmt.Fprintf(out, "%s acceptance-test workflow completed with failures\n", summary.Mode); err != nil {
			return err
		}
	}
	_, err := fmt.Fprintln(out, "destructive cleanup was not run")
	return err
}

func returnE2ESummary(out, errOut io.Writer, format outputFormat, summary e2eSummary) int {
	if err := emitE2ESummary(out, format, summary); err != nil {
		fmt.Fprintln(errOut, "writing summary:", err)
		return 1
	}
	return summary.ExitCode
}

// stateHasFailures reports whether the persisted state left by the run step
// has any failure. Results holds the FINAL (post-retry) status of every test
// - runWithSpecTriage persists retryResult.Final, see engine/retry.go - so
// this answers "does the run still have a failure after retries".
func stateHasFailures(root string) bool {
	state, err := engine.Load(filepath.Join(root, ".pulsar-state.json"))
	if err != nil {
		return false
	}
	return state.BuildFailed || state.PreRunFailed || len(state.FailedTopLevel()) > 0
}

// stateHadInitialFailures reports whether the run's INITIAL attempt (before
// any retries) had a failure. engine.RunWithRetries only ever appends a
// PersistFailure entry for a top-level test that failed on the first,
// pre-retry attempt - see classifyInitialOnly and the main retry loop's use
// of failedTopLevelResults(initial) in engine/retry.go - regardless of
// whether a later retry then passed. So a non-empty state.Failures here is
// authoritative proof the initial run failed at least one test. This lets
// e2e recover the initial-vs-final distinction from state runWithSpecTriage
// already computed and persisted, instead of re-running tests or
// duplicating RunWithRetries' retry bookkeeping in a second engine.
func stateHadInitialFailures(root string) bool {
	state, err := engine.Load(filepath.Join(root, ".pulsar-state.json"))
	if err != nil {
		return false
	}
	return state.BuildFailed || state.PreRunFailed || len(state.Failures) > 0
}

func lastRunCompleted(tracker *e2eRunTracker) bool {
	return len(tracker.errs) > 0 && tracker.errs[len(tracker.errs)-1] == nil
}

func invalidateFinalOrphanAccounting(accounting *engine.OrphanAccounting, mode string) *engine.OrphanAccounting {
	if accounting == nil {
		accounting = &engine.OrphanAccounting{Mode: mode}
	}
	accounting.Final = nil
	accounting.PreExisting = nil
	accounting.New = nil
	accounting.FinalCaptured = false
	accounting.CleanupStatus = engine.CleanupUnknown
	return accounting
}

func finalizeOrphanAccountingWithRedactor(
	parent context.Context,
	prov provider.Provider,
	mode string,
	state *engine.State,
	red *redact.Redactor,
) error {
	if state.Orphans == nil || !state.Orphans.BaselineCaptured {
		state.Orphans = invalidateFinalOrphanAccounting(state.Orphans, mode)
		return errCapturedOrphanBaselineUnavailable
	}
	state.Orphans.Baseline = redactOrphanResources(state.Orphans.Baseline, red)

	ctx, cancel := context.WithTimeout(parent, finalOrphanCheckTimeout)
	defer cancel()

	final, err := prov.Orphans(ctx, mode)
	if err != nil {
		state.Orphans = invalidateFinalOrphanAccounting(state.Orphans, mode)
		return err
	}
	final = redactOrphanResources(final, red)

	preExisting, newlyLeaked := engine.ComputeOrphanDelta(state.Orphans.Baseline, final)
	state.Orphans.Mode = mode
	state.Orphans.Final = final
	state.Orphans.PreExisting = preExisting
	state.Orphans.New = newlyLeaked
	state.Orphans.FinalCaptured = true
	state.Orphans.CleanupStatus = engine.CleanupComplete
	return nil
}

type finalOrphanContextFactory func() (context.Context, context.CancelFunc)

func freshFinalOrphanContext() (context.Context, context.CancelFunc) {
	return context.WithCancel(context.Background())
}

// runE2E is the guided end-to-end workflow: shared planning, plan-aware
// preflight, then one caller-locked orphan baseline and a single bounded
// retry+triage run over exactly plan.Eligible. It never runs a test the plan
// excludes and never sweeps.
func runE2E(args []string, out, errOut io.Writer, d deps) int {
	parent, stopRun := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stopRun()
	return runE2EContextWithFinalContext(parent, args, out, errOut, d, func() (context.Context, context.CancelFunc) {
		stopRun()
		return signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	})
}

func runE2EContext(parent context.Context, args []string, out, errOut io.Writer, d deps) int {
	return runE2EContextWithFinalContext(parent, args, out, errOut, d, freshFinalOrphanContext)
}

func runE2EContextWithFinalContext(
	parent context.Context,
	args []string,
	out, errOut io.Writer,
	d deps,
	finalContext finalOrphanContextFactory,
) int {
	fs := flag.NewFlagSet("e2e", flag.ContinueOnError)
	fs.SetOutput(errOut)
	var (
		repoRoot          string
		timeout           time.Duration
		mode              string
		allowUnclassified bool
	)
	format := formatText
	addJSONFlag(fs, &format)
	fs.StringVar(&mode, "mode", e2eDefaultMode, "guided mode: anonymous, individual, or organization")
	fs.BoolVar(&allowUnclassified, "allow-unclassified", false, "run unclassified tests with conservative requirements")
	fs.StringVar(&repoRoot, "repo-root", "", "repo root override")
	fs.DurationVar(&timeout, "timeout", 120*time.Minute, "test timeout")
	if err := fs.Parse(args); err != nil {
		return 2
	}

	// Guided e2e supports only anonymous, individual, and organization.
	// team/enterprise usage is rejected before any discovery, planning, or
	// preflight work happens.
	if !e2eAllowedModes[mode] {
		fmt.Fprintf(errOut, "e2e: --mode %q is not supported by guided e2e; use anonymous, individual, or organization\n", mode)
		return 2
	}

	root, err := d.resolveRoot(repoRoot)
	if err != nil {
		fmt.Fprintln(errOut, "resolving repo root:", err)
		return 1
	}

	getenv := getenvFunc(d)
	setenv := setenvFunc(d)
	// The provider test suite's TestMain still reads GH_TEST_AUTH_MODE from the
	// ambient environment, so guided e2e exports the selected mode for the
	// duration of the child go test run and restores it afterward.
	previousMode := getenv("GH_TEST_AUTH_MODE")
	if err := setenv("GH_TEST_AUTH_MODE", mode); err != nil {
		fmt.Fprintln(errOut, "setting auth mode:", err)
		return 1
	}
	defer func() {
		_ = setenv("GH_TEST_AUTH_MODE", previousMode)
	}()

	setupCtx, cancelSetup := context.WithTimeout(parent, guidedE2ESetupTimeout)
	defer cancelSetup()

	summary := e2eSummary{
		SchemaVersion: 1,
		Type:          "summary",
		Command:       "e2e",
		Mode:          mode,
	}

	prov := d.provider
	plan, _, planErr := planForCommand(setupCtx, d, root, planningOptions{
		Mode:              mode,
		AllowUnclassified: allowUnclassified,
	})
	if planErr == nil {
		planErr = setupCtx.Err()
	}
	if planErr != nil {
		fmt.Fprintln(errOut, "planning:", planErr)
		summary.ExitCode = planExitCode(planErr)
		return returnE2ESummary(out, errOut, format, summary)
	}

	red := redactorForProvider(prov, getenv)
	plan = redactPlan(plan, red)

	var jw *jsonWriter
	if format == formatJSON {
		jw = newJSONWriter(out)
	}
	if err := emitPlan(out, jw, format, plan); err != nil {
		fmt.Fprintln(errOut, "writing JSON:", err)
		return 1
	}

	// Plan-aware preflight, inlined (rather than delegated to the run/
	// preflight subcommands) so e2e can emit the same "check" JSON shape the
	// standalone preflight subcommand uses, and so the single run call below
	// can carry e2e's own fixed triage/retry configuration.
	requirements := provider.TestRequirements{
		Scopes:       plan.Scopes,
		Capabilities: plan.Capabilities,
		SideEffects:  plan.SideEffects,
	}
	report := prov.Preflight(setupCtx, mode, requirements)
	if format == formatJSON {
		for _, check := range report.Checks {
			if err := jw.Encode(jsonCheckEvent{
				SchemaVersion: 1,
				Type:          "check",
				Mode:          mode,
				Name:          check.Name,
				Status:        check.Status.String(),
				Detail:        red.String(check.Detail),
				Fix:           stringPtr(red.String(check.Fix)),
			}); err != nil {
				fmt.Fprintln(errOut, "writing JSON:", err)
				return 1
			}
		}
	} else {
		// Match the standalone preflight subcommand (runPreflightContext):
		// print the human checklist unconditionally in text mode, on both
		// success and failure, rather than only on failure. In JSON mode
		// every check was already reported as its own jsonCheckEvent above,
		// so printing the human checklist there too would put non-JSON text
		// on the stdout NDJSON stream.
		printPreflight(out, report, red)
	}

	summary.PreflightOK = report.OK()
	if setupErr := setupCtx.Err(); setupErr != nil {
		fmt.Fprintln(errOut, "guided setup:", setupErr)
		summary.ExitCode = 1
		return returnE2ESummary(out, errOut, format, summary)
	}
	if !report.OK() {
		summary.ExitCode = 1
		return returnE2ESummary(out, errOut, format, summary)
	}

	tracker := &e2eRunTracker{}
	d = trackE2ERuns(d, tracker)
	overallExit := 0
	statePath := filepath.Join(root, ".pulsar-state.json")
	lockPath := statePath + ".lock"
	release, lockErr := engine.AcquireLock(lockPath)
	if lockErr != nil {
		fmt.Fprintln(errOut, "acquiring lock:", lockErr)
		summary.ExitCode = 1
		return returnE2ESummary(out, errOut, format, summary)
	}
	defer release() //nolint:errcheck

	prevState := engine.State{Plan: &plan}
	if e2eCredentialedModes[mode] {
		summary.OrphanCheckAttempted = true
		baseline, err := prov.Orphans(setupCtx, mode)
		if err == nil {
			err = setupCtx.Err()
		}
		if err != nil {
			fmt.Fprintln(errOut, "capturing orphan baseline:", err)
			summary.OrphanCheckExitCode = intPtr(1)
			summary.CleanupStatus = engine.CleanupUnknown
			summary.ExitCode = 1
			return returnE2ESummary(out, errOut, format, summary)
		}
		baseline = redactOrphanResources(baseline, red)
		summary.OrphanCheckExitCode = intPtr(0)
		prevState.Orphans = &engine.OrphanAccounting{
			Mode:             mode,
			Baseline:         baseline,
			BaselineCaptured: true,
			CleanupStatus:    engine.CleanupBaselineOnly,
		}
	} else {
		prevState.Orphans = &engine.OrphanAccounting{
			Mode:          "anonymous",
			CleanupStatus: engine.CleanupNotApplicable,
		}
	}
	cancelSetup()
	if err := prevState.Save(statePath); err != nil {
		fmt.Fprintln(errOut, "saving state:", err)
		summary.ExitCode = 1
		return returnE2ESummary(out, errOut, format, summary)
	}
	accountingState := prevState

	// engine.RunPattern returns "" for zero eligible tests, and callers must
	// treat that as "run nothing", never "run all" (an empty go test -run
	// matches everything) - never running a test the plan excludes requires
	// this guard.
	if pattern := engine.RunPattern(plan.Eligible); pattern != "" {
		runArgs := []string{
			"--repo-root", root,
			"--mode", mode,
			"--timeout", timeout.String(),
			"--retries", "1",
			"--retry-policy", "retryable",
			"--retry-backoff", (30 * time.Second).String(),
			"--retry-max-backoff", (5 * time.Minute).String(),
		}
		if format == formatJSON {
			runArgs = append(runArgs, "--json")
		}
		runExit := runRunContext(parent, runArgs, out, errOut, d, runContextOptions{
			Preflight:     false,
			PreserveState: false,
			LockHeld:      true,
			Plan:          &plan,
		})

		summary.RunAttempted = tracker.calls > 0
		summary.RetryAttempted = tracker.calls > 1

		// initial_exit_code must reflect the FIRST attempt, which can differ
		// from the final/retry outcome (e.g. a retryable failure the retry
		// then clears) - runExit alone is runWithSpecTriage's single FINAL
		// exit code and cannot make that distinction. Recover the initial
		// outcome from the state runWithSpecTriage already persisted (see
		// stateHadInitialFailures), trusting it only when the last run call
		// actually completed: the same lastRunCompleted(tracker) guard used
		// for TriageAttempted below. An interrupted run persists partial state,
		// but it has no completed attempt from which to derive the initial
		// outcome, so fall back to runExit in that case.
		initialExit := runExit
		if lastRunCompleted(tracker) {
			initialExit = 0
			if stateHadInitialFailures(root) {
				initialExit = 1
			}
		}
		summary.InitialExitCode = intPtr(initialExit)
		if summary.RetryAttempted {
			// retry_exit_code represents the retry/final attempt, which is
			// exactly what runExit (exitFromResult applied to
			// retryResult.Final inside runWithSpecTriage) already is.
			summary.RetryExitCode = intPtr(runExit)
		}
		overallExit = runExit
	}

	if summary.RunAttempted && e2eCredentialedModes[mode] {
		finalParent, stopFinal := finalContext()
		finalState, err := engine.Load(statePath)
		if err != nil {
			fmt.Fprintln(errOut, "loading state for final orphan accounting:", err)
			finalState = prevState
			overallExit = 1
		}
		finalErr := finalizeOrphanAccountingWithRedactor(finalParent, prov, mode, &finalState, red)
		stopFinal()
		if finalErr != nil {
			fmt.Fprintln(errOut, "capturing final orphan state:", finalErr)
			summary.OrphanCheckExitCode = intPtr(1)
			overallExit = 1
		} else {
			summary.OrphanCheckExitCode = intPtr(0)
		}
		if err := finalState.Save(statePath); err != nil {
			fmt.Fprintln(errOut, "saving state:", err)
			finalState.Orphans = invalidateFinalOrphanAccounting(finalState.Orphans, mode)
			overallExit = 1
		}
		accountingState = finalState
	}

	summary.TriageAttempted = summary.RunAttempted && lastRunCompleted(tracker) && stateHasFailures(root)
	if err := emitE2EOrphanDelta(out, jw, format, root, accountingState.Orphans, red, &summary); err != nil {
		fmt.Fprintln(errOut, "writing JSON:", err)
		overallExit = 1
	}
	summary.ExitCode = overallExit
	return returnE2ESummary(out, errOut, format, summary)
}
