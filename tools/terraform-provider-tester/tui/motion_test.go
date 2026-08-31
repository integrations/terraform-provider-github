package tui

import (
	"strings"
	"testing"
	"time"
)

// base is a fixed clock for deterministic elapsed/stall assertions. The liveness
// layer must derive all timing from the time carried in spinnerTickMsg, never
// from a hidden time.Now(), so these tests pin exact instants.
var base = time.Date(2026, 6, 24, 12, 0, 0, 0, time.UTC)

// TestRunStartedSetsLabel verifies the producer-supplied verb label appears in
// the status line as "running <label>".
func TestRunStartedSetsLabel(t *testing.T) {
	m := newTestModel()
	m2, _ := m.Update(RunStartedMsg{Total: 1, Label: "ip-ranges"})
	got := m2.(Model)
	if got.runLabel != "ip-ranges" {
		t.Fatalf("runLabel = %q, want %q", got.runLabel, "ip-ranges")
	}
	line := got.renderStatus(80)
	if !strings.Contains(line, "running ip-ranges") {
		t.Errorf("status line = %q, want it to contain %q", line, "running ip-ranges")
	}
}

// TestElapsedTicksWhileRunning verifies the elapsed time is computed from the
// tick timestamps and rendered in gh-style compact duration.
func TestElapsedTicksWhileRunning(t *testing.T) {
	m := newTestModel()
	m2, _ := m.Update(RunStartedMsg{Total: 1, Label: "ip-ranges"})
	m = m2.(Model)

	// First tick anchors the start; elapsed is 0 and must not show yet.
	m2, _ = m.Update(spinnerTickMsg{t: base})
	m = m2.(Model)
	if strings.Contains(m.renderStatus(80), "0s") {
		t.Errorf("did not expect a 0s elapsed on the anchoring tick: %q", m.renderStatus(80))
	}

	// 90 seconds later the line shows 1m30s.
	m2, _ = m.Update(spinnerTickMsg{t: base.Add(90 * time.Second)})
	m = m2.(Model)
	line := m.renderStatus(80)
	if !strings.Contains(line, "1m30s") {
		t.Errorf("status line = %q, want it to contain %q", line, "1m30s")
	}
}

// TestElapsedHiddenWhenIdle is the golden-safety guard: a model that never ticks
// renders no elapsed text, so golden snapshots cannot move.
func TestElapsedHiddenWhenIdle(t *testing.T) {
	m := newTestModel()
	m2, _ := m.Update(RunStartedMsg{Total: 1, Label: "ip-ranges"})
	m = m2.(Model)
	line := m.renderStatus(80)
	if strings.Contains(line, " · ") {
		t.Errorf("expected no elapsed/stall separators before any tick, got %q", line)
	}
}

// TestStallHintAppearsAfterThreshold verifies a quiet period past the threshold
// surfaces a stall hint, and that a new test result clears it.
func TestStallHintAppearsAfterThreshold(t *testing.T) {
	m := newTestModel()
	m2, _ := m.Update(RunStartedMsg{Total: 2, Label: "repos"})
	m = m2.(Model)
	m2, _ = m.Update(spinnerTickMsg{t: base}) // anchor
	m = m2.(Model)

	// 50s of silence (> 45s threshold) → stall hint.
	m2, _ = m.Update(spinnerTickMsg{t: base.Add(50 * time.Second)})
	m = m2.(Model)
	if !strings.Contains(m.renderStatus(120), "quiet") {
		t.Errorf("expected a stall hint after 50s of silence, got %q", m.renderStatus(120))
	}

	// A new result arrives, then a tick: the stall hint clears.
	m2, _ = m.Update(TestUpdateMsg{Result: makeResult("TestAccRepoA", "", 0, 1.0)})
	m = m2.(Model)
	m2, _ = m.Update(spinnerTickMsg{t: base.Add(51 * time.Second)})
	m = m2.(Model)
	if strings.Contains(m.renderStatus(120), "quiet") {
		t.Errorf("expected stall hint cleared after a new result, got %q", m.renderStatus(120))
	}
}

// TestStallHintHiddenBelowThreshold verifies a short quiet period shows no hint.
func TestStallHintHiddenBelowThreshold(t *testing.T) {
	m := newTestModel()
	m2, _ := m.Update(RunStartedMsg{Total: 1, Label: "repos"})
	m = m2.(Model)
	m2, _ = m.Update(spinnerTickMsg{t: base})
	m = m2.(Model)
	m2, _ = m.Update(spinnerTickMsg{t: base.Add(30 * time.Second)})
	m = m2.(Model)
	if strings.Contains(m.renderStatus(120), "quiet") {
		t.Errorf("did not expect a stall hint at 30s (< 45s threshold), got %q", m.renderStatus(120))
	}
}

// TestStatusLineComposesSpinnerLabelElapsed verifies left-to-right ordering:
// spinner glyph first, then "running <label>", then the elapsed time.
func TestStatusLineComposesSpinnerLabelElapsed(t *testing.T) {
	m := newTestModel()
	m2, _ := m.Update(RunStartedMsg{Total: 1, Label: "ip-ranges"})
	m = m2.(Model)
	m2, _ = m.Update(spinnerTickMsg{t: base})
	m = m2.(Model)
	m2, _ = m.Update(spinnerTickMsg{t: base.Add(5 * time.Second)})
	m = m2.(Model)

	line := m.renderStatus(120)
	iLabel := strings.Index(line, "running ip-ranges")
	iElapsed := strings.Index(line, "5s")
	if iLabel <= 0 {
		t.Fatalf("expected the spinner glyph to precede the label, got %q", line)
	}
	if iElapsed < 0 {
		t.Fatalf("status line missing elapsed: %q", line)
	}
	if iElapsed < iLabel {
		t.Errorf("expected elapsed after label, got %q", line)
	}
}
