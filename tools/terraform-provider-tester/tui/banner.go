package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// renderBanner builds the welcome card shown at the top of the Preflight tab.
// It mirrors the gh / Copilot CLI welcome box: a small face mark on the left,
// the product name + version, a one-line tagline, and a tip line, wrapped in a
// rounded border. The function is pure so it can be golden-tested; it never
// reads the clock or environment.
func renderBanner(version string, ascii bool) string {
	icon := ghTFMark(ascii)
	titleLine := renderConsoleTitle(true, version)
	tagline := Tagline
	tip := Styles.StatusLine.Render("Tip: press ? for keys · docs " + DocsPath)

	// The footer signals that Terraform Provider Tester is GitHub property and
	// states its license. '©' degrades to '(c)' on the ASCII profile.
	mark := "©"
	if ascii {
		mark = "(c)"
	}
	footer := lipgloss.NewStyle().Foreground(Muted).
		Render(mark + " " + Vendor + " · " + License + " License")

	text := lipgloss.JoinVertical(lipgloss.Left, titleLine, tagline, "", tip, footer)
	body := lipgloss.JoinHorizontal(lipgloss.Top, icon, "   ", text)

	border := lipgloss.RoundedBorder()
	if ascii {
		border = lipgloss.NormalBorder()
	}
	card := lipgloss.NewStyle().
		Border(border).
		BorderForeground(Border).
		Padding(0, 2).
		Render(body)

	return card
}

// ghTFMark returns a compact stacked GitHub/Terraform mark. It uses only
// terminal-safe box drawing and has a strict ASCII fallback.
func ghTFMark(ascii bool) string {
	top, side, divider, bottom := "╭─────╮", "│", "│  ╱  │", "╰─────╯"
	if ascii {
		top, side, divider, bottom = "+-----+", "|", "|  /  |", "+-----+"
	}
	gh := lipgloss.NewStyle().Bold(true).Foreground(Accent).Render("GH")
	tf := lipgloss.NewStyle().Bold(true).Foreground(TerraformPurple).Render("TF")
	return strings.Join([]string{
		top,
		side + " " + gh + "  " + side,
		divider,
		side + " " + tf + "  " + side,
		bottom,
	}, "\n")
}
