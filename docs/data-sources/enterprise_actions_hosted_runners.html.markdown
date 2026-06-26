---
layout: "github"
page_title: "GitHub: github_enterprise_actions_hosted_runners"
description: |-
  Get information about all GitHub-hosted runners for an enterprise
---

# github\_enterprise\_actions\_hosted\_runners

Use this data source to retrieve information about all GitHub-hosted runners in a GitHub enterprise.

## Example Usage

```hcl
data "github_enterprise_actions_hosted_runners" "all" {
  enterprise_slug = "my-enterprise"
}

# Output all runner names
output "runner_names" {
  value = [for runner in data.github_enterprise_actions_hosted_runners.all.runners : runner.name]
}

```

## Argument Reference

The following arguments are supported:

* `enterprise_slug` - (Required) The slug of the enterprise.

## Attributes Reference

* `runners` - List of hosted runners for the enterprise. Each runner contains:
  * `id` - The ID of the hosted runner.
  * `name` - The name of the hosted runner.
  * `runner_group_id` - The runner group ID this runner belongs to.
  * `platform` - The platform of the runner (e.g., `linux-x64`, `win-x64`, `macos-arm64`).
  * `status` - Current status of the runner (e.g., `Ready`, `Provisioning`, `Deleting`).
  * `maximum_runners` - Maximum number of runners to scale up to.
  * `public_ip_enabled` - Whether static public IP is enabled for this runner.
  * `last_active_on` - RFC3339 timestamp indicating when the runner was last active.
  * `image_details` - Details about the runner's image:
    * `id` - The image ID.
    * `source` - The image source (`github`, `partner`, or `custom`).
    * `version` - The image version.
    * `size_gb` - The size of the image in GB.
    * `display_name` - Human-readable display name for the image.
  * `machine_size_details` - Details about the runner's machine size:
    * `id` - Machine size identifier (e.g., `4-core`, `8-core`).
    * `cpu_cores` - Number of CPU cores.
    * `memory_gb` - Amount of memory in GB.
    * `storage_gb` - Amount of SSD storage in GB.
