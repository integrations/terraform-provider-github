---
page_title: "github_organization_members (Data Source) - GitHub"
subcategory: ""
description: |-
  Data source to list all organization members.
---

# github_organization_members (Data Source)

Data source to list all organization members.

## Example Usage

```terraform
data "github_organization_members" "example" {}
```

<!--
## Schema

### Read-Only

- `id` (String) The ID of this resource.
- `members` (List of Object) Organization members. (see [below for nested schema](#nestedatt--members))

<a id="nestedatt--members"></a>
### Nested Schema for `members`

Read-Only:

- `id` (Number)
- `login` (String)
- `node_id` (String)
-->

## Schema

### Read-Only

- `id` (String) The ID of this resource.
- `members` (List of Object) Organization members. (see [below for nested schema](#nestedatt--members))

<a id="nestedatt--members"></a>
### Nested Schema for `members`

Read-Only:

- `id` (Number) ID of the member.
- `login` (String) Login of the member.
- `node_id` (String) Node ID of the member.
