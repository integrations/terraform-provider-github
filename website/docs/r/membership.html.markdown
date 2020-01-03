---
layout: "github"
page_title: "GitHub: github_membership"
description: |-
  Provides a GitHub membership resource.
---

# github_membership

Provides a GitHub membership resource.

This resource allows you to add/remove users from your organization. When applied,
an invitation will be sent to the user to become part of the organization. When
destroyed, either the invitation will be cancelled or the user will be removed.

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
```

## Argument Reference

The following arguments are supported:

* `user_id` - (Required) The GitHub user ID to add to the organization.
* `role` - (Optional) The role of the user within the organization.
            Must be one of `member` or `admin`. Defaults to `member`.

## Import

GitHub Membership can be imported using an ID made up of two parts; an
organization name and a user identifier. The user can be specified using its
name (login) or numeric ID.

```
$ terraform import github_membership.member org:1

$ terraform import github_membership.member org:octocat
```
