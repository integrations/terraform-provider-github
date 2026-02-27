---
page_title: "GitHub: github_actions_variable"
description: |-
  Creates and manages an Action variable within a GitHub repository
---

# github_actions_variable

This resource allows you to create and manage GitHub Actions variables within your GitHub repositories. You must have write access to a repository to use this resource.

## Example Usage

```terraform
resource "github_actions_variable" "example_variable" {
  repository       = "example_repository"
  variable_name    = "example_variable_name"
  value            = "example_variable_value"
}
```

## Argument Reference

The following arguments are supported:

- `repository` - (Required) Name of the repository.
- `variable_name` - (Required) Name of the variable.
- `value` - (Required) Value of the variable.

## Attributes Reference

- `repository_id` - ID of the repository.
- `created_at` - Date the variable was created.
- `updated_at` - Date the variable was last updated.

## Import

This resource can be imported using an ID made of the repository name, and variable name separated by a `:`.

### Import Block

The following import imports a GitHub actions variable named `myvariable` for the repo `myrepo` to a `github_actions_variable` resource named `example`.

```terraform
import {
  to = github_actions_variable.example
  id = "myrepo:myvariable"
}
```

### Import Command

The following command imports a GitHub actions variable named `myvariable` for the repo `myrepo` to a `github_actions_variable` resource named `example`.

```shell
terraform import github_actions_variable.example myrepo:myvariable
```
