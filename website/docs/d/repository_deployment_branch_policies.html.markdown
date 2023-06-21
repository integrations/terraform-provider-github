---
layout: "github"
page_title: "GitHub: github_repository_deployment_branch_policies"
description: |-
  Get the list of deployment branch policies for a given repo / env.
---

# github_repository_deployment_branch_policies

Use this data source to retrieve deployment branch policies for a repository / environment.

## Example Usage

```hcl
data "github_repository_deployment_branch_policies" "example" {
    repository = "example-repository"
    environment_name = "env_name"
}
```

## Argument Reference

* `repository` - (Required) Name of the repository to retrieve the deployment branch policies from.

* `environment_name` - (Required) Name of the environment to retrieve the deployment branch policies  from.

## Attributes Reference

* `deployment_branch_policies` - The list of this repository / environment deployment policies. Each element of `deployment_branch_policies` has the following attributes:
    * `id` - Id of the policy.
    * `name` - The name pattern that branches must match in order to deploy to the environment.
