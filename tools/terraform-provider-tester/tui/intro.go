package tui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// The intro is the one-shot "ignition" splash shown when the dashboard opens on
// a color TTY: the GH/TF wordmark resolves out of a dot scatter, holds for a
// beat, then the dashboard takes over. It is skippable (any key) and is never
// shown in ascii/NO_COLOR mode (the dashboard opens directly). Like every other
// frame in this package, renderIntro is PURE: progress comes from a frame count
// the Update loop advances via introTickMsg, never from the clock inside View.

const (
	// The intro plays in three phases, all frame-driven (no clock in View):
	//   reveal  - the wordmark resolves out of a dot scatter (progress 0→1)
	//   hold    - it lingers fully resolved (progress 1)
	//   dissolve- it dissolves back into a scatter and clears (progress 1→0)
	introRevealFrames = 30
	introHoldFrames   = 12
	introExitFrames   = 22
	// introTotalFrames is the full intro length, after which it auto-dismisses.
	introTotalFrames = introRevealFrames + introHoldFrames + introExitFrames
	// introPeakFrame is the first frame at which the wordmark is fully resolved
	// (start of the hold). Used by tests and as the "settled" reference.
	introPeakFrame = introRevealFrames
	// introInterval is the per-frame cadence (~64 frames * 55ms ≈ 3.5s total).
	introInterval = 55 * time.Millisecond
)

// introProgress maps the frame counter to wordmark resolve progress in [0,1]
// across the three phases: ramps up during reveal, holds at 1, then ramps back
// down during the dissolve-out so the splash "goes away" by un-resolving rather
// than blinking off.
func introProgress(frame int) float64 {
	switch {
	case frame <= 0:
		return 0
	case frame < introRevealFrames:
		return float64(frame) / float64(introRevealFrames)
	case frame <= introRevealFrames+introHoldFrames:
		return 1
	case frame < introTotalFrames:
		out := frame - (introRevealFrames + introHoldFrames)
		return 1 - float64(out)/float64(introExitFrames)
	default:
		return 0
	}
}

// renderIntro builds the full-screen splash for the given frame, centered in a
// width×height box. version is shown beneath the wordmark.
func renderIntro(frame, width, height int, version string, ascii bool) string {
	wordmark := renderWordmark(introProgress(frame), ascii)

	tagline := lipgloss.NewStyle().Foreground(Muted).Render(Tagline)

	ver := renderConsoleTitle(true, version) +
		lipgloss.NewStyle().Foreground(Muted).Render(" · terraform-provider-tester")

	hint := lipgloss.NewStyle().Foreground(Muted).Render("press any key to skip")

	content := lipgloss.JoinVertical(
		lipgloss.Center,
		wordmark,
		"",
		ver,
		tagline,
		"",
		hint,
	)

	return lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, content)
}

// introTickMsg advances the intro animation one frame. It carries the tick time
// only for symmetry with the spinner; the intro derives nothing from the clock.
type introTickMsg struct{ t time.Time }

// introTickCmd schedules the next intro frame.
func introTickCmd() tea.Cmd {
	return tea.Tick(introInterval, func(t time.Time) tea.Msg {
		return introTickMsg{t: t}
	})
}
