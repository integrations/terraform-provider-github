---
layout: "github"
page_title: "GitHub: github_organization_security_manager"
description: |-
  Manages the Security manager teams for a GitHub Organization.
---

# github_organization_security_manager

## Example Usage

```hcl
resource "github_team" "some_team" {
  name        = "SomeTeam"
  description = "Some cool team"
}

resource "github_organization_security_manager" "some_team" {
  team_slug = github_team.some_team.slug
}
```

## Argument Reference

The following arguments are supported:

* `team_slug` - (Required) The slug of the team to manage.

## Import

GitHub Security Manager Teams can be imported using the GitHub team ID e.g.

```
$ terraform import github_organization_security_manager.core 1234567
```
