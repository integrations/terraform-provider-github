---
page_title: "github_organization_role_users (Data Source) - GitHub"
subcategory: ""
description: |-
  Data source to list all users assigned to a custom organization role.
---

# github_organization_role_users (Data Source)

Data source to list all users assigned to a custom organization role.

## Example Usage

```terraform
data "github_organization_role_users" "example" {
  role_id = 1234
}
```

<!--
## Schema

### Required

- `role_id` (Number) ID of the organization role.

### Read-Only

- `id` (String) The ID of this resource.
- `users` (List of Object) Users assigned to the organization role. (see [below for nested schema](#nestedatt--users))

<a id="nestedatt--users"></a>
### Nested Schema for `users`

Read-Only:

- `assignment` (String)
- `login` (String)
- `user_id` (Number)
-->

## Schema

### Required

- `role_id` (Number) ID of the organization role.

### Read-Only

- `id` (String) The ID of this resource.
- `users` (List of Object) Users assigned to the organization role. (see [below for nested schema](#nestedatt--users))

<a id="nestedatt--users"></a>
### Nested Schema for `users`

Read-Only:

- `assignment` (String) Relationship a user has with a role; one of `direct`, `indirect`, or `mixed`.
- `login` (String) Login of the user.
- `user_id` (Number) ID of the user.
