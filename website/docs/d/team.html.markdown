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
 * `membership_type` - (Optional) Type of membershp to be requested to fill the list of members. Can be either "all" or "immediate". Default: "all"

## Attributes Reference

 * `id` - the ID of the team.
 * `node_id` - the Node ID of the team.
 * `name` - the team's full name.
 * `description` - the team's description.
 * `privacy` - the team's privacy type.
 * `permission` - the team's permission level.
 * `members` - List of team members (list of GitHub usernames)
 * `repositories` - List of team repositories (list of repo names)

