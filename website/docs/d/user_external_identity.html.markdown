---
layout: "github"
page_title: "GitHub: github_user_external_identity"
description: |-
  Get information on user's identity in extrenal IDP.
---

# github_user_external_identity

Use this data source to find out the external IDP identity of a user inside a specific GitHub Organization. See more about external IDP:s [here](https://docs.github.com/en/enterprise-cloud@latest/organizations/managing-saml-single-sign-on-for-your-organization/connecting-your-identity-provider-to-your-organization).

## Example Usage

```hcl
data "github_user_external_identity" "test" {
  username        = "username"
  organization    = "github-org"
}
```

## Argument Reference

 * `username` - (Required) The username to lookup in the organization.

 * `organization` - (Optional) The organization to check for the above username. If left empty, the Organization configured in the provider itself will be used

## Attributes Reference

 * `saml_identity` - The SAML information from the external IDP
 * `scim_identity` - The SCIM information from the external IDP
