---
page_title: "github_enterprise_app_installable_organizations (Data Source) - GitHub"
subcategory: ""
description: |-
  Use this data source to retrieve the enterprise-owned organizations that GitHub Apps can be installed on. This data source requires GitHub Enterprise Cloud or GitHub Enterprise Server 3.19+ and an authenticated user that is an enterprise owner.
---

# github_enterprise_app_installable_organizations (Data Source)

Use this data source to retrieve the enterprise-owned organizations that GitHub Apps can be installed on. This data source requires GitHub Enterprise Cloud or GitHub Enterprise Server 3.19+ and an authenticated user that is an enterprise owner.

## Example Usage

```terraform
data "github_enterprise_app_installable_organizations" "example" {
  enterprise_slug = "my-enterprise"
}
```

<!--
## Schema

### Required

- `enterprise_slug` (String) The slug of the enterprise.

### Read-Only

- `id` (String) The ID of this resource.
- `organizations` (List of Object) List of organizations in the enterprise that GitHub Apps can be installed on. (see [below for nested schema](#nestedatt--organizations))

<a id="nestedatt--organizations"></a>
### Nested Schema for `organizations`

Read-Only:

- `accessible_repositories_url` (String)
- `id` (Number)
- `login` (String)
-->

## Schema

### Required

- `enterprise_slug` (String) The slug of the enterprise.

### Read-Only

- `id` (String) The ID of this resource.
- `organizations` (List of Object) List of organizations in the enterprise that GitHub Apps can be installed on. (see [below for nested schema](#nestedatt--organizations))

<a id="nestedatt--organizations"></a>
### Nested Schema for `organizations`

Read-Only:

- `accessible_repositories_url` (String) The API URL listing the repositories that can be made accessible to a GitHub App installed on the organization.
- `id` (Number) The ID of the organization.
- `login` (String) The login of the organization.
