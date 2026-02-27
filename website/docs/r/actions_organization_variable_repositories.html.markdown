---
layout: "github"
page_title: "GitHub: github_actions_organization_variable_repositories"
description: |-
  Manages repository allow list for an Actions Variable within a GitHub organization.
---

# github_actions_organization_variable_repositories

This resource allows you to manage the repositories allowed to access an actions variable within your GitHub organization.
You must have write access to an organization variable to use this resource.

This resource is only applicable when `visibility` of the existing organization variable has been set to `selected`.

## Example Usage

```hcl
resource "github_actions_organization_variable" "example" {
	variable_name = "myvariable"
	value         = "foo"
	visibility    = "selected"
}

resource "github_repository" "example" {
	name       = "myrepo"
	visibility = "public"
}

resource "github_actions_organization_variable_repositories" "example" {
  variable_name             = github_actions_organization_variable.example.name
  selected_repository_ids = [ github_repository.example.repo_id ]
}
```

## Argument Reference

The following arguments are supported:

- `variable_name` - (Required) Name of the actions organization variable.
- `selected_repository_ids` - (Required) List of IDs for the repositories that should be able to access the variable.

## Import

This resource can be imported using the variable name as the ID.

### Import Block

The following import block imports the repositories able to access the actions organization variable named `myvariable` to a `github_actions_organization_variable_repositories` resource named `example`.

```hcl
import {
  to = github_actions_organization_variable_repositories.example
  id = "myvariable"
}
```

### Import Command

The following command imports the repositories able to access the actions organization variable named `myvariable` to a `github_actions_organization_variable_repositories` resource named `example`.

```shell
terraform import github_actions_organization_variable_repositories.example myvariable
```
