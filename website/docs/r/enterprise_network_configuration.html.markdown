---
layout: "github"
page_title: "GitHub: github_enterprise_network_configuration"
description: |-
  Creates and manages hosted compute network configurations for a GitHub enterprise.
---

# github_enterprise_network_configuration

This resource allows you to create and manage hosted compute network configurations for a GitHub Enterprise. Network configurations allow GitHub-hosted compute services, such as Actions hosted runners, to connect to your private network resources.

~> **Note:** This resource is enterprise-only and is available for GitHub Enterprise Cloud enterprises. See the [GitHub documentation](https://docs.github.com/en/enterprise-cloud@latest/rest/enterprise-admin/network-configurations) for more information.

## Example Usage

```hcl
data "github_enterprise" "enterprise" {
  slug = "my-enterprise"
}

resource "github_enterprise_network_configuration" "example" {
  enterprise_slug      = data.github_enterprise.enterprise.slug
  name                 = "my-network-config"
  compute_service      = "actions"
  network_settings_ids = ["23456789ABCDEF1"]
}
```

## Argument Reference

The following arguments are supported:

* `enterprise_slug` - (Required) The slug of the enterprise.

* `name` - (Required) The name of the network configuration. Must be between 1 and 100 characters and may only contain upper and lowercase letters a-z, numbers 0-9, `.`, `-`, and `_`.

* `compute_service` - (Optional) The hosted compute service to use for the network configuration. Can be one of `none` or `actions`. Defaults to `none`.

* `network_settings_ids` - (Required) An array containing exactly one network settings ID. Network settings resources are configured separately through your cloud provider. For Azure private networking, use the `GitHubId` returned by the Azure `GitHub.Network/networkSettings` resource, not the Azure ARM resource ID. A network settings resource can only be associated with one network configuration at a time.

## Attributes Reference

In addition to the arguments above, the following attributes are exported:

* `id` - The ID of the network configuration.

* `created_on` - The timestamp when the network configuration was created.

## Notes

* This resource can only be used with enterprise accounts.
* GitHub currently allows exactly one `network_settings_ids` value per enterprise network configuration.
* The `network_settings_ids` value must reference an existing hosted compute network settings resource configured outside this provider.
* Use `github_enterprise_actions_runner_group.network_configuration_id` to associate this configuration with GitHub-hosted runners through an enterprise runner group.

## Import

Enterprise network configurations can be imported using the enterprise slug and network configuration ID:

```shell
terraform import github_enterprise_network_configuration.example enterprise-slug/1234567890ABCDEF
```

The network configuration ID can be found using the [list hosted compute network configurations for an enterprise](https://docs.github.com/en/enterprise-cloud@latest/rest/enterprise-admin/network-configurations#list-hosted-compute-network-configurations-for-an-enterprise) API.
