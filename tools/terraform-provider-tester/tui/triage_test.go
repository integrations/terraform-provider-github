package tui

import (
	"fmt"
	"strings"
	"testing"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/github/terraform-provider-tester/engine"
)

func sampleFailures() []engine.PersistFailure {
	return []engine.PersistFailure{
		{Test: "TestAccKnown", Classification: engine.ClassificationReal, Class: "api", Fingerprint: "sha256:aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", ShortFingerprint: "aaaaaaaaaaaaaaaa", Attempts: 1, KnownIssue: 42, IssueAction: "known", Reasons: []string{"known upstream issue"}, LogPath: ".pulsar-failures/github_TestAccKnown.log"},
		{Test: "TestAccFlake", Classification: engine.ClassificationFlakeConfirmed, Class: "rate-limit", Fingerprint: "sha256:bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb", ShortFingerprint: "bbbbbbbbbbbbbbbb", Attempts: 2, Retryable: true, Reasons: []string{"later retry passed"}, LogPath: ".pulsar-failures/github_TestAccFlake.log"},
		{Test: "TestAccReal", Classification: engine.ClassificationReal, Class: "assertion", Fingerprint: "sha256:cccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccc", ShortFingerprint: "cccccccccccccccc", Attempts: 1, Retryable: false, Reasons: []string{"signature is not retryable"}, LogPath: ".pulsar-failures/github_TestAccReal.log"},
		{Test: "TestAccUnstable", Classification: engine.ClassificationRealUnstable, Class: "conflict", Fingerprint: "sha256:dddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddd", ShortFingerprint: "dddddddddddddddd", Attempts: 2, Retryable: true, Reasons: []string{"failed attempts had different fingerprints"}, LogPath: ".pulsar-failures/github_TestAccUnstable.log"},
	}
}

func TestTriageLoadedRendersListAndDetail(t *testing.T) {
	m := fixedModel()
	m.section = sectionTriage
	next, _ := m.Update(TriageLoadedMsg{Failures: sampleFailures(), CacheAvailable: true})
	m = next.(Model)
	out := m.View()
	assertNoANSI(t, out)
	compareOrUpdate(t, "triage_list", out)

	next, _ = m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	m = next.(Model)
	out = m.View()
	assertNoANSI(t, out)
	compareOrUpdateRightTrimmed(t, "triage_detail", out)
}

func TestTriageCounts(t *testing.T) {
	flakes, real, unstable, known, unknown := triageCounts(sampleFailures())
	if flakes != 1 || real != 2 || unstable != 1 || known != 1 || unknown != 2 {
		t.Fatalf("counts = flakes:%d real:%d unstable:%d known:%d unknown:%d", flakes, real, unstable, known, unknown)
	}
}

func TestTriageListShowsIssueLabels(t *testing.T) {
	m := fixedModel()
	m.section = sectionTriage
	next, _ := m.Update(TriageLoadedMsg{Failures: sampleFailures(), CacheAvailable: true})
	out := next.(Model).View()
	for _, want := range []string{"known #42", "eligible"} {
		if !strings.Contains(out, want) {
			t.Fatalf("triage list missing %q:\n%s", want, out)
		}
	}
}

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

func TestTriageViewsFitRepresentativeWidths(t *testing.T) {
	m := fixedModel()
	m.triageFailures = sampleFailures()
	for _, width := range []int{48, 80, 120} {
		list := m.renderTriageList(width)
		assertRenderedWidth(t, list, width)
		for _, want := range []string{"known #42", "eligible"} {
			if !strings.Contains(stripANSI(list), want) {
				t.Fatalf("width %d triage list missing %q:\n%s", width, want, list)
			}
		}
		m.triageDetailActive = true
		assertRenderedWidth(t, m.renderTriageDetail(width), width)
		m.triageDetailActive = false
	}
}

func TestTriageDetailOnlyShowsFileIssueForEligibleFailure(t *testing.T) {
	m := fixedModel()
	m.section = sectionTriage
	next, _ := m.Update(TriageLoadedMsg{Failures: sampleFailures()})
	m = next.(Model)

	next, _ = m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	knownDetail := next.(Model).View()
	if strings.Contains(knownDetail, "i file issue") {
		t.Fatalf("known issue detail should not offer filing:\n%s", knownDetail)
	}

	m.triageDetailActive = false
	m.triageCursor = 2
	next, _ = m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	eligibleDetail := next.(Model).View()
	if !strings.Contains(eligibleDetail, "i file issue") {
		t.Fatalf("eligible detail should offer filing:\n%s", eligibleDetail)
	}
}

func TestTriageDetailShowsCanonicalSignature(t *testing.T) {
	failures := sampleFailures()
	failures[0].Canonical = "TestAccKnown api canonical signature with ghp_CANONICAL_SECRET"
	m := fixedModel()
	m.section = sectionTriage
	next, _ := m.Update(TriageLoadedMsg{Failures: failures})
	m = next.(Model)
	next, _ = m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	out := next.(Model).View()
	if !strings.Contains(out, "canonical") || !strings.Contains(out, "TestAccKnown api canonical signature") {
		t.Fatalf("triage detail missing canonical signature:\n%s", out)
	}
}

func TestTriageSectionEmptyDetailFallsBackToList(t *testing.T) {
	m := fixedModel()
	m.section = sectionTriage
	m.triageDetailActive = true

	got := stripANSI(m.View())
	for _, want := range []string{
		"REAL 0", "UNSTABLE 0", "FLAKES 0", "KNOWN 0", "ELIGIBLE 0",
		"No classified failures", "t refresh triage", "K sync known issues",
	} {
		if !strings.Contains(got, want) {
			t.Fatalf("empty triage detail should render the list missing %q:\n%s", want, got)
		}
	}
	if strings.Contains(got, "Failure dossier") {
		t.Fatalf("empty triage detail should not render a dossier:\n%s", got)
	}
}

func TestTriageDetailTinyWidthsFitAndUseBoundNavigationKeys(t *testing.T) {
	m := fixedModel()
	m.ascii = true
	m.section = sectionTriage
	m.triageFailures = sampleFailures()[2:3]
	m.triageDetailActive = true

	tests := []struct {
		width    int
		backHint string
		backKey  tea.KeyMsg
	}{
		{width: 1, backHint: "h", backKey: tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("h")}},
		{width: 2, backHint: "h", backKey: tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("h")}},
		{width: 4, backHint: "h", backKey: tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("h")}},
		{width: 8, backHint: "esc back", backKey: tea.KeyMsg{Type: tea.KeyEsc}},
		{width: 11, backHint: "esc back", backKey: tea.KeyMsg{Type: tea.KeyEsc}},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%d", tc.width), func(t *testing.T) {
			defer func() {
				if recovered := recover(); recovered != nil {
					t.Fatalf("renderTriageSection(%d) panicked: %v", tc.width, recovered)
				}
			}()
			got := stripANSI(m.renderTriageSection(tc.width))
			assertRenderedWidth(t, got, tc.width)
			if !strings.Contains(got, tc.backHint) {
				t.Fatalf("width %d compact detail missing bound back hint %q:\n%s", tc.width, tc.backHint, got)
			}
			if !strings.Contains(got, "i") || !key.Matches(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("i")}, m.keys.FileIssue) {
				t.Fatalf("width %d compact detail displayed unbound file-issue key:\n%s", tc.width, got)
			}
			if !key.Matches(tc.backKey, m.keys.Back) {
				t.Fatalf("width %d compact detail displayed unbound back key %q", tc.width, tc.backHint)
			}
			next, _ := m.Update(tc.backKey)
			if next.(Model).triageDetailActive {
				t.Fatalf("width %d displayed back key %q did not leave detail", tc.width, tc.backHint)
			}
			for _, r := range got {
				if r > 127 {
					t.Fatalf("width %d rendered non-ASCII rune %q:\n%s", tc.width, r, got)
				}
			}
		})
	}
}
