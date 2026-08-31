package tui

import (
	"strings"
	"testing"

	"github.com/charmbracelet/lipgloss"

	"github.com/github/terraform-provider-tester/engine"
	"github.com/github/terraform-provider-tester/provider"
)

// headerLine returns the first rendered row (the header bar).
func headerLine(t *testing.T, m Model) string {
	t.Helper()
	return strings.SplitN(m.View(), "\n", 2)[0]
}

func TestHeaderShowsConsoleIdentityAndVersion(t *testing.T) {
	m := fixedModel().WithVersion("9.9.9")
	head := headerLine(t, m)
	for _, want := range []string{"GH/TF", "Provider Acceptance"} {
		if !strings.Contains(head, want) {
			t.Errorf("header missing console identity %q: %q", want, head)
		}
	}
	if strings.Contains(head, "TESTER") {
		t.Errorf("header still shows retired TESTER mark: %q", head)
	}
	if !strings.Contains(head, "9.9.9") {
		t.Errorf("header missing version: %q", head)
	}
	if !strings.Contains(head, "provider: github") {
		t.Errorf("header missing provider: %q", head)
	}
	if !strings.Contains(head, "mode: organization") {
		t.Errorf("header missing mode: %q", head)
	}
}

func TestHeaderShowsOwnerChipOnlyOnWideLayouts(t *testing.T) {
	m := fixedModel() // owner = "my-test-org"
	if head := headerLine(t, m); strings.Contains(head, "owner") {
		t.Errorf("standard header should omit owner context: %q", head)
	}
	m.width = 120
	if head := headerLine(t, m); !strings.Contains(head, "[owner: my-test-org]") {
		t.Errorf("wide header should include owner chip: %q", head)
	}
}

func TestHeaderWithoutVersionHasNoTrailingVersion(t *testing.T) {
	m := fixedModel().WithVersion("")
	head := headerLine(t, m)
	if !strings.Contains(head, "GH/TF Provider Acceptance [provider: github]") {
		t.Errorf("empty version should collapse cleanly: %q", head)
	}
}

func TestActiveTabHasUnderlineRule(t *testing.T) {
	m := fixedModel() // section = Preflight
	lines := strings.Split(m.View(), "\n")
	if len(lines) < 3 {
		t.Fatalf("expected at least 3 rows (header, tabs, underline), got %d", len(lines))
	}
	if !strings.Contains(lines[2], "-") {
		t.Errorf("expected an underline rule under the active tab, got %q", lines[2])
	}
}

func TestConsoleSelectionUsesLineWithoutPurpleFill(t *testing.T) {
	if _, ok := Styles.TabActive.GetBackground().(lipgloss.NoColor); !ok {
		t.Fatalf("active tab background = %#v, want no color", Styles.TabActive.GetBackground())
	}
	if Styles.SelectedRow.GetBackground() == Selected {
		t.Fatal("selected row still uses the purple selection color as a full background")
	}
}

func runConsoleModel() Model {
	m := fixedModel()
	m.section = sectionRun
	m.groups = []engine.Group{
		{Name: "repositories", Tests: []string{"TestAccGithubRepoA", "TestAccGithubRepoB"}},
		{Name: "teams", Tests: []string{"TestAccGithubTeamA"}},
	}
	m.results = []engine.TestResult{
		{Name: "TestAccGithubRepoA", Status: provider.StatusPass},
		{Name: "TestAccGithubRepoB", Status: provider.StatusFail},
	}
	return m
}

func TestRunConsoleShowsTargetTalliesAndDocs(t *testing.T) {
	out := runConsoleModel().View()
	for _, want := range []string{
		TargetRepo,
		"mode: organization",
		"my-test-org", // owner surfaces here, not in the header
		"passed",
		"failed",
		"not run",
		DocsPath,
	} {
		if !strings.Contains(out, want) {
			t.Errorf("run console missing %q\n---\n%s", want, out)
		}
	}
}

func TestRunConsoleCountsAreCorrect(t *testing.T) {
	out := runConsoleModel().View()
	// 3 total tests, 1 passed, 1 failed, 0 running, 1 not run.
	for _, want := range []string{"3 total", "1 passed", "1 failed", "1 not run"} {
		if !strings.Contains(out, want) {
			t.Errorf("run console tally missing %q\n---\n%s", want, out)
		}
	}
}

func TestRunConsoleEmptyStateGuidesUser(t *testing.T) {
	m := fixedModel()
	m.section = sectionRun
	m.groups = nil
	out := m.View()
	if !strings.Contains(out, "No tests discovered") {
		t.Errorf("empty Run console should guide the user, got:\n%s", out)
	}
}

func TestBannerShowsBrandVersionTagline(t *testing.T) {
	out := renderBanner("1.2.3", true)
	for _, want := range []string{"GH/TF", "Provider Acceptance", "1.2.3", "legible", DocsPath} {
		if !strings.Contains(out, want) {
			t.Errorf("banner missing %q\n---\n%s", want, out)
		}
	}
	if strings.Contains(out, "TESTER") {
		t.Errorf("banner still shows retired TESTER mark\n%s", out)
	}
}

func TestBannerWithoutVersionOmitsV(t *testing.T) {
	out := renderBanner("", true)
	if !strings.Contains(out, "GH/TF Provider Acceptance") {
		t.Fatalf("banner should still show console identity\n%s", out)
	}
	if strings.Contains(out, "Provider Acceptance v") {
		t.Errorf("banner should omit version suffix when version is empty\n%s", out)
	}
}

func TestBannerAsciiVariantHasNoBoxDrawing(t *testing.T) {
	ascii := renderBanner("1.2.3", true)
	// The ascii icon must avoid half-block glyphs so dumb / NO_COLOR terminals
	// render cleanly (the ascii border itself may use box-drawing).
	for _, g := range []string{"▀", "▄", "█"} {
		if strings.Contains(ascii, g) {
			t.Errorf("ascii banner should not contain block glyph %q\n%s", g, ascii)
		}
	}
	tty := renderBanner("1.2.3", false)
	if !strings.Contains(tty, "GH") || !strings.Contains(tty, "TF") {
		t.Errorf("tty banner should render the GH/TF mark\n%s", tty)
	}
}

func TestBannerUsesGHTFMarkNotLegacyMarks(t *testing.T) {
	tty := renderBanner("1.2.3", false)
	if strings.Contains(tty, "▄██▄") {
		t.Errorf("banner still renders the legacy ear glyphs\n%s", tty)
	}
	if strings.Contains(tty, "█████") {
		t.Errorf("banner still renders the TESTER beacon core\n%s", tty)
	}
}

func TestBannerSignifiesGitHubOwnershipAndLicense(t *testing.T) {
	for _, ascii := range []bool{true, false} {
		out := renderBanner("1.2.3", ascii)
		for _, want := range []string{Vendor, License} {
			if !strings.Contains(out, want) {
				t.Errorf("banner (ascii=%v) missing ownership/license token %q\n%s", ascii, want, out)
			}
		}
	}
}

func TestPreflightTabShowsBanner(t *testing.T) {
	m := fixedModel()
	m.section = sectionPreflight
	out := m.View()
	for _, want := range []string{"GH/TF", "Provider Acceptance", Tagline} {
		if !strings.Contains(out, want) {
			t.Errorf("preflight view missing banner content %q\n---\n%s", want, out)
		}
	}
}
