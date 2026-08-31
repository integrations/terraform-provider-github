# Contributing to Terraform Provider Tester

Terraform Provider Tester is the acceptance-test harness for
[terraform-provider-github](https://github.com/integrations/terraform-provider-github).
It wraps the provider's `TestAcc*` suite so a run is legible: preflight checks,
per-group pass and fail, single-test retry, and resume. It does not change the
provider or its tests.

## Before you start

- Go: the version pinned in `go.mod` (currently 1.26).
- A local checkout of the provider if you want to run anything past `--help`.
  Terraform Provider Tester drives `go test ./github/...` inside a provider tree, so point it at
  one with `--repo-root`, or run Terraform Provider Tester from inside the provider checkout.

## Build and test

```
make build      # builds bin/terraform-provider-tester, version-stamped
make test       # unit tests, no GitHub calls
make vet        # go vet
make checkdocs  # docs and code drift check, no provider pipeline
```

Before opening a pull request, run the full local gate and make sure it is
green:

```
gofmt -l .          # prints nothing when formatting is clean
go vet ./...
go build ./...
go test ./... -count=1
```

CI runs the same checks plus a gitleaks secret scan on every pull request.

## Pull requests

- Keep changes focused and describe the user impact in the body.
- Record user-facing changes in `CHANGELOG.md` under `## [Unreleased]`.
- Add or update tests when you change behavior. Terraform Provider Tester is test-driven; most
  packages have unit coverage and golden files.
- Never commit real tokens, secrets, or PEM keys. The harness reads credentials
  from the environment and redacts them. Run `gitleaks detect --no-git` locally
  if you touch anything credential-shaped.

## AI-assisted work

Terraform Provider Tester was built with AI assistance. That is fine, but a human must understand
and test any change before it merges, and the provider's AI Use Policy applies
to anything ever proposed upstream. Mark AI-assisted pull requests as such.

## Docs

Docs live in `docs/` and are plain Markdown. The drift check (`make checkdocs`)
keeps them honest about real flags and environment variables. See
[`docs/contributing.md`](./docs/contributing.md) for how the docs are written
and validated.

## Maintainers

See [`MAINTAINERS.md`](./MAINTAINERS.md). Owning team: Enterprise Primitives.
Open an issue or ping a maintainer with questions.
