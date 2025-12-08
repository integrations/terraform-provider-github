---
layout: "github"
page_title: "GitHub: github_codespaces_secrets"
description: |-
  Get codespaces secrets for a repository
---

# github\_codespaces\_secrets

Use this data source to retrieve the list of codespaces secrets for a GitHub repository.

## Example Usage

```hcl
data "github_codespaces_secrets" "example" {
  name = "example_repository"
}

data "github_codespaces_secrets" "example_2" {
  full_name = "org/example_repository"
}
```

## Argument Reference

 * `name` - (Optional) The name of the repository.
 * `full_name` - (Optional) Full name of the repository (in `org/name` format).

## Attributes Reference

 * `secrets` - list of codespaces secrets for the repository
   * `name` - Secret name
   * `created_at` - Timestamp of the secret creation
   * `updated_at` - Timestamp of the secret last update
 
