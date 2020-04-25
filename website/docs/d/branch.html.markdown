---
layout: "github"
page_title: "GitHub: github_branch"
description: |-
  Get information about a repository branch.
---

# github\_branch

Use this data source to retrieve information about a repository branch.

## Example Usage

```hcl
data "github_branch" "development" {
  repository = "example"
  branch     = "development"
}
```

## Argument Reference

The following arguments are supported:

* `repository` - (Required) The GitHub repository name.

* `branch` - (Required) The repository branch to create.

## Attribute Reference

The following additional attributes are exported:

* `etag` - An etag representing the Branch object.

* `ref` - A string representing a branch reference, in the form of `refs/heads/<branch>`.

* `sha` - A string storing the reference's `HEAD` commit's SHA1.
