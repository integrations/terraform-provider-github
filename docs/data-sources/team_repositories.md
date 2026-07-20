---
page_title: "github_team_repositories (Data Source) - GitHub"
subcategory: ""
description: |-
  Data source to list all team repositories.
---

# github_team_repositories (Data Source)

Data source to list all team repositories.

## Example Usage

```terraform
data "github_team_repositories" "example" {
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
- `repositories` (List of Object) Team repositories. (see [below for nested schema](#nestedatt--repositories))

<a id="nestedatt--repositories"></a>
### Nested Schema for `repositories`

Read-Only:

- `archived` (Boolean)
- `id` (Number)
- `name` (String)
- `node_id` (String)
- `visibility` (String)
-->

## Schema

### Optional

- `slug` (String) Slug of the team name. One of `team_id` or `slug` must be specified.
- `team_id` (Number) ID of the team. One of `team_id` or `slug` must be specified.

### Read-Only

- `id` (String) The ID of this resource.
- `repositories` (List of Object) Team repositories. (see [below for nested schema](#nestedatt--repositories))

<a id="nestedatt--repositories"></a>
### Nested Schema for `repositories`

Read-Only:

- `archived` (Boolean) Whether the repository is archived.
- `id` (Number) ID of the repository.
- `name` (String) Name of the repository.
- `node_id` (String) Node ID of the repository.
- `visibility` (String) Visibility of the repository; one of `public`, `private`, or `internal`.
