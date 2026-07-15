---
page_title: "github_code_security_configuration (Resource) - GitHub"
description: |-
  Manages a GitHub Code Security Configuration at the organization or enterprise level.
---

# github_code_security_configuration (Resource)

This resource allows you to create and manage [Code Security Configurations](https://docs.github.com/en/code-security/securing-your-organization/enabling-security-features-in-your-organization/about-enabling-security-features-at-scale) for a GitHub organization or enterprise. Code security configurations bundle security feature settings (Advanced Security, Dependabot, code scanning default setup, secret scanning, private vulnerability reporting) so they can be applied to repositories at scale.

By default the configuration is created at the organization level, using the organization configured as the provider `owner`. Set `enterprise_slug` to create the configuration at the enterprise level instead.

Organization-level usage requires an organization admin token. Enterprise-level usage requires enterprise admin access.

## Example Usage

```terraform
# Organization-level configuration, set as default for new private/internal
# repos and attached to all repositories without an existing configuration
resource "github_code_security_configuration" "org_baseline" {
  name        = "org-security-baseline"
  description = "Baseline security configuration for all repositories"

  advanced_security                     = "enabled"
  dependency_graph                      = "enabled"
  dependabot_alerts                     = "enabled"
  dependabot_security_updates           = "enabled"
  code_scanning_default_setup           = "enabled"
  secret_scanning                       = "enabled"
  secret_scanning_push_protection       = "enabled"
  secret_scanning_validity_checks       = "enabled"
  secret_scanning_non_provider_patterns = "disabled"
  private_vulnerability_reporting       = "enabled"
  enforcement                           = "enforced"

  default_for_new_repos = "private_and_internal"
  attach_scope          = "all_without_configurations"
}

# Enterprise-level configuration
resource "github_code_security_configuration" "enterprise_baseline" {
  enterprise_slug = "my-enterprise"
  name            = "enterprise-security-baseline"
  description     = "Enterprise-wide security baseline"

  dependabot_alerts               = "enabled"
  secret_scanning                 = "enabled"
  secret_scanning_push_protection = "enabled"

  default_for_new_repos = "all"
}
```

## Argument Reference

The following arguments are supported:

- `name` - (Required) The name of the code security configuration. Must be unique within the organization or enterprise.

- `description` - (Optional) A description of the code security configuration.

- `enterprise_slug` - (Optional) The slug of the enterprise to create the configuration in. If omitted, the configuration is created at the organization level using the provider's configured owner. Changing this forces a new resource.

- `advanced_security` - (Optional) The enablement status of GitHub Advanced Security. Can be `enabled`, `disabled` or `not_set`. Defaults to `disabled`.

- `dependency_graph` - (Optional) The enablement status of Dependency Graph. Can be `enabled`, `disabled` or `not_set`. Defaults to `enabled`.

- `dependabot_alerts` - (Optional) The enablement status of Dependabot alerts. Can be `enabled`, `disabled` or `not_set`. Defaults to `disabled`.

- `dependabot_security_updates` - (Optional) The enablement status of Dependabot security updates. Can be `enabled`, `disabled` or `not_set`. Defaults to `disabled`.

- `code_scanning_default_setup` - (Optional) The enablement status of code scanning default setup. Can be `enabled`, `disabled` or `not_set`. Defaults to `disabled`.

- `secret_scanning` - (Optional) The enablement status of secret scanning. Can be `enabled`, `disabled` or `not_set`. Defaults to `disabled`.

- `secret_scanning_push_protection` - (Optional) The enablement status of secret scanning push protection. Can be `enabled`, `disabled` or `not_set`. Defaults to `disabled`.

- `secret_scanning_validity_checks` - (Optional) The enablement status of secret scanning validity checks. Can be `enabled`, `disabled` or `not_set`. Defaults to `disabled`.

- `secret_scanning_non_provider_patterns` - (Optional) The enablement status of secret scanning non-provider patterns. Can be `enabled`, `disabled` or `not_set`. Defaults to `disabled`.

- `private_vulnerability_reporting` - (Optional) The enablement status of private vulnerability reporting. Can be `enabled`, `disabled` or `not_set`. Defaults to `disabled`.

- `enforcement` - (Optional) The enforcement status of the configuration. Can be `enforced` or `unenforced`. Defaults to `enforced`.

- `default_for_new_repos` - (Optional) Which types of new repositories this configuration should be applied to by default. Can be `all`, `none`, `private_and_internal` or `public`. If omitted, the configuration is not set as a default.

- `attach_scope` - (Optional) The scope of repositories to attach the configuration to. Can be `all` or `all_without_configurations`. The attach operation runs on create and whenever this value changes. The GitHub API does not expose the scope a configuration was attached with, so this value cannot be read back; removing the attribute does not detach repositories.

## Attributes Reference

- `id` - The ID of the code security configuration.

- `configuration_id` - The numeric ID of the code security configuration.

- `target_type` - The target type of the configuration (`organization` or `enterprise`).

- `html_url` - The URL of the configuration in the GitHub UI.

## Import

Organization-level code security configurations can be imported using the configuration ID:

```shell
terraform import github_code_security_configuration.org_baseline 1234
```

Enterprise-level configurations use `<enterprise_slug>:<configuration_id>`:

```shell
terraform import github_code_security_configuration.enterprise_baseline my-enterprise:1234
```

~> **Note** `attach_scope` cannot be recovered on import and will need to be re-specified in configuration if desired.
