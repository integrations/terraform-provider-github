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

~> **Note** This resource is not compatible with `github_team_membership`. Use either `github_team_members` or `github_team_membership`.

~> **Note** You can accidentally lock yourself out of your team using this resource. Deleting a `github_team_members` resource removes access from anyone without organization-level access to the team. Proceed with caution. It should generally only be used with teams fully managed by Terraform.

~> **Note** Attempting to set a user who is an organization owner to "member" will result in the user being granted "maintainer" instead; this can result in a perpetual `terraform plan` diff that changes their status back to "member".

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

* `team_id` - (Required) The team id or the team slug

~> **Note** Although the team id or team slug can be used it is recommended to use the team id.  Using the team slug will cause the team members associations to the team to be destroyed and recreated if the team name is updated.

* `members` - (Required) List of team members. See [Members](#members) below for details.

### Members

`members` supports the following arguments:

* `username` - (Required) The user to add to the team.
* `role` - (Optional) The role of the user within the team.
            Must be one of `member` or `maintainer`. Defaults to `member`.

## Import

~> **Note** Although the team id or team slug can be used it is recommended to use the team id.  Using the team slug will result in terraform doing conversions between the team slug and team id.  This will cause team members associations to the team to be destroyed and recreated on import.

GitHub Team Membership can be imported using the team ID team id or team slug, e.g.

```
$ terraform import github_team_members.some_team 1234567
$ terraform import github_team_members.some_team Administrators
```
