---
page_title: "github_organization_security_managers Data Source - terraform-provider-github
description: |-
  Get the security managers for an organization.
---

# github_organization_security_managers (Data Source)

~> **Note:** This data source is deprecated, please use the `github_organization_role_team` resource instead.

Use this data source to retrieve the security managers for an organization.

## Example Usage

```terraform
data "github_organization_security_managers" "test" {}
```

## Attributes Reference

- `teams` - An list of GitHub teams. Each `team` block consists of the fields documented below.

---___

The `team` block consists of:

- `id` - Unique identifier of the team.
- `slug` - Name based identifier of the team.
- `name` - Name of the team.
- `permission` - Permission that the team will have for its repositories.
