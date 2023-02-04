---
layout: "github"
page_title: "Github: github_enterprise_organization"
description: |-
  Create and manages a GitHub enterprise organization.
---

# github_enterprise_organization

This resource allows you to create and manage a GitHub enterprise organization.

~> **Note** This resource cannot delete an organization. Organizations must be deleted through the GitHub UI and remove them from the state using `terraform state rm`.

## Example Usage

```
resource "github_enterprise_organization" "org" {
  enterprise_id = data.github_enterprise.enterprise.id
  name          = "some-awesome-org"
  description   = "Organization created with terraform"
  billing_email = "jon@winteriscoming.com"
  admin_logins  = [
    "jon-snow"
  ]
}
```

## Argument Reference

* `enterprise_id` - (Required) The ID of the enterprise.
* `name` - (Required) The name of the organization.
* `description` - (Optional) The description of the organization.
* `billing_email` - (Required) The billing email address.
* `admin_logins` - (Required) List of organization owner usernames.

## Attributes Reference

The following additional attributes are exported:

* `id` - The ID of the organization.

## Import

Support for importing organizations is not currently supported.
