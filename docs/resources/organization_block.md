---
page_title: "github_organization_block Resource - terraform-provider-github
description: |-
  Creates and manages blocks for GitHub organizations
---

# github_organization_block (Resource)

This resource allows you to create and manage blocks for GitHub organizations.

## Example Usage

```terraform
resource "github_organization_block" "example" {
  username = "paultyng"
}
```

## Argument Reference

The following arguments are supported:

- `username` - (Required) The name of the user to block.

## Import

GitHub organization block can be imported using a username, e.g.

```sh
$ terraform import github_github_organization_block.example someuser
```
