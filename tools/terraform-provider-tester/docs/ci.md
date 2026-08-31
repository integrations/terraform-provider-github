# CI integration

This page shows how to call terraform-provider-tester's guided workflow from a CI pipeline and parse its JSON output for a pass/fail decision. See [CLI reference](./cli-reference.md) for every subcommand and JSON event field, and [prerequisites](./prerequisites.md) for the current OS support matrix.

## Recommended command

Run individual mode first: it needs no dedicated organization, so it is the lowest-setup credentialed check for a new CI pipeline. It still needs `GITHUB_OWNER`, `GITHUB_USERNAME`, `GITHUB_TOKEN` (or the `GITHUB_APP_*` trio), and `GH_TEST_ORG_TEMPLATE_REPOSITORY` naming a template repository owned by `GITHUB_OWNER`. A classic PAT needs the `repo`, `delete_repo`, `user`, `admin:public_key`, `admin:gpg_key`, and `codespace` scopes; see [prerequisites](./prerequisites.md).

```sh
terraform-provider-tester e2e --mode individual --json
```

Add `--mode organization` (or omit `--mode`, since `organization` is `e2e`'s default) once individual mode is green and the pipeline has organization credentials configured for full resource coverage.

Credentialed `e2e --json` writes one JSON object per line (NDJSON): a `plan` event, zero or more `excluded` events, one `check` event per preflight check, one `test` event per completed test, a `triage` event classifying any remaining failures, an `orphan_delta` event with baseline/final/pre-existing/new counts and cleanup status, and a final `summary` event for the whole workflow. Exact cleanup commands appear only when cleanup status supports an attributed delta; status `unknown` leaves both command fields empty. In `baseline-only` JSON, zero `final` and `pre_existing` values are unmeasured placeholders, and `cleanup_status` is authoritative. Anonymous `e2e --json` omits `orphan_delta` and cleanup commands because no orphan scan applies.

A guided run also embeds the run step's summary, so a stream can contain more than one `summary` object: the run step emits `"command": "run"`, and the workflow's own `"command": "e2e"` summary is written last. For credentialed accounting, `orphan_delta` is the penultimate object; the final object remains the workflow summary. Likewise, `resume --json` may emit an inner run summary and a final resume summary after orphan finalization. For `resume`, the last `summary` object is authoritative. Only the terminal resume summary is workflow-final; an earlier summary describes the run phase and does not include the final accounting verdict.

## Parsing the result

CI only needs the final `summary` event, not every intermediate event. Capture the full NDJSON stream, then select the last object whose `type` is `summary`:

```sh
terraform-provider-tester e2e --mode individual --json | jq -s 'map(select(.type == "summary")) | last'
```

`jq -s` (slurp) reads every NDJSON line into one array; `map(select(.type == "summary"))` keeps only `summary` objects; `last` returns the workflow summary, which is always written last. To assert that explicitly rather than relying on ordering, match on the command instead:

```sh
terraform-provider-tester e2e --mode individual --json | jq -s 'map(select(.type == "summary" and .command == "e2e")) | last'
```

CI consumers must ignore unknown NDJSON event types. Treat any `type` that is not `summary` as an event to skip, not as an error: `plan`, `excluded`, `check`, `test`, `triage`, and `orphan_delta` all appear before the workflow summary, and a future harness release can add more. Apply the same rule to `summary` objects for commands you do not handle. Filtering on `.type == "summary"` (optionally with `.command == "e2e"`) and ignoring everything else keeps the parser forward-compatible.

## Deciding pass or fail

Prefer the harness's own process exit code as the authoritative pass/fail signal alongside the parsed summary: `0` is success, `1` is a valid command with a failed result (preflight failure, test failure, build failure, lock error, or similar runtime problem), and `2` is a usage error such as a bad flag. If the pipeline needs failure detail for a report or notification, read `totals`, `failures`, and `exit_code` off the parsed `summary` object. Read `totals` and `failures` from the `"command": "run"` summary, and the workflow's overall verdict from the `"command": "e2e"` summary's `exit_code`.

## No automatic cleanup

Guided runs never delete resources automatically, in CI or anywhere else. If a job leaves an attributed delta with a supported cleanup status, run `terraform-provider-tester orphans --repo-root "$PROVIDER_ROOT" --mode <mode> --run-delta --json` as a separate read-only preview. Run `terraform-provider-tester sweep --repo-root "$PROVIDER_ROOT" --mode <mode> --run-delta --confirm` only as a separate cleanup step with explicit operator intent. Status `unknown` rejects run-delta operations; use full non-delta `orphans` for read-only inspection instead.
