---
page_title: "github_organization_network_configuration (Resource) - GitHub"
description: |-
  Creates and manages a hosted compute network configuration within a GitHub organization
---

# github_organization_network_configuration (Resource)

This resource allows you to create and manage hosted compute network configurations for a GitHub organization. A network configuration associates an Azure virtual network with GitHub-hosted runners so that workflow jobs run inside your own private network. Assign the configuration to a runner group with `github_actions_runner_group.network_configuration_id`.

Before using this resource, create the backing `GitHub.Network/networkSettings` resource in Azure and register it against the same organization. The `GitHubId` returned by Azure is the value to pass in `network_settings_ids`.

## Example Usage

```terraform
resource "github_organization_network_configuration" "example" {
  name                 = "my-network-configuration"
  compute_service      = "actions"
  network_settings_ids = ["23456789ABCDEF1"]
}

resource "github_actions_runner_group" "example" {
  name                     = "my-runner-group"
  visibility               = "all"
  network_configuration_id = github_organization_network_configuration.example.id
}
```

## Argument Reference

The following arguments are supported:

- `name` - (Required) Name of the network configuration. Must be between 1 and 100 characters and may only contain upper and lowercase letters a-z, numbers 0-9, `.`, `-`, and `_`.
- `network_settings_ids` - (Required) A list containing exactly one network settings ID. A network settings resource can only be associated with one network configuration at a time.
- `compute_service` - (Optional) The hosted compute service the network configuration supports. Can be `none` or `actions`. Defaults to `none`.

## Attributes Reference

- `id` - The ID of the network configuration.
- `created_on` - Timestamp of when the network configuration was created, in RFC3339 format.

## Import

This resource can be imported using the ID of the network configuration:

```shell
terraform import github_organization_network_configuration.example 123456789ABCDEF
```
