---
layout: "github"
page_title: "Github: github_enterprise_scim_group"
description: |-
  Get SCIM provisioning information for a GitHub enterprise group.
---

# github_enterprise_scim_group

Use this data source to retrieve SCIM provisioning information for a single enterprise group.

## Example Usage

```
data "github_enterprise_scim_group" "example" {
  enterprise     = "example-co"
  scim_group_id  = "123456"
}
```

## Argument Reference

* `enterprise` - (Required) The enterprise slug.
* `scim_group_id` - (Required) The SCIM group ID.
* `excluded_attributes` - (Optional) SCIM `excludedAttributes` query parameter.

## Attributes Reference

* `schemas`, `id`, `external_id`, `display_name`, `members`, `meta`.
