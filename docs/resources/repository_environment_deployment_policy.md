---
page_title: "GitHub: github_repository_environment_deployment_policy"
description: |-
  Creates and manages environment deployment branch policies for GitHub repositories
---

# github_repository_environment_deployment_policy

This resource allows you to create and manage environment deployment branch policies for a GitHub repository.

## Example Usage

Create a branch-based deployment policy:

```terraform
data "github_user" "current" {
  username = ""
}

resource "github_repository" "test" {
  name      = "tf-acc-test-%s"
}

resource "github_repository_environment" "test" {
  repository    = github_repository.test.name
  environment   = "environment/test"
  wait_timer    = 10000
  reviewers {
    users = [data.github_user.current.id]
  }
  deployment_branch_policy {
    protected_branches     = false
    custom_branch_policies = true
  }
}

resource "github_repository_environment_deployment_policy" "test" {
  repository     = github_repository.test.name
  environment    = github_repository_environment.test.environment
  branch_pattern = "releases/*"
}
```

Create a tag-based deployment policy:

```terraform
data "github_user" "current" {
  username = ""
}

resource "github_repository" "test" {
  name      = "tf-acc-test-%s"
}

resource "github_repository_environment" "test" {
  repository  = github_repository.test.name
  environment = "environment/test"
  wait_timer  = 10000
  reviewers {
    users = [data.github_user.current.id]
  }
  deployment_branch_policy {
    protected_branches     = false
    custom_branch_policies = true
  }
}

resource "github_repository_environment_deployment_policy" "test" {
  repository  = github_repository.test.name
  environment = github_repository_environment.test.environment
  tag_pattern = "v*"
}
```

## Argument Reference

The following arguments are supported:

- `environment` - (Required) The name of the environment.

- `repository` - (Required) The repository of the environment.

- `branch_pattern` - (Optional) The name pattern that branches must match in order to deploy to the environment. If not specified, `tag_pattern` must be specified.

- `tag_pattern` - (Optional) The name pattern that tags must match in order to deploy to the environment. If not specified, `branch_pattern` must be specified.

## Attributes Reference

- `repository_id` - The ID of the repository.
- `policy_id` - The ID of the deployment policy.

## Import

This resource can be imported using an ID made of the repository name, environment name (any `:` in the environment name need to be escaped as `??`), and deployment policy ID name all separated by a `:`.

### Import Block

The following import block imports a deployment policy with the ID `123456` for the repo `myrepo` and environment `myenv` to a `github_repository_environment_deployment_policy` resource named `example`.

```terraform
import {
  to = github_repository_environment_deployment_policy.example
  id = "myrepo:myenv:123456"
}
```

### Import Command

The following command imports a deployment policy with the ID `123456` for the repo `myrepo` and environment `myenv` to a `github_repository_environment_deployment_policy` resource named `example`.

```shell
terraform import github_repository_environment_deployment_policy.example myrepo:myenv:123456
```
