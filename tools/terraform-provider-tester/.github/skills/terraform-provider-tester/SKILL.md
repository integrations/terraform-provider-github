---
name: terraform-provider-tester
description: Use when the user asks to run terraform-provider-github acceptance tests, run an individual or organization e2e test, run a specific TestAcc group, check acceptance-test prerequisites, retry failed acceptance tests, resume an interrupted run, or triage a failing acceptance test.
version: 0.1.0
---

# Terraform Provider Tester

Use Terraform Provider Tester to run terraform-provider-github acceptance tests from any provider worktree. The tester is a standalone private repo, not vendored provider code. For Copilot and automation, prefer `--json`, which writes newline-delimited JSON and bypasses the optional dashboard.

## Safety rules

- Acceptance tests can create and delete real GitHub resources. Run credentialed modes only when the user asked for them or the environment is dedicated to acceptance testing.
- Respect the provider AI Use Policy. This is AI-assisted work for human review. Do not open an upstream provider PR from tester output unless a human understands and validates it.
- Never print secrets. Do not cat `.pulsar.env`. Do not echo tokens or PEM contents. Do not use `--allow-sensitive-logs` unless the user explicitly asks.
- Default scripted commands must omit `--env-file`. The CLI auto-detects, in order, non-empty `PULSAR_ENV_FILE`, cwd `.pulsar.env`, user config `terraform-provider-tester/.pulsar.env`, home `.pulsar.env`, then no env file with no error. Pass an explicit env file only when the user names one.
- Preflight the selected mode with JSON must pass before any resource-creating acceptance test. Individual mode needs no dedicated organization, so prefer it for a first credentialed check; organization mode is `e2e`'s default when the user does not name a mode. If preflight exits non-zero, stop and report every emitted check, status, detail, and fix.
- Never fall back to raw `make testacc`, `gotestsum`, or provider make targets when the tester is unavailable. Fix or report tester bootstrap instead.
- Preview before suggesting cleanup. Resolve `RUN_MODE` from the persisted logical run (prefer `orphan_delta.mode`); never guess it from ambient credentials. Use the run-delta preview below first.
- When cleanup status is `unknown`, persisted `New` is unavailable. Rerun `resume` to recompute final accounting against the original baseline, or use full non-delta `orphans` for read-only inspection. Do not use `--run-delta` until final accounting is complete.
- Never execute cleanup without explicit user intent. Guided `e2e` and `resume` never sweep, and an attributed sweep requires both `--run-delta` and `--confirm`.
- Triage defaults to dry-run issue filing. Creating issues requires both `--file-issues` and `--confirm-file-issues`; use `--known-issues-offline` for local-only known-issue lookup, not with filing.

## Bootstrap from any provider worktree

Run this once at the start of a fresh terraform-provider-github session. It locates the provider root, resolves an absolute tester binary path, and clones or builds the private tester repo only under `$HOME/.local/share/terraform-provider-tester`.

```sh
set -eu

normalize_existing_dir() {
  path="$1"
  label="$2"
  case "$path" in
    /*) ;;
    *) echo "$label must be an absolute path: $path" >&2; exit 1 ;;
  esac
  test -d "$path" || { echo "$label is not a directory: $path" >&2; exit 1; }
  (cd -P "$path" && pwd)
}

normalize_child_path() {
  path="$1"
  label="$2"
  case "$path" in
    /*) ;;
    *) echo "$label must be an absolute path: $path" >&2; exit 1 ;;
  esac
  parent="$(dirname "$path")"
  base="$(basename "$path")"
  test -d "$parent" || { echo "$label parent is not a directory: $parent" >&2; exit 1; }
  parent_real="$(cd -P "$parent" && pwd)" || { echo "cannot normalize $label parent: $parent" >&2; exit 1; }
  printf '%s/%s\n' "$parent_real" "$base"
}

fail_invalid_tpt() {
  echo "Invalid tester checkout at $TPT_HOME: $1. Move it aside or reinstall github/terraform-provider-tester there." >&2
  exit 1
}

prepare_tpt_home_path() {
  if test -e "$TPT_HOME" && ! test -d "$TPT_HOME"; then
    echo "Tester cache exists but is not a directory: $TPT_HOME" >&2
    exit 1
  fi
  if test -d "$TPT_HOME"; then
    TPT_HOME_REAL="$(normalize_existing_dir "$TPT_HOME" "tester cache")"
  else
    mkdir -p "$(dirname "$TPT_HOME")"
    TPT_HOME_REAL="$(normalize_child_path "$TPT_HOME" "tester cache")"
  fi
  if test "$TPT_HOME_REAL" = "$PROVIDER_ROOT"; then
    echo "Tester cache must not be the provider checkout; use $HOME/.local/share/terraform-provider-tester." >&2
    exit 1
  fi
}

validate_tpt_checkout() {
  TPT_HOME_REAL="$(normalize_existing_dir "$TPT_HOME" "tester cache")"
  TPT_TOP="$(git -C "$TPT_HOME" rev-parse --show-toplevel 2>/dev/null)" || fail_invalid_tpt "not a git checkout"
  TPT_TOP="$(normalize_existing_dir "$TPT_TOP" "tester git root")"
  test "$TPT_TOP" = "$TPT_HOME_REAL" || fail_invalid_tpt "path is not the checkout root"
  TPT_MODULE="$(sed -n 's/^module //p' "$TPT_HOME/go.mod" 2>/dev/null)" || fail_invalid_tpt "cannot read go.mod"
  test "$TPT_MODULE" = "github.com/github/terraform-provider-tester" || fail_invalid_tpt "go.mod must declare module github.com/github/terraform-provider-tester"
  TPT_ORIGIN="$(git -C "$TPT_HOME" config --get remote.origin.url 2>/dev/null)" || fail_invalid_tpt "missing origin remote"
  case "$TPT_ORIGIN" in
    https://github.com/github/terraform-provider-tester|https://github.com/github/terraform-provider-tester.git|git@github.com:github/terraform-provider-tester|git@github.com:github/terraform-provider-tester.git|ssh://git@github.com/github/terraform-provider-tester|ssh://git@github.com/github/terraform-provider-tester.git) ;;
    *) fail_invalid_tpt "origin must be github/terraform-provider-tester" ;;
  esac
  test -f "$TPT_HOME/Makefile" || fail_invalid_tpt "missing Makefile"
  grep -Eq '^build:' "$TPT_HOME/Makefile" || fail_invalid_tpt "missing Makefile build target"
}

normalize_path_binary() {
  TPT_PATH="$1"
  case "$TPT_PATH" in
    */*) ;;
    *) echo "terraform-provider-tester on PATH did not resolve to a file path" >&2; exit 1 ;;
  esac
  TPT_PATH_DIR="$(dirname "$TPT_PATH")"
  TPT_PATH_BASE="$(basename "$TPT_PATH")"
  TPT_PATH_DIR="$(cd -P "$TPT_PATH_DIR" && pwd)" || { echo "cannot normalize terraform-provider-tester from PATH" >&2; exit 1; }
  TPT_BIN="$TPT_PATH_DIR/$TPT_PATH_BASE"
  test -x "$TPT_BIN" || { echo "terraform-provider-tester from PATH is not executable after normalization" >&2; exit 1; }
}

PROVIDER_ROOT="$(git rev-parse --show-toplevel)"
PROVIDER_ROOT="$(normalize_existing_dir "$PROVIDER_ROOT" "provider root")"
test -f "$PROVIDER_ROOT/go.mod"
test -d "$PROVIDER_ROOT/github"

case "$HOME" in
  /*) ;;
  *) echo "HOME must be an absolute path for tester bootstrap" >&2; exit 1 ;;
esac

TPT_HOME="$HOME/.local/share/terraform-provider-tester"
case "$TPT_HOME" in
  /*) ;;
  *) echo "Tester cache path must be absolute: $TPT_HOME" >&2; exit 1 ;;
esac
TPT_BIN=""

if TPT_ON_PATH="$(command -v terraform-provider-tester 2>/dev/null)"; then
  normalize_path_binary "$TPT_ON_PATH"
else
  prepare_tpt_home_path
  if test -d "$TPT_HOME"; then
    validate_tpt_checkout
  else
    command -v gh >/dev/null
    gh repo clone github/terraform-provider-tester "$TPT_HOME" -- --depth 1
    validate_tpt_checkout
  fi
  if ! test -x "$TPT_HOME_REAL/bin/terraform-provider-tester"; then
    command -v go >/dev/null
    make -C "$TPT_HOME" build
  fi
  TPT_BIN="$TPT_HOME_REAL/bin/terraform-provider-tester"
fi

test -x "$TPT_BIN"
```

Do not mutate global Git credential config. Do not use `go install` for this private repo. Do not delete, overwrite, or build an invalid directory at `$HOME/.local/share/terraform-provider-tester`.

## Standard JSON workflow

When the user says "Run the individual e2e test," or asks for a first credentialed check without naming a mode, run individual mode first - it is the lowest-setup credentialed mode because it needs no dedicated organization:

```sh
"$TPT_BIN" e2e --repo-root "$PROVIDER_ROOT" --mode individual --json
```

Individual mode is not setup-free. It needs `GITHUB_OWNER`, `GITHUB_USERNAME`, one auth method, and `GH_TEST_ORG_TEMPLATE_REPOSITORY` naming a template repository owned by `GITHUB_OWNER`. A classic PAT needs the `repo`, `delete_repo`, `user`, `admin:public_key`, `admin:gpg_key`, and `codespace` scopes, because the user SSH key, user GPG key, and user Codespaces tests stay enabled; fine-grained PATs and GitHub Apps need the equivalent permissions, which preflight reports as unverified. Let preflight report what is missing instead of guessing.

When the user says "Run the organization e2e test," or once individual mode is clean and the user wants full resource coverage, run organization mode. `--mode` defaults to `organization` when omitted:

```sh
"$TPT_BIN" e2e --repo-root "$PROVIDER_ROOT" --json
```

Either command performs the preflight gate for that mode, its acceptance-test run, one failed-test retry, remaining-failure triage, and read-only orphan accounting (credentialed modes only; `anonymous` makes no provider orphan calls and emits no `orphan_delta` or cleanup command). Credentialed execution captures one baseline before tests; failure blocks execution. The cancellation-independent final check has a 30-second timeout. If preflight exits non-zero, report every emitted check and fix; the command will not have attempted tests. For credentialed modes, read `orphan_delta` immediately before the last `summary`; always use that final summary for the overall workflow result. Do not paste full logs unless the user asks, and still redact secrets.

If a command fails during planning because a test has no requirements classification, report it rather than working around it: the harness fails closed on purpose. Re-run with `--allow-unclassified` only when the user asks to include the unknown test; `e2e`, `preflight`, and `run` all accept the flag and then plan it with conservative authenticated requirements, which means it still does not run in `anonymous` mode.

Use `"$TPT_BIN" known-issues list|sync|add` to inspect, refresh, or edit the known-issues cache.

If the run was interrupted or some tests did not run, resume it. Only guided `e2e` installs the graceful signal handler. A graceful first interruption of guided `e2e` can persist partial results, the execution plan, and final orphan accounting before returning. When that save succeeds, guided resume reuses the original baseline; never recapture or replace it. SIGKILL, or interruption of `run`, `retry`, or `resume` outside that guided signal path, may leave no resumable plan. If the plan is missing, start a new run or e2e workflow. Omit `--mode`: `retry` and `resume` take the mode from that persisted plan, so they always continue in the mode the interrupted run actually used. Pass `--mode` only when the user names one explicitly, and only when it matches that plan - a conflicting value exits `2`.

```sh
"$TPT_BIN" resume --repo-root "$PROVIDER_ROOT" --json
"$TPT_BIN" retry --repo-root "$PROVIDER_ROOT" --failed --json
```

Both commands need the previous run's persisted execution plan. Legacy state written before execution plans existed has no v2 plan, so both commands exit `2` with `state/plan-required` without starting a test. Rerun `run` or `e2e` for the requested mode to establish resumable state.

## NDJSON interpretation

`--json` writes one JSON object per stdout line. Each object has `schema_version` and `type`.

Event types:

- `plan`: the resolved execution plan, emitted first by `preflight`, `run`, and `e2e`. Fields include `selected`, `eligible`, `excluded`, `unclassified`, `scopes`, `capabilities`, and `side_effects`.
- `excluded`: one test the plan will not run this mode. Fields include `test`, `code`, `detail`, `fix`, and `required_modes`. Zero or more follow `plan`.
- `test`: one completed test or subtest. Fields include `name`, `sub`, `group`, `status`, `elapsed`, and `failure_log_path`.
- `check`: one preflight check from standalone `preflight` or from `e2e`'s own preflight step. Fields include `name`, `status`, `detail`, and `fix`.
- `capability`: one preflight check from `run`'s embedded plan-aware preflight. Same fields as `check`, under a different `type`.
- `group`: one discovered test group. Fields include `name`, `tests`, and `total`.
- `orphan`: one possible leaked acceptance-test resource.
- `orphan_delta`: credentialed guided-run `mode`, baseline/final/pre-existing/new counts, `cleanup_status`, `preview_command`, and `cleanup_command`. It is immediately before the final `e2e` summary. When `cleanup_status` is `unknown`, numeric zero fields are schema placeholders, not proven counts; `cleanup_status` is authoritative, and both command fields are empty.
- `summary`: final command result. Use the last `summary` object for the final answer.

Run summary fields include `totals`, `groups`, `build_failed`, `pre_run_failed`, `failures`, `go_test_command`, and `exit_code`.

Ignore unknown NDJSON event types so future event additions do not break parsing.

## Orphan preview and explicit cleanup

Set `RUN_MODE` to the persisted logical-run mode from `orphan_delta.mode` or state. Preview the attributed `New` resources:

```sh
"$TPT_BIN" orphans --repo-root "$PROVIDER_ROOT" --mode "$RUN_MODE" --run-delta --json
```

Inspect every resource and verify its name starts with `tf-acc-test-`. Report the preview before suggesting deletion. Only when the user explicitly asks to delete that attributed set, run:

```sh
"$TPT_BIN" sweep --repo-root "$PROVIDER_ROOT" --mode "$RUN_MODE" --run-delta --confirm
```

The first command is read-only. The second is destructive. A plain confirmed sweep omitting `--run-delta` means all prefixed resources in the selected target kinds, so never substitute it for attributed cleanup.

Exit codes:

- `0`: success.
- `1`: valid command with failed result or runtime problem, such as preflight failure, test failure, build failure, lock error, repo root error, list error, or run start error.
- `2`: usage error, such as bad flags, missing `retry --failed`, invalid regex, a `retry`/`resume` `--mode` that conflicts with the persisted execution plan, or unknown subcommand.

## Natural language to command map

Always use the resolved `"$TPT_BIN"` and `--repo-root "$PROVIDER_ROOT"` for commands that read provider source, state, failures, reports, or leaked resources.

| Ask | Command sequence |
| --- | --- |
| Run the individual e2e test | `"$TPT_BIN" e2e --repo-root "$PROVIDER_ROOT" --mode individual --json`. |
| Run the organization e2e test | `"$TPT_BIN" e2e --repo-root "$PROVIDER_ROOT" --json`. |
| Run the organization group | Preflight organization, then `"$TPT_BIN" run --repo-root "$PROVIDER_ROOT" --mode organization --group organization --json`. |
| Run the issue label tests | Preflight organization, then `"$TPT_BIN" run --repo-root "$PROVIDER_ROOT" --mode organization --run '^TestAccGithubIssueLabel' --json`. |
| Run the issues tests | Preflight organization, then `"$TPT_BIN" run --repo-root "$PROVIDER_ROOT" --mode organization --group issues --json`. |
| Retry what failed | `"$TPT_BIN" retry --repo-root "$PROVIDER_ROOT" --failed --json`. The mode comes from the persisted execution plan. |
| Resume the last run | `"$TPT_BIN" resume --repo-root "$PROVIDER_ROOT" --json`. The mode comes from the persisted execution plan. |
| What groups exist? | `"$TPT_BIN" groups --repo-root "$PROVIDER_ROOT" --json`. |
| Check my prereqs | `"$TPT_BIN" preflight --repo-root "$PROVIDER_ROOT" --mode organization --json`. Replace the mode if the user names one. |
| Run anonymous smoke | `"$TPT_BIN" preflight --repo-root "$PROVIDER_ROOT" --mode anonymous --json`, then `"$TPT_BIN" run --repo-root "$PROVIDER_ROOT" --mode anonymous --run '^TestAccGithubIpRangesDataSource$' --json`. |
| List this run's attributed resources | Set `RUN_MODE` from persisted state, then run `"$TPT_BIN" orphans --repo-root "$PROVIDER_ROOT" --mode "$RUN_MODE" --run-delta --json`. |
| Clean up this run's attributed resources | Preview first. Only after explicit user intent, run `"$TPT_BIN" sweep --repo-root "$PROVIDER_ROOT" --mode "$RUN_MODE" --run-delta --confirm`. |
| Make a report | `"$TPT_BIN" report --repo-root "$PROVIDER_ROOT" --md terraform-provider-tester-report.md` or `"$TPT_BIN" report --repo-root "$PROVIDER_ROOT" --html terraform-provider-tester-report.html`. |

## Explicit env files

Only when the user names an env file, add the global flag before the subcommand:

```sh
"$TPT_BIN" --env-file /path/to/env preflight --repo-root "$PROVIDER_ROOT" --mode organization --json
```

## Optional interactive dashboard

Humans can request the dashboard for `run`, `retry`, or `resume` with `--tui`, or run the binary with no subcommand on a TTY. Keep the same resolved binary and provider root, for example `"$TPT_BIN" run --repo-root "$PROVIDER_ROOT" --mode organization --tui`. Automation should use `--json`. The precedence is `--json`, then `--no-tui`, then `--tui`, then `PULSAR_NO_TUI` or `PULSAR_FORCE_TTY`, then TTY auto-detect.
