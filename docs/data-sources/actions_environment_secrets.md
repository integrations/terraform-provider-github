---
page_title: "github_actions_environment_secrets Data Source - terraform-provider-github
description: |-
  Get Actions secrets of the repository environment
---

# github_actions_environment_secrets (Data Source)

Use this data source to retrieve the list of secrets of the repository environment.

## Example Usage

```terraform
data "github_actions_environment_secrets" "example" {
    name        = "exampleRepo"
    environment = "exampleEnvironment"
}
```

## Argument Reference

## Attributes Reference

- `secrets` - list of secrets for the environment
    - `name` - Name of the secret
    - `created_at` - Timestamp of the secret creation
    - `updated_at` - Timestamp of the secret last update
