---
name: terraform-provider-pr-review
description: "Review pull requests for the Terraform GitHub provider. Use when asked to review a PR, audit Terraform provider changes, check schema or state behavior, verify acceptance tests, or catch regressions in resources/data sources/docs/examples."
argument-hint: "PR number/URL or 'active/open PR' plus focus areas (schema, tests, docs, migration, security)"
user-invocable: true
---

# Terraform Provider PR Review

Use this skill to perform a high-signal code review for this repository (`integrations/terraform-provider-github`).

## Review Goals

- Find correctness bugs, regressions, and provider behavior changes.
- Validate schema/state compatibility for Terraform users.
- Check test coverage (unit and acceptance), docs, and examples.
- Identify risk around GitHub API usage, permissions, and error handling.

## Terraform Background for Reviewers

This section summarizes the Terraform module and provider concepts most relevant to PR review. Refer back here when evaluating schema, examples, or state changes.

### Key Concepts

- **Resources and Data Sources**: A *resource* (`resource` block) manages a lifecycle object (create, read, update, delete). A *data source* (`data` block) only reads existing objects. Both declare a *schema* of typed attributes.
- **Schema Attributes**: Each attribute has a `Type` (string, int, bool, list, set, map), and flags like `Required`, `Optional`, `Computed`, `ForceNew`, `Default`, `Sensitive`, and `Description`. Changing any flag alters user-visible behavior.
- **State**: Terraform stores the last-known attribute values of every managed resource in *state*. The *read* function must produce output consistent with state or Terraform will detect drift and propose changes.
- **Plan and Apply**: `terraform plan` computes a diff between desired config and current state. `terraform apply` executes that diff. Bugs in schema or read logic cause perpetual diffs, surprise replacements, or silent data loss.
- **Imports**: Users can adopt existing infrastructure with `terraform import`. The resource's read function must be able to populate full state from just the resource ID.

### Module Structure (for reviewing `examples/`)

- A module is a directory of `.tf` files. The recommended layout is `main.tf`, `variables.tf`, `outputs.tf`, and a `README.md`.
- Variables and outputs should always include `description` and `type`.
- Child modules must **not** contain `provider` blocks — provider configuration belongs exclusively in root modules.
- Each module should declare `required_providers` with a minimum version constraint (`>=`).

### Backward Compatibility Rules

- **Safe (minor/patch)**: Adding new optional attributes with defaults, adding new resources or data sources.
- **Breaking (major)**: Removing or renaming attributes, changing `Optional` to `Required`, changing `Type`, adding `ForceNew` to an existing attribute.
- When renaming resources or restructuring modules, Terraform's `moved` block lets users migrate state without destroying infrastructure. If a PR restructures resources, check that migration guidance is provided.
- Removing a `moved` block is itself a breaking change.

## Inputs

- PR identifier: URL, number, or `active/open PR`.
- Optional focus areas: `schema`, `state`, `tests`, `docs`, `examples`, `security`, `performance`.

## Procedure

1. Gather PR context.
2. Inspect changed files and classify them:
	- Provider/resource/data-source code under `github/`
	- Tests (`*_test.go`, especially acceptance tests)
	- Docs/site content under `website/`
	- Example configurations under `examples/`
3. For schema changes, verify backward compatibility using the rules in *Terraform Background*: flag attribute removal/renames, type changes, new `ForceNew`, or `Optional`→`Required` transitions. If a PR restructures resources or renames modules, consult [Refactoring Modules](./references/refactoring.mdx) for `moved` block requirements.
4. For example/configuration changes, confirm they follow [Standard Module Structure](./references/structure.mdx) (`main.tf`, `variables.tf`, `outputs.tf`, `README.md`), do not embed `provider` blocks in child modules (see [Providers Within Modules](./references/providers.mdx)), and include `description` on variables/outputs.
5. Review against the checklist in [Terraform Provider Review Checklist](./references/review-checklist.md).
6. Prioritize findings by severity and provide actionable fixes.
7. Report residual risk and testing gaps when uncertain.

## Output Format

Return findings first, ordered by severity.

1. `HIGH`/`MEDIUM`/`LOW` title - short impact statement
2. File reference: `path/to/file.go:line`
3. Why this is a problem (runtime behavior, Terraform UX, upgrade risk)
4. Suggested fix (concise)

Then include:

- `Open Questions / Assumptions`
- `Residual Risk`
- `Change Summary` (brief)

If no issues are found, explicitly state `No blocking findings found` and list remaining risk areas.

## Repository Notes

- Acceptance and manual validation are important in this provider. See contribution guidance in `CONTRIBUTING.md`.
- Prefer matching resource/data source tests when implementation behavior changes.
- When schema or state semantics change, treat as high-risk unless compatibility is clearly preserved.
- Breaking changes must follow semantic versioning: attribute removal/rename or new required fields warrant a major version bump.
- Create/update functions in this provider intentionally do **not** call the read function afterward. This reduces API call volume against GitHub's rate limits and avoids stale reads from eventually-consistent endpoints (see [#2892](https://github.com/integrations/terraform-provider-github/issues/2892)). Do not flag this as an issue.
- Example configurations under `examples/` should be self-contained, follow standard module structure, and not include `provider` blocks inside nested modules.
- Sensitive attributes (tokens, secrets, private keys) must be marked `Sensitive: true` in the schema and must not appear in log output.
