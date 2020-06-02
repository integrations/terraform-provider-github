---
layout: "github"
page_title: "GitHub: github_ip_ranges"
description: |-
  Get information on GitHub's IP addresses.
---

# github_ip_ranges

Use this data source to retrieve information about GitHub's IP addresses.

## Example Usage

```hcl
data "github_ip_ranges" "test" {}
```

## Attributes Reference

 * `hooks` - An Array of IP addresses in CIDR format specifying the addresses that incoming service hooks will originate from.
 * `web` - An Array of IP addresses in CIDR format specifying the Web servers.
 * `api` - An Array of IP addresses in CIDR format specifying the API servers.
 * `git` - An Array of IP addresses in CIDR format specifying the Git servers.
 * `pages` - An Array of IP addresses in CIDR format specifying the A records for GitHub Pages.
 * `importer` - An Array of IP addresses in CIDR format specifying the A records for GitHub Importer.
