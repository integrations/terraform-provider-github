---
applyTo: "website/**"
---

# Website / Docs Review

These rules apply to documentation under `website/`. Combine with the
repo-wide checklist in `.github/copilot-instructions.md`.

## Keep Docs in Sync with Schema

- Any schema change in `github/` (new attribute, renamed attribute,
  changed `Required`/`Optional`/`Computed`/`Default`, new `ForceNew`,
  removed attribute) must have a matching docs update.
- Argument tables should list attributes with the same name, type, and
  required/optional status as the schema.
- Deprecated attributes must be clearly marked and include guidance on the
  replacement.

## Imports

- Resources that support `terraform import` must document the import ID
  format with at least one example command.

## Permissions and Scopes

- For any GitHub API call that requires a specific token scope or app
  permission, the docs should call this out so users can configure their
  authentication correctly.

## Examples in Docs

- Inline example HCL should reflect current argument names and be
  syntactically valid.
- Sensitive attributes should not appear with real-looking secrets, even as
  examples.

## Style and Links

- Internal links between docs pages should resolve.
- New resources/data sources need at least one usage example and a list of
  every supported argument and attribute.
