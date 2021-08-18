---
layout: "github"
page_title: "GitHub: github_actions_organization_secret_repositories"
description: |-
  Manages repository allow list for an Action Secret within a GitHub organization
---

# github_actions_organization_secret_repositories

This resource allows you to manage repository allow list for existing GitHub Actions secrets within your GitHub organization.
You must have write access to an organization secret to use this resource.

This resource is only applicable when `visibility` of the existing organization secret has been set to `selected`.

## Example Usage

```hcl
data "github_repository" "repo" {
  full_name = "my-org/repo"
}

resource "github_actions_organization_secret_repositories" "org_secret_repos" {
  secret_name = "existing_secret_name"
  selected_repository_ids = [data.github_repository.repo.repo_id]
}
```

## Argument Reference

The following arguments are supported:

* `secret_name`             - (Required) Name of the existing secret
* `selected_repository_ids` - (Required) An array of repository ids that can access the organization secret.

## Import

This resource can be imported using an ID made up of the secret name:

```
$ terraform import github_actions_organization_secret_repositories.test_secret_repos test_secret_name
```
