# Private validation runbook

Use this runbook to prove Terraform Provider Tester works against **real GitHub**, on infrastructure **you own**, with **your own GitHub authentication**, before you trust it more widely. The primary path uses `--json` so Copilot, CI, and humans can parse the same output.

This is the "does it actually work in production" pass. It is deliberately small and private: you validate the harness end to end without touching the shared upstream repository or a shared test org.

> [!IMPORTANT]
> This is an AI-assisted tool built in a fork. It is not an upstream contribution. Do not open an upstream pull request from this work, and do not run it against shared upstream infrastructure, without a human understanding and testing it first. Respect the provider's AI Use Policy.

## What you are actually testing

The provider's acceptance tests exercise the **provider code** against the live GitHub API. They create real resources named `tf-acc-test-*`, assert on them, then destroy them. So "validate a recent push" means: take a recent change to `terraform-provider-github`, check it out privately, and run the acceptance group for the resource that change touched. The harness wraps that run so you can see it per group and per test, retry only the failures, and clean up.

You provide three things:

- **A checkout** of the provider at the commit you want to validate, such as your fork branch or a fetched pull request.
- **GitHub auth** you control, such as a personal access token or `gh auth token`.
- **An owner** to create throwaway resources under, such as your personal account for `individual` mode or a private org for `organization` mode.

## Before you start

- Put `go` on `PATH`. The harness shells out to `go test`.
- Build the harness once with `make -C ~/.local/share/terraform-provider-tester build`, and put `$HOME/.local/share/terraform-provider-tester/bin` on your `PATH`. Run the commands below from your provider checkout, where terraform-provider-tester finds the repo root automatically.
- Never commit a token. Export it into your shell only. The harness reads secrets from the environment, redacts them from reports, and never writes them to `.pulsar-state.json`.
- Acceptance runs cost real API calls and create real resources. Keep the first runs small.

## Step 1 - Zero-credential confidence

Confirm the harness builds, discovers the suite, and runs against your checkout before you involve any credentials.

```shell
# from the provider repo root, on the branch you want to validate
terraform-provider-tester preflight --mode anonymous --json
terraform-provider-tester groups --json
terraform-provider-tester run --mode anonymous --run '^TestAccGithubIpRangesDataSource$' --json
```

The groups listing is just `go test -list`: no credentials, no resources. The anonymous run exercises a public read-only endpoint and creates nothing. If these pass, the harness is wired to your checkout.

## Step 2 - Authenticate with your GitHub account

Pick one auth method. A classic personal access token is the most predictable because preflight can read its scopes.

Quick path, reuse your `gh` login:

```shell
export GH_TEST_AUTH_MODE=individual
export GITHUB_OWNER="$(gh api user --jq .login)"
export GITHUB_USERNAME="$GITHUB_OWNER"
# Acceptance teardown deletes repos, which needs the delete_repo scope a default gh login omits.
gh auth refresh -s delete_repo
export GITHUB_TOKEN="$(gh auth token)"
```

Robust path, use a dedicated classic token with the `repo`, `delete_repo`, and `user` scopes that preflight expects for `individual` mode. Set `GITHUB_OWNER`, `GITHUB_USERNAME`, and `GITHUB_TOKEN` in your shell or in a private `.pulsar.env` file. Never paste the token into docs, issues, or chat.

> [!NOTE]
> `gh auth token` returns whatever scopes your `gh` session has, and a default `gh` login does not include `delete_repo`. Without it, acceptance tests create repositories they cannot delete, so teardown returns 403 and leaks `tf-acc-test-*` repos. Add it with `gh auth refresh -s delete_repo`. That command only works when `gh` stores the token itself; when the token comes from a `GH_TOKEN` environment variable, create a classic token with `repo`, `delete_repo`, and `user` and set `GITHUB_TOKEN` to it instead. If preflight reports scopes as unverified, you are on a fine-grained PAT. For GitHub App auth instead of a token, see [Prerequisites](./prerequisites.md).

Validate before running anything:

```shell
terraform-provider-tester preflight --mode individual --json
```

Preflight prints names and statuses only, never your token. In JSON mode, use `check` events for details and the final `summary` object for the result.

## Step 3 - Run a real test privately

Run one real `individual`-mode test against your account. This one creates a repository under your account, reads it through a data source, then destroys it.

```shell
terraform-provider-tester run --mode individual --run '^TestAccDataSourceGithubRepository$' --json
```

The exit code is non-zero if anything failed. The last-run state is saved to `.pulsar-state.json` at the repo root.

Then explore the features you came to validate:

```shell
terraform-provider-tester report
terraform-provider-tester retry --failed --json
```

## Optional interactive dashboard

For a human dashboard, run the binary with no arguments on a TTY, or pass `--tui` to `run`, `retry`, or `resume`. Automation should keep using `--json`.

## Step 4 - Validate a specific recent push

To validate an actual change someone made, fetch it into your private checkout and run the group for the resource it touched.

```shell
# fetch a pull request by number into a local branch (origin = your fork or a private mirror)
git fetch origin pull/<PR_NUMBER>/head:validate-pr-<PR_NUMBER>
git checkout validate-pr-<PR_NUMBER>

# find the group for the touched resource, then run just that group
terraform-provider-tester groups --json
terraform-provider-tester run --mode individual --group <resource-group> --json
```

Most resources need an organization. If the change touches one of those, run `organization` mode against a **private** test org instead, and see [Prerequisites](./prerequisites.md) for the org variables `GH_TEST_ORG_USER1`, `GH_TEST_ORG_REPOSITORY`, `GH_TEST_ORG_TEMPLATE_REPOSITORY`, and `GH_TEST_ORG_SECRET_NAME`. The template repository must exist in your org and be marked as a template.

You can also point the harness at a checkout in another directory without changing directory:

```shell
terraform-provider-tester run --mode individual --group repositories --repo-root /path/to/private/clone --json
```

## Step 5 - Validate interruption and explicit cleanup

Run this destructive-safety QA only in a disposable individual account. Do not use a production account or a shared organization. The goal is to prove forced interruption, original-baseline reuse, read-only preview, and confirmation-gated attributed cleanup.

### Force one interruption and resume

Start a guided individual run, wait until acceptance-test execution begins, then interrupt it once with `Ctrl+C`:

```shell
terraform-provider-tester e2e --mode individual --json
```

Only guided `e2e` installs the graceful signal handler. A graceful first interruption of guided `e2e` can persist partial results, the execution plan, and final orphan accounting before returning. It can perform the cancellation-independent final check with a 30-second timeout, emit `orphan_delta`, and exit non-zero; when the save succeeds, `resume` reuses the original baseline. SIGKILL, or interruption of `run`, `retry`, or `resume` outside that guided signal path, may leave no resumable plan. In that case, start a new run or e2e workflow rather than assuming the original baseline survived. Inspect `.pulsar-state.json` without printing any secret environment values. Confirm that it contains the execution plan, one captured baseline, partial results, and final accounting. Then resume without overriding the persisted mode:

```shell
terraform-provider-tester resume --json
terraform-provider-tester report --md terraform-provider-tester-interruption-report.md
```

Verify that resume retained the original baseline rather than capturing a second one. A resource created before the interruption must remain `New` relative to that original baseline.

### Preview before an attributed sweep

Read the final `orphan_delta` and its exact commands. Preview the attributed resources:

```shell
terraform-provider-tester orphans --mode individual --run-delta --json
```

The preview is read-only. Before any confirmed cleanup, inspect every resource and verify that every name starts with the hard-coded `tf-acc-test-` prefix. Confirm that no pre-existing resource appears as `New` and no new orphan is missing from the preview.

Only with explicit cleanup intent, and only in the disposable account, run:

```shell
terraform-provider-tester sweep --mode individual --run-delta --confirm
```

The persisted run-delta remains historical after cleanup; sweep does not clear it. Do not rerun `orphans --run-delta` to prove current emptiness. Verify the live account with a full read-only scan:

```shell
terraform-provider-tester orphans --mode individual --json
```

Confirm that this full scan is empty. Guided `e2e` and `resume` never execute cleanup. A plain `sweep --mode individual --confirm` remains broader: it selects all prefixed resources in the provider's target kinds, not only this run's delta.

## Step 6 - Capture what you learned

Export a shareable, redacted artifact:

```shell
terraform-provider-tester report --md terraform-provider-tester-private-validation.md
```

Then note, for the production-readiness conversation:

- Did preflight catch a misconfiguration before the run, or did the run die on a missing variable?
- Were per-group and per-test results legible? Did retry re-run only the failures?
- Did resume pick up where an interrupted run stopped and reuse the original orphan baseline?
- Did the preview identify only `New` resources, and did a full non-delta orphan scan confirm that the explicitly swept resources are gone while the persisted historical delta remains?
- Was rate-limit headroom enough for the groups you ran?

These are the signals that tell you whether the harness is ready to put in front of other engineers.

## See also

- [Quickstart for terraform-provider-tester](./quickstart.md) - the zero-config first run.
- [Prerequisites](./prerequisites.md) - per-mode variables, including the org setup.
- [Test mode reference](./test-modes.md) - what each mode covers.
- [terraform-provider-tester CLI reference](./cli-reference.md) - every subcommand, flag, and exit code.
