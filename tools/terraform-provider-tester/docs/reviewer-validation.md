# Reviewer validation packet

This packet gives a reviewer or Copilot a disposable, credential-free way to
test this pull request against the current `terraform-provider-github` main
branch. It runs one read-only anonymous acceptance test and validates the
tester's machine-readable output.

The test reads GitHub's public IP-ranges metadata. It does not create or delete
GitHub resources.

## Copilot prompt

Point Copilot at this file and use:

> Follow `docs/reviewer-validation.md` exactly. Use the repository's
> `.github/skills/terraform-provider-tester/SKILL.md`. Run the disposable
> anonymous validation, show the tester and provider commit SHAs, both command
> exit codes, and the final compact JSON summary from each command. Do not read
> credential files, print environment values, run a credentialed mode, sweep
> resources, or file issues. Do not report success unless every assertion in
> the script exits zero.

## Automated anonymous validation

Requirements:

- `bash`
- `git`
- SSH access to the private tester repository
- `go`
- `jq`

Run from any directory:

```bash
bash <<'VALIDATE'
set -euo pipefail

RUN_ROOT="$(mktemp -d)"
cleanup() {
  chmod -R u+w "$RUN_ROOT" 2>/dev/null || true
  rm -rf "$RUN_ROOT"
}
trap cleanup EXIT

PROVIDER_ROOT="$RUN_ROOT/provider"
TESTER_ROOT="$RUN_ROOT/tester"
EMPTY_ENV="$RUN_ROOT/empty.env"
PREFLIGHT_LOG="$RUN_ROOT/preflight.ndjson"
RUN_LOG="$RUN_ROOT/run.ndjson"
touch "$EMPTY_ENV"

git clone --depth 1 https://github.com/integrations/terraform-provider-github "$PROVIDER_ROOT"
git clone git@github.com:github/terraform-provider-tester.git "$TESTER_ROOT"
git -C "$TESTER_ROOT" fetch origin pull/18/head:review-pr-18
git -C "$TESTER_ROOT" checkout review-pr-18
make -C "$TESTER_ROOT" build

TPT_BIN="$TESTER_ROOT/bin/terraform-provider-tester"
test -x "$TPT_BIN"

echo "tester_head=$(git -C "$TESTER_ROOT" rev-parse HEAD)"
echo "provider_head=$(git -C "$PROVIDER_ROOT" rev-parse HEAD)"

run_without_credentials() {
  env \
    -u GITHUB_TOKEN \
    -u GITHUB_APP_ID \
    -u GITHUB_APP_INSTALLATION_ID \
    -u GITHUB_APP_PEM_FILE \
    -u GITHUB_OWNER \
    -u GITHUB_USERNAME \
    -u GH_TEST_EXTERNAL_USER1_TOKEN \
    PULSAR_ENV_FILE="$EMPTY_ENV" \
    GH_TEST_AUTH_MODE=anonymous \
    PULSAR_NO_TUI=1 \
    "$@"
}

set +e
run_without_credentials \
  "$TPT_BIN" preflight \
  --repo-root "$PROVIDER_ROOT" \
  --mode anonymous \
  --json | tee "$PREFLIGHT_LOG"
preflight_exit="${PIPESTATUS[0]}"
set -e

echo "preflight_exit=$preflight_exit"
test "$preflight_exit" -eq 0
jq -e -s '
  map(select(.type == "summary")) | last |
  .command == "preflight" and
  .mode == "anonymous" and
  .ok == true and
  .checks.fail == 0 and
  .exit_code == 0
' "$PREFLIGHT_LOG" >/dev/null

set +e
run_without_credentials \
  "$TPT_BIN" run \
  --repo-root "$PROVIDER_ROOT" \
  --mode anonymous \
  --run '^TestAccGithubIpRangesDataSource$' \
  --json | tee "$RUN_LOG"
run_exit="${PIPESTATUS[0]}"
set -e

echo "run_exit=$run_exit"
test "$run_exit" -eq 0
jq -e '
  select(
    .type == "test" and
    .name == "TestAccGithubIpRangesDataSource" and
    .status == "pass"
  )
' "$RUN_LOG" >/dev/null
jq -e -s '
  map(select(.type == "summary")) | last |
  .command == "run" and
  .mode == "anonymous" and
  .build_failed == false and
  .pre_run_failed == false and
  .totals.total >= 1 and
  .totals.pass >= 1 and
  .totals.fail == 0 and
  .totals.panic == 0 and
  .totals.timeout == 0 and
  .exit_code == 0
' "$RUN_LOG" >/dev/null

echo "preflight_summary:"
jq -c 'select(.type == "summary")' "$PREFLIGHT_LOG"
echo "run_summary:"
jq -c 'select(.type == "summary")' "$RUN_LOG"
echo "anonymous_validation=passed"
VALIDATE
```

## Truthful result contract

Report **passed** only when:

1. `preflight_exit=0`;
2. `run_exit=0`;
3. both `jq -e` summary assertions exit zero;
4. the selected test emits a top-level `pass` event; and
5. the script prints `anonymous_validation=passed`.

If any command fails, report its exit code and final emitted summary. Do not
convert a build failure, pre-run failure, missing summary, or assertion failure
into a passing result.

## Optional human console check

After automated validation passes, open the same pull-request build against the
disposable provider checkout:

```bash
PULSAR_ENV_FILE="$EMPTY_ENV" \
GH_TEST_AUTH_MODE=anonymous \
PULSAR_FORCE_TTY=1 \
"$TPT_BIN" run \
  --repo-root "$PROVIDER_ROOT" \
  --mode anonymous \
  --run '^TestAccGithubIpRangesDataSource$' \
  --tui
```

Verify the GH/TF Provider Acceptance header, Preflight, Groups, Run, and Triage
tabs. Press `?` for the key reference and `q` to quit.
