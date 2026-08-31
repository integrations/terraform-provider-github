package tui

import (
	"strings"
	"testing"

	"github.com/charmbracelet/lipgloss"
)

// countSubstr counts non-overlapping occurrences of sub in s.
func countSubstr(s, sub string) int { return strings.Count(s, sub) }

// TestWordmarkFinalFrameAllSolid asserts that at full progress every lit pixel
// of the wordmark is the solid block glyph (the ignition has fully resolved):
// no dither or scatter glyphs remain.
func TestWordmarkFinalFrameAllSolid(t *testing.T) {
	got := renderWordmark(1.0, false)
	if strings.Contains(got, ditherHeavy) || strings.Contains(got, ditherLight) || strings.Contains(got, scatterDot) {
		t.Errorf("final frame still contains transition glyphs; want all solid.\n%s", got)
	}
	// "GH/TF" has a known number of lit pixels in the 5x3 block font; the final
	// frame must render at least that many solid cells (sanity that it's not blank).
	if n := countSubstr(got, pixelSolid); n < 40 {
		t.Errorf("final frame has too few solid pixels (%d); wordmark looks empty.\n%s", n, got)
	}
}

// TestWordmarkFirstFrameNoSolids asserts that at zero progress nothing has
// resolved yet: there are no solid block pixels (only scatter/blank).
func TestWordmarkFirstFrameNoSolids(t *testing.T) {
	got := renderWordmark(0.0, false)
	if strings.Contains(got, pixelSolid) {
		t.Errorf("first frame should have no solid pixels yet.\n%s", got)
	}
}

// TestWordmarkDeterministic asserts the renderer is a pure function of its
// inputs: identical args yield identical output (required for golden stability).
func TestWordmarkDeterministic(t *testing.T) {
	for _, p := range []float64{0.0, 0.33, 0.5, 0.77, 1.0} {
		a := renderWordmark(p, false)
		b := renderWordmark(p, false)
		if a != b {
			t.Errorf("renderWordmark(%v) not deterministic", p)
		}
	}
}

// TestWordmarkRowsEqualWidth asserts every rendered row measures the same
// display width, so the block letters stay aligned (lipgloss.Width is
// ANSI/rune aware, so color codes do not affect the measurement).
func TestWordmarkRowsEqualWidth(t *testing.T) {
	for _, p := range []float64{0.0, 0.4, 1.0} {
		rows := strings.Split(renderWordmark(p, false), "\n")
		if len(rows) != wordmarkRows {
			t.Fatalf("progress %v: got %d rows, want %d", p, len(rows), wordmarkRows)
		}
		want := lipgloss.Width(rows[0])
		for i, r := range rows {
			if w := lipgloss.Width(r); w != want {
				t.Errorf("progress %v: row %d width %d != %d", p, i, w, want)
			}
		}
		if want != wordmarkCellWidth {
			t.Errorf("progress %v: row width %d != expected %d", p, want, wordmarkCellWidth)
		}
	}
}

// TestWordmarkASCIIVariantDiffers asserts the ascii fallback uses different
// glyphs from the block variant at the same progress and is total (renders).
func TestWordmarkASCIIVariantDiffers(t *testing.T) {
	block := renderWordmark(1.0, false)
	ascii := renderWordmark(1.0, true)
	if ascii == block {
		t.Errorf("ascii variant should differ from block variant")
	}
	if strings.Contains(ascii, pixelSolid) {
		t.Errorf("ascii variant must not contain block solid glyph %q", pixelSolid)
	}
	if !strings.Contains(ascii, pixelSolidASCII) {
		t.Errorf("ascii final frame should contain ascii solid glyph %q", pixelSolidASCII)
	}
}

func TestWordmarkUsesGHTFBrand(t *testing.T) {
	if wordmarkText != "GH/TF" {
		t.Fatalf("wordmarkText = %q, want GH/TF", wordmarkText)
	}
	got := renderWordmark(1, true)
	if strings.Contains(got, "TESTER") || strings.Contains(got, "PULSAR") {
		t.Fatalf("resolved wordmark contains a retired dashboard brand:\n%s", got)
	}
	compareOrUpdate(t, "wordmark_full", got)
}

// TestWordmarkProgressMonotoneSolids asserts the ignition fills in over time:
// later frames have at least as many solid pixels as earlier frames.
func TestWordmarkProgressMonotoneSolids(t *testing.T) {
	prev := -1
	for _, p := range []float64{0.0, 0.2, 0.4, 0.6, 0.8, 1.0} {
		n := countSubstr(renderWordmark(p, false), pixelSolid)
		if n < prev {
			t.Errorf("solids decreased at progress %v: %d < %d", p, n, prev)
		}
		prev = n
	}
}

// TestWordmarkGoldens pins two representative frames so the visual shape is
// reviewed on change. Color is stripped by the Ascii profile (TestMain), so the
// goldens show the raw block/dither glyphs only.
func TestWordmarkGoldens(t *testing.T) {
	cases := []struct {
		name     string
		progress float64
		ascii    bool
	}{
		{"wordmark_mid", 0.45, false},
		{"wordmark_full_tty", 1.0, false},
		{"wordmark_full", 1.0, true},
	}
	for _, tc := range cases {
		got := renderWordmark(tc.progress, tc.ascii)
		assertNoANSI(t, got)
		compareOrUpdate(t, tc.name, got)
	}
}
