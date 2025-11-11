---
layout: "github"
page_title: "GitHub: github_repository_pre_receive_hook"
description: |-
  Creates and manages repository pre-receive hooks.
---

# github_repository_pre_receive_hook

This resource allows you to create and manage pre-receive hooks for repositories.

~> **Note** Repository pre-receive hooks are currently only available in GitHub Enterprise Server.

## Example usage

```
resource "github_repository_pre_receive_hook" "example" {
  repository  = "test-repo"
  name        = "ensure-conventional-commits"
  enforcement = "enabled"
}
```

## Argument Reference

The following arguments are supported:

* `repository` - (Required) The repository of the pre-receive hook.

* `name` - (Required) The name of the pre-receive hook.

* `enforcement` - (Required) The state of enforcement for the hook on the repository. Possible values for enforcement are `enabled`, `disabled` and `testing`. `disabled` indicates the pre-receive hook will not run. `enabled` indicates it will run and reject any pushes that result in a non-zero status. `testing` means the script will run but will not cause any pushes to be rejected.

## Attributes Reference

The following additional attributes are exported:

* `configuration_url` - The URL for the endpoint where enforcement is set.
