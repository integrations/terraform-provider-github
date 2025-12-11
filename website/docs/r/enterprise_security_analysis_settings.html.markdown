---
layout: "github"
page_title: "GitHub: github_enterprise_security_analysis_settings"
description: |-
  Manages GitHub Enterprise security analysis settings.
---

# github_enterprise_security_analysis_settings

This resource allows you to manage code security and analysis settings for a GitHub Enterprise account. This controls Advanced Security, Secret Scanning, and related security features that are automatically enabled for new repositories in the enterprise.

You must have enterprise admin access to use this resource.

## Example Usage

```hcl
# Basic security settings - enable secret scanning only
resource "github_enterprise_security_analysis_settings" "basic" {
  enterprise_slug = "my-enterprise"
  
  secret_scanning_enabled_for_new_repositories = true
}

# Full security configuration with all features enabled
resource "github_enterprise_security_analysis_settings" "comprehensive" {
  enterprise_slug = "my-enterprise"
  
  advanced_security_enabled_for_new_repositories             = true
  secret_scanning_enabled_for_new_repositories               = true
  secret_scanning_push_protection_enabled_for_new_repositories = true
  secret_scanning_validity_checks_enabled                   = true
  secret_scanning_push_protection_custom_link               = "https://octokit.com/security-guidelines"
}
```

## Argument Reference

The following arguments are supported:

* `enterprise_slug` - (Required) The slug of the enterprise.

* `advanced_security_enabled_for_new_repositories` - (Optional) Whether GitHub Advanced Security is automatically enabled for new repositories. Defaults to `false`. Requires Advanced Security license.

* `secret_scanning_enabled_for_new_repositories` - (Optional) Whether secret scanning is automatically enabled for new repositories. Defaults to `false`.

* `secret_scanning_push_protection_enabled_for_new_repositories` - (Optional) Whether secret scanning push protection is automatically enabled for new repositories. Defaults to `false`.

* `secret_scanning_push_protection_custom_link` - (Optional) Custom URL for secret scanning push protection bypass instructions.

* `secret_scanning_validity_checks_enabled` - (Optional) Whether secret scanning validity checks are enabled. Defaults to `false`.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The enterprise slug.

## Import

Enterprise security analysis settings can be imported using the enterprise slug:

```
terraform import github_enterprise_security_analysis_settings.example my-enterprise
```

## Notes

~> **Note:** This resource requires a GitHub Enterprise account and enterprise admin permissions.

~> **Note:** Advanced Security features require a GitHub Advanced Security license.

When this resource is destroyed, all security analysis settings will be reset to disabled defaults for security reasons.

## Dependencies

This resource manages the following security features:

- **Advanced Security**: Code scanning, secret scanning, and dependency review
- **Secret Scanning**: Automatic detection of secrets in code
- **Push Protection**: Prevents secrets from being committed to repositories
- **Validity Checks**: Verifies that detected secrets are actually valid

These settings only apply to **new repositories** created after the settings are enabled. Existing repositories are not affected and must be configured individually.