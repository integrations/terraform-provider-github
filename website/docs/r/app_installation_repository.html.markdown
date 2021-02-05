---
layout: "github"
page_title: "GitHub: github_app_installation_repository"
description: |-
  Manages the associations between app installations and repositories.
---

# github_app_installation_repository

This resource manages relationships between app installations and repositories
in your GitHub organization.

Creating this resource installs a particular app on a particular repository.

The app installation and the repository must both belong to the same
organization on GitHub.

## Example Usage

```hcl
# Create a repository
resource "github_repository" "some_repo" {
  name = "some-repo"
}

resource "github_app_installation_repository" "some_app_repo" {
  # The installation id of the app.
  installation_id    = "1234567"
  repository         = "${github_repository.some_repo.name}"
}
```

## Argument Reference

The following arguments are supported:

* `installation_id` - (Required) The GitHub app installation id.
* `repository`      - (Required) The repository to install the app on

## Import

GitHub App Installation Repository can be imported
using an ID made up of `installation_id:repository`, e.g.

```
$ terraform import github_app_installation_repository.terraform_repo 1234567:terraform
```
