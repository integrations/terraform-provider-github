---
page_title: "GitHub: github_organization_security_manager"
description: |-
  Manages the Security manager teams for a GitHub Organization.
---

# github_organization_security_manager

~> **Note:** This resource is deprecated, please use the `github_organization_role_team` resource instead.

## Example Usage

```terraform
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

```text
$ terraform import github_organization_security_manager.core 1234567
```
