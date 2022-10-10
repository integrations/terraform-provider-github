---
layout: "github"
page_title: "GitHub: github_dependabot_secrets"
description: |-
  Get dependabot secrets for a repository
---

# github\dependabot\_secrets

Use this data source to retrieve the list of dependabot secrets for a GitHub repository.

## Example Usage

```hcl
data "github_dependabot_secrets" "example" {
  name = "example"
}
```

## Argument Reference

 * `name` - (Optional) The name of the repository.
 * `full_name` - (Optional) Full name of the repository (in `org/name` format).

## Attributes Reference

 * `secrets` - list of dependabot secrets for the repository
   * `name` - Secret name
   * `created_at` - Timestamp of the secret creation
   * `updated_at` - Timestamp of the secret last update
 
