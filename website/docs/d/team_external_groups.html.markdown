---
layout: "github"
page_title: "GitHub: github_team_external_groups"
description: |-
  Retrieve external groups for a specific GitHub team.
---

# github\_team\_external\_groups

Use this data source to retrieve external groups for a specific GitHub team.

## Example Usage

```hcl
data "github_team_external_groups" "example" {
  slug = "example"
}

output "groups" {
  value = data.github_team_external_groups.example.external_groups
}
```

## Argument Reference

* `slug` - (Required) The slug of the GitHub team.

## Attributes Reference

* `external_groups` - An array of external groups for the team. Each group consists of the fields documented below.
  * `group_id` - The ID of the external group.
  * `group_name` - The name of the external group.
  * `updated_at` - The date the group was last updated.
