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
resource "github_actions_environment_variable" "example_variable" {
  environment       = "example_environment"
  variable_name     = "example_variable_name"
  value             = "example_variable_value"
}
```

```hcl
data "github_repository" "repo" {
  full_name = "my-org/repo"
}

resource "github_repository_environment" "repo_environment" {
  repository       = data.github_repository.repo.name
  environment      = "example_environment"
}

resource "github_actions_environment_variable" "example_variable" {
  repository       = data.github_repository.repo.name
  environment      = github_repository_environment.repo_environment.environment
  variable_name    = "example_variable_name"
  value            = "example_variable_value"
}
```

## Argument Reference

The following arguments are supported:


* `repository`              - (Required) Name of the repository.
* `environment`             - (Required) Name of the environment.
* `variable_name`           - (Required) Name of the variable.
* `value`                   - (Required) Value of the variable

## Attributes Reference

* `created_at`      - Date of actions_environment_secret creation.
* `updated_at`      - Date of actions_environment_secret update.

## Import

This resource can be imported using an ID made up of the repository name, environment name, and variable name:

```
$ terraform import github_actions_environment_variable.test_variable myrepo:myenv:myvariable
```
