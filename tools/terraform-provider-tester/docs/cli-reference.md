# terraform-provider-tester CLI reference

Use `terraform-provider-tester` to run and inspect the terraform-provider-github acceptance-test harness from a command line. The command wraps the provider's existing `go test` acceptance flow; it does not rewrite or replace any test.

```shell
terraform-provider-tester [--env-file <path>] <subcommand> [flags]
```

For automation and Copilot, pass `--json` to `e2e`, `preflight`, `groups`, `run`, `retry`, `resume`, or `orphans`. JSON mode writes newline-delimited JSON to stdout, sends human diagnostics to stderr, and bypasses the optional dashboard. `triage --json` prints one stable JSON object with `type: "triage"` for the last-run failures.

For mode requirements, see [Test mode reference](./test-modes.md). For environment variables, see [Environment variable reference](./environment-variables.md). For failures, see [Troubleshooting terraform-provider-tester](./troubleshooting.md).

## Global flags

### --env-file

Loads environment variables from a file before running any subcommand. Variables already set in the process environment are never overridden.

```shell
terraform-provider-tester --env-file .pulsar.env run --mode organization --json
```

Env-file precedence is `--env-file` > non-empty `PULSAR_ENV_FILE` > cwd `.pulsar.env` > `<os.UserConfigDir()>/terraform-provider-tester/.pulsar.env` > `<os.UserHomeDir()>/.pulsar.env` > none (load no env file, no error). For the file format and mode-scoped keys, see [Environment variable reference](./environment-variables.md#single-env-file).

## JSON output

`e2e`, `preflight`, `groups`, `run`, `retry`, `resume`, and `orphans` support `--json`. Each stdout line is one JSON object with `schema_version` and `type`. `triage --json` prints one JSON object with `version`, `type`, `mode`, `issues_repo`, and `failures`.

| Type | Meaning |
| --- | --- |
| `plan` | The resolved execution plan: selected, eligible, excluded, and unclassified counts, plus aggregated scopes, capabilities, and side effects. `preflight`, `run`, and `e2e` emit this first (`retry` and `resume` reuse the plan already persisted from an earlier `run` or `e2e` and do not re-emit it). |
| `excluded` | One test the plan will not run this mode, with a reason code, detail, optional fix, and the modes that do allow it, if any. Zero or more of these follow `plan`. |
| `check` | One preflight check from standalone `preflight --json` or from `e2e --json`'s own preflight step, including status, detail, and optional fix. |
| `capability` | One preflight check from the plan-aware preflight embedded in `run --json`. Same shape as `check`, with `type: "capability"` so callers can tell the source apart. |
| `group` | One discovered test group and its tests. |
| `test` | One completed test or subtest. |
| `orphan` | One possible leaked acceptance-test resource. |
| `orphan_delta` | Credentialed guided-run accounting with `mode`, `baseline`, `final`, `pre_existing`, `new`, `cleanup_status`, `preview_command`, and `cleanup_command`. It appears immediately before the final `e2e` summary. When `cleanup_status` is `unknown`, numeric zero fields are schema placeholders, not proven counts; `cleanup_status` is authoritative, and `preview_command` and `cleanup_command` stay empty. |
| `summary` | Final command result. Use the last one for reporting. A guided `e2e` stream embeds the run step's summary, so it can contain more than one; match on `command` (`run` or `e2e`) when the distinction matters. |

Every `--json` stream ends in a `summary` object on the runtime failure paths too: for `run` and `preflight`, a discovery or planning failure still writes a final `summary` carrying exit code `1`, without fabricating a `plan` event for a plan that was never built. Usage errors are the exception and write no JSON at all - a flag parse error exits `2` before any JSON is written because the output format is not yet known, and an invalid `--run` regex exits `2` reported on stderr only.

Exit codes are `0` for success, `1` for a valid command with failed checks or tests, and `2` for usage errors.

## Commands

Use these subcommands to validate the environment, discover tests, run tests, resume tests, triage failures, manage known issues, report results, and clean up leaked acceptance-test resources.

### e2e

Runs the guided preflight, run, one failed-test retry, triage, and orphan-check workflow for one mode at a time.

The command resolves an execution plan for the selected mode, then runs read-only preflight and stops before execution if any check fails. For a credentialed mode (`individual` or `organization`), baseline capture is read-only, and a baseline failure blocks execution. After capturing the baseline, `e2e` runs eligible tests, retries only failed tests once, and triages remaining failures without rerunning. Every attempted credentialed run gets a 30-second final orphan check, including successful, failed, and interrupted runs. A final-check failure records cleanup status `unknown` and forces a non-zero result. `anonymous` skips orphan inspection because it cannot create a resource.

Interruption does not start a new accounting window: resume reuses the original baseline. Only `New` is attributed to the run. The run-delta preview is read-only. Credentialed guided commands print preview and cleanup commands but never invoke `sweep`; `anonymous` emits no `orphan_delta` or cleanup command.

Use `groups` and `run --group` or `run --run` for selective suites and
individual acceptance tests. Selective runs still require a successful
preflight for the chosen mode.

`--mode` (default `organization`)

Sets the guided mode: `anonymous`, `individual`, or `organization`. `team` and `enterprise` are rejected before any discovery, planning, or preflight work. Start with `individual` when validating a personal setup: it needs no dedicated organization, which makes it the lowest-setup credentialed mode. It is not zero-setup - it needs `GITHUB_OWNER`, `GITHUB_USERNAME`, one auth method, and `GH_TEST_ORG_TEMPLATE_REPOSITORY` naming a template repository owned by `GITHUB_OWNER`, plus the `repo`, `delete_repo`, `user`, `admin:public_key`, `admin:gpg_key`, and `codespace` classic scopes for the complete suite. See [Prerequisites](./prerequisites.md).

`--allow-unclassified` (default `false`)

Allows the plan to proceed when the selection includes a test with no known classification, using conservative (widest) requirements and printing a warning for each one, instead of failing the plan closed. Those conservative requirements are authenticated-only, so an allowed unclassified test is still excluded in `anonymous` mode rather than run.

`--repo-root` (default `""`)

Sets the provider repository root.

`--timeout` (default `120m`)

Sets the timeout for the initial run and retry.

`--json` (default `false`)

Writes NDJSON for each workflow step: a `plan` object, zero or more `excluded` objects, zero or more `check` objects from e2e's own preflight step, `test` events, the run step's own `summary`, and a `triage` object. Credentialed modes then emit an `orphan_delta` object immediately before the final `e2e` summary; `anonymous` emits neither that event nor cleanup commands. For credentialed guided runs, the final summary repeats `baseline_orphans`, `new_orphans`, `cleanup_status`, `preview_command`, and `cleanup_command`. The command fields include the shell-quoted resolved repository root and persisted logical-run mode when cleanup status is `complete` or `baseline-only`; they stay empty when status is `unknown`. In anonymous mode, the status is `not-applicable`, the accounting counts are zero, and the command fields are empty. The final object remains the `e2e` summary.

```shell
terraform-provider-tester e2e --mode individual --json
terraform-provider-tester e2e --json
```

### preflight

Validates the environment and PAT for a test mode without creating resources.

`--mode` (default `anonymous`)

Sets the test mode to validate.

`--group` (default `""`)

Limits validation to the aggregated requirements of tests in this discovered group.

`--run` (default `""`)

Limits validation to the aggregated requirements of tests matching this regex. Overrides `--group`.

`--allow-unclassified` (default `false`)

Allows the plan to proceed when the selection includes a test with no known classification, using conservative (widest) requirements and printing a warning for each one, instead of failing the plan closed. Those conservative requirements are authenticated-only, so an allowed unclassified test is still excluded in `anonymous` mode rather than run.

`--repo-root` (default `""`)

Sets the provider repository root.

`--json` (default `false`)

Writes a JSON `plan` object, zero or more `excluded` objects, one JSON `check` object per preflight check, then a final JSON `summary`.

```shell
terraform-provider-tester preflight --mode anonymous --json
terraform-provider-tester preflight --mode organization --json
```

### groups

Lists discovered test groups and their tests.

`--unmatched` (default `false`)

Also lists tests mapped to no group and exits `1` if any unmatched tests exist.

`--repo-root` (default `""`)

Sets the provider repository root.

`--json` (default `false`)

Writes one JSON `group` object per group, followed by a final JSON `summary`.

```shell
terraform-provider-tester groups --json
terraform-provider-tester groups --unmatched --json
```

### run

Runs selected acceptance tests.

A standalone `run` starts a new logical run without guided orphan accounting.
It clears any prior guided delta before writing the new run state. Use `e2e` for
baseline capture, run-delta attribution, and finalization. An empty
`cleanup_status` means no accounting window opened before execution, for example
because planning, preflight, or lock acquisition failed.

`--mode` (default `anonymous`)

Sets the test mode for the run.

`--group` (default `""`)

Selects a discovered test group.

`--run` (default `""`)

Selects tests by regex. `--run` overrides `--group`.

`--timeout` (default `120m`)

Sets the `go test` timeout.

`--repo-root` (default `""`)

Sets the provider repository root.

`--allow-sensitive-logs` (default `false`)

Allows `TF_LOG*` environment variables to reach the child `go test` process.

`--allow-unclassified` (default `false`)

Runs tests the provider catalog does not classify, using its conservative requirements. Without this flag an unclassified test fails the plan before preflight and before the runner. An allowed unclassified test still obeys mode compatibility: the conservative requirements are authenticated-only, so it is excluded in `anonymous` mode rather than run.

`--json` (default `false`)

Writes a JSON `plan` object, zero or more `excluded` objects, zero or more `capability` preflight-check objects, then JSON `test` events as results complete, followed by a final JSON `summary`. The summary includes `go_test_command`, totals, groups, failure log paths, and `exit_code`. Stdout is JSON only.

`--tui` (default `false`)

Requests the optional interactive dashboard for a human. If stdout is not a TTY, the command falls back to text and prints a short note to stderr.

`--no-tui` (default `false`)

Forces the non-interactive text path even when stdout is a TTY.

`--triage` (default `false`)

Classifies failures after the run and stores safe triage metadata in `.pulsar-state.json`.

`--retries` (default `0`)

Retries eligible failed top-level tests one at a time. The hard cap is `2`. Setting this to a value above zero implies `--triage`.

`--retry-backoff` (default `30s`)

Initial retry backoff. Later retries use exponential backoff.

`--retry-max-backoff` (default `5m`)

Caps retry waits, including waits from `Retry-After` and `X-RateLimit-Reset` headers in failure output.

`--retry-policy` (default `retryable`)

Use `retryable` to retry only safe retryable classes, or `none` to classify without retrying.

`--retry-timeouts` (default `false`)

Allows a bare go test timeout to be retried. Without this flag, timeouts are classified but not retried.

In text mode, `terraform-provider-tester run` prints the exact `go test` command before running and streams one line per completed top-level test. In JSON mode, the command appears in the final `summary`. Failed top-level tests get redacted logs under `.pulsar-failures/<pkg>_<Test>.log`.

```shell
terraform-provider-tester run --mode anonymous --run '^TestAccGithubIpRangesDataSource$' --json
terraform-provider-tester run --mode organization --json
terraform-provider-tester run --mode organization --tui
```

### retry

Re-runs only failed top-level tests from the last run.

`--failed` (default `false`, required)

Selects failed tests from the last run. Without `--failed`, `retry` returns a usage error.

`--mode` (default: the persisted execution plan's mode)

Omit `--mode` and the retry runs in the mode recorded in the persisted execution plan. Pass it only to state the mode explicitly: a value that conflicts with the persisted plan mode is rejected with exit `2` rather than silently retrying in the wrong mode.

`--timeout` (default `120m`)

Sets the `go test` timeout.

`--repo-root` (default `""`)

Sets the provider repository root.

`--allow-sensitive-logs` (default `false`)

Allows `TF_LOG*` environment variables to reach the child `go test` process.

`--json` (default `false`)

Writes JSON `test` events as results complete, followed by a final JSON `summary`.

`--tui` (default `false`)

Requests the optional interactive dashboard for a human. If stdout is not a TTY, the command falls back to text and prints a short note to stderr.

`--no-tui` (default `false`)

Forces the non-interactive text path even when stdout is a TTY.

`--triage` (default `false`)

Classifies failures after the run and stores safe triage metadata in `.pulsar-state.json`.

`--retries` (default `0`)

Retries eligible failed top-level tests one at a time. The hard cap is `2`. Setting this to a value above zero implies `--triage`.

`--retry-backoff` (default `30s`)

Initial retry backoff. Later retries use exponential backoff.

`--retry-max-backoff` (default `5m`)

Caps retry waits, including waits from `Retry-After` and `X-RateLimit-Reset` headers in failure output.

`--retry-policy` (default `retryable`)

Use `retryable` to retry only safe retryable classes, or `none` to classify without retrying.

`--retry-timeouts` (default `false`)

Allows a bare go test timeout to be retried. Without this flag, timeouts are classified but not retried.

```shell
terraform-provider-tester retry --failed --json
terraform-provider-tester retry --failed --tui
```

Manual `retry --failed` does not reopen or finalize orphan accounting. The persisted run delta remains the snapshot from the last guided finalization. To renew accounting around another test attempt, start or resume the guided workflow instead.

### resume

Re-runs failed and not-yet-run tests from the last run.

Both `retry` and `resume` need the previous run's persisted execution plan. Only guided `e2e` installs the graceful signal handler. A graceful first interruption of guided `e2e` can persist partial results, the execution plan, and final orphan accounting before returning. When that save succeeds, `resume` uses the original baseline for final attribution rather than taking a new snapshot. SIGKILL, or interruption of `run`, `retry`, or `resume` outside that guided signal path, may leave no resumable plan. If the plan is missing, start a new run or e2e workflow to establish fresh state. State written before execution plans existed still has no persisted plan; both commands reject that legacy state with `state/plan-required`.

`--mode` (default: the persisted execution plan's mode)

Omit `--mode` and the resumed run uses the mode recorded in the persisted execution plan. Pass it only to state the mode explicitly: a value that conflicts with the persisted plan mode is rejected with exit `2` rather than silently resuming in the wrong mode.

`--timeout` (default `120m`)

Sets the `go test` timeout.

`--repo-root` (default `""`)

Sets the provider repository root.

`--allow-sensitive-logs` (default `false`)

Allows `TF_LOG*` environment variables to reach the child `go test` process.

`--json` (default `false`)

Writes JSON `test` events as results complete. A credentialed resume may therefore contain an inner run summary and a final resume summary, with a redacted `orphan_delta` between them. For automation, the last `summary` object is authoritative: its `cleanup_status` and `exit_code` include final orphan accounting and match the process outcome, including a final-check failure after the tests pass. Only the terminal resume summary is workflow-final; the inner summary describes the run phase before final accounting.

`--tui` (default `false`)

Requests the optional interactive dashboard for a human. If stdout is not a TTY, the command falls back to text and prints a short note to stderr.

`--no-tui` (default `false`)

Forces the non-interactive text path even when stdout is a TTY.

`--triage` (default `false`)

Classifies failures after the run and stores safe triage metadata in `.pulsar-state.json`.

`--retries` (default `0`)

Retries eligible failed top-level tests one at a time. The hard cap is `2`. Setting this to a value above zero implies `--triage`.

`--retry-backoff` (default `30s`)

Initial retry backoff. Later retries use exponential backoff.

`--retry-max-backoff` (default `5m`)

Caps retry waits, including waits from `Retry-After` and `X-RateLimit-Reset` headers in failure output.

`--retry-policy` (default `retryable`)

Use `retryable` to retry only safe retryable classes, or `none` to classify without retrying.

`--retry-timeouts` (default `false`)

Allows a bare go test timeout to be retried. Without this flag, timeouts are classified but not retried.

```shell
terraform-provider-tester resume --json
terraform-provider-tester resume --tui
```

### triage

Classifies the last run from `.pulsar-state.json` and `.pulsar-failures/*.log` without rerunning tests. By default, unknown real failures are shown as dry-run issue drafts and no issue is filed. Triage can optionally look up known issues from GitHub, from a local YAML cache, or from both. Use `--known-issues-offline` when triage must use only the local cache and make no network calls.

`--repo-root` (default `""`)

Sets the provider repository root.

`--known-issues` (default `""`)

Reads a YAML known-issues cache. Matching entries can suppress filing previews and set `known_issue` metadata.

`--known-issues-offline` (default `false`)

Uses only the local `--known-issues` YAML cache and makes no network calls. This disables live known-issue lookup and de-duplication calls, so it cannot be combined with `--file-issues`.

`--issues-repo` (default `integrations/terraform-provider-github`)

Sets the GitHub issue repo used for live known-issue lookup, de-duplication, dry-run previews, and optional filing.

`--file-issues` (default `false`)

Creates issues for unknown real failures only when `--confirm-file-issues` is also set. Without both flags, filing stays a dry-run.

`--confirm-file-issues` (default `false`)

Confirms non-interactive issue filing. On its own it is ignored with a warning; with `--file-issues` it allows de-duplicated issue creation.

`--json` (default `false`)

Writes one stable JSON object with `version`, `type`, `mode`, `issues_repo`, and `failures`. Each failure includes `test`, `package`, `status`, `classification`, `fingerprint`, `short_fingerprint`, `class`, `canonical`, `retryable`, `attempts`, `known_issue`, and `issue_action`.

```shell
terraform-provider-tester triage --json
terraform-provider-tester triage --known-issues .pulsar-known-issues.yaml --known-issues-offline --json
terraform-provider-tester triage --file-issues --confirm-file-issues --issues-repo integrations/terraform-provider-github
```

The same known-issue and issue-filing flags are accepted by `run`, `retry`, and `resume` when triage is enabled with `--triage` or implied by `--retries`. Do not combine `--known-issues-offline` with `--file-issues`.

### known-issues

Lists, syncs, or edits the YAML known-issues cache used by triage.

`known-issues list`

Lists known issues from the local cache, live GitHub issues, or both.

Flags: `--known-issues`, `--known-issues-offline`, `--issues-repo`, `--json`.

`known-issues sync`

Reads live GitHub issues from `--issues-repo` and writes a YAML cache.

Flags: `--issues-repo`, `--out`.

`known-issues add`

Adds or replaces one fingerprint entry in a YAML cache without contacting GitHub.

Flags: `--known-issues`, `--fingerprint`, `--issue`, `--mode`, `--test`, `--class`.

Synced entries are marked with `origin: github`; entries created by `add` are
marked with `origin: local`. Online lookup replaces GitHub-origin snapshots with
the current open-issue list and keeps local entries. `sync` also preserves local
entries already in its output file. Offline lookup uses the cached snapshot as-is.

```shell
terraform-provider-tester known-issues list --known-issues .pulsar-known-issues.yaml --known-issues-offline --json
terraform-provider-tester known-issues sync --issues-repo integrations/terraform-provider-github --out .pulsar-known-issues.yaml
terraform-provider-tester known-issues add --known-issues .pulsar-known-issues.yaml --fingerprint sha256:<fp> --issue 123 --test TestAccThing --class api/422-leftover-state
```

### report

Exports a report from the last-run state.

`--html` (default `""`)

Writes an HTML report to the path.

`--md` (default `""`)

Writes a Markdown report to the path.

`--repo-root` (default `""`)

Sets the provider repository root.

With neither `--html` nor `--md`, `report` prints a terminal summary and lists any `.pulsar-failures/*.log` files. Filtered runs preserve logs for failures they did not re-run, so this list reflects the cumulative pass/fail status. State-aware terminal, Markdown, and HTML reports show baseline, pre-existing, final, and new orphan counts plus cleanup status. State without orphan fields says `orphan accounting unavailable`.

```shell
terraform-provider-tester report --md acceptance-report.md
```

### orphans

Lists leaked `tf-acc-test-*` resources without deleting anything.

`--mode` (default: the persisted cleanup mode)

Uses the explicit cleanup mode. When omitted, `orphans` resolves the mode from the last persisted run state in this order: `orphans.mode`, execution-plan mode, then state mode. If none is available, it exits `2` with `cleanup mode is unknown; pass --mode or start a plan-bearing run`.

`--repo-root` (default `""`)

Sets the provider repository root.

`--run-delta` (default `false`)

Prints only the persisted `orphans.new` subset from state, without making a provider API call. This is a no-op preview when that persisted subset is empty. Cleanup status `unknown` rejects this operation before any provider call and directs the operator to resume final accounting or use full non-delta `orphans` for read-only inspection.

`--json` (default `false`)

Writes one `orphan` object per resource followed by an `orphans` summary. It does not change the exit behavior: finding resources succeeds, while a listing error exits `1`.

`orphans` acquires `.pulsar-state.json.lock` before loading state, is read-only, and prints leaked resources as `KIND`, `NAME`, and `URL`.

For attributed cleanup, first run the exact shape `terraform-provider-tester orphans --mode <mode> --run-delta`. This preview is read-only and uses the mode persisted for the logical run.

```shell
terraform-provider-tester orphans
terraform-provider-tester orphans --mode individual --run-delta --json
```

### sweep

Deletes leaked `tf-acc-test-*` repositories and teams.

> [!WARNING]
> `sweep` is destructive and deletes real GitHub resources.

`--confirm` (default `false`, required)

Confirms deletion. Without `--confirm`, `sweep` refuses to run and exits `1`.

`--mode` (default: the persisted cleanup mode)

Uses the explicit cleanup mode. When omitted, `sweep` resolves the mode from the last persisted run state in this order: `orphans.mode`, execution-plan mode, then state mode. If none is available, it exits `2` with `cleanup mode is unknown; pass --mode or start a plan-bearing run`.

`--repo-root` (default `""`)

Sets the provider repository root.

`--run-delta` (default `false`)

Deletes only the persisted `orphans.new` subset from state by passing an exact non-nil `SweepOpts.Resources` slice. An explicit `--mode` that disagrees with the persisted logical-run mode is rejected before any provider call when `--run-delta` is set. Cleanup status `unknown` also rejects the operation before any provider or sweep call.

Attributed cleanup requires both protections: `sweep` requires both `--run-delta` and `--confirm`. Its exact shape is `terraform-provider-tester sweep --mode <mode> --run-delta --confirm`. By contrast, a plain confirmed sweep still means all prefixed resources in the selected target kinds; it does not narrow itself to the prior run's delta. Every sweep acquires `.pulsar-state.json.lock` before loading state and remains limited to resources named `tf-acc-test-*`.

```shell
terraform-provider-tester sweep --confirm
terraform-provider-tester sweep --mode individual --run-delta --confirm
```

### unlock

Clears a stale `.pulsar-state.json.lock` left behind by a crashed or killed run.

`--repo-root` (default `""`)

Sets the provider repository root.

Most stale locks clear themselves: when a run starts, the harness reclaims a lock automatically if it is empty or if its owner process is gone on this host. `unlock` is the manual escape hatch for the cases it leaves in place on purpose, such as a lock written by a different host. Clearing an absent lock is a no-op and still exits `0`.

> [!CAUTION]
> Only run `unlock` when no other harness run is active. Clearing the lock while a run holds it can corrupt the last-run state.

```shell
terraform-provider-tester unlock
```

### version

Prints the `terraform-provider-tester` version.

This command has no flags. You can also print the version with `terraform-provider-tester --version`.

```shell
terraform-provider-tester version
```

### discover

Lists the orgs, enterprises, and, with `--org`, template repos the authenticated token can reach. Use this to find values for `GITHUB_OWNER`, `GITHUB_ENTERPRISE_SLUG`, and `GH_TEST_ORG_TEMPLATE_REPOSITORY` before running `terraform-provider-tester preflight`.

`--org` (default `""`)

When set, also lists template repositories in the given org.

`discover` degrades gracefully when the token lacks the required scope or is a fine-grained PAT: it prints a `note:` line for each unverified section and continues rather than returning an error.

```shell
terraform-provider-tester discover
terraform-provider-tester discover --org my-org
```

## Exit codes

Use these exit codes to distinguish successful runs, failing conditions, and usage errors.

| Exit code | Meaning |
| --- | --- |
| `0` | Command succeeded. For runs, no build, pre-run, panic, timeout, or failed top-level tests. |
| `1` | Valid command but failed result or runtime problem: preflight failed, build failed, tests failed, lock held, repo root missing, list failed, run start failed, `sweep` without `--confirm`, or `groups --unmatched` found unmatched tests. |
| `2` | Usage error: flag parse error, missing required flag (`retry` without `--failed`), invalid `--run` regex, invalid retry flags, a `retry`/`resume` `--mode` that conflicts with the persisted execution plan, or unknown subcommand. |

## Output selection

For `run`, `retry`, and `resume`, output selection precedence is:

1. `--json`: writes NDJSON and bypasses the dashboard.
2. `--no-tui`: forces non-interactive text.
3. `--tui`: requests the dashboard.
4. `PULSAR_NO_TUI` or `PULSAR_FORCE_TTY`: legacy env controls when no explicit flag is passed.
5. TTY auto-detect.

With no subcommand, `terraform-provider-tester` opens the dashboard on a TTY and prints help on a non-TTY.

## Optional interactive dashboard

The dashboard is for humans. After each dashboard run, failure output is written to `.pulsar-failures/<pkg>_<Test>.log` in the same redacted format as `terraform-provider-tester run`. When you relaunch the dashboard, the log pane restores prior failure content for each failed test. Use `y` to copy the visible failure log to the clipboard, and `c` to copy the exact `go test` command. Both keys use OSC 52 and are SSH-safe.

Press `m` to open the mode picker and switch auth modes without leaving the dashboard. After selecting a mode, the dashboard re-runs preflight automatically. Press `v` to open the variable editor for the current mode and set the non-secret environment variables that mode requires. Submitting a variable also re-runs preflight. Secret fields are shown as `<set>` or `<unset>` and cannot be edited from the dashboard. Press `p` to re-run preflight for the current mode at any time.

## Behavior environment variables

The following environment variables change CLI behavior. For the full environment variable list, see [Environment variable reference](./environment-variables.md).

| Environment variable | Behavior |
| --- | --- |
| `GH_TEST_AUTH_MODE` | Sets the default mode when `--mode` is not passed. The default is `anonymous`. |
| `PULSAR_FORCE_TTY=1` | Forces the dashboard path when no explicit output flag is passed. |
| `PULSAR_NO_TUI=1` | Never launches the dashboard when no explicit output flag is passed. |
| `NO_COLOR=1` | Uses ASCII glyphs and no color. It also disables the spinner braille. |
