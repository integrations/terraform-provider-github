---
layout: "github"
page_title: "GitHub: github_organization"
description: |-
  Get an organization.
---

# github_organization

Use this data source to retrieve basic information about a GitHub Organization.

## Example Usage

```hcl
data "github_organization" "test" {
  name = "github"
}

locals {
  admins = data.github_organization.test.admins
}
```

## Attributes Reference

 * `plan` - The plan name for the organization account
 * `repositories` - (`list`) A list with the repositories on the organization
 * `admins` - (`list`) A list with the admins on the organization
 * `members` - (`list`) A list with the members on the organization