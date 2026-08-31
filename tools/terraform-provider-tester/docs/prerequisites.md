# Prerequisites

## Always required

Before you run acceptance tests with `terraform-provider-tester`, make sure your environment has the required tools and variables. Preflight checks these before `go test` runs so you do not hit the provider's blind `os.Exit(1)` when configuration is missing.

| Requirement | Details |
| --- | --- |
| Go toolchain | Put `go` on `PATH`. The harness shells out to `go test`. |
| `TF_ACC=1` | Acceptance tests require `TF_ACC=1`. The harness sets it for runs. |
| `GITHUB_BASE_URL` | Optional. Set it when you run tests against GitHub Enterprise Server. |
| Operating system | macOS and Linux are supported. Automated CI currently validates Linux. Windows is unverified and not supported today; cross-builds fail because lock identity uses POSIX-only `syscall.Stat_t`. |

For a first run without credentials, see [Quickstart for terraform-provider-tester](./quickstart.md).

## Prerequisites by mode

Each non-anonymous mode needs either a PAT with `GITHUB_TOKEN` or the GitHub App trio `GITHUB_APP_ID`, `GITHUB_APP_INSTALLATION_ID`, and `GITHUB_APP_PEM_FILE`. These auth methods are mutually exclusive. `GITHUB_APP_PEM_FILE` is the PEM contents, not a path. Use `\n` for newlines when exporting it through a shell or secret manager.

> [!IMPORTANT]
> Every mode except anonymous needs `GH_TEST_ORG_TEMPLATE_REPOSITORY` to name a repository that exists under `GITHUB_OWNER` and is marked as a template repository, or tests fail. In individual mode that repository lives under the personal account, not an organization.

### Anonymous

Anonymous mode uses public endpoints only.

| Type | Environment variables |
| --- | --- |
| Required | None |
| Optional | None |

### Individual

Individual mode runs tests for a user account. It is the lowest-setup credentialed mode and needs no dedicated organization, but it is not satisfied by a bare personal account and a generic PAT: the complete individual suite keeps the user SSH key, user GPG key, and user Codespaces tests enabled, and the repository resource and data source exercise a configured template repository.

| Type | Environment variables |
| --- | --- |
| Required | `GITHUB_OWNER`, `GITHUB_USERNAME`, one auth method, and `GH_TEST_ORG_TEMPLATE_REPOSITORY` naming a template repository owned by `GITHUB_OWNER` |
| Optional | `GH_TEST_USER_REPOSITORY` |

A classic PAT for the complete individual suite needs `repo`, `delete_repo`, `user`, `admin:public_key`, `admin:gpg_key`, and `codespace`. Preflight requests exactly the scopes and capabilities the plan aggregates, so a narrower selection (`--group` or `--run`) requires less. With a fine-grained PAT or a GitHub App, grant the equivalent permissions: repository administration and contents, user email/keys and GPG keys, and Codespaces user secrets. Preflight cannot read scopes for those credential types and reports them as unverified, so confirm them yourself.

### Organization

Organization mode runs tests for an organization.

| Type | Environment variables |
| --- | --- |
| Required | `GITHUB_OWNER`, one auth method, `GH_TEST_ORG_USER1`, `GH_TEST_ORG_REPOSITORY`, `GH_TEST_ORG_TEMPLATE_REPOSITORY`, `GH_TEST_ORG_SECRET_NAME` |
| Optional | `GH_TEST_ORG_USER3`, `GH_TEST_ORG_APP_INSTALLATION_ID` |

### Team

Team mode includes the organization mode requirements.

| Type | Environment variables |
| --- | --- |
| Required | `GITHUB_OWNER`, one auth method, `GH_TEST_ORG_USER1`, `GH_TEST_ORG_REPOSITORY`, `GH_TEST_ORG_TEMPLATE_REPOSITORY`, `GH_TEST_ORG_SECRET_NAME`, `GH_TEST_ORG_USER2`, `GH_TEST_EXTERNAL_USER1`, `GH_TEST_EXTERNAL_USER1_TOKEN`, `GH_TEST_EXTERNAL_USER2` |
| Optional | `GH_TEST_ORG_USER3`, `GH_TEST_ORG_APP_INSTALLATION_ID` |

### Enterprise

Enterprise mode includes the organization mode requirements.

| Type | Environment variables |
| --- | --- |
| Required | `GITHUB_OWNER`, one auth method, `GH_TEST_ORG_USER1`, `GH_TEST_ORG_REPOSITORY`, `GH_TEST_ORG_TEMPLATE_REPOSITORY`, `GH_TEST_ORG_SECRET_NAME`, `GITHUB_ENTERPRISE_SLUG` |
| Optional | `GH_TEST_ORG_USER3`, `GH_TEST_ORG_APP_INSTALLATION_ID`, `GH_TEST_ENTERPRISE_IS_EMU`, `GH_TEST_ENTERPRISE_EMU_GROUP_ID`, `GH_TEST_ADVANCED_SECURITY`, `GITHUB_BASE_URL` |

## Token requirements

Preflight checks token access before a long run. It calls `GET /user` to confirm the token works and capture the login, reads `X-OAuth-Scopes` for classic PATs, and checks `GET /rate_limit` for headroom. Fine-grained PATs and GitHub Apps do not expose scopes, so preflight reports their capabilities as unverified.

Classic PATs need `repo` and `delete_repo` in every mode, because tests create and then delete `tf-acc-test-*` repositories. Individual mode also needs `user`, `admin:public_key`, `admin:gpg_key`, and `codespace` for the complete suite; organization, team, and enterprise modes also need `read:org` and `admin:org` (enterprise adds `admin:enterprise`). A default `gh auth login` token omits `delete_repo`; add it with `gh auth refresh -s delete_repo`.

## Next steps

- For the full environment variable table, see [Environment variable reference](./environment-variables.md).
- For mode-specific behavior, see [Test mode reference](./test-modes.md).
- For a zero-configuration first run, see [Quickstart for terraform-provider-tester](./quickstart.md).
- To discover which orgs and enterprises your token can reach, run `terraform-provider-tester discover`.
