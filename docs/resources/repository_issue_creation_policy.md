---
page_title: "github_repository_issue_creation_policy (Resource) - GitHub"
description: |-
  Manages the issue creation policy for a repository
---

# github_repository_issue_creation_policy (Resource)

This resource allows you to manage the issue creation policy for a repository. The policy controls who is allowed to create issues in the repository.

Destroying this resource does not delete anything on GitHub; it resets the repository's issue creation policy to `all`.

## Example Usage

```terraform
resource "github_repository" "example" {
  name       = "example-repo"
  visibility = "private"
}

resource "github_repository_issue_creation_policy" "example" {
  repository = github_repository.example.name
  policy     = "collaborators_only"
}
```

## Argument Reference

The following arguments are supported:

- `repository` - (Required) The name of the GitHub repository. Renaming the repository is supported without recreating this resource.

- `policy` - (Required) Controls who can create issues in the repository. Can be `all` or `collaborators_only`.

## Attribute Reference

In addition to the above arguments, the following attributes are exported:

- `repository_id` - The numeric ID of the GitHub repository.

## Import

The issue creation policy can be imported using the repository name.

```shell
terraform import github_repository_issue_creation_policy.example my-repo
```
