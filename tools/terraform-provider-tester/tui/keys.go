package tui

import (
	"github.com/charmbracelet/bubbles/key"
)

// keyMap holds all key bindings for the TUI.
// Bindings for later slices are defined here so the help footer is complete;
// their Update handlers will be wired in PR10b/c.
type keyMap struct {
	Up           key.Binding
	Down         key.Binding
	Top          key.Binding
	Bottom       key.Binding
	NextSection  key.Binding
	PrevSection  key.Binding
	PreflightTab key.Binding
	GroupsTab    key.Binding
	RunTab       key.Binding
	TriageTab    key.Binding
	DrillIn      key.Binding
	Back         key.Binding
	Help         key.Binding
	Quit         key.Binding

	// PR10b/c bindings - defined now so help footer is complete.
	Retry         key.Binding
	RetryAll      key.Binding
	FailuresFirst key.Binding
	ShowAll       key.Binding
	Orphans       key.Binding
	Export        key.Binding
	Sweep         key.Binding
	Preflight     key.Binding
	CopyCmd       key.Binding
	CopyLog       key.Binding
	Mode          key.Binding
	Vars          key.Binding

	TriageRefresh   key.Binding
	SyncKnownIssues key.Binding
	FileIssue       key.Binding
}

// defaultKeys returns the default key map.
func defaultKeys() keyMap {
	return keyMap{
		Up: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "move up"),
		),
		Down: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "move down"),
		),
		Top: key.NewBinding(
			key.WithKeys("g"),
			key.WithHelp("g", "top"),
		),
		Bottom: key.NewBinding(
			key.WithKeys("G"),
			key.WithHelp("G", "bottom"),
		),
		NextSection: key.NewBinding(
			key.WithKeys("tab", "right"),
			key.WithHelp("→/tab", "next tab"),
		),
		PrevSection: key.NewBinding(
			key.WithKeys("shift+tab", "left"),
			key.WithHelp("←/shift+tab", "previous tab"),
		),
		PreflightTab: key.NewBinding(
			key.WithKeys("1"),
			key.WithHelp("1", "preflight"),
		),
		GroupsTab: key.NewBinding(
			key.WithKeys("2"),
			key.WithHelp("2", "groups"),
		),
		RunTab: key.NewBinding(
			key.WithKeys("3"),
			key.WithHelp("3", "run"),
		),
		TriageTab: key.NewBinding(
			key.WithKeys("4"),
			key.WithHelp("4", "triage"),
		),
		DrillIn: key.NewBinding(
			key.WithKeys("enter", "l"),
			key.WithHelp("enter/l", "drill in"),
		),
		Back: key.NewBinding(
			key.WithKeys("esc", "h"),
			key.WithHelp("esc/h", "back"),
		),
		Help: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "help"),
		),
		Quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("q", "quit"),
		),
		Retry: key.NewBinding(
			key.WithKeys("r"),
			key.WithHelp("r", "retry"),
		),
		RetryAll: key.NewBinding(
			key.WithKeys("R"),
			key.WithHelp("R", "retry all failures"),
		),
		FailuresFirst: key.NewBinding(
			key.WithKeys("f"),
			key.WithHelp("f", "failures first"),
		),
		ShowAll: key.NewBinding(
			key.WithKeys("a"),
			key.WithHelp("a", "show all"),
		),
		Orphans: key.NewBinding(
			key.WithKeys("o"),
			key.WithHelp("o", "orphans"),
		),
		Export: key.NewBinding(
			key.WithKeys("e"),
			key.WithHelp("e", "export"),
		),
		Sweep: key.NewBinding(
			key.WithKeys("s"),
			key.WithHelp("s", "sweep"),
		),
		Preflight: key.NewBinding(
			key.WithKeys("p"),
			key.WithHelp("p", "preflight"),
		),
		CopyCmd: key.NewBinding(
			key.WithKeys("c"),
			key.WithHelp("c", "copy cmd"),
		),
		CopyLog: key.NewBinding(
			key.WithKeys("y"),
			key.WithHelp("y", "copy log"),
		),
		Mode: key.NewBinding(
			key.WithKeys("m"),
			key.WithHelp("m", "mode"),
		),
		Vars: key.NewBinding(
			key.WithKeys("v"),
			key.WithHelp("v", "vars"),
		),
		TriageRefresh: key.NewBinding(
			key.WithKeys("t"),
			key.WithHelp("t", "refresh triage"),
		),
		SyncKnownIssues: key.NewBinding(
			key.WithKeys("K"),
			key.WithHelp("K", "sync known issues"),
		),
		FileIssue: key.NewBinding(
			key.WithKeys("i"),
			key.WithHelp("i", "file selected issue"),
		),
	}
}

// ShortHelp returns the bindings shown in the compact footer.
func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Up, k.Down, k.NextSection, k.PrevSection,
		k.PreflightTab, k.GroupsTab, k.RunTab, k.TriageTab,
		k.DrillIn, k.Back, k.Help, k.Quit,
	}
}

// FullHelp returns all bindings grouped into columns for the expanded overlay.
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Top, k.Bottom},
		{k.NextSection, k.PrevSection, k.PreflightTab, k.GroupsTab},
		{k.RunTab, k.TriageTab, k.DrillIn, k.Back},
		{k.Retry, k.RetryAll, k.FailuresFirst, k.ShowAll},
		{k.Orphans, k.Export, k.Sweep, k.Preflight},
		{k.CopyCmd, k.CopyLog, k.Mode, k.Vars},
		{k.TriageRefresh, k.SyncKnownIssues, k.FileIssue},
		{k.Help, k.Quit},
	}
}
