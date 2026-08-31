package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type runSummary struct {
	Total, Passed, Failed, Running, NotRun int
	Duration                               float64
}

func (m Model) runSummary() runSummary {
	var out runSummary
	for _, group := range m.groups {
		summary := summarize(group, m.results)
		out.Total += summary.Total
		out.Passed += summary.Passed
		out.Failed += summary.Failed
		out.Running += summary.Running
		out.NotRun += summary.NotRun
		out.Duration += summary.Duration
	}
	return out
}

func (m Model) renderRunConsole(width int) string {
	width = max(1, width)
	summary := m.runSummary()
	metrics := renderMetricRow(width, []metricSpec{
		{Label: "PASSED", Value: fmt.Sprint(summary.Passed), Tone: panelSuccess},
		{Label: "FAILED", Value: fmt.Sprint(summary.Failed), Tone: panelDanger},
		{Label: "RUNNING", Value: fmt.Sprint(summary.Running), Tone: panelAttention},
		{Label: "REMAINING", Value: fmt.Sprint(summary.NotRun), Tone: panelNeutral},
	}, m.ascii)

	state, stateTone, stateDetail := m.runVisualState(summary)
	progressBody := m.renderRunProgressBody(summary, state, stateDetail, max(1, width-4))
	progress := renderPanel(width, "Run progress", state, progressBody, stateTone, m.ascii)

	targetBody := m.renderRunTargetBody(max(1, width-4))
	actionsBody := m.renderRunActionsBody(max(1, width-4))

	mode := layoutForWidth(width)
	if mode == layoutWide {
		gap := 2
		leftW := max(34, (width-gap)/2)
		rightW := max(34, width-gap-leftW)
		if leftW+gap+rightW > width {
			leftW = max(12, width-gap-rightW)
		}
		target := renderPanel(leftW, "Target and scope", "Provider acceptance-test target.", m.renderRunTargetBody(max(1, leftW-4)), panelAccent, m.ascii)
		actions := renderPanel(rightW, "Safe actions", "No automatic commands run from this view.", m.renderRunActionsBody(max(1, rightW-4)), panelNeutral, m.ascii)
		return stackMissionBlocks(width, m.ascii, metrics, progress, lipgloss.JoinHorizontal(lipgloss.Top, target, strings.Repeat(" ", gap), actions))
	}

	target := renderPanel(width, "Target and scope", "Provider acceptance-test target.", targetBody, panelAccent, m.ascii)
	actions := renderPanel(width, "Safe actions", "No automatic commands run from this view.", actionsBody, panelNeutral, m.ascii)
	return stackMissionBlocks(width, m.ascii, metrics, progress, target, actions)
}

func (m Model) runVisualState(summary runSummary) (string, panelTone, string) {
	switch {
	case m.err != nil && m.errOperation == "run" && m.errIsOperationFailure:
		return "Failed", panelDanger, m.err.Error()
	case m.running:
		label := strings.TrimSpace(m.runLabel)
		if label == "" {
			label = "selection"
		}
		return "Running", panelAttention, "running " + label
	case summary.Total == 0:
		return "Ready for discovery", panelAttention, "No tests discovered."
	case summary.Running > 0:
		return "Interrupted", panelAttention, "run stopped while tests were still marked running"
	case summary.Total > 0 && summary.Passed+summary.Failed == 0:
		return "Ready to run", panelAccent, "choose a group or test, then start the run"
	case summary.NotRun == 0:
		if summary.Failed > 0 {
			return "Completed", panelDanger, "completed with failures ready for retry or triage"
		}
		return "Completed", panelSuccess, "all discovered tests reached a terminal state"
	default:
		return "Ready to run", panelAccent, "partial results are loaded; continue or retry selection"
	}
}

func (m Model) renderRunProgressBody(summary runSummary, state, detail string, width int) string {
	width = max(1, width)
	completed := summary.Passed + summary.Failed
	pct := 0
	if summary.Total > 0 {
		pct = (completed * 100) / summary.Total
	}
	barWidth := min(28, max(10, width-18))
	status := aggregateStatus(summary.Passed, summary.Failed, summary.Running, summary.Total)
	lines := []string{
		truncateUI(fmt.Sprintf("%s%s%d of %d terminal%s%d%%", state, uiSeparator(m.ascii), completed, summary.Total, uiSeparator(m.ascii), pct), width, m.ascii),
		truncateUI(fmt.Sprintf("%d total%s%d passed%s%d failed%s%d running%s%d not run", summary.Total, uiSeparator(m.ascii), summary.Passed, uiSeparator(m.ascii), summary.Failed, uiSeparator(m.ascii), summary.Running, uiSeparator(m.ascii), summary.NotRun), width, m.ascii),
	}
	if summary.Total == 0 {
		lines = append(lines, truncateUI("2 open Groups to discover and select acceptance tests.", width, m.ascii))
	} else {
		lines = append(lines, truncateUI(renderProgressBar(completed, summary.Total, barWidth, status, m.ascii)+" "+fmt.Sprintf("%d%%", pct), width, m.ascii))
	}
	if detail != "" {
		lines = append(lines, truncateUI(detail, width, m.ascii))
	}
	lines = append(lines, truncateUI(fmt.Sprintf("duration %.1fs", summary.Duration), width, m.ascii))
	return strings.Join(lines, "\n")
}

func (m Model) renderRunTargetBody(width int) string {
	width = max(1, width)
	scope := "mode: " + m.mode
	if m.owner != "" {
		scope += uiSeparator(m.ascii) + "owner: " + m.owner
	}
	selection := strings.TrimSpace(m.runLabel)
	if selection == "" {
		selection = "current Groups selection"
	}
	command := "go test ./github -run '<selection>' -json -timeout 120m"
	return strings.Join([]string{
		missionField("target", TargetRepo, width, m.ascii),
		missionField("scope", scope, width, m.ascii),
		missionField("selection", selection, width, m.ascii),
		missionField("command", command, width, m.ascii),
	}, "\n")
}

func (m Model) renderRunActionsBody(width int) string {
	width = max(1, width)
	activity := "Last activity: idle"
	if m.running {
		parts := []string{"Last activity: running"}
		if m.runLabel != "" {
			parts[0] += " " + m.runLabel
		}
		parts = append(parts, "elapsed "+formatRunDuration(m.runElapsed))
		if m.stalledFor >= stallThreshold {
			parts = append(parts, "quiet "+formatRunDuration(m.stalledFor))
		}
		activity = strings.Join(parts, uiSeparator(m.ascii))
	} else if m.err != nil && m.errOperation == "run" && m.errIsOperationFailure {
		activity = "Last activity: failed - " + m.err.Error()
	} else if len(m.results) > 0 {
		activity = "Last activity: results loaded"
	}
	actions := []string{
		truncateUI("r run selection", width, m.ascii),
		truncateUI("R retry failures", width, m.ascii),
		truncateUI("e export report", width, m.ascii),
		truncateUI("docs "+DocsPath, width, m.ascii),
		"",
		truncateUI(activity, width, m.ascii),
	}
	if m.runSummary().Total == 0 {
		actions = []string{
			truncateUI("2 open Groups", width, m.ascii),
			truncateUI("p run preflight", width, m.ascii),
			truncateUI("docs "+DocsPath, width, m.ascii),
			"",
			truncateUI(activity, width, m.ascii),
		}
	}
	return strings.Join(actions, "\n")
}

func missionField(label, value string, width int, ascii bool) string {
	width = max(1, width)
	prefix := lipgloss.NewStyle().Foreground(Muted).Render(label + " ")
	valueW := max(1, width-lipgloss.Width(prefix))
	return truncateUI(prefix+truncateUI(value, valueW, ascii), width, ascii)
}
