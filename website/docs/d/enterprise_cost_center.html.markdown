---
layout: "github"
page_title: "Github: github_enterprise_cost_center"
description: |-
  Get a GitHub enterprise cost center by ID.
---

# github_enterprise_cost_center

Use this data source to retrieve a GitHub enterprise cost center by ID.

## Example Usage

```
data "github_enterprise_cost_center" "example" {
  enterprise_slug = "example-enterprise"
  cost_center_id  = "cc_123456"
}
```

## Argument Reference

* `enterprise_slug` - (Required) The slug of the enterprise.
* `cost_center_id` - (Required) The ID of the cost center.

## Attributes Reference

* `name` - The name of the cost center.
* `state` - The state of the cost center.
* `azure_subscription` - The Azure subscription associated with the cost center.
* `resources` - A list of assigned resources.
  * `type` - The resource type.
  * `name` - The resource identifier (username, organization login, or repository full name).
* `users` - The usernames currently assigned to the cost center.
* `organizations` - The organization logins currently assigned to the cost center.
* `repositories` - The repositories currently assigned to the cost center.

