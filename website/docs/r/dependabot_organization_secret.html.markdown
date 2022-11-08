---
layout: "github"
page_title: "GitHub: github_dependabot_organization_secret"
description: |-
  Creates and manages an Dependabot Secret within a GitHub organization
---

# github_dependabot_organization_secret

This resource allows you to create and manage GitHub Dependabot secrets within your GitHub organization.
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
resource "github_dependabot_organization_secret" "example_secret" {
  secret_name     = "example_secret_name"
  visibility      = "private"
  plaintext_value = var.some_secret_string
}

resource "github_dependabot_organization_secret" "example_secret" {
  secret_name     = "example_secret_name"
  visibility      = "private"
  encrypted_value = var.some_encrypted_secret_string
}
```

```hcl
data "github_repository" "repo" {
  full_name = "my-org/repo"
}

resource "github_dependabot_organization_secret" "example_secret" {
  secret_name             = "example_secret_name"
  visibility              = "selected"
  plaintext_value         = var.some_secret_string
  selected_repository_ids = [data.github_repository.repo.repo_id]
}

resource "github_dependabot_organization_secret" "example_secret" {
  secret_name             = "example_secret_name"
  visibility              = "selected"
  encrypted_value         = var.some_encrypted_secret_string
  selected_repository_ids = [data.github_repository.repo.repo_id]
}
```

## Argument Reference

The following arguments are supported:

* `secret_name`             - (Required) Name of the secret
* `encrypted_value`         - (Optional) Encrypted value of the secret using the Github public key in Base64 format.
* `plaintext_value`         - (Optional) Plaintext value of the secret to be encrypted
* `visibility`              - (Required) Configures the access that repositories have to the organization secret.
                              Must be one of `all`, `private`, `selected`. `selected_repository_ids` is required if set to `selected`.
* `selected_repository_ids` - (Optional) An array of repository ids that can access the organization secret.

## Attributes Reference

* `created_at`      - Date of dependabot_secret creation.
* `updated_at`      - Date of dependabot_secret update.

## Import

This resource can be imported using an ID made up of the secret name:

```
terraform import github_dependabot_organization_secret.test_secret test_secret_name
```

NOTE: the implementation is limited in that it won't fetch the value of the
`plaintext_value` or `encrypted_value` fields when importing. You may need to ignore changes for these as a workaround.
