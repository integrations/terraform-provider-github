---
layout: "github"
page_title: "GitHub: github_actions_organization_oidc_subject_claim_customization_template"
description: |-
Creates and manages an OpenID Connect subject claim customization template for an organization
---

# github_actions_organization_oidc_subject_claim_customization_template

This resource allows you to create and manage an OpenID Connect subject claim customization template within a GitHub 
organization.

More information on integrating GitHub with cloud providers using OpenID Connect and a list of available claims is
available in the [Actions documentation](https://docs.github.com/en/actions/deployment/security-hardening-your-deployments/about-security-hardening-with-openid-connect).

## Example Usage

```hcl
resource "github_actions_organization_oidc_subject_claim_customization_template" "example_template" {
  include_claim_keys = ["actor", "context", "repository_owner"]
}
```

## Argument Reference

The following arguments are supported:

* `include_claim_keys` - (Required) A list of OpenID Connect claims.

## Import

This resource can be imported using the organization's name.

```
$ terraform import github_actions_organization_oidc_subject_claim_customization_template.test example_organization
```