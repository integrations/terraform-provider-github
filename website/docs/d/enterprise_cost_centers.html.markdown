---
layout: "github"
page_title: "GitHub: github_enterprise_cost_centers"
description: |-
  List GitHub enterprise cost centers.
---

# github_enterprise_cost_centers

Use this data source to list GitHub enterprise cost centers.

## Example Usage

```
data "github_enterprise_cost_centers" "active" {
  enterprise_slug = "example-enterprise"
  state           = "active"
}
```

## Argument Reference

* `enterprise_slug` - (Required) The slug of the enterprise.
* `state` - (Optional) Filter cost centers by state. Valid values are `active` and `deleted`.

## Attributes Reference

* `cost_centers` - A set of cost centers.
  * `id` - The cost center ID.
  * `name` - The name of the cost center.
  * `state` - The state of the cost center.
  * `azure_subscription` - The Azure subscription associated with the cost center.

