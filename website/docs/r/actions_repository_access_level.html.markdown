---
layout: "github"
page_title: "GitHub: github_actions_repository_access_level"
description: |-
  Manages Actions and Reusable Workflow access for a GitHub repository
---

# github_actions_repository_access_level

This resource allows you to set the access level of a non-public repositories actions and reusable workflows for use in other repositories.
You must have admin access to a repository to use this resource.

## Example Usage

```hcl
resource "github_repository" "example" {
  name       = "my-repository"
  visibility = "private"
}

resource "github_actions_repository_access_level" "test" {
  access_level = "user"
  repository   = github_repository.example.name
}
```

## Argument Reference

The following arguments are supported:

* `repository`   - (Required) The GitHub repository
* `access_level` - (Required) Where the actions or reusable workflows of the repository may be used. Possible values are `none`, `user`, `organization`, or `enterprise`. 

## Import

This resource can be imported using the name of the GitHub repository:

```
$ terraform import github_actions_repository_access_level.test <github_repository_name>
```
