package tui

import (
	"strings"

	"github.com/charmbracelet/x/ansi"
)

func uiSeparator(ascii bool) string {
	if ascii {
		return " | "
	}
	return " · "
}

func uiEllipsis(ascii bool) string {
	if ascii {
		return "..."
	}
	return "…"
}

func truncateUI(s string, width int, ascii bool) string {
	if width <= 0 {
		return ""
	}
	if ansi.StringWidth(s) <= width {
		return s
	}
	suffix := uiEllipsis(ascii)
	if ansi.StringWidth(suffix) > width {
		return ansi.Truncate(suffix, width, "")
	}
	return ansi.Truncate(s, width, suffix)
}

func truncateBlockUI(block string, width int, ascii bool) string {
	lines := strings.Split(block, "\n")
	for i, line := range lines {
		lines[i] = truncateUI(line, width, ascii)
	}
	return strings.Join(lines, "\n")
}

func asciiKeyHint(hint string) string {
	return strings.NewReplacer(
		"←/→", "left/right",
		"↑/↓", "up/down",
		"↺", "reset",
		"↻", "resume",
	).Replace(hint)
}
