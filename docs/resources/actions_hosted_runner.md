---
page_title: "github_actions_hosted_runner (Resource) - GitHub"
description: |-
  Creates and manages GitHub-hosted runners within a GitHub organization
---

# github_actions_hosted_runner (Resource)

This resource allows you to create and manage GitHub-hosted runners within your GitHub organization. You must have admin access to an organization to use this resource.

GitHub-hosted runners are fully managed virtual machines that run your GitHub Actions workflows. Unlike self-hosted runners, GitHub handles the infrastructure, maintenance, and scaling.

## Example Usage

### Basic Usage

```terraform
resource "github_actions_runner_group" "example" {
  name       = "example-runner-group"
  visibility = "all"
}

resource "github_actions_hosted_runner" "example" {
  name = "example-hosted-runner"

  image {
    id     = "2306"
    source = "github"
  }

  size            = "4-core"
  runner_group_id = github_actions_runner_group.example.id
}
```

### Advanced Usage with Optional Parameters

```terraform
resource "github_actions_runner_group" "advanced" {
  name       = "advanced-runner-group"
  visibility = "selected"
}

resource "github_actions_hosted_runner" "advanced" {
  name = "advanced-hosted-runner"

  image {
    id     = "2306"
    source = "github"
  }

  size              = "8-core"
  runner_group_id   = github_actions_runner_group.advanced.id
  maximum_runners   = 10
  public_ip_enabled = true
}
```

## Argument Reference

The following arguments are supported:

- `name` - (Required) Name of the hosted runner. Must be between 1 and 64 characters and may only contain alphanumeric characters, '.', '-', and '_'.
- `image` - (Required) Image configuration for the hosted runner. Cannot be changed after creation. Block supports:
  - `id` - (Required) The image ID. For GitHub-owned images, use numeric IDs like "2306" for Ubuntu Latest 24.04. To get available images, use the GitHub API: `GET /orgs/{org}/actions/hosted-runners/images/github-owned`.
  - `source` - (Optional) The image source. Valid values are "github", "partner", or "custom". Defaults to "github".
- `size` - (Required) Machine size for the hosted runner (e.g., "4-core", "8-core"). Can be updated to scale the runner. To list available sizes, use the GitHub API: `GET /orgs/{org}/actions/hosted-runners/machine-sizes`.
- `runner_group_id` - (Required) The ID of the runner group to assign this runner to.
- `maximum_runners` - (Optional) Maximum number of runners to scale up to. Runners will not auto-scale above this number. Use this setting to limit costs.
- `public_ip_enabled` - (Optional) Whether to enable static public IP for the runner. Note there are account limits. To list limits, use the GitHub API: `GET /orgs/{org}/actions/hosted-runners/limits`. Defaults to false.
- `image_version` - (Optional) The version of the runner image to deploy. This is only relevant for runners using custom images.

## Timeouts

The `timeouts` block allows you to specify timeouts for certain actions:

- `delete` - (Defaults to 10 minutes) Used for waiting for the hosted runner deletion to complete.

Example:

```terraform
resource "github_actions_hosted_runner" "example" {
  name = "example-hosted-runner"

  image {
    id     = "2306"
    source = "github"
  }

  size            = "4-core"
  runner_group_id = github_actions_runner_group.example.id

  timeouts {
    delete = "15m"
  }
}
```

## Attributes Reference

In addition to the arguments above, the following attributes are exported:

- `id` - The ID of the hosted runner.
- `status` - Current status of the runner (e.g., "Ready", "Provisioning").
- `platform` - Platform of the runner (e.g., "linux-x64", "win-x64").
- `image` - In addition to the arguments above, the image block exports:
  - `size_gb` - The size of the image in gigabytes.
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

Hosted runners can be imported using the runner ID:

```hcl
$ terraform import github_actions_hosted_runner.example 123456
```

## Notes

- This resource is **organization-only*- and cannot be used with individual accounts.
- The `image` field cannot be changed after the runner is created. Changing it will force recreation of the runner.
- The `size` field can be updated to scale the runner up or down as needed.
- Image IDs for GitHub-owned images are numeric strings (e.g., "2306" for Ubuntu Latest 24.04), not names like "ubuntu-latest".
- Deletion of hosted runners is asynchronous. The provider will poll for up to 10 minutes (configurable via timeouts) to confirm deletion.
- Runner creation and updates may take several minutes as GitHub provisions the infrastructure.
- Static public IPs are subject to account limits. Check your organization's limits before enabling.

## Getting Available Images and Sizes

To get a list of available images:

```bash
curl -H "Authorization: Bearer YOUR_TOKEN" \
     -H "Accept: application/vnd.github+json" \
     https://api.github.com/orgs/YOUR_ORG/actions/hosted-runners/images/github-owned
```

To get available machine sizes:

```bash
curl -H "Authorization: Bearer YOUR_TOKEN" \
     -H "Accept: application/vnd.github+json" \
     https://api.github.com/orgs/YOUR_ORG/actions/hosted-runners/machine-sizes
```
