---
page_title: "github_organization_custom_properties (Data Source) - GitHub"
description: |-
  Get information about a GitHub organization custom property
---

# github_organization_custom_properties (Data Source)

~> **Deprecated:** Use [`github_organization_repository_custom_property`](organization_repository_custom_property) to look up one property, or [`github_organization_repository_custom_properties`](organization_repository_custom_properties) to list all properties. This data source will be removed in a future major release.

Use this data source to retrieve information about a GitHub organization custom property.

## Migration

For a single-property lookup, replace the data source type and keep `property_name`:

```hcl
data "github_organization_repository_custom_property" "environment" {
  property_name = "environment"
}
```

The single-property data source keeps the property attributes at the top level. Its `default_value` is now a list of strings; use `default_value[0]` where one value was expected, and keep the full list for `multi_select` properties.

To retrieve every property, use the plural data source, which returns a `properties` set instead of a single property's top-level attributes:

```hcl
data "github_organization_repository_custom_properties" "all" {}
```

Reference individual entries through `data.github_organization_repository_custom_properties.all.properties`, filtering or iterating by `property_name` as needed.

## Example Usage

```terraform
data "github_organization_custom_properties" "environment" {
  property_name = "environment"
}
```

## Argument Reference

The following arguments are supported:

- `property_name` - (Required) The name of the custom property to retrieve.

## Attributes Reference

- `property_name` - The name of the custom property.

- `value_type` - The type of the custom property. Can be one of `string`, `single_select`, `multi_select`, or `true_false`.

- `required` - Whether the custom property is required.

- `description` - The description of the custom property.

- `default_value` - The default value of the custom property.

- `allowed_values` - List of allowed values for the custom property. Only populated when `value_type` is `single_select` or `multi_select`.

- `values_editable_by` - Who can edit the values of the custom property. Can be one of `org_actors` or `org_and_repo_actors`.
