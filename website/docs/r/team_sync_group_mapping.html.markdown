---
layout: "github"
page_title: "GitHub: github_team_sync_group_mapping"
description: |-
  Creates and manages the connections between a team and its IdP group(s).
---

# github_team_sync_group_mapping

This resource allows you to create and manage Identity Provider (IdP) group connections within your GitHub teams.
You must have team synchronization enabled for organizations owned by enterprise accounts.

To learn more about team synchronization between IdPs and Github, please refer to:
https://help.github.com/en/github/setting-up-and-managing-organizations-and-teams/synchronizing-teams-between-your-identity-provider-and-github

## Example Usage

```hcl

data "github_team_sync_groups" "example_groups" {}

resource "github_team_sync_group_mapping" "example_group_mapping" {
  retrieve_by      = "slug" 
  team_slug        = "example"
  
  dynamic "group" {
    for_each = [for g in data.example_groups.groups : g if g.name == "some_team_group"]
    content {
      group_id          = each.value.group_id
      group_name        = each.value.group_name
      group_description = each.value.group_description
    }
  } 
}
```

## Argument Reference

The following arguments are supported:

* `retrieve_by`     - (Required) Name of the repository
* `team_id`         - (Optional) ID of the team  
* `team_slug`       - (Optional) Slug of the team
* `group`           - (Optional) An Array of GitHub Identity Provider Groups.  Each `group` block consists of the fields documented below.
___

The `group` block consists of:

* `group_id` - The ID of the IdP group.

* `group_name` - The name of the IdP group. 

* `group_description` - The description of the IdP group.

## Attribute Reference

In addition to the above arguments, the following attributes are exported:

 * `org_name` - The name of the team's organization.

 * `org_id` - The ID of the team's organization.
