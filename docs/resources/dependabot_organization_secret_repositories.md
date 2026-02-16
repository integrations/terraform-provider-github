---
page_title: "github_dependabot_organization_secret_repositories (Resource) - GitHub"
description: |-
  Manages repository allow list for an Dependabot Secret within a GitHub organization.
---

# github_dependabot_organization_secret_repositories (Resource)

This resource allows you to manage the repositories allowed to access a Dependabot secret within your GitHub organization. You must have write access to an organization secret to use this resource.

This resource is only applicable when `visibility` of the existing organization secret has been set to `selected`.

## Example Usage

```terraform
resource "github_dependabot_organization_secret" "example" {
  secret_name     = "mysecret"
  plaintext_value = "foo"
  visibility      = "selected"
}

resource "github_repository" "example" {
  name       = "myrepo"
  visibility = "public"
}

resource "github_dependabot_organization_secret_repositories" "example" {
  secret_name             = github_dependabot_organization_secret.example.name
  selected_repository_ids = [github_repository.example.repo_id]
}
```

## Argument Reference

The following arguments are supported:

- `secret_name` - (Required) Name of the Dependabot organization secret.
- `selected_repository_ids` - (Required) List of IDs for the repositories that should be able to access the secret.

## Import

This resource can be imported using the secret name as the ID.

### Import Block

The following import block imports the repositories able to access the Dependabot organization secret named `mysecret` to a `github_dependabot_organization_secret_repositories` resource named `example`.

```terraform
import {
  to = github_dependabot_organization_secret_repositories.example
  id = "mysecret"
}
```

### Import Command

The following command imports the repositories able to access the Dependabot organization secret named `mysecret` to a `github_dependabot_organization_secret_repositories` resource named `example`.

```shell
terraform import github_dependabot_organization_secret_repositories.example mysecret
```
