---
layout: "github"
page_title: "GitHub: github_organization_custom_properties"
description: |-
  Creates and manages custom properties for a GitHub organization
---

# github_organization_custom_properties

This resource allows you to create and manage custom properties for a GitHub organization.

Custom properties enable you to add metadata to repositories within your organization. You can use custom properties to add context about repositories, such as who owns them, when they expire, or compliance requirements.

## Example Usage

```hcl
resource "github_organization_custom_properties" "environment" {
  property_name = "environment"
  value_type    = "single_select"
  required      = true
  description   = "The deployment environment for this repository"
  default_value = "development"
  allowed_values = [
    "development",
    "staging", 
    "production"
  ]
}
```

## Example Usage - Text Property

```hcl
resource "github_organization_custom_properties" "owner" {
  property_name = "owner"
  value_type    = "string"
  required      = true
  description   = "The team or individual responsible for this repository"
}
```

## Example Usage - Boolean Property

```hcl
resource "github_organization_custom_properties" "archived" {
  property_name = "archived"
  value_type    = "true_false"
  required      = false
  description   = "Whether this repository is archived"
  default_value = "false"
}
```

## Argument Reference

The following arguments are supported:

* `property_name` - (Required) The name of the custom property.

* `value_type` - (Optional) The type of the custom property. Can be one of `string`, `single_select`, `multi_select`, or `true_false`. Defaults to `string`.

* `required` - (Optional) Whether the custom property is required. Defaults to `false`.

* `description` - (Optional) The description of the custom property.

* `default_value` - (Optional) The default value of the custom property.

* `allowed_values` - (Optional) List of allowed values for the custom property. Only applicable when `value_type` is `single_select` or `multi_select`.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `property_name` - The name of the custom property.

## Import

Organization custom properties can be imported using the property name:

```
terraform import github_organization_custom_properties.environment environment
```