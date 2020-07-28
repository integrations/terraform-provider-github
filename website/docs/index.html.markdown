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

* `base_url` - (Optional) This is the target GitHub base API endpoint. Providing a value is a requirement when working with GitHub Enterprise.  It is optional to provide this value and it can also be sourced from the `GITHUB_BASE_URL` environment variable.  The value must end with a slash, for example: `https://terraformtesting-ghe.westus.cloudapp.azure.com/`

* `owner` - (Optional) This is the target GitHub individual account to manage.  It is optional to provide this value and it can also be sourced from the `GITHUB_OWNER` environment variable. For example, `torvalds` is a valid owner. When not provided and a `token` is available, the individual account owning the `token` will be used. When not provided and no `token` is available, the provider may not function correctly. Conflicts with `organization`.

* `organization` - (Optional) This is the target GitHub organization account to manage. It is optional to provide this value and it can also be sourced from the `GITHUB_ORGANIZATION` environment variable. For example, `github` is a valid organization. Conflicts with `owner`and requires `token`, as the individual account corresponding to provided `token` will need "owner" privileges for this organization.
