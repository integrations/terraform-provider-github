---
page_title: "GitHub: github_actions_organization_workflow_permissions"
description: |-
  Manages GitHub Actions workflow permissions for a GitHub Organization.
---

# github_actions_organization_workflow_permissions

This resource allows you to manage GitHub Actions workflow permissions for a GitHub Organization account. This controls the default permissions granted to the GITHUB_TOKEN when running workflows and whether GitHub Actions can approve pull request reviews.

You must have organization admin access to use this resource.

## Example Usage

```terraform
# Basic workflow permissions configuration
resource "github_actions_organization_workflow_permissions" "example" {
  organization_slug = "my-organization"

  default_workflow_permissions     = "read"
  can_approve_pull_request_reviews = false
}

# Allow write permissions and PR approvals
resource "github_actions_organization_workflow_permissions" "permissive" {
  organization_slug = "my-organization"

  default_workflow_permissions     = "write"
  can_approve_pull_request_reviews = true
}
```

## Argument Reference

The following arguments are supported:

* `organization_slug` - (Required) The slug of the organization.

* `default_workflow_permissions` - (Optional) The default workflow permissions granted to the GITHUB_TOKEN when running workflows. Can be `read` or `write`. Defaults to `read`.

* `can_approve_pull_request_reviews` - (Optional) Whether GitHub Actions can approve pull request reviews. Defaults to `false`.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The organization slug.

## Import

Organization Actions workflow permissions can be imported using the organization slug:

```sh
terraform import github_actions_organization_workflow_permissions.example my-organization
```

## Notes

~> **Note:** This resource requires a GitHub Organization account and organization admin permissions.

When this resource is destroyed, the workflow permissions will be reset to safe defaults:

* `default_workflow_permissions` = `read`
* `can_approve_pull_request_reviews` = `false`
