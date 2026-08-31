package tui

import "github.com/charmbracelet/lipgloss"

// Semantic color tokens.
// Each token is an AdaptiveColor that works on both light and dark terminals.
// Light values use GitHub Primer light-mode hexes; dark values use Primer dark-mode hexes.
var (
	// Success - pass, ok, reachable.
	Success = lipgloss.AdaptiveColor{Light: "#1a7f37", Dark: "#3fb950"}
	// Danger - fail, panic, timeout, blocked.
	Danger = lipgloss.AdaptiveColor{Light: "#cf222e", Dark: "#f85149"}
	// Attention - pending, running, unverified.
	Attention = lipgloss.AdaptiveColor{Light: "#9a6700", Dark: "#d29922"}
	// Flaky - passed on retry.
	Flaky = lipgloss.AdaptiveColor{Light: "#8250df", Dark: "#a371f7"}
	// Muted - metadata, timings, not-run.
	Muted = lipgloss.AdaptiveColor{Light: "#656d76", Dark: "#8b949e"}
	// Accent - headers, links, identifiers.
	Accent = lipgloss.AdaptiveColor{Light: "#0969da", Dark: "#2f81f7"}
	// TerraformPurple identifies the Terraform half of GH/TF and marks the
	// current navigation position without becoming a large filled surface.
	TerraformPurple = lipgloss.AdaptiveColor{Light: "#844fba", Dark: "#a16be8"}
	// Selected is kept as an internal alias for the previous token name.
	Selected = TerraformPurple
	// SelectedSurface is the neutral background used for the current table row.
	SelectedSurface = lipgloss.AdaptiveColor{Light: "#f6f8fa", Dark: "#21262d"}
	// Border - rules, pane borders.
	Border = lipgloss.AdaptiveColor{Light: "#d0d7de", Dark: "#30363d"}
	// Text is intentionally unset so prose inherits the terminal default foreground.
)
