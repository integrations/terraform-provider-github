# Environment variable reference

Use this reference to configure `terraform-provider-tester` and the provider acceptance tests that it runs. For mode requirements, see [Test mode reference](./test-modes.md). For setup requirements, see [Prerequisites](./prerequisites.md).

## Single env file

You can store all variables for every mode in one file and load it with the global `--env-file` flag.

```shell
terraform-provider-tester --env-file .pulsar.env run --mode organization --json
```

A `.pulsar.env` in the working directory can be loaded automatically without the flag. The file name is retained legacy naming.

Env-file precedence uses the first selected source in this order:

1. Explicit `--env-file`.
2. `PULSAR_ENV_FILE`, when set to a non-empty path.
3. `.pulsar.env` in the working directory.
4. `<os.UserConfigDir()>/terraform-provider-tester/.pulsar.env`.
5. `<os.UserHomeDir()>/.pulsar.env`.
6. None. No env file is loaded, and this is not an error.

The user config directory comes from the operating system. On Linux this is usually `~/.config`, so the persistent file is usually `~/.config/terraform-provider-tester/.pulsar.env`. On macOS it is usually `~/Library/Application Support`, so the persistent file is usually `~/Library/Application Support/terraform-provider-tester/.pulsar.env`.

If `PULSAR_ENV_FILE` is set but the file cannot be opened, the command exits with an `env-file:` error.

**No-override rule.** Variables already set to a non-empty value in the process environment are never overridden; the file only fills in variables that are absent. A variable exported as an empty string counts as absent, so any non-empty shell value always wins over the file.

**Mode-scoped keys.** Because `GITHUB_OWNER` means a personal account in `individual` mode but an organization in `organization` mode, one file can carry both values using a `PULSAR_<MODE>_` prefix. The prefix for the active mode is stripped and the bare key is applied, overriding any plain value in the same file. The mode in the prefix is matched case-insensitively, so writing the mode in lowercase works the same as the uppercase `PULSAR_<MODE>_` form.

```shell
# shared across all modes
GITHUB_TOKEN=<token>

# GITHUB_OWNER per mode, replace <MODE> with INDIVIDUAL or ORGANIZATION
PULSAR_<MODE>_GITHUB_OWNER=<owner>
```

Switch modes by changing only `--mode`:

```shell
terraform-provider-tester --env-file .pulsar.env run --mode individual --json
terraform-provider-tester --env-file .pulsar.env run --mode organization --json
```

**Parser rules.** Each line is `KEY=VALUE`. Blank lines and `#` comments are ignored. An optional `export ` prefix is stripped. Values may be unquoted, single-quoted, or double-quoted. No shell or `$VAR` expansion is performed. Only real mode names (`anonymous`, `individual`, `organization`, `team`, `enterprise`) are reserved after `PULSAR_` for mode-scoped keys (`PULSAR_<MODE>_<KEY>`); any other `PULSAR_*` name, such as `PULSAR_FORCE_TTY`, is applied to the environment unchanged.

## Behavior variables

These variables control the harness output path and default mode. They do not configure provider authentication. Explicit flags win over env vars: `--json`, then `--no-tui`, then `--tui`, then these env vars, then TTY auto-detect.

| Variable | Required | Secret | Description |
| --- | --- | --- | --- |
| `GH_TEST_AUTH_MODE` | No | No | Sets the default mode when you do not pass `--mode`. The default is `anonymous`. |
| `PULSAR_ENV_FILE` | No | No | Sets the env file path to load when `--env-file` is omitted. |
| `PULSAR_FORCE_TTY` | No | No | Set to `1` to force the optional dashboard path when no explicit output flag is passed. |
| `PULSAR_NO_TUI` | No | No | Set to `1` to prevent the optional dashboard from launching when no explicit output flag is passed. |
| `NO_COLOR` | No | No | Set to `1` to use ASCII glyphs and no color. This also disables the spinner braille. |
| `TPT_PROVIDER_ROOT` | Only for the tagged live smoke test | No | Points `TestSmokeAnonymousIpRanges` at the terraform-provider-github checkout to exercise. It is not used by normal CLI commands. |

`PULSAR_FORCE_TTY` and `PULSAR_NO_TUI` are not mode names, so they are not reserved and can be set through the single env file like any other variable.

## Provider variables

These variables affect the provider test process.

| Variable | Required | Secret | Description |
| --- | --- | --- | --- |
| `TF_ACC` | Yes | No | Enables acceptance tests. The harness sets `TF_ACC=1` for runs. |
| `GITHUB_BASE_URL` | No | No | Sets the GitHub Enterprise Server base URL. |
| `GITHUB_LEGACY_CLIENT` | No | No | Optional provider-level variable. |

## Authentication

Each non-anonymous mode requires exactly one authentication method: either `GITHUB_TOKEN` or the GitHub App trio `GITHUB_APP_ID`, `GITHUB_APP_INSTALLATION_ID`, and `GITHUB_APP_PEM_FILE`. Token auth and GitHub App auth are mutually exclusive. `GITHUB_APP_PEM_FILE` is the PEM contents, not a path. Use `\n` for newlines when exporting it through a shell or secret manager.

| Variable | Required | Secret | Description |
| --- | --- | --- | --- |
| `GITHUB_TOKEN` | Required for non-anonymous modes unless you use GitHub App auth | Yes | Personal access token for token auth. |
| `GITHUB_APP_ID` | Required for non-anonymous modes when you use GitHub App auth | No | GitHub App ID for GitHub App auth. |
| `GITHUB_APP_INSTALLATION_ID` | Required for non-anonymous modes when you use GitHub App auth | No | GitHub App installation ID for GitHub App auth. |
| `GITHUB_APP_PEM_FILE` | Required for non-anonymous modes when you use GitHub App auth | Yes | GitHub App PEM contents. Use `\n` for newlines. |

## Mode-specific variables

Anonymous mode requires no environment variables and uses public endpoints only.

### Individual

Individual mode also requires one authentication method. For more information, see [Test mode reference](./test-modes.md).

| Variable | Required | Secret | Description |
| --- | --- | --- | --- |
| `GITHUB_OWNER` | Yes | No | Required owner for individual mode. |
| `GITHUB_USERNAME` | Yes | No | Required username for individual mode. |
| `GH_TEST_ORG_TEMPLATE_REPOSITORY` | No | No | Template repository owned by `GITHUB_OWNER`; required when the selected tests request the `template-repository` capability. |
| `GH_TEST_USER_REPOSITORY` | No | No | Optional user repository for individual mode. |

### Organization

Organization mode also requires one authentication method.

| Variable | Required | Secret | Description |
| --- | --- | --- | --- |
| `GITHUB_OWNER` | Yes | No | Required organization owner. |
| `GH_TEST_ORG_USER1` | Yes | No | Required organization user. |
| `GH_TEST_ORG_REPOSITORY` | Yes | No | Required organization repository. |
| `GH_TEST_ORG_TEMPLATE_REPOSITORY` | Yes | No | Required organization template repository. It must exist and be marked as a template. |
| `GH_TEST_ORG_SECRET_NAME` | Yes | No | Required organization secret name. |
| `GH_TEST_ORG_USER3` | No | No | Optional organization mode variable. |
| `GH_TEST_ORG_APP_INSTALLATION_ID` | No | No | Optional organization mode variable. |

### Team

Team mode includes the organization variables and also requires the team variables in this table.

| Variable | Required | Secret | Description |
| --- | --- | --- | --- |
| `GITHUB_OWNER` | Yes | No | Required organization owner. |
| `GH_TEST_ORG_USER1` | Yes | No | Required organization user. |
| `GH_TEST_ORG_REPOSITORY` | Yes | No | Required organization repository. |
| `GH_TEST_ORG_TEMPLATE_REPOSITORY` | Yes | No | Required organization template repository. It must exist and be marked as a template. |
| `GH_TEST_ORG_SECRET_NAME` | Yes | No | Required organization secret name. |
| `GH_TEST_ORG_USER2` | Yes | No | Required team mode variable. |
| `GH_TEST_EXTERNAL_USER1` | Yes | No | Required team mode variable. |
| `GH_TEST_EXTERNAL_USER1_TOKEN` | Yes | Yes | Required team mode token. |
| `GH_TEST_EXTERNAL_USER2` | Yes | No | Required team mode variable. |
| `GH_TEST_ORG_USER3` | No | No | Optional organization mode variable. |
| `GH_TEST_ORG_APP_INSTALLATION_ID` | No | No | Optional organization mode variable. |

### Enterprise

Enterprise mode includes the organization variables and also requires `GITHUB_ENTERPRISE_SLUG`.

| Variable | Required | Secret | Description |
| --- | --- | --- | --- |
| `GITHUB_OWNER` | Yes | No | Required organization owner. |
| `GH_TEST_ORG_USER1` | Yes | No | Required organization user. |
| `GH_TEST_ORG_REPOSITORY` | Yes | No | Required organization repository. |
| `GH_TEST_ORG_TEMPLATE_REPOSITORY` | Yes | No | Required organization template repository. It must exist and be marked as a template. |
| `GH_TEST_ORG_SECRET_NAME` | Yes | No | Required organization secret name. |
| `GITHUB_ENTERPRISE_SLUG` | Yes | No | Required enterprise slug. |
| `GH_TEST_ORG_USER3` | No | No | Optional organization mode variable. |
| `GH_TEST_ORG_APP_INSTALLATION_ID` | No | No | Optional organization mode variable. |
| `GH_TEST_ENTERPRISE_IS_EMU` | No | No | Optional enterprise mode variable. |
| `GH_TEST_ENTERPRISE_EMU_GROUP_ID` | No | No | Optional enterprise mode variable. |
| `GH_TEST_ADVANCED_SECURITY` | No | No | Optional enterprise mode variable. |
| `GITHUB_BASE_URL` | No | No | Optional GitHub Enterprise Server base URL. |

## Redacted variables

The harness masks secret values in every persisted artifact, report, failure log, and dashboard message.

| Variable | Required | Secret | Description |
| --- | --- | --- | --- |
| `GITHUB_TOKEN` | Required for non-anonymous modes unless you use GitHub App auth | Yes | Redacted token auth value. |
| `GH_TEST_EXTERNAL_USER1_TOKEN` | Required for team mode | Yes | Redacted external user token value. |
