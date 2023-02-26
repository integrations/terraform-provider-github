---
layout: "github"
page_title: "GitHub: github_actions_repository_variable"
description: |-
  Creates and manages a GitHub Actions variable within a GitHub repository
---

# github_actions_repository_variable

This resource allows you to create and manage GitHub Actions variables within your GitHub repositories. You must have write access to a repository to use this resource.

GitHub Actions variables are not encrypted or secret. To store secret values, use `github_actions_secret` instead.

## Example Usage

```hcl
resource "github_repository" "repo" {
  name = "my-repo"
}

resource "github_actions_repository_variable" "username" {
  repository = github_repository.repo.name
  name       = "username"
  value      = "defunkt"
}
```

## Argument Reference

The following arguments are supported:

- `repository` - (Required) Name of the repository in which to create the variable.

- `name` - (Required) Name of the variable to create. See GitHub's [naming constraints][naming-constraints] for more information.

- `value` - (Required) Value of the variable. This value is stored in plaintext and available to anyone with access to read repository variables.


## Attributes Reference

- `created_at` - Timestamp of when the repository variable was created.

- `updated_at` - Timestamp of when the repository variable was last updated.

## Import

This resource can be imported using an ID made up of the repo name  and variable name separated by a colon:

```shell
terraform import github_actions_repository_variable.username my-repo:username
```


[naming-constraints]: https://docs.github.com/en/actions/learn-github-actions/variables#naming-conventions-for-configuration-variables
