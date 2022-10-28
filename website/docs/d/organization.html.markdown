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
data "github_organization" "example" {
  name = "github"
}
```

## Attributes Reference

 * `id` - The ID of the organization
 * `node_id` - GraphQL global node id for use with v4 API
 * `name` - The organization's public profile name
 * `orgname` - The organization's name as used in URLs and the API
 * `login` - The login of the organization account
 * `description` - The description the organization account
 * `plan` - The plan name for the organization account
 * `repositories` - (`list`) A list with the repositories on the organization
 * `members` - (`list`) A list with the members of the organization
