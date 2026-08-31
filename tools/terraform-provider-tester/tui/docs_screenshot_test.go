package tui

import (
	"image/png"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"

	"github.com/github/terraform-provider-tester/engine"
	"github.com/github/terraform-provider-tester/provider"
)

const (
	// docsScreenshotBackground is the GitHub-dark canvas fill behind each frame.
	docsScreenshotBackground = "#0d1117"
	// docsScreenshotWidth and docsScreenshotHeight are the reference frame size
	// the synthetic models are laid out for. The freeze renderer emits a
	// higher-resolution image sized to the content, so these are treated as
	// minimum acceptable dimensions rather than exact ones.
	docsScreenshotWidth     = 1600
	docsScreenshotHeight    = 900
	docsScreenshotMinWidth  = 1200
	docsScreenshotMinHeight = 600
	docsSyntheticIdentity   = "authenticated as test-user-1"
	docsIntroWordmarkLine   = "######  ##  ##      ##  ######  ######"
	docsIntroSkipHint       = "press any key to skip"
	docsPreflightSlugFix    = "set GITHUB_ENTERPRISE_SLUG=<slug>"
	docsPreflightScopeFix   = "use classic PAT for scope verification"
)

var credentialLikeRE = regexp.MustCompile(`(?:gh[pousr]_[A-Za-z0-9_]+|github_pat_[A-Za-z0-9_]+)`)

func docsImageDir(t *testing.T) string {
	t.Helper()
	dir, err := filepath.Abs(filepath.Join("..", "docs", "images"))
	if err != nil {
		t.Fatalf("resolve docs image dir: %v", err)
	}
	return dir
}

// docsFreezeBin locates the charmbracelet/freeze renderer used to turn the
// console's ANSI views into PNGs. It checks PATH first, then the default
// `go install` location, and reports whether a binary was found.
func docsFreezeBin() (string, bool) {
	if p, err := exec.LookPath("freeze"); err == nil {
		return p, true
	}
	if home, err := os.UserHomeDir(); err == nil && home != "" {
		cand := filepath.Join(home, "go", "bin", "freeze")
		if _, err := os.Stat(cand); err == nil {
			return cand, true
		}
	}
	return "", false
}

func assertSyntheticScreenshotSource(t *testing.T, name, view string) string {
	t.Helper()
	plain := stripANSI(view)
	if !strings.Contains(plain, ConsoleMark) {
		t.Fatalf("%s view missing console mark %q\n%s", name, ConsoleMark, plain)
	}
	if name != "intro" {
		for _, want := range []string{
			"* GH/TF Provider Acceptance 1.2.3",
			"1 Preflight",
			"2 Groups",
			"3 Run",
			"4 Triage",
			"left/right tabs",
			"? help",
			"q quit",
		} {
			if !strings.Contains(plain, want) {
				t.Fatalf("%s view missing chrome %q\n%s", name, want, plain)
			}
		}
	}
	if credentialLikeRE.MatchString(plain) {
		t.Fatalf("%s view contains credential-shaped token\n%s", name, plain)
	}
	if strings.Contains(plain, "-----BEGIN ") {
		t.Fatalf("%s view contains PEM marker\n%s", name, plain)
	}
	if strings.Contains(plain, "owner:") && !strings.Contains(plain, "my-test-org") {
		t.Fatalf("%s view contains non-synthetic owner\n%s", name, plain)
	}
	if strings.Contains(plain, "github.com/") || strings.Contains(plain, "/Users/") {
		t.Fatalf("%s view contains non-synthetic path or host data\n%s", name, plain)
	}

	switch name {
	case "intro":
		for _, want := range []string{
			docsIntroWordmarkLine,
			ConsoleTitle + " 1.2.3",
			"terraform-provider-tester",
			Tagline,
			docsIntroSkipHint,
		} {
			if !strings.Contains(plain, want) {
				t.Fatalf("intro view missing %q\n%s", want, plain)
			}
		}
	case "preflight":
		if strings.Contains(plain, "octocat") {
			t.Fatalf("preflight view must not contain octocat\n%s", plain)
		}
		for _, want := range []string{
			"identity",
			docsSyntheticIdentity,
			"token-scopes",
			docsPreflightScopeFix,
			"enterprise",
			"GITHUB_ENTERPRISE_SLUG not set",
			docsPreflightSlugFix,
			"READY 1",
			"WARNINGS 1",
			"BLOCKED 1",
			"Next action",
		} {
			if !strings.Contains(plain, want) {
				t.Fatalf("preflight view missing %q\n%s", want, plain)
			}
		}
	case "groups":
		for _, want := range []string{
			"GROUP",
			"PASS",
			"TIME",
			"PROGRESS",
			"TOTAL",
			"repositories",
			"teams",
			"actions-secrets",
			"GROUPS 3",
			"TESTS 6",
			"enter tests",
			"r run/retry",
			"f failures first",
			"a all",
		} {
			if !strings.Contains(plain, want) {
				t.Fatalf("groups view missing %q\n%s", want, plain)
			}
		}
	case "run":
		for _, want := range []string{
			"Run progress",
			"Running",
			"PASSED 1",
			"FAILED 1",
			"RUNNING 1",
			"REMAINING 3",
			"Target and scope",
			"Safe actions",
			"repositories",
			"33%",
			"r run selection",
			"R retry failures",
			"e export report",
		} {
			if !strings.Contains(plain, want) {
				t.Fatalf("run view missing %q\n%s", want, plain)
			}
		}
	case "triage":
		for _, want := range []string{
			"Triage",
			"REAL 2",
			"UNSTABLE 1",
			"FLAKES 1",
			"KNOWN 1",
			"ELIGIBLE 2",
			"TestAccKnown",
			"TestAccReal",
			"known #42",
			"eligible",
			"bbbbbbbbbbbbbbbb",
		} {
			if !strings.Contains(plain, want) {
				t.Fatalf("triage view missing %q\n%s", want, plain)
			}
		}
	}
	return plain
}

func writeScreenshot(t *testing.T, freezeBin, dir, name, view string) {
	t.Helper()
	assertSyntheticScreenshotSource(t, name, view)

	ansiFile, err := os.CreateTemp(dir, "."+name+"-*.ansi")
	if err != nil {
		t.Fatalf("create temp ansi for %s: %v", name, err)
	}
	ansiPath := ansiFile.Name()
	defer func() { _ = os.Remove(ansiPath) }()
	if err := ansiFile.Chmod(0o600); err != nil {
		_ = ansiFile.Close()
		t.Fatalf("chmod temp ansi for %s: %v", name, err)
	}
	if _, err := ansiFile.WriteString(view); err != nil {
		_ = ansiFile.Close()
		t.Fatalf("write temp ansi for %s: %v", name, err)
	}
	if err := ansiFile.Close(); err != nil {
		t.Fatalf("close temp ansi for %s: %v", name, err)
	}

	tempPNG, err := os.CreateTemp(dir, "."+name+"-*.png")
	if err != nil {
		t.Fatalf("create temp png for %s: %v", name, err)
	}
	tempPNGPath := tempPNG.Name()
	defer func() { _ = os.Remove(tempPNGPath) }()
	if err := tempPNG.Close(); err != nil {
		t.Fatalf("close temp png for %s: %v", name, err)
	}

	out := filepath.Join(dir, name+".png")
	cmd := exec.Command(
		freezeBin,
		ansiPath,
		"--output", tempPNGPath,
		"--background", docsScreenshotBackground,
		"--padding", "32",
		"--margin", "0",
	)
	if output, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("render %s: %v\n%s", name, err, output)
	}
	if err := validateAndPromotePNG(tempPNGPath, out, docsScreenshotMinWidth, docsScreenshotMinHeight); err != nil {
		t.Fatalf("promote %s screenshot: %v", name, err)
	}
}

func docsPreflightModel() Model {
	m := fixedModel()
	m.section = sectionPreflight
	m.checks = []provider.Check{
		{Name: "identity", Status: provider.CheckOK, Detail: docsSyntheticIdentity},
		{Name: "token-scopes", Status: provider.CheckWarn, Detail: "scopes unverified", Fix: docsPreflightScopeFix},
		{Name: "enterprise", Status: provider.CheckFail, Detail: "GITHUB_ENTERPRISE_SLUG not set", Fix: docsPreflightSlugFix},
	}
	return m
}

func docsGroupsModel() Model {
	m := fixedModel()
	m.section = sectionGroups
	m.groups = []engine.Group{
		{Name: "repositories", Tests: []string{"TestAccA", "TestAccB", "TestAccC"}},
		{Name: "teams", Tests: []string{"TestAccD", "TestAccE"}},
		{Name: "actions-secrets", Tests: []string{"TestAccF"}},
	}
	return m
}

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

func docsTriageModel(t *testing.T) Model {
	t.Helper()
	m := fixedModel()
	m.section = sectionTriage
	next, _ := m.Update(TriageLoadedMsg{Failures: sampleFailures(), CacheAvailable: true})
	got, ok := next.(Model)
	if !ok {
		t.Fatal("triage update did not return a Model")
	}
	return got
}

func TestUpdateDocsScreenshots(t *testing.T) {
	if os.Getenv("UPDATE_DOC_IMAGES") != "1" {
		t.Skip("set UPDATE_DOC_IMAGES=1 to regenerate docs screenshots")
	}
	freezeBin, ok := docsFreezeBin()
	if !ok {
		t.Skip("freeze is not installed; run: go install github.com/charmbracelet/freeze@latest")
	}

	// The synthetic models render plain text under the Ascii profile that a
	// non-TTY `go test` selects. Force a truecolor, dark-background renderer so
	// the console's real GH/TF palette (Terraform purple, GitHub blue, and the
	// pass/fail/attention tones) survives into the ANSI freeze consumes.
	prevProfile := lipgloss.ColorProfile()
	prevDark := lipgloss.HasDarkBackground()
	lipgloss.SetColorProfile(termenv.TrueColor)
	lipgloss.SetHasDarkBackground(true)
	defer func() {
		lipgloss.SetColorProfile(prevProfile)
		lipgloss.SetHasDarkBackground(prevDark)
	}()

	dir := docsImageDir(t)
	writeScreenshot(t, freezeBin, dir, "intro", renderIntro(introPeakFrame, 100, 30, "1.2.3", true))
	writeScreenshot(t, freezeBin, dir, "preflight", docsPreflightModel().View())
	writeScreenshot(t, freezeBin, dir, "groups", docsGroupsModel().View())
	writeScreenshot(t, freezeBin, dir, "run", docsRunModel().View())
	writeScreenshot(t, freezeBin, dir, "triage", docsTriageModel(t).View())
}

func TestDocsScreenshotSources(t *testing.T) {
	assertSyntheticScreenshotSource(t, "intro", renderIntro(introPeakFrame, 100, 30, "1.2.3", true))
	assertSyntheticScreenshotSource(t, "preflight", docsPreflightModel().View())
	assertSyntheticScreenshotSource(t, "groups", docsGroupsModel().View())
	assertSyntheticScreenshotSource(t, "run", docsRunModel().View())
	assertSyntheticScreenshotSource(t, "triage", docsTriageModel(t).View())
}

func TestDocsScreenshots(t *testing.T) {
	dir := docsImageDir(t)
	wantNames := []string{"intro.png", "preflight.png", "groups.png", "run.png", "triage.png"}
	want := make(map[string]bool, len(wantNames))
	for _, name := range wantNames {
		want[name] = true
		path := filepath.Join(dir, name)
		f, err := os.Open(path)
		if err != nil {
			t.Fatalf("open docs screenshot %s: %v", name, err)
		}
		cfg, err := png.DecodeConfig(f)
		closeErr := f.Close()
		if err != nil {
			t.Fatalf("decode docs screenshot %s: %v", name, err)
		}
		if closeErr != nil {
			t.Fatalf("close docs screenshot %s: %v", name, closeErr)
		}
		if cfg.Width < docsScreenshotMinWidth || cfg.Height < docsScreenshotMinHeight {
			t.Fatalf("%s dimensions = %dx%d, want at least %dx%d", name, cfg.Width, cfg.Height, docsScreenshotMinWidth, docsScreenshotMinHeight)
		}
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatalf("read docs image dir: %v", err)
	}
	var got []string
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".png" {
			continue
		}
		got = append(got, entry.Name())
		if !want[entry.Name()] {
			t.Fatalf("unexpected docs screenshot %s in %s", entry.Name(), dir)
		}
	}
	if len(got) != len(wantNames) {
		t.Fatalf("docs screenshot inventory = %v, want %v", got, wantNames)
	}
}
