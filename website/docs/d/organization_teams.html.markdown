---
layout: "github"
page_title: "GitHub: github_organization_teams"
description: |-
  Get information on all GitHub teams of an organization.
---

# github\_organization\_teams

Use this data source to retrieve information about all GitHub teams in an organization.

## Example Usage

```hcl
data "github_organization_teams" "all" {}
```

## Attributes Reference

* `teams` - An Array of GitHub Teams.  Each `team` block consists of the fields documented below.

___

The `team` block consists of:

 * `id` - the ID of the team.
 * `node_id` - the Node ID of the team.
 * `slug` - the slug of the team.
 * `name` - the team's full name.
 * `description` - the team's description.
 * `privacy` - the team's privacy type.
 * `members` - List of team members.
