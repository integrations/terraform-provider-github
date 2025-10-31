# GitHub Enterprise Settings Example

This example demonstrates how to configure GitHub Enterprise settings using the Terraform GitHub provider.

## Overview

The `github_enterprise_settings` resource allows you to manage various enterprise-level settings for a GitHub Enterprise account, including:

- Actions permissions (which organizations can run GitHub Actions)
- Allowed actions policies (which actions are allowed to run)
- Workflow permissions (default permissions for GITHUB_TOKEN)
- Pull request review approval settings

## Requirements

- GitHub Enterprise account
- Personal access token with enterprise admin permissions
- Terraform >= 0.14

## Usage

1. Set your environment variables:

```bash
export TF_VAR_github_token="your_github_token"
export TF_VAR_enterprise_slug="your-enterprise-slug"
```

2. Initialize and apply:

```bash
terraform init
terraform plan
terraform apply
```

## Examples

### Basic Configuration

```terraform
resource "github_enterprise_settings" "basic" {
  enterprise_slug = "my-enterprise"

  actions_enabled_organizations = "all"
  actions_allowed_actions       = "all"
  
  default_workflow_permissions     = "read"
  can_approve_pull_request_reviews = false
}
```

### Advanced Configuration with Selective Permissions

```terraform
resource "github_enterprise_settings" "advanced" {
  enterprise_slug = "my-enterprise"

  # Only selected organizations can run actions
  actions_enabled_organizations = "selected"
  
  # Only allow specific actions
  actions_allowed_actions = "selected"
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

## Configuration Reference

### Actions Settings

- **`actions_enabled_organizations`**: Controls which organizations can run GitHub Actions
  - `"all"` - All organizations in the enterprise
  - `"none"` - No organizations
  - `"selected"` - Only specified organizations (requires additional configuration)

- **`actions_allowed_actions`**: Controls which actions can be run
  - `"all"` - All actions and reusable workflows
  - `"local_only"` - Only actions and workflows in the same repository/organization
  - `"selected"` - Only specified actions (requires additional configuration)

When `actions_allowed_actions` is set to `"selected"`, you can specify:

- **`actions_github_owned_allowed`**: Allow GitHub-owned actions (e.g., `actions/checkout`)
- **`actions_verified_allowed`**: Allow verified Marketplace actions
- **`actions_patterns_allowed`**: List of specific action patterns to allow

### Workflow Settings

- **`default_workflow_permissions`**: Default permissions for the GITHUB_TOKEN
  - `"read"` - Read-only permissions (recommended for security)
  - `"write"` - Read and write permissions

- **`can_approve_pull_request_reviews`**: Whether GitHub Actions can approve pull request reviews
  - `true` - Actions can approve PR reviews  
  - `false` - Actions cannot approve PR reviews (recommended for security)

## Security Considerations

1. **Workflow Permissions**: Use `"read"` permissions by default and grant `"write"` only when necessary
2. **PR Approvals**: Disable `can_approve_pull_request_reviews` to prevent automated approval bypasses
3. **Action Restrictions**: Use `"selected"` for `actions_allowed_actions` to limit which actions can run
4. **Token Security**: Store your GitHub token securely and use environment variables

## Limitations

This resource currently supports a subset of enterprise settings available through the GitHub API. Additional settings like fork PR workflows, artifact retention, and self-hosted runner permissions are not yet supported by the go-github version used in this provider and will be added in future versions.

## Import

You can import existing enterprise settings:

```bash
terraform import github_enterprise_settings.example my-enterprise
```

## Troubleshooting

### Common Issues

1. **Authentication**: Ensure your token has enterprise admin permissions
2. **Enterprise Access**: Verify you have access to the specified enterprise
3. **API Limits**: GitHub API has rate limits; consider adding delays for large configurations

### Verification

After applying, verify settings in the GitHub Enterprise dashboard:
1. Go to your enterprise settings
2. Navigate to "Policies" > "Actions"  
3. Check that the configured settings match your Terraform configuration