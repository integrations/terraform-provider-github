package tui

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/charmbracelet/lipgloss"

	"github.com/github/terraform-provider-tester/engine"
	"github.com/github/terraform-provider-tester/provider"
)

func TestFinalReviewRunStateOnlyUsesRunScopedErrors(t *testing.T) {
	completedRun := func() Model {
		m := fixedModel()
		m.groups = []engine.Group{{Name: "repositories", Tests: []string{"TestA"}}}
		m.results = []engine.TestResult{{Name: "TestA", Status: provider.StatusPass}}
		m.resultIndex = map[resultKey]int{{Name: "TestA"}: 0}
		return m
	}

	t.Run("starting a run replaces an unrelated failure", func(t *testing.T) {
		m := completedRun()
		next, _ := m.Update(OperationErrMsg{Op: "export", Err: errors.New("export failed")})
		next, _ = next.(Model).Update(RunStartedMsg{Label: "repositories"})
		got := next.(Model)
		if got.err != nil {
			t.Fatalf("RunStartedMsg left stale error: %v", got.err)
		}
		out := stripANSI(got.renderRunConsole(100))
		if !strings.Contains(out, "Running") || strings.Contains(out, "Failed") {
			t.Fatalf("run view should be running after unrelated error:\n%s", out)
		}
	})

	t.Run("clean completion replaces an unrelated failure", func(t *testing.T) {
		m := completedRun()
		next, _ := m.Update(OperationErrMsg{Op: "orphans", Err: errors.New("list failed")})
		next, _ = next.(Model).Update(RunDoneMsg{})
		got := next.(Model)
		if got.err != nil {
			t.Fatalf("RunDoneMsg left stale error: %v", got.err)
		}
		out := stripANSI(got.renderRunConsole(100))
		if !strings.Contains(out, "Completed") || strings.Contains(out, "Failed") {
			t.Fatalf("run view should derive completed state after clean completion:\n%s", out)
		}
	})

	t.Run("late run failure remains visible", func(t *testing.T) {
		m := completedRun()
		next, _ := m.Update(RunStartedMsg{Label: "repositories"})
		next, _ = next.(Model).Update(RunDoneMsg{})
		next, _ = next.(Model).Update(OperationErrMsg{Op: "run", Err: errors.New("run failed: <redacted>")})
		got := next.(Model)
		out := stripANSI(got.renderRunConsole(100))
		if got.err == nil || !strings.Contains(out, "Failed") || !strings.Contains(out, "run failed: <redacted>") {
			t.Fatalf("late run failure was hidden:\n%s", out)
		}
	})

	t.Run("successful operation replaces stale error", func(t *testing.T) {
		m := completedRun()
		next, _ := m.Update(OperationErrMsg{Op: "export", Err: errors.New("export failed")})
		next, _ = next.(Model).Update(ReportExportDoneMsg{MarkdownPath: "report.md"})
		if got := next.(Model); got.err != nil {
			t.Fatalf("successful export left stale error: %v", got.err)
		}
	})
}

func TestFinalReviewEmptyGroupsAndTriageKeepMissionControlStructure(t *testing.T) {
	groups := fixedModel()
	groups.section = sectionGroups
	groupsOut := stripANSI(groups.renderGroupsSection(100))
	for _, want := range []string{
		"GROUPS 0", "TESTS 0", "PASSED 0", "FAILED 0", "RUNNING 0",
		"No groups discovered", "p run preflight", "provider root",
	} {
		if !strings.Contains(groupsOut, want) {
			t.Fatalf("empty groups view missing %q:\n%s", want, groupsOut)
		}
	}

	triage := fixedModel()
	triage.section = sectionTriage
	triageOut := stripANSI(triage.renderTriageSection(100))
	for _, want := range []string{
		"REAL 0", "UNSTABLE 0", "FLAKES 0", "KNOWN 0", "ELIGIBLE 0",
		"No classified failures", "t refresh triage", "K sync known issues",
	} {
		if !strings.Contains(triageOut, want) {
			t.Fatalf("empty triage view missing %q:\n%s", want, triageOut)
		}
	}
}

func TestFinalReviewPanelsAndCompleteViewsFitTinyPositiveWidths(t *testing.T) {
	for _, width := range []int{1, 2, 4, 8, 11} {
		t.Run(fmt.Sprintf("panel_%d", width), func(t *testing.T) {
			defer func() {
				if recovered := recover(); recovered != nil {
					t.Fatalf("renderPanel(%d) panicked: %v", width, recovered)
				}
			}()
			assertRenderedWidth(t, renderPanel(width, "Panel title", "Panel subtitle", "Panel body", panelAccent, true), width)
		})
	}

	views := []struct {
		name  string
		model func() Model
	}{
		{
			name: "preflight",
			model: func() Model {
				m := fixedModel()
				m.checks = []provider.Check{{Name: "identity", Status: provider.CheckOK, Detail: "ok"}}
				return m
			},
		},
		{
			name: "groups",
			model: func() Model {
				m := fixedModel()
				m.section = sectionGroups
				m.groups = []engine.Group{{Name: "repositories", Tests: []string{"TestA"}}}
				return m
			},
		},
		{
			name: "run",
			model: func() Model {
				m := fixedModel()
				m.section = sectionRun
				m.groups = []engine.Group{{Name: "repositories", Tests: []string{"TestA"}}}
				return m
			},
		},
		{
			name: "triage",
			model: func() Model {
				m := fixedModel()
				m.section = sectionTriage
				m.triageFailures = sampleFailures()
				return m
			},
		},
		{
			name: "help",
			model: func() Model {
				m := fixedModel()
				m.help.ShowAll = true
				return m
			},
		},
		{
			name: "orphan overlay",
			model: func() Model {
				m := fixedModel()
				m.orphanOverlayActive = true
				m.orphanOwner = "acme"
				return m
			},
		},
		{
			name: "sweep overlay",
			model: func() Model {
				m := fixedModel()
				m.sweepConfirmActive = true
				m.orphanOwner = "acme"
				return m
			},
		},
	}
	for _, width := range []int{1, 2, 4, 8, 11} {
		for _, tc := range views {
			t.Run(fmt.Sprintf("%s_%d", tc.name, width), func(t *testing.T) {
				defer func() {
					if recovered := recover(); recovered != nil {
						t.Fatalf("View(%d) panicked: %v", width, recovered)
					}
				}()
				m := tc.model()
				m.width = width
				assertRenderedWidth(t, m.View(), width)
			})
		}
	}
}

func TestFinalReviewASCIIProfileEmitsOnlyASCIIUI(t *testing.T) {
	assertASCII := func(t *testing.T, rendered string) {
		t.Helper()
		for _, r := range []rune(stripANSI(rendered)) {
			if r > 127 {
				t.Fatalf("ASCII profile rendered non-ASCII rune %q in:\n%s", r, rendered)
			}
		}
	}

	base := func() Model {
		long := strings.Repeat("acceptance-test-context-", 8)
		m := fixedModel()
		m.checks = []provider.Check{{Name: long, Status: provider.CheckFail, Detail: long, Fix: long}}
		m.groups = []engine.Group{{Name: long, Tests: []string{"TestA"}}}
		m.results = []engine.TestResult{{Name: "TestA", Status: provider.StatusPass, Output: []string{"redacted log"}}}
		m.resultIndex = map[resultKey]int{{Name: "TestA"}: 0}
		return m
	}
	views := []struct {
		name  string
		model func() Model
	}{
		{"preflight", base},
		{"groups", func() Model { m := base(); m.section = sectionGroups; return m }},
		{"run", func() Model { m := base(); m.section = sectionRun; return m }},
		{"triage", func() Model {
			m := base()
			m.section = sectionTriage
			m.triageFailures = sampleFailures()
			m.triageFailures[0].Test = strings.Repeat("TestAccLongName", 10)
			return m
		}},
		{"expanded help", func() Model { m := base(); m.help.ShowAll = true; return m }},
		{"groups log", func() Model {
			m := base()
			m.section = sectionGroups
			m.focus = focusLog
			m.logVP.SetContent("redacted log")
			return m
		}},
		{"orphan overlay", func() Model {
			m := base()
			m.orphanOverlayActive = true
			m.orphanOwner = strings.Repeat("acme-owner-", 12)
			m.orphans = sampleResources()
			m.orphans[0].Name = strings.Repeat("orphan-resource-", 12)
			return m
		}},
		{"sweep overlay", func() Model {
			m := base()
			m.sweepConfirmActive = true
			m.orphanOwner = strings.Repeat("acme-owner-", 12)
			m.orphans = sampleResources()
			return m
		}},
		{"file issue overlay", func() Model {
			m := base()
			m.fileIssueActive = true
			m.fileIssuePreview = IssuePreviewMsg{
				IssuesRepo:       strings.Repeat("integrations/repository-", 8),
				Title:            strings.Repeat("redacted-title-", 12),
				Classification:   engine.ClassificationReal,
				ShortFingerprint: "0123456789abcdef",
			}
			return m
		}},
	}
	for _, tc := range views {
		for _, width := range []int{4, 40, 100} {
			t.Run(fmt.Sprintf("%s_%d", tc.name, width), func(t *testing.T) {
				m := tc.model()
				m.width = width
				assertASCII(t, m.View())
			})
		}
	}
}

func TestTruncateUIUsesTerminalDisplayWidth(t *testing.T) {
	tests := []struct {
		name  string
		input string
		width int
		want  string
	}{
		{
			name:  "ANSI styled text",
			input: "\x1b[31mabcdef\x1b[0m",
			width: 4,
			want:  "abc…",
		},
		{
			name:  "wide graphemes",
			input: "界界界",
			width: 3,
			want:  "界…",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := truncateUI(tt.input, tt.width, false)
			if plain := stripANSI(got); plain != tt.want {
				t.Fatalf("truncateUI() = %q, want visible text %q", plain, tt.want)
			}
			if gotWidth := lipgloss.Width(got); gotWidth != tt.width {
				t.Fatalf("truncateUI() width = %d, want %d", gotWidth, tt.width)
			}
		})
	}
}

func TestTruncateUIZeroWidthReturnsEmpty(t *testing.T) {
	if got := truncateUI("must not escape", 0, false); got != "" {
		t.Fatalf("truncateUI width 0 = %q, want empty", got)
	}
	if got := truncateUI("must not escape", -1, true); got != "" {
		t.Fatalf("truncateUI negative width = %q, want empty", got)
	}
}

func TestFinalReviewHeaderChipsAndStatusRail(t *testing.T) {
	m := fixedModel()
	m.ascii = false
	header := stripANSI(m.renderHeader(120))
	for _, want := range []string{"‹provider: github›", "‹mode: organization›", "‹owner: my-test-org›"} {
		if !strings.Contains(header, want) {
			t.Fatalf("header missing chip %q:\n%s", want, header)
		}
	}
	if strings.Contains(header, " · ") {
		t.Fatalf("header should render context as chips, not breadcrumb prose:\n%s", header)
	}

	m.ascii = true
	m.operation = "export"
	m.status = "report exported"
	status := stripANSI(m.renderStatus(40))
	if !strings.Contains(status, "report exported") || !strings.HasSuffix(status, "export") {
		t.Fatalf("status rail should preserve left status and right operation:\n%s", status)
	}
	next, _ := m.Update(OperationErrMsg{Op: "export", Err: errors.New("export failed: <redacted>")})
	status = stripANSI(next.(Model).renderStatus(80))
	if !strings.Contains(status, "export failed: <redacted>") {
		t.Fatalf("status rail hid redacted error:\n%s", status)
	}
}

func TestFinalReviewRunProgressUsesASCIIGlyphs(t *testing.T) {
	m := fixedModel()
	m.section = sectionRun
	m.groups = []engine.Group{{Name: "repositories", Tests: []string{"TestA", "TestB"}}}
	m.results = []engine.TestResult{{Name: "TestA", Status: provider.StatusPass}}
	m.resultIndex = map[resultKey]int{{Name: "TestA"}: 0}
	out := stripANSI(m.renderRunConsole(100))
	if !strings.Contains(out, "#") || !strings.Contains(out, ".") {
		t.Fatalf("ASCII run progress should use # and .:\n%s", out)
	}
}
