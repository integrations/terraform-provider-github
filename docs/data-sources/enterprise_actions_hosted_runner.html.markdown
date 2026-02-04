---
layout: "github"
page_title: "GitHub: github_enterprise_actions_hosted_runner"
description: |-
  Get information about a specific GitHub-hosted runner in an enterprise
---

# github\_enterprise\_actions\_hosted\_runner

Use this data source to retrieve information about a specific GitHub-hosted runner in a GitHub enterprise by name.

## Example Usage

```hcl
# Get a specific runner by name
data "github_enterprise_actions_hosted_runner" "prod_runner" {
  enterprise_slug = "my-enterprise"
  name            = "prod-runner-01"
}
```

## Argument Reference

The following arguments are supported:

* `enterprise_slug` - (Required) The slug of the enterprise.
* `name` - (Required) The name of the hosted runner to lookup.

## Attributes Reference

* `runner_id` - The numeric ID of the hosted runner.
* `runner_group_id` - The runner group ID this runner belongs to.
* `platform` - The platform of the runner (e.g., `linux-x64`, `win-x64`, `macos-arm64`).
* `status` - Current status of the runner (e.g., `Ready`, `Provisioning`, `Deleting`).
* `maximum_runners` - Maximum number of runners to scale up to.
* `public_ip_enabled` - Whether static public IP is enabled for this runner.
* `last_active_on` - RFC3339 timestamp indicating when the runner was last active.
* `image_details` - Details about the runner's image:
  * `id` - The image ID (e.g., `2306` for Ubuntu Latest 24.04).
  * `source` - The image source (`github`, `partner`, or `custom`).
  * `version` - The image version (e.g., `latest`).
  * `size_gb` - The size of the image in GB.
  * `display_name` - Human-readable display name for the image (e.g., `Ubuntu Latest 24.04`).
* `machine_size_details` - Details about the runner's machine size:
  * `id` - Machine size identifier (e.g., `4-core`, `8-core`, `16-core`).
  * `cpu_cores` - Number of CPU cores.
  * `memory_gb` - Amount of memory in GB.
  * `storage_gb` - Amount of SSD storage in GB.
* `public_ips` - List of public IP ranges assigned to this runner (only populated if `public_ip_enabled` is true):
  * `enabled` - Whether this IP range is enabled.
  * `prefix` - IP address prefix (e.g., `20.80.208.150`).
  * `length` - Subnet length in CIDR notation (e.g., `28`).
