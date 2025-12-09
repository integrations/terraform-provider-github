---
page_title: "GitHub: github_repository_environment_deployment_policies"
description: |-
  Get the list of environment deployment policies for a given repository environment.
---

# github_repository_environment_deployment_policies

Use this data source to retrieve deployment branch policies for a repository environment.

## Example Usage

```terraform
data "github_repository_environment_deployment_policies" "example" {
    repository  = "example-repository"
    environment = "env-name"
}
```

## Argument Reference

- `repository` - (Required) Name of the repository to retrieve the deployment branch policies from.

- `environment` - (Required) Name of the environment to retrieve the deployment branch policies from.

## Attributes Reference

- `policies` - The list of deployment policies for the repository environment. Each element of `policies` has the following attributes:
    - `type` - Type of the policy; this could be `branch` or `tag`.
    - `pattern` - The pattern that branch or tag names must match in order to deploy to the environment.
