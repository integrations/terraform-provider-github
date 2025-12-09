---
page_title: "GitHub: github_enterprise_actions_permissions"
description: |-
  Creates and manages Actions permissions within a GitHub enterprise
---

# github_enterprise_actions_permissions

This resource allows you to create and manage GitHub Actions permissions within your GitHub enterprise. You must have admin access to an enterprise to use this resource.

## Example Usage

```terraform
data "github_organization" "example-org" {
  name = "my-org"
}

resource "github_enterprise_actions_permissions" "test" {
  enterprise_slug = "my-enterprise"
  allowed_actions = "selected"
  enabled_organizations = "selected"
  allowed_actions_config {
    github_owned_allowed = true 
    patterns_allowed     = ["actions/cache@*", "actions/checkout@*"]
    verified_allowed     = true
  }
  enabled_organizations_config {
    organization_ids = [data.github_organization.example-org.id]
  }
}
```

## Argument Reference

The following arguments are supported:

- `enterprise_slug` - (Required) The slug of the enterprise.
- `allowed_actions` - (Optional) The permissions policy that controls the actions that are allowed to run. Can be one of: `all`, `local_only`, or `selected`.
- `enabled_organizations` - (Required) The policy that controls the organizations in the enterprise that are allowed to run GitHub Actions. Can be one of: `all`, `none`, or `selected`.
- `allowed_actions_config` - (Optional) Sets the actions that are allowed in an enterprise. Only available when `allowed_actions` = `selected`. See [Allowed Actions Config](#allowed-actions-config) below for details.
- `enabled_organizations_config` - (Optional) Sets the list of selected organizations that are enabled for GitHub Actions in an enterprise. Only available when `enabled_organizations` = `selected`. See [Enabled Organizations Config](#enabled-organizations-config) below for details.

### Allowed Actions Config

The `allowed_actions_config` block supports the following:

- `github_owned_allowed` - (Required) Whether GitHub-owned actions are allowed in the organization.
- `patterns_allowed` - (Optional) Specifies a list of string-matching patterns to allow specific action(s). Wildcards, tags, and SHAs are allowed. For example, `monalisa/octocat@*`, `monalisa/octocat@v2`, `monalisa/*`.
- `verified_allowed` - (Optional) Whether actions in GitHub Marketplace from verified creators are allowed. Set to `true` to allow all GitHub Marketplace actions by verified creators.

### Enabled Organizations Config

The `enabled_organizations_config` block supports the following:

- `organization_ids` - (Required) List of organization IDs to enable for GitHub Actions.

## Import

This resource can be imported using the name of the GitHub enterprise:

```sh
$ terraform import github_enterprise_actions_permissions.test github_enterprise_name
```
