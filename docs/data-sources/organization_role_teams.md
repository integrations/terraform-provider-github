---
page_title: "github_organization_role_teams (Data Source) - GitHub"
subcategory: ""
description: |-
  Data source to list all teams assigned to a custom organization role.
---

# github_organization_role_teams (Data Source)

Data source to list all teams assigned to a custom organization role.

## Example Usage

```terraform
data "github_organization_role_teams" "example" {
  role_id = 1234
}
```

<!--
## Schema

### Required

- `role_id` (Number) ID of the organization role.

### Read-Only

- `id` (String) The ID of this resource.
- `teams` (List of Object) Teams assigned to the organization role. (see [below for nested schema](#nestedatt--teams))

<a id="nestedatt--teams"></a>
### Nested Schema for `teams`

Read-Only:

- `assignment` (String)
- `description` (String)
- `id` (Number)
- `name` (String)
- `parent_team` (List of Object) (see [below for nested schema](#nestedobjatt--teams--parent_team))
- `permission` (String)
- `privacy` (String)
- `slug` (String)
- `team_id` (Number)
- `type` (String)

<a id="nestedobjatt--teams--parent_team"></a>
### Nested Schema for `teams.parent_team`

Read-Only:

- `id` (Number)
- `slug` (String)
-->

## Schema

### Required

- `role_id` (Number) ID of the organization role.

### Read-Only

- `id` (String) The ID of this resource.
- `teams` (List of Object) Teams assigned to the organization role. (see [below for nested schema](#nestedatt--teams))

<a id="nestedatt--teams"></a>
### Nested Schema for `teams`

Read-Only:

- `assignment` (String) Relationship a team has with a role; one of `direct`, `indirect`, or `mixed`.
- `description` (String) Description of the team.
- `id` (Number) ID of the team.
- `name` (String) Name of the team.
- `parent_team` (List of Object) Parent team; only set if this team is not a root team. (see [below for nested schema](#nestedobjatt--teams--parent_team))
- `permission` (String) Legacy default repository permission for the team (typically pull, push, or admin), used when adding a repository without specifying an explicit permission. This does not represent effective access for all repositories or custom repository roles.
- `privacy` (String) Privacy level of the team; one of `secret` or `closed`.
- `slug` (String) Slug of the team name.
- `team_id` (Number, Deprecated) ID of the team.
- `type` (String) Ownership type of the team; one of `enterprise` or `organization`.

<a id="nestedobjatt--teams--parent_team"></a>
### Nested Schema for `teams.parent_team`

Read-Only:

- `id` (Number) ID of the parent team.
- `slug` (String) Slug of the parent team name.
