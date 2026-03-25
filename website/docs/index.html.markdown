---
layout: "github"
page_title: "Provider: GitHub"
description: |-
  The GitHub provider is used to interact with GitHub resources.
---

# GitHub Provider

The GitHub provider is used to interact with GitHub resources.

The provider allows you to manage your GitHub organization's members and teams easily.
It needs to be configured with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

Terraform 0.13 and later:

```terraform
terraform {
  required_providers {
    github = {
      source  = "integrations/github"
      version = "~> 6.0"
    }
  }
}

# Configure the GitHub Provider
provider "github" {}

# Add a user to the organization
resource "github_membership" "membership_for_user_x" {
  # ...
}
```

- You **must** add a `required_providers` block to every module that will create resources with this provider. If you do not explicitly require `integrations/github` in a submodule, your terraform run may [break in hard-to-troubleshoot ways](https://github.com/integrations/terraform-provider-github/issues/876#issuecomment-1303790559).

Terraform 0.12 and earlier:

```terraform
# Configure the GitHub Provider
provider "github" {
  version = "~> 5.0"
}

# Add a user to the organization
resource "github_membership" "membership_for_user_x" {
  # ...
}
```
~> **Note:** When upgrading from `hashicorp/github` to `integrations/github`, use `terraform state replace-provider`. Otherwise, Terraform will still require the old provider to interact with the state file.

## Authentication

The GitHub provider offers multiple ways to authenticate with GitHub API. You can explicitly set the authentication mode using the `auth_mode` argument, or let the provider auto-detect based on the provided credentials.

### Explicit Authentication Mode (Recommended)

Setting `auth_mode` explicitly is recommended for clarity and to avoid unexpected behavior.

#### Anonymous

```terraform
provider "github" {
  auth_mode = "anonymous"
}
```

When `auth_mode` is set to `"anonymous"`, the provider will not use any credentials, even if `GITHUB_TOKEN` or other authentication environment variables are set.

#### OAuth / Personal Access Token

```terraform
provider "github" {
  auth_mode = "token"
  token     = var.token # or `GITHUB_TOKEN`
}
```

When `auth_mode` is set to `"token"`, the provider requires the `token` argument or `GITHUB_TOKEN` environment variable. An error will be returned if no token is provided.

#### GitHub App Installation

```terraform
provider "github" {
  auth_mode           = "app"
  owner               = var.github_owner
  app_id              = var.app_id              # or `GITHUB_APP_ID`
  app_installation_id = var.app_installation_id # or `GITHUB_APP_INSTALLATION_ID`
  app_private_key     = var.app_private_key     # or `GITHUB_APP_PRIVATE_KEY`
}
```

When `auth_mode` is set to `"app"`, the provider requires all three app credential arguments (`app_id`, `app_installation_id`, `app_private_key`) or their corresponding environment variables. The `owner` argument is also required when using app authentication.

Some API operations may not be available when using a GitHub App installation configuration. For more information, refer to the list of [supported endpoints](https://docs.github.com/en/rest/overview/endpoints-available-for-github-apps).

### Auto-detected Authentication (Default)

When `auth_mode` is not set, the provider auto-detects the authentication method using the following priority:

1. If `token` (or `GITHUB_TOKEN`) is set, use token-based authentication.
2. If app credentials (`app_id`, `app_installation_id`, `app_private_key`, or the deprecated `app_auth` block) are set, use GitHub App authentication.
3. If none of the above, fall back to the GitHub CLI (`gh auth token`).
4. If no credentials are found, operate in anonymous mode.

This is equivalent to the pre-existing behavior and is preserved for backward compatibility.

~> **Note:** Using `auth_mode = "anonymous"` is the only way to ensure the provider runs anonymously when `GITHUB_TOKEN` or GitHub CLI credentials are present in the environment.

### GitHub CLI (Deprecated)

~> The GitHub CLI token fallback is deprecated and will be removed in a future major release. Please set the `token` provider argument or `GITHUB_TOKEN` environment variable explicitly. You can use `export GITHUB_TOKEN=$(gh auth token)` as a replacement.

### Deprecated: `app_auth` block

~> The `app_auth` block is deprecated. Use the top-level `app_id`, `app_installation_id`, and `app_private_key` arguments instead. Note that the deprecated `app_auth` block reads from `GITHUB_APP_PEM_FILE`, while the new `app_private_key` argument reads from `GITHUB_APP_PRIVATE_KEY`.

The `app_auth` block is still supported for backward compatibility:

```terraform
provider "github" {
  owner = var.github_organization
  app_auth {
    id              = var.app_id              # or `GITHUB_APP_ID`
    installation_id = var.app_installation_id # or `GITHUB_APP_INSTALLATION_ID`
    pem_file        = var.app_private_key     # or `GITHUB_APP_PEM_FILE`
  }
}
```

## Argument Reference

The following arguments are supported in the `provider` block:

* `auth_mode` - (Optional) Explicit authentication mode. Valid values are `anonymous`, `token`, and `app`. When not set, the provider auto-detects the mode based on provided credentials for backward compatibility. Can also be sourced from the `GITHUB_AUTH_MODE` environment variable.

* `token` - (Optional) A GitHub OAuth / Personal Access Token. When not provided or made available via the `GITHUB_TOKEN` environment variable, the provider can only access resources available anonymously.

* `base_url` - (Optional) This is the target GitHub base API endpoint. Providing a value is a requirement when working with GitHub Enterprise. It is optional to provide this value and it can also be sourced from the `GITHUB_BASE_URL` environment variable. The value must end with a slash, for example: `https://terraformtesting-ghe.westus.cloudapp.azure.com/`

* `owner` - (Optional) This is the target GitHub organization or individual user account to manage. For example, `torvalds` and `github` are valid owners. It is optional to provide this value and it can also be sourced from the `GITHUB_OWNER` environment variable. When not provided and a `token` is available, the individual user account owning the `token` will be used. When not provided and no `token` is available, the provider may not function correctly. It is required in case of GitHub App Installation.

* `organization` - (Deprecated) This behaves the same as `owner`, which should be used instead. This value can also be sourced from the `GITHUB_ORGANIZATION` environment variable.

* `app_id` - (Optional) This is the ID of the GitHub App. It can also be sourced from the `GITHUB_APP_ID` environment variable.

* `app_installation_id` - (Optional) This is the ID of the GitHub App installation. It can also be sourced from the `GITHUB_APP_INSTALLATION_ID` environment variable.

* `app_private_key` - (Optional) This is the contents of the GitHub App private key in PEM format. It can also be sourced from the `GITHUB_APP_PRIVATE_KEY` environment variable and may use `\n` instead of actual new lines. If you have a PEM file on disk, you can pass it in via `app_private_key = file("path/to/file.pem")`.

* `app_auth` - (Optional, **Deprecated**: Use top-level `app_id`, `app_installation_id`, and `app_private_key` instead.) Configuration block to use GitHub App installation token.
  * `id` - (Required) This is the ID of the GitHub App. It can sourced from the `GITHUB_APP_ID` environment variable.
  * `installation_id` - (Required) This is the ID of the GitHub App installation. It can sourced from the `GITHUB_APP_INSTALLATION_ID` environment variable.
  * `pem_file` - (Required) This is the contents of the GitHub App private key PEM file. It can also be sourced from the `GITHUB_APP_PEM_FILE` environment variable and may use `\n` instead of actual new lines.

* `write_delay_ms` - (Optional) The number of milliseconds to sleep in between write operations in order to satisfy the GitHub API rate limits. Note that requests to the GraphQL API are implemented as ``POST`` requests under the hood, so this setting affects those calls as well. Defaults to 1000ms or 1 second if not provided.

* `retry_delay_ms` - (Optional) Amount of time in milliseconds to sleep in between requests to GitHub API after an error response. Defaults to 1000ms or 1 second if not provided, the max_retries must be set to greater than zero.

* `read_delay_ms` - (Optional) The number of milliseconds to sleep in between non-write operations in order to satisfy the GitHub API rate limits. Defaults to 0ms.

* `retryable_errors` - (Optional) "Allow the provider to retry after receiving an error status code, the max_retries should be set for this to work. Defaults to [500, 502, 503, 504]

* `max_retries` - (Optional) Number of times to retry a request after receiving an error status code. Defaults to 3

For backwards compatibility, if more than one of `owner`, `organization`,
`GITHUB_OWNER` and `GITHUB_ORGANIZATION` are set, the first in this
list takes priority.

1. Setting `organization` in the GitHub provider configuration.
2. Setting the `GITHUB_ORGANIZATION` environment variable.
3. Setting the `GITHUB_OWNER` environment variable.
4. Setting `owner` in the GitHub provider configuration.

~> It is a bug that `GITHUB_OWNER` takes precedence over `owner`, which may
be fixed in a future major release. For compatibility with future releases,
please set only one of `GITHUB_OWNER` and `owner`.
