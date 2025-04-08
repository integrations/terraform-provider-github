---
layout: "github"
page_title: "GitHub: github_organization_role_team_assignment"
description: |-
  Manages the associations between teams and organization roles.
---

# github_organization_role_team_assignment

This resource manages relationships between teams and organization roles
in your GitHub organization. This works on predefined roles, and custom roles, where the latter is an Enterprise feature.

Creating this resource assigns the role to a team.

The organization role and team must both belong to the same organization
on GitHub.

## Example Usage

```hcl
resource "github_team" "test-team" {
  name     = "test-team"
}

resource "github_organization_role_team_assignment" "test-team-role-assignment" {
  team_slug = github_team.test-team.slug
  role_id   = "8132" # all_repo_read (predefined)
}
```

## Argument Reference

The following arguments are supported:

* `team_id` - (Required) The GitHub team id or the GitHub team slug
* `role_id` - (Required) The GitHub Organization role id or role name

## Import

GitHub Team Organization Role Assignment can be imported using an ID made up of `team_id:role_id` where both name or id works, e.g.

```
$ terraform import github_organization_role_team_assignment.role_assignment 1234567:8132
$ terraform import github_organization_role_team_assignment.role_assignment test-team:8132
$ terraform import github_organization_role_team_assignment.role_assignment test-team:all_repo_read
```
