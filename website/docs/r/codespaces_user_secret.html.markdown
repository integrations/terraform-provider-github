---
layout: "github"
page_title: "GitHub: github_codespaces_user_secret"
description: |-
  Creates and manages an Codespaces Secret within a GitHub user
---

# github_codespaces_user_secret

This resource allows you to create and manage GitHub Codespaces secrets within your GitHub user.
You must have write access to a repository to use this resource.

Secret values are encrypted using the [Go '/crypto/box' module](https://godoc.org/golang.org/x/crypto/nacl/box) which is
interoperable with [libsodium](https://libsodium.gitbook.io/doc/). Libsodium is used by GitHub to decrypt secret values.

For the purposes of security, the contents of the `plaintext_value` field have been marked as `sensitive` to Terraform,
but it is important to note that **this does not hide it from state files**. You should treat state as sensitive always.
It is also advised that you do not store plaintext values in your code but rather populate the `encrypted_value`
using fields from a resource, data source or variable as, while encrypted in state, these will be easily accessible
in your code. See below for an example of this abstraction.

## Example Usage

```hcl
data "github_repository" "repo" {
  full_name = "my-org/repo"
}

resource "github_codespaces_user_secret" "example_secret" {
  secret_name             = "example_secret_name"
  plaintext_value         = var.some_secret_string
  selected_repository_ids = [data.github_repository.repo.repo_id]
}

resource "github_codespaces_user_secret" "example_secret" {
  secret_name             = "example_secret_name"
  encrypted_value         = var.some_encrypted_secret_string
  selected_repository_ids = [data.github_repository.repo.repo_id]
}
```

## Argument Reference

The following arguments are supported:

* `secret_name`             - (Required) Name of the secret
* `encrypted_value`         - (Optional) Encrypted value of the secret using the GitHub public key in Base64 format.
* `plaintext_value`         - (Optional) Plaintext value of the secret to be encrypted
* `selected_repository_ids` - (Optional) An array of repository ids that can access the user secret.

## Attributes Reference

* `created_at`      - Date of codespaces_secret creation.
* `updated_at`      - Date of codespaces_secret update.

## Import

This resource can be imported using an ID made up of the secret name:

```
terraform import github_codespaces_user_secret.test_secret test_secret_name
```

NOTE: the implementation is limited in that it won't fetch the value of the
`plaintext_value` or `encrypted_value` fields when importing. You may need to ignore changes for these as a workaround.
