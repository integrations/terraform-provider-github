---
layout: "github"
page_title: "GitHub: github_repository_pull_request_creation_policy"
description: |-
  Manages the pull request creation policy for a repository
---

# github_repository_pull_request_creation_policy

This resource allows you to manage the pull request creation policy for a repository. The policy controls who is allowed to create pull requests.

## Example Usage

```hcl
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

* `repository` - (Required) The name of the GitHub repository. Changing this forces a new resource.

* `policy` - (Required) Controls who can create pull requests in the repository. Supported values are `all` and `collaborators_only`.

## Import

The pull request creation policy can be imported using the repository name.

```sh
terraform import github_repository_pull_request_creation_policy.example my-repo
```
