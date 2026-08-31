package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type layoutMode int

const (
	layoutCompact layoutMode = iota
	layoutStandard
	layoutWide
)

func layoutForWidth(width int) layoutMode {
	switch {
	case width >= 110:
		return layoutWide
	case width >= 72:
		return layoutStandard
	default:
		return layoutCompact
	}
}

type panelTone int

const (
	panelNeutral panelTone = iota
	panelAccent
	panelSuccess
	panelAttention
	panelDanger
	panelFlaky
)

type metricSpec struct {
	Label string
	Value string
	Tone  panelTone
}

func toneColor(tone panelTone) lipgloss.TerminalColor {
	switch tone {
	case panelAccent:
		return Accent
	case panelSuccess:
		return Success
	case panelAttention:
		return Attention
	case panelDanger:
		return Danger
	case panelFlaky:
		return Flaky
	default:
		return Border
	}
}

var asciiPanelBorder = lipgloss.Border{
	Top:          "-",
	Bottom:       "-",
	Left:         "|",
	Right:        "|",
	TopLeft:      "+",
	TopRight:     "+",
	BottomLeft:   "+",
	BottomRight:  "+",
	MiddleLeft:   "+",
	MiddleRight:  "+",
	Middle:       "-",
	MiddleTop:    "+",
	MiddleBottom: "+",
}

func renderPanel(width int, title, subtitle, body string, tone panelTone, ascii bool) string {
	width = max(1, width)
	if width < 6 {
		parts := []string{title}
		if subtitle != "" {
			parts = append(parts, subtitle)
		}
		if body != "" {
			parts = append(parts, body)
		}
		return truncateBlockUI(strings.Join(parts, "\n"), width, ascii)
	}
	border := lipgloss.RoundedBorder()
	if ascii {
		border = asciiPanelBorder
	}
	inner := max(1, width-4)
	parts := []string{
		lipgloss.NewStyle().Bold(true).Foreground(toneColor(tone)).Render(truncateUI(title, inner, ascii)),
	}
	if subtitle != "" {
		parts = append(parts, lipgloss.NewStyle().Foreground(Muted).Render(truncateUI(subtitle, inner, ascii)))
	}
	if body != "" {
		parts = append(parts, lipgloss.NewStyle().MaxWidth(inner).Render(truncateBlockUI(body, inner, ascii)))
	}
	return lipgloss.NewStyle().
		Width(inner).
		Border(border).
		BorderForeground(toneColor(tone)).
		Padding(0, 1).
		Render(lipgloss.JoinVertical(lipgloss.Left, parts...))
}

func renderMetricRow(width int, metrics []metricSpec, ascii bool) string {
	if width <= 0 || len(metrics) == 0 {
		return ""
	}
	gap := 2
	cells := make([]string, 0, len(metrics))
	totalWidth := 0
	for i, metric := range metrics {
		cell := renderMetricCell(metric, width, ascii)
		cells = append(cells, cell)
		totalWidth += lipgloss.Width(cell)
		if i > 0 {
			totalWidth += gap
		}
	}
	if totalWidth <= width {
		return strings.Join(cells, strings.Repeat(" ", gap))
	}
	return strings.Join(cells, "\n")
}

func renderMetricCell(metric metricSpec, width int, ascii bool) string {
	if width <= 0 {
		return ""
	}
	plain := strings.TrimSpace(metric.Label + " " + metric.Value)
	if lipgloss.Width(plain) > width {
		value := truncateUI(metric.Value, width, ascii)
		valueW := lipgloss.Width(value)
		labelW := max(0, width-valueW-1)
		if labelW == 0 {
			plain = truncateUI(plain, width, ascii)
		} else {
			plain = truncateUI(metric.Label, labelW, ascii) + " " + value
		}
	}
	if metric.Label == "" {
		return lipgloss.NewStyle().Bold(true).Foreground(toneColor(metric.Tone)).Render(truncateUI(metric.Value, width, ascii))
	}
	label, value, ok := strings.Cut(plain, " ")
	if !ok {
		return lipgloss.NewStyle().MaxWidth(width).Render(plain)
	}
	return lipgloss.NewStyle().Foreground(Muted).Render(label) + " " +
		lipgloss.NewStyle().Bold(true).Foreground(toneColor(metric.Tone)).Render(value)
}

func renderHelpGroup(title string, lines []string) string {
	return Styles.SectionTitle.Render(title) + "\n  " + strings.Join(lines, "\n  ")
}
