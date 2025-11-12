---
layout: "github"
page_title: "GitHub: github_team_repository"
description: |-
  Manages the associations between teams and repositories.
---

# github_team_repository

~> Note: github_team_repository cannot be used in conjunction with github_repository_collaborators or
they will fight over what your policy should be.

This resource manages relationships between teams and repositories
in your GitHub organization.

Creating this resource grants a particular team permissions on a
particular repository.

The repository and the team must both belong to the same organization
on GitHub. This resource does not actually *create* any repositories;
to do that, see [`github_repository`](repository.html).

~> **Note on Archived Repositories**: When a repository is archived, GitHub makes it read-only, preventing team permission modifications. If you attempt to destroy resources associated with archived repositories, the provider will gracefully handle the operation by logging an informational message and removing the resource from Terraform state without attempting to modify the archived repository.

This resource is non-authoritative, for managing ALL collaborators of a repo, use github_repository_collaborators
instead.

## Example Usage

```hcl
# Add a repository to the team
resource "github_team" "some_team" {
  name        = "SomeTeam"
  description = "Some cool team"
}

resource "github_repository" "some_repo" {
  name = "some-repo"
}

resource "github_team_repository" "some_team_repo" {
  team_id    = github_team.some_team.id
  repository = github_repository.some_repo.name
  permission = "pull"
}
```

## Argument Reference

The following arguments are supported:

* `team_id` - (Required) The GitHub team id or the GitHub team slug
* `repository` - (Required) The repository to add to the team.
* `permission` - (Optional) The permissions of team members regarding the repository.
  Must be one of `pull`, `triage`, `push`, `maintain`, `admin` or the name of an existing [custom repository role](https://docs.github.com/en/enterprise-cloud@latest/organizations/managing-peoples-access-to-your-organization-with-roles/managing-custom-repository-roles-for-an-organization) within the organisation. Defaults to `pull`.


## Import

GitHub Team Repository can be imported using an ID made up of `team_id:repository` or `team_name:repository`, e.g.

```
$ terraform import github_team_repository.terraform_repo 1234567:terraform
$ terraform import github_team_repository.terraform_repo Administrators:terraform
```
