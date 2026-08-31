package cli

import (
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"testing"

	"github.com/github/terraform-provider-tester/engine"
	githubprovider "github.com/github/terraform-provider-tester/provider/github"
)

// This file wires the Terraform Provider Tester/terraform-provider-tester
// documentation to the code so the two cannot drift apart silently. It is a
// fork-safe check: it runs in the harness's own `go test ./...`, imports only
// harness packages, and never touches the provider's tfplugindocs pipeline.
//
// It enforces three properties of every covered operator-doc file in the
// repository:
//   1. Every environment variable the docs name is one the harness actually
//      reads (derived from the provider's EnvFor + a small, cited set of
//      behavior/provider-level vars).
//   2. Every documented tester binary invocation names a real subcommand.
//   3. The CLI reference documents every real subcommand.

// canonicalEnv builds the set of env vars the harness legitimately reads.
// Mode-specific vars come straight from the provider so this stays correct as
// the provider evolves; the rest are documented behavior/provider-level vars,
// each cited to its source.
func canonicalEnv(t *testing.T) map[string]bool {
	t.Helper()
	set := map[string]bool{}
	prov := githubprovider.New()
	for _, m := range prov.Modes() {
		for _, ev := range prov.EnvFor(m.Name) {
			set[ev.Key] = true
		}
	}
	// Documented vars that are not mode inputs (see cli/cli.go and
	// provider/github/preflight.go). Keep this list in sync with the docs.
	for _, k := range []string{
		"GH_TEST_AUTH_MODE",    // cli.go: default mode when --mode is absent
		"PULSAR_ENV_FILE",      // cli.go: env file path when --env-file is absent
		"PULSAR_FORCE_TTY",     // cli.go: force the dashboard on a non-TTY
		"PULSAR_NO_TUI",        // cli.go: never launch the dashboard
		"NO_COLOR",             // cli.go: ASCII glyphs, no color
		"TPT_PROVIDER_ROOT",    // smoke_test.go: provider checkout for the tagged live smoke
		"TF_ACC",               // harness sets it for the child go test run
		"TF_LOG",               // forwarded only with --allow-sensitive-logs
		"GITHUB_BASE_URL",      // preflight.go: GHES base URL
		"GITHUB_LEGACY_CLIENT", // provider-level toggle
	} {
		set[k] = true
	}
	return set
}

// envTokenRE matches whole env-var-looking identifiers with the prefixes the
// harness uses, plus the three exact provider-level names.
var envTokenRE = regexp.MustCompile(`(GH_TEST_[A-Z0-9_]*[A-Z0-9]|GITHUB_[A-Z0-9_]*[A-Z0-9]|PULSAR_[A-Z0-9_]*[A-Z0-9]|TPT_PROVIDER_[A-Z0-9_]*[A-Z0-9]|TF_ACC|TF_LOG|NO_COLOR)`)

// unknownEnvTokens returns env tokens in content that are not in canonical.
// Wildcard stems (a token immediately followed by '_' or '*', e.g. the docs'
// "GITHUB_APP_*" shorthand) are skipped so prose globs do not trip the check.
func unknownEnvTokens(content string, canonical map[string]bool) []string {
	var unknown []string
	seen := map[string]bool{}
	for _, loc := range envTokenRE.FindAllStringIndex(content, -1) {
		tok := content[loc[0]:loc[1]]
		if end := loc[1]; end < len(content) {
			if next := content[end]; next == '_' || next == '*' {
				continue // part of a longer name or a wildcard stem
			}
		}
		if canonical[tok] || seen[tok] {
			continue
		}
		seen[tok] = true
		unknown = append(unknown, tok)
	}
	sort.Strings(unknown)
	return unknown
}

var subcommandTokenRE = regexp.MustCompile(`^[a-z][a-z0-9-]*$`)

func inlineSubcommands(content string) []string {
	var out []string
	for _, snippet := range markdownCommandSnippets(content) {
		out = append(out, subcommandsInCommandSnippet(snippet)...)
	}
	return out
}

func markdownCommandSnippets(content string) []string {
	var snippets []string
	for i := 0; i < len(content); {
		if strings.HasPrefix(content[i:], "```") {
			infoStart := i + 3
			lineEndRel := strings.IndexByte(content[infoStart:], '\n')
			if lineEndRel < 0 {
				break
			}
			infoEnd := infoStart + lineEndRel
			codeStart := infoEnd + 1
			endRel := strings.Index(content[codeStart:], "```")
			if endRel < 0 {
				break
			}
			if scanFenceInfo(content[infoStart:infoEnd]) {
				snippets = append(snippets, content[codeStart:codeStart+endRel])
			}
			i = codeStart + endRel + 3
			continue
		}
		if content[i] == '`' {
			endRel := strings.IndexByte(content[i+1:], '`')
			if endRel < 0 {
				break
			}
			snippets = append(snippets, content[i+1:i+1+endRel])
			i += endRel + 2
			continue
		}
		i++
	}
	return snippets
}

func scanFenceInfo(info string) bool {
	fields := strings.Fields(strings.TrimSpace(info))
	if len(fields) == 0 {
		return true
	}
	switch fields[0] {
	case "sh", "shell", "bash":
		return true
	default:
		return false
	}
}

func subcommandsInCommandSnippet(snippet string) []string {
	var out []string
	for _, line := range strings.Split(snippet, "\n") {
		for _, segment := range shellCommandSegments(line) {
			if sub, ok := testerSubcommandFromSegment(segment); ok {
				out = append(out, sub)
			}
		}
	}
	return out
}

func shellCommandSegments(line string) []string {
	var segments []string
	start := 0
	var quote rune
	for i, r := range line {
		switch {
		case quote != 0:
			if r == quote {
				quote = 0
			}
		case r == '\'' || r == '"':
			quote = r
		case r == ';':
			segments = append(segments, line[start:i])
			start = i + len(string(r))
		case r == '&' && i+1 < len(line) && line[i+1] == '&':
			segments = append(segments, line[start:i])
			start = i + 2
		case r == '|' && i+1 < len(line) && line[i+1] == '|':
			segments = append(segments, line[start:i])
			start = i + 2
		}
	}
	segments = append(segments, line[start:])
	return segments
}

func testerSubcommandFromSegment(segment string) (string, bool) {
	fields := strings.Fields(strings.TrimSpace(segment))
	for len(fields) > 0 {
		tok := cleanShellToken(fields[0])
		if tok == "" {
			fields = fields[1:]
			continue
		}
		switch tok {
		case "if", "then", "do", "!", "time":
			fields = fields[1:]
			continue
		}
		if strings.Contains(tok, "=") && !strings.HasPrefix(tok, "$") {
			fields = fields[1:]
			continue
		}
		break
	}
	if len(fields) == 0 || !isTesterBinaryToken(fields[0]) {
		return "", false
	}
	for i := 1; i < len(fields); i++ {
		tok := cleanShellToken(fields[i])
		if tok == "" {
			continue
		}
		if strings.HasPrefix(tok, "-") {
			if flagTakesValue(tok) && !strings.Contains(tok, "=") {
				i++
			}
			continue
		}
		if subcommandTokenRE.MatchString(tok) {
			return tok, true
		}
		return "", false
	}
	return "", false
}

func cleanShellToken(tok string) string {
	return strings.Trim(tok, "\"'`")
}

func isTesterBinaryToken(tok string) bool {
	tok = cleanShellToken(tok)
	if tok == "$TPT_BIN" {
		return true
	}
	return tok == "terraform-provider-tester" ||
		tok == "./bin/terraform-provider-tester" ||
		tok == "bin/terraform-provider-tester" ||
		strings.HasSuffix(tok, "/terraform-provider-tester")
}

func flagTakesValue(flag string) bool {
	flag = cleanShellToken(flag)
	switch flag {
	case "--env-file":
		return true
	default:
		return false
	}
}

func realSubcommands() map[string]bool {
	set := map[string]bool{}
	for _, s := range []string{
		"e2e", "preflight", "groups", "run", "retry", "resume",
		"report", "triage", "known-issues", "orphans", "sweep", "unlock", "version", "discover",
	} {
		set[s] = true
	}
	return set
}

// harnessDocs returns the paths of every Markdown file the check inspects:
// the harness README, the whole docs/ tree, and Copilot-facing skill docs.
// Paths are relative to the cli package directory (the test working directory).
func harnessDocs(t *testing.T) []string {
	t.Helper()
	var paths []string
	if _, err := os.Stat("../docs"); err != nil {
		t.Skipf("docs tree not found from %s: %v", mustWD(t), err)
	}
	if _, err := os.Stat("../README.md"); err == nil {
		paths = append(paths, "../README.md")
	}
	if _, err := os.Stat("../llms.txt"); err == nil {
		paths = append(paths, "../llms.txt")
	}
	entries, err := filepath.Glob("../docs/*.md")
	if err != nil {
		t.Fatalf("glob docs: %v", err)
	}
	paths = append(paths, entries...)
	for _, path := range skillDocs() {
		if _, err := os.Stat(path); err == nil {
			paths = append(paths, path)
		}
	}
	return paths
}

func skillDocs() []string {
	return []string{
		"../.github/copilot-instructions.md",
		"../.github/skills/terraform-provider-tester/README.md",
		"../.github/skills/terraform-provider-tester/SKILL.md",
	}
}

func mustWD(t *testing.T) string {
	t.Helper()
	wd, _ := os.Getwd()
	return wd
}

func readDoc(t *testing.T, path string) string {
	t.Helper()
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	return string(b)
}

func TestDocsReferenceOnlyRealEnvVars(t *testing.T) {
	canonical := canonicalEnv(t)
	for _, path := range harnessDocs(t) {
		content := readDoc(t, path)
		if unknown := unknownEnvTokens(content, canonical); len(unknown) > 0 {
			t.Errorf("%s names env vars the harness does not read: %v\n"+
				"If these are real, add them to the provider's EnvFor or the cited list in canonicalEnv.",
				path, unknown)
		}
	}
}

func TestDocsReferenceOnlyRealSubcommands(t *testing.T) {
	real := realSubcommands()
	for _, path := range harnessDocs(t) {
		content := readDoc(t, path)
		for _, sub := range inlineSubcommands(content) {
			if !real[sub] {
				t.Errorf("%s shows `terraform-provider-tester %s`, which is not a real subcommand", path, sub)
			}
		}
	}
}

func TestDocsCopilotFacingResolvedBinarySubcommandsAreValidated(t *testing.T) {
	real := realSubcommands()
	for _, path := range []string{
		"../.github/skills/terraform-provider-tester/SKILL.md",
		"../.github/copilot-instructions.md",
	} {
		content := readDoc(t, path)
		subs := inlineSubcommands(content)
		if len(subs) == 0 {
			t.Fatalf("%s yielded no documented tester subcommands", path)
		}
		unique := map[string]bool{}
		for _, sub := range subs {
			if !real[sub] {
				t.Fatalf("%s shows non-real tester subcommand %q", path, sub)
			}
			unique[sub] = true
		}
		names := make([]string, 0, len(unique))
		for sub := range unique {
			names = append(names, sub)
		}
		sort.Strings(names)
		t.Logf("%s extracted %d tester subcommands (%d unique): %s", path, len(subs), len(names), strings.Join(names, ", "))
	}
}

func TestDocsSubcommandScannerCatchesResolvedBinaryDrift(t *testing.T) {
	content := strings.Join([]string{
		"`terraform-provider-tester preflight --json`",
		"`./bin/terraform-provider-tester groups --json`",
		"`\"$TPT_BIN\" run --repo-root \"$PROVIDER_ROOT\" --json`",
		"`$TPT_BIN retry --failed --json`",
		"```sh\n\"$TPT_BIN\" --env-file /path/to/env resume --repo-root \"$PROVIDER_ROOT\"\n```",
	}, "\n")
	got := inlineSubcommands(content)
	want := []string{"preflight", "groups", "run", "retry", "resume"}
	if strings.Join(got, ",") != strings.Join(want, ",") {
		t.Fatalf("scanner missed resolved binary invocations: got %v want %v", got, want)
	}

	bad := "`\"$TPT_BIN\" made-up --repo-root \"$PROVIDER_ROOT\" --json`"
	got = inlineSubcommands(bad)
	if len(got) != 1 || got[0] != "made-up" || realSubcommands()[got[0]] {
		t.Fatalf("scanner failed to expose bogus resolved-binary subcommand: got %v", got)
	}
}

func TestDocsSubcommandScannerRecognizesGuidedE2EAndRejectsTypo(t *testing.T) {
	good := "`\"$TPT_BIN\" e2e --repo-root \"$PROVIDER_ROOT\" --json`"
	got := inlineSubcommands(good)
	if len(got) != 1 || got[0] != "e2e" {
		t.Fatalf("scanner missed guided e2e command: got %v", got)
	}

	bad := "`\"$TPT_BIN\" e3e --repo-root \"$PROVIDER_ROOT\" --json`"
	got = inlineSubcommands(bad)
	if len(got) != 1 || got[0] != "e3e" {
		t.Fatalf("scanner failed to expose typo subcommand: got %v", got)
	}
	if realSubcommands()[got[0]] {
		t.Fatalf("scanner treated typo %q as a real subcommand", got[0])
	}
}

func TestDocsSubcommandScannerIgnoresProseAndAssignments(t *testing.T) {
	content := strings.Join([]string{
		"`TPT_BIN=\"$TPT_HOME_REAL/bin/terraform-provider-tester\"`",
		"`command -v terraform-provider-tester`",
		"```sh\nTPT_BIN=\"$TPT_HOME_REAL/bin/terraform-provider-tester\"\necho \"terraform-provider-tester on PATH did not resolve\"\n```",
	}, "\n")
	if got := inlineSubcommands(content); len(got) != 0 {
		t.Fatalf("scanner extracted non-command subcommands: %v", got)
	}
}

func TestCLIReferenceCoversEverySubcommand(t *testing.T) {
	const ref = "../docs/cli-reference.md"
	if _, err := os.Stat(ref); err != nil {
		t.Skipf("cli-reference not found: %v", err)
	}
	content := readDoc(t, ref)
	for sub := range realSubcommands() {
		if !strings.Contains(content, "`"+sub+"`") && !strings.Contains(content, "terraform-provider-tester "+sub) {
			t.Errorf("cli-reference.md does not document the %q subcommand", sub)
		}
	}
}

// TestEnvScannerCatchesDrift proves the scanner actually fails on a bad token,
// so a green run of the checks above is meaningful and not vacuous.
func TestEnvScannerCatchesDrift(t *testing.T) {
	canonical := canonicalEnv(t)
	good := "Set `GITHUB_TOKEN`, `GH_TEST_ORG_USER1`, and `TPT_PROVIDER_ROOT`. `TPT_BIN` and `TPT_HOME` are shell locals. The `GITHUB_APP_*` trio is optional."
	if got := unknownEnvTokens(good, canonical); len(got) != 0 {
		t.Fatalf("scanner flagged real vars / wildcard stem: %v", got)
	}
	bad := "Set `GH_TEST_MADE_UP`, `GITHUB_NOT_A_REAL_VAR`, and `TPT_PROVIDER_NOT_A_REAL_VAR`."
	got := unknownEnvTokens(bad, canonical)
	want := []string{"GH_TEST_MADE_UP", "GITHUB_NOT_A_REAL_VAR", "TPT_PROVIDER_NOT_A_REAL_VAR"}
	if strings.Join(got, ",") != strings.Join(want, ",") {
		t.Fatalf("scanner missed drift: got %v want %v", got, want)
	}
}

func TestSkillBootstrapsTesterOutsideProviderCheckout(t *testing.T) {
	skill := readDoc(t, "../.github/skills/terraform-provider-tester/SKILL.md")
	want := []string{
		`PROVIDER_ROOT="$(git rev-parse --show-toplevel)"`,
		`test -f "$PROVIDER_ROOT/go.mod"`,
		`test -d "$PROVIDER_ROOT/github"`,
		`TPT_HOME="$HOME/.local/share/terraform-provider-tester"`,
		`command -v terraform-provider-tester`,
		`gh repo clone github/terraform-provider-tester "$TPT_HOME" -- --depth 1`,
		`make -C "$TPT_HOME" build`,
		`TPT_BIN=`,
	}
	for _, snippet := range want {
		if !strings.Contains(skill, snippet) {
			t.Errorf("skill bootstrap is missing %q", snippet)
		}
	}
}

func TestSkillBootstrapValidatesCachedTesterCheckoutIdentity(t *testing.T) {
	skill := readDoc(t, "../.github/skills/terraform-provider-tester/SKILL.md")
	want := []string{
		`TPT_HOME="$HOME/.local/share/terraform-provider-tester"`,
		`case "$HOME" in`,
		`git -C "$TPT_HOME" rev-parse --show-toplevel`,
		`test "$TPT_TOP" = "$TPT_HOME_REAL"`,
		`if test "$TPT_HOME_REAL" = "$PROVIDER_ROOT"; then`,
		`test "$TPT_MODULE" = "github.com/github/terraform-provider-tester"`,
		`git -C "$TPT_HOME" config --get remote.origin.url`,
		`case "$TPT_ORIGIN" in`,
		`https://github.com/github/terraform-provider-tester.git`,
		`git@github.com:github/terraform-provider-tester.git`,
		`ssh://git@github.com/github/terraform-provider-tester.git`,
		`grep -Eq '^build:' "$TPT_HOME/Makefile"`,
		`cd -P "$TPT_PATH_DIR"`,
		`TPT_BIN="$TPT_PATH_DIR/$TPT_PATH_BASE"`,
	}
	for _, snippet := range want {
		if !strings.Contains(skill, snippet) {
			t.Errorf("skill bootstrap is missing identity/path guard %q", snippet)
		}
	}
	for _, stale := range []string{
		`TPT_HOME="${TPT_HOME:-$HOME/.local/share/terraform-provider-tester}"`,
		`TPT_BIN="$TPT_ON_PATH"`,
	} {
		if strings.Contains(skill, stale) {
			t.Errorf("skill bootstrap still allows stale path behavior %q", stale)
		}
	}

	mirror := readDoc(t, "../.github/copilot-instructions.md")
	for _, snippet := range []string{
		"fixed `$HOME/.local/share/terraform-provider-tester` cache",
		"not `$PROVIDER_ROOT`",
		"`go.mod` declares the expected module",
		"`origin` identifies `github/terraform-provider-tester`",
		"Makefile build target exists",
	} {
		if !strings.Contains(mirror, snippet) {
			t.Errorf("copilot instructions are missing cached tester guard %q", snippet)
		}
	}
}

func TestSkillCommandsUseResolvedBinaryAndProviderRoot(t *testing.T) {
	skill := readDoc(t, "../.github/skills/terraform-provider-tester/SKILL.md")
	want := []string{
		`"$TPT_BIN" e2e --repo-root "$PROVIDER_ROOT" --json`,
		`"$TPT_BIN" preflight --repo-root "$PROVIDER_ROOT" --mode organization --json`,
		`"$TPT_BIN" retry --repo-root "$PROVIDER_ROOT" --failed --json`,
		`"$TPT_BIN" resume --repo-root "$PROVIDER_ROOT" --json`,
		`"$TPT_BIN" groups --repo-root "$PROVIDER_ROOT" --json`,
		`"$TPT_BIN" orphans --repo-root "$PROVIDER_ROOT" --mode "$RUN_MODE" --run-delta --json`,
		`"$TPT_BIN" sweep --repo-root "$PROVIDER_ROOT" --mode "$RUN_MODE" --run-delta --confirm`,
		`"$TPT_BIN" report --repo-root "$PROVIDER_ROOT" --md terraform-provider-tester-report.md`,
	}
	for _, snippet := range want {
		if !strings.Contains(skill, snippet) {
			t.Errorf("skill command is missing %q", snippet)
		}
	}
}

func TestCopilotFacingDocsAvoidStaleOperationalCommands(t *testing.T) {
	for _, path := range []string{
		"../README.md",
		"../.github/copilot-instructions.md",
		"../.github/skills/terraform-provider-tester/README.md",
		"../.github/skills/terraform-provider-tester/SKILL.md",
	} {
		content := readDoc(t, path)
		for _, stale := range []string{
			"./bin/terraform-provider-tester",
			"--env-file .pulsar.env",
			"go install github.com/github/terraform-provider-tester",
			"\nmake build\n",
			"Assume `make build`",
			`${TPT_BIN:-`,
			"set -euo pipefail",
		} {
			if strings.Contains(content, stale) {
				t.Errorf("%s contains stale operational command %q", path, stale)
			}
		}
	}
}

func fencedCodeBlocks(content, lang string) []string {
	fence := "```" + lang
	var blocks []string
	for {
		start := strings.Index(content, fence)
		if start < 0 {
			return blocks
		}
		content = content[start+len(fence):]
		if strings.HasPrefix(content, "\r\n") {
			content = content[2:]
		} else if strings.HasPrefix(content, "\n") {
			content = content[1:]
		}
		end := strings.Index(content, "```")
		if end < 0 {
			return blocks
		}
		blocks = append(blocks, content[:end])
		content = content[end+3:]
	}
}

func TestDocsReadmePrimaryWorkflowDefersToValidatedSkillBootstrap(t *testing.T) {
	readme := readDoc(t, "../README.md")
	for _, stale := range []string{
		`TPT_ROOT="$(cd -P "$HOME/.local/share/terraform-provider-tester" && pwd)"`,
		`make -C "$TPT_ROOT" build`,
		`TPT_BIN="$TPT_ROOT/bin/terraform-provider-tester"`,
	} {
		if strings.Contains(readme, stale) {
			t.Errorf("README reintroduces weaker cache-path bootstrap %q", stale)
		}
	}
	for _, snippet := range []string{
		"bundled skill",
		"validated portable bootstrap",
		".github/skills/terraform-provider-tester/SKILL.md",
		"one-time personal-skill setup",
	} {
		if !strings.Contains(readme, snippet) {
			t.Errorf("README primary workflow is missing validated-skill reference %q", snippet)
		}
	}
}

func TestSkillShBlocksStayPortablePOSIX(t *testing.T) {
	skill := readDoc(t, "../.github/skills/terraform-provider-tester/SKILL.md")
	blocks := fencedCodeBlocks(skill, "sh")
	if len(blocks) == 0 {
		t.Fatal("skill has no sh code blocks")
	}
	for i, block := range blocks {
		for _, bashOnly := range []string{
			"set -euo pipefail",
			"[[",
			"]]",
			"<(",
			">(",
			"function ",
			"=~",
		} {
			if strings.Contains(block, bashOnly) {
				t.Errorf("skill sh block %d contains Bash-only syntax %q", i+1, bashOnly)
			}
		}
		for _, line := range strings.Split(block, "\n") {
			trimmed := strings.TrimSpace(line)
			for _, bashCmd := range []string{"declare ", "local ", "mapfile ", "readarray ", "source "} {
				if strings.HasPrefix(trimmed, bashCmd) {
					t.Errorf("skill sh block %d contains Bash-only command %q", i+1, bashCmd)
				}
			}
		}
	}
}

func TestReadmesDocumentUserLevelSkillInstall(t *testing.T) {
	want := []string{
		`gh repo clone github/terraform-provider-tester ~/.local/share/terraform-provider-tester -- --depth 1`,
		`copilot skill add ~/.local/share/terraform-provider-tester/.github/skills/terraform-provider-tester/SKILL.md`,
	}
	for _, path := range []string{
		"../README.md",
		"../.github/skills/terraform-provider-tester/README.md",
	} {
		content := readDoc(t, path)
		for _, snippet := range want {
			if !strings.Contains(content, snippet) {
				t.Errorf("%s is missing user install recipe %q", path, snippet)
			}
		}
	}
}

// TestReadmeAndQuickstartDocumentIndividualGuidedRun verifies that both the
// README and the quickstart doc present the individual-mode guided e2e run -
// the recommended first credentialed check, since it needs only a personal
// account rather than a dedicated organization.
func TestReadmeAndQuickstartDocumentIndividualGuidedRun(t *testing.T) {
	const want = "e2e --mode individual --json"
	for _, path := range []string{"../README.md", "../docs/quickstart.md"} {
		content := readDoc(t, path)
		if !strings.Contains(content, want) {
			t.Errorf("%s does not document the individual-first guided run %q", path, want)
		}
	}
}

func TestLLMsDocumentGuidedE2EWorkflow(t *testing.T) {
	content := readDoc(t, "../llms.txt")
	for _, want := range []string{
		"`terraform-provider-tester e2e --mode individual --json`",
		"Guided `e2e` supports `anonymous`, `individual`, and `organization`.",
		"`--mode` defaults to `organization`.",
		"runs plan + preflight + tests + one failed-test retry",
		"does not run `sweep` automatically",
	} {
		if !strings.Contains(content, want) {
			t.Errorf("llms.txt does not document the guided e2e workflow: missing %q", want)
		}
	}
}

// TestCLIReferenceDocumentsPlanningFlags verifies the CLI reference documents
// the planning-related flags shared across preflight/run/e2e.
func TestCLIReferenceDocumentsPlanningFlags(t *testing.T) {
	content := readDoc(t, "../docs/cli-reference.md")
	for _, flag := range []string{"--mode", "--group", "--run", "--allow-unclassified"} {
		if !strings.Contains(content, "`"+flag+"`") {
			t.Errorf("cli-reference.md does not document the %q flag", flag)
		}
	}
}

// TestSkillMapsIndividualRequestToE2E verifies the skill's natural-language
// map resolves an individual-mode e2e request to the individual-mode
// resolved-binary command, alongside the existing organization row.
func TestSkillMapsIndividualRequestToE2E(t *testing.T) {
	skill := readDoc(t, "../.github/skills/terraform-provider-tester/SKILL.md")
	for _, want := range []string{
		`"$TPT_BIN" e2e --repo-root "$PROVIDER_ROOT" --mode individual --json`,
		"Run the individual e2e test",
	} {
		if !strings.Contains(skill, want) {
			t.Errorf("SKILL.md does not map an individual e2e request to the individual-mode command: missing %q", want)
		}
	}
}

// TestDocsDocumentNewPlanExcludedCapabilityEvents verifies the CLI reference
// and README document the plan/excluded/capability JSON event types (see
// cli/json.go's jsonPlanEvent, jsonExcludedEvent, jsonCapabilityEvent).
func TestDocsDocumentNewPlanExcludedCapabilityEvents(t *testing.T) {
	for _, path := range []string{"../docs/cli-reference.md", "../README.md"} {
		content := readDoc(t, path)
		for _, event := range []string{"`plan`", "`excluded`", "`capability`"} {
			if !strings.Contains(content, event) {
				t.Errorf("%s does not document the %s JSON event type", path, event)
			}
		}
	}
}

// TestDocsDocumentPlanlessStateMigration verifies architecture.md documents
// how a legacy, pre-state-v2 state file (state.Plan == nil) is handled: never
// synthesized, with planning counts explicitly unavailable (see
// engine.Load/engine.TerminalStateSummary).
func TestDocsDocumentPlanlessStateMigration(t *testing.T) {
	content := readDoc(t, "../docs/architecture.md")
	for _, want := range []string{
		"legacy state: planning counts unavailable",
		"does not synthesize a plan",
	} {
		if !strings.Contains(content, want) {
			t.Errorf("architecture.md does not document the planless-state migration behavior: missing %q", want)
		}
	}
}

// TestCIDocsDocumentSummaryParser verifies docs/ci.md exists and documents
// the exact jq parser that selects the last summary event from an individual
// guided e2e run, explicitly calling out that intermediate event types of
// unknown/unhandled type must be ignored.
func TestCIDocsDocumentSummaryParser(t *testing.T) {
	const path = "../docs/ci.md"
	if _, err := os.Stat(path); err != nil {
		t.Fatalf("%s does not exist: %v", path, err)
	}
	content := readDoc(t, path)
	for _, want := range []string{
		"e2e --mode individual --json",
		`jq -s 'map(select(.type == "summary")) | last'`,
		"ignore",
		"unknown",
	} {
		if !strings.Contains(content, want) {
			t.Errorf("docs/ci.md is missing %q", want)
		}
	}
}

// TestDocsDescribeWindowsAsUnverified verifies prerequisites.md keeps the
// required explicit Windows-unverified statement without overstating current
// cross-platform support or promising Windows validation.
func TestDocsDescribeWindowsAsUnverified(t *testing.T) {
	content := readDoc(t, "../docs/prerequisites.md")
	for _, want := range []string{
		"macOS and Linux are supported.",
		"Automated CI currently validates Linux.",
		"Windows is unverified",
		"not supported today",
		"cross-builds fail because lock identity uses POSIX-only",
		"syscall.Stat_t",
	} {
		if !strings.Contains(content, want) {
			t.Errorf("prerequisites.md does not describe the OS support matrix: missing %q", want)
		}
	}
}

// TestCLIReferenceDocumentsRetryResumeModeDefault verifies the CLI reference
// describes retry/resume `--mode` truthfully: it defaults to the persisted
// execution plan's mode (not the environment default `anonymous`), and an
// explicit value that disagrees with the persisted plan is rejected with exit
// 2. See resolveRetryResumeMode in cli/cli.go.
func TestCLIReferenceDocumentsRetryResumeModeDefault(t *testing.T) {
	content := readDoc(t, "../docs/cli-reference.md")
	for _, want := range []string{
		"`--mode` (default: the persisted execution plan's mode)",
		"conflicts with the persisted plan mode",
	} {
		if !strings.Contains(content, want) {
			t.Errorf("cli-reference.md does not document retry/resume mode resolution: missing %q", want)
		}
	}
	if strings.Contains(content, "Sets the test mode for the retry.") ||
		strings.Contains(content, "Sets the test mode for the resumed run.") {
		t.Error("cli-reference.md still describes retry/resume --mode as a free choice; it must defer to the persisted plan")
	}
}

// TestCopilotFacingDocsAvoidHardcodedRetryResumeMode verifies no doc or skill
// example pins a mode onto retry/resume. Those commands derive the mode from
// the persisted plan, so a hardcoded `--mode organization` example exits 2
// with a mode-mismatch error whenever the previous run used another mode.
func TestCopilotFacingDocsAvoidHardcodedRetryResumeMode(t *testing.T) {
	bad := regexp.MustCompile("(?:terraform-provider-tester|\\$TPT_BIN\")\\s+(?:retry|resume)\\b[^\n`]*--mode\\s+\\S+")
	for _, path := range harnessDocs(t) {
		content := readDoc(t, path)
		for _, line := range strings.Split(content, "\n") {
			if bad.MatchString(line) {
				t.Errorf("%s pins a mode onto retry/resume, which must default to the persisted plan: %q", path, strings.TrimSpace(line))
			}
		}
	}
}

// TestSkillDocumentsPersistedPlanModeForRetryResume verifies the skill tells
// Copilot to omit `--mode` for retry/resume so the persisted plan's mode is
// used, instead of guessing a mode that may not match the last run.
func TestSkillDocumentsPersistedPlanModeForRetryResume(t *testing.T) {
	skill := readDoc(t, "../.github/skills/terraform-provider-tester/SKILL.md")
	for _, want := range []string{
		`"$TPT_BIN" resume --repo-root "$PROVIDER_ROOT" --json`,
		`"$TPT_BIN" retry --repo-root "$PROVIDER_ROOT" --failed --json`,
		"persisted execution plan",
	} {
		if !strings.Contains(skill, want) {
			t.Errorf("SKILL.md does not use the persisted-plan mode for retry/resume: missing %q", want)
		}
	}
}

// TestCIDocsDescribeMultipleSummaryEvents verifies docs/ci.md tells the truth
// about the guided e2e NDJSON stream: a guided run embeds the `run` step's
// own summary and the orphan check's summary before its final `e2e` summary,
// so "exactly one summary" is wrong. The documented parser stays correct
// because the e2e summary is last, and it is identifiable by `command`.
func TestCIDocsDescribeMultipleSummaryEvents(t *testing.T) {
	content := readDoc(t, "../docs/ci.md")
	for _, want := range []string{
		`"command": "e2e"`,
		"more than one",
		"`triage`",
	} {
		if !strings.Contains(content, want) {
			t.Errorf("docs/ci.md does not describe the multi-summary NDJSON stream: missing %q", want)
		}
	}
	for _, banned := range []string{
		"exactly one final `summary` event",
		"there is exactly one per run",
	} {
		if strings.Contains(content, banned) {
			t.Errorf("docs/ci.md still claims a single summary event: %q", banned)
		}
	}
}

func TestDocsDescribeTerminalResumeSummaryAuthority(t *testing.T) {
	requirements := map[string][]string{
		"../docs/ci.md": {
			"`resume --json` may emit an inner run summary and a final resume summary",
			"For `resume`, the last `summary` object is authoritative",
			"Only the terminal resume summary is workflow-final",
		},
		"../docs/cli-reference.md": {
			"inner run summary and a final resume summary",
			"the last `summary` object is authoritative",
			"Only the terminal resume summary is workflow-final",
		},
	}
	for path, wants := range requirements {
		content := readDoc(t, path)
		for _, want := range wants {
			if !strings.Contains(content, want) {
				t.Errorf("%s does not document resume summary authority: missing %q", path, want)
			}
		}
	}
}

func TestDocsExplainRetryDoesNotRefreshOrphanAccounting(t *testing.T) {
	for _, path := range []string{
		"../docs/cli-reference.md",
		"../docs/troubleshooting.md",
	} {
		content := readDoc(t, path)
		for _, want := range []string{
			"`retry --failed` does not reopen or finalize orphan accounting",
			"persisted run delta remains the snapshot from the last guided finalization",
			"start or resume the guided workflow",
		} {
			if !strings.Contains(content, want) {
				t.Errorf("%s does not explain retry accounting limits: missing %q", path, want)
			}
		}
	}
}

// TestDocsDocumentPlanningFailureSummaries verifies the CLI reference states
// that `run --json` and `preflight --json` still end their NDJSON stream with
// a terminal summary when discovery or planning fails, and is explicit that a
// flag-parse usage error exits 2 before any JSON is written.
func TestDocsDocumentPlanningFailureSummaries(t *testing.T) {
	content := readDoc(t, "../docs/cli-reference.md")
	for _, want := range []string{
		"a discovery or planning failure still writes a final `summary`",
		"Usage errors are the exception and write no JSON at all",
		"flag parse error exits `2` before any JSON is written",
		"invalid `--run` regex exits `2` reported on stderr only",
	} {
		if !strings.Contains(content, want) {
			t.Errorf("cli-reference.md does not document planning-failure JSON summaries: missing %q", want)
		}
	}
}

// TestCLIReferenceDocumentsRunAllowUnclassified verifies the CLI reference
// documents `run --allow-unclassified` and no longer claims `run` always
// tolerates unclassified tests.
func TestCLIReferenceDocumentsRunAllowUnclassified(t *testing.T) {
	content := readDoc(t, "../docs/cli-reference.md")
	if !strings.Contains(content, "`--allow-unclassified` (default `false`)") {
		t.Error("cli-reference.md does not document --allow-unclassified as a default-false flag")
	}
	if strings.Contains(content, "there is no `--allow-unclassified` flag here because `run` never fails a plan closed") {
		t.Error("cli-reference.md still claims run has no --allow-unclassified flag")
	}
}

func TestCleanupDocsDescribePersistedModeAndRunDelta(t *testing.T) {
	cliRef := readDoc(t, "../docs/cli-reference.md")
	for _, want := range []string{
		"`--mode` (default: the persisted cleanup mode)",
		"`--run-delta` (default `false`)",
		"cleanup mode is unknown; pass --mode or start a plan-bearing run",
		"terraform-provider-tester orphans --mode individual --run-delta --json",
		"terraform-provider-tester sweep --mode individual --run-delta --confirm",
	} {
		if !strings.Contains(cliRef, want) {
			t.Errorf("cli-reference.md does not document cleanup mode protection: missing %q", want)
		}
	}

	readme := readDoc(t, "../README.md")
	if strings.Contains(readme, "`orphans` and `sweep` scope by `GH_TEST_AUTH_MODE`") {
		t.Error("README.md still claims cleanup commands scope directly by GH_TEST_AUTH_MODE")
	}
	if !strings.Contains(readme, "resolve cleanup mode from `--mode` first") {
		t.Error("README.md does not explain persisted cleanup mode resolution")
	}
}

func TestDocsCoverOrphanDeltaOperationsAndState(t *testing.T) {
	requirements := map[string][]string{
		"../README.md": {
			"`orphan_delta`",
			"Only `New` is this run's cleanup obligation.",
			"Guided runs do not run cleanup automatically.",
		},
		"../CHANGELOG.md": {
			"`orphan_delta`",
			"`--run-delta`",
		},
		"../docs/index.md": {
			"original baseline",
			"`orphan_delta`",
			"No cleanup runs automatically.",
		},
		"../docs/quickstart.md": {
			"terraform-provider-tester orphans --mode <mode> --run-delta",
			"terraform-provider-tester sweep --mode <mode> --run-delta --confirm",
			"preview is read-only",
		},
		"../docs/cli-reference.md": {
			"`orphan_delta`",
			"30-second final orphan check",
			"requires both `--run-delta` and `--confirm`",
			"all prefixed resources in the selected target kinds",
		},
		"../docs/troubleshooting.md": {
			"cleanup status `unknown`",
			"30-second",
			"original baseline",
		},
		"../docs/architecture.md": {
			"optional `orphans` fields",
			"orphan accounting unavailable",
			"reuses the original baseline",
			"30-second",
		},
		"../docs/private-validation.md": {
			"disposable individual account",
			"forced interruption",
			"terraform-provider-tester orphans --mode individual --run-delta",
			"terraform-provider-tester sweep --mode individual --run-delta --confirm",
			"explicit cleanup intent",
		},
		"../docs/ci.md": {
			"`orphan_delta`",
			"ignore unknown NDJSON event types",
		},
	}
	for path, wants := range requirements {
		content := readDoc(t, path)
		for _, want := range wants {
			if !strings.Contains(content, want) {
				t.Errorf("%s does not cover orphan-delta operations/state: missing %q", path, want)
			}
		}
	}
}

func TestDocsCoverOrphanBaselineAndExplicitCleanupSafety(t *testing.T) {
	cliRef := readDoc(t, "../docs/cli-reference.md")
	for _, want := range []string{
		"baseline capture is read-only",
		"baseline failure blocks execution",
		"resume reuses the original baseline",
		"Only `New` is attributed to the run.",
		"The run-delta preview is read-only.",
	} {
		if !strings.Contains(cliRef, want) {
			t.Errorf("cli-reference.md is missing orphan safety rule %q", want)
		}
	}
}

func TestDocsExplainStandaloneRunOrphanAccountingBoundary(t *testing.T) {
	cliRef := readDoc(t, "../docs/cli-reference.md")
	for _, want := range []string{
		"A standalone `run` starts a new logical run without guided orphan accounting.",
		"Use `e2e` for",
		"baseline capture, run-delta attribution, and finalization.",
		"`cleanup_status` means no accounting window opened before execution",
	} {
		if !strings.Contains(cliRef, want) {
			t.Errorf("cli-reference.md is missing standalone accounting guidance %q", want)
		}
	}
}

func TestDocsQualifyCredentialedAccountingAndUnknownRecovery(t *testing.T) {
	requirements := map[string][]string{
		"../README.md": {
			"Credentialed text and NDJSON output",
			"Anonymous mode emits no `orphan_delta` or cleanup command",
			"only when cleanup status supports an attributed delta",
			"`baseline-only`",
			"`cleanup_status` is authoritative",
		},
		"../docs/index.md": {
			"For a credentialed mode, the command",
			"Anonymous mode emits no `orphan_delta`",
			"only when cleanup status supports an attributed delta",
			"`baseline-only`",
			"`cleanup_status` is authoritative",
		},
		"../docs/cli-reference.md": {
			"Credentialed guided-run accounting",
			"`anonymous` emits no `orphan_delta` or cleanup command",
			"For credentialed guided runs, the final summary repeats",
			"numeric zero fields are schema placeholders, not proven counts",
			"`cleanup_status` is authoritative",
			"`preview_command` and `cleanup_command` stay empty",
		},
		"../docs/ci.md": {
			"Credentialed `e2e --json`",
			"Anonymous `e2e --json` omits `orphan_delta`",
			"only when cleanup status supports an attributed delta",
			"`baseline-only`",
			"`cleanup_status` is authoritative",
		},
		"../docs/quickstart.md": {
			"only when cleanup status supports an attributed delta",
			"`baseline-only`",
			"`cleanup_status` is authoritative",
		},
		"../docs/troubleshooting.md": {
			"Persisted `New` is unavailable while cleanup status is `unknown`.",
			"resume --json",
			"orphans --mode <mode> --json",
			"Do not pass `--run-delta` while cleanup status is `unknown`.",
		},
		"../.github/skills/terraform-provider-tester/SKILL.md": {
			"When cleanup status is `unknown`, persisted `New` is unavailable.",
			"Do not use `--run-delta` until final accounting is complete.",
			"`orphan_delta`: credentialed guided-run",
			"numeric zero fields are schema placeholders, not proven counts",
			"`cleanup_status` is authoritative",
		},
		"../.github/copilot-instructions.md": {
			"When cleanup status is `unknown`, persisted `New` is unavailable.",
			"Do not use `--run-delta` until final accounting is complete.",
		},
	}
	for path, wants := range requirements {
		content := readDoc(t, path)
		for _, want := range wants {
			if !strings.Contains(content, want) {
				t.Errorf("%s is missing accounting safety guidance %q", path, want)
			}
		}
	}
}

func TestPrivateValidationVerifiesLiveOrphansAfterHistoricalRunDeltaSweep(t *testing.T) {
	content := readDoc(t, "../docs/private-validation.md")
	for _, want := range []string{
		"terraform-provider-tester orphans --mode individual --json",
		"persisted run-delta remains historical",
		"sweep does not clear it",
	} {
		if !strings.Contains(content, want) {
			t.Errorf("private-validation.md is missing post-sweep verification guidance %q", want)
		}
	}
}

func TestDocsSkillPreviewsBeforeExplicitCleanup(t *testing.T) {
	for _, path := range []string{
		"../.github/skills/terraform-provider-tester/SKILL.md",
		"../.github/copilot-instructions.md",
	} {
		content := readDoc(t, path)
		for _, want := range []string{
			`"$TPT_BIN" orphans --repo-root "$PROVIDER_ROOT" --mode "$RUN_MODE" --run-delta --json`,
			`"$TPT_BIN" sweep --repo-root "$PROVIDER_ROOT" --mode "$RUN_MODE" --run-delta --confirm`,
			"Preview before suggesting cleanup.",
			"Never execute cleanup without explicit user intent.",
		} {
			if !strings.Contains(content, want) {
				t.Errorf("%s is missing explicit cleanup guidance %q", path, want)
			}
		}
	}
}

// individualPlanFromCatalog builds the full individual-mode execution plan
// from the REAL provider catalog and the pinned acceptance-test fixture, so
// the documentation assertions below are tied to code rather than to a
// hand-maintained list. It is the same aggregate
// provider/github's TestIndividualPlanAggregatesRealCatalogRequirements pins.
func individualPlanFromCatalog(t *testing.T) engine.ExecutionPlan {
	t.Helper()
	data, err := os.ReadFile("../engine/testdata/list_github.txt")
	if err != nil {
		t.Fatalf("read pinned test catalog: %v", err)
	}
	var names []string
	for _, line := range strings.Split(string(data), "\n") {
		if strings.HasPrefix(line, "TestAcc") {
			names = append(names, line)
		}
	}
	plan, err := engine.BuildExecutionPlan(names, githubprovider.New(), engine.PlanOptions{Mode: "individual"})
	if err != nil {
		t.Fatalf("build individual plan: %v", err)
	}
	return plan
}

// TestDocsStateExactIndividualModePrerequisites ties the individual-mode
// prerequisites documentation to the catalog: every classic scope the real
// individual plan aggregates, and the template repository its
// template-repository capability probe needs, must be named in
// docs/prerequisites.md. The docs previously claimed individual mode "needs
// only a personal account and PAT", which understated six aggregate scopes
// and a required template repository.
func TestDocsStateExactIndividualModePrerequisites(t *testing.T) {
	plan := individualPlanFromCatalog(t)
	content := readDoc(t, "../docs/prerequisites.md")

	individual := docSection(content, "### Individual")
	if individual == "" {
		t.Fatal("prerequisites.md has no '### Individual' section")
	}
	for _, scope := range plan.Scopes {
		if !strings.Contains(individual, "`"+scope+"`") {
			t.Errorf("prerequisites.md's Individual section does not name the aggregate classic scope %q", scope)
		}
	}
	if !strings.Contains(individual, "GH_TEST_ORG_TEMPLATE_REPOSITORY") {
		t.Error("prerequisites.md's Individual section does not name GH_TEST_ORG_TEMPLATE_REPOSITORY, which the template-repository capability probe requires")
	}
}

// docSection returns the body of the Markdown section starting at heading, up
// to the next heading of the same or a higher level (fewer '#'). Returns ""
// when heading is absent.
func docSection(content, heading string) string {
	idx := strings.Index(content, heading)
	if idx < 0 {
		return ""
	}
	level := len(heading) - len(strings.TrimLeft(heading, "#"))
	rest := content[idx+len(heading):]
	for offset := 0; ; {
		next := strings.Index(rest[offset:], "\n#")
		if next < 0 {
			return rest
		}
		start := offset + next + 1
		line := rest[start:]
		if end := strings.IndexByte(line, '\n'); end >= 0 {
			line = line[:end]
		}
		if depth := len(line) - len(strings.TrimLeft(line, "#")); depth <= level {
			return rest[:start]
		}
		offset = start + 1
	}
}

// TestDocsDoNotUnderstateIndividualModeSetup verifies no operator- or
// Copilot-facing doc still claims individual mode needs only a personal
// account and a PAT. It remains the lowest-setup credentialed mode and needs
// no dedicated organization, but it does need specific classic scopes and a
// personal-account template repository (see
// TestDocsStateExactIndividualModePrerequisites).
func TestDocsDoNotUnderstateIndividualModeSetup(t *testing.T) {
	understatements := []string{
		"only a personal account and PAT",
		"only a personal account and a PAT",
		"just a personal account and PAT",
		"just a personal account and a PAT",
		"needs only a personal account",
	}
	for _, path := range harnessDocs(t) {
		content := readDoc(t, path)
		for _, phrase := range understatements {
			if strings.Contains(content, phrase) {
				t.Errorf("%s claims individual mode %q; state the real prerequisites instead (see docs/prerequisites.md)", path, phrase)
			}
		}
	}
}

// TestDocsRecommendingIndividualModePointAtPrerequisites verifies that every place
// recommending individual mode as the first credentialed check also points
// the reader at the exact prerequisites, so the recommendation can never be
// read as "no setup required".
func TestDocsRecommendingIndividualModePointAtPrerequisites(t *testing.T) {
	for _, path := range []string{
		"../README.md",
		"../docs/index.md",
		"../docs/quickstart.md",
		"../docs/ci.md",
		"../docs/cli-reference.md",
		"../docs/test-modes.md",
		"../llms.txt",
		"../.github/skills/terraform-provider-tester/SKILL.md",
	} {
		content := readDoc(t, path)
		if !strings.Contains(content, "no dedicated organization") {
			t.Errorf("%s does not say individual mode needs no dedicated organization", path)
		}
		if !strings.Contains(content, "GH_TEST_ORG_TEMPLATE_REPOSITORY") {
			t.Errorf("%s recommends individual mode without naming GH_TEST_ORG_TEMPLATE_REPOSITORY", path)
		}
	}
}

// TestDocsChangelogDescribesRunAllowUnclassifiedAccurately verifies the changelog
// no longer claims `run` always tolerates an unclassified test. `run` gained
// the same default-false `--allow-unclassified` flag as `preflight` and `e2e`
// (see runRunContext), and the same entry's "Changed" section already says so.
func TestDocsChangelogDescribesRunAllowUnclassifiedAccurately(t *testing.T) {
	content := readDoc(t, "../CHANGELOG.md")
	if strings.Contains(content, "`run` always allows it") {
		t.Error("CHANGELOG.md still claims `run` always allows an unclassified test")
	}
	if !strings.Contains(content, "`run` requires the same explicit `--allow-unclassified`") {
		t.Error("CHANGELOG.md does not say `run` requires the explicit --allow-unclassified flag too")
	}
}

// TestDocsScopeInterruptionPersistenceToGuidedE2E verifies the operator- and
// Copilot-facing docs describe the only installed graceful signal path without
// promising resumable state for direct run, retry, or resume interruption.
func TestDocsScopeInterruptionPersistenceToGuidedE2E(t *testing.T) {
	for _, path := range []string{
		"../docs/troubleshooting.md",
		"../docs/cli-reference.md",
		"../docs/architecture.md",
		"../docs/private-validation.md",
		"../.github/skills/terraform-provider-tester/SKILL.md",
	} {
		content := readDoc(t, path)
		for _, want := range []string{
			"Only guided `e2e` installs the graceful signal handler.",
			"A graceful first interruption of guided `e2e` can persist partial results",
			"final orphan accounting",
			"original baseline",
			"SIGKILL, or interruption of `run`, `retry`, or `resume` outside that guided signal path, may leave no resumable plan.",
			"start a new run or e2e",
		} {
			if !strings.Contains(content, want) {
				t.Errorf("%s does not accurately scope interruption persistence: missing %q", path, want)
			}
		}
		for _, overclaim := range []string{
			"Graceful cancellation saves partial results",
			"Graceful cancellation persists partial results",
			"A single graceful interruption persists partial results",
			"second interruption before the save",
		} {
			if strings.Contains(content, overclaim) {
				t.Errorf("%s overclaims interruption persistence with %q", path, overclaim)
			}
		}
	}
}
