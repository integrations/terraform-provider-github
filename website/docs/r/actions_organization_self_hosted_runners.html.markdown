---
layout: "github"
page_title: "GitHub: github_actions_organization_self_hosted_runners"
description: |-
  Creates and manages self-hosted runners settings within a GitHub organization
---

# github_actions_organization_self_hosted_runners

This resource allows you to manage self-hosted runners settings within your GitHub organization.
It controls which repositories are allowed to create repository-level self-hosted runners.
You must have admin access to an organization to use this resource.

## Example Usage

```hcl
resource "github_actions_organization_self_hosted_runners" "example" {
  enabled_repositories = "all"
}
```

```hcl
resource "github_repository" "example" {
  name = "my-repository"
}

resource "github_actions_organization_self_hosted_runners" "example" {
  enabled_repositories = "selected"
  enabled_repositories_config {
    repository_ids = [github_repository.example.repo_id]
  }
}
```

## Argument Reference

The following arguments are supported:

* `enabled_repositories`        - (Required) The policy that controls which repositories in the organization can create self-hosted runners. Can be one of: `all`, `selected`, or `none`.
* `enabled_repositories_config` - (Optional) Sets the list of selected repositories that are allowed to create self-hosted runners. Only available when `enabled_repositories` = `selected`. See [Enabled Repositories Config](#enabled-repositories-config) below for details.

### Enabled Repositories Config

The `enabled_repositories_config` block supports the following:

* `repository_ids` - (Required) List of repository IDs allowed to create self-hosted runners.

## Import

This resource can be imported using the name of the GitHub organization:

```
$ terraform import github_actions_organization_self_hosted_runners.test github_organization_name
```
