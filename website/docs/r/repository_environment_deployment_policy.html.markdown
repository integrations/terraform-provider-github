---
layout: "github"
page_title: "GitHub: github_repository_environment_deployment_policy"
description: |-
  Creates and manages environment deployment branch policies for GitHub repositories
---

# github_repository_environment_deployment_policy

This resource allows you to create and manage environment deployment branch policies for a GitHub repository.

## Example Usage

```hcl
data "github_user" "current" {
  username = ""
}

resource "github_repository" "test" {
  name      = "tf-acc-test-%s"
}

resource "github_repository_environment" "test" {
  repository 	= github_repository.test.name
  environment	= "environment/test"
  wait_timer	= 10000
  reviewers {
    users = [data.github_user.current.id]
  }
  deployment_branch_policy {
    protected_branches     = false
    custom_branch_policies = true
  }
}

resource "github_repository_environment_deployment_policy" "test" {
  repository 	   = github_repository.test.name
  environment	   = github_repository_environment.test.environment
  branch_pattern = "releases/*"
}
```

## Argument Reference

The following arguments are supported:

* `environment` - (Required) The name of the environment.

* `repository` - (Required) The repository of the environment.

* `branch_pattern` - (Required) The name pattern that branches must match in order to deploy to the environment.


## Import

GitHub Repository Environment Deployment Policy can be imported using an ID made up of `name` of the repository combined with the `environment` name of the environment with the `Id` of the deployment policy, separated by a `:` character, e.g.

```
$ terraform import github_repository_environment.daily terraform:daily:123456
```
