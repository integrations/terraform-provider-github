---
page_title: "GitHub Provider"
description: |-
  The GitHub Terraform provider is used to interact with GitHub resources either as an authenticated client or anonymously.
---

# GitHub Provider

The GitHub Terraform provider is used to interact with GitHub resources either as an authenticated client or anonymously.

!> You **must** add a `required_providers` block to every module that will create resources with this provider. If you do not explicitly require `integrations/github` in a submodule, your Terraform code run may [break in hard-to-troubleshoot ways](https://github.com/integrations/terraform-provider-github/issues/876#issuecomment-1303790559).

## Example Usage

```terraform
terraform {
  required_providers {
    github = {
      source  = "integrations/github"
      version = "~> 6.0"
    }
  }
}

provider "github" {
  owner = "integrations"
}

data "github_repository" "example" {
  name = "terraform-provider-github"
}
```

## Owner

For backwards compatibility; if more than one of `owner`, `organization`, `GITHUB_OWNER` and `GITHUB_ORGANIZATION` are set the first in this list takes priority.

1. Setting `organization` in the GitHub provider configuration.
2. Setting the `GITHUB_ORGANIZATION` environment variable.
3. Setting the `GITHUB_OWNER` environment variable.
4. Setting `owner` in the GitHub provider configuration.

!> It is a bug that `GITHUB_OWNER` takes precedence over `owner`; this will be fixed in a future major release. For compatibility with future releases, please set only one of `GITHUB_OWNER` and `owner`.

## Authentication

The GitHub provider can be authenticated with the GitHub API via a GitHub App, an OAuth Token, or a Personal Access Token (PAT); it can also operate anonymously (in a limited manner) if no authentication is provided. The provider selects the authentication used based on the `auth_mode` argument, with the ability to explicitly set the authentication mode and falling back to `auto` mode when a specific mode is not explicitly set. Auto mode uses the following authentication fallback chain (first match wins):

1. [**GitHub App Installation**](#github-app-installation) — `app_auth` block (or environment) with `id`, `installation_id`, and `pem_file`.
2. [**Explicit Token**](#oauth-or-personal-access-token-pat) — `token` argument or `GITHUB_TOKEN` environment variable.
3. [**GitHub CLI**](#github-cli-authentication) — Falls back to `gh auth token` if neither token nor app_auth is set.
4. **Anonymous** — Read-only access when no credentials are available.

### GitHub App Installation

GitHub App authentication requires the `owner` argument to be set and is supported by the `app_auth` provider configuration block and/or the related environment variables. Authenticating the provider with a GitHub App requires all three `app_auth` arguments to be set; `id`, `installation_id`, and `pem_file`. If you want to make sure that the provider is using GitHub App authentication, you can set the `auth_mode` argument to `app` (setting the `app_auth` block also requires the provider to use GitHub App authentication). When using environment variables the provider defaults to using the `GITHUB_APP_` prefix, but this can be overridden with the `app_auth_env_prefix` argument.

By default, the provider will look for the following environment variables: `GITHUB_APP_ID`, `GITHUB_APP_INSTALLATION_ID`, and `GITHUB_APP_PEM_FILE`. If you want to use a different prefix, you can set the `app_auth_env_prefix` argument to the desired prefix. For example, if you set `app_auth_env_prefix = "MYAPP_"`, the provider will look for the following environment variables: `MYAPP_ID`, `MYAPP_INSTALLATION_ID`, and `MYAPP_PEM_FILE`. This is useful if you want to use multiple GitHub providers authenticating with different GitHub Apps.

Some API operations may not be available when using a GitHub App installation configuration. For more information, refer to the list of [supported endpoints](https://docs.github.com/en/rest/overview/endpoints-available-for-github-apps).

#### Block with values

```terraform
provider "github" {
  owner = var.github_organization

  auth_mode = "app" # or `GITHUB_AUTH_MODE=app`

  app_auth {
    id              = var.app_id              # or `GITHUB_APP_ID`
    installation_id = var.app_installation_id # or `GITHUB_APP_INSTALLATION_ID`
    pem_file        = var.app_pem_file        # or `GITHUB_APP_PEM_FILE`
  }
}
```

#### Environment variables only

```shell
export GITHUB_APP_ID="12332432" # Required: The GitHub App ID for authentication
export GITHUB_APP_INSTALLATION_ID="12435523" # Required: The GitHub App Installation ID for authentication
export GITHUB_APP_PEM_FILE="..." # Required: Contents of the PEM file for the GitHub App, not the path to the PEM file
```

```terraform
provider "github" {
  owner = var.github_organization

  auth_mode = "app" # Credentials required to come from the `GITHUB_APP_XXX` environment variables.
}
```

#### Non-secret values in configuration, secret from the environment

```shell
export GITHUB_APP_PEM_FILE="..." # Required: Contents of the PEM file for the GitHub App, not the path to the PEM file
```

```terraform
provider "github" {
  owner = var.github_organization

  auth_mode = "app"

  app_auth {
    id              = "123456"   # the App ID and installation ID are not secret,
    installation_id = "78901234" # so they can be set directly in configuration
    # pem_file is omitted; it falls back to the GITHUB_APP_PEM_FILE variable,
    # keeping the secret out of configuration
  }
}
```

### Opt-In Enterprise App Lifecycle

The optional `enterprise_app_auth` block enables an App-only path for
`github_enterprise_organization`. It does not change default/PAT authentication
or route other organization resources dynamically. Use a separate normally
configured provider for repositories, teams, and other organization resources.
This feature requires a provider build containing this opt-in implementation.

Set `auth_mode = "app"`. The existing `app_auth` credentials are the enterprise
creation App; its installation must target the configured enterprise. `owner`
is not required or looked up in this mode. No PAT or GitHub CLI fallback is
allowed: remove `token` and the environment variable selected by `token_env_name`
(normally `GITHUB_TOKEN`). Clients renew installation tokens automatically.

```terraform
terraform {
  required_providers {
    github = {
      source = "integrations/github"
    }
  }
}

provider "github" {
  auth_mode = "app"

  app_auth {
    id              = var.creation_app_id
    installation_id = var.creation_installation_id
    pem_file        = var.creation_private_key
  }

  enterprise_app_auth {
    enterprise_slug           = var.enterprise_slug
    installer_app_id          = var.installer_app_id
    installer_installation_id = var.installer_installation_id
    installer_pem_file        = var.installer_private_key
    automation_app_id         = var.automation_app_id
    automation_client_id      = var.automation_client_id
    automation_pem_file       = var.automation_private_key
  }
}

data "github_enterprise" "selected" {
  slug = var.enterprise_slug
}

resource "github_enterprise_organization" "managed" {
  for_each = var.organizations

  enterprise_id = data.github_enterprise.selected.id
  name          = each.key
  billing_email = each.value.billing_email
  admin_logins  = each.value.admin_logins

  lifecycle {
    prevent_destroy = true
  }
}

variable "enterprise_slug" {
  type        = string
  description = "Enterprise managed by the creation and installer Apps."
}

variable "creation_app_id" {
  type        = string
  description = "Enterprise creation App ID or client ID."
}

variable "creation_installation_id" {
  type        = string
  description = "Enterprise creation App installation ID."
}

variable "creation_private_key" {
  type        = string
  sensitive   = true
  description = "Creation App private key PEM content."
}

variable "installer_app_id" {
  type        = string
  description = "Installer App ID or client ID."
}

variable "installer_installation_id" {
  type        = string
  description = "Installer enterprise installation ID."
}

variable "installer_private_key" {
  type        = string
  sensitive   = true
  description = "Installer App private key PEM content."
}

variable "automation_app_id" {
  type        = string
  description = "Numeric automation App ID."
}

variable "automation_client_id" {
  type        = string
  description = "Automation App client ID matching the numeric ID and private key."
}

variable "automation_private_key" {
  type        = string
  sensitive   = true
  description = "Automation App private key PEM content."
}

variable "organizations" {
  type = map(object({
    billing_email = string
    admin_logins  = set(string)
  }))
  description = "Organizations to create or import into the configured enterprise."
}
```

Supply the new block's credentials as attributes or with these dedicated
environment defaults. Environment variables alone do not enable the feature;
the block with `enterprise_slug` must be present.

| Attribute | Environment Default |
| --- | --- |
| `installer_app_id` | `INSTALLER_APP_CLIENT_ID` |
| `installer_installation_id` | `INSTALLER_APP_INSTALLATION_ID` |
| `installer_pem_file` | `INSTALLER_APP_PRIVATE_KEY` |
| `automation_app_id` | `ORG_APP_ID` |
| `automation_client_id` | `ORG_APP_CLIENT_ID` |
| `automation_pem_file` | `ORG_APP_PRIVATE_KEY` |

`installer_app_id` accepts the numeric App ID or client ID JWT issuer.
`automation_app_id` is the numeric App ID; `automation_client_id` must identify
the same App. PEM attributes/environment variables contain private key content,
not file paths. Keep all private key inputs sensitive in HCP Terraform.
All installer and automation credentials are required, whether supplied as
attributes or environment variables. Missing credentials fail configuration.

The installer must have access to the enterprise's installable organization
inventory and permission to install organization Apps. The automation App must
have `organization_administration = write` and `members = read` (or `write`).
Repository permissions other than metadata are rejected because bootstrap uses
`repository_selection = "none"`. The creation App needs the enterprise
organization-creation permission. See the GitHub
[enterprise organization APIs](https://docs.github.com/en/rest/enterprise-admin/organizations)
and [App installation APIs](https://docs.github.com/en/rest/apps/installations)
for endpoint permission requirements.

Read/import never install, update, or unsuspend Apps. Creation installs the
automation App only if its organization installation is absent. Existing
installations must be active and already have the required permissions; this
provider does not upgrade them. The `automation_installation_id` resource
attribute exposes the discovered installation ID, but does not configure other
provider instances in the same run.

Updates and deletion of enterprise organizations are explicitly blocked in
this mode. For partial-create recovery and import behavior, see
[github_enterprise_organization](resources/enterprise_organization).

### OAuth or Personal Access Token (PAT)

To authenticate using OAuth tokens, ensure that the `token` argument or the `GITHUB_TOKEN` environment variable is set.

```terraform
provider "github" {
  owner = "octocat"

  auth_mode = "token" # or `GITHUB_AUTH_MODE=token`

  token = var.token # or `GITHUB_TOKEN`
}
```

### GitHub CLI Authentication

~> The GitHub CLI authentication fallback is only available when using the legacy client implementation and will be removed in a future major release. To use the legacy client implementation, set `legacy_client = true` in the provider configuration or set the `GITHUB_LEGACY_CLIENT` environment variable to `true`.

When using the GitHub CLI authentication fallback, you can optionally specify the path to the `gh` executable using the `GH_PATH` environment variable. This is useful when the provider cannot properly determine the path to GitHub CLI, such as in cygwin terminals. If not specified, the provider looks for `gh` in your system PATH.

#### .env

```shell
export GH_PATH="/path/to/gh" # Optional: Specify the path to the GitHub CLI executable if not in system PATH
```

#### main.tf

```terraform
provider "github" {
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `app_auth` (Block List, Max: 1) Authenticate using a GitHub App. (see [below for nested schema](#nestedblock--app_auth))
- `app_auth_env_prefix` (String) The environment variable prefix for the GitHub App authentication used to determine the environment variable names for the GitHub App's ID (`<PREFIX>_ID`), installation ID (`<PREFIX>_INSTALLATION_ID`), and PEM file content (`<PREFIX>_PEM_FILE`). This defaults to `GITHUB_APP_`.
- `auth_mode` (String) The authentication mode to use; this can be one of `auto`, `app`, `token` or `none` and defaults to `auto` which will detect the highest priority authentication mode available (`app` -> `token` -> `none`). This can also be set by the `GITHUB_AUTH_MODE` environment variable.
- `base_url` (String) The base URL for the GitHub API; this defaults to the GitHub API URL. If you are using GitHub Enterprise Server (GHES) or GitHub Enterprise Cloud with Data Residency (GHEC-DR), this is required. This can also be set by the `GITHUB_BASE_URL` environment variable.
- `cache_path` (String) The path to the cache directory for persisting GitHub API requests between runs; if not set there will be no caching between runs. This can also be set by the `GITHUB_CACHE_PATH` environment variable.
- `enterprise_app_auth` (Block List, Max: 1) Opt-in enterprise App lifecycle credentials. All credentials must be supplied here or through their dedicated environment variables. (see [below for nested schema](#nestedblock--enterprise_app_auth))
- `insecure` (Boolean, Deprecated) Allow insecure server connections when using SSL.
- `legacy_client` (Boolean) Use the legacy GitHub client implementation; if set to `false`, the new client implementation is used. This can also be set by the `GITHUB_LEGACY_CLIENT` environment variable.
- `max_per_page` (Number) The maximum number of results per page for paginated API requests; this defaults to `100`. This can also be set by the `GITHUB_MAX_PER_PAGE` environment variable.
- `max_retries` (Number) The maximum number of retries for failed requests; this defaults to `3`.
- `organization` (String, Deprecated) GitHub organization to manage. This can also be set by the `GITHUB_ORGANIZATION` environment variable.
- `owner` (String) GitHub organization or user account to manage; this is required when authenticating using a GitHub App. If the owner is not provided and a token is provided, the provider will attempt to auto-detect the owner associated with the token. This can also be set by the `GITHUB_OWNER` environment variable.
- `parallel_requests` (Boolean) Allow the provider to make parallel API calls; this is experimental and may cause concurrency and rate limiting issues. This is ignored for the REST API when `legacy_client` is `false` since the new client implementation is designed to safely handle parallel requests.
- `read_delay_ms` (Number) The delay in milliseconds between read operations; this defaults to `0`. This can be used to mitigate rate limiting issues when performing a large number of read operations. This is ignored for the REST API when `legacy_client` is `false` since the new client implementation is GitHub rate limit aware.
- `retry_delay_ms` (Number) The delay in milliseconds between retry attempts; this defaults to `1000`. This setting only applies when `max_retries` is greater than `0`.
- `retryable_errors` (List of Number) List of HTTP status codes that should be retried; if not set this uses the provider defaults. This setting only applies when `max_retries` is greater than `0`. This is ignored for the REST API when `legacy_client` is `false` since the new client implementation handles the retry logic.
- `token` (String) GitHub OAuth or Personal Access Token (PAT) to use for authentication. This can also be set by the environment variable specified in `token_env_name` which defaults to `GITHUB_TOKEN`.
- `token_env_name` (String) The environment variable name for the GitHub token. This defaults to `GITHUB_TOKEN`.
- `write_delay_ms` (Number) The delay in milliseconds between write operations; this defaults to `1000`. This is used to mitigate the GitHub API's abuse rate limits when writing. Note that **ALL** requests to the GraphQL API are implemented as `POST` requests under the hood, so this setting affects those calls as well. This is ignored for the REST API when `legacy_client` is `false` since the new client implementation is GitHub rate limit aware.

<a id="nestedblock--app_auth"></a>
### Nested Schema for `app_auth`

Optional:

- `id` (String) The GitHub App's identifier. This can also be set by the `GITHUB_APP_ID` environment variable when `app_auth_env_prefix` is `GITHUB_APP_` (modify the prefix as needed).
- `installation_id` (String) The GitHub App's installation identifier. This can also be set by the `GITHUB_APP_INSTALLATION_ID` environment variable when `app_auth_env_prefix` is `GITHUB_APP_` (modify the prefix as needed).
- `pem_file` (String, Sensitive) The GitHub App's PEM file content; `\n` can be used for newlines. This can also be set by the `GITHUB_APP_PEM_FILE` environment variable when `app_auth_env_prefix` is `GITHUB_APP_` (modify the prefix as needed).


<a id="nestedblock--enterprise_app_auth"></a>
### Nested Schema for `enterprise_app_auth`

Required:

- `enterprise_slug` (String) Enterprise slug to manage. Opts into App-only enterprise organization lifecycle; deletion and updates are blocked.

Optional:

- `automation_app_id` (String) Automation App numeric ID. Defaults to ORG_APP_ID.
- `automation_client_id` (String) Automation App client ID. Defaults to ORG_APP_CLIENT_ID.
- `automation_pem_file` (String, Sensitive) Automation private key PEM content, not a filename. Defaults to ORG_APP_PRIVATE_KEY.
- `installer_app_id` (String) Installer App numeric ID or client ID. Defaults to INSTALLER_APP_CLIENT_ID.
- `installer_installation_id` (String) Installer enterprise installation ID. Defaults to INSTALLER_APP_INSTALLATION_ID.
- `installer_pem_file` (String, Sensitive) Installer private key PEM content, not a filename. Defaults to INSTALLER_APP_PRIVATE_KEY.
