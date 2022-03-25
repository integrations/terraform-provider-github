---
layout: "github"
page_title: "GitHub: github_actions_organization_permissions"
description: |-
  Creates and manages Actions permissions within a GitHub organization
---

# github_actions_organization_permissions

This resource allows you to create and manage GitHub Actions permissions within your GitHub enterprise organizations.
You must have admin access to an organization to use this resource.

## Example Usage

```hcl
resource "github_repository" "example" {
  name = "my-repository"
}

resource "github_actions_organization_permissions" "test" {
  allowed_actions = "selected"
  enabled_repositories = "selected"
  allowed_actions_config {
    github_owned_allowed = true 
    patterns_allowed     = ["actions/cache@*", "actions/checkout@*"]
    verified_allowed     = true
  }
  enabled_repositories_config {
    repository_ids = [github_repository.example.repo_id]
  }
}
```

## Argument Reference

The following arguments are supported:

* `allowed_actions`             - (Optional) The permissions policy that controls the actions that are allowed to run. Can be one of: `all`, `local_only`, or `selected`.
* `enabled_repositories`        - (Required) The policy that controls the repositories in the organization that are allowed to run GitHub Actions. Can be one of: `all`, `none`, or `selected`.
* `allowed_actions_config`      - (Optional) Sets the actions that are allowed in an organization. Only available when `allowed_actions` = `selected`. See [Allowed Actions Config](#allowed-actions-config) below for details.
* `enabled_repositories_config` - (Optional) Sets the list of selected repositories that are enabled for GitHub Actions in an organization. Only available when `enabled_repositories` = `selected`. See [Enabled Repositories Config](#enabled-repositories-config) below for details.

### Allowed Actions Config

The `allowed_actions_config` block supports the following:

* `github_owned_allowed` - (Required) Whether GitHub-owned actions are allowed in the organization.
* `patterns_allowed` - (Optional) Specifies a list of string-matching patterns to allow specific action(s). Wildcards, tags, and SHAs are allowed. For example, monalisa/octocat@*, monalisa/octocat@v2, monalisa/*."
* `verified_allowed` - (Optional) Whether actions in GitHub Marketplace from verified creators are allowed. Set to true to allow all GitHub Marketplace actions by verified creators.

### Enabled Repositories Config

The `enabled_repositories_config` block supports the following:

* `repository_ids` - (Required) List of repository IDs to enable for GitHub Actions.

## Import

This resource can be imported using the ID of the GitHub organization:

```
$ terraform import github_actions_organization_permissions.test <github_organization_name>
```
