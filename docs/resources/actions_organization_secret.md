---
page_title: "github_actions_organization_secret (Resource) - GitHub"
description: |-
  Creates and manages an Action Secret within a GitHub organization
---

# github_actions_organization_secret (Resource)

This resource allows you to create and manage GitHub Actions secrets within your GitHub organization. You must have write access to a repository to use this resource.

Secret values are encrypted using the [Go '/crypto/box' module](https://godoc.org/golang.org/x/crypto/nacl/box) which is interoperable with [libsodium](https://libsodium.gitbook.io/doc/). Libsodium is used by GitHub to decrypt secret values.

For the purposes of security, the contents of the `plaintext_value` field have been marked as `sensitive` to Terraform, but it is important to note that **this does not hide it from state files**. You should treat state as sensitive always. It is also advised that you do not store plaintext values in your code but rather populate the `encrypted_value` using fields from a resource, data source or variable as, while encrypted in state, these will be easily accessible in your code. See below for an example of this abstraction.

## Example Usage

```terraform
resource "github_actions_organization_secret" "example_plaintext" {
  secret_name     = "example_secret_name"
  visibility      = "all"
  plaintext_value = var.some_secret_string
}

resource "github_actions_organization_secret" "example_encrypted" {
  secret_name     = "example_secret_name"
  visibility      = "all"
  encrypted_value = var.some_encrypted_secret_string
}
```

```terraform
data "github_repository" "repo" {
  full_name = "my-org/repo"
}

resource "github_actions_organization_secret" "example_encrypted" {
  secret_name             = "example_secret_name"
  visibility              = "selected"
  plaintext_value         = var.some_secret_string
  selected_repository_ids = [data.github_repository.repo.repo_id]
}

resource "github_actions_organization_secret" "example_secret" {
  secret_name             = "example_secret_name"
  visibility              = "selected"
  encrypted_value         = var.some_encrypted_secret_string
  selected_repository_ids = [data.github_repository.repo.repo_id]
}
```

## Example Lifecycle Ignore Changes

This resource supports using the `lifecycle` `ignore_changes` block on `remote_updated_at` to support use cases where a secret value is created using a placeholder value and then modified after creation outside the scope of Terraform. This approach ensures only the initial placeholder value is referenced in your code and in the resulting state file.

```terraform
resource "github_actions_organization_secret" "example_allow_drift" {
  secret_name     = "example_secret_name"
  visibility      = "all"
  plaintext_value = "placeholder"

  lifecycle {
    ignore_changes = [remote_updated_at]
  }
}
```

## Argument Reference

The following arguments are supported:

- `secret_name` - (Required) Name of the secret.
- `key_id` - (Optional) ID of the public key used to encrypt the secret. This should be provided when setting `encrypted_value`; if it isn't then the current public key will be looked up, which could cause a missmatch. This conflicts with `plaintext_value`.
- `encrypted_value` - (Optional) Encrypted value of the secret using the GitHub public key in Base64 format.
- `plaintext_value` - (Optional) Plaintext value of the secret to be encrypted.
- `visibility` - (Required) Configures the access that repositories have to the organization secret; must be one of `all`, `private`, or `selected`.
- `selected_repository_ids` - (Optional) An array of repository IDs that can access the organization variable; this requires `visibility` to be set to `selected`.
- `destroy_on_drift` - (**DEPRECATED**) (Optional) This is ignored as drift detection is built into the resource.

~> **Note**: One of either `encrypted_value` or `plaintext_value` must be specified.

## Attributes Reference

- `created_at` - Date the secret was created.
- `updated_at` - Date the secret was last updated by the provider.
- `remote_updated_at` - Date the secret was last updated in GitHub.

## Import

This resource can be imported using the secret name as the ID.

~> **Note**: When importing secrets, the `plaintext_value` or `encrypted_value` fields will not be populated in the state. You may need to ignore changes for these as a workaround if you're not planning on updating the secret through Terraform.

### Import Block

The following import imports a GitHub actions organization secret named `mysecret` to a `github_actions_organization_secret` resource named `example`.

```terraform
import {
  to = github_actions_organization_secret.example
  id = "mysecret"
}
```

### Import Command

The following command imports a GitHub actions organization secret named `mysecret` to a `github_actions_organization_secret` resource named `example`.

```shell
terraform import github_actions_organization_secret.example mysecret
```
