---
layout: "github"
page_title: "GitHub: github_actions_organization_secret_repository"
description: |-
  Adds/remove a repository to an organization secret when the visibility for repository access is set to selected.
---

# github_actions_organization_secret_repository

This resource help you to allow/unallow a repository to use an existing GitHub Actions secrets within your GitHub organization.
You must have write access to an organization secret to use this resource.

This resource is only applicable when `visibility` of the existing organization secret has been set to `selected`.

## Example Usage

```hcl
data "github_repository" "repo" {
  full_name = "my-org/repo"
}

resource "github_actions_organization_secret_repository" "org_secret_repos" {
  secret_name = "EXAMPLE_SECRET_NAME"
  repository_id = github_repository.repo.repo_id
}
```

## Argument Reference

The following arguments are supported:

* `secret_name`   - (Required) Name of the existing secret
* `repository_id` - (Required) Repository id that can access the organization secret.

## Import

This resource can be imported using an ID made up of the secret name:

```
$ terraform import github_actions_organization_secret_repository.test_secret_repos test_secret_name:repo_id
```
