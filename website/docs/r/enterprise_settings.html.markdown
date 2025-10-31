---
layout: "github"
page_title: "GitHub: github_enterprise_settings"
description: |-
  Creates and manages settings for a GitHub Enterprise.
---

# github_enterprise_settings

This resource allows you to create and manage settings for a GitHub Enterprise, including Actions permissions, workflow permissions, and security policies.
You must have admin access to an enterprise to use this resource.

## Example Usage

### Basic Configuration

```hcl
resource "github_enterprise_settings" "example" {
  enterprise_slug = "my-enterprise"
  
  actions_enabled_organizations = "all"
  actions_allowed_actions      = "all"
  
  default_workflow_permissions     = "read"
  can_approve_pull_request_reviews = false
}
```

### Advanced Configuration with Selective Permissions

```hcl
resource "github_enterprise_settings" "advanced" {
  enterprise_slug = "my-enterprise"
  
  # Only selected organizations can run actions
  actions_enabled_organizations = "selected"
  
  # Only allow specific actions
  actions_allowed_actions      = "selected"
  actions_github_owned_allowed = true
  actions_verified_allowed     = true
  actions_patterns_allowed = [
    "actions/cache@*",
    "actions/checkout@*",
    "my-org/custom-action@v1"
  ]
  
  # Workflow permissions
  default_workflow_permissions     = "write"
  can_approve_pull_request_reviews = true
}
```

## Argument Reference

The following arguments are supported:

* `enterprise_slug` - (Required) The slug of the enterprise.
* `actions_enabled_organizations` - (Optional) The policy that controls which organizations in the enterprise are allowed to run GitHub Actions. Can be one of: `all`, `none`, or `selected`. Defaults to `all`.
* `actions_allowed_actions` - (Optional) The permissions policy that controls the actions and reusable workflows that are allowed to run. Can be one of: `all`, `local_only`, or `selected`. Defaults to `all`.
* `actions_github_owned_allowed` - (Optional) Whether GitHub-owned actions are allowed. Only used when `actions_allowed_actions` is set to `selected`. Defaults to `false`.
* `actions_verified_allowed` - (Optional) Whether actions from verified creators are allowed. Only used when `actions_allowed_actions` is set to `selected`. Defaults to `false`.
* `actions_patterns_allowed` - (Optional) Specifies a list of string-matching patterns to allow specific action(s) and reusable workflow(s). Wildcards, tags, and SHAs are allowed. For example, `monalisa/octocat@*`, `monalisa/octocat@v2`, `monalisa/*`. Only used when `actions_allowed_actions` is set to `selected`.
* `default_workflow_permissions` - (Optional) The default permissions granted to the GITHUB_TOKEN when running workflows. Can be `read` (recommended) or `write`. Defaults to `read`.
* `can_approve_pull_request_reviews` - (Optional) Whether GitHub Actions can approve pull request reviews. Defaults to `false`.

## Attributes Reference

The following additional attributes are exported:

* `id` - The ID of the enterprise settings.

## Import

Enterprise settings can be imported using the enterprise slug:

```
$ terraform import github_enterprise_settings.example my-enterprise
```

## Notes

### Actions Policies

When `actions_allowed_actions` is set to `selected`, you can control which actions are allowed to run by configuring:

- `actions_github_owned_allowed`: Allow all GitHub-owned actions (like `actions/checkout`, `actions/upload-artifact`, etc.)
- `actions_verified_allowed`: Allow actions from verified creators in the GitHub Marketplace
- `actions_patterns_allowed`: Specify exact action patterns using wildcards and version constraints

### Security Considerations

For maximum security, consider:
- Setting `default_workflow_permissions` to `read` to limit GITHUB_TOKEN permissions
- Setting `can_approve_pull_request_reviews` to `false` to prevent automated approval bypasses
- Using `selected` for `actions_allowed_actions` to restrict which actions can run
- Regularly reviewing and updating `actions_patterns_allowed` patterns

### API Limitations

Some newer enterprise settings like fork pull request workflows from outside collaborators, artifact retention policies, and self-hosted runner permissions are not yet supported and will be added in future versions when the go-github dependency is updated.