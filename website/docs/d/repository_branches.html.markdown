---
layout: "github"
page_title: "GitHub: repository_branches"
description: |-
  Get information on a Github repository's branches.
---

# github_repository_branches

Use this data source to retrieve information about branches in a repository.

## Example Usage

```hcl
data "github_repository_branches" "example" {
    repository = "example-repository"
}
```

## Argument Reference

* `repository` - (Required) Name of the repository to retrieve the branches from.

## Attributes Reference

* `branches` - The list of this repository's branches. Each element of `branches` has the following attributes:
    * `name` - Name of the branch.
    * `protected` - Whether the branch is protected.
