---
layout: "github"
page_title: "Github: github_enterprise_team_organizations"
description: |-
  Manages organization assignments for a GitHub enterprise team.
---

# github_enterprise_team_organizations

This resource manages which organizations an enterprise team is assigned to. It will reconcile
the current assignments with the desired `organization_slugs`, adding and removing as needed.

~> **Note:** Requires GitHub Enterprise Cloud with a classic PAT that has enterprise admin scope.

## Example Usage

```hcl
data "github_enterprise" "enterprise" {
  slug = "my-enterprise"
}

resource "github_enterprise_team" "team" {
  enterprise_slug             = data.github_enterprise.enterprise.slug
  name                        = "Platform"
  organization_selection_type = "selected"
}

resource "github_enterprise_team_organizations" "assignments" {
  enterprise_slug = data.github_enterprise.enterprise.slug
  enterprise_team = github_enterprise_team.team.slug

  organization_slugs = [
    "my-org",
    "another-org",
  ]
}
```

## Argument Reference

The following arguments are supported:

* `enterprise_slug` - (Required) The slug of the enterprise.
* `enterprise_team` - (Required) The slug or ID of the enterprise team.
* `organization_slugs` - (Optional) Set of organization slugs to assign the team to.

## Import

This resource can be imported using:

```
$ terraform import github_enterprise_team_organizations.assignments enterprise-slug/ent:platform
```
