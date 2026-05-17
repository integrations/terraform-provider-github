---
layout: "github"
page_title: "GitHub: github_enterprise_private_repository_forking_setting"
description: |-
  Creates and manages the private repository forking policy for a GitHub Enterprise.
---

# github_enterprise_private_repository_forking_setting

This resource allows you to create and manage the private repository forking policy for a GitHub Enterprise.
You must have enterprise admin access to use this resource.

When `setting` is `ENABLED`, the `policy` attribute controls where forks
can be created. When `DISABLED`, forking of private repositories is not allowed.
When `NO_POLICY`, individual organizations within the enterprise control their own
forking settings.

## Example Usage

### Restrict forking to same organization only

```hcl
resource "github_enterprise_private_repository_forking_setting" "example" {
  enterprise_slug = "my-enterprise"
  setting   = "ENABLED"
  policy    = "SAME_ORGANIZATION"
}
```

### Allow forking to enterprise-managed user accounts or enterprise organizations

```hcl
resource "github_enterprise_private_repository_forking_setting" "example" {
  enterprise_slug = "my-enterprise"
  setting   = "ENABLED"
  policy    = "ENTERPRISE_ORGANIZATIONS_USER_ACCOUNTS"
}
```

### Disable private repository forking entirely

```hcl
resource "github_enterprise_private_repository_forking_setting" "example" {
  enterprise_slug = "my-enterprise"
  setting   = "DISABLED"
}
```

### Allow organizations to set their own policy

```hcl
resource "github_enterprise_private_repository_forking_setting" "example" {
  enterprise_slug = "my-enterprise"
  setting   = "NO_POLICY"
}
```

## Argument Reference

The following arguments are supported:

* `enterprise_slug` - (Required) The slug of the enterprise.
* `setting` - (Required) Whether private repository forking is enabled for the enterprise. Must be one of `ENABLED`, `DISABLED`, or `NO_POLICY`.
* `policy` - (Optional) Where members can fork private repositories. Required when `setting` is `ENABLED`. Must be one of:
  * `ENTERPRISE_ORGANIZATIONS` - Members can fork to an organization within this enterprise.
  * `SAME_ORGANIZATION` - Members can fork only within the same organization (intra-org).
  * `SAME_ORGANIZATION_USER_ACCOUNTS` - Members can fork to their user account or within the same organization.
  * `ENTERPRISE_ORGANIZATIONS_USER_ACCOUNTS` - Members can fork to their enterprise-managed user account or an organization inside this enterprise.
  * `USER_ACCOUNTS` - Members can fork to their user account.
  * `EVERYWHERE` - Members can fork to their user account or an organization, either inside or outside of this enterprise.

**Note:** Destroying this resource sets the enterprise policy to `NO_POLICY`, which allows individual organizations to control their own forking settings. It does not set the policy to `DISABLED`.

## Attributes Reference

No additional attributes are exported.

## Import

Enterprise private repository forking settings can be imported using the enterprise slug:

```
$ terraform import github_enterprise_private_repository_forking_setting.example my-enterprise
```
