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

func runViewModel() Model {
	m := fixedModel()
	m.groups = []engine.Group{{Name: "repositories", Tests: []string{"TestA", "TestB", "TestC", "TestD"}}}
	m.results = []engine.TestResult{
		{Name: "TestA", Status: provider.StatusPass, Elapsed: 1.2},
		{Name: "TestB", Status: provider.StatusFail, Elapsed: 2.3},
		{Name: "TestC", Status: provider.StatusRunning},
	}
	m.resultIndex = map[resultKey]int{
		{Name: "TestA"}: 0,
		{Name: "TestB"}: 1,
		{Name: "TestC"}: 2,
	}
	return m
}

func TestRunSummaryDerivesCurrentState(t *testing.T) {
	got := runViewModel().runSummary()
	if got.Total != 4 || got.Passed != 1 || got.Failed != 1 || got.Running != 1 || got.NotRun != 1 {
		t.Fatalf("summary = %#v", got)
	}
}

func TestRunViewShowsProgressScopeAndActions(t *testing.T) {
	m := runViewModel()
	m.running = true
	m.runLabel = "repositories"
	got := stripANSI(m.renderRunConsole(120))
	for _, want := range []string{"Run progress", "50%", "PASSED 1", "FAILED 1", "RUNNING 1", "REMAINING 1", "Target and scope", "Safe actions", "repositories"} {
		if !strings.Contains(got, want) {
			t.Fatalf("run view missing %q:\n%s", want, got)
		}
	}
}

func TestRunEmptyStatePointsToGroups(t *testing.T) {
	got := stripANSI(fixedModel().renderRunConsole(80))
	if !strings.Contains(got, "No tests discovered") || !strings.Contains(got, "2 open Groups") {
		t.Fatalf("unexpected Run empty state:\n%s", got)
	}
}

func TestRunEmptyStateShowsMissionControlPanelsAtWidths(t *testing.T) {
	for _, width := range []int{48, 80, 120} {
		t.Run(fmt.Sprintf("%d", width), func(t *testing.T) {
			got := stripANSI(fixedModel().renderRunConsole(width))
			for _, want := range []string{
				"Run progress",
				"Target and scope",
				"Safe actions",
				"No tests discovered",
				"2 open Groups",
				"target",
				"scope",
				"selection",
				"command",
				"No automatic commands run from this view.",
				"Last activity: idle",
			} {
				if !strings.Contains(got, want) {
					t.Fatalf("width %d missing %q:\n%s", width, want, got)
				}
			}
			if strings.Contains(got, "r run selection") {
				t.Fatalf("width %d should not advertise a run action before tests are discovered:\n%s", width, got)
			}
			for i, line := range strings.Split(got, "\n") {
				if gotWidth := lipgloss.Width(line); gotWidth > width {
					t.Fatalf("width %d line %d is %d columns: %q", width, i+1, gotWidth, line)
				}
			}
		})
	}
}

func TestRunViewShowsZeroElapsedWhileRunning(t *testing.T) {
	m := runViewModel()
	m.running = true
	m.runLabel = "repositories"
	m.runElapsed = 0
	got := stripANSI(m.renderRunConsole(100))
	if !strings.Contains(got, "elapsed 0s") {
		t.Fatalf("running run view should show zero elapsed:\n%s", got)
	}
}

func TestRunViewShowsLifecycleStates(t *testing.T) {
	tests := []struct {
		name  string
		model func() Model
		want  string
	}{
		{
			name: "idle",
			model: func() Model {
				m := fixedModel()
				m.groups = []engine.Group{{Name: "repositories", Tests: []string{"TestA"}}}
				return m
			},
			want: "Ready to run",
		},
		{
			name: "running",
			model: func() Model {
				m := runViewModel()
				m.running = true
				m.runLabel = "repositories"
				return m
			},
			want: "Running",
		},
		{
			name: "completed",
			model: func() Model {
				m := fixedModel()
				m.groups = []engine.Group{{Name: "repositories", Tests: []string{"TestA", "TestB"}}}
				m.results = []engine.TestResult{{Name: "TestA", Status: provider.StatusPass}, {Name: "TestB", Status: provider.StatusPass}}
				return m
			},
			want: "Completed",
		},
		{
			name: "interrupted",
			model: func() Model {
				m := runViewModel()
				m.running = false
				return m
			},
			want: "Interrupted",
		},
		{
			name: "failed",
			model: func() Model {
				m := runViewModel()
				next, _ := m.Update(OperationErrMsg{Op: "run", Err: errors.New("go test exited 1")})
				return next.(Model)
			},
			want: "Failed",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := stripANSI(tc.model().renderRunConsole(100)); !strings.Contains(got, tc.want) {
				t.Fatalf("run state missing %q:\n%s", tc.want, got)
			}
		})
	}
}

func TestRunViewFitsRepresentativeWidths(t *testing.T) {
	m := runViewModel()
	m.running = true
	m.runLabel = "repositories"
	for _, width := range []int{48, 80, 120} {
		for i, line := range strings.Split(stripANSI(m.renderRunConsole(width)), "\n") {
			if got := lipgloss.Width(line); got > width {
				t.Fatalf("width %d line %d is %d columns: %q", width, i+1, got, line)
			}
		}
	}
}

func TestRunConsoleNarrowASCIIUsesOnlyASCIIRunes(t *testing.T) {
	m := fixedModel()
	m.ascii = true
	m.runLabel = strings.Repeat("selected acceptance target ", 4)
	m.err = errors.New(strings.Repeat("long failed run status ", 4))

	got := stripANSI(m.renderRunConsole(12))
	for _, r := range got {
		if r > 127 {
			t.Fatalf("narrow ASCII Run console rendered non-ASCII rune %q:\n%s", r, got)
		}
	}
}
