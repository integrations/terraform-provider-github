---
layout: "github"
page_title: "Github: github_enterprise_cost_center"
description: |-
  Create and manage a GitHub enterprise cost center.
---

# github_enterprise_cost_center

This resource allows you to create and manage a GitHub enterprise cost center.

Deleting this resource archives the cost center (GitHub calls this state `deleted`).

## Example Usage

```
resource "github_enterprise_cost_center" "example" {
  enterprise_slug = "example-enterprise"
  name            = "platform-cost-center"

  # Authoritatively manage assignments (Terraform will add/remove to match).
  users         = ["alice", "bob"]
  organizations = ["octo-org"]
  repositories  = ["octo-org/app"]
}
```

## Argument Reference

* `enterprise_slug` - (Required) The slug of the enterprise.
* `name` - (Required) The name of the cost center.
* `users` - (Optional) Set of usernames to assign to the cost center. Assignment is authoritative.
* `organizations` - (Optional) Set of organization logins to assign to the cost center. Assignment is authoritative.
* `repositories` - (Optional) Set of repositories (full name, e.g. `org/repo`) to assign to the cost center. Assignment is authoritative.

## Attributes Reference

The following additional attributes are exported:

* `id` - The cost center ID.
* `state` - The state of the cost center.
* `azure_subscription` - The Azure subscription associated with the cost center.
* `resources` - A list of assigned resources.
  * `type` - The resource type.
  * `name` - The resource identifier (username, organization login, or repository full name).
* `users` - The usernames currently assigned to the cost center (mirrors the authoritative input).
* `organizations` - The organization logins currently assigned to the cost center.
* `repositories` - The repositories currently assigned to the cost center.

## Import

GitHub Enterprise Cost Center can be imported using the `enterprise_slug` and the `cost_center_id`, separated by a `/` character.

```
$ terraform import github_enterprise_cost_center.example example-enterprise/<cost_center_id>
```

