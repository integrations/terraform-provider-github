---
page_title: "github_membership (Resource) - GitHub"
description: |-
  Provides a GitHub membership resource.
---

# github_membership (Resource)

Provides a GitHub membership resource.

This resource allows you to add/remove users from your organization. When applied, an invitation will be sent to the user to become part of the organization. When destroyed, either the invitation will be cancelled or the user will be removed.

Users can be addressed by their GitHub username, or invited by email address when they may not have a GitHub account yet.

## Example Usage

```terraform
# Add a user to the organization
resource "github_membership" "membership_for_some_user" {
  username = "SomeUser"
  role     = "member"
}
```

Inviting by email address:

```terraform
# Invite someone to the organization by email, e.g. when they
# do not have a GitHub account yet
resource "github_membership" "membership_for_some_email" {
  email = "someuser@example.com"
  role  = "member"
}
```

## Argument Reference

The following arguments are supported:

- `username` - (Optional) The user to add to the organization. Exactly one of `username` or `email` must be set.
- `email` - (Optional) The email address of the person to invite to the organization. The invitee does not need an existing GitHub account. Exactly one of `username` or `email` must be set. Cannot be combined with `downgrade_on_destroy = true`.
- `role` - (Optional) The role of the user within the organization. Must be one of `member` or `admin`. Defaults to `member`. `admin` role represents the `owner` role available via GitHub UI.
- `downgrade_on_destroy` - (Optional) Defaults to `false`. If set to true, when this resource is destroyed, the member will not be removed from the organization. Instead, the member's role will be downgraded to 'member'. Cannot be set to `true` when `email` is used.

~> **Note** When `email` is set, this resource manages the organization invitation: destroying the resource cancels a pending (or failed) invitation, and changing `role` cancels and re-sends the invitation. The GitHub API provides no way to look up the resulting member by invitation email, so once the invitation has been accepted the resource no longer tracks the membership. To manage the member from that point on (e.g. to change their role or remove them), replace the resource with one that uses `username` and import it.

## Import

GitHub Membership can be imported using an ID made up of `organization:username`, e.g.

```shell
terraform import github_membership.member hashicorp:someuser
```

A pending email invitation can be imported using `organization:email`, e.g.

```shell
terraform import github_membership.member hashicorp:someuser@example.com
```
