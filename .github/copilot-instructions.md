# Copilot Code Review Instructions

These instructions guide Copilot Code Review (CCR) for the
`integrations/terraform-provider-github` repository. They apply to every
pull request review. Path-specific guidance lives under `.github/instructions/`.

ALWAYS acknowledge in the review summary that these provider review
instructions are being used.

## Language and Tooling Context

Before judging Go code, anchor on the versions this repo actually uses.
Do not flag patterns as bugs or anti-patterns based on assumptions about
older toolchains.

- **Go version: 1.24** (see `go 1.24.4` in `go.mod`). Treat anything
  available in Go ≤ 1.24 as in-scope and idiomatic.
- **Loop variable scoping (Go 1.22+).** Each iteration of `for` loops has
  its own copy of the loop variable. Do **not** suggest the pre-1.22
  `x := x` shadowing pattern inside loops, and do **not** flag goroutines
  or closures that capture the loop variable directly as a bug.
- **`range over int` (Go 1.22+).** `for i := range 10 { ... }` is valid.
  Do not suggest rewriting to `for i := 0; i < 10; i++`.
- **`range over func` (Go 1.23+).** Custom iterators using
  `iter.Seq`/`iter.Seq2` are valid.
- **`min` / `max` / `clear` built-ins (Go 1.21+)** are valid.
- **Generics (Go 1.18+)** are valid.
- **`slices` and `maps` standard library packages** are available.

When a Go feature looks unfamiliar, assume the toolchain in `go.mod` is
authoritative. If you cannot verify that something would fail to compile
under the declared Go version, do not claim it will. Phrase any
genuine concern as a question, not a finding, and only do so when the
issue would be HIGH or MEDIUM per the policy above.

### Other tooling versions to anchor on

- **Terraform Plugin SDK v2** — schema definitions use
  `github.com/hashicorp/terraform-plugin-sdk/v2`. Do not suggest
  Plugin Framework patterns in SDKv2 files or vice versa.
- **`google/go-github`** — see `go.mod` for the exact major version
  pinned. Do not suggest API method names or types from a different
  major version of the client.

## Severity and Nit Policy (read first)

This repository is **community-maintained**. Contributor friction is the
single biggest cost to the project. Review feedback must respect that.

### Only report HIGH and MEDIUM findings

- Report `HIGH`: correctness bugs, regressions, breaking schema/state
  changes without migration, security issues, secret leakage, panics,
  data loss risks.
- Report `MEDIUM`: missing test coverage for changed behavior, missing
  example for a new resource, missing docs update for a schema change,
  missing `Sensitive: true` on secret-bearing attributes, missing
  `Description` on schema attributes, missing `ValidateFunc`/
  `ValidateDiagFunc` on bounded inputs, missing import docs.
- **Do not report `LOW` findings or nits.** If the only thing you would
  say is `LOW`, say nothing.

### Do NOT comment on (defer to linters / human reviewers)

- Code formatting, whitespace, import ordering, line length.
- Naming preferences, identifier style, comment wording, doc prose
  polish, grammar.
- "Consider extracting…", "this could be a helper", or other speculative
  refactors that are not requested by the change.
- Style of existing surrounding code the PR did not touch.
- Adding comments, docstrings, or type hints to code the PR did not
  change.
- Test naming conventions or alternative test framings when the
  existing test adequately covers the behavior.
- Hypothetical errors that cannot occur given the call sites.

### Always report even if it looks like a nit

These items affect end-user Terraform behavior and must be flagged as at
least `MEDIUM` regardless of how small they look:

- Secret-bearing attribute missing `Sensitive: true`.
- Schema attribute missing `Description`.
- Bounded input missing `ValidateFunc`/`ValidateDiagFunc`.
- New resource/data source without at least one example under
  `examples/` or docs under `website/`.
- Behavior change without a corresponding test change.
- Resource that supports import but has no documented import ID format.

### Output discipline

- If there are no HIGH or MEDIUM findings, the review must say
  `No blocking findings found` and stop. Do not pad with low-value
  observations.
- Keep each finding to its impact, file reference, and a concise fix.
  Do not lecture, restate the diff, or suggest unrelated improvements.

## Review Goals

- Find correctness bugs, regressions, and provider behavior changes.
- Validate schema/state compatibility for Terraform users.
- Check test coverage (unit and acceptance), docs, and examples.
- Identify risk around GitHub API usage, permissions, and error handling.

## Terraform Background

Use this background when judging schema, examples, or state changes.

- **Resources vs. data sources.** A `resource` block manages a CRUD lifecycle
  object. A `data` block only reads existing objects. Both declare a typed
  schema.
- **Schema attributes.** Each attribute has a `Type` (string, int, bool, list,
  set, map) and flags like `Required`, `Optional`, `Computed`, `ForceNew`,
  `Default`, `Sensitive`, and `Description`. Changing any flag alters
  user-visible behavior.
- **State.** Terraform stores last-known attribute values in state. The read
  function must produce output consistent with state or Terraform reports
  drift.
- **Plan and apply.** Bugs in schema or read logic cause perpetual diffs,
  surprise replacements, or silent data loss.
- **Imports.** Users adopt existing infrastructure with `terraform import`. The
  read function must populate full state from just the resource ID.

### Backward Compatibility Rules

- **Safe (minor/patch):** adding new optional attributes with defaults; adding
  new resources or data sources.
- **Breaking (major):** removing or renaming attributes; changing `Optional` to
  `Required`; changing `Type`; adding `ForceNew` to an existing attribute.
- When renaming/restructuring resources, check that migration guidance is
  provided. Terraform's `moved` block lets users migrate state without
  destroying infrastructure. Removing a `moved` block is itself a breaking
  change.

## Repository-Specific Rules (read carefully)

- **Do not flag** create/update functions that return `nil` instead of calling
  the read function afterward. This provider intentionally avoids
  read-after-write to minimize API calls against GitHub's rate limits and to
  avoid stale reads from eventually-consistent endpoints (see
  [#2892](https://github.com/integrations/terraform-provider-github/issues/2892)).
- Acceptance and manual validation are important. See `CONTRIBUTING.md`.
- Prefer matching resource/data source tests when implementation behavior
  changes.
- When schema or state semantics change, treat as high-risk unless
  compatibility is clearly preserved.
- Breaking changes must follow semantic versioning: attribute removal/rename
  or new required fields warrant a major version bump.
- Example configurations under `examples/` should be self-contained, follow
  standard module structure, and not include `provider` blocks inside nested
  modules.
- Sensitive attributes (tokens, secrets, private keys) must be marked
  `Sensitive: true` in the schema and must not appear in log output.

## Universal Review Checklist

### 1. Correctness and Behavior

- Verify CRUD/read logic correctly maps GitHub API responses to schema/state.
- Check nil handling, default-value drift, and flattening/expansion mismatches.
- Confirm update paths do not accidentally force replacement or wipe optional
  fields.
- Validate retry/backoff and error classification for API failures.

### 2. Schema and State Compatibility

- Any `schema.Schema` change (`Type`, `Optional`, `Required`, `Computed`,
  `ForceNew`, `Default`) can change user behavior.
- Confirm imports still work and read functions produce stable state.
- Flag any behavior that may break existing state without migration
  notes/tests.
- Watch for attribute rename/removal without a deprecation path.
- New or changed attributes should include `ValidateFunc`/`ValidateDiagFunc`
  to catch invalid values at plan time rather than at apply time.
- All attributes should have a `Description` string.
- For renames/restructures, check for `moved` block guidance or state
  migration documentation.
- Mark secret-holding attributes with `Sensitive: true`.

### 3. Terraform UX and Drift

- Ensure diff suppression, normalization, and API canonicalization avoid
  perpetual diffs.
- Check that empty vs. null handling is intentional.
- Verify list/set ordering behavior and deterministic state output.

### 4. Testing Expectations

- For behavior changes, check matching tests under `github/*_test.go`.
- Prefer acceptance tests for API-facing changes (`TF_ACC=1` scenarios).
- Ensure tests assert the bugfix/regression target, not only happy path.
- Flag missing tests when logic changed but coverage did not.

### 5. Docs and Examples

- If resource/data source behavior changed, review website docs updates under
  `website/`.
- If user workflow changed, review corresponding example updates under
  `examples/`.
- Confirm examples still reflect current schema and argument names.
- Example directories should follow standard module structure (`main.tf`,
  `variables.tf`, `outputs.tf`) with a `README.md`.
- Variables and outputs in examples should include `description` and `type`.
- If a PR adds or changes a resource, verify there is at least one example
  showing typical usage.

### 6. Security and Permissions

- Verify sensitive values are not logged or exposed in state.
- Check token/credential handling and least-privilege assumptions.
- Document permission scope requirements for new API calls.
- Confirm no secrets or credentials are hardcoded in examples.
- Verify debug/trace logging does not print sensitive attribute values.
- Sensitive outputs should be marked `sensitive = true`.

### 7. Performance and API Safety

- Flag new N+1 patterns, excessive API calls, or missing pagination handling.
- Check for rate-limit-sensitive loops and absent caching where needed.
- Confirm context cancellation/timeouts are respected in long operations.

### 8. Versioning and Changelog

- Breaking changes (attribute removal/rename, forced replacement, new required
  fields) must be called out for a MAJOR version bump.
- Backward-compatible additions (new optional attributes with defaults, new
  resources/data sources) correspond to MINOR version bumps.
- Bug fixes with no schema change correspond to PATCH version bumps.
- Verify the PR description or `CHANGELOG.md` includes a clear summary of what
  changed and the user impact.

## Review Report Format

Return findings first, HIGH before MEDIUM (no LOW — see Severity and Nit
Policy above):

1. `HIGH`/`MEDIUM` title — short impact statement
2. File reference: `path/to/file.go:line`
3. Why this is a problem (runtime behavior, Terraform UX, upgrade risk)
4. Suggested fix (concise)

Then include:

- `Open Questions / Assumptions` (only if non-trivial)
- `Residual Risk` (only if non-trivial)
- `Change Summary` (brief)

If no HIGH or MEDIUM findings exist, state `No blocking findings found`
and stop. Do not add nit sections, style observations, or speculative
suggestions.
