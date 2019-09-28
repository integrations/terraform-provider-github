---
layout: "github"
page_title: "GitHub: github_organization_block"
description: |-
  Creates and manages blocks for GitHub organizations
---

# github_organization_block

This resource allows you to create and manage blocks for GitHub organizations.

## Example Usage

```hcl
resource "github_organization_block" "example" {
  username = "paultyng"
}
```

## Argument Reference

The following arguments are supported:

* `username` - (Required) The name of the user to block.
