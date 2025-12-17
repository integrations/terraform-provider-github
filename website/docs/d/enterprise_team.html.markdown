---
layout: "github"
page_title: "Github: github_enterprise_team"
description: |-
  Get information about a GitHub enterprise team.
---

# github_enterprise_team

Use this data source to retrieve information about an enterprise team.

~> **Note:** Requires GitHub Enterprise Cloud with a classic PAT that has enterprise admin scope.

## Example Usage

Lookup by slug:

```hcl
data "github_enterprise_team" "example" {
  enterprise_slug = "my-enterprise"
  slug            = "ent:platform"
}
```

Lookup by numeric ID:

```hcl
data "github_enterprise_team" "example" {
  enterprise_slug = "my-enterprise"
  team_id         = 123456
}
```

## Argument Reference

The following arguments are supported:

* `enterprise_slug` - (Required) The slug of the enterprise.
* `slug` - (Optional) The slug of the enterprise team. Conflicts with `team_id`.
* `team_id` - (Optional) The numeric ID of the enterprise team. Conflicts with `slug`.

## Attributes Reference

The following additional attributes are exported:

* `name` - The name of the enterprise team.
* `description` - The description of the enterprise team.
* `organization_selection_type` - Which organizations in the enterprise should have access to this team.
* `group_id` - The ID of the IdP group to assign team membership with.
