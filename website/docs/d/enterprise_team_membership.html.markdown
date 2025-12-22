---
layout: "github"
page_title: "Github: github_enterprise_team_membership"
description: |-
  Check if a user is a member of a GitHub enterprise team.
---

# github_enterprise_team_membership

Use this data source to check whether a user belongs to an enterprise team.

~> **Note:** Requires GitHub Enterprise Cloud with a classic PAT that has enterprise admin scope.

## Example Usage

```hcl
data "github_enterprise_team_membership" "example" {
  enterprise_slug = "my-enterprise"
  enterprise_team = "ent:platform"
  username        = "octocat"
}
```

## Argument Reference

The following arguments are supported:

* `enterprise_slug` - (Required) The slug of the enterprise.
* `enterprise_team` - (Required) The slug or ID of the enterprise team.
* `username` - (Required) The GitHub username.

## Attributes Reference

The following additional attributes are exported:

* `role` - The membership role, if returned by the API.
* `state` - The membership state, if returned by the API.
* `etag` - The response ETag.
