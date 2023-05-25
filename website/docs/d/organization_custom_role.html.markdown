---
layout: "github"
page_title: "GitHub: github_organization_custom_role"
description: |-
  Get a custom role from a GitHub Organization for use in repositories.
---

# github\_organization\_custom\_role

Use this data source to retrieve information about a custom role in a GitHub Organization.

~> Note: Custom roles are currently only available in GitHub Enterprise Cloud.

## Example Usage

```hcl
data "github_organization_custom_role" "example" {
  name = "example"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the custom role.
* `description` - (Optional) The description for the custom role.
* `base_role` - (Required) The system role from which the role inherits permissions. Can be one of: `read`, `triage`, `write`, or `maintain`.
* `permissions` - (Required) A list of additional permissions included in this role. Must have a minimum of 1 additional permission. The list of available permissions can be found using the [list repository fine-grained permissions for an organization](https://docs.github.com/en/enterprise-cloud@latest/rest/orgs/custom-roles?apiVersion=2022-11-28#list-repository-fine-grained-permissions-for-an-organization) API.

## Attributes Reference

The following additional attributes are exported:

* `id` - The ID of the custom role.
* `description` - The description for the custom role.
* `base_role` - (The system role from which the role inherits permissions.
* `permissions` - A list of additional permissions included in this role.