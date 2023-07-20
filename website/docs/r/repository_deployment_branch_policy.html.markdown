---
layout: "github"
page_title: "GitHub: github_repository_deployment_branch_policy"
description: |-
  Creates and manages deployment branch policies
---

# github_repository_deployment_branch_policy

This resource allows you to create and manage deployment branch policies.


## Example Usage

```hcl
resource "github_repository_environment" "env" {
  repository  = "my_repo"
  environment = "my_env"
  deployment_branch_policy {
    protected_branches     = false
    custom_branch_policies = true
  }
}

resource "github_repository_deployment_branch_policy" "foo" {
  repository = "my_repo"
  environment_name = "my_env"
  name = "foo"
}
```


## Argument Reference

The following arguments are supported:

* `repository` - (Required) The repository to create the policy in.

* `environment_name` - (Required) The name of the environment. This environment must have `deployment_branch_policy.custom_branch_policies` set to true.

* `name` - (Required) The name pattern that branches must match in order to deploy to the environment.

## Attributes Reference

The following additional attributes are exported:

* `id` - The ID of the deployment branch policy.

## Import

```
$ terraform import github_repository_deployment_branch_policy.foo repo:env:id
```
