---
applyTo: "examples/**"
---

# Example Configuration Review

These rules apply to Terraform configurations under `examples/`. Combine
with the repo-wide checklist in `.github/copilot-instructions.md`.

## Standard Module Structure

Each example directory should follow the
[standard module structure](https://developer.hashicorp.com/terraform/language/modules/develop/structure):

- `main.tf` — primary resources/module calls
- `variables.tf` — input variable declarations
- `outputs.tf` — output declarations
- `README.md` — purpose and usage of the example

Empty stub files are acceptable when an example legitimately has no inputs
or outputs.

## Variables and Outputs

- Every `variable` and `output` block should include a `description` and
  `type`.
- Outputs that surface sensitive values must be marked `sensitive = true`.
- Variables that accept secrets should be marked `sensitive = true`.

## Provider Configuration

- **Do not** embed `provider` blocks inside nested or child modules.
  Provider configuration belongs in root modules only. A module with a
  `provider` block is not compatible with `count`, `for_each`, or
  `depends_on`.
- Each example module should declare `required_providers` with a minimum
  version constraint (`>=`) for the `integrations/github` provider.

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
