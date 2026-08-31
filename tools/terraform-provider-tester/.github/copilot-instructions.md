# Copilot instructions for Terraform Provider Tester

Terraform Provider Tester is a Go 1.26 CLI harness for running terraform-provider-github acceptance tests. It wraps `go test -json`, can create real GitHub resources in credentialed modes, and is AI-assisted work for human review under the provider AI Use Policy.

## Use the skill

When the user asks to run acceptance tests, e2e tests, organization tests, groups, preflight, retry, resume, or triage a failing TestAcc result, use the repository skill:

`.github/skills/terraform-provider-tester/SKILL.md`

Examples of matching asks:

- "Run the individual e2e test"
- "Run the organization e2e test"
- "Run just the issue label tests"
- "Retry what failed"
- "What groups exist?"
- "Check my prereqs"
- "Resume the last run"

## Setup

1. Start in the terraform-provider-github checkout and set `PROVIDER_ROOT` with `git rev-parse --show-toplevel`.
2. Validate that `$PROVIDER_ROOT/go.mod` and `$PROVIDER_ROOT/github` exist.
3. Resolve `TPT_BIN` with the skill bootstrap: normalize an existing `terraform-provider-tester` on `PATH` to an absolute physical path, or use the fixed `$HOME/.local/share/terraform-provider-tester` cache.
4. Before using or building the cache, validate that it is not `$PROVIDER_ROOT`, that it is the top-level `github/terraform-provider-tester` checkout, that `go.mod` declares the expected module, that `origin` identifies `github/terraform-provider-tester`, and that the Makefile build target exists. If the cache is missing, shallow clone `github/terraform-provider-tester` there with `gh repo clone`, validate it, then build it.
5. Never build in the provider checkout, never use `go install` for the private tester repo, never overwrite or build an invalid cache directory, and never mutate global Git credential config.
6. Default commands omit an explicit env file. The CLI auto-detects `PULSAR_ENV_FILE`, cwd `.pulsar.env`, user config, home `.pulsar.env`, then no file. Pass an explicit env file only when the user names one.
7. Prefer `--json` for `e2e`, `preflight`, `groups`, `run`, `retry`, `resume`, and `orphans`.
8. Never print secrets from the environment or `.pulsar.env`.

## Agent-safe defaults

- Always run JSON preflight before credentialed resource-creating runs.
- For `run` failures, retry once with `retry --failed`, then report remaining failures with failure log paths.
- Do not use `--allow-sensitive-logs` unless the user explicitly asks.
- Credentialed `e2e` captures a read-only baseline before tests and performs a cancellation-independent final check with a 30-second timeout. `resume` reuses the original baseline.
- Preview before suggesting cleanup. Set `RUN_MODE` from `orphan_delta.mode` or the persisted logical-run mode, then use the exact run-delta preview command below.
- When cleanup status is `unknown`, persisted `New` is unavailable. Rerun `resume` against the original baseline, or use full non-delta `orphans` for read-only inspection. Do not use `--run-delta` until final accounting is complete.
- Never execute cleanup without explicit user intent. Attributed cleanup requires both `--run-delta` and `--confirm`; guided workflows never invoke sweep.
- Triage defaults to dry-run issue filing. Creating issues requires both `--file-issues` and `--confirm-file-issues`; do not combine filing with `--known-issues-offline`. `known-issues list|sync|add` manages the known-issues cache.
- Respect the provider AI Use Policy. Do not open an upstream provider PR from this output without human validation.

## Common commands

```sh
"$TPT_BIN" e2e --repo-root "$PROVIDER_ROOT" --mode individual --json
"$TPT_BIN" e2e --repo-root "$PROVIDER_ROOT" --json
"$TPT_BIN" preflight --repo-root "$PROVIDER_ROOT" --mode organization --json
"$TPT_BIN" run --repo-root "$PROVIDER_ROOT" --mode organization --json
"$TPT_BIN" retry --repo-root "$PROVIDER_ROOT" --failed --json
"$TPT_BIN" resume --repo-root "$PROVIDER_ROOT" --json
"$TPT_BIN" groups --repo-root "$PROVIDER_ROOT" --json
"$TPT_BIN" triage --repo-root "$PROVIDER_ROOT" --json
"$TPT_BIN" known-issues list --known-issues .pulsar-known-issues.yaml --known-issues-offline --json
"$TPT_BIN" orphans --repo-root "$PROVIDER_ROOT" --mode "$RUN_MODE" --run-delta --json
"$TPT_BIN" sweep --repo-root "$PROVIDER_ROOT" --mode "$RUN_MODE" --run-delta --confirm
```

The orphan command is read-only. The sweep command is destructive and belongs only in an explicitly authorized cleanup step after the preview has been inspected.
