---
layout: "github"
page_title: "GitHub: github_branch_default"
description: |-
  Provides a GitHub branch default for a given repository.
---

# github_branch_default

Provides a GitHub branch default resource.

This resource allows you to set the default branch for a given repository. 

## Example Usage

```hcl
# Add a collaborator to a repository
resource "github_repository_collaborator" "a_repo_collaborator" {
  repository = "our-cool-repo"
  branch     = "my-default-branch"
}
```

## Argument Reference

The following arguments are supported:

* `repository` - (Required) The GitHub repository
* `branch` - (Required) The branch (e.g. `main`)

## Import

GitHub Branch Defaults can be imported using an ID made up of `repository`, e.g.

```
$ terraform import github_branch_default.branch_default my-repo
```