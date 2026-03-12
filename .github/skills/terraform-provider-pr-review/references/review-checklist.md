# Terraform Provider Review Checklist

Use this checklist when reviewing PRs in `terraform-provider-github`.

## 1. Correctness and Behavior

- Verify CRUD/read logic correctly maps GitHub API responses to Terraform schema/state.
- Check for nil handling, default-value drift, and state flattening/expansion mismatches.
- Confirm update paths do not accidentally force replacement or wipe optional fields.
- Validate retry/backoff and error classification for API failures.

## 2. Schema and State Compatibility

- Any `schema.Schema` changes (`Type`, `Optional`, `Required`, `Computed`, `ForceNew`, `Default`) can change user behavior.
- Confirm imports still work and read functions produce stable state.
- Flag any behavior that may break existing state without migration notes/tests.
- Watch for attribute rename/removal without deprecation path.
- Adding an optional attribute with a default is backward-compatible; removing or renaming an attribute is a breaking change that needs a deprecation cycle.
- New or changed attributes should include `ValidateFunc`/`ValidateDiagFunc` to catch invalid values at plan time rather than at apply time.
- All attributes should have a `Description` string for documentation generation.
- When resources are renamed or restructured, check for `moved` block guidance or state migration documentation so existing users don't face resource destruction on upgrade.
- Mark attributes that hold secrets with `Sensitive: true` in the schema to prevent leaking values in plan output and state.

## 3. Terraform UX and Drift

- Ensure diff suppression, normalization, and API canonicalization avoid perpetual diffs.
- Check that empty vs null handling is intentional.
- Verify list/set ordering behavior and deterministic state output.

## 4. Testing Expectations

- For behavior changes, check matching tests under `github/*_test.go`.
- Prefer acceptance tests for API-facing changes (`TF_ACC=1` scenarios).
- Ensure tests assert the bugfix/regression target, not only happy path.
- Flag missing tests when logic changed but test coverage did not.

## 5. Docs and Examples

- If resource/data source behavior changed, review website docs updates under `website/`.
- If user workflow changed, review corresponding example updates under `examples/`.
- Confirm examples still reflect current schema and argument names.
- Example directories should follow standard module structure (`main.tf`, `variables.tf`, `outputs.tf`) with a `README.md` explaining purpose and usage.
- Example configurations should not embed `provider` blocks inside child modules; provider configuration belongs in root modules only.
- Variables and outputs in examples should include `description` and `type`.
- If a PR adds or changes a resource, verify there is at least one example showing typical usage.

## 6. Security and Permissions

- Verify sensitive values are not logged or accidentally exposed in state.
- Check token/credential handling and least-privilege assumptions.
- Ensure permission scope requirements are documented for new API calls.
- Confirm no secrets or credentials are hardcoded in example configurations.
- Verify that debug/trace logging does not print sensitive attribute values.
- Sensitive outputs should be marked `sensitive = true` so Terraform redacts them in CLI output.

## 7. Performance and API Safety

- Flag new N+1 patterns, excessive API calls, or missing pagination handling.
- Check for rate-limit-sensitive loops and absent caching where needed.
- Confirm context cancellation/timeouts are respected in long operations.

## 8. Versioning and Changelog

- Breaking changes (attribute removal/rename, forced replacement, new required fields) must be called out for a MAJOR version bump under semantic versioning.
- Backward-compatible additions (new optional attributes with defaults, new resources/data sources) correspond to MINOR version bumps.
- Bug fixes with no schema change correspond to PATCH version bumps.
- Verify the PR description or CHANGELOG includes a clear summary of what changed and the user impact.

## 9. Review Report Rules

- Report findings first, ordered by severity.
- Include precise file references like `github/resource_x.go:123`.
- Explain impact in Terraform terms: plan/apply behavior, drift, state compatibility.
- Mention residual risks if tests or docs are incomplete.
