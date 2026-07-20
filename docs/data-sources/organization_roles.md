---
page_title: "github_organization_roles (Data Source) - GitHub"
subcategory: ""
description: |-
  Data source to list all custom roles in an organization.
---

# github_organization_roles (Data Source)

Data source to list all custom roles in an organization.

## Example Usage

```terraform
data "github_organization_roles" "example" {
}
```

<!--
## Schema

### Read-Only

- `id` (String) The ID of this resource.
- `roles` (List of Object) Available organization roles. (see [below for nested schema](#nestedatt--roles))

<a id="nestedatt--roles"></a>
### Nested Schema for `roles`

Read-Only:

- `base_role` (String)
- `description` (String)
- `id` (Number)
- `name` (String)
- `permissions` (Set of String)
- `role_id` (Number)
- `source` (String)
-->

## Schema

### Read-Only

- `id` (String) The ID of this resource.
- `roles` (List of Object) Available organization roles. (see [below for nested schema](#nestedatt--roles))

<a id="nestedatt--roles"></a>
### Nested Schema for `roles`

Read-Only:

- `base_role` (String) System role from which this role inherits permissions.
- `description` (String) Description of the organization role.
- `id` (Number) ID of the organization role.
- `name` (String) Name of the organization role.
- `permissions` (Set of String) Additional permissions included in this role.
- `role_id` (Number, Deprecated) ID of the organization role.
- `source` (String) Source of this role; one of `Predefined`, `Organization`, or `Enterprise`.
