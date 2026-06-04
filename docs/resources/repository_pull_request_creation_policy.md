---
page_title: "github_repository_pull_request_creation_policy (Resource) - GitHub"
description: |-
  Manages the pull request creation policy for a repository
---

# github_repository_pull_request_creation_policy (Resource)

This resource allows you to manage the pull request creation policy for a repository. The policy controls who is allowed to create pull requests in the repository.

Destroying this resource does not delete anything on GitHub; it resets the repository's pull request creation policy to `all`.

## Example Usage

```terraform
resource "github_repository" "example" {
  name       = "example-repo"
  visibility = "private"
}

resource "github_repository_pull_request_creation_policy" "example" {
  repository = github_repository.example.name
  policy     = "collaborators_only"
}
```

## Argument Reference

The following arguments are supported:

- `repository` - (Required) The name of the GitHub repository. Changing this forces a new resource to be created.

- `policy` - (Required) Controls who can create pull requests in the repository. Can be `all` or `collaborators_only`.

## Import

The pull request creation policy can be imported using the repository name.

```shell
terraform import github_repository_pull_request_creation_policy.example my-repo
```
