---
layout: "github"
page_title: "GitHub: github_team_repository"
description: |-
  Manages the associations between teams and repositories.
---

# github_team_repository

This resource manages relationships between teams and repositories
in your GitHub organization.

Creating this resource grants a particular team permissions on a
particular repository.

The repository and the team must both belong to the same organization
on GitHub. This resource does not actually *create* any repositories;
to do that, see [`github_repository`](repository.html).

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
  team_id    = "${github_team.some_team.id}"
  repository = "${github_repository.some_repo.name}"
  permission = "pull"
}
```

## Argument Reference

The following arguments are supported:

* `team_id` - (Required) The GitHub team id or the GitHub team slug
* `repository` - (Required) The repository to add to the team.
* `permission` - (Optional) The permissions of team members regarding the repository.
  Must be one of `pull`, `triage`, `push`, `maintain`, or `admin`. Defaults to `pull`.


## Import

GitHub Team Repository can be imported using an ID made up of `teamid:repository`, e.g.

```
$ terraform import github_team_repository.terraform_repo 1234567:terraform
```
