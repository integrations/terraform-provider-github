package tui

import (
	"strings"
	"testing"

	"github.com/github/terraform-provider-tester/engine"
)

func TestOrphanOverlayGolden(t *testing.T) {
	m := fixedModel()
	m.orphanOverlayActive = true
	m.orphanOwner = "acme-test"
	m.orphans = sampleResources()

	out := m.View()
	assertNoANSI(t, out)
	for _, want := range []string{"Orphan resources", "owner", "acme-test", "KIND", "NAME", "URL", "tf-acc-repo-a"} {
		if !strings.Contains(out, want) {
			t.Fatalf("orphan overlay missing %q:\n%s", want, out)
		}
	}
	compareOrUpdate(t, "orphan_overlay", out)
}

func TestOrphanOverlayUsesSharedPanel(t *testing.T) {
	m := fixedModel()
	m.orphanOwner = "acme-test"
	m.orphans = sampleResources()

	got := stripANSI(m.renderOrphanOverlay(100))
	for _, want := range []string{"+", "|", "Orphan resources", "owner", "count", "s sweep", "esc close"} {
		if !strings.Contains(got, want) {
			t.Fatalf("orphan overlay missing shared panel marker %q:\n%s", want, got)
		}
	}
	assertRenderedWidth(t, got, 100)
}

func TestSweepConfirmGolden(t *testing.T) {
	m := fixedModel()
	m.orphanOverlayActive = true
	m.orphanOwner = "acme-test"
	m.orphans = sampleResources()
	next, _ := m.Update(keyRunes("s"))
	m = next.(Model)

	out := m.View()
	assertNoANSI(t, out)
	if !strings.Contains(out, "SWEEP acme-test") {
		t.Fatalf("sweep confirmation missing exact phrase:\n%s", out)
	}
	compareOrUpdate(t, "sweep_confirm", out)
}

func TestSweepConfirmationUsesDangerPanelAndExactPhrase(t *testing.T) {
	m := fixedModel()
	m.orphanOwner = "my-test-org"
	m.orphans = sampleResources()
	m.sweepConfirmActive = true

	got := stripANSI(m.renderSweepConfirm(100))
	for _, want := range []string{"Destructive action", "Confirm sweep", "SWEEP my-test-org", "Targets", "enter confirm", "esc close"} {
		if !strings.Contains(got, want) {
			t.Fatalf("sweep confirm missing %q:\n%s", want, got)
		}
	}
}

func TestFileIssueConfirmGolden(t *testing.T) {
	m := fixedModel()
	m.section = sectionTriage
	next, _ := m.Update(IssuePreviewMsg{
		Fingerprint:      "sha256:" + strings.Repeat("c", 64),
		ShortFingerprint: "cccccccccccccccc",
		IssuesRepo:       "integrations/terraform-provider-github",
		Title:            "redacted failure title",
		Labels:           []string{"bug", "acceptance-test"},
		Classification:   engine.ClassificationReal,
	})
	m = next.(Model)

	out := m.View()
	assertNoANSI(t, out)
	for _, want := range []string{"File issue", "integrations/terraform-provider-github", "redacted failure title", "bug, acceptance-test", "FILE cccccccccccccccc"} {
		if !strings.Contains(out, want) {
			t.Fatalf("file issue overlay missing %q:\n%s", want, out)
		}
	}
	if strings.Contains(strings.ToLower(out), "issue body") {
		t.Fatalf("file issue overlay must not render an issue body:\n%s", out)
	}
	compareOrUpdate(t, "file_issue_confirm", out)
}

func TestFileIssueConfirmationUsesDangerPanelAndExactPhrase(t *testing.T) {
	m := fixedModel()
	m.section = sectionTriage
	next, _ := m.Update(IssuePreviewMsg{
		Fingerprint:      "sha256:" + strings.Repeat("c", 64),
		ShortFingerprint: "cccccccccccccccc",
		IssuesRepo:       "integrations/terraform-provider-github",
		Title:            "redacted failure title",
		Labels:           []string{"bug", "acceptance-test"},
		Classification:   engine.ClassificationReal,
	})
	m = next.(Model)

	got := stripANSI(m.renderFileIssueConfirm(100))
	for _, want := range []string{"Destructive action", "File issue", "FILE cccccccccccccccc", "repo", "fingerprint", "enter confirm", "esc close"} {
		if !strings.Contains(got, want) {
			t.Fatalf("file issue confirm missing %q:\n%s", want, got)
		}
	}
}

func TestExportResultOverlay(t *testing.T) {
	m := fixedModel()
	next, _ := m.Update(ReportExportDoneMsg{MarkdownPath: ".pulsar-report.md", HTMLPath: ".pulsar-report.html"})
	m = next.(Model)

	out := m.View()
	assertNoANSI(t, out)
	for _, want := range []string{"Report exported", ".pulsar-report.md", ".pulsar-report.html", "esc close"} {
		if !strings.Contains(out, want) {
			t.Fatalf("export result overlay missing %q:\n%s", want, out)
		}
	}
}

func TestResultOverlayUsesSharedPanel(t *testing.T) {
	m := fixedModel()
	m.resultTitle = "Report exported"
	m.resultLines = []string{"markdown: report.md", "html: report.html"}

	got := stripANSI(m.renderResultOverlay(80))
	for _, want := range []string{"+", "|", "Report exported", "markdown: report.md", "html: report.html", "esc close"} {
		if !strings.Contains(got, want) {
			t.Fatalf("result overlay missing %q:\n%s", want, got)
		}
	}
	assertRenderedWidth(t, got, 80)
}

func TestCleanupOverlaysFitRepresentativeWidths(t *testing.T) {
	for _, width := range []int{48, 80, 120} {
		m := fixedModel()
		m.orphanOwner = "acme-test"
		m.orphans = sampleResources()
		assertRenderedWidth(t, m.renderOrphanOverlay(width), width)

		m.sweepConfirmActive = true
		assertRenderedWidth(t, m.renderSweepConfirm(width), width)

		next, _ := m.Update(IssuePreviewMsg{
			Fingerprint:      "sha256:" + strings.Repeat("c", 64),
			ShortFingerprint: "cccccccccccccccc",
			IssuesRepo:       "integrations/terraform-provider-github",
			Title:            "redacted failure title",
			Labels:           []string{"bug", "acceptance-test"},
			Classification:   engine.ClassificationReal,
		})
		m = next.(Model)
		assertRenderedWidth(t, m.renderFileIssueConfirm(width), width)

		m.setResultOverlay("Report exported", "markdown: a-very-long-report-path-that-must-fit.md")
		assertRenderedWidth(t, m.renderResultOverlay(width), width)
	}
}
