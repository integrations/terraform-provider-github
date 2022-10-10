---
layout: "github"
page_title: "GitHub: github_actions_organization_secrets"
description: |-
  Get actions secrets of the organization
---

# github\_actions\_organization\_secrets

Use this data source to retrieve the list of secrets of the organization.

## Example Usage

```hcl
data "github_actions_organization_secrets" "example" {
}
```

## Argument Reference

## Attributes Reference

 * `secrets` - list of secrets for the repository
   * `name` - Secret name
   * `visibility` - Secret visibility
   * `created_at` - Timestamp of the secret creation
   * `updated_at` - Timestamp of the secret last update
 
