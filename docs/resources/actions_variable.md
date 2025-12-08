---
layout: "github"
page_title: "GitHub: github_actions_variable"
description: |-
  Creates and manages an Action variable within a GitHub repository
---

# github_actions_variable

This resource allows you to create and manage GitHub Actions variables within your GitHub repositories.
You must have write access to a repository to use this resource.


## Example Usage

```hcl
resource "github_actions_variable" "example_variable" {
  repository       = "example_repository"
  variable_name    = "example_variable_name"
  value            = "example_variable_value"
}
```

## Argument Reference

The following arguments are supported:

* `repository`      - (Required) Name of the repository
* `variable_name`   - (Required) Name of the variable
* `value`           - (Required) Value of the variable

## Attributes Reference

* `created_at`      - Date of actions_variable creation.
* `updated_at`      - Date of actions_variable update.

## Import

GitHub Actions variables can be imported using an ID made up of `repository:variable_name`, e.g.

```
$ terraform import github_actions_variable.myvariable myrepo:myvariable
```
