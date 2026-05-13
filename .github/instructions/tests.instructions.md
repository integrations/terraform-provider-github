---
applyTo: "github/**/*_test.go"
---

# Provider Test Review

These rules apply to test files under `github/`. Combine with the repo-wide
checklist in `.github/copilot-instructions.md`.

## Coverage Expectations

- When behavior in a resource or data source changes, there must be a
  matching test change. Flag PRs where production code changed but tests
  did not.
- Prefer acceptance tests (`TF_ACC=1`) for any API-facing change. Unit-only
  coverage is rarely sufficient for schema or CRUD changes.
- Tests should assert the specific bugfix or regression being targeted, not
  only the happy path.

## Test Structure

- Acceptance tests should exercise create → read → import → update → destroy
  where applicable. Import steps are particularly valuable because they
  validate that the read function can reconstruct state from the ID alone.
- Use realistic fixture data; avoid asserting on transient or
  environment-specific fields without normalization.
- Avoid hardcoded secrets or tokens in test files; use environment variables
  or test helpers.

## terraform-plugin-testing Conventions

Tests in this repo use
[`github.com/hashicorp/terraform-plugin-testing`](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-testing)
and follow its modern conventions. Hold new and changed tests to these
patterns:

- Prefer `ConfigStateChecks` (state-check functions from the
  `statecheck` package) over the older `Check` /
  `resource.ComposeTestCheckFunc` pattern. Flag new `TestStep`s that
  use `Check` with the legacy `resource.TestCheckResourceAttr*` helpers
  when an equivalent `statecheck` exists.
- For comparing values across steps (e.g. to assert a value did or did
  not change between Apply runs), use
  [`ValueComparers`](https://developer.hashicorp.com/terraform/plugin/testing/acceptance-tests/value-comparers)
  via `compare.ValuesSame` / `compare.ValuesDiffer` instead of the older
  pattern of stashing pointers to attribute values in custom structs and
  comparing them by hand.
- Existing tests using the legacy `Check` pattern do not need to be
  rewritten as part of an unrelated change. Only flag the legacy pattern
  in new tests, or in tests that are being substantially modified.

## When Reviewing Test Changes

- If a test was deleted or weakened, explain why in the report and flag as
  at least MEDIUM unless the corresponding production code was also removed.
- New skip conditions or `t.Skip` calls must include a clear justification.
- Tests that depend on specific organization/repo names should use the
  shared test helpers/config, not hardcoded values.
