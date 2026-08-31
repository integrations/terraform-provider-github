# Terraform Provider Tester TUI Full Parity Design

**Status:** Approved by delegated autonomy on 2026-07-09
**Delivery branch:** `tui-full-parity`, stacked on `persist-env-config`
**Primary interface:** CLI and NDJSON remain the automation interface; the TUI remains optional

## Problem

The dashboard exposes preflight, group navigation, test logs, retry, resume, mode selection, and environment editing. It does not expose the newer triage, known-issue, report, orphan, sweep, or issue-filing workflows. Several keys appear in help but do nothing. The header and banner use the Terraform Provider Tester name, but the startup animation still spells `PULSAR`.

This mismatch makes the TUI look complete while hiding current capabilities and advertising inactive actions.

## Goals

1. Replace visible Pulsar branding with Terraform Provider Tester branding.
2. Show persisted triage classifications and known-issue results.
3. Let a user refresh triage, sync known issues, export reports, list orphans, sweep confirmed resources, and file one confirmed issue from the TUI.
4. Reuse the same behavior as the CLI instead of copying command logic.
5. Preserve redaction, state safety, retry behavior, and destructive-action gates.
6. Keep generated reports and caches outside provider worktrees.
7. Make every advertised key functional.

## Non-goals

- Replace the CLI or NDJSON workflow.
- Rename compatibility identifiers such as `.pulsar.env`, `.pulsar-state.json`, `.pulsar-failures`, `PULSAR_*`, `pulsar-fp-v1`, or known-issue markers.
- Change acceptance tests or their grouping rules.
- Add automatic retries beyond the existing manual retry and resume controls.
- File several issues in one action.
- Run cleanup without a fresh orphan preview and an exact confirmation phrase.
- Change provider behavior unrelated to dashboard actions.

## Chosen Approach

Add a fourth Triage tab and implement report and cleanup actions as overlays. Keep the `tui` package pure. The dashboard producer performs asynchronous work through structured services shared with CLI handlers.

This approach adds the requested capabilities without turning the TUI into a second command implementation. A separate Cleanup tab would spend permanent screen space on occasional work. Folding triage into Run would make failure detail unreadable. A fourth Triage tab and temporary overlays preserve the current navigation model.

## Branding

The startup animation will render `TESTER`, which has the same six-letter width as `PULSAR`. The existing deterministic animation and terminal dimensions remain unchanged. The implementation will add `T` and `E` glyphs, remove unused visible-brand glyphs, and update comments and golden fixtures.

The full product name, `Terraform Provider Tester`, remains in the header, banner, and startup subtitle. The beacon graphic may remain as a neutral test-signal motif, but code and user-facing text will call it the tester mark rather than the Pulsar mark.

Compatibility identifiers retain their existing names. Renaming them would break env discovery, saved state, failure logs, fingerprints, known-issue matching, and existing automation.

## User Interface

### Tabs

The dashboard will have four tabs:

1. **Preflight** shows credential and environment checks.
2. **Groups** shows groups, tests, and logs.
3. **Run** shows the current target, scope, counts, command shape, and run actions.
4. **Triage** shows classified failures and known-issue state.

Tab and shift-tab continue to cycle through all tabs.

### Triage Tab

The top of the Triage tab shows these counts:

- total classified failures
- confirmed or historical flakes
- real failures
- unstable real failures
- known or deduplicated issues
- unknown failures eligible for issue filing

The failure list shows:

- test name
- classification
- failure class
- attempt count
- short fingerprint
- known issue number, when present
- issue action: `known`, `dedup`, `filed`, or `dry-run`. The TUI never runs a
  dry-run filing mode; it shows `eligible` for an unfiled failure that
  passes the filing rules, corresponding to the CLI's `dry-run` outcome.

Enter opens a detail view with the canonical signature, retryability, reasons, log path, and issue target. The detail view never includes raw unredacted output.

### Key Actions

| Key | Scope | Action |
| --- | --- | --- |
| `f` | Groups | Order failed and running groups or tests first without hiding other entries |
| `a` | Groups | Restore complete discovery order |
| `t` | Triage | Reconstruct or refresh last-run triage without rerunning tests |
| `K` | Triage | Sync known issues to the persistent cache, then rematch current failures |
| `i` | Triage detail | Preview and confirm filing for the selected eligible failure |
| `e` | Any main tab | Export timestamped Markdown and HTML reports from last-run state |
| `o` | Any main tab | List orphaned `tf-acc-test-*` resources in a read-only overlay |
| `s` | Orphan overlay | Review and confirm a sweep |

Existing retry, resume, preflight, mode, variable, copy, navigation, help, and quit keys keep their behavior.

## Architecture

### Pure TUI Boundary

The `tui` package owns:

- sections and view state
- list and detail cursors
- overlays and text inputs
- intents emitted by key actions
- structured messages received from the producer
- deterministic rendering

It performs no filesystem, network, provider, GitHub issue, or test-run work.

New intents will represent user requests, not implementation details:

- `RefreshTriageIntent`
- `SyncKnownIssuesIntent`
- `ExportReportIntent`
- `ListOrphansIntent`
- `ConfirmSweepIntent`
- `ConfirmFileIssueIntent`

New messages will carry typed, already-redacted results:

- triage started and completed
- known-issue sync completed
- report export completed
- orphan listing completed
- sweep completed
- issue preview and filing completed
- operation failure

### Shared Services

Current CLI flag handlers mix parsing, orchestration, output formatting, persistence, and side effects. The implementation will extract internal services in the `cli` package. CLI handlers and the dashboard producer will call the same services.

The services will cover:

1. **Triage last run:** load state, reconstruct missing failures, match known issues, persist updated failures, and return structured results.
2. **Known-issue sync:** fetch issue entries and atomically write a cache file.
3. **Report export:** reconstruct the last run, write redacted Markdown and HTML reports, and return their paths.
4. **Orphan listing:** call the provider and return structured resources.
5. **Sweep:** validate a confirmed resource snapshot, call the provider with `Confirm: true`, and return the refreshed orphan list.
6. **Issue filing:** build a redacted preview, deduplicate immediately before filing, file one selected failure, and persist the resulting issue action.

Text and NDJSON formatting remain adapters around these structured results. The extraction must preserve current CLI output and exit behavior.

### Environment Consistency

The dashboard will use the existing synchronized environment overlay. Env-file values and successful TUI edits update both the overlay and the production process environment before publication. Provider cleanup methods and GitHub clients therefore see the same credentials and owner as preflight and test runs.

Services will receive the dashboard's environment accessor and redactor. Tests will inject both. The implementation will not read or display secret values.

### Operation Gate

The dashboard will replace the run-only busy guard with one operation gate. Only one run, triage refresh, sync, export, orphan request, sweep, or filing action may execute at a time.

The TUI disables conflicting actions while an operation runs. The producer enforces the same rule, so delayed or synthetic messages cannot bypass it. Mode and environment changes remain serialized through their existing configuration lock.

## Data Flow

### Dashboard Startup

1. Run preflight and send its report.
2. Send discovered groups.
3. Load persisted test results and safe failure logs.
4. Send persisted triage failures.
5. Offer resume when the prior state contains failed or unrun tests.

### Test Run

1. Resolve the selected tests and acquire the state lock.
2. Run tests and stream redacted test updates.
3. Merge selected results into state, replacing stale results for those tests.
4. Classify current failures with zero automatic retries.
5. Remove stale failure records for selected tests that now pass.
6. Match known issues from the persistent cache without an implicit network call.
7. Save state atomically.
8. Send run completion and updated triage messages.

Manual retry and resume remain the mechanisms for rerunning tests. Historical mixed pass and fail results may produce the existing historical-flake classification.

### Triage Refresh

1. Load last-run state.
2. Use persisted failures or reconstruct them from state and redacted failure logs.
3. Match the persistent known-issue cache.
4. Save updated failures.
5. Refresh the Triage tab.

No acceptance test or network request runs during this flow. A missing cache yields
unmatched failures and a non-fatal status note.

### Known-Issue Sync

1. Resolve the issues repository to `integrations/terraform-provider-github`.
2. Fetch known issues through the existing GitHub issue client.
3. Write the YAML cache atomically under the OS user config directory:
   `terraform-provider-tester/known-issues.yaml`.
4. Rematch current failures against the fresh cache.
5. Show the entry count and cache path.

A failed sync preserves the last good cache and current triage view. Sync is the
explicit live-network action; routine runs and triage refreshes use the cache.

### Report Export

1. Load last-run state and reconstruct the run result.
2. Create `terraform-provider-tester/reports` under the OS user config directory.
3. Export both Markdown and HTML with a UTC timestamp in each filename.
4. Write through temporary files and rename them into place.
5. Show both paths in the result overlay.

Reports remain outside the provider worktree and never make it dirty.
The implementation resolves this directory through the injected user-config
dependency, creates private directories, and writes private files.

### Orphan Listing and Sweep

1. `o` fetches and displays orphan resources without changing them.
2. `s` opens a confirmation view only after a successful listing.
3. The view displays the owner, resource count, kinds, names, and URLs.
4. The user must type `SWEEP <owner>`.
5. The TUI validates the phrase, then emits a confirmation intent.
6. The producer validates the phrase again and re-lists resources.
7. The producer compares sorted `{kind, name, URL}` identities. If the set changed,
   it aborts and returns the new list for another review.
8. If the set matches, the producer calls `Sweep` with `Confirm: true`.
9. The producer re-lists resources and shows any residual resources.

The action remains unavailable during a test run or any other operation.

### Issue Filing

1. `i` works only from a selected Triage detail record.
2. The failure must be real or real-unstable, unknown, and eligible under existing filing rules.
3. The preview shows the issues repository, redacted title, labels, classification, and short fingerprint.
4. The user must type `FILE <short-fingerprint>`.
5. The TUI and producer validate the phrase.
6. The producer performs a live fingerprint lookup and issue search again.
7. A new match changes the action to `known` or `dedup` and cancels creation.
8. Otherwise, the producer files one issue, records its number and `filed` action, saves state, and refreshes Triage.

Selected-only filing avoids ambiguous partial success across several issue creations.
The preview omits the issue body; it shows only redacted metadata needed to make
the decision. If GitHub creates the issue but state persistence fails, the TUI
reports an error rather than completion. A retry performs live deduplication and
recovers the created issue number.

## Safety

### Secrets

- TUI messages carry structured values, not environment maps or raw subprocess output.
- The producer redacts test output before sending it.
- Reports and issue drafts use the existing redactor.
- Secret editor fields remain masked and non-readable.
- Error text passes through redaction before display.

### Destructive and External Actions

- Orphan listing, triage refresh, report export, and known-issue sync need no confirmation.
- Sweep requires a current orphan snapshot and `SWEEP <owner>`.
- Filing requires an eligible selected failure and `FILE <short-fingerprint>`.
- Both confirmation flows validate in the TUI and producer.
- Filing repeats deduplication immediately before creation.
- Sweep repeats discovery immediately before deletion.

### State

- Test runs continue to use the existing state lock.
- Triage and filing update state atomically.
- Local report and cache writes use temporary files and rename.
- Read-only actions preserve prior good model state on failure.
- Compatibility state and fingerprint identifiers remain unchanged.

## Error Handling

Every asynchronous action sends an operation-specific error. The TUI keeps prior results visible and adds a concise failure status. It does not convert failures into empty success results.

Sweep errors trigger a best-effort read-only refresh so the user can see remaining resources. The original sweep error remains visible even if refresh succeeds. If refresh also fails, the TUI reports both errors.

Issue filing handles deduplication as a normal result, not an error. Network, authorization, persistence, and issue-creation failures remain errors. A successfully created issue must be persisted before the TUI reports completion.

## Testing

Implementation will follow test-driven development.

### TUI Tests

- tab cycling with four sections
- TESTER wordmark at start, peak, and exit frames
- Triage empty, summary, list, and detail golden fixtures
- failure-focused and discovery-order views
- all new intent emissions
- operation-busy behavior
- wrong, partial, and exact confirmation phrases
- stale orphan snapshot cancellation
- ineligible issue-filing rejection
- narrow and ASCII terminal rendering

### Dashboard and Service Tests

- startup loads persisted failures
- TUI run replaces stale selected failures
- triage reconstruction and known-issue matching
- sync preserves the previous cache on failure
- reports use redaction and stay outside the provider root
- orphan listing remains read-only
- sweep receives `Confirm: true` only after both validations
- changed orphan snapshots abort sweep
- issue filing deduplicates immediately before creation
- filing updates state before sending completion
- the operation gate prevents concurrent actions
- environment overlay values reach provider and issue operations

### Regression and Safety Tests

- existing CLI text and NDJSON behavior
- existing report, triage, retry, resume, known-issue, orphan, and sweep tests
- `go test -race ./cli ./tui`
- full `go test ./...`
- `go vet ./...`
- `gofmt` and documentation consistency checks
- pseudo-TTY dashboard smoke test with fake data
- secret canary checks across messages, state, reports, previews, and errors

Credentialed acceptance tests require a dedicated test organization and scoped credentials. Without them, verification stops at preflight and uses fake provider integrations for side-effect flows.

## Documentation

Update:

- README dashboard section and key summary
- docs index and architecture description
- troubleshooting for confirmation and cache paths
- changelog with visible rebranding and TUI parity
- dashboard screenshots using synthetic data

Documentation will continue to present CLI and NDJSON first.

## Acceptance Criteria

1. No visible TUI screen spells `PULSAR`.
2. Compatibility files, env vars, and fingerprint markers remain readable.
3. Every key shown in help performs its documented action.
4. Triage state appears after startup, runs, retries, resume, refresh, sync, dedup, and filing.
5. Reports and known-issue cache persist across provider worktrees without dirtying them.
6. Orphan listing never deletes resources.
7. Sweep cannot run without a fresh matching snapshot and exact owner phrase.
8. Issue creation cannot run without an eligible selected failure, fresh deduplication, and exact fingerprint phrase.
9. Secrets do not appear in TUI messages, state, reports, previews, errors, or committed fixtures.
10. CLI text and NDJSON behavior remain compatible.
11. Tests, race checks, vet, formatting, and documentation checks pass.
