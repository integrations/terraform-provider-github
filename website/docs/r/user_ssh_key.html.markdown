---
layout: "github"
page_title: "GitHub: github_user_ssh_key"
description: |-
  Provides a GitHub user's SSH key resource.
---

# github_user_ssh_key

Provides a GitHub user's SSH key resource.

This resource allows you to add/remove SSH keys from your user account.

## Example Usage

```hcl
resource "github_user_ssh_key" "example" {
  title = "example title"
  key   = file("~/.ssh/id_rsa.pub")
}
```

## Argument Reference

The following arguments are supported:

* `title` - (Required) A descriptive name for the new key. e.g. `Personal MacBook Air`
* `key` - (Required) The public SSH key to add to your GitHub account.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the SSH key
* `url` - The URL of the SSH key

## Import

SSH keys can be imported using their ID e.g.

```
$ terraform import github_user_ssh_key.example 1234567
```
