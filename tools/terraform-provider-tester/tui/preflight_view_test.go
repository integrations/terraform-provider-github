package tui

import (
	"strings"
	"testing"

	"github.com/charmbracelet/lipgloss"

	"github.com/github/terraform-provider-tester/provider"
)

func TestPreflightSummaryCountsStatuses(t *testing.T) {
	m := fixedModel()
	m.checks = []provider.Check{
		{Status: provider.CheckOK},
		{Status: provider.CheckWarn},
		{Status: provider.CheckFail},
		{Status: provider.CheckFail},
	}
	if got := m.preflightSummary(); got.OK != 1 || got.Warn != 1 || got.Fail != 2 {
		t.Fatalf("summary = %#v", got)
	}
}

func TestPreflightViewShowsMetricsAndFirstBlockingFix(t *testing.T) {
	m := fixedModel()
	m.checks = []provider.Check{
		{Name: "identity", Status: provider.CheckOK, Detail: "authenticated"},
		{Name: "owner", Status: provider.CheckFail, Detail: "GITHUB_OWNER not set", Fix: "set GITHUB_OWNER=<org>"},
		{Name: "template", Status: provider.CheckFail, Detail: "template missing", Fix: "create terraform-template-module"},
	}
	got := stripANSI(m.renderPreflight(120))
	for _, want := range []string{"READY 1", "BLOCKED 2", "Prerequisite checks", "Next action", "set GITHUB_OWNER=<org>", "m mode", "v variables", "p rerun"} {
		if !strings.Contains(got, want) {
			t.Fatalf("preflight missing %q:\n%s", want, got)
		}
	}
}

func TestPreflightEmptyStateHasNextAction(t *testing.T) {
	got := stripANSI(fixedModel().renderPreflight(80))
	if !strings.Contains(got, "Preflight has not run yet") || !strings.Contains(got, "p run preflight") {
		t.Fatalf("unexpected empty state:\n%s", got)
	}
}

func TestPreflightViewFitsRepresentativeWidths(t *testing.T) {
	m := fixedModel()
	m.checks = []provider.Check{
		{Name: "identity", Status: provider.CheckOK, Detail: "authenticated as a very long identity that must be clipped safely"},
		{Name: "enterprise-slug", Status: provider.CheckWarn, Detail: "slug is optional for this mode but recommended", Fix: "set GITHUB_ENTERPRISE_SLUG=<slug> when running enterprise groups"},
	}
	for _, width := range []int{48, 80, 120} {
		for i, line := range strings.Split(stripANSI(m.renderPreflight(width)), "\n") {
			if got := lipgloss.Width(line); got > width {
				t.Fatalf("width %d line %d is %d columns: %q", width, i+1, got, line)
			}
		}
	}
}
