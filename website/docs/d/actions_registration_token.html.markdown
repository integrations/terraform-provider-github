---
layout: "github"
page_title: "GitHub: actions_registration_token"
description: |-
  Get a GitHub Actions repository registration token.
---

# actions_registration_token

Use this data source to retrieve a GitHub Actions repository registration token. This token can then be used to register a self-hosted runner.

## Example Usage

```hcl
data "github_actions_registration_token" "example" {
  repository = "example_repo"
}
```

## Argument Reference

 * `repository` - (Required) Name of the repository to get a GitHub Actions registration token for.

## Attributes Reference

 * `token` - The token that has been retrieved.
 * `expires_at` - The token expiration date.
