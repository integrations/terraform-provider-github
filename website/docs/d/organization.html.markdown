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
 * `node_id` - GraphQL global node ID for use with the v4 API
 * `name` - The organization's public profile name
 * `orgname` - The organization's name as used in URLs and the API
 * `login` - The organization account login
 * `description` - The organization account description
 * `plan` - The organization account plan name
 * `repositories` - (`list`) A list of the full names of the repositories in the organization formatted as `owner/name` strings
 * `members` - **Deprecated**: use `users` instead by replacing `github_organization.example.members` to `github_organization.example.users[*].login` which will give you the same value, expect this field to be removed in next major version
 * `users` - (`list`) A list with the members of the organization with following fields:
   * `id` - The ID of the member
   * `login` - The members login
   * `email` - Publicly available email
   * `role` - Member role `ADMIN`, `MEMBER`
