---
page_title: "github_organization_immutable_releases (Resource) - GitHub"
description: |-
  Manages immutable releases settings within a GitHub organization.
---

# github_organization_immutable_releases (Resource)

Manages immutable releases settings within a GitHub organization. You must have admin access to the organization to use this resource.

When immutable releases are enforced, release assets and metadata cannot be modified or deleted after publication.

## Example Usage

```terraform
resource "github_repository" "example" {
  name = "my-repository"
}

resource "github_organization_immutable_releases" "example" {
  enforced_repositories   = "selected"
  selected_repository_ids = [github_repository.example.repo_id]
}
```

## Schema

### Required

- `enforced_repositories` (String) The policy that controls which repositories in the organization have immutable releases enforced. Can be one of: 'all', 'none', or 'selected'.

### Optional

- `selected_repository_ids` (Set of Number) An array of repository IDs for which immutable releases enforcement should be applied. Only valid when 'enforced_repositories' is set to 'selected'.

## Import

This resource can be imported using the name of the GitHub organization.

```shell
terraform import github_organization_immutable_releases.example github_organization_name
```
