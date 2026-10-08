---
page_title: "github_membership (Resource) - GitHub"
description: |-
  Provides a GitHub membership resource.
---

# github_membership (Resource)

Provides a GitHub membership resource.

This resource allows you to add/remove users from your organization. When applied, an invitation will be sent to the user to become part of the organization. When destroyed, either the invitation will be cancelled or the user will be removed.

## Example Usage

```terraform
# Add a user to the organization
resource "github_membership" "membership_for_some_user" {
  username = "SomeUser"
  role     = "member"
}
```

## Example Usage with Downgrade on Destroy

```terraform
# Downgrade a member to an outside collaborator when the resource is destroyed
resource "github_membership" "outside_collaborator_on_destroy" {
  username = "SomeUser"
  role     = "member"

  downgrade_on_destroy = true
  downgrade_to         = "outside_collaborator"
}
```

## Argument Reference

The following arguments are supported:

- `username` - (Required) The user to add to the organization.
- `role` - (Optional) The role of the user within the organization. Must be one of `member` or `admin`. Defaults to `member`. `admin` role represents the `owner` role available via GitHub UI.
- `downgrade_on_destroy` - (Optional) Defaults to `false`. Instead of removing the member from the org, you can choose to downgrade their membership when this resource is destroyed. This is useful when wanting to downgrade admins while keeping them in the organization, or to downgrade members while keeping their access to public repositories intact.
- `downgrade_to` - (Optional) The target membership state when `downgrade_on_destroy` is true. Must be one of `member` or `outside_collaborator`. Defaults to `member`.

## Import

GitHub Membership can be imported using an ID made up of `organization:username`, e.g.

```shell
terraform import github_membership.member hashicorp:someuser
```
