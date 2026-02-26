---
layout: "github"
page_title: "GitHub: github_organization_immutable_releases"
description: |-
  Creates and manages immutable releases settings within a GitHub organization
---

# github_organization_immutable_releases

This resource allows you to create and manage immutable releases settings within your GitHub organization.
You must have admin access to an organization to use this resource.

When immutable releases are enforced, release assets and metadata cannot be modified or deleted after publication.

## Example Usage

### Enforce immutable releases for all repositories

```hcl
resource "github_organization_immutable_releases" "example" {
  enforced_repositories = "all"
}
```

### Enforce immutable releases for selected repositories

```hcl
resource "github_repository" "example" {
  name = "my-repository"
}

resource "github_organization_immutable_releases" "example" {
  enforced_repositories  = "selected"
  selected_repository_ids = [github_repository.example.repo_id]
}
```

### Disable immutable releases enforcement

```hcl
resource "github_organization_immutable_releases" "example" {
  enforced_repositories = "none"
}
```

## Argument Reference

The following arguments are supported:

* `enforced_repositories` - (Required) The policy that controls which repositories in the organization have immutable releases enforced. Can be one of: `all`, `none`, or `selected`.
* `selected_repository_ids` - (Optional) An array of repository IDs for which immutable releases enforcement should be applied. Only valid when `enforced_repositories` = `selected`.

## Import

This resource can be imported using the name of the GitHub organization:

```
$ terraform import github_organization_immutable_releases.example github_organization_name
```
