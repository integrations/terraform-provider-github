---
layout: "github"
page_title: "Github: github_enterprise_team_membership"
description: |-
  Manages membership in a GitHub enterprise team.
---

# github_enterprise_team_membership

This resource manages a user's membership in an enterprise team.

~> **Note:** Requires GitHub Enterprise Cloud with a classic PAT that has enterprise admin scope.

## Example Usage

```hcl
data "github_enterprise" "enterprise" {
  slug = "my-enterprise"
}

resource "github_enterprise_team" "team" {
  enterprise_slug = data.github_enterprise.enterprise.slug
  name            = "Platform"
}

resource "github_enterprise_team_membership" "member" {
  enterprise_slug = data.github_enterprise.enterprise.slug
  enterprise_team = github_enterprise_team.team.slug
  username        = "octocat"
}
```

## Argument Reference

The following arguments are supported:

* `enterprise_slug` - (Required) The slug of the enterprise.
* `enterprise_team` - (Required) The slug or ID of the enterprise team.
* `username` - (Required) The GitHub username to manage.

## Import

This resource can be imported using:

```
$ terraform import github_enterprise_team_membership.member enterprise-slug/ent:platform/octocat
```
