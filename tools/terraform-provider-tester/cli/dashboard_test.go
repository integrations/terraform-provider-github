package cli

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/github/terraform-provider-tester/engine"
	"github.com/github/terraform-provider-tester/fakeprovider"
	"github.com/github/terraform-provider-tester/internal/redact"
	"github.com/github/terraform-provider-tester/provider"
	"github.com/github/terraform-provider-tester/tui"
)

// ── 1. decideTUI ─────────────────────────────────────────────────────────────

func TestDecideTUI(t *testing.T) {
	cases := []struct {
		isTTY, forceTTY, noTUI bool
		want                   bool
	}{
		// noTUI always wins (false regardless)
		{true, true, true, false},
		{true, false, true, false},
		{false, true, true, false},
		{false, false, true, false},
		// noTUI=false: open when TTY or force
		{true, true, false, true},
		{true, false, false, true},
		{false, true, false, true},
		{false, false, false, false},
	}
	for _, tc := range cases {
		got := decideTUI(tc.isTTY, tc.forceTTY, tc.noTUI)
		if got != tc.want {
			t.Errorf("decideTUI(%v,%v,%v) = %v, want %v",
				tc.isTTY, tc.forceTTY, tc.noTUI, got, tc.want)
		}
	}
}

// ── 2. testsForIntent ─────────────────────────────────────────────────────
//
// patternForIntent (a thin ok-bool wrapper over testsForIntent with no
// production callers — runIntent calls testsForIntent directly so it can
// report the exact state/plan-required error) was removed as dead code; its
// coverage lives here, directly against testsForIntent's real (tests, err)
// return shape.

func TestTestsForIntent(t *testing.T) {
	groups := []engine.Group{
		{Name: "repos", Tests: []string{"TestAccRepoA", "TestAccRepoB", "TestAccRepoC"}},
		{Name: "teams", Tests: []string{"TestAccTeamX"}},
	}
	allNames := []string{"TestAccRepoA", "TestAccRepoB", "TestAccRepoC", "TestAccTeamX"}

	t.Run("RetryGroupIntent with prior failures uses failed subset", func(t *testing.T) {
		prev := engine.State{
			Plan: planWithEligible("anonymous", []string{"TestAccRepoA", "TestAccRepoB", "TestAccRepoC"}),
			Results: []engine.PersistResult{
				{Test: "TestAccRepoA", Status: "fail"},
				{Test: "TestAccRepoC", Status: "pass"},
			},
		}
		tests, err := testsForIntent(tui.RetryGroupIntent{Group: "repos"}, groups, prev, allNames)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		want := []string{"TestAccRepoA"}
		if !reflect.DeepEqual(tests, want) {
			t.Errorf("tests = %v, want %v", tests, want)
		}
	})

	t.Run("RetryGroupIntent with no prior failures uses whole group", func(t *testing.T) {
		prev := engine.State{Plan: planWithEligible("anonymous", []string{"TestAccRepoA", "TestAccRepoB", "TestAccRepoC"})}
		tests, err := testsForIntent(tui.RetryGroupIntent{Group: "repos"}, groups, prev, allNames)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		want := []string{"TestAccRepoA", "TestAccRepoB", "TestAccRepoC"}
		if !reflect.DeepEqual(tests, want) {
			t.Errorf("tests = %v, want %v", tests, want)
		}
	})

	t.Run("RetryGroupIntent excludes ineligible group members", func(t *testing.T) {
		// Only TestAccRepoA is eligible in the persisted plan; TestAccRepoB
		// and TestAccRepoC belong to the group but the current mode excludes
		// them, so the whole-group fallback must retry only TestAccRepoA.
		prev := engine.State{Plan: planWithEligible("anonymous", []string{"TestAccRepoA"})}
		tests, err := testsForIntent(tui.RetryGroupIntent{Group: "repos"}, groups, prev, allNames)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		want := []string{"TestAccRepoA"}
		if !reflect.DeepEqual(tests, want) {
			t.Errorf("tests = %v, want %v", tests, want)
		}
	})

	t.Run("RetryGroupIntent unknown group returns no tests, no error", func(t *testing.T) {
		prev := engine.State{Plan: planWithEligible("anonymous", allNames)}
		tests, err := testsForIntent(tui.RetryGroupIntent{Group: "unknown"}, groups, prev, allNames)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(tests) != 0 {
			t.Errorf("tests = %v, want none for an unknown group", tests)
		}
	})

	t.Run("RetryGroupIntent with nil persisted plan returns state/plan-required error", func(t *testing.T) {
		prev := engine.State{}
		tests, err := testsForIntent(tui.RetryGroupIntent{Group: "repos"}, groups, prev, allNames)
		if err == nil || err.Error() != planRequiredError {
			t.Errorf("err = %v, want %q", err, planRequiredError)
		}
		if len(tests) != 0 {
			t.Errorf("tests = %v, want none alongside the error", tests)
		}
	})

	t.Run("RetryTestIntent produces single test when it is currently eligible", func(t *testing.T) {
		// RetryTestIntent has no CLI equivalent and is bound to whatever the
		// dashboard is showing right now (allNames, the current-mode eligible
		// set), independent of any persisted plan — a persisted plan is not
		// required to run a single displayed test.
		prev := engine.State{}
		tests, err := testsForIntent(tui.RetryTestIntent{Test: "TestAccRepoA"}, groups, prev, allNames)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		want := []string{"TestAccRepoA"}
		if !reflect.DeepEqual(tests, want) {
			t.Errorf("tests = %v, want %v", tests, want)
		}
	})

	t.Run("RetryTestIntent for a test outside the current eligible set returns no tests, no error", func(t *testing.T) {
		prev := engine.State{Plan: planWithEligible("anonymous", allNames)}
		currentEligible := []string{"TestAccRepoB", "TestAccRepoC", "TestAccTeamX"} // excludes TestAccRepoA
		tests, err := testsForIntent(tui.RetryTestIntent{Test: "TestAccRepoA"}, groups, prev, currentEligible)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(tests) != 0 {
			t.Errorf("tests = %v, want none for a test the current mode excludes", tests)
		}
	})

	t.Run("RetryAllIntent uses all failed top-level tests", func(t *testing.T) {
		prev := engine.State{
			Plan: planWithEligible("anonymous", allNames),
			Results: []engine.PersistResult{
				{Test: "TestAccRepoA", Status: "fail"},
				{Test: "TestAccRepoB", Status: "pass"},
				{Test: "TestAccTeamX", Status: "fail"},
			},
		}
		tests, err := testsForIntent(tui.RetryAllIntent{}, groups, prev, allNames)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		want := []string{"TestAccRepoA", "TestAccTeamX"}
		if !reflect.DeepEqual(tests, want) {
			t.Errorf("tests = %v, want %v", tests, want)
		}
	})

	t.Run("RetryAllIntent with no failures returns no tests, no error", func(t *testing.T) {
		prev := engine.State{Plan: planWithEligible("anonymous", allNames)}
		tests, err := testsForIntent(tui.RetryAllIntent{}, groups, prev, allNames)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(tests) != 0 {
			t.Errorf("tests = %v, want none when no failures", tests)
		}
	})

	t.Run("RetryAllIntent with nil persisted plan returns state/plan-required error", func(t *testing.T) {
		prev := engine.State{
			Results: []engine.PersistResult{{Test: "TestAccRepoA", Status: "fail"}},
		}
		tests, err := testsForIntent(tui.RetryAllIntent{}, groups, prev, allNames)
		if err == nil || err.Error() != planRequiredError {
			t.Errorf("err = %v, want %q", err, planRequiredError)
		}
		if len(tests) != 0 {
			t.Errorf("tests = %v, want none alongside the error", tests)
		}
	})

	t.Run("ResumeIntent combines failed and not-run", func(t *testing.T) {
		prev := engine.State{
			Plan: planWithEligible("anonymous", allNames),
			Results: []engine.PersistResult{
				{Test: "TestAccRepoA", Status: "pass"},
				{Test: "TestAccRepoB", Status: "fail"},
			},
		}
		// TestAccRepoC and TestAccTeamX are not-run
		tests, err := testsForIntent(tui.ResumeIntent{}, groups, prev, allNames)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		// failed first, then not-run (deduped, deterministic order)
		want := []string{"TestAccRepoB", "TestAccRepoC", "TestAccTeamX"}
		if !reflect.DeepEqual(tests, want) {
			t.Errorf("tests = %v, want %v", tests, want)
		}
	})

	t.Run("ResumeIntent on empty state and allNames returns no tests, no error", func(t *testing.T) {
		// A plan with an empty eligible set + empty state → both failed and
		// notRun are empty → nothing to run (distinct from a nil Plan, which
		// is the separate state/plan-required case exercised below).
		prev := engine.State{Plan: planWithEligible("anonymous", nil)}
		tests, err := testsForIntent(tui.ResumeIntent{}, groups, prev, []string{})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(tests) != 0 {
			t.Errorf("tests = %v, want none for empty state and empty allNames", tests)
		}
	})

	t.Run("ResumeIntent with nil persisted plan returns state/plan-required error", func(t *testing.T) {
		tests, err := testsForIntent(tui.ResumeIntent{}, groups, engine.State{}, allNames)
		if err == nil || err.Error() != planRequiredError {
			t.Errorf("err = %v, want %q", err, planRequiredError)
		}
		if len(tests) != 0 {
			t.Errorf("tests = %v, want none alongside the error", tests)
		}
	})

	t.Run("unknown msg type returns no tests, no error", func(t *testing.T) {
		prev := engine.State{Plan: planWithEligible("anonymous", allNames)}
		tests, err := testsForIntent(struct{}{}, groups, prev, allNames)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(tests) != 0 {
			t.Errorf("tests = %v, want none for an unknown message type", tests)
		}
	})
}

// ── 3. resumePromptInfo ──────────────────────────────────────────────────────

func TestResumePromptInfo(t *testing.T) {
	allNames := []string{"TestAccA", "TestAccB", "TestAccC"}

	t.Run("prior run with failures shows prompt", func(t *testing.T) {
		prev := engine.State{
			Results: []engine.PersistResult{
				{Test: "TestAccA", Status: "pass"},
				{Test: "TestAccB", Status: "fail"},
			},
		}
		failed, notRun, show := resumePromptInfo(prev, allNames)
		if !show {
			t.Error("expected show=true")
		}
		if failed != 1 {
			t.Errorf("failed = %d, want 1", failed)
		}
		if notRun != 1 { // TestAccC not run
			t.Errorf("notRun = %d, want 1", notRun)
		}
	})

	t.Run("empty state (no results) does not show prompt", func(t *testing.T) {
		prev := engine.State{}
		_, _, show := resumePromptInfo(prev, allNames)
		if show {
			t.Error("expected show=false for empty state")
		}
	})

	t.Run("all passed no not-run does not show prompt", func(t *testing.T) {
		prev := engine.State{
			Results: []engine.PersistResult{
				{Test: "TestAccA", Status: "pass"},
				{Test: "TestAccB", Status: "pass"},
				{Test: "TestAccC", Status: "pass"},
			},
		}
		_, _, show := resumePromptInfo(prev, allNames)
		if show {
			t.Error("expected show=false when all passed and no not-run")
		}
	})
}

// ── 4. redactResult (SECURITY) ───────────────────────────────────────────────

func TestRedactResult(t *testing.T) {
	const secret = "ghp_FAKEFAKEFAKE"
	red := redact.New([]string{secret})

	original := engine.TestResult{
		Name:   "TestAccFoo",
		Status: provider.StatusFail,
		Output: []string{"running test", "token=" + secret, "error occurred"},
	}

	got := redactResult(red, original)

	// The returned result must have the secret replaced.
	for _, line := range got.Output {
		if strings.Contains(line, secret) {
			t.Errorf("raw secret found in redacted output line: %q", line)
		}
	}
	// The mask must appear.
	found := false
	for _, line := range got.Output {
		if strings.Contains(line, "***REDACTED***") {
			found = true
		}
	}
	if !found {
		t.Error("redaction mask not found in returned Output")
	}

	// The ORIGINAL must be unchanged (no mutation).
	if original.Output[1] != "token="+secret {
		t.Errorf("original Output was mutated: %q", original.Output[1])
	}

	// Non-secret fields must be preserved.
	if got.Name != original.Name {
		t.Errorf("Name changed: %q vs %q", got.Name, original.Name)
	}
	if got.Status != original.Status {
		t.Errorf("Status changed")
	}
}

// ── 5. runIntent streams REDACTED results (SECURITY) ─────────────────────────

func TestRunIntentRedactsOutput(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")

	// Seed state so RetryAllIntent has something to retry.
	st := engine.State{Provider: "test", Mode: "anonymous"}
	st.Plan = planWithEligible("anonymous", []string{"TestAccFoo"})
	st.Results = []engine.PersistResult{
		{Test: "TestAccFoo", Status: "fail"},
	}
	if err := st.Save(statePath); err != nil {
		t.Fatalf("seeding state: %v", err)
	}

	const secret = "ghp_FAKEFAKEFAKE"
	fr := &fakeRunner{
		result: engine.RunResult{
			Tests: []engine.TestResult{
				{
					Name:   "TestAccFoo",
					Status: provider.StatusFail,
					Output: []string{"token=" + secret},
				},
			},
			BuildOutput:  []string{"build token=" + secret},
			PreRunOutput: []string{"prerun token=" + secret},
			RawLog:       []string{"raw token=" + secret},
		},
	}

	var mu sync.Mutex
	var msgs []tea.Msg
	send := func(m tea.Msg) {
		mu.Lock()
		msgs = append(msgs, m)
		mu.Unlock()
	}

	pf := &fakeprovider.Fake{
		NameVal:  "test",
		Packages: []string{"./..."},
	}

	db := &dashboard{
		prov:      pf,
		root:      root,
		mode:      "anonymous",
		statePath: statePath,
		newRunner: func(_ io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccFoo"}),
		red:       redact.New([]string{secret}),
		groups:    []engine.Group{{Name: "tests", Tests: []string{"TestAccFoo"}}},
		allNames:  []string{"TestAccFoo"},
		send:      send,
		logSink:   io.Discard,
	}

	db.runIntent(context.Background(), tui.RetryAllIntent{})

	mu.Lock()
	captured := make([]tea.Msg, len(msgs))
	copy(captured, msgs)
	mu.Unlock()

	var gotStart, gotDone bool
	var doneMsg tui.RunDoneMsg
	var updateMsgs []tui.TestUpdateMsg
	for _, m := range captured {
		switch v := m.(type) {
		case tui.RunStartedMsg:
			gotStart = true
		case tui.TestUpdateMsg:
			updateMsgs = append(updateMsgs, v)
		case tui.RunDoneMsg:
			gotDone = true
			doneMsg = v
		}
	}

	if !gotStart {
		t.Error("expected RunStartedMsg")
	}
	if len(updateMsgs) == 0 {
		t.Error("expected at least one TestUpdateMsg")
	}
	if !gotDone {
		t.Error("expected RunDoneMsg")
	}

	// SECURITY: raw secret must NOT appear in any TestUpdateMsg.Output.
	for _, um := range updateMsgs {
		for _, line := range um.Result.Output {
			if strings.Contains(line, secret) {
				t.Errorf("raw secret found in TestUpdateMsg.Output: %q", line)
			}
		}
	}

	// SECURITY: redaction mask must appear in the output.
	maskFound := false
	for _, um := range updateMsgs {
		for _, line := range um.Result.Output {
			if strings.Contains(line, "***REDACTED***") {
				maskFound = true
			}
		}
	}
	if !maskFound {
		t.Error("redaction mask not found in any TestUpdateMsg.Output")
	}

	// SECURITY: the on-disk state file must not contain the secret.
	data, err := os.ReadFile(statePath)
	if err != nil {
		t.Fatalf("reading state file: %v", err)
	}
	if strings.Contains(string(data), secret) {
		t.Error("raw secret found in on-disk state file")
	}

	// SECURITY: original TestResult.Output must be unchanged (no mutation).
	origOutput := fr.result.Tests[0].Output
	if len(origOutput) == 0 || origOutput[0] != "token="+secret {
		t.Errorf("original TestResult.Output was mutated: %v", origOutput)
	}

	// SECURITY: RunDoneMsg carries the full RunResult; the raw secret must not
	// appear in ANY of its output-bearing fields (Tests[].Output, BuildOutput,
	// PreRunOutput, RawLog), and a redaction mask must be present.
	if !gotDone {
		t.Fatal("expected RunDoneMsg")
	}
	var doneLines []string
	for _, tr := range doneMsg.Result.Tests {
		doneLines = append(doneLines, tr.Output...)
	}
	doneLines = append(doneLines, doneMsg.Result.BuildOutput...)
	doneLines = append(doneLines, doneMsg.Result.PreRunOutput...)
	doneLines = append(doneLines, doneMsg.Result.RawLog...)
	maskInDone := false
	for _, line := range doneLines {
		if strings.Contains(line, secret) {
			t.Errorf("raw secret found in RunDoneMsg.Result: %q", line)
		}
		if strings.Contains(line, "***REDACTED***") {
			maskInDone = true
		}
	}
	if !maskInDone {
		t.Error("redaction mask not found in RunDoneMsg.Result")
	}

	// SECURITY: redacting the done payload must not mutate the runner's result.
	if got := fr.result.RawLog; len(got) == 0 || got[0] != "raw token="+secret {
		t.Errorf("original RunResult.RawLog was mutated: %v", got)
	}
	if got := fr.result.BuildOutput; len(got) == 0 || got[0] != "build token="+secret {
		t.Errorf("original RunResult.BuildOutput was mutated: %v", got)
	}
	if got := fr.result.PreRunOutput; len(got) == 0 || got[0] != "prerun token="+secret {
		t.Errorf("original RunResult.PreRunOutput was mutated: %v", got)
	}
}

// TestRunIntentLocksBeforeReadingState verifies runIntent acquires the state
// lock BEFORE deriving its run decision from prior state. With the lock already
// held, it must fail at lock acquisition (not proceed to a state-derived
// "nothing to run" verdict), and must not invoke the runner. This pins the
// lock-before-load ordering that prevents a stale-merge TOCTOU race.
func TestRunIntentLocksBeforeReadingState(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")

	// Seed state with a single PASSING result: prev exists but has no failures,
	// so a RetryAll would resolve to "nothing to run" if state were read first.
	st := engine.State{Provider: "test", Mode: "anonymous"}
	st.Results = []engine.PersistResult{{Test: "TestAccFoo", Status: "pass"}}
	if err := st.Save(statePath); err != nil {
		t.Fatalf("seeding state: %v", err)
	}

	// Pre-hold the lock to simulate a concurrent harness run.
	release, err := engine.AcquireLock(statePath + ".lock")
	if err != nil {
		t.Fatalf("pre-acquiring lock: %v", err)
	}
	defer release() //nolint:errcheck

	fr := &fakeRunner{}
	var mu sync.Mutex
	var errs []error
	send := func(m tea.Msg) {
		if em, ok := m.(tui.ErrMsg); ok {
			mu.Lock()
			errs = append(errs, em.Err)
			mu.Unlock()
		}
	}

	db := &dashboard{
		prov:      &fakeprovider.Fake{NameVal: "test", Packages: []string{"./..."}},
		root:      root,
		mode:      "anonymous",
		statePath: statePath,
		newRunner: func(_ io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccFoo"}),
		red:       redact.New(nil),
		groups:    []engine.Group{{Name: "tests", Tests: []string{"TestAccFoo"}}},
		allNames:  []string{"TestAccFoo"},
		send:      send,
		logSink:   io.Discard,
	}

	db.runIntent(context.Background(), tui.RetryAllIntent{})

	if fr.calls != 0 {
		t.Errorf("runner was invoked %d times while the lock was held; expected 0", fr.calls)
	}
	mu.Lock()
	defer mu.Unlock()
	if len(errs) != 1 {
		t.Fatalf("expected exactly one ErrMsg, got %d: %v", len(errs), errs)
	}
	got := errs[0].Error()
	if !strings.Contains(got, "lock") {
		t.Errorf("expected a lock-acquisition error (lock acquired before reading state); got %q", got)
	}
	if strings.Contains(got, "nothing to run") {
		t.Errorf("state was read before acquiring the lock (TOCTOU): got %q", got)
	}
}

// ── 6. single-run guard ──────────────────────────────────────────────────────

func TestRunIntentSingleRunGuard(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")

	fr := &fakeRunner{}
	pf := &fakeprovider.Fake{NameVal: "test", Packages: []string{"./..."}}

	var msgs []tea.Msg
	send := func(m tea.Msg) { msgs = append(msgs, m) }

	db := &dashboard{
		prov:      pf,
		root:      root,
		mode:      "anonymous",
		statePath: statePath,
		newRunner: func(_ io.Writer) testRunner { return fr },
		list:      stubList(nil),
		red:       redact.New(nil),
		groups:    []engine.Group{},
		allNames:  []string{},
		send:      send,
		logSink:   io.Discard,
	}

	db.busy.Store(true) // simulate a run already in progress

	db.runIntent(context.Background(), tui.RetryAllIntent{})

	if len(msgs) != 1 {
		t.Fatalf("expected exactly 1 message, got %d: %v", len(msgs), msgs)
	}
	errMsg, ok := msgs[0].(tui.ErrMsg)
	if !ok {
		t.Fatalf("expected ErrMsg, got %T", msgs[0])
	}
	if !strings.Contains(errMsg.Err.Error(), "already in progress") {
		t.Errorf("expected 'already in progress', got %q", errMsg.Err.Error())
	}
	if fr.calls != 0 {
		t.Errorf("runner should not be invoked when busy, got %d calls", fr.calls)
	}
}

// ── 7. nothing-to-run ────────────────────────────────────────────────────────

func TestRunIntentNothingToRun(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")

	// Seed a state with a persisted (empty-eligible) plan so this exercises
	// the "eligible but nothing to run" path rather than the separate
	// nil-plan/state-plan-required path (covered by
	// TestRunIntentPlanRequiredWhenPersistedPlanNil).
	st := engine.State{Provider: "test", Mode: "anonymous"}
	st.Plan = planWithEligible("anonymous", nil)
	if err := st.Save(statePath); err != nil {
		t.Fatalf("seeding state: %v", err)
	}

	fr := &fakeRunner{}
	pf := &fakeprovider.Fake{NameVal: "test", Packages: []string{"./..."}}

	var msgs []tea.Msg
	send := func(m tea.Msg) { msgs = append(msgs, m) }

	db := &dashboard{
		prov:      pf,
		root:      root,
		mode:      "anonymous",
		statePath: statePath,
		newRunner: func(_ io.Writer) testRunner { return fr },
		list:      stubList(nil),
		red:       redact.New(nil),
		groups:    []engine.Group{},
		allNames:  []string{},
		send:      send,
		logSink:   io.Discard,
	}

	// ResumeIntent on empty state (no results, no allNames) → nothing to run
	db.runIntent(context.Background(), tui.ResumeIntent{})

	for _, m := range msgs {
		if _, ok := m.(tui.RunStartedMsg); ok {
			t.Error("RunStartedMsg must not be sent when nothing to run")
		}
	}

	var gotErr bool
	for _, m := range msgs {
		if errMsg, ok := m.(tui.ErrMsg); ok {
			gotErr = true
			if !strings.Contains(errMsg.Err.Error(), "nothing to run") {
				t.Errorf("expected 'nothing to run', got %q", errMsg.Err.Error())
			}
		}
	}
	if !gotErr {
		t.Error("expected ErrMsg('nothing to run')")
	}
	if fr.calls != 0 {
		t.Errorf("runner must not be invoked, got %d calls", fr.calls)
	}
}

// ── 8. seed ──────────────────────────────────────────────────────────────────

func TestSeed(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")

	runAt := time.Now().Add(-1 * time.Hour)
	st := engine.State{
		Provider: "test",
		Mode:     "anonymous",
		RunAt:    runAt,
		Results: []engine.PersistResult{
			{Test: "TestAccA", Status: "fail"},
			{Test: "TestAccB", Status: "pass"},
		},
	}
	if err := st.Save(statePath); err != nil {
		t.Fatalf("seeding state: %v", err)
	}

	pf := &fakeprovider.Fake{
		NameVal:  "test",
		Packages: []string{"./..."},
		PreflightFn: func(_ context.Context, _ string, _ provider.TestRequirements) provider.PreflightReport {
			return provider.PreflightReport{Mode: "anonymous"}
		},
	}

	allNames := []string{"TestAccA", "TestAccB"}
	groups := []engine.Group{{Name: "tests", Tests: allNames}}

	var msgs []tea.Msg
	send := func(m tea.Msg) { msgs = append(msgs, m) }

	db := &dashboard{
		prov:      pf,
		root:      root,
		mode:      "anonymous",
		statePath: statePath,
		newRunner: func(_ io.Writer) testRunner { return &fakeRunner{} },
		list:      stubList(allNames),
		red:       redact.New(nil),
		plan:      planWithEligible("anonymous", allNames),
		groups:    groups,
		allNames:  allNames,
		send:      send,
		logSink:   io.Discard,
	}

	db.seed(context.Background())

	var gotPreflight, gotGroups, gotResume bool
	var updateCount int
	for _, m := range msgs {
		switch m.(type) {
		case tui.PreflightMsg:
			gotPreflight = true
		case tui.GroupsMsg:
			gotGroups = true
		case tui.TestUpdateMsg:
			updateCount++
		case tui.ResumePromptMsg:
			gotResume = true
		}
	}

	if !gotPreflight {
		t.Error("expected PreflightMsg")
	}
	if !gotGroups {
		t.Error("expected GroupsMsg")
	}
	if updateCount != 2 {
		t.Errorf("expected 2 TestUpdateMsg (one per persisted result), got %d", updateCount)
	}
	if !gotResume {
		t.Error("expected ResumePromptMsg (one failed test in prior state)")
	}

	// Verify TestUpdateMsgs have empty Output (status-only).
	for _, m := range msgs {
		if um, ok := m.(tui.TestUpdateMsg); ok {
			if len(um.Result.Output) != 0 {
				t.Errorf("seed TestUpdateMsg must have empty Output, got %v", um.Result.Output)
			}
		}
	}
}

// ── 9. runIntent writes failure logs ──────────────────────────────────────────

func TestRunIntentWritesFailureLogs(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")

	const secret = "ghp_TESTSECRET"
	st := engine.State{Provider: "test", Mode: "anonymous"}
	st.Plan = planWithEligible("anonymous", []string{"TestAccBar"})
	st.Results = []engine.PersistResult{{Test: "TestAccBar", Status: "fail"}}
	if err := st.Save(statePath); err != nil {
		t.Fatalf("seeding state: %v", err)
	}

	const pkg = "github.com/integrations/terraform-provider-github/github"
	fr := &fakeRunner{
		result: engine.RunResult{
			Tests: []engine.TestResult{
				{
					Package: pkg,
					Name:    "TestAccBar",
					Status:  provider.StatusFail,
					Output:  []string{"token=" + secret + "\n"},
				},
			},
		},
	}

	var mu sync.Mutex
	var msgs []tea.Msg
	send := func(m tea.Msg) {
		mu.Lock()
		msgs = append(msgs, m)
		mu.Unlock()
	}

	pf := &fakeprovider.Fake{NameVal: "test", Packages: []string{pkg}}
	db := &dashboard{
		prov:      pf,
		root:      root,
		mode:      "anonymous",
		statePath: statePath,
		newRunner: func(_ io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccBar"}),
		red:       redact.New([]string{secret}),
		groups:    []engine.Group{{Name: "tests", Tests: []string{"TestAccBar"}}},
		allNames:  []string{"TestAccBar"},
		send:      send,
		logSink:   io.Discard,
	}

	db.runIntent(context.Background(), tui.RetryAllIntent{})

	failDir := filepath.Join(root, ".pulsar-failures")
	wantLog := engine.FailureLogPath(failDir, pkg, "TestAccBar")

	data, err := os.ReadFile(wantLog)
	if err != nil {
		t.Fatalf("failure log not written: %v", err)
	}
	content := string(data)
	if strings.Contains(content, secret) {
		t.Errorf("failure log contains unredacted secret: %q", content)
	}
	if !strings.Contains(content, "***REDACTED***") {
		t.Errorf("failure log missing redaction marker: %q", content)
	}
}

func TestRunIntentPersistsBuildFailureStateAndLog(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")
	failDir := filepath.Join(root, ".pulsar-failures")

	const secret = "ghp_TESTSECRET"
	const pkg = "github.com/integrations/terraform-provider-github/github"
	st := engine.State{Provider: "test", Mode: "anonymous"}
	st.Plan = planWithEligible("anonymous", []string{"TestAccBar"})
	st.Results = []engine.PersistResult{{Package: pkg, Test: "TestAccBar", Status: "fail"}}
	if err := st.Save(statePath); err != nil {
		t.Fatalf("seeding state: %v", err)
	}

	db := &dashboard{
		prov:      &fakeprovider.Fake{NameVal: "test", Packages: []string{pkg}},
		root:      root,
		mode:      "anonymous",
		statePath: statePath,
		newRunner: func(_ io.Writer) testRunner {
			return &fakeRunner{result: engine.RunResult{
				BuildFailed: true,
				BuildOutput: []string{"compile failed token=" + secret + "\n"},
				Tests: []engine.TestResult{{
					Package: pkg,
					Name:    "TestAccBar",
					Status:  provider.StatusPass,
				}},
			}}
		},
		list:     stubList([]string{"TestAccBar"}),
		red:      redact.New([]string{secret}),
		groups:   []engine.Group{{Name: "tests", Tests: []string{"TestAccBar"}}},
		allNames: []string{"TestAccBar"},
		send:     func(tea.Msg) {},
		logSink:  io.Discard,
	}

	db.runIntent(context.Background(), tui.RetryAllIntent{})

	state, err := engine.Load(statePath)
	if err != nil {
		t.Fatalf("loading state: %v", err)
	}
	wantLog := filepath.Join(failDir, "build.log")
	if !state.BuildFailed || state.BuildLog != wantLog {
		t.Fatalf("suite build failure state = %+v, want BuildFailed with log %q", state, wantLog)
	}
	if state.PreRunFailed || state.PreRunLog != "" {
		t.Fatalf("pre-run state = failed:%v log:%q, want cleared", state.PreRunFailed, state.PreRunLog)
	}
	data, err := os.ReadFile(state.BuildLog)
	if err != nil {
		t.Fatalf("reading build failure log: %v", err)
	}
	content := string(data)
	if strings.Contains(content, secret) {
		t.Fatalf("build failure log contains unredacted secret: %q", content)
	}
	if !strings.Contains(content, "***REDACTED***") {
		t.Fatalf("build failure log missing redaction marker: %q", content)
	}
}

func TestRunIntentPersistsPreRunFailureStateAndLog(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")
	failDir := filepath.Join(root, ".pulsar-failures")

	const secret = "ghp_TESTSECRET"
	const pkg = "github.com/integrations/terraform-provider-github/github"
	st := engine.State{Provider: "test", Mode: "anonymous"}
	st.Plan = planWithEligible("anonymous", []string{"TestAccBar"})
	st.Results = []engine.PersistResult{{Package: pkg, Test: "TestAccBar", Status: "fail"}}
	if err := st.Save(statePath); err != nil {
		t.Fatalf("seeding state: %v", err)
	}

	db := &dashboard{
		prov:      &fakeprovider.Fake{NameVal: "test", Packages: []string{pkg}},
		root:      root,
		mode:      "anonymous",
		statePath: statePath,
		newRunner: func(_ io.Writer) testRunner {
			return &fakeRunner{result: engine.RunResult{
				PreRunFailed: true,
				PreRunOutput: []string{"provider setup failed token=" + secret + "\n"},
				Tests: []engine.TestResult{{
					Package: pkg,
					Name:    "TestAccBar",
					Status:  provider.StatusPass,
				}},
			}}
		},
		list:     stubList([]string{"TestAccBar"}),
		red:      redact.New([]string{secret}),
		groups:   []engine.Group{{Name: "tests", Tests: []string{"TestAccBar"}}},
		allNames: []string{"TestAccBar"},
		send:     func(tea.Msg) {},
		logSink:  io.Discard,
	}

	db.runIntent(context.Background(), tui.RetryAllIntent{})

	state, err := engine.Load(statePath)
	if err != nil {
		t.Fatalf("loading state: %v", err)
	}
	wantLog := filepath.Join(failDir, "pre-run.log")
	if !state.PreRunFailed || state.PreRunLog != wantLog {
		t.Fatalf("suite pre-run failure state = %+v, want PreRunFailed with log %q", state, wantLog)
	}
	if state.BuildFailed || state.BuildLog != "" {
		t.Fatalf("build state = failed:%v log:%q, want cleared", state.BuildFailed, state.BuildLog)
	}
	data, err := os.ReadFile(state.PreRunLog)
	if err != nil {
		t.Fatalf("reading pre-run failure log: %v", err)
	}
	content := string(data)
	if strings.Contains(content, secret) {
		t.Fatalf("pre-run failure log contains unredacted secret: %q", content)
	}
	if !strings.Contains(content, "***REDACTED***") {
		t.Fatalf("pre-run failure log missing redaction marker: %q", content)
	}
}

func TestRunIntentSuccessfulRunClearsStaleSuiteFailureStateAndLogs(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")
	failDir := filepath.Join(root, ".pulsar-failures")
	buildLog := filepath.Join(failDir, "build.log")
	preRunLog := filepath.Join(failDir, "pre-run.log")

	const pkg = "github.com/integrations/terraform-provider-github/github"
	if err := os.MkdirAll(failDir, 0o700); err != nil {
		t.Fatalf("mkdir failDir: %v", err)
	}
	for path, content := range map[string]string{
		buildLog:  "stale build failure\n",
		preRunLog: "stale pre-run failure\n",
	} {
		if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
			t.Fatalf("writing %q: %v", path, err)
		}
	}

	st := engine.State{
		Provider:     "test",
		Mode:         "anonymous",
		BuildFailed:  true,
		BuildLog:     buildLog,
		PreRunFailed: true,
		PreRunLog:    preRunLog,
		Results:      []engine.PersistResult{{Package: pkg, Test: "TestAccBar", Status: "fail"}},
	}
	st.Plan = planWithEligible("anonymous", []string{"TestAccBar"})
	if err := st.Save(statePath); err != nil {
		t.Fatalf("seeding state: %v", err)
	}

	db := &dashboard{
		prov:      &fakeprovider.Fake{NameVal: "test", Packages: []string{pkg}},
		root:      root,
		mode:      "anonymous",
		statePath: statePath,
		newRunner: func(_ io.Writer) testRunner {
			return &fakeRunner{result: engine.RunResult{Tests: []engine.TestResult{{
				Package: pkg,
				Name:    "TestAccBar",
				Status:  provider.StatusPass,
			}}}}
		},
		list:     stubList([]string{"TestAccBar"}),
		red:      redact.New(nil),
		groups:   []engine.Group{{Name: "tests", Tests: []string{"TestAccBar"}}},
		allNames: []string{"TestAccBar"},
		send:     func(tea.Msg) {},
		logSink:  io.Discard,
	}

	db.runIntent(context.Background(), tui.RetryAllIntent{})

	state, err := engine.Load(statePath)
	if err != nil {
		t.Fatalf("loading state: %v", err)
	}
	if state.BuildFailed || state.BuildLog != "" || state.PreRunFailed || state.PreRunLog != "" {
		t.Fatalf("suite failure state = %+v, want cleared flags and log paths", state)
	}
	for _, path := range []string{buildLog, preRunLog} {
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			t.Fatalf("stale suite failure log %q still present; stat err=%v", path, err)
		}
	}
}

// ── 10. seed loads failure log for failed test ────────────────────────────────

func TestSeedLoadsFailureLogForFailedTest(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")

	const pkg = "github.com/integrations/terraform-provider-github/github"
	const logContent = "assertion error line 1\nassertion error line 2\n"

	// Pre-write a failure log file.
	failDir := filepath.Join(root, ".pulsar-failures")
	if err := os.MkdirAll(failDir, 0o700); err != nil {
		t.Fatalf("mkdir failDir: %v", err)
	}
	logPath := engine.FailureLogPath(failDir, pkg, "TestAccFail")
	if err := os.WriteFile(logPath, []byte(logContent), 0o600); err != nil {
		t.Fatalf("writing log: %v", err)
	}

	// Seed state with one failed and one passing test.
	st := engine.State{
		Provider: "test",
		Mode:     "anonymous",
		Results: []engine.PersistResult{
			{Test: "TestAccFail", Package: pkg, Status: "fail"},
			{Test: "TestAccPass", Package: pkg, Status: "pass"},
		},
	}
	if err := st.Save(statePath); err != nil {
		t.Fatalf("seeding state: %v", err)
	}

	pf := &fakeprovider.Fake{
		NameVal:  "test",
		Packages: []string{pkg},
		PreflightFn: func(_ context.Context, _ string, _ provider.TestRequirements) provider.PreflightReport {
			return provider.PreflightReport{Mode: "anonymous"}
		},
	}

	allNames := []string{"TestAccFail", "TestAccPass"}
	var msgs []tea.Msg
	send := func(m tea.Msg) { msgs = append(msgs, m) }

	db := &dashboard{
		prov:      pf,
		root:      root,
		mode:      "anonymous",
		statePath: statePath,
		newRunner: func(_ io.Writer) testRunner { return &fakeRunner{} },
		list:      stubList(allNames),
		red:       redact.New(nil),
		plan:      planWithEligible("anonymous", allNames),
		groups:    []engine.Group{{Name: "tests", Tests: allNames}},
		allNames:  allNames,
		send:      send,
		logSink:   io.Discard,
	}

	db.seed(context.Background())

	var failMsg, passMsg *tui.TestUpdateMsg
	for i := range msgs {
		if um, ok := msgs[i].(tui.TestUpdateMsg); ok {
			um := um
			if um.Result.Name == "TestAccFail" {
				failMsg = &um
			} else if um.Result.Name == "TestAccPass" {
				passMsg = &um
			}
		}
	}

	if failMsg == nil {
		t.Fatal("expected TestUpdateMsg for TestAccFail")
	}
	if passMsg == nil {
		t.Fatal("expected TestUpdateMsg for TestAccPass")
	}

	// Failed test must have output loaded from the log file.
	joinedOutput := strings.Join(failMsg.Result.Output, "\n")
	if !strings.Contains(joinedOutput, "assertion error line 1") {
		t.Errorf("seed failed test output = %q; want to contain log content", joinedOutput)
	}

	// Passing test must have empty output (status-only).
	if len(passMsg.Result.Output) != 0 {
		t.Errorf("seed passing test must have empty Output, got %v", passMsg.Result.Output)
	}
}

// ── 11. bare-invocation seam ─────────────────────────────────────────────────

func TestBareInvocationNonTTYPrintsHelp(t *testing.T) {
	pf := &fakeprovider.Fake{}
	d := deps{
		provider:  pf,
		newRunner: func(_ io.Writer) testRunner { return &fakeRunner{} },
		list:      stubList(nil),
		cwd:       func() (string, error) { return t.TempDir(), nil },
		isTTY:     func() bool { return false },
	}

	done := make(chan int, 1)
	go func() {
		var out, errOut bytes.Buffer
		done <- runWithDeps([]string{}, &out, &errOut, d)
	}()

	select {
	case code := <-done:
		if code != 0 {
			t.Errorf("expected exit 0, got %d", code)
		}
	case <-time.After(3 * time.Second):
		t.Fatal("runWithDeps blocked; expected immediate return for non-TTY")
	}
}

// ── 12. seed: unreadable failure log sends ErrMsg (Finding 7) ────────────────

func TestSeedFailureLogUnreadableSendsErrMsg(t *testing.T) {
	if os.Getuid() == 0 {
		t.Skip("root can read files with mode 0o000; permission test not meaningful")
	}
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")

	const pkg = "github.com/integrations/terraform-provider-github/github"

	failDir := filepath.Join(root, ".pulsar-failures")
	if err := os.MkdirAll(failDir, 0o700); err != nil {
		t.Fatalf("mkdir failDir: %v", err)
	}
	logPath := engine.FailureLogPath(failDir, pkg, "TestAccFail")
	if err := os.WriteFile(logPath, []byte("some content"), 0o600); err != nil {
		t.Fatalf("writing log: %v", err)
	}
	if err := os.Chmod(logPath, 0o000); err != nil {
		t.Fatalf("chmod log: %v", err)
	}

	st := engine.State{
		Provider: "test",
		Mode:     "anonymous",
		Results: []engine.PersistResult{
			{Test: "TestAccFail", Package: pkg, Status: "fail"},
		},
	}
	if err := st.Save(statePath); err != nil {
		t.Fatalf("seeding state: %v", err)
	}

	pf := &fakeprovider.Fake{
		NameVal:  "test",
		Packages: []string{pkg},
		PreflightFn: func(_ context.Context, _ string, _ provider.TestRequirements) provider.PreflightReport {
			return provider.PreflightReport{Mode: "anonymous"}
		},
	}

	allNames := []string{"TestAccFail"}
	var msgs []tea.Msg
	send := func(m tea.Msg) { msgs = append(msgs, m) }

	db := &dashboard{
		prov:      pf,
		root:      root,
		mode:      "anonymous",
		statePath: statePath,
		newRunner: func(_ io.Writer) testRunner { return &fakeRunner{} },
		list:      stubList(allNames),
		red:       redact.New(nil),
		plan:      planWithEligible("anonymous", allNames),
		groups:    []engine.Group{{Name: "tests", Tests: allNames}},
		allNames:  allNames,
		send:      send,
		logSink:   io.Discard,
	}

	db.seed(context.Background())

	var gotErr bool
	for _, m := range msgs {
		if em, ok := m.(tui.ErrMsg); ok {
			gotErr = true
			if em.Err == nil || !strings.Contains(em.Err.Error(), "reading failure log") {
				t.Errorf("expected 'reading failure log' in ErrMsg, got %v", em.Err)
			}
		}
	}
	if !gotErr {
		t.Error("expected ErrMsg for unreadable failure log")
	}
}

// TestLabelForIntent verifies the status-line label derived for each run intent
// (shown next to the run spinner during a run).
func TestLabelForIntent(t *testing.T) {
	cases := []struct {
		name  string
		msg   tea.Msg
		count int
		want  string
	}{
		{"group", tui.RetryGroupIntent{Group: "repos"}, 3, "repos"},
		{"test", tui.RetryTestIntent{Test: "TestAccGithubRepository"}, 1, "TestAccGithubRepository"},
		{"all", tui.RetryAllIntent{}, 5, "all failures"},
		{"resume", tui.ResumeIntent{}, 7, "resume"},
		{"unknown-singular", struct{ tea.Msg }{}, 1, "1 test"},
		{"unknown-plural", struct{ tea.Msg }{}, 4, "4 tests"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := labelForIntent(c.msg, c.count); got != c.want {
				t.Errorf("labelForIntent(%T, %d) = %q, want %q", c.msg, c.count, got, c.want)
			}
		})
	}
}

// ── 12. config intents ───────────────────────────────────────────────────────
//
// SwitchModeIntent is handled by beginSwitchMode/applySwitchMode, not
// handleConfigIntent (which now serves only SetEnvVarIntent and
// PreflightIntent), so these names say applySwitchMode.

func TestApplySwitchMode(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")

	var preflightModes []string
	pf := &fakeprovider.Fake{
		NameVal:  "test",
		Packages: []string{"./..."},
		ModesVal: []provider.Mode{
			{Name: "anonymous"}, {Name: "enterprise"},
		},
		EnvByMode: map[string][]provider.EnvVar{
			"enterprise": {{Key: "GITHUB_ENTERPRISE_SLUG", Required: true}},
		},
		PreflightFn: func(_ context.Context, m string, _ provider.TestRequirements) provider.PreflightReport {
			preflightModes = append(preflightModes, m)
			return provider.PreflightReport{Mode: m}
		},
	}

	var msgs []tea.Msg
	send := func(m tea.Msg) { msgs = append(msgs, m) }

	db := &dashboard{
		prov:      pf,
		root:      root,
		mode:      "anonymous",
		statePath: statePath,
		newRunner: func(_ io.Writer) testRunner { return &fakeRunner{} },
		list:      stubList(nil),
		red:       redact.New(nil),
		send:      send,
		logSink:   io.Discard,
	}

	db.applySwitchMode(context.Background(), tui.SwitchModeIntent{Mode: "enterprise"}, db.beginSwitchMode())

	if db.mode != "enterprise" {
		t.Errorf("db.mode = %q, want enterprise", db.mode)
	}

	var gotPreflight bool
	var gotMode string
	var gotEnvVars []provider.EnvVar
	for _, m := range msgs {
		if pm, ok := m.(tui.PreflightMsg); ok {
			gotPreflight = true
			gotMode = pm.Report.Mode
			gotEnvVars = pm.EnvVars
		}
	}
	if !gotPreflight {
		t.Fatal("expected PreflightMsg after SwitchModeIntent")
	}
	if gotMode != "enterprise" {
		t.Errorf("PreflightMsg.Report.Mode = %q, want enterprise", gotMode)
	}
	if len(gotEnvVars) == 0 {
		t.Error("expected non-empty EnvVars in PreflightMsg after mode switch")
	}
}

func TestHandleConfigIntentSetEnvVar(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")

	var preflightCalls int
	pf := &fakeprovider.Fake{
		NameVal:  "test",
		Packages: []string{"./..."},
		EnvByMode: map[string][]provider.EnvVar{
			"organization": {{Key: "GITHUB_OWNER", Required: true}},
		},
		PreflightFn: func(_ context.Context, m string, _ provider.TestRequirements) provider.PreflightReport {
			preflightCalls++
			return provider.PreflightReport{Mode: m}
		},
	}

	var msgs []tea.Msg
	send := func(m tea.Msg) { msgs = append(msgs, m) }
	env := map[string]string{"GITHUB_OWNER": ""}

	db := &dashboard{
		prov:      pf,
		root:      root,
		mode:      "organization",
		getenv:    func(key string) string { return env[key] },
		setenv:    func(key, val string) error { env[key] = val; return nil },
		statePath: statePath,
		newRunner: func(_ io.Writer) testRunner { return &fakeRunner{} },
		list:      stubList(nil),
		red:       redact.New(nil),
		plan:      planWithEligible("organization", nil),
		send:      send,
		logSink:   io.Discard,
	}

	const testOwner = "my-test-org"
	db.handleConfigIntent(context.Background(), tui.SetEnvVarIntent{
		Key:   "GITHUB_OWNER",
		Value: testOwner,
	})

	if env["GITHUB_OWNER"] != testOwner {
		t.Errorf("GITHUB_OWNER not set: got %q, want %q", env["GITHUB_OWNER"], testOwner)
	}
	if preflightCalls != 1 {
		t.Errorf("expected 1 preflight call, got %d", preflightCalls)
	}

	var gotPreflight bool
	for _, m := range msgs {
		if _, ok := m.(tui.PreflightMsg); ok {
			gotPreflight = true
		}
	}
	if !gotPreflight {
		t.Fatal("expected PreflightMsg after SetEnvVarIntent")
	}
}

func TestDashboardSetEnvVarUpdatesChildRunEnvOverlay(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")
	st := engine.State{Provider: "test", Mode: "organization"}
	st.Plan = planWithEligible("organization", []string{"TestAccEnv"})
	st.Results = []engine.PersistResult{{Test: "TestAccEnv", Status: "fail"}}
	if err := st.Save(statePath); err != nil {
		t.Fatalf("seeding state: %v", err)
	}

	env := map[string]string{"GITHUB_OWNER": "old-org"}
	pf := &fakeprovider.Fake{
		NameVal:  "test",
		Packages: []string{"./..."},
		EnvByMode: map[string][]provider.EnvVar{
			"organization": {{Key: "GITHUB_OWNER", Required: true}},
		},
		PreflightFn: func(_ context.Context, m string, _ provider.TestRequirements) provider.PreflightReport {
			return provider.PreflightReport{Mode: m}
		},
	}

	fr := &fakeRunner{}
	var msgs []tea.Msg
	send := func(m tea.Msg) { msgs = append(msgs, m) }
	db := &dashboard{
		prov:      pf,
		root:      root,
		mode:      "organization",
		getenv:    func(key string) string { return env[key] },
		setenv:    func(key, val string) error { env[key] = val; return nil },
		statePath: statePath,
		newRunner: func(_ io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccEnv"}),
		red:       redact.New(nil),
		plan:      planWithEligible("organization", []string{"TestAccEnv"}),
		groups:    []engine.Group{{Name: "tests", Tests: []string{"TestAccEnv"}}},
		allNames:  []string{"TestAccEnv"},
		send:      send,
		logSink:   io.Discard,
	}

	db.handleConfigIntent(context.Background(), tui.SetEnvVarIntent{
		Key:   "GITHUB_OWNER",
		Value: "new-org",
	})
	db.runIntent(context.Background(), tui.RetryAllIntent{})

	if !extraEnvContains(fr.spec.ExtraEnv, "GITHUB_OWNER=new-org") {
		t.Fatalf("child ExtraEnv = %v, want latest GITHUB_OWNER=new-org", fr.spec.ExtraEnv)
	}
	if extraEnvContains(fr.spec.ExtraEnv, "GITHUB_OWNER=old-org") {
		t.Fatalf("child ExtraEnv used stale env-file value: %v", fr.spec.ExtraEnv)
	}
}

func TestDashboardGetenvNilIgnoresProcessEnv(t *testing.T) {
	t.Setenv("TPT_DASHBOARD_SENTINEL", "process-value")

	if got := dashboardGetenv(&dashboard{})("TPT_DASHBOARD_SENTINEL"); got != "" {
		t.Fatalf("nil dashboard getenv read process env: got %q, want empty", got)
	}
}

func TestHandleConfigIntentSetEnvVarSetterError(t *testing.T) {
	env := map[string]string{"GITHUB_OWNER": "old-org"}
	setErr := errors.New("setter failed")
	pf := &fakeprovider.Fake{
		NameVal:  "test",
		Packages: []string{"./..."},
		EnvByMode: map[string][]provider.EnvVar{
			"organization": {{Key: "GITHUB_OWNER", Required: true}},
		},
		PreflightFn: func(_ context.Context, m string, _ provider.TestRequirements) provider.PreflightReport {
			return provider.PreflightReport{Mode: m}
		},
	}

	var msgs []tea.Msg
	send := func(m tea.Msg) { msgs = append(msgs, m) }
	db := &dashboard{
		prov:      pf,
		root:      t.TempDir(),
		mode:      "organization",
		getenv:    func(key string) string { return env[key] },
		setenv:    func(string, string) error { return setErr },
		statePath: filepath.Join(t.TempDir(), ".pulsar-state.json"),
		newRunner: func(_ io.Writer) testRunner { return &fakeRunner{} },
		list:      stubList(nil),
		red:       redact.New(nil),
		send:      send,
		logSink:   io.Discard,
	}

	db.handleConfigIntent(context.Background(), tui.SetEnvVarIntent{
		Key:   "GITHUB_OWNER",
		Value: "new-org",
	})

	if got := env["GITHUB_OWNER"]; got != "old-org" {
		t.Fatalf("setter error published override: got %q, want old-org", got)
	}
	var gotErr bool
	for _, m := range msgs {
		switch v := m.(type) {
		case tui.ErrMsg:
			gotErr = true
			if v.Err == nil || !strings.Contains(v.Err.Error(), "setting GITHUB_OWNER") {
				t.Fatalf("ErrMsg.Err = %v, want setting GITHUB_OWNER error", v.Err)
			}
		case tui.PreflightMsg:
			t.Fatal("PreflightMsg must not be sent when setter fails")
		}
	}
	if !gotErr {
		t.Fatal("expected ErrMsg when setter fails")
	}
}

func TestEnvOverlaySetGetPrecedenceAndConcurrentUse(t *testing.T) {
	base := map[string]string{
		"GITHUB_OWNER": "base-org",
		"BASE_ONLY":    "base-only",
	}
	var baseMu sync.Mutex
	var setCalls int
	overlay := newEnvOverlay(
		func(key string) string {
			baseMu.Lock()
			defer baseMu.Unlock()
			return base[key]
		},
		func(key, val string) error {
			baseMu.Lock()
			defer baseMu.Unlock()
			setCalls++
			base[key] = val
			return nil
		},
	)

	if got := overlay.Get("GITHUB_OWNER"); got != "base-org" {
		t.Fatalf("initial Get = %q, want base-org", got)
	}
	if err := overlay.Set("GITHUB_OWNER", "overlay-org"); err != nil {
		t.Fatalf("Set returned error: %v", err)
	}
	if got := overlay.Get("GITHUB_OWNER"); got != "overlay-org" {
		t.Fatalf("override Get = %q, want overlay-org", got)
	}
	if got := overlay.Get("BASE_ONLY"); got != "base-only" {
		t.Fatalf("base fallback Get = %q, want base-only", got)
	}
	if setCalls != 1 {
		t.Fatalf("base setter calls = %d, want 1", setCalls)
	}

	errOverlay := newEnvOverlay(
		func(key string) string {
			if key == "ERR_KEY" {
				return "base-value"
			}
			return ""
		},
		func(string, string) error { return errors.New("base setter failed") },
	)
	if err := errOverlay.Set("ERR_KEY", "new-value"); err == nil {
		t.Fatal("Set error = nil, want error")
	}
	if got := errOverlay.Get("ERR_KEY"); got != "base-value" {
		t.Fatalf("failed Set published override: got %q, want base-value", got)
	}

	var wg sync.WaitGroup
	for i := range 100 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := "KEY_" + string(rune('A'+(i%10)))
			val := "value"
			if err := overlay.Set(key, val); err != nil {
				t.Errorf("Set(%q) error: %v", key, err)
			}
			_ = overlay.Get(key)
			_ = overlay.Get("BASE_ONLY")
		}(i)
	}
	wg.Wait()
}

func TestHandleConfigIntentPreflight(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")

	var preflightCalls int
	pf := &fakeprovider.Fake{
		NameVal:  "test",
		Packages: []string{"./..."},
		PreflightFn: func(_ context.Context, m string, _ provider.TestRequirements) provider.PreflightReport {
			preflightCalls++
			return provider.PreflightReport{Mode: m}
		},
	}

	var msgs []tea.Msg
	send := func(m tea.Msg) { msgs = append(msgs, m) }

	db := &dashboard{
		prov:      pf,
		root:      root,
		mode:      "organization",
		statePath: statePath,
		newRunner: func(_ io.Writer) testRunner { return &fakeRunner{} },
		list:      stubList(nil),
		red:       redact.New(nil),
		plan:      planWithEligible("organization", nil),
		send:      send,
		logSink:   io.Discard,
	}

	db.handleConfigIntent(context.Background(), tui.PreflightIntent{})

	if preflightCalls != 1 {
		t.Errorf("expected 1 preflight call, got %d", preflightCalls)
	}

	var gotPreflight bool
	for _, m := range msgs {
		if _, ok := m.(tui.PreflightMsg); ok {
			gotPreflight = true
		}
	}
	if !gotPreflight {
		t.Fatal("expected PreflightMsg after PreflightIntent")
	}
}

// TestHandleConfigIntentSecretValueNeverLeaks ensures that the value set via
// SetEnvVarIntent never appears in any TUI message, including the PreflightMsg.
func TestHandleConfigIntentSecretValueNeverLeaks(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")

	const nonSecretVal = "value-set-via-intent"
	pf := &fakeprovider.Fake{
		NameVal:  "test",
		Packages: []string{"./..."},
		Secrets:  []string{"GITHUB_TOKEN"},
		PreflightFn: func(_ context.Context, m string, _ provider.TestRequirements) provider.PreflightReport {
			return provider.PreflightReport{Mode: m}
		},
	}

	var msgs []tea.Msg
	send := func(m tea.Msg) { msgs = append(msgs, m) }

	db := &dashboard{
		prov:      pf,
		root:      root,
		mode:      "organization",
		statePath: statePath,
		newRunner: func(_ io.Writer) testRunner { return &fakeRunner{} },
		list:      stubList(nil),
		red:       redact.New(nil),
		plan:      planWithEligible("organization", nil),
		send:      send,
		logSink:   io.Discard,
	}

	// Use a non-secret key; the value should not leak in any message.
	db.handleConfigIntent(context.Background(), tui.SetEnvVarIntent{
		Key:   "GITHUB_OWNER",
		Value: nonSecretVal,
	})

	for _, m := range msgs {
		switch pm := m.(type) {
		case tui.PreflightMsg:
			// PreflightReport detail must not echo the set value.
			for _, c := range pm.Report.Checks {
				if strings.Contains(c.Detail, nonSecretVal) {
					t.Errorf("set value leaked into PreflightMsg check detail: %q", c.Detail)
				}
			}
		}
	}
}

// ── 13. concurrent config intent does not race with runIntent ─────────────────

// firstSinkBarrierRunner pauses immediately after the first streamed result.
// Tests use the two channels to make an overlapping config operation
// deterministic instead of relying on scheduler timing.
type firstSinkBarrierRunner struct {
	result    engine.RunResult
	firstSink chan struct{}
	proceed   chan struct{}
}

func (s *firstSinkBarrierRunner) Run(_ context.Context, _ engine.RunSpec, sink func(engine.TestResult)) (engine.RunResult, error) {
	for i, tr := range s.result.Tests {
		if sink != nil {
			sink(tr)
		}
		if i == 0 {
			close(s.firstSink)
			<-s.proceed
		}
	}
	return s.result, nil
}

// TestRunIntentConcurrentConfigIsSafe fires config intents while a run is
// streaming results and asserts no data race occurs and no secret leaks into
// any captured message. Run under -race to exercise the cfgMu snapshot path.
func TestRunIntentConcurrentConfigIsSafe(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")

	const secret = "ghp_RACE_TEST_SECRET"
	st := engine.State{Provider: "test", Mode: "anonymous"}
	st.Plan = planWithEligible("anonymous", []string{"TestAccRace1", "TestAccRace2"})
	st.Results = []engine.PersistResult{
		{Test: "TestAccRace1", Status: "fail"},
		{Test: "TestAccRace2", Status: "fail"},
	}
	if err := st.Save(statePath); err != nil {
		t.Fatalf("seeding state: %v", err)
	}

	slow := &firstSinkBarrierRunner{
		result: engine.RunResult{
			Tests: []engine.TestResult{
				{Name: "TestAccRace1", Status: provider.StatusFail, Output: []string{"token=" + secret}},
				{Name: "TestAccRace2", Status: provider.StatusFail, Output: []string{"line2"}},
			},
		},
		firstSink: make(chan struct{}),
		proceed:   make(chan struct{}),
	}

	var mu sync.Mutex
	var captured []tea.Msg
	send := func(m tea.Msg) {
		mu.Lock()
		captured = append(captured, m)
		mu.Unlock()
	}

	pf := &fakeprovider.Fake{
		NameVal:  "test",
		Packages: []string{"./..."},
		Secrets:  []string{"RACE_SECRET_KEY"},
		ModesVal: []provider.Mode{{Name: "anonymous"}, {Name: "enterprise"}},
		PreflightFn: func(_ context.Context, m string, _ provider.TestRequirements) provider.PreflightReport {
			return provider.PreflightReport{Mode: m}
		},
		RequirementsFn: allowAllRequirements,
	}

	db := &dashboard{
		prov:      pf,
		root:      root,
		mode:      "anonymous",
		statePath: statePath,
		newRunner: func(_ io.Writer) testRunner { return slow },
		list:      stubList([]string{"TestAccRace1", "TestAccRace2"}),
		red:       redact.New([]string{secret}),
		plan:      planWithEligible("anonymous", []string{"TestAccRace1", "TestAccRace2"}),
		groups:    []engine.Group{{Name: "tests", Tests: []string{"TestAccRace1", "TestAccRace2"}}},
		allNames:  []string{"TestAccRace1", "TestAccRace2"},
		send:      send,
		logSink:   io.Discard,
	}

	done := make(chan struct{})
	go func() {
		defer close(done)
		db.runIntent(context.Background(), tui.RetryAllIntent{})
	}()

	// Fire config intents while the run is deterministically paused after its
	// first streamed result.
	<-slow.firstSink
	db.handleConfigIntent(context.Background(), tui.SetEnvVarIntent{
		Key:   "RACE_SECRET_KEY",
		Value: "new-secret-value",
	})
	db.applySwitchMode(context.Background(), tui.SwitchModeIntent{Mode: "enterprise"}, db.beginSwitchMode())
	close(slow.proceed)

	<-done

	mu.Lock()
	msgs := make([]tea.Msg, len(captured))
	copy(msgs, captured)
	mu.Unlock()

	var gotDone bool
	for _, m := range msgs {
		if _, ok := m.(tui.RunDoneMsg); ok {
			gotDone = true
		}
	}
	if !gotDone {
		t.Error("expected RunDoneMsg; run did not complete")
	}

	// Secret from the run must not appear in any TestUpdateMsg.
	for _, m := range msgs {
		if um, ok := m.(tui.TestUpdateMsg); ok {
			for _, line := range um.Result.Output {
				if strings.Contains(line, secret) {
					t.Errorf("raw secret found in TestUpdateMsg: %q", line)
				}
			}
		}
	}
}

// ── C12.4: SetEnvVarIntent rejects secret keys ────────────────────────────────

// TestHandleConfigIntentSetEnvVarRejectsSecretKey verifies that
// handleConfigIntent refuses to set a key listed in SecretEnvKeys,
// sends an ErrMsg, and never emits a PreflightMsg for that intent.
func TestHandleConfigIntentSetEnvVarRejectsSecretKey(t *testing.T) {
	const secretKey = "GITHUB_TOKEN"
	const originalVal = "original-token-value"
	env := map[string]string{secretKey: originalVal}
	var setCalled bool

	pf := &fakeprovider.Fake{
		NameVal:  "test",
		Packages: []string{"./..."},
		Secrets:  []string{secretKey},
		PreflightFn: func(_ context.Context, m string, _ provider.TestRequirements) provider.PreflightReport {
			return provider.PreflightReport{Mode: m}
		},
	}

	var msgs []tea.Msg
	send := func(m tea.Msg) { msgs = append(msgs, m) }

	db := &dashboard{
		prov:      pf,
		root:      t.TempDir(),
		mode:      "organization",
		getenv:    func(key string) string { return env[key] },
		setenv:    func(key, val string) error { setCalled = true; env[key] = val; return nil },
		statePath: filepath.Join(t.TempDir(), ".pulsar-state.json"),
		newRunner: func(_ io.Writer) testRunner { return &fakeRunner{} },
		list:      stubList(nil),
		red:       redact.New(nil),
		send:      send,
		logSink:   io.Discard,
	}

	db.handleConfigIntent(context.Background(), tui.SetEnvVarIntent{
		Key:   secretKey,
		Value: "should-not-be-set",
	})

	if got := env[secretKey]; got != originalVal {
		t.Errorf("env mutated: %s=%q, want %q", secretKey, got, originalVal)
	}
	if setCalled {
		t.Error("setenv must not be called for a secret key")
	}

	var gotErr bool
	for _, m := range msgs {
		if em, ok := m.(tui.ErrMsg); ok {
			gotErr = true
			if em.Err == nil || !strings.Contains(em.Err.Error(), "secret") {
				t.Errorf("ErrMsg.Err = %v, expected it to mention 'secret'", em.Err)
			}
		}
	}
	if !gotErr {
		t.Error("expected ErrMsg when secret key is rejected")
	}

	for _, m := range msgs {
		if _, ok := m.(tui.PreflightMsg); ok {
			t.Error("PreflightMsg must not be sent when secret key is rejected")
		}
	}
}

// TestHandleConfigIntentSetEnvVarNonSecretStillWorks confirms the non-secret
// path is unaffected when the provider has secret keys configured.
func TestHandleConfigIntentSetEnvVarNonSecretStillWorks(t *testing.T) {
	const nonSecretKey = "GITHUB_OWNER"
	const wantVal = "my-org-via-fix"
	env := map[string]string{nonSecretKey: ""}

	pf := &fakeprovider.Fake{
		NameVal:  "test",
		Packages: []string{"./..."},
		Secrets:  []string{"GITHUB_TOKEN"},
		PreflightFn: func(_ context.Context, m string, _ provider.TestRequirements) provider.PreflightReport {
			return provider.PreflightReport{Mode: m}
		},
	}

	var msgs []tea.Msg
	send := func(m tea.Msg) { msgs = append(msgs, m) }

	db := &dashboard{
		prov:      pf,
		root:      t.TempDir(),
		mode:      "organization",
		getenv:    func(key string) string { return env[key] },
		setenv:    func(key, val string) error { env[key] = val; return nil },
		statePath: filepath.Join(t.TempDir(), ".pulsar-state.json"),
		newRunner: func(_ io.Writer) testRunner { return &fakeRunner{} },
		list:      stubList(nil),
		red:       redact.New(nil),
		plan:      planWithEligible("organization", nil),
		send:      send,
		logSink:   io.Discard,
	}

	db.handleConfigIntent(context.Background(), tui.SetEnvVarIntent{
		Key:   nonSecretKey,
		Value: wantVal,
	})

	if got := env[nonSecretKey]; got != wantVal {
		t.Errorf("GITHUB_OWNER = %q, want %q", got, wantVal)
	}

	var gotPreflight bool
	for _, m := range msgs {
		if _, ok := m.(tui.PreflightMsg); ok {
			gotPreflight = true
		}
	}
	if !gotPreflight {
		t.Error("expected PreflightMsg after non-secret SetEnvVarIntent")
	}
}

// ── 14. Task 10: dashboard plans before preflight/groups/retry/resume ────────

// TestPlanDashboardFiltersToEligibleTests verifies that planDashboard builds
// an ExecutionPlan for the requested mode and returns groups/names already
// narrowed to Plan.Eligible: a test whose requirements restrict it to a
// different mode must not appear in either return value, while a compatible
// test must.
func TestPlanDashboardFiltersToEligibleTests(t *testing.T) {
	requirements := func(name string) (provider.TestRequirements, bool) {
		if name == "TestAccEnterpriseOnly" {
			return provider.TestRequirements{Modes: []string{"enterprise"}}, true
		}
		return provider.TestRequirements{}, true
	}
	pf := &fakeprovider.Fake{
		NameVal:        "test",
		Packages:       []string{"./..."},
		ModesVal:       []provider.Mode{{Name: "anonymous"}, {Name: "enterprise"}},
		RequirementsFn: requirements,
	}
	names := []string{"TestAccOpen", "TestAccEnterpriseOnly"}

	groups, allNames, plan, err := planDashboard(context.Background(), stubList(names), pf, t.TempDir(), "anonymous", false, redact.New(nil))
	if err != nil {
		t.Fatalf("planDashboard error: %v", err)
	}

	if plan.Mode != "anonymous" {
		t.Errorf("plan.Mode = %q, want anonymous", plan.Mode)
	}
	for _, n := range allNames {
		if n == "TestAccEnterpriseOnly" {
			t.Errorf("allNames = %v; must not include mode-incompatible TestAccEnterpriseOnly", allNames)
		}
	}
	for _, g := range groups {
		for _, n := range g.Tests {
			if n == "TestAccEnterpriseOnly" {
				t.Errorf("groups = %+v; must not include mode-incompatible TestAccEnterpriseOnly", groups)
			}
		}
	}
	if !reflect.DeepEqual(allNames, []string{"TestAccOpen"}) {
		t.Errorf("allNames = %v, want [TestAccOpen]", allNames)
	}
}

// TestPlanDashboardDiscoveryFailureReturnsError verifies that a listing
// failure (e.g. the provider package failing to build) is surfaced as an
// error rather than an empty, silently-truncated plan.
func TestPlanDashboardDiscoveryFailureReturnsError(t *testing.T) {
	pf := &fakeprovider.Fake{NameVal: "test", Packages: []string{"./..."}, ModesVal: []provider.Mode{{Name: "anonymous"}}}
	wantErr := errors.New("boom")
	_, _, _, err := planDashboard(context.Background(), failList(wantErr), pf, t.TempDir(), "anonymous", false, redact.New(nil))
	if err == nil {
		t.Fatal("expected an error from planDashboard on discovery failure")
	}
}

// TestSeedPreflightUsesRequirementsFromPlan verifies that seed calls Preflight
// with requirements sourced from db.plan (Scopes/Capabilities/SideEffects),
// not an empty provider.TestRequirements{}.
func TestSeedPreflightUsesRequirementsFromPlan(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")

	var gotReq provider.TestRequirements
	pf := &fakeprovider.Fake{
		NameVal:  "test",
		Packages: []string{"./..."},
		PreflightFn: func(_ context.Context, _ string, req provider.TestRequirements) provider.PreflightReport {
			gotReq = req
			return provider.PreflightReport{Mode: "anonymous"}
		},
	}

	allNames := []string{"TestAccA"}
	plan := &engine.ExecutionPlan{
		Mode:         "anonymous",
		Eligible:     allNames,
		Scopes:       []string{"repo"},
		Capabilities: []string{"organization"},
		SideEffects:  []string{"repository"},
	}

	var msgs []tea.Msg
	send := func(m tea.Msg) { msgs = append(msgs, m) }

	db := &dashboard{
		prov:      pf,
		root:      root,
		mode:      "anonymous",
		statePath: statePath,
		newRunner: func(_ io.Writer) testRunner { return &fakeRunner{} },
		list:      stubList(allNames),
		red:       redact.New(nil),
		plan:      plan,
		groups:    []engine.Group{{Name: "tests", Tests: allNames}},
		allNames:  allNames,
		send:      send,
		logSink:   io.Discard,
	}

	db.seed(context.Background())

	if !reflect.DeepEqual(gotReq.Scopes, plan.Scopes) {
		t.Errorf("Preflight Scopes = %v, want %v", gotReq.Scopes, plan.Scopes)
	}
	if !reflect.DeepEqual(gotReq.Capabilities, plan.Capabilities) {
		t.Errorf("Preflight Capabilities = %v, want %v", gotReq.Capabilities, plan.Capabilities)
	}
	if !reflect.DeepEqual(gotReq.SideEffects, plan.SideEffects) {
		t.Errorf("Preflight SideEffects = %v, want %v", gotReq.SideEffects, plan.SideEffects)
	}
}

// TestSeedSendsGroupsMsgWithPlan verifies that seed's GroupsMsg carries the
// dashboard's current plan alongside the (already eligible-filtered) groups.
func TestSeedSendsGroupsMsgWithPlan(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")

	pf := &fakeprovider.Fake{NameVal: "test", Packages: []string{"./..."}}
	allNames := []string{"TestAccA"}
	plan := &engine.ExecutionPlan{Mode: "anonymous", Eligible: allNames}
	groups := []engine.Group{{Name: "tests", Tests: allNames}}

	var msgs []tea.Msg
	send := func(m tea.Msg) { msgs = append(msgs, m) }

	db := &dashboard{
		prov:      pf,
		root:      root,
		mode:      "anonymous",
		statePath: statePath,
		newRunner: func(_ io.Writer) testRunner { return &fakeRunner{} },
		list:      stubList(allNames),
		red:       redact.New(nil),
		plan:      plan,
		groups:    groups,
		allNames:  allNames,
		send:      send,
		logSink:   io.Discard,
	}

	db.seed(context.Background())

	var gm *tui.GroupsMsg
	for i := range msgs {
		if v, ok := msgs[i].(tui.GroupsMsg); ok {
			gm = &v
		}
	}
	if gm == nil {
		t.Fatal("expected GroupsMsg")
	}
	if gm.Plan != plan {
		t.Errorf("GroupsMsg.Plan = %p, want %p (db.plan)", gm.Plan, plan)
	}
	if !reflect.DeepEqual(gm.Groups, groups) {
		t.Errorf("GroupsMsg.Groups = %+v, want %+v", gm.Groups, groups)
	}
}

// TestRunIntentRetryAllIntersectsPlanEligibility verifies that RetryAllIntent
// only retries failed tests that are also in the persisted plan's Eligible
// set: a failed test excluded by the persisted plan must be skipped.
func TestRunIntentRetryAllIntersectsPlanEligibility(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")

	st := engine.State{Provider: "test", Mode: "anonymous"}
	st.Plan = planWithEligible("anonymous", []string{"TestAccEligible"}) // TestAccExcluded is NOT eligible
	st.Results = []engine.PersistResult{
		{Test: "TestAccEligible", Status: "fail"},
		{Test: "TestAccExcluded", Status: "fail"},
	}
	if err := st.Save(statePath); err != nil {
		t.Fatalf("seeding state: %v", err)
	}

	fr := &fakeRunner{}
	pf := &fakeprovider.Fake{NameVal: "test", Packages: []string{"./..."}}

	var msgs []tea.Msg
	send := func(m tea.Msg) { msgs = append(msgs, m) }

	db := &dashboard{
		prov:      pf,
		root:      root,
		mode:      "anonymous",
		statePath: statePath,
		newRunner: func(_ io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccEligible", "TestAccExcluded"}),
		red:       redact.New(nil),
		groups:    []engine.Group{{Name: "tests", Tests: []string{"TestAccEligible", "TestAccExcluded"}}},
		allNames:  []string{"TestAccEligible", "TestAccExcluded"},
		send:      send,
		logSink:   io.Discard,
	}

	db.runIntent(context.Background(), tui.RetryAllIntent{})

	wantPattern := engine.RunPattern([]string{"TestAccEligible"})
	if fr.spec.Pattern != wantPattern {
		t.Errorf("runner pattern = %q, want %q", fr.spec.Pattern, wantPattern)
	}
	for _, m := range msgs {
		if rs, ok := m.(tui.RunStartedMsg); ok && rs.Total != 1 {
			t.Errorf("RunStartedMsg.Total = %d, want 1", rs.Total)
		}
	}
}

// TestRunIntentRetryGroupIntersectsPlanEligibility verifies that
// RetryGroupIntent's whole-group fallback (no prior failures recorded for the
// group) only includes members that are in the persisted plan's Eligible set.
func TestRunIntentRetryGroupIntersectsPlanEligibility(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")

	st := engine.State{Provider: "test", Mode: "anonymous"}
	st.Plan = planWithEligible("anonymous", []string{"TestAccGroupA"}) // TestAccGroupB is NOT eligible
	if err := st.Save(statePath); err != nil {
		t.Fatalf("seeding state: %v", err)
	}

	fr := &fakeRunner{}
	pf := &fakeprovider.Fake{NameVal: "test", Packages: []string{"./..."}}

	send := func(tea.Msg) {}

	db := &dashboard{
		prov:      pf,
		root:      root,
		mode:      "anonymous",
		statePath: statePath,
		newRunner: func(_ io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccGroupA", "TestAccGroupB"}),
		red:       redact.New(nil),
		groups:    []engine.Group{{Name: "grp", Tests: []string{"TestAccGroupA", "TestAccGroupB"}}},
		allNames:  []string{"TestAccGroupA", "TestAccGroupB"},
		send:      send,
		logSink:   io.Discard,
	}

	db.runIntent(context.Background(), tui.RetryGroupIntent{Group: "grp"})

	wantPattern := engine.RunPattern([]string{"TestAccGroupA"})
	if fr.spec.Pattern != wantPattern {
		t.Errorf("runner pattern = %q, want %q", fr.spec.Pattern, wantPattern)
	}
}

// TestRunIntentResumeIntersectsPlanEligibility verifies that ResumeIntent's
// not-run calculation is bounded by the persisted plan's Eligible set: a
// never-run test that the persisted plan excluded must not be resumed.
func TestRunIntentResumeIntersectsPlanEligibility(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")

	st := engine.State{Provider: "test", Mode: "anonymous"}
	// TestAccExcluded was never run and is excluded from the persisted plan;
	// TestAccEligible was never run and remains eligible.
	st.Plan = planWithEligible("anonymous", []string{"TestAccEligible"})
	if err := st.Save(statePath); err != nil {
		t.Fatalf("seeding state: %v", err)
	}

	fr := &fakeRunner{}
	pf := &fakeprovider.Fake{NameVal: "test", Packages: []string{"./..."}}
	send := func(tea.Msg) {}

	db := &dashboard{
		prov:      pf,
		root:      root,
		mode:      "anonymous",
		statePath: statePath,
		newRunner: func(_ io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccEligible", "TestAccExcluded"}),
		red:       redact.New(nil),
		groups:    []engine.Group{{Name: "tests", Tests: []string{"TestAccEligible", "TestAccExcluded"}}},
		allNames:  []string{"TestAccEligible", "TestAccExcluded"},
		send:      send,
		logSink:   io.Discard,
	}

	db.runIntent(context.Background(), tui.ResumeIntent{})

	wantPattern := engine.RunPattern([]string{"TestAccEligible"})
	if fr.spec.Pattern != wantPattern {
		t.Errorf("runner pattern = %q, want %q", fr.spec.Pattern, wantPattern)
	}
}

// TestRunIntentRetryTestOnlyWhenEligible verifies that RetryTestIntent runs
// the requested test when it belongs to the dashboard's current eligible set
// (db.allNames) and sends an ErrMsg without invoking the runner otherwise.
// RetryTestIntent has no CLI equivalent and is bound to whatever the
// dashboard is currently displaying, so it does not require a persisted
// plan — only current-mode eligibility.
func TestRunIntentRetryTestOnlyWhenEligible(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")
	// No state file at all: proves RetryTestIntent does not need a persisted
	// plan when the test is currently eligible.

	fr := &fakeRunner{}
	pf := &fakeprovider.Fake{NameVal: "test", Packages: []string{"./..."}}

	var msgs []tea.Msg
	send := func(m tea.Msg) { msgs = append(msgs, m) }

	db := &dashboard{
		prov:      pf,
		root:      root,
		mode:      "anonymous",
		statePath: statePath,
		newRunner: func(_ io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccEligible"}),
		red:       redact.New(nil),
		groups:    []engine.Group{{Name: "tests", Tests: []string{"TestAccEligible"}}},
		allNames:  []string{"TestAccEligible"},
		send:      send,
		logSink:   io.Discard,
	}

	db.runIntent(context.Background(), tui.RetryTestIntent{Test: "TestAccEligible"})

	wantPattern := engine.RunPattern([]string{"TestAccEligible"})
	if fr.spec.Pattern != wantPattern {
		t.Errorf("runner pattern = %q, want %q", fr.spec.Pattern, wantPattern)
	}
	if fr.calls != 1 {
		t.Fatalf("expected the runner to be invoked once, got %d calls", fr.calls)
	}

	// Now retry a test outside the current eligible set: the runner must not
	// be invoked again, and an ErrMsg must be sent.
	msgs = nil
	db.runIntent(context.Background(), tui.RetryTestIntent{Test: "TestAccOutside"})
	if fr.calls != 1 {
		t.Errorf("runner must not be invoked for an ineligible test, got %d calls", fr.calls)
	}
	var gotErr bool
	for _, m := range msgs {
		if _, ok := m.(tui.ErrMsg); ok {
			gotErr = true
		}
		if _, ok := m.(tui.RunStartedMsg); ok {
			t.Error("RunStartedMsg must not be sent for an ineligible test")
		}
	}
	if !gotErr {
		t.Error("expected ErrMsg for a test outside the current eligible set")
	}
}

// TestRunIntentRetryTestBootstrapsPlanWhenStateHasNone verifies that saving
// state after a RetryTestIntent run on a root with no prior state (no
// persisted Plan to carry forward) attaches the dashboard's current plan,
// rather than leaving Plan nil forever. Without this, a dashboard used only
// via single-test retries could never satisfy RetryAllIntent/RetryGroupIntent/
// ResumeIntent's persisted-plan requirement afterward.
func TestRunIntentRetryTestBootstrapsPlanWhenStateHasNone(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")

	fr := &fakeRunner{}
	pf := &fakeprovider.Fake{NameVal: "test", Packages: []string{"./..."}}
	send := func(tea.Msg) {}

	currentPlan := planWithEligible("anonymous", []string{"TestAccEligible"})
	db := &dashboard{
		prov:      pf,
		root:      root,
		mode:      "anonymous",
		statePath: statePath,
		newRunner: func(_ io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccEligible"}),
		red:       redact.New(nil),
		plan:      currentPlan,
		groups:    []engine.Group{{Name: "tests", Tests: []string{"TestAccEligible"}}},
		allNames:  []string{"TestAccEligible"},
		send:      send,
		logSink:   io.Discard,
	}

	db.runIntent(context.Background(), tui.RetryTestIntent{Test: "TestAccEligible"})

	saved, err := engine.Load(statePath)
	if err != nil {
		t.Fatalf("loading saved state: %v", err)
	}
	if saved.Plan == nil {
		t.Fatal("expected saved state to carry a bootstrapped Plan, got nil")
	}
	if !reflect.DeepEqual(saved.Plan.Eligible, currentPlan.Eligible) {
		t.Errorf("saved.Plan.Eligible = %v, want %v", saved.Plan.Eligible, currentPlan.Eligible)
	}
}

// TestRunIntentRetryAllPreservesPersistedPlanUnchanged verifies that a
// successful RetryAllIntent run saves state with the ORIGINAL persisted Plan
// carried forward unchanged — never overwritten by the dashboard's current
// plan — exactly mirroring the CLI retry/resume convention in runWithSpec.
func TestRunIntentRetryAllPreservesPersistedPlanUnchanged(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")

	persistedPlan := planWithEligible("anonymous", []string{"TestAccEligible"})
	st := engine.State{Provider: "test", Mode: "anonymous", Plan: persistedPlan}
	st.Results = []engine.PersistResult{{Test: "TestAccEligible", Status: "fail"}}
	if err := st.Save(statePath); err != nil {
		t.Fatalf("seeding state: %v", err)
	}

	fr := &fakeRunner{}
	pf := &fakeprovider.Fake{NameVal: "test", Packages: []string{"./..."}}
	send := func(tea.Msg) {}

	// db.plan deliberately differs from the persisted plan (as it would after
	// a mode switch) to prove the persisted plan wins for retry/resume saves.
	currentPlan := planWithEligible("enterprise", []string{"TestAccEligible", "TestAccEnterprise"})
	db := &dashboard{
		prov:      pf,
		root:      root,
		mode:      "anonymous",
		statePath: statePath,
		newRunner: func(_ io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccEligible"}),
		red:       redact.New(nil),
		plan:      currentPlan,
		groups:    []engine.Group{{Name: "tests", Tests: []string{"TestAccEligible"}}},
		allNames:  []string{"TestAccEligible"},
		send:      send,
		logSink:   io.Discard,
	}

	db.runIntent(context.Background(), tui.RetryAllIntent{})

	saved, err := engine.Load(statePath)
	if err != nil {
		t.Fatalf("loading saved state: %v", err)
	}
	if saved.Plan == nil {
		t.Fatal("expected saved state to carry a Plan, got nil")
	}
	if saved.Plan.Mode != "anonymous" {
		t.Errorf("saved.Plan.Mode = %q, want unchanged %q (persisted, not db.plan's %q)", saved.Plan.Mode, "anonymous", currentPlan.Mode)
	}
	if !reflect.DeepEqual(saved.Plan.Eligible, persistedPlan.Eligible) {
		t.Errorf("saved.Plan.Eligible = %v, want unchanged %v", saved.Plan.Eligible, persistedPlan.Eligible)
	}
}

func TestRunIntentPreservesPersistedOrphanAccounting(t *testing.T) {
	for _, tc := range []struct {
		name   string
		intent tea.Msg
	}{
		{name: "retry test", intent: tui.RetryTestIntent{Test: "TestAccEligible"}},
		{name: "retry all", intent: tui.RetryAllIntent{}},
		{name: "resume", intent: tui.ResumeIntent{}},
	} {
		t.Run(tc.name, func(t *testing.T) {
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
				Mode:             "organization",
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
				Mode:     "organization",
				Plan:     planWithEligible("organization", []string{"TestAccEligible"}),
				Results:  []engine.PersistResult{{Test: "TestAccEligible", Status: "fail"}},
				Orphans:  accounting,
			}
			if err := state.Save(statePath); err != nil {
				t.Fatalf("seeding state: %v", err)
			}

			runner := &fakeRunner{}
			db := &dashboard{
				prov:      &fakeprovider.Fake{NameVal: "test", Packages: []string{"./..."}},
				root:      root,
				mode:      "organization",
				statePath: statePath,
				newRunner: func(_ io.Writer) testRunner { return runner },
				list:      stubList([]string{"TestAccEligible"}),
				red:       redact.New(nil),
				plan:      planWithEligible("organization", []string{"TestAccEligible"}),
				groups:    []engine.Group{{Name: "tests", Tests: []string{"TestAccEligible"}}},
				allNames:  []string{"TestAccEligible"},
				send:      func(tea.Msg) {},
				logSink:   io.Discard,
			}

			db.runIntent(context.Background(), tc.intent)

			if runner.calls != 1 {
				t.Fatalf("runner calls = %d, want 1", runner.calls)
			}
			saved, err := engine.Load(statePath)
			if err != nil {
				t.Fatalf("loading saved state: %v", err)
			}
			if !reflect.DeepEqual(saved.Orphans, accounting) {
				t.Fatalf("saved orphan accounting = %+v, want unchanged %+v", saved.Orphans, accounting)
			}
		})
	}
}

// TestApplySwitchModeRebuildsPlanAndPreflights verifies that
// SwitchModeIntent rebuilds db.plan/db.groups/db.allNames for the new mode
// (excluding mode-incompatible tests) and emits PreflightMsg before
// GroupsMsg, both reflecting the new mode.
func TestApplySwitchModeRebuildsPlanAndPreflights(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")

	requirements := func(name string) (provider.TestRequirements, bool) {
		if name == "TestAccEnterpriseOnly" {
			return provider.TestRequirements{Modes: []string{"enterprise"}}, true
		}
		return provider.TestRequirements{}, true
	}
	var preflightModes []string
	pf := &fakeprovider.Fake{
		NameVal:  "test",
		Packages: []string{"./..."},
		ModesVal: []provider.Mode{{Name: "anonymous"}, {Name: "enterprise"}},
		EnvByMode: map[string][]provider.EnvVar{
			"enterprise": {{Key: "GITHUB_ENTERPRISE_SLUG", Required: true}},
		},
		RequirementsFn: requirements,
		PreflightFn: func(_ context.Context, m string, _ provider.TestRequirements) provider.PreflightReport {
			preflightModes = append(preflightModes, m)
			return provider.PreflightReport{Mode: m}
		},
	}

	var msgs []tea.Msg
	send := func(m tea.Msg) { msgs = append(msgs, m) }

	db := &dashboard{
		prov:      pf,
		root:      root,
		mode:      "anonymous",
		statePath: statePath,
		newRunner: func(_ io.Writer) testRunner { return &fakeRunner{} },
		list:      stubList([]string{"TestAccOpen", "TestAccEnterpriseOnly"}),
		red:       redact.New(nil),
		plan:      planWithEligible("anonymous", []string{"TestAccOpen", "TestAccEnterpriseOnly"}),
		groups:    []engine.Group{{Name: "tests", Tests: []string{"TestAccOpen", "TestAccEnterpriseOnly"}}},
		allNames:  []string{"TestAccOpen", "TestAccEnterpriseOnly"},
		send:      send,
		logSink:   io.Discard,
	}

	db.applySwitchMode(context.Background(), tui.SwitchModeIntent{Mode: "enterprise"}, db.beginSwitchMode())

	if db.mode != "enterprise" {
		t.Errorf("db.mode = %q, want enterprise", db.mode)
	}
	if db.plan == nil || db.plan.Mode != "enterprise" {
		t.Fatalf("db.plan = %+v, want Mode=enterprise", db.plan)
	}
	if !reflect.DeepEqual(db.allNames, []string{"TestAccOpen", "TestAccEnterpriseOnly"}) {
		t.Errorf("db.allNames = %v, want both tests eligible under enterprise", db.allNames)
	}

	// Message order: PreflightMsg must be emitted before GroupsMsg.
	var preflightIdx, groupsIdx = -1, -1
	for i, m := range msgs {
		switch m.(type) {
		case tui.PreflightMsg:
			if preflightIdx == -1 {
				preflightIdx = i
			}
		case tui.GroupsMsg:
			if groupsIdx == -1 {
				groupsIdx = i
			}
		}
	}
	if preflightIdx == -1 || groupsIdx == -1 {
		t.Fatalf("expected both PreflightMsg and GroupsMsg, got %#v", msgs)
	}
	if preflightIdx > groupsIdx {
		t.Errorf("PreflightMsg (index %d) must be sent before GroupsMsg (index %d)", preflightIdx, groupsIdx)
	}
	if len(preflightModes) == 0 || preflightModes[len(preflightModes)-1] != "enterprise" {
		t.Errorf("Preflight was not called with the new mode: %v", preflightModes)
	}
}

// TestApplySwitchModePlanFailureSendsErrMsg verifies that a
// planning failure during SwitchModeIntent sends ErrMsg, never calls
// Preflight, and never mutates db.mode/db.plan/db.groups/db.allNames.
func TestApplySwitchModePlanFailureSendsErrMsg(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")

	pf := &fakeprovider.Fake{
		NameVal:  "test",
		Packages: []string{"./..."},
		ModesVal: []provider.Mode{{Name: "anonymous"}, {Name: "enterprise"}},
		PreflightFn: func(context.Context, string, provider.TestRequirements) provider.PreflightReport {
			t.Fatal("Preflight must not be called when planning fails")
			return provider.PreflightReport{}
		},
	}

	origPlan := planWithEligible("anonymous", []string{"TestAccOpen"})
	origGroups := []engine.Group{{Name: "tests", Tests: []string{"TestAccOpen"}}}
	origNames := []string{"TestAccOpen"}

	var msgs []tea.Msg
	send := func(m tea.Msg) { msgs = append(msgs, m) }

	wantErr := errors.New("listing tests: boom")
	db := &dashboard{
		prov:      pf,
		root:      root,
		mode:      "anonymous",
		statePath: statePath,
		newRunner: func(_ io.Writer) testRunner { return &fakeRunner{} },
		list:      failList(wantErr),
		red:       redact.New(nil),
		plan:      origPlan,
		groups:    origGroups,
		allNames:  origNames,
		send:      send,
		logSink:   io.Discard,
	}

	db.applySwitchMode(context.Background(), tui.SwitchModeIntent{Mode: "enterprise"}, db.beginSwitchMode())

	if db.mode != "anonymous" {
		t.Errorf("db.mode = %q, want unchanged anonymous", db.mode)
	}
	if db.plan != origPlan {
		t.Errorf("db.plan changed on planning failure: %p, want unchanged %p", db.plan, origPlan)
	}
	if !reflect.DeepEqual(db.groups, origGroups) {
		t.Errorf("db.groups changed on planning failure: %+v", db.groups)
	}
	if !reflect.DeepEqual(db.allNames, origNames) {
		t.Errorf("db.allNames changed on planning failure: %v", db.allNames)
	}

	var gotErr bool
	for _, m := range msgs {
		switch m.(type) {
		case tui.ErrMsg:
			gotErr = true
		case tui.PreflightMsg:
			t.Error("PreflightMsg must not be sent when planning fails")
		case tui.GroupsMsg:
			t.Error("GroupsMsg must not be sent when planning fails")
		}
	}
	if !gotErr {
		t.Error("expected ErrMsg when planning fails during SwitchModeIntent")
	}
}

// TestRunIntentPlanRequiredWhenPersistedPlanNil verifies that RetryAllIntent,
// RetryGroupIntent, and ResumeIntent each send the exact shared
// state/plan-required error (never the generic "nothing to run" message) and
// never invoke the runner when the persisted state has no execution plan
// (legacy pre-Task-6 state, or a root with no run history at all).
func TestRunIntentPlanRequiredWhenPersistedPlanNil(t *testing.T) {
	cases := []struct {
		name string
		msg  tea.Msg
	}{
		{"RetryAllIntent", tui.RetryAllIntent{}},
		{"RetryGroupIntent", tui.RetryGroupIntent{Group: "tests"}},
		{"ResumeIntent", tui.ResumeIntent{}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			root := t.TempDir()
			statePath := filepath.Join(root, ".pulsar-state.json")

			// Persist a legacy (planless) state with a recorded failure, so a
			// naive nil-check-free implementation would find "something to
			// retry" and only the Plan-nil gate prevents it.
			st := engine.State{Provider: "test", Mode: "anonymous"}
			st.Results = []engine.PersistResult{{Test: "TestAccFoo", Status: "fail"}}
			if err := st.Save(statePath); err != nil {
				t.Fatalf("seeding state: %v", err)
			}

			fr := &fakeRunner{}
			pf := &fakeprovider.Fake{NameVal: "test", Packages: []string{"./..."}}
			var msgs []tea.Msg
			send := func(m tea.Msg) { msgs = append(msgs, m) }

			db := &dashboard{
				prov:      pf,
				root:      root,
				mode:      "anonymous",
				statePath: statePath,
				newRunner: func(_ io.Writer) testRunner { return fr },
				list:      stubList([]string{"TestAccFoo"}),
				red:       redact.New(nil),
				groups:    []engine.Group{{Name: "tests", Tests: []string{"TestAccFoo"}}},
				allNames:  []string{"TestAccFoo"},
				send:      send,
				logSink:   io.Discard,
			}

			db.runIntent(context.Background(), tc.msg)

			if fr.calls != 0 {
				t.Errorf("runner must not be invoked, got %d calls", fr.calls)
			}
			var gotErr bool
			for _, m := range msgs {
				if rs, ok := m.(tui.RunStartedMsg); ok {
					t.Errorf("RunStartedMsg must not be sent, got Total=%d", rs.Total)
				}
				if em, ok := m.(tui.ErrMsg); ok {
					gotErr = true
					if em.Err == nil || !strings.Contains(em.Err.Error(), "state/plan-required") {
						t.Errorf("ErrMsg = %v, want it to contain %q", em.Err, "state/plan-required")
					}
				}
			}
			if !gotErr {
				t.Error("expected ErrMsg(state/plan-required)")
			}
		})
	}
}

// TestRunIntentTotalCountsEligibleTestsOnly verifies that RunStartedMsg.Total
// counts only the eligible selected tests (the persisted plan's Eligible
// intersected with the group), not every test physically in the group.
func TestRunIntentTotalCountsEligibleTestsOnly(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")

	st := engine.State{Provider: "test", Mode: "anonymous"}
	// Group "trio" has 3 members; only 2 are eligible under the persisted plan.
	st.Plan = planWithEligible("anonymous", []string{"TestAccOne", "TestAccTwo"})
	if err := st.Save(statePath); err != nil {
		t.Fatalf("seeding state: %v", err)
	}

	fr := &fakeRunner{}
	pf := &fakeprovider.Fake{NameVal: "test", Packages: []string{"./..."}}

	var msgs []tea.Msg
	send := func(m tea.Msg) { msgs = append(msgs, m) }

	db := &dashboard{
		prov:      pf,
		root:      root,
		mode:      "anonymous",
		statePath: statePath,
		newRunner: func(_ io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccOne", "TestAccTwo", "TestAccThree"}),
		red:       redact.New(nil),
		groups:    []engine.Group{{Name: "trio", Tests: []string{"TestAccOne", "TestAccTwo", "TestAccThree"}}},
		allNames:  []string{"TestAccOne", "TestAccTwo", "TestAccThree"},
		send:      send,
		logSink:   io.Discard,
	}

	db.runIntent(context.Background(), tui.RetryGroupIntent{Group: "trio"})

	var gotTotal int
	var gotStart bool
	for _, m := range msgs {
		if rs, ok := m.(tui.RunStartedMsg); ok {
			gotStart = true
			gotTotal = rs.Total
		}
	}
	if !gotStart {
		t.Fatal("expected RunStartedMsg")
	}
	if gotTotal != 2 {
		t.Errorf("RunStartedMsg.Total = %d, want 2 (eligible-only)", gotTotal)
	}
}

// ── 15. Task 10 fix: preserve dashboard plan mode ─────────────────────────────
//
// Critical finding: RetryAllIntent/RetryGroupIntent/ResumeIntent selected
// tests from prev.Plan.Eligible (the persisted plan) but ran under the
// dashboard's CURRENT db.mode while also carrying prev.Plan forward
// unchanged, so a SwitchModeIntent between the last save and a later
// retry/resume produced State.Mode != State.Plan.Mode and executed the
// wrong-mode ExtraEnv. The fix (modeForIntent) mirrors CLI's
// resolveRetryResumeMode: an implicit/omitted mode for a persisted-plan
// intent always derives from prev.Plan.Mode, never the dashboard's current
// mode. RetryTestIntent (no CLI equivalent; bound to whatever single test
// row the dashboard currently shows) keeps using the current mode, and
// runIntent's curPlan fallback is hardened so it never carries forward a
// persisted plan whose Mode disagrees with the mode actually used.
//
// Important finding: concurrent SwitchModeIntent dispatches can complete out
// of order. beginSwitchMode/applySwitchMode add a monotonic sequence
// (cfgSeq, guarded by cfgMu) assigned synchronously at dispatch time
// (mirroring runDashboard's exec, which itself runs synchronously inside
// tui.Model.Update — see tui/update.go's m.exec call), so a stale
// (earlier-dispatched) completion is detected after planning and never
// mutates dashboard state or emits PreflightMsg/GroupsMsg once a later
// dispatch has already applied.

// TestModeForIntent verifies modeForIntent in isolation: persisted-plan
// intents (RetryAllIntent, RetryGroupIntent, ResumeIntent) always resolve to
// prev.Plan.Mode, matching CLI's resolveRetryResumeMode omitted-mode
// convention; every other intent (RetryTestIntent, and any other message)
// keeps the dashboard's current mode unchanged.
func TestModeForIntent(t *testing.T) {
	prev := engine.State{Plan: planWithEligible("anonymous", []string{"TestAccFoo"})}

	cases := []struct {
		name string
		msg  tea.Msg
		want string
	}{
		{"RetryAllIntent uses persisted plan mode", tui.RetryAllIntent{}, "anonymous"},
		{"RetryGroupIntent uses persisted plan mode", tui.RetryGroupIntent{Group: "tests"}, "anonymous"},
		{"ResumeIntent uses persisted plan mode", tui.ResumeIntent{}, "anonymous"},
		{"RetryTestIntent keeps current mode", tui.RetryTestIntent{Test: "TestAccFoo"}, "enterprise"},
		{"unknown intent keeps current mode", struct{}{}, "enterprise"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := modeForIntent(tc.msg, "enterprise", prev)
			if got != tc.want {
				t.Errorf("modeForIntent(%#v, %q, prev) = %q, want %q", tc.msg, "enterprise", got, tc.want)
			}
		})
	}

	t.Run("persisted-plan intent with nil prev.Plan defensively keeps current mode", func(t *testing.T) {
		// runIntent never reaches modeForIntent with a nil prev.Plan for
		// these intents (testsForIntent already returns planRequiredError
		// and runIntent returns before computing mode), but modeForIntent
		// must not panic if ever called that way.
		got := modeForIntent(tui.RetryAllIntent{}, "enterprise", engine.State{})
		if got != "enterprise" {
			t.Errorf("modeForIntent with nil prev.Plan = %q, want unchanged current mode %q", got, "enterprise")
		}
	})
}

// TestBeginSwitchModeIncrementsMonotonically verifies that beginSwitchMode
// returns a strictly increasing sequence on every call, matching dispatch
// order (it is called synchronously at dispatch time, never inside the
// async planning goroutine).
func TestBeginSwitchModeIncrementsMonotonically(t *testing.T) {
	db := &dashboard{}
	first := db.beginSwitchMode()
	second := db.beginSwitchMode()
	third := db.beginSwitchMode()
	if !(first < second && second < third) {
		t.Errorf("beginSwitchMode sequence = %d, %d, %d; want strictly increasing", first, second, third)
	}
}

// TestRunIntentRetryAllRetryGroupResumeUseModeFromPersistedPlan is the
// end-to-end regression test for the Critical finding: it switches the
// dashboard to a different current mode than the persisted plan's mode,
// then runs RetryAllIntent / RetryGroupIntent / ResumeIntent from that
// persisted plan, and asserts the runner's ExtraEnv and the saved
// State.Mode match State.Plan.Mode (the persisted plan's mode) — and that
// the incompatible, wrong (current-mode) env is never used.
func TestRunIntentRetryAllRetryGroupResumeUseModeFromPersistedPlan(t *testing.T) {
	cases := []struct {
		name string
		msg  tea.Msg
	}{
		{"RetryAllIntent", tui.RetryAllIntent{}},
		{"RetryGroupIntent", tui.RetryGroupIntent{Group: "tests"}},
		{"ResumeIntent", tui.ResumeIntent{}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			root := t.TempDir()
			statePath := filepath.Join(root, ".pulsar-state.json")

			// The persisted plan/state are from a run under "anonymous".
			persistedPlan := planWithEligible("anonymous", []string{"TestAccFoo"})
			st := engine.State{Provider: "test", Mode: "anonymous", Plan: persistedPlan}
			st.Results = []engine.PersistResult{{Test: "TestAccFoo", Status: "fail"}}
			if err := st.Save(statePath); err != nil {
				t.Fatalf("seeding state: %v", err)
			}

			fr := &fakeRunner{}
			pf := &fakeprovider.Fake{
				NameVal:  "test",
				Packages: []string{"./..."},
				EnvByMode: map[string][]provider.EnvVar{
					// enterprise-only: must never appear in ExtraEnv when the
					// run actually executes under the persisted anonymous mode.
					"enterprise": {{Key: "GITHUB_ENTERPRISE_SLUG"}},
				},
			}
			getenv := func(k string) string {
				if k == "GITHUB_ENTERPRISE_SLUG" {
					return "wrong-mode-value"
				}
				return ""
			}
			send := func(tea.Msg) {}

			// db.mode/plan/groups/allNames reflect a SwitchModeIntent to
			// "enterprise" having already happened since the persisted save.
			db := &dashboard{
				prov:      pf,
				root:      root,
				mode:      "enterprise",
				getenv:    getenv,
				statePath: statePath,
				newRunner: func(_ io.Writer) testRunner { return fr },
				list:      stubList([]string{"TestAccFoo"}),
				red:       redact.New(nil),
				plan:      planWithEligible("enterprise", []string{"TestAccFoo"}),
				groups:    []engine.Group{{Name: "tests", Tests: []string{"TestAccFoo"}}},
				allNames:  []string{"TestAccFoo"},
				send:      send,
				logSink:   io.Discard,
			}

			db.runIntent(context.Background(), tc.msg)

			if fr.calls != 1 {
				t.Fatalf("expected the runner to be invoked once, got %d calls", fr.calls)
			}

			var gotAnonymous, gotEnterprise, gotWrongModeSecret bool
			for _, e := range fr.spec.ExtraEnv {
				switch e {
				case "GH_TEST_AUTH_MODE=anonymous":
					gotAnonymous = true
				case "GH_TEST_AUTH_MODE=enterprise":
					gotEnterprise = true
				}
				if strings.HasPrefix(e, "GITHUB_ENTERPRISE_SLUG=") {
					gotWrongModeSecret = true
				}
			}
			if !gotAnonymous {
				t.Errorf("ExtraEnv = %v, want GH_TEST_AUTH_MODE=anonymous (the persisted plan's mode)", fr.spec.ExtraEnv)
			}
			if gotEnterprise {
				t.Errorf("ExtraEnv = %v, must not contain the dashboard's stale current mode enterprise", fr.spec.ExtraEnv)
			}
			if gotWrongModeSecret {
				t.Errorf("ExtraEnv = %v, must never leak enterprise-only env while running under anonymous mode", fr.spec.ExtraEnv)
			}

			saved, err := engine.Load(statePath)
			if err != nil {
				t.Fatalf("loading saved state: %v", err)
			}
			if saved.Mode != "anonymous" {
				t.Errorf("saved.Mode = %q, want %q (the persisted plan's mode)", saved.Mode, "anonymous")
			}
			if saved.Plan == nil || saved.Plan.Mode != "anonymous" {
				t.Fatalf("saved.Plan.Mode = %+v, want anonymous", saved.Plan)
			}
			if saved.Mode != saved.Plan.Mode {
				t.Errorf("saved.Mode (%q) != saved.Plan.Mode (%q): dashboard must never persist a mismatched mode", saved.Mode, saved.Plan.Mode)
			}
		})
	}
}

// TestRunIntentRetryTestKeepsCurrentModeConsistentWithStalePersistedPlan
// verifies that RetryTestIntent — which always runs under the dashboard's
// CURRENT mode, never a persisted one — never carries a stale persisted
// Plan from a different, earlier mode forward. Carrying it forward would
// save State.Plan.Mode != State.Mode: the same class of bug the Critical
// Task 10 finding covers for RetryAllIntent/RetryGroupIntent/ResumeIntent.
func TestRunIntentRetryTestKeepsCurrentModeConsistentWithStalePersistedPlan(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")

	// Persisted plan/state are stale: from a run under "anonymous" mode.
	stalePlan := planWithEligible("anonymous", []string{"TestAccEligible"})
	st := engine.State{Provider: "test", Mode: "anonymous", Plan: stalePlan}
	if err := st.Save(statePath); err != nil {
		t.Fatalf("seeding state: %v", err)
	}

	fr := &fakeRunner{}
	pf := &fakeprovider.Fake{NameVal: "test", Packages: []string{"./..."}}
	send := func(tea.Msg) {}

	// db.mode/plan reflect a SwitchModeIntent to "enterprise" since the save.
	currentPlan := planWithEligible("enterprise", []string{"TestAccEligible"})
	db := &dashboard{
		prov:      pf,
		root:      root,
		mode:      "enterprise",
		statePath: statePath,
		newRunner: func(_ io.Writer) testRunner { return fr },
		list:      stubList([]string{"TestAccEligible"}),
		red:       redact.New(nil),
		plan:      currentPlan,
		groups:    []engine.Group{{Name: "tests", Tests: []string{"TestAccEligible"}}},
		allNames:  []string{"TestAccEligible"},
		send:      send,
		logSink:   io.Discard,
	}

	db.runIntent(context.Background(), tui.RetryTestIntent{Test: "TestAccEligible"})

	saved, err := engine.Load(statePath)
	if err != nil {
		t.Fatalf("loading saved state: %v", err)
	}
	if saved.Mode != "enterprise" {
		t.Errorf("saved.Mode = %q, want %q (RetryTestIntent always uses the current mode)", saved.Mode, "enterprise")
	}
	if saved.Plan == nil || saved.Plan.Mode != "enterprise" {
		t.Fatalf("saved.Plan.Mode = %+v, want enterprise: must not carry the stale anonymous-mode persisted plan forward", saved.Plan)
	}
	if saved.Mode != saved.Plan.Mode {
		t.Errorf("saved.Mode (%q) != saved.Plan.Mode (%q)", saved.Mode, saved.Plan.Mode)
	}
}

// gatedList returns a lister whose FIRST call blocks until proceed is
// closed, returning namesA; every subsequent call returns immediately with
// namesB. started is closed right before the first call blocks, letting a
// test deterministically wait for that call to be in-flight before
// completing a second, faster switch — without relying on time.Sleep.
func gatedList(namesA, namesB []string, started, proceed chan struct{}) lister {
	var mu sync.Mutex
	calls := 0
	return func(_ context.Context, _ string, _ []string, _ string) ([]string, error) {
		mu.Lock()
		calls++
		n := calls
		mu.Unlock()
		if n == 1 {
			close(started)
			<-proceed
			return namesA, nil
		}
		return namesB, nil
	}
}

// TestApplySwitchModeLatestDispatchWinsOverStaleSlowCompletion verifies the
// Important Task 10 review finding: two concurrent SwitchModeIntent
// dispatches can finish out of order (a slower, EARLIER-dispatched switch
// completing AFTER a faster, LATER-dispatched one). The stale (earlier)
// completion must never mutate db.mode/plan/groups/allNames, nor emit
// PreflightMsg/GroupsMsg, once a later dispatch has already applied — the
// latest dispatch, by sequence (assigned at dispatch time via
// beginSwitchMode), always wins regardless of completion order.
func TestApplySwitchModeLatestDispatchWinsOverStaleSlowCompletion(t *testing.T) {
	root := t.TempDir()

	started := make(chan struct{})
	proceed := make(chan struct{})

	pf := &fakeprovider.Fake{
		NameVal:  "test",
		Packages: []string{"./..."},
		ModesVal: []provider.Mode{{Name: "modeA"}, {Name: "modeB"}},
		PreflightFn: func(_ context.Context, m string, _ provider.TestRequirements) provider.PreflightReport {
			return provider.PreflightReport{Mode: m}
		},
		RequirementsFn: allowAllRequirements,
	}

	var mu sync.Mutex
	var msgs []tea.Msg
	send := func(m tea.Msg) {
		mu.Lock()
		msgs = append(msgs, m)
		mu.Unlock()
	}

	db := &dashboard{
		prov:      pf,
		root:      root,
		mode:      "anonymous",
		statePath: filepath.Join(root, ".pulsar-state.json"),
		newRunner: func(_ io.Writer) testRunner { return &fakeRunner{} },
		list:      gatedList([]string{"TestAccModeA"}, []string{"TestAccModeB"}, started, proceed),
		red:       redact.New(nil),
		send:      send,
		logSink:   io.Discard,
	}

	// Dispatch A first (lowest sequence): its planning call blocks.
	seqA := db.beginSwitchMode()
	doneA := make(chan struct{})
	go func() {
		defer close(doneA)
		db.applySwitchMode(context.Background(), tui.SwitchModeIntent{Mode: "modeA"}, seqA)
	}()
	<-started // A's planning call is confirmed blocked mid-flight.

	// Dispatch B second (higher sequence): the gated list only blocks on its
	// very FIRST call, so B's planning call returns immediately and this
	// completes synchronously, right here, before A is ever unblocked.
	seqB := db.beginSwitchMode()
	db.applySwitchMode(context.Background(), tui.SwitchModeIntent{Mode: "modeB"}, seqB)

	if db.mode != "modeB" || db.plan == nil || db.plan.Mode != "modeB" {
		t.Fatalf("after B applies: db.mode=%q db.plan=%+v, want modeB", db.mode, db.plan)
	}

	// Now allow A's (already-stale) blocked completion to finish.
	close(proceed)
	<-doneA

	if db.mode != "modeB" {
		t.Errorf("db.mode = %q after the stale A completion, want unchanged %q", db.mode, "modeB")
	}
	if db.plan == nil || db.plan.Mode != "modeB" {
		t.Fatalf("db.plan = %+v after the stale A completion, want unchanged Mode=modeB", db.plan)
	}
	if !reflect.DeepEqual(db.allNames, []string{"TestAccModeB"}) {
		t.Errorf("db.allNames = %v after the stale A completion, want unchanged [TestAccModeB]", db.allNames)
	}
	if len(db.groups) != 1 || !reflect.DeepEqual(db.groups[0].Tests, []string{"TestAccModeB"}) {
		t.Errorf("db.groups = %+v after the stale A completion, want unchanged single group with [TestAccModeB]", db.groups)
	}

	mu.Lock()
	defer mu.Unlock()
	var preflightModes, groupsModes []string
	for _, m := range msgs {
		switch pm := m.(type) {
		case tui.PreflightMsg:
			preflightModes = append(preflightModes, pm.Report.Mode)
		case tui.GroupsMsg:
			if pm.Plan != nil {
				groupsModes = append(groupsModes, pm.Plan.Mode)
			}
		}
	}
	if !reflect.DeepEqual(preflightModes, []string{"modeB"}) {
		t.Errorf("PreflightMsg modes = %v, want exactly [modeB]: the stale A completion must emit none", preflightModes)
	}
	if !reflect.DeepEqual(groupsModes, []string{"modeB"}) {
		t.Errorf("GroupsMsg plan modes = %v, want exactly [modeB]: the stale A completion must emit none", groupsModes)
	}
}

// TestApplySwitchModeStaleFailureEmitsNoErrMsgOrMutation verifies that a
// STALE SwitchModeIntent completion whose planning FAILS emits no ErrMsg and
// mutates no dashboard state — exactly like a stale SUCCESSFUL completion
// emits nothing — so an operator is never shown a confusing error banner
// about a mode switch they have already superseded. Uses direct sequential
// calls (deterministic; no goroutines needed) to simulate B arriving before
// stale A.
func TestApplySwitchModeStaleFailureEmitsNoErrMsgOrMutation(t *testing.T) {
	root := t.TempDir()

	pf := &fakeprovider.Fake{
		NameVal:  "test",
		Packages: []string{"./..."},
		ModesVal: []provider.Mode{{Name: "modeA"}, {Name: "modeB"}},
		PreflightFn: func(_ context.Context, m string, _ provider.TestRequirements) provider.PreflightReport {
			return provider.PreflightReport{Mode: m}
		},
		RequirementsFn: allowAllRequirements,
	}

	var msgs []tea.Msg
	send := func(m tea.Msg) { msgs = append(msgs, m) }

	callCount := 0
	list := func(_ context.Context, _ string, _ []string, _ string) ([]string, error) {
		callCount++
		if callCount == 1 {
			return []string{"TestAccFoo"}, nil // modeB's (latest) planning call
		}
		return nil, errors.New("boom") // modeA's (stale) planning call
	}

	db := &dashboard{
		prov:      pf,
		root:      root,
		mode:      "anonymous",
		statePath: filepath.Join(root, ".pulsar-state.json"),
		newRunner: func(_ io.Writer) testRunner { return &fakeRunner{} },
		list:      list,
		red:       redact.New(nil),
		send:      send,
		logSink:   io.Discard,
	}

	seqA := db.beginSwitchMode() // dispatched first, will complete LAST (stale)
	seqB := db.beginSwitchMode() // dispatched second (latest)

	// modeB "arrives" (completes) first: applies normally.
	db.applySwitchMode(context.Background(), tui.SwitchModeIntent{Mode: "modeB"}, seqB)
	if db.mode != "modeB" {
		t.Fatalf("db.mode = %q, want modeB once the latest dispatch applies", db.mode)
	}
	msgs = nil // only inspect what A's stale completion below produces.

	// modeA "arrives" (completes) second, now stale; its planning also fails.
	db.applySwitchMode(context.Background(), tui.SwitchModeIntent{Mode: "modeA"}, seqA)

	if len(msgs) != 0 {
		t.Errorf("stale completion must emit nothing (no ErrMsg), got %#v", msgs)
	}
	if db.mode != "modeB" {
		t.Errorf("db.mode = %q, want unchanged modeB after a stale completion", db.mode)
	}
}

// unknownDashboardProvider returns a provider that reports every name in
// unknown as unclassified, with the conservative authenticated requirements
// the real catalog uses as its fail-closed fallback.
func unknownDashboardProvider(unknown map[string]bool) *fakeprovider.Fake {
	return &fakeprovider.Fake{
		NameVal:  "test",
		Packages: []string{"./..."},
		ModesVal: []provider.Mode{{Name: "anonymous"}, {Name: "individual"}, {Name: "organization"}},
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
		PreflightFn: func(_ context.Context, m string, _ provider.TestRequirements) provider.PreflightReport {
			return provider.PreflightReport{Mode: m}
		},
	}
}

// TestApplySwitchModeFailsClosedOnUnclassified is the fail-closed regression
// for the dashboard, which hardcoded AllowUnclassified: true and therefore
// silently accepted unknown tests with over-broad conservative requirements
// while every non-TUI surface fails such a plan closed. Without the operator
// opting in, a mode switch that discovers an unknown test must report the
// planning error and leave dashboard state untouched.
func TestApplySwitchModeFailsClosedOnUnclassified(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")

	pf := unknownDashboardProvider(map[string]bool{"TestAccMystery": true})
	pf.PreflightFn = func(context.Context, string, provider.TestRequirements) provider.PreflightReport {
		t.Fatal("Preflight must not be called when planning fails closed")
		return provider.PreflightReport{}
	}

	origPlan := planWithEligible("anonymous", []string{"TestAccOpen"})
	var msgs []tea.Msg
	db := &dashboard{
		prov:      pf,
		root:      root,
		mode:      "anonymous",
		statePath: statePath,
		newRunner: func(_ io.Writer) testRunner { return &fakeRunner{} },
		list:      stubList([]string{"TestAccMystery"}),
		red:       redact.New(nil),
		plan:      origPlan,
		groups:    []engine.Group{{Name: "tests", Tests: []string{"TestAccOpen"}}},
		allNames:  []string{"TestAccOpen"},
		send:      func(m tea.Msg) { msgs = append(msgs, m) },
		logSink:   io.Discard,
	}

	db.applySwitchMode(context.Background(), tui.SwitchModeIntent{Mode: "organization"}, db.beginSwitchMode())

	if db.mode != "anonymous" {
		t.Errorf("db.mode = %q, want unchanged anonymous after a fail-closed plan", db.mode)
	}
	if db.plan != origPlan {
		t.Errorf("db.plan changed on a fail-closed plan: %+v", db.plan)
	}
	var gotErr bool
	for _, m := range msgs {
		switch msg := m.(type) {
		case tui.ErrMsg:
			gotErr = true
			if !strings.Contains(msg.Err.Error(), "has no requirements classification") {
				t.Errorf("ErrMsg = %v, want the fail-closed unclassified planning error", msg.Err)
			}
		case tui.GroupsMsg:
			t.Error("GroupsMsg must not be sent when planning fails closed")
		}
	}
	if !gotErr {
		t.Error("expected ErrMsg for an unclassified test without the opt-in")
	}
}

// TestPlanDashboardAllowUnclassifiedOptIn verifies the dashboard honors the
// operator's explicit --allow-unclassified opt-in exactly like the non-TUI
// planning path: fail closed by default, and plan the unknown test with its
// conservative requirements when allowed.
func TestPlanDashboardAllowUnclassifiedOptIn(t *testing.T) {
	pf := unknownDashboardProvider(map[string]bool{"TestAccMystery": true})
	names := []string{"TestAccMystery"}

	if _, _, _, err := planDashboard(context.Background(), stubList(names), pf, t.TempDir(), "organization", false, redact.New(nil)); err == nil {
		t.Fatal("expected planDashboard to fail closed on an unclassified test")
	}

	_, allNames, plan, err := planDashboard(context.Background(), stubList(names), pf, t.TempDir(), "organization", true, redact.New(nil))
	if err != nil {
		t.Fatalf("planDashboard with the opt-in: %v", err)
	}
	if !reflect.DeepEqual(allNames, []string{"TestAccMystery"}) {
		t.Errorf("allNames = %v, want the allowed unknown test", allNames)
	}
	if !reflect.DeepEqual(plan.Unclassified, []string{"TestAccMystery"}) {
		t.Errorf("plan.Unclassified = %v, want [TestAccMystery]", plan.Unclassified)
	}
}

// TestApplySwitchModeAllowUnclassifiedCarriesOptIn verifies the opt-in the
// dashboard was launched with is carried into mode-switch planning, so a
// switch does not fail closed on tests the initial plan already accepted.
func TestApplySwitchModeAllowUnclassifiedCarriesOptIn(t *testing.T) {
	root := t.TempDir()
	pf := unknownDashboardProvider(map[string]bool{"TestAccMystery": true})

	var msgs []tea.Msg
	db := &dashboard{
		prov:              pf,
		root:              root,
		mode:              "individual",
		statePath:         filepath.Join(root, ".pulsar-state.json"),
		newRunner:         func(_ io.Writer) testRunner { return &fakeRunner{} },
		list:              stubList([]string{"TestAccMystery"}),
		red:               redact.New(nil),
		plan:              planWithEligible("individual", []string{"TestAccMystery"}),
		groups:            []engine.Group{{Name: "tests", Tests: []string{"TestAccMystery"}}},
		allNames:          []string{"TestAccMystery"},
		allowUnclassified: true,
		send:              func(m tea.Msg) { msgs = append(msgs, m) },
		logSink:           io.Discard,
	}

	db.applySwitchMode(context.Background(), tui.SwitchModeIntent{Mode: "organization"}, db.beginSwitchMode())

	if db.mode != "organization" {
		t.Errorf("db.mode = %q, want organization", db.mode)
	}
	if db.plan == nil || !reflect.DeepEqual(db.plan.Unclassified, []string{"TestAccMystery"}) {
		t.Fatalf("db.plan = %+v, want the allowed unknown test carried into the new plan", db.plan)
	}
	for _, m := range msgs {
		if errMsg, ok := m.(tui.ErrMsg); ok {
			t.Errorf("unexpected ErrMsg with the opt-in carried forward: %v", errMsg.Err)
		}
	}
}

// TestSeedDoesNotClobberCompletedModeSwitch is the deterministic regression
// for seed's unsynchronized config reads and, with them, seed's ability to
// overwrite a newer mode switch.
//
// seed read db.mode/db.plan for Preflight and then read db.groups/db.allNames
// again afterwards, so a SwitchModeIntent landing in between made seed emit a
// PreflightMsg for one mode and a GroupsMsg for another - an inconsistent
// opening view, and an unguarded read of fields applySwitchMode writes under
// cfgMu. Worse, seed's launch-mode pair lands AFTER the switch's pair, and
// tui.Model replaces mode/checks/groups wholesale on each message, so the
// operator ends up looking at the launch mode while db.mode - and therefore
// every run the dashboard starts - uses the switched mode.
//
// seed therefore takes ONE snapshot under cfgMu (including the cfgSeq the
// switch dispatch sequence uses) and, exactly like a superseded
// applySwitchMode, sends no mode-tied message once a newer switch exists.
// The provider's Preflight blocks until the test has completed a full mode
// switch, so the interleaving is deterministic rather than timing-dependent.
func TestSeedDoesNotClobberCompletedModeSwitch(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")

	st := engine.State{
		Provider: "test",
		Mode:     "anonymous",
		RunAt:    time.Now().Add(-time.Hour),
		Results:  []engine.PersistResult{{Test: "TestAccOpen", Status: "fail"}},
	}
	if err := st.Save(statePath); err != nil {
		t.Fatalf("seeding state: %v", err)
	}

	seedStarted := make(chan struct{})
	switchDone := make(chan struct{})
	var preflightModes []string
	var preflightMu sync.Mutex

	pf := &fakeprovider.Fake{
		NameVal:        "test",
		Packages:       []string{"./..."},
		ModesVal:       []provider.Mode{{Name: "anonymous"}, {Name: "organization"}},
		RequirementsFn: allowAllRequirements,
		PreflightFn: func(_ context.Context, m string, _ provider.TestRequirements) provider.PreflightReport {
			preflightMu.Lock()
			preflightModes = append(preflightModes, m)
			first := len(preflightModes) == 1
			preflightMu.Unlock()
			if first {
				// This is seed's Preflight: let the mode switch complete
				// entirely before seed continues.
				close(seedStarted)
				<-switchDone
			}
			return provider.PreflightReport{Mode: m}
		},
	}

	var mu sync.Mutex
	var msgs []tea.Msg
	send := func(m tea.Msg) {
		mu.Lock()
		msgs = append(msgs, m)
		mu.Unlock()
	}

	db := &dashboard{
		prov:      pf,
		root:      root,
		mode:      "anonymous",
		statePath: statePath,
		newRunner: func(_ io.Writer) testRunner { return &fakeRunner{} },
		list:      stubList([]string{"TestAccOpen"}),
		red:       redact.New(nil),
		plan:      planWithEligible("anonymous", []string{"TestAccOpen"}),
		groups:    []engine.Group{{Name: "anonymous-tests", Tests: []string{"TestAccOpen"}}},
		allNames:  []string{"TestAccOpen"},
		send:      send,
		logSink:   io.Discard,
	}

	seeded := make(chan struct{})
	go func() {
		defer close(seeded)
		db.seed(context.Background())
	}()

	<-seedStarted
	db.applySwitchMode(context.Background(), tui.SwitchModeIntent{Mode: "organization"}, db.beginSwitchMode())
	close(switchDone)
	<-seeded

	preflightMu.Lock()
	gotModes := append([]string(nil), preflightModes...)
	preflightMu.Unlock()
	if len(gotModes) != 2 || gotModes[0] != "anonymous" || gotModes[1] != "organization" {
		t.Fatalf("Preflight modes = %v, want seed's snapshot mode then the switch's", gotModes)
	}

	mu.Lock()
	defer mu.Unlock()
	var lastPreflight *tui.PreflightMsg
	var lastGroups *tui.GroupsMsg
	var updates int
	var gotResume bool
	for _, m := range msgs {
		switch typed := m.(type) {
		case tui.PreflightMsg:
			cp := typed
			lastPreflight = &cp
		case tui.GroupsMsg:
			cp := typed
			lastGroups = &cp
		case tui.TestUpdateMsg:
			updates++
		case tui.ResumePromptMsg:
			gotResume = true
		}
	}
	if lastPreflight == nil || lastGroups == nil {
		t.Fatalf("expected a PreflightMsg and a GroupsMsg, got %#v", msgs)
	}
	// The switch is newer, so its view must be the one left standing.
	if lastPreflight.Report.Mode != "organization" {
		t.Errorf("final PreflightMsg mode = %q, want organization (seed must not clobber a newer switch)", lastPreflight.Report.Mode)
	}
	if lastGroups.Plan == nil || lastGroups.Plan.Mode != "organization" {
		t.Errorf("final GroupsMsg plan = %+v, want the switched-to organization plan", lastGroups.Plan)
	}
	// The persisted-result restore is mode-independent and must still arrive.
	if updates != 1 {
		t.Errorf("TestUpdateMsg count = %d, want 1 restored persisted result", updates)
	}
	if !gotResume {
		t.Error("expected the resume prompt for the prior failed run")
	}
}

// TestSeedEmitsSnapshotPairWithoutModeSwitch pins the ordinary path: with no
// concurrent mode switch, seed emits its PreflightMsg and GroupsMsg from one
// consistent snapshot of mode/plan/groups.
func TestSeedEmitsSnapshotPairWithoutModeSwitch(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")

	pf := &fakeprovider.Fake{
		NameVal:        "test",
		Packages:       []string{"./..."},
		ModesVal:       []provider.Mode{{Name: "anonymous"}, {Name: "organization"}},
		RequirementsFn: allowAllRequirements,
		PreflightFn: func(_ context.Context, m string, _ provider.TestRequirements) provider.PreflightReport {
			return provider.PreflightReport{Mode: m}
		},
	}

	var msgs []tea.Msg
	db := &dashboard{
		prov:      pf,
		root:      root,
		mode:      "anonymous",
		statePath: statePath,
		newRunner: func(_ io.Writer) testRunner { return &fakeRunner{} },
		list:      stubList([]string{"TestAccOpen"}),
		red:       redact.New(nil),
		plan:      planWithEligible("anonymous", []string{"TestAccOpen"}),
		groups:    []engine.Group{{Name: "anonymous-tests", Tests: []string{"TestAccOpen"}}},
		allNames:  []string{"TestAccOpen"},
		send:      func(m tea.Msg) { msgs = append(msgs, m) },
		logSink:   io.Discard,
	}

	db.seed(context.Background())

	var gotPreflight, gotGroups bool
	for _, m := range msgs {
		switch typed := m.(type) {
		case tui.PreflightMsg:
			gotPreflight = true
			if typed.Report.Mode != "anonymous" {
				t.Errorf("PreflightMsg mode = %q, want anonymous", typed.Report.Mode)
			}
		case tui.GroupsMsg:
			gotGroups = true
			if typed.Plan == nil || typed.Plan.Mode != "anonymous" {
				t.Errorf("GroupsMsg plan = %+v, want the anonymous snapshot", typed.Plan)
			}
			if len(typed.Groups) != 1 || typed.Groups[0].Name != "anonymous-tests" {
				t.Errorf("GroupsMsg groups = %+v, want the snapshot groups", typed.Groups)
			}
		}
	}
	if !gotPreflight || !gotGroups {
		t.Fatalf("expected both PreflightMsg and GroupsMsg, got %#v", msgs)
	}
}

// TestDashboardPlanRedactionKeepsSecretsOutOfPersistedState is the
// defense-in-depth regression for dashboard plans, which bypassed the shared
// redactPlan step every non-TUI surface applies before emitting or persisting
// a plan. Exclusion detail/fix is the only interpolated text on an
// ExecutionPlan and therefore the only field redaction governs, on every
// surface: Selected/Eligible/Excluded[].Test and RequiredModes carry
// provider-discovered `go test -list` names and catalog mode vocabulary, which
// are static source identifiers rather than operator-configured values. The
// test drives a configured secret marker through the interpolated detail/fix
// path and asserts the RAW persisted state carries the redaction marker and no
// marker in any exclusion detail or fix - exactly the invariant run,
// preflight, and e2e already hold.
func TestDashboardPlanRedactionKeepsSecretsOutOfPersistedState(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")

	const secret = "ghp_DASHBOARDPLANSECRET"
	leaky := "TestAccLeak" + secret
	pf := &fakeprovider.Fake{
		NameVal:  "test",
		Packages: []string{"./..."},
		ModesVal: []provider.Mode{{Name: "anonymous"}, {Name: "organization"}},
		Secrets:  []string{"GITHUB_TOKEN"},
		RequirementsFn: func(name string) (provider.TestRequirements, bool) {
			if name == leaky {
				return provider.TestRequirements{Modes: []string{"organization"}}, true
			}
			return provider.TestRequirements{}, true
		},
		PreflightFn: func(_ context.Context, m string, _ provider.TestRequirements) provider.PreflightReport {
			return provider.PreflightReport{Mode: m}
		},
	}

	red := redact.New([]string{secret})
	groups, allNames, plan, err := planDashboard(
		context.Background(), stubList([]string{"TestAccOpen", leaky}), pf, root, "anonymous", false, red)
	if err != nil {
		t.Fatalf("planDashboard: %v", err)
	}
	if len(plan.Excluded) != 1 {
		t.Fatalf("plan.Excluded = %+v, want the mode-incompatible test", plan.Excluded)
	}
	if strings.Contains(plan.Excluded[0].Detail, secret) || strings.Contains(plan.Excluded[0].Fix, secret) {
		t.Fatalf("dashboard plan exclusion still carries the raw secret: %+v", plan.Excluded[0])
	}

	var msgs []tea.Msg
	db := &dashboard{
		prov:      pf,
		root:      root,
		mode:      "anonymous",
		statePath: statePath,
		newRunner: func(_ io.Writer) testRunner {
			return &fakeRunner{result: engine.RunResult{Tests: []engine.TestResult{
				{Package: "./github", Name: "TestAccOpen", Status: provider.StatusPass},
			}}}
		},
		list:     stubList([]string{"TestAccOpen", leaky}),
		getenv:   getenvFromMap(map[string]string{"GITHUB_TOKEN": secret}),
		red:      red,
		plan:     &plan,
		groups:   groups,
		allNames: allNames,
		send:     func(m tea.Msg) { msgs = append(msgs, m) },
		logSink:  io.Discard,
	}

	db.runIntent(context.Background(), tui.RetryTestIntent{Test: "TestAccOpen"})

	raw, err := os.ReadFile(statePath)
	if err != nil {
		t.Fatalf("reading persisted state: %v", err)
	}
	if !strings.Contains(string(raw), "***REDACTED***") {
		t.Fatalf("persisted state was not written from a redacted plan:\n%s", raw)
	}
	var persisted engine.State
	if err := json.Unmarshal(raw, &persisted); err != nil {
		t.Fatalf("decoding persisted state: %v", err)
	}
	if persisted.Plan == nil {
		t.Fatal("persisted state has no plan")
	}
	if len(persisted.Plan.Excluded) != 1 {
		t.Fatalf("persisted plan exclusions = %+v, want the mode-incompatible test", persisted.Plan.Excluded)
	}
	for _, ex := range persisted.Plan.Excluded {
		if strings.Contains(ex.Detail, secret) || strings.Contains(ex.Fix, secret) {
			t.Fatalf("raw persisted plan exclusion carries the configured marker: %+v", ex)
		}
	}

	// The GroupsMsg the TUI renders must carry the same redacted plan.
	for _, m := range msgs {
		gm, ok := m.(tui.GroupsMsg)
		if !ok || gm.Plan == nil {
			continue
		}
		for _, ex := range gm.Plan.Excluded {
			if strings.Contains(ex.Detail, secret) || strings.Contains(ex.Fix, secret) {
				t.Fatalf("GroupsMsg plan exclusion carries the configured marker: %+v", ex)
			}
		}
	}
}

// gatedPreflight returns a PreflightFn whose FIRST call blocks: it closes
// started, waits for proceed, then returns a report for the requested mode.
// Every later call returns immediately. It is the Preflight counterpart of
// gatedList, used to pin a producer inside provider network I/O while a
// later dispatch completes.
func gatedPreflight(started, proceed chan struct{}) func(context.Context, string, provider.TestRequirements) provider.PreflightReport {
	var mu sync.Mutex
	calls := 0
	return func(_ context.Context, mode string, _ provider.TestRequirements) provider.PreflightReport {
		mu.Lock()
		calls++
		n := calls
		mu.Unlock()
		if n == 1 {
			close(started)
			<-proceed
		}
		return provider.PreflightReport{Mode: mode}
	}
}

// modeTiedMessages splits captured messages into the PreflightMsg report modes
// and GroupsMsg plan modes they carry, in send order.
func modeTiedMessages(msgs []tea.Msg) (preflightModes, groupsModes []string) {
	for _, m := range msgs {
		switch pm := m.(type) {
		case tui.PreflightMsg:
			preflightModes = append(preflightModes, pm.Report.Mode)
		case tui.GroupsMsg:
			if pm.Plan != nil {
				groupsModes = append(groupsModes, pm.Plan.Mode)
			}
		}
	}
	return preflightModes, groupsModes
}

// TestApplySwitchModeLatestDispatchWinsOverStalePreflight covers the
// latest-dispatch hole the sequence check alone left open: applySwitchMode
// verified cfgSeq BEFORE calling the provider's Preflight, then sent its
// PreflightMsg/GroupsMsg pair unconditionally once that (slow, networked)
// call returned. An earlier switch pinned inside Preflight therefore
// clobbered a later switch's messages, leaving the operator looking at mode A
// while db.mode — and every run the dashboard starts — already used mode B.
// The freshness check must gate the sends themselves, not just planning.
func TestApplySwitchModeLatestDispatchWinsOverStalePreflight(t *testing.T) {
	root := t.TempDir()

	started := make(chan struct{})
	proceed := make(chan struct{})

	pf := &fakeprovider.Fake{
		NameVal:        "test",
		Packages:       []string{"./..."},
		ModesVal:       []provider.Mode{{Name: "modeA"}, {Name: "modeB"}},
		PreflightFn:    gatedPreflight(started, proceed),
		RequirementsFn: allowAllRequirements,
	}

	var mu sync.Mutex
	var msgs []tea.Msg
	send := func(m tea.Msg) {
		mu.Lock()
		msgs = append(msgs, m)
		mu.Unlock()
	}

	db := &dashboard{
		prov:      pf,
		root:      root,
		mode:      "anonymous",
		statePath: filepath.Join(root, ".pulsar-state.json"),
		newRunner: func(_ io.Writer) testRunner { return &fakeRunner{} },
		list: func(_ context.Context, _ string, _ []string, _ string) ([]string, error) {
			return []string{"TestAccThing"}, nil
		},
		red:     redact.New(nil),
		send:    send,
		logSink: io.Discard,
	}

	// A is dispatched first and blocks inside Preflight (planning is fast).
	seqA := db.beginSwitchMode()
	doneA := make(chan struct{})
	go func() {
		defer close(doneA)
		db.applySwitchMode(context.Background(), tui.SwitchModeIntent{Mode: "modeA"}, seqA)
	}()
	<-started // A is confirmed pinned inside the provider Preflight call.

	// B is dispatched second and completes end-to-end while A is still
	// inside Preflight: its own Preflight call is not the first, so it
	// returns immediately.
	seqB := db.beginSwitchMode()
	db.applySwitchMode(context.Background(), tui.SwitchModeIntent{Mode: "modeB"}, seqB)

	db.cfgMu.Lock()
	gotMode, gotPlan := db.mode, db.plan
	db.cfgMu.Unlock()
	if gotMode != "modeB" || gotPlan == nil || gotPlan.Mode != "modeB" {
		t.Fatalf("after B applies: db.mode=%q db.plan=%+v, want modeB", gotMode, gotPlan)
	}

	// Release A's superseded Preflight.
	close(proceed)
	<-doneA

	db.cfgMu.Lock()
	gotMode, gotPlan = db.mode, db.plan
	db.cfgMu.Unlock()
	if gotMode != "modeB" {
		t.Errorf("db.mode = %q after the stale A completion, want unchanged modeB", gotMode)
	}
	if gotPlan == nil || gotPlan.Mode != "modeB" {
		t.Fatalf("db.plan = %+v after the stale A completion, want unchanged Mode=modeB", gotPlan)
	}

	mu.Lock()
	defer mu.Unlock()
	preflightModes, groupsModes := modeTiedMessages(msgs)
	if !reflect.DeepEqual(preflightModes, []string{"modeB"}) {
		t.Errorf("PreflightMsg modes = %v, want exactly [modeB]: A's superseded preflight must emit none", preflightModes)
	}
	if !reflect.DeepEqual(groupsModes, []string{"modeB"}) {
		t.Errorf("GroupsMsg plan modes = %v, want exactly [modeB]: A's superseded preflight must emit none", groupsModes)
	}
}

// TestHandleConfigIntentPreflightSuppressedByLaterModeSwitch verifies that a
// PreflightIntent which snapshotted the pre-switch mode/plan and then blocked
// inside the provider Preflight call does not report that stale mode after a
// later SwitchModeIntent has already applied and rendered its own pair.
func TestHandleConfigIntentPreflightSuppressedByLaterModeSwitch(t *testing.T) {
	root := t.TempDir()

	started := make(chan struct{})
	proceed := make(chan struct{})

	pf := &fakeprovider.Fake{
		NameVal:        "test",
		Packages:       []string{"./..."},
		ModesVal:       []provider.Mode{{Name: "modeA"}, {Name: "modeB"}},
		PreflightFn:    gatedPreflight(started, proceed),
		RequirementsFn: allowAllRequirements,
	}

	var mu sync.Mutex
	var msgs []tea.Msg
	send := func(m tea.Msg) {
		mu.Lock()
		msgs = append(msgs, m)
		mu.Unlock()
	}

	plan := engine.ExecutionPlan{Mode: "modeA", Eligible: []string{"TestAccThing"}}
	db := &dashboard{
		prov:      pf,
		root:      root,
		mode:      "modeA",
		statePath: filepath.Join(root, ".pulsar-state.json"),
		newRunner: func(_ io.Writer) testRunner { return &fakeRunner{} },
		list: func(_ context.Context, _ string, _ []string, _ string) ([]string, error) {
			return []string{"TestAccThing"}, nil
		},
		red:     redact.New(nil),
		plan:    &plan,
		send:    send,
		logSink: io.Discard,
	}

	doneCfg := make(chan struct{})
	go func() {
		defer close(doneCfg)
		db.handleConfigIntent(context.Background(), tui.PreflightIntent{})
	}()
	<-started // the config intent is pinned inside Preflight with modeA.

	seqB := db.beginSwitchMode()
	db.applySwitchMode(context.Background(), tui.SwitchModeIntent{Mode: "modeB"}, seqB)

	close(proceed)
	<-doneCfg

	db.cfgMu.Lock()
	gotMode := db.mode
	db.cfgMu.Unlock()
	if gotMode != "modeB" {
		t.Errorf("db.mode = %q, want modeB", gotMode)
	}

	mu.Lock()
	defer mu.Unlock()
	preflightModes, _ := modeTiedMessages(msgs)
	if !reflect.DeepEqual(preflightModes, []string{"modeB"}) {
		t.Errorf("PreflightMsg modes = %v, want exactly [modeB]: the stale config-intent report must be suppressed", preflightModes)
	}
}

// TestHandleConfigIntentSetEnvVarSuppressedByLaterModeSwitch is the
// SetEnvVarIntent half of the same contract: the env write still happens, but
// the report it would render for the pre-switch mode must not reach the TUI
// once a later mode switch has applied.
func TestHandleConfigIntentSetEnvVarSuppressedByLaterModeSwitch(t *testing.T) {
	root := t.TempDir()

	started := make(chan struct{})
	proceed := make(chan struct{})

	pf := &fakeprovider.Fake{
		NameVal:        "test",
		Packages:       []string{"./..."},
		ModesVal:       []provider.Mode{{Name: "modeA"}, {Name: "modeB"}},
		PreflightFn:    gatedPreflight(started, proceed),
		RequirementsFn: allowAllRequirements,
	}

	var mu sync.Mutex
	var msgs []tea.Msg
	send := func(m tea.Msg) {
		mu.Lock()
		msgs = append(msgs, m)
		mu.Unlock()
	}

	var envMu sync.Mutex
	env := map[string]string{}
	plan := engine.ExecutionPlan{Mode: "modeA", Eligible: []string{"TestAccThing"}}
	db := &dashboard{
		prov:      pf,
		root:      root,
		mode:      "modeA",
		statePath: filepath.Join(root, ".pulsar-state.json"),
		newRunner: func(_ io.Writer) testRunner { return &fakeRunner{} },
		list: func(_ context.Context, _ string, _ []string, _ string) ([]string, error) {
			return []string{"TestAccThing"}, nil
		},
		red:  redact.New(nil),
		plan: &plan,
		getenv: func(k string) string {
			envMu.Lock()
			defer envMu.Unlock()
			return env[k]
		},
		setenv: func(k, v string) error {
			envMu.Lock()
			defer envMu.Unlock()
			env[k] = v
			return nil
		},
		send:    send,
		logSink: io.Discard,
	}

	doneCfg := make(chan struct{})
	go func() {
		defer close(doneCfg)
		db.handleConfigIntent(context.Background(), tui.SetEnvVarIntent{Key: "GITHUB_OWNER", Value: "acme"})
	}()
	<-started

	seqB := db.beginSwitchMode()
	db.applySwitchMode(context.Background(), tui.SwitchModeIntent{Mode: "modeB"}, seqB)

	close(proceed)
	<-doneCfg

	envMu.Lock()
	gotEnv := env["GITHUB_OWNER"]
	envMu.Unlock()
	if gotEnv != "acme" {
		t.Errorf("GITHUB_OWNER = %q, want acme: the env write must still happen", gotEnv)
	}

	mu.Lock()
	defer mu.Unlock()
	preflightModes, _ := modeTiedMessages(msgs)
	if !reflect.DeepEqual(preflightModes, []string{"modeB"}) {
		t.Errorf("PreflightMsg modes = %v, want exactly [modeB]: the stale config-intent report must be suppressed", preflightModes)
	}
}

// TestHandleConfigIntentPreflightSuppressedByModeSwitchThatCommitsFirst covers
// the REVERSED ordering from TestHandleConfigIntentPreflightSuppressedByLaterModeSwitch
// above: there, the config intent dispatches (and snapshots) BEFORE the
// switch, so its snapshotted seq is strictly older than the switch's and the
// existing cfgSeq-changed check alone suppresses it.
//
// Here the SwitchModeIntent is dispatched FIRST via beginSwitchMode, which
// bumps cfgSeq synchronously, and is then pinned inside its own Preflight
// call before it can commit. Only THEN does the config intent (PreflightIntent)
// dispatch: it snapshots the CURRENT cfgSeq — which the pending switch has
// already bumped — together with the still-uncommitted, pre-switch mode. The
// config intent's own Preflight call is pinned too, so the test can force the
// switch to commit (mode -> modeB, and its PreflightMsg/GroupsMsg pair sent)
// strictly BEFORE releasing the config intent's Preflight and letting it
// reach withLatestDispatch.
//
// At that point cfgSeq is unchanged from what the config intent snapshotted
// (beginSwitchMode is the only thing that bumps it, and no second switch was
// dispatched), so the plain cfgSeq-changed check cannot tell the two apart:
// it would accept the config intent's report and append a stale, old-mode
// PreflightMsg AFTER the switch's already-rendered pair, misleading the
// operator into thinking the dashboard reverted to the old mode. The fix
// must additionally re-check the COMMITTED mode when emitting the report and
// drop it if the mode has since moved on.
func TestHandleConfigIntentPreflightSuppressedByModeSwitchThatCommitsFirst(t *testing.T) {
	root := t.TempDir()

	startedSwitch := make(chan struct{})
	proceedSwitch := make(chan struct{})
	startedCfg := make(chan struct{})
	proceedCfg := make(chan struct{})

	pf := &fakeprovider.Fake{
		NameVal:  "test",
		Packages: []string{"./..."},
		ModesVal: []provider.Mode{{Name: "modeA"}, {Name: "modeB"}},
		PreflightFn: func(_ context.Context, mode string, _ provider.TestRequirements) provider.PreflightReport {
			// Gate each call on its OWN pair of channels, keyed by mode, so
			// the switch's (modeB) and the config intent's (modeA) Preflight
			// calls can be released independently and in a deterministic
			// order, regardless of goroutine scheduling.
			switch mode {
			case "modeB":
				close(startedSwitch)
				<-proceedSwitch
			case "modeA":
				close(startedCfg)
				<-proceedCfg
			}
			return provider.PreflightReport{Mode: mode}
		},
		RequirementsFn: allowAllRequirements,
	}

	var mu sync.Mutex
	var msgs []tea.Msg
	send := func(m tea.Msg) {
		mu.Lock()
		msgs = append(msgs, m)
		mu.Unlock()
	}

	plan := engine.ExecutionPlan{Mode: "modeA", Eligible: []string{"TestAccThing"}}
	db := &dashboard{
		prov:      pf,
		root:      root,
		mode:      "modeA",
		statePath: filepath.Join(root, ".pulsar-state.json"),
		newRunner: func(_ io.Writer) testRunner { return &fakeRunner{} },
		list: func(_ context.Context, _ string, _ []string, _ string) ([]string, error) {
			return []string{"TestAccThing"}, nil
		},
		red:     redact.New(nil),
		plan:    &plan,
		send:    send,
		logSink: io.Discard,
	}

	// The switch is dispatched FIRST: beginSwitchMode bumps cfgSeq before the
	// config intent ever reads it.
	seqB := db.beginSwitchMode()
	doneSwitch := make(chan struct{})
	go func() {
		defer close(doneSwitch)
		db.applySwitchMode(context.Background(), tui.SwitchModeIntent{Mode: "modeB"}, seqB)
	}()
	<-startedSwitch // the switch is pinned inside Preflight, not yet committed.

	// The config intent dispatches SECOND, while the switch is still
	// pending: it snapshots the switch's own cfgSeq together with the
	// still-current, pre-switch mode.
	doneCfg := make(chan struct{})
	go func() {
		defer close(doneCfg)
		db.handleConfigIntent(context.Background(), tui.PreflightIntent{})
	}()
	<-startedCfg // the config intent is pinned inside Preflight with modeA.

	// Let the switch commit FIRST — this is the ordering the existing tests
	// do not cover.
	close(proceedSwitch)
	<-doneSwitch

	db.cfgMu.Lock()
	gotMode := db.mode
	db.cfgMu.Unlock()
	if gotMode != "modeB" {
		t.Fatalf("db.mode = %q after the switch commits, want modeB", gotMode)
	}

	// Only now release the config intent's stale Preflight call.
	close(proceedCfg)
	<-doneCfg

	db.cfgMu.Lock()
	gotMode = db.mode
	db.cfgMu.Unlock()
	if gotMode != "modeB" {
		t.Errorf("db.mode = %q after the stale config-intent completion, want unchanged modeB", gotMode)
	}

	mu.Lock()
	defer mu.Unlock()
	preflightModes, groupsModes := modeTiedMessages(msgs)
	if !reflect.DeepEqual(preflightModes, []string{"modeB"}) {
		t.Errorf("PreflightMsg modes = %v, want exactly [modeB]: the stale config-intent report (computed for modeA before the switch committed) must not land after the switch's pair", preflightModes)
	}
	if !reflect.DeepEqual(groupsModes, []string{"modeB"}) {
		t.Errorf("GroupsMsg plan modes = %v, want exactly [modeB]", groupsModes)
	}
}

// TestSeedPreflightSuppressedByLaterModeSwitch pins seed inside its provider
// Preflight call while a later SwitchModeIntent applies, proving seed's
// launch-mode pair is gated by the same latest-dispatch rule after the
// network call, not only before it.
func TestSeedPreflightSuppressedByLaterModeSwitch(t *testing.T) {
	root := t.TempDir()

	started := make(chan struct{})
	proceed := make(chan struct{})

	pf := &fakeprovider.Fake{
		NameVal:        "test",
		Packages:       []string{"./..."},
		ModesVal:       []provider.Mode{{Name: "modeA"}, {Name: "modeB"}},
		PreflightFn:    gatedPreflight(started, proceed),
		RequirementsFn: allowAllRequirements,
	}

	var mu sync.Mutex
	var msgs []tea.Msg
	send := func(m tea.Msg) {
		mu.Lock()
		msgs = append(msgs, m)
		mu.Unlock()
	}

	plan := engine.ExecutionPlan{Mode: "modeA", Eligible: []string{"TestAccThing"}}
	db := &dashboard{
		prov:      pf,
		root:      root,
		mode:      "modeA",
		statePath: filepath.Join(root, ".pulsar-state.json"),
		newRunner: func(_ io.Writer) testRunner { return &fakeRunner{} },
		list: func(_ context.Context, _ string, _ []string, _ string) ([]string, error) {
			return []string{"TestAccThing"}, nil
		},
		red:      redact.New(nil),
		plan:     &plan,
		groups:   []engine.Group{{Name: "misc", Tests: []string{"TestAccThing"}}},
		allNames: []string{"TestAccThing"},
		send:     send,
		logSink:  io.Discard,
	}

	doneSeed := make(chan struct{})
	go func() {
		defer close(doneSeed)
		db.seed(context.Background())
	}()
	<-started

	seqB := db.beginSwitchMode()
	db.applySwitchMode(context.Background(), tui.SwitchModeIntent{Mode: "modeB"}, seqB)

	close(proceed)
	<-doneSeed

	mu.Lock()
	defer mu.Unlock()
	preflightModes, groupsModes := modeTiedMessages(msgs)
	if !reflect.DeepEqual(preflightModes, []string{"modeB"}) {
		t.Errorf("PreflightMsg modes = %v, want exactly [modeB]: seed's superseded pair must be suppressed", preflightModes)
	}
	if !reflect.DeepEqual(groupsModes, []string{"modeB"}) {
		t.Errorf("GroupsMsg plan modes = %v, want exactly [modeB]", groupsModes)
	}
}

// TestModeTiedSendsAreAtomicAgainstLaterDispatch pins producer A INSIDE its
// first mode-tied send — after its freshness check has already passed — and
// then dispatches and runs a later switch B. Two properties are asserted:
//
//  1. Deadlock freedom: beginSwitchMode runs on the Bubble Tea event loop and
//     must never wait on a lock a producer holds while blocked in db.send
//     (tea.Program.Send's channel is unbuffered, so a send blocks until that
//     very event loop drains it). If the freshness check and the sends were
//     serialized under cfgMu instead of the dedicated send lock, this call
//     would block forever and the test would hang.
//  2. Atomicity: B cannot slip its own PreflightMsg/GroupsMsg pair between
//     A's check and A's sends, so the captured order is A's complete pair
//     followed by B's complete pair, and the operator's last view is B's.
func TestModeTiedSendsAreAtomicAgainstLaterDispatch(t *testing.T) {
	root := t.TempDir()

	sendStarted := make(chan struct{})
	sendProceed := make(chan struct{})

	pf := &fakeprovider.Fake{
		NameVal:  "test",
		Packages: []string{"./..."},
		ModesVal: []provider.Mode{{Name: "modeA"}, {Name: "modeB"}},
		PreflightFn: func(_ context.Context, m string, _ provider.TestRequirements) provider.PreflightReport {
			return provider.PreflightReport{Mode: m}
		},
		RequirementsFn: allowAllRequirements,
	}

	var mu sync.Mutex
	var msgs []tea.Msg
	sends := 0
	send := func(m tea.Msg) {
		mu.Lock()
		sends++
		first := sends == 1
		msgs = append(msgs, m)
		mu.Unlock()
		if first {
			close(sendStarted)
			<-sendProceed
		}
	}

	db := &dashboard{
		prov:      pf,
		root:      root,
		mode:      "anonymous",
		statePath: filepath.Join(root, ".pulsar-state.json"),
		newRunner: func(_ io.Writer) testRunner { return &fakeRunner{} },
		list: func(_ context.Context, _ string, _ []string, _ string) ([]string, error) {
			return []string{"TestAccThing"}, nil
		},
		red:     redact.New(nil),
		send:    send,
		logSink: io.Discard,
	}

	seqA := db.beginSwitchMode()
	doneA := make(chan struct{})
	go func() {
		defer close(doneA)
		db.applySwitchMode(context.Background(), tui.SwitchModeIntent{Mode: "modeA"}, seqA)
	}()
	<-sendStarted // A is past its freshness check, blocked mid-send.

	// Property 1: this is the event loop's call and must not block.
	seqB := db.beginSwitchMode()
	doneB := make(chan struct{})
	go func() {
		defer close(doneB)
		db.applySwitchMode(context.Background(), tui.SwitchModeIntent{Mode: "modeB"}, seqB)
	}()

	close(sendProceed)
	<-doneA
	<-doneB

	mu.Lock()
	defer mu.Unlock()
	preflightModes, groupsModes := modeTiedMessages(msgs)
	if !reflect.DeepEqual(preflightModes, []string{"modeA", "modeB"}) {
		t.Errorf("PreflightMsg modes = %v, want [modeA modeB]", preflightModes)
	}
	if !reflect.DeepEqual(groupsModes, []string{"modeA", "modeB"}) {
		t.Errorf("GroupsMsg plan modes = %v, want [modeA modeB]", groupsModes)
	}
	// Property 2: A's pair must be contiguous — B may not interleave.
	var order []string
	for _, m := range msgs {
		switch pm := m.(type) {
		case tui.PreflightMsg:
			order = append(order, "preflight:"+pm.Report.Mode)
		case tui.GroupsMsg:
			if pm.Plan != nil {
				order = append(order, "groups:"+pm.Plan.Mode)
			}
		}
	}
	want := []string{"preflight:modeA", "groups:modeA", "preflight:modeB", "groups:modeB"}
	if !reflect.DeepEqual(order, want) {
		t.Errorf("mode-tied send order = %v, want %v: a later dispatch must not interleave with an in-flight pair", order, want)
	}
	db.cfgMu.Lock()
	defer db.cfgMu.Unlock()
	if db.mode != "modeB" || db.plan == nil || db.plan.Mode != "modeB" {
		t.Errorf("db.mode=%q db.plan=%+v, want the latest dispatch modeB", db.mode, db.plan)
	}
}

func waitForDashboardSignal(t *testing.T, signal <-chan struct{}, description string) {
	t.Helper()
	select {
	case <-signal:
	case <-time.After(3 * time.Second):
		t.Fatalf("timed out waiting for %s", description)
	}
}

type testRunnerFunc func(context.Context, engine.RunSpec, func(engine.TestResult)) (engine.RunResult, error)

func (f testRunnerFunc) Run(ctx context.Context, spec engine.RunSpec, sink func(engine.TestResult)) (engine.RunResult, error) {
	return f(ctx, spec, sink)
}

// TestApplySwitchModeSupersededPlanningIsCanceled proves a newer dispatch
// actively cancels an older switch while it is blocked in planning. The
// planning lister waits on its context rather than a clock, and the test checks
// that the canceled switch neither commits nor emits a mode-tied result.
func TestApplySwitchModeSupersededPlanningIsCanceled(t *testing.T) {
	root := t.TempDir()
	started := make(chan struct{})
	canceled := make(chan struct{})
	release := make(chan struct{})
	defer close(release)

	pf := &fakeprovider.Fake{
		NameVal:        "test",
		Packages:       []string{"./..."},
		ModesVal:       []provider.Mode{{Name: "modeA"}, {Name: "modeB"}},
		RequirementsFn: allowAllRequirements,
		PreflightFn: func(_ context.Context, mode string, _ provider.TestRequirements) provider.PreflightReport {
			return provider.PreflightReport{Mode: mode}
		},
	}
	list := func(ctx context.Context, _ string, _ []string, _ string) ([]string, error) {
		if ctx == nil {
			t.Fatal("planning received a nil context")
		}
		select {
		case <-started:
			return []string{"TestAccModeB"}, nil
		default:
			close(started)
		}
		select {
		case <-ctx.Done():
			close(canceled)
			return nil, ctx.Err()
		case <-release:
			return nil, errors.New("test cleanup released planning")
		}
	}

	var mu sync.Mutex
	var msgs []tea.Msg
	db := &dashboard{
		prov:      pf,
		root:      root,
		mode:      "anonymous",
		statePath: filepath.Join(root, ".pulsar-state.json"),
		newRunner: func(io.Writer) testRunner { return &fakeRunner{} },
		list:      list,
		red:       redact.New(nil),
		send: func(msg tea.Msg) {
			mu.Lock()
			msgs = append(msgs, msg)
			mu.Unlock()
		},
		logSink: io.Discard,
	}

	seqA := db.beginSwitchMode()
	doneA := make(chan struct{})
	go func() {
		defer close(doneA)
		db.applySwitchMode(context.Background(), tui.SwitchModeIntent{Mode: "modeA"}, seqA)
	}()
	waitForDashboardSignal(t, started, "modeA planning to start")

	seqB := db.beginSwitchMode()
	waitForDashboardSignal(t, canceled, "modeA planning cancellation")
	db.applySwitchMode(context.Background(), tui.SwitchModeIntent{Mode: "modeB"}, seqB)
	waitForDashboardSignal(t, doneA, "modeA switch to finish")

	db.cfgMu.Lock()
	mode := db.mode
	plan := db.plan
	db.cfgMu.Unlock()
	if mode != "modeB" || plan == nil || plan.Mode != "modeB" {
		t.Fatalf("dashboard committed mode=%q plan=%+v, want modeB only", mode, plan)
	}
	mu.Lock()
	defer mu.Unlock()
	preflightModes, groupsModes := modeTiedMessages(msgs)
	if !reflect.DeepEqual(preflightModes, []string{"modeB"}) || !reflect.DeepEqual(groupsModes, []string{"modeB"}) {
		t.Errorf("mode-tied messages = preflight %v, groups %v; want modeB only", preflightModes, groupsModes)
	}
	for _, msg := range msgs {
		if _, ok := msg.(tui.ErrMsg); ok {
			t.Errorf("superseded planning must not send ErrMsg: %#v", msg)
		}
	}
}

// TestApplySwitchModeSupersededPreflightIsCanceled proves a newer dispatch
// cancels an older switch after planning but while provider Preflight is
// blocked. The stale switch cannot commit or send any result after that
// cancellation.
func TestApplySwitchModeSupersededPreflightIsCanceled(t *testing.T) {
	root := t.TempDir()
	started := make(chan struct{})
	canceled := make(chan struct{})
	release := make(chan struct{})
	defer close(release)

	pf := &fakeprovider.Fake{
		NameVal:        "test",
		Packages:       []string{"./..."},
		ModesVal:       []provider.Mode{{Name: "modeA"}, {Name: "modeB"}},
		RequirementsFn: allowAllRequirements,
		PreflightFn: func(ctx context.Context, mode string, _ provider.TestRequirements) provider.PreflightReport {
			if mode != "modeA" {
				return provider.PreflightReport{Mode: mode}
			}
			close(started)
			select {
			case <-ctx.Done():
				close(canceled)
				return provider.PreflightReport{Mode: mode}
			case <-release:
				return provider.PreflightReport{Mode: mode}
			}
		},
	}

	var mu sync.Mutex
	var msgs []tea.Msg
	db := &dashboard{
		prov:      pf,
		root:      root,
		mode:      "anonymous",
		statePath: filepath.Join(root, ".pulsar-state.json"),
		newRunner: func(io.Writer) testRunner { return &fakeRunner{} },
		list:      stubList([]string{"TestAccSwitch"}),
		red:       redact.New(nil),
		send: func(msg tea.Msg) {
			mu.Lock()
			msgs = append(msgs, msg)
			mu.Unlock()
		},
		logSink: io.Discard,
	}

	seqA := db.beginSwitchMode()
	doneA := make(chan struct{})
	go func() {
		defer close(doneA)
		db.applySwitchMode(context.Background(), tui.SwitchModeIntent{Mode: "modeA"}, seqA)
	}()
	waitForDashboardSignal(t, started, "modeA preflight to start")

	seqB := db.beginSwitchMode()
	waitForDashboardSignal(t, canceled, "modeA preflight cancellation")
	db.applySwitchMode(context.Background(), tui.SwitchModeIntent{Mode: "modeB"}, seqB)
	waitForDashboardSignal(t, doneA, "modeA switch to finish")

	db.cfgMu.Lock()
	mode := db.mode
	plan := db.plan
	db.cfgMu.Unlock()
	if mode != "modeB" || plan == nil || plan.Mode != "modeB" {
		t.Fatalf("dashboard committed mode=%q plan=%+v, want modeB only", mode, plan)
	}
	mu.Lock()
	defer mu.Unlock()
	preflightModes, groupsModes := modeTiedMessages(msgs)
	if !reflect.DeepEqual(preflightModes, []string{"modeB"}) || !reflect.DeepEqual(groupsModes, []string{"modeB"}) {
		t.Errorf("mode-tied messages = preflight %v, groups %v; want modeB only", preflightModes, groupsModes)
	}
}

func TestClearSwitchContextOnlyClearsItsOwnCompletedSwitch(t *testing.T) {
	db := &dashboard{}
	seqA := db.beginSwitchMode()
	ctxA, ok := db.switchContext(seqA)
	if !ok {
		t.Fatal("expected first switch context")
	}
	seqB := db.beginSwitchMode()
	ctxB, ok := db.switchContext(seqB)
	if !ok {
		t.Fatal("expected second switch context")
	}
	waitForDashboardSignal(t, ctxA.Done(), "superseded switch cancellation")

	// A's deferred cleanup runs after B has taken ownership. It must neither
	// clear nor cancel B's in-flight context.
	db.clearSwitchContext(seqA)
	if _, ok := db.switchContext(seqB); !ok {
		t.Fatal("stale cleanup cleared the newer switch ownership")
	}
	select {
	case <-ctxB.Done():
		t.Fatal("stale cleanup canceled the newer switch")
	default:
	}

	db.clearSwitchContext(seqB)
	if _, ok := db.switchContext(seqB); ok {
		t.Fatal("completed switch ownership was not cleared")
	}
}

// TestDashboardLifecycleCancellationPropagatesToWorkers proves dashboard-owned
// lifetime cancellation reaches workers even when their callers otherwise pass
// context.Background. Each worker blocks on the context it receives, so the
// test's channels establish the cancellation ordering without sleeps.
func TestDashboardLifecycleCancellationPropagatesToWorkers(t *testing.T) {
	t.Run("seed", func(t *testing.T) {
		lifetime, cancel := context.WithCancel(context.Background())
		defer cancel()
		started := make(chan struct{})
		canceled := make(chan struct{})
		done := make(chan struct{})
		pf := &fakeprovider.Fake{
			NameVal: "test",
			PreflightFn: func(ctx context.Context, _ string, _ provider.TestRequirements) provider.PreflightReport {
				close(started)
				<-ctx.Done()
				close(canceled)
				return provider.PreflightReport{}
			},
		}
		plan := planWithEligible("anonymous", []string{"TestAccLifecycle"})
		db := &dashboard{
			prov:         pf,
			mode:         "anonymous",
			statePath:    filepath.Join(t.TempDir(), ".pulsar-state.json"),
			lifecycleCtx: lifetime,
			plan:         plan,
			red:          redact.New(nil),
			send:         func(tea.Msg) {},
		}
		go func() {
			defer close(done)
			db.seed(context.Background())
		}()
		waitForDashboardSignal(t, started, "seed preflight to start")
		cancel()
		waitForDashboardSignal(t, canceled, "seed lifetime cancellation")
		waitForDashboardSignal(t, done, "seed to finish")
	})

	t.Run("config preflight", func(t *testing.T) {
		lifetime, cancel := context.WithCancel(context.Background())
		defer cancel()
		started := make(chan struct{})
		canceled := make(chan struct{})
		done := make(chan struct{})
		pf := &fakeprovider.Fake{
			NameVal: "test",
			PreflightFn: func(ctx context.Context, _ string, _ provider.TestRequirements) provider.PreflightReport {
				close(started)
				<-ctx.Done()
				close(canceled)
				return provider.PreflightReport{}
			},
		}
		plan := planWithEligible("anonymous", []string{"TestAccLifecycle"})
		db := &dashboard{
			prov:         pf,
			mode:         "anonymous",
			lifecycleCtx: lifetime,
			plan:         plan,
			red:          redact.New(nil),
			send:         func(tea.Msg) {},
		}
		go func() {
			defer close(done)
			db.handleConfigIntent(context.Background(), tui.PreflightIntent{})
		}()
		waitForDashboardSignal(t, started, "config preflight to start")
		cancel()
		waitForDashboardSignal(t, canceled, "config preflight lifetime cancellation")
		waitForDashboardSignal(t, done, "config preflight to finish")
	})

	t.Run("switch", func(t *testing.T) {
		lifetime, cancel := context.WithCancel(context.Background())
		defer cancel()
		started := make(chan struct{})
		canceled := make(chan struct{})
		done := make(chan struct{})
		pf := &fakeprovider.Fake{
			NameVal:        "test",
			Packages:       []string{"./..."},
			ModesVal:       []provider.Mode{{Name: "anonymous"}, {Name: "enterprise"}},
			RequirementsFn: allowAllRequirements,
			PreflightFn: func(ctx context.Context, _ string, _ provider.TestRequirements) provider.PreflightReport {
				close(started)
				<-ctx.Done()
				close(canceled)
				return provider.PreflightReport{}
			},
		}
		plan := planWithEligible("anonymous", []string{"TestAccLifecycle"})
		db := &dashboard{
			prov:         pf,
			root:         t.TempDir(),
			mode:         "anonymous",
			lifecycleCtx: lifetime,
			list:         stubList([]string{"TestAccLifecycle"}),
			plan:         plan,
			red:          redact.New(nil),
			send:         func(tea.Msg) {},
		}
		seq := db.beginSwitchMode()
		go func() {
			defer close(done)
			db.applySwitchMode(context.Background(), tui.SwitchModeIntent{Mode: "enterprise"}, seq)
		}()
		waitForDashboardSignal(t, started, "switch preflight to start")
		cancel()
		waitForDashboardSignal(t, canceled, "switch lifetime cancellation")
		waitForDashboardSignal(t, done, "switch to finish")
		if db.mode != "anonymous" {
			t.Errorf("canceled switch committed mode %q, want anonymous", db.mode)
		}
	})

	t.Run("run", func(t *testing.T) {
		lifetime, cancel := context.WithCancel(context.Background())
		defer cancel()
		started := make(chan struct{})
		canceled := make(chan struct{})
		done := make(chan struct{})
		runner := testRunnerFunc(func(ctx context.Context, _ engine.RunSpec, _ func(engine.TestResult)) (engine.RunResult, error) {
			close(started)
			<-ctx.Done()
			close(canceled)
			return engine.RunResult{}, ctx.Err()
		})
		plan := planWithEligible("anonymous", []string{"TestAccLifecycle"})
		db := &dashboard{
			prov:         &fakeprovider.Fake{NameVal: "test", Packages: []string{"./..."}},
			root:         t.TempDir(),
			mode:         "anonymous",
			statePath:    filepath.Join(t.TempDir(), ".pulsar-state.json"),
			lifecycleCtx: lifetime,
			newRunner:    func(io.Writer) testRunner { return runner },
			plan:         plan,
			groups:       []engine.Group{{Name: "tests", Tests: []string{"TestAccLifecycle"}}},
			allNames:     []string{"TestAccLifecycle"},
			red:          redact.New(nil),
			send:         func(tea.Msg) {},
			logSink:      io.Discard,
		}
		go func() {
			defer close(done)
			db.runIntent(context.Background(), tui.RetryTestIntent{Test: "TestAccLifecycle"})
		}()
		waitForDashboardSignal(t, started, "run to start")
		cancel()
		waitForDashboardSignal(t, canceled, "run lifetime cancellation")
		waitForDashboardSignal(t, done, "run to finish")
	})
}

// TestRunIntentCanceledAfterStreamedResultPersistsPartialState proves that a
// run canceled AFTER the runner has already streamed a result through sink
// still merges and persists that partial delta — state.Save, the suite
// failure state, and per-test failure logs must not be skipped just because
// ctx was canceled while the runner was still in flight. It also asserts the
// outward-facing completion signal (RunDoneMsg) is suppressed for the
// canceled run, so the operator is never shown a misleading "run finished"
// banner for work that was actually cut short.
func TestRunIntentCanceledAfterStreamedResultPersistsPartialState(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")
	failDir := filepath.Join(root, ".pulsar-failures")
	const pkg = "github.com/integrations/terraform-provider-github/github"

	st := engine.State{Provider: "test", Mode: "anonymous"}
	st.Plan = planWithEligible("anonymous", []string{"TestAccOne", "TestAccTwo"})
	st.Results = []engine.PersistResult{
		{Package: pkg, Test: "TestAccOne", Status: "fail"},
		{Package: pkg, Test: "TestAccTwo", Status: "fail"},
	}
	if err := st.Save(statePath); err != nil {
		t.Fatalf("seeding state: %v", err)
	}

	// The pre-seeded TestAccOne result has Elapsed 0; the streamed result
	// below reports a distinct Elapsed so persistence (or the lack of it)
	// after cancellation is unambiguous regardless of status.
	const streamedElapsed = 42.5
	streamed := make(chan struct{})
	runner := testRunnerFunc(func(ctx context.Context, _ engine.RunSpec, sink func(engine.TestResult)) (engine.RunResult, error) {
		tr := engine.TestResult{Package: pkg, Name: "TestAccOne", Status: provider.StatusFail, Elapsed: streamedElapsed}
		sink(tr)
		close(streamed)
		<-ctx.Done()
		// The runner itself returns only its own partial result plus the
		// cancellation error — runIntent must not discard that partial
		// delta just because runErr wraps ctx.Err().
		return engine.RunResult{Tests: []engine.TestResult{tr}}, ctx.Err()
	})

	var mu sync.Mutex
	var msgs []tea.Msg
	send := func(m tea.Msg) {
		mu.Lock()
		msgs = append(msgs, m)
		mu.Unlock()
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db := &dashboard{
		prov:      &fakeprovider.Fake{NameVal: "test", Packages: []string{pkg}},
		root:      root,
		mode:      "anonymous",
		statePath: statePath,
		newRunner: func(io.Writer) testRunner { return runner },
		list:      stubList([]string{"TestAccOne", "TestAccTwo"}),
		red:       redact.New(nil),
		groups:    []engine.Group{{Name: "tests", Tests: []string{"TestAccOne", "TestAccTwo"}}},
		allNames:  []string{"TestAccOne", "TestAccTwo"},
		send:      send,
		logSink:   io.Discard,
	}

	done := make(chan struct{})
	go func() {
		defer close(done)
		db.runIntent(ctx, tui.RetryAllIntent{})
	}()

	waitForDashboardSignal(t, streamed, "run to stream its first result")
	cancel()
	waitForDashboardSignal(t, done, "runIntent to finish after cancellation")

	state, err := engine.Load(statePath)
	if err != nil {
		t.Fatalf("loading state: %v", err)
	}
	var gotOne bool
	for _, r := range state.Results {
		if r.Test == "TestAccOne" && r.Elapsed == streamedElapsed {
			gotOne = true
		}
	}
	if !gotOne {
		t.Errorf("state.Results = %+v, want the streamed TestAccOne result (Elapsed=%v) merged and persisted despite cancellation", state.Results, streamedElapsed)
	}

	wantLog := engine.FailureLogPath(failDir, pkg, "TestAccOne")
	if _, err := os.Stat(wantLog); err != nil {
		t.Errorf("failure log %q not written for the streamed canceled result: %v", wantLog, err)
	}

	mu.Lock()
	defer mu.Unlock()
	for _, m := range msgs {
		if _, ok := m.(tui.RunDoneMsg); ok {
			t.Error("expected no RunDoneMsg after cancellation, run was cut short")
		}
	}
}
