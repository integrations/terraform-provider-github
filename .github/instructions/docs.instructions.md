---
applyTo: "templates/**"
---

# Docs and Templates Review

Provider docs under `docs/**` are **auto-generated**. Do not edit files
under `docs/**` directly. A CI workflow regenerates docs and will fail
if the checked-in `docs/**` differs from the generated output.

Manual documentation edits belong in one of three places:

- `templates/**` - Markdown templates that drive doc generation. This is
  where most narrative doc edits live.
- `examples/**` - example HCL referenced by the templates.
- Resource and data source `Description` fields in `github/**/*.go`.

These rules apply to changes under `templates/`. Combine with the
repo-wide checklist in `.github/copilot-instructions.md`.

## Flag These as HIGH

- Manual edits to `docs/**`. The doc generation workflow will revert
  them on the next run and the PR will fail CI. Move the change to the
  appropriate source (`templates/`, `examples/`, or the resource
  `Description` field) instead.

## Keep Docs in Sync with Schema

- Any schema change in `github/` (new attribute, renamed attribute,
  changed `Required`/`Optional`/`Computed`/`Default`, new `ForceNew`,
  removed attribute) must have a matching update either in the
  resource's `Description` fields (preferred for argument descriptions)
  or in the corresponding template under `templates/` (for narrative
  prose, import docs, and examples).
- Deprecated attributes must be clearly marked and include guidance on
  the replacement.

## Imports

- Resources that support `terraform import` must document the import ID
  format with at least one example command in the relevant template.

## Permissions and Scopes

- For any GitHub API call that requires a specific token scope or app
  permission, the template should call this out so users can configure
  their authentication correctly.

## Examples in Templates

- Inline example HCL should reflect current argument names and be
  syntactically valid.
- Sensitive attributes should not appear with real-looking secrets, even
  as examples.

## Style and Links

- Internal links between docs pages should resolve.
- New resources/data sources need at least one usage example. The list
  of supported arguments and attributes is rendered from the schema, so
  the `Description` fields in `github/**/*.go` are the source of truth
  for those rows.
