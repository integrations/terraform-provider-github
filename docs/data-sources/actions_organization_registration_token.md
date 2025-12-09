---
page_title: "github_actions_organization_registration_token Data Source - terraform-provider-github
description: |-
  Get a GitHub Actions organization registration token.
---

# github_actions_organization_registration_token (Data Source)

Use this data source to retrieve a GitHub Actions organization registration token. This token can then be used to register a self-hosted runner.

## Example Usage

```terraform
data "github_actions_organization_registration_token" "example" {
}
```

## Argument Reference

## Attributes Reference

- `token` - The token that has been retrieved.
- `expires_at` - The token expiration date.
