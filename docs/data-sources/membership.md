---
page_title: "github_membership (Data Source) - GitHub"
description: |-
  Get information on user membership in an organization.
---

# github_membership (Data Source)

Use this data source to find out if a user is a member of your organization, as well as what role they have within it. If the user's membership in the organization is pending their acceptance of an invite, the role they would have once they accept will be returned.

## Example Usage

```terraform
data "github_membership" "membership_for_some_user" {
  username = "SomeUser"
}
```

### Lookup by stable user ID

```terraform
# Look up a membership by the stable GitHub user ID.
# The numeric ID does not change when the user renames their account.
data "github_membership" "by_user_id" {
  user_id = 1
}
```

## Argument Reference

Exactly one of the following must be set:

- `username` - (Optional) The username (login) to lookup in the organization.
- `user_id` - (Optional) The GitHub numeric user ID. Stable across username changes; prefer this for lookups that should survive renames.

Other arguments:

- `organization` - (Optional) The organization to check for the above user.

## Attributes Reference

- `username` - The username (login). Always reflects the user's current login at refresh time.
- `user_id` - The GitHub numeric user ID.
- `role` - `admin` or `member` -- the role the user has within the organization.
- `etag` - An etag representing the membership object.
- `state` - `active` or `pending` -- the state of membership within the organization. `active` if the member has accepted the invite, or `pending` if the invite is still pending.
