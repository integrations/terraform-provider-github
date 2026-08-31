# Terraform Provider Tester TUI Full Parity Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Give the optional dashboard full, safe access to current triage, known-issue, report, orphan, sweep, and selected issue-filing workflows while completing the Terraform Provider Tester rebrand.

**Architecture:** Keep `tui/` as a pure Bubble Tea state machine. Add focused internal services under `cli/` so CLI handlers and the dashboard producer reuse state, redaction, GitHub, report, and cleanup behavior. Persist reports and known-issue cache under the OS user config directory, and guard every asynchronous operation with one producer-side gate.

**Tech Stack:** Go 1.26, Bubble Tea, Bubbles, Lip Gloss, Terraform Provider Tester engine/provider packages, GitHub REST client, YAML, Go golden tests, and the existing Makefile.

## Global Constraints

- CLI and NDJSON remain the primary automation interfaces; the TUI remains optional.
- Render `TESTER` in the startup animation and `Terraform Provider Tester` in product text.
- Preserve `.pulsar.env`, `.pulsar-state.json`, `.pulsar-failures`, `PULSAR_*`, `pulsar-fp-v1`, and known-issue markers.
- Keep manual retry and resume behavior; add zero automatic retries.
- File only the selected eligible failure, never a batch.
- Require exact, case-sensitive `FILE <short-fingerprint>` and `SWEEP <owner>` phrases.
- Re-list orphans before sweep and abort when the normalized resource set changes.
- Use the persistent known-issue cache during routine runs; only explicit `K` sync and issue filing may call GitHub.
- Store reports and the known-issue cache outside provider worktrees.
- Never put secrets or raw unredacted logs in TUI messages, state, reports, previews, errors, fixtures, or screenshots.
- Do not change the `provider.Provider` interface; the synchronized environment overlay already writes through to the process environment.
- Use only existing dependencies and tools.
- Work on `tui-full-parity`, stacked on `persist-env-config`; do not expand PR #17.
- Do not run credentialed acceptance tests without a dedicated test organization and scoped credentials.
- In this environment, export `PATH="/opt/homebrew/bin:$PATH"` before Go commands.

---

## Context Map

### Files to Modify

| File | Purpose | Changes Needed |
| --- | --- | --- |
| `tui/wordmark.go`, `tui/intro.go`, `tui/banner.go` | Startup and banner branding | Render `TESTER`; rename visible/internal tester-mark copy |
| `tui/model.go`, `tui/messages.go`, `tui/intents.go`, `tui/keys.go` | Pure dashboard state and protocol | Add Triage state, operation state, intents, messages, and `t`/`K`/`i` bindings |
| `tui/update.go`, `tui/view.go` | Reducer and rendering | Add fourth tab, ordering, overlays, confirmations, and result handling |
| `tui/triage.go`, `tui/confirm.go`, `tui/cleanup.go` | Focused new UI units | Render triage, issue preview, orphan list, and exact-phrase confirmations |
| `engine/atomic.go`, `engine/report.go`, `engine/known_issues.go` | Persistent artifacts | Write reports and cache atomically |
| `cli/triage_service.go` | Shared state triage | Reconstruct, match, and persist failures with configurable live/offline options |
| `cli/known_issues_service.go` | Known-issue cache sync | Fetch entries and atomically persist the default cache |
| `cli/report_service.go` | TUI report export | Reconstruct last run and export timestamped Markdown and HTML |
| `cli/cleanup_service.go` | Orphan and sweep orchestration | Normalize snapshots, re-list, confirm, sweep, and return residuals |
| `cli/issue_service.go` | Selected issue workflow | Build redacted preview, live-dedup, file one issue, and persist |
| `cli/dashboard.go`, `cli/cli.go`, `cli/triage.go`, `cli/known_issues.go` | Producers and CLI adapters | Wire services, operation gate, state updates, and dependency factories |
| `fakeprovider/fake.go` | Provider test double | Add orphan error injection |
| `README.md`, `CHANGELOG.md`, `docs/*.md`, `docs/images/*` | User documentation | Document functional keys, paths, safety, architecture, and synthetic screens |

### Dependencies

| File | Relationship |
| --- | --- |
| `engine/state.go` | Supplies atomic state save, lock acquisition, persisted results, and failures |
| `engine/triage.go`, `engine/signature.go`, `engine/retry.go` | Supplies classification, fingerprints, and retry history |
| `provider/types.go`, `provider/provider.go` | Supplies `Resource`, `SweepOpts`, and provider operations |
| `provider/github/issues.go` | Supplies known-issue listing, draft construction, dedup, and creation |
| `cli/env_overlay.go` | Keeps env-file values, TUI edits, process env, providers, and GitHub clients synchronized |
| `internal/redact/*` | Redacts all output before persistence or TUI delivery |

### Test Files

| Test | Coverage |
| --- | --- |
| `tui/wordmark_test.go`, `tui/intro_test.go`, `tui/branding_test.go` | Rebranding |
| `tui/view_test.go`, `tui/update_test.go`, `tui/triage_test.go`, `tui/cleanup_test.go`, `tui/confirm_test.go` | Tabs, rendering, ordering, overlays, messages, and confirmations |
| `engine/report_test.go`, `engine/known_issues_test.go`, `engine/atomic_test.go` | Atomic persistent writes and redaction |
| `cli/triage_test.go`, `cli/known_issues_test.go`, `cli/report_service_test.go`, `cli/cleanup_service_test.go`, `cli/issue_service_test.go` | Shared services and CLI compatibility |
| `cli/dashboard_test.go` | Producer operation gate, startup state, post-run triage, and all intents |
| `provider/github/sweep_test.go`, `provider/github/issues_test.go` | Existing provider and issue client behavior |

### Reference Patterns

| File | Pattern |
| --- | --- |
| `engine/state.go:81-111` | Same-directory temp file, close, rename, and cleanup |
| `tui/update.go:137-240` | Overlay key interception and text-input forwarding |
| `tui/view_test.go:55-75` | Golden update and comparison |
| `cli/dashboard.go:226-298` | Lock-before-load, streamed redaction, and state persistence |
| `cli/triage.go:492-572` | Known-issue matching, dedup, and filing rules |
| `cli/dashboard_test.go` | Captured `send` callback and fake runner |
| `cli/triage_test.go` | Fake registry/filer and secret canaries |

### Risk Assessment

- [x] Internal protocol changes: new TUI messages and intents.
- [x] Persistent artifact changes: atomic write implementation and new user-config paths.
- [x] External side effects: sweep and GitHub issue creation require dual validation.
- [ ] Public CLI flag or output changes.
- [ ] Provider interface changes.
- [ ] Database migrations.

---

### Task 1: Complete Rebranding and Add the Fourth Tab Shell

**Files:**
- Modify: `tui/wordmark.go`
- Modify: `tui/intro.go`
- Modify: `tui/banner.go`
- Modify: `tui/model.go`
- Modify: `tui/view.go`
- Modify: `tui/update.go`
- Modify: `tui/wordmark_test.go`
- Modify: `tui/update_test.go`
- Modify: `tui/branding_test.go`
- Modify: `tui/testdata/*.golden`

**Interfaces:**
- Consumes: existing `renderWordmark`, `section`, `Model.View`, and golden helpers.
- Produces: `sectionTriage`, `numSections == 4`, and a visible `TESTER` startup mark for later tasks.

- [ ] **Step 1: Run the branch baseline**

Run:

```bash
export PATH="/opt/homebrew/bin:$PATH"
make test
make vet
make checkdocs
```

Expected: all commands exit `0`.

- [ ] **Step 2: Write failing branding and tab tests**

Add these assertions to `tui/wordmark_test.go` and `tui/update_test.go`:

```go
func TestWordmarkUsesTesterBrand(t *testing.T) {
	if wordmarkText != "TESTER" {
		t.Fatalf("wordmarkText = %q, want TESTER", wordmarkText)
	}
	got := renderWordmark(1, true)
	if strings.Contains(got, "PULSAR") {
		t.Fatalf("resolved wordmark contains legacy brand:\n%s", got)
	}
	compareOrUpdate(t, "wordmark_full", got)
}

func TestUpdateTabCyclesFourSections(t *testing.T) {
	m := newTestModel()
	want := []section{sectionGroups, sectionRun, sectionTriage, sectionPreflight}
	for i, sectionWant := range want {
		next, _ := m.Update(tea.KeyMsg{Type: tea.KeyTab})
		m = next.(Model)
		if m.section != sectionWant {
			t.Fatalf("step %d section = %d, want %d", i, m.section, sectionWant)
		}
	}
}
```

Update the tab-name assertion in `tui/view_test.go` to require `Triage`.

- [ ] **Step 3: Run the focused tests and verify failure**

Run:

```bash
go test ./tui -run 'TestWordmarkUsesTesterBrand|TestUpdateTabCyclesFourSections' -count=1
```

Expected: FAIL because `wordmarkText` is `PULSAR` and `numSections` is `3`.

- [ ] **Step 4: Implement the TESTER glyphs and fourth section**

Use this exact section declaration in `tui/model.go`:

```go
const (
	sectionPreflight section = iota
	sectionGroups
	sectionRun
	sectionTriage
	numSections = 4
)
```

Add these glyphs and change the word in `tui/wordmark.go`:

```go
'T': {
	"###",
	" # ",
	" # ",
	" # ",
	" # ",
},
'E': {
	"###",
	"#  ",
	"###",
	"#  ",
	"###",
},

const wordmarkText = "TESTER"
```

Keep only the `S`, `T`, `E`, and `R` glyph entries. Rename `pulsarMark` to
`testerMark`, update its call site, and change visible-brand comments in
`tui/intro.go`, `tui/banner.go`, and `tui/wordmark.go`.

Update the tab array and body switch in `tui/view.go`:

```go
names := [numSections]string{"Preflight", "Groups", "Run", "Triage"}
```

```go
case sectionTriage:
	return Styles.Body.Render("  no triage results")
```

Add `case sectionTriage: return 0` to `itemCount` in `tui/update.go`.

- [ ] **Step 5: Regenerate and inspect golden fixtures**

Run:

```bash
go test ./tui -run 'TestView|TestWordmark' -update -count=1
go test ./tui -count=1
git diff -- tui/testdata
```

Expected: goldens show `TESTER` and four tabs; `go test ./tui` passes.

- [ ] **Step 6: Commit the rebrand and tab shell**

```bash
git add tui
git commit -m "feat(tui): complete tester rebrand" \
  -m "Co-authored-by: Copilot App <223556219+Copilot@users.noreply.github.com>"
```

---

### Task 2: Add Triage State, Rendering, Keys, and Group Ordering

**Files:**
- Create: `engine/issue_eligibility.go`
- Create: `engine/issue_eligibility_test.go`
- Modify: `cli/triage.go`
- Modify: `cli/triage_test.go`
- Create: `tui/triage.go`
- Create: `tui/triage_test.go`
- Modify: `tui/model.go`
- Modify: `tui/messages.go`
- Modify: `tui/intents.go`
- Modify: `tui/keys.go`
- Modify: `tui/update.go`
- Modify: `tui/view.go`
- Modify: `tui/update_test.go`
- Modify: `tui/testdata/triage_empty.golden`
- Modify: `tui/testdata/triage_list.golden`
- Modify: `tui/testdata/triage_detail.golden`

**Interfaces:**
- Consumes: `engine.PersistFailure`, `sectionTriage`, existing `exec` intent seam.
- Produces: `TriageLoadedMsg`, `RefreshTriageIntent`, `SyncKnownIssuesIntent`, `PreviewFileIssueIntent`, triage list/detail state, and deterministic group ordering.

- [ ] **Step 1: Write failing message, key, ordering, and render tests**

Create `tui/triage_test.go` with fixtures that use all four classifications:

```go
func sampleFailures() []engine.PersistFailure {
	return []engine.PersistFailure{
		{Test: "TestAccKnown", Classification: engine.ClassificationReal, Class: "api", Fingerprint: "sha256:aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", ShortFingerprint: "aaaaaaaaaaaaaaaa", Attempts: 1, KnownIssue: 42, IssueAction: "known"},
		{Test: "TestAccFlake", Classification: engine.ClassificationFlakeConfirmed, Class: "rate-limit", Fingerprint: "sha256:bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb", ShortFingerprint: "bbbbbbbbbbbbbbbb", Attempts: 2},
		{Test: "TestAccReal", Classification: engine.ClassificationReal, Class: "assertion", Fingerprint: "sha256:cccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccc", ShortFingerprint: "cccccccccccccccc", Attempts: 1, Retryable: false},
		{Test: "TestAccUnstable", Classification: engine.ClassificationRealUnstable, Class: "conflict", Fingerprint: "sha256:dddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddd", ShortFingerprint: "dddddddddddddddd", Attempts: 2},
	}
}

func TestTriageLoadedRendersListAndDetail(t *testing.T) {
	m := fixedModel()
	m.section = sectionTriage
	next, _ := m.Update(TriageLoadedMsg{Failures: sampleFailures(), CacheAvailable: true})
	m = next.(Model)
	compareOrUpdate(t, "triage_list", m.View())

	next, _ = m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	m = next.(Model)
	compareOrUpdate(t, "triage_detail", m.View())
}
```

Add update tests that prove:

```go
// t emits RefreshTriageIntent only in sectionTriage.
// K emits SyncKnownIssuesIntent only in sectionTriage.
// i emits PreviewFileIssueIntent only for an eligible selected failure.
// f sets failuresFirst=true and reorders groups/tests.
// a restores the exact discovery order.
```

Add `TestEligibleForIssueFiling` as an engine test matrix proving that only
`real` and `real-unstable` failures with a valid `sha256:` fingerprint, its
matching 16-hex short fingerprint, and no known issue/action are eligible.
Include empty, truncated, non-hex, and mismatched fingerprint cases. Add a
`TestShouldFileIssueUsesEngineEligibility` CLI regression test proving the
existing filing path uses the same predicate.

- [ ] **Step 2: Run focused tests and verify failure**

Run:

```bash
go test ./engine ./cli ./tui -run 'TestEligibleForIssueFiling|TestShouldFileIssueUsesEngineEligibility|TestTriage|TestUpdate.*(Triage|KnownIssues|FailuresFirst|ShowAll)' -count=1
```

Expected: FAIL because the shared predicate, types, fields, keys, and renderer
do not exist.

- [ ] **Step 3: Add the protocol types and model fields**

Add to `tui/messages.go`:

```go
type OperationStartedMsg struct{ Name string }

type TriageLoadedMsg struct {
	Failures       []engine.PersistFailure
	CacheAvailable bool
}

type OperationErrMsg struct {
	Op  string
	Err error
}
```

Add to `tui/intents.go`:

```go
type RefreshTriageIntent struct{}
type SyncKnownIssuesIntent struct{}
type PreviewFileIssueIntent struct{ Fingerprint string }
```

Add to `Model`:

```go
triageFailures       []engine.PersistFailure
triageCursor         int
triageDetailActive   bool
triageCacheAvailable bool
failuresFirst        bool
discoveryGroups      []engine.Group
operation            string
```

Add key bindings:

```go
TriageRefresh   key.Binding
SyncKnownIssues key.Binding
FileIssue       key.Binding
```

with:

```go
TriageRefresh: key.NewBinding(key.WithKeys("t"), key.WithHelp("t", "refresh triage")),
SyncKnownIssues: key.NewBinding(key.WithKeys("K"), key.WithHelp("K", "sync known issues")),
FileIssue: key.NewBinding(key.WithKeys("i"), key.WithHelp("i", "file selected issue")),
```

- [ ] **Step 4: Implement triage rendering and eligibility**

Create `engine/issue_eligibility.go`:

```go
func validFingerprintPair(full, short string) bool {
	const prefix = "sha256:"
	if !strings.HasPrefix(full, prefix) || len(full) != len(prefix)+64 {
		return false
	}
	digest := strings.TrimPrefix(full, prefix)
	if len(short) != 16 || short != digest[:16] {
		return false
	}
	_, err := hex.DecodeString(digest)
	return err == nil
}

func EligibleForIssueFiling(f PersistFailure) bool {
	switch f.Classification {
	case ClassificationReal, ClassificationRealUnstable:
		return validFingerprintPair(f.Fingerprint, f.ShortFingerprint) &&
			f.KnownIssue == 0 && f.IssueAction == ""
	default:
		return false
	}
}
```

Import `encoding/hex` and `strings`. Change CLI `shouldFileIssue` to delegate
to this helper. Create
`tui/triage.go` with:

```go
func triageCounts(failures []engine.PersistFailure) (flakes, real, unstable, known, unknown int) {
	for _, f := range failures {
		switch f.Classification {
		case engine.ClassificationFlakeConfirmed, engine.ClassificationFlakeHistorical:
			flakes++
		case engine.ClassificationRealUnstable:
			unstable++
		default:
			real++
		}
		if f.KnownIssue != 0 || f.IssueAction == "known" || f.IssueAction == "dedup" || f.IssueAction == "filed" {
			known++
		} else if engine.EligibleForIssueFiling(f) {
			unknown++
		}
	}
	return
}

func (m Model) renderTriageSection(width int) string {
	if len(m.triageFailures) == 0 {
		return Styles.Header.Render("Triage") + "\n\n  no classified failures"
	}
	if m.triageDetailActive {
		return m.renderTriageDetail(width)
	}
	return m.renderTriageList(width)
}
```

`renderTriageList` must show six summary counts (`total` plus the five returned
counts) and one row per failure.
`renderTriageDetail` must show classification, class, attempts, retryability,
short fingerprint, known issue/action, reasons, and log path. Apply `truncate`
to every unbounded value and show `i file issue` only when
`engine.EligibleForIssueFiling` is true.

- [ ] **Step 5: Implement reducer behavior and stable ordering**

On `GroupsMsg`, deep-copy groups into both `m.groups` and
`m.discoveryGroups`. Add:

```go
func cloneGroups(in []engine.Group) []engine.Group {
	out := make([]engine.Group, len(in))
	for i, g := range in {
		out[i] = g
		out[i].Tests = append([]string(nil), g.Tests...)
	}
	return out
}
```

For `f`, stable-sort groups and each group's tests by current aggregate status,
with failed, panic, timeout, and running entries before pass, skip, and not-run.
For `a`, restore `m.groups = cloneGroups(m.discoveryGroups)`. Reapply the sort
after each `TestUpdateMsg` while `m.failuresFirst` is true.

Handle `TriageLoadedMsg`, cursor clamping, Enter/detail, Esc/back, `t`, `K`,
and eligible `i`. Add the new intent types to the existing executor dispatch
case. Block all three action keys while `m.operation != ""`.

- [ ] **Step 6: Run TUI tests and regenerate only new goldens**

Run:

```bash
go test ./engine ./cli -run 'TestEligibleForIssueFiling|TestShouldFileIssueUsesEngineEligibility' -count=1
go test ./tui -run 'TestTriage|TestUpdate.*(Triage|KnownIssues|FailuresFirst|ShowAll)' -update -count=1
go test ./tui -count=1
```

Expected: all TUI tests pass; new triage fixtures contain no ANSI escapes or raw secrets.

- [ ] **Step 7: Commit triage display and ordering**

```bash
git add engine/issue_eligibility.go engine/issue_eligibility_test.go \
  cli/triage.go cli/triage_test.go tui
git commit -m "feat(tui): add triage dashboard" \
  -m "Co-authored-by: Copilot App <223556219+Copilot@users.noreply.github.com>"
```

---

### Task 3: Add Pure TUI Result and Confirmation Overlays

**Files:**
- Create: `tui/confirm.go`
- Create: `tui/confirm_test.go`
- Create: `tui/cleanup.go`
- Create: `tui/cleanup_test.go`
- Modify: `tui/model.go`
- Modify: `tui/messages.go`
- Modify: `tui/intents.go`
- Modify: `tui/update.go`
- Modify: `tui/view.go`
- Modify: `tui/testdata/orphan_overlay.golden`
- Modify: `tui/testdata/sweep_confirm.golden`
- Modify: `tui/testdata/file_issue_confirm.golden`

**Interfaces:**
- Consumes: `provider.Resource`, `engine.PersistFailure`, `textinput.Model`, and Task 2's operation field.
- Produces: typed result messages, exact-phrase intent payloads, and overlay state used by dashboard services.

- [ ] **Step 1: Write failing confirmation and overlay tests**

Create `tui/confirm_test.go`:

```go
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
```

Create update tests that press `o`, `e`, `s`, and `i`, type wrong and exact
phrases, and assert that no confirmation intent appears for wrong input. Also
prove that an `OperationErrMsg` for a rejected second action does not clear the
first action's `m.operation`.

- [ ] **Step 2: Run focused tests and verify failure**

Run:

```bash
go test ./tui -run 'TestConfirmation|TestOrphan|TestSweep|TestFileIssue|TestExport' -count=1
```

Expected: FAIL because overlay fields, messages, intents, and renderers are absent.

- [ ] **Step 3: Add exact protocol types**

Add to `tui/intents.go`:

```go
type ExportReportIntent struct{}
type ListOrphansIntent struct{}

type ConfirmSweepIntent struct {
	Owner     string
	Phrase    string
	Resources []provider.Resource
}

type ConfirmFileIssueIntent struct {
	Fingerprint string
	Phrase      string
}
```

Add to `tui/messages.go`:

```go
type KnownIssueSyncDoneMsg struct {
	Count     int
	CachePath string
	Failures []engine.PersistFailure
}

type ReportExportDoneMsg struct {
	MarkdownPath string
	HTMLPath     string
}

type OrphansListedMsg struct {
	Owner     string
	Resources []provider.Resource
}

type SweepDoneMsg struct {
	Remaining       []provider.Resource
	SnapshotChanged bool
}

type IssuePreviewMsg struct {
	Fingerprint      string
	ShortFingerprint string
	IssuesRepo       string
	Title            string
	Labels           []string
	Classification   string
}

type IssueFiledMsg struct {
	Fingerprint string
	IssueNumber int
	Action      string
	Failures    []engine.PersistFailure
}
```

- [ ] **Step 4: Add model fields and phrase helpers**

Add to `Model`:

```go
orphanOverlayActive bool
orphanOwner         string
orphans             []provider.Resource
sweepConfirmActive  bool
sweepInput          textinput.Model
fileIssueActive     bool
fileIssuePreview    IssuePreviewMsg
fileIssueInput      textinput.Model
resultOverlayActive bool
resultTitle         string
resultLines         []string
```

Create `tui/confirm.go`:

```go
func phraseMatches(input, expected string) bool { return input == expected }
func sweepPhrase(owner string) string           { return "SWEEP " + owner }
func fileIssuePhrase(short string) string       { return "FILE " + short }
```

Initialize each text input with `textinput.New()`, focus it when the
confirmation overlay opens, and clear it whenever the overlay closes.

- [ ] **Step 5: Implement cleanup and result overlays**

Create `tui/cleanup.go` with:

```go
func (m Model) renderOrphanOverlay(width int) string
func (m Model) renderSweepConfirm(width int) string
func (m Model) renderFileIssueConfirm(width int) string
func (m Model) renderResultOverlay(width int) string
```

The orphan overlay must show owner, count, `KIND`, `NAME`, and `URL`. The
sweep overlay must show every target and the exact expected phrase. The issue
overlay must show only repository, redacted title, labels, classification, and
fingerprint. It must not render an issue body.

Mirror this overlay priority in both `Update` and `View`:

```go
switch {
case m.pickerActive:
	// existing
case m.editorActive:
	// existing
case m.sweepConfirmActive:
	// sweep text input
case m.orphanOverlayActive:
	// orphan list
case m.fileIssueActive:
	// issue text input
case m.resultOverlayActive:
	// completed paths/counts
default:
	// main body
}
```

Each overlay must consume all keys and close on Esc. `s` works only in a
non-empty orphan overlay. Enter emits `ConfirmSweepIntent` or
`ConfirmFileIssueIntent` only after exact local validation.

- [ ] **Step 6: Handle all completion and error messages**

`OperationStartedMsg` sets `m.operation`. A completion clears only its matching
operation: `RunDoneMsg` clears `"run"`, `TriageLoadedMsg` clears `"triage"`,
`KnownIssueSyncDoneMsg` clears `"known-issues-sync"`,
`ReportExportDoneMsg` clears `"export"`, `OrphansListedMsg` clears
`"orphans"`, `SweepDoneMsg` clears `"sweep"`, `IssuePreviewMsg` clears
`"issue-preview"`, and `IssueFiledMsg` clears `"issue-file"`.
`OperationErrMsg` clears only when `msg.Op == m.operation`; a
rejected second operation must not clear the operation already in progress.
Preserve the prior triage and orphan data on errors.
`SweepDoneMsg{SnapshotChanged:true}` replaces the displayed orphan list and
requires a new confirmation. `IssueFiledMsg` replaces the triage failure list
and closes the filing overlay.

While `m.operation != ""`, do not emit retry/resume, `t`, `K`, `e`, `o`, `s`,
or `i` intents. Keep navigation, copy, `f`, `a`, help, and quit responsive.

- [ ] **Step 7: Run and update TUI fixtures**

Run:

```bash
go test ./tui -run 'TestConfirmation|TestOrphan|TestSweep|TestFileIssue|TestExport' -update -count=1
go test ./tui -count=1
```

Expected: all tests pass; wrong phrases emit no intent.

- [ ] **Step 8: Commit pure TUI action surfaces**

```bash
git add tui
git commit -m "feat(tui): add safe action overlays" \
  -m "Co-authored-by: Copilot App <223556219+Copilot@users.noreply.github.com>"
```

---

### Task 4: Make Known-Issue and Report Writes Atomic

**Files:**
- Create: `engine/atomic.go`
- Create: `engine/atomic_test.go`
- Modify: `engine/known_issues.go`
- Modify: `engine/known_issues_test.go`
- Modify: `engine/report.go`
- Modify: `engine/report_test.go`

**Interfaces:**
- Consumes: existing report and YAML encoders.
- Produces: unchanged public `SaveKnownIssuesFile`, `ExportMarkdown`, and `ExportHTML` APIs with private-file atomic writes.

- [ ] **Step 1: Write failing atomicity tests**

Add tests that pre-create a target, force the writer callback to fail, and
assert the original target remains:

```go
func TestWriteAtomicPreservesTargetOnWriteError(t *testing.T) {
	path := filepath.Join(t.TempDir(), "target")
	if err := os.WriteFile(path, []byte("old"), 0o600); err != nil {
		t.Fatal(err)
	}
	wantErr := errors.New("write failed")
	err := writeAtomic(path, 0o600, func(io.Writer) error { return wantErr })
	if !errors.Is(err, wantErr) {
		t.Fatalf("error = %v, want %v", err, wantErr)
	}
	got, _ := os.ReadFile(path)
	if string(got) != "old" {
		t.Fatalf("target = %q, want old", got)
	}
}
```

Add success tests for mode `0600` and absence of `*.tmp` files.

- [ ] **Step 2: Run tests and verify failure**

Run:

```bash
go test ./engine -run 'TestWriteAtomic|TestSaveKnownIssuesFile|TestExport' -count=1
```

Expected: FAIL because `writeAtomic` does not exist.

- [ ] **Step 3: Implement the shared atomic writer**

Create `engine/atomic.go`:

```go
package engine

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

func writeAtomic(path string, perm fs.FileMode, write func(io.Writer) error) (err error) {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return err
	}
	f, err := os.CreateTemp(dir, "."+filepath.Base(path)+".*.tmp")
	if err != nil {
		return err
	}
	tmp := f.Name()
	defer func() {
		_ = f.Close()
		if err != nil {
			_ = os.Remove(tmp)
		}
	}()
	if err = f.Chmod(perm); err != nil {
		return err
	}
	if err = write(f); err != nil {
		return err
	}
	if err = f.Close(); err != nil {
		return err
	}
	err = os.Rename(tmp, path)
	return err
}
```

- [ ] **Step 4: Refactor YAML and report exporters onto writers**

In `engine/known_issues.go`, marshal first, then:

```go
return writeAtomic(path, 0o600, func(w io.Writer) error {
	_, err := w.Write(data)
	return err
})
```

In `engine/report.go`, extract:

```go
func writeMarkdown(w io.Writer, groups []Group, res RunResult, red *redact.Redactor) error
func writeHTML(w io.Writer, groups []Group, res RunResult, red *redact.Redactor) error
```

Keep all existing rendering logic in those functions. Make the public
exporters call `writeAtomic(path, 0o600, ...)`.

- [ ] **Step 5: Run engine and repository tests**

Run:

```bash
go test ./engine -count=1
go test ./... -count=1
```

Expected: PASS with unchanged report content and CLI behavior.

- [ ] **Step 6: Commit atomic persistence**

```bash
git add engine
git commit -m "fix(engine): write reports and cache atomically" \
  -m "Co-authored-by: Copilot App <223556219+Copilot@users.noreply.github.com>"
```

---

### Task 5: Extract the Shared Triage State Service

**Files:**
- Create: `cli/triage_service.go`
- Create: `cli/triage_service_test.go`
- Modify: `cli/triage.go`
- Modify: `cli/triage_test.go`

**Interfaces:**
- Consumes: `reconstructFailures`, `processTriageIssues`, state lock/save, registry/filer factories.
- Produces: `triageStateService.apply` for an already-locked run and `triageStateService.Refresh` for standalone CLI/TUI refresh.

- [ ] **Step 1: Write failing offline, force-reconstruct, and lock tests**

Create tests for:

```go
func TestTriageStateServiceRefreshUsesCacheWithoutLiveLookup(t *testing.T)
func TestTriageStateServiceForceReconstructClearsPassingFailure(t *testing.T)
func TestTriageStateServiceLocksBeforeLoad(t *testing.T)
func TestPersistSuiteFailureLogsRedactsAndClears(t *testing.T)
func TestRunTriageKeepsExistingJSONAndTextBehavior(t *testing.T)
func TestRunTriageRefusesConcurrentStateMutation(t *testing.T)
```

Use a cache fixture whose fingerprint matches one failure. Set registry and
filer factories that fail the test if called during offline refresh.

- [ ] **Step 2: Run focused tests and verify failure**

Run:

```bash
go test ./cli -run 'TestTriageStateService|TestPersistSuiteFailureLogs|TestRunTriage(Keeps|Refuses)' -count=1
```

Expected: FAIL because the service does not exist.

- [ ] **Step 3: Implement the service types**

Create `cli/triage_service.go`:

```go
type triageStateResult struct {
	Failures       []engine.PersistFailure
	CacheAvailable bool
}

type triageStateService struct {
	root            string
	statePath       string
	red             *redact.Redactor
	registryFactory func(triageOptions) (issueRegistry, error)
	filerFactory    func(string) (issueFiler, error)
	getenv          func(string) string
}

func (s triageStateService) apply(
	ctx context.Context,
	st *engine.State,
	opts triageOptions,
	forceReconstruct bool,
) (triageStateResult, error) {
	failures := st.Failures
	if forceReconstruct || len(failures) == 0 {
		failures = reconstructFailures(s.root, *st)
	}
	cacheAvailable := opts.KnownIssuesPath != ""
	if cacheAvailable {
		if _, err := os.Stat(opts.KnownIssuesPath); err != nil {
			if !os.IsNotExist(err) {
				return triageStateResult{}, err
			}
			cacheAvailable = false
		}
	}
	updated, err := processTriageIssues(
		ctx, failures, st.Mode, s.root, opts, s.red,
		s.registryFactory, s.filerFactory, s.getenv,
	)
	if err != nil {
		return triageStateResult{}, err
	}
	st.Failures = updated
	return triageStateResult{Failures: updated, CacheAvailable: cacheAvailable}, nil
}

func (s triageStateService) Refresh(
	ctx context.Context,
	opts triageOptions,
	forceReconstruct bool,
) (triageStateResult, error) {
	release, err := engine.AcquireLock(s.statePath + ".lock")
	if err != nil {
		return triageStateResult{}, err
	}
	defer release()
	st, err := engine.Load(s.statePath)
	if err != nil {
		return triageStateResult{}, err
	}
	result, err := s.apply(ctx, &st, opts, forceReconstruct)
	if err != nil {
		return triageStateResult{}, err
	}
	if err := st.Save(s.statePath); err != nil {
		return triageStateResult{}, err
	}
	return result, nil
}
```

Extract the existing build/pre-run log block in `runWithSpecTriage` into:

```go
func persistSuiteFailureLogs(
	root string,
	st *engine.State,
	res engine.RunResult,
	red *redact.Redactor,
) error
```

The helper must set `BuildFailed` and `PreRunFailed`, write redacted `build.log`
and `pre-run.log` files under `.pulsar-failures` when present, remove stale
suite logs and clear their state paths after successful runs, and return all
filesystem errors. Keep `runWithSpecTriage` behavior unchanged by calling the
helper from the existing location.

- [ ] **Step 4: Refactor `runTriage` to use the service**

Keep flag parsing and output formatting in `runTriage`. Build the service with
the existing factories and call:

```go
result, err := svc.Refresh(context.Background(), opts, false)
```

Then pass `result.Failures` to existing JSON/text formatters. Preserve live
lookup and confirmed filing when CLI flags request them. Preserve normal CLI
text, JSON, and exit behavior. The one intentional concurrency change is that
`triage` now returns an explicit lock error rather than racing another process
and potentially overwriting newer state; cover that behavior with the lock
test from Step 1.

- [ ] **Step 5: Run all triage tests**

Run:

```bash
go test ./cli -run 'Test.*Triage|TestRunRetries' -count=1
go test ./cli -count=1
```

Expected: all existing and new triage tests pass with byte-compatible JSON fields.

- [ ] **Step 6: Commit the triage service**

```bash
git add cli/triage_service.go cli/triage_service_test.go cli/triage.go cli/triage_test.go
git commit -m "refactor(cli): share triage state workflow" \
  -m "Co-authored-by: Copilot App <223556219+Copilot@users.noreply.github.com>"
```

---

### Task 6: Add the Dashboard Operation Gate and Post-Run Triage

**Files:**
- Modify: `cli/dashboard.go`
- Modify: `cli/dashboard_test.go`

**Interfaces:**
- Consumes: Task 5's `triageStateService.apply` and Task 2's TUI protocol.
- Produces: one operation gate, persisted startup failures, cached post-run classification, and `RefreshTriageIntent` handling.

- [ ] **Step 1: Write failing producer tests**

Add:

```go
func TestSeedSendsPersistedTriageFailures(t *testing.T)
func TestRunIntentClassifiesFailuresOfflineAndSendsTriage(t *testing.T)
func TestRunIntentRemovesStaleFailureWhenRetryPasses(t *testing.T)
func TestRunIntentPreservesFiledMetadataAfterCachedRematch(t *testing.T)
func TestRunIntentPersistsSuiteFailureForTriage(t *testing.T)
func TestRunIntentAlwaysSendsRunDoneAfterPersistenceError(t *testing.T)
func TestRefreshTriageIntentUsesOperationGate(t *testing.T)
func TestOperationErrorsAreRedacted(t *testing.T)
```

For the gate test, block the first operation on a channel, start a second, and
assert one `OperationErrMsg` with `"another operation is in progress"`.

- [ ] **Step 2: Run focused tests and verify failure**

Run:

```bash
go test ./cli -run 'TestSeedSendsPersistedTriage|TestRunIntent(Classifies|Removes|Preserves|Persists|Always)|TestRefreshTriageIntent|TestOperationErrors' -count=1
```

Expected: FAIL because seed, post-run triage, and the generalized gate are absent.

- [ ] **Step 3: Replace `busy` with the operation gate**

In `dashboard`:

```go
opGate atomic.Bool
```

Add:

```go
func (db *dashboard) beginOperation(name string) bool {
	if !db.opGate.CompareAndSwap(false, true) {
		db.send(tui.OperationErrMsg{
			Op:  name,
			Err: errors.New("another operation is in progress"),
		})
		return false
	}
	db.send(tui.OperationStartedMsg{Name: name})
	return true
}

func (db *dashboard) endOperation() { db.opGate.Store(false) }

func (db *dashboard) sendOperationError(name string, err error) {
	db.cfgMu.Lock()
	red := db.red
	db.cfgMu.Unlock()
	if red != nil {
		err = errors.New(red.String(err.Error()))
	}
	db.send(tui.OperationErrMsg{Op: name, Err: err})
}
```

Use `beginOperation("run")` and `defer endOperation()` in `runIntent`.
Replace every test reference to `db.busy` with `db.opGate`, including the
existing concurrent-run test. Every later dashboard handler that successfully
calls `beginOperation` must immediately `defer db.endOperation()`.

- [ ] **Step 4: Send persisted triage during seed**

Add these helpers:

```go
func (db *dashboard) knownIssuesPath() (string, error) {
	base, err := db.userConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(base, "terraform-provider-tester", "known-issues.yaml"), nil
}

func (db *dashboard) triageService(red *redact.Redactor) triageStateService {
	return triageStateService{
		root:      db.root,
		statePath: db.statePath,
		red:       red,
		getenv:    db.getenv,
	}
}
```

After persisted `TestUpdateMsg` values:

```go
cachePath, cacheErr := db.knownIssuesPath()
cacheAvailable := false
if cacheErr != nil {
	db.sendOperationError("known-issues-cache", cacheErr)
} else if _, err := os.Stat(cachePath); err == nil {
	cacheAvailable = true
} else if !os.IsNotExist(err) {
	db.sendOperationError("known-issues-cache", err)
}
db.send(tui.TriageLoadedMsg{
	Failures:       append([]engine.PersistFailure(nil), prev.Failures...),
	CacheAvailable: cacheAvailable,
})
```

Add `userConfigDir func() (string, error)` to `dashboard` and pass
`d.userConfigDir` from `runDashboard`.

- [ ] **Step 5: Classify before the final state save**

In `runIntent`, keep the state lock held and replace the existing block from
the early `cur.Save` through `RunDoneMsg`; do not retain the old pre-log save or
its `tui.ErrMsg` sends. After the runner returns, collect `runErr` and all
post-run errors, then perform this sequence:

1. write per-test failure logs;
2. merge fresh results;
3. persist redacted suite failure logs;
4. reconstruct current failures and carry forward issue metadata by
   fingerprint;
5. rematch the cache without reconstructing again;
6. save state exactly once;
7. send triage, run completion, and one joined redacted error.

```go
var postRunErrs []error
if runErr != nil {
	postRunErrs = append(postRunErrs, fmt.Errorf("run error: %w", runErr))
}

failDir := filepath.Join(filepath.Dir(db.statePath), ".pulsar-failures")
if _, err := engine.WriteFailures(failDir, res, red); err != nil {
	postRunErrs = append(postRunErrs, fmt.Errorf("writing failure logs: %w", err))
}
cur.Results = mergeResults(prev.Results, cur.Results)
if err := persistSuiteFailureLogs(db.root, &cur, res, red); err != nil {
	postRunErrs = append(postRunErrs, fmt.Errorf("writing suite failure logs: %w", err))
}
cur.Failures = preserveIssueMetadata(
	prev.Failures,
	reconstructFailures(db.root, cur),
)

cachePath, pathErr := db.knownIssuesPath()
svc := db.triageService(red)
triageResult := triageStateResult{
	Failures:       cur.Failures,
	CacheAvailable: cacheFileExists(cachePath),
}
triageErr := pathErr
if triageErr == nil {
	triageOpts := triageOptions{
		IssuesRepo:         defaultIssuesRepo,
		KnownIssuesPath:    cachePath,
		KnownIssuesOffline: true,
	}
	if result, err := svc.apply(ctx, &cur, triageOpts, false); err != nil {
		triageErr = err
	} else {
		triageResult = result
	}
}
if err := cur.Save(db.statePath); err != nil {
	postRunErrs = append(postRunErrs, fmt.Errorf("saving state: %w", err))
}
db.send(tui.TriageLoadedMsg{
	Failures:       triageResult.Failures,
	CacheAvailable: triageResult.CacheAvailable,
})
if triageErr != nil {
	postRunErrs = append(postRunErrs, fmt.Errorf("triage: %w", triageErr))
}
db.send(tui.RunDoneMsg{Result: redactRunResult(red, res)})
if err := errors.Join(postRunErrs...); err != nil {
	db.sendOperationError("run", err)
}
```

`preserveIssueMetadata` matches by full fingerprint and copies only
`KnownIssue` and `IssueAction` into reconstructed failures whose corresponding
fields are still empty. This keeps fresh local classification, removes failures
whose retry passed, and retains valid `known`, `dedup`, and `filed` metadata
even when the local cache does not yet contain a newly created issue.
`cacheFileExists` returns false for an empty path or any `os.Stat` error.
After `RunStartedMsg`, every path must still send one `RunDoneMsg`, including
persistence and triage errors, so the spinner cannot stick.

- [ ] **Step 6: Wire standalone triage refresh**

Add a handler that uses `beginOperation("triage")`, resolves `cachePath`, and
calls:

```go
result, err := svc.Refresh(ctx, triageOptions{
	IssuesRepo:         defaultIssuesRepo,
	KnownIssuesPath:    cachePath,
	KnownIssuesOffline: true,
}, false)
```

Send `TriageLoadedMsg` on success and a redacted `OperationErrMsg` on failure.
Extend the executor switch with `RefreshTriageIntent`.

- [ ] **Step 7: Run dashboard, race, and CLI tests**

Run:

```bash
go test ./cli -run 'TestSeed|TestRunIntent|TestRefreshTriage|TestOperation' -count=1
go test -race ./cli -count=1
```

Expected: PASS; no stale failures or races.

- [ ] **Step 8: Commit dashboard triage integration**

```bash
git add cli/dashboard.go cli/dashboard_test.go
git commit -m "feat(tui): classify dashboard runs" \
  -m "Co-authored-by: Copilot App <223556219+Copilot@users.noreply.github.com>"
```

---

### Task 7: Add Persistent Known-Issue Sync

**Files:**
- Create: `cli/known_issues_service.go`
- Create: `cli/known_issues_service_test.go`
- Modify: `cli/cli.go`
- Modify: `cli/known_issues.go`
- Modify: `cli/known_issues_test.go`
- Modify: `cli/dashboard.go`
- Modify: `cli/dashboard_test.go`

**Interfaces:**
- Consumes: atomic `engine.SaveKnownIssuesFile`, `deps.userConfigDir`, and Task 5's offline rematch.
- Produces: `knownIssueSyncService.Sync`, injectable known-issue lister factory, CLI compatibility, and `SyncKnownIssuesIntent` handling.

- [ ] **Step 1: Write failing cache-path, preservation, and rematch tests**

Create:

```go
func TestKnownIssueSyncServiceWritesConfiguredPath(t *testing.T)
func TestKnownIssueSyncServicePreservesExistingCacheOnFailure(t *testing.T)
func TestKnownIssueSyncCommandKeepsExplicitOutBehavior(t *testing.T)
func TestDashboardKnownIssueSyncUsesPersistentDefaultPath(t *testing.T)
func TestDashboardSyncKnownIssuesRematchesFailures(t *testing.T)
```

Use a fake lister that returns one `engine.KnownIssueEntry` and another that
returns an error. Seed the failure state with the same fingerprint.

- [ ] **Step 2: Run focused tests and verify failure**

Run:

```bash
go test ./cli -run 'TestKnownIssueSync|TestDashboard.*KnownIssueSync' -count=1
```

Expected: FAIL because the service and dependency seam are absent.

- [ ] **Step 3: Add the lister seam and service**

Define:

```go
type knownIssueLister interface {
	ListKnownIssues(context.Context) ([]engine.KnownIssueEntry, error)
}

type knownIssueSyncResult struct {
	Count     int
	CachePath string
	SyncedAt  time.Time
}

type knownIssueSyncService struct {
	issuesRepo string
	cachePath  string
	now        func() time.Time
	newLister  func(string) (knownIssueLister, error)
}

func (s knownIssueSyncService) Sync(ctx context.Context) (knownIssueSyncResult, error) {
	client, err := s.newLister(s.issuesRepo)
	if err != nil {
		return knownIssueSyncResult{}, err
	}
	entries, err := client.ListKnownIssues(ctx)
	if err != nil {
		return knownIssueSyncResult{}, err
	}
	now := s.now().UTC()
	file := engine.KnownIssueFile{
		Version: 1, Source: "github", IssuesRepo: s.issuesRepo,
		Label: "acctest-failure", SyncedAt: &now, Entries: entries,
	}
	if err := engine.SaveKnownIssuesFile(s.cachePath, file); err != nil {
		return knownIssueSyncResult{}, err
	}
	return knownIssueSyncResult{Count: len(entries), CachePath: s.cachePath, SyncedAt: now}, nil
}
```

Add `newKnownIssueLister` to `deps`. Production returns
`githubprovider.NewIssueClientFromEnv(repo)`. Tests inject a fake.

- [ ] **Step 4: Preserve CLI behavior through the service**

Change `runKnownIssuesSync` to accept `deps`, keep `--out` required, and call
the service with the explicit path. Do not introduce an implicit CLI output
path.

Expected text remains:

```text
synced <count> known issue(s) to <path>
```

- [ ] **Step 5: Wire explicit TUI sync and offline rematch**

Resolve the TUI cache path with Task 6's `db.knownIssuesPath`. The handler must:

1. acquire `opGate` as `"known-issues-sync"`;
2. call `Sync`;
3. call `triageStateService.Refresh` with `KnownIssuesOffline:true`;
4. send one `KnownIssueSyncDoneMsg` containing count, path, and rematched failures;
5. send a redacted `OperationErrMsg` without clearing prior state on error.

- [ ] **Step 6: Run service, dashboard, and CLI regression tests**

Run:

```bash
go test ./cli -run 'TestKnownIssue|TestDashboard.*(KnownIssueSync|SyncKnownIssues)' -count=1
go test -race ./cli -count=1
```

Expected: PASS; a failed sync leaves the old cache bytes unchanged.

- [ ] **Step 7: Commit known-issue sync**

```bash
git add cli
git commit -m "feat(tui): sync persistent known issues" \
  -m "Co-authored-by: Copilot App <223556219+Copilot@users.noreply.github.com>"
```

---

### Task 8: Add Persistent Markdown and HTML Report Export

**Files:**
- Create: `cli/report_service.go`
- Create: `cli/report_service_test.go`
- Modify: `cli/cli.go`
- Modify: `cli/dashboard.go`
- Modify: `cli/dashboard_test.go`

**Interfaces:**
- Consumes: atomic engine exporters, persisted state, discovered groups, redactor, and user config directory.
- Produces: `reportExportService.Export` and `ExportReportIntent` handling without changing `report` CLI behavior.

- [ ] **Step 1: Write failing path, redaction, and no-dirty-worktree tests**

Create:

```go
func TestReportExportServiceWritesBothFormatsOutsideRoot(t *testing.T)
func TestReportExportServiceUsesUTCNamesAndRedacts(t *testing.T)
func TestReportCommandWithoutPathsStillPrintsSummary(t *testing.T)
func TestDashboardExportSendsBothPaths(t *testing.T)
```

Inject a fixed clock:

```go
now := func() time.Time {
	return time.Date(2026, 7, 9, 16, 30, 0, 0, time.UTC)
}
```

Expected filenames:

```text
20260709T163000Z-report.md
20260709T163000Z-report.html
```

- [ ] **Step 2: Run focused tests and verify failure**

Run:

```bash
go test ./cli -run 'TestReportExportService|TestReportCommandWithoutPaths|TestDashboardExport' -count=1
```

Expected: FAIL because the service does not exist.

- [ ] **Step 3: Extract state reconstruction**

Move `runReport`'s state-to-result mapping into:

```go
func runResultFromState(st engine.State) engine.RunResult {
	res := engine.RunResult{
		BuildFailed:  st.BuildFailed,
		PreRunFailed: st.PreRunFailed,
	}
	if st.BuildLog != "" {
		res.BuildOutput = []string{"see " + st.BuildLog + "\n"}
	}
	if st.PreRunLog != "" {
		res.PreRunOutput = []string{"see " + st.PreRunLog + "\n"}
	}
	for _, r := range st.Results {
		status, _ := provider.ParseStatus(r.Status)
		res.Tests = append(res.Tests, engine.TestResult{
			Package: r.Package, Name: r.Test, Sub: r.Sub,
			Status: status, Elapsed: r.Elapsed,
		})
	}
	return res
}
```

Use it from both `runReport` and the new service.

- [ ] **Step 4: Implement timestamped export**

Create:

```go
type reportExportResult struct {
	MarkdownPath string
	HTMLPath     string
	RunAt        time.Time
	TestCount    int
}

type reportExportService struct {
	statePath     string
	groups        []engine.Group
	red           *redact.Redactor
	userConfigDir func() (string, error)
	now           func() time.Time
}

func (s reportExportService) Export(ctx context.Context) (reportExportResult, error) {
	if err := ctx.Err(); err != nil {
		return reportExportResult{}, err
	}
	release, err := engine.AcquireLock(s.statePath + ".lock")
	if err != nil {
		return reportExportResult{}, err
	}
	defer release()
	st, err := engine.Load(s.statePath)
	if err != nil {
		return reportExportResult{}, err
	}
	base, err := s.userConfigDir()
	if err != nil {
		return reportExportResult{}, err
	}
	dir := filepath.Join(base, "terraform-provider-tester", "reports")
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return reportExportResult{}, err
	}
	stamp := s.now().UTC().Format("20060102T150405Z")
	md := filepath.Join(dir, stamp+"-report.md")
	html := filepath.Join(dir, stamp+"-report.html")
	res := runResultFromState(st)
	if err := engine.ExportMarkdown(md, s.groups, res, s.red); err != nil {
		return reportExportResult{}, err
	}
	if err := engine.ExportHTML(html, s.groups, res, s.red); err != nil {
		_ = os.Remove(md)
		return reportExportResult{}, err
	}
	return reportExportResult{
		MarkdownPath: md, HTMLPath: html, RunAt: st.RunAt,
		TestCount: len(st.Results),
	}, nil
}
```

- [ ] **Step 5: Wire the dashboard**

Handle `ExportReportIntent` under `beginOperation("export")`, snapshot `db.red`
under `cfgMu`, call the service, and send `ReportExportDoneMsg`. Keep
`runReport`'s no-path terminal summary unchanged.

- [ ] **Step 6: Run tests**

Run:

```bash
go test ./cli -run 'TestReport|TestDashboardExport' -count=1
go test ./engine ./cli -count=1
```

Expected: PASS; no file is created under the provider root.

- [ ] **Step 7: Commit report export**

```bash
git add cli
git commit -m "feat(tui): export persistent reports" \
  -m "Co-authored-by: Copilot App <223556219+Copilot@users.noreply.github.com>"
```

---

### Task 9: Add Orphan Preview and Snapshot-Safe Sweep

**Files:**
- Create: `cli/cleanup_service.go`
- Create: `cli/cleanup_service_test.go`
- Modify: `cli/dashboard.go`
- Modify: `cli/dashboard_test.go`
- Modify: `fakeprovider/fake.go`

**Interfaces:**
- Consumes: `provider.Orphans`, `provider.Sweep`, current `GITHUB_OWNER`, and Task 3's snapshot-bearing intent.
- Produces: normalized snapshot comparison, double phrase validation, residual resources, and no provider interface change.

- [ ] **Step 1: Add fake-provider error injection and failing service tests**

Change the fake:

```go
OrphansVal []provider.Resource
OrphansErr error
```

```go
func (f *Fake) Orphans(context.Context) ([]provider.Resource, error) {
	return f.OrphansVal, f.OrphansErr
}
```

Create:

```go
func TestCleanupListIsReadOnly(t *testing.T)
func TestCleanupSweepRejectsWrongPhrase(t *testing.T)
func TestCleanupSweepRejectsChangedOwner(t *testing.T)
func TestCleanupSweepAbortsWhenSnapshotChanges(t *testing.T)
func TestCleanupSweepPassesConfirmTrueAndReturnsResidual(t *testing.T)
func TestCleanupSweepReturnsOriginalAndRefreshErrors(t *testing.T)
```

- [ ] **Step 2: Run focused tests and verify failure**

Run:

```bash
go test ./cli -run 'TestCleanup' -count=1
```

Expected: FAIL because the cleanup service does not exist.

- [ ] **Step 3: Implement normalized snapshots**

Create:

```go
type cleanupSnapshot struct {
	Owner     string
	Resources []provider.Resource
}

type sweepRequest struct {
	Owner     string
	Phrase    string
	Resources []provider.Resource
}

type sweepResult struct {
	Remaining       []provider.Resource
	SnapshotChanged bool
}

func normalizedResources(in []provider.Resource) []provider.Resource {
	out := append([]provider.Resource(nil), in...)
	sort.Slice(out, func(i, j int) bool {
		if out[i].Kind != out[j].Kind {
			return out[i].Kind < out[j].Kind
		}
		if out[i].Name != out[j].Name {
			return out[i].Name < out[j].Name
		}
		return out[i].URL < out[j].URL
	})
	return out
}

func sameResources(a, b []provider.Resource) bool {
	return reflect.DeepEqual(normalizedResources(a), normalizedResources(b))
}
```

- [ ] **Step 4: Implement list and sweep**

```go
type cleanupService struct {
	prov   provider.Provider
	getenv func(string) string
}

func (s cleanupService) List(ctx context.Context) (cleanupSnapshot, error) {
	owner := s.getenv("GITHUB_OWNER")
	if owner == "" {
		return cleanupSnapshot{}, errors.New("GITHUB_OWNER environment variable not set")
	}
	resources, err := s.prov.Orphans(ctx)
	if err != nil {
		return cleanupSnapshot{}, err
	}
	return cleanupSnapshot{Owner: owner, Resources: resources}, nil
}

func (s cleanupService) Sweep(ctx context.Context, req sweepRequest) (sweepResult, error) {
	currentOwner := s.getenv("GITHUB_OWNER")
	if req.Owner == "" || currentOwner != req.Owner {
		return sweepResult{}, errors.New("sweep owner changed; list orphans again")
	}
	if req.Phrase != "SWEEP "+currentOwner {
		return sweepResult{}, errors.New("sweep confirmation phrase does not match")
	}
	if len(req.Resources) == 0 {
		return sweepResult{}, errors.New("no orphaned resources to sweep")
	}
	fresh, err := s.prov.Orphans(ctx)
	if err != nil {
		return sweepResult{}, err
	}
	if !sameResources(req.Resources, fresh) {
		return sweepResult{Remaining: fresh, SnapshotChanged: true}, nil
	}
	sweepErr := s.prov.Sweep(ctx, provider.SweepOpts{
		Targets: []string{"repositories", "teams"},
		Confirm: true,
	})
	remaining, refreshErr := s.prov.Orphans(ctx)
	return sweepResult{Remaining: remaining}, errors.Join(sweepErr, refreshErr)
}
```

- [ ] **Step 5: Wire list and sweep handlers**

`ListOrphansIntent` uses `beginOperation("orphans")`, calls `List`, and sends
`OrphansListedMsg`. `ConfirmSweepIntent` uses `beginOperation("sweep")`, passes
the phrase, owner, and displayed resources to `Sweep`, and sends
`SweepDoneMsg`.

On a changed snapshot, send `SnapshotChanged:true` and do not call
`provider.Sweep`. When `Sweep` returns both a result and an error, send
`SweepDoneMsg` with the residual resources first, then send the redacted
`OperationErrMsg`; never discard the residual snapshot. Redact all error
strings before delivery.

- [ ] **Step 6: Run cleanup, provider, and race tests**

Run:

```bash
go test ./cli -run 'TestCleanup|TestDashboard.*(Orphan|Sweep)' -count=1
go test ./provider/github -run 'Test.*(Orphan|Sweep)' -count=1
go test -race ./cli -count=1
```

Expected: PASS; `Confirm:true` is observed only after both phrase and snapshot checks.

- [ ] **Step 7: Commit cleanup parity**

```bash
git add cli fakeprovider
git commit -m "feat(tui): add guarded cleanup workflow" \
  -m "Co-authored-by: Copilot App <223556219+Copilot@users.noreply.github.com>"
```

---

### Task 10: Add Selected Issue Preview and Confirmed Filing

**Files:**
- Create: `cli/issue_service.go`
- Create: `cli/issue_service_test.go`
- Modify: `cli/dashboard.go`
- Modify: `cli/dashboard_test.go`
- Modify: `cli/triage.go`

**Interfaces:**
- Consumes: authoritative persisted failure by fingerprint, live registry/filer factories, failure logs, redactor, and Task 3's exact phrase.
- Produces: redacted preview metadata, one-failure filing, dedup recovery, and persisted `known`/`dedup`/`filed` actions.

- [ ] **Step 1: Write failing preview, eligibility, dedup, and recovery tests**

Create:

```go
func TestIssueServicePreviewOmitsBodyAndRedactsTitle(t *testing.T)
func TestIssueServiceFileOneRejectsWrongPhrase(t *testing.T)
func TestIssueServiceFileOneRejectsFlakeAndKnownFailure(t *testing.T)
func TestIssueServiceFileOneUsesLiveRegistryBeforeDedup(t *testing.T)
func TestIssueServiceFileOneDedupsBeforeCreate(t *testing.T)
func TestIssueServiceFileOneCreatesOnlySelectedFailure(t *testing.T)
func TestIssueServiceFileOneReportsSaveFailureAfterCreate(t *testing.T)
func TestIssueServiceFileOneRetryRecoversCreatedIssueByDedup(t *testing.T)
```

Seed two eligible failures and assert only the requested fingerprint reaches
`CreateFailureIssue`.

- [ ] **Step 2: Run focused tests and verify failure**

Run:

```bash
go test ./cli -run 'TestIssueService' -count=1
```

Expected: FAIL because the service does not exist.

- [ ] **Step 3: Define authoritative request and result types**

```go
type issuePreviewResult struct {
	Fingerprint      string
	ShortFingerprint string
	IssuesRepo       string
	Title            string
	Labels           []string
	Classification   string
}

type issueFileResult struct {
	Fingerprint string
	IssueNumber int
	Action      string
	Failures    []engine.PersistFailure
}

type issueService struct {
	root            string
	statePath       string
	issuesRepo      string
	red             *redact.Redactor
	registryFactory func(triageOptions) (issueRegistry, error)
	filerFactory    func(string) (issueFiler, error)
	getenv          func(string) string
}
```

Requests carry only the fingerprint and phrase. The service reloads the
authoritative failure from state instead of trusting the TUI copy.

- [ ] **Step 4: Implement preview**

`Preview` loads state, finds the exact fingerprint, checks `shouldFileIssue`,
builds a draft with `BuildFailureIssueDraft`, and returns title and labels only:

```go
func (s issueService) Preview(fingerprint string) (issuePreviewResult, error) {
	st, err := engine.Load(s.statePath)
	if err != nil {
		return issuePreviewResult{}, err
	}
	f, ok := findFailure(st.Failures, fingerprint)
	if !ok || !shouldFileIssue(f) {
		return issuePreviewResult{}, errors.New("selected failure is not eligible for issue filing")
	}
	draft := ghissues.BuildFailureIssueDraft(f, ghissues.IssueDraftOptions{
		IssuesRepo: s.issuesRepo,
		LogLines:   readFailureLogLines(s.root, f),
		Redactor:   s.red,
	})
	return issuePreviewResult{
		Fingerprint: f.Fingerprint, ShortFingerprint: f.ShortFingerprint,
		IssuesRepo: s.issuesRepo, Title: s.red.String(draft.Title),
		Labels: append([]string(nil), draft.Labels...),
		Classification: f.Classification,
	}, nil
}
```

- [ ] **Step 5: Implement single-failure filing under the state lock**

`FileOne` must:

1. acquire `statePath+".lock"` before loading;
2. reload and find the selected fingerprint;
3. validate `phrase == "FILE "+ShortFingerprint`;
4. validate `shouldFileIssue`;
5. run live registry lookup;
6. run `FindIssueByFingerprint`;
7. create one issue only if both miss;
8. replace only that failure in state;
9. save state before returning success.

Use:

```go
func replaceFailure(all []engine.PersistFailure, updated engine.PersistFailure) []engine.PersistFailure {
	out := append([]engine.PersistFailure(nil), all...)
	for i := range out {
		if out[i].Fingerprint == updated.Fingerprint {
			out[i] = updated
			return out
		}
	}
	return out
}
```

If creation succeeds and save fails, return the save error. A retry must call
`FindIssueByFingerprint` and persist the dedup match.

- [ ] **Step 6: Wire preview and confirmation intents**

`PreviewFileIssueIntent` calls `Preview` under
`beginOperation("issue-preview")` and sends `IssuePreviewMsg`.
`ConfirmFileIssueIntent` calls `FileOne` under
`beginOperation("issue-file")` and sends `IssueFiledMsg` only after state save.
Both paths send redacted operation errors with the matching operation name.

- [ ] **Step 7: Run issue, dashboard, and race tests**

Run:

```bash
go test ./cli -run 'TestIssueService|TestDashboard.*Issue' -count=1
go test ./provider/github -run 'Test.*Issue' -count=1
go test -race ./cli -count=1
```

Expected: PASS; fake filer `createCalls == 1` for the selected failure and `0` for every rejected/dedup case.

- [ ] **Step 8: Commit confirmed filing**

```bash
git add cli
git commit -m "feat(tui): file selected triage issues" \
  -m "Co-authored-by: Copilot App <223556219+Copilot@users.noreply.github.com>"
```

---

### Task 11: Document and Capture the Functional Dashboard

**Files:**
- Create: `tui/docs_screenshot_test.go`
- Create: `docs/images/triage.png`
- Modify: `docs/images/intro.png`
- Modify: `docs/images/preflight.png`
- Modify: `docs/images/groups.png`
- Modify: `README.md`
- Modify: `CHANGELOG.md`
- Modify: `docs/index.md`
- Modify: `docs/architecture.md`
- Modify: `docs/quickstart.md`
- Modify: `docs/troubleshooting.md`

**Interfaces:**
- Consumes: synthetic TUI models and completed key/action behavior.
- Produces: reproducible, secret-free screenshots and CLI-first user documentation.

- [ ] **Step 1: Write a gated synthetic screenshot generator**

Create `tui/docs_screenshot_test.go`. The test must skip unless
`UPDATE_DOC_IMAGES=1`. It renders fixed synthetic views into escaped `<pre>`
HTML, then invokes an already-installed headless Chrome:

```go
func writeScreenshot(t *testing.T, chrome, name, view string) {
	t.Helper()
	plain := stripANSI(view)
	page := `<!doctype html><html><head><meta charset="utf-8"><style>
html,body{margin:0;background:#0d1117;color:#f0f6fc}
body{padding:32px}
pre{font:20px/1.35 ui-monospace,SFMono-Regular,Menlo,Consolas,monospace;
white-space:pre;margin:0}
</style></head><body><pre>` + html.EscapeString(plain) + `</pre></body></html>`
	htmlPath := filepath.Join(t.TempDir(), name+".html")
	if err := os.WriteFile(htmlPath, []byte(page), 0o600); err != nil {
		t.Fatal(err)
	}
	out, err := filepath.Abs(filepath.Join("..", "docs", "images", name+".png"))
	if err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command(
		chrome,
		"--headless=new",
		"--disable-gpu",
		"--hide-scrollbars",
		"--window-size=1600,900",
		"--screenshot="+out,
		"file://"+htmlPath,
	)
	if output, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("render %s: %v\n%s", name, err, output)
	}
}

func TestUpdateDocsScreenshots(t *testing.T) {
	if os.Getenv("UPDATE_DOC_IMAGES") != "1" {
		t.Skip("set UPDATE_DOC_IMAGES=1 to regenerate docs screenshots")
	}
	chrome := "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome"
	if _, err := os.Stat(chrome); err != nil {
		t.Skip("headless Chrome is not installed")
	}
	m := fixedModel()
	writeScreenshot(t, chrome, "intro",
		renderIntro(introPeakFrame, 100, 30, "v0.3.0", true))

	m.section = sectionPreflight
	writeScreenshot(t, chrome, "preflight", m.View())

	m.section = sectionGroups
	writeScreenshot(t, chrome, "groups", m.View())

	m.section = sectionTriage
	m.triageFailures = sampleFailures()
	writeScreenshot(t, chrome, "triage", m.View())
}
```

Import `html`, `os`, `os/exec`, `path/filepath`, and `testing`. Reuse the
existing test-only `stripANSI` helper. Use only the synthetic data in
`fixedModel` and `sampleFailures`.

- [ ] **Step 2: Generate images**

Run:

```bash
UPDATE_DOC_IMAGES=1 go test ./tui -run TestUpdateDocsScreenshots -count=1
file docs/images/*.png
```

Expected: four valid PNGs, including `triage.png`; none contains credentials or real account names.

- [ ] **Step 3: Update CLI-first documentation**

Make these exact content changes:

- README feature list includes the Triage tab and guarded actions.
- README TUI section lists `t`, `K`, `i`, `e`, `o`, `s`, `f`, and `a`.
- README includes `docs/images/triage.png`.
- README keeps CLI and NDJSON before the optional TUI section.
- Architecture describes pure TUI, shared CLI services, operation gate, and user-config artifacts.
- Quickstart explains that routine triage is cache-only and `K` performs live sync.
- Troubleshooting includes exact phrase failures, changed orphan snapshots, cache path, and report directory.
- Changelog `[Unreleased]` records TESTER wordmark, fourth tab, functional keys, atomic artifacts, and guarded side effects.

- [ ] **Step 4: Run docs checks**

Run:

```bash
make checkdocs
git diff --check
```

Expected: PASS; no stale claim says export, orphan, or sweep are CLI-only.

- [ ] **Step 5: Commit docs and images**

```bash
git add README.md CHANGELOG.md docs tui/docs_screenshot_test.go
git commit -m "docs: show full dashboard parity" \
  -m "Co-authored-by: Copilot App <223556219+Copilot@users.noreply.github.com>"
```

---

### Task 12: Run End-to-End Local Verification and Review Gates

**Files:**
- Modify only files required by failures found during this task.

**Interfaces:**
- Consumes: the complete implementation.
- Produces: a clean, reviewed branch ready for a stacked AI-assisted pull request.

- [ ] **Step 1: Run formatting and static checks**

Run:

```bash
export PATH="/opt/homebrew/bin:$PATH"
gofmt -w $(git ls-files '*.go')
test -z "$(gofmt -l .)"
go vet ./...
make checkdocs
git diff --check
```

Expected: every command exits `0`.

- [ ] **Step 2: Run all tests and race-sensitive packages**

Run:

```bash
go test ./... -count=1
go test -race ./cli ./tui -count=1
```

Expected: all tests pass and the race detector reports no races.

- [ ] **Step 3: Run secret-canary regression tests**

Run:

```bash
go test ./cli ./engine ./tui -run 'Redact|Secret|Sensitive|Issue|Report|Triage' -count=1
```

Expected: PASS; no canary appears in state, report, TUI message, preview, error, or fixture assertions.

- [ ] **Step 4: Build and run a pseudo-TTY smoke session**

Build:

```bash
make build
```

Launch from the existing provider worktree with no acceptance-test credentials:

```bash
TPT_ROOT="$(pwd)"
cd "${PROVIDER_ROOT:?set PROVIDER_ROOT to a terraform-provider-github checkout}"
PULSAR_FORCE_TTY=1 "$TPT_ROOT/bin/terraform-provider-tester"
```

Verify without starting acceptance tests:

1. the splash spells `TESTER`;
2. all four tabs render;
3. `t`, `K`, `e`, and `o` produce clear success or credential errors;
4. wrong sweep and issue phrases produce no side effect;
5. Esc closes every overlay;
6. no secret value appears.

Expected: the session exits cleanly with `q`; no resource-creating test starts.

- [ ] **Step 5: Run independent spec and code reviews**

Dispatch one spec-compliance reviewer against:

```text
docs/superpowers/specs/2026-07-09-tui-full-parity-design.md
```

Dispatch one code-quality reviewer against the complete branch diff from
`persist-env-config`. Fix only verified findings, rerun the affected focused
tests, then rerun Steps 1 through 3.

- [ ] **Step 6: Commit verification fixes, if any**

If review or verification changed files:

```bash
git add -A
git commit -m "fix: address TUI parity review" \
  -m "Co-authored-by: Copilot App <223556219+Copilot@users.noreply.github.com>"
```

If no files changed, do not create an empty commit.

- [ ] **Step 7: Confirm branch scope**

Run:

```bash
git status --short --branch
git log --oneline persist-env-config..HEAD
git diff --stat persist-env-config...HEAD
git diff --check persist-env-config...HEAD
```

Expected: clean `tui-full-parity` branch containing the design, plan, and TUI parity commits only.
