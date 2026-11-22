---
page_title: "github_team_members (Data Source) - GitHub"
subcategory: ""
description: |-
  Data source to list all team members.
---

# github_team_members (Data Source)

Data source to list all team members.

## Example Usage

```terraform
data "github_team_members" "example" {
  slug = "example"
}
```

<!--
## Schema

### Optional

- `slug` (String) Slug of the team name. One of `team_id` or `slug` must be specified.
- `team_id` (Number) ID of the team. One of `team_id` or `slug` must be specified.

### Read-Only

- `id` (String) The ID of this resource.
- `members` (List of Object) Team members. (see [below for nested schema](#nestedatt--members))

<a id="nestedatt--members"></a>
### Nested Schema for `members`

Read-Only:

- `email` (String)
- `id` (Number)
- `inherited` (Boolean)
- `login` (String)
- `node_id` (String)
- `role` (String)
-->

## Schema

### Optional

- `slug` (String) Slug of the team name. One of `team_id` or `slug` must be specified.
- `team_id` (Number) ID of the team. One of `team_id` or `slug` must be specified.

### Read-Only

- `id` (String) The ID of this resource.
- `members` (List of Object) Team members. (see [below for nested schema](#nestedatt--members))

<a id="nestedatt--members"></a>
### Nested Schema for `members`

Read-Only:

- `email` (String) Email of the member.
- `id` (Number) ID of the member.
- `inherited` (Boolean) Whether the member is inherited from a parent team.
- `login` (String) Login of the member.
- `node_id` (String) Node ID of the member.
- `role` (String) Role of the member in the team; can be one of `member` or `maintainer`.
