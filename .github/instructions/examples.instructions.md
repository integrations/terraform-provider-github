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

   - `examples/provider/**/*.tf` - snippets used by the provider
     landing page (`templates/index.md.tmpl`). These are full,
     standalone snippets and **do** declare a `terraform { required_providers }`
     block and a `provider "github"` block, because the landing page
     is the first thing users see. Files matching
     `examples/provider/provider*.tf` can be automated via the
     `.Examples` templating capabilities; other files under
     `examples/provider/**/` (such as `examples/provider/app_auth/main.tf`)
     are referenced explicitly from the template.
   - `examples/resources/<resource_name>/resource*.tf` - per-resource
     snippets (plus optional `import.sh` for the import section).
   - `examples/data-sources/<data_source_name>/data-source*.tf` -
     per-data-source snippets.

   The per-resource and per-data-source snippets are intentionally
   minimal HCL fragments meant to be rendered inline by
   `tfplugindocs`. By default each resource or data source should
   have a single example file (`resource*.tf` or `data-source*.tf`)
   with any additional context placed in a comment at the top of the
   file. They should **not** carry their own `variables.tf`,
   `outputs.tf`, or `README.md`, and they should not declare
   `required_providers` or a `provider` block (those would re-render
   in every resource and data-source doc page). The provider
   landing-page snippets under `examples/provider/**/*.tf` are the
   deliberate exception.

   The `example*.tf` pattern should only appear where the docs
   template explicitly calls for it; in that case the snippet should
   still be a single file with any additional context placed directly
   in the template.

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
- `tfplugindocs` per-resource and per-data-source snippets (single-file
  snippets under `examples/resources/<name>/` and
  `examples/data-sources/<name>/`) should **not** declare
  `required_providers` or a `provider` block - they are rendered into
  the docs as fragments. The provider landing-page snippets under
  `examples/provider/**/*.tf` are the deliberate exception and do
  include both.

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
