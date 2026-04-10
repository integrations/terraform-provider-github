---
layout: "github"
page_title: "GitHub: github_organization_network_configuration"
description: |-
  Creates and manages hosted compute network configurations for a GitHub organization.
---

# github_organization_network_configuration

This resource allows you to create and manage hosted compute network configurations for a GitHub Organization. Network configurations allow GitHub-hosted compute services, such as Actions hosted runners, to connect to your private network resources.

~> **Note:** This resource is organization-only and is available for GitHub Enterprise Cloud organizations. See the [GitHub documentation](https://docs.github.com/en/enterprise-cloud@latest/rest/orgs/network-configurations) for more information.

~> **Note:** Organization-level network configurations are only available when enterprise policy allows organizations to create their own hosted compute network configurations. Otherwise, organizations can only inherit enterprise-level network configurations.

## Example Usage

```hcl
resource "github_organization_network_configuration" "example" {
  name                 = "my-network-config"
  compute_service      = "actions"
  network_settings_ids = ["23456789ABCDEF1"]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the network configuration. Must be between 1 and 100 characters and may only contain upper and lowercase letters a-z, numbers 0-9, `.`, `-`, and `_`.

* `compute_service` - (Optional) The hosted compute service to use for the network configuration. Can be one of `none` or `actions`. Defaults to `none`.

* `network_settings_ids` - (Required) An array containing exactly one network settings ID. Network settings resources are configured separately through your cloud provider. For Azure private networking, use the `GitHubId` returned by the Azure `GitHub.Network/networkSettings` resource, not the Azure ARM resource ID. A network settings resource can only be associated with one network configuration at a time.

## Attributes Reference

In addition to the arguments above, the following attributes are exported:

* `id` - The ID of the network configuration.

* `created_on` - The timestamp when the network configuration was created.

## Notes

* This resource can only be used with organization accounts.
* GitHub currently allows exactly one `network_settings_ids` value per organization network configuration.
* The `network_settings_ids` value must reference an existing hosted compute network settings resource configured outside this provider.
* For organization-scoped configurations backed by Azure private networking, create the Azure `GitHub.Network/networkSettings` resource using the GitHub organization's `databaseId`. Using a mismatched scope, such as an enterprise `databaseId` for an organization configuration, can cause GitHub to reject the configuration.

## Import

Organization network configurations can be imported using the network configuration ID:

```shell
terraform import github_organization_network_configuration.example 1234567890ABCDEF
```

The network configuration ID can be found using the [list hosted compute network configurations for an organization](https://docs.github.com/en/enterprise-cloud@latest/rest/orgs/network-configurations#list-hosted-compute-network-configurations-for-an-organization) API.
