---
layout: "github"
page_title: "GitHub: github_actions_environment_variable"
description: |-
  Creates and manages an Action variable within a GitHub repository environment
---

# github_actions_environment_variable

This resource allows you to create and manage GitHub Actions variables within your GitHub repository environments.
You must have write access to a repository to use this resource.

## Example Usage

```hcl
resource "github_actions_environment_variable" "example" {
  repository    = "example-repo"
  environment   = "example-environment"
  variable_name = "example_variable_name"
  value         = "example-value"
}
```

```hcl
data "github_repository" "example" {
  full_name = "my-org/repo"
}

resource "github_repository_environment" "example" {
  repository       = data.github_repository.example.name
  environment      = "example_environment"
}

resource "github_actions_environment_variable" "example" {
  repository    = data.github_repository.example.name
  environment   = github_repository_environment.example.environment
  variable_name = "example_variable_name"
  value         = "example-value"
}
```

## Argument Reference

The following arguments are supported:

- `repository` - (Required) Name of the repository.
- `environment` - (Required) Name of the environment.
- `variable_name` - (Required) Name of the variable.
- `value` - (Required) Value of the variable.

## Attributes Reference

- `repository_id` - ID of the repository.
- `created_at` - Date the variable was created.
- `updated_at` - Date the variable was last updated.

## Import

This resource can be imported using an ID made of the repository name, environment name (any `:` in the environment name need to be escaped as `??`), and variable name all separated by a `:`.

### Import Block

The following import imports a GitHub actions environment variable named `myvariable` for the repo `myrepo` and environment `myenv` to a `github_actions_environment_variable` resource named `example`.

```hcl
import {
  to = github_actions_environment_variable.example
  id = "myrepo:myenv:myvariable"
}
```

### Import Command

The following command imports a GitHub actions environment variable named `myvariable` for the repo `myrepo` and environment `myenv` to a `github_actions_environment_variable` resource named `example`.

```shell
terraform import github_actions_environment_variable.example myrepo:myenv:myvariable
```
