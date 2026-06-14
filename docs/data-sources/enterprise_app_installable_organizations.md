---
page_title: "github_enterprise_app_installable_organizations (Data Source) - GitHub"
description: |-
  Get the organizations in an enterprise that a GitHub App can be installed on.
---

# github_enterprise_app_installable_organizations (Data Source)

Use this data source to retrieve the enterprise-owned organizations that a GitHub
App can be installed on.

## Example Usage

```terraform
data "github_enterprise_app_installable_organizations" "example" {
  enterprise_slug = "my-enterprise"
}
```

## Argument Reference

- `enterprise_slug` - (Required) The slug of the enterprise.

## Attributes Reference

- `organizations` - List of organizations. Each `organization` block consists of the fields documented below.

---

The `organization` block consists of:

- `id` - The ID of the organization.
- `login` - The login (slug) of the organization.
- `accessible_repositories_url` - The URL for the repositories the app can access on the organization.
