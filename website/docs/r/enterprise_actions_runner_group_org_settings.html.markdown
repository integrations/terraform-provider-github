---
layout: "github"
page_title: "GitHub: github_enterprise_actions_runner_group_org_settings"
description: |-
  Manages repository access for an enterprise Actions Runner Group at the organization level.
---

# github_enterprise_actions_runner_group_org_settings

This resource allows you to manage repository access for **enterprise** Actions runner groups that are inherited by an organization.
When an enterprise runner group is shared with an organization (via `selected_organization_ids` in `github_enterprise_actions_runner_group`),
this resource allows you to configure which repositories within that organization can use the runner group.

**Important:** This resource is specifically for managing inherited enterprise runner groups. It will not work with organization-level runner groups created directly in the organization. For organization-level runner groups, use the `github_actions_runner_group` resource instead.

You must have admin access to the organization to use this resource.

## Example Usage

### Basic Usage

```hcl
data "github_enterprise" "enterprise" {
  slug = "my-enterprise"
}

data "github_organization" "org" {
  name = "my-organization"
}

# Create a repository
resource "github_repository" "example" {
  name        = "example-repo"
  description = "Example repository"
  visibility  = "private"
}

# Create an enterprise runner group and share it with the organization
resource "github_enterprise_actions_runner_group" "example" {
  name                      = "my-runner-group"
  enterprise_slug           = data.github_enterprise.enterprise.slug
  visibility                = "selected"
  selected_organization_ids = [data.github_organization.org.id]
}

# Configure repository access for the runner group in the organization
resource "github_enterprise_actions_runner_group_org_settings" "example" {
  organization                 = data.github_organization.org.name
  enterprise_runner_group_name = github_enterprise_actions_runner_group.example.name
  selected_repository_ids      = [github_repository.example.repo_id]
  allows_public_repositories   = true
  restricted_to_workflows      = true
  selected_workflows           = ["${github_repository.example.full_name}/.github/workflows/ci.yml@refs/heads/main"]
}
```

### Multiple Repositories

```hcl
data "github_enterprise" "enterprise" {
  slug = "my-enterprise"
}

data "github_organization" "org" {
  name = "my-organization"
}

resource "github_repository" "repo1" {
  name       = "repo-1"
  visibility = "private"
}

resource "github_repository" "repo2" {
  name       = "repo-2"
  visibility = "private"
}

resource "github_enterprise_actions_runner_group" "example" {
  name                      = "my-runner-group"
  enterprise_slug           = data.github_enterprise.enterprise.slug
  visibility                = "selected"
  selected_organization_ids = [data.github_organization.org.id]
}

resource "github_enterprise_actions_runner_group_org_settings" "example" {
  organization                 = data.github_organization.org.name
  enterprise_runner_group_name = github_enterprise_actions_runner_group.example.name
  selected_repository_ids = [
    github_repository.repo1.repo_id,
    github_repository.repo2.repo_id,
  ]
}
```

### All Repositories (visibility = "all")

```hcl
data "github_enterprise" "enterprise" {
  slug = "my-enterprise"
}

data "github_organization" "org" {
  name = "my-organization"
}

resource "github_enterprise_actions_runner_group" "example" {
  name                      = "my-runner-group"
  enterprise_slug           = data.github_enterprise.enterprise.slug
  visibility                = "selected"
  selected_organization_ids = [data.github_organization.org.id]
}

# Make the runner group available to all repositories in the organization
resource "github_enterprise_actions_runner_group_org_settings" "example" {
  organization                 = data.github_organization.org.name
  enterprise_runner_group_name = github_enterprise_actions_runner_group.example.name
  visibility                   = "all"
}
```

## Argument Reference

The following arguments are supported:

* `organization`                 - (Required) The GitHub organization name.
* `enterprise_runner_group_name` - (Required) The name of the enterprise runner group inherited by the organization.
* `visibility`                   - (Optional) The visibility of the runner group. Can be `all`, `selected`, or `private`. Defaults to `selected`.
* `selected_repository_ids`      - (Optional) List of repository IDs that can access the runner group. Required when `visibility` is set to `selected`.
* `allows_public_repositories`   - (Optional) Whether public repositories can be added to the runner group. Defaults to `false`.
* `restricted_to_workflows`      - (Optional) If `true`, the runner group will be restricted to running only the workflows specified in the `selected_workflows` array. Defaults to `false`.
* `selected_workflows`           - (Optional) List of workflows the runner group should be allowed to run. This setting will be ignored unless `restricted_to_workflows` is set to `true`. The format is `{repo_full_name}/.github/workflows/{workflow_file}@{ref}` (e.g., `my-org/my-repo/.github/workflows/ci.yml@refs/heads/main`).

## Attributes Reference

The following additional attributes are exported:

* `id`              - The ID of the resource in the format `organization:runner_group_id`.
* `runner_group_id` - The ID of the inherited enterprise runner group in the organization.
* `inherited`       - Whether this runner group is inherited from the enterprise (always `true` for this resource).

## Import

This resource can be imported using the organization name and either the runner group ID or name:

```
# Import using runner group ID
$ terraform import github_enterprise_actions_runner_group_org_settings.example my-organization:123

# Import using runner group name
$ terraform import github_enterprise_actions_runner_group_org_settings.example my-organization:my-runner-group
```

## Notes

* This resource **only** manages inherited enterprise runner groups. It will automatically verify that the runner group is inherited from the enterprise.
* The runner group must already exist at the enterprise level and be shared with the organization (via `selected_organization_ids` in `github_enterprise_actions_runner_group`).
* For organization-level runner groups (not inherited from enterprise), use the `github_actions_runner_group` resource instead.
* When this resource is destroyed, the runner group visibility is reset to `all`, making it available to all repositories in the organization.
* The runner group itself is not deleted when this resource is destroyed - only the repository access configuration is reset.
