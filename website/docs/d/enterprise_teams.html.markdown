---
layout: "github"
page_title: "Github: github_enterprise_teams"
description: |-
  List all enterprise teams in a GitHub enterprise.
---

# github_enterprise_teams

Use this data source to retrieve all enterprise teams for an enterprise.

~> **Note:** Requires GitHub Enterprise Cloud with a classic PAT that has enterprise admin scope.

## Example Usage

```hcl
data "github_enterprise_teams" "all" {
  enterprise_slug = "my-enterprise"
}

output "enterprise_team_slugs" {
  value = [for t in data.github_enterprise_teams.all.teams : t.slug]
}
```

## Argument Reference

The following arguments are supported:

* `enterprise_slug` - (Required) The slug of the enterprise.

## Attributes Reference

The following additional attributes are exported:

* `teams` - List of enterprise teams in the enterprise.

Each `teams` element exports:

* `team_id` - The numeric ID of the enterprise team.
* `slug` - The slug of the enterprise team.
* `name` - The name of the enterprise team.
* `description` - The description of the enterprise team.
* `organization_selection_type` - Which organizations in the enterprise should have access to this team.
* `group_id` - The ID of the IdP group to assign team membership with.
