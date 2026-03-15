---
layout: "github"
page_title: "GitHub: github_actions_organization_variable"
description: |-
  Creates and manages an Action variable within a GitHub organization
---

# github_actions_organization_variable

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
  variable_name           = "example_variable_name"
  visibility              = "selected"
  value                   = "example_variable_value"
  selected_repository_ids = [data.github_repository.repo.repo_id]
}
```

## Argument Reference

The following arguments are supported:

- `variable_name` - (Required) Name of the variable.
- `value` - (Required) Value of the variable.
- `visibility` - (Required) Configures the access that repositories have to the organization variable; must be one of `all`, `private`, or `selected`.
- `selected_repository_ids` - (Optional) An array of repository IDs that can access the organization variable; this requires `visibility` to be set to `selected`.

## Attributes Reference

- `created_at` - Date the variable was created.
- `updated_at` - Date the variable was last updated.

## Import

This resource can be imported using the variable name as the ID.

### Import Block

The following import imports a GitHub actions organization variable named `myvariable`to a `github_actions_organization_variable` resource named `example`.

```hcl
import {
  to = github_actions_organization_variable.example
  id = "myvariable"
}
```

### Import Command

The following command imports a GitHub actions organization variable named `myvariable` to a `github_actions_organization_variable` resource named `example`.

```shell
terraform import github_actions_organization_variable.example myvariable
```
