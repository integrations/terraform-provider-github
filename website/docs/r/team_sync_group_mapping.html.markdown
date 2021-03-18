---
layout: "github"
page_title: "GitHub: github_team_sync_group_mapping"
description: |-
  Creates and manages the connections between a team and its IdP group(s).
---

# github_team_sync_group_mapping

This resource allows you to create and manage Identity Provider (IdP) group connections within your GitHub teams.
You must have team synchronization enabled for organizations owned by enterprise accounts.

To learn more about team synchronization between IdPs and GitHub, please refer to:
https://help.github.com/en/github/setting-up-and-managing-organizations-and-teams/synchronizing-teams-between-your-identity-provider-and-github

## Example Usage

```hcl

data "github_organization_team_sync_groups" "example_groups" {}

resource "github_team_sync_group_mapping" "example_group_mapping" {
  team_slug        = "example"

  dynamic "group" {
    for_each = [for g in data.github_organization_team_sync_groups.example_groups.groups : g if g.group_name == "some_team_group"]
    content {
      group_id          = group.value.group_id
      group_name        = group.value.group_name
      group_description = group.value.group_description
    }
  }
}
```

## Argument Reference

The following arguments are supported:

* `team_slug`       - (Required) Slug of the team
* `group`           - (Required) An Array of GitHub Identity Provider Groups (or empty []).  Each `group` block consists of the fields documented below.
___

The `group` block consists of:

* `group_id` - The ID of the IdP group.

* `group_name` - The name of the IdP group.

* `group_description` - The description of the IdP group.

## Import

GitHub Team Sync Group Mappings can be imported using the GitHub team `slug` e.g.

```
$ terraform import github_team_sync_group_mapping.example some_team
```
