package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/github/terraform-provider-tester/engine"
)

func triageCounts(failures []engine.PersistFailure) (flakes, real, unstable, known, unknown int) {
	for _, f := range failures {
		switch f.Classification {
		case engine.ClassificationFlakeConfirmed, engine.ClassificationFlakeHistorical:
			flakes++
		case engine.ClassificationRealUnstable:
			unstable++
		default:
			real++
		}
		if f.KnownIssue != 0 || f.IssueAction == "known" || f.IssueAction == "dedup" || f.IssueAction == "filed" {
			known++
		} else if engine.EligibleForIssueFiling(f) {
			unknown++
		}
	}
	return
}

func (m Model) renderTriageSection(width int) string {
	if m.triageDetailActive {
		return m.renderTriageDetail(width)
	}
	return m.renderTriageList(width)
}

func (m Model) renderTriageList(width int) string {
	width = max(1, width)
	flakes, real, unstable, known, unknown := triageCounts(m.triageFailures)
	if width < 12 {
		return renderPanel(width, "Triage", "", fmt.Sprintf("%d classified failures", len(m.triageFailures)), panelFlaky, m.ascii)
	}
	cache := "unavailable"
	if m.triageCacheAvailable {
		cache = "available"
	}

	var sb strings.Builder
	sb.WriteString(Styles.Header.Render(truncateUI("Triage", width, m.ascii)) + "\n\n")
	metrics := []metricSpec{
		{Label: "REAL", Value: fmt.Sprintf("%d", real), Tone: panelDanger},
		{Label: "UNSTABLE", Value: fmt.Sprintf("%d", unstable), Tone: panelAttention},
		{Label: "FLAKES", Value: fmt.Sprintf("%d", flakes), Tone: panelFlaky},
		{Label: "KNOWN", Value: fmt.Sprintf("%d", known), Tone: panelAccent},
		{Label: "ELIGIBLE", Value: fmt.Sprintf("%d", unknown), Tone: panelSuccess},
	}
	sb.WriteString(renderMetricRow(width, metrics, m.ascii))
	sb.WriteString("\n")
	sb.WriteString("  " + Styles.StatusLine.Render("cache") + " " + cache + "\n\n")
	if len(m.triageFailures) == 0 {
		panel := renderPanel(width, "No classified failures", "Triage is clear until a run reports failures.",
			"Refresh local triage after a run, or sync known issues.\n\nt refresh triage\nK sync known issues",
			panelSuccess, m.ascii)
		sb.WriteString(panel)
		return strings.TrimRight(sb.String(), "\n")
	}
	if width < 88 {
		sb.WriteString("    CLASSIFICATION TEST\n")
	} else {
		sb.WriteString("    CLASSIFICATION   TEST                             CLASS      ATTEMPTS FINGERPRINT       ISSUE\n")
	}

	for i, f := range m.triageFailures {
		gutter := "  "
		if i == m.triageCursor {
			arrow := ">"
			if !m.ascii {
				arrow = "❯"
			}
			gutter = lipgloss.NewStyle().Foreground(TerraformPurple).Bold(true).Render(arrow) + " "
		}
		issue := issueLabel(f)
		classification, classTone := classificationBadge(f.Classification)
		_, issueTone := issueStateBadge(f)
		if width < 88 {
			classCell := lipgloss.NewStyle().Bold(true).Foreground(toneColor(classTone)).
				Width(10).Render(truncateUI(classification, 10, m.ascii))
			firstLine := gutter + classCell + " " + truncateUI(f.Test, max(1, width-lipgloss.Width(gutter)-11), m.ascii)
			sb.WriteString(selectedTableRow(firstLine, i == m.triageCursor, width))
			sb.WriteByte('\n')
			meta := fmt.Sprintf("issue %s%sclass %s%sattempts %d%sfingerprint %s",
				issue, uiSeparator(m.ascii), fallback(f.Class, "-"), uiSeparator(m.ascii), f.Attempts, uiSeparator(m.ascii), fallback(f.ShortFingerprint, "-"))
			sb.WriteString("    " + lipgloss.NewStyle().Foreground(toneColor(issueTone)).Render(truncateUI(meta, max(1, width-4), m.ascii)))
			sb.WriteByte('\n')
			continue
		}
		classCell := lipgloss.NewStyle().Bold(true).Foreground(toneColor(classTone)).
			Width(16).Render(truncateUI(classification, 16, m.ascii))
		issueCell := lipgloss.NewStyle().Bold(true).Foreground(toneColor(issueTone)).
			Render(truncateUI(issue, max(1, width-88), m.ascii))
		line := fmt.Sprintf("%s%-16s %-32s %-10s %-8d %-16s %s",
			gutter,
			classCell,
			truncateUI(f.Test, 32, m.ascii),
			truncateUI(f.Class, 10, m.ascii),
			f.Attempts,
			truncateUI(f.ShortFingerprint, 16, m.ascii),
			issueCell,
		)
		sb.WriteString(selectedTableRow(line, i == m.triageCursor, width))
		sb.WriteByte('\n')
	}
	sb.WriteString("\n")
	sb.WriteString(renderActionStrip(width, []string{"enter detail", "t refresh triage", "K sync known issues"}, m.ascii))
	return strings.TrimRight(sb.String(), "\n")
}

func (m Model) renderTriageDetail(width int) string {
	width = max(1, width)
	if len(m.triageFailures) == 0 {
		return m.renderTriageList(width)
	}
	idx := m.triageCursor
	if idx < 0 {
		idx = 0
	}
	if idx >= len(m.triageFailures) {
		idx = len(m.triageFailures) - 1
	}
	f := m.triageFailures[idx]
	if width < 12 {
		return renderTriageDetailTiny(width, f)
	}

	known := "-"
	if f.KnownIssue != 0 {
		known = fmt.Sprintf("#%d", f.KnownIssue)
	}
	action := f.IssueAction
	if action == "" {
		action = "-"
	}
	reasons := "-"
	if len(f.Reasons) > 0 {
		reasons = strings.Join(f.Reasons, "; ")
	}
	logPath := f.LogPath
	if logPath == "" {
		logPath = "-"
	}
	canonical := f.Canonical
	if canonical == "" {
		canonical = "-"
	}
	retryable := "false"
	if f.Retryable {
		retryable = "true"
	}

	classification, classTone := classificationBadge(f.Classification)
	issueState, issueTone := issueStateBadge(f)
	panelW := max(1, width-4)
	classificationBody := strings.Join([]string{
		"test: " + truncateUI(f.Test, max(1, panelW-10), m.ascii),
		"classification: " + f.Classification,
		"class: " + fallback(f.Class, "-"),
		"attempts: " + fmt.Sprintf("%d", f.Attempts),
		"retryable: " + retryable,
	}, "\n")
	evidenceBody := strings.Join([]string{
		"fingerprint: " + fallback(f.ShortFingerprint, "-"),
		"canonical: " + canonical,
		"reasons: " + reasons,
		"log: " + logPath,
	}, "\n")
	issueBody := strings.Join([]string{
		"state: " + issueLabel(f),
		"known issue: " + known,
		"issue action: " + action,
	}, "\n")

	var sb strings.Builder
	sb.WriteString(Styles.Header.Render("Failure dossier") + "\n\n")
	sb.WriteString(renderPanel(panelW, "Classification", classification, classificationBody, classTone, m.ascii))
	sb.WriteString("\n")
	sb.WriteString(renderPanel(panelW, "Evidence", "canonical signature and fingerprint", evidenceBody, panelNeutral, m.ascii))
	sb.WriteString("\n")
	sb.WriteString(renderPanel(panelW, "Issue state", issueState, issueBody, issueTone, m.ascii))
	sb.WriteString("\n\n")

	actions := "esc back"
	if engine.EligibleForIssueFiling(f) {
		actions = "i file issue  " + actions
	}
	sb.WriteString("  " + lipgloss.NewStyle().Foreground(Muted).Render(actions))
	return strings.TrimRight(sb.String(), "\n")
}

func renderTriageDetailTiny(width int, failure engine.PersistFailure) string {
	label := "Dossier"
	if width < len(label) {
		label = "D"
	}

	back := "h"
	switch {
	case width >= 8:
		back = "esc back"
	}

	lines := []string{label}
	if engine.EligibleForIssueFiling(failure) {
		issueAction := "i"
		if width >= 6 {
			issueAction = "i file"
		}
		lines = append(lines, issueAction)
	}
	lines = append(lines, back)
	return strings.Join(lines, "\n")
}

func classificationBadge(classification string) (string, panelTone) {
	switch classification {
	case engine.ClassificationFlakeConfirmed, engine.ClassificationFlakeHistorical:
		return "FLAKE", panelFlaky
	case engine.ClassificationRealUnstable:
		return "UNSTABLE", panelAttention
	default:
		return "REAL", panelDanger
	}
}

func issueStateBadge(f engine.PersistFailure) (string, panelTone) {
	switch {
	case f.IssueAction == "dedup":
		return "DEDUP", panelAccent
	case f.IssueAction == "filed":
		return "FILED", panelSuccess
	case f.IssueAction == "known" || f.KnownIssue != 0:
		return "KNOWN", panelAccent
	case engine.EligibleForIssueFiling(f):
		return "ELIGIBLE", panelSuccess
	default:
		return "BLOCKED", panelNeutral
	}
}

func fallback(value, empty string) string {
	if value == "" {
		return empty
	}
	return value
}

func issueLabel(f engine.PersistFailure) string {
	switch {
	case f.IssueAction == "dedup":
		return "dedup"
	case f.IssueAction == "filed":
		return "filed"
	case f.IssueAction == "known" || f.KnownIssue != 0:
		if f.KnownIssue != 0 {
			return fmt.Sprintf("known #%d", f.KnownIssue)
		}
		return "known"
	case engine.EligibleForIssueFiling(f):
		return "eligible"
	default:
		return "-"
	}
}
