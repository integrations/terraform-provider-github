---
subcategory: "Repositories"
layout: "github"
page_title: "GitHub: github_repository_immutable_releases"
description: |-
  Manages immutable releases settings for a GitHub repository.
---

# github_repository_immutable_releases

Manages whether [immutable releases](https://docs.github.com/en/code-security/concepts/supply-chain-security/immutable-releases) are enabled for a repository.

## Example Usage

```hcl
resource "github_repository_immutable_releases" "example" {
  repository = github_repository.example.name
  enabled    = true
}
```

## Argument Reference

* `repository` - (Required) The name of the repository.
* `enabled` - (Required) Whether immutable releases are enabled.

## Import

Repositories can be imported using the repository name, e.g.

```shell
terraform import github_repository_immutable_releases.example myrepo
```
