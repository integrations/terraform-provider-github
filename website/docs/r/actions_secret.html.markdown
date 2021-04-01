---
layout: "github"
page_title: "GitHub: github_actions_secret"
description: |-
  Creates and manages an Action Secret within a GitHub repository
---

# github_actions_secret

This resource allows you to create and manage GitHub Actions secrets within your GitHub repositories.
You must have write access to a repository to use this resource.

Secret values are encrypted using the [Go '/crypto/box' module](https://godoc.org/golang.org/x/crypto/nacl/box) which is
interoperable with [libsodium](https://libsodium.gitbook.io/doc/). Libsodium is used by GitHub to decrypt secret values.

For the purposes of security, the contents of the `plaintext_value` field have been marked as `sensitive` to Terraform,
but it is important to note that **this does not hide it from state files**. You should treat state as sensitive always.
It is also advised that you do not store plaintext values in your code but rather populate the `plaintext_value`
using fields from a resource, data source or variable as, while encrypted in state, these will be easily accessible
in your code. See below for an example of this abstraction.

## Example Usage

```hcl
data "github_actions_public_key" "example_public_key" {
  repository = "example_repository"
}

resource "github_actions_secret" "example_secret" {
  repository       = "example_repository"
  secret_name      = "example_secret_name"
  plaintext_value  = var.some_secret_string
}
```

## Argument Reference

The following arguments are supported:

* `repository`      - (Required) Name of the repository
* `secret_name`     - (Required) Name of the secret
* `plaintext_value` - (Required) Plaintext value of the secret to be encrypted

## Attributes Reference

* `created_at`      - Date of actions_secret creation.
* `updated_at`      - Date of actions_secret update.
