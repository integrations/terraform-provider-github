---
layout: "github"
page_title: "Provider: GitHub"
sidebar_current: "docs-github-index"
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

* `token` - (Optional) This is the GitHub personal access token. It can also be
  sourced from the `GITHUB_TOKEN` environment variable. If `anonymous` is false,
  token is required.

* `owner` - (Required) This is the target GitHub organization or a user to manage.
  The account corresponding to the token will need "owner" privileges for this
  organization. It can also be sourced from the `GITHUB_OWNER`
  environment variable.

* `organization` - (DEPRICATED) This is the target GitHub organization or a user to manage. The account
  corresponding to the token will need "organization" privileges for this organization. It must be provided, but
  it can also be sourced from the `GITHUB_ORGANIZATION` environment variable.

* `base_url` - (Optional) This is the target GitHub base API endpoint. Providing a value is a
  requirement when working with GitHub Enterprise.  It is optional to provide this value and
  it can also be sourced from the `GITHUB_BASE_URL` environment variable.  The value must end with a slash,
  and generally includes the API version, for instance `https://github.someorg.example/api/v3/`.

* `insecure` - (Optional) Whether server should be accessed without verifying the TLS certificate.
  As the name suggests **this is insecure** and should not be used beyond experiments,
  accessing local (non-production) GHE instance etc.
  There is a number of ways to obtain trusted certificate for free, e.g. from [Let's Encrypt](https://letsencrypt.org/).
  Such trusted certificate *does not require* this option to be enabled.
  Defaults to `false`.

* `individual`: (Optional) Run outside an organization.  When `individual` is true, the provider will run outside
  the scope of an organization. Defaults to `false`.

* `anonymous`: (Optional) Authenticate without a token.  When `anonymous` is true, the provider will not be able to
  access resources that require authentication. Setting to true will lead the GitHub provider to work in an anonymous
  mode with the corresponding API [rate limits](https://developer.github.com/v3/#rate-limiting).  Defaults to `false`.
