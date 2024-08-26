---
page_title: "github_user_ssh_signing_key (Resource) - GitHub"
description: |-
  Provides a GitHub user's SSH signing key resource.
---

# github_user_ssh_signing_key (Resource)

Provides a GitHub user's SSH signing key resource.

This resource allows you to add/remove SSH signing keys from your user account.

## Example Usage

```terraform
resource "github_user_ssh_signing_key" "example" {
  title = "example title"
  key   = file("~/.ssh/id_rsa.pub")
}
```

## Argument Reference

The following arguments are supported:

- `title` - (Required) A descriptive name for the new key.
- `key` - (Required) The public SSH signing key to add to your GitHub account.

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the SSH signing key
- `key_id` - The unique identifier of the SSH signing key.
- `etag`

## Import

SSH signing keys can be imported using their ID e.g.

```shell
terraform import github_user_ssh_signing_key.example 1234567
```
