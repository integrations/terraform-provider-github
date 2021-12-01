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
      version = "~> 4.0"
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

Terraform 0.12 and earlier:

```terraform
# Configure the GitHub Provider
provider "github" {
  version = "~> 4.0"
}

# Add a user to the organization
resource "github_membership" "membership_for_user_x" {
  # ...
}
```
~> **Note:** When upgrading from `hashicorp/github` to `integrations/github`, use `terraform state replace-provider`. Otherwise, Terraform will still require the old provider to interact with the state file.

## Authentication

The GitHub provider offers multiple ways to authenticate with GitHub API.

### OAuth / Personal Access Token

To authenticate using OAuth tokens, ensure that the `token` argument or the `GITHUB_TOKEN` environment variable is set.

```terraform
provider "github" {
  token = var.token # or `GITHUB_TOKEN`
}
```

### GitHub App Installation

To authenticate using a GitHub App installation, ensure that arguments in the `app_auth` block or the `GITHUB_APP_XXX` environment variables are set.

Some API operations may not be available when using a GitHub App installation configuration. For more information, refer to the list of [supported endpoints](https://docs.github.com/en/rest/overview/endpoints-available-for-github-apps).

```terraform
provider "github" {
  app_auth {
    id              = var.app_id              # or `GITHUB_APP_ID`
    installation_id = var.app_installation_id # or `GITHUB_APP_INSTALLATION_ID`
    pem_file        = var.app_pem_file        # or `GITHUB_APP_PEM_FILE`
  }
}
```

~> **Note:** When using environment variables, an empty `app_auth` block is required to allow provider configurations from environment variables to be specified. See: https://github.com/hashicorp/terraform-plugin-sdk/issues/142

```terraform
provider "github" {
  app_auth {} # When using `GITHUB_APP_XXX` environment variables
}
```

## Argument Reference

The following arguments are supported in the `provider` block:

* `token` - (Optional) A GitHub OAuth / Personal Access Token. When not provided or made available via the `GITHUB_TOKEN` environment variable, the provider can only access resources available anonymously.

* `base_url` - (Optional) This is the target GitHub base API endpoint. Providing a value is a requirement when working with GitHub Enterprise. It is optional to provide this value and it can also be sourced from the `GITHUB_BASE_URL` environment variable. The value must end with a slash, for example: `https://terraformtesting-ghe.westus.cloudapp.azure.com/`

* `owner` - (Optional) This is the target GitHub organization or individual user account to manage. For example, `torvalds` and `github` are valid owners. It is optional to provide this value and it can also be sourced from the `GITHUB_OWNER` environment variable. When not provided and a `token` is available, the individual user account owning the `token` will be used. When not provided and no `token` is available, the provider may not function correctly. If both `owner` and `GITHUB_OWNER` are set, `owner` takes priority. // TODO(kfcampbell): is this true? (no) should it be true? (yes)

* `app_auth` - (Optional) Configuration block to use GitHub App installation token. When not provided, the provider can only access resources available anonymously.
  * `id` - (Required) This is the ID of the GitHub App. It can sourced from the `GITHUB_APP_ID` environment variable.
  * `installation_id` - (Required) This is the ID of the GitHub App installation. It can sourced from the `GITHUB_APP_INSTALLATION_ID` environment variable.
  * `pem_file` - (Required) This is the contents of the GitHub App private key PEM file. It can also be sourced from the `GITHUB_APP_PEM_FILE` environment variable and may use `\n` instead of actual new lines.

* `write_delay_ms` - (Optional) The number of milliseconds to sleep in between write operations in order to satisfy the GitHub API rate limits. Defaults to 1000ms or 1 second if not provided.

Note: If you have a PEM file on disk, you can pass it in via `pem_file = file("path/to/file.pem")`.

