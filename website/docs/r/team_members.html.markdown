---
layout: "github"
page_title: "GitHub: github_team_members"
description: |-
  Provides an authoritative GitHub team members resource.
---

# github_team_members

Provides a GitHub team members resource.

This resource allows you to manage members of teams in your organization. It sets the requested team members for the team and removes all users not managed by Terraform.

When applied, if the user hasn't accepted their invitation to the organization, they won't be part of the team until they do.

When destroyed, all users will be removed from the team.

~> **Note**: This resource is not compatible with `github_team_membership`. Use either `github_team_members` or `github_team_membership`.

~> **Note**:Note: You can accidentally lock yourself out of your team using this resource. Deleting a `github_team_members` resource removes access from anyone without organization-level access to the team. Proceed with caution. It should generally only be used with teams fully managed by Terraform.

## Example Usage

```hcl
# Add a user to the organization
resource "github_membership" "membership_for_some_user" {
  username = "SomeUser"
  role     = "member"
}

resource "github_membership" "membership_for_another_user" {
  username = "AnotherUser"
  role     = "member"
}

resource "github_team" "some_team" {
  name        = "SomeTeam"
  description = "Some cool team"
}

resource "github_team_members" "some_team_members" {
  team_id  = github_team.some_team.id

  members {
    username = "SomeUser"
    role     = "maintainer"
  }

  members {
    username = "AnotherUser"
    role     = "member"
  }
}
```

## Argument Reference

The following arguments are supported:

* `team_id` - (Required) The GitHub team id
* `members` - (Optional) List of team members. See [Members](#members) below for details.

### Members

`members` supports the following arguments:

* `username` - (Required) The user to add to the team.
* `role` - (Optional) The role of the user within the team.
            Must be one of `member` or `maintainer`. Defaults to `member`.

## Import

GitHub Team Membership can be imported using the team ID `teamid`, e.g.

```
$ terraform import github_team_member.some_team 1234567
```
