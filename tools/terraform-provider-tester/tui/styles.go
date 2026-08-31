package tui

import "github.com/charmbracelet/lipgloss"

// Styles holds the lipgloss styles for the TUI chrome.
// All styles are built from the semantic color tokens in theme.go.
// Selected rows use a neutral surface; Terraform purple is reserved for the
// active-tab rule and cursor glyphs.
var Styles = struct {
	Header       lipgloss.Style
	TabActive    lipgloss.Style
	TabInactive  lipgloss.Style
	Body         lipgloss.Style
	StatusLine   lipgloss.Style
	HelpFooter   lipgloss.Style
	SelectedRow  lipgloss.Style
	PaneBorder   lipgloss.Style
	SectionTitle lipgloss.Style
	Chip         lipgloss.Style
	MetricLabel  lipgloss.Style
	Breadcrumb   lipgloss.Style
	Action       lipgloss.Style
	DangerPanel  lipgloss.Style
}{
	// Header bar - bold, accent foreground for the title portion.
	Header: lipgloss.NewStyle().Bold(true),

	// Active tab - bold text; renderTabRule carries the Terraform-purple line.
	TabActive: lipgloss.NewStyle().
		Foreground(lipgloss.NoColor{}).
		Bold(true).
		Padding(0, 1),

	// Inactive tab - Muted foreground.
	TabInactive: lipgloss.NewStyle().
		Foreground(Muted).
		Padding(0, 1),

	// Section body - default terminal foreground; no background.
	Body: lipgloss.NewStyle(),

	// Status line - muted, single line at the bottom.
	StatusLine: lipgloss.NewStyle().Foreground(Muted),

	// Help footer - muted, rendered by the Bubbles help component.
	HelpFooter: lipgloss.NewStyle().Foreground(Muted),

	// Selected row cursor highlight - neutral background, never brand purple.
	SelectedRow: lipgloss.NewStyle().Background(SelectedSurface),

	// Pane border - Border color.
	PaneBorder: lipgloss.NewStyle().BorderForeground(Border),

	SectionTitle: lipgloss.NewStyle().Bold(true),
	Chip:         lipgloss.NewStyle().Padding(0, 1).Foreground(Muted).Border(lipgloss.RoundedBorder()).BorderForeground(Border),
	MetricLabel:  lipgloss.NewStyle().Foreground(Muted),
	Breadcrumb:   lipgloss.NewStyle().Foreground(Muted),
	Action:       lipgloss.NewStyle().Foreground(Accent).Bold(true),
	DangerPanel:  lipgloss.NewStyle().BorderForeground(Danger),
}
