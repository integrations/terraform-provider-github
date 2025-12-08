---
layout: "github"
page_title: "GitHub: actions_organization_oidc_subject_claim_customization_template"
description: |-
  Get a GitHub Actions organization OpenID Connect customization template
---

# actions_organization_oidc_subject_claim_customization_template

Use this data source to retrieve the OpenID Connect subject claim customization template for an organization

## Example Usage

```hcl
data "github_actions_organization_oidc_subject_claim_customization_template" "example" {
}
```

## Argument Reference

## Attributes Reference

 * `include_claim_keys` - The list of OpenID Connect claim keys.
