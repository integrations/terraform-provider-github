---
page_title: "github_organization_repositories (Data Source) - GitHub"
subcategory: ""
description: |-
  Data source to list all organization repositories.
---

# github_organization_repositories (Data Source)

Data source to list all organization repositories.

## Example Usage

```terraform
data "github_organization_repositories" "example" {}
```

<!--
## Schema

### Read-Only

- `id` (String) The ID of this resource.
- `repositories` (List of Object) Organization repositories. (see [below for nested schema](#nestedatt--repositories))

<a id="nestedatt--repositories"></a>
### Nested Schema for `repositories`

Read-Only:

- `archived` (Boolean)
- `id` (Number)
- `name` (String)
- `node_id` (String)
- `visibility` (String)
-->

## Schema

### Read-Only

- `id` (String) The ID of this resource.
- `repositories` (List of Object) Organization repositories. (see [below for nested schema](#nestedatt--repositories))

<a id="nestedatt--repositories"></a>
### Nested Schema for `repositories`

Read-Only:

- `archived` (Boolean) Whether the repository is archived.
- `id` (Number) ID of the repository.
- `name` (String) Name of the repository.
- `node_id` (String) Node ID of the repository.
- `visibility` (String) Visibility of the repository; one of `public`, `private`, or `internal`.
