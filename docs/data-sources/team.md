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

- `membership_type` (String) If `summary_only` is `false` this controls which members are returned; this can be set to either `all` or `immediate`.
- `results_per_page` (Number, Deprecated) This is unused and will be removed in a future version of the provider.
- `slug` (String) Slug of the team name. One of `team_id` or `slug` must be specified.
- `summary_only` (Boolean) If true, non-default team details such as `members` & `repositories` will be omitted.
- `team_id` (Number) ID of the team. One of `team_id` or `slug` must be specified.

### Read-Only

- `description` (String) Description of the team.
- `id` (String) The ID of this resource.
- `members` (List of String) List of members of the team.
- `name` (String) Name of the team.
- `node_id` (String) Node ID of the team.
- `notification_setting` (String) Notification setting for the team; one of `notifications_enabled`, or `notifications_disabled`.
- `parent_team` (List of Object) Parent team; only set if this team is not a root team. (see [below for nested schema](#nestedatt--parent_team))
- `permission` (String) Legacy default repository permission for the team (typically pull, push, or admin), used when adding a repository without specifying an explicit permission. This does not represent effective access for all repositories or custom repository roles.
- `privacy` (String) Privacy level of the team; one of `secret` or `closed`.
- `repositories` (List of String, Deprecated) List of repositories the team has access to.
- `repositories_detailed` (List of Object) List of repositories the team has access to. (see [below for nested schema](#nestedatt--repositories_detailed))
- `type` (String) Ownership type of the team; one of `enterprise` or `organization`.

<a id="nestedatt--parent_team"></a>
### Nested Schema for `parent_team`

Read-Only:

- `id` (Number)
- `slug` (String)


<a id="nestedatt--repositories_detailed"></a>
### Nested Schema for `repositories_detailed`

Read-Only:

- `repo_id` (Number)
- `repo_name` (String)
- `role_name` (String)
-->

## Schema

### Optional

- `membership_type` (String) If `summary_only` is `false` this controls which members are returned; this can be set to either `all` or `immediate`.
- `results_per_page` (Number, Deprecated) This is unused and will be removed in a future version of the provider.
- `slug` (String) Slug of the team name. One of `team_id` or `slug` must be specified.
- `summary_only` (Boolean) If true, non-default team details such as `members` & `repositories` will be omitted.
- `team_id` (Number) ID of the team. One of `team_id` or `slug` must be specified.

### Read-Only

- `description` (String) Description of the team.
- `id` (String) The ID of this resource.
- `members` (List of String) List of members of the team.
- `name` (String) Name of the team.
- `node_id` (String) Node ID of the team.
- `notification_setting` (String) Notification setting for the team; one of `notifications_enabled`, or `notifications_disabled`.
- `parent_team` (List of Object) Parent team; only set if this team is not a root team. (see [below for nested schema](#nestedatt--parent_team))
- `permission` (String) Legacy default repository permission for the team (typically pull, push, or admin), used when adding a repository without specifying an explicit permission. This does not represent effective access for all repositories or custom repository roles.
- `privacy` (String) Privacy level of the team; one of `secret` or `closed`.
- `repositories` (List of String, Deprecated) List of repositories the team has access to.
- `repositories_detailed` (List of Object) List of repositories the team has access to. (see [below for nested schema](#nestedatt--repositories_detailed))
- `type` (String) Ownership type of the team; one of `enterprise` or `organization`.

<a id="nestedatt--parent_team"></a>
### Nested Schema for `parent_team`

Read-Only:

- `id` (Number) ID of the parent team.
- `slug` (String) Slug of the parent team name.


<a id="nestedatt--repositories_detailed"></a>
### Nested Schema for `repositories_detailed`

Read-Only:

- `repo_id` (Number) ID of the repository.
- `repo_name` (String) Name of the repository.
- `role_name` (String) Role the team has for the repository.
