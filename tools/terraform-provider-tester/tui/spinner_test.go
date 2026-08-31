package tui

import "testing"

// TestPulsarAnimatesInRenderedView drives the real render path: starting from a
// running model, each spinnerTickMsg must produce a status line whose leading
// glyph advances through every Pulsar frame in gh's order. This proves the
// animation is visible in rendered View output, not just in the frame counter.
func TestPulsarAnimatesInRenderedView(t *testing.T) {
	m := New("github", "anonymous", "", false) // ascii=false → braille frames
	m.running = true
	m.status = "running…"

	got := make([]string, 0, len(PulsarFrames))
	for range PulsarFrames {
		got = append(got, firstBraille(m.renderStatus(80)))
		m2, _ := m.Update(spinnerTickMsg{})
		m = m2.(Model)
	}

	for i, want := range PulsarFrames {
		if got[i] != want {
			t.Errorf("rendered frame %d = %q, want %q (full sequence %v)", i, got[i], want, got)
		}
	}
}

// firstBraille returns the first braille-range rune in s as a string, or "" if
// the string contains none.
func firstBraille(s string) string {
	for _, r := range s {
		if r >= 0x2800 && r <= 0x28FF {
			return string(r)
		}
	}
	return ""
}

// TestPulsarFramesMatchGH locks Pulsar's TTY frames to the exact GitHub CLI
// spinner set (briandowns/spinner CharSets[11]). If gh's animation is the
// reference, these glyphs and their order must not drift.
func TestPulsarFramesMatchGH(t *testing.T) {
	ghCharSet11 := []string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"}
	if len(PulsarFrames) != len(ghCharSet11) {
		t.Fatalf("PulsarFrames has %d frames, want %d (gh CharSets[11])", len(PulsarFrames), len(ghCharSet11))
	}
	for i, want := range ghCharSet11 {
		if PulsarFrames[i] != want {
			t.Errorf("PulsarFrames[%d] = %q, want %q (gh CharSets[11])", i, PulsarFrames[i], want)
		}
	}
}

// TestSpinnerGlyphWraps verifies spinnerGlyph indexes frames[frame % len] for
// both the TTY and ASCII sets, including wrap-around past the end.
func TestSpinnerGlyphWraps(t *testing.T) {
	cases := []struct {
		frame    int
		ascii    bool
		wantFrom []string
	}{
		{0, false, PulsarFrames},
		{3, false, PulsarFrames},
		{len(PulsarFrames), false, PulsarFrames},     // wraps to index 0
		{len(PulsarFrames) + 2, false, PulsarFrames}, // wraps to index 2
		{0, true, PulsarFramesASCII},
		{1, true, PulsarFramesASCII},
		{len(PulsarFramesASCII) + 1, true, PulsarFramesASCII}, // wraps to index 1
	}
	for _, c := range cases {
		got := spinnerGlyph(c.frame, c.ascii)
		want := c.wantFrom[c.frame%len(c.wantFrom)]
		if got != want {
			t.Errorf("spinnerGlyph(%d, ascii=%v) = %q, want %q", c.frame, c.ascii, got, want)
		}
	}
}

// TestSpinnerGlyphASCIIIsNotBraille guards the no-color fallback: ASCII frames
// must be plain single-byte glyphs so they render in dumb / NO_COLOR terminals.
func TestSpinnerGlyphASCIIIsNotBraille(t *testing.T) {
	want := []string{"|", "/", "-", "\\"}
	if len(PulsarFramesASCII) != len(want) {
		t.Fatalf("PulsarFramesASCII has %d frames, want %d", len(PulsarFramesASCII), len(want))
	}
	for i, w := range want {
		if PulsarFramesASCII[i] != w {
			t.Errorf("PulsarFramesASCII[%d] = %q, want %q", i, PulsarFramesASCII[i], w)
		}
	}
}
