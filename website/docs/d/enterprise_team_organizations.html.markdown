---
layout: "github"
page_title: "GitHub: github_enterprise_team_organizations"
description: |-
  Get organizations assigned to a GitHub enterprise team.
---

# github_enterprise_team_organizations

Use this data source to retrieve the organizations that an enterprise team has access to.

~> **Note:** Requires GitHub Enterprise Cloud with a classic PAT that has enterprise admin scope.

## Example Usage

```hcl
data "github_enterprise_team_organizations" "example" {
  enterprise_slug = "my-enterprise"
  team_slug       = "ent:platform"
}

output "assigned_orgs" {
  value = data.github_enterprise_team_organizations.example.organization_slugs
}
```

## Argument Reference

The following arguments are supported:

* `enterprise_slug` - (Required) The slug of the enterprise.
* `team_slug` - (Required) The slug of the enterprise team.

## Attributes Reference

The following additional attributes are exported:

* `organization_slugs` - Set of organization slugs the enterprise team is assigned to.
