---
layout: "github"
page_title: "GitHub: github_actions_variables"
description: |-
  Creates and manages multiple Action variables within a GitHub repository
---

# github_actions_variables

This resource allows you to create and manage multiple GitHub Actions variables within your GitHub repositories.
You must have write access to a repository to use this resource.

~> Note: github_actions_variables cannot be used in conjunction with github_actions_variable or
they will fight over what your policy should be.

## Example Usage

```hcl
resource "github_actions_variables" "repo_vars" {
  repository = "example_repository"

  variable {
    name  = "first_variable"
    value = "first_value"
  }

  variable {
    name  = "second_variable"
    value = "second_value"
  }
}
```

## Argument Reference

The following arguments are supported:

* `repository`   - (Required) Name of the repository.
* `variable`     - (Optional) Set of variables to manage. Limited to a maximum of 500 variables per repository. Each variable block supports the following:
  * `name`       - (Required) Name of the variable.
  * `value`      - (Required) Value of the variable.

## Attributes Reference

In addition to the arguments above, each variable block exports the following read-only attributes:

* `created_at`   - Date of variable creation.
* `updated_at`   - Date of variable update.

## Import

This resource can be imported using the repository name:

```
$ terraform import github_actions_variables.vars repository
