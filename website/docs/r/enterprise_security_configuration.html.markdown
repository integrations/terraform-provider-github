---
layout: "github"
page_title: "GitHub: github_enterprise_security_configuration"
description: |-
  Manages a code security configuration for a GitHub Enterprise.
---

# github_enterprise_security_configuration

This resource allows you to create and manage code security configurations for a GitHub Enterprise.

## Example Usage

```hcl
resource "github_enterprise_security_configuration" "default" {
  enterprise_slug                 = "my-enterprise"
  name                            = "default-config"
  description                     = "Default security configuration"
  advanced_security               = "enabled"
  dependency_graph                = "enabled"
  dependabot_alerts               = "enabled"
  dependabot_security_updates     = "enabled"
  code_scanning_default_setup     = "enabled"
  secret_scanning                 = "enabled"
  secret_scanning_push_protection = "enabled"
  private_vulnerability_reporting = "enabled"
  enforcement                     = "enforced"
}
```

## Argument Reference

The following arguments are supported:

* `enterprise_slug` - (Required) The slug of the enterprise. Changing this forces a new resource to be created.
* `name` - (Required) The name of the code security configuration.
* `description` - (Required) A description of the code security configuration.
* `advanced_security` - (Optional) The advanced security configuration. Can be one of `enabled`, `disabled`.
* `dependency_graph` - (Optional) The dependency graph configuration. Can be one of `enabled`, `disabled`, `not_set`.
* `dependency_graph_autosubmit_action` - (Optional) The dependency graph autosubmit action configuration. Can be one of `enabled`, `disabled`, `not_set`.
* `dependency_graph_autosubmit_action_options` - (Optional) The dependency graph autosubmit action options. See [Dependency Graph Autosubmit Action Options](#dependency-graph-autosubmit-action-options) below for details.
* `dependabot_alerts` - (Optional) The dependabot alerts configuration. Can be one of `enabled`, `disabled`, `not_set`.
* `dependabot_security_updates` - (Optional) The dependabot security updates configuration. Can be one of `enabled`, `disabled`, `not_set`.
* `code_scanning_default_setup` - (Optional) The code scanning default setup configuration. Can be one of `enabled`, `disabled`, `not_set`.
* `code_scanning_default_setup_options` - (Optional) The code scanning default setup options. See [Code Scanning Default Setup Options](#code-scanning-default-setup-options) below for details.
* `code_scanning_options` - (Optional) The code scanning options. See [Code Scanning Options](#code-scanning-options) below for details.
* `code_security` - (Optional) The code security configuration. Can be one of `enabled`, `disabled`, `not_set`.
* `secret_scanning` - (Optional) The secret scanning configuration. Can be one of `enabled`, `disabled`, `not_set`.
* `secret_scanning_push_protection` - (Optional) The secret scanning push protection configuration. Can be one of `enabled`, `disabled`, `not_set`.
* `secret_scanning_validity_checks` - (Optional) The secret scanning validity checks configuration. Can be one of `enabled`, `disabled`, `not_set`.
* `secret_scanning_non_provider_patterns` - (Optional) The secret scanning non provider patterns configuration. Can be one of `enabled`, `disabled`, `not_set`.
* `secret_scanning_generic_secrets` - (Optional) The secret scanning generic secrets configuration. Can be one of `enabled`, `disabled`, `not_set`.
* `secret_protection` - (Optional) The secret protection configuration. Can be one of `enabled`, `disabled`, `not_set`.
* `private_vulnerability_reporting` - (Optional) The private vulnerability reporting configuration. Can be one of `enabled`, `disabled`, `not_set`.
* `enforcement` - (Optional) The enforcement configuration. Can be one of `enforced`, `unenforced`.

## Attributes Reference

* `target_type` - The target type of the code security configuration.

### Dependency Graph Autosubmit Action Options

The `dependency_graph_autosubmit_action_options` block supports:

* `labeled_runners` - (Optional) Whether to use labeled runners for the dependency graph autosubmit action.

### Code Scanning Default Setup Options

The `code_scanning_default_setup_options` block supports:

* `runner_type` - (Optional) The type of runner to use for code scanning default setup. Can be one of `standard`, `labeled`.
* `runner_label` - (Optional) The label of the runner to use for code scanning default setup.

### Code Scanning Options

The `code_scanning_options` block supports:

* `allow_advanced` - (Optional) Whether to allow advanced security for code scanning.

## Import

GitHub Enterprise Code Security Configurations can be imported using the enterprise slug and the configuration ID separated by a colon, e.g.

```text
$ terraform import github_enterprise_security_configuration.example my-enterprise:123
```
