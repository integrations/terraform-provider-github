---
page_title: "github_team (Data Source) - GitHub"
subcategory: ""
description: |-
  Data source to lookup a team.
---

# github_team (Data Source)

Data source to lookup a team.

## Example Usage

```terraform
data "github_team" "example" {
  slug = "example"
}
```

<!--
## Schema

### Optional

- `lookup_child_teams` (Boolean) If `true`, child teams will be looked up and returned in the `child_teams` attribute.
- `membership_type` (String) If `summary_only` is `false` this controls which members are returned; this can be set to either `all` or `immediate`.
- `results_per_page` (Number, Deprecated) This is unused and will be removed in a future version of the provider.
- `slug` (String) Slug of the team name. One of `team_id` or `slug` must be specified.
- `summary_only` (Boolean) If `true`, non-default team details such as `members` & `repositories` will be omitted.
- `team_id` (Number) ID of the team. One of `team_id` or `slug` must be specified.

### Read-Only

- `child_teams` (List of Object) List of child teams; only set if `lookup_child_teams` is `true`. (see [below for nested schema](#nestedatt--child_teams))
- `description` (String) Description of the team.
- `id` (String) The ID of this resource.
- `members` (List of String, Deprecated) List of members of the team.
- `name` (String) Name of the team.
- `node_id` (String) Node ID of the team.
- `notification_setting` (String) Notification setting for the team; one of `notifications_enabled`, or `notifications_disabled`.
- `parent_team` (List of Object) Parent team; only set if this team is not a root team. (see [below for nested schema](#nestedatt--parent_team))
- `permission` (String) Legacy default repository permission for the team (typically pull, push, or admin), used when adding a repository without specifying an explicit permission. This does not represent effective access for all repositories or custom repository roles.
- `privacy` (String) Privacy level of the team; one of `secret` or `closed`.
- `repositories` (List of String, Deprecated) List of repositories the team has access to.
- `repositories_detailed` (List of Object, Deprecated) List of repositories the team has access to. (see [below for nested schema](#nestedatt--repositories_detailed))
- `type` (String) Ownership type of the team; one of `enterprise` or `organization`.

<a id="nestedatt--child_teams"></a>
### Nested Schema for `child_teams`

Read-Only:

- `description` (String)
- `id` (Number)
- `name` (String)
- `node_id` (String)
- `notification_setting` (String)
- `permission` (String)
- `privacy` (String)
- `slug` (String)
- `type` (String)


<a id="nestedatt--parent_team"></a>
### Nested Schema for `parent_team`

Read-Only:

- `description` (String)
- `id` (Number)
- `name` (String)
- `node_id` (String)
- `notification_setting` (String)
- `permission` (String)
- `privacy` (String)
- `slug` (String)
- `type` (String)


<a id="nestedatt--repositories_detailed"></a>
### Nested Schema for `repositories_detailed`

Read-Only:

- `repo_id` (Number)
- `repo_name` (String)
- `role_name` (String)
-->

## Schema

### Optional

- `lookup_child_teams` (Boolean) If `true`, child teams will be looked up and returned in the `child_teams` attribute.
- `membership_type` (String) If `summary_only` is `false` this controls which members are returned; this can be set to either `all` or `immediate`.
- `results_per_page` (Number, Deprecated) This is unused and will be removed in a future version of the provider.
- `slug` (String) Slug of the team name. One of `team_id` or `slug` must be specified.
- `summary_only` (Boolean) If `true`, non-default team details such as `members` & `repositories` will be omitted.
- `team_id` (Number) ID of the team. One of `team_id` or `slug` must be specified.

### Read-Only

- `child_teams` (List of Object) List of child teams; only set if `lookup_child_teams` is `true`. (see [below for nested schema](#nestedatt--child_teams))
- `description` (String) Description of the team.
- `id` (String) The ID of this resource.
- `members` (List of String, Deprecated) List of members of the team.
- `name` (String) Name of the team.
- `node_id` (String) Node ID of the team.
- `notification_setting` (String) Notification setting for the team; one of `notifications_enabled`, or `notifications_disabled`.
- `parent_team` (List of Object) Parent team; only set if this team is not a root team. (see [below for nested schema](#nestedatt--parent_team))
- `permission` (String) Legacy default repository permission for the team (typically pull, push, or admin), used when adding a repository without specifying an explicit permission. This does not represent effective access for all repositories or custom repository roles.
- `privacy` (String) Privacy level of the team; one of `secret` or `closed`.
- `repositories` (List of String, Deprecated) List of repositories the team has access to.
- `repositories_detailed` (List of Object, Deprecated) List of repositories the team has access to. (see [below for nested schema](#nestedatt--repositories_detailed))
- `type` (String) Ownership type of the team; one of `enterprise` or `organization`.

<a id="nestedatt--child_teams"></a>
### Nested Schema for `child_teams`

Read-Only:

- `description` (String) Description of the child team.
- `id` (Number) ID of the child team.
- `name` (String) Name of the child team.
- `node_id` (String) Node ID of the child team.
- `notification_setting` (String) Notification setting for the child team; one of `notifications_enabled`, or `notifications_disabled`.
- `permission` (String) Legacy default repository permission for the child team (typically pull, push, or admin), used when adding a repository without specifying an explicit permission. This does not represent effective access for all repositories or custom repository roles.
- `privacy` (String) Privacy level of the child team; one of `secret` or `closed`.
- `slug` (String) Slug of the child team name.
- `type` (String) Ownership type of the child team; one of `enterprise` or `organization`.

<a id="nestedatt--parent_team"></a>
### Nested Schema for `parent_team`

Read-Only:

- `description` (String) Description of the parent team.
- `id` (Number) ID of the parent team.
- `name` (String) Name of the parent team.
- `node_id` (String) Node ID of the parent team.
- `notification_setting` (String)
- `permission` (String) Legacy default repository permission for the parent team (typically pull, push, or admin), used when adding a repository without specifying an explicit permission. This does not represent effective access for all repositories or custom repository roles.
- `privacy` (String) Privacy level of the parent team; one of `secret` or `closed`.
- `slug` (String) Slug of the parent team name.
- `type` (String) Ownership type of the parent team; one of `enterprise` or `organization`.


<a id="nestedatt--repositories_detailed"></a>
### Nested Schema for `repositories_detailed`

Read-Only:

- `repo_id` (Number) ID of the repository.
- `repo_name` (String) Name of the repository.
- `role_name` (String) Role the team has for the repository.
