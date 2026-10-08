---
page_title: "github_team_external_groups (Data Source) - GitHub"
subcategory: ""
description: |-
  Data source to retrieve external groups for a specific GitHub team.
---

# github_team_external_groups (Data Source)

Data source to retrieve external groups for a specific GitHub team.

## Example Usage

```terraform
data "github_team_external_groups" "example" {
  slug = "example"
}
```

<!--
## Schema

### Required

- `slug` (String) The slug of the GitHub team.

### Read-Only

- `external_groups` (List of Object) List of external groups connected to the team. (see [below for nested schema](#nestedatt--external_groups))
- `id` (String) The ID of this resource.

<a id="nestedatt--external_groups"></a>
### Nested Schema for `external_groups`

Read-Only:

- `group_id` (Number)
- `group_name` (String)
- `updated_at` (String)
-->

## Schema

### Read-Only

- `external_groups` (List of Object) List of external groups connected to the team. (see [below for nested schema](#nestedatt--external_groups))
- `id` (String) The ID of this resource.

<a id="nestedatt--external_groups"></a>
### Nested Schema for `external_groups`

Read-Only:

- `group_id` (Number) ID of the external group.
- `group_name` (String) Name of the external group.
- `updated_at` (String) Timestamp of the last update to the external group.
