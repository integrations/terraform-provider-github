---
layout: "github"
page_title: "GitHub: github_membership"
description: |-
  Get information on user membership in an organization.
---

# github_membership

Use this data source to find out if a user is a member of your organization, as well
as what role they have within it.
If the user's membership in the organization is pending their acceptance of an invite,
the role they would have once they accept will be returned.

## Example Usage

```hcl
data "github_membership" "membership_for_some_user" {
    username = "SomeUser"
}
```

## Argument Reference

 * `username` - (Required) The username to lookup in the organization.

 * `organization` - (Optional) The organization to check for the above username.

## Attributes Reference

 * `username` - The username.
 * `role` - `admin` or `member` -- the role the user has within the organization.
 * `etag` - An etag representing the membership object.
