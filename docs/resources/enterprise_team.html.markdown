---
layout: "github"
page_title: "GitHub: github_enterprise_team"
description: |-
  Creates and manages a GitHub enterprise team.
---

# github_enterprise_team

This resource allows you to create and manage a GitHub enterprise team.

~> **Note:** These API endpoints are in public preview for GitHub Enterprise Cloud and require a classic personal access token with enterprise admin permissions.

## Example Usage

```hcl
data "github_enterprise" "enterprise" {
  slug = "my-enterprise"
}

resource "github_enterprise_team" "example" {
  enterprise_slug              = data.github_enterprise.enterprise.slug
  name                         = "Platform"
  description                  = "Platform Engineering"
  organization_selection_type  = "selected"
}
```

## Argument Reference

The following arguments are supported:

* `enterprise_slug` - (Required) The slug of the enterprise.
* `name` - (Required) The name of the enterprise team.
* `description` - (Optional) A description of the enterprise team.
* `organization_selection_type` - (Optional) Which organizations in the enterprise should have access to this team. One of `disabled`, `selected`, or `all`. Defaults to `disabled`.
* `group_id` - (Optional) The ID of the IdP group to assign team membership with.

## Attributes Reference

The following additional attributes are exported:

* `id` - The numeric ID of the enterprise team.
* `team_id` - The numeric ID of the enterprise team.
* `slug` - The slug of the enterprise team (GitHub generates it and adds the `ent:` prefix).

## Import

This resource can be imported using the enterprise slug and the enterprise team numeric ID:

```
$ terraform import github_enterprise_team.example enterprise-slug/42
```
