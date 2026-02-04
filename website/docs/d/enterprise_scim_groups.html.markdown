---
layout: "github"
page_title: "GitHub: github_enterprise_scim_groups"
description: |-
  Get SCIM groups provisioned for a GitHub enterprise.
---

# github_enterprise_scim_groups

Use this data source to retrieve SCIM groups provisioned for a GitHub enterprise.

## Example Usage

```hcl
data "github_enterprise_scim_groups" "example" {
  enterprise = "example-co"
}
```

## Argument Reference

* `enterprise` - (Required) The enterprise slug.
* `filter` - (Optional) SCIM filter string.
* `results_per_page` - (Optional) Page size used while auto-fetching all pages (mapped to SCIM `count`).

### Notes on `filter`

`filter` is passed to the GitHub SCIM API as-is (server-side filtering). It is **not** a Terraform expression and it does **not** understand provider schema paths.

GitHub supports **only one** filter expression and only for these attributes on the enterprise `Groups` listing endpoint:

* `externalId`
* `id`
* `displayName`

Example:

```hcl
filter = "displayName eq \"Engineering\""
```

## Attributes Reference

* `schemas` - SCIM response schemas.
* `total_results` - Total number of SCIM groups.
* `start_index` - Start index from the first page.
* `items_per_page` - Items per page from the first page.
* `resources` - List of SCIM groups. Each entry includes:
  * `schemas`, `id`, `external_id`, `display_name`, `members`, `meta`.
