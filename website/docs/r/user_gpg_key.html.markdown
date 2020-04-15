---
layout: "github"
page_title: "GitHub: github_user_gpg_key"
description: |-
  Provides a GitHub user's GPG key resource.
---

# github_user_gpg_key

Provides a GitHub user's GPG key resource.

This resource allows you to add/remove GPG keys from your user account.

## Example Usage

```hcl
resource "github_user_gpg_key" "example" {
  armored_public_key = "-----BEGIN PGP PUBLIC KEY BLOCK-----\n...\n-----END PGP PUBLIC KEY BLOCK-----"
}
```

## Argument Reference

The following arguments are supported:

* `armored_public_key` - (Required) Your public GPG key, generated in ASCII-armored format.
  See [Generating a new GPG key](https://help.github.com/articles/generating-a-new-gpg-key/) for help on creating a GPG key.

## Attributes Reference

The following attributes are exported:

* `id` - The GitHub ID of the GPG key, e.g. `401586`
* `key_id` - The key ID of the GPG key, e.g. `3262EFF25BA0D270`

## Import

GPG keys are not importable due to the fact that [API](https://developer.github.com/v3/users/gpg_keys/#gpg-keys)
does not return previously uploaded GPG key.
