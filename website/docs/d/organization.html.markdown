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

 * `plan` - The plan name for the organization account
