---
page_title: "github_enterprise_app_installations (Data Source) - GitHub"
subcategory: ""
description: |-
  Use this data source to retrieve the GitHub App installations on an enterprise-owned organization. This data source requires GitHub Enterprise Cloud or GitHub Enterprise Server 3.19+ and an authenticated user that is an enterprise owner.
---

# github_enterprise_app_installations (Data Source)

Use this data source to retrieve the GitHub App installations on an enterprise-owned organization. This data source requires GitHub Enterprise Cloud or GitHub Enterprise Server 3.19+ and an authenticated user that is an enterprise owner.

## Example Usage

```terraform
data "github_enterprise_app_installations" "example" {
  enterprise_slug = "my-enterprise"
  organization    = "my-org"
}
```

<!--
## Schema

### Required

- `enterprise_slug` (String) The slug of the enterprise that owns the organization.
- `organization` (String) The login of the enterprise-owned organization.

### Read-Only

- `id` (String) The ID of this resource.
- `installations` (List of Object) List of GitHub App installations on the organization. (see [below for nested schema](#nestedatt--installations))

<a id="nestedatt--installations"></a>
### Nested Schema for `installations`

Read-Only:

- `app_id` (Number)
- `app_slug` (String)
- `client_id` (String)
- `created_at` (String)
- `events` (List of String)
- `id` (Number)
- `permissions` (Map of String)
- `repository_selection` (String)
- `single_file_paths` (List of String)
- `suspended` (Boolean)
- `target_id` (Number)
- `target_type` (String)
- `updated_at` (String)
-->

## Schema

### Required

- `enterprise_slug` (String) The slug of the enterprise that owns the organization.
- `organization` (String) The login of the enterprise-owned organization.

### Read-Only

- `id` (String) The ID of this resource.
- `installations` (List of Object) List of GitHub App installations on the organization. (see [below for nested schema](#nestedatt--installations))

<a id="nestedatt--installations"></a>
### Nested Schema for `installations`

Read-Only:

- `app_id` (Number) The ID of the GitHub App.
- `app_slug` (String) The URL-friendly name of the GitHub App.
- `client_id` (String) The OAuth client ID of the GitHub App.
- `created_at` (String) The date the GitHub App installation was created.
- `events` (List of String) The list of events the GitHub App installation subscribes to.
- `id` (Number) The ID of the GitHub App installation.
- `permissions` (Map of String) The permissions granted to the GitHub App installation.
- `repository_selection` (String) Whether the installation has access to all repositories or only selected ones. Possible values are `all` or `selected`.
- `single_file_paths` (List of String) The list of single file paths the GitHub App installation has access to.
- `suspended` (Boolean) Whether the GitHub App installation is currently suspended.
- `target_id` (Number) The ID of the account the GitHub App is installed on.
- `target_type` (String) The type of account the GitHub App is installed on. Possible values are `Organization` or `User`.
- `updated_at` (String) The date the GitHub App installation was last updated.
