package tui

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

// TestWithIntroActivates asserts WithIntro() turns the splash on and Init()
// kicks the first tick, while a plain model stays inert (golden tests rely on
// the intro being OFF by default).
func TestWithIntroActivates(t *testing.T) {
	plain := New("github", "organization", "org", false)
	if plain.introActive {
		t.Errorf("intro should be off by default")
	}
	if plain.Init() != nil {
		t.Errorf("Init() should be nil when intro is off")
	}
	m := plain.WithIntro()
	if !m.introActive {
		t.Errorf("WithIntro() should activate the intro")
	}
	if m.Init() == nil {
		t.Errorf("Init() should return a tick cmd when intro is active")
	}
}

// TestIntroTickAdvancesAndStops asserts the intro tick increments the frame and
// self-stops at introTotalFrames (mirrors the spinner's self-stopping loop).
func TestIntroTickAdvancesAndStops(t *testing.T) {
	m := New("github", "organization", "org", false).WithIntro()
	m.width, m.height = 100, 30

	// One tick advances the frame and reschedules.
	next, cmd := m.Update(introTickMsg{})
	nm := next.(Model)
	if nm.introFrame != 1 {
		t.Fatalf("frame after one tick = %d, want 1", nm.introFrame)
	}
	if cmd == nil {
		t.Errorf("intro should reschedule while active")
	}

	// Drive to the end: the loop must stop and clear introActive. On completion
	// it returns tea.ClearScreen (not another tick), so the tick loop ends.
	nm.introFrame = introTotalFrames - 1
	end, endCmd := nm.Update(introTickMsg{})
	em := end.(Model)
	if em.introActive {
		t.Errorf("intro should deactivate at the final frame")
	}
	if endCmd != nil {
		if _, isTick := endCmd().(introTickMsg); isTick {
			t.Errorf("intro should not reschedule a tick after completing")
		}
	}
}

// TestKeyDismissesIntro asserts any key press skips the splash immediately.
func TestKeyDismissesIntro(t *testing.T) {
	m := New("github", "organization", "org", false).WithIntro()
	m.width, m.height = 100, 30
	next, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'x'}})
	if next.(Model).introActive {
		t.Errorf("a key press should dismiss the intro")
	}
}

// TestViewShowsIntroThenDashboard asserts View renders the splash while the
// intro is active and the normal dashboard once it finishes.
func TestViewShowsIntroThenDashboard(t *testing.T) {
	m := New("github", "organization", "my-test-org", false).WithIntro()
	m.width, m.height = 100, 30
	m.version = "1.2.3"
	m.introFrame = introPeakFrame

	intro := m.View()
	if !strings.Contains(intro, pixelSolid) {
		t.Errorf("active intro View should show the wordmark")
	}
	if !strings.Contains(intro, "press any key to skip") {
		t.Errorf("active intro View should show the splash hint")
	}
	if strings.Contains(intro, "Preflight") {
		t.Errorf("active intro View should not show the dashboard tabs")
	}

	m.introActive = false
	dash := m.View()
	if strings.Contains(dash, "press any key to skip") {
		t.Errorf("dashboard View should not show the splash hint")
	}
	if !strings.Contains(dash, "Preflight") {
		t.Errorf("dashboard View should show the tabs")
	}
}
