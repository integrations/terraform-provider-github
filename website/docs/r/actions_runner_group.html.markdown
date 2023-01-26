---
layout: "github"
page_title: "GitHub: github_actions_runner_group"
description: |-
  Creates and manages an Actions Runner Group within a GitHub organization
---

# github_actions_runner_group

This resource allows you to create and manage GitHub Actions runner groups within your GitHub enterprise organizations.
You must have admin access to an organization to use this resource.

## Example Usage

```hcl
resource "github_repository" "example" {
  name = "my-repository"
}

resource "github_actions_runner_group" "example" {
  name                    = github_repository.example.name
  visibility              = "selected"
  selected_repository_ids = [github_repository.example.repo_id]
}
```

## Argument Reference

The following arguments are supported:

* `name`                    - (Required) Name of the runner group
* `restricted_to_workflows` - (Optional) If true, the runner group will be restricted to running only the workflows specified in the selected_workflows array. Defaults to false.
* `selected_repository_ids` - (Optional) IDs of the repositories which should be added to the runner group
* `selected_workflows`      - (Optional) List of workflows the runner group should be allowed to run. This setting will be ignored unless restricted_to_workflows is set to true.
* `visibility`              - (Optional) Visibility of a runner group. Whether the runner group can include `all`, `selected`, or `private` repositories. A value of `private` is not currently supported due to limitations in the GitHub API.

## Attributes Reference

* `allows_public_repositories` - Whether public repositories can be added to the runner group
* `default`                    - Whether this is the default runner group
* `etag`                       - An etag representing the runner group object
* `inherited`                  - Whether the runner group is inherited from the enterprise level
* `runners_url`                - The GitHub API URL for the runner group's runners
* `selected_repository_ids`    - List of repository IDs that can access the runner group
* `selected_repositories_url`  - GitHub API URL for the runner group's repositories
* `visibility`                 - The visibility of the runner group
* `restricted_to_workflows`    - If true, the runner group will be restricted to running only the workflows specified in the selected_workflows array. Defaults to false.
* `selected_workflows`         - List of workflows the runner group should be allowed to run. This setting will be ignored unless restricted_to_workflows is set to true.

## Import

This resource can be imported using the ID of the runner group:

```
$ terraform import github_actions_runner_group.test 7
```
