---
layout: "github"
page_title: "GitHub: github_actions_hosted_runner"
description: |-
  Creates and manages GitHub-hosted runners within a GitHub organization
---

# github_actions_hosted_runner

This resource allows you to create and manage GitHub-hosted runners within your GitHub organization.
You must have admin access to an organization to use this resource.

GitHub-hosted runners are fully managed virtual machines that run your GitHub Actions workflows. Unlike self-hosted runners, GitHub handles the infrastructure, maintenance, and scaling.

## Example Usage

### Basic Usage

```hcl
resource "github_actions_runner_group" "example" {
  name       = "example-runner-group"
  visibility = "all"
}

resource "github_actions_hosted_runner" "example" {
  name = "example-hosted-runner"
  
  image {
    id     = "ubuntu-latest"
    source = "github"
  }

  size            = "4-core"
  runner_group_id = github_actions_runner_group.example.id
}
```

### Advanced Usage with Optional Parameters

```hcl
resource "github_actions_runner_group" "advanced" {
  name       = "advanced-runner-group"
  visibility = "selected"
}

resource "github_actions_hosted_runner" "advanced" {
  name = "advanced-hosted-runner"
  
  image {
    id     = "ubuntu-latest"
    source = "github"
  }

  size             = "8-core"
  runner_group_id  = github_actions_runner_group.advanced.id
  maximum_runners  = 10
  enable_static_ip = true
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the hosted runner.
* `image` - (Required) Image configuration for the hosted runner. Block supports:
  * `id` - (Required) The image ID (e.g., "ubuntu-latest").
  * `source` - (Optional) The image source. Defaults to "github".
* `size` - (Required) Machine size for the hosted runner (e.g., "4-core", "8-core").
* `runner_group_id` - (Required) The ID of the runner group to assign this runner to.
* `maximum_runners` - (Optional) Maximum number of runners to scale up to.
* `enable_static_ip` - (Optional) Whether to enable static IP for the runner. Defaults to false.

## Attributes Reference

In addition to the arguments above, the following attributes are exported:

* `id` - The ID of the hosted runner.
* `status` - Current status of the runner.
* `platform` - Platform of the runner (e.g., "linux", "windows").

## Import

Hosted runners can be imported using the runner ID:

```
$ terraform import github_actions_hosted_runner.example 123456
```

## Notes

* This resource is **organization-only** and cannot be used with individual accounts.
* Deletion of hosted runners is asynchronous. The provider will poll for up to 10 minutes to confirm deletion.
* Runner creation and updates may take several minutes as GitHub provisions the infrastructure.
