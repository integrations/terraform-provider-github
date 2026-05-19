---
applyTo: "examples/**"
---

# Example Configuration Review

These rules apply to Terraform configurations under `examples/`. Combine
with the repo-wide checklist in [`.github/copilot-instructions.md`](../copilot-instructions.md).

## Two Kinds of Examples

This repo has two distinct kinds of example under `examples/`. The
expectations differ - apply the right set:

1. **`tfplugindocs` examples** - the single-file snippets that get
   embedded into the auto-generated docs. These live under:

   - `examples/provider/provider.tf`
   - `examples/resources/<resource_name>/example_N.tf`
     (plus optional `import.sh` for the import section)
   - `examples/data-sources/<data_source_name>/example_N.tf`

   These are not standalone Terraform modules. They are intentionally
   minimal HCL fragments meant to be rendered inline by
   `tfplugindocs`. They should **not** carry their own
   `variables.tf`, `outputs.tf`, or `README.md`, and they should not
   declare `required_providers` or a `provider` block.

2. **Root-module examples** - standalone, runnable Terraform
   configurations that demonstrate a more complete workflow (for
   instance `examples/release/`, `examples/emu/`,
   `examples/repository_collaborator/`). These should follow the
   standard module structure described below.

When reviewing changes, first identify which kind of example is being
touched and apply the matching ruleset.

## Standard Module Structure (Root-Module Examples Only)

Each root-module example directory should follow the
[standard module structure](https://developer.hashicorp.com/terraform/language/modules/develop/structure):

- `main.tf` - primary resources/module calls
- `variables.tf` - input variable declarations
- `outputs.tf` - output declarations
- `README.md` - purpose and usage of the example

Empty stub files are acceptable when an example legitimately has no inputs
or outputs.

## Variables and Outputs (Root-Module Examples Only)

- Every `variable` and `output` block should include a `description` and
  `type`.
- Outputs that surface sensitive values must be marked `sensitive = true`.
- Variables that accept secrets should be marked `sensitive = true`.

## Provider Configuration (Root-Module Examples Only)

- **Do not** embed `provider` blocks inside nested or child modules.
  Provider configuration belongs in root modules only. A module with a
  `provider` block is not compatible with `count`, `for_each`, or
  `depends_on`.
- Each root-module example should declare `required_providers` with a
  minimum version constraint (`>=`) for the `integrations/github`
  provider.
- `tfplugindocs` examples (single-file snippets under
  `examples/resources/<name>/` etc.) should **not** declare
  `required_providers` or a `provider` block - they are rendered into
  the docs as fragments.

## Coverage

- If a PR adds or meaningfully changes a resource or data source, verify
  there is at least one example demonstrating typical usage.
- Examples must reflect the current schema: argument names, required vs.
  optional, default values.

## Security

- No hardcoded tokens, passwords, or other secrets. Reference variables or
  environment-sourced values instead.
- Example READMEs should call out any non-obvious permission requirements.

## Refactoring and `moved` Blocks

When an example demonstrates resource renames or restructures, prefer
`moved` blocks over destroy-and-recreate. Removing a previously published
`moved` block is itself a breaking change for downstream users.
