package tui

import (
	"bytes"
	"flag"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/github/terraform-provider-tester/engine"
	"github.com/github/terraform-provider-tester/provider"
)

var updateGolden = flag.Bool("update", false, "update golden files")

// TestMain pins the lipgloss renderer to the Ascii (no-color) profile so
// goldens are machine-stable and contain no ANSI escape sequences.
func TestMain(m *testing.M) {
	flag.Parse()
	lipgloss.SetColorProfile(termenv.Ascii)
	lipgloss.SetHasDarkBackground(false)
	os.Exit(m.Run())
}

// fixedModel returns a Model with a fixed clock for deterministic View output.
// resetAt is stored on the model so View() formats it as an absolute string
// instead of calling time.Now().
func fixedModel() Model {
	resetAt := time.Date(2026, 6, 24, 21, 0, 0, 0, time.UTC)
	m := New("github", "organization", "my-test-org", true)
	m.width = 100
	m.height = 30
	m.version = "1.2.3"
	m.rate = RateMsg{Remaining: 4870, Limit: 5000, Reset: resetAt}
	return m
}

func assertNoANSI(t *testing.T, s string) {
	t.Helper()
	if bytes.ContainsRune([]byte(s), '\x1b') {
		t.Errorf("output contains ANSI escape bytes; profile pin may have failed.\nOutput:\n%s", s)
	}
}

func goldenFile(name string) string {
	return "testdata/" + name + ".golden"
}

func compareOrUpdate(t *testing.T, name, got string) {
	t.Helper()
	got = rightTrimLines(got)
	path := goldenFile(name)
	if *updateGolden || os.Getenv("UPDATE_GOLDEN") == "1" {
		if err := os.MkdirAll("testdata", 0o755); err != nil {
			t.Fatalf("mkdir testdata: %v", err)
		}
		if err := os.WriteFile(path, []byte(got), 0o644); err != nil {
			t.Fatalf("write golden %s: %v", path, err)
		}
		t.Logf("updated golden: %s", path)
		return
	}
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read golden %s: %v (run with -update to create)", path, err)
	}
	want := rightTrimLines(string(data))
	if want != got {
		t.Errorf("golden mismatch for %s\n--- want ---\n%s\n--- got ---\n%s", name, want, got)
	}
}

func compareOrUpdateRightTrimmed(t *testing.T, name, got string) {
	t.Helper()
	compareOrUpdate(t, name, rightTrimLines(got))
}

func rightTrimLines(s string) string {
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimRight(line, " 	")
	}
	return strings.Join(lines, "\n")
}

func TestViewPreflightChromeGolden(t *testing.T) {
	m := fixedModel()
	m.section = sectionPreflight
	m.checks = []provider.Check{
		{Name: "identity", Status: provider.CheckOK, Detail: docsSyntheticIdentity},
		{Name: "token-scopes", Status: provider.CheckWarn, Detail: "scopes unverified", Fix: "use classic PAT for scope verification"},
		{Name: "enterprise", Status: provider.CheckFail, Detail: "GITHUB_ENTERPRISE_SLUG not set", Fix: "set GITHUB_ENTERPRISE_SLUG=<slug>"},
	}

	out := m.View()
	assertNoANSI(t, out)
	compareOrUpdateRightTrimmed(t, "preflight_chrome", out)
}

func TestViewTabBarIncludesTriage(t *testing.T) {
	m := fixedModel()
	out := m.View()
	for _, want := range []string{"Preflight", "Groups", "Run", "Triage"} {
		if !strings.Contains(out, want) {
			t.Fatalf("expected tab bar to include %q\n%s", want, out)
		}
	}
}

func TestViewTriageEmptyGolden(t *testing.T) {
	m := fixedModel()
	m.section = sectionTriage

	out := m.View()
	assertNoANSI(t, out)
	if !strings.Contains(out, "No classified failures") {
		t.Fatalf("expected triage placeholder in output\n%s", out)
	}
	compareOrUpdate(t, "triage_empty", out)
}

func TestViewGroupsChromeGolden(t *testing.T) {
	m := fixedModel()
	m.section = sectionGroups
	m.groups = []engine.Group{
		{Name: "repositories", Tests: []string{"TestAccA", "TestAccB", "TestAccC"}},
		{Name: "teams", Tests: []string{"TestAccD", "TestAccE"}},
		{Name: "actions-secrets", Tests: []string{"TestAccF"}},
	}

	out := m.View()
	assertNoANSI(t, out)
	if !strings.Contains(out, "repositories") {
		t.Error("expected group 'repositories' in output")
	}
	if !strings.Contains(out, "teams") {
		t.Error("expected group 'teams' in output")
	}
	compareOrUpdate(t, "groups_chrome", out)
}

func TestViewGroupsMasterGolden(t *testing.T) {
	m := fixedModel()
	m.section = sectionGroups
	m.focus = focusGroups
	m.groups = []engine.Group{
		{Name: "repositories", Tests: []string{"TestAccA", "TestAccB", "TestAccC", "TestAccD"}},
		{Name: "teams", Tests: []string{"TestAccE", "TestAccF"}},
		{Name: "actions-secrets", Tests: []string{"TestAccG"}},
		{Name: "webhooks", Tests: []string{"TestAccH", "TestAccI"}},
	}
	m.results = []engine.TestResult{
		{Package: "pkg", Name: "TestAccA", Sub: "", Status: provider.StatusPass, Elapsed: 1.2},
		{Package: "pkg", Name: "TestAccB", Sub: "", Status: provider.StatusPass, Elapsed: 0.8},
		{Package: "pkg", Name: "TestAccC", Sub: "", Status: provider.StatusFail, Elapsed: 2.1},
		// TestAccD not yet run
		{Package: "pkg", Name: "TestAccE", Sub: "", Status: provider.StatusPass, Elapsed: 0.5},
		{Package: "pkg", Name: "TestAccF", Sub: "", Status: provider.StatusRunning, Elapsed: 0.0},
		// actions-secrets: TestAccG pass
		{Package: "pkg", Name: "TestAccG", Sub: "", Status: provider.StatusPass, Elapsed: 0.3},
		// webhooks: both not run
	}
	m.resultIndex = map[resultKey]int{
		{Name: "TestAccA", Sub: ""}: 0,
		{Name: "TestAccB", Sub: ""}: 1,
		{Name: "TestAccC", Sub: ""}: 2,
		{Name: "TestAccE", Sub: ""}: 3,
		{Name: "TestAccF", Sub: ""}: 4,
		{Name: "TestAccG", Sub: ""}: 5,
	}
	m.groupCursor = 0

	out := m.View()
	assertNoANSI(t, out)
	if !strings.Contains(out, "2/4") && !strings.Contains(out, "2 / 4") {
		t.Error("expected passed/total count for repositories group")
	}
	if !strings.Contains(out, "1") {
		t.Error("expected failed count annotation for repositories group")
	}
	compareOrUpdate(t, "groups_master", out)
}

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

func TestGroupsViewsFitRepresentativeWidths(t *testing.T) {
	m := fixedModel()
	m.section = sectionGroups
	m.groups = []engine.Group{
		{Name: "repositories-with-a-long-name", Tests: []string{"TestAccA", "TestAccB"}},
		{Name: "teams", Tests: []string{"TestAccC"}},
	}
	m.results = []engine.TestResult{
		{Name: "TestAccA", Status: provider.StatusPass, Elapsed: 1.2},
		{Name: "TestAccB", Status: provider.StatusFail, Elapsed: 2.3},
	}
	m.resultIndex = map[resultKey]int{{Name: "TestAccA"}: 0, {Name: "TestAccB"}: 1}
	for _, width := range []int{48, 80, 120} {
		assertRenderedWidth(t, m.renderGroups(width), width)
		m.focus = focusTests
		assertRenderedWidth(t, m.renderGroupsSection(width), width)
		m.focus = focusLog
		m.testCursor = 0
		m.logVP.Width = max(1, width-4)
		m.logVP.Height = 6
		m.logVP.SetContent("redacted log")
		assertRenderedWidth(t, m.renderGroupsSection(width), width)
		m.focus = focusGroups
	}
}

func TestContextualFooterAdvertisesArrowTabNavigation(t *testing.T) {
	m := fixedModel()
	got := stripANSI(m.renderFooter(m.width))
	for _, want := range []string{"left/right", "tabs", "? help", "q quit"} {
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

func TestExpandedHelpFooterOverridesNestedViewHints(t *testing.T) {
	tests := []struct {
		name      string
		model     func() Model
		notWanted []string
	}{
		{
			name: "groups log focus",
			model: func() Model {
				m := fixedModel()
				m.section = sectionGroups
				m.focus = focusLog
				m.help.ShowAll = true
				return m
			},
			notWanted: []string{"j/k scroll", "esc back", "enter log", "r run/retry"},
		},
		{
			name: "triage detail",
			model: func() Model {
				m := fixedModel()
				m.section = sectionTriage
				m.triageFailures = sampleFailures()
				m.triageCursor = 2
				m.triageDetailActive = true
				m.help.ShowAll = true
				return m
			},
			notWanted: []string{"i file issue", "t refresh", "K sync", "esc back", "enter detail"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m := tc.model()
			view := stripANSI(m.View())
			footer := stripANSI(m.renderFooter(m.width))
			for _, want := range []string{"Keyboard reference", "? or esc close", "q quit"} {
				if !strings.Contains(view, want) {
					t.Fatalf("expanded help view missing %q:\n%s", want, view)
				}
			}
			for _, want := range []string{"? or esc close", "q quit"} {
				if !strings.Contains(footer, want) {
					t.Fatalf("expanded help footer missing %q:\n%s", want, footer)
				}
			}
			for _, unwanted := range tc.notWanted {
				if strings.Contains(footer, unwanted) {
					t.Fatalf("expanded help footer should not contain nested hint %q:\n%s", unwanted, footer)
				}
			}
		})
	}
}

func TestExpandedHelpDefersToHigherPriorityBodies(t *testing.T) {
	tests := []struct {
		name          string
		msg           tea.Msg
		wantBody      []string
		wantFooter    []string
		notWantBody   []string
		notWantFooter []string
	}{
		{
			name: "resume prompt",
			msg: ResumePromptMsg{
				When:   time.Date(2026, 1, 15, 10, 30, 0, 0, time.UTC),
				Failed: 3,
				NotRun: 7,
			},
			wantBody:      []string{"resume last run"},
			wantFooter:    []string{"enter resume", "esc dismiss"},
			notWantBody:   []string{"Keyboard reference"},
			notWantFooter: []string{"? or esc close"},
		},
		{
			name: "result overlay",
			msg: ReportExportDoneMsg{
				MarkdownPath: ".pulsar-report.md",
				HTMLPath:     ".pulsar-report.html",
			},
			wantBody:      []string{"Report exported", "markdown: .pulsar-report.md"},
			wantFooter:    []string{"esc close"},
			notWantBody:   []string{"Keyboard reference"},
			notWantFooter: []string{"? or esc close"},
		},
		{
			name: "orphan overlay",
			msg: OrphansListedMsg{
				Owner:     "acme-test",
				Resources: sampleResources(),
			},
			wantBody:      []string{"Orphan resources", "tf-acc-repo-a"},
			wantFooter:    []string{"s sweep", "esc close"},
			notWantBody:   []string{"Keyboard reference"},
			notWantFooter: []string{"? or esc close"},
		},
	}

	apply := func(t *testing.T, msg tea.Msg) Model {
		t.Helper()
		m := fixedModel()
		m.help.ShowAll = true
		next, cmd := m.Update(msg)
		if cmd != nil {
			t.Fatalf("unexpected command for %T", msg)
		}
		got := next.(Model)
		if !got.help.ShowAll {
			t.Fatalf("%T unexpectedly closed help state", msg)
		}
		return got
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := apply(t, tc.msg)
			view := stripANSI(got.View())
			footer := stripANSI(got.renderFooter(got.width))
			for _, want := range tc.wantBody {
				if !strings.Contains(view, want) {
					t.Fatalf("view missing %q:\n%s", want, view)
				}
			}
			for _, want := range tc.wantFooter {
				if !strings.Contains(footer, want) {
					t.Fatalf("footer missing %q:\n%s", want, footer)
				}
			}
			for _, unwanted := range tc.notWantBody {
				if strings.Contains(view, unwanted) {
					t.Fatalf("view should not contain %q:\n%s", unwanted, view)
				}
			}
			for _, unwanted := range tc.notWantFooter {
				if strings.Contains(footer, unwanted) {
					t.Fatalf("footer should not contain %q:\n%s", unwanted, footer)
				}
			}
		})
	}
}

func TestExpandedHelpUsesASCIIBorderWhenASCIIEnabled(t *testing.T) {
	m := fixedModel()
	m.help.ShowAll = true

	got := stripANSI(m.renderHelpPanel(m.width))
	for _, want := range []string{"+", "-", "|"} {
		if !strings.Contains(got, want) {
			t.Fatalf("ascii help panel missing %q:\n%s", want, got)
		}
	}
	for _, unwanted := range []string{"╭", "╮", "╰", "╯", "│", "─"} {
		if strings.Contains(got, unwanted) {
			t.Fatalf("ascii help panel should not contain %q:\n%s", unwanted, got)
		}
	}
}

func TestFooterHintsMatchCurrentView(t *testing.T) {
	tests := []struct {
		name      string
		model     func() Model
		want      []string
		notWanted []string
	}{
		{
			name: "log viewport",
			model: func() Model {
				m := fixedModel()
				m.section = sectionGroups
				m.focus = focusLog
				return m
			},
			want:      []string{"j/k scroll", "esc back"},
			notWanted: []string{"enter detail", "enter confirm"},
		},
		{
			name: "triage detail",
			model: func() Model {
				m := fixedModel()
				m.section = sectionTriage
				m.triageFailures = sampleFailures()
				m.triageCursor = 2
				m.triageDetailActive = true
				return m
			},
			want:      []string{"i file issue", "esc back"},
			notWanted: []string{"enter detail", "enter confirm"},
		},
		{
			name: "file issue confirmation",
			model: func() Model {
				m := fixedModel()
				m.section = sectionTriage
				m.fileIssueActive = true
				m.fileIssuePreview = IssuePreviewMsg{
					Fingerprint:      "sha256:" + strings.Repeat("c", 64),
					ShortFingerprint: "cccccccccccccccc",
				}
				return m
			},
			want:      []string{"enter confirm", "esc close"},
			notWanted: []string{"enter detail", "q quit", "←/→ tabs"},
		},
		{
			name: "sweep confirmation",
			model: func() Model {
				m := fixedModel()
				m.sweepConfirmActive = true
				m.orphanOwner = "acme-test"
				return m
			},
			want:      []string{"enter confirm", "esc close"},
			notWanted: []string{"enter detail", "q quit", "←/→ tabs"},
		},
		{
			name: "expanded help",
			model: func() Model {
				m := fixedModel()
				m.help.ShowAll = true
				return m
			},
			want:      []string{"? or esc close", "q quit"},
			notWanted: []string{"enter detail", "←/→ tabs"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := stripANSI(tc.model().renderFooter(fixedModel().width))
			for _, want := range tc.want {
				if !strings.Contains(got, want) {
					t.Fatalf("footer missing %q:\n%s", want, got)
				}
			}
			for _, unwanted := range tc.notWanted {
				if strings.Contains(got, unwanted) {
					t.Fatalf("footer should not contain %q:\n%s", unwanted, got)
				}
			}
		})
	}
}

func TestTriageFooterReflectsFailureAvailability(t *testing.T) {
	m := fixedModel()
	m.section = sectionTriage

	empty := stripANSI(m.renderFooter(m.width))
	if strings.Contains(empty, "enter detail") {
		t.Fatalf("empty triage footer advertises unavailable detail action:\n%s", empty)
	}

	m.triageFailures = sampleFailures()
	withFailures := stripANSI(m.renderFooter(m.width))
	if !strings.Contains(withFailures, "enter detail") {
		t.Fatalf("triage footer with failures missing detail action:\n%s", withFailures)
	}
}

func TestRunFooterReflectsGroupAvailability(t *testing.T) {
	m := fixedModel()
	m.section = sectionRun

	empty := stripANSI(m.renderFooter(m.width))
	for _, unavailable := range []string{"r run", "R retry failed"} {
		if strings.Contains(empty, unavailable) {
			t.Fatalf("run footer without groups advertises unavailable action %q:\n%s", unavailable, empty)
		}
	}

	m.groups = []engine.Group{{Name: "repositories", Tests: []string{"TestA"}}}
	withGroups := stripANSI(m.renderFooter(m.width))
	for _, available := range []string{"r run", "R retry failed"} {
		if !strings.Contains(withGroups, available) {
			t.Fatalf("run footer with groups missing action %q:\n%s", available, withGroups)
		}
	}
}

func TestViewTestsDrilldownGolden(t *testing.T) {
	m := fixedModel()
	m.section = sectionGroups
	m.focus = focusTests
	m.groups = []engine.Group{
		{Name: "repositories", Tests: []string{"TestAccA", "TestAccB", "TestAccC"}},
	}
	m.groupCursor = 0
	m.testCursor = 1
	m.results = []engine.TestResult{
		{Package: "pkg", Name: "TestAccA", Sub: "", Status: provider.StatusPass, Elapsed: 1.2},
		{Package: "pkg", Name: "TestAccB", Sub: "", Status: provider.StatusFail, Elapsed: 2.5},
		// TestAccC not run
	}
	m.resultIndex = map[resultKey]int{
		{Name: "TestAccA", Sub: ""}: 0,
		{Name: "TestAccB", Sub: ""}: 1,
	}

	out := m.View()
	assertNoANSI(t, out)
	if !strings.Contains(out, "TestAccA") {
		t.Error("expected TestAccA in drill-down view")
	}
	if !strings.Contains(out, "TestAccB") {
		t.Error("expected TestAccB in drill-down view")
	}
	compareOrUpdate(t, "tests_drilldown", out)
}

func TestViewLogViewportGolden(t *testing.T) {
	m := fixedModel()
	m.section = sectionGroups
	m.focus = focusLog
	m.groups = []engine.Group{
		{Name: "repositories", Tests: []string{"TestAccA", "TestAccB"}},
	}
	m.groupCursor = 0
	m.testCursor = 0
	m.results = []engine.TestResult{
		{
			Package: "pkg",
			Name:    "TestAccA",
			Sub:     "",
			Status:  provider.StatusFail,
			Elapsed: 2.5,
			Output:  []string{"=== RUN TestAccA", "    github_repository.test: creating...", "    token: ••••redacted••••", "    FAIL: assertion failed", "--- FAIL: TestAccA (2.50s)"},
		},
	}
	m.resultIndex = map[resultKey]int{
		{Name: "TestAccA", Sub: ""}: 0,
	}
	// Initialize the logVP with content by simulating drill-in
	m2, _ := m.Update(tea.WindowSizeMsg{Width: 100, Height: 30})
	m = m2.(Model)
	m.focus = focusTests
	m.testCursor = 0
	m3, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("l")})
	m = m3.(Model)

	out := m.View()
	assertNoANSI(t, out)
	if !strings.Contains(out, "••••redacted••••") {
		t.Error("expected redacted placeholder in log viewport")
	}
	compareOrUpdate(t, "log_viewport", out)
}

// TestViewResumeBanner verifies that a model with resumePrompt=true renders the
// banner containing "resume last run" and the accept/dismiss hints.
func TestViewResumeBanner(t *testing.T) {
	m := fixedModel()
	m.section = sectionGroups
	m.resumePrompt = true
	m.resumeWhen = time.Date(2026, 1, 15, 10, 30, 0, 0, time.UTC)
	m.resumeFailed = 3
	m.resumeNotRun = 7

	out := m.View()
	assertNoANSI(t, out)
	if !strings.Contains(out, "resume last run") {
		t.Errorf("expected 'resume last run' in output; got:\n%s", out)
	}
	if !strings.Contains(out, "enter resume") {
		t.Errorf("expected 'enter resume' hint in output; got:\n%s", out)
	}
	if !strings.Contains(out, "esc dismiss") {
		t.Errorf("expected 'esc dismiss' hint in output; got:\n%s", out)
	}
}

// TestViewResumeBannerNonGroupsSection guards the consistency between the global
// resume-prompt key interception and the banner: while the prompt is active the
// banner must be visible regardless of which section is focused, otherwise keys
// are silently consumed with no on-screen explanation.
func TestViewResumeBannerNonGroupsSection(t *testing.T) {
	for _, sec := range []section{sectionPreflight, sectionRun} {
		m := fixedModel()
		m.section = sec
		m.resumePrompt = true
		m.resumeWhen = time.Date(2026, 1, 15, 10, 30, 0, 0, time.UTC)
		m.resumeFailed = 1
		m.resumeNotRun = 2

		out := m.View()
		assertNoANSI(t, out)
		if !strings.Contains(out, "resume last run") {
			t.Errorf("section %d: expected resume banner regardless of section; got:\n%s", sec, out)
		}
	}
}

// ── overlay golden tests ───────────────────────────────────────────────────────

func TestViewPickerOverlayGolden(t *testing.T) {
	m := fixedModel().WithModes([]provider.Mode{
		{Name: "anonymous", Description: "No credentials; public endpoints only"},
		{Name: "individual", Description: "Personal account tests requiring GITHUB_OWNER and credentials"},
		{Name: "organization", Description: "Organization tests requiring org owner and credentials"},
		{Name: "team", Description: "Team tests; adds external user fixtures"},
		{Name: "enterprise", Description: "Enterprise tests; adds GITHUB_ENTERPRISE_SLUG"},
	})
	m.pickerActive = true
	m.pickerCursor = 2 // organization selected

	out := m.View()
	assertNoANSI(t, out)
	if !strings.Contains(out, "anonymous") {
		t.Error("expected 'anonymous' in picker output")
	}
	if !strings.Contains(out, "enterprise") {
		t.Error("expected 'enterprise' in picker output")
	}
	compareOrUpdate(t, "picker_overlay", out)
}

func TestViewEditorOverlayGolden(t *testing.T) {
	const secretValue = "ghp_FAKESECRET"
	envVars := []provider.EnvVar{
		{Key: "GITHUB_OWNER", Required: true, Secret: false, Doc: "GitHub org owner"},
		{Key: "GITHUB_TOKEN", Required: false, Secret: true, Doc: "PAT; or use the GITHUB_APP_* trio"},
		{Key: "GITHUB_APP_ID", Required: false, Secret: false, Doc: "GitHub App ID"},
	}
	getenv := func(k string) string {
		switch k {
		case "GITHUB_OWNER":
			return "my-test-org"
		case "GITHUB_TOKEN":
			return secretValue // must NOT appear in output
		default:
			return ""
		}
	}
	m := fixedModel().WithEnvVars(envVars).WithGetenv(getenv)
	m.editorActive = true
	m.editorCursor = 0
	m.editorFields = buildEditorFields(envVars, getenv)

	out := m.View()
	assertNoANSI(t, out)

	// Redaction contract: secret value must NEVER appear.
	if strings.Contains(out, secretValue) {
		t.Errorf("secret value leaked into editor View(): %q", out)
	}
	// Should show <set> for GITHUB_TOKEN (it is set).
	if !strings.Contains(out, "<set>") {
		t.Errorf("expected <set> for set secret field; got: %q", out)
	}
	// Should show current value for non-secret GITHUB_OWNER.
	if !strings.Contains(out, "my-test-org") {
		t.Errorf("expected current value 'my-test-org' for non-secret field; got: %q", out)
	}
	compareOrUpdate(t, "editor_overlay", out)
}

// TestPickerDescriptionNarrowTerminalDoesNotOverflow verifies that a terminal
// width < 20 does not let the full mode description leak into the overlay output.
// Without the max(1, ...) clamp, truncate receives a non-positive maxLen and
// returns the string unchanged, overflowing the layout.
func TestPickerDescriptionNarrowTerminalDoesNotOverflow(t *testing.T) {
	const longDesc = "No credentials; public endpoints only"
	m := fixedModel().WithModes([]provider.Mode{
		{Name: "anon", Description: longDesc},
	})
	m.width = 10
	m.pickerActive = true

	out := m.renderPickerOverlay(m.width)
	if strings.Contains(out, longDesc) {
		t.Errorf("full description appeared in narrow-terminal picker (width=%d); expected it to be truncated", m.width)
	}
}

func TestPickerAndEditorOverlaysUseSharedPanels(t *testing.T) {
	getenv := func(key string) string {
		if key == "GITHUB_OWNER" {
			return "my-test-org"
		}
		return ""
	}
	envVars := []provider.EnvVar{{Key: "GITHUB_OWNER", Required: true}}
	tests := []struct {
		name string
		out  string
		want []string
	}{
		{
			name: "picker",
			out: fixedModel().WithModes([]provider.Mode{
				{Name: "anonymous", Description: "No credentials; public endpoints only"},
			}).renderPickerOverlay(100),
			want: []string{"+", "|", "Select mode", "anonymous", "enter select", "esc/q close"},
		},
		{
			name: "editor",
			out: func() string {
				m := fixedModel().WithEnvVars(envVars).WithGetenv(getenv)
				m.editorFields = buildEditorFields(envVars, getenv)
				return m.renderEditorOverlay(100)
			}(),
			want: []string{"+", "|", "Config: organization", "GITHUB_OWNER", "enter edit", "esc/q close", "p preflight"},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := stripANSI(tc.out)
			for _, want := range tc.want {
				if !strings.Contains(got, want) {
					t.Fatalf("%s overlay missing shared panel marker %q:\n%s", tc.name, want, got)
				}
			}
			assertRenderedWidth(t, got, 100)
		})
	}
}

func TestPickerAndEditorOverlaysFitRepresentativeWidths(t *testing.T) {
	getenv := func(key string) string {
		if key == "GITHUB_OWNER" {
			return "my-test-org"
		}
		return ""
	}
	envVars := []provider.EnvVar{
		{Key: "GITHUB_OWNER", Required: true},
		{Key: "GITHUB_ENTERPRISE_SLUG_WITH_LONG_NAME"},
	}
	for _, width := range []int{48, 80, 120} {
		m := fixedModel().WithModes([]provider.Mode{
			{Name: "anonymous", Description: "No credentials; public endpoints only"},
			{Name: "organization", Description: "Organization tests requiring org owner and credentials"},
		})
		assertRenderedWidth(t, m.renderPickerOverlay(width), width)

		m = fixedModel().WithEnvVars(envVars).WithGetenv(getenv)
		m.editorFields = buildEditorFields(envVars, getenv)
		assertRenderedWidth(t, m.renderEditorOverlay(width), width)
	}
}
