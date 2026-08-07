---
page_title: "github_enterprise_actions_hosted_runner (Resource) - GitHub"
description: |-
  Creates and manages GitHub-hosted runners within a GitHub enterprise
---

# github_enterprise_actions_hosted_runner (Resource)

This resource allows you to create and manage GitHub-hosted runners within your GitHub enterprise. You must have enterprise admin permissions to use this resource.

GitHub-hosted runners are fully managed virtual machines that run your GitHub Actions workflows. Unlike self-hosted runners, GitHub handles the infrastructure, maintenance, and scaling.

## Example Usage

### Basic Usage

```terraform
data "github_enterprise" "example" {
  slug = "example-co"
}

resource "github_enterprise_actions_runner_group" "example" {
  enterprise_slug = data.github_enterprise.example.slug
  name            = "example-runner-group"
  visibility      = "all"
}

resource "github_enterprise_actions_hosted_runner" "example" {
  enterprise_slug = data.github_enterprise.example.slug
  name            = "example-hosted-runner"

  image {
    id     = "2306"
    source = "github"
  }

  size            = "4-core"
  runner_group_id = github_enterprise_actions_runner_group.example.id
}
```

### Advanced Usage with Optional Parameters

```terraform
data "github_enterprise" "example" {
  slug = "example-co"
}

resource "github_enterprise_actions_runner_group" "advanced" {
  enterprise_slug = data.github_enterprise.example.slug
  name            = "advanced-runner-group"
  visibility      = "selected"
}

resource "github_enterprise_actions_hosted_runner" "advanced" {
  enterprise_slug = data.github_enterprise.example.slug
  name            = "advanced-hosted-runner"

  image {
    id     = "2306"
    source = "github"
  }

  size              = "8-core"
  runner_group_id   = github_enterprise_actions_runner_group.advanced.id
  maximum_runners   = 10
  public_ip_enabled = true
}
```

## Argument Reference

The following arguments are supported:

- `enterprise_slug` - (Required) The slug of the enterprise. Cannot be changed after creation.
- `name` - (Required) Name of the hosted runner. Must be between 1 and 64 characters and may only contain alphanumeric characters, '.', '-', and '_'.
- `image` - (Required) Image configuration for the hosted runner. Cannot be changed after creation. Block supports:
  - `id` - (Required) The image ID. For GitHub-owned images, use numeric IDs like "2306" for Ubuntu Latest 24.04. To get available images, use the GitHub API: `GET /enterprises/{enterprise}/actions/hosted-runners/images/github-owned`.
  - `source` - (Optional) The image source. Valid values are "github", "partner", or "custom". Defaults to "github".
  - `version` - (Optional) The version of the runner image to deploy. For GitHub-owned images, this must be "latest" (default). For custom images, you can specify a specific version.
- `size` - (Required) Machine size for the hosted runner (e.g., "4-core", "8-core"). Can be updated to scale the runner. To list available sizes, use the GitHub API: `GET /enterprises/{enterprise}/actions/hosted-runners/machine-sizes`.
- `runner_group_id` - (Required) The ID of the runner group to assign this runner to.
- `maximum_runners` - (Optional) Maximum number of runners to scale up to. Runners will not auto-scale above this number. Use this setting to limit costs.
- `public_ip_enabled` - (Optional) Whether to enable static public IP for the runner. Note there are account-level limits. To check limits, use the GitHub API: `GET /enterprises/{enterprise}/actions/hosted-runners/limits`. Defaults to false.
- `image_gen` - (Optional) Whether this runner should be used to generate custom images. Cannot be changed after creation. Defaults to false.

## Timeouts

The `timeouts` block allows you to specify timeouts for certain actions:

- `delete` - (Defaults to 15 minutes) Used for waiting for the hosted runner deletion to complete.

Example:

```terraform
data "github_enterprise" "example" {
  slug = "example-co"
}

resource "github_enterprise_actions_runner_group" "example" {
  enterprise_slug = data.github_enterprise.example.slug
  name            = "example-runner-group"
  visibility      = "all"
}

resource "github_enterprise_actions_hosted_runner" "example" {
  enterprise_slug = data.github_enterprise.example.slug
  name            = "example-hosted-runner"

  image {
    id     = "2306"
    source = "github"
  }

  size            = "4-core"
  runner_group_id = github_enterprise_actions_runner_group.example.id

  timeouts {
    delete = "20m"
  }
}
```

## Attributes Reference

In addition to the arguments above, the following attributes are exported:

- `id` - The ID of the hosted runner in the format `{enterprise_slug}:{runner_id}`.
- `runner_id` - The numeric ID of the hosted runner.
- `status` - Current status of the runner (e.g., "Ready", "Provisioning").
- `platform` - Platform of the runner (e.g., "linux-x64", "win-x64").
- `image` - In addition to the arguments above, the image block exports:
  - `size_gb` - The size of the image in gigabytes.
  - `display_name` - Human-readable display name for the image.
- `machine_size_details` - Detailed specifications of the machine size:
  - `id` - Machine size identifier.
  - `cpu_cores` - Number of CPU cores.
  - `memory_gb` - Amount of memory in gigabytes.
  - `storage_gb` - Amount of storage in gigabytes.
- `public_ips` - List of public IP ranges assigned to this runner (only if `public_ip_enabled` is true):
  - `enabled` - Whether this IP range is enabled.
  - `prefix` - IP address prefix.
  - `length` - Subnet length.
- `last_active_on` - Timestamp (RFC3339) when the runner was last active.

## Import

Enterprise hosted runners can be imported using the enterprise slug and runner ID, separated by a colon:

```shell
terraform import github_enterprise_actions_hosted_runner.example my-enterprise:123456
```
