---
applyTo: "{templates,docs}/**"
---

# Docs and Templates Review

These rules apply to changes under `templates/` and `docs/`. Combine with the
repo-wide checklist in [`.github/copilot-instructions.md`](../copilot-instructions.md).

## How Doc Generation Works

This repo uses
[`hashicorp/terraform-plugin-docs`](https://github.com/hashicorp/terraform-plugin-docs)
(`tfplugindocs`) to generate the contents of `docs/**` from three
sources:

- **Templates** under `templates/`, organized as:
  - `templates/index.md.tmpl` - provider landing page
  - `templates/resources.md.tmpl` and `templates/data-sources.md.tmpl` -
    optional category landing pages
  - `templates/resources/<name>.md.tmpl` and
    `templates/data-sources/<name>.md.tmpl` - per-resource and
    per-data-source pages
- **Example HCL** under `examples/`, organized as:
  - `examples/example_*.tf` - root-level snippets used by the provider
    landing page (`templates/index.md.tmpl`). These intentionally
    declare `required_providers` and a `provider` block because the
    landing page is the first thing users see.
  - `examples/resources/<name>/*.tf` - per-resource snippets (no
    `required_providers` or `provider` blocks; those would re-render in
    every resource doc).
  - `examples/data-sources/<name>/*.tf` - per-data-source snippets
    (same: no `required_providers` or `provider` blocks).
  - Per-resource `import.sh` & `import-by-string-id.tf` files for the import section.
- **Schema `Description` fields** in `github/**/*.go`, which become the
  argument and attribute reference rows.

If a resource or data source does not need any narrative beyond what
the schema provides, it does not need a template file at all -
`tfplugindocs` falls back to the default template.

Provider docs under `docs/**` are **auto-generated**. Do not edit files
under `docs/**` directly. A CI workflow regenerates docs and will fail
if the checked-in `docs/**` differs from the generated output.

Manual documentation edits belong in one of three places:

- `templates/**` - Markdown templates that drive doc generation. This is
  where most narrative doc edits live.
- `examples/**` - example HCL referenced by the templates.
- Resource and data source `Description` fields in `github/**/*.go`.

## Local Doc Workflow

Contributors can regenerate and validate the docs locally with the
make targets defined in `GNUmakefile`:

- `make generatedocs` - regenerate `docs/**` from templates, examples,
  and schema.
- `make validatedocs` - run `tfplugindocs validate` on the generated
  docs.
- `make checkdocs` - regenerate and fail if the working tree differs
  from what was committed (this is what CI runs).
- `make fmtdocs` / `make lintdocs` - format and lint the rendered
  Markdown under `docs/`.

## Flag These as HIGH

- Manual edits to `docs/**`. The doc generation workflow will revert
  them on the next run and the PR will fail CI. Move the change to the
  appropriate source (`templates/`, `examples/`, or the resource
  `Description` field) instead.

## Keep Docs in Sync with Schema (Fix Forward)

Doc fixes for argument descriptions, defaults, required/optional
status, or any other rendered-from-schema content should be made at
the source: the resource's `Description` fields in `github/**/*.go`
and/or the templates under `templates/`. Do not paper over a wrong
description by hand-editing `docs/**`; the next `make generatedocs`
run will revert it.

- Any schema change in `github/` (new attribute, renamed attribute,
  changed `Required`/`Optional`/`Computed`/`Default`, new `ForceNew`,
  removed attribute) must have a matching update either in the
  resource's `Description` fields (preferred for argument descriptions)
  or in the corresponding template under `templates/` (for narrative
  prose, import docs, and examples).
- Deprecated attributes must be clearly marked and include guidance on
  the replacement.

## Imports

- For resources that support `terraform import`, check that the
  resource's template renders an import section. `tfplugindocs`
  generates this automatically from a sibling `examples/resources/<name>/import.sh`
  when one exists, so the usual fix is to add or update the `import.sh`
  example rather than to hand-write the section in the template.

## Permissions and Scopes

- For any GitHub API call that requires a specific token scope or app
  permission, the template should call this out so users can configure
  their authentication correctly. Prefer linking to the relevant
  [GitHub REST API reference](https://docs.github.com/en/rest) page
  for the underlying endpoint rather than re-documenting the exact
  scope and permission list inline; that way the docs do not go stale
  when GitHub changes the required scopes or fine-grained permissions.

## Examples in Templates

- Inline example HCL should reflect current argument names and be
  syntactically valid. `terraform fmt` is the formatting baseline -
  examples committed in this repo are expected to be `terraform fmt`
  clean.
- Sensitive attributes should not appear with real-looking secrets, even
  as examples.

## Style and Links

- Internal links between docs pages must be relative paths (no leading
  slash, no absolute URL). `tfplugindocs` renders templates into a flat
  `docs/resources/` and `docs/data-sources/` layout, so the conventions
  are:

  ```markdown
  <!-- from docs/resources/branch.md, link to the repository resource -->
  [github_repository](repository)

  <!-- from docs/resources/branch.md, link to the repository data source -->
  [github_repository (datasource)](../data-sources/repository)
  ```

- New resources/data sources need at least one usage example. The list
  of supported arguments and attributes is rendered from the schema, so
  the `Description` fields in `github/**/*.go` are the source of truth
  for those rows.
