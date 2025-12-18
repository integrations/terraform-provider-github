---
layout: "github"
page_title: "Github: github_enterprise_scim_user"
description: |-
  Get SCIM provisioning information for a GitHub enterprise user.
---

# github_enterprise_scim_user

Use this data source to retrieve SCIM provisioning information for a single enterprise user.

## Example Usage

```
data "github_enterprise_scim_user" "example" {
  enterprise   = "example-co"
  scim_user_id = "123456"
}
```

## Argument Reference

* `enterprise` - (Required) The enterprise slug.
* `scim_user_id` - (Required) The SCIM user ID.
* `excluded_attributes` - (Optional) SCIM `excludedAttributes` query parameter.

## Attributes Reference

* `schemas`, `id`, `external_id`, `user_name`, `display_name`, `active`, `name`, `emails`, `roles`, `meta`.
