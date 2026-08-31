package tui

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/github/terraform-provider-tester/engine"
	"github.com/github/terraform-provider-tester/provider"
)

func newTestModel() Model {
	return New("github", "organization", "my-test-org", true)
}

func TestUpdateWindowSizeSetsDimensions(t *testing.T) {
	m := newTestModel()
	m2, _ := m.Update(tea.WindowSizeMsg{Width: 120, Height: 40})
	got := m2.(Model)
	if got.width != 120 || got.height != 40 {
		t.Fatalf("expected width=120 height=40, got width=%d height=%d", got.width, got.height)
	}
}

func TestUpdateTabCyclesSections(t *testing.T) {
	m := newTestModel()
	// should start at Preflight
	if m.section != sectionPreflight {
		t.Fatalf("expected section=Preflight, got %d", m.section)
	}

	advance := func(m Model) Model {
		m2, _ := m.Update(tea.KeyMsg{Type: tea.KeyTab})
		return m2.(Model)
	}
	prev := func(m Model) Model {
		m2, _ := m.Update(tea.KeyMsg{Type: tea.KeyShiftTab})
		return m2.(Model)
	}

	m = advance(m)
	if m.section != sectionGroups {
		t.Fatalf("after tab: expected sectionGroups, got %d", m.section)
	}
	m = advance(m)
	if m.section != sectionRun {
		t.Fatalf("after tab: expected sectionRun, got %d", m.section)
	}
	m = advance(m)
	if m.section != sectionTriage {
		t.Fatalf("after tab: expected sectionTriage, got %d", m.section)
	}
	m = advance(m)
	if m.section != sectionPreflight {
		t.Fatalf("after tab: expected sectionPreflight (wrap), got %d", m.section)
	}

	// shift+tab reverses
	m = prev(m)
	if m.section != sectionTriage {
		t.Fatalf("after shift+tab: expected sectionTriage, got %d", m.section)
	}
	m = prev(m)
	if m.section != sectionRun {
		t.Fatalf("after shift+tab: expected sectionRun, got %d", m.section)
	}
	m = prev(m)
	if m.section != sectionGroups {
		t.Fatalf("after shift+tab: expected sectionGroups, got %d", m.section)
	}
}

func TestUpdateTabCyclesFourSections(t *testing.T) {
	m := newTestModel()
	want := []section{sectionGroups, sectionRun, sectionTriage, sectionPreflight}
	for i, sectionWant := range want {
		next, _ := m.Update(tea.KeyMsg{Type: tea.KeyTab})
		m = next.(Model)
		if m.section != sectionWant {
			t.Fatalf("step %d section = %d, want %d", i, m.section, sectionWant)
		}
	}
}

func TestUpdateArrowKeysCycleTabs(t *testing.T) {
	m := newTestModel()

	next, _ := m.Update(tea.KeyMsg{Type: tea.KeyRight})
	m = next.(Model)
	if m.section != sectionGroups {
		t.Fatalf("Right section = %d, want Groups", m.section)
	}

	next, _ = m.Update(tea.KeyMsg{Type: tea.KeyLeft})
	m = next.(Model)
	if m.section != sectionPreflight {
		t.Fatalf("Left section = %d, want Preflight", m.section)
	}

	next, _ = m.Update(tea.KeyMsg{Type: tea.KeyLeft})
	m = next.(Model)
	if m.section != sectionTriage {
		t.Fatalf("wrapped Left section = %d, want Triage", m.section)
	}
}

func TestUpdateNumberKeysJumpToTabs(t *testing.T) {
	tests := []struct {
		key  rune
		want section
	}{{'1', sectionPreflight}, {'2', sectionGroups}, {'3', sectionRun}, {'4', sectionTriage}}

	for _, tc := range tests {
		m := newTestModel()
		next, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{tc.key}})
		if got := next.(Model).section; got != tc.want {
			t.Fatalf("key %q section = %d, want %d", tc.key, got, tc.want)
		}
	}
}

func TestUpdateQuitReturnsQuitCmd(t *testing.T) {
	m := newTestModel()
	_, cmd := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("q")})
	if cmd == nil {
		t.Fatal("expected non-nil cmd for q")
	}
	if msg := cmd(); msg != tea.QuitMsg(struct{}{}) {
		t.Fatalf("expected QuitMsg, got %T", msg)
	}

	_, cmd2 := m.Update(tea.KeyMsg{Type: tea.KeyCtrlC})
	if cmd2 == nil {
		t.Fatal("expected non-nil cmd for ctrl+c")
	}
	if msg := cmd2(); msg != tea.QuitMsg(struct{}{}) {
		t.Fatalf("expected QuitMsg from ctrl+c, got %T", msg)
	}
}

func TestUpdatePreflightMsgPopulatesChecks(t *testing.T) {
	m := newTestModel()
	report := provider.PreflightReport{
		Mode: "organization",
		Checks: []provider.Check{
			{Name: "identity", Status: provider.CheckOK, Detail: "ok"},
			{Name: "rate-limit", Status: provider.CheckWarn, Detail: "limited", Fix: "wait"},
		},
	}
	m2, _ := m.Update(PreflightMsg{Report: report})
	got := m2.(Model)
	if got.mode != "organization" {
		t.Fatalf("expected mode=organization, got %q", got.mode)
	}
	if len(got.checks) != 2 {
		t.Fatalf("expected 2 checks, got %d", len(got.checks))
	}
}

func TestUpdateGroupsMsgPopulatesGroups(t *testing.T) {
	m := newTestModel()
	groups := []engine.Group{
		{Name: "repositories", Tests: []string{"TestAccA", "TestAccB"}},
		{Name: "teams", Tests: []string{"TestAccC"}},
	}
	m2, _ := m.Update(GroupsMsg{Groups: groups})
	got := m2.(Model)
	if len(got.groups) != 2 {
		t.Fatalf("expected 2 groups, got %d", len(got.groups))
	}
}

// TestUpdateGroupsMsgWithPlanStillPopulatesGroups proves GroupsMsg carries an
// optional Plan (the ExecutionPlan the CLI wiring built for the current mode)
// without changing how the model handles Groups: the TUI stores no new
// screen or field for Plan, so a message that sets it must behave exactly
// like one that does not.
func TestUpdateGroupsMsgWithPlanStillPopulatesGroups(t *testing.T) {
	m := newTestModel()
	groups := []engine.Group{
		{Name: "repositories", Tests: []string{"TestAccA", "TestAccB"}},
	}
	plan := &engine.ExecutionPlan{Mode: "organization", Eligible: []string{"TestAccA", "TestAccB"}}
	m2, _ := m.Update(GroupsMsg{Groups: groups, Plan: plan})
	got := m2.(Model)
	if len(got.groups) != 1 {
		t.Fatalf("expected 1 group, got %d", len(got.groups))
	}
	if got.groups[0].Name != "repositories" {
		t.Fatalf("expected group %q, got %q", "repositories", got.groups[0].Name)
	}
}

func TestUpdateTestUpdateMsgUpsertsResult(t *testing.T) {
	m := newTestModel()

	r1 := engine.TestResult{Package: "pkg", Name: "TestFoo", Sub: "", Status: provider.StatusRunning}
	m2, _ := m.Update(TestUpdateMsg{Result: r1})
	m = m2.(Model)
	if len(m.results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(m.results))
	}

	// Same (Name, Sub) → upsert, not duplicate
	r2 := engine.TestResult{Package: "pkg", Name: "TestFoo", Sub: "", Status: provider.StatusPass}
	m3, _ := m.Update(TestUpdateMsg{Result: r2})
	m = m3.(Model)
	if len(m.results) != 1 {
		t.Fatalf("expected 1 result after upsert, got %d", len(m.results))
	}
	if m.results[0].Status != provider.StatusPass {
		t.Fatalf("expected StatusPass after upsert, got %v", m.results[0].Status)
	}
}

// TestUpdateTestResultZeroModelNilMapSafe guards against a panic when a
// zero-valued Model (resultIndex == nil) receives a TestUpdateMsg. Update must
// lazily initialize the map rather than assigning into a nil map.
func TestUpdateTestResultZeroModelNilMapSafe(t *testing.T) {
	var m Model // zero value: resultIndex is nil
	r := engine.TestResult{Package: "pkg", Name: "TestFoo", Status: provider.StatusPass}
	m2, _ := m.Update(TestUpdateMsg{Result: r})
	got := m2.(Model)
	if len(got.results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(got.results))
	}
}

func TestUpdateRateMsgSetsRate(t *testing.T) {
	m := newTestModel()
	reset := time.Now().Add(15 * time.Minute)
	m2, _ := m.Update(RateMsg{Remaining: 4870, Limit: 5000, Reset: reset})
	got := m2.(Model)
	if got.rate.Remaining != 4870 || got.rate.Limit != 5000 {
		t.Fatalf("expected rate 4870/5000, got %d/%d", got.rate.Remaining, got.rate.Limit)
	}
	if !got.rate.Reset.Equal(reset) {
		t.Fatalf("expected reset %v, got %v", reset, got.rate.Reset)
	}
}

func TestUpdateErrMsgSetsStatus(t *testing.T) {
	m := newTestModel()
	m2, _ := m.Update(ErrMsg{Err: errors.New("something went wrong")})
	got := m2.(Model)
	if got.err == nil {
		t.Fatal("expected err to be set")
	}
	if got.status == "" {
		t.Fatal("expected non-empty status after ErrMsg")
	}
}

// TestUpdateErrMsgNilDoesNotPanic guards against a nil-interface dereference
// when a producer emits ErrMsg{Err: nil}. Update must not call Error() on a nil
// error.
func TestUpdateErrMsgNilDoesNotPanic(t *testing.T) {
	m := newTestModel()
	m2, _ := m.Update(ErrMsg{Err: nil})
	got := m2.(Model)
	if got.err != nil {
		t.Fatalf("expected nil err, got %v", got.err)
	}
}

func TestResultOverlayEscClearsDangerOverlayState(t *testing.T) {
	m := newTestModel()
	m.setResultOverlayWithTone(panelDanger, "Sweep failed", "residual state: unknown")

	next, cmd := m.Update(tea.KeyMsg{Type: tea.KeyEsc})
	got := next.(Model)
	if cmd != nil {
		t.Fatal("result overlay esc emitted a command")
	}

	assertResultOverlayCleared(t, got)
}

func TestOrphansListedMsgClearsDangerResultOverlayState(t *testing.T) {
	m := newTestModel()
	m.setResultOverlayWithTone(panelDanger, "Sweep failed", "residual state: unknown")

	next, _ := m.Update(OrphansListedMsg{Owner: "acme", Resources: sampleResources()})
	got := next.(Model)

	if !got.orphanOverlayActive {
		t.Fatal("OrphansListedMsg should open the orphan overlay")
	}
	if got.orphanOwner != "acme" {
		t.Fatalf("orphanOwner = %q, want acme", got.orphanOwner)
	}
	if len(got.orphans) != len(sampleResources()) {
		t.Fatalf("orphans = %d, want %d", len(got.orphans), len(sampleResources()))
	}

	assertResultOverlayCleared(t, got)
}

func TestUpdateHelpToggles(t *testing.T) {
	m := newTestModel()
	if m.help.ShowAll {
		t.Fatal("expected ShowAll=false initially")
	}
	m2, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("?")})
	got := m2.(Model)
	if !got.help.ShowAll {
		t.Fatal("expected ShowAll=true after ?")
	}
	m3, _ := got.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("?")})
	got2 := m3.(Model)
	if got2.help.ShowAll {
		t.Fatal("expected ShowAll=false after second ?")
	}
}

func TestUpdateEscapeClosesExpandedHelp(t *testing.T) {
	m := newTestModel()
	next, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("?")})
	m = next.(Model)
	if !m.help.ShowAll {
		t.Fatal("help did not open")
	}

	next, _ = m.Update(tea.KeyMsg{Type: tea.KeyEsc})
	if next.(Model).help.ShowAll {
		t.Fatal("Escape did not close help")
	}
}

func comparableModel(m Model) Model {
	m.exec = nil
	m.copyFn = nil
	m.cmdFor = nil
	m.getenv = nil
	return m
}

func assertModelUnchanged(t *testing.T, before, after Model) {
	t.Helper()
	if !reflect.DeepEqual(comparableModel(before), comparableModel(after)) {
		t.Fatalf("model changed while help was open")
	}
}

func TestUpdateExpandedHelpBlocksDashboardShortcuts(t *testing.T) {
	tests := []struct {
		name  string
		model func() Model
		msg   tea.KeyMsg
	}{
		{
			name: "right arrow",
			model: func() Model {
				m := newTestModel()
				m.help.ShowAll = true
				return m
			},
			msg: tea.KeyMsg{Type: tea.KeyRight},
		},
		{
			name: "number jump",
			model: func() Model {
				m := newTestModel()
				m.help.ShowAll = true
				return m
			},
			msg: tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("2")},
		},
		{
			name: "enter drill in",
			model: func() Model {
				m := newTestModel()
				m.section = sectionGroups
				m.groups = []engine.Group{{Name: "repositories", Tests: []string{"TestAccA"}}}
				m.help.ShowAll = true
				return m
			},
			msg: tea.KeyMsg{Type: tea.KeyEnter},
		},
		{
			name: "mode picker",
			model: func() Model {
				m := newTestModel()
				m.pickerModes = []provider.Mode{{Name: "organization"}, {Name: "enterprise"}}
				m.help.ShowAll = true
				return m
			},
			msg: tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("m")},
		},
		{
			name: "variables editor",
			model: func() Model {
				m := newTestModel()
				m.help.ShowAll = true
				return m
			},
			msg: tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("v")},
		},
		{
			name: "retry command",
			model: func() Model {
				m := newTestModel()
				m.section = sectionGroups
				m.groups = []engine.Group{{Name: "repositories", Tests: []string{"TestAccA"}}}
				m.help.ShowAll = true
				return m
			},
			msg: tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("r")},
		},
		{
			name: "triage refresh",
			model: func() Model {
				m := newTestModel()
				m.section = sectionTriage
				m.help.ShowAll = true
				return m
			},
			msg: tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("t")},
		},
		{
			name: "triage sync",
			model: func() Model {
				m := newTestModel()
				m.section = sectionTriage
				m.help.ShowAll = true
				return m
			},
			msg: tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("K")},
		},
		{
			name: "triage file issue",
			model: func() Model {
				m := newTestModel()
				m.section = sectionTriage
				m.triageFailures = sampleFailures()
				m.triageCursor = 2
				m.triageDetailActive = true
				m.help.ShowAll = true
				return m
			},
			msg: tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("i")},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			before := tc.model()
			next, cmd := before.Update(tc.msg)
			if cmd != nil {
				t.Fatalf("unexpected command for %s while help is open", tc.name)
			}
			assertModelUnchanged(t, before, next.(Model))
		})
	}
}

func TestUpdateExpandedHelpCloseKeys(t *testing.T) {
	tests := []struct {
		name string
		msg  tea.KeyMsg
	}{
		{name: "question mark", msg: tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("?")}},
		{name: "escape", msg: tea.KeyMsg{Type: tea.KeyEsc}},
		{name: "h", msg: tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("h")}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			before := newTestModel()
			before.help.ShowAll = true

			next, cmd := before.Update(tc.msg)
			if cmd != nil {
				t.Fatalf("unexpected command for %s", tc.name)
			}
			got := next.(Model)
			want := before
			want.help.ShowAll = false
			if !reflect.DeepEqual(comparableModel(got), comparableModel(want)) {
				t.Fatalf("%s did not only close help", tc.name)
			}
		})
	}
}

func TestUpdateExpandedHelpQuitKeys(t *testing.T) {
	tests := []struct {
		name string
		msg  tea.KeyMsg
	}{
		{name: "q", msg: tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("q")}},
		{name: "ctrl+c", msg: tea.KeyMsg{Type: tea.KeyCtrlC}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m := newTestModel()
			m.help.ShowAll = true

			_, cmd := m.Update(tc.msg)
			if cmd == nil {
				t.Fatalf("expected quit command for %s", tc.name)
			}
			if msg := cmd(); msg != tea.QuitMsg(struct{}{}) {
				t.Fatalf("%s returned %T, want QuitMsg", tc.name, msg)
			}
		})
	}
}

func TestUpdateExpandedHelpStillProcessesNonKeyMessages(t *testing.T) {
	t.Run("window size", func(t *testing.T) {
		m := newTestModel()
		m.help.ShowAll = true

		next, cmd := m.Update(tea.WindowSizeMsg{Width: 120, Height: 40})
		if cmd != nil {
			t.Fatal("unexpected command for window size")
		}
		got := next.(Model)
		if got.width != 120 || got.height != 40 {
			t.Fatalf("expected width=120 height=40, got width=%d height=%d", got.width, got.height)
		}
		if !got.help.ShowAll {
			t.Fatal("window size should not close help")
		}
	})

	t.Run("triage loaded", func(t *testing.T) {
		m := newTestModel()
		m.section = sectionTriage
		m.help.ShowAll = true

		wantFailures := sampleFailures()
		next, cmd := m.Update(TriageLoadedMsg{Failures: wantFailures, CacheAvailable: true})
		if cmd != nil {
			t.Fatal("unexpected command for triage loaded")
		}
		got := next.(Model)
		if len(got.triageFailures) != len(wantFailures) {
			t.Fatalf("triageFailures = %d, want %d", len(got.triageFailures), len(wantFailures))
		}
		if !got.triageCacheAvailable {
			t.Fatal("triageCacheAvailable = false, want true")
		}
		if !got.help.ShowAll {
			t.Fatal("triage load should not close help")
		}
	})
}

func TestUpdateCursorClampsOnEmpty(t *testing.T) {
	m := newTestModel()
	// cursor should be 0 on empty list; j should not panic
	m2, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("j")})
	got := m2.(Model)
	if got.cursor != 0 {
		t.Fatalf("expected cursor=0 on empty list after j, got %d", got.cursor)
	}
	m3, _ := got.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("k")})
	got2 := m3.(Model)
	if got2.cursor != 0 {
		t.Fatalf("expected cursor=0 on empty list after k, got %d", got2.cursor)
	}
}

// helper: send a key rune to a model.
func pressKey(m Model, r rune) (Model, tea.Cmd) {
	m2, cmd := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{r}})
	return m2.(Model), cmd
}

func assertResultOverlayCleared(t *testing.T, got Model) {
	t.Helper()
	if got.resultOverlayActive {
		t.Fatal("resultOverlayActive = true, want false")
	}
	if got.resultTitle != "" {
		t.Fatalf("resultTitle = %q, want empty", got.resultTitle)
	}
	if got.resultLines != nil {
		t.Fatalf("resultLines = %v, want nil", got.resultLines)
	}
	if got.resultTone != panelSuccess {
		t.Fatalf("resultTone = %v, want %v", got.resultTone, panelSuccess)
	}
}

func TestDrillInOutCyclesFocus(t *testing.T) {
	m := newTestModel()
	m.section = sectionGroups
	m.groups = []engine.Group{
		{Name: "repositories", Tests: []string{"TestAccA", "TestAccB"}},
	}
	m.results = []engine.TestResult{
		{Package: "pkg", Name: "TestAccA", Sub: "", Status: provider.StatusPass, Elapsed: 1.0},
	}
	m.resultIndex = map[resultKey]int{
		{Name: "TestAccA", Sub: ""}: 0,
	}
	// Start at focusGroups
	if m.focus != focusGroups {
		t.Fatalf("expected focusGroups initially, got %d", m.focus)
	}

	// DrillIn: focusGroups → focusTests
	m, _ = pressKey(m, 'l')
	if m.focus != focusTests {
		t.Fatalf("after DrillIn: expected focusTests, got %d", m.focus)
	}
	if m.testCursor != 0 {
		t.Fatalf("after DrillIn: expected testCursor=0, got %d", m.testCursor)
	}

	// DrillIn: focusTests → focusLog
	m, _ = pressKey(m, 'l')
	if m.focus != focusLog {
		t.Fatalf("after DrillIn: expected focusLog, got %d", m.focus)
	}

	// DrillIn: focusLog → no-op (stays focusLog)
	m, _ = pressKey(m, 'l')
	if m.focus != focusLog {
		t.Fatalf("after DrillIn at log: expected focusLog (no-op), got %d", m.focus)
	}

	// Back: focusLog → focusTests
	m, _ = pressKey(m, 'h')
	if m.focus != focusTests {
		t.Fatalf("after Back: expected focusTests, got %d", m.focus)
	}

	// Back: focusTests → focusGroups
	m, _ = pressKey(m, 'h')
	if m.focus != focusGroups {
		t.Fatalf("after Back: expected focusGroups, got %d", m.focus)
	}

	// Back: focusGroups → no-op
	m, _ = pressKey(m, 'h')
	if m.focus != focusGroups {
		t.Fatalf("after Back at groups: expected focusGroups (no-op), got %d", m.focus)
	}
}

func TestCursorMovesPerFocusLevel(t *testing.T) {
	m := newTestModel()
	m.section = sectionGroups
	m.groups = []engine.Group{
		{Name: "g1", Tests: []string{"TestA"}},
		{Name: "g2", Tests: []string{"TestB", "TestC"}},
	}
	m.results = []engine.TestResult{
		{Package: "pkg", Name: "TestA", Sub: "", Status: provider.StatusPass},
		{Package: "pkg", Name: "TestB", Sub: "", Status: provider.StatusPass},
		{Package: "pkg", Name: "TestC", Sub: "", Status: provider.StatusFail},
	}
	m.resultIndex = map[resultKey]int{
		{Name: "TestA", Sub: ""}: 0,
		{Name: "TestB", Sub: ""}: 1,
		{Name: "TestC", Sub: ""}: 2,
	}

	// At focusGroups, Down moves groupCursor
	m, _ = pressKey(m, 'j')
	if m.groupCursor != 1 {
		t.Fatalf("focusGroups Down: want groupCursor=1, got %d", m.groupCursor)
	}
	// Clamp at end
	m, _ = pressKey(m, 'j')
	if m.groupCursor != 1 {
		t.Fatalf("focusGroups Down clamp: want groupCursor=1, got %d", m.groupCursor)
	}
	// Up moves back
	m, _ = pressKey(m, 'k')
	if m.groupCursor != 0 {
		t.Fatalf("focusGroups Up: want groupCursor=0, got %d", m.groupCursor)
	}
	// Up clamp at 0
	m, _ = pressKey(m, 'k')
	if m.groupCursor != 0 {
		t.Fatalf("focusGroups Up clamp: want groupCursor=0, got %d", m.groupCursor)
	}

	// Drill into group 1 (index 1) which has 2 tests
	m.groupCursor = 1
	m, _ = pressKey(m, 'l') // → focusTests
	if m.focus != focusTests {
		t.Fatalf("expected focusTests after drill, got %d", m.focus)
	}
	// At focusTests, Down moves testCursor
	m, _ = pressKey(m, 'j')
	if m.testCursor != 1 {
		t.Fatalf("focusTests Down: want testCursor=1, got %d", m.testCursor)
	}
	// Clamp at end
	m, _ = pressKey(m, 'j')
	if m.testCursor != 1 {
		t.Fatalf("focusTests Down clamp: want testCursor=1, got %d", m.testCursor)
	}
	// Up
	m, _ = pressKey(m, 'k')
	if m.testCursor != 0 {
		t.Fatalf("focusTests Up: want testCursor=0, got %d", m.testCursor)
	}

	// Empty group: no panic on Up/Down
	mEmpty := newTestModel()
	mEmpty.section = sectionGroups
	mEmpty.groups = []engine.Group{}
	mEmpty, _ = pressKey(mEmpty, 'j')
	mEmpty, _ = pressKey(mEmpty, 'k')
	if mEmpty.groupCursor != 0 {
		t.Fatalf("empty groups: expected groupCursor=0, got %d", mEmpty.groupCursor)
	}
}

func TestRetryKeyEmitsRetryGroupIntent(t *testing.T) {
	m := newTestModel()
	m.section = sectionGroups
	m.groups = []engine.Group{
		{Name: "repositories", Tests: []string{"TestA"}},
		{Name: "teams", Tests: []string{"TestB"}},
	}
	m.groupCursor = 1
	m.focus = focusGroups

	_, cmd := pressKey(m, 'r')
	if cmd == nil {
		t.Fatal("expected non-nil cmd for 'r' at focusGroups")
	}
	msg := cmd()
	intent, ok := msg.(RetryGroupIntent)
	if !ok {
		t.Fatalf("expected RetryGroupIntent, got %T", msg)
	}
	if intent.Group != "teams" {
		t.Errorf("expected Group=teams, got %q", intent.Group)
	}
}

func TestRetryKeyEmitsRetryTestIntent(t *testing.T) {
	m := newTestModel()
	m.section = sectionGroups
	m.groups = []engine.Group{
		{Name: "repositories", Tests: []string{"TestA", "TestB"}},
	}
	m.groupCursor = 0
	m.focus = focusTests
	m.testCursor = 1

	_, cmd := pressKey(m, 'r')
	if cmd == nil {
		t.Fatal("expected non-nil cmd for 'r' at focusTests")
	}
	msg := cmd()
	intent, ok := msg.(RetryTestIntent)
	if !ok {
		t.Fatalf("expected RetryTestIntent, got %T", msg)
	}
	if intent.Test != "TestB" {
		t.Errorf("expected Test=TestB, got %q", intent.Test)
	}
}

func TestRetryAllEmitsRetryAllIntent(t *testing.T) {
	m := newTestModel()
	m.section = sectionGroups
	m.groups = []engine.Group{
		{Name: "repositories", Tests: []string{"TestA"}},
	}
	m.focus = focusGroups

	_, cmd := pressKey(m, 'R')
	if cmd == nil {
		t.Fatal("expected non-nil cmd for 'R'")
	}
	msg := cmd()
	if _, ok := msg.(RetryAllIntent); !ok {
		t.Fatalf("expected RetryAllIntent, got %T", msg)
	}
}

func TestRunSectionRetryKeyEmitsSelectedGroupIntent(t *testing.T) {
	m := newTestModel()
	m.section = sectionRun
	m.groups = []engine.Group{
		{Name: "repositories", Tests: []string{"TestA"}},
		{Name: "teams", Tests: []string{"TestB"}},
	}
	m.groupCursor = 1

	_, cmd := pressKey(m, 'r')
	if cmd == nil {
		t.Fatal("expected non-nil cmd for 'r' in Run")
	}
	intent, ok := cmd().(RetryGroupIntent)
	if !ok {
		t.Fatalf("expected RetryGroupIntent, got %T", cmd())
	}
	if intent.Group != "teams" {
		t.Fatalf("expected Group=teams, got %q", intent.Group)
	}
}

func TestRunSectionRetryAllKeyEmitsRetryAllIntent(t *testing.T) {
	m := newTestModel()
	m.section = sectionRun
	m.groups = []engine.Group{
		{Name: "repositories", Tests: []string{"TestA"}},
	}

	_, cmd := pressKey(m, 'R')
	if cmd == nil {
		t.Fatal("expected non-nil cmd for 'R' in Run")
	}
	if _, ok := cmd().(RetryAllIntent); !ok {
		t.Fatalf("expected RetryAllIntent, got %T", cmd())
	}
}

func TestRetryWithNoSelectionIsNoCmd(t *testing.T) {
	// Empty groups: 'r' and 'R' must not panic and must return nil cmd.
	m := newTestModel()
	m.section = sectionGroups
	m.groups = []engine.Group{}
	m.focus = focusGroups

	_, cmd := pressKey(m, 'r')
	if cmd != nil {
		t.Fatal("expected nil cmd for 'r' with no groups")
	}

	_, cmd = pressKey(m, 'R')
	if cmd != nil {
		t.Fatal("expected nil cmd for 'R' with no groups")
	}

	m.section = sectionRun
	for _, retryKey := range []rune{'r', 'R'} {
		t.Run("Run/"+string(retryKey), func(t *testing.T) {
			_, runCmd := pressKey(m, retryKey)
			if runCmd != nil {
				t.Fatalf("expected nil cmd for %q in Run with no groups", retryKey)
			}
		})
	}

	// Empty tests in focusTests: 'r' must return nil cmd, no panic.
	m2 := newTestModel()
	m2.section = sectionGroups
	m2.groups = []engine.Group{{Name: "g", Tests: []string{}}}
	m2.groupCursor = 0
	m2.focus = focusTests

	_, cmd = pressKey(m2, 'r')
	if cmd != nil {
		t.Fatal("expected nil cmd for 'r' with empty test list")
	}
}

// TestExecDispatch verifies that intent messages are routed through the
// injected executor seam and the returned cmd is non-nil.
func TestExecDispatch(t *testing.T) {
	intents := []tea.Msg{
		RetryGroupIntent{Group: "repositories"},
		RetryTestIntent{Test: "TestAccX"},
		RetryAllIntent{},
		ResumeIntent{},
		RefreshTriageIntent{},
		SyncKnownIssuesIntent{},
		PreviewFileIssueIntent{Fingerprint: "sha256:" + strings.Repeat("a", 64)},
		ExportReportIntent{},
		ListOrphansIntent{},
		ConfirmSweepIntent{Owner: "org", Phrase: "SWEEP org", Resources: []provider.Resource{{Kind: "repository", Name: "tf-acc"}}},
		ConfirmFileIssueIntent{Fingerprint: "sha256:" + strings.Repeat("a", 64), Phrase: "FILE aaaaaaaaaaaaaaaa"},
	}

	for _, want := range intents {
		want := want
		var got tea.Msg
		stub := func(msg tea.Msg) tea.Cmd {
			got = msg
			return func() tea.Msg { return nil }
		}
		m := New("github", "organization", "org", true).WithExec(stub)
		_, cmd := m.Update(want)
		if cmd == nil {
			t.Fatalf("%T: expected non-nil cmd from exec seam", want)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("%T: exec saw %v, want %v", want, got, want)
		}
	}
}

// TestExecNilIsNoOp verifies that a model without an injected executor safely
// ignores all intent messages without panicking.
func TestExecNilIsNoOp(t *testing.T) {
	m := New("github", "organization", "org", true) // no WithExec
	intents := []tea.Msg{
		RetryGroupIntent{Group: "g"},
		RetryTestIntent{Test: "T"},
		RetryAllIntent{},
		ResumeIntent{},
		RefreshTriageIntent{},
		SyncKnownIssuesIntent{},
		PreviewFileIssueIntent{Fingerprint: "sha256:" + strings.Repeat("a", 64)},
		ExportReportIntent{},
		ListOrphansIntent{},
		ConfirmSweepIntent{Owner: "org", Phrase: "SWEEP org", Resources: []provider.Resource{{Kind: "repository", Name: "tf-acc"}}},
		ConfirmFileIssueIntent{Fingerprint: "sha256:" + strings.Repeat("a", 64), Phrase: "FILE aaaaaaaaaaaaaaaa"},
	}
	for _, intent := range intents {
		_, cmd := m.Update(intent)
		if cmd != nil {
			t.Errorf("%T: expected nil cmd with nil exec, got non-nil", intent)
		}
	}
}

func TestUpdateTriageLoadedMsgSetsStateAndClampsCursor(t *testing.T) {
	m := newTestModel()
	m.section = sectionTriage
	m.triageCursor = 9
	next, _ := m.Update(TriageLoadedMsg{Failures: sampleFailures()[:2], CacheAvailable: true})
	got := next.(Model)
	if len(got.triageFailures) != 2 {
		t.Fatalf("triageFailures = %d, want 2", len(got.triageFailures))
	}
	if got.triageCursor != 1 {
		t.Fatalf("triageCursor = %d, want clamped to 1", got.triageCursor)
	}
	if !got.triageCacheAvailable {
		t.Fatal("triageCacheAvailable = false, want true")
	}
}

func TestUpdateTriageLoadedMsgCompletesActiveRefresh(t *testing.T) {
	m := newTestModel()
	m.operation = "triage"
	m.status = "triage…"

	next, _ := m.Update(TriageLoadedMsg{})
	got := next.(Model)

	if got.operation != "" {
		t.Fatalf("operation = %q, want empty", got.operation)
	}
	if got.status != "triage refreshed" {
		t.Fatalf("status = %q, want %q", got.status, "triage refreshed")
	}
}

func TestCompletionClearsOnlyMatchingOperationError(t *testing.T) {
	t.Run("triage completion keeps a different operation error", func(t *testing.T) {
		m := newTestModel()
		next, _ := m.Update(OperationErrMsg{Op: "known-issues-cache", Err: errors.New("known issue cache failed")})
		next, _ = next.(Model).Update(TriageLoadedMsg{})
		got := next.(Model)

		if got.err == nil || got.err.Error() != "known issue cache failed" {
			t.Fatalf("TriageLoadedMsg erased unrelated error: %v", got.err)
		}
		if !strings.Contains(got.status, "known-issues-cache: known issue cache failed") {
			t.Fatalf("TriageLoadedMsg de-styled unrelated error status: %q", got.status)
		}
	})

	t.Run("triage completion clears its own stale error and status", func(t *testing.T) {
		m := newTestModel()
		m.operation = "triage"
		m.err = errors.New("stale triage error")
		m.errOperation = "triage"
		m.status = "triage: stale triage error"
		next, _ := m.Update(TriageLoadedMsg{})
		got := next.(Model)

		if got.operation != "" || got.err != nil || got.status != "triage refreshed" {
			t.Fatalf("TriageLoadedMsg = operation:%q err:%v status:%q, want cleared operation/error and refreshed status", got.operation, got.err, got.status)
		}
	})

	t.Run("run start clears an unrelated stale error", func(t *testing.T) {
		m := newTestModel()
		next, _ := m.Update(OperationErrMsg{Op: "export", Err: errors.New("export failed")})
		next, _ = next.(Model).Update(RunStartedMsg{})
		if got := next.(Model); got.err != nil {
			t.Fatalf("RunStartedMsg left stale error: %v", got.err)
		}
	})

	t.Run("clean run completion clears an unrelated stale error", func(t *testing.T) {
		m := newTestModel()
		next, _ := m.Update(OperationErrMsg{Op: "export", Err: errors.New("export failed")})
		next, _ = next.(Model).Update(RunDoneMsg{})
		if got := next.(Model); got.err != nil {
			t.Fatalf("RunDoneMsg left stale error: %v", got.err)
		}
	})

	t.Run("post-completion run error remains visible", func(t *testing.T) {
		m := newTestModel()
		next, _ := m.Update(RunDoneMsg{})
		next, _ = next.(Model).Update(OperationErrMsg{Op: "run", Err: errors.New("run failed")})
		if got := next.(Model); got.err == nil || got.err.Error() != "run failed" {
			t.Fatalf("post-RunDone run error was hidden: %v", got.err)
		}
	})
}

func TestRunDonePreservesTerminalSuiteFailure(t *testing.T) {
	tests := []struct {
		name   string
		result engine.RunResult
		want   string
	}{
		{
			name:   "build failure",
			result: engine.RunResult{BuildFailed: true, BuildOutput: []string{"redacted compiler output"}},
			want:   "build failed before tests ran",
		},
		{
			name:   "pre-run failure",
			result: engine.RunResult{PreRunFailed: true, PreRunOutput: []string{"redacted pre-run output"}},
			want:   "pre-run failed before tests ran",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := newTestModel()
			next, _ := m.Update(RunStartedMsg{Label: "repositories"})
			next, _ = next.(Model).Update(RunDoneMsg{Result: tt.result})
			got := next.(Model)

			if got.running {
				t.Fatalf("terminal suite failure left the run active")
			}
			if got.err == nil || got.errOperation != "run" || !got.errIsOperationFailure {
				t.Fatalf("terminal suite failure = err:%v owner:%q failure:%t, want run-owned failure", got.err, got.errOperation, got.errIsOperationFailure)
			}
			if got.err.Error() != tt.want {
				t.Fatalf("terminal suite failure = %q, want %q", got.err, tt.want)
			}
			if state, _, _ := got.runVisualState(got.runSummary()); state != "Failed" {
				t.Fatalf("terminal suite failure state = %q, want Failed", state)
			}
		})
	}
}

func TestErrorPresentationOwnership(t *testing.T) {
	t.Run("rejected duplicate run does not fail the active run", func(t *testing.T) {
		m := newTestModel()
		next, _ := m.Update(OperationStartedMsg{Name: "run"})
		next, _ = next.(Model).Update(OperationRejectedMsg{Op: "run", Err: errors.New("another operation is in progress")})
		got := next.(Model)

		if got.err == nil || got.errOperation != "run" {
			t.Fatalf("rejection = err:%v owner:%q, want danger presentation owned by active run", got.err, got.errOperation)
		}
		if state, _, _ := got.runVisualState(got.runSummary()); state == "Failed" {
			t.Fatalf("duplicate rejection must not fail the active Run view")
		}
		if activity := got.renderRunActionsBody(80); strings.Contains(activity, "Last activity: failed") {
			t.Fatalf("duplicate rejection must not render a failed Run activity:\n%s", activity)
		}
	})

	t.Run("rejected run belongs to active export until export succeeds", func(t *testing.T) {
		m := newTestModel()
		next, _ := m.Update(OperationStartedMsg{Name: "export"})
		next, _ = next.(Model).Update(OperationRejectedMsg{Op: "run", Err: errors.New("run is blocked while report exports")})
		got := next.(Model)

		if got.err == nil || got.errOperation != "export" {
			t.Fatalf("rejection = err:%v owner:%q, want danger presentation owned by export", got.err, got.errOperation)
		}
		state, _, _ := got.runVisualState(got.runSummary())
		if state == "Failed" {
			t.Fatalf("rejected run must not fail the Run view")
		}
		status := stripANSI(got.renderStatus(80))
		if !strings.Contains(status, "run is blocked") || !strings.HasSuffix(status, "export") {
			t.Fatalf("rejection status should stay danger-visible with active export on the right:\n%s", status)
		}
		for _, width := range []int{1, 2, 4, 8, 11} {
			assertRenderedWidth(t, got.renderStatus(width), width)
		}

		next, _ = got.Update(ReportExportDoneMsg{MarkdownPath: "report.md"})
		got = next.(Model)
		if got.err != nil || got.errOperation != "" {
			t.Fatalf("successful export left rejection presentation: err:%v owner:%q", got.err, got.errOperation)
		}
		if status := stripANSI(got.renderStatus(80)); !strings.Contains(status, "report exported") || strings.Contains(status, "run is blocked") {
			t.Fatalf("successful export did not supersede rejection status:\n%s", status)
		}
	})

	t.Run("ownerless error survives unrelated completion until an explicit reset", func(t *testing.T) {
		const errText = "background status failed"
		ownerlessError := func() Model {
			m := newTestModel()
			next, _ := m.Update(OperationStartedMsg{Name: "export"})
			next, _ = next.(Model).Update(ErrMsg{Err: errors.New(errText)})
			next, _ = next.(Model).Update(ReportExportDoneMsg{MarkdownPath: "report.md"})
			return next.(Model)
		}

		got := ownerlessError()
		if got.err == nil || got.err.Error() != errText || got.errOperation != "" {
			t.Fatalf("export completion cleared ownerless error: err:%v owner:%q", got.err, got.errOperation)
		}
		if status := stripANSI(got.renderStatus(80)); !strings.Contains(status, errText) || strings.Contains(status, "report exported") {
			t.Fatalf("ownerless error was not retained as the danger status:\n%s", status)
		}

		next, _ := got.Update(OperationStartedMsg{Name: "triage"})
		if got := next.(Model); got.err != nil {
			t.Fatalf("OperationStartedMsg did not clear ownerless error: %v", got.err)
		}

		next, _ = ownerlessError().Update(RunStartedMsg{Label: "repositories"})
		if got := next.(Model); got.err != nil {
			t.Fatalf("RunStartedMsg did not clear ownerless error: %v", got.err)
		}
	})

	t.Run("only real run errors fail Run and run resets clear them", func(t *testing.T) {
		m := newTestModel()
		next, _ := m.Update(OperationErrMsg{Op: "run", Err: errors.New("run failed")})
		got := next.(Model)
		if state, _, _ := got.runVisualState(got.runSummary()); state != "Failed" {
			t.Fatalf("run-owned error state = %q, want Failed", state)
		}

		next, _ = got.Update(RunStartedMsg{Label: "repositories"})
		got = next.(Model)
		if got.err != nil {
			t.Fatalf("RunStartedMsg did not clear real run error: %v", got.err)
		}
		if state, _, _ := got.runVisualState(got.runSummary()); state == "Failed" {
			t.Fatalf("RunStartedMsg left the Run view failed")
		}

		next, _ = got.Update(OperationErrMsg{Op: "run", Err: errors.New("run failed")})
		next, _ = next.(Model).Update(RunDoneMsg{})
		got = next.(Model)
		if got.err != nil {
			t.Fatalf("clean RunDoneMsg did not clear run error: %v", got.err)
		}

		next, _ = got.Update(OperationErrMsg{Op: "run", Err: errors.New("late run failed")})
		got = next.(Model)
		if state, _, _ := got.runVisualState(got.runSummary()); state != "Failed" {
			t.Fatalf("late run error state = %q, want Failed", state)
		}
	})

	t.Run("real run error stops an active run and renders failed activity", func(t *testing.T) {
		m := newTestModel()
		next, _ := m.Update(RunStartedMsg{Label: "repositories"})
		next, _ = next.(Model).Update(OperationErrMsg{Op: "run", Err: errors.New("run failed")})
		got := next.(Model)

		if got.running {
			t.Fatalf("terminal run error left the run active")
		}
		if activity := got.renderRunActionsBody(80); !strings.Contains(activity, "Last activity: failed - run failed") {
			t.Fatalf("terminal run error did not render failed activity:\n%s", activity)
		}
	})

	t.Run("only matching successful completion clears owned error", func(t *testing.T) {
		m := newTestModel()
		next, _ := m.Update(OperationErrMsg{Op: "export", Err: errors.New("export failed")})
		next, _ = next.(Model).Update(KnownIssueSyncDoneMsg{})
		got := next.(Model)
		if got.err == nil || got.errOperation != "export" {
			t.Fatalf("unrelated completion cleared export error: err:%v owner:%q", got.err, got.errOperation)
		}
		if status := stripANSI(got.renderStatus(80)); !strings.Contains(status, "export failed") || strings.Contains(status, "known issues synced") {
			t.Fatalf("unrelated success replaced owned danger status:\n%s", status)
		}

		next, _ = got.Update(ReportExportDoneMsg{MarkdownPath: "report.md"})
		if got := next.(Model); got.err != nil || got.errOperation != "" {
			t.Fatalf("matching export completion left error: err:%v owner:%q", got.err, got.errOperation)
		}
	})
}

// TestKnownIssueSyncDoneMsgUpdatesCacheAvailable proves the reducer transitions
// m.triageCacheAvailable from an initially unavailable cache to available
// using the producer-reported CacheAvailable value on the completion message,
// while preserving the count/path/failure overlay behavior.
func TestKnownIssueSyncDoneMsgUpdatesCacheAvailable(t *testing.T) {
	m := newTestModel()
	m.triageCacheAvailable = false // starts unavailable, e.g. before first sync

	refreshed := sampleFailures()[:2]
	next, _ := m.Update(KnownIssueSyncDoneMsg{
		Count:          3,
		CachePath:      "/tmp/known-issues.yaml",
		Failures:       refreshed,
		CacheAvailable: true,
	})
	got := next.(Model)

	if !got.triageCacheAvailable {
		t.Fatal("triageCacheAvailable = false, want true after a successful sync")
	}
	if got.status != "known issues synced" {
		t.Fatalf("status = %q, want %q", got.status, "known issues synced")
	}
	joined := strings.Join(got.resultLines, "\n")
	if !strings.Contains(joined, "count: 3") {
		t.Fatalf("resultLines missing count line: %v", got.resultLines)
	}
	if !strings.Contains(joined, "cache: /tmp/known-issues.yaml") {
		t.Fatalf("resultLines missing cache path line: %v", got.resultLines)
	}
	if !strings.Contains(joined, fmt.Sprintf("failures: %d", len(refreshed))) {
		t.Fatalf("resultLines missing failure count line: %v", got.resultLines)
	}
}

// TestKnownIssueSyncDoneMsgCacheUnavailablePreserved proves a sync that
// reports no cache leaves m.triageCacheAvailable false (error/offline path).
func TestKnownIssueSyncDoneMsgCacheUnavailablePreserved(t *testing.T) {
	m := newTestModel()
	m.triageCacheAvailable = true

	next, _ := m.Update(KnownIssueSyncDoneMsg{CacheAvailable: false})
	got := next.(Model)

	if got.triageCacheAvailable {
		t.Fatal("triageCacheAvailable = true, want false when the message reports no cache")
	}
}

func TestUpdateTriageRefreshKeyOnlyInTriageSection(t *testing.T) {
	m := newTestModel()
	m.section = sectionRun
	if _, cmd := pressKey(m, 't'); cmd != nil {
		t.Fatal("t outside triage emitted a command")
	}
	m.section = sectionTriage
	_, cmd := pressKey(m, 't')
	if cmd == nil {
		t.Fatal("t in triage should emit RefreshTriageIntent")
	}
	if _, ok := cmd().(RefreshTriageIntent); !ok {
		t.Fatalf("t emitted %T, want RefreshTriageIntent", cmd())
	}
}

func TestUpdateKnownIssuesSyncKeyOnlyInTriageSection(t *testing.T) {
	m := newTestModel()
	m.section = sectionGroups
	if _, cmd := pressKey(m, 'K'); cmd != nil {
		t.Fatal("K outside triage emitted a command")
	}
	m.section = sectionTriage
	_, cmd := pressKey(m, 'K')
	if cmd == nil {
		t.Fatal("K in triage should emit SyncKnownIssuesIntent")
	}
	if _, ok := cmd().(SyncKnownIssuesIntent); !ok {
		t.Fatalf("K emitted %T, want SyncKnownIssuesIntent", cmd())
	}
}

func TestUpdateTriageFileIssueKeyOnlyForEligibleSelection(t *testing.T) {
	m := newTestModel()
	m.section = sectionTriage
	m.triageFailures = sampleFailures()
	m.triageDetailActive = true

	m.triageCursor = 0
	if _, cmd := pressKey(m, 'i'); cmd != nil {
		t.Fatal("i should not emit for known issue selection")
	}

	m.triageCursor = 2
	_, cmd := pressKey(m, 'i')
	if cmd == nil {
		t.Fatal("i should emit for eligible selected failure")
	}
	intent, ok := cmd().(PreviewFileIssueIntent)
	if !ok {
		t.Fatalf("i emitted %T, want PreviewFileIssueIntent", cmd())
	}
	if intent.Fingerprint != sampleFailures()[2].Fingerprint {
		t.Fatalf("fingerprint = %q, want %q", intent.Fingerprint, sampleFailures()[2].Fingerprint)
	}
}

// TestUpdateTriageFileIssueKeyRequiresDetailView proves 'i' emits
// PreviewFileIssueIntent only when the Triage detail view is active for the
// selected eligible failure; pressing 'i' on the Triage list (detail not
// active) must emit no command and start no operation, even though the
// same failure/cursor would be eligible from detail.
func TestUpdateTriageFileIssueKeyRequiresDetailView(t *testing.T) {
	m := newTestModel()
	m.section = sectionTriage
	m.triageFailures = sampleFailures()
	m.triageCursor = 2 // eligible failure (see sampleFailures/TestAccReal)

	// List view: detail not active. 'i' must be a no-op.
	m.triageDetailActive = false
	got, cmd := pressKey(m, 'i')
	if cmd != nil {
		t.Fatalf("i on the Triage list emitted %T, want no command", cmd())
	}
	if got.operation != "" {
		t.Fatalf("i on the Triage list started operation %q, want none", got.operation)
	}

	// Detail view for the same selection: 'i' must emit the preview intent.
	m.triageDetailActive = true
	_, cmd = pressKey(m, 'i')
	if cmd == nil {
		t.Fatal("i in Triage detail should emit PreviewFileIssueIntent")
	}
	if _, ok := cmd().(PreviewFileIssueIntent); !ok {
		t.Fatalf("i emitted %T, want PreviewFileIssueIntent", cmd())
	}
}

func TestUpdateTriageActionKeysBlockedDuringOperation(t *testing.T) {
	m := newTestModel()
	m.section = sectionTriage
	m.operation = "refresh triage"
	m.triageFailures = sampleFailures()
	m.triageCursor = 2
	for _, r := range []rune{'t', 'K', 'i'} {
		if _, cmd := pressKey(m, r); cmd != nil {
			t.Fatalf("%c emitted command while operation was active", r)
		}
	}
}

func TestUpdateFailuresFirstReordersGroupsAndTests(t *testing.T) {
	m := newTestModel()
	m.section = sectionGroups
	next, _ := m.Update(GroupsMsg{Groups: []engine.Group{
		{Name: "passing", Tests: []string{"TestPass"}},
		{Name: "mixed", Tests: []string{"TestSkip", "TestFail", "TestRunning", "TestNotRun"}},
		{Name: "empty", Tests: []string{"TestOtherNotRun"}},
	}})
	m = next.(Model)
	for _, result := range []engine.TestResult{
		{Name: "TestPass", Status: provider.StatusPass},
		{Name: "TestSkip", Status: provider.StatusSkip},
		{Name: "TestFail", Status: provider.StatusFail},
		{Name: "TestRunning", Status: provider.StatusRunning},
	} {
		next, _ = m.Update(TestUpdateMsg{Result: result})
		m = next.(Model)
	}

	m, _ = pressKey(m, 'f')
	if !m.failuresFirst {
		t.Fatal("failuresFirst = false after f")
	}
	if got := []string{m.groups[0].Name, m.groups[1].Name, m.groups[2].Name}; strings.Join(got, ",") != "mixed,passing,empty" {
		t.Fatalf("group order = %v, want mixed,passing,empty", got)
	}
	if got := strings.Join(m.groups[0].Tests, ","); got != "TestFail,TestRunning,TestSkip,TestNotRun" {
		t.Fatalf("test order = %s, want failures/running before skip/not-run", got)
	}
}

func TestUpdateShowAllRestoresDiscoveryOrder(t *testing.T) {
	discovery := []engine.Group{
		{Name: "passing", Tests: []string{"TestPass"}},
		{Name: "mixed", Tests: []string{"TestSkip", "TestFail", "TestRunning", "TestNotRun"}},
		{Name: "empty", Tests: []string{"TestOtherNotRun"}},
	}
	m := newTestModel()
	m.section = sectionGroups
	next, _ := m.Update(GroupsMsg{Groups: discovery})
	m = next.(Model)
	for _, result := range []engine.TestResult{
		{Name: "TestPass", Status: provider.StatusPass},
		{Name: "TestSkip", Status: provider.StatusSkip},
		{Name: "TestFail", Status: provider.StatusFail},
		{Name: "TestRunning", Status: provider.StatusRunning},
	} {
		next, _ = m.Update(TestUpdateMsg{Result: result})
		m = next.(Model)
	}
	m, _ = pressKey(m, 'f')
	m, _ = pressKey(m, 'a')
	if m.failuresFirst {
		t.Fatal("failuresFirst = true after a")
	}
	for i := range discovery {
		if m.groups[i].Name != discovery[i].Name || strings.Join(m.groups[i].Tests, ",") != strings.Join(discovery[i].Tests, ",") {
			t.Fatalf("groups restored to %+v, want %+v", m.groups, discovery)
		}
	}
}

// TestResumePromptMsgSetsState verifies that ResumePromptMsg populates the
// resumePrompt flag and all associated fields.
func TestResumePromptMsgSetsState(t *testing.T) {
	m := newTestModel()
	when := time.Date(2026, 1, 15, 10, 30, 0, 0, time.UTC)
	m2, _ := m.Update(ResumePromptMsg{When: when, Failed: 2, NotRun: 5})
	got := m2.(Model)
	if !got.resumePrompt {
		t.Fatal("expected resumePrompt=true after ResumePromptMsg")
	}
	if !got.resumeWhen.Equal(when) {
		t.Errorf("resumeWhen: want %v, got %v", when, got.resumeWhen)
	}
	if got.resumeFailed != 2 {
		t.Errorf("resumeFailed: want 2, got %d", got.resumeFailed)
	}
	if got.resumeNotRun != 5 {
		t.Errorf("resumeNotRun: want 5, got %d", got.resumeNotRun)
	}
}

// TestResumePromptAccept verifies that DrillIn (enter/l) and RetryAll (R)
// while the resume prompt is visible emit ResumeIntent and clear the prompt.
func TestResumePromptAccept(t *testing.T) {
	when := time.Date(2026, 1, 15, 10, 30, 0, 0, time.UTC)

	type sender struct {
		name string
		send func(Model) (Model, tea.Cmd)
	}
	senders := []sender{
		{"enter", func(m Model) (Model, tea.Cmd) {
			m2, cmd := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
			return m2.(Model), cmd
		}},
		{"l", func(m Model) (Model, tea.Cmd) { return pressKey(m, 'l') }},
		{"R", func(m Model) (Model, tea.Cmd) { return pressKey(m, 'R') }},
	}

	for _, s := range senders {
		s := s
		m := newTestModel()
		m.resumePrompt = true
		m.resumeWhen = when
		m.section = sectionGroups
		m.groups = []engine.Group{{Name: "g", Tests: []string{"T"}}}

		got, cmd := s.send(m)
		if got.resumePrompt {
			t.Errorf("%s: resumePrompt still true after accept", s.name)
		}
		if cmd == nil {
			t.Fatalf("%s: expected non-nil cmd for resume accept", s.name)
		}
		if msg := cmd(); msg != (ResumeIntent{}) {
			t.Errorf("%s: expected ResumeIntent, got %T", s.name, msg)
		}
	}
}

// TestResumePromptDismiss verifies that Back (esc/h) clears the resume prompt
// without emitting a cmd, and that unrelated keys are consumed (no nav change).
func TestResumePromptDismiss(t *testing.T) {
	when := time.Date(2026, 1, 15, 10, 30, 0, 0, time.UTC)

	// esc dismisses
	m := newTestModel()
	m.resumePrompt = true
	m.resumeWhen = when
	m2, cmd := m.Update(tea.KeyMsg{Type: tea.KeyEsc})
	got := m2.(Model)
	if got.resumePrompt {
		t.Error("esc: expected resumePrompt=false after dismiss")
	}
	if cmd != nil {
		t.Errorf("esc: expected nil cmd, got non-nil")
	}

	// h (also Back binding) dismisses and does NOT forward to navigation
	m = newTestModel()
	m.resumePrompt = true
	m.resumeWhen = when
	m.section = sectionGroups
	m.focus = focusTests // would normally navigate back to focusGroups
	got, cmd = pressKey(m, 'h')
	if got.resumePrompt {
		t.Error("h: expected resumePrompt=false after dismiss")
	}
	if cmd != nil {
		t.Errorf("h: expected nil cmd, got non-nil")
	}
	if got.focus != focusTests {
		t.Errorf("h: focus leaked to nav; expected focusTests, got %d", got.focus)
	}

	// unrelated key (j) while prompt is up is consumed; no navigation change
	m = newTestModel()
	m.resumePrompt = true
	m.resumeWhen = when
	m.section = sectionGroups
	m.groups = []engine.Group{{Name: "g1"}, {Name: "g2"}}
	m.groupCursor = 0
	got, cmd = pressKey(m, 'j')
	if !got.resumePrompt {
		t.Error("j: resumePrompt was cleared by unrelated key")
	}
	if got.groupCursor != 0 {
		t.Error("j: navigation leaked through prompt; groupCursor moved")
	}
	if cmd != nil {
		t.Errorf("j: expected nil cmd, got non-nil")
	}
}

// TestRunLifecycle verifies the RunStartedMsg/RunDoneMsg state transitions.
func TestRunLifecycle(t *testing.T) {
	m := newTestModel()

	m2, _ := m.Update(RunStartedMsg{Total: 10})
	got := m2.(Model)
	if !got.running {
		t.Fatal("expected running=true after RunStartedMsg")
	}
	if !strings.Contains(got.status, "running") {
		t.Errorf("expected status to contain 'running', got %q", got.status)
	}

	m3, _ := got.Update(RunDoneMsg{})
	got2 := m3.(Model)
	if got2.running {
		t.Fatal("expected running=false after RunDoneMsg")
	}
}

// TestRetryRunsSelectedGroupAfterNavigation proves that pressing 'r' on the
// Groups list runs the group under the cursor after the user navigates with
// individual 'j' presses - not always the first group. This is the
// deterministic counterpart to a PTY walkthrough.
func TestRetryRunsSelectedGroupAfterNavigation(t *testing.T) {
	m := newTestModel()
	m2, _ := m.Update(GroupsMsg{Groups: []engine.Group{
		{Name: "actions"}, {Name: "apps"}, {Name: "branches"}, {Name: "codespaces"},
		{Name: "dependabot"}, {Name: "enterprise"}, {Name: "ip-ranges"},
	}})
	mm := m2.(Model)
	// Tab from Preflight to Groups.
	m3, _ := mm.Update(tea.KeyMsg{Type: tea.KeyTab})
	mm = m3.(Model)
	// Six individual 'j' presses → cursor on index 6 (ip-ranges).
	for i := 0; i < 6; i++ {
		mm, _ = pressKey(mm, 'j')
	}
	if mm.groupCursor != 6 {
		t.Fatalf("expected groupCursor=6 (ip-ranges), got %d", mm.groupCursor)
	}
	// 'r' should emit a RetryGroupIntent for the SELECTED group.
	_, cmd := pressKey(mm, 'r')
	if cmd == nil {
		t.Fatal("expected a command from 'r', got nil")
	}
	intent, ok := cmd().(RetryGroupIntent)
	if !ok {
		t.Fatalf("expected RetryGroupIntent, got %T", cmd())
	}
	if intent.Group != "ip-ranges" {
		t.Fatalf("expected RetryGroupIntent{Group:\"ip-ranges\"}, got %q", intent.Group)
	}
}

// TestCoalescedRuneBurstDoesNotNavigate documents that a single KeyMsg carrying
// multiple runes (as produced when several keystrokes arrive in one terminal
// read, e.g. a scripted burst) matches no single-key binding and therefore does
// not move the cursor. Real interactive typing delivers one KeyMsg per key.
func TestCoalescedRuneBurstDoesNotNavigate(t *testing.T) {
	m := newTestModel()
	m2, _ := m.Update(GroupsMsg{Groups: []engine.Group{
		{Name: "actions"}, {Name: "ip-ranges"},
	}})
	mm := m2.(Model)
	m3, _ := mm.Update(tea.KeyMsg{Type: tea.KeyTab})
	mm = m3.(Model)
	// One KeyMsg with six 'j' runes (a coalesced burst).
	m4, _ := mm.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("jjjjjj")})
	mm = m4.(Model)
	if mm.groupCursor != 0 {
		t.Fatalf("coalesced burst should not navigate; groupCursor=%d", mm.groupCursor)
	}
}

// TestRunStartedStartsSpinner verifies RunStartedMsg resets the Pulsar frame to
// 0 and returns a non-nil command (the tick that drives the animation).
func TestRunStartedStartsSpinner(t *testing.T) {
	m := newTestModel()
	m.spinnerFrame = 7 // stale value from a previous run
	m2, cmd := m.Update(RunStartedMsg{Total: 5})
	got := m2.(Model)
	if !got.running {
		t.Fatal("expected running=true after RunStartedMsg")
	}
	if got.spinnerFrame != 0 {
		t.Errorf("expected spinnerFrame reset to 0, got %d", got.spinnerFrame)
	}
	if cmd == nil {
		t.Fatal("expected a non-nil spinner tick command after RunStartedMsg")
	}
}

// TestSpinnerTickAdvancesWhileRunning verifies a spinnerTickMsg advances the
// frame and re-arms the tick while a run is in progress.
func TestSpinnerTickAdvancesWhileRunning(t *testing.T) {
	m := newTestModel()
	m.running = true
	m.spinnerFrame = 2
	m2, cmd := m.Update(spinnerTickMsg{})
	got := m2.(Model)
	if got.spinnerFrame != 3 {
		t.Errorf("expected spinnerFrame=3 after tick, got %d", got.spinnerFrame)
	}
	if cmd == nil {
		t.Fatal("expected the tick to re-arm with a non-nil command while running")
	}
}

// TestSpinnerTickStopsWhenNotRunning verifies the tick loop self-terminates once
// the run is done: no frame advance and no further command.
func TestSpinnerTickStopsWhenNotRunning(t *testing.T) {
	m := newTestModel()
	m.running = false
	m.spinnerFrame = 4
	m2, cmd := m.Update(spinnerTickMsg{})
	got := m2.(Model)
	if got.spinnerFrame != 4 {
		t.Errorf("expected spinnerFrame unchanged at 4 when not running, got %d", got.spinnerFrame)
	}
	if cmd != nil {
		t.Error("expected nil command when not running (tick loop stops)")
	}
}

// ── copy keys ─────────────────────────────────────────────────────────────────

// modelWithFailedTest returns a model drilled into the log pane for a test
// that has output, with a fake copyFn injected.
func modelWithFailedTest(copyCapture *string, copyErr error) (Model, string) {
	const testName = "TestAccFoo"
	const logLine = "failed: assertion error"

	var captured string
	copyFn := func(s string) error {
		captured = s
		if copyCapture != nil {
			*copyCapture = s
		}
		return copyErr
	}

	m := New("github", "organization", "org", true).WithClipboard(copyFn)
	m.section = sectionGroups
	m.groups = []engine.Group{
		{Name: "repos", Tests: []string{testName}},
	}
	m.groupCursor = 0
	m.results = []engine.TestResult{
		{Name: testName, Sub: "", Status: provider.StatusFail, Output: []string{logLine + "\n"}},
	}
	m.resultIndex = map[resultKey]int{{Name: testName, Sub: ""}: 0}
	_ = captured
	return m, logLine
}

func TestCopyLogCopiesOutput(t *testing.T) {
	var captured string
	m, logLine := modelWithFailedTest(&captured, nil)

	// Drill to focusTests then focusLog so selectedTestName() works.
	m, _ = pressKey(m, 'l') // focusGroups -> focusTests
	m, _ = pressKey(m, 'l') // focusTests -> focusLog

	m, _ = pressKey(m, 'y')

	if !strings.Contains(captured, logLine) {
		t.Errorf("copyFn captured %q; want it to contain %q", captured, logLine)
	}
	if m.status != "copied log" {
		t.Errorf("status = %q; want %q", m.status, "copied log")
	}
}

func TestCopyLogEmptyOutputIsNothingToCopy(t *testing.T) {
	var called bool
	m := New("github", "organization", "org", true).WithClipboard(func(s string) error {
		called = true
		return nil
	})
	m.section = sectionGroups
	m.groups = []engine.Group{{Name: "repos", Tests: []string{"TestAccFoo"}}}
	m.groupCursor = 0
	m.results = []engine.TestResult{
		{Name: "TestAccFoo", Sub: "", Status: provider.StatusFail, Output: nil},
	}
	m.resultIndex = map[resultKey]int{{Name: "TestAccFoo", Sub: ""}: 0}

	m, _ = pressKey(m, 'l') // focusTests
	m, _ = pressKey(m, 'l') // focusLog
	m, _ = pressKey(m, 'y')

	if called {
		t.Error("copyFn should not be called when there is no output")
	}
	if m.status != "nothing to copy" {
		t.Errorf("status = %q; want %q", m.status, "nothing to copy")
	}
}

func TestCopyCmdCopiesCommand(t *testing.T) {
	const wantCmd = "go test ./github/... -run '^(TestAccFoo)$' -v -count=1"
	var captured string

	m := New("github", "organization", "org", true).
		WithClipboard(func(s string) error { captured = s; return nil }).
		WithCmdFunc(func(test string) string { return wantCmd })

	m.section = sectionGroups
	m.groups = []engine.Group{{Name: "repos", Tests: []string{"TestAccFoo"}}}
	m.groupCursor = 0
	m.focus = focusTests

	m, _ = pressKey(m, 'c')

	if captured != wantCmd {
		t.Errorf("copyFn captured %q; want %q", captured, wantCmd)
	}
	if m.status != "copied cmd" {
		t.Errorf("status = %q; want %q", m.status, "copied cmd")
	}
}

func TestCopyLogNilCopyFnIsNoOp(t *testing.T) {
	// Model with no copyFn injected (default oscCopy overridden to nil for test).
	m := New("github", "organization", "org", true)
	m.copyFn = nil
	m.section = sectionGroups
	m.groups = []engine.Group{{Name: "repos", Tests: []string{"TestAccFoo"}}}
	m.groupCursor = 0
	m.results = []engine.TestResult{
		{Name: "TestAccFoo", Sub: "", Status: provider.StatusFail, Output: []string{"line\n"}},
	}
	m.resultIndex = map[resultKey]int{{Name: "TestAccFoo", Sub: ""}: 0}
	m, _ = pressKey(m, 'l')
	m, _ = pressKey(m, 'l')

	// Should not panic; status stays empty.
	m2, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'y'}})
	if got := m2.(Model).status; got != "" {
		t.Errorf("nil copyFn: status = %q; want empty", got)
	}
}

// ── focus-guard tests (Findings 4+5) ─────────────────────────────────────────

func TestCopyLogFocusGroupsSelectTestFirst(t *testing.T) {
	var called bool
	m, _ := modelWithFailedTest(nil, nil)
	m.copyFn = func(s string) error {
		called = true
		return nil
	}
	// focus is already focusGroups (default after modelWithFailedTest).
	if m.focus != focusGroups {
		t.Fatalf("precondition: expected focusGroups, got %d", m.focus)
	}

	m, _ = pressKey(m, 'y')

	if called {
		t.Error("copyFn must not be called when focus is on the Groups list")
	}
	if m.status != "select a test first" {
		t.Errorf("status = %q; want %q", m.status, "select a test first")
	}
}

func TestCopyCmdFocusGroupsSelectTestFirst(t *testing.T) {
	var called bool
	m, _ := modelWithFailedTest(nil, nil)
	m.copyFn = func(s string) error {
		called = true
		return nil
	}
	m.cmdFor = func(test string) string { return "go test ./... -run '" + test + "' -v -count=1" }
	if m.focus != focusGroups {
		t.Fatalf("precondition: expected focusGroups, got %d", m.focus)
	}

	m, _ = pressKey(m, 'c')

	if called {
		t.Error("copyFn must not be called when focus is on the Groups list")
	}
	if m.status != "select a test first" {
		t.Errorf("status = %q; want %q", m.status, "select a test first")
	}
}

func TestCopyLogFocusTestsCopiesOutput(t *testing.T) {
	var captured string
	m, logLine := modelWithFailedTest(&captured, nil)
	m.focus = focusTests

	m, _ = pressKey(m, 'y')

	if !strings.Contains(captured, logLine) {
		t.Errorf("copyFn captured %q; want it to contain %q", captured, logLine)
	}
	if m.status != "copied log" {
		t.Errorf("status = %q; want %q", m.status, "copied log")
	}
}

func TestCopyCmdFocusLogCopiesCommand(t *testing.T) {
	const wantCmd = "go test ./github/... -run '^(TestAccFoo)$' -v -count=1"
	var captured string

	m, _ := modelWithFailedTest(nil, nil)
	m.copyFn = func(s string) error { captured = s; return nil }
	m.cmdFor = func(_ string) string { return wantCmd }
	m.focus = focusLog

	m, _ = pressKey(m, 'c')

	if captured != wantCmd {
		t.Errorf("copyFn captured %q; want %q", captured, wantCmd)
	}
	if m.status != "copied cmd" {
		t.Errorf("status = %q; want %q", m.status, "copied cmd")
	}
}

// ── mode picker overlay ────────────────────────────────────────────────────────

func testModes() []provider.Mode {
	return []provider.Mode{
		{Name: "anonymous", Description: "No credentials; public endpoints only"},
		{Name: "individual", Description: "Personal account tests"},
		{Name: "organization", Description: "Organization tests"},
		{Name: "team", Description: "Team tests; adds external user fixtures"},
		{Name: "enterprise", Description: "Enterprise tests; adds GITHUB_ENTERPRISE_SLUG"},
	}
}

func TestModePickerOpenWithM(t *testing.T) {
	m := newTestModel().WithModes(testModes())
	got, _ := pressKey(m, 'm')
	if !got.pickerActive {
		t.Fatal("expected pickerActive=true after pressing m")
	}
}

func TestModePickerNoModesDoesNotOpen(t *testing.T) {
	m := newTestModel() // no WithModes
	got, _ := pressKey(m, 'm')
	if got.pickerActive {
		t.Fatal("expected pickerActive=false when no modes configured")
	}
}

func TestModePickerNavAndSelectEnterprise(t *testing.T) {
	var capturedMsg tea.Msg
	exec := func(msg tea.Msg) tea.Cmd {
		capturedMsg = msg
		return func() tea.Msg { return msg }
	}
	m := newTestModel().WithModes(testModes()).WithExec(exec)
	m, _ = pressKey(m, 'm')
	if !m.pickerActive {
		t.Fatal("expected picker open after m")
	}
	// Navigate down to enterprise (index 4).
	for i := 0; i < 4; i++ {
		m, _ = pressKey(m, 'j')
	}
	if m.pickerCursor != 4 {
		t.Fatalf("expected pickerCursor=4, got %d", m.pickerCursor)
	}
	// Select with enter.
	got, cmd := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	m = got.(Model)
	if m.pickerActive {
		t.Fatal("expected pickerActive=false after selecting")
	}
	if m.mode != "enterprise" {
		t.Errorf("expected mode=enterprise (optimistic), got %q", m.mode)
	}
	if cmd == nil {
		t.Fatal("expected non-nil cmd after picker selection")
	}
	// The exec receives SwitchModeIntent.
	if intent, ok := capturedMsg.(SwitchModeIntent); !ok || intent.Mode != "enterprise" {
		t.Errorf("expected SwitchModeIntent{Mode:enterprise}, got %T %v", capturedMsg, capturedMsg)
	}
}

func TestModePickerCursorStartsAtCurrentMode(t *testing.T) {
	m := New("github", "organization", "org", true).WithModes(testModes())
	got, _ := pressKey(m, 'm')
	if !got.pickerActive {
		t.Fatal("expected picker open")
	}
	// organization is index 2 in testModes()
	if got.pickerCursor != 2 {
		t.Errorf("expected pickerCursor=2 (organization), got %d", got.pickerCursor)
	}
}

func TestModePickerEscClosesOverlay(t *testing.T) {
	m := newTestModel().WithModes(testModes())
	m, _ = pressKey(m, 'm')
	got, cmd := m.Update(tea.KeyMsg{Type: tea.KeyEsc})
	m = got.(Model)
	if m.pickerActive {
		t.Fatal("esc: expected pickerActive=false")
	}
	// esc must not quit the app
	if cmd != nil {
		if msg := cmd(); msg == (tea.QuitMsg{}) {
			t.Error("esc: must not emit QuitMsg while overlay is open")
		}
	}
}

func TestModePickerQKeyClosesNotQuits(t *testing.T) {
	m := newTestModel().WithModes(testModes())
	m, _ = pressKey(m, 'm')
	got, cmd := pressKey(m, 'q')
	m = got
	if m.pickerActive {
		t.Fatal("q: expected pickerActive=false")
	}
	// q must NOT quit the app while the overlay is open
	if cmd != nil {
		if msg := cmd(); msg == (tea.QuitMsg{}) {
			t.Error("q: must not emit QuitMsg while picker overlay is open")
		}
	}
}

func TestModePickerUnrelatedKeyConsumed(t *testing.T) {
	// An unrelated key (e.g. 't') while picker is open should not navigate the
	// underlying dashboard sections.
	m := newTestModel().WithModes(testModes())
	m, _ = pressKey(m, 'm')
	m.section = sectionPreflight
	got, _ := pressKey(m, 't')
	if got.section != sectionPreflight {
		t.Error("unrelated key leaked through picker overlay into navigation")
	}
	if !got.pickerActive {
		t.Error("unrelated key closed the picker overlay unexpectedly")
	}
}

// ── variable editor overlay ────────────────────────────────────────────────────

func testEnvVars() []provider.EnvVar {
	return []provider.EnvVar{
		{Key: "GITHUB_OWNER", Required: true, Secret: false, Doc: "GitHub org owner"},
		{Key: "GITHUB_TOKEN", Required: false, Secret: true, Doc: "PAT"},
		{Key: "GITHUB_APP_ID", Required: false, Secret: false, Doc: "App ID"},
	}
}

func testGetenv(k string) string {
	switch k {
	case "GITHUB_OWNER":
		return "my-test-org"
	case "GITHUB_TOKEN":
		return "ghp_SUPERSECRET123"
	default:
		return ""
	}
}

func TestEditorOpensWithVKey(t *testing.T) {
	m := newTestModel().
		WithEnvVars(testEnvVars()).
		WithGetenv(testGetenv)
	got, _ := pressKey(m, 'v')
	if !got.editorActive {
		t.Fatal("expected editorActive=true after pressing v")
	}
	if len(got.editorFields) != 3 {
		t.Fatalf("expected 3 editor fields, got %d", len(got.editorFields))
	}
	if got.editorFields[0].Key != "GITHUB_OWNER" {
		t.Errorf("field[0].Key = %q, want GITHUB_OWNER", got.editorFields[0].Key)
	}
	if got.editorFields[1].Key != "GITHUB_TOKEN" {
		t.Errorf("field[1].Key = %q, want GITHUB_TOKEN", got.editorFields[1].Key)
	}
	if !got.editorFields[1].Secret {
		t.Error("expected GITHUB_TOKEN to be secret")
	}
}

func TestEditorOpensWithEmptyEnvVarsWhenNoneSet(t *testing.T) {
	m := newTestModel() // no WithEnvVars
	got, _ := pressKey(m, 'v')
	if !got.editorActive {
		t.Fatal("expected editorActive=true (even with no vars)")
	}
	if len(got.editorFields) != 0 {
		t.Fatalf("expected 0 editor fields, got %d", len(got.editorFields))
	}
}

func TestEditorEscClosesOverlay(t *testing.T) {
	m := newTestModel().WithEnvVars(testEnvVars()).WithGetenv(testGetenv)
	m, _ = pressKey(m, 'v')
	got, cmd := m.Update(tea.KeyMsg{Type: tea.KeyEsc})
	m = got.(Model)
	if m.editorActive {
		t.Fatal("esc: expected editorActive=false")
	}
	if cmd != nil {
		if msg := cmd(); msg == (tea.QuitMsg{}) {
			t.Error("esc: must not emit QuitMsg while editor overlay is open")
		}
	}
}

func TestEditorQKeyClosesNotQuits(t *testing.T) {
	m := newTestModel().WithEnvVars(testEnvVars()).WithGetenv(testGetenv)
	m, _ = pressKey(m, 'v')
	got, cmd := pressKey(m, 'q')
	m = got
	if m.editorActive {
		t.Fatal("q: expected editorActive=false")
	}
	if cmd != nil {
		if msg := cmd(); msg == (tea.QuitMsg{}) {
			t.Error("q: must not emit QuitMsg while editor overlay is open")
		}
	}
}

func TestEditorSecretFieldNotEditable(t *testing.T) {
	// Field 1 is GITHUB_TOKEN (secret). Pressing enter on it must not focus
	// its input.
	m := newTestModel().WithEnvVars(testEnvVars()).WithGetenv(testGetenv)
	m, _ = pressKey(m, 'v')
	// Move down once to GITHUB_TOKEN.
	m, _ = pressKey(m, 'j')
	if m.editorCursor != 1 {
		t.Fatalf("expected editorCursor=1, got %d", m.editorCursor)
	}
	got, _ := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	m = got.(Model)
	if m.editorFocused {
		t.Error("secret field must not become editorFocused on enter")
	}
}

func TestEditorSetNonSecretEmitsSetEnvVarIntent(t *testing.T) {
	var capturedMsg tea.Msg
	exec := func(msg tea.Msg) tea.Cmd {
		capturedMsg = msg
		return func() tea.Msg { return msg }
	}
	m := newTestModel().
		WithEnvVars([]provider.EnvVar{
			{Key: "GITHUB_OWNER", Required: true, Secret: false},
		}).
		WithGetenv(func(string) string { return "" }).
		WithExec(exec)
	// Open editor.
	m, _ = pressKey(m, 'v')
	// Focus field 0 (GITHUB_OWNER) with enter.
	got, _ := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	m = got.(Model)
	if !m.editorFocused {
		t.Fatal("expected editorFocused=true after enter on non-secret field")
	}
	// Type "my-org".
	for _, r := range "my-org" {
		got2, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{r}})
		m = got2.(Model)
	}
	// Submit with enter.
	got3, cmd := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	m = got3.(Model)
	if m.editorFocused {
		t.Error("expected editorFocused=false after submit")
	}
	if cmd == nil {
		t.Fatal("expected non-nil cmd after submit")
	}
	cmd() // trigger exec
	intent, ok := capturedMsg.(SetEnvVarIntent)
	if !ok {
		t.Fatalf("expected SetEnvVarIntent, got %T", capturedMsg)
	}
	if intent.Key != "GITHUB_OWNER" {
		t.Errorf("intent.Key = %q, want GITHUB_OWNER", intent.Key)
	}
	if intent.Value != "my-org" {
		t.Errorf("intent.Value = %q, want my-org", intent.Value)
	}
}

func TestEditorEscOnFocusedFieldClosesEditor(t *testing.T) {
	m := newTestModel().
		WithEnvVars([]provider.EnvVar{{Key: "GITHUB_OWNER", Required: true}}).
		WithGetenv(func(string) string { return "" })
	m, _ = pressKey(m, 'v')
	// Focus field.
	got, _ := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	m = got.(Model)
	if !m.editorFocused {
		t.Fatal("expected focus after enter")
	}
	// Esc on focused field closes editor.
	got2, _ := m.Update(tea.KeyMsg{Type: tea.KeyEsc})
	m = got2.(Model)
	if m.editorFocused {
		t.Error("expected editorFocused=false after esc")
	}
	if m.editorActive {
		t.Error("expected editorActive=false after esc")
	}
}

// TestEditorSecretValueNeverInView is the redaction contract test: the raw
// value of a secret field must never appear in View() output.
func TestEditorSecretValueNeverInView(t *testing.T) {
	const secretValue = "ghp_SUPERSECRET123"
	m := newTestModel().
		WithEnvVars(testEnvVars()).
		WithGetenv(func(k string) string {
			if k == "GITHUB_TOKEN" {
				return secretValue
			}
			return ""
		})
	m, _ = pressKey(m, 'v')
	view := m.View()
	if strings.Contains(view, secretValue) {
		t.Errorf("secret value leaked into View() output: %q", view)
	}
	if !strings.Contains(view, "<set>") {
		t.Error("expected <set> placeholder for set secret field")
	}
}

func TestEditorUnsetSecretShowsUnset(t *testing.T) {
	m := newTestModel().
		WithEnvVars([]provider.EnvVar{
			{Key: "GITHUB_TOKEN", Required: false, Secret: true},
		}).
		WithGetenv(func(string) string { return "" })
	m, _ = pressKey(m, 'v')
	view := m.View()
	if !strings.Contains(view, "<unset>") {
		t.Errorf("expected <unset> for empty secret field; view=%q", view)
	}
}

// ── new intent dispatch ────────────────────────────────────────────────────────

func TestSwitchModeIntentRoutedViaExec(t *testing.T) {
	var got tea.Msg
	exec := func(msg tea.Msg) tea.Cmd {
		got = msg
		return func() tea.Msg { return msg }
	}
	m := New("github", "organization", "org", true).WithExec(exec)
	_, cmd := m.Update(SwitchModeIntent{Mode: "enterprise"})
	if cmd == nil {
		t.Fatal("expected non-nil cmd for SwitchModeIntent")
	}
	cmd()
	if intent, ok := got.(SwitchModeIntent); !ok || intent.Mode != "enterprise" {
		t.Errorf("expected SwitchModeIntent{Mode:enterprise}, got %T %v", got, got)
	}
}

func TestSetEnvVarIntentRoutedViaExec(t *testing.T) {
	var got tea.Msg
	exec := func(msg tea.Msg) tea.Cmd {
		got = msg
		return func() tea.Msg { return msg }
	}
	m := New("github", "organization", "org", true).WithExec(exec)
	_, cmd := m.Update(SetEnvVarIntent{Key: "GITHUB_OWNER", Value: "myorg"})
	if cmd == nil {
		t.Fatal("expected non-nil cmd for SetEnvVarIntent")
	}
	cmd()
	if intent, ok := got.(SetEnvVarIntent); !ok || intent.Key != "GITHUB_OWNER" {
		t.Errorf("expected SetEnvVarIntent, got %T %v", got, got)
	}
}

func TestPreflightIntentRoutedViaExec(t *testing.T) {
	var got tea.Msg
	exec := func(msg tea.Msg) tea.Cmd {
		got = msg
		return func() tea.Msg { return msg }
	}
	m := New("github", "organization", "org", true).WithExec(exec)
	_, cmd := m.Update(PreflightIntent{})
	if cmd == nil {
		t.Fatal("expected non-nil cmd for PreflightIntent")
	}
	cmd()
	if _, ok := got.(PreflightIntent); !ok {
		t.Errorf("expected PreflightIntent, got %T", got)
	}
}

func TestPreflightKeyEmitsPreflightIntent(t *testing.T) {
	var got tea.Msg
	exec := func(msg tea.Msg) tea.Cmd {
		got = msg
		return func() tea.Msg { return msg }
	}
	m := New("github", "organization", "org", true).WithExec(exec)
	_, cmd := pressKey(m, 'p')
	if cmd == nil {
		t.Fatal("expected non-nil cmd from p key")
	}
	cmd()
	if _, ok := got.(PreflightIntent); !ok {
		t.Errorf("expected PreflightIntent from p key, got %T", got)
	}
}

func TestPreflightMsgUpdatesEnvVars(t *testing.T) {
	m := newTestModel()
	newVars := []provider.EnvVar{
		{Key: "GITHUB_ENTERPRISE_SLUG", Required: true},
	}
	m2, _ := m.Update(PreflightMsg{
		Report:  provider.PreflightReport{Mode: "enterprise"},
		EnvVars: newVars,
	})
	got := m2.(Model)
	if got.mode != "enterprise" {
		t.Errorf("mode = %q, want enterprise", got.mode)
	}
	if len(got.envVars) != 1 || got.envVars[0].Key != "GITHUB_ENTERPRISE_SLUG" {
		t.Errorf("envVars not updated: %v", got.envVars)
	}
}

func TestPreflightMsgNilEnvVarsLeavesEnvVarsUnchanged(t *testing.T) {
	existing := []provider.EnvVar{{Key: "GITHUB_OWNER", Required: true}}
	m := newTestModel().WithEnvVars(existing)
	m2, _ := m.Update(PreflightMsg{
		Report:  provider.PreflightReport{Mode: "organization"},
		EnvVars: nil, // nil = no update
	})
	got := m2.(Model)
	if len(got.envVars) != 1 || got.envVars[0].Key != "GITHUB_OWNER" {
		t.Errorf("envVars changed when nil was passed: %v", got.envVars)
	}
}

// ── C12.2: p in editor nav emits PreflightIntent ───────────────────────────────

func TestEditorNavPressPreflightEmitsPreflightIntent(t *testing.T) {
	var capturedMsg tea.Msg
	exec := func(msg tea.Msg) tea.Cmd {
		capturedMsg = msg
		return func() tea.Msg { return msg }
	}
	m := newTestModel().
		WithEnvVars(testEnvVars()).
		WithGetenv(testGetenv).
		WithExec(exec)
	m, _ = pressKey(m, 'v')
	if !m.editorActive || m.editorFocused {
		t.Fatal("expected editor open in nav mode (not focused)")
	}
	_, cmd := pressKey(m, 'p')
	if cmd == nil {
		t.Fatal("expected non-nil cmd from p in editor nav")
	}
	cmd()
	if _, ok := capturedMsg.(PreflightIntent); !ok {
		t.Errorf("expected PreflightIntent from p in editor nav, got %T", capturedMsg)
	}
}

// ── C12.3: ctrl+c always quits from the mode-picker overlay ────────────────────

func TestPickerCtrlCAlwaysQuits(t *testing.T) {
	m := newTestModel().WithModes(testModes())
	m, _ = pressKey(m, 'm')
	if !m.pickerActive {
		t.Fatal("expected picker open after m")
	}
	_, cmd := m.Update(tea.KeyMsg{Type: tea.KeyCtrlC})
	if cmd == nil {
		t.Fatal("expected non-nil cmd from ctrl+c with picker open")
	}
	if msg := cmd(); msg != (tea.QuitMsg{}) {
		t.Fatalf("expected QuitMsg from ctrl+c in picker, got %T", msg)
	}
}

// ── C12.5: ctrl+c always quits from the variable editor (focused and nav) ──────

func TestEditorFocusedCtrlCAlwaysQuits(t *testing.T) {
	m := newTestModel().
		WithEnvVars([]provider.EnvVar{{Key: "GITHUB_OWNER", Required: true}}).
		WithGetenv(func(string) string { return "" })
	m, _ = pressKey(m, 'v')
	got, _ := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	m = got.(Model)
	if !m.editorFocused {
		t.Fatal("expected editorFocused after enter on non-secret field")
	}
	_, cmd := m.Update(tea.KeyMsg{Type: tea.KeyCtrlC})
	if cmd == nil {
		t.Fatal("expected non-nil cmd from ctrl+c with editor focused")
	}
	if msg := cmd(); msg != (tea.QuitMsg{}) {
		t.Fatalf("expected QuitMsg from ctrl+c in editor focused, got %T", msg)
	}
}

func TestEditorNavCtrlCAlwaysQuits(t *testing.T) {
	m := newTestModel().
		WithEnvVars(testEnvVars()).
		WithGetenv(testGetenv)
	m, _ = pressKey(m, 'v')
	if !m.editorActive || m.editorFocused {
		t.Fatal("expected editor open in nav mode (not focused)")
	}
	_, cmd := m.Update(tea.KeyMsg{Type: tea.KeyCtrlC})
	if cmd == nil {
		t.Fatal("expected non-nil cmd from ctrl+c in editor nav")
	}
	if msg := cmd(); msg != (tea.QuitMsg{}) {
		t.Fatalf("expected QuitMsg from ctrl+c in editor nav, got %T", msg)
	}
}
