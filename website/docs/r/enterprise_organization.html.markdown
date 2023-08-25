---
layout: "github"
page_title: "Github: github_enterprise_organization"
description: |-
  Create and manages a GitHub enterprise organization.
---

# github_enterprise_organization

This resource allows you to create and manage a GitHub enterprise organization.

## Example Usage

```
resource "github_enterprise_organization" "org" {
  enterprise_id = data.github_enterprise.enterprise.id
  name          = "some-awesome-org"
  display_name  = "Some Awesome Org"
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
* `display_name` - (Optional) The display name of the organization.
* `billing_email` - (Required) The billing email address.
* `admin_logins` - (Required) List of organization owner usernames.

## Attributes Reference

The following additional attributes are exported:

* `id` - The ID of the organization.

## Import

GitHub Enterprise Organization can be imported using the `slug` of the enterprise, combined with the `orgname` of the organization, separated by a `/` character.

```
$ terraform import github_enterprise_organization.org enterp/some-awesome-org
```
