# Quickstart for terraform-provider-tester

In this quickstart, you run your first acceptance test in a few minutes with no credentials. The examples use `--json` because it is the safest path for Copilot, CI, and repeatable local checks.

## Prerequisites

Before you start, make sure you have:

- Go on `PATH`, because the harness shells out to `go test`.
- A clone of the provider repository.

You do not need a PAT or org for anonymous mode.

> [!NOTE]
> Anonymous mode needs no credentials, only reads the public `api.github.com/meta` endpoint, and creates no GitHub resources.

## Step 1 - Install terraform-provider-tester

1. Clone and build the private tester repo in the shared user location.

   ```shell
   gh repo clone github/terraform-provider-tester ~/.local/share/terraform-provider-tester -- --depth 1
   make -C ~/.local/share/terraform-provider-tester build
   export PATH="$HOME/.local/share/terraform-provider-tester/bin:$PATH"
   ```

2. Clone the provider. Terraform Provider Tester runs its `go test` acceptance flow against this checkout.

   ```shell
   git clone https://github.com/integrations/terraform-provider-github
   cd terraform-provider-github
   ```

## Step 2 - Run preflight as NDJSON

Run preflight in anonymous mode.

```shell
terraform-provider-tester preflight --mode anonymous --json
```

Preflight is read-only and validates config before any test runs. In JSON mode, stdout is one JSON object per line. Use the final `summary` object for the result.

## Step 3 - Run a smoke test as NDJSON

Run the anonymous smoke test.

```shell
terraform-provider-tester run --mode anonymous --run '^TestAccGithubIpRangesDataSource$' --json
```

A cold build can take around 30-60 seconds. In JSON mode, the exact `go_test_command` appears in the final `summary` object.

## Step 4 - See the results

The final `summary` object includes totals, the exit code, failure log paths, and any build or pre-run failure fields. You can export a report with:

```shell
terraform-provider-tester report --md PATH
```

For more information, see [terraform-provider-tester CLI reference](./cli-reference.md).

## Step 5 - Run a guided credentialed flow

When you are ready to exercise real GitHub resources, run the guided end-to-end workflow. Start with individual mode - the lowest-setup credentialed mode, because it needs no dedicated organization:

```shell
terraform-provider-tester e2e --mode individual --json
```

Individual mode is still not zero-setup: it needs `GITHUB_OWNER`, `GITHUB_USERNAME`, one auth method, and `GH_TEST_ORG_TEMPLATE_REPOSITORY` naming a template repository owned by `GITHUB_OWNER`. A classic PAT needs `repo`, `delete_repo`, `user`, `admin:public_key`, `admin:gpg_key`, and `codespace`.

`e2e` runs preflight, that mode's acceptance tests, one retry of failures, triage, and, in the credentialed modes, a read-only orphan check. `--mode` defaults to `organization` when omitted. See [Prerequisites](./prerequisites.md) for the credentials each mode needs, and [Test mode reference](./test-modes.md) for what individual mode covers relative to organization mode.

After a credentialed run, `e2e` prints exact commands for the persisted logical-run mode and repository root only when cleanup status supports an attributed delta. The attributed command shapes are `terraform-provider-tester orphans --mode <mode> --run-delta` and `terraform-provider-tester sweep --mode <mode> --run-delta --confirm`; status `unknown` suppresses both. In `baseline-only` JSON, zero `final` and `pre_existing` values are unmeasured placeholders, and `cleanup_status` is authoritative. The preview is read-only. Inspect it first, and run the cleanup command only when you intend to delete this run's `New` resources; guided execution never runs cleanup for you.

## Optional interactive dashboard

For the **GH/TF Provider Acceptance** dashboard, run `terraform-provider-tester` with no subcommand on a TTY, or pass `--tui` to `run`, `retry`, or `resume`. Automation should keep using `--json`. Inside the dashboard, ordinary `t` triage refreshes and post-run triage updates stay cache-only. Press uppercase `K` when you want an explicit live known-issue sync. The TUI remains optional.

### Explore the dashboard

The dashboard has four tabs: Preflight, Groups, Run, and Triage. Press Left/Right or Tab/Shift+Tab to switch tabs. Press `1` through `4` to jump directly. Press `?` for help. Press Escape to close expanded help or an active overlay.

## Next steps

To run credentialed modes, see [Prerequisites](./prerequisites.md) and [Test mode reference](./test-modes.md). To load credentials from a file instead of exporting each variable manually, see [Environment variable reference](./environment-variables.md#single-env-file). For full command detail, see [terraform-provider-tester CLI reference](./cli-reference.md).
