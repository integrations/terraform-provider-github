---
layout: "github"
page_title: "GitHub: github_organization_network_configuration"
description: |-
  Creates and manages network configurations for GitHub Actions hosted runners in an organization.
---

# github_organization_network_configuration

This resource allows you to create and manage network configurations for GitHub Actions hosted runners in a GitHub Organization. Network configurations enable you to configure networking settings for hosted compute services.

~> **Note:** This resource is only available for GitHub Enterprise Cloud organizations. See the [GitHub documentation](https://docs.github.com/en/enterprise-cloud@latest/rest/orgs/network-configurations) for more information.

## Example Usage

```hcl
resource "github_organization_network_configuration" "example" {
  name                 = "my-network-config"
  compute_service      = "actions"
  network_settings_ids = ["23456789ABDCEF1"]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the network configuration. Must be between 1 and 100 characters and may only contain upper and lowercase letters a-z, numbers 0-9, '.', '-', and '_'.

* `compute_service` - (Optional) The hosted compute service to use for the network configuration. Can be one of `none` or `actions`. Defaults to `none`.

* `network_settings_ids` - (Required) An array containing exactly one network settings ID. Network settings resources are configured separately through your cloud provider (e.g., Azure). **Note:** A network settings resource can only be associated with one network configuration at a time.

## Attributes Reference

The following additional attributes are exported:

* `id` - The ID of the network configuration.

* `created_on` - The timestamp when the network configuration was created.

## Import

Organization network configurations can be imported using the network configuration ID:

```
$ terraform import github_organization_network_configuration.example 1234567890ABCDEF
```

The network configuration ID can be found using the [List network configurations](https://docs.github.com/en/enterprise-cloud@latest/rest/orgs/network-configurations#list-hosted-compute-network-configurations-for-an-organization) API endpoint.
