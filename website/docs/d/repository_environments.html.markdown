---
layout: "github"
page_title: "GitHub: repository_environments"
description: |-
  Get information on a GitHub repository's environments.
---

# github_repository_environments

Use this data source to retrieve information about environments for a repository.

> [!NOTE]
> Verify you have the correct permissions set up from the [GitHub API docs](https://docs.github.com/en/rest/deployments/environments?apiVersion=2022-11-28#get-an-environment--fine-grained-access-tokens) 

## Example Usage

```hcl
data "github_repository_environments" "example" {
    repository = "example-repository"
}
```

## Argument Reference

* `repository` - (Required) Name of the repository to retrieve the environments from.

## Attributes Reference

* `environments` - The list of this repository's environments. Each element of `environments` has the following attributes:
    * `name` - Environment name.
    * `node_id` - Environment node id.
