---
layout: "github"
page_title: "GitHub: github_dependabot_secret"
description: |-
  Creates and manages an Dependabot Secret within a GitHub repository
---

# github_dependabot_secret

This resource allows you to create and manage GitHub Dependabot secrets within your GitHub repositories.
You must have write access to a repository to use this resource.

Secret values are encrypted using the [Go '/crypto/box' module](https://godoc.org/golang.org/x/crypto/nacl/box) which is
interoperable with [libsodium](https://libsodium.gitbook.io/doc/). Libsodium is used by GitHub to decrypt secret values.

For the purposes of security, the contents of the `value` field have been marked as `sensitive` to Terraform,
but it is important to note that **this does not hide it from state files**. You should treat state as sensitive always.
It is also advised that you do not store plaintext values in your code but rather populate the `value_encrypted`
using fields from a resource, data source or variable as, while encrypted in state, these will be easily accessible
in your code. See below for an example of this abstraction.

## Example Usage

```hcl
resource "github_dependabot_secret" "example_plaintext" {
  repository  = "example_repository"
  secret_name = "example_secret_name"
  value       = var.some_secret_string
}

resource "github_dependabot_secret" "example_encrypted" {
  repository       = "example_repository"
  secret_name      = "example_secret_name"
  value_encrypted  = var.some_encrypted_secret_string
}
```

## Example Lifecycle Ignore Changes

This resource supports using the `lifecycle` `ignore_changes` block on `remote_updated_at` to support use cases where a secret value is created using a placeholder value and then modified after creation outside the scope of Terraform. This approach ensures only the initial placeholder value is referenced in your code and in the resulting state file.

```hcl
resource "github_dependabot_secret" "example_allow_drift" {
  repository  = "example_repository"
  secret_name = "example_secret_name"
  value       = "placeholder"

  lifecycle {
    ignore_changes = [remote_updated_at]
  }
}
```

## Argument Reference

The following arguments are supported:

- `repository` - (Required) Name of the repository.
- `secret_name` - (Required) Name of the secret.
- `key_id` - (Optional) ID of the public key used to encrypt the secret, required when setting `encrypted_value`.
- `value` - (Optional) Plaintext value of the secret to be encrypted. This conflicts with `value_encrypted`, `encrypted_value` & `plaintext_value`.
- `value_encrypted` - (Optional) Encrypted value of the secret using the GitHub public key in Base64 format, `key_id` is required with this value. This conflicts with `value`, `encrypted_value` & `plaintext_value`.
- `encrypted_value` - (**DEPRECATED**)(Optional) Please use `value_encrypted`.
- `plaintext_value` - (**DEPRECATED**)(Optional) Please use `value`.

~> **Note**: One of either `value`, `value_encrypted`, `encrypted_value`, or `plaintext_value` must be specified.

## Attributes Reference

- `repository_id` - ID of the repository.
- `created_at` - Date the secret was created.
- `updated_at` - Date the secret was last updated by the provider.
- `remote_updated_at` - Date the secret was last updated in GitHub.

## Import

This resource can be imported using an ID made of the repository name, and secret name separated by a `:`.

~> **Note**: When importing secrets, the `value`, `value_encrypted`, `encrypted_value`, or `plaintext_value` fields will not be populated in the state. You may need to ignore changes for these as a workaround if you're not planning on updating the secret through Terraform.

### Import Block

The following import imports a GitHub Dependabot secret named `mysecret` for the repo `myrepo` to a `github_dependabot_secret` resource named `example`.

```hcl
import {
  to = github_dependabot_secret.example
  id = "myrepo:mysecret"
}
```

### Import Command

The following command imports a GitHub Dependabot secret named `mysecret` for the repo `myrepo` to a `github_dependabot_secret` resource named `example`.

```shell
terraform import github_dependabot_secret.example myrepo:mysecret
```
