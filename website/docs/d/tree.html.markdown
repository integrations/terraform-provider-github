---
layout: "github"
page_title: "GitHub: github_tree"
description: |-
  Returns a single tree using the SHA1 value for that tree.
---

# github_tree

Use this data source to retrieve information about a single tree.

## Example Usage

```hcl
data "github_repository" "this" {
  name = "example"
}

data "github_branch" "this" {
  branch     = data.github_repository.this.default_branch
  repository = data.github_repository.this.name
}

data "github_tree" "this" {
  recursive  = false
  repository = data.github_repository.this.name
  tree_sha   = data.github_branch.this.sha
}

output "entries" {
  value = data.github_tree.this.entries
}

```

## Argument Reference

- `recursive` - (Optional) Setting this parameter to `true` returns the objects or subtrees referenced by the tree specified in `tree_sha`.
- `repository` - (Required) The name of the repository.
- `tree_sha` - (Required) The SHA1 value for the tree.

## Attributes Reference

- `entries` - Objects (of `path`, `mode`, `type`, `size`, and `sha`) specifying a tree structure.
