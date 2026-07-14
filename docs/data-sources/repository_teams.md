---
page_title: "github_repository_teams (Data Source) - GitHub"
subcategory: ""
description: |-
  Data source to list all teams with access to a repository.
---

# github_repository_teams (Data Source)

Data source to list all teams with access to a repository.

## Example Usage

```terraform
data "github_repository_teams" "example" {
  name = "example"
}
```

<!--
## Schema

### Optional

- `full_name` (String, Deprecated) The full name of the repository (e.g. `owner/repo`).
- `name` (String) The name of the repository.

### Read-Only

- `id` (String) The ID of this resource.
- `teams` (List of Object) Teams with access to the repository. (see [below for nested schema](#nestedatt--teams))

<a id="nestedatt--teams"></a>
### Nested Schema for `teams`

Read-Only:

- `access_source` (String)
- `description` (String)
- `id` (Number)
- `name` (String)
- `node_id` (String)
- `permission` (String)
- `privacy` (String)
- `slug` (String)
- `type` (String)
-->

## Schema

### Optional

- `full_name` (String, Deprecated) The full name of the repository (e.g. `owner/repo`).
- `name` (String) The name of the repository.

### Read-Only

- `id` (String) The ID of this resource.
- `teams` (List of Object) Teams with access to the repository. (see [below for nested schema](#nestedatt--teams))

<a id="nestedatt--teams"></a>
### Nested Schema for `teams`

Read-Only:

- `access_source` (String) Source of the team's access to the repository; one of `direct`, `organization`, or `enterprise`.
- `description` (String) Description of the team.
- `id` (Number) ID of the team.
- `name` (String) Name of the team.
- `node_id` (String) Node ID of the team.
- `permission` (String) Legacy default repository permission for the team (typically pull, push, or admin), used when adding a repository without specifying an explicit permission. This does not represent effective access for all repositories or custom repository roles.
- `privacy` (String) Privacy level of the team; one of `secret` or `closed`.
- `slug` (String) Slug of the team name.
- `type` (String) Ownership type of the team; one of `enterprise` or `organization`.
