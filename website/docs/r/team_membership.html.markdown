---
layout: "github"
page_title: "GitHub: github_team_membership"
description: |-
  Provides a GitHub team membership resource.
---

# github_team_membership

Provides a GitHub team membership resource.

This resource allows you to add/remove users from teams in your organization. When applied,
the user will be added to the team. If the user hasn't accepted their invitation to the
organization, they won't be part of the team until they do. When
destroyed, the user will be removed from the team.

## Example Usage

```hcl
# Add a data source to obtain user details for the user
data "github_user" "some_user" {
  username = "SomeUser"
}

# Add a user to the organization
resource "github_membership" "membership_for_some_user" {
  user_id  = data.github_user.some_user.id
  role     = "member"
}

resource "github_team" "some_team" {
  name        = "SomeTeam"
  description = "Some cool team"
}

resource "github_team_membership" "some_team_membership" {
  team_id  = github_team.some_team.id
  user_id  = data.github_user.some_user.id
  role     = "member"
}
```

## Argument Reference

The following arguments are supported:

* `team_id` - (Required) The GitHub team ID
* `user_id` - (Required) The GitHub user ID to add to the team.
* `role` - (Optional) The role of the user within the team.
Must be one of `member` or `maintainer`. Defaults to `member`.

## Attribute Reference

The following attributes are exported:

* `username` - The username (login) of the user specified by user_id

## Import

GitHub Team Membership can be imported using an ID made up of two parts; a
team identifier and a user identifier. The team can be specified using its
'slug' (short name) or numeric ID, and the user can be specified using its
name (login) or numeric ID.

```
$ terraform import github_team_membership.member 1234567:1

$ terraform import github_team_membership.member 1234567:octocat

$ terraform import github_team_membership.member some-team:1

$ terraform import github_team_membership.member some-team:octocat
```
