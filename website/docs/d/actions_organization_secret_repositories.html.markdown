---
layout: "github"
page_title: "GitHub: github_actions_organization_secret_repositories"
description: |-
  Get actions repositories of secret of the organization
---

# github\_actions\_organization\_secret\_repositories

Use this data source to retrieve the list of repositories of secret of the organization.

## Example Usage

```hcl
data "github_actions_organization_secret_repositories" "example_secrets" {
  secret_name = "example_secret"
}
```

## Argument Reference

  * `secret_name` - (Required) The name of the secret.

## Attributes Reference

 * `repositories` - list of repositories of secret
   * `name` - Secret name
   * `full_name` - Full name of the repository
