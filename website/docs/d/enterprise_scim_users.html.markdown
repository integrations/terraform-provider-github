---
layout: "github"
page_title: "GitHub: github_enterprise_scim_users"
description: |-
  Get SCIM users provisioned for a GitHub enterprise.
---

# github_enterprise_scim_users

Use this data source to retrieve SCIM users provisioned for a GitHub enterprise.

## Example Usage

```hcl
data "github_enterprise_scim_users" "example" {
  enterprise = "example-co"
}
```

## Argument Reference

* `enterprise` - (Required) The enterprise slug.
* `filter` - (Optional) SCIM filter string.
* `results_per_page` - (Optional) Page size used while auto-fetching all pages (mapped to SCIM `count`).

### Notes on `filter`

`filter` is passed to the GitHub SCIM API as-is (server-side filtering). It is **not** a Terraform expression and it does **not** understand provider schema paths such as `name[0].family_name`.

GitHub supports **only one** filter expression and only for these attributes on the enterprise `Users` listing endpoint:

* `userName`
* `externalId`
* `id`
* `displayName`

Examples:

```hcl
filter = "userName eq \"E012345\""
```

```hcl
filter = "externalId eq \"9138790-10932-109120392-12321\""
```

If you need to filter by other values that only exist in the Terraform schema (for example `name[0].family_name`), retrieve the users and filter locally in Terraform.

## Attributes Reference

* `schemas` - SCIM response schemas.
* `total_results` - Total number of SCIM users.
* `start_index` - Start index from the first page.
* `items_per_page` - Items per page from the first page.
* `resources` - List of SCIM users. Each entry includes:
  * `schemas`, `id`, `external_id`, `user_name`, `display_name`, `active`, `name`, `emails`, `roles`, `meta`.
