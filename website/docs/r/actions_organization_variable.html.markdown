---
layout: "github"
page_title: "GitHub: github_actions_organization_variable"
description: |-
  Creates and manages an Action variable within a GitHub organization
---

# github_actions_organization_secret

This resource allows you to create and manage GitHub Actions variables within your GitHub organization.
You must have write access to a repository to use this resource.

## Example Usage

```hcl
resource "github_actions_organization_variable" "example_variable" {
  variable_name   = "example_variable_name"
  visibility      = "private"
  value           = "example_variable_value"
}
```

```hcl
data "github_repository" "repo" {
  full_name = "my-org/repo"
}

resource "github_actions_organization_variable" "example_variable" {
  secret_name             = "example_variable_name"
  visibility              = "selected"
  plaintext_value         = "example_variable_value"
  selected_repository_ids = [data.github_repository.repo.repo_id]
}
```

## Argument Reference

The following arguments are supported:

* `variable_name`           - (Required) Name of the variable
* `value`                   - (Required) Value of the variable
* `visibility`              - (Required) Configures the access that repositories have to the organization variable.
                              Must be one of `all`, `private`, `selected`. `selected_repository_ids` is required if set to `selected`.
* `selected_repository_ids` - (Optional) An array of repository ids that can access the organization variable.

## Attributes Reference

* `created_at`      - Date of actions_variable creation.
* `updated_at`      - Date of actions_variable update.

## Import

This resource can be imported using an ID made up of the variable name:

```
$ terraform import github_actions_organization_variable.test_variable test_variable_name
```
