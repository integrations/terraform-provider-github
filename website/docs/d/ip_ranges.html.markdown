---
layout: "github"
page_title: "Github: github_ip_ranges"
sidebar_current: "docs-github-datasource-ip-ranges"
description: |-
  Get information on a GitHub's IP addresses.
---

# github_ip_ranges

Use this data source to retrieve information about a GitHub's IP addresses.
## Example Usage

```
data "github_ip_ranges" "test" {}
```

## Attributes Reference

 * `hooks` - An Array of IP addresses in CIDR format specifying the addresses that incoming service hooks will originate from.
 * `git` - An Array of IP addresses in CIDR format specifying the Git servers.
 * `pages` - An Array of IP addresses in CIDR format specifying the A records for GitHub Pages.
