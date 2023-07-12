---
layout: "github"
page_title: "GitHub: github_actions_default_workflow_repository_permissions"
description: |-
  Sets the default workflow permissions granted to the GITHUB_TOKEN when running workflows in a repository, and sets if GitHub Actions can submit approving pull request reviews.
---

# github_actions_default_workflow_repository_permissions

This resource allows you to manage default workflow permissions granted to GITHUB_TOKEN for a given repository.
You must have admin access to an repository to use this resource.

## Example Usage

```hcl
resource "github_repository" "example" {
  name = "my-repository"
}

resource "github_actions_default_workflow_repository_permissions" "test" {
  default_workflow_permissions     = "write"
  can_approve_pull_request_reviews = true
  repository                       = github_repository.example.name
}
```

## Argument Reference

The following arguments are supported:

* `repository`                       - (Required) The GitHub repository
* `default_workflow_permissions`     - (Optional) The default workflow permissions granted to the GITHUB_TOKEN when running workflows. Can be one of: `read`, or `write`.
* `can_approve_pull_request_reviews` - (Optional) Whether GitHub Actions can approve pull requests.

## Import

This resource can be imported using the name of the GitHub repository:

```
$ terraform import github_actions_default_workflow_repository_permissions.test my-repository
```
