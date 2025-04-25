---
layout: "github"
page_title: "GitHub: github_actions_environment_variables"
description: |-
  Creates and manages multiple Action variables within a GitHub repository environment
---

# github_actions_environment_variables

This resource allows you to create and manage multiple GitHub Actions variables within your GitHub repository environments.
You must have write access to a repository to use this resource.

~> Note: github_actions_environment_variables cannot be used in conjunction with github_actions_environment_variable or
they will fight over what your policy should be.

## Example Usage

```hcl
data "github_repository" "repo" {
  full_name = "my-org/repo"
}

resource "github_repository_environment" "repo_environment" {
  repository       = data.github_repository.repo.name
  environment      = "example_environment"
}

resource "github_actions_environment_variables" "environment_vars" {
  repository  = data.github_repository.repo.name
  environment = github_repository_environment.repo_environment.environment

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
* `environment`  - (Required) Name of the environment.
* `variable`     - (Optional) Set of variables to manage. Limited to a maximum of 100 variables per environment. Each variable block supports the following:
  * `name`       - (Required) Name of the variable.
  * `value`      - (Required) Value of the variable.

## Attributes Reference

In addition to the arguments above, each variable block exports the following read-only attributes:

* `created_at`   - Date of variable creation.
* `updated_at`   - Date of variable update.

## Import

This resource can be imported using an ID made up of the repository and environment name:

```
$ terraform import github_actions_environment_variables.vars repository:environment
