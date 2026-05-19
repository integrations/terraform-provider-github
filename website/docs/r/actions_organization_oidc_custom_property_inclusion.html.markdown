---
layout: "github"
page_title: "GitHub: github_actions_organization_oidc_custom_property_inclusion"
description: |-
  Manages a custom property inclusion in OIDC tokens for a GitHub organization
---

# github_actions_organization_oidc_custom_property_inclusion

This resource allows you to add a repository custom property to be included in the OIDC token for repository actions
in a GitHub organization.

When a custom property is included, its value will be available as a claim in the OIDC token issued to GitHub Actions
workflows. This enables cloud providers to make authorization decisions based on repository custom properties.

More information on integrating GitHub with cloud providers using OpenID Connect and a list of available claims is
available in the [Actions documentation](https://docs.github.com/en/actions/deployment/security-hardening-your-deployments/about-security-hardening-with-openid-connect).

~> **Note:** This resource requires the organization to have custom properties already defined. The custom property
referenced by `custom_property_name` must exist as an organization-level custom property.

## Example Usage

### Single custom property inclusion

```hcl
resource "github_actions_organization_oidc_custom_property_inclusion" "environment" {
  custom_property_name = "environment"
}
```

### Multiple custom property inclusions with repository custom properties

```hcl
resource "github_organization_custom_properties" "props" {
  property {
    property_name = "environment"
    value_type    = "single_select"
    required      = true
    allowed_values = ["production", "staging", "development"]
  }

  property {
    property_name = "team"
    value_type    = "string"
  }
}

# Include custom properties in OIDC tokens
resource "github_actions_organization_oidc_custom_property_inclusion" "environment" {
  custom_property_name = "environment"
}

resource "github_actions_organization_oidc_custom_property_inclusion" "team" {
  custom_property_name = "team"
}

# Set custom property values on a repository
resource "github_repository" "example" {
  name = "example-repository"
}

resource "github_repository_custom_property" "env" {
  repository     = github_repository.example.name
  property_name  = "environment"
  property_type  = "single_select"
  property_value = ["production"]
}

# Configure OIDC subject claim to include the custom property
resource "github_actions_repository_oidc_subject_claim_customization_template" "example" {
  repository        = github_repository.example.name
  use_default       = false
  include_claim_keys = ["repo", "context", "repository_custom_property_environment"]
}
```

## Argument Reference

The following arguments are supported:

- `custom_property_name` - (Required) The name of the custom property to include in the OIDC token. This must match
  an existing organization-level custom property.

## Import

This resource can be imported using the organization name and custom property name separated by a colon (`:`).

```
$ terraform import github_actions_organization_oidc_custom_property_inclusion.environment organization-name:environment
```
