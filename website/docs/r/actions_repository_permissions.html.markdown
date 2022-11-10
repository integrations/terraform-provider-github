---
layout: "github"
page_title: "GitHub: github_actions_repository_permissions"
description: |-
  Enables and manages Actions permissions for a GitHub repository
---

# github_actions_repository_permissions

This resource allows you to enable and manage GitHub Actions permissions for a given repository.
You must have admin access to an repository to use this resource.

## Example Usage

```hcl
resource "github_repository" "example" {
  name = "my-repository"
}

resource "github_actions_repository_permissions" "test" {
  allowed_actions = "selected"
  allowed_actions_config {
    github_owned_allowed = true
    patterns_allowed     = ["actions/cache@*", "actions/checkout@*"]
    verified_allowed     = true
  }
  repository = github_repository.example.name
}
```

## Argument Reference

The following arguments are supported:

* `repository`             - (Required) The GitHub repository
* `allowed_actions`        - (Optional) The permissions policy that controls the actions that are allowed to run. Can be one of: `all`, `local_only`, or `selected`.
* `enabled`                - (Optional) Should GitHub actions be enabled on this repository?
* `allowed_actions_config` - (Optional) Sets the actions that are allowed in an repository. Only available when `allowed_actions` = `selected`. See [Allowed Actions Config](#allowed-actions-config) below for details.

### Allowed Actions Config

The `allowed_actions_config` block supports the following:

* `github_owned_allowed` - (Required) Whether GitHub-owned actions are allowed in the repository.
* `patterns_allowed` - (Optional) Specifies a list of string-matching patterns to allow specific action(s). Wildcards, tags, and SHAs are allowed. For example, monalisa/octocat@*, monalisa/octocat@v2, monalisa/*."
* `verified_allowed` - (Optional) Whether actions in GitHub Marketplace from verified creators are allowed. Set to true to allow all GitHub Marketplace actions by verified creators.

## Import

This resource can be imported using the name of the GitHub repository:

```
$ terraform import github_actions_repository_permissions.test <github_repository_name>
```
