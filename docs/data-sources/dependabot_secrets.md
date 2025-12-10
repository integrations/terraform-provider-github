---
page_title: "github_dependabot_secrets Data Source - terraform-provider-github
description: |-
  Get dependabot secrets for a repository
---

# github_dependabot_secrets (Data Source)

Use this data source to retrieve the list of dependabot secrets for a GitHub repository.

## Example Usage

```terraform
data "github_dependabot_secrets" "example" {
  name = "example"
}
```

## Argument Reference

- `name` - (Optional) The name of the repository.
- `full_name` - (Optional) Full name of the repository (in `org/name` format).

## Attributes Reference

- `secrets` - list of dependabot secrets for the repository
    - `name` - Secret name
    - `created_at` - Timestamp of the secret creation
    - `updated_at` - Timestamp of the secret last update
