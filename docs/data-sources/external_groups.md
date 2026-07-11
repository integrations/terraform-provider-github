---
page_title: "github_external_groups (Data Source) - GitHub"
description: |-
  Retrieve external groups belonging to an organization.
---

# github_external_groups (Data Source)

Use this data source to retrieve external groups belonging to an organization.

## Example Usage

### All external groups

```terraform
data "github_external_groups" "example_external_groups" {}

locals {
  local_groups = data.github_external_groups.example_external_groups
}

output "groups" {
  value = local.local_groups
}
```

### Filtered by display name

```terraform
data "github_external_groups" "example_external_groups_filtered" {
  display_name = "my-group"
}

locals {
  filtered_groups = data.github_external_groups.example_external_groups_filtered
}

output "groups" {
  value = local.filtered_groups
}
```

## Argument Reference

- `display_name` - (Optional) Filter the list of external groups by display name. Only groups whose name contains this value will be returned.

## Attributes Reference

- `external_groups` - an array of external groups belonging to the organization. Each group consists of the fields documented below.

---

- `group_id` - the ID of the group.
- `group_name` - the name of the group.
- `updated_at` - the date the group was last updated.
