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
data "github_branch" "master" {
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

* `ref` - A string representing a GitHub reference, in the form of `refs/heads/<branch>`.

* `sha` - A string storing the reference starting hash.
  _Note: This is not populated when imported_.
