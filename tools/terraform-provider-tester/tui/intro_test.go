package tui

import (
	"strings"
	"testing"

	"github.com/charmbracelet/lipgloss"
)

// TestIntroProgressClamps asserts the frame→progress mapping ramps up to 1.0 at
// the peak, holds, then ramps back to 0 as the splash dissolves out.
func TestIntroProgressClamps(t *testing.T) {
	if p := introProgress(0); p != 0 {
		t.Errorf("frame 0 progress = %v, want 0", p)
	}
	if p := introProgress(introPeakFrame); p != 1 {
		t.Errorf("frame %d progress = %v, want 1 (peak)", introPeakFrame, p)
	}
	if p := introProgress(introRevealFrames + introHoldFrames); p != 1 {
		t.Errorf("end-of-hold progress = %v, want 1", p)
	}
	if p := introProgress(introTotalFrames); p != 0 {
		t.Errorf("end-of-dissolve progress = %v, want 0 (fully gone)", p)
	}
	if p := introProgress(introTotalFrames + 99); p != 0 {
		t.Errorf("overrun progress = %v, want 0", p)
	}
}

// TestRenderIntroFillsBox asserts the splash exactly fills the given width and
// height (lipgloss.Place pads to the box), so it owns the alt-screen cleanly.
func TestRenderIntroFillsBox(t *testing.T) {
	const w, h = 100, 30
	out := renderIntro(introPeakFrame, w, h, "1.2.3", false)
	lines := strings.Split(out, "\n")
	if len(lines) != h {
		t.Fatalf("got %d lines, want %d", len(lines), h)
	}
	for i, ln := range lines {
		if got := lipgloss.Width(ln); got != w {
			t.Errorf("line %d width %d, want %d", i, got, w)
		}
	}
}

// TestRenderIntroPeakResolved asserts the peak frame shows the fully resolved
// wordmark (solid blocks) plus the tagline and version.
func TestRenderIntroPeakResolved(t *testing.T) {
	out := renderIntro(introPeakFrame, 100, 30, "1.2.3", false)
	if !strings.Contains(out, pixelSolid) {
		t.Errorf("peak intro frame should contain solid wordmark pixels")
	}
	if !strings.Contains(out, Tagline) {
		t.Errorf("intro should show the tagline %q", Tagline)
	}
	if !strings.Contains(out, "1.2.3") {
		t.Errorf("intro should show the version")
	}
}

// TestRenderIntroDissolvesOut asserts the splash un-resolves on the way out:
// the final dissolve frame has strictly fewer solid pixels than the peak, so it
// "goes away" by dissolving rather than blinking off.
func TestRenderIntroDissolvesOut(t *testing.T) {
	peak := countSubstr(renderIntro(introPeakFrame, 100, 30, "1.2.3", false), pixelSolid)
	tail := countSubstr(renderIntro(introTotalFrames-1, 100, 30, "1.2.3", false), pixelSolid)
	if tail >= peak {
		t.Errorf("dissolve-out should reduce solids: tail=%d peak=%d", tail, peak)
	}
}

// TestRenderIntroDeterministic asserts the splash is a pure function of inputs.
func TestRenderIntroDeterministic(t *testing.T) {
	for _, f := range []int{0, 5, introPeakFrame, introTotalFrames} {
		a := renderIntro(f, 100, 30, "1.2.3", false)
		b := renderIntro(f, 100, 30, "1.2.3", false)
		if a != b {
			t.Errorf("renderIntro(frame=%d) not deterministic", f)
		}
	}
}

// TestRenderIntroASCII asserts the ascii splash avoids block glyphs and stays
// within the box, so dumb / NO_COLOR terminals render cleanly.
func TestRenderIntroASCII(t *testing.T) {
	out := renderIntro(introPeakFrame, 100, 30, "1.2.3", true)
	if strings.Contains(out, pixelSolid) {
		t.Errorf("ascii intro must not contain block glyph %q", pixelSolid)
	}
	if !strings.Contains(out, pixelSolidASCII) {
		t.Errorf("ascii intro should contain ascii solid glyph %q", pixelSolidASCII)
	}
	assertNoANSI(t, out)
}
