---
layout: "github"
page_title: "GitHub: github_enterprise_billing_usage_summary"
description: |-
  Gets a billing usage summary for a GitHub enterprise.
---

# github_enterprise_billing_usage_summary

Use this data source to retrieve a billing usage summary for a GitHub enterprise.
To use this data source, you must be an administrator or billing manager of the enterprise.

~> **Note:** This endpoint is in public preview and is subject to change.

~> **Note:** Only data from the past 24 months is accessible via this data source.

## Example Usage

```hcl
data "github_enterprise_billing_usage_summary" "example" {
  enterprise_slug = "my-enterprise"
}

# Filter by a specific product
data "github_enterprise_billing_usage_summary" "actions" {
  enterprise_slug = "my-enterprise"
  year            = 2025
  product         = "Actions"
}

# Filter by organization and repository
data "github_enterprise_billing_usage_summary" "repo" {
  enterprise_slug = "my-enterprise"
  organization    = "my-org"
  repository      = "my-org/my-repo"
}
```

## Argument Reference

* `enterprise_slug` - (Required) The slug of the enterprise.
* `year` - (Optional) If specified, only return results for a single year.
* `month` - (Optional) If specified, only return results for a single month. Value between 1 and 12.
* `day` - (Optional) If specified, only return results for a single day. Value between 1 and 31.
* `organization` - (Optional) The organization name to query usage for.
* `repository` - (Optional) The repository name to query for usage in the format `owner/repository`.
* `product` - (Optional) The product name to query usage for.
* `sku` - (Optional) The SKU to query for usage.
* `cost_center_id` - (Optional) The ID corresponding to a cost center. Use `none` to target usage not associated to any cost center.

## Attributes Reference

* `time_period` - The time period of the report.
  * `year` - The year of the time period.
  * `month` - The month of the time period.
  * `day` - The day of the time period.
* `enterprise` - The enterprise name from the report.
* `usage_items` - The list of usage summary items. Each item has the following attributes:
  * `product` - The product name.
  * `sku` - The SKU name.
  * `unit_type` - The type of unit for the usage.
  * `price_per_unit` - The price per unit of usage.
  * `gross_quantity` - The gross quantity of usage.
  * `gross_amount` - The gross amount of usage.
  * `discount_quantity` - The discount quantity applied.
  * `discount_amount` - The discount amount applied.
  * `net_quantity` - The net quantity after discounts.
  * `net_amount` - The net amount after discounts.
