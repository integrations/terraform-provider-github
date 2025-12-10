---
page_title: "github_organization_custom_role Data Source - terraform-provider-github
description: |-
  Get a custom role from a GitHub Organization for use in repositories.
---

# github_organization_custom_role (Data Source)

~> **Note:** This data source is deprecated, please use the `github_organization_repository_role` data source instead.

Use this data source to retrieve information about a custom role in a GitHub Organization.

~> Note: Custom roles are currently only available in GitHub Enterprise Cloud.

## Example Usage

```terraform
data "github_organization_custom_role" "example" {
  name = "example"
}
```

## Argument Reference

The following arguments are supported:

- `name` - (Required) The name of the custom role.

## Attributes Reference

The following additional attributes are exported:

- `id` - The ID of the custom role.
- `description` - The description for the custom role.
- `base_role` - The system role from which the role inherits permissions.
- `permissions` - A list of additional permissions included in this role.
