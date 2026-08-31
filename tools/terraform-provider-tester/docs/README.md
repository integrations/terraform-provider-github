# Terraform Provider Tester documentation (terraform-provider-tester)

**Terraform Provider Tester** is the acceptance-test harness for terraform-provider-github; you
invoke it with the **`terraform-provider-tester`** command. It wraps the provider's existing
`go test` acceptance flow so you can prepare, run, inspect, retry, and clean up
acceptance tests without rewriting any test.

Start with [About Terraform Provider Tester](./index.md) for an overview, or jump to a topic below.

## Get started

- [About Terraform Provider Tester](./index.md) - what the harness does and why it exists.
- [Quickstart for terraform-provider-tester](./quickstart.md) - run your first test in anonymous mode with no credentials.
- [Prerequisites](./prerequisites.md) - tools, credentials, and per-mode setup.
- [Private validation runbook](./private-validation.md) - validate a real push privately on your own account or org with your GitHub auth.

## Reference

- [terraform-provider-tester CLI reference](./cli-reference.md) - every subcommand, flag, default, and exit code.
- [Environment variable reference](./environment-variables.md) - every variable, with required and secret status.
- [Test mode reference](./test-modes.md) - what each mode covers and the variables it needs.

## Concepts and help

- [Architecture](./architecture.md) - layers, data flow, and state model.
- [Troubleshooting terraform-provider-tester](./troubleshooting.md) - fixes for common failures.

## Maintainers

- [Contributing to the Terraform Provider Tester docs](./contributing.md) - where the docs live and publish, the drift-check, and the maintainer handoff.

> This is an AI-assisted tool built in a fork. It is not an upstream
> contribution. Per the provider's AI Use Policy, do not open an upstream pull
> request from this without a human understanding and testing it first.
