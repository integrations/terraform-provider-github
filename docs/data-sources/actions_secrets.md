---
layout: "github"
page_title: "GitHub: github_actions_secrets"
description: |-
  Get actions secrets for a repository
---

# github\_actions\_secrets

Use this data source to retrieve the list of secrets for a GitHub repository.

## Example Usage

```hcl
data "github_actions_secrets" "example" {
  name = "example"
}
```

## Argument Reference

 * `name` - (Optional) The name of the repository.
 * `full_name` - (Optional) Full name of the repository (in `org/name` format).

## Attributes Reference

 * `secrets` - list of secrets for the repository
   * `name` - Secret name
   * `created_at` - Timestamp of the secret creation
   * `updated_at` - Timestamp of the secret last update
 
