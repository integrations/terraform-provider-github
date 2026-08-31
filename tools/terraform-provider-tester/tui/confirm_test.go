package tui

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/github/terraform-provider-tester/engine"
	"github.com/github/terraform-provider-tester/provider"
)

func TestConfirmationPhrasesAreExact(t *testing.T) {
	tests := []struct {
		input, expected string
		want            bool
	}{
		{"SWEEP acme-test", "SWEEP acme-test", true},
		{"sweep acme-test", "SWEEP acme-test", false},
		{"SWEEP acme-test ", "SWEEP acme-test", false},
		{"FILE abcdef0123456789", "FILE abcdef0123456789", true},
		{"FILE abcdef", "FILE abcdef0123456789", false},
	}
	for _, tc := range tests {
		if got := phraseMatches(tc.input, tc.expected); got != tc.want {
			t.Fatalf("phraseMatches(%q, %q) = %v, want %v", tc.input, tc.expected, got, tc.want)
		}
	}
}

func TestConfirmationPhraseHelpersUseSecurityBoundaryStrings(t *testing.T) {
	if got, want := sweepPhrase("acme-test"), "SWEEP acme-test"; got != want {
		t.Fatalf("sweepPhrase = %q, want %q", got, want)
	}
	if got, want := fileIssuePhrase("abcdef0123456789"), "FILE abcdef0123456789"; got != want {
		t.Fatalf("fileIssuePhrase = %q, want %q", got, want)
	}
}

func TestOrphanAndExportKeysEmitIntentsWhenIdle(t *testing.T) {
	m := newTestModel()
	if _, cmd := pressKey(m, 'o'); cmd == nil {
		t.Fatal("o should emit ListOrphansIntent")
	} else if _, ok := cmd().(ListOrphansIntent); !ok {
		t.Fatalf("o emitted %T, want ListOrphansIntent", cmd())
	}
	if _, cmd := pressKey(m, 'e'); cmd == nil {
		t.Fatal("e should emit ExportReportIntent")
	} else if _, ok := cmd().(ExportReportIntent); !ok {
		t.Fatalf("e emitted %T, want ExportReportIntent", cmd())
	}
}

func TestSweepConfirmationRequiresExactPhrase(t *testing.T) {
	resources := sampleResources()
	m := newTestModel()
	next, _ := m.Update(OrphansListedMsg{Owner: "githubq", Resources: resources})
	m = next.(Model)

	m, cmd := pressKey(m, 's')
	if cmd != nil {
		t.Fatal("s should open sweep confirmation without emitting an intent")
	}
	if !m.sweepConfirmActive || !m.sweepInput.Focused() {
		t.Fatal("s should focus the sweep confirmation input")
	}

	m, _ = typeIntoFocusedInput(t, m, "sweep githubq")
	next, cmd = m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	m = next.(Model)
	if cmd != nil {
		t.Fatal("wrong sweep phrase emitted confirmation intent")
	}
	if !m.sweepConfirmActive {
		t.Fatal("wrong sweep phrase closed confirmation overlay")
	}

	m.sweepInput.SetValue("")
	phrase := "SWEEP githubq"
	for i, r := range phrase {
		next, cmd = m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{r}})
		m = next.(Model)
		if !m.sweepConfirmActive {
			t.Fatalf("typing rune %q closed sweep confirmation", r)
		}
		if got, want := m.sweepInput.Value(), string([]rune(phrase)[:i+1]); got != want {
			t.Fatalf("after typing rune %q input = %q, want %q", r, got, want)
		}
	}

	next, cmd = m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	got := next.(Model)
	if cmd == nil {
		t.Fatal("exact sweep phrase should emit ConfirmSweepIntent")
	}
	intent, ok := cmd().(ConfirmSweepIntent)
	if !ok {
		t.Fatalf("enter emitted %T, want ConfirmSweepIntent", cmd())
	}
	if intent.Owner != "githubq" || intent.Phrase != phrase {
		t.Fatalf("intent owner/phrase = %q/%q", intent.Owner, intent.Phrase)
	}
	if len(intent.Resources) != len(resources) {
		t.Fatalf("intent resources = %d, want %d", len(intent.Resources), len(resources))
	}
	if got.sweepConfirmActive || got.sweepInput.Value() != "" {
		t.Fatal("successful sweep confirmation should close and clear input")
	}
}

func TestFileIssueConfirmationRequiresExactPhrase(t *testing.T) {
	m := newTestModel()
	m.section = sectionTriage
	m.triageFailures = sampleFailures()
	m.triageCursor = 2
	m.triageDetailActive = true // i files only from the Triage detail view

	_, cmd := pressKey(m, 'i')
	if cmd == nil {
		t.Fatal("i should emit PreviewFileIssueIntent for eligible failure")
	}
	if _, ok := cmd().(PreviewFileIssueIntent); !ok {
		t.Fatalf("i emitted %T, want PreviewFileIssueIntent", cmd())
	}

	next, _ := m.Update(IssuePreviewMsg{
		Fingerprint:      "sha256:" + strings.Repeat("c", 64),
		ShortFingerprint: "cccccccccccccccc",
		IssuesRepo:       "integrations/terraform-provider-github",
		Title:            "redacted failure title",
		Labels:           []string{"bug", "acceptance-test"},
		Classification:   engine.ClassificationReal,
	})
	m = next.(Model)
	if !m.fileIssueActive || !m.fileIssueInput.Focused() {
		t.Fatal("IssuePreviewMsg should open and focus the file issue confirmation")
	}

	m, _ = typeIntoFocusedInput(t, m, "FILE cccccc")
	next, cmd = m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	m = next.(Model)
	if cmd != nil {
		t.Fatal("wrong file issue phrase emitted confirmation intent")
	}
	if !m.fileIssueActive {
		t.Fatal("wrong file issue phrase closed confirmation overlay")
	}

	m.fileIssueInput.SetValue("")
	m, _ = typeIntoFocusedInput(t, m, "FILE cccccccccccccccc")
	next, cmd = m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	got := next.(Model)
	if cmd == nil {
		t.Fatal("exact file issue phrase should emit ConfirmFileIssueIntent")
	}
	intent, ok := cmd().(ConfirmFileIssueIntent)
	if !ok {
		t.Fatalf("enter emitted %T, want ConfirmFileIssueIntent", cmd())
	}
	if intent.Fingerprint != "sha256:"+strings.Repeat("c", 64) || intent.Phrase != "FILE cccccccccccccccc" {
		t.Fatalf("intent fingerprint/phrase = %q/%q", intent.Fingerprint, intent.Phrase)
	}
	if got.fileIssueActive || got.fileIssueInput.Value() != "" {
		t.Fatal("successful file issue confirmation should close and clear input")
	}
}

// TestIssueFiledMsgIsActionAware verifies the completion overlay/status wording
// reflects issueService.FileOne's actual outcome (filed/dedup/known) instead
// of always claiming a new issue was created, while every action still
// updates fingerprint/action/failure state, clamps the cursor, releases the
// operation, and closes the confirmation overlay.
func TestIssueFiledMsgIsActionAware(t *testing.T) {
	tests := []struct {
		name        string
		action      string
		issueNumber int
		wantTitle   string
		wantStatus  string
		wantIssue   bool
	}{
		{"filed", "filed", 505, "Issue filed", "issue filed", true},
		{"dedup", "dedup", 42, "Issue deduplicated", "issue deduplicated", true},
		{"known", "known", 0, "Issue already known", "issue already known", false},
		{"unknown action", "suppressed", 0, "", "", false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m := newTestModel()
			m.section = sectionTriage
			m.operation = "issue-file"
			m.fileIssueActive = true
			m.fileIssueInput.Focus()
			m.triageFailures = sampleFailures()
			m.triageCursor = 9 // out of range; must clamp after the update

			refreshed := sampleFailures()[:2]
			next, _ := m.Update(IssueFiledMsg{
				Fingerprint: "sha256:" + strings.Repeat("c", 64),
				IssueNumber: tc.issueNumber,
				Action:      tc.action,
				Failures:    refreshed,
			})
			got := next.(Model)

			if got.operation != "" {
				t.Fatalf("operation = %q, want empty after completion", got.operation)
			}
			if got.fileIssueActive || got.fileIssueInput.Focused() {
				t.Fatal("IssueFiledMsg should close the file issue confirmation overlay")
			}
			if len(got.triageFailures) != len(refreshed) {
				t.Fatalf("triageFailures = %d, want %d refreshed failures", len(got.triageFailures), len(refreshed))
			}
			if got.triageCursor != len(refreshed)-1 {
				t.Fatalf("triageCursor = %d, want clamped to %d", got.triageCursor, len(refreshed)-1)
			}
			if !got.resultOverlayActive {
				t.Fatal("IssueFiledMsg should open the result overlay")
			}
			if tc.wantTitle != "" && got.resultTitle != tc.wantTitle {
				t.Fatalf("resultTitle = %q, want %q", got.resultTitle, tc.wantTitle)
			}
			if tc.wantStatus != "" && got.status != tc.wantStatus {
				t.Fatalf("status = %q, want %q", got.status, tc.wantStatus)
			}
			if tc.action != "filed" && got.resultTitle == "Issue filed" {
				t.Fatalf("action %q must not render the filed title", tc.action)
			}
			if tc.action != "filed" && got.status == "issue filed" {
				t.Fatalf("action %q must not render the filed status", tc.action)
			}
			joined := strings.Join(got.resultLines, "\n")
			if !strings.Contains(joined, "fingerprint: sha256:"+strings.Repeat("c", 64)) {
				t.Fatalf("resultLines missing fingerprint line: %v", got.resultLines)
			}
			if !strings.Contains(joined, "action: "+tc.action) {
				t.Fatalf("resultLines missing action line: %v", got.resultLines)
			}
			if !strings.Contains(joined, fmt.Sprintf("failures: %d", len(refreshed))) {
				t.Fatalf("resultLines missing failure count line: %v", got.resultLines)
			}
			if tc.wantIssue {
				if !strings.Contains(joined, fmt.Sprintf("issue: #%d", tc.issueNumber)) {
					t.Fatalf("resultLines missing issue number line for action %q: %v", tc.action, got.resultLines)
				}
			} else {
				if strings.Contains(joined, "issue: #0") {
					t.Fatalf("resultLines must never render issue: #0: %v", got.resultLines)
				}
				if strings.Contains(joined, "issue:") {
					t.Fatalf("resultLines must omit the issue line when IssueNumber <= 0: %v", got.resultLines)
				}
			}
		})
	}
}

func TestOverlaysConsumeKeysAndCloseOnEsc(t *testing.T) {
	t.Run("orphan", func(t *testing.T) {
		m := newTestModel()
		m.section = sectionGroups
		m.focus = focusTests
		m.groupCursor = 1
		m.testCursor = 1
		next, _ := m.Update(OrphansListedMsg{Owner: "acme", Resources: sampleResources()})
		m = next.(Model)
		baseline := m
		for _, r := range []rune{'q', 'h'} {
			got, cmd := pressKey(baseline, r)
			if cmd != nil {
				t.Fatalf("%q emitted a command", r)
			}
			if !got.orphanOverlayActive {
				t.Fatalf("%q closed orphan overlay", r)
			}
			if got.section != baseline.section || got.focus != baseline.focus || got.groupCursor != baseline.groupCursor || got.testCursor != baseline.testCursor {
				t.Fatalf("%q leaked underlying navigation: got section=%d focus=%d groupCursor=%d testCursor=%d", r, got.section, got.focus, got.groupCursor, got.testCursor)
			}
		}
		got, cmd := pressKey(baseline, 's')
		if cmd != nil {
			t.Fatal("orphan overlay should consume sweep key without emitting a command")
		}
		if !got.sweepConfirmActive {
			t.Fatal("s should open sweep confirmation from a non-empty orphan overlay")
		}
		next, cmd = baseline.Update(tea.KeyMsg{Type: tea.KeyEsc})
		got = next.(Model)
		if cmd != nil {
			t.Fatal("orphan esc emitted a command")
		}
		if got.orphanOverlayActive {
			t.Fatal("esc should close orphan overlay")
		}
	})

	t.Run("sweep", func(t *testing.T) {
		m := newTestModel()
		next, _ := m.Update(OrphansListedMsg{Owner: "acme", Resources: sampleResources()})
		m = next.(Model)
		m, _ = pressKey(m, 's')
		got, _ := pressKey(m, 'j')
		if got.orphanOverlayActive != m.orphanOverlayActive {
			t.Fatal("sweep overlay leaked to orphan overlay")
		}
		next, cmd := got.Update(tea.KeyMsg{Type: tea.KeyEsc})
		got = next.(Model)
		if cmd != nil {
			t.Fatal("sweep esc emitted a command")
		}
		if got.sweepConfirmActive || got.sweepInput.Value() != "" {
			t.Fatal("esc should close and clear sweep overlay")
		}
	})

	t.Run("file issue", func(t *testing.T) {
		m := modelWithFileIssuePreview()
		got, _ := pressKey(m, 'j')
		if got.triageCursor != m.triageCursor {
			t.Fatal("file issue overlay leaked triage navigation")
		}
		next, cmd := got.Update(tea.KeyMsg{Type: tea.KeyEsc})
		got = next.(Model)
		if cmd != nil {
			t.Fatal("file issue esc emitted a command")
		}
		if got.fileIssueActive || got.fileIssueInput.Value() != "" {
			t.Fatal("esc should close and clear file issue overlay")
		}
	})

	t.Run("result", func(t *testing.T) {
		m := newTestModel()
		m.section = sectionGroups
		m.focus = focusTests
		m.groupCursor = 1
		m.testCursor = 1
		next, _ := m.Update(ReportExportDoneMsg{MarkdownPath: ".pulsar-report.md", HTMLPath: ".pulsar-report.html"})
		m = next.(Model)
		baseline := m
		for _, r := range []rune{'q', 'h'} {
			got, cmd := pressKey(baseline, r)
			if cmd != nil {
				t.Fatalf("%q emitted a command", r)
			}
			if !got.resultOverlayActive {
				t.Fatalf("%q closed result overlay", r)
			}
			if got.section != baseline.section || got.focus != baseline.focus || got.groupCursor != baseline.groupCursor || got.testCursor != baseline.testCursor {
				t.Fatalf("%q leaked underlying navigation: got section=%d focus=%d groupCursor=%d testCursor=%d", r, got.section, got.focus, got.groupCursor, got.testCursor)
			}
		}
		next, cmd := baseline.Update(tea.KeyMsg{Type: tea.KeyEsc})
		got := next.(Model)
		if cmd != nil {
			t.Fatal("result esc emitted a command")
		}
		if got.resultOverlayActive {
			t.Fatal("esc should close result overlay")
		}
	})
}

func TestOperationMessagesClearOnlyMatchingOperation(t *testing.T) {
	tests := []struct {
		name string
		op   string
		msg  tea.Msg
	}{
		{"run", "run", RunDoneMsg{}},
		{"triage", "triage", TriageLoadedMsg{}},
		{"known issues", "known-issues-sync", KnownIssueSyncDoneMsg{}},
		{"export", "export", ReportExportDoneMsg{}},
		{"orphans", "orphans", OrphansListedMsg{}},
		{"sweep", "sweep", SweepDoneMsg{}},
		{"issue preview", "issue-preview", IssuePreviewMsg{}},
		{"issue file", "issue-file", IssueFiledMsg{}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m := newTestModel()
			next, _ := m.Update(OperationStartedMsg{Name: tc.op})
			m = next.(Model)
			if m.operation != tc.op {
				t.Fatalf("operation after OperationStartedMsg = %q, want %q", m.operation, tc.op)
			}
			next, _ = m.Update(tc.msg)
			got := next.(Model)
			if got.operation != "" {
				t.Fatalf("operation after matching completion = %q, want empty", got.operation)
			}
		})
	}
}

func TestNonMatchingCompletionAndRejectedSecondOperationDoNotClearActiveOperation(t *testing.T) {
	m := newTestModel()
	next, _ := m.Update(OperationStartedMsg{Name: "orphans"})
	m = next.(Model)

	next, _ = m.Update(OperationErrMsg{Op: "export", Err: errors.New("operation already running")})
	got := next.(Model)
	if got.operation != "orphans" {
		t.Fatalf("rejected second operation cleared active operation: got %q", got.operation)
	}

	next, _ = got.Update(ReportExportDoneMsg{MarkdownPath: ".pulsar-report.md"})
	got = next.(Model)
	if got.operation != "orphans" {
		t.Fatalf("non-matching completion cleared active operation: got %q", got.operation)
	}

	next, _ = got.Update(OperationErrMsg{Op: "orphans", Err: errors.New("list failed")})
	got = next.(Model)
	if got.operation != "" {
		t.Fatalf("matching error left operation active: %q", got.operation)
	}
}

func TestActionIntentsBlockedDuringOperation(t *testing.T) {
	m := newTestModel()
	next, _ := m.Update(OperationStartedMsg{Name: "run"})
	m = next.(Model)
	m.section = sectionGroups
	m.groups = []engine.Group{{Name: "g", Tests: []string{"T"}}}
	for _, r := range []rune{'r', 'R', 'e', 'o'} {
		if _, cmd := pressKey(m, r); cmd != nil {
			t.Fatalf("%c emitted a command while operation was active", r)
		}
	}
	m.section = sectionTriage
	m.triageFailures = sampleFailures()
	m.triageCursor = 2
	for _, r := range []rune{'t', 'K', 'i'} {
		if _, cmd := pressKey(m, r); cmd != nil {
			t.Fatalf("%c emitted a command while operation was active", r)
		}
	}
	m.resumePrompt = true
	if next, cmd := m.Update(tea.KeyMsg{Type: tea.KeyEnter}); cmd != nil {
		t.Fatal("resume emitted a command while operation was active")
	} else if !next.(Model).resumePrompt {
		t.Fatal("blocked resume should keep prompt visible")
	}
}

func TestSweepDoneOverlayStateAndDismissBehavior(t *testing.T) {
	t.Run("snapshot changed", func(t *testing.T) {
		m := newTestModel()
		next, _ := m.Update(OrphansListedMsg{Owner: "acme", Resources: sampleResources()})
		m = next.(Model)
		m, _ = pressKey(m, 's')
		m.setResultOverlay("stale result", "should be cleared")
		next, _ = m.Update(OperationStartedMsg{Name: "sweep"})
		m = next.(Model)

		remaining := []provider.Resource{{Kind: "repository", Name: "tf-acc-left", URL: "https://github.com/acme/tf-acc-left"}}
		next, _ = m.Update(SweepDoneMsg{Remaining: remaining, SnapshotChanged: true})
		got := next.(Model)
		if got.operation != "" {
			t.Fatalf("sweep completion left operation %q", got.operation)
		}
		if got.sweepConfirmActive || got.sweepInput.Value() != "" {
			t.Fatal("snapshot change should require a fresh sweep confirmation")
		}
		if !got.orphanOverlayActive {
			t.Fatal("snapshot change should show refreshed orphan overlay")
		}
		if got.resultOverlayActive {
			t.Fatal("snapshot change should clear stale result overlay")
		}
		if len(got.orphans) != 1 || got.orphans[0].Name != "tf-acc-left" {
			t.Fatalf("remaining orphans = %+v", got.orphans)
		}

		next, cmd := got.Update(tea.KeyMsg{Type: tea.KeyEsc})
		got = next.(Model)
		if cmd != nil {
			t.Fatal("orphan overlay esc emitted a command")
		}
		if got.orphanOverlayActive || got.resultOverlayActive {
			t.Fatal("one esc should return to the main view after snapshot changes")
		}
	})

	t.Run("snapshot unchanged", func(t *testing.T) {
		m := newTestModel()
		next, _ := m.Update(OrphansListedMsg{Owner: "acme", Resources: sampleResources()})
		m = next.(Model)
		m, _ = pressKey(m, 's')
		next, _ = m.Update(OperationStartedMsg{Name: "sweep"})
		m = next.(Model)

		next, _ = m.Update(SweepDoneMsg{Remaining: nil, SnapshotChanged: false})
		got := next.(Model)
		if got.operation != "" {
			t.Fatalf("sweep completion left operation %q", got.operation)
		}
		if got.sweepConfirmActive || got.sweepInput.Value() != "" {
			t.Fatal("successful sweep should close and clear confirmation input")
		}
		if got.orphanOverlayActive {
			t.Fatal("snapshot unchanged should not leave orphan overlay open")
		}
		if !got.resultOverlayActive {
			t.Fatal("snapshot unchanged should preserve sweep result overlay")
		}
		if got.resultTone != panelSuccess {
			t.Fatalf("successful sweep result tone = %v, want success", got.resultTone)
		}

		next, cmd := got.Update(tea.KeyMsg{Type: tea.KeyEsc})
		got = next.(Model)
		if cmd != nil {
			t.Fatal("result overlay esc emitted a command")
		}
		if got.resultOverlayActive || got.orphanOverlayActive {
			t.Fatal("esc should dismiss the lone sweep result overlay")
		}
	})
}

func typeIntoFocusedInput(t *testing.T, m Model, input string) (Model, tea.Cmd) {
	t.Helper()
	next, cmd := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune(input)})
	return next.(Model), cmd
}

func keyRunes(s string) tea.KeyMsg {
	return tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune(s)}
}

func sampleResources() []provider.Resource {
	return []provider.Resource{
		{Kind: "repository", Name: "tf-acc-repo-a", URL: "https://github.com/acme/tf-acc-repo-a"},
		{Kind: "team", Name: "tf-acc-team-b", URL: "https://github.com/orgs/acme/teams/tf-acc-team-b"},
	}
}

func modelWithFileIssuePreview() Model {
	m := newTestModel()
	m.section = sectionTriage
	m.triageFailures = sampleFailures()
	m.triageCursor = 2
	next, _ := m.Update(IssuePreviewMsg{
		Fingerprint:      "sha256:" + strings.Repeat("c", 64),
		ShortFingerprint: "cccccccccccccccc",
		IssuesRepo:       "integrations/terraform-provider-github",
		Title:            "redacted failure title",
		Labels:           []string{"bug", "acceptance-test"},
		Classification:   engine.ClassificationReal,
	})
	return next.(Model)
}

func TestRejectedDuplicateOperationDoesNotClearSameNamedActiveOperation(t *testing.T) {
	m := newTestModel()
	next, _ := m.Update(OperationStartedMsg{Name: "export"})
	m = next.(Model)

	next, _ = m.Update(OperationRejectedMsg{Op: "export", Err: errors.New("another operation is in progress")})
	got := next.(Model)
	if got.operation != "export" {
		t.Fatalf("duplicate rejection cleared active operation: got %q, want export", got.operation)
	}
	if got.status != "export: another operation is in progress" {
		t.Fatalf("status = %q, want useful busy status", got.status)
	}
}

func TestFailedSweepResultNeverRendersSuccessShape(t *testing.T) {
	m := newTestModel()
	next, _ := m.Update(OperationStartedMsg{Name: "sweep"})
	m = next.(Model)
	next, _ = m.Update(SweepDoneMsg{Failed: true, ResidualUnknown: true})
	got := next.(Model)
	if got.resultTone != panelDanger {
		t.Fatalf("failed sweep result tone = %v, want danger", got.resultTone)
	}
	view := got.View()
	if strings.Contains(view, "Sweep complete") || strings.Contains(view, "remaining: 0") {
		t.Fatalf("failed residual-unknown sweep rendered success-shaped result:\n%s", view)
	}
	if !strings.Contains(view, "Sweep failed") || !strings.Contains(view, "residual state: unknown") {
		t.Fatalf("failed residual-unknown sweep missing explicit failure/unknown state:\n%s", view)
	}
}
