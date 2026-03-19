---
layout: "github"
page_title: "GitHub: github_actions_organization_secret_repository"
description: |-
  Add access for a repository to an Actions Secret within a GitHub organization.
---

# github_actions_organization_secret_repository

This resource adds permission for a repository to use an actions secret within your GitHub organization.
You must have write access to an organization secret to use this resource.

This resource is only applicable when `visibility` of the existing organization secret has been set to `selected`.

## Example Usage

```hcl
resource "github_actions_organization_secret" "example" {
	secret_name = "mysecret"
	value       = "foo"
	visibility  = "selected"
}

resource "github_repository" "example" {
	name       = "myrepo"
	visibility = "public"
}

resource "github_actions_organization_secret_repository" "example" {
  secret_name   = github_actions_organization_secret.example.name
  repository_id = github_repository.example.repo_id
}
```

## Argument Reference

The following arguments are supported:

- `secret_name` - (Required) Name of the actions organization secret.
- `repository_id` - (Required) ID of the repository that should be able to access the secret.

## Import

This resource can be imported using an ID made of the secret name and repository name separated by a `:`.

### Import Block

The following import block imports the access of repository ID `123456` for the actions organization secret named `mysecret` to a `github_actions_organization_secret_repository` resource named `example`.

```hcl
import {
  to = github_actions_organization_secret_repository.example
  id = "mysecret:123456"
}
```

### Import Command

The following command imports the access of repository ID `123456` for the actions organization secret named `mysecret` to a `github_actions_organization_secret_repository` resource named `example`.

```shell
terraform import github_actions_organization_secret_repository.example mysecret:123456
```
