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

```hcl
# Configure the GitHub Provider
provider "github" {
  token = "${var.github_token}"
  owner = "${var.github_owner}"
}

# Add a user to the organization
resource "github_membership" "membership_for_user_x" {
  # ...
}
```

## Argument Reference

The following arguments are supported in the `provider` block:

* `token` - (Optional) A GitHub OAuth / Personal Access Token. When not provided or made available via the `GITHUB_TOKEN` environment variable, the provider can only access resources available anonymously.

* `base_url` - (Optional) This is the target GitHub base API endpoint. Providing a value is a requirement when working with GitHub Enterprise. It is optional to provide this value and it can also be sourced from the `GITHUB_BASE_URL` environment variable. The value must end with a slash, for example: `https://terraformtesting-ghe.westus.cloudapp.azure.com/`

* `owner` - (Optional) This is the target GitHub organization or individual user account to manage. For example, `torvalds` and `github` are valid owners. It is optional to provide this value and it can also be sourced from the `GITHUB_OWNER` environment variable. When not provided and a `token` is available, the individual user account owning the `token` will be used. When not provided and no `token` is available, the provider may not function correctly.

* `organization` - (Deprecated) This behaves the same as `owner`, which should be used instead. This value can also be sourced from the `GITHUB_ORGANIZATION` environment variable.

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
