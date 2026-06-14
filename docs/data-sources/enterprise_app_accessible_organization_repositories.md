---
page_title: "github_enterprise_app_accessible_organization_repositories (Data Source) - GitHub"
description: |-
  Get the repositories of an enterprise-owned organization that a GitHub App can be granted access to.
---

# github_enterprise_app_accessible_organization_repositories (Data Source)

Use this data source to retrieve the repositories of an enterprise-owned organization
that a GitHub App can be granted access to.

## Example Usage

```terraform
data "github_enterprise_app_accessible_organization_repositories" "example" {
  enterprise_slug = "my-enterprise"
  organization    = "my-org"
}
```

## Argument Reference

- `enterprise_slug` - (Required) The slug of the enterprise.
- `organization` - (Required) The login of the enterprise-owned organization.

## Attributes Reference

- `repositories` - List of repositories. Each `repository` block consists of the fields documented below.

---

The `repository` block consists of:

- `id` - The ID of the repository.
- `name` - The name of the repository.
- `full_name` - The full name of the repository (`org/repo`).
