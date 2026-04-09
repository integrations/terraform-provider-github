---
layout: "github"
page_title: "GitHub: github_enterprise_custom_property"
description: |-
  Creates and manages custom property definitions for a GitHub enterprise
---

# github_enterprise_custom_property

This resource allows you to create and manage custom property definitions for a GitHub enterprise.

Custom properties enable you to add metadata to repositories across your enterprise. Properties defined at the enterprise level are available to all organizations within the enterprise. You can use them to add context about repositories, such as security classification, compliance requirements, or team ownership.

~> **Note** You must be an enterprise owner to manage enterprise custom properties.

## Example Usage

```hcl
resource "github_enterprise_custom_property" "security_tier" {
  enterprise_slug = "my-enterprise"
  property_name   = "securityTier"
  value_type      = "single_select"
  required        = true
  description     = "Security classification tier for the repository"
  allowed_values  = ["tier1", "tier2", "tier3"]
}
```

## Example Usage - String Property

```hcl
resource "github_enterprise_custom_property" "owner" {
  enterprise_slug = "my-enterprise"
  property_name   = "owningTeam"
  value_type      = "string"
  required        = true
  description     = "The team responsible for this repository"
}
```

## Example Usage - Boolean Property

```hcl
resource "github_enterprise_custom_property" "contains_pii" {
  enterprise_slug = "my-enterprise"
  property_name   = "containsPII"
  value_type      = "true_false"
  required        = false
  description     = "Whether this repository contains personally identifiable information"
  default_values  = ["false"]
}
```

## Example Usage - Allow Repository Actors to Edit

```hcl
resource "github_enterprise_custom_property" "team_contact" {
  enterprise_slug    = "my-enterprise"
  property_name      = "teamContact"
  value_type         = "string"
  required           = false
  description        = "Contact information for the team managing this repository"
  values_editable_by = "org_and_repo_actors"
}
```

## Argument Reference

The following arguments are supported:

* `enterprise_slug` - (Required, Forces new resource) The slug of the enterprise.

* `property_name` - (Required, Forces new resource) The name of the custom property.

* `value_type` - (Required) The type of value for the property. Can be one of `string`, `single_select`, `multi_select`, `true_false`, or `url`.

* `required` - (Optional) Whether the custom property is required on repositories. Defaults to `false`.

* `description` - (Optional) A short description of the custom property.

* `default_values` - (Optional) The default value(s) of the custom property. For `multi_select` properties, multiple values may be specified (e.g. `["b1", "b2"]`). For all other types, provide a single value in a list (e.g. `["value"]`).

* `allowed_values` - (Optional) An ordered list of allowed values for the property. Only applicable when `value_type` is `single_select` or `multi_select`. Can have up to 200 values.

* `values_editable_by` - (Optional) Who can edit the values of the property on repositories. Can be one of `org_actors` or `org_and_repo_actors`. When set to `org_actors` (the default), only organization owners can edit property values. When set to `org_and_repo_actors`, repository administrators with the custom properties permission can also edit values.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `property_name` - The name of the custom property.

## Import

Enterprise custom properties can be imported using `<enterprise_slug>:<property_name>`:

```shell
terraform import github_enterprise_custom_property.security_tier my-enterprise:securityTier
```
