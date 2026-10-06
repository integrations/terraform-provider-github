---
applyTo: "github/**/*.go"
---

# Provider Source Review (Schema, State, API)

These rules apply to all provider Go source files under `github/`. Combine
with the repo-wide checklist in [`.github/copilot-instructions.md`](../copilot-instructions.md)
and the idiomatic-Go reference in
[`go.instructions.md`](go.instructions.md). Where this file disagrees
with `go.instructions.md`, this file wins for provider source under
`github/`.

The `applyTo` glob above also matches test files
(`github/**/*_test.go`). That overlap is intentional: tests do exercise
schema-shaped code paths and the rules here are still useful context.
On test files, however, the test-focused guidance in
[`tests.instructions.md`](tests.instructions.md) takes precedence, and
the schema, CRUD, and API-safety bullets in this file should be treated
as background - flag them only when the test really is touching that
code (for example, hand-rolled pagination in an acceptance test).

## Schema Changes Are User-Visible

Any change to `schema.Schema` (`Type`, `Optional`, `Required`, `Computed`,
`ForceNew`, `Default`, `Sensitive`, `Description`) is potentially breaking.
Flag all schema diffs and verify:

- Attribute removals or renames have a deprecation cycle or `moved`/state
  migration guidance.
- `Optional` → `Required` transitions are called out as breaking.
- New `ForceNew` flags on existing attributes are called out as breaking
  (forces resource replacement).
- New attributes are `Optional` with a `Default` where reasonable to avoid
  forcing existing users to update their configs.
- All attributes carry a `Description` string (used for docs generation).
- Attributes accepting bounded values declare `ValidateDiagFunc` so
  invalid input fails at plan time, not apply time. `ValidateFunc` is
  deprecated and not allowed in this repo - do not suggest it.
- Attributes holding tokens, secrets, or private keys are marked
  `Sensitive: true`.

## Repository as a Required Argument

When a resource accepts a repository name as a required argument, follow
the provider's rename-safe convention so users can rename a repository
without triggering a destroy/recreate cycle:

- Name the attribute `repository` (not `repo`, not `repository_name`).
- Do **not** mark `repository` as `ForceNew`, even when the underlying
  resource needs to be recreated on most changes. The rename handling
  below decides when replacement is actually required.
- Add a `Computed` attribute called `repository_id` that holds the
  GitHub repository's numeric ID.
- Set `CustomizeDiff: diffRepository` on the resource (or include it via
  `customdiff.All(...)` when multiple `CustomizeDiff` funcs are needed).
  This compares the stored `repository_id` against the current ID for
  the named repository and only forces replacement when the underlying
  repository actually changed, not when it was merely renamed.

Flag any new resource that takes a repository as required input but is
missing this pattern.

## State and Drift

- Read functions must populate every state attribute from API responses so
  `terraform import` works from the resource ID alone.
- Verify the read path does not produce values that differ from what create/
  update wrote (causes perpetual diffs).
- Watch list/set ordering: prefer `schema.TypeSet` or stable sort when the
  GitHub API does not return deterministic order.
- Empty vs. null handling must be intentional and consistent between create,
  read, and update.
- Diff suppression (`DiffSuppressFunc`) and normalization should be reviewed
  for correctness whenever schema is touched.

## CRUD Behavior

- All CRUD functions (`Create`, `Read`, `Update`, `Delete`, and the
  importer) must use their `*Context` variants
  (`CreateContext`/`ReadContext`/`UpdateContext`/`DeleteContext` and
  `StateContext` on importers) and return `diag.Diagnostics`. Flag any
  new resource that uses the deprecated non-context variants.
- Update paths must not accidentally force resource replacement or wipe
  optional fields that the user did not change.
- Nil-check API response fields before dereferencing.
- Classify errors:
  - In **Read**, a 404 from GitHub means "remove from state" - call
    `d.SetId("")` and return nil. Other errors should bubble up.
  - In **Delete**, a 404 from GitHub means the object is already gone -
    treat it as success and return nil, not as an error.
  - Other unexpected status codes should bubble up.
- Respect `context.Context` cancellation and any configured timeouts.

## Repository-Specific: No Read-After-Write

**Do not flag** create or update functions that return `nil` instead of
calling the resource's read function. This is intentional in this provider to
minimize API calls against GitHub rate limits and to avoid stale reads from
eventually-consistent endpoints. See
[#2892](https://github.com/integrations/terraform-provider-github/issues/2892).

## API Safety and Performance

- Flag new N+1 access patterns over GitHub APIs.
- Verify pagination is handled on any endpoint that returns a list.
  Prefer the iterator pattern that `google/go-github` exposes via its
  generated `*Iter` methods, which return `iter.Seq2[T, error]` and
  walk every page automatically:

  ```go
  for item, err := range client.SomeService.ListSomethingIter(ctx, owner, repo, opts) {
      if err != nil {
          return err
      }
      // use item
  }
  ```

  The older `ListOptions{} + resp.NextPage` loop is still acceptable
  in existing code, but flag new pagination code that hand-rolls the
  `NextPage` loop when a corresponding `*Iter` method exists on the
  client.
- Check for rate-limit-sensitive loops; consider caching or batching where
  appropriate.
- Sensitive values must never appear in log output, even at debug/trace
  level.

## Security

- Token, credential, and webhook secret handling must follow least
  privilege.
- New API calls should document the GitHub permission scope they require.
- Do not hardcode secrets anywhere in source.
