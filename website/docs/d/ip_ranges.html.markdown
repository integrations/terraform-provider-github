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

 * `actions` - An array of IP addresses in CIDR format specifying the addresses that incoming requests from GitHub actions will originate from.
 * `dependabot` - An array of IP addresses in CIDR format specifying the A records for dependabot.
 * `hooks` - An Array of IP addresses in CIDR format specifying the addresses that incoming service hooks will originate from.
 * `git` - An Array of IP addresses in CIDR format specifying the Git servers.
 * `pages` - An Array of IP addresses in CIDR format specifying the A records for GitHub Pages.
 * `importer` - An Array of IP addresses in CIDR format specifying the A records for GitHub Importer.
