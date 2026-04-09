---
layout: "github"
page_title: "GitHub: github_enterprise_custom_property"
description: |-
  Get information about a GitHub enterprise custom property definition
---

# github_enterprise_custom_property

Use this data source to retrieve information about a custom property definition for a GitHub enterprise.

## Example Usage

```hcl
data "github_enterprise_custom_property" "security_tier" {
  enterprise_slug = "my-enterprise"
  property_name   = "securityTier"
}
```

## Example Usage - Reference in a Repository

```hcl
data "github_enterprise_custom_property" "security_tier" {
  enterprise_slug = "my-enterprise"
  property_name   = "securityTier"
}

resource "github_repository" "example" {
  name       = "example"
  visibility = "private"

  custom_properties = {
    (data.github_enterprise_custom_property.security_tier.property_name) = "tier1"
  }
}
```

## Argument Reference

The following arguments are supported:

* `enterprise_slug` - (Required) The slug of the enterprise.

* `property_name` - (Required) The name of the custom property to retrieve.

## Attributes Reference

* `value_type` - The type of the value for the property. Can be one of `string`, `single_select`, `multi_select`, `true_false`, or `url`.

* `required` - Whether the custom property is required on repositories.

* `description` - A short description of the custom property.

* `default_values` - The default value(s) of the custom property. For `multi_select` properties this is a list of values; for all other types it is a single-element list.

* `allowed_values` - An ordered list of allowed values for the property. Only populated when `value_type` is `single_select` or `multi_select`.

* `values_editable_by` - Who can edit the values of the property. Can be one of `org_actors` or `org_and_repo_actors`.
