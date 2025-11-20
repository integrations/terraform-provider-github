---
layout: "github"
page_title: "GitHub: github_enterprise_actions_workflow_permissions"
description: |-
  Manages GitHub Actions workflow permissions for a GitHub Enterprise.
---

# github_enterprise_actions_workflow_permissions

This resource allows you to manage GitHub Actions workflow permissions for a GitHub Enterprise account. This controls the default permissions granted to the GITHUB_TOKEN when running workflows and whether GitHub Actions can approve pull request reviews.

You must have enterprise admin access to use this resource.

## Example Usage

```hcl
# Basic workflow permissions configuration
resource "github_enterprise_actions_workflow_permissions" "example" {
  enterprise_slug = "my-enterprise"
  
  default_workflow_permissions     = "read"
  can_approve_pull_request_reviews = false
}

# Allow write permissions and PR approvals
resource "github_enterprise_actions_workflow_permissions" "permissive" {
  enterprise_slug = "my-enterprise"
  
  default_workflow_permissions     = "write"
  can_approve_pull_request_reviews = true
}
```

## Argument Reference

The following arguments are supported:

* `enterprise_slug` - (Required) The slug of the enterprise.

* `default_workflow_permissions` - (Optional) The default workflow permissions granted to the GITHUB_TOKEN when running workflows. Can be `read` or `write`. Defaults to `read`.

* `can_approve_pull_request_reviews` - (Optional) Whether GitHub Actions can approve pull request reviews. Defaults to `false`.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The enterprise slug.

## Import

Enterprise Actions workflow permissions can be imported using the enterprise slug:

```
terraform import github_enterprise_actions_workflow_permissions.example my-enterprise
```

## Notes

~> **Note:** This resource requires a GitHub Enterprise account and enterprise admin permissions.

When this resource is destroyed, the workflow permissions will be reset to safe defaults:
- `default_workflow_permissions` = `read`
- `can_approve_pull_request_reviews` = `false`