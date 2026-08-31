# Changelog

All notable changes to `terraform-provider-tester`, the acceptance test harness for
terraform-provider-github, are documented here.

The format follows [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and the project aims to follow [Semantic Versioning](https://semver.org/spec/v2.0.0.html).
Releases are tagged `vX.Y.Z`.

## [Unreleased]

### Changed

- `run` no longer accepts an unclassified test implicitly. It now has an
  `--allow-unclassified` flag (default `false`) matching `preflight` and `e2e`,
  and the optional dashboard fails closed the same way. An allowed unclassified
  test is planned with the provider's conservative authenticated requirements
  and therefore runs only in a mode those requirements support.
- `retry` and `resume` derive their mode from the persisted execution plan. The
  documentation and Copilot skill no longer show a hardcoded `--mode`, which
  exits `2` whenever it disagrees with the plan the previous run recorded.
- Guided `e2e` runs the read-only orphan check only in the credentialed modes
  (`individual`, `organization`). The scan needs `GITHUB_OWNER` and an
  authenticated client, so running it after an anonymous run failed a run that
  had otherwise succeeded.
- Renamed the harness from `pulsar` to `terraform-provider-tester` so it is easy
  to find and run. The Go module is now
  `github.com/github/terraform-provider-tester` and the binary builds to
  `bin/terraform-provider-tester`. On-disk state (`.pulsar-state.json`,
  `.pulsar-failures/`, `.pulsar.env`) and the `PULSAR_*` environment variables
  keep their names for backward compatibility with existing checkouts.
- `.pulsar.env` now persists across Copilot-created worktrees through a stable
  user-config fallback. Explicit `--env-file` and `PULSAR_ENV_FILE` settings
  still take precedence over working-directory, user-config, and home files.
- The optional dashboard now uses the GH/TF Provider Acceptance identity, shows
  four tabs (Preflight, Groups, Run, and Triage), and stays in sync with the
  CLI-first workflow instead of trailing it. The repository, module, binary,
  command, and compatibility identifiers keep their existing names.
- Terraform purple is now limited to the `TF` mark and current navigation
  signal. GitHub blue remains the action and identifier color, and selected
  rows use a neutral surface instead of a full purple fill.
- The optional dashboard now supports direct tab navigation with `1` through
  `4`, Left/Right, and Tab/Shift+Tab, plus contextual help that Escape closes.
- Dashboard tabs now share responsive mission-control panels and metric strips
  across compact, standard, and wide terminal layouts.

### Fixed

- Dashboard truncation now measures terminal cells and preserves ANSI sequences,
  so color panels, wide characters, and ASCII-only fallbacks stay within the
  terminal width without broken borders or color bleed.
- Online known-issue lookup now replaces GitHub-synced cache snapshots with
  current open issues, while preserving local rules. Persisted `known`
  annotations are revalidated before de-duplication or guarded issue filing.

### Added

- Guided credentialed runs now persist one orphan baseline, reuse it across
  interruption and resume, and emit an `orphan_delta` event before the final
  `e2e` summary. Text, NDJSON, terminal, Markdown, and HTML reports distinguish
  baseline, pre-existing, final, and new resources.
- `orphans --run-delta` previews only the persisted resources attributed to the
  logical run. The `--run-delta` flag also narrows sweep to that subset.
  Deletion remains a separate command and requires both `sweep --run-delta`
  and `--confirm`; guided workflows never sweep.
- `e2e`, a guided workflow for one mode at a time (`anonymous`, `individual`,
  or `organization`; defaults to `organization`) that gates on preflight,
  retries only current failed tests once, triages remaining failures, checks
  for orphans after an attempted run in the credentialed modes, and never
  performs destructive cleanup.
- `--json` machine-readable output for `preflight`, `groups`, `run`, `retry`,
  `resume`, and `triage`, emitting stable NDJSON so Copilot and other tools can
  drive the harness without scraping text. The TUI dashboard is now opt-in via
  `--tui` (or `PULSAR_FORCE_TTY`); the CLI is the primary interface.
- Copilot skill under `.github/skills/terraform-provider-tester/` and
  `.github/copilot-instructions.md` so an agent can run preflight, groups, runs,
  retries, resume, and triage from a plain-language request.
- Personal Copilot skill setup can bootstrap and reuse the standalone tester
  from `~/.local/share/terraform-provider-tester`, allowing fresh provider
  worktrees to run preflight without rebuilding setup from scratch.
- Failure triage: `run`/`retry`/`resume --triage` and a standalone `triage`
  subcommand classify each failure with a stable fingerprint and a canonical
  error class.
- Guarded retries: `--retries N` (capped at 2) re-runs only retryable failure
  signatures before classifying, separating flakes from real failures without
  rerunning the whole suite.
- Known-issues registry: new `known-issues list|sync|add` subcommand plus
  `--known-issues <path>` and `--known-issues-offline` lookup on triage, matching
  a failure against a local YAML cache and/or open GitHub issues.
- Optional issue filing for unknown real failures, gated behind BOTH
  `--file-issues` and `--confirm-file-issues` (default is an offline-safe
  dry-run). Filed issues are de-duplicated by fingerprint, labelled
  `acctest-failure` and `ai-assisted`, and carry an AI-assisted disclosure.
- Dashboard triage actions: `t` for cache-only refresh, `K` for live
  known-issue sync, `i` for selected issue preview plus confirmation, `f` for
  failures-first group ordering, and `a` for the full group list.
- Dashboard artifact actions: `e` exports both report formats to
  `<user-config-dir>/terraform-provider-tester/reports/<UTC timestamp>-report.{md,html}`,
  `o` previews orphaned resources, and `s` opens sweep confirmation from that
  snapshot.
- Persistent user-config artifacts: live known-issue sync now writes
  `<user-config-dir>/terraform-provider-tester/known-issues.yaml`, while report
  export writes the timestamped report pair in the reports directory atomically
  with private permissions.
- Full synthetic screenshot set for the optional dashboard: intro, preflight,
  groups, run, and triage.
- Plan-aware preflight: `preflight`, `run`, and `e2e --json` resolve an
  execution plan first and emit it as a `plan` event, followed by one
  `excluded` event per test the plan will not run this mode. `--allow-unclassified`
  on `preflight` and `e2e` lets an unclassified test into the plan with
  conservative requirements instead of failing closed; `run` requires the same explicit `--allow-unclassified`
  opt-in and otherwise fails the plan closed.
  `run --json`'s embedded preflight emits `capability` events, the same shape
  as standalone `preflight`'s `check` events under a different `type`.
- `report` and the terminal summary show planning counts (selected, eligible,
  excluded, unclassified) and retry counts (retried, recovered, remaining)
  when the loaded state has a plan. State written before this addition shows
  the existing totals plus a `legacy state: planning counts unavailable` note
  instead of guessed counts.

### Security

- Issue filing never runs without explicit confirmation flags, never files a
  duplicate for a known fingerprint, and redacts secrets from every sink (issue
  title and body, dry-run output, JSON, and logs). `--known-issues-offline`
  cannot be combined with `--file-issues`.
- Dashboard actions now run one at a time, redact operation errors before they
  reach the UI, require exact `SWEEP <owner>` or `FILE <short-fingerprint>`
  phrases, and re-check the live orphan or issue state before they mutate
  anything.

### Fixed

- Report export now preserves an existing same-second Markdown and HTML pair
  instead of overwriting it.
- Issue filing now confines log excerpts to tracked redacted failure logs under
  `.pulsar-failures/`.

## [0.2.0] - 2026-07-07

### Added

- In-dashboard mode picker (`m`) and variable editor (`v`): switch auth mode or
  set non-secret env vars without leaving the dashboard. Selecting a mode or
  submitting a variable re-runs preflight automatically. Secret fields show
  `<set>`/`<unset>` and are read-only; raw secret values are never displayed.
- `p` key re-runs preflight for the current mode from anywhere in the dashboard.
- CI workflow that runs gofmt, vet, build, unit tests, the docs drift check, and
  a gitleaks secret scan on every pull request.
- Repository hygiene: `CONTRIBUTING.md`, `SECURITY.md`, issue templates, and a
  pull request template.
- `--env-file <path>` global flag loads variables from a dotenv file before any
  subcommand runs. `.pulsar.env` in the working directory is loaded automatically
  when present. Variables already in the process environment are never overridden.
  Mode-scoped keys (`PULSAR_<MODE>_KEY`) let one file serve every mode and resolve
  the `GITHUB_OWNER` collision between `individual` and `organization` modes.
- `pulsar discover` lists accessible orgs and enterprises so users can pick values for `GITHUB_OWNER`, `GITHUB_ENTERPRISE_SLUG`, and `GH_TEST_ORG_TEMPLATE_REPOSITORY`. Degrades gracefully with a `note:` line when the token lacks enumeration scope.
- TUI dashboard now writes redacted `.pulsar-failures/<pkg>_<Test>.log` files
  after each run, matching the behavior of `pulsar run`. Relaunching the
  dashboard restores prior failure output in the log pane for failed tests.
- `y` (copy log) and `c` (copy cmd) keys in the dashboard now copy the visible
  failure log or the exact `go test` command to the clipboard via SSH-safe
  OSC 52. A status-line message confirms the action.

### Fixed

- Individual-mode environment editing now exposes the template repository
  variable when a selected test requires it.
- JSON run summaries no longer count mode-incompatible tests in group totals.
- Env-file mode-scoped keys now match the active mode case-insensitively, so a
  `PULSAR_organization_GITHUB_OWNER` written with a lowercase mode is applied
  instead of being silently dropped.
- Env-file mode scoping now follows the same mode the run resolves. When the
  shell and the file both set `GH_TEST_AUTH_MODE`, the shell value wins for
  scoping too, so the file no longer promotes keys for a mode the run does not use.
- Env-file now reserves only real mode names after `PULSAR_`, so behavior
  variables such as `PULSAR_FORCE_TTY` and `PULSAR_NO_TUI` can be set through the
  single env file instead of being ignored.
- Stale run-lock reclaim no longer removes a replacement lock on Linux. The lock
  identity now includes a content fingerprint, so a reused inode at the same
  path is not mistaken for the file judged stale.
- `y` (copy log) and `c` (copy cmd) keys were stubbed no-ops; they now work.
- TUI runs lost failure detail on quit; failure logs are now persisted to disk.
- Mode-picker overlay: a terminal narrower than 20 columns no longer overflows
  the description line (truncation is now clamped to at least 1 rune).
- `ctrl+c` now always quits from the mode-picker and variable-editor overlays
  instead of just closing them.
- `p` now re-runs preflight from inside the variable editor (nav mode).
- Variable editor refuses to set secret env vars via `SetEnvVarIntent`,
  preventing accidental exposure through that path.

## [0.1.0] - 2026-06-26

### Added

- Preflight validates PAT auth mode, scopes, and rate limits before test runs.
- Groups discovers and buckets `TestAccGithub*` tests by resource.
- Run executes grouped `go test ./github -json` and tracks per-group pass and fail.
- Retry and resume support single-test, per-group reruns, and cross-session resume.
- Report exports shareable Markdown and HTML run summaries.
- Orphans and sweep add orphan detection plus confirm-gated cleanup.
- State and logs add redaction, minimal persisted run state, stale-lock recovery,
  and redacted failure log capture.
- TUI adds the Pulsar dashboard, spinner, and live run status UI.
- Command rename: user-facing identifiers moved from `tfacc` to `pulsar`,
  including command name, binary path, state files, and environment variables.

[Unreleased]: https://github.com/github/terraform-provider-tester/compare/v0.2.0...HEAD
[0.2.0]: https://github.com/github/terraform-provider-tester/releases/tag/v0.2.0
[0.1.0]: https://github.com/github/terraform-provider-tester/releases/tag/v0.1.0
