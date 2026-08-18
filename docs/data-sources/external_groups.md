---
page_title: "github_external_groups (Data Source) - GitHub"
subcategory: ""
description: |-
  Data source to list all external groups in an organization.
---

# github_external_groups (Data Source)

Data source to list all external groups in an organization.

## Example Usage

```terraform
data "github_external_groups" "example" {}
```

<!--
## Schema

### Optional

- `display_name_filter` (String) Filter external groups by display name.

### Read-Only

- `external_groups` (List of Object) List of external groups in the organization. (see [below for nested schema](#nestedatt--external_groups))
- `id` (String) The ID of this resource.

<a id="nestedatt--external_groups"></a>
### Nested Schema for `external_groups`

Read-Only:

- `group_id` (Number)
- `group_name` (String)
- `updated_at` (String)
-->

## Schema

### Optional

- `display_name_filter` (String) Filter external groups by display name.

### Read-Only

- `external_groups` (List of Object) List of external groups in the organization. (see [below for nested schema](#nestedatt--external_groups))
- `id` (String) The ID of this resource.

<a id="nestedatt--external_groups"></a>
### Nested Schema for `external_groups`

Read-Only:

- `group_id` (Number) ID of the external group.
- `group_name` (String) Name of the external group.
- `updated_at` (String) Timestamp of the last update to the external group.
