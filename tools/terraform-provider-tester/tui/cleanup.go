package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) renderOrphanOverlay(width int) string {
	label := func(s string) string { return Styles.StatusLine.Render(s) }
	bodyW := max(1, width-4)
	var sb strings.Builder
	sb.WriteString("  " + label("owner") + " " + truncateUI(m.orphanOwner, max(1, bodyW-10), m.ascii) + "\n")
	sb.WriteString("  " + label("count") + " " + fmt.Sprintf("%d", len(m.orphans)) + "\n\n")
	sb.WriteString("    KIND            NAME                            URL\n")
	for _, r := range m.orphans {
		sb.WriteString(fmt.Sprintf("  %-14s  %-30s  %s\n",
			truncateUI(r.Kind, 14, m.ascii),
			truncateUI(r.Name, 30, m.ascii),
			truncateUI(r.URL, max(1, bodyW-52), m.ascii),
		))
	}
	hint := "esc close"
	if len(m.orphans) > 0 {
		hint = "s sweep  " + hint
	}
	sb.WriteString("\n  " + lipgloss.NewStyle().Foreground(Muted).Render(hint))
	return renderPanel(width, "Orphan resources", "", strings.TrimRight(sb.String(), "\n"), panelNeutral, m.ascii)
}

func (m Model) renderSweepConfirm(width int) string {
	label := func(s string) string { return Styles.StatusLine.Render(s) }
	phrase := sweepPhrase(m.orphanOwner)
	bodyW := max(1, width-4)
	var sb strings.Builder
	sb.WriteString("  " + label("owner ") + " " + truncateUI(m.orphanOwner, max(1, bodyW-11), m.ascii) + "\n")
	sb.WriteString("  " + label("phrase") + " " + phrase + "\n\n")
	sb.WriteString("  Targets:\n")
	for _, r := range m.orphans {
		target := fmt.Sprintf("%s %s %s", r.Kind, r.Name, r.URL)
		sb.WriteString("  - " + truncateUI(target, max(1, bodyW-4), m.ascii) + "\n")
	}
	sb.WriteString("\n  Type " + phrase + " to confirm:\n")
	sb.WriteString("  " + m.sweepInput.View() + "\n\n")
	sb.WriteString("  " + lipgloss.NewStyle().Foreground(Muted).Render("enter confirm  esc close"))
	return renderPanel(width, "Confirm sweep", destructiveSubtitle(), strings.TrimRight(sb.String(), "\n"), panelDanger, m.ascii)
}

func (m Model) renderFileIssueConfirm(width int) string {
	label := func(s string) string { return Styles.StatusLine.Render(s) }
	preview := m.fileIssuePreview
	phrase := fileIssuePhrase(preview.ShortFingerprint)
	bodyW := max(1, width-4)
	labels := "-"
	if len(preview.Labels) > 0 {
		labels = strings.Join(preview.Labels, ", ")
	}
	var sb strings.Builder
	sb.WriteString("  " + label("repo          ") + " " + truncateUI(preview.IssuesRepo, max(1, bodyW-18), m.ascii) + "\n")
	sb.WriteString("  " + label("title redacted") + " " + truncateUI(preview.Title, max(1, bodyW-18), m.ascii) + "\n")
	sb.WriteString("  " + label("labels        ") + " " + truncateUI(labels, max(1, bodyW-18), m.ascii) + "\n")
	sb.WriteString("  " + label("classification") + " " + truncateUI(preview.Classification, max(1, bodyW-18), m.ascii) + "\n")
	sb.WriteString("  " + label("fingerprint   ") + " " + truncateUI(preview.ShortFingerprint, max(1, bodyW-18), m.ascii) + "\n\n")
	sb.WriteString("  Type " + phrase + " to confirm:\n")
	sb.WriteString("  " + m.fileIssueInput.View() + "\n\n")
	sb.WriteString("  " + lipgloss.NewStyle().Foreground(Muted).Render("enter confirm  esc close"))
	return renderPanel(width, "File issue", destructiveSubtitle(), strings.TrimRight(sb.String(), "\n"), panelDanger, m.ascii)
}

func (m Model) renderResultOverlay(width int) string {
	bodyW := max(1, width-4)
	var sb strings.Builder
	for _, line := range m.resultLines {
		sb.WriteString("  " + truncateUI(line, max(1, bodyW-4), m.ascii) + "\n")
	}
	sb.WriteString("\n  " + lipgloss.NewStyle().Foreground(Muted).Render("esc close"))
	return renderPanel(width, m.resultTitle, "", strings.TrimRight(sb.String(), "\n"), m.resultTone, m.ascii)
}

func destructiveSubtitle() string {
	return "Destructive action. Review the snapshot and type the exact phrase."
}
