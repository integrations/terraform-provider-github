---
page_title: "github_actions_organization_permissions (Data Source) - GitHub"
description: |-
  Get GitHub Actions permissions for an organization
---

# github_actions_organization_permissions (Data Source)

Use this data source to retrieve the GitHub Actions permissions for an organization, including which actions are allowed and which repositories are enabled.

## Example Usage

```terraform
data "github_actions_organization_permissions" "example" {
}
```

## Argument Reference

No arguments are required. The organization is determined by the provider configuration.

## Attributes Reference

- `allowed_actions` - The permissions policy that controls the actions that are allowed to run. Can be one of: `all`, `local_only`, or `selected`.
- `enabled_repositories` - The policy that controls the repositories in the organization that are allowed to run GitHub Actions. Can be one of: `all`, `none`, or `selected`.
- `allowed_actions_config` - (Set when `allowed_actions` is `selected`) The actions that are allowed in the organization.
  - `github_owned_allowed` - Whether GitHub-owned actions are allowed in the organization.
  - `patterns_allowed` - Specifies a list of string-matching patterns to allow specific action(s).
  - `verified_allowed` - Whether actions in GitHub Marketplace from verified creators are allowed.
- `enabled_repositories_config` - (Set when `enabled_repositories` is `selected`) The list of selected repositories enabled for GitHub Actions.
  - `repository_ids` - List of repository IDs enabled for GitHub Actions.
- `sha_pinning_required` - Whether pinning to a specific SHA is required for all actions and reusable workflows in an organization.
