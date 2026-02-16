---
page_title: "github_repository_environments (Data Source) - GitHub"
description: |-
  Get information on a GitHub repository's environments.
---

# github_repository_environments (Data Source)

Use this data source to retrieve information about environments for a repository.

## Example Usage

```terraform
data "github_repository_environments" "example" {
  repository = "example-repository"
}
```

## Argument Reference

- `repository` - (Required) Name of the repository to retrieve the environments from.

## Attributes Reference

- `environments` - The list of this repository's environments. Each element of `environments` has the following attributes:
  - `name` - Environment name.
  - `node_id` - Environment node id.
