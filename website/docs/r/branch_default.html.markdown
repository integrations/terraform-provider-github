---
layout: "github"
page_title: "GitHub: github_branch_default"
description: |-
  Provides a GitHub branch default for a given repository.
---

# github_branch_default

Provides a GitHub branch default resource.

This resource allows you to set the default branch for a given repository. 

Note that use of this resource is incompatible with the `default_branch` option of the `github_repository` resource.  Using both will result in plans always showing a diff.

## Example Usage

Basic usage:

```hcl
resource "github_repository" "example" {
  name        = "example"
  description = "My awesome codebase"
  auto_init   = true
}

resource "github_branch" "development" {
  repository = github_repository.example.name
  branch     = "development"
}

resource "github_branch_default" "default"{
  repository = github_repository.example.name
  branch     = github_branch.development.branch
}
```

Renaming to a branch that doesn't exist:

```hcl
resource "github_repository" "example" {
  name        = "example"
  description = "My awesome codebase"
  auto_init   = true
}

resource "github_branch_default" "default"{
  repository = github_repository.example.name
  branch     = "development"
  rename     = true
}
```

## Argument Reference

The following arguments are supported:

* `repository` - (Required) The GitHub repository
* `branch` - (Required) The branch (e.g. `main`)
* `rename` - (Optional) Indicate if it should rename the branch rather than use an existing branch. Defaults to `false`. 

## Import

GitHub Branch Defaults can be imported using an ID made up of `repository`, e.g.

```
$ terraform import github_branch_default.branch_default my-repo
```
