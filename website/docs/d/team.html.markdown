---
layout: "github"
page_title: "GitHub: github_team"
description: |-
  Get information on a GitHub team.
---

# github\_team

Use this data source to retrieve information about a GitHub team.

## Example Usage

```hcl
data "github_team" "example" {
  slug = "example"
}
```

## Argument Reference

* `slug` - (Required) The team slug.
* `membership_type` - (Optional) Type of membership to be requested to fill the list of members. Can be either `all` _(default)_ or `immediate`.
* `summary_only` - (Optional) Exclude the members and repositories of the team from the returned result. Defaults to `false`.
* `results_per_page` - (**DEPRECATED**) (Optional) Set the number of results per REST API query. Accepts a value between 0 - 100 _(defaults to `100`)_.

## Attributes Reference

* `id` - ID of the team.
* `node_id` - Node ID of the team.
* `name` - Team's full name.
* `description` - Team's description.
* `privacy` - Team's privacy type. Can either be `closed` or `secret`.
* `notification_setting` - Teams's notification setting. Can be either `notifications_enabled` or `notifications_disabled`.
* `permission` - (**DEPRECATED**) The permission that new repositories will be added to the team with when none is specified.
* `members` - List of team members (list of GitHub usernames). Not returned if `summary_only = true`.
* `repositories` - (**DEPRECATED**) List of team repositories (list of repo names). Not returned if `summary_only = true`.
* `repositories_detailed` - List of team repositories (each item comprises of `repo_id`, `repo_name` & [`role_name`](https://registry.terraform.io/providers/integrations/github/latest/docs/resources/team_repository#permission)). Not returned if `summary_only = true`.
