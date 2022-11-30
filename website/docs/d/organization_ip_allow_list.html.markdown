---
layout: "github"
page_title: "GitHub: github_organization_ip_allow_list"
description: |-
  Get the IP allow list of an organization.
---

# github_organization_ip_allow_list

Use this data source to retrieve information about the IP allow list of an organization.
The allow list for IP addresses will block access to private resources via the web, API,
and Git from any IP addresses that are not on the allow list.

## Example Usage

```hcl
data "github_organization_ip_allow_list" "all" {}
```

## Attributes Reference

* `ip_allow_list` - An Array of allowed IP addresses.
___

Each element in the `ip_allow_list` block consists of:

 * `id` - The ID of the IP allow list entry.
 * `name` - The name of the IP allow list entry.
 * `allow_list_value` - A single IP address or range of IP addresses in CIDR notation.
 * `is_active` -  Whether the entry is currently active.
 * `created_at` - Identifies the date and time when the object was created.
 * `updated_at` - Identifies the date and time when the object was last updated.
