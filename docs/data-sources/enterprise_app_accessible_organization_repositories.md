---
page_title: "github_enterprise_app_accessible_organization_repositories (Data Source) - GitHub"
subcategory: ""
description: |-
  Use this data source to retrieve the repositories of an enterprise-owned organization that GitHub Apps can be granted access to. This data source requires GitHub Enterprise Cloud or GitHub Enterprise Server 3.19+ and an authenticated user that is an enterprise owner.
---

# github_enterprise_app_accessible_organization_repositories (Data Source)

Use this data source to retrieve the repositories of an enterprise-owned organization that GitHub Apps can be granted access to. This data source requires GitHub Enterprise Cloud or GitHub Enterprise Server 3.19+ and an authenticated user that is an enterprise owner.

## Example Usage

```terraform
data "github_enterprise_app_accessible_organization_repositories" "example" {
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
- `repositories` (List of Object) List of repositories of the organization that GitHub Apps can be granted access to. (see [below for nested schema](#nestedatt--repositories))

<a id="nestedatt--repositories"></a>
### Nested Schema for `repositories`

Read-Only:

- `full_name` (String)
- `id` (Number)
- `name` (String)
-->

## Schema

### Required

- `enterprise_slug` (String) The slug of the enterprise that owns the organization.
- `organization` (String) The login of the enterprise-owned organization.

### Read-Only

- `id` (String) The ID of this resource.
- `repositories` (List of Object) List of repositories of the organization that GitHub Apps can be granted access to. (see [below for nested schema](#nestedatt--repositories))

<a id="nestedatt--repositories"></a>
### Nested Schema for `repositories`

Read-Only:

- `full_name` (String) The full name of the repository, in the format `<organization>/<name>`.
- `id` (Number) The ID of the repository.
- `name` (String) The name of the repository.
