package cli

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/github/terraform-provider-tester/engine"
	"github.com/github/terraform-provider-tester/internal/redact"
	"github.com/github/terraform-provider-tester/provider"
	"github.com/github/terraform-provider-tester/tui"
)

// decideTUI reports whether to open the dashboard for a bare invocation.
// noTUI always wins; otherwise open when the output is a TTY or forced by env.
func decideTUI(isTTY, forceTTY, noTUI bool) bool {
	return !noTUI && (isTTY || forceTTY)
}

// testsForIntent resolves the ordered list of top-level test names for an
// intent, and an error when the intent cannot be resolved.
//
// RetryAllIntent and ResumeIntent delegate to the same retrySelection /
// resumeSelection helpers CLI retry/resume commands use (cli.go), so both
// surfaces intersect the persisted plan's Eligible set identically and share
// the exact state/plan-required error when prev.Plan is nil (a legacy,
// planless state). RetryGroupIntent applies the same persisted-eligibility
// intersection manually, scoped to the named group's current members.
//
// RetryTestIntent has no CLI equivalent: it is bound to whatever single test
// row the dashboard currently shows, so it only checks membership in the
// CURRENT eligible set (allNames) and never requires a persisted plan - this
// keeps "run this one test" usable on a brand-new root with no run history.
func testsForIntent(msg tea.Msg, groups []engine.Group, prev engine.State, allNames []string) ([]string, error) {
	switch m := msg.(type) {
	case tui.RetryGroupIntent:
		if prev.Plan == nil {
			return nil, fmt.Errorf("%s", planRequiredError)
		}
		var grpTests []string
		for _, g := range groups {
			if g.Name == m.Group {
				grpTests = g.Tests
				break
			}
		}
		if len(grpTests) == 0 {
			return nil, nil
		}
		// Restrict the group to tests the persisted plan still allows before
		// choosing between the failed subset and the whole-group fallback.
		eligible := intersectOrdered(prev.Plan.Eligible, grpTests)
		if len(eligible) == 0 {
			return nil, nil
		}
		failedSet := make(map[string]bool, len(prev.FailedTopLevel()))
		for _, f := range prev.FailedTopLevel() {
			failedSet[f] = true
		}
		var sel []string
		for _, t := range eligible {
			if failedSet[t] {
				sel = append(sel, t)
			}
		}
		if len(sel) == 0 {
			// No prior failures in this group; retry the whole eligible group.
			sel = eligible
		}
		return sel, nil

	case tui.RetryTestIntent:
		for _, n := range allNames {
			if n == m.Test {
				return []string{m.Test}, nil
			}
		}
		return nil, nil

	case tui.RetryAllIntent:
		return retrySelection(prev)

	case tui.ResumeIntent:
		return resumeSelection(prev)
	}
	return nil, nil
}

// labelForIntent returns a short human description of what a run intent will
// execute, shown in the TUI status line next to the run spinner.
// count is the number of top-level tests the intent resolves to, used only for
// the generic fallback.
func labelForIntent(msg tea.Msg, count int) string {
	switch m := msg.(type) {
	case tui.RetryGroupIntent:
		return m.Group
	case tui.RetryTestIntent:
		return m.Test
	case tui.RetryAllIntent:
		return "all failures"
	case tui.ResumeIntent:
		return "resume"
	}
	if count == 1 {
		return "1 test"
	}
	return fmt.Sprintf("%d tests", count)
}

// modeForIntent returns the run mode runIntent must use for msg. For the
// persisted-plan intents (RetryAllIntent, RetryGroupIntent, ResumeIntent) it
// is prev.Plan.Mode — mirroring CLI's resolveRetryResumeMode, which always
// derives an omitted --mode from the persisted plan for retry/resume — so
// ExtraEnv, the saved State.Mode, and State.Plan.Mode never disagree after a
// SwitchModeIntent has changed currentMode since the plan was persisted.
// RetryTestIntent (no CLI equivalent; bound to whatever single test row the
// dashboard currently shows) and any other intent keep currentMode
// unchanged. Callers only reach this after testsForIntent has already
// confirmed prev.Plan is non-nil for the three persisted-plan intents (it
// returns the shared planRequiredError otherwise, and runIntent returns
// before computing mode); the nil check here is kept as a defensive
// fallback so this helper can never panic if that invariant is ever broken.
func modeForIntent(msg tea.Msg, currentMode string, prev engine.State) string {
	switch msg.(type) {
	case tui.RetryAllIntent, tui.RetryGroupIntent, tui.ResumeIntent:
		if prev.Plan != nil {
			return prev.Plan.Mode
		}
	}
	return currentMode
}

// requirementsFromPlan converts an ExecutionPlan's derived requirements into
// the provider.TestRequirements shape Preflight expects.
func requirementsFromPlan(plan engine.ExecutionPlan) provider.TestRequirements {
	return provider.TestRequirements{
		Scopes:       plan.Scopes,
		Capabilities: plan.Capabilities,
		SideEffects:  plan.SideEffects,
	}
}

// eligibleGroups returns groups filtered down to only the tests present in
// eligible, preserving group order and each group's internal test order.
// A group left with no eligible tests is dropped entirely.
func eligibleGroups(groups []engine.Group, eligible []string) []engine.Group {
	allowed := make(map[string]bool, len(eligible))
	for _, n := range eligible {
		allowed[n] = true
	}
	out := make([]engine.Group, 0, len(groups))
	for _, g := range groups {
		var tests []string
		for _, t := range g.Tests {
			if allowed[t] {
				tests = append(tests, t)
			}
		}
		if len(tests) == 0 {
			continue
		}
		cp := g
		cp.Tests = tests
		out = append(out, cp)
	}
	return out
}

// planDashboard discovers tests, groups them, and builds an ExecutionPlan for
// mode, mirroring the interactive run path's planning convention (see cli.go's
// contextOpts.Preflight branch). allowUnclassified carries the operator's
// explicit opt-in, so the dashboard fails closed on an unknown test exactly
// like every non-TUI planning path instead of silently accepting it. The plan
// is passed through the shared redactPlan step before it is returned, so the
// GroupsMsg the TUI renders and the plan runIntent persists into
// .pulsar-state.json carry the same already-redacted exclusion text the
// run/preflight/e2e surfaces emit. It returns groups and names already
// narrowed to the plan's Eligible set, so callers (seed, SwitchModeIntent)
// never see a mode-incompatible test.
func planDashboard(ctx context.Context, list lister, prov provider.Provider, root, mode string, allowUnclassified bool, red *redact.Redactor) ([]engine.Group, []string, engine.ExecutionPlan, error) {
	plan, groups, err := planForCommand(ctx, deps{provider: prov, list: list}, root, planningOptions{
		Mode:              mode,
		AllowUnclassified: allowUnclassified,
	})
	if err != nil {
		return nil, nil, engine.ExecutionPlan{}, err
	}
	if red == nil {
		red = redact.New(nil)
	}
	plan = redactPlan(plan, red)
	return eligibleGroups(groups, plan.Eligible), plan.Eligible, plan, nil
}

// isFailureStatus reports whether s is a terminal failure status.
func isFailureStatus(s provider.Status) bool {
	return s == provider.StatusFail || s == provider.StatusPanic || s == provider.StatusTimeout
}

// resumePromptInfo decides whether to offer a resume prompt on open.
// show = (failed+notRun) > 0 AND prev has at least one recorded result (prior run exists).
func resumePromptInfo(prev engine.State, allNames []string) (failed, notRun int, show bool) {
	failed = len(prev.FailedTopLevel())
	notRun = len(prev.NotRun(allNames))
	show = len(prev.Results) > 0 && (failed+notRun) > 0
	return
}

// redactResult returns a copy of tr with Output redacted. Never mutates the input.
func redactResult(red *redact.Redactor, tr engine.TestResult) engine.TestResult {
	cp := tr
	if len(tr.Output) > 0 {
		cp.Output = red.Lines(tr.Output)
	}
	return cp
}

// redactRunResult returns a copy of res with every output-bearing field
// redacted (per-test Output, BuildOutput, PreRunOutput, RawLog). It never
// mutates the input, so the raw RunResult that crosses to the TUI via
// RunDoneMsg can never carry a secret value.
func redactRunResult(red *redact.Redactor, res engine.RunResult) engine.RunResult {
	cp := res
	if len(res.Tests) > 0 {
		cp.Tests = make([]engine.TestResult, len(res.Tests))
		for i, tr := range res.Tests {
			cp.Tests[i] = redactResult(red, tr)
		}
	}
	if len(res.BuildOutput) > 0 {
		cp.BuildOutput = red.Lines(res.BuildOutput)
	}
	if len(res.PreRunOutput) > 0 {
		cp.PreRunOutput = red.Lines(res.PreRunOutput)
	}
	if len(res.RawLog) > 0 {
		cp.RawLog = red.Lines(res.RawLog)
	}
	return cp
}

// dashboard is the stateful producer that drives the TUI.
type dashboard struct {
	prov      provider.Provider
	root      string
	mode      string
	owner     string
	getenv    func(string) string
	setenv    func(string, string) error
	statePath string
	newRunner func(io.Writer) testRunner
	list      lister
	red       *redact.Redactor
	// allowUnclassified is the operator's --allow-unclassified opt-in,
	// captured at launch and reused for every later re-plan (mode switches)
	// so the dashboard's fail-closed posture never drifts from the flag the
	// run was started with. It is immutable after construction.
	allowUnclassified bool
	// lifecycleCtx is canceled when the Bubble Tea program exits. Every
	// dashboard worker receives it instead of context.Background so shutdown
	// reaches in-flight discovery, preflight, and run work.
	lifecycleCtx    context.Context
	cancelLifecycle context.CancelFunc
	// cfgMu guards the configuration snapshot and the active switch's
	// cancellation ownership. A later dispatch cancels the previous switch
	// while holding this short critical section; no blocking I/O or send occurs
	// while cfgMu is held.
	cfgMu        sync.Mutex // guards mode, red, plan, groups, allNames, cfgSeq, and switch cancellation
	cfgSeq       uint64     // monotonic SwitchModeIntent dispatch sequence; see beginSwitchMode
	switchSeq    uint64
	switchCtx    context.Context
	cancelSwitch context.CancelFunc
	// sendMu serializes the freshness check, state commit, and message sends
	// of every mode-tied producer (seed, applySwitchMode, handleConfigIntent)
	// — see withLatestDispatch. It is deliberately a SEPARATE lock from
	// cfgMu, never held during provider network I/O, and never acquired on
	// the Bubble Tea event loop's path: db.send is tea.Program.Send, whose
	// message channel is UNBUFFERED, so a send blocks until the event loop
	// drains it, and that event loop synchronously calls beginSwitchMode
	// (which takes cfgMu) from tui.Model.Update. Holding cfgMu across a send
	// would therefore deadlock the program; holding sendMu cannot, because
	// nothing Update reaches ever waits on sendMu.
	sendMu   sync.Mutex
	plan     *engine.ExecutionPlan
	groups   []engine.Group
	allNames []string
	send     func(tea.Msg) // seam: tea.Program.Send in prod; a capture func in tests
	busy     atomic.Bool   // single-run guard
	opGate   atomic.Bool   // serializes all user-triggered dashboard operations
	logSink  io.Writer     // where runner streams raw text; io.Discard in prod
	wg       sync.WaitGroup

	userConfigDir       func() (string, error)
	now                 func() time.Time
	newKnownIssueLister func(string) (knownIssueLister, error)
	newIssueRegistry    func(triageOptions) (issueRegistry, error)
	newIssueFiler       func(string) (issueFiler, error)
}

// workerContext returns the dashboard lifetime context when available. Direct
// unit callers can still supply their own context for dashboards constructed
// without a lifecycle.
func (db *dashboard) workerContext(ctx context.Context) context.Context {
	// lifecycleCtx is set once, at construction, and never reassigned, so
	// reading it needs no lock.
	if db.lifecycleCtx != nil {
		return db.lifecycleCtx
	}
	return ctx
}

// seed runs preflight, emits opening messages: PreflightMsg, GroupsMsg, one
// TestUpdateMsg per persisted result (failed top-level tests get their log
// content loaded from .pulsar-failures/<slug>.log when the file exists), and
// ResumePromptMsg when resumePromptInfo says show. db.plan must already be
// set (runDashboard builds it before constructing db); Preflight is called
// with the requirements the plan derived, and GroupsMsg carries that same
// plan alongside the already eligible-filtered groups.
//
// seed runs concurrently with the TUI (runDashboard starts it in its own
// goroutine), so a SwitchModeIntent can land while seed is still inside
// Preflight, which for a credentialed mode makes real network calls. seed
// therefore takes ONE snapshot of cfgSeq/mode/plan/groups under cfgMu up
// front - the same discipline runIntent uses - so the mode and plan it
// preflights are always a consistent pair, and it observes the SAME staleness
// rule as a superseded applySwitchMode: once a newer SwitchModeIntent has been
// dispatched, seed sends no mode-tied message at all. The check and the pair
// of sends happen together inside withLatestDispatch, so a newer dispatch can
// never slip its own pair between them. Without that rule seed's launch-mode
// PreflightMsg/GroupsMsg land after the switch's pair, and tui.Model replaces
// mode, checks, and groups wholesale on each message, so the operator would be
// left looking at the launch mode while db.mode - and therefore every run the
// dashboard starts - already uses the switched mode.
//
// The persisted-result restore and the resume prompt are mode-independent
// (they come from .pulsar-state.json) and are still emitted either way; the
// prompt's not-run count is computed from the eligible set currently in
// effect, which is the snapshot's whenever no switch intervened.
func (db *dashboard) seed(ctx context.Context) {
	ctx = db.workerContext(ctx)
	db.cfgMu.Lock()
	seq := db.cfgSeq
	mode := db.mode
	plan := db.plan
	groups := db.groups
	db.cfgMu.Unlock()

	report := db.prov.Preflight(ctx, mode, requirementsFromPlan(*plan))

	db.withLatestDispatch(
		func() bool { return db.admittedLocked(ctx, seq) },
		func() []tea.Msg {
			return []tea.Msg{
				tui.PreflightMsg{Report: report},
				tui.GroupsMsg{Groups: groups, Plan: plan},
			}
		},
	)

	prev, _ := engine.Load(db.statePath)
	failDir := filepath.Join(filepath.Dir(db.statePath), ".pulsar-failures")
	cachePath, cacheErr := db.knownIssuesPath()
	if cacheErr != nil {
		db.sendOperationError(ctx, "known-issues-cache", cacheErr)
	}
	db.sendIfActive(ctx, tui.TriageLoadedMsg{
		Failures:       append([]engine.PersistFailure(nil), prev.Failures...),
		CacheAvailable: cacheErr == nil && cacheFileExists(cachePath),
	})
	for _, r := range prev.Results {
		status, _ := provider.ParseStatus(r.Status)
		result := engine.TestResult{
			Package: r.Package,
			Name:    r.Test,
			Sub:     r.Sub,
			Status:  status,
			Elapsed: r.Elapsed,
		}
		// Load prior log content for failed top-level tests when the file exists.
		if r.Sub == "" && isFailureStatus(status) {
			logPath := engine.FailureLogPath(failDir, r.Package, r.Test)
			if data, err := os.ReadFile(logPath); err == nil {
				result.Output = strings.Split(string(data), "\n")
			} else if !os.IsNotExist(err) {
				db.sendIfActive(ctx, tui.ErrMsg{Err: fmt.Errorf("reading failure log %q: %w", logPath, err)})
			}
		}
		db.sendIfActive(ctx, tui.TestUpdateMsg{Result: result})
	}

	db.cfgMu.Lock()
	allNames := db.allNames
	db.cfgMu.Unlock()
	failed, notRun, show := resumePromptInfo(prev, allNames)
	if show {
		db.sendIfActive(ctx, tui.ResumePromptMsg{
			When:   prev.RunAt,
			Failed: failed,
			NotRun: notRun,
		})
	}
}

// runIntent executes one intent: guards against concurrent runs, resolves the
// eligible test selection, acquires the state lock, streams REDACTED results
// to the TUI, and emits RunDoneMsg. Secret values in test output are never
// forwarded. A nil persisted plan or a planning-derived error stops before
// the runner is ever invoked (see testsForIntent).
func (db *dashboard) runIntent(ctx context.Context, msg tea.Msg) {
	ctx = db.workerContext(ctx)
	if !db.busy.CompareAndSwap(false, true) {
		db.sendIfActive(ctx, tui.ErrMsg{Err: errors.New("a run is already in progress")})
		return
	}
	defer db.busy.Store(false)

	// Snapshot config under cfgMu so a concurrent config intent (which may
	// rebuild plan/groups/allNames on a mode switch) cannot affect the
	// in-flight run. The subprocess is launched, and eligibility resolved,
	// from this snapshot throughout.
	db.cfgMu.Lock()
	currentMode := db.mode
	red := db.red
	plan := db.plan
	groups := db.groups
	allNames := db.allNames
	db.cfgMu.Unlock()

	// Acquire the state lock BEFORE reading state so the snapshot we merge
	// against is the one protected by the lock (no stale-merge TOCTOU window).
	release, err := engine.AcquireLock(db.statePath + ".lock")
	if err != nil {
		db.sendIfActive(ctx, tui.ErrMsg{Err: fmt.Errorf("acquiring lock: %w", err)})
		return
	}
	defer release() //nolint:errcheck

	// Load fresh state so retries see the latest results.
	prev, _ := engine.Load(db.statePath)

	tests, testsErr := testsForIntent(msg, groups, prev, allNames)
	if testsErr != nil {
		// e.g. state/plan-required for a legacy, planless state: never reach
		// the runner.
		db.sendIfActive(ctx, tui.ErrMsg{Err: testsErr})
		return
	}
	if len(tests) == 0 {
		db.sendIfActive(ctx, tui.ErrMsg{Err: errors.New("nothing to run")})
		return
	}
	pattern := engine.RunPattern(tests)

	// Resolve the mode to actually run under: for the persisted-plan intents
	// this is prev.Plan.Mode (never currentMode, which may have moved on via
	// a SwitchModeIntent since prev was last saved); RetryTestIntent keeps
	// currentMode. See modeForIntent.
	mode := modeForIntent(msg, currentMode, prev)

	spec := engine.RunSpec{
		Dir:      db.root,
		Packages: db.prov.TestPackages(),
		Pattern:  pattern,
		ExtraEnv: buildExtraEnv(mode, db.prov.EnvFor(mode), dashboardGetenv(db)),
		// Timeout zero → engine defaults to 120m; AllowSensitiveLogs stays false.
	}

	// Total counts only the eligible selected tests (tests already reflects
	// the persisted-plan/current-eligible intersection from testsForIntent).
	total := len(tests)
	db.sendIfActive(ctx, tui.RunStartedMsg{Total: total, Label: labelForIntent(msg, total)})

	// Carry the persisted plan forward unchanged, exactly mirroring CLI's
	// retry/resume convention (runWithSpec sets cur.Plan = prev.Plan). Fall
	// back to the dashboard's current plan when there was none to carry
	// forward (e.g. a first-ever run bootstrapped via RetryTestIntent, the
	// one intent that does not require a persisted plan) so a later
	// RetryAllIntent/RetryGroupIntent/ResumeIntent has an Eligible set to
	// intersect against instead of being permanently stuck on plan-required
	// — or when prev.Plan is stale for a DIFFERENT mode than the one this
	// run actually used (mode, above): carrying it forward would save
	// State.Plan.Mode != State.Mode, the same bug class this fix closes.
	// db.plan (plan, above) always has Mode == db.mode by construction (both
	// are only ever updated together, atomically, under cfgMu), so it is
	// always self-consistent with a currentMode-tied run (RetryTestIntent).
	curPlan := prev.Plan
	if curPlan == nil || curPlan.Mode != mode {
		curPlan = plan
	}
	cur := engine.State{
		Provider: db.prov.Name(),
		Mode:     mode,
		RunAt:    time.Now(),
		History:  prev.History,
		Plan:     curPlan,
		Orphans:  prev.Orphans,
	}

	runner := db.newRunner(db.logSink)
	res, runErr := runner.Run(ctx, spec, func(tr engine.TestResult) {
		cur.Record(tr)
		// SECURITY: only the redacted copy is ever sent outward.
		db.sendIfActive(ctx, tui.TestUpdateMsg{Result: redactResult(red, tr)})
	})
	// Cancellation (dashboard lifetime ending, or a run made moot mid-stream)
	// must still merge and persist whatever partial results the runner
	// already streamed through cur.Record above — the operator's disk state
	// must never regress behind what the TUI already showed for this run.
	// Only the OUTWARD completion messaging (the run-error report for a
	// cancellation, and RunDoneMsg) is suppressed: sendIfActive already gates
	// every send below on ctx being live, so nothing is emitted to a
	// dashboard that has moved on, but the state/log side effects below
	// always run to completion.
	canceled := ctx.Err() != nil
	if !canceled && runErr != nil {
		db.sendIfActive(ctx, tui.ErrMsg{Err: fmt.Errorf("run error: %w", runErr)})
	}

	failDir := filepath.Join(filepath.Dir(db.statePath), ".pulsar-failures")
	cur.Results = mergeResults(prev.Results, cur.Results)
	persistSuiteFailureState(&cur, failDir, res, res, red, func(action string, err error) {
		db.sendIfActive(ctx, tui.ErrMsg{Err: fmt.Errorf("%s: %w", action, err)})
	})
	if _, err := engine.WriteFailures(failDir, res, red); err != nil {
		db.sendIfActive(ctx, tui.ErrMsg{Err: fmt.Errorf("writing failure logs: %w", err)})
	}

	cur.Failures = preserveIssueMetadata(prev.Failures, reconstructFailures(db.root, cur))
	cachePath, pathErr := db.knownIssuesPath()
	triageResult := triageStateResult{
		Failures:       cur.Failures,
		CacheAvailable: cacheFileExists(cachePath),
	}
	if pathErr == nil {
		result, err := db.triageService(red).apply(ctx, &cur, triageOptions{
			IssuesRepo:         defaultIssuesRepo,
			KnownIssuesPath:    cachePath,
			KnownIssuesOffline: true,
		}, false)
		if err != nil {
			db.sendIfActive(ctx, tui.ErrMsg{Err: fmt.Errorf("triage: %w", err)})
		} else {
			triageResult = result
		}
	} else {
		db.sendIfActive(ctx, tui.ErrMsg{Err: fmt.Errorf("triage: %w", pathErr)})
	}

	if saveErr := cur.Save(db.statePath); saveErr != nil {
		db.sendIfActive(ctx, tui.ErrMsg{Err: fmt.Errorf("saving state: %w", saveErr)})
	}
	db.sendIfActive(ctx, tui.TriageLoadedMsg{
		Failures:       triageResult.Failures,
		CacheAvailable: triageResult.CacheAvailable,
	})

	if !canceled {
		db.sendIfActive(ctx, tui.RunDoneMsg{Result: redactRunResult(red, res)})
	}
}

// beginSwitchMode assigns and returns the next monotonic SwitchModeIntent
// dispatch sequence (guarded by cfgMu, alongside mode/plan/groups/allNames).
// It must be called synchronously, at dispatch time, BEFORE the
// corresponding applySwitchMode is scheduled — runDashboard's exec closure
// does this in its own synchronous body, never inside the goroutine it
// spawns, because tui.Model.Update calls exec synchronously (see
// tui/update.go's m.exec call). That guarantees the sequence this produces
// always matches true user-dispatch order, regardless of goroutine
// scheduling or how long each switch's planning takes.
//
// It runs ON the Bubble Tea event loop, so it must never block on anything a
// producer holds while sending: it takes cfgMu only, for a short CPU-bound
// critical section, and never sendMu.
func (db *dashboard) beginSwitchMode() uint64 {
	db.cfgMu.Lock()
	if db.cancelSwitch != nil {
		db.cancelSwitch()
	}
	db.cfgSeq++
	seq := db.cfgSeq
	parent := db.lifecycleCtx
	if parent == nil {
		parent = context.Background()
	}
	db.switchCtx, db.cancelSwitch = context.WithCancel(parent)
	db.switchSeq = seq
	db.cfgMu.Unlock()
	return seq
}

// switchContext returns the active context owned by seq. A switch that has
// already been superseded before its goroutine starts has no work to perform.
func (db *dashboard) switchContext(seq uint64) (context.Context, bool) {
	db.cfgMu.Lock()
	defer db.cfgMu.Unlock()
	if db.switchSeq != seq || db.switchCtx == nil {
		return nil, false
	}
	return db.switchCtx, true
}

// clearSwitchContext drops completion ownership only when it still belongs to
// seq. A stale completion must never cancel or clear a newer switch.
func (db *dashboard) clearSwitchContext(seq uint64) {
	db.cfgMu.Lock()
	defer db.cfgMu.Unlock()
	if db.switchSeq != seq {
		return
	}
	if db.cancelSwitch != nil {
		db.cancelSwitch()
	}
	db.switchSeq = 0
	db.switchCtx = nil
	db.cancelSwitch = nil
}

// configSeqUnchanged reports whether cfgSeq is still seq, i.e. whether no
// newer SwitchModeIntent has been dispatched since the caller snapshotted it.
// It is a cheap, ADVISORY staleness check a producer may use to bail out of
// expensive work (e.g. skipping a superseded switch's Preflight call) before
// ever reaching withLatestDispatch. It is NOT the authoritative admission
// decision — see admittedLocked and withLatestDispatch — because cfgSeq can
// change again between this call returning and that producer's eventual
// commit.
func (db *dashboard) configSeqUnchanged(seq uint64) bool {
	db.cfgMu.Lock()
	defer db.cfgMu.Unlock()
	return db.cfgSeq == seq
}

// admittedLocked is the shared freshness predicate behind every mode-tied
// producer's (seed's, applySwitchMode's, sendConfigPreflight's) latest-
// dispatch rule: seq is admitted only while ctx is not canceled and no newer
// SwitchModeIntent has been dispatched since seq was snapshotted. Callers
// must hold cfgMu, and must decide admission and perform any accompanying
// state mutation in that SAME critical section (see withLatestDispatch) so a
// later beginSwitchMode — which only needs cfgMu, not sendMu — can never land
// between the decision and the mutation.
func (db *dashboard) admittedLocked(ctx context.Context, seq uint64) bool {
	return ctx.Err() == nil && db.cfgSeq == seq
}

// withLatestDispatch admits one producer's dispatch sequence exactly once:
// decide runs under cfgMu and reports whether this producer is still the
// latest dispatched configuration, performing any accompanying commit (e.g.
// applySwitchMode's mode/plan/groups/allNames write) in that same critical
// section. Only when admitted does withLatestDispatch send every message
// publish returns, IN ORDER, with NO further freshness or cancellation check
// in between — once a producer is admitted, its whole pair goes out
// together, so a later dispatch's own pair can never interleave with, or
// tear, this one, and a switch that commits always publishes a coherent
// result. A superseded or canceled producer commits nothing and sends
// nothing.
//
// Holding sendMu across the ENTIRE call — the admission decision, the
// commit, and every send — is what makes the latest-dispatch rule atomic. A
// later dispatch's beginSwitchMode only needs cfgMu (never sendMu; see the
// sendMu field comment for why that must stay true to avoid deadlocking the
// Bubble Tea event loop), so it can bump cfgSeq and cancel this producer's
// context at any moment; but once cfgMu's brief critical section above has
// admitted this producer, nothing changes that decision. Rechecking freshness
// again before each individual send re-opened exactly that window — a later
// dispatch landing between the pair's two sends — and produced a torn
// PreflightMsg/GroupsMsg pair or, worse, a committed config with zero
// messages sent at all.
func (db *dashboard) withLatestDispatch(decide func() bool, publish func() []tea.Msg) {
	db.sendMu.Lock()
	defer db.sendMu.Unlock()

	db.cfgMu.Lock()
	admitted := decide()
	db.cfgMu.Unlock()

	if !admitted {
		return
	}
	for _, msg := range publish() {
		db.send(msg)
	}
}

// sendIfActive prevents background workers from enqueueing a result after the
// dashboard lifetime has ended.
func (db *dashboard) sendIfActive(ctx context.Context, msg tea.Msg) {
	if ctx == nil || ctx.Err() == nil {
		db.send(msg)
	}
}

// applySwitchMode plans, preflights, and — only if still current — applies
// the SwitchModeIntent m dispatched with sequence seq (from beginSwitchMode).
// Re-discovery, planning, and the provider Preflight call all happen outside
// both locks so this never blocks a concurrent runIntent/handleConfigIntent
// config snapshot during discovery or network I/O.
//
// The commit and the PreflightMsg/GroupsMsg pair then happen together, in a
// single admission decision, inside withLatestDispatch: a LATER
// SwitchModeIntent already dispatched (via beginSwitchMode) means this
// completion has been superseded, so it mutates no dashboard state and sends
// nothing at all — no ErrMsg, no PreflightMsg, no GroupsMsg — whether its own
// planning succeeded or failed. A slow, superseded switch can therefore never
// clobber a newer one's mode/plan/groups/allNames, and an operator is never
// shown a confusing message about a mode switch they have already moved
// past. Once admitted, the mutation and BOTH sends happen without rechecking
// in between, so a newer switch can never tear this pair or interleave its
// own between the two messages (see withLatestDispatch).
//
// Dashboard state is committed only AFTER Preflight returns, so a switch is
// never half-applied: for the whole duration of a slow, credentialed
// Preflight, db.mode/plan/groups/allNames remain the previous mode's
// self-consistent set, and a run started meanwhile uses that coherent pair.
func (db *dashboard) applySwitchMode(_ context.Context, m tui.SwitchModeIntent, seq uint64) {
	ctx, ok := db.switchContext(seq)
	if !ok {
		return
	}
	defer db.clearSwitchContext(seq)

	// Re-discover and re-plan fresh for the new mode (not a re-filter of the
	// already-narrowed db.groups) so a second mode switch still sees the
	// full universe of tests, not just the previous mode's subset. The
	// launch-time --allow-unclassified posture and the current redactor are
	// carried into the re-plan so a mode switch never widens what the
	// dashboard accepts, and the new plan is redacted like the first one.
	db.cfgMu.Lock()
	red := db.red
	db.cfgMu.Unlock()
	groups, allNames, plan, err := planDashboard(ctx, db.list, db.prov, db.root, m.Mode, db.allowUnclassified, red)

	if err != nil {
		// Planning failure: never call Preflight, never reach the runner.
		// The ErrMsg is still gated on this dispatch still being the latest.
		db.withLatestDispatch(
			func() bool { return db.admittedLocked(ctx, seq) },
			func() []tea.Msg {
				return []tea.Msg{tui.ErrMsg{Err: fmt.Errorf("planning: %w", err)}}
			},
		)
		return
	}

	// Cheap advisory bail-out: skip a superseded switch's Preflight network
	// call entirely. withLatestDispatch below is still the authoritative
	// check, since a newer switch can be dispatched at any moment.
	if !db.configSeqUnchanged(seq) {
		return
	}

	report := db.prov.Preflight(ctx, m.Mode, requirementsFromPlan(plan))
	envVars := db.prov.EnvFor(m.Mode)

	db.withLatestDispatch(
		func() bool {
			if !db.admittedLocked(ctx, seq) {
				return false
			}
			db.mode = m.Mode
			db.plan = &plan
			db.groups = groups
			db.allNames = allNames
			return true
		},
		func() []tea.Msg {
			return []tea.Msg{
				tui.PreflightMsg{Report: report, EnvVars: envVars},
				tui.GroupsMsg{Groups: groups, Plan: &plan},
			}
		},
	)
}

// handleConfigIntent handles the non-run, non-mode-switch config intents:
// SetEnvVarIntent and PreflightIntent. (SwitchModeIntent is handled by
// beginSwitchMode/applySwitchMode instead, so its dispatch sequence can be
// assigned synchronously at dispatch time; see runDashboard's exec.)
//
// Mode/plan reads and redactor updates are guarded by cfgMu; Preflight
// happens outside every lock to avoid holding one during network I/O. The
// cfgSeq is snapshotted together with mode/plan, and the resulting report is
// emitted through withLatestDispatch, so a report computed for the
// pre-switch mode is suppressed once a later SwitchModeIntent has been
// dispatched — the same latest-dispatch rule seed and applySwitchMode obey.
// A SetEnvVarIntent's env write and redactor refresh still happen either
// way; only the now-stale report is dropped.
func (db *dashboard) handleConfigIntent(ctx context.Context, msg tea.Msg) {
	ctx = db.workerContext(ctx)
	switch m := msg.(type) {
	case tui.SetEnvVarIntent:
		secretKeys := make(map[string]bool, len(db.prov.SecretEnvKeys()))
		for _, k := range db.prov.SecretEnvKeys() {
			secretKeys[k] = true
		}
		if secretKeys[m.Key] {
			db.sendIfActive(ctx, tui.ErrMsg{Err: fmt.Errorf("refusing to set secret env var %q from the editor", m.Key)})
			return
		}
		if err := dashboardSetenv(db)(m.Key, m.Value); err != nil {
			db.sendIfActive(ctx, tui.ErrMsg{Err: fmt.Errorf("setting %s: %w", m.Key, err)})
			return
		}
		db.refreshRedactor()
		db.sendConfigPreflight(ctx)
	case tui.PreflightIntent:
		db.sendConfigPreflight(ctx)
	}
}

// sendConfigPreflight snapshots the current dispatch sequence together with
// the mode and plan it belongs to, runs Preflight outside every lock, and
// emits the report only if that dispatch is still the latest one.
//
// cfgSeq alone cannot detect every stale report: beginSwitchMode bumps cfgSeq
// synchronously, at dispatch time, BEFORE its applySwitchMode goroutine ever
// runs Preflight or commits. So a config intent that snapshots mode/cfgSeq
// while that switch is dispatched but still in flight observes the SAME
// cfgSeq the switch itself owns, together with the pre-switch mode. If the
// switch then commits (mode advances, its own PreflightMsg/GroupsMsg pair is
// sent) before this config intent's Preflight call returns, cfgSeq is still
// unchanged — nothing bumps it again until a NEWER SwitchModeIntent is
// dispatched — so withLatestDispatch's admittedLocked check (cfgSeq plus
// ctx cancellation) alone would still accept this report and append it,
// stale and for the wrong mode, after the switch's pair.
//
// To close that window, withLatestDispatch's decide closure re-reads db.mode
// under cfgMu (the same critical section admittedLocked runs in, so the
// lock order stays sendMu → cfgMu, exactly like applySwitchMode's own
// commit) and compares it against the mode this report was computed for. A
// mismatch means some switch sharing this cfgSeq has already committed a
// different mode, so the report is dropped before it is ever published. When
// no switch intervenes, or an intervening switch fails to plan/preflight
// (mode never changes), the committed mode still matches and the report is
// sent exactly as before.
func (db *dashboard) sendConfigPreflight(ctx context.Context) {
	db.cfgMu.Lock()
	seq := db.cfgSeq
	mode := db.mode
	plan := db.plan
	db.cfgMu.Unlock()

	report := db.prov.Preflight(ctx, mode, requirementsFromPlan(*plan))
	envVars := db.prov.EnvFor(mode)

	db.withLatestDispatch(
		func() bool {
			if !db.admittedLocked(ctx, seq) {
				return false
			}
			if db.mode != mode {
				// A mode switch sharing this same cfgSeq has already
				// committed a different mode; this report was computed for
				// the mode it superseded, so drop it instead of rendering it
				// after that switch's own pair.
				return false
			}
			return true
		},
		func() []tea.Msg {
			return []tea.Msg{tui.PreflightMsg{Report: report, EnvVars: envVars}}
		},
	)
}

// refreshRedactor rebuilds db.red from the current values of all secret env
// keys. Call after any secret env var may have changed so later redaction
// covers the new value.
func (db *dashboard) refreshRedactor() {
	var secretVals []string
	getenv := dashboardGetenv(db)
	for _, k := range db.prov.SecretEnvKeys() {
		if v := getenv(k); v != "" {
			secretVals = append(secretVals, v)
		}
	}
	db.cfgMu.Lock()
	db.red = redact.New(secretVals)
	db.cfgMu.Unlock()
}

func dashboardGetenv(db *dashboard) func(string) string {
	if db.getenv != nil {
		return db.getenv
	}
	return func(string) string { return "" }
}

func dashboardSetenv(db *dashboard) func(string, string) error {
	if db.setenv != nil {
		return db.setenv
	}
	return func(string, string) error { return nil }
}

func runDashboard(d deps, root, mode, owner string, ascii, allowUnclassified bool, out, errOut io.Writer) int {
	prov := d.provider
	getenv := getenvFunc(d)
	setenv := setenvFunc(d)

	// Build redactor from secret values (mirrors cli.go:540-548 pattern).
	var secretVals []string
	for _, key := range prov.SecretEnvKeys() {
		if val := getenv(key); val != "" {
			secretVals = append(secretVals, val)
		}
	}
	red := redact.New(secretVals)

	// Plan+group up front (mirrors the interactive run path's planning
	// convention) so seed/runIntent share the same eligible-only view. A
	// planning failure (e.g. the provider package failing to compile) aborts
	// here rather than opening an empty dashboard with no explanation.
	ctx, cancelLifecycle := context.WithCancel(context.Background())
	groups, allNames, plan, err := planDashboard(ctx, d.list, prov, root, mode, allowUnclassified, red)
	if err != nil {
		cancelLifecycle()
		fmt.Fprintln(errOut, err)
		return 1
	}

	statePath := filepath.Join(root, ".pulsar-state.json")

	db := &dashboard{
		prov:                prov,
		root:                root,
		mode:                mode,
		owner:               owner,
		getenv:              getenv,
		setenv:              setenv,
		statePath:           statePath,
		newRunner:           d.newRunner,
		list:                d.list,
		red:                 red,
		allowUnclassified:   allowUnclassified,
		lifecycleCtx:        ctx,
		cancelLifecycle:     cancelLifecycle,
		plan:                &plan,
		groups:              groups,
		allNames:            allNames,
		logSink:             io.Discard,
		userConfigDir:       d.userConfigDir,
		now:                 time.Now,
		newKnownIssueLister: d.newKnownIssueLister,
		newIssueRegistry:    d.newIssueRegistry,
		newIssueFiler:       d.newIssueFiler,
	}

	var p *tea.Program
	exec := func(msg tea.Msg) tea.Cmd {
		if sm, ok := msg.(tui.SwitchModeIntent); ok {
			// Assign the dispatch sequence HERE, synchronously — exec is
			// called synchronously by tui.Model.Update (see tui/update.go's
			// m.exec call) — so it reflects true dispatch order regardless
			// of goroutine scheduling or how long each switch's planning
			// takes. See beginSwitchMode/applySwitchMode.
			seq := db.beginSwitchMode()
			finish := db.beginTrackedOperation()
			return func() tea.Msg {
				finish(func() { db.applySwitchMode(db.lifecycleCtx, sm, seq) })
				return nil
			}
		}
		finish := db.beginTrackedOperation()
		return func() tea.Msg {
			finish(func() { db.dispatchIntent(db.lifecycleCtx, msg) })
			return nil
		}
	}
	// cmdFor intentionally emits the human-runnable form (-v, no -json) so
	// the copied command is readable when pasted into a terminal.
	cmdFor := func(test string) string {
		pkgs := strings.Join(prov.TestPackages(), " ")
		pattern := engine.RunPattern([]string{test})
		return fmt.Sprintf("go test %s -run '%s' -v -count=1", pkgs, pattern)
	}

	model := tui.New(prov.Name(), mode, owner, ascii).
		WithExec(exec).
		WithVersion(version).
		WithCmdFunc(cmdFor).
		WithClipboard(tui.OscCopy).
		WithModes(prov.Modes()).
		WithEnvVars(prov.EnvFor(mode)).
		WithGetenv(getenv)
	// Play the one-shot ignition splash on color TTYs only; ascii / NO_COLOR
	// terminals open straight to the dashboard (keeps dumb terminals clean).
	if !ascii {
		model = model.WithIntro()
	}
	p = tea.NewProgram(model, tea.WithOutput(out), tea.WithAltScreen())
	db.send = p.Send

	finishSeed := db.beginTrackedOperation()
	go finishSeed(func() { db.seed(db.lifecycleCtx) })
	_, runErr := p.Run()
	db.cancelLifecycle()
	db.waitForActiveOperations()
	if runErr != nil {
		fmt.Fprintln(errOut, "tui error:", runErr)
		return 1
	}
	return 0
}
