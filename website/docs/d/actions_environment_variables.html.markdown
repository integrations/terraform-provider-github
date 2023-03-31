---
layout: "github"
page_title: "GitHub: github_actions_environment_variables"
description: |-
  Get Actions variables of the repository environment
---

# github\_actions\_environment\_variables

Use this data source to retrieve the list of variables of the repository environment.

## Example Usage

```hcl
data "github_actions_environment_variables" "example" {
    name        = "exampleRepo"
    environment = "exampleEnvironment"
}
```

## Argument Reference

## Attributes Reference

 * `variables` - list of variables for the environment
   * `name`         - Name of the variable
   * `value`        - Value of the variable
   * `created_at`   - Timestamp of the variable creation
   * `updated_at`   - Timestamp of the variable last update