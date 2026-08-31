# Troubleshooting terraform-provider-tester

Use this guide when `terraform-provider-tester` stops before, during, or after an acceptance-test run. For command details, see [terraform-provider-tester CLI reference](./cli-reference.md). For required configuration, see [Prerequisites](./prerequisites.md), [Test mode reference](./test-modes.md), and [Environment variable reference](./environment-variables.md).

## Preflight reports missing variables

### Cause

The selected mode requires environment variables that you have not set. `terraform-provider-tester preflight` checks the mode before the provider's tests start.

### Fix

Review the required variables for your mode. For more information, see [Prerequisites](./prerequisites.md) and [Test mode reference](./test-modes.md).

To smoke-test without credentials, use anonymous mode:

```shell
terraform-provider-tester preflight --mode anonymous --json
```

## Tests die immediately with no per-test output

### Cause

The provider's `TestMain` calls `os.Exit(1)` when a required environment variable for the mode is missing. This can stop the run before you see per-test output.

### Fix

Run `terraform-provider-tester preflight --json` before `terraform-provider-tester run --json` so you can fix missing configuration before `go test` starts. For more information, see [Environment variable reference](./environment-variables.md).

## Preflight says token capabilities are unverified

### Cause

Fine-grained PATs and GitHub Apps do not expose scopes through `X-OAuth-Scopes`. This is expected and is not an error. Classic PATs show scopes.

### Fix

Continue only if the token or GitHub App has the permissions your selected mode needs. For more information, see [Prerequisites](./prerequisites.md).

## Rate limit is low before a long run

### Cause

A full acceptance-test run makes many GitHub API calls. `terraform-provider-tester preflight` checks `GET /rate_limit` and reports rate-limit headroom before the run.

### Fix

Wait for the rate limit to reset, or use a token with more headroom before you start a long run.

## Diagnosing a failed test

### Cause

A failed top-level test can emit useful output that is too noisy for `.pulsar-state.json` and too sensitive to keep unredacted.

### Fix

Look in `.pulsar-failures/` for the failed test's redacted `.log` file instead of re-running with `go test -v`. In JSON mode, use the `failure_log_path` fields from `test` and `summary` events. `terraform-provider-tester run`, `terraform-provider-tester retry --failed`, and `terraform-provider-tester resume` refresh logs for top-level tests they re-run and preserve logs for failures they did not re-run, so the files stay consistent with the cumulative status shown by `terraform-provider-tester report`. A full run that passes clears the failure logs.


## Classify a failed run without rerunning tests

Run `terraform-provider-tester triage --json` after a failed run. It reads `.pulsar-state.json` and redacted logs under `.pulsar-failures/`, then prints stable fingerprints and classifications such as `real`, `real-unstable`, `flake-confirmed`, or `flake-historical`. By default, issue filing is a dry-run. Triage can look up known issues live or from `--known-issues`; pass `--known-issues-offline` to use only the local YAML cache and make no network calls. Filing a de-duplicated issue requires both `--file-issues` and `--confirm-file-issues`, and cannot be combined with `--known-issues-offline`.

Use `--retries N` on `run`, `retry`, or `resume` to retry only safe retryable failures before triage. `N` defaults to `0` and is capped at `2`.

The dashboard follows the same rules. `t` refreshes triage from `.pulsar-state.json` and the local cache only. Press uppercase `K` to sync live known issues into `<user-config-dir>/terraform-provider-tester/known-issues.yaml`, then refresh again.

## `terraform-provider-tester report` says planning counts are unavailable

### Cause

`.pulsar-state.json` was written before state version 2 added the `plan` field, so it has no persisted execution plan. `terraform-provider-tester report` never guesses at planning or retry counts it cannot verify from state; it prints the literal note `legacy state: planning counts unavailable` alongside the existing totals instead.

### Fix

This is expected for state left over from an older harness version, and it is not an error. `retry` and `resume` require a persisted plan, so they reject this state; run `terraform-provider-tester run` or `e2e` again to write a fresh state with a plan. After that, `report` shows planning (`selected`, `eligible`, `excluded`, `unclassified`) and retry (`retried`, `recovered`, `remaining`) counts, and `retry`/`resume` work again too.

## `resume` exits 2 with `state/plan-required`

### Cause

`retry` and `resume` re-run exactly what the previous run's persisted execution plan marked eligible, so they need a state file that already carries a version-2 `plan`. Only guided `e2e` installs the graceful signal handler. A graceful first interruption of guided `e2e` can persist partial results, the execution plan, and final orphan accounting before returning. When that save succeeds, `resume` reuses the original baseline. SIGKILL, or interruption of `run`, `retry`, or `resume` outside that guided signal path, may leave no resumable plan. This error therefore means the state predates execution plans or the prior process ended before saving its plan.

### Fix

If the plan is missing, start a new run or e2e workflow with `terraform-provider-tester run --mode <mode>` or `terraform-provider-tester e2e --mode <mode>` to establish resumable state. The rerun writes a fresh plan, after which `retry` and `resume` work normally. When resumable guided state includes orphan accounting, `resume` keeps the original baseline; it never substitutes a new baseline after interruption.

## Final orphan accounting is unknown

### Cause

After an attempted credentialed run, the 30-second final orphan check failed or timed out. The harness preserves the original baseline, records cleanup status `unknown`, and forces a non-zero result because it cannot prove the final or new counts.

### Fix

Treat the cleanup obligation as unknown. Persisted `New` is unavailable while cleanup status is `unknown`. Resolve the API or credential failure, then rerun final accounting against the original baseline:

```shell
terraform-provider-tester resume --json
```

If resume cannot complete and you need read-only inspection, list the full current prefixed set without claiming run attribution:

```shell
terraform-provider-tester orphans --mode <mode> --json
```

Do not pass `--run-delta` while cleanup status is `unknown`. Inspect the full preview before considering any action, and never infer a safe deletion set from incomplete final accounting.

## Manual retry does not refresh orphan accounting

`retry --failed` does not reopen or finalize orphan accounting. The persisted run delta remains the snapshot from the last guided finalization, even after retry runs more tests. To renew accounting around another test attempt, start or resume the guided workflow instead of treating manual retry as a new accounting window.

## `retry --failed` prints `nothing to run`

### Cause

The last-run state has no failed top-level tests for `retry --failed` to select.

### Fix

Run tests first so the harness writes last-run state. If you want to run failed tests and tests that have not run yet, use `resume`.

## A run refuses to start because of a lock

### Cause

The `.pulsar-state.json.lock` file exists because another harness run holds it, or because a previous run did not exit cleanly. The lock prevents concurrent reads and writes to `.pulsar-state.json`.

### Fix

A lock left by a crashed or killed run usually clears itself. The next run reclaims the lock automatically when it is empty or when the process that wrote it is no longer running on this host, so you can just run the command again.

If the lock persists, confirm that no other harness run is active and then clear it with `terraform-provider-tester unlock`.

```shell
terraform-provider-tester unlock
```

> [!CAUTION]
> Only run `terraform-provider-tester unlock` when no other harness run is active. Clearing the lock while a run holds it can corrupt the last-run state.

## `sweep` refuses to run

### Cause

`sweep` requires `--confirm` because it deletes leaked `tf-acc-test-*` repositories and teams.

### Fix

Use the run-delta preview first when you mean to inspect only resources attributed to the persisted logical run:

```shell
terraform-provider-tester orphans --mode <mode> --run-delta
terraform-provider-tester sweep --mode <mode> --run-delta --confirm
```

In the dashboard, press `o` first. The sweep confirmation accepts only the exact phrase `SWEEP <owner>` shown in the overlay.

The first command is read-only. Run the second only with explicit cleanup intent. A plain `terraform-provider-tester sweep --mode <mode> --confirm` is broader: it deletes all `tf-acc-test-*` resources in the provider's selected target kinds.

## Leftover `tf-acc-test-*` resources after an interrupted run

### Cause

An interrupted credentialed run can leave `tf-acc-test-*` resources behind.

### Fix

Preview only this logical run's attributed `New` resources without deleting them:

```shell
terraform-provider-tester orphans --mode <mode> --run-delta
```

Verify every previewed resource has the `tf-acc-test-` prefix. Then delete that exact persisted subset only when cleanup is intended:

```shell
terraform-provider-tester sweep --mode <mode> --run-delta --confirm
```

If the dashboard reports that the orphan snapshot changed, no deletion happened. Refresh the orphan view, confirm that the owner and resource list still match, and then confirm again.

## Dashboard issue filing refuses the confirmation phrase

### Cause

The selected failure changed, became ineligible, or the typed phrase does not match the current short fingerprint.

### Fix

Open the failure detail, press `i`, and type the exact phrase `FILE <short-fingerprint>` shown in the confirmation overlay. The dashboard reloads authoritative state, checks eligibility again, looks for a live known-issue or dedup match, and creates at most one issue before it saves state and reports success.

## Dashboard report export says the destination already exists

### Cause

Another export already wrote `<user-config-dir>/terraform-provider-tester/reports/<UTC timestamp>-report.md` and `.html` in the same second.

### Fix

The existing pair is preserved. Wait until the timestamp changes, then export again. The dashboard and CLI both write reports into `<user-config-dir>/terraform-provider-tester/reports/`.

## Optional dashboard does not open

### Cause

When you run `terraform-provider-tester` with no subcommand on a non-TTY, such as in CI or through a pipe, it prints help instead of opening the dashboard. `--json` always bypasses the dashboard.

### Fix

For automation, use `--json`. For a human dashboard, use a real TTY and pass `--tui` to `run`, `retry`, or `resume`, or set `PULSAR_FORCE_TTY=1` when no explicit output flag is passed. Set `PULSAR_NO_TUI=1` or `--no-tui` to force the non-interactive path. Set `NO_COLOR=1` to force ASCII glyphs and no color.

## Optional dashboard looks cramped or keys seem stuck

### Cause

The dashboard adapts to terminal width. Widths below 72 columns use compact layout, widths from 72 to 109 columns use standard layout, and widths of 110 columns or more use the wide mission-control layout. If an overlay is open, it owns focus until you close it. `NO_COLOR=1` also changes the dashboard to ASCII glyphs and no color.

### Fix

Widen the terminal when you want full labels and side-by-side panels. Press Left/Right or Tab/Shift+Tab to switch between Preflight, Groups, Run, and Triage. Press `1` through `4` to jump directly. Press `?` for help. Press Escape to close expanded help or the active overlay before using the underlying view's keys.

## Optional dashboard loses failure detail on quit

### Cause

The dashboard did not persist per-test failure output between sessions.

### Fix

After each dashboard run, `terraform-provider-tester` writes redacted `.pulsar-failures/<pkg>_<Test>.log` files to the provider repository root. Relaunch the dashboard to restore failure detail in the log pane for failed tests. Press `y` to copy the visible log to the clipboard or `c` to copy the exact `go test` command. All clipboard operations use OSC 52 and are SSH-safe.

If no log file is present for a failed test, the test has not yet been run in a session that writes logs, or its log was cleared by a subsequent passing run.
