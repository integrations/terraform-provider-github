---
layout: "github"
page_title: "GitHub: github_organization_role_team_assignment"
description: |-
  Manages the associations between teams and organization roles.
---

# github_organization_role_team_assignment

~> **Note:** This resource is deprecated, please use the `github_organization_role_team` resource instead.

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

* `team_slug` - (Required) The GitHub team slug
* `role_id` - (Required) The GitHub organization role id

## Import

GitHub Team Organization Role Assignment can be imported using an ID made up of `team_slug:role_id`

```text
$ terraform import github_organization_role_team_assignment.role_assignment test-team:8132
```
