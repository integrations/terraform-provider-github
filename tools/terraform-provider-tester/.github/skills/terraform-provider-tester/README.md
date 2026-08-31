# Terraform Provider Tester skill

## When to Use This Skill

Use this skill when a user asks Copilot to run terraform-provider-github acceptance tests, run an individual or organization e2e test, check prerequisites, list groups, retry failed tests, resume an interrupted run, or triage a failing TestAcc result.

## User-level install

Because this repository is private, use existing `gh` authentication instead of an unauthenticated raw URL:

```sh
gh repo clone github/terraform-provider-tester ~/.local/share/terraform-provider-tester -- --depth 1
copilot skill add ~/.local/share/terraform-provider-tester/.github/skills/terraform-provider-tester/SKILL.md
```

The skill bootstraps the CLI from `~/.local/share/terraform-provider-tester` when a fresh provider worktree does not have the binary on `PATH`.

For repository use, keep this directory at `.github/skills/terraform-provider-tester`.

## Safety summary

Terraform Provider Tester can run credentialed acceptance tests that create real GitHub resources. Prefer `--json`, preflight before credentialed runs, never print secrets or `.pulsar.env`, retry failures once, and run `sweep --confirm` only when cleanup is explicitly requested. Default commands omit an explicit env file so the CLI can auto-detect `PULSAR_ENV_FILE`, cwd `.pulsar.env`, user config, home `.pulsar.env`, or no file.
