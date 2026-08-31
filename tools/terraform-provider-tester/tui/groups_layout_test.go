package tui

import (
	"regexp"
	"strings"
	"testing"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"

	"github.com/github/terraform-provider-tester/engine"
)

var ansiRE = regexp.MustCompile("\x1b\\[[0-9;]*m")

func stripANSI(s string) string { return ansiRE.ReplaceAllString(s, "") }

// firstBarColumn returns the rune column of the first progress-bar cell on a
// (already ANSI-stripped) line, or -1 if the line has no bar. Only the filled
// and empty bar glyphs count; the fixtures here always have total > 0.
func firstBarColumn(line string) int {
	if idx := strings.Index(line, strings.Repeat(".", 10)); idx >= 0 {
		return len([]rune(line[:idx]))
	}
	for i, r := range []rune(line) {
		if r == '█' || r == '░' || r == '#' {
			return i
		}
	}
	return -1
}

// On a color terminal every group row is ANSI-styled. The progress bars must
// still line up in a single column and be equal width. This is the regression
// guard for the "%-Ns on an already-styled string" padding bug that left the
// bars in a ragged staircase (the styled name's ANSI bytes were counted toward
// the field width, so no visible padding was applied).
func TestRenderGroupsBarsAlignOnColorProfile(t *testing.T) {
	lipgloss.SetColorProfile(termenv.TrueColor)
	defer lipgloss.SetColorProfile(termenv.Ascii) // restore the TestMain default

	m := New("github", "anonymous", "", false)
	m.width = 100
	m.section = sectionGroups
	m.focus = focusGroups
	// Deliberately mixed-length names: this is what knocks the bars out of
	// alignment when padding is computed on the styled (ANSI-wrapped) string.
	m.groups = []engine.Group{
		{Name: "actions", Tests: []string{"TestAccA"}},
		{Name: "ip-ranges", Tests: []string{"TestAccB"}},
		{Name: "organization", Tests: []string{"TestAccC"}},
		{Name: "teams", Tests: []string{"TestAccD"}},
		{Name: "repositories", Tests: []string{"TestAccE"}},
	}
	m.groupCursor = 0 // exercise the selected-row gutter too

	out := m.renderGroups(m.width)
	lines := strings.Split(out, "\n")

	want := -1
	for _, ln := range lines {
		plain := stripANSI(ln)
		col := firstBarColumn(plain)
		if col < 0 {
			continue // header label and divider rule carry no bar
		}
		if want == -1 {
			want = col
			continue
		}
		if col != want {
			t.Errorf("progress bar misaligned: got column %d, want %d\nrow: %q", col, want, plain)
		}
	}
	if want == -1 {
		t.Fatal("no progress bars found in any row")
	}
}

// The passed/total counts must be right-aligned to a common width so the
// numbers read cleanly down the column.
func TestRenderGroupsCountsRightAligned(t *testing.T) {
	lipgloss.SetColorProfile(termenv.TrueColor)
	defer lipgloss.SetColorProfile(termenv.Ascii)

	m := New("github", "anonymous", "", false)
	m.width = 100
	m.section = sectionGroups
	m.focus = focusGroups
	m.groups = []engine.Group{
		{Name: "actions", Tests: makeNTests(34)},      // total width 2
		{Name: "ip-ranges", Tests: makeNTests(1)},     // total width 1
		{Name: "repositories", Tests: makeNTests(45)}, // total width 2
	}

	out := m.renderGroups(m.width)
	// Build the column of the bar; right-aligned counts keep it constant. Only
	// rows that carry a bar (the data rows and the totals row) are checked -
	// the header label and divider rule intentionally have neither bar nor
	// count column.
	cols := map[int]struct{}{}
	for _, ln := range strings.Split(out, "\n") {
		plain := stripANSI(ln)
		col := firstBarColumn(plain)
		if col < 0 {
			continue
		}
		if !strings.Contains(plain, "/") {
			t.Fatalf("row missing count column: %q", plain)
		}
		cols[col] = struct{}{}
	}
	if len(cols) != 1 {
		t.Errorf("expected a single bar column for right-aligned counts, got %d distinct columns", len(cols))
	}
}

func TestMissionPanelsStayIntactOnColorProfile(t *testing.T) {
	lipgloss.SetColorProfile(termenv.TrueColor)
	defer lipgloss.SetColorProfile(termenv.Ascii)

	const width = 70
	m := New("github", "organization", "", false)
	views := map[string]string{
		"preflight": m.renderPreflight(width),
		"groups":    m.renderGroups(width),
		"run":       m.renderRunConsole(width),
		"triage":    m.renderTriageSection(width),
	}
	for name, out := range views {
		t.Run(name, func(t *testing.T) {
			borderLines := 0
			for _, line := range strings.Split(out, "\n") {
				if got := lipgloss.Width(line); got > width {
					t.Fatalf("line width = %d, exceeds %d: %q", got, width, stripANSI(line))
				}
				plain := stripANSI(line)
				switch {
				case strings.HasPrefix(plain, "╭"):
					borderLines++
					if !strings.HasSuffix(plain, "╮") {
						t.Fatalf("truncated top panel border: %q", plain)
					}
				case strings.HasPrefix(plain, "│"):
					borderLines++
					if !strings.HasSuffix(plain, "│") {
						t.Fatalf("truncated panel body border: %q", plain)
					}
				case strings.HasPrefix(plain, "╰"):
					borderLines++
					if !strings.HasSuffix(plain, "╯") {
						t.Fatalf("truncated bottom panel border: %q", plain)
					}
				}
			}
			if borderLines == 0 {
				t.Fatal("rendered view has no complete panel border lines")
			}
		})
	}
}

func makeNTests(n int) []string {
	out := make([]string, n)
	for i := range out {
		out[i] = "TestAccGenerated" + strings.Repeat("x", i%3)
	}
	return out
}
