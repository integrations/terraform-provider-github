---
layout: "github"
page_title: "GitHub: github_actions_runner_group_repository_access"
description: |-
  Manages a repository's access to an Actions Runner Group within a GitHub organization
---

# github_actions_runner_group_repository_access

This resource allows you to manage repository access to GitHub Actions runner groups within your GitHub (enterprise) organizations independently for each repository.
You must have runner group admin access to an organization to use this resource.

~> **Note:** The action runners group's `visibility` must be `selected` and if also managing the runner group via terraform: `selected_repository_ids` must **not** be set.

## Example Usage

```hcl
resource "github_repository" "example" {
  name = "my-repository"
}

resource "github_actions_runner_group" "example" {
  name                    = github_repository.example.name
  visibility              = "selected"
}

resource "github_actions_runner_group_repository_access" "example" {
  runner_group_id = github_actions_runner_group.id
  repository_id   = github_repository.example.repo_id
}
```

## Argument Reference

The following arguments are supported:

* `runner_group_id`                       - (Required) Id of the runner group
* `repository_id`                         - (Required) Id of the repository to give access to the runner group


## Attributes Reference

* `id` - id of this resource, formed as `<runner_group_id>/<repository_id>`

## Import

This resource can be imported using the ID of the runner group and the repository ID:

```
$ terraform import github_actions_runner_group_repository_access.test <runner_group_id>/<repository_id>
```
