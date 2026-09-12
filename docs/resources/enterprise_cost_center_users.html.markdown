---
layout: "github"
page_title: "GitHub: github_enterprise_cost_center_users"
description: |-
  Manages user assignments to a GitHub enterprise cost center.
---

# github_enterprise_cost_center_users

This resource manages user assignments to a GitHub enterprise cost center.

~> **Note:** This resource is authoritative. It will manage the full set of users assigned to the cost center. To add users without affecting other assignments, you must include all desired users in the `usernames` set.

## Example Usage

```hcl
resource "github_enterprise_cost_center" "example" {
  enterprise_slug = "example-enterprise"
  name            = "platform-cost-center"
}

resource "github_enterprise_cost_center_users" "example" {
  enterprise_slug = "example-enterprise"
  cost_center_id  = github_enterprise_cost_center.example.id
  usernames       = ["alice", "bob", "charlie"]
}
```

## Argument Reference

* `enterprise_slug` - (Required) The slug of the enterprise.
* `cost_center_id` - (Required) The ID of the cost center.
* `usernames` - (Required) Set of usernames to assign to the cost center. Must contain at least one username.

## Attributes Reference

This resource exports no additional attributes.

## Import

GitHub Enterprise Cost Center User assignments can be imported using the `enterprise_slug` and the `cost_center_id`, separated by a `:` character.

```
$ terraform import github_enterprise_cost_center_users.example example-enterprise:<cost_center_id>
```
