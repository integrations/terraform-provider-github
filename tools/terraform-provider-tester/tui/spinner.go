package tui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// Pulsar is the Pulsar harness loading animation. A pulsar is a rotating beacon
// that emits regular pulses, which is exactly a periodic spinner. The frames are
// GitHub CLI's: the 8-dot braille set briandowns/spinner exposes as CharSets[11],
// advanced every PulsarInterval (gh's cadence). The name is ours; the motion is
// gh's, so the harness feels familiar to anyone who uses the GitHub CLI.
var (
	// PulsarFrames are the TTY frames, identical to gh's CharSets[11] in order.
	PulsarFrames = []string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"}

	// PulsarFramesASCII is the no-color fallback (gh disables its spinner on
	// non-TTY; we keep motion with a plain ASCII line so dumb / NO_COLOR
	// terminals still show liveness).
	PulsarFramesASCII = []string{"|", "/", "-", "\\"}

	// PulsarInterval is gh's spinner cadence (briandowns/spinner, 120ms ~ 8 FPS).
	PulsarInterval = 120 * time.Millisecond
)

// spinnerGlyph returns the Pulsar frame for the given counter, picking the ASCII
// fallback set when ascii is true. The counter wraps via modulo, so any frame
// value is valid; the extra term keeps it total even for negative inputs.
func spinnerGlyph(frame int, ascii bool) string {
	frames := PulsarFrames
	if ascii {
		frames = PulsarFramesASCII
	}
	n := len(frames)
	return frames[((frame%n)+n)%n]
}

// spinnerTickMsg drives the Pulsar animation. One is delivered every
// PulsarInterval while a run is in progress; the loop self-stops when the run
// ends (see the Update handler). It carries the tick time so the liveness layer
// (elapsed time, stall detection) derives all timing from an explicit input
// rather than a hidden time.Now() inside Update or View, keeping both testable
// and golden-stable.
type spinnerTickMsg struct{ t time.Time }

// spinnerTickCmd schedules the next Pulsar frame.
func spinnerTickCmd() tea.Cmd {
	return tea.Tick(PulsarInterval, func(t time.Time) tea.Msg {
		return spinnerTickMsg{t: t}
	})
}
