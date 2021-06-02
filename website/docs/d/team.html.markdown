---
layout: "github"
page_title: "GitHub: github_team"
description: |-
  Get information on a GitHub team.
---

# github\_team

Use this data source to retrieve information about a GitHub team.

## Example Usage

```hcl
data "github_team" "example" {
  slug = "example"
}
```

## Argument Reference

 * `slug` - (Required) The team slug.

## Attributes Reference

 * `id` - the ID of the team.
 * `node_id` - the Node ID of the team.
 * `name` - the team's full name.
 * `description` - the team's description.
 * `privacy` - the team's privacy type.
 * `permission` - the team's permission level.
 * `members` - List of team members
 * `repositories` - List of team repositories
 
