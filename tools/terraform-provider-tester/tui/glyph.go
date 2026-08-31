package tui

import "github.com/github/terraform-provider-tester/provider"

// rowStatus is a presentation-layer status enum for the TUI.
// It extends provider.Status with view-only states (notRun, flaky).
type rowStatus int

const (
	rowPass rowStatus = iota
	rowFail
	rowPanic
	rowTimeout
	rowSkip
	rowRunning
	rowNotRun
	rowFlaky
)

// fromStatus converts a provider.Status to the presentation rowStatus.
func fromStatus(s provider.Status) rowStatus {
	switch s {
	case provider.StatusPass:
		return rowPass
	case provider.StatusFail:
		return rowFail
	case provider.StatusPanic:
		return rowPanic
	case provider.StatusTimeout:
		return rowTimeout
	case provider.StatusSkip:
		return rowSkip
	case provider.StatusRunning:
		return rowRunning
	default:
		return rowNotRun
	}
}

// glyphEntry pairs a display glyph with its semantic color token.
type glyphEntry struct {
	Glyph string
	Color interface{} // lipgloss.TerminalColor (AdaptiveColor or NoColor)
}

// ttyGlyphs maps rowStatus → TTY glyph + color.
var ttyGlyphs = map[rowStatus]glyphEntry{
	rowPass:    {Glyph: "✓", Color: Success},
	rowFail:    {Glyph: "✗", Color: Danger},
	rowPanic:   {Glyph: "‼", Color: Danger},
	rowTimeout: {Glyph: "⏱", Color: Danger},
	rowSkip:    {Glyph: "-", Color: Muted},
	rowRunning: {Glyph: "*", Color: Attention}, // spinner in real use; star for static frames
	rowNotRun:  {Glyph: "·", Color: Muted},
	rowFlaky:   {Glyph: "~", Color: Flaky},
}

// asciiGlyphs maps rowStatus → ASCII glyph + color.
var asciiGlyphs = map[rowStatus]glyphEntry{
	rowPass:    {Glyph: "+", Color: Success},
	rowFail:    {Glyph: "X", Color: Danger},
	rowPanic:   {Glyph: "!!", Color: Danger},
	rowTimeout: {Glyph: "T", Color: Danger},
	rowSkip:    {Glyph: "-", Color: Muted},
	rowRunning: {Glyph: "*", Color: Attention},
	rowNotRun:  {Glyph: ".", Color: Muted},
	rowFlaky:   {Glyph: "~", Color: Flaky},
}

// glyphFor returns the glyph entry for the given rowStatus using the selected glyph set.
// When ascii is true, the ASCII set is used; otherwise the TTY set is used.
func glyphFor(rs rowStatus, ascii bool) glyphEntry {
	if ascii {
		if e, ok := asciiGlyphs[rs]; ok {
			return e
		}
		return glyphEntry{Glyph: ".", Color: Muted}
	}
	if e, ok := ttyGlyphs[rs]; ok {
		return e
	}
	return glyphEntry{Glyph: "·", Color: Muted}
}

// checkGlyph maps a provider.CheckStatus to a glyph entry.
// Preflight reuses the same color tokens as the status table (§2.2).
func checkGlyph(cs provider.CheckStatus, ascii bool) glyphEntry {
	switch cs {
	case provider.CheckOK:
		if ascii {
			return glyphEntry{Glyph: "+", Color: Success}
		}
		return glyphEntry{Glyph: "✓", Color: Success}
	case provider.CheckWarn:
		if ascii {
			return glyphEntry{Glyph: "!", Color: Attention}
		}
		return glyphEntry{Glyph: "⚠", Color: Attention}
	default: // CheckFail
		if ascii {
			return glyphEntry{Glyph: "X", Color: Danger}
		}
		return glyphEntry{Glyph: "✗", Color: Danger}
	}
}
