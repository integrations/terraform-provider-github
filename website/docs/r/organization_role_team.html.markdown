---
layout: "github"
page_title: "GitHub: github_organization_role_team Resource"
description: |-
  Manage an association between an organization role and a team.
---

# github_organization_role_team (Resource)

Manage an association between an organization role and a team.

## Example Usage

```terraform
resource "github_organization_role_team" "example" {
  role_id   = 1234
  team_slug = "example-team"
}
```

## Example Usage Security Manager Role

```terraform
locals {
  security_manager_id = one([for x in data.github_organization_roles.all_roles.roles : x.role_id if x.name == "security_manager"][*])
}

data "github_organization_roles" "all_roles" {}

resource "github_organization_role_team" "security_managers" {
  role_id   = local.security_manager_id
  team_slug = "example-team"
}

## Schema

### Required

- `role_id` (Number) The ID of the organization role.
- `team_slug` (String) The slug of the team name.

## Import

An organization role team association can be imported using the role ID and the team slug separated by a `:`.

```shell
terraform import github_organization_role_team.example "1234:example-team"
```
