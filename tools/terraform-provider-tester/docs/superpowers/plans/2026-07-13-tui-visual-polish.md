# Terraform Provider Tester TUI Visual Polish Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Polish the optional four-tab dashboard into a responsive GitHub-style mission-control TUI, fix direct horizontal navigation, refresh all dashboard documentation, and deliver the rebased branch as an AI-assisted PR for Robert.

**Architecture:** Preserve the Bubble Tea pure-reducer boundary and existing CLI services. Add pure layout/chrome primitives, derive all metrics and badges from existing model state, split large view responsibilities into focused files, and leave provider, runner, state, redaction, sweep, and issue semantics untouched.

**Tech Stack:** Go 1.26, Bubble Tea, Bubbles, Lip Gloss, golden tests, deterministic synthetic Chrome screenshots, GNU Make, GitHub CLI.

## Global Constraints

- CLI and NDJSON remain the primary automation interfaces.
- The TUI remains an optional pure reducer; network, filesystem, provider, and run side effects remain in CLI services.
- Exact confirmations remain case-sensitive: `SWEEP <owner>` and `FILE <short-fingerprint>`.
- Preserve `.pulsar*`, `PULSAR_*`, `pulsar-fp-v1`, `pulsar/pre-run`, and the known-issue HTML marker.
- Use no Nerd Font glyphs and add no runtime dependency.
- `NO_COLOR` and ASCII rendering must preserve all meaning.
- Color supplements text and glyphs; it never carries meaning alone.
- Secret values never enter rendered model state, screenshots, docs fixtures, status messages, reports, or errors.
- Existing single-operation, cancellation, locking, redaction, exact-snapshot sweep, and issue-filing behavior must not change.
- Use strict TDD: add one failing behavioral test, observe the expected failure, implement the smallest change, and rerun the focused suite.
- Do not run credentialed acceptance tests, sweep resources, file issues, create releases, or modify `integrations/terraform-provider-github/test/harness`.
- Every commit includes:
  `Co-authored-by: Copilot App <223556219+Copilot@users.noreply.github.com>`.

## File Structure

- Create `tui/layout.go`: responsive breakpoints, metrics, panel composition, width helpers.
- Create `tui/layout_test.go`: breakpoint, panel, metric, and width invariants.
- Create `tui/chrome.go`: product header, tab summaries, tabs, contextual footer, expanded help, status rail.
- Create `tui/chrome_test.go`: tab badges, responsive chrome, contextual help, and width invariants.
- Create `tui/preflight_view.go`: Preflight metrics, checks, next-action panel, empty state.
- Create `tui/preflight_view_test.go`: ready/warn/blocked and responsive Preflight views.
- Create `tui/run_view.go`: Run metrics, progress card, target/actions/last-activity panels.
- Create `tui/run_view_test.go`: idle/running/completed/failed Run views.
- Modify `tui/view.go`: top-level composition plus Groups/test/log rendering only; remove moved chrome, Preflight, and Run functions.
- Modify `tui/keys.go`: Left/Right and direct tab bindings; contextual help key maps.
- Modify `tui/update.go`: direct navigation and Escape-to-close expanded help.
- Modify `tui/model.go`: remove the static `help.Model` only if contextual rendering no longer needs it; otherwise retain it with dynamic key maps.
- Modify `tui/styles.go`, `tui/theme.go`: semantic panel, chip, metric, breadcrumb, action, and destructive styles.
- Modify `tui/triage.go`: metric strip, badges, selected row, detail dossier.
- Modify `tui/cleanup.go`, `tui/confirm.go`: shared panel system for overlays and confirmations.
- Modify `tui/docs_screenshot_test.go`: add Run screenshot and assert new chrome.
- Modify `tui/testdata/*.golden`: update only through existing golden generators.
- Modify `README.md`, `docs/index.md`, `docs/quickstart.md`, `docs/architecture.md`, `docs/troubleshooting.md`, `CHANGELOG.md`.
- Add `docs/images/run.png`; regenerate the other four dashboard PNGs.

---

### Task 1: Direct Navigation and Contextual Help

**Files:**
- Modify: `tui/keys.go`
- Modify: `tui/update.go`
- Modify: `tui/view.go`
- Test: `tui/update_test.go`
- Test: `tui/view_test.go`

**Interfaces:**
- Consumes: existing `section`, `numSections`, overlay interception, and Bubbles key bindings.
- Produces: `keyMap.PreflightTab`, `GroupsTab`, `RunTab`, `TriageTab`; Left/Right aliases on section navigation; `Model.contextualKeys() contextualKeyMap`; `Model.renderHelpPanel(int) string`.

- [ ] **Step 1: Add failing direct-navigation tests**

Add to `tui/update_test.go`:

```go
func TestUpdateArrowKeysCycleTabs(t *testing.T) {
	m := newTestModel()
	next, _ := m.Update(tea.KeyMsg{Type: tea.KeyRight})
	m = next.(Model)
	if m.section != sectionGroups {
		t.Fatalf("Right section = %d, want Groups", m.section)
	}
	next, _ = m.Update(tea.KeyMsg{Type: tea.KeyLeft})
	m = next.(Model)
	if m.section != sectionPreflight {
		t.Fatalf("Left section = %d, want Preflight", m.section)
	}
	next, _ = m.Update(tea.KeyMsg{Type: tea.KeyLeft})
	m = next.(Model)
	if m.section != sectionTriage {
		t.Fatalf("wrapped Left section = %d, want Triage", m.section)
	}
}

func TestUpdateNumberKeysJumpToTabs(t *testing.T) {
	tests := []struct {
		key  rune
		want section
	}{{'1', sectionPreflight}, {'2', sectionGroups}, {'3', sectionRun}, {'4', sectionTriage}}
	for _, tc := range tests {
		m := newTestModel()
		next, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{tc.key}})
		if got := next.(Model).section; got != tc.want {
			t.Fatalf("key %q section = %d, want %d", tc.key, got, tc.want)
		}
	}
}

func TestUpdateEscapeClosesExpandedHelp(t *testing.T) {
	m := newTestModel()
	next, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("?")})
	m = next.(Model)
	if !m.help.ShowAll {
		t.Fatal("help did not open")
	}
	next, _ = m.Update(tea.KeyMsg{Type: tea.KeyEsc})
	if next.(Model).help.ShowAll {
		t.Fatal("Escape did not close help")
	}
}
```

- [ ] **Step 2: Run navigation tests and observe RED**

Run:

```bash
go test ./tui -run 'TestUpdate(ArrowKeysCycleTabs|NumberKeysJumpToTabs|EscapeClosesExpandedHelp)$' -count=1
```

Expected: failures because Right/Left and number keys do not change sections and
Escape does not close expanded help.

- [ ] **Step 3: Add navigation bindings**

Update `keyMap` and `defaultKeys()` in `tui/keys.go`:

```go
type keyMap struct {
	Up, Down, Top, Bottom             key.Binding
	NextSection, PrevSection          key.Binding
	PreflightTab, GroupsTab           key.Binding
	RunTab, TriageTab                 key.Binding
	DrillIn, Back, Help, Quit         key.Binding
	Retry, RetryAll                   key.Binding
	FailuresFirst, ShowAll            key.Binding
	Orphans, Export, Sweep, Preflight key.Binding
	CopyCmd, CopyLog, Mode, Vars      key.Binding
	TriageRefresh, SyncKnownIssues    key.Binding
	FileIssue                         key.Binding
}

NextSection: key.NewBinding(
	key.WithKeys("tab", "right"),
	key.WithHelp("→/tab", "next tab"),
),
PrevSection: key.NewBinding(
	key.WithKeys("shift+tab", "left"),
	key.WithHelp("←/shift+tab", "previous tab"),
),
PreflightTab: key.NewBinding(key.WithKeys("1"), key.WithHelp("1", "preflight")),
GroupsTab:    key.NewBinding(key.WithKeys("2"), key.WithHelp("2", "groups")),
RunTab:       key.NewBinding(key.WithKeys("3"), key.WithHelp("3", "run")),
TriageTab:    key.NewBinding(key.WithKeys("4"), key.WithHelp("4", "triage")),
```

Keep existing bindings unchanged.

- [ ] **Step 4: Implement direct navigation and help close**

In `tui/update.go`, immediately after resume interception and before the normal
switch:

```go
if m.help.ShowAll && key.Matches(msg, m.keys.Back) {
	m.help.ShowAll = false
	return m, nil
}
```

Add direct cases after previous-section handling:

```go
case key.Matches(msg, m.keys.PreflightTab):
	m.switchSection(sectionPreflight)
case key.Matches(msg, m.keys.GroupsTab):
	m.switchSection(sectionGroups)
case key.Matches(msg, m.keys.RunTab):
	m.switchSection(sectionRun)
case key.Matches(msg, m.keys.TriageTab):
	m.switchSection(sectionTriage)
```

Add this pure helper:

```go
func (m *Model) switchSection(next section) {
	m.section = next
	m.cursor = 0
	m.focus = focusGroups
	m.triageDetailActive = false
}
```

Use `switchSection` for next and previous cases as well.

- [ ] **Step 5: Add failing contextual-help view test**

Add to `tui/view_test.go`:

```go
func TestContextualFooterAdvertisesArrowTabNavigation(t *testing.T) {
	m := fixedModel()
	got := stripANSI(m.renderFooter(m.width))
	for _, want := range []string{"←/→", "tabs", "? help", "q quit"} {
		if !strings.Contains(got, want) {
			t.Fatalf("footer missing %q:\n%s", want, got)
		}
	}
}

func TestExpandedHelpOwnsBodyAndGroupsCommands(t *testing.T) {
	m := fixedModel()
	m.help.ShowAll = true
	got := stripANSI(m.View())
	for _, want := range []string{
		"Keyboard reference",
		"Navigation",
		"Run and retry",
		"Triage and known issues",
		"Reports and cleanup",
		"Configuration",
		"esc close",
	} {
		if !strings.Contains(got, want) {
			t.Fatalf("help missing %q:\n%s", want, got)
		}
	}
}
```

- [ ] **Step 6: Run contextual-help tests and observe RED**

Run:

```bash
go test ./tui -run 'Test(ContextualFooterAdvertisesArrowTabNavigation|ExpandedHelpOwnsBodyAndGroupsCommands)$' -count=1
```

Expected: footer lacks the compact navigation copy and expanded help does not
own the body.

- [ ] **Step 7: Implement contextual footer and help panel**

Add to `tui/view.go`:

```go
func (m Model) renderFooter(width int) string {
	hints := []string{"←/→ tabs", "↑/↓ move"}
	switch m.section {
	case sectionPreflight:
		hints = append(hints, "p preflight", "m mode", "v vars")
	case sectionGroups:
		hints = append(hints, "enter detail", "r run/retry", "f failures")
	case sectionRun:
		hints = append(hints, "r run", "R retry failed", "e export")
	case sectionTriage:
		hints = append(hints, "enter detail", "t refresh", "K sync")
	}
	hints = append(hints, "? help", "q quit")
	return renderKeyHints(width, hints)
}

func renderKeyHints(width int, hints []string) string {
	line := strings.Join(hints, "  •  ")
	return Styles.HelpFooter.Render(truncate(line, max(1, width)))
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
		blocks = append(blocks, renderHelpGroup(group.title, group.lines))
	}
	return renderPanel(width, "Keyboard reference", "All actions are scoped to the current view.", strings.Join(blocks, "\n\n"), panelNeutral, m.ascii) +
		"\n" + Styles.HelpFooter.Render("  ? or esc close")
}
```

In `View()`, render expanded help before normal body selection:

```go
if m.help.ShowAll {
	body = m.renderHelpPanel(width)
} else if m.pickerActive {
	// existing overlay precedence
}
```

`renderPanel`, `panelNeutral`, and `renderHelpGroup` are introduced in Task 2;
until Task 2 lands, use a local plain bordered style with the same signature and
move it in Task 2 without changing output.

- [ ] **Step 8: Run Task 1 tests**

Run:

```bash
go test ./tui -run 'TestUpdate(TabCyclesSections|ArrowKeysCycleTabs|NumberKeysJumpToTabs|EscapeClosesExpandedHelp)|Test(ContextualFooterAdvertisesArrowTabNavigation|ExpandedHelpOwnsBodyAndGroupsCommands)' -count=1
```

Expected: PASS.

- [ ] **Step 9: Commit Task 1**

```bash
git add tui/keys.go tui/update.go tui/update_test.go tui/view.go tui/view_test.go
git commit -m "feat(tui): add direct tab navigation and contextual help" \
  -m "Co-authored-by: Copilot App <223556219+Copilot@users.noreply.github.com>"
```

---

### Task 2: Responsive Layout Primitives and Product Chrome

**Files:**
- Create: `tui/layout.go`
- Create: `tui/layout_test.go`
- Create: `tui/chrome.go`
- Create: `tui/chrome_test.go`
- Modify: `tui/theme.go`
- Modify: `tui/styles.go`
- Modify: `tui/view.go`

**Interfaces:**
- Consumes: `Model`, semantic colors, `summarize`, `aggregateStatus`, existing dimensions.
- Produces:
  - `type layoutMode int`
  - `layoutForWidth(int) layoutMode`
  - `type panelTone int`
  - `renderPanel(width int, title, subtitle, body string, tone panelTone, ascii bool) string`
  - `type metricSpec struct { Label, Value string; Tone panelTone }`
  - `renderMetricRow(width int, metrics []metricSpec, ascii bool) string`
  - `Model.tabSummaries() [numSections]tabSummary`
  - moved `renderHeader`, `renderTabs`, `renderStatus`, `renderFooter`, `renderHelpPanel`.

- [ ] **Step 1: Write failing breakpoint and width tests**

Create `tui/layout_test.go`:

```go
package tui

import (
	"strings"
	"testing"

	"github.com/charmbracelet/lipgloss"
)

func TestLayoutForWidthUsesDeterministicBreakpoints(t *testing.T) {
	tests := []struct {
		width int
		want  layoutMode
	}{{71, layoutCompact}, {72, layoutStandard}, {109, layoutStandard}, {110, layoutWide}}
	for _, tc := range tests {
		if got := layoutForWidth(tc.width); got != tc.want {
			t.Fatalf("layoutForWidth(%d) = %d, want %d", tc.width, got, tc.want)
		}
	}
}

func TestPanelAndMetricRowsFitWidth(t *testing.T) {
	for _, width := range []int{52, 80, 120} {
		got := renderPanel(width, "Ready", "Environment status", "All checks passed.", panelSuccess, true)
		assertRenderedWidth(t, got, width)
		metrics := renderMetricRow(width, []metricSpec{
			{Label: "READY", Value: "6", Tone: panelSuccess},
			{Label: "BLOCKED", Value: "2", Tone: panelDanger},
		}, true)
		assertRenderedWidth(t, metrics, width)
	}
}

func assertRenderedWidth(t *testing.T, rendered string, width int) {
	t.Helper()
	for i, line := range strings.Split(stripANSI(rendered), "\n") {
		if got := lipgloss.Width(line); got > width {
			t.Fatalf("line %d width = %d, want <= %d:\n%s", i, got, width, rendered)
		}
	}
}
```

- [ ] **Step 2: Run layout tests and observe RED**

Run:

```bash
go test ./tui -run 'Test(LayoutForWidthUsesDeterministicBreakpoints|PanelAndMetricRowsFitWidth)$' -count=1
```

Expected: compile failure because the layout primitives do not exist.

- [ ] **Step 3: Implement layout primitives**

Create `tui/layout.go` with:

```go
package tui

import (
	"fmt"
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

func renderPanel(width int, title, subtitle, body string, tone panelTone, ascii bool) string {
	width = max(12, width)
	border := lipgloss.RoundedBorder()
	if ascii {
		border = lipgloss.NormalBorder()
	}
	inner := max(1, width-4)
	titleLine := lipgloss.NewStyle().Bold(true).Foreground(toneColor(tone)).Render(truncate(title, inner))
	parts := []string{titleLine}
	if subtitle != "" {
		parts = append(parts, lipgloss.NewStyle().Foreground(Muted).Render(truncate(subtitle, inner)))
	}
	if body != "" {
		parts = append(parts, body)
	}
	return lipgloss.NewStyle().
		Width(inner).
		Border(border).
		BorderForeground(toneColor(tone)).
		Padding(0, 1).
		Render(lipgloss.JoinVertical(lipgloss.Left, parts...))
}

func renderMetricRow(width int, metrics []metricSpec, ascii bool) string {
	if len(metrics) == 0 {
		return ""
	}
	gap := 2
	cellWidth := max(8, (width-gap*(len(metrics)-1))/len(metrics))
	cells := make([]string, 0, len(metrics))
	for _, metric := range metrics {
		value := lipgloss.NewStyle().Bold(true).Foreground(toneColor(metric.Tone)).Render(metric.Value)
		label := lipgloss.NewStyle().Foreground(Muted).Render(metric.Label)
		cells = append(cells, lipgloss.NewStyle().Width(cellWidth).Render(fmt.Sprintf("%s %s", label, value)))
	}
	return truncate(strings.Join(cells, strings.Repeat(" ", gap)), width)
}

func renderHelpGroup(title string, lines []string) string {
	return Styles.SectionTitle.Render(title) + "\n  " + strings.Join(lines, "\n  ")
}
```

- [ ] **Step 4: Add semantic styles**

Extend `Styles` in `tui/styles.go` with:

```go
SectionTitle lipgloss.Style
Chip         lipgloss.Style
MetricLabel  lipgloss.Style
Breadcrumb   lipgloss.Style
Action       lipgloss.Style
DangerPanel  lipgloss.Style
```

Initialize them:

```go
SectionTitle: lipgloss.NewStyle().Bold(true),
Chip:         lipgloss.NewStyle().Padding(0, 1).Foreground(Muted).Border(lipgloss.RoundedBorder()).BorderForeground(Border),
MetricLabel:  lipgloss.NewStyle().Foreground(Muted),
Breadcrumb:   lipgloss.NewStyle().Foreground(Muted),
Action:       lipgloss.NewStyle().Foreground(Accent).Bold(true),
DangerPanel:  lipgloss.NewStyle().BorderForeground(Danger),
```

- [ ] **Step 5: Write failing chrome tests**

Create `tui/chrome_test.go`:

```go
package tui

import (
	"strings"
	"testing"

	"github.com/github/terraform-provider-tester/engine"
	"github.com/github/terraform-provider-tester/provider"
)

func TestTabSummariesDeriveBadges(t *testing.T) {
	m := fixedModel()
	m.checks = []provider.Check{{Status: provider.CheckOK}, {Status: provider.CheckFail}}
	m.groups = []engine.Group{{Name: "a", Tests: []string{"TestA", "TestB"}}}
	m.triageFailures = sampleFailures()
	got := m.tabSummaries()
	if got[0].Badge != "1 blocked" || got[1].Badge != "1" || got[2].Badge != "2" || got[3].Badge != "4" {
		t.Fatalf("tab badges = %#v", got)
	}
}

func TestChromeFitsResponsiveWidths(t *testing.T) {
	for _, width := range []int{60, 90, 140} {
		m := fixedModel()
		m.width = width
		assertRenderedWidth(t, m.renderHeader(width), width)
		assertRenderedWidth(t, m.renderTabs(width), width)
		assertRenderedWidth(t, m.renderFooter(width), width)
	}
}

func TestTabsAdvertiseNumberedSections(t *testing.T) {
	got := stripANSI(fixedModel().renderTabs(100))
	for _, want := range []string{"1 Preflight", "2 Groups", "3 Run", "4 Triage"} {
		if !strings.Contains(got, want) {
			t.Fatalf("tabs missing %q:\n%s", want, got)
		}
	}
}
```

- [ ] **Step 6: Run chrome tests and observe RED**

Run:

```bash
go test ./tui -run 'Test(TabSummariesDeriveBadges|ChromeFitsResponsiveWidths|TabsAdvertiseNumberedSections)$' -count=1
```

Expected: compile failure because `tabSummaries` and new chrome behavior do not
exist.

- [ ] **Step 7: Extract and implement chrome**

Create `tui/chrome.go`. Move `renderHeader`, `renderRateGauge`, `renderTabs`,
`renderStatus`, `renderFooter`, and `renderHelpPanel` from `view.go`.

Define:

```go
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
```

Render a compact header with `◆ TESTER`, version, provider/mode/owner chips, and
the existing rate gauge. Compact layout hides owner and reset time.

Render numbered tabs from `tabSummaries`; compact layout shortens labels to
`Pre`, `Grp`, `Run`, `Tri` but preserves numbers and badges.

- [ ] **Step 8: Remove duplicate chrome from `view.go`**

Delete the moved functions and unused imports from `tui/view.go`. Keep
`Model.View`, section dispatch, Groups/test/log rendering, and shared
`truncate`.

- [ ] **Step 9: Run Task 2 tests and package suite**

Run:

```bash
gofmt -w tui/layout.go tui/layout_test.go tui/chrome.go tui/chrome_test.go tui/theme.go tui/styles.go tui/view.go
go test ./tui -run 'Test(Layout|Panel|Tab|Chrome|Contextual|Expanded)' -count=1
UPDATE_GOLDEN=1 go test ./tui -run 'TestView|TestTriage|Test.*Golden' -count=1
go test ./tui -count=1
```

Expected: PASS. Update only goldens directly changed by the new shared chrome;
Task 6 performs the final complete regeneration after every polished view lands.

- [ ] **Step 10: Commit Task 2**

```bash
git add tui/layout.go tui/layout_test.go tui/chrome.go tui/chrome_test.go tui/theme.go tui/styles.go tui/view.go
git commit -m "feat(tui): add responsive mission control chrome" \
  -m "Co-authored-by: Copilot App <223556219+Copilot@users.noreply.github.com>"
```

---

### Task 3: Preflight and Run Mission-Control Views

**Files:**
- Create: `tui/preflight_view.go`
- Create: `tui/preflight_view_test.go`
- Create: `tui/run_view.go`
- Create: `tui/run_view_test.go`
- Modify: `tui/view.go`

**Interfaces:**
- Consumes: Task 2 panel/metric primitives and existing `Model` state.
- Produces:
  - `type preflightSummary struct { OK, Warn, Fail int }`
  - `Model.preflightSummary() preflightSummary`
  - `Model.renderPreflight(int) string`
  - `type runSummary struct { Total, Passed, Failed, Running, NotRun int; Duration float64 }`
  - `Model.runSummary() runSummary`
  - `Model.renderRunConsole(int) string`

- [ ] **Step 1: Add failing Preflight tests**

Create `tui/preflight_view_test.go`:

```go
package tui

import (
	"strings"
	"testing"

	"github.com/github/terraform-provider-tester/provider"
)

func TestPreflightSummaryCountsStatuses(t *testing.T) {
	m := fixedModel()
	m.checks = []provider.Check{
		{Status: provider.CheckOK},
		{Status: provider.CheckWarn},
		{Status: provider.CheckFail},
		{Status: provider.CheckFail},
	}
	if got := m.preflightSummary(); got.OK != 1 || got.Warn != 1 || got.Fail != 2 {
		t.Fatalf("summary = %#v", got)
	}
}

func TestPreflightViewShowsMetricsAndFirstBlockingFix(t *testing.T) {
	m := fixedModel()
	m.checks = []provider.Check{
		{Name: "identity", Status: provider.CheckOK, Detail: "authenticated"},
		{Name: "owner", Status: provider.CheckFail, Detail: "GITHUB_OWNER not set", Fix: "set GITHUB_OWNER=<org>"},
		{Name: "template", Status: provider.CheckFail, Detail: "template missing", Fix: "create terraform-template-module"},
	}
	got := stripANSI(m.renderPreflight(120))
	for _, want := range []string{"READY 1", "BLOCKED 2", "Prerequisite checks", "Next action", "set GITHUB_OWNER=<org>", "m mode", "v variables", "p rerun"} {
		if !strings.Contains(got, want) {
			t.Fatalf("preflight missing %q:\n%s", want, got)
		}
	}
}

func TestPreflightEmptyStateHasNextAction(t *testing.T) {
	got := stripANSI(fixedModel().renderPreflight(80))
	if !strings.Contains(got, "Preflight has not run yet") || !strings.Contains(got, "p run preflight") {
		t.Fatalf("unexpected empty state:\n%s", got)
	}
}
```

- [ ] **Step 2: Run Preflight tests and observe RED**

Run:

```bash
go test ./tui -run 'TestPreflight(SummaryCountsStatuses|ViewShowsMetricsAndFirstBlockingFix|EmptyStateHasNextAction)$' -count=1
```

Expected: compile failure because `preflightSummary` does not exist and the
view lacks metrics/next action.

- [ ] **Step 3: Implement Preflight view**

Create `tui/preflight_view.go`. Move `renderPreflight` from `view.go` and
implement:

```go
type preflightSummary struct {
	OK, Warn, Fail int
}

func (m Model) preflightSummary() preflightSummary {
	var summary preflightSummary
	for _, check := range m.checks {
		switch check.Status {
		case provider.CheckOK:
			summary.OK++
		case provider.CheckWarn:
			summary.Warn++
		default:
			summary.Fail++
		}
	}
	return summary
}

func (m Model) firstPreflightAction() (string, panelTone) {
	for _, check := range m.checks {
		if check.Status == provider.CheckFail {
			if check.Fix != "" {
				return check.Fix, panelDanger
			}
			return check.Detail, panelDanger
		}
	}
	for _, check := range m.checks {
		if check.Status == provider.CheckWarn {
			if check.Fix != "" {
				return check.Fix, panelAttention
			}
			return check.Detail, panelAttention
		}
	}
	return "Environment is ready. Open Groups to select a run.", panelSuccess
}
```

`renderPreflight` must:

- show a four-item metric row;
- render checks in a panel;
- render first action plus `m mode  v variables  p rerun`;
- use `lipgloss.JoinHorizontal` at `layoutWide`, otherwise stack;
- fit every line to width.

- [ ] **Step 4: Add failing Run tests**

Create `tui/run_view_test.go`:

```go
package tui

import (
	"strings"
	"testing"

	"github.com/github/terraform-provider-tester/engine"
	"github.com/github/terraform-provider-tester/provider"
)

func runViewModel() Model {
	m := fixedModel()
	m.groups = []engine.Group{{Name: "repositories", Tests: []string{"TestA", "TestB", "TestC", "TestD"}}}
	m.results = []engine.TestResult{
		{Name: "TestA", Status: provider.StatusPass, Elapsed: 1.2},
		{Name: "TestB", Status: provider.StatusFail, Elapsed: 2.3},
		{Name: "TestC", Status: provider.StatusRunning},
	}
	m.resultIndex = map[resultKey]int{
		{Name: "TestA"}: 0,
		{Name: "TestB"}: 1,
		{Name: "TestC"}: 2,
	}
	return m
}

func TestRunSummaryDerivesCurrentState(t *testing.T) {
	got := runViewModel().runSummary()
	if got.Total != 4 || got.Passed != 1 || got.Failed != 1 || got.Running != 1 || got.NotRun != 1 {
		t.Fatalf("summary = %#v", got)
	}
}

func TestRunViewShowsProgressScopeAndActions(t *testing.T) {
	m := runViewModel()
	m.running = true
	m.runLabel = "repositories"
	got := stripANSI(m.renderRunConsole(120))
	for _, want := range []string{"Run progress", "50%", "PASSED 1", "FAILED 1", "RUNNING 1", "REMAINING 1", "Target and scope", "Safe actions", "repositories"} {
		if !strings.Contains(got, want) {
			t.Fatalf("run view missing %q:\n%s", want, got)
		}
	}
}

func TestRunEmptyStatePointsToGroups(t *testing.T) {
	got := stripANSI(fixedModel().renderRunConsole(80))
	if !strings.Contains(got, "No tests discovered") || !strings.Contains(got, "2 open Groups") {
		t.Fatalf("unexpected Run empty state:\n%s", got)
	}
}
```

- [ ] **Step 5: Run Run tests and observe RED**

Run:

```bash
go test ./tui -run 'TestRun(SummaryDerivesCurrentState|ViewShowsProgressScopeAndActions|EmptyStatePointsToGroups)$' -count=1
```

Expected: compile failure because `runSummary` does not exist and current Run
view lacks progress panels.

- [ ] **Step 6: Implement Run view**

Create `tui/run_view.go`. Move `renderRunConsole` from `view.go`.

```go
type runSummary struct {
	Total, Passed, Failed, Running, NotRun int
	Duration                              float64
}

func (m Model) runSummary() runSummary {
	var out runSummary
	for _, group := range m.groups {
		summary := summarize(group, m.results)
		out.Total += summary.Total
		out.Passed += summary.Passed
		out.Failed += summary.Failed
		out.Running += summary.Running
		out.NotRun += summary.NotRun
		out.Duration += summary.Duration
	}
	return out
}
```

`renderRunConsole` must render:

- metrics for passed, failed, running, remaining;
- a progress panel whose percentage and existing progress bar use
  `(passed + failed) / total`; running and not-run tests are incomplete;
- target/scope/command panel;
- actions panel (`r run selection`, `R retry failures`, `e export report`);
- current `runLabel`, elapsed, and quiet time while running;
- wide horizontal panels and stacked standard/compact panels.

- [ ] **Step 7: Remove moved views from `view.go`**

Delete `renderPreflight` and `renderRunConsole` from `tui/view.go`; retain
section dispatch.

- [ ] **Step 8: Run Task 3 tests**

Run:

```bash
gofmt -w tui/preflight_view.go tui/preflight_view_test.go tui/run_view.go tui/run_view_test.go tui/view.go
go test ./tui -run 'Test(Preflight|Run)' -count=1
UPDATE_GOLDEN=1 go test ./tui -run 'TestViewPreflight|TestViewRun' -count=1
go test ./tui -count=1
```

Expected: focused tests and the package suite pass. Update only Preflight and
Run goldens directly changed by this task.

- [ ] **Step 9: Commit Task 3**

```bash
git add tui/preflight_view.go tui/preflight_view_test.go tui/run_view.go tui/run_view_test.go tui/view.go
git commit -m "feat(tui): redesign preflight and run dashboards" \
  -m "Co-authored-by: Copilot App <223556219+Copilot@users.noreply.github.com>"
```

---

### Task 4: Groups, Logs, and Triage Presentation

**Files:**
- Modify: `tui/view.go`
- Modify: `tui/triage.go`
- Test: `tui/view_test.go`
- Test: `tui/triage_test.go`

**Interfaces:**
- Consumes: Task 2 metric, panel, badge, and breadcrumb styles.
- Produces: Groups metric strip, full-row selection, breadcrumbs, Triage metric strip/badges/dossier.

- [ ] **Step 1: Add failing Groups presentation tests**

Add to `tui/view_test.go`:

```go
func TestGroupsViewShowsMissionControlMetricsAndActions(t *testing.T) {
	m := fixedModel()
	m.section = sectionGroups
	m.groups = []engine.Group{
		{Name: "repositories", Tests: []string{"TestA", "TestB"}},
		{Name: "teams", Tests: []string{"TestC"}},
	}
	m.results = []engine.TestResult{
		{Name: "TestA", Status: provider.StatusPass},
		{Name: "TestB", Status: provider.StatusFail},
	}
	m.resultIndex = map[resultKey]int{{Name: "TestA"}: 0, {Name: "TestB"}: 1}
	got := stripANSI(m.renderGroupsSection(100))
	for _, want := range []string{"GROUPS 2", "TESTS 3", "PASSED 1", "FAILED 1", "enter tests", "r run/retry", "f failures first"} {
		if !strings.Contains(got, want) {
			t.Fatalf("groups missing %q:\n%s", want, got)
		}
	}
}

func TestGroupsDrilldownAndLogShowBreadcrumbs(t *testing.T) {
	m := fixedModel()
	m.groups = []engine.Group{{Name: "repositories", Tests: []string{"TestA"}}}
	m.groupCursor = 0
	m.focus = focusTests
	if got := stripANSI(m.renderGroupsSection(100)); !strings.Contains(got, "Groups / repositories / tests") {
		t.Fatalf("tests breadcrumb missing:\n%s", got)
	}
	m.focus = focusLog
	m.testCursor = 0
	m.logVP.SetContent("redacted log")
	if got := stripANSI(m.renderGroupsSection(100)); !strings.Contains(got, "Groups / repositories / TestA / log") || !strings.Contains(got, "REDACTED") {
		t.Fatalf("log breadcrumb or badge missing:\n%s", got)
	}
}
```

- [ ] **Step 2: Run Groups tests and observe RED**

Run:

```bash
go test ./tui -run 'TestGroups(ViewShowsMissionControlMetricsAndActions|DrilldownAndLogShowBreadcrumbs)$' -count=1
```

Expected: missing metrics, action strip, breadcrumbs, and redacted badge.

- [ ] **Step 3: Implement Groups polish**

In `tui/view.go`:

- prepend `renderMetricRow` with groups/tests/passed/failed/running;
- pad selected rows to the table width and render with `Styles.SelectedRow`;
- add `Groups / <group> / tests` breadcrumb;
- add `Groups / <group> / <test> / log` breadcrumb and `REDACTED` chip;
- append contextual action strip;
- reduce progress bar width from 14 to 10 under compact layout.

Use:

```go
func selectedTableRow(line string, selected bool, width int) string {
	if !selected {
		return line
	}
	return Styles.SelectedRow.Width(max(1, width)).Render(line)
}
```

- [ ] **Step 4: Add failing Triage presentation tests**

Add to `tui/triage_test.go`:

```go
func TestTriageListShowsMetricStripAndSemanticBadges(t *testing.T) {
	m := fixedModel()
	m.section = sectionTriage
	m.triageFailures = sampleFailures()
	got := stripANSI(m.renderTriageList(120))
	for _, want := range []string{"REAL", "UNSTABLE", "FLAKES", "KNOWN", "ELIGIBLE", "known #42", "eligible"} {
		if !strings.Contains(got, want) {
			t.Fatalf("triage list missing %q:\n%s", want, got)
		}
	}
}

func TestTriageDetailRendersDossier(t *testing.T) {
	m := fixedModel()
	m.triageFailures = sampleFailures()
	m.triageCursor = 1
	m.triageDetailActive = true
	got := stripANSI(m.renderTriageDetail(100))
	for _, want := range []string{"Failure dossier", "Classification", "Evidence", "Issue state", "canonical", "fingerprint", "reasons", "log"} {
		if !strings.Contains(got, want) {
			t.Fatalf("triage detail missing %q:\n%s", want, got)
		}
	}
}
```

- [ ] **Step 5: Run Triage tests and observe RED**

Run:

```bash
go test ./tui -run 'TestTriage(ListShowsMetricStripAndSemanticBadges|DetailRendersDossier)$' -count=1
```

Expected: missing metric labels and dossier panel titles.

- [ ] **Step 6: Implement Triage polish**

In `tui/triage.go`:

- replace the prose summary line with five `metricSpec` entries;
- render classification and issue action with semantic glyph plus text;
- apply `selectedTableRow` to the selected list row;
- render detail as three panels:
  - Classification
  - Evidence
  - Issue state
- preserve all existing values and issue eligibility behavior.

Do not change `triageCounts`, `issueLabel`, or eligibility semantics except to
add pure label/tone helpers.

- [ ] **Step 7: Run Task 4 tests**

Run:

```bash
gofmt -w tui/view.go tui/view_test.go tui/triage.go tui/triage_test.go
go test ./tui -run 'Test(Groups|Triage)' -count=1
UPDATE_GOLDEN=1 go test ./tui -run 'TestViewGroups|TestTriage' -count=1
go test ./tui -count=1
```

Expected: PASS. Update only Groups, log, and Triage goldens directly changed by
this task.

- [ ] **Step 8: Commit Task 4**

```bash
git add tui/view.go tui/view_test.go tui/triage.go tui/triage_test.go
git commit -m "feat(tui): polish groups logs and triage views" \
  -m "Co-authored-by: Copilot App <223556219+Copilot@users.noreply.github.com>"
```

---

### Task 5: Unified Overlays and Safety Presentation

**Files:**
- Modify: `tui/cleanup.go`
- Modify: `tui/confirm.go`
- Modify: `tui/view.go`
- Test: `tui/cleanup_test.go`
- Test: `tui/confirm_test.go`
- Test: `tui/view_test.go`

**Interfaces:**
- Consumes: `renderPanel`, `panelDanger`, `panelSuccess`, exact phrase helpers.
- Produces: consistent mode/config/result/orphan/confirmation panels with unchanged intent behavior.

- [ ] **Step 1: Add failing overlay consistency tests**

Add focused assertions:

```go
func TestSweepConfirmationUsesDangerPanelAndExactPhrase(t *testing.T) {
	m := fixedModel()
	m.orphanOwner = "my-test-org"
	m.orphans = cleanupTestResources()
	m.sweepConfirmActive = true
	got := stripANSI(m.renderSweepConfirm(100))
	for _, want := range []string{"Destructive action", "Confirm sweep", "SWEEP my-test-org", "Targets", "enter confirm", "esc close"} {
		if !strings.Contains(got, want) {
			t.Fatalf("sweep confirm missing %q:\n%s", want, got)
		}
	}
}

func TestResultOverlayUsesSharedPanel(t *testing.T) {
	m := fixedModel()
	m.resultTitle = "Report exported"
	m.resultLines = []string{"markdown: report.md", "html: report.html"}
	got := stripANSI(m.renderResultOverlay(80))
	for _, want := range []string{"Report exported", "markdown: report.md", "html: report.html", "esc close"} {
		if !strings.Contains(got, want) {
			t.Fatalf("result overlay missing %q:\n%s", want, got)
		}
	}
}
```

- [ ] **Step 2: Run overlay tests and observe RED**

Run:

```bash
go test ./tui -run 'Test(SweepConfirmationUsesDangerPanelAndExactPhrase|ResultOverlayUsesSharedPanel)$' -count=1
```

Expected: sweep confirmation lacks the explicit destructive label and shared
panel structure.

- [ ] **Step 3: Implement shared overlay rendering**

Refactor `tui/cleanup.go`, `tui/confirm.go`, and picker/editor/result rendering
to call `renderPanel`.

Safety rules:

```go
func destructiveSubtitle() string {
	return "Destructive action. Review the snapshot and type the exact phrase."
}
```

- Sweep and issue confirmations use `panelDanger`.
- Orphan list, mode, variables, and help use `panelNeutral` or `panelAccent`.
- Successful result overlays use `panelSuccess`; failed sweep remains
  `panelDanger`.
- Exact phrase text and input remain visible and unchanged.
- No overlay calls CLI, provider, filesystem, or network code.

- [ ] **Step 4: Verify key isolation and exact confirmations**

Run:

```bash
go test ./tui -run 'Test.*(Overlay|Confirmation|Sweep|FileIssue|Picker|Editor).*' -count=1
go test ./cli -run 'TestDashboard(Sweep|IssueFile).*' -count=1
```

Expected: PASS; wrong phrases emit no intent or side effect.

- [ ] **Step 5: Commit Task 5**

```bash
git add tui/cleanup.go tui/cleanup_test.go tui/confirm.go tui/confirm_test.go tui/view.go tui/view_test.go
git commit -m "feat(tui): unify guarded action panels" \
  -m "Co-authored-by: Copilot App <223556219+Copilot@users.noreply.github.com>"
```

---

### Task 6: Goldens, Screenshots, and Documentation

**Files:**
- Modify: `tui/docs_screenshot_test.go`
- Modify: `tui/docs_screenshot_promote_test.go`
- Modify: `tui/view_test.go`
- Modify: `tui/triage_test.go`
- Modify: `tui/testdata/*.golden`
- Modify: `docs/images/intro.png`
- Modify: `docs/images/preflight.png`
- Modify: `docs/images/groups.png`
- Create: `docs/images/run.png`
- Modify: `docs/images/triage.png`
- Modify: `README.md`
- Modify: `docs/index.md`
- Modify: `docs/quickstart.md`
- Modify: `docs/architecture.md`
- Modify: `docs/troubleshooting.md`
- Modify: `CHANGELOG.md`

**Interfaces:**
- Consumes: completed polished views.
- Produces: deterministic five-image documentation set and accurate CLI-first docs.

- [ ] **Step 1: Update screenshot assertions before regenerating images**

In `tui/docs_screenshot_test.go`:

- add `docsRunModel() Model`;
- require compact product header, numbered tabs, and contextual footer;
- require Preflight metrics and Next action;
- require Groups metrics and actions;
- require Run progress/target/actions;
- require Triage metrics and semantic issue states.

Add:

```go
func docsRunModel() Model {
	m := docsGroupsModel()
	m.section = sectionRun
	m.results = []engine.TestResult{
		{Name: "TestAccA", Status: provider.StatusPass, Elapsed: 1.2},
		{Name: "TestAccB", Status: provider.StatusFail, Elapsed: 2.3},
		{Name: "TestAccD", Status: provider.StatusRunning},
	}
	m.resultIndex = map[resultKey]int{
		{Name: "TestAccA"}: 0,
		{Name: "TestAccB"}: 1,
		{Name: "TestAccD"}: 2,
	}
	m.running = true
	m.runLabel = "repositories"
	return m
}
```

Update `TestUpdateDocsScreenshots`:

```go
writeScreenshot(t, docsChromePath, dir, "run", docsRunModel().View())
```

- [ ] **Step 2: Run screenshot source tests and observe RED**

Run:

```bash
go test ./tui -run 'TestUpdateDocsScreenshots|TestDocsScreenshots' -count=1
```

Expected: skipped generator plus failing static image inventory because
`run.png` does not exist or current images do not match the new assertions.

- [ ] **Step 3: Regenerate golden files**

Run:

```bash
UPDATE_GOLDEN=1 go test ./tui -run 'TestView|TestTriage' -count=1
go test ./tui -run 'TestView|TestTriage' -count=1
```

Expected: updated goldens, then PASS without update mode.

- [ ] **Step 4: Regenerate five deterministic PNGs**

Run:

```bash
UPDATE_DOC_IMAGES=1 go test ./tui -run TestUpdateDocsScreenshots -count=1
go test ./tui -run 'TestDocsScreenshots|TestUpdateDocsScreenshots' -count=1
```

Expected: five valid 1600×900 PNGs and no temporary files.

- [ ] **Step 5: Update documentation**

Make these exact content changes:

- README dashboard section:
  - state `Left/Right` or `Tab/Shift+Tab` switch tabs;
  - state `1`–`4` jump directly;
  - state Escape closes expanded help;
  - add `run.png` between Groups and Triage;
  - retain CLI/NDJSON precedence and exact confirmation text.
- `docs/quickstart.md`:
  - add a “Explore the dashboard” block with the same navigation keys.
- `docs/architecture.md`:
  - document `layoutCompact <72`, `layoutStandard 72–109`,
    `layoutWide >=110`;
  - document pure derived metrics/tab badges and shared panel primitives.
- `docs/troubleshooting.md`:
  - add terminal width, `NO_COLOR=1`, overlay focus, and Left/Right notes.
- `docs/index.md`:
  - link directly to the README dashboard section.
- CHANGELOG `[Unreleased]`:
  - direct tab navigation;
  - responsive mission-control panels;
  - contextual help;
  - full synthetic screenshot set.

- [ ] **Step 6: Run docs and image gates**

Run:

```bash
make checkdocs
go test ./tui -run 'TestDocs|TestView|TestTriage|TestUpdate' -count=1
git diff --check
```

Expected: PASS.

- [ ] **Step 7: Commit Task 6**

```bash
git add README.md CHANGELOG.md docs tui/docs_screenshot_test.go tui/docs_screenshot_promote_test.go tui/testdata
git commit -m "docs: show polished tester mission control" \
  -m "Co-authored-by: Copilot App <223556219+Copilot@users.noreply.github.com>"
```

---

### Task 7: Full Verification and Independent Reviews

**Files:**
- Modify only files required by verified failures.

**Interfaces:**
- Consumes: complete polished branch.
- Produces: clean reviewed branch ready for rebase and PR.

- [ ] **Step 1: Run formatting and static gates**

```bash
export PATH="/opt/homebrew/bin:$PATH"
gofmt -w $(git ls-files '*.go')
test -z "$(gofmt -l .)"
go vet ./...
make checkdocs
git diff --check
```

Expected: every command exits 0.

- [ ] **Step 2: Run full tests and race tests**

```bash
go test ./... -count=1
go test -race ./cli ./tui -count=1
```

Expected: all tests pass and no race is reported.

- [ ] **Step 3: Run secret-canary and safety regressions**

```bash
go test ./cli ./engine ./tui -run 'Redact|Secret|Sensitive|Issue|Report|Triage|Confirmation|Sweep' -count=1
```

Expected: PASS with no canary or secret-shaped value in output.

- [ ] **Step 4: Build**

```bash
make build
./bin/terraform-provider-tester version
```

Expected: binary builds and reports Terraform Provider Tester with the current
git-describe version.

- [ ] **Step 5: Run isolated pseudo-TTY smoke**

Use the existing synthetic-provider/tmux pattern from Task 12. Verify:

1. TESTER ignition renders;
2. Right, Left, Tab, Shift+Tab, and `1`–`4` reach all tabs;
3. `?` opens help and Escape closes it;
4. wide and compact terminal captures remain legible;
5. Preflight, Groups, Run, and Triage mission-control panels render;
6. safe `t`, `K`, `e`, and `o` outcomes render;
7. wrong sweep/file phrases have zero side effects;
8. `q` exits cleanly;
9. no acceptance-test child process starts;
10. no secret canary appears.

- [ ] **Step 6: Dispatch independent focused review**

Review the diff from `fee3575` to current HEAD against:

`docs/superpowers/specs/2026-07-13-tui-visual-polish-design.md`

Require explicit:

- Spec compliance PASS/FAIL
- Task quality APPROVED/NEEDS FIXES
- Critical/Important/Minor counts

Fix all validated Critical and Important findings with TDD, then rerun Steps 1
through 5 and re-review.

- [ ] **Step 7: Dispatch final whole-branch review**

Review `c5e431b...HEAD` for:

- TUI/CLI boundary regressions;
- navigation ordering and overlay isolation;
- width overflow;
- secret sinks;
- exact destructive gates;
- cancellation and operation cardinality;
- documentation accuracy;
- release readiness.

Fix validated Critical/Important findings and rerun all gates.

- [ ] **Step 8: Record final evidence**

Record:

- final SHA;
- all gate commands and exit codes;
- pseudo-TTY captures;
- review reports;
- clean `git status --short --branch`.

Do not create an empty verification commit.

---

### Task 8: Merge PR #17, Rebase, Push, and Request Robert

**Files:**
- No source edits unless the rebase produces real conflicts.
- Update PR title/body only through GitHub.

**Interfaces:**
- Consumes: reviewed polished branch; open PR #17 at `c5e431b`.
- Produces: one clean AI-assisted PR from `tui-full-parity` to `main`, with Robert requested.

- [ ] **Step 1: Re-verify PR #17**

Run:

```bash
gh pr view 17 --repo github/terraform-provider-tester \
  --json state,isDraft,mergeStateStatus,reviewDecision,statusCheckRollup,headRefOid
```

Expected:

- state `OPEN`;
- draft `false`;
- merge state `CLEAN`;
- head `c5e431b3aff77ae0e95db0abc59278fc9beccff5`;
- every required check completed successfully.

- [ ] **Step 2: Merge PR #17**

Use squash merge and delete the remote branch:

```bash
gh pr merge 17 --repo github/terraform-provider-tester --squash --delete-branch
```

Then verify:

```bash
gh pr view 17 --repo github/terraform-provider-tester --json state,mergedAt,mergeCommit
```

Expected: state `MERGED`, non-null merge timestamp and commit.

- [ ] **Step 3: Fetch and rebase only parity/polish commits**

From the tester checkout:

```bash
git fetch origin --prune
git rebase --onto origin/main c5e431b3aff77ae0e95db0abc59278fc9beccff5 tui-full-parity
```

This replays only commits after PR #17’s head and prevents duplicate persistence
history after a squash merge.

If conflicts occur:

- preserve `origin/main` persistence behavior;
- preserve all parity/polish behavior from the rebased commit;
- never discard a conflict wholesale;
- rerun focused tests after each resolved logical area.

- [ ] **Step 4: Re-run post-rebase gates**

```bash
gofmt -w $(git ls-files '*.go')
test -z "$(gofmt -l .)"
go vet ./...
go test ./... -count=1
go test -race ./cli ./tui -count=1
make checkdocs
make build
git diff --check
git status --short --branch
```

Expected: all pass and branch is clean.

- [ ] **Step 5: Confirm branch scope**

```bash
git log --oneline origin/main..HEAD
git diff --stat origin/main...HEAD
git diff --check origin/main...HEAD
```

Expected: only TUI parity, visual polish, docs, tests, and their safety fixes.

- [ ] **Step 6: Push branch**

```bash
git push -u origin tui-full-parity
```

Expected: remote branch created without force push.

- [ ] **Step 7: Create the AI-assisted PR**

Create the PR with:

```bash
gh pr create \
  --repo github/terraform-provider-tester \
  --base main \
  --head tui-full-parity \
  --title "[AI-assisted] Complete and polish the TESTER dashboard" \
  --body-file "$PR_BODY_FILE"
```

The PR body must include:

- summary of complete CLI/TUI parity;
- mission-control visual/navigation improvements;
- safety invariants;
- five screenshots;
- exact local verification commands;
- independent review results;
- statement that no credentialed acceptance run, sweep, or issue creation ran;
- follow-up gate for a real organization run before v0.3.0.

- [ ] **Step 8: Request Robert and Copilot reviews**

```bash
gh pr edit "$PR_NUMBER" \
  --repo github/terraform-provider-tester \
  --add-reviewer robert-crandall
```

Copilot automatic review is supplied by the active repository ruleset. Verify
both:

```bash
gh pr view "$PR_NUMBER" --repo github/terraform-provider-tester \
  --json reviewRequests,latestReviews,statusCheckRollup,mergeStateStatus,url
```

Expected: Robert appears in review requests, automatic review is pending or
present, checks are queued/running/successful, and merge state is not blocked by
conflicts.

- [ ] **Step 9: Stop for human review**

Do not merge the polished TUI PR, create a tag, publish v0.3.0, run a
credentialed acceptance suite, sweep, or file issues. Report the PR URL and
current checks/review requests.
