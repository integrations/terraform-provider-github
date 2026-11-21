---
page_title: "github_repository_environment_deployment_protection_rule (Resource) - GitHub"
description: |-
  Creates and manages custom deployment protection rules for GitHub repository environments
---

# github_repository_environment_deployment_protection_rule (Resource)

Manages a custom deployment protection rule for a GitHub repository environment. A custom deployment protection rule is a GitHub App that approves or rejects deployments to the environment.

When a workflow job targets the environment, GitHub sends the App a `deployment_protection_rule` webhook and waits for the App to approve or reject the deployment. The App must be installed on the repository and available as a custom deployment protection rule for the environment before it can be enabled.

## Example Usage

```terraform
resource "github_repository" "example" {
  name = "example"
}

resource "github_repository_environment" "example" {
  repository  = github_repository.example.name
  environment = "production"
}

resource "github_repository_environment_deployment_protection_rule" "example" {
  repository     = github_repository.example.name
  environment    = github_repository_environment.example.environment
  integration_id = 123456 # The ID of the GitHub App
}
```

## Argument Reference

The following arguments are supported:

- `repository` - (Required) The name of the GitHub repository.

- `environment` - (Required) The name of the environment. Changing this forces a new resource.

- `integration_id` - (Required) The ID of the GitHub App to enable as a custom deployment protection rule. Changing this forces a new resource.

## Attributes Reference

- `repository_id` - The ID of the repository.
- `protection_rule_id` - The ID of the custom deployment protection rule.

## Import

This resource can be imported using an ID made of the repository name, environment name (any `:` in the environment name needs to be escaped as `??`), and custom deployment protection rule ID, all separated by a `:`.

### Import Block

The following import block imports a custom deployment protection rule with the ID `123456` for the repo `myrepo` and environment `myenv` to a `github_repository_environment_deployment_protection_rule` resource named `example`.

```terraform
import {
  to = github_repository_environment_deployment_protection_rule.example
  id = "myrepo:myenv:123456"
}
```

### Import Command

The following command imports a custom deployment protection rule with the ID `123456` for the repo `myrepo` and environment `myenv` to a `github_repository_environment_deployment_protection_rule` resource named `example`.

```shell
terraform import github_repository_environment_deployment_protection_rule.example myrepo:myenv:123456
```
