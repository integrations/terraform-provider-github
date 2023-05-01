---
layout: "github"
page_title: "GitHub: github_actions_repository_oidc_subject_claim_customization_template"
description: |-
Creates and manages an OpenID Connect subject claim customization template for a repository
---

# github_actions_repository_oidc_subject_claim_customization_template

This resource allows you to create and manage an OpenID Connect subject claim customization template for a GitHub
repository.

More information on integrating GitHub with cloud providers using OpenID Connect and a list of available claims is
available in the [Actions documentation](https://docs.github.com/en/actions/deployment/security-hardening-your-deployments/about-security-hardening-with-openid-connect).

The following table lists the behaviour of `use_default`:

| `use_default` | `include_claim_keys` | Template used                                             |
|---------------|----------------------|-----------------------------------------------------------|
| `true`        | Unset                | GitHub's default                                          |
| `false`       | Set                  | `include_claim_keys`                                      |
| `false`       | Unset                | Organization's default if set, otherwise GitHub's default |

## Example Usage

```hcl
resource "github_repository" "example" {
  name = "example-repository"
}

resource "github_actions_repository_oidc_subject_claim_customization_template" "example_template" {
  repository = github_repository.example.name
  use_default = false
  include_claim_keys = ["actor", "context", "repository_owner"]
}
```

## Argument Reference

The following arguments are supported:

* `use_default`        - (Required) Whether to use the default template or not. If `true`, `include_claim_keys` must not
be set.
* `include_claim_keys` - (Optional) A list of OpenID Connect claims.

## Import

This resource can be imported using the repository's name.

```
$ terraform import github_actions_repository_oidc_subject_claim_customization_template.test example_repository
```