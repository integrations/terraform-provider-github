package tui

import (
	"fmt"
	"strings"
	"testing"

	"github.com/github/terraform-provider-tester/engine"
	"github.com/github/terraform-provider-tester/provider"
)

func TestTabSummariesDeriveBadges(t *testing.T) {
	m := fixedModel()
	m.checks = []provider.Check{{Status: provider.CheckOK}, {Status: provider.CheckFail}}
	m.groups = []engine.Group{{Name: "a", Tests: []string{"TestA", "TestB"}}}
	m.triageFailures = sampleFailures()
	got := m.tabSummaries()
	if got[0].Badge != "1 blocked" || got[1].Badge != "1" || got[2].Badge != "2" || got[3].Badge != "4" {
		t.Fatalf("tab badges = %#v", got)
	}
}

func TestChromeFitsResponsiveWidths(t *testing.T) {
	for _, width := range []int{60, 90, 140} {
		m := fixedModel()
		m.width = width
		assertRenderedWidth(t, m.renderHeader(width), width)
		assertRenderedWidth(t, m.renderTabs(width), width)
		assertRenderedWidth(t, m.renderFooter(width), width)
	}
}

func TestTabsAdvertiseNumberedSections(t *testing.T) {
	got := stripANSI(fixedModel().renderTabs(100))
	for _, want := range []string{"1 Preflight", "2 Groups", "3 Run", "4 Triage"} {
		if !strings.Contains(got, want) {
			t.Fatalf("tabs missing %q:\n%s", want, got)
		}
	}
}

func TestTabsUseASCIIUnderlineWhenASCIIEnabled(t *testing.T) {
	got := stripANSI(fixedModel().renderTabs(100))
	if strings.Contains(got, "─") {
		t.Fatalf("ascii tabs should not contain Unicode rule glyphs:\n%s", got)
	}
	if !strings.Contains(got, "-") {
		t.Fatalf("ascii tabs missing underline rule:\n%s", got)
	}
}

func TestChromeDoesNotPadTrailingSpaces(t *testing.T) {
	m := fixedModel()
	m.groups = []engine.Group{{Name: "a", Tests: []string{"TestA"}}}
	for _, rendered := range []string{m.renderHeader(100), m.renderTabs(100), m.renderFooter(100)} {
		for _, line := range strings.Split(stripANSI(rendered), "\n") {
			if strings.HasSuffix(line, " ") {
				t.Fatalf("chrome line has trailing spaces %q in:\n%s", line, rendered)
			}
		}
	}
}

func TestTabsFitVeryNarrowPositiveWidths(t *testing.T) {
	for _, width := range []int{1, 8, 20, 40, 71} {
		t.Run(fmt.Sprintf("width_%d", width), func(t *testing.T) {
			m := fixedModel()
			m.section = sectionTriage
			m.checks = []provider.Check{{Status: provider.CheckFail}, {Status: provider.CheckWarn}}
			m.groups = []engine.Group{
				{Name: "repositories", Tests: []string{"TestAccOne", "TestAccTwo", "TestAccThree"}},
				{Name: "actions", Tests: []string{"TestAccFour", "TestAccFive"}},
			}
			m.triageFailures = sampleFailures()

			var got string
			func() {
				defer func() {
					if r := recover(); r != nil {
						t.Fatalf("renderTabs(%d) panicked: %v", width, r)
					}
				}()
				got = stripANSI(m.renderTabs(width))
			}()
			assertRenderedWidth(t, got, width)
			if !strings.Contains(got, "4") {
				t.Fatalf("tabs at width %d should preserve active tab number:\n%s", width, got)
			}
			if width >= len("4 Tri") && !strings.Contains(got, "4 Tri") {
				t.Fatalf("tabs at width %d should preserve active tab label when it fits:\n%s", width, got)
			}
		})
	}
}
