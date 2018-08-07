---
layout: "github"
page_title: "Provider: GitHub"
sidebar_current: "docs-github-index"
description: |-
  The GitHub provider is used to interact with GitHub organization resources.
---

# GitHub Provider

The GitHub provider is used to interact with GitHub organization resources.

The provider allows you to manage your GitHub organization's members and teams easily.
It needs to be configured with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
# Configure the GitHub Provider
provider "github" {
  token        = "${var.github_token}"
  organization = "${var.github_organization}"
}

# Add a user to the organization
resource "github_membership" "membership_for_user_x" {
  # ...
}
```

## Argument Reference

The following arguments are supported in the `provider` block:

* `token` - (Optional) This is the GitHub personal access token. It must be provided, but
  it can also be sourced from the `GITHUB_TOKEN` environment variable.

* `organization` - (Optional) This is the target GitHub organization to manage. The account
  corresponding to the token will need "owner" privileges for this organization. It must be provided, but
  it can also be sourced from the `GITHUB_ORGANIZATION` environment variable.

* `base_url` - (Optional) This is the target GitHub base API endpoint. Providing a value is a
  requirement when working with GitHub Enterprise.  It is optional to provide this value and
  it can also be sourced from the `GITHUB_BASE_URL` environment variable.  The value must end with a slash.

* `insecure` - (Optional) Whether server should be accessed without verifying the TLS certificate.
  As the name suggests **this is insecure** and should not be used beyond experiments,
  accessing local (non-production) GHE instance etc.
  There is a number of ways to obtain trusted certificate for free, e.g. from [Let's Encrypt](https://letsencrypt.org/).
  Such trusted certificate *does not require* this option to be enabled.
  Defaults to `false`.
