---
layout: "github"
page_title: "GitHub: github_enterprise_cost_center_repositories"
description: |-
  Manages repository assignments to a GitHub enterprise cost center.
---

# github_enterprise_cost_center_repositories

This resource manages repository assignments to a GitHub enterprise cost center.

~> **Note:** This resource is authoritative. It will manage the full set of repositories assigned to the cost center. To add repositories without affecting other assignments, you must include all desired repositories in the `repository_names` set.

## Example Usage

```hcl
resource "github_enterprise_cost_center" "example" {
  enterprise_slug = "example-enterprise"
  name            = "platform-cost-center"
}

resource "github_enterprise_cost_center_repositories" "example" {
  enterprise_slug  = "example-enterprise"
  cost_center_id   = github_enterprise_cost_center.example.id
  repository_names = ["octo-org/my-app", "acme-corp/backend-service"]
}
```

## Argument Reference

* `enterprise_slug` - (Required) The slug of the enterprise.
* `cost_center_id` - (Required) The ID of the cost center.
* `repository_names` - (Required) Set of repository names (in `owner/repo` format) to assign to the cost center. Must contain at least one repository.

## Attributes Reference

This resource exports no additional attributes.

## Import

GitHub Enterprise Cost Center Repository assignments can be imported using the `enterprise_slug` and the `cost_center_id`, separated by a `:` character.

```
$ terraform import github_enterprise_cost_center_repositories.example example-enterprise:<cost_center_id>
```
