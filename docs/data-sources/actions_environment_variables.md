---
page_title: "github_actions_environment_variables Data Source - terraform-provider-github
description: |-
  Get Actions variables of the repository environment
---

# github_actions_environment_variables (Data Source)

Use this data source to retrieve the list of variables of the repository environment.

## Example Usage

```terraform
data "github_actions_environment_variables" "example" {
    name        = "exampleRepo"
    environment = "exampleEnvironment"
}
```

## Argument Reference

## Attributes Reference

- `variables` - list of variables for the environment
    - `name` - Name of the variable
    - `value` - Value of the variable
    - `created_at` - Timestamp of the variable creation
    - `updated_at` - Timestamp of the variable last update
