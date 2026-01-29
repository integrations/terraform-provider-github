---
layout: "github"
page_title: "GitHub: github_repository_environment_custom_deployment_protection_rules"
description: |-
  Gets all custom deployment protection rule integrations that are available for an environment.
---

# github_repository_environment_custom_deployment_protection_rules

Use this data source to retrieve all custom deployment protection rules available for an environment.

## Example Usage

```hcl
data "github_repository_environment_custom_deployment_protection_rules" "example" {
    repository        = "example-repository"
    environment  = "env_name"
}
```

## Argument Reference

* `repository` - (Required) Name of the repository to retrieve the custom deployment protection rules from.

* `environment` - (Required) Name of the environment to retrieve the custom deployment protection rules from.

## Attributes Reference

* `custom_deployment_protection_rules` - The list of this repository's environments. Each element of `custom_deployment_protection_rules` has the following attributes:
    * `id` - The ID of the custom deployment protection rule.
    * `slug` - The URL-friendly name of the GitHub App.
    * `integration_url` - The API endpoint for the GitHub App.
    * `node_id` - The Node ID of the custom deployment protection rule.
