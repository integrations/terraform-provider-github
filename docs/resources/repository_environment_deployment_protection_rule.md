---
page_title: "github_repository_environment_deployment_protection_rule (Resource) - GitHub"
description: |-
  Creates and manages custom deployment protection rules for GitHub repository environments
---

# github_repository_environment_deployment_protection_rule (Resource)

This resource enables a custom deployment protection rule — a GitHub App that gates
deployments to an environment — on a GitHub repository environment. When a workflow job
targets the environment, GitHub sends the App a `deployment_protection_rule` webhook and
waits for the App to approve or reject the deployment.

The App must be installed on the repository and available as a custom deployment
protection rule for the environment before it can be enabled.

## Example Usage

```terraform
resource "github_repository" "test" {
  name = "tf-acc-test"
}

resource "github_repository_environment" "test" {
  repository  = github_repository.test.name
  environment = "prod"
}

resource "github_repository_environment_deployment_protection_rule" "test" {
  repository     = github_repository.test.name
  environment    = github_repository_environment.test.environment
  integration_id = 4358979 # the GitHub App's ID
}
```

## Argument Reference

The following arguments are supported:

- `repository` - (Required) The name of the GitHub repository.

- `environment` - (Required) The name of the environment. Changing this forces a new resource.

- `integration_id` - (Required) The ID of the custom deployment protection rule integration — the GitHub App that gates deployments. Changing this forces a new resource.

## Attributes Reference

- `id` - An ID made of the repository name, environment name, and deployment protection rule ID, separated by `:`.

## Import

This resource can be imported using an ID made of the repository name, environment name (any `:` in the environment name needs to be escaped as `??`), and deployment protection rule ID, all separated by a `:`.

### Import Block

```terraform
import {
  to = github_repository_environment_deployment_protection_rule.example
  id = "myrepo:myenv:123456"
}
```

### Import Command

```shell
terraform import github_repository_environment_deployment_protection_rule.example myrepo:myenv:123456
```
