---
layout: "github"
page_title: "GitHub: actions_organization_remove_token"
description: |-
  Get a GitHub Actions organization remove token.
---

# actions_regmove_token

Use this data source to retrieve a GitHub Actions organization remove token. This token can then be used to register a self-hosted runner.

## Example Usage

```hcl
data "github_actions_organization_remove_token" "example" {
}
```

## Argument Reference

## Attributes Reference

 * `token` - The token that has been retrieved.
 * `expires_at` - The token expiration date.