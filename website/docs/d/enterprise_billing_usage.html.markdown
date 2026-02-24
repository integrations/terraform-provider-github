---
layout: "github"
page_title: "GitHub: github_enterprise_billing_usage"
description: |-
  Gets a billing usage report for a GitHub enterprise.
---

# github_enterprise_billing_usage

Use this data source to retrieve a billing usage report for a GitHub enterprise.
To use this data source, you must be an administrator or billing manager of the enterprise.

~> **Note:** This data source is only available to enterprises with access to the enhanced billing platform.

## Example Usage

```hcl
data "github_enterprise_billing_usage" "example" {
  enterprise_slug = "my-enterprise"
}

# Filter by a specific month
data "github_enterprise_billing_usage" "monthly" {
  enterprise_slug = "my-enterprise"
  year            = 2025
  month           = 6
}

# Filter by cost center
data "github_enterprise_billing_usage" "cost_center" {
  enterprise_slug = "my-enterprise"
  cost_center_id  = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
}
```

## Argument Reference

* `enterprise_slug` - (Required) The slug of the enterprise.
* `year` - (Optional) If specified, only return results for a single year.
* `month` - (Optional) If specified, only return results for a single month. Value between 1 and 12.
* `day` - (Optional) If specified, only return results for a single day. Value between 1 and 31.
* `cost_center_id` - (Optional) The ID corresponding to a cost center. Use `none` to target usage not associated to any cost center.

## Attributes Reference

* `usage_items` - The list of billing usage items. Each item has the following attributes:
  * `date` - The date of the usage item.
  * `product` - The product name.
  * `sku` - The SKU name.
  * `quantity` - The quantity of usage.
  * `unit_type` - The type of unit for the usage.
  * `price_per_unit` - The price per unit of usage.
  * `gross_amount` - The gross amount of usage.
  * `discount_amount` - The discount amount applied.
  * `net_amount` - The net amount after discounts.
  * `organization_name` - The organization name associated with the usage.
  * `repository_name` - The repository name associated with the usage.
