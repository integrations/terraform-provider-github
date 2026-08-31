# GH/TF Console Identity Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Replace the generic TESTER dashboard identity with GH/TF Provider Acceptance while preserving every CLI, state, safety, and compatibility contract.

**Architecture:** Keep `AppName = "Terraform Provider Tester"` as the automation-facing product name. Add dashboard-only `ConsoleMark`, `ConsoleName`, and `ConsoleTitle` constants, then route header, splash, and banner rendering through them. Add a dedicated Terraform-purple semantic token for identity and selection signals while GitHub blue remains the action color.

**Tech Stack:** Go 1.26, Bubble Tea, Lip Gloss, ANSI-aware terminal rendering, golden tests.

## Global Constraints

- Preserve `terraform-provider-tester` as repository, module, binary, and command.
- Preserve `.pulsar.env`, `.pulsar-state.json`, `.pulsar-failures`, `PULSAR_*`, `pulsar-fp-v1`, `pulsar/pre-run`, and the known-issue marker.
- Preserve exact confirmations `SWEEP <owner>` and `FILE <short-fingerprint>`.
- Do not change commands, keys, acceptance-test execution, persistence, or destructive behavior.
- Terraform purple appears only on the TF mark, active-tab bottom rule, and selection cursor.
- GitHub blue remains the action, link, command, and identifier color.
- Preserve ASCII/NO_COLOR output and responsive width guarantees.

---

### Task 1: Pin the GH/TF identity contract

**Files:**
- Modify: `tui/branding_test.go`
- Modify: `tui/wordmark_test.go`
- Modify: `tui/version_render_test.go`

**Interfaces:**
- Consumes: current `renderHeader`, `renderIntro`, `renderBanner`, and `renderWordmark`.
- Produces: failing tests for `ConsoleMark`, `ConsoleTitle`, GH/TF glyphs, and removal of the visible TESTER mark.

- [ ] **Step 1: Change the wordmark expectation**

Replace `TestWordmarkUsesTesterBrand` with a GH/TF contract:

```go
func TestWordmarkUsesGHTFBrand(t *testing.T) {
	if wordmarkText != "GH/TF" {
		t.Fatalf("wordmarkText = %q, want GH/TF", wordmarkText)
	}
	got := renderWordmark(1, true)
	if strings.Contains(got, "TESTER") || strings.Contains(got, "PULSAR") {
		t.Fatalf("resolved wordmark contains a retired dashboard brand:\n%s", got)
	}
}
```

- [ ] **Step 2: Change header, intro, and banner expectations**

Require the full layouts to contain `ConsoleTitle`, compact layouts to contain
`ConsoleMark`, and every dashboard surface to omit the standalone `TESTER`
wordmark. Keep existing version, provider, mode, owner, vendor, license, and
tagline assertions.

- [ ] **Step 3: Add style contract tests**

Assert:

```go
if !Styles.TabActive.GetBorderBottom() {
	t.Fatal("active tab must use a bottom rule")
}
if Styles.TabActive.GetBackground() != lipgloss.NoColor{} {
	t.Fatal("active tab must not use a filled background")
}
if Styles.SelectedRow.GetBackground() == TerraformPurple {
	t.Fatal("selected rows must not use a Terraform-purple background")
}
```

Use the actual `TerminalColor` zero-value comparison supported by the installed
Lip Gloss version.

- [ ] **Step 4: Run the focused tests and verify RED**

Run:

```bash
go test ./tui -run 'TestWordmarkUsesGHTFBrand|TestHeader|TestBanner|TestIntro|TestConsoleStyle' -count=1
```

Expected: failures mention missing `GH/TF`, visible `TESTER`, and the active-tab
background.

- [ ] **Step 5: Commit the red tests**

```bash
git add tui/branding_test.go tui/wordmark_test.go tui/version_render_test.go
git commit -m "test(tui): define GH/TF console identity"
```

### Task 2: Add dashboard-only identity and color semantics

**Files:**
- Modify: `tui/branding.go`
- Modify: `tui/theme.go`
- Modify: `tui/styles.go`
- Modify: `tui/chrome.go`
- Modify: `tui/view.go`
- Modify: `tui/triage.go`

**Interfaces:**
- Produces:
  - `ConsoleMark = "GH/TF"`
  - `ConsoleName = "Provider Acceptance"`
  - `ConsoleTitle = ConsoleMark + " " + ConsoleName`
  - `TerraformPurple` and `SelectedSurface` adaptive colors.

- [ ] **Step 1: Add dashboard identity constants**

```go
const (
	AppName     = "Terraform Provider Tester"
	ConsoleMark = "GH/TF"
	ConsoleName = "Provider Acceptance"
	ConsoleTitle = ConsoleMark + " " + ConsoleName
)
```

Document that `AppName` remains the CLI-facing name and `ConsoleTitle` is visual
dashboard identity only.

- [ ] **Step 2: Add semantic color tokens**

```go
TerraformPurple = lipgloss.AdaptiveColor{Light: "#844FBA", Dark: "#A16BE8"}
SelectedSurface = lipgloss.AdaptiveColor{Light: "#f6f8fa", Dark: "#21262d"}
```

Do not repurpose `Accent`, `Success`, `Danger`, `Attention`, or `Flaky`.

- [ ] **Step 3: Replace filled tabs and rows**

Make the active tab bold with a Terraform-purple bottom border and no background.
Give inactive tabs an equal-height blank bottom border. Make selected rows use
`SelectedSurface`, not Terraform purple.

- [ ] **Step 4: Render the console title compositionally**

Add a helper that renders:

- `GH` in GitHub blue,
- `/` muted,
- `TF` in Terraform purple,
- `Provider Acceptance` in terminal-default bold text.

Use the compact form in narrow headers and the full title in standard/wide
headers. Keep version and chips unchanged.

- [ ] **Step 5: Change selected cursors from blue to Terraform purple**

Update group, test, triage, and picker cursor glyphs only. Leave action labels,
commands, links, and identifiers on `Accent`.

- [ ] **Step 6: Run focused chrome tests**

```bash
go test ./tui -run 'TestHeader|TestActiveTab|TestConsoleStyle|TestMissionPanels|TestRenderGroupsBarsAlign' -count=1
```

Expected: PASS.

- [ ] **Step 7: Commit**

```bash
git add tui/branding.go tui/theme.go tui/styles.go tui/chrome.go tui/view.go tui/triage.go
git commit -m "feat(tui): adopt GH/TF console identity"
```

### Task 3: Replace the splash and welcome mark

**Files:**
- Modify: `tui/wordmark.go`
- Modify: `tui/intro.go`
- Modify: `tui/banner.go`
- Modify: `tui/testdata/wordmark_full.golden`
- Modify: `tui/testdata/wordmark_full_tty.golden`
- Modify: `tui/testdata/wordmark_mid.golden`

**Interfaces:**
- Produces: a deterministic five-glyph `GH/TF` dot-matrix wordmark and stacked
  GH/TF welcome-card mark.

- [ ] **Step 1: Add 5x3 glyphs**

Use:

```go
'G': {"###", "#  ", "# #", "# #", "###"},
'H': {"# #", "# #", "###", "# #", "# #"},
'/': {"  #", "  #", " # ", "#  ", "#  "},
'T': {"###", " # ", " # ", " # ", " # "},
'F': {"###", "#  ", "## ", "#  ", "#  "},
```

Set `wordmarkText = "GH/TF"` and calculate `wordmarkCellWidth` from
`len(wordmarkText)` rather than a hard-coded six-letter count.

- [ ] **Step 2: Color the final mark by identity role**

At the solid stage, render G/H with `Accent`, slash with `Muted`, and T/F with
`TerraformPurple`. Preserve existing dither stages and deterministic hashing.

- [ ] **Step 3: Update intro copy**

Render:

```text
GH/TF Provider Acceptance <version> · terraform-provider-tester
```

Keep the tagline and skip hint.

- [ ] **Step 4: Replace the beacon**

Rename `testerMark` to `ghTFMark` and render a five-line stacked GH/TF terminal
mark. The ASCII variant must use only ASCII; the TTY variant may use box
drawing but no Nerd Font glyphs.

- [ ] **Step 5: Verify green, then regenerate goldens**

```bash
UPDATE_GOLDEN=1 go test ./tui -run 'TestWordmark|TestBanner|TestIntro' -count=1
go test ./tui -run 'TestWordmark|TestBanner|TestIntro' -count=1
```

Expected: both commands PASS.

- [ ] **Step 6: Commit**

```bash
git add tui/wordmark.go tui/intro.go tui/banner.go tui/wordmark_test.go tui/branding_test.go tui/version_render_test.go tui/testdata/wordmark_*.golden
git commit -m "feat(tui): render GH/TF splash and mark"
```

### Task 4: Refresh documentation visuals and copy

**Files:**
- Modify: `README.md`
- Modify: `CHANGELOG.md`
- Modify: `docs/index.md`
- Modify: `docs/quickstart.md`
- Modify: `docs/troubleshooting.md`
- Modify: `tui/docs_screenshot_test.go`
- Modify: `docs/images/intro.png`
- Modify: `docs/images/preflight.png`
- Modify: `docs/images/groups.png`
- Modify: `docs/images/run.png`
- Modify: `docs/images/triage.png`

**Interfaces:**
- Consumes: deterministic screenshot generator and promotion tests.
- Produces: documentation that names GH/TF as the dashboard identity without
  renaming the tool.

- [ ] **Step 1: Update visible identity copy**

Describe GH/TF Provider Acceptance as the optional dashboard inside Terraform
Provider Tester. Remove statements that call TESTER the dashboard wordmark.

- [ ] **Step 2: Update screenshot fixture expectations**

Replace synthetic `TESTER` header/wordmark assertions with `GH/TF` and
`Provider Acceptance`. Preserve the five-screen completeness and corruption
guards.

- [ ] **Step 3: Generate and promote screenshots**

Run the repository's documented screenshot generation and promotion tests.
Verify all five PNG files are valid and changed intentionally.

- [ ] **Step 4: Run docs checks**

```bash
make checkdocs
git diff --check -- '*.go' '*.md' 'docs/**' 'README.md' 'CHANGELOG.md'
```

Expected: PASS.

- [ ] **Step 5: Commit**

```bash
git add README.md CHANGELOG.md docs/ tui/docs_screenshot_test.go
git commit -m "docs: show GH/TF provider acceptance console"
```

### Task 5: Verify and present the rebuilt console

**Files:**
- No new production files.
- Evidence: session artifact directory only.

- [ ] **Step 1: Run complete local gates**

```bash
go vet ./...
make checkdocs
go test ./... -count=1
go test -race ./cli ./tui -count=1
make build
```

Expected: all commands exit 0.

- [ ] **Step 2: Run isolated pseudo-TTY smoke**

Exercise intro, all four tabs, compact/wide layouts, help, redaction canaries,
exact confirmations, report export, and cleanup. Do not run credentialed
acceptance tests, sweep, or issue filing.

- [ ] **Step 3: Relaunch the local preview**

Launch `bin/terraform-provider-tester` against the provider worktree in
anonymous mode with credentials unset and an isolated empty env file.

- [ ] **Step 4: Human visual gate**

Confirm:

- GH/TF wordmark is legible.
- Terraform purple is a line/cursor signal, not a filled block.
- GitHub blue still identifies actions.
- No visible TESTER or PULSAR wordmark remains.
- Compact and 120-column layouts remain readable.

- [ ] **Step 5: Commit any evidence-only test adjustments separately**

Do not merge or tag. Resume the v0.3.0 release path only after human approval.
