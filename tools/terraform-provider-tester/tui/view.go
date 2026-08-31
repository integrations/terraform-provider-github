package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/github/terraform-provider-tester/engine"
	"github.com/github/terraform-provider-tester/provider"
)

type contextualKeyMap struct {
	navigation []string
	section    []string
	global     []string
}

func (k contextualKeyMap) footerHints() []string {
	hints := append([]string{}, k.navigation...)
	hints = append(hints, k.section...)
	hints = append(hints, k.global...)
	return hints
}

// View renders the TUI chrome as a plain string.
//
// Design contract (TUI_STYLE §3):
//
//	header bar  (1 row)
//	tab bar     (1 row)
//	body        (flex)
//	status line (1 row)
//	help footer (1 row)
//
// The View is a pure function: same Model → same string.
// time.Now() is NEVER called here. The rate-reset time is formatted as an
// absolute UTC timestamp (e.g. "21:00 UTC") so the output is deterministic
// and golden-testable without a clock field.
func (m Model) View() string {
	width := m.width
	if width <= 0 {
		width = 80
	}

	// One-shot ignition splash: while active it owns the whole screen.
	if m.introActive {
		height := m.height
		if height <= 0 {
			height = 24
		}
		return renderIntro(m.introFrame, width, height, m.version, m.ascii)
	}

	header := m.renderHeader(width)
	tabs := m.renderTabs(width)
	status := m.renderStatus(width)
	footer := m.renderFooter(width)

	var body string
	if m.pickerActive {
		body = m.renderPickerOverlay(width)
	} else if m.editorActive {
		body = m.renderEditorOverlay(width)
	} else if m.sweepConfirmActive {
		body = m.renderSweepConfirm(width)
	} else if m.orphanOverlayActive {
		body = m.renderOrphanOverlay(width)
	} else if m.fileIssueActive {
		body = m.renderFileIssueConfirm(width)
	} else if m.resultOverlayActive {
		body = m.renderResultOverlay(width)
	} else if m.help.ShowAll && !m.resumePrompt {
		body = m.renderHelpPanel(width)
	} else {
		body = m.renderBody(width)
	}

	parts := []string{header, tabs}
	// The resume prompt is rendered globally (not per-section) so it stays
	// visible while its key interception is active in any section.
	if m.resumePrompt {
		parts = append(parts, m.renderResumeBanner(width))
	}
	parts = append(parts, body, status, footer)

	return lipgloss.JoinVertical(lipgloss.Left, parts...)
}

// renderBody renders the section body.
func (m Model) renderBody(width int) string {
	switch m.section {
	case sectionPreflight:
		return m.renderPreflight(width)
	case sectionGroups:
		return m.renderGroupsSection(width)
	case sectionTriage:
		return m.renderTriageSection(width)
	default:
		return m.renderRunConsole(width)
	}
}

// renderResumeBanner returns a one-line banner offering to resume a prior run.
// The timestamp is formatted as a fixed UTC string so View remains pure
// (no time.Now() call).
func (m Model) renderResumeBanner(width int) string {
	when := m.resumeWhen.UTC().Format("2006-01-02 15:04 UTC")
	resume := "↻"
	if m.ascii {
		resume = "resume"
	}
	banner := fmt.Sprintf(
		"%s resume last run (%s) - %d failed, %d not run%senter resume%sesc dismiss",
		resume, when, m.resumeFailed, m.resumeNotRun, uiSeparator(m.ascii), uiSeparator(m.ascii),
	)
	banner = truncateUI(banner, width, m.ascii)
	return lipgloss.NewStyle().Foreground(Accent).Render(banner)
}

// renderGroupsSection dispatches to the correct drill-down level view.
func (m Model) renderGroupsSection(width int) string {
	switch m.focus {
	case focusTests:
		return m.renderTestsDrilldown(width)
	case focusLog:
		return m.renderLogViewport(width)
	default:
		return m.renderGroups(width)
	}
}

// groupNameMax caps the group-name column so an unusually long name cannot
// push the numbers and progress bars off-screen on a narrow terminal.
const groupNameMax = 22

// renderGroups renders the groups master table.
// Columns: cursor · status glyph · group name · passed/total · duration ·
// progress bar · failed annotation.
//
// Column widths are measured from the data and applied with lipgloss .Width()
// (which is ANSI-aware) rather than fmt's %-Ns, so the cells stay aligned even
// once they carry color escapes. The result: every progress bar starts in the
// same column and the passed/total counts right-align into a clean number
// column.
func (m Model) renderGroups(width int) string {
	width = max(1, width)
	if len(m.groups) == 0 {
		metrics := renderMetricRow(width, []metricSpec{
			{Label: "GROUPS", Value: "0", Tone: panelAccent},
			{Label: "TESTS", Value: "0", Tone: panelNeutral},
			{Label: "PASSED", Value: "0", Tone: panelSuccess},
			{Label: "FAILED", Value: "0", Tone: panelDanger},
			{Label: "RUNNING", Value: "0", Tone: panelAttention},
		}, m.ascii)
		panel := renderPanel(width, "No groups discovered", "Discovery has not returned any provider test groups.",
			"Set a provider root, then run discovery through preflight.\n\np run preflight",
			panelAttention, m.ascii)
		return stackMissionBlocks(width, m.ascii, metrics, panel)
	}
	if width < 12 {
		body := fmt.Sprintf("%d groups available.\nUse enter to inspect tests.", len(m.groups))
		return renderPanel(width, "Groups", "", body, panelAccent, m.ascii)
	}

	mode := layoutForWidth(width)
	barWidth := 14
	if mode == layoutCompact {
		barWidth = 10
	}
	gap := "   " // three-space column gutter for breathing room
	compact := mode == layoutCompact
	if compact {
		gap = "  "
	}

	// First pass: summarize each group once, measure the dynamic columns, and
	// accumulate the overall totals for the footer row. Column widths start at
	// their header-label widths so a header never overflows its column.
	sums := make([]groupSummary, len(m.groups))
	nameW, countW, durW := len("GROUP"), len("PASS"), len("TIME")
	var totPassed, totFailed, totRunning, totTotal, maxFailed int
	var totDur float64
	for i, g := range m.groups {
		sum := summarize(g, m.results)
		sums[i] = sum
		if w := lipgloss.Width(g.Name); w > nameW {
			nameW = w
		}
		if w := len(fmt.Sprintf("%d/%d", sum.Passed, sum.Total)); w > countW {
			countW = w
		}
		if w := len(fmt.Sprintf("%.1fs", sum.Duration)); w > durW {
			durW = w
		}
		totPassed += sum.Passed
		totFailed += sum.Failed
		totRunning += sum.Running
		totTotal += sum.Total
		totDur += sum.Duration
		if sum.Failed > maxFailed {
			maxFailed = sum.Failed
		}
	}
	if nameW > groupNameMax {
		nameW = groupNameMax
	}
	// The totals row can be wider than any single group (its count and summed
	// duration aggregate every group), so widen the columns to fit it too or
	// the TOTAL row would shift out of alignment with the data rows.
	if w := len(fmt.Sprintf("%d/%d", totPassed, totTotal)); w > countW {
		countW = w
	}
	if w := len(fmt.Sprintf("%.1fs", totDur)); w > durW {
		durW = w
	}
	maxFailedLabelW := 0
	if maxFailed > 0 || totFailed > 0 {
		maxFailedLabelW = len(groupFailureText(max(maxFailed, totFailed), compact))
	}
	pctW := len("100%")
	compactBaseW := lipgloss.Width("  "+" ") + countW + durW + barWidth + pctW + 4*len(gap)
	if maxFailedLabelW > 0 {
		compactBaseW += maxFailedLabelW + len(gap)
	}
	if compact && compactBaseW+nameW >= width {
		nameW = max(len("TOTAL"), width-compactBaseW-1)
	}

	// A blank two-column gutter plus a blank one-column glyph slot keeps the
	// header and totals rows aligned with the data rows below them.
	pad := "  " + " " + gap

	var sb strings.Builder
	metrics := []metricSpec{
		{Label: "GROUPS", Value: fmt.Sprintf("%d", len(m.groups)), Tone: panelAccent},
		{Label: "TESTS", Value: fmt.Sprintf("%d", totTotal), Tone: panelNeutral},
		{Label: "PASSED", Value: fmt.Sprintf("%d", totPassed), Tone: panelSuccess},
		{Label: "FAILED", Value: fmt.Sprintf("%d", totFailed), Tone: panelDanger},
		{Label: "RUNNING", Value: fmt.Sprintf("%d", totRunning), Tone: panelAttention},
	}
	sb.WriteString(renderMetricRow(width, metrics, m.ascii))
	sb.WriteString("\n\n")

	// Header row - dim labels that explain the columns.
	hdr := lipgloss.NewStyle().Foreground(Muted)
	sb.WriteString(pad +
		hdr.Width(nameW).Render("GROUP") + gap +
		hdr.Width(countW).Align(lipgloss.Right).Render("PASS") + gap +
		hdr.Width(durW).Align(lipgloss.Right).Render("TIME") + gap +
		hdr.Render("PROGRESS"))
	sb.WriteByte('\n')

	for i, g := range m.groups {
		sum := sums[i]

		// Cursor gutter - a gh-style accent arrow marks the selected row and
		// always occupies two columns so the table never shifts between rows.
		gutter := "  "
		if i == m.groupCursor {
			arrow := "❯"
			if m.ascii {
				arrow = ">"
			}
			gutter = lipgloss.NewStyle().Foreground(TerraformPurple).Bold(true).Render(arrow) + " "
		}

		glyphStr := m.renderRowGlyph(sum.Status)
		name := lipgloss.NewStyle().Foreground(rowStatusColor(sum.Status)).Width(nameW).
			Render(truncateUI(g.Name, nameW, m.ascii))
		counts := lipgloss.NewStyle().Bold(true).Width(countW).Align(lipgloss.Right).
			Render(fmt.Sprintf("%d/%d", sum.Passed, sum.Total))
		dur := lipgloss.NewStyle().Foreground(Muted).Width(durW).Align(lipgloss.Right).
			Render(fmt.Sprintf("%.1fs", sum.Duration))
		bar := renderProgressBar(sum.Passed, sum.Total, barWidth, sum.Status, m.ascii)

		line := gutter + glyphStr + gap + name + gap + counts + gap + dur + gap + bar
		if sum.Failed > 0 {
			line += gap + lipgloss.NewStyle().Foreground(Danger).
				Render(groupFailureText(sum.Failed, compact))
		}
		sb.WriteString(selectedTableRow(line, i == m.groupCursor, width))
		sb.WriteByte('\n')
	}

	// Divider rule spanning the table body, then the TOTAL row: an at-a-glance
	// overall gauge with a completion percentage so the operator never has to
	// sum the groups by hand.
	ruleW := lipgloss.Width(pad) + nameW + len(gap) + countW + len(gap) + durW + len(gap) + barWidth
	ruleGlyph := "─"
	if m.ascii {
		ruleGlyph = "-"
	}
	sb.WriteString(lipgloss.NewStyle().Foreground(Muted).Render(strings.Repeat(ruleGlyph, ruleW)))
	sb.WriteByte('\n')

	overall := aggregateStatus(totPassed, totFailed, totRunning, totTotal)
	pct := 0
	if totTotal > 0 {
		pct = totPassed * 100 / totTotal
	}
	totLine := pad +
		lipgloss.NewStyle().Bold(true).Width(nameW).Render("TOTAL") + gap +
		lipgloss.NewStyle().Bold(true).Width(countW).Align(lipgloss.Right).
			Render(fmt.Sprintf("%d/%d", totPassed, totTotal)) + gap +
		lipgloss.NewStyle().Foreground(Muted).Width(durW).Align(lipgloss.Right).
			Render(fmt.Sprintf("%.1fs", totDur)) + gap +
		renderProgressBar(totPassed, totTotal, barWidth, overall, m.ascii) + gap +
		lipgloss.NewStyle().Bold(true).Render(fmt.Sprintf("%d%%", pct))
	if totFailed > 0 {
		totLine += gap + lipgloss.NewStyle().Foreground(Danger).
			Render(groupFailureText(totFailed, compact))
	}
	sb.WriteString(totLine)
	sb.WriteString("\n\n")
	sb.WriteString(renderActionStrip(width, []string{"enter tests", "r run/retry", "f failures first", "a all"}, m.ascii))

	return strings.TrimRight(sb.String(), "\n")
}

func groupFailureText(failed int, compact bool) string {
	if compact {
		return fmt.Sprintf("%d fail", failed)
	}
	return fmt.Sprintf("%d failed", failed)
}

func selectedTableRow(line string, selected bool, width int) string {
	if !selected {
		return line
	}
	if lipgloss.Width(line) >= width {
		return Styles.SelectedRow.Render(line)
	}
	return Styles.SelectedRow.Width(max(1, width)).Render(line)
}

func rowStatusColor(rs rowStatus) lipgloss.TerminalColor {
	switch rs {
	case rowPass:
		return Success
	case rowFail:
		return Danger
	case rowRunning:
		return Attention
	default:
		return Accent
	}
}

// renderProgressBar renders a fixed-width progress bar filled by passed/total.
// The filled portion uses the status color token.
func renderProgressBar(passed, total, width int, status rowStatus, ascii bool) string {
	if total == 0 || width <= 0 {
		return strings.Repeat("-", width)
	}
	filled := (passed * width) / total
	if filled > width {
		filled = width
	}
	empty := width - filled

	var fillColor lipgloss.TerminalColor
	switch status {
	case rowPass:
		fillColor = Success
	case rowFail:
		fillColor = Danger
	case rowRunning:
		fillColor = Attention
	default:
		fillColor = Muted
	}
	fill, emptyGlyph := "█", "░"
	if ascii {
		fill, emptyGlyph = "#", "."
	}
	bar := lipgloss.NewStyle().Foreground(fillColor).Render(strings.Repeat(fill, filled))
	bar += lipgloss.NewStyle().Foreground(Muted).Render(strings.Repeat(emptyGlyph, empty))
	return bar
}

// renderTestsDrilldown renders the test list for the selected group (§4.3).
func (m Model) renderTestsDrilldown(width int) string {
	width = max(1, width)
	if len(m.groups) == 0 || m.groupCursor >= len(m.groups) {
		return Styles.Body.Render("  no group selected")
	}
	if width < 12 {
		return renderPanel(width, "Tests", "", fmt.Sprintf("%d tests", len(m.groups[m.groupCursor].Tests)), panelAccent, m.ascii)
	}
	g := m.groups[m.groupCursor]
	sum := summarize(g, m.results)

	header := fmt.Sprintf("  %s   %d/%d %s  %d %s  %.1fs",
		lipgloss.NewStyle().Foreground(Accent).Render(g.Name),
		sum.Passed, sum.Total, m.renderRowGlyph(rowPass), sum.Failed, m.renderRowGlyph(rowFail), sum.Duration)
	header = truncateUI(header, max(1, width), m.ascii)
	breadcrumb := renderBreadcrumb(width, fmt.Sprintf("Groups / %s / tests", g.Name), m.ascii)
	nameW := min(40, max(8, width-16))

	var sb strings.Builder
	sb.WriteString(breadcrumb)
	sb.WriteByte('\n')
	sb.WriteString(header)
	sb.WriteByte('\n')

	for i, testName := range g.Tests {
		var rs rowStatus
		var elapsed float64
		k := resultKey{Name: testName, Sub: ""}
		if idx, ok := m.resultIndex[k]; ok {
			r := m.results[idx]
			rs = fromStatus(r.Status)
			elapsed = r.Elapsed
		} else {
			rs = rowNotRun
		}

		glyphStr := m.renderRowGlyph(rs)

		elapsedStr := "-"
		if rs != rowNotRun {
			elapsedStr = lipgloss.NewStyle().Foreground(Muted).Render(fmt.Sprintf("%.2fs", elapsed))
		}

		name := truncateUI(testName, nameW, m.ascii)
		gutter := "  "
		if i == m.testCursor {
			arrow := "❯"
			if m.ascii {
				arrow = ">"
			}
			gutter = lipgloss.NewStyle().Foreground(TerraformPurple).Bold(true).Render(arrow) + " "
		}
		line := gutter + glyphStr + "  " + lipgloss.NewStyle().Width(nameW).Render(name) + "  " + elapsedStr
		sb.WriteString(selectedTableRow(line, i == m.testCursor, width))
		sb.WriteByte('\n')
	}
	sb.WriteString("\n")
	sb.WriteString(renderActionStrip(width, []string{"enter log", "r run/retry", "esc back"}, m.ascii))
	return strings.TrimRight(sb.String(), "\n")
}

// renderLogViewport renders the scrollable log pane for the selected test (§4.3).
//
// REDACTION CONTRACT: Output is rendered VERBATIM. The producer (CLI, PR10c) is
// responsible for redacting secrets before sending TestResult.Output to the TUI.
// This function never imports or invokes redact; it only echoes what it receives.
func (m Model) renderLogViewport(width int) string {
	width = max(1, width)
	if width < 12 {
		return renderPanel(width, "Log", "REDACTED", "j/k scroll\nesc back", panelAttention, m.ascii)
	}
	testName := m.selectedTestName()
	groupName := ""
	if len(m.groups) > 0 && m.groupCursor >= 0 && m.groupCursor < len(m.groups) {
		groupName = m.groups[m.groupCursor].Name
	}
	badge := renderSemanticBadge("REDACTED", panelAttention)
	scroll := fmt.Sprintf("scroll %3.0f%%", m.logVP.ScrollPercent()*100)
	meta := badge + "  " + lipgloss.NewStyle().Foreground(Muted).Render(scroll)
	breadcrumb := renderBreadcrumb(max(1, width-lipgloss.Width(meta)-4), fmt.Sprintf("Groups / %s / %s / log", groupName, testName), m.ascii)
	footer := "j/k scroll" + uiSeparator(m.ascii) + "esc back"

	content := m.logVP.View()
	border := lipgloss.RoundedBorder()
	if m.ascii {
		border = asciiPanelBorder
	}
	pane := Styles.PaneBorder.
		Border(border).
		Width(max(1, width-4)).
		Render(content)

	return lipgloss.JoinVertical(lipgloss.Left,
		"  "+breadcrumb+"  "+meta,
		pane,
		"  "+lipgloss.NewStyle().Foreground(Muted).Render(footer),
	)
}

func renderBreadcrumb(width int, text string, ascii bool) string {
	return Styles.Breadcrumb.Render(truncateUI(text, max(1, width-2), ascii))
}

func renderActionStrip(width int, actions []string, ascii bool) string {
	if len(actions) == 0 {
		return ""
	}
	return "  " + Styles.Action.Render(truncateUI(strings.Join(actions, "  "+uiSeparator(ascii)+"  "), max(1, width-2), ascii))
}

func renderSemanticBadge(text string, tone panelTone) string {
	return lipgloss.NewStyle().Bold(true).Foreground(toneColor(tone)).Render(text)
}

// renderRowGlyph returns the colored leading glyph for a status row. Running
// rows animate with the current Pulsar frame; every other status keeps its
// static glyph. The color token comes from glyphFor (Attention for running), so
// the spinner stays on-theme with the running progress bar beside it.
func (m Model) renderRowGlyph(rs rowStatus) string {
	ge := glyphFor(rs, m.ascii)
	glyph := ge.Glyph
	if rs == rowRunning {
		glyph = spinnerGlyph(m.spinnerFrame, m.ascii)
	}
	if c, ok := ge.Color.(lipgloss.TerminalColor); ok {
		return lipgloss.NewStyle().Foreground(c).Render(glyph)
	}
	return glyph
}

func (m Model) contextualKeys() contextualKeyMap {
	if keys, ok := m.overlayContextualKeys(); ok {
		return keys
	}
	if m.help.ShowAll {
		return contextualKeyMap{
			section: []string{"? or esc close"},
			global:  []string{"q quit"},
		}
	}
	if keys, ok := m.groupsContextualKeys(); ok {
		return keys
	}
	if m.section == sectionTriage && m.triageDetailActive {
		return m.triageDetailContextualKeys()
	}
	keys := contextualKeyMap{
		navigation: []string{"←/→ tabs", "↑/↓ move"},
		global:     []string{"? help", "q quit"},
	}
	switch m.section {
	case sectionPreflight:
		keys.section = []string{"p preflight", "m mode", "v vars"}
	case sectionGroups:
		keys.section = []string{"enter detail", "r run/retry", "f failures"}
	case sectionRun:
		keys.section = []string{"e export"}
		if len(m.groups) > 0 {
			keys.section = []string{"r run", "R retry failed", "e export"}
		}
	case sectionTriage:
		keys.section = []string{"t refresh", "K sync"}
		if len(m.triageFailures) > 0 {
			keys.section = []string{"enter detail", "t refresh", "K sync"}
		}
	}
	return keys
}

func (m Model) overlayContextualKeys() (contextualKeyMap, bool) {
	switch {
	case m.pickerActive:
		return contextualKeyMap{section: []string{"↑/↓ move", "enter select", "esc/q close"}}, true
	case m.editorActive:
		if m.editorFocused {
			return contextualKeyMap{section: []string{"enter save", "esc/q close"}}, true
		}
		return contextualKeyMap{section: []string{"↑/↓ move", "enter edit", "esc/q close", "p preflight"}}, true
	case m.sweepConfirmActive:
		return contextualKeyMap{section: []string{"enter confirm", "esc close"}}, true
	case m.orphanOverlayActive:
		hints := []string{"esc close"}
		if len(m.orphans) > 0 {
			hints = append([]string{"s sweep"}, hints...)
		}
		return contextualKeyMap{section: hints}, true
	case m.fileIssueActive:
		return contextualKeyMap{section: []string{"enter confirm", "esc close"}}, true
	case m.resultOverlayActive:
		return contextualKeyMap{section: []string{"esc close"}}, true
	case m.resumePrompt:
		return contextualKeyMap{section: []string{"enter resume", "esc dismiss"}}, true
	default:
		return contextualKeyMap{}, false
	}
}

func (m Model) groupsContextualKeys() (contextualKeyMap, bool) {
	if m.section != sectionGroups {
		return contextualKeyMap{}, false
	}
	switch m.focus {
	case focusTests:
		return contextualKeyMap{
			navigation: []string{"←/→ tabs", "↑/↓ move"},
			section:    []string{"enter log", "r run/retry", "esc back"},
			global:     []string{"? help", "q quit"},
		}, true
	case focusLog:
		return contextualKeyMap{
			navigation: []string{"←/→ tabs"},
			section:    []string{"j/k scroll", "esc back"},
			global:     []string{"? help", "q quit"},
		}, true
	default:
		return contextualKeyMap{}, false
	}
}

func (m Model) triageDetailContextualKeys() contextualKeyMap {
	keys := contextualKeyMap{
		navigation: []string{"←/→ tabs", "↑/↓ move"},
		section:    []string{"t refresh", "K sync", "esc back"},
		global:     []string{"? help", "q quit"},
	}
	if f, ok := m.selectedTriageFailure(); ok && engine.EligibleForIssueFiling(f) {
		keys.section = append([]string{"i file issue"}, keys.section...)
	}
	return keys
}

func renderKeyHints(width int, hints []string, ascii bool) string {
	hints = append([]string(nil), hints...)
	if ascii {
		for i, hint := range hints {
			hints[i] = asciiKeyHint(hint)
		}
	}
	separator := "  " + uiSeparator(ascii) + "  "
	line := strings.Join(hints, separator)
	for len(hints) > 1 && lipgloss.Width(line) > width {
		drop := 1
		for i := 1; i < len(hints); i++ {
			switch hints[i] {
			case "↑/↓ move", "up/down move", "m mode", "v vars", "f failures", "R retry failed", "e export", "t refresh", "K sync":
				drop = i
				goto remove
			}
		}
	remove:
		hints = append(hints[:drop], hints[drop+1:]...)
		line = strings.Join(hints, separator)
	}
	return Styles.HelpFooter.Render(truncateUI(line, max(1, width), ascii))
}

// renderPickerOverlay renders the mode-selection overlay shown when the user
// presses 'm'. Each mode is listed with its description; the active row is
// marked with a cursor arrow. Secret values are not involved here.
func (m Model) renderPickerOverlay(width int) string {
	bodyW := max(1, width-4)
	var sb strings.Builder

	for i, pm := range m.pickerModes {
		gutter := "  "
		if i == m.pickerCursor {
			arrow := ">"
			if !m.ascii {
				arrow = "❯"
			}
			gutter = lipgloss.NewStyle().Foreground(TerraformPurple).Bold(true).Render(arrow) + " "
		}
		name := lipgloss.NewStyle().Foreground(Accent).Render(pm.Name)
		desc := truncateUI(pm.Description, max(1, bodyW-20), m.ascii)
		sb.WriteString(fmt.Sprintf("%s%-14s  %s\n", gutter, name, desc))
	}

	sb.WriteString("\n  ")
	sb.WriteString(lipgloss.NewStyle().Foreground(Muted).Render("enter select  esc/q close"))
	return renderPanel(width, "Select mode", "", strings.TrimRight(sb.String(), "\n"), panelAccent, m.ascii)
}

// renderEditorOverlay renders the variable-config overlay shown when the user
// presses 'v'. Non-secret fields show their current value (from getenv) in an
// editable text input. Secret fields render <set>/<unset> and are read-only;
// the raw secret value is NEVER shown.
func (m Model) renderEditorOverlay(width int) string {
	var sb strings.Builder

	if len(m.editorFields) == 0 {
		sb.WriteString("  No variables for this mode.\n")
	}

	for i, f := range m.editorFields {
		gutter := "  "
		if i == m.editorCursor {
			arrow := ">"
			if !m.ascii {
				arrow = "❯"
			}
			gutter = lipgloss.NewStyle().Foreground(TerraformPurple).Bold(true).Render(arrow) + " "
		}

		reqTag := ""
		if f.Required {
			reqTag = " required"
		}

		var valStr string
		if f.Secret {
			// Redaction contract: secret values are NEVER displayed.
			if m.getenv != nil && m.getenv(f.Key) != "" {
				valStr = lipgloss.NewStyle().Foreground(Muted).Render("<set>")
			} else {
				valStr = lipgloss.NewStyle().Foreground(Muted).Render("<unset>")
			}
			valStr += lipgloss.NewStyle().Foreground(Muted).Render("  secret")
		} else if m.editorFocused && i == m.editorCursor {
			valStr = f.input.View()
		} else {
			v := f.input.Value()
			if v == "" {
				v = lipgloss.NewStyle().Foreground(Muted).Render("<unset>")
			}
			valStr = v
		}

		keyLabel := lipgloss.NewStyle().Width(34).Render(f.Key)
		line := fmt.Sprintf("%s%s  %s%s", gutter, keyLabel, valStr, reqTag)
		sb.WriteString(line + "\n")
	}

	sb.WriteString("\n  ")
	sb.WriteString(lipgloss.NewStyle().Foreground(Muted).Render("enter edit  esc/q close  p preflight"))
	return renderPanel(width, "Config: "+m.mode, "", strings.TrimRight(sb.String(), "\n"), panelAccent, m.ascii)
}

// Compile-time interface check.
var _ provider.CheckStatus = provider.CheckOK
