---
layout: "github"
page_title: "GitHub: github_enterprise_team_membership"
description: |-
  Checks if a user is a member of a GitHub enterprise team.
---

# github_enterprise_team_membership

Use this data source to check whether a user belongs to an enterprise team.

~> **Note:** Requires GitHub Enterprise Cloud with a classic PAT that has enterprise admin scope.

## Example Usage

```hcl
data "github_enterprise_team_membership" "example" {
  enterprise_slug = "my-enterprise"
  team_slug       = "ent:platform"
  username        = "octocat"
}
```

## Argument Reference

The following arguments are supported:

* `enterprise_slug` - (Required) The slug of the enterprise.
* `team_slug` - (Required) The slug of the enterprise team.
* `username` - (Required) The GitHub username.

## Attributes Reference

The following additional attributes are exported:

* `user_id` - The ID of the user.
