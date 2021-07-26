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
```

## Attributes Reference

 * `name` - The name of the organization account
 * `login` - The login of the organization account
 * `description` - The description the organization account
 * `plan` - The plan name for the organization account
 * `repositories` - (`list`) A list with the repositories on the organization
 * `members` - (`list`) A list with the members of the organization
