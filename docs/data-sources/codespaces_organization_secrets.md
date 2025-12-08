---
layout: "github"
page_title: "GitHub: github_codespaces_organization_secrets"
description: |-
  Get codespaces secrets of the organization
---

# github\_codespaces\_organization\_secrets

Use this data source to retrieve the list of codespaces secrets of the organization.

## Example Usage

```hcl
data "github_codespaces_organization_secrets" "example" {
}
```

## Argument Reference

## Attributes Reference

 * `secrets` - list of secrets for the repository
   * `name` - Secret name
   * `visibility` - Secret visibility
   * `created_at` - Timestamp of the secret creation
   * `updated_at` - Timestamp of the secret last update
 
