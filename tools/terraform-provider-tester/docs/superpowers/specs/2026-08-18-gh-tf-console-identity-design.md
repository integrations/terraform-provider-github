# GH/TF Console Identity Design

**Date:** 2026-08-18
**Status:** Approved for implementation
**Scope:** Optional interactive dashboard only

## Decision

The dashboard identity becomes **GH/TF Provider Acceptance**.

- `GH` names the GitHub-owned provider.
- `TF` names Terraform and carries the Terraform-purple accent.
- `Provider Acceptance` states the console's job without creating another
  standalone product brand.

The repository, Go module, binary, command, and automation-facing product name
remain `terraform-provider-tester` / Terraform Provider Tester. Existing
`.pulsar*`, `PULSAR_*`, fingerprint, state, and confirmation compatibility
identifiers remain unchanged.

## Visual System

The console stays industrial and data-dense. It should look like an operator
surface for `integrations/terraform-provider-github`, not a generic testing app.

### Identity

- Splash wordmark: `GH/TF` in the existing deterministic 5x3 dot-matrix motion.
- Full header: `◆ GH/TF Provider Acceptance <version>`.
- Compact header: `◆ GH/TF <version>`.
- Welcome card: stacked `GH` / `TF` mark and the full console title.
- Supporting copy continues to identify
  `integrations/terraform-provider-github` explicitly.

### Color Roles

- Terraform purple: the `TF` mark, active-tab bottom rule, and selection cursor.
- GitHub Primer blue: actions, links, commands, and identifiers.
- Green/red/yellow: pass/fail/running semantics only.
- Neutral gray: borders, inactive tabs, metadata, and selected-row surface.

Terraform purple must not fill an entire active tab or selected table row.
Color remains meaningful rather than decorative.

### Typography and Motion

- Preserve terminal-native monospace typography.
- Preserve the deterministic reveal/hold/dissolve intro sequence.
- Shorten the wordmark from six glyphs to five so it remains legible in compact
  terminals.
- Preserve the strict ASCII/NO_COLOR fallback.

## Interaction

No command, key binding, data flow, acceptance-test behavior, or destructive
confirmation changes. The four-tab information architecture remains:

1. Preflight
2. Groups
3. Run
4. Triage

## Documentation

Regenerate the deterministic intro, preflight, groups, run, and triage
screenshots after implementation. Update copy that describes the visible
`TESTER` wordmark, but do not rename the repository, binary, command examples,
state files, environment variables, or public compatibility identifiers.

## Validation

- Add failing tests for the `GH/TF` wordmark and console title before changing
  production code.
- Pin the active-tab bottom rule and absence of a selected-row purple
  background.
- Update ASCII and TTY wordmark goldens.
- Run all TUI tests, full repository tests, race tests, vet, docs checks, build,
  and the isolated pseudo-TTY smoke.
- Present the rebuilt dashboard for human visual review before merging or
  releasing.
