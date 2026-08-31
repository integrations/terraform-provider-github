package tui

import (
	"strings"
	"testing"

	"github.com/github/terraform-provider-tester/engine"
	"github.com/github/terraform-provider-tester/provider"
)

// The Groups table is the heart of the harness, but without a header row a
// newcomer cannot tell what the "2/4" count or the bar represent. A dim header
// must label the columns and align to the same columns as the data rows.
func TestGroupsTableShowsColumnHeader(t *testing.T) {
	m := New("github", "anonymous", "", true)
	m.width = 100
	m.section = sectionGroups
	m.focus = focusGroups
	m.groupCursor = -1 // no selection: keep every gutter ASCII for stable columns
	m.groups = []engine.Group{
		{Name: "repositories", Tests: []string{"TestAccA", "TestAccB"}},
		{Name: "teams", Tests: []string{"TestAccC"}},
	}

	out := stripANSI(m.renderGroups(m.width))
	lines := strings.Split(out, "\n")

	header := ""
	for _, line := range lines {
		if strings.Contains(line, "GROUP") && strings.Contains(line, "PROGRESS") {
			header = line
			break
		}
	}
	if header == "" {
		t.Fatalf("no groups table header found:\n%s", out)
	}
	for _, want := range []string{"GROUP", "PASS", "TIME", "PROGRESS"} {
		if !strings.Contains(header, want) {
			t.Fatalf("header row missing %q: %q", want, header)
		}
	}

	var dataLine string
	for _, ln := range lines {
		if strings.Contains(ln, "repositories") {
			dataLine = ln
			break
		}
	}
	if dataLine == "" {
		t.Fatal("no data row rendered for the repositories group")
	}

	// The GROUP label must start in the same column as the group names.
	if g, d := strings.Index(header, "GROUP"), strings.Index(dataLine, "repositories"); g != d {
		t.Errorf("GROUP header at col %d but group names start at col %d", g, d)
	}
	// The PROGRESS label must sit at the progress-bar column.
	if p, b := strings.Index(header, "PROGRESS"), firstBarColumn(dataLine); p != b {
		t.Errorf("PROGRESS header at col %d but bars start at col %d", p, b)
	}
}

// The overall totals can be wider than any single group's count (e.g. 15
// groups of <=45 tests sum to 174). The totals row must still align with the
// data rows, which means the measured count/duration columns have to account
// for the totals row too -- not just the per-group values.
func TestGroupsTableTotalsWiderThanGroupsStillAligns(t *testing.T) {
	m := New("github", "anonymous", "", true)
	m.width = 100
	m.section = sectionGroups
	m.focus = focusGroups
	m.groupCursor = -1
	// Each group has a single-digit total ("0/9"), but twelve of them sum to a
	// three-digit total ("0/108") that is wider than any per-group count.
	var groups []engine.Group
	for i := 0; i < 12; i++ {
		groups = append(groups, engine.Group{Name: "grp", Tests: makeNTests(9)})
	}
	m.groups = groups

	out := stripANSI(m.renderGroups(m.width))
	cols := map[int]struct{}{}
	for _, ln := range strings.Split(out, "\n") {
		col := firstBarColumn(ln)
		if col < 0 {
			continue
		}
		cols[col] = struct{}{}
	}
	if len(cols) != 1 {
		t.Errorf("totals row knocked the bars out of alignment: got %d distinct bar columns, want 1", len(cols))
	}
}

// The table must end with a TOTAL row so the operator sees overall progress at
// a glance instead of summing groups in their head: total counts, an overall
// percentage gauge, and the aggregate failure count.
func TestGroupsTableShowsTotalsRow(t *testing.T) {
	m := New("github", "anonymous", "", true)
	m.width = 100
	m.section = sectionGroups
	m.focus = focusGroups
	m.groupCursor = -1
	m.groups = []engine.Group{
		{Name: "alpha", Tests: []string{"TestAccA", "TestAccB"}},
		{Name: "beta", Tests: []string{"TestAccC"}},
	}
	m.results = []engine.TestResult{
		{Name: "TestAccA", Status: provider.StatusPass, Elapsed: 1.0},
		{Name: "TestAccB", Status: provider.StatusFail, Elapsed: 2.0},
		{Name: "TestAccC", Status: provider.StatusPass, Elapsed: 0.5},
	}

	out := stripANSI(m.renderGroups(m.width))
	lines := strings.Split(out, "\n")
	total := ""
	for _, line := range lines {
		if strings.Contains(line, "TOTAL") {
			total = line
			break
		}
	}
	if total == "" {
		t.Fatalf("no totals row rendered:\n%s", out)
	}

	// 2 of 3 passed (66%), 1 failed.
	for _, want := range []string{"TOTAL", "2/3", "66%", "1 failed"} {
		if !strings.Contains(total, want) {
			t.Errorf("totals row missing %q: %q", want, total)
		}
	}
}
