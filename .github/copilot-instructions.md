# Project Guidelines

## What maintainers consistently push for

- Keep changes tightly scoped. Prefer the smallest viable fix over broad cleanup or multi-issue refactors.
- Match existing patterns in the surrounding file instead of introducing new abstractions unless there is a clear repeated need.
- Favor readability and explicit behavior over cleverness; reviewers routinely call out structure, naming, and unnecessary complexity.

## Provider implementation

- Follow the modern Terraform Plugin SDK patterns already used in this repo, including `ReadContext`, `schema.ImportStatePassthroughContext`, and `ValidateDiagFunc`.
- Avoid deprecated SDK APIs when an established replacement already exists in the codebase.
- Preserve existing behavior unless the change is intentional and covered by tests.
- Be careful with GitHub API updates: only send fields when they actually changed, avoid widening side effects, and handle expected `404`/not-found cases gracefully when the resource should disappear from state.
- Keep helpers local and focused. Prefer a small helper in the same file over a sweeping refactor.

## Tests

- For bug fixes and features, add or update tests. Acceptance tests are the default for resource and data source behavior.
- Reuse the acceptance-test harness in `github/acc_test.go`, especially helpers like `skipUnlessMode`, `skipUnlessEnterprise`, `skipUnlessHasOrgs`, and related mode gates.
- When a resource supports import, add or preserve import verification coverage.
- Keep test coverage close to the behavior change instead of relying on unrelated broad test rewrites.

## Docs

- Update docs whenever schema, import behavior, or user-visible behavior changes.
- Check the pull request checklist in `.github/pull_request_template.md` and keep docs in sync with code.
- Be concise but complete. Reviewers do flag wording, capitalization, headings, and missing examples in docs.

## Build and validation

- Use the repo commands from `GNUmakefile` when validating changes: `make lintcheck`, `make website-lint`, `make build`, and `make test`.
- Follow the contribution workflow in `CONTRIBUTING.md` for local setup and acceptance-test expectations.

## Pull request mindset

- Assume maintainers will prefer a narrowly targeted PR with tests and docs over a larger “while I’m here” rewrite.
- If a change touches a brittle or actively changing area, keep the implementation incremental and avoid mixing compatibility cleanups with behavior changes in the same patch.
