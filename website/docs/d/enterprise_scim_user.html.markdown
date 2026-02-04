---
layout: "github"
page_title: "GitHub: github_enterprise_scim_user"
description: |-
  Get SCIM provisioning information for a GitHub enterprise user.
---

# github_enterprise_scim_user

Use this data source to retrieve SCIM provisioning information for a single enterprise user.

## Example Usage

```hcl
data "github_enterprise_scim_user" "example" {
  enterprise   = "example-co"
  scim_user_id = "123456"
}
```

## Argument Reference

* `enterprise` - (Required) The enterprise slug.
* `scim_user_id` - (Required) The SCIM user ID.

## Attributes Reference

* `schemas`, `id`, `external_id`, `user_name`, `display_name`, `active`, `name`, `emails`, `roles`, `meta`.
