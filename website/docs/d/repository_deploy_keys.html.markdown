---
layout: "github"
page_title: "GitHub: repository_deploy_keys"
description: |-
  Get all deploy keys of a repository
---

# github_repository_deploy_keys

Use this data source to retrieve all deploy keys of a repository.

## Example Usage

```hcl
data "github_repository_deploy_keys" "example" {
    repository = "example-repository"
}
```

## Argument Reference

* `repository` - (Required) Name of the repository to retrieve the branches from.

## Attributes Reference

* `keys` - The list of this repository's deploy keys. Each element of `keys` has the following attributes:
    * `id` - Key id
    * `title` - Key title
    * `key` - Key itself
    * `verified` - `true` if the key was verified.
