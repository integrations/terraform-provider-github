# Contributing to the Terraform Provider Tester docs

This page explains how the Terraform Provider Tester (`terraform-provider-tester`) documentation is written, where it
is published, how it is kept honest, and what a maintainer needs before this
tool is adopted upstream. It is for contributors and maintainers, not end users.

## Where the docs live and publish

All Terraform Provider Tester documentation lives in this repository, next to the harness it
documents:

- `README.md` - the landing page and install summary.
- `docs/` - the full doc set (this directory).

There is no separate publishing step. The docs are plain Markdown, read on
GitHub in this repository. The CLI, skill, and optional dashboard point readers
here. This keeps the docs versioned with the code that they describe and adds no
hosting surface to maintain.

If the tool is ever folded into the provider, the right surface is a short
pointer from the provider's `CONTRIBUTING.md` near its manual-testing section,
with the detailed docs staying in this repository's `docs/`.

## How this differs from the provider's own docs

Do not confuse these docs with the provider's user documentation:

- The **provider** generates its resource and data-source docs with
  `tfplugindocs` from `templates/`, `examples/`, and schema descriptions. Those
  are built and validated by `make generatedocs`, `make validatedocs`, and
  `make checkdocs` at the repository root, and published to the
  [Terraform Registry](https://registry.terraform.io/providers/integrations/github/).
- **Terraform Provider Tester** is a maintainer tool, not a Terraform resource, so it has no place
  in the Registry. Its docs are hand-written Markdown under `docs/`
  and never run through `tfplugindocs`.

The two pipelines are independent. Editing Terraform Provider Tester docs never touches the
provider's generated docs, and the provider's docs jobs never read this tree.

## Editing the docs

1. Edit the relevant Markdown file under `docs/` (or the README).
2. Run `make checkdocs`.
3. Fix anything it reports, then commit.

Keep the writing task-oriented and factual. State only behavior that the harness
actually has: real subcommands, real flags, real environment variables. When in
doubt, leave it out.

## The drift-check

`make checkdocs` runs a fork-safe consistency check (in `cli/`) that fails if the
docs and the code disagree. It verifies that every Markdown file in this
repository:

- names only environment variables the harness reads, derived from the
  provider's `EnvFor` plus a small, cited set of behavior and provider-level
  variables;
- shows only real `terraform-provider-tester` subcommands in inline code; and
- documents every real subcommand in the CLI reference.

The check imports only harness packages and is independent of the provider's
`tfplugindocs` pipeline, so it is safe to run in fork CI. Run it before opening
any docs change.

## Maintainer handoff and upstreaming

Terraform Provider Tester is an AI-assisted tool built in a fork. Before any of it is proposed
upstream:

- A human must understand and test the change, and the pull request must be
  marked AI-assisted, per the provider's AI Use Policy in `CONTRIBUTING.md`.
- Record user-facing changes in `CHANGELOG.md`.
- Decide ownership: `CODEOWNERS` lists the current owners; confirm them before
  wider adoption.

These are maintainer decisions and are intentionally left open here. See the
project's open questions before upstreaming.

## Reading order

New readers should start at [About Terraform Provider Tester](./index.md), then
[Quickstart for terraform-provider-tester](./quickstart.md) and [Prerequisites](./prerequisites.md).
Reference material lives in the
[CLI reference](./cli-reference.md),
[environment variable reference](./environment-variables.md), and
[test mode reference](./test-modes.md).
