---
layout: "github"
page_title: "GitHub: github_actions_environment_secret"
description: |-
  Creates and manages an Action Secret within a GitHub repository environment
---

# github_actions_environment_secret

This resource allows you to create and manage GitHub Actions secrets within your GitHub repository environments.
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
resource "github_actions_environment_secret" "example_secret" {
  environment       = "example_environment"
  secret_name       = "example_secret_name"
  plaintext_value   = var.some_secret_string
}

resource "github_actions_environment_secret" "example_secret" {
  environment       = "example_environment"
  secret_name       = "example_secret_name"
  encrypted_value   = var.some_encrypted_secret_string
}
```

```hcl
data "github_repository" "repo" {
  full_name = "my-org/repo"
}

resource "github_repository_environment" "repo_environment" {
  repository       = data.github_repository.repo.name
  environment      = "example_environment"
}

resource "github_actions_environment_secret" "test_secret" {
  repository       = data.github_repository.repo.name
  environment      = github_repository_environment.repo_environment.environment
  secret_name      = "test_secret_name"
  plaintext_value  = "%s"
}
```

## Example Lifecycle Ignore Changes

This resource supports the `lifecycle` `ignore_changes` block. This is for use cases where a secret value is created
using a placeholder value and then modified after creation outside the scope of Terraform. This approach ensures only
the initial placeholder value is referenced in your code and in the resulting state file.

```hcl
resource "github_actions_environment_secret" "example_secret" {
  environment       = "example_environment"
  secret_name       = "example_secret_name"
  plaintext_value   = "placeholder"

  lifecycle {
    ignore_changes = [plaintext_value]
  }
}

resource "github_actions_environment_secret" "example_secret" {
  environment       = "example_environment"
  secret_name       = "example_secret_name"
  encrypted_value   = base64sha256("placeholder")

  lifecycle {
    ignore_changes = [encrypted_value]
  }
}
```

## Argument Reference

The following arguments are supported:


* `repository`              - (Required) Name of the repository.
* `environment`             - (Required) Name of the environment.
* `secret_name`             - (Required) Name of the secret.
* `encrypted_value`         - (Optional) Encrypted value of the secret using the GitHub public key in Base64 format.
* `plaintext_value`         - (Optional) Plaintext value of the secret to be encrypted.

## Attributes Reference

* `created_at`      - Date of actions_environment_secret creation.
* `updated_at`      - Date of actions_environment_secret update.

## Import

This resource does not support importing. If you'd like to help contribute it, please visit our [GitHub page](https://github.com/integrations/terraform-provider-github)!
