---
page_title: "github_organization_ip_allow_list_entry (Resource) - GitHub"
description: |-
  Creates and manages IP allow list entries within a GitHub Organization
---

# github_organization_ip_allow_list_entry (Resource)

This resource allows you to create and manage IP allow list entries for a GitHub Organization. IP allow list entries define IP addresses or ranges that are permitted to access private and internal resources owned by the organization.

The organization is taken from the `owner` configured on the provider. The organization must be on a GitHub Enterprise Cloud plan, and the IP allow list itself must be enabled in the organization's security settings before entries can be created.

## Example Usage

```terraform
resource "github_organization_ip_allow_list_entry" "test" {
  ip        = "192.168.1.0/20"
  name      = "My IP Range Name"
  is_active = true
}
```

## Argument Reference

The following arguments are supported:

- `ip` - (Required) An IP address or range of IP addresses in CIDR notation.
- `name` - (Optional) A descriptive name for the IP allow list entry.
- `is_active` - (Optional) Whether the entry is currently active. Default: true.

## Import

This resource can be imported using the ID of the IP allow list entry:

```shell
terraform import github_organization_ip_allow_list_entry.test IALE_kwHOC1234567890a
```
