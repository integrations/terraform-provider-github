---
layout: "github"
page_title: "GitHub: github_enterprise_cost_center"
description: |-
  Create and manage a GitHub enterprise cost center.
---

# github_enterprise_cost_center

This resource allows you to create and manage a GitHub enterprise cost center.

~> **Note:** This resource manages only the cost center entity itself. To assign users, organizations, or repositories, use the separate `github_enterprise_cost_center_users`, `github_enterprise_cost_center_organizations`, and `github_enterprise_cost_center_repositories` resources.

Deleting this resource archives the cost center (GitHub calls this state `deleted`).

## Example Usage

```hcl
resource "github_enterprise_cost_center" "example" {
  enterprise_slug = "example-enterprise"
  name            = "platform-cost-center"
}

# Use separate resources to manage assignments
resource "github_enterprise_cost_center_users" "example" {
  enterprise_slug = "example-enterprise"
  cost_center_id  = github_enterprise_cost_center.example.id
  usernames       = ["alice", "bob"]
}
```

## Argument Reference

* `enterprise_slug` - (Required) The slug of the enterprise.
* `name` - (Required) The name of the cost center.

## Attributes Reference

The following additional attributes are exported:

* `id` - The cost center ID.
* `state` - The state of the cost center.
* `azure_subscription` - The Azure subscription associated with the cost center.

## Import

GitHub Enterprise Cost Center can be imported using the `enterprise_slug` and the `cost_center_id`, separated by a `:` character.

```
$ terraform import github_enterprise_cost_center.example example-enterprise:<cost_center_id>
```

