---
layout: "github"
page_title: "GitHub: github_organization_invitation Resource"
description: |-
  Invite a user to a GitHub organization by email address or GitHub user ID.
---

# github_organization_invitation (Resource)

Invite a user to a GitHub organization by email address or GitHub user ID. This
resource creates a pending organization invitation, which the invitee must accept
to become a member.

~> **Note:** Once the invitation is accepted, this resource will be removed from
state on the next `terraform plan`/`apply`. To manage ongoing organization
membership after acceptance, use [`github_membership`](membership.html).

~> **Note:** The `role` attribute uses `direct_member` (not `member` as used by
`github_membership`). Make sure to use the correct value when specifying a role.

## Example Usage

### Invite by email address

```terraform
resource "github_organization_invitation" "example" {
  email = "user@example.com"
}
```

### Invite by GitHub user ID

```terraform
data "github_user" "example" {
  username = "example-user"
}

resource "github_organization_invitation" "example" {
  invitee_id = data.github_user.example.id
}
```

### Invite as admin

```terraform
resource "github_organization_invitation" "admin" {
  email = "admin@example.com"
  role  = "admin"
}
```

## Argument Reference

The following arguments are supported:

* `email` - (Optional) The email address of the person to invite to the
  organization. Exactly one of `email` or `invitee_id` must be set.

* `invitee_id` - (Optional) The GitHub user ID of the person to invite.
  Exactly one of `email` or `invitee_id` must be set.

* `role` - (Optional) The role for the new member. Must be one of `admin`,
  `direct_member`, or `billing_manager`. Defaults to `direct_member`.

## Attribute Reference

The following additional attributes are exported:

* `invitation_id` - The ID of the invitation that was created.

* `login` - The GitHub username of the invited user, if the invitee has a
  GitHub account.
