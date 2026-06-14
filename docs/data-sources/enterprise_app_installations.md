---
page_title: "github_enterprise_app_installations (Data Source) - GitHub"
description: |-
  Get the GitHub App installations of an enterprise-owned organization.
---

# github_enterprise_app_installations (Data Source)

Use this data source to retrieve the GitHub App installations on an enterprise-owned
organization.

## Example Usage

```terraform
data "github_enterprise_app_installations" "example" {
  enterprise_slug = "my-enterprise"
  organization    = "my-org"
}
```

## Argument Reference

- `enterprise_slug` - (Required) The slug of the enterprise that owns the organization.
- `organization` - (Required) The login of the enterprise-owned organization.

## Attributes Reference

- `installations` - List of GitHub App installations on the organization. Each `installation` block consists of the fields documented below.

---

The `installation` block consists of:

- `id` - The ID of the GitHub App installation.
- `app_id` - The ID of the GitHub App.
- `app_slug` - The URL-friendly name of the GitHub App.
- `client_id` - The OAuth client ID of the GitHub App.
- `target_id` - The ID of the account the GitHub App is installed on.
- `target_type` - The type of account the GitHub App is installed on.
- `repository_selection` - Whether the installation has access to `all` repositories or only `selected` ones.
- `permissions` - A map of the permissions granted to the GitHub App installation.
- `events` - The list of events the GitHub App installation subscribes to.
- `suspended` - Whether the GitHub App installation is currently suspended.
- `single_file_paths` - The list of single file paths the GitHub App installation has access to.
- `created_at` - The date the GitHub App installation was created.
- `updated_at` - The date the GitHub App installation was last updated.
