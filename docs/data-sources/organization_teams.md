---
page_title: "github_organization_teams (Data Source) - GitHub"
subcategory: ""
description: |-
  Data source to list all organization teams.
---

# github_organization_teams (Data Source)

Data source to list all organization teams.

## Example Usage

```terraform
# All teams

data "github_organization_teams" "example" {}
```

```terraform
# Only find teams without a parent team (root teams only)

data "github_organization_teams" "example" {
  root_teams_only = true
}
```

<!--
## Schema

### Optional

- `results_per_page` (Number, Deprecated) This is unused and will be removed in a future version of the provider.
- `root_teams_only` (Boolean) If true, only root teams (teams without a parent) will be returned.
- `summary_only` (Boolean) If true, non-default team details such as `members` & `repositories` will be omitted.

### Read-Only

- `id` (String) The ID of this resource.
- `teams` (List of Object) Organization teams. (see [below for nested schema](#nestedatt--teams))

<a id="nestedatt--teams"></a>
### Nested Schema for `teams`

Read-Only:

- `description` (String)
- `id` (Number)
- `members` (List of String)
- `name` (String)
- `node_id` (String)
- `notification_setting` (String)
- `parent` (Map of String)
- `parent_team` (List of Object) (see [below for nested schema](#nestedobjatt--teams--parent_team))
- `parent_team_id` (String)
- `parent_team_slug` (String)
- `permission` (String)
- `privacy` (String)
- `repositories` (List of String)
- `slug` (String)
- `type` (String)

<a id="nestedobjatt--teams--parent_team"></a>
### Nested Schema for `teams.parent_team`

Read-Only:

- `id` (Number)
- `slug` (String)
-->

## Schema

### Optional

- `results_per_page` (Number, Deprecated) This is unused and will be removed in a future version of the provider.
- `root_teams_only` (Boolean) If true, only root teams (teams without a parent) will be returned.
- `summary_only` (Boolean) If true, non-default team details such as `members` & `repositories` will be omitted.

### Read-Only

- `id` (String) The ID of this resource.
- `teams` (List of Object) Organization teams. (see [below for nested schema](#nestedatt--teams))

<a id="nestedatt--teams"></a>
### Nested Schema for `teams`

Read-Only:

- `description` (String) Description of the team.
- `id` (Number) ID of the team.
- `members` (List of String, Deprecated) List of members in the team.
- `name` (String) Name of the team.
- `node_id` (String) Node ID of the team.
- `notification_setting` (String) Notification setting for the team; one of `notifications_enabled`, or `notifications_disabled`.
- `parent` (Map of String, Deprecated) Map of parent team attributes; only set if this team is not a root team.
- `parent_team` (List of Object) Parent team; only set if this team is not a root team. (see [below for nested schema](#nestedobjatt--teams--parent_team))
- `parent_team_id` (String, Deprecated) ID of the parent team; only set if this team is not a root team.
- `parent_team_slug` (String, Deprecated) Slug of the parent team; only set if this team is not a root team.
- `permission` (String) Legacy default repository permission for the team (typically pull, push, or admin), used when adding a repository without specifying an explicit permission. This does not represent effective access for all repositories or custom repository roles.
- `privacy` (String) Privacy level of the team; one of `secret` or `closed`.
- `repositories` (List of String, Deprecated) List of repositories the team has access to.
- `slug` (String) Slug of the team name.
- `type` (String) Ownership type of the team; one of `enterprise` or `organization`.

<a id="nestedobjatt--teams--parent_team"></a>
### Nested Schema for `teams.parent_team`

Read-Only:

- `id` (Number) ID of the parent team.
- `slug` (String) Slug of the parent team name.
