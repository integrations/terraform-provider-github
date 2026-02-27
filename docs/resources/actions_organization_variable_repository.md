---
page_title: "GitHub: github_actions_organization_variable_repository"
description: |-
  Add access for a repository to an Actions Variable within a GitHub organization.
---

# github_actions_organization_variable_repository

This resource adds permission for a repository to use an actions variables within your GitHub organization. You must have write access to an organization variable to use this resource.

This resource is only applicable when `visibility` of the existing organization variable has been set to `selected`.

## Example Usage

```terraform
resource "github_actions_organization_variable" "example" {
	variable_name     = "myvariable"
	plaintext_value = "foo"
	visibility      = "selected"
}

resource "github_repository" "example" {
	name       = "myrepo"
	visibility = "public"
}

resource "github_actions_organization_variable_repository" "example" {
  variable_name   = github_actions_organization_variable.example.name
  repository_id = github_repository.example.repo_id
}
```

## Argument Reference

The following arguments are supported:

- `variable_name` - (Required) Name of the actions organization variable.
- `repository_id` - (Required) ID of the repository that should be able to access the variable.

## Import

This resource can be imported using an ID made of the variable name and repository name separated by a `:`.

### Import Block

The following import block imports the access of repository ID `123456` for the actions organization variable named `myvariable` to a `github_actions_organization_variable_repository` resource named `example`.

```terraform
import {
  to = github_actions_organization_variable_repository.example
  id = "myvariable:123456"
}
```

### Import Command

The following command imports the access of repository ID `123456` for the actions organization variable named `myvariable` to a `github_actions_organization_variable_repository` resource named `example`.

```shell
terraform import github_actions_organization_variable_repository.example myvariable:123456
```
