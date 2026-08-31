# Terraform Provider Tester TUI Visual Polish Design

**Date:** 2026-07-13

**Status:** Approved for implementation
**Builds on:** `2026-07-09-tui-full-parity-design.md`

## Goal

Turn the complete four-tab dashboard into a polished, immediately legible
mission-control surface without weakening the CLI-first workflow, safety
gates, redaction, deterministic rendering, or narrow-terminal support.

The dashboard should feel intentionally designed rather than like styled
command output. A first-time user should know where they are, how to move
between tabs, whether the environment is ready, what will run, and what action
is safe to take next.

## User-reported papercut

The current tab bar looks horizontally navigable, but Left and Right do
nothing. Only Tab and Shift+Tab switch tabs. The redesign makes both interaction
models work and advertises them consistently.

## Constraints

- CLI and NDJSON remain the primary automation interfaces.
- The TUI remains an optional pure reducer. Network, filesystem, provider, and
  run side effects remain in CLI services.
- Existing exact confirmations remain case-sensitive:
  - `SWEEP <owner>`
  - `FILE <short-fingerprint>`
- Existing legacy compatibility identifiers remain unchanged:
  `.pulsar*`, `PULSAR_*`, `pulsar-fp-v1`, `pulsar/pre-run`, and the
  known-issue HTML marker.
- No Nerd Font dependency. UTF-8 terminals get box drawing and standard
  symbols; `NO_COLOR` and ASCII profiles remain complete.
- Color is supplementary. Every status also has text, a glyph, or both.
- Secret values never enter rendered model state, screenshots, docs fixtures,
  status messages, reports, or errors.
- The dashboard continues to serialize operations and cancel plus await active
  work on exit.

## Explored directions

### A. GitHub Mission Control — selected

A compact product header, numbered status-bearing tabs, summary metrics,
bordered content panels, strong selected-row treatment, contextual actions,
and responsive wide/narrow layouts.

**Why selected:** It improves navigation and hierarchy while preserving the
existing model, messages, services, and four-tab mental model. Most new display
data can be derived from current state.

### B. Guided wizard

A linear Preflight → Select → Run → Triage flow with explicit Next and Back
steps.

**Trade-off:** Easier for first use, but worse for retry, resume, inspection,
and operators who need to jump directly to Triage. It would also obscure the
CLI parity the dashboard was designed to expose.

### C. Dense observability console

An always-visible multi-pane dashboard showing checks, groups, run logs, and
triage simultaneously.

**Trade-off:** Visually dramatic on large terminals, but cramped below 140
columns, harder to test, and likely to overload first-time users.

## Visual language

### Product identity

The animated TESTER ignition splash remains. After ignition, the large welcome
card is replaced by a compact command-deck header so normal operation gives
more space to useful state.

Header layout:

```text
◆ TESTER  v0.3.0       github  organization  my-test-org       API 4,870/5,000
```

- `◆ TESTER` is the compact product mark.
- Version is muted.
- Provider, mode, and owner render as small labeled context chips.
- API headroom remains right-aligned and changes semantic color at the existing
  warning thresholds.
- On narrow terminals, owner and reset time disappear before core context.

### Palette

Keep GitHub Primer adaptive colors and add explicit surface roles:

- Accent blue: navigation, links, selected identifiers.
- Success green: ready and passed.
- Attention amber: warning, running, retryable.
- Danger red: blocked, failed, destructive.
- Flaky purple: retry-confirmed flakes.
- Muted gray: metadata and inactive controls.
- Border gray: panels and separators.
- Selected surface: subtle purple/blue background only for the active tab and
  selected row.

No rainbow decoration, gradients, blinking, or color-only meaning.

### Panels and spacing

- Section content uses one-column cards below 110 columns.
- At 110 columns and above, Preflight and Run use a two-column layout.
- Cards have a title, optional status glyph, one-line purpose, and concise
  contents.
- Empty states use an icon, a plain explanation, and the next useful key.
- Tables keep stable columns and use a full-row selected surface rather than
  only a small cursor glyph.

## Global chrome and navigation

### Tabs

Tabs become compact numbered segments with derived badges:

```text
  1 Preflight ✓    2 Groups 15    3 Run 168    4 Triage 3
```

- Preflight badge: ready, warning count, or blocked count.
- Groups badge: discovered group count.
- Run badge: total tests, or running indicator while active.
- Triage badge: classified failure count.
- Active tab uses a selected background and underline.
- Inactive tabs stay muted.

Supported navigation:

- Right or Tab: next tab.
- Left or Shift+Tab: previous tab.
- Number keys `1` through `4`: jump directly to a tab when no overlay or text
  input is active.
- Existing `h`/`l` drill navigation remains unchanged.
- Switching tabs resets only section-local drill focus, never run or triage
  state.

The compact footer explicitly shows `←/→ tabs`.

### Contextual footer

The footer stops showing every global key all the time. It shows:

1. tab navigation;
2. movement/drill controls relevant to the current view;
3. two or three primary actions for the active tab;
4. `? help` and `q quit`.

Expanded help becomes a real full-body panel grouped by:

- Navigation
- Run and retry
- Triage and known issues
- Reports and cleanup
- Configuration

Both `?` and Escape close expanded help.

### Status rail

The status line becomes a visually separated rail:

- left: spinner or semantic glyph plus current status;
- right: current operation name when busy;
- errors use danger color but remain redacted;
- success is concise and does not persist over newer state.

## Tab designs

### Preflight

Top metrics:

```text
READY  6      WARNINGS  1      BLOCKED  2      MODE  organization
```

Wide layout:

- Left panel: complete prerequisite checklist.
- Right panel: “Next action” summary containing the first blocking fix, config
  source state, and shortcuts for mode, variables, and rerun.

Narrow layout stacks these panels.

Each check renders:

- semantic glyph;
- check name;
- short detail;
- a visually indented fix line when present.

Empty state: “Preflight has not run yet” with `p run preflight`.

### Groups

Top metrics:

- groups;
- tests;
- passed;
- failed;
- running.

The existing aligned table remains, with:

- full-row selected surface;
- status-aware group name;
- stable pass, duration, progress, and failure columns;
- contextual action strip for drill-in, run/retry, failures-first, and show-all.

The Tests drill-down gains a breadcrumb:

```text
Groups / repositories / tests
```

The log pane gains a matching breadcrumb, an explicit redacted badge, and a
scroll-position indicator.

### Run

Run becomes the strongest mission-control screen.

Top progress card:

- overall progress bar and percentage;
- pass, fail, running, and remaining counters;
- elapsed and quiet time while running;
- clear idle, running, completed, interrupted, and failed states.

Wide layout:

- Left panel: target, scope, selected group/test, and safe command preview.
- Right panel: primary actions and last activity.

When running, the current run label and spinner lead the page. When idle, the
primary next action is explicit. No command starts automatically.

### Triage

Top metrics:

- real;
- unstable;
- flakes;
- known;
- eligible.

The list uses semantic classification badges and full-row selection. The
detail view is a bordered dossier with:

- test and classification;
- canonical signature and fingerprint;
- retry evidence;
- reasons;
- known-issue state;
- redacted log path;
- issue-filing eligibility.

Issue actions continue to require entering detail first.

## Overlays

Mode, variables, orphan list, sweep confirmation, issue confirmation, and
result messages share one panel system:

- clear title and semantic icon;
- concise description;
- content;
- bottom action bar.

Destructive confirmations use a danger border and a dedicated phrase box.
They never use success coloring before completion.

Escape closes every non-text-entry overlay. While a text field is focused,
Escape exits the field or closes the overlay according to existing editor
semantics. `ctrl+c` always quits.

## Responsive behavior

Three deterministic breakpoints:

- **Wide:** 110 columns and above. Two-column panels and full metrics.
- **Standard:** 72–109 columns. Stacked panels and full tables where possible.
- **Compact:** below 72 columns. Short product header, shortened tab labels,
  condensed metrics, no optional owner/reset metadata, and narrower progress
  bars.

Rendering must not exceed the model width after ANSI stripping. Height remains
content-driven, but long lists retain existing cursor/viewport behavior.

## Documentation

Update the standalone tester repository only. The removed
`terraform-provider-github/test/harness` docs are stale and remain untouched.

Required docs:

- README:
  - updated keyboard navigation;
  - refreshed dashboard description;
  - all five dashboard screenshots.
- `docs/quickstart.md`:
  - launching the TUI;
  - arrow, Tab, and number-key navigation.
- `docs/architecture.md`:
  - responsive chrome and pure derived-view helpers.
- `docs/troubleshooting.md`:
  - terminal width, `NO_COLOR`, and keyboard-focus notes.
- `docs/index.md`:
  - link the polished dashboard guidance.
- CHANGELOG:
  - visual polish, direct tab navigation, contextual help, responsive panels.

Regenerate deterministic synthetic screenshots:

1. ignition;
2. preflight;
3. groups;
4. run;
5. triage.

Screenshots remain synthetic, credential-free, and validated before promotion.

## Architecture

Keep the reducer and side-effect boundaries unchanged.

New pure presentation helpers may include:

- layout breakpoint selection;
- dashboard aggregate metrics;
- tab badge derivation;
- card/panel rendering;
- contextual key groups;
- status badge rendering.

No service, provider, state, lock, redaction, issue, sweep, or runner semantics
change as part of visual polish.

If `view.go` becomes harder to understand, extract focused pure files such as
`chrome.go`, `cards.go`, or `run_view.go`. Do not refactor unrelated CLI code.

## Testing

Use strict TDD for every behavior change.

Required automated coverage:

- Left/Right, Tab/Shift+Tab, and `1`–`4` navigation.
- Overlay and text-input key isolation.
- Escape closes expanded help.
- Contextual footer keys per tab and drill level.
- Tab badge derivation for empty, warning, running, and failed states.
- Wide, standard, and compact width rendering.
- No rendered line exceeds the target width after ANSI stripping.
- ASCII and color-independent status meaning.
- Existing exact confirmation and side-effect tests remain unchanged and pass.
- Existing secret-canary tests pass.
- Updated golden tests and five deterministic PNG screenshots.
- Full `go test ./...`, race tests for `cli` and `tui`, vet, docs checks, and
  build.
- Safe pseudo-TTY smoke covers direct tab navigation, help close, all tabs,
  overlays, clean quit, no secret output, and no acceptance-test child process.

## Delivery

1. Implement and review the visual polish on `tui-full-parity`.
2. Verify and merge PR #17 (`persist-env-config`).
3. Rebase only commits after `c5e431b` onto updated `origin/main`, avoiding
   duplicate persistence commits.
4. Push `tui-full-parity`.
5. Open an AI-assisted PR to `main`, include screenshots and verification
   evidence, request `@robert-crandall`, and stop for human review.

No release, tag, credentialed acceptance run, sweep, or issue filing is part of
this delivery slice.

## Acceptance criteria

1. Left/Right, Tab/Shift+Tab, and `1`–`4` switch tabs predictably.
2. Every tab has a clear title, summary state, primary next action, and
   contextual footer.
3. Wide and narrow layouts remain legible without horizontal overflow.
4. Preflight communicates ready, warning, and blocked states at a glance.
5. Groups preserves aligned progress data and improves selected-row clarity.
6. Run shows overall progress, live state, and safe next actions.
7. Triage clearly distinguishes real, unstable, flaky, known, and eligible
   failures.
8. Expanded help is grouped and closes with either `?` or Escape.
9. All overlays share a consistent panel system; destructive overlays remain
   unmistakably guarded.
10. `NO_COLOR` and ASCII output preserve all meaning without Nerd Fonts.
11. CLI, NDJSON, services, state, locks, redaction, exact confirmations, sweep,
    and issue-filing behavior remain unchanged.
12. README and docs show five current, synthetic dashboard screenshots.
13. Full verification and independent reviews are clean before the branch is
    rebased and pushed for Robert.
