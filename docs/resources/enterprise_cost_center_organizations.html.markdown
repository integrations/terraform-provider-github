---
layout: "github"
page_title: "GitHub: github_enterprise_cost_center_organizations"
description: |-
  Manages organization assignments to a GitHub enterprise cost center.
---

# github_enterprise_cost_center_organizations

This resource manages organization assignments to a GitHub enterprise cost center.

~> **Note:** This resource is authoritative. It will manage the full set of organizations assigned to the cost center. To add organizations without affecting other assignments, you must include all desired organizations in the `organization_logins` set.

## Example Usage

```hcl
resource "github_enterprise_cost_center" "example" {
  enterprise_slug = "example-enterprise"
  name            = "platform-cost-center"
}

resource "github_enterprise_cost_center_organizations" "example" {
  enterprise_slug     = "example-enterprise"
  cost_center_id      = github_enterprise_cost_center.example.id
  organization_logins = ["octo-org", "acme-corp"]
}
```

## Argument Reference

* `enterprise_slug` - (Required) The slug of the enterprise.
* `cost_center_id` - (Required) The ID of the cost center.
* `organization_logins` - (Required) Set of organization logins to assign to the cost center. Must contain at least one organization.

## Attributes Reference

This resource exports no additional attributes.

## Import

GitHub Enterprise Cost Center Organization assignments can be imported using the `enterprise_slug` and the `cost_center_id`, separated by a `:` character.

```
$ terraform import github_enterprise_cost_center_organizations.example example-enterprise:<cost_center_id>
```
