package tui

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/github/terraform-provider-tester/engine"
	"github.com/github/terraform-provider-tester/provider"
)

// itemCount returns the number of items in the current section's list.
// Used to clamp the cursor; returns 0 when a section has no list.
func (m Model) itemCount() int {
	switch m.section {
	case sectionPreflight:
		return len(m.checks)
	case sectionGroups:
		return len(m.groups)
	case sectionRun:
		return len(m.results)
	case sectionTriage:
		return len(m.triageFailures)
	default:
		return 0
	}
}

// clampCursor ensures cursor stays within [0, itemCount-1].
// When the list is empty cursor is forced to 0.
func (m *Model) clampCursor() {
	n := m.itemCount()
	if n == 0 || m.cursor < 0 {
		m.cursor = 0
		return
	}
	if m.cursor >= n {
		m.cursor = n - 1
	}
}

// clampGroupCursor keeps groupCursor in [0, len(groups)-1].
func (m *Model) clampGroupCursor() {
	n := len(m.groups)
	if n == 0 || m.groupCursor < 0 {
		m.groupCursor = 0
		return
	}
	if m.groupCursor >= n {
		m.groupCursor = n - 1
	}
}

// clampTestCursor keeps testCursor in [0, len(tests)-1] for the selected group.
func (m *Model) clampTestCursor() {
	if m.groupCursor < 0 || m.groupCursor >= len(m.groups) {
		m.testCursor = 0
		return
	}
	n := len(m.groups[m.groupCursor].Tests)
	if n == 0 || m.testCursor < 0 {
		m.testCursor = 0
		return
	}
	if m.testCursor >= n {
		m.testCursor = n - 1
	}
}

func (m *Model) clampTriageCursor() {
	n := len(m.triageFailures)
	if n == 0 || m.triageCursor < 0 {
		m.triageCursor = 0
		return
	}
	if m.triageCursor >= n {
		m.triageCursor = n - 1
	}
}

func cloneGroups(in []engine.Group) []engine.Group {
	out := make([]engine.Group, len(in))
	for i, g := range in {
		out[i] = g
		out[i].Tests = append([]string(nil), g.Tests...)
	}
	return out
}

func (m *Model) applyFailuresFirstOrdering() {
	sort.SliceStable(m.groups, func(i, j int) bool {
		return m.groupFailurePriority(m.groups[i]) < m.groupFailurePriority(m.groups[j])
	})
	for i := range m.groups {
		sort.SliceStable(m.groups[i].Tests, func(a, b int) bool {
			return m.testFailurePriority(m.groups[i].Tests[a]) < m.testFailurePriority(m.groups[i].Tests[b])
		})
	}
	m.clampGroupCursor()
	m.clampTestCursor()
}

func (m Model) groupFailurePriority(g engine.Group) int {
	best := statusFailurePriority(rowNotRun)
	for _, test := range g.Tests {
		if p := m.testFailurePriority(test); p < best {
			best = p
		}
	}
	return best
}

func (m Model) testFailurePriority(name string) int {
	if idx, ok := m.resultIndex[resultKey{Name: name, Sub: ""}]; ok {
		return providerStatusFailurePriority(m.results[idx].Status)
	}
	return statusFailurePriority(rowNotRun)
}

func providerStatusFailurePriority(s provider.Status) int {
	switch s {
	case provider.StatusFail, provider.StatusPanic, provider.StatusTimeout:
		return 0
	case provider.StatusRunning:
		return 1
	case provider.StatusPass:
		return 2
	case provider.StatusSkip:
		return 3
	default:
		return 4
	}
}

func statusFailurePriority(s rowStatus) int {
	switch s {
	case rowFail, rowPanic, rowTimeout:
		return 0
	case rowRunning:
		return 1
	case rowPass:
		return 2
	case rowSkip:
		return 3
	default:
		return 4
	}
}

func (m Model) selectedTriageFailure() (engine.PersistFailure, bool) {
	if len(m.triageFailures) == 0 || m.triageCursor < 0 || m.triageCursor >= len(m.triageFailures) {
		return engine.PersistFailure{}, false
	}
	return m.triageFailures[m.triageCursor], true
}

func cloneResources(in []provider.Resource) []provider.Resource {
	return append([]provider.Resource(nil), in...)
}

func (m *Model) clearErrorPresentation() {
	m.err = nil
	m.errOperation = ""
	m.errIsOperationFailure = false
	m.statusError = false
}

func (m *Model) setErrorPresentation(err error, owner string) {
	m.err = err
	m.errOperation = owner
	m.errIsOperationFailure = false
	m.statusError = err != nil
}

func (m *Model) clearOperation(op string) {
	if m.operation == op {
		m.operation = ""
	}
	if m.errOperation != "" && m.errOperation == op {
		m.clearErrorPresentation()
	}
}

func (m *Model) openSweepConfirm() {
	m.sweepInput.SetValue("")
	m.sweepInput.Focus()
	m.sweepConfirmActive = true
}

func (m *Model) closeSweepConfirm() {
	m.sweepInput.SetValue("")
	m.sweepInput.Blur()
	m.sweepConfirmActive = false
}

func (m *Model) openFileIssueConfirm(preview IssuePreviewMsg) {
	m.fileIssuePreview = preview
	m.fileIssueInput.SetValue("")
	m.fileIssueInput.Focus()
	m.fileIssueActive = true
}

func (m *Model) closeFileIssueConfirm() {
	m.fileIssueInput.SetValue("")
	m.fileIssueInput.Blur()
	m.fileIssueActive = false
}

// issueFiledWording maps issueService.FileOne's action outcome to the
// overlay title and status line. Only "filed" claims a new issue was
// created; "dedup" and "known" report that filing was skipped because the
// failure was already tracked. Any other/unknown action gets neutral
// wording so the TUI never falsely claims creation.
func issueFiledWording(action string) (title, status string) {
	switch action {
	case "filed":
		return "Issue filed", "issue filed"
	case "dedup":
		return "Issue deduplicated", "issue deduplicated"
	case "known":
		return "Issue already known", "issue already known"
	default:
		return "Issue action complete", "issue action complete"
	}
}

func (m *Model) setResultOverlay(title string, lines ...string) {
	m.setResultOverlayWithTone(panelSuccess, title, lines...)
}

func (m *Model) setResultOverlayWithTone(tone panelTone, title string, lines ...string) {
	m.resultOverlayActive = true
	m.resultTitle = title
	m.resultLines = append([]string(nil), lines...)
	m.resultTone = tone
}

func (m *Model) clearResultOverlay() {
	m.resultOverlayActive = false
	m.resultTitle = ""
	m.resultLines = nil
	m.resultTone = panelSuccess
}

func (m *Model) switchSection(next section) {
	m.section = next
	m.cursor = 0
	m.focus = focusGroups
	m.triageDetailActive = false
}

// selectedGroupName returns the name of the currently selected group, or "".
func (m Model) selectedGroupName() string {
	if len(m.groups) == 0 || m.groupCursor < 0 || m.groupCursor >= len(m.groups) {
		return ""
	}
	return m.groups[m.groupCursor].Name
}

// selectedTestName returns the name of the currently selected test in the
// selected group, or "".
func (m Model) selectedTestName() string {
	if len(m.groups) == 0 || m.groupCursor < 0 || m.groupCursor >= len(m.groups) {
		return ""
	}
	tests := m.groups[m.groupCursor].Tests
	if len(tests) == 0 || m.testCursor < 0 || m.testCursor >= len(tests) {
		return ""
	}
	return tests[m.testCursor]
}

// loadLogContent loads the selected test's Output into the log viewport.
func (m *Model) loadLogContent() {
	name := m.selectedTestName()
	if name == "" {
		m.logVP.SetContent("")
		return
	}
	k := resultKey{Name: name, Sub: ""}
	if idx, ok := m.resultIndex[k]; ok {
		m.logVP.SetContent(strings.Join(m.results[idx].Output, "\n"))
	} else {
		m.logVP.SetContent("no output")
	}
}

// Update processes a tea.Msg and returns the updated model and optional command.
// It performs NO I/O and contains no business logic.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.help.Width = msg.Width
		// Resize the log viewport to fit inside the body area.
		vpW := msg.Width - 4
		vpH := msg.Height - 6
		if vpW < 1 {
			vpW = 1
		}
		if vpH < 1 {
			vpH = 1
		}
		m.logVP.Width = vpW
		m.logVP.Height = vpH

	case tea.KeyMsg:
		// Intro splash: any key skips the one-shot ignition animation. Handle
		// before everything else so the key is consumed, not leaked into the UI.
		// Quit still quits so ctrl+c is never swallowed by the splash.
		if m.introActive {
			if key.Matches(msg, m.keys.Quit) {
				return m, tea.Quit
			}
			m.introActive = false
			// Clear the splash so the dashboard paints from a clean screen.
			return m, tea.ClearScreen
		}

		// Mode picker overlay: intercept all keys while the picker is open so
		// navigation (and quit) do not leak into the underlying dashboard.
		if m.pickerActive {
			if msg.Type == tea.KeyCtrlC {
				return m, tea.Quit
			}
			switch {
			case key.Matches(msg, m.keys.Quit), key.Matches(msg, m.keys.Back):
				m.pickerActive = false
			case key.Matches(msg, m.keys.Up):
				if m.pickerCursor > 0 {
					m.pickerCursor--
				}
			case key.Matches(msg, m.keys.Down):
				if m.pickerCursor < len(m.pickerModes)-1 {
					m.pickerCursor++
				}
			case key.Matches(msg, m.keys.DrillIn):
				if len(m.pickerModes) > 0 {
					mode := m.pickerModes[m.pickerCursor].Name
					m.mode = mode // optimistic update; confirmed by PreflightMsg
					m.pickerActive = false
					if m.exec != nil {
						return m, m.exec(SwitchModeIntent{Mode: mode})
					}
				}
			}
			return m, nil
		}

		// Variable editor overlay: intercept all keys while the editor is open.
		if m.editorActive {
			if m.editorFocused {
				// A textinput is active. Esc/q close the whole editor; Enter
				// submits the value and emits SetEnvVarIntent; everything else
				// is forwarded to the focused textinput.
				if msg.Type == tea.KeyCtrlC {
					return m, tea.Quit
				}
				switch {
				case key.Matches(msg, m.keys.Quit), key.Matches(msg, m.keys.Back):
					if m.editorCursor < len(m.editorFields) {
						f := m.editorFields[m.editorCursor]
						f.input.Blur()
						m.editorFields[m.editorCursor] = f
					}
					m.editorFocused = false
					m.editorActive = false
					return m, nil
				}
				if msg.Type == tea.KeyEnter {
					if m.editorCursor < len(m.editorFields) {
						f := m.editorFields[m.editorCursor]
						envKey := f.Key
						val := f.input.Value()
						f.input.Blur()
						m.editorFields[m.editorCursor] = f
						m.editorFocused = false
						if m.exec != nil {
							return m, m.exec(SetEnvVarIntent{Key: envKey, Value: val})
						}
					}
					return m, nil
				}
				// Forward all other keys to the focused textinput.
				if m.editorCursor < len(m.editorFields) {
					f := m.editorFields[m.editorCursor]
					var cmd tea.Cmd
					f.input, cmd = f.input.Update(msg)
					m.editorFields[m.editorCursor] = f
					return m, cmd
				}
				return m, nil
			}
			// No textinput is active: handle navigation and field activation.
			if msg.Type == tea.KeyCtrlC {
				return m, tea.Quit
			}
			switch {
			case key.Matches(msg, m.keys.Quit), key.Matches(msg, m.keys.Back):
				m.editorActive = false
			case key.Matches(msg, m.keys.Up):
				if m.editorCursor > 0 {
					m.editorCursor--
				}
			case key.Matches(msg, m.keys.Down):
				if m.editorCursor < len(m.editorFields)-1 {
					m.editorCursor++
				}
			case key.Matches(msg, m.keys.DrillIn):
				if m.editorCursor < len(m.editorFields) {
					f := m.editorFields[m.editorCursor]
					if !f.Secret {
						f.input.Focus()
						m.editorFields[m.editorCursor] = f
						m.editorFocused = true
					}
				}
			case key.Matches(msg, m.keys.Preflight):
				if m.exec != nil {
					return m, m.exec(PreflightIntent{})
				}
			}
			return m, nil
		}

		if m.sweepConfirmActive {
			if msg.Type == tea.KeyCtrlC {
				return m, tea.Quit
			}
			if msg.Type == tea.KeyEsc {
				m.closeSweepConfirm()
				return m, nil
			}
			if msg.Type == tea.KeyEnter {
				phrase := sweepPhrase(m.orphanOwner)
				if phraseMatches(m.sweepInput.Value(), phrase) {
					resources := cloneResources(m.orphans)
					owner := m.orphanOwner
					m.closeSweepConfirm()
					return m, func() tea.Msg {
						return ConfirmSweepIntent{Owner: owner, Phrase: phrase, Resources: resources}
					}
				}
				m.status = "confirmation phrase mismatch"
				return m, nil
			}
			var cmd tea.Cmd
			m.sweepInput, cmd = m.sweepInput.Update(msg)
			return m, cmd
		}

		if m.orphanOverlayActive {
			if msg.Type == tea.KeyCtrlC {
				return m, tea.Quit
			}
			switch {
			case msg.Type == tea.KeyEsc:
				m.orphanOverlayActive = false
			case key.Matches(msg, m.keys.Sweep):
				if len(m.orphans) > 0 && m.operation == "" {
					m.openSweepConfirm()
				}
			}
			return m, nil
		}

		if m.fileIssueActive {
			if msg.Type == tea.KeyCtrlC {
				return m, tea.Quit
			}
			if msg.Type == tea.KeyEsc {
				m.closeFileIssueConfirm()
				return m, nil
			}
			if msg.Type == tea.KeyEnter {
				phrase := fileIssuePhrase(m.fileIssuePreview.ShortFingerprint)
				if phraseMatches(m.fileIssueInput.Value(), phrase) {
					fingerprint := m.fileIssuePreview.Fingerprint
					m.closeFileIssueConfirm()
					return m, func() tea.Msg {
						return ConfirmFileIssueIntent{Fingerprint: fingerprint, Phrase: phrase}
					}
				}
				m.status = "confirmation phrase mismatch"
				return m, nil
			}
			var cmd tea.Cmd
			m.fileIssueInput, cmd = m.fileIssueInput.Update(msg)
			return m, cmd
		}

		if m.resultOverlayActive {
			if msg.Type == tea.KeyCtrlC {
				return m, tea.Quit
			}
			if msg.Type == tea.KeyEsc {
				m.clearResultOverlay()
			}
			return m, nil
		}

		// Resume-prompt key interception: handle before normal navigation so keys
		// don't leak into the rest of the UI while the prompt is displayed.
		if m.resumePrompt {
			switch {
			case key.Matches(msg, m.keys.Quit):
				return m, tea.Quit
			case (key.Matches(msg, m.keys.DrillIn) || key.Matches(msg, m.keys.RetryAll)) && m.operation == "":
				m.resumePrompt = false
				return m, func() tea.Msg { return ResumeIntent{} }
			case key.Matches(msg, m.keys.Back):
				m.resumePrompt = false
				return m, nil
			default:
				// consume all other keys while the prompt is visible
				return m, nil
			}
		}
		if m.help.ShowAll {
			switch {
			case key.Matches(msg, m.keys.Quit):
				return m, tea.Quit
			case key.Matches(msg, m.keys.Help), key.Matches(msg, m.keys.Back):
				m.help.ShowAll = false
				return m, nil
			default:
				return m, nil
			}
		}
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit

		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll

		case key.Matches(msg, m.keys.NextSection):
			m.switchSection((m.section + 1) % numSections)

		case key.Matches(msg, m.keys.PrevSection):
			m.switchSection((m.section + numSections - 1) % numSections)

		case key.Matches(msg, m.keys.PreflightTab):
			m.switchSection(sectionPreflight)

		case key.Matches(msg, m.keys.GroupsTab):
			m.switchSection(sectionGroups)

		case key.Matches(msg, m.keys.RunTab):
			m.switchSection(sectionRun)

		case key.Matches(msg, m.keys.TriageTab):
			m.switchSection(sectionTriage)

		case key.Matches(msg, m.keys.Up):
			if m.section == sectionGroups {
				switch m.focus {
				case focusGroups:
					m.groupCursor--
					m.clampGroupCursor()
				case focusTests:
					m.testCursor--
					m.clampTestCursor()
				case focusLog:
					m.logVP.LineUp(1)
				}
			} else if m.section == sectionTriage {
				m.triageCursor--
				m.clampTriageCursor()
			} else {
				m.cursor--
				m.clampCursor()
			}

		case key.Matches(msg, m.keys.Down):
			if m.section == sectionGroups {
				switch m.focus {
				case focusGroups:
					m.groupCursor++
					m.clampGroupCursor()
				case focusTests:
					m.testCursor++
					m.clampTestCursor()
				case focusLog:
					m.logVP.LineDown(1)
				}
			} else if m.section == sectionTriage {
				m.triageCursor++
				m.clampTriageCursor()
			} else {
				m.cursor++
				m.clampCursor()
			}

		case key.Matches(msg, m.keys.Top):
			if m.section == sectionGroups {
				switch m.focus {
				case focusGroups:
					m.groupCursor = 0
				case focusTests:
					m.testCursor = 0
				case focusLog:
					m.logVP.GotoTop()
				}
			} else if m.section == sectionTriage {
				m.triageCursor = 0
			} else {
				m.cursor = 0
			}

		case key.Matches(msg, m.keys.Bottom):
			if m.section == sectionGroups {
				switch m.focus {
				case focusGroups:
					if n := len(m.groups); n > 0 {
						m.groupCursor = n - 1
					}
				case focusTests:
					if m.groupCursor >= 0 && m.groupCursor < len(m.groups) {
						if n := len(m.groups[m.groupCursor].Tests); n > 0 {
							m.testCursor = n - 1
						}
					}
				case focusLog:
					m.logVP.GotoBottom()
				}
			} else if m.section == sectionTriage {
				if n := len(m.triageFailures); n > 0 {
					m.triageCursor = n - 1
				}
			} else {
				n := m.itemCount()
				if n > 0 {
					m.cursor = n - 1
				}
			}

		case key.Matches(msg, m.keys.DrillIn):
			if m.section == sectionGroups {
				switch m.focus {
				case focusGroups:
					if len(m.groups) > 0 {
						m.focus = focusTests
						m.testCursor = 0
					}
				case focusTests:
					m.loadLogContent()
					m.focus = focusLog
				case focusLog:
					// no-op at deepest level
				}
			} else if m.section == sectionTriage && len(m.triageFailures) > 0 {
				m.triageDetailActive = true
			}

		case key.Matches(msg, m.keys.Back):
			if m.section == sectionGroups {
				switch m.focus {
				case focusLog:
					m.focus = focusTests
				case focusTests:
					m.focus = focusGroups
				case focusGroups:
					// no-op at top level
				}
			} else if m.section == sectionTriage {
				m.triageDetailActive = false
			}

		case key.Matches(msg, m.keys.Retry):
			if m.operation == "" {
				switch m.section {
				case sectionRun:
					if name := m.selectedGroupName(); name != "" {
						return m, func() tea.Msg { return RetryGroupIntent{Group: name} }
					}
				case sectionGroups:
					switch m.focus {
					case focusGroups:
						if name := m.selectedGroupName(); name != "" {
							return m, func() tea.Msg { return RetryGroupIntent{Group: name} }
						}
					case focusTests, focusLog:
						if name := m.selectedTestName(); name != "" {
							return m, func() tea.Msg { return RetryTestIntent{Test: name} }
						}
					}
				}
			}

		case key.Matches(msg, m.keys.RetryAll):
			if (m.section == sectionGroups || m.section == sectionRun) && len(m.groups) > 0 && m.operation == "" {
				return m, func() tea.Msg { return RetryAllIntent{} }
			}

		case key.Matches(msg, m.keys.FailuresFirst):
			if m.section == sectionGroups {
				m.failuresFirst = true
				m.applyFailuresFirstOrdering()
			}

		case key.Matches(msg, m.keys.ShowAll):
			if m.section == sectionGroups {
				m.failuresFirst = false
				m.groups = cloneGroups(m.discoveryGroups)
				m.clampGroupCursor()
				m.clampTestCursor()
			}

		case key.Matches(msg, m.keys.Orphans):
			if m.operation == "" {
				return m, func() tea.Msg { return ListOrphansIntent{} }
			}

		case key.Matches(msg, m.keys.Export):
			if m.operation == "" {
				return m, func() tea.Msg { return ExportReportIntent{} }
			}

		case key.Matches(msg, m.keys.Sweep):
			// Sweep is intentionally gated by the orphan overlay only.

		case key.Matches(msg, m.keys.Preflight):
			if m.exec != nil {
				return m, m.exec(PreflightIntent{})
			}

		case key.Matches(msg, m.keys.Mode):
			if len(m.pickerModes) > 0 {
				m.pickerActive = true
				// Pre-position cursor at the currently active mode.
				for i, pm := range m.pickerModes {
					if pm.Name == m.mode {
						m.pickerCursor = i
						break
					}
				}
			}

		case key.Matches(msg, m.keys.Vars):
			m.editorActive = true
			m.editorCursor = 0
			m.editorFocused = false
			getenv := m.getenv
			if getenv == nil {
				getenv = func(string) string { return "" }
			}
			m.editorFields = buildEditorFields(m.envVars, getenv)

		case key.Matches(msg, m.keys.TriageRefresh):
			if m.section == sectionTriage && m.operation == "" {
				return m, func() tea.Msg { return RefreshTriageIntent{} }
			}

		case key.Matches(msg, m.keys.SyncKnownIssues):
			if m.section == sectionTriage && m.operation == "" {
				return m, func() tea.Msg { return SyncKnownIssuesIntent{} }
			}

		case key.Matches(msg, m.keys.FileIssue):
			if m.section == sectionTriage && m.triageDetailActive && m.operation == "" {
				if f, ok := m.selectedTriageFailure(); ok && engine.EligibleForIssueFiling(f) {
					return m, func() tea.Msg { return PreviewFileIssueIntent{Fingerprint: f.Fingerprint} }
				}
			}

		case key.Matches(msg, m.keys.CopyLog):
			if m.copyFn == nil {
				return m, nil
			}
			if m.focus == focusGroups {
				m.status = "select a test first"
				return m, nil
			}
			name := m.selectedTestName()
			k := resultKey{Name: name, Sub: ""}
			idx, ok := m.resultIndex[k]
			if !ok || len(m.results[idx].Output) == 0 {
				m.status = "nothing to copy"
				return m, nil
			}
			content := strings.Join(m.results[idx].Output, "\n")
			if err := m.copyFn(content); err != nil {
				m.status = "copy failed: " + err.Error()
				return m, nil
			}
			m.status = "copied log"

		case key.Matches(msg, m.keys.CopyCmd):
			if m.copyFn == nil || m.cmdFor == nil {
				return m, nil
			}
			if m.focus == focusGroups {
				m.status = "select a test first"
				return m, nil
			}
			name := m.selectedTestName()
			if name == "" {
				m.status = "nothing to copy"
				return m, nil
			}
			cmd := m.cmdFor(name)
			if err := m.copyFn(cmd); err != nil {
				m.status = "copy failed: " + err.Error()
				return m, nil
			}
			m.status = "copied cmd"
		}

	case PreflightMsg:
		m.clearErrorPresentation()
		m.mode = msg.Report.Mode
		m.checks = msg.Report.Checks
		if msg.EnvVars != nil {
			m.envVars = msg.EnvVars
		}

	case GroupsMsg:
		m.clearErrorPresentation()
		m.groups = cloneGroups(msg.Groups)
		m.discoveryGroups = cloneGroups(msg.Groups)
		if m.failuresFirst {
			m.applyFailuresFirstOrdering()
		}

	case TestUpdateMsg:
		m.clearErrorPresentation()
		if m.resultIndex == nil {
			m.resultIndex = make(map[resultKey]int)
		}
		k := resultKey{Name: msg.Result.Name, Sub: msg.Result.Sub}
		if idx, ok := m.resultIndex[k]; ok {
			m.results[idx] = msg.Result
		} else {
			m.resultIndex[k] = len(m.results)
			m.results = append(m.results, msg.Result)
		}
		// Record that progress happened; the next tick resets the stall clock.
		m.sawActivity = true
		if m.failuresFirst {
			m.applyFailuresFirstOrdering()
		}

	case TriageLoadedMsg:
		m.triageFailures = append([]engine.PersistFailure(nil), msg.Failures...)
		m.triageCacheAvailable = msg.CacheAvailable
		triageRefreshComplete := m.operation == "triage"
		m.clearOperation("triage")
		if triageRefreshComplete {
			m.status = "triage refreshed"
		}
		m.clampTriageCursor()
		if len(m.triageFailures) == 0 {
			m.triageDetailActive = false
		}

	case OperationStartedMsg:
		m.clearErrorPresentation()
		m.operation = msg.Name
		if msg.Name != "" {
			m.status = msg.Name + uiEllipsis(m.ascii)
		}

	case OperationErrMsg:
		if msg.Op == m.operation {
			m.operation = ""
		}
		if msg.Op == "run" && msg.Err != nil {
			m.running = false
		}
		m.setErrorPresentation(msg.Err, msg.Op)
		m.errIsOperationFailure = msg.Err != nil
		if msg.Err != nil {
			if msg.Op != "" {
				m.status = msg.Op + ": " + msg.Err.Error()
			} else {
				m.status = msg.Err.Error()
			}
		}

	case OperationRejectedMsg:
		m.setErrorPresentation(msg.Err, m.operation)
		if msg.Err != nil {
			if msg.Op != "" {
				m.status = msg.Op + ": " + msg.Err.Error()
			} else {
				m.status = msg.Err.Error()
			}
		}

	case RunDoneMsg:
		m.running = false
		m.clearOperation("run")
		m.clearErrorPresentation()
		switch {
		case msg.Result.BuildFailed:
			m.setErrorPresentation(errors.New("build failed before tests ran"), "run")
			m.errIsOperationFailure = true
			m.status = "run: " + m.err.Error()
		case msg.Result.PreRunFailed:
			m.setErrorPresentation(errors.New("pre-run failed before tests ran"), "run")
			m.errIsOperationFailure = true
			m.status = "run: " + m.err.Error()
		default:
			m.status = "run complete"
		}
		m.runLabel = ""
		m.stalledFor = 0

	case RunStartedMsg:
		m.clearErrorPresentation()
		m.running = true
		m.status = "running" + uiEllipsis(m.ascii)
		m.runLabel = msg.Label
		m.spinnerFrame = 0
		// Reset the liveness clocks; the first tick anchors them.
		m.runStartedAt = time.Time{}
		m.runElapsed = 0
		m.lastActivityAt = time.Time{}
		m.sawActivity = false
		m.stalledFor = 0
		return m, spinnerTickCmd()

	case introTickMsg:
		if m.introActive {
			m.introFrame++
			if m.introFrame >= introTotalFrames {
				m.introActive = false
				// Clear the splash so the dashboard paints from a clean screen
				// in every terminal (no leftover wordmark rows).
				return m, tea.ClearScreen
			}
			return m, introTickCmd()
		}
		return m, nil

	case spinnerTickMsg:
		if m.running {
			m.spinnerFrame++
			if m.runStartedAt.IsZero() {
				// First tick of the run: anchor both clocks to it.
				m.runStartedAt = msg.t
				m.lastActivityAt = msg.t
			}
			m.runElapsed = msg.t.Sub(m.runStartedAt)
			if m.sawActivity {
				m.lastActivityAt = msg.t
				m.sawActivity = false
			}
			m.stalledFor = msg.t.Sub(m.lastActivityAt)
			return m, spinnerTickCmd()
		}
		return m, nil

	case ResumePromptMsg:
		m.resumePrompt = true
		m.resumeWhen = msg.When
		m.resumeFailed = msg.Failed
		m.resumeNotRun = msg.NotRun

	case KnownIssueSyncDoneMsg:
		m.clearOperation("known-issues-sync")
		m.triageFailures = append([]engine.PersistFailure(nil), msg.Failures...)
		m.triageCacheAvailable = msg.CacheAvailable
		m.clampTriageCursor()
		lines := []string{fmt.Sprintf("count: %d", msg.Count)}
		if msg.CachePath != "" {
			lines = append(lines, "cache: "+msg.CachePath)
		}
		lines = append(lines, fmt.Sprintf("failures: %d", len(msg.Failures)))
		m.setResultOverlay("Known issues synced", lines...)
		m.status = "known issues synced"

	case ReportExportDoneMsg:
		m.clearOperation("export")
		var lines []string
		if msg.MarkdownPath != "" {
			lines = append(lines, "markdown: "+msg.MarkdownPath)
		}
		if msg.HTMLPath != "" {
			lines = append(lines, "html: "+msg.HTMLPath)
		}
		if len(lines) == 0 {
			lines = append(lines, "no report paths returned")
		}
		m.setResultOverlay("Report exported", lines...)
		m.status = "report exported"

	case OrphansListedMsg:
		m.clearOperation("orphans")
		m.orphanOwner = msg.Owner
		m.orphans = cloneResources(msg.Resources)
		m.orphanOverlayActive = true
		m.clearResultOverlay()
		m.status = fmt.Sprintf("%d orphan resources", len(m.orphans))

	case SweepDoneMsg:
		m.clearOperation("sweep")
		m.closeSweepConfirm()
		m.orphans = cloneResources(msg.Remaining)
		if msg.SnapshotChanged {
			m.orphanOverlayActive = true
			m.clearResultOverlay()
			m.status = "orphan list refreshed"
		} else if msg.Failed || msg.ResidualUnknown {
			m.orphanOverlayActive = false
			lines := []string{"snapshot changed: false"}
			if msg.ResidualUnknown {
				lines = append(lines, "residual state: unknown")
			} else {
				lines = append(lines, fmt.Sprintf("remaining: %d", len(msg.Remaining)))
			}
			m.setResultOverlayWithTone(panelDanger, "Sweep failed", lines...)
			m.status = "sweep failed"
		} else {
			m.orphanOverlayActive = false
			lines := []string{
				fmt.Sprintf("remaining: %d", len(msg.Remaining)),
				"snapshot changed: false",
			}
			m.setResultOverlay("Sweep complete", lines...)
			m.status = "sweep complete"
		}

	case IssuePreviewMsg:
		m.clearOperation("issue-preview")
		m.openFileIssueConfirm(msg)
		m.status = "issue preview ready"

	case IssueFiledMsg:
		m.clearOperation("issue-file")
		m.closeFileIssueConfirm()
		m.triageFailures = append([]engine.PersistFailure(nil), msg.Failures...)
		m.clampTriageCursor()
		title, status := issueFiledWording(msg.Action)
		lines := []string{"fingerprint: " + msg.Fingerprint}
		if msg.IssueNumber > 0 {
			lines = append(lines, fmt.Sprintf("issue: #%d", msg.IssueNumber))
		}
		lines = append(lines,
			"action: "+msg.Action,
			fmt.Sprintf("failures: %d", len(msg.Failures)),
		)
		m.setResultOverlay(title, lines...)
		m.status = status

	case RetryGroupIntent, RetryTestIntent, RetryAllIntent, ResumeIntent,
		SwitchModeIntent, SetEnvVarIntent, PreflightIntent,
		RefreshTriageIntent, SyncKnownIssuesIntent, PreviewFileIssueIntent,
		ExportReportIntent, ListOrphansIntent, ConfirmSweepIntent,
		ConfirmFileIssueIntent:
		if m.exec != nil {
			return m, m.exec(msg)
		}
		// no executor injected (e.g. headless tests): intents are no-ops.

	case RateMsg:
		m.rate = msg

	case ErrMsg:
		m.setErrorPresentation(msg.Err, "")
		if msg.Err != nil {
			m.status = msg.Err.Error()
		}
	}

	return m, nil
}
