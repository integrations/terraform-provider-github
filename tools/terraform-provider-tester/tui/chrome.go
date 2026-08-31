package tui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"

	"github.com/github/terraform-provider-tester/provider"
)

type tabSummary struct {
	Label string
	Badge string
	Tone  panelTone
}

func (m Model) tabSummaries() [numSections]tabSummary {
	var ok, warn, fail int
	for _, check := range m.checks {
		switch check.Status {
		case provider.CheckOK:
			ok++
		case provider.CheckWarn:
			warn++
		default:
			fail++
		}
	}
	preflight := tabSummary{Label: "Preflight", Badge: "-", Tone: panelNeutral}
	switch {
	case fail > 0:
		preflight.Badge, preflight.Tone = fmt.Sprintf("%d blocked", fail), panelDanger
	case warn > 0:
		preflight.Badge, preflight.Tone = fmt.Sprintf("%d warn", warn), panelAttention
	case ok > 0:
		preflight.Badge, preflight.Tone = "ready", panelSuccess
	}
	total := 0
	for _, group := range m.groups {
		total += len(group.Tests)
	}
	runTone := panelNeutral
	if m.running {
		runTone = panelAttention
	}
	return [numSections]tabSummary{
		preflight,
		{Label: "Groups", Badge: fmt.Sprintf("%d", len(m.groups)), Tone: panelAccent},
		{Label: "Run", Badge: fmt.Sprintf("%d", total), Tone: runTone},
		{Label: "Triage", Badge: fmt.Sprintf("%d", len(m.triageFailures)), Tone: panelFlaky},
	}
}

func (m Model) renderHeader(width int) string {
	width = max(1, width)
	mode := layoutForWidth(width)
	mark := "◆"
	if m.ascii {
		mark = "*"
	}
	title := lipgloss.NewStyle().Bold(true).Foreground(Accent).Render(mark) +
		" " + renderConsoleTitle(mode != layoutCompact, m.version)

	left := []string{title}
	for _, chip := range m.headerChips(mode) {
		left = append(left, renderChromeChip(chip, panelNeutral, m.ascii))
	}
	right := m.renderRateGauge(mode != layoutCompact)
	return joinChromeLine(width, left, right)
}

func (m Model) headerChips(mode layoutMode) []string {
	if mode == layoutCompact {
		return []string{m.provider, m.mode}
	}
	chips := []string{"provider: " + m.provider, "mode: " + m.mode}
	if mode == layoutWide && m.owner != "" {
		chips = append(chips, "owner: "+m.owner)
	}
	return chips
}

func renderChromeChip(label string, tone panelTone, ascii bool) string {
	left, right := "‹", "›"
	if ascii {
		left, right = "[", "]"
	}
	return lipgloss.NewStyle().
		Foreground(Muted).
		Render(left + label + right)
}

func joinChromeLine(width int, left []string, right string) string {
	if right != "" && chromeLineWidth(left, right) > width {
		right = ""
	}
	for len(left) > 1 && chromeLineWidth(left, right) > width {
		left = left[:len(left)-1]
	}
	line := strings.Join(left, " ")
	if right == "" {
		if lipgloss.Width(line) > width {
			return lipgloss.NewStyle().MaxWidth(width).Render(line)
		}
		return line
	}
	if leftW := lipgloss.Width(line); leftW > 0 {
		space := width - leftW - lipgloss.Width(right)
		if space < 1 {
			if leftW >= width {
				return lipgloss.NewStyle().MaxWidth(width).Render(line)
			}
			space = 1
		}
		return lipgloss.NewStyle().MaxWidth(width).Render(line + strings.Repeat(" ", space) + right)
	}
	return lipgloss.NewStyle().MaxWidth(width).Render(right)
}

func chromeLineWidth(left []string, right string) int {
	line := strings.Join(left, " ")
	if right == "" {
		return lipgloss.Width(line)
	}
	if line == "" {
		return lipgloss.Width(right)
	}
	return lipgloss.Width(line) + 1 + lipgloss.Width(right)
}

func (m Model) renderRateGauge(showReset bool) string {
	if m.rate.Limit == 0 {
		return ""
	}
	resetStr := ""
	if showReset && !m.rate.Reset.IsZero() {
		reset := "↺"
		if m.ascii {
			reset = "reset"
		}
		resetStr = " " + reset + " " + m.rate.Reset.UTC().Format("15:04 UTC")
	}
	gauge := fmt.Sprintf("rate: %d/%d%s", m.rate.Remaining, m.rate.Limit, resetStr)

	headroom := float64(m.rate.Remaining) / float64(m.rate.Limit)
	switch {
	case headroom < 0.05:
		return lipgloss.NewStyle().Foreground(Danger).Render(gauge)
	case headroom < 0.15:
		return lipgloss.NewStyle().Foreground(Attention).Render(gauge)
	default:
		return gauge
	}
}

func (m Model) renderTabs(width int) string {
	width = max(1, width)
	mode := layoutForWidth(width)
	summaries := m.tabSummaries()
	labels := tabLabels(summaries, mode == layoutCompact)
	candidates := []tabRenderPlan{
		{labels: labels, badges: true, padded: true},
		{labels: labels, badges: false, padded: true},
		{labels: labels, badges: false, numbersOnlyInactive: true, separator: " "},
		{labels: labels, activeOnly: true},
	}
	if mode != layoutCompact {
		compactLabels := tabLabels(summaries, true)
		candidates = append(candidates,
			tabRenderPlan{labels: compactLabels, badges: true, padded: true},
			tabRenderPlan{labels: compactLabels, badges: false, padded: true},
			tabRenderPlan{labels: compactLabels, badges: false, numbersOnlyInactive: true, separator: " "},
			tabRenderPlan{labels: compactLabels, activeOnly: true},
		)
	}
	for _, candidate := range candidates {
		bar, start, activeW := m.renderTabBar(candidate, summaries, width)
		if lipgloss.Width(bar) <= width {
			return bar + "\n" + m.renderTabRule(width, start, activeW)
		}
	}
	bar, start, activeW := m.renderTabBar(tabRenderPlan{labels: labels, activeOnly: true}, summaries, width)
	return bar + "\n" + m.renderTabRule(width, start, activeW)
}

func (m Model) renderTabSegment(i int, text string) string {
	style := Styles.TabInactive.Padding(0, 0)
	if section(i) == m.section {
		style = Styles.TabActive.Padding(0, 0)
	}
	padded := " " + text
	if i < numSections-1 {
		padded += " "
	}
	return style.Render(padded)
}

type tabRenderPlan struct {
	labels              [numSections]string
	badges              bool
	padded              bool
	numbersOnlyInactive bool
	activeOnly          bool
	separator           string
}

func tabLabels(summaries [numSections]tabSummary, compact bool) [numSections]string {
	if compact {
		return [numSections]string{"Pre", "Grp", "Run", "Tri"}
	}
	var labels [numSections]string
	for i, summary := range summaries {
		labels[i] = summary.Label
	}
	return labels
}

func (m Model) renderTabBar(plan tabRenderPlan, summaries [numSections]tabSummary, width int) (string, int, int) {
	active := int(m.section)
	if active < 0 || active >= numSections {
		active = 0
	}
	if plan.activeOnly {
		text := activeTabText(active+1, plan.labels[active], width, m.ascii)
		rendered := m.renderTabText(active, text)
		return rendered, 0, lipgloss.Width(rendered)
	}

	parts := make([]string, 0, numSections)
	start := 0
	for i, summary := range summaries {
		text := fmt.Sprintf("%d", i+1)
		if !plan.numbersOnlyInactive || i == active {
			text += " " + plan.labels[i]
		}
		if plan.badges && summary.Badge != "" {
			text += " " + summary.Badge
		}
		if plan.padded {
			part := m.renderTabSegment(i, text)
			if i < active {
				start += lipgloss.Width(part)
			}
			parts = append(parts, part)
			continue
		}
		part := m.renderTabText(i, text)
		if i < active {
			start += lipgloss.Width(part)
			if plan.separator != "" {
				start += lipgloss.Width(plan.separator)
			}
		}
		parts = append(parts, part)
	}
	bar := strings.Join(parts, plan.separator)
	activeW := lipgloss.Width(parts[active])
	return bar, start, activeW
}

func (m Model) renderTabText(i int, text string) string {
	style := Styles.TabInactive.Padding(0, 0)
	if section(i) == m.section {
		style = Styles.TabActive.Padding(0, 0)
	}
	return style.Render(text)
}

func activeTabText(number int, label string, width int, ascii bool) string {
	prefix := fmt.Sprintf("%d", number)
	if width <= lipgloss.Width(prefix) {
		return truncateUI(prefix, max(1, width), ascii)
	}
	labelWidth := max(0, width-lipgloss.Width(prefix)-1)
	if labelWidth == 0 {
		return prefix
	}
	return prefix + " " + truncateUI(label, labelWidth, ascii)
}

func (m Model) renderTabRule(width, start, activeW int) string {
	start = max(0, start)
	if start > width {
		start = width
	}
	count := max(0, min(activeW, width-start))
	ruleGlyph := "─"
	if m.ascii {
		ruleGlyph = "-"
	}
	return strings.Repeat(" ", start) +
		lipgloss.NewStyle().Foreground(TerraformPurple).Render(strings.Repeat(ruleGlyph, count))
}

func (m Model) renderStatus(width int) string {
	width = max(1, width)
	s := m.status
	if m.statusError && m.err != nil {
		s = m.err.Error()
	}
	if m.running {
		if m.runLabel != "" {
			s = strings.TrimSuffix(s, uiEllipsis(m.ascii)) + " " + m.runLabel
		}
		if m.runElapsed >= time.Second {
			s += uiSeparator(m.ascii) + formatRunDuration(m.runElapsed)
		}
		if m.stalledFor >= stallThreshold {
			s += uiSeparator(m.ascii) + "quiet " + formatRunDuration(m.stalledFor)
		}
		s = spinnerGlyph(m.spinnerFrame, m.ascii) + " " + s
	}
	s = truncateUI(s, width, m.ascii)
	style := Styles.StatusLine
	if m.statusError {
		style = lipgloss.NewStyle().Foreground(Danger)
	}
	if m.operation == "" {
		return style.Render(s)
	}
	operation := truncateUI(m.operation, width, m.ascii)
	if lipgloss.Width(s)+1+lipgloss.Width(operation) > width {
		return style.Render(s)
	}
	return style.Render(s) + strings.Repeat(" ", width-lipgloss.Width(s)-lipgloss.Width(operation)) +
		Styles.StatusLine.Render(operation)
}

func (m Model) renderFooter(width int) string {
	return renderKeyHints(width, m.contextualKeys().footerHints(), m.ascii)
}

func (m Model) renderHelpPanel(width int) string {
	groups := []struct {
		title string
		lines []string
	}{
		{"Navigation", []string{"←/→ or tab/shift+tab  switch tabs", "1-4  jump to tab", "↑/↓ or j/k  move", "enter/l  open", "esc/h  back"}},
		{"Run and retry", []string{"r  run or retry selection", "R  retry all failures", "f  failures first", "a  show all"}},
		{"Triage and known issues", []string{"t  refresh local triage", "K  sync known issues", "i  preview selected issue"}},
		{"Reports and cleanup", []string{"e  export report", "o  list orphans", "s  sweep from reviewed snapshot"}},
		{"Configuration", []string{"p  preflight", "m  mode", "v  variables", "c  copy command", "y  copy log"}},
	}
	var blocks []string
	for _, group := range groups {
		lines := group.lines
		if m.ascii {
			lines = make([]string, len(group.lines))
			for i, line := range group.lines {
				lines[i] = asciiKeyHint(line)
			}
		}
		blocks = append(blocks, renderHelpGroup(group.title, lines))
	}
	return renderPanel(
		width,
		"Keyboard reference",
		"All actions are scoped to the current view.",
		strings.Join(blocks, "\n\n"),
		panelAccent,
		m.ascii,
	) + "\n" + Styles.HelpFooter.Render(truncateUI("  ? or esc close", max(1, width), m.ascii))
}
