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
resource "github_repository" "example" {
  name        = "example"
  description = "My awesome codebase"

  visibility = "private"

  template {
    owner = "github"
    repository = "terraform-module-template"
  }
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

## Argument Reference

The following arguments are supported:

* `repository` - (Required) The GitHub repository
* `branch` - (Required) The branch (e.g. `main`)

## Import

GitHub Branch Defaults can be imported using an ID made up of `repository`, e.g.

```
$ terraform import github_branch_default.branch_default my-repo
```
