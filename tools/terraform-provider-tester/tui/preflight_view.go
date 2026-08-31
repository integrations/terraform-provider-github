package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/github/terraform-provider-tester/provider"
)

type preflightSummary struct {
	OK, Warn, Fail int
}

func (m Model) preflightSummary() preflightSummary {
	var summary preflightSummary
	for _, check := range m.checks {
		switch check.Status {
		case provider.CheckOK:
			summary.OK++
		case provider.CheckWarn:
			summary.Warn++
		default:
			summary.Fail++
		}
	}
	return summary
}

func (m Model) firstPreflightAction() (string, panelTone) {
	for _, check := range m.checks {
		if check.Status == provider.CheckFail {
			if check.Fix != "" {
				return check.Fix, panelDanger
			}
			return check.Detail, panelDanger
		}
	}
	for _, check := range m.checks {
		if check.Status == provider.CheckWarn {
			if check.Fix != "" {
				return check.Fix, panelAttention
			}
			return check.Detail, panelAttention
		}
	}
	return "Environment is ready. Open Groups to select a run.", panelSuccess
}

func (m Model) renderPreflight(width int) string {
	width = max(1, width)
	summary := m.preflightSummary()
	metrics := renderMetricRow(width, []metricSpec{
		{Label: "READY", Value: fmt.Sprint(summary.OK), Tone: panelSuccess},
		{Label: "WARNINGS", Value: fmt.Sprint(summary.Warn), Tone: panelAttention},
		{Label: "BLOCKED", Value: fmt.Sprint(summary.Fail), Tone: panelDanger},
		{Label: "MODE", Value: m.mode, Tone: panelAccent},
	}, m.ascii)

	if len(m.checks) == 0 {
		body := strings.Join([]string{
			Tagline,
			"",
			"Preflight has not run yet.",
			"Run prerequisite checks before launching acceptance tests.",
			"",
			"p run preflight",
		}, "\n")
		panel := renderPanel(width, "Prerequisite checks", "No check data yet.", truncateBlockUI(body, max(1, width-4), m.ascii), panelAttention, m.ascii)
		return stackMissionBlocks(width, m.ascii, metrics, panel)
	}

	checksTone := panelSuccess
	if summary.Fail > 0 {
		checksTone = panelDanger
	} else if summary.Warn > 0 {
		checksTone = panelAttention
	}

	action, actionTone := m.firstPreflightAction()
	actionBody := strings.Join([]string{
		truncateUI(action, max(1, width-4), m.ascii),
		"",
		"m mode  v variables  p rerun",
	}, "\n")

	mode := layoutForWidth(width)
	if mode == layoutWide {
		gap := 2
		leftW := max(28, (width-gap)*2/3)
		rightW := max(24, width-gap-leftW)
		if leftW+gap+rightW > width {
			leftW = max(12, width-gap-rightW)
		}
		checks := renderPanel(leftW, "Prerequisite checks", preflightSubtitle(summary, m.ascii), m.renderPreflightCheckList(max(1, leftW-4)), checksTone, m.ascii)
		next := renderPanel(rightW, "Next action", "Safe before provider writes.", truncateBlockUI(actionBody, max(1, rightW-4), m.ascii), actionTone, m.ascii)
		return stackMissionBlocks(width, m.ascii, metrics, lipgloss.JoinHorizontal(lipgloss.Top, checks, strings.Repeat(" ", gap), next))
	}

	checks := renderPanel(width, "Prerequisite checks", preflightSubtitle(summary, m.ascii), m.renderPreflightCheckList(max(1, width-4)), checksTone, m.ascii)
	next := renderPanel(width, "Next action", "Safe before provider writes.", truncateBlockUI(actionBody, max(1, width-4), m.ascii), actionTone, m.ascii)
	return stackMissionBlocks(width, m.ascii, metrics, checks, next)
}

func preflightSubtitle(summary preflightSummary, ascii bool) string {
	return fmt.Sprintf("%d ready%s%d warnings%s%d blocked", summary.OK, uiSeparator(ascii), summary.Warn, uiSeparator(ascii), summary.Fail)
}

func (m Model) renderPreflightCheckList(width int) string {
	width = max(1, width)
	maxName := len("CHECK")
	for _, check := range m.checks {
		if w := lipgloss.Width(check.Name); w > maxName {
			maxName = w
		}
	}
	maxName = min(maxName, max(8, width/3))

	lines := make([]string, 0, len(m.checks)*2)
	for _, check := range m.checks {
		glyph := checkGlyph(check.Status, m.ascii)
		glyphStr := glyph.Glyph
		if color, ok := glyph.Color.(lipgloss.TerminalColor); ok {
			glyphStr = lipgloss.NewStyle().Foreground(color).Render(glyph.Glyph)
		}
		status := preflightStatusText(check.Status)
		prefix := fmt.Sprintf("%s %-7s %-*s", glyphStr, status, maxName, truncateUI(check.Name, maxName, m.ascii))
		detailW := max(1, width-lipgloss.Width(prefix)-1)
		line := prefix
		if check.Detail != "" && detailW > 1 {
			line += " " + truncateUI(check.Detail, detailW, m.ascii)
		}
		lines = append(lines, truncateUI(line, width, m.ascii))
		if check.Fix != "" {
			lines = append(lines, truncateUI("  fix: "+check.Fix, width, m.ascii))
		}
	}
	return strings.Join(lines, "\n")
}

func preflightStatusText(status provider.CheckStatus) string {
	switch status {
	case provider.CheckOK:
		return "READY"
	case provider.CheckWarn:
		return "WARNING"
	default:
		return "BLOCKED"
	}
}

func stackMissionBlocks(width int, ascii bool, blocks ...string) string {
	width = max(1, width)
	out := make([]string, 0, len(blocks))
	for _, block := range blocks {
		block = strings.TrimRight(block, "\n")
		if block == "" {
			continue
		}
		out = append(out, truncateBlockUI(block, width, ascii))
	}
	return strings.Join(out, "\n\n")
}
