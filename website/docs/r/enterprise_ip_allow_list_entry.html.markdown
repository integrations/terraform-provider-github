---
layout: "github"
page_title: "GitHub: github_enterprise_ip_allow_list_entry"
description: |-
  Creates and manages IP allow list entries within a GitHub Enterprise
---

# github_enterprise_ip_allow_list_entry

This resource allows you to create and manage IP allow list entries for a GitHub Enterprise account. IP allow list entries define IP addresses or ranges that are permitted to access private resources in the enterprise.

## Example Usage

```hcl
resource "github_enterprise_ip_allow_list_entry" "test" {
  enterprise_slug = "my-enterprise"
  ip              = "192.168.1.0/20"
  name            = "My IP Range Name"
  is_active       = true
}
```

## Argument Reference

The following arguments are supported:

* `enterprise_slug` - (Required) The slug of the enterprise.
* `ip`              - (Required) An IP address or range of IP addresses in CIDR notation.
* `name`            - (Optional) A descriptive name for the IP allow list entry.
* `is_active`       - (Optional) Whether the entry is currently active. Default: true.

## Import

This resource can be imported using the enterprise slug and ID of the IP allow list entry:

```bash
$ terraform import github_enterprise_ip_allow_list_entry.test enterprise-slug:IALE_kwHOC1234567890a
```