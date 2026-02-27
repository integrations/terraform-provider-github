---
page_title: "GitHub: github_organization_custom_properties"
description: |-
  Get information about a GitHub organization custom property
---

# github_organization_custom_properties

Use this data source to retrieve information about a GitHub organization custom property.

## Example Usage

```terraform
data "github_organization_custom_properties" "environment" {
  property_name = "environment"
}
```

## Argument Reference

The following arguments are supported:

* `property_name` - (Required) The name of the custom property to retrieve.

## Attributes Reference

* `property_name` - The name of the custom property.

* `value_type` - The type of the custom property. Can be one of `string`, `single_select`, `multi_select`, or `true_false`.

* `required` - Whether the custom property is required.

* `description` - The description of the custom property.

* `default_value` - The default value of the custom property.

* `allowed_values` - List of allowed values for the custom property. Only populated when `value_type` is `single_select` or `multi_select`.

* `values_editable_by` - Who can edit the values of the custom property. Can be one of `org_actors` or `org_and_repo_actors`.
