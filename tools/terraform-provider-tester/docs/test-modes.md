# Test mode reference

`terraform-provider-tester` selects a test mode with `--mode` or `GH_TEST_AUTH_MODE`. If neither is set, it uses `anonymous`. The mode determines which tests run and which environment variables are required. Preflight validates the selected mode before `go test`; the provider's `TestMain` exits immediately if a required variable for the mode is missing, which is why preflight runs first.

For more information, see [Environment variable reference](./environment-variables.md), [Prerequisites](./prerequisites.md), and [terraform-provider-tester CLI reference](./cli-reference.md).

## Modes at a glance

| Mode | What it covers | Auth needed |
| --- | --- | --- |
| `anonymous` | Public read-only tests | No auth |
| `individual` | Individual account tests | PAT or GitHub App |
| `organization` | Organization resource coverage | PAT or GitHub App |
| `team` | Team coverage on top of organization setup | PAT or GitHub App |
| `enterprise` | Enterprise-only resources on top of organization setup | PAT or GitHub App |

## Mode details

Non-anonymous modes need either `GITHUB_TOKEN` or the GitHub App trio: `GITHUB_APP_ID`, `GITHUB_APP_INSTALLATION_ID`, and `GITHUB_APP_PEM_FILE`. `GITHUB_APP_PEM_FILE` is the PEM contents, not a path. Use `\n` for newlines when exporting it through a shell or secret manager. Token and App authentication are mutually exclusive.

### `anonymous`

Public endpoints only.

| Environment variable type | Environment variables |
| --- | --- |
| Required | None |
| Optional | None |

### `individual`

Individual mode uses an owner, username, auth, and a personal-account template repository. The complete individual suite keeps the user SSH key, user GPG key, and user Codespaces tests enabled, so a classic PAT needs `repo`, `delete_repo`, `user`, `admin:public_key`, `admin:gpg_key`, and `codespace`.

| Environment variable type | Environment variables |
| --- | --- |
| Required | `GITHUB_OWNER`, `GITHUB_USERNAME`, `GH_TEST_ORG_TEMPLATE_REPOSITORY`, and either `GITHUB_TOKEN` or `GITHUB_APP_ID` + `GITHUB_APP_INSTALLATION_ID` + `GITHUB_APP_PEM_FILE` |
| Optional | `GH_TEST_USER_REPOSITORY` |

### `organization`

Organization mode uses organization ownership, organization users, repositories, and a secret name.

| Environment variable type | Environment variables |
| --- | --- |
| Required | `GITHUB_OWNER`, `GH_TEST_ORG_USER1`, `GH_TEST_ORG_REPOSITORY`, `GH_TEST_ORG_TEMPLATE_REPOSITORY`, `GH_TEST_ORG_SECRET_NAME`, and either `GITHUB_TOKEN` or `GITHUB_APP_ID` + `GITHUB_APP_INSTALLATION_ID` + `GITHUB_APP_PEM_FILE` |
| Optional | `GH_TEST_ORG_USER3`, `GH_TEST_ORG_APP_INSTALLATION_ID` |

`GITHUB_OWNER` must be the organization. `GH_TEST_ORG_TEMPLATE_REPOSITORY` must exist and be marked as a template.

### `team`

Team mode builds on organization mode with another organization user and external users.

| Environment variable type | Environment variables |
| --- | --- |
| Required | All required `organization` variables, plus `GH_TEST_ORG_USER2`, `GH_TEST_EXTERNAL_USER1`, `GH_TEST_EXTERNAL_USER1_TOKEN`, and `GH_TEST_EXTERNAL_USER2` |
| Optional | `GH_TEST_ORG_USER3`, `GH_TEST_ORG_APP_INSTALLATION_ID` |

### `enterprise`

Enterprise mode builds on organization mode with an enterprise slug.

| Environment variable type | Environment variables |
| --- | --- |
| Required | All required `organization` variables, plus `GITHUB_ENTERPRISE_SLUG` |
| Optional | `GH_TEST_ORG_USER3`, `GH_TEST_ORG_APP_INSTALLATION_ID`, `GH_TEST_ENTERPRISE_IS_EMU`, `GH_TEST_ENTERPRISE_EMU_GROUP_ID`, `GH_TEST_ADVANCED_SECURITY`, `GITHUB_BASE_URL` |

Set `GITHUB_BASE_URL` for GHES.

## Choosing a mode

Start with `anonymous` to smoke-test the harness with zero configuration. Use `individual` next to validate credentialed access before you need a dedicated organization: it is the lowest-setup credentialed mode because it needs no dedicated organization, though it still needs a personal-account template repository in `GH_TEST_ORG_TEMPLATE_REPOSITORY` and the `repo`, `delete_repo`, `user`, `admin:public_key`, `admin:gpg_key`, and `codespace` classic scopes. Use `organization` for most resource coverage. Use `enterprise` for enterprise-only resources.

For Copilot and CI, pass `--mode` explicitly and use `--json` so the result is parseable. If a human uses the optional dashboard, they can press `m` to open the mode picker and switch auth modes without restarting. After switching, the dashboard re-runs preflight for the new mode automatically. Press `v` to view and set the non-secret variables the new mode requires.
