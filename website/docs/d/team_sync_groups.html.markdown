---
layout: "github"
page_title: "GitHub: github_team_sync_groups"
description: |-
  Get the external identity provider (IdP) groups for a given team.
---

# github_team_sync_groups

Use this data source to retrieve the identity provider (IdP) groups for a given team.

## Example Usage

```hcl
data "github_team_sync_groups" "test" {
  retrieve_by = "slug"
  team_slug = "test-team"
}
```

## Arguments Reference

 * `retrieve_by` - (Required) The type of identifier to specify a team by (i.e. id or slug)

 * `team_id` - (Optional) The ID of the team if `retrieve_by` set to `id`.

 * `team_slug` - (Optional) The slug of the team if `retrieve_by` set to `slug`. The value may or may not differ from `name` depending on whether name contains "URL-unsafe" characters.

## Attributes Reference

 * `org_name` - The name of the team's organization.

 * `org_id` - The ID of the team's organization.

 * `group` - An Array of GitHub Identity Provider Groups.  Each `group` block consists of the fields documented below.

___

The `group` block consists of:

* `group_id` - The ID of the IdP group.

* `group_name` - The name of the IdP group. 

* `group_description` - The description of the IdP group.

