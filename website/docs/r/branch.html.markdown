---
layout: "github"
page_title: "GitHub: github_branch"
description: |-
  Creates and manages branches within GitHub repositories.
---

# github\_branch

This resource allows you to create and manage branches within your repository.

Additional constraints can be applied to ensure your branch is created from
another branch or commit.

## Example Usage

```hcl
resource "github_branch" "development" {
  repository = "example"
  branch     = "development"
}
```

## Argument Reference

The following arguments are supported:

* `repository` - (Required) The GitHub repository name.

* `branch` - (Required) The repository branch to create.

* `source_branch` - (Optional) The branch name to start from. Defaults to `main`.

* `source_sha` - (Optional) The commit hash to start from. Defaults to the tip of `source_branch`. If provided, `source_branch` is ignored.

## Attribute Reference

The following additional attributes are exported:

* `source_sha` - A string storing the commit this branch was started from. Not populated when imported.

* `etag` - An etag representing the Branch object.

* `ref` - A string representing a branch reference, in the form of `refs/heads/<branch>`.

* `sha` - A string storing the reference's `HEAD` commit's SHA1.

## Import

GitHub Branch can be imported using an ID made up of `repository:branch`, e.g.

```
$ terraform import github_branch.terraform terraform:master
```

Optionally, a source branch may be specified using an ID of `repository:branch:source_branch`.
This is useful for importing branches that do not branch directly off master.

```
$ terraform import github_branch.terraform terraform:feature-branch:dev
```
