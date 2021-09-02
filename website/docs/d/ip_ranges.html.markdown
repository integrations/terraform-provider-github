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
   * `actions_ipv4` - A subset of the `actions` array that contains IP addresses in IPv4 CIDR format.
   * `actions_ipv6` - A subset of the `actions` array that contains IP addresses in IPv6 CIDR format.
 * `dependabot` - An array of IP addresses in CIDR format specifying the A records for dependabot.
   * `dependabot_ipv4` - A subset of the `dependabot` array that contains IP addresses in IPv4 CIDR format.
   * `dependabot_ipv6` - A subset of the `dependabot` array that contains IP addresses in IPv6 CIDR format.
 * `hooks` - An Array of IP addresses in CIDR format specifying the addresses that incoming service hooks will originate from.
   * `hooks_ipv4` - A subset of the `hooks` array that contains IP addresses in IPv4 CIDR format.
   * `hooks_ipv6` - A subset of the `hooks` array that contains IP addresses in IPv6 CIDR format.
 * `git` - An Array of IP addresses in CIDR format specifying the Git servers.
   * `git_ipv4` - A subset of the `git` array that contains IP addresses in IPv4 CIDR format.
   * `git_ipv6` - A subset of the `git` array that contains IP addresses in IPv6 CIDR format.
 * `pages` - An Array of IP addresses in CIDR format specifying the A records for GitHub Pages.
   * `pages_ipv4` - A subset of the `pages` array that contains IP addresses in IPv4 CIDR format.
   * `pages_ipv6` - A subset of the `pages` array that contains IP addresses in IPv6 CIDR format.
 * `importer` - An Array of IP addresses in CIDR format specifying the A records for GitHub Importer.
   * `importer_ipv4` - A subset of the `importer` array that contains IP addresses in IPv4 CIDR format.
   * `importer_ipv6` - A subset of the `importer` array that contains IP addresses in IPv6 CIDR format.
