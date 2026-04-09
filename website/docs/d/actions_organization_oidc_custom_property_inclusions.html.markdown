---
layout: "github"
page_title: "GitHub: github_actions_organization_oidc_custom_property_inclusions"
description: |-
  Lists the repository custom properties included in OIDC tokens for a GitHub organization
---

# github_actions_organization_oidc_custom_property_inclusions

Use this data source to retrieve the list of repository custom properties that are included in the OIDC token for
repository actions in a GitHub organization.

## Example Usage

```hcl
data "github_actions_organization_oidc_custom_property_inclusions" "example" {}

output "included_properties" {
  value = data.github_actions_organization_oidc_custom_property_inclusions.example.custom_property_names
}
```

## Argument Reference

This data source has no required arguments.

## Attributes Reference

* `custom_property_names` - A list of custom property names that are included in the OIDC token.
