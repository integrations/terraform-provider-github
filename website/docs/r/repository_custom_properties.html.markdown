---
layout: "github"
page_title: "GitHub: github_repository_custom_properties"
description: |-
  Manages multiple custom property values for a GitHub repository
---

# github_repository_custom_properties

This resource allows you to manage multiple custom property values for a GitHub repository in a single resource block. Property values are updated in-place when changed, without recreating the resource.

~> **Note:** This resource is **non-authoritative**: it only manages the custom property values explicitly declared in the `property` blocks. Any other custom properties set on the repository by other sources (UI, API, or other tooling) will be ignored and left unchanged.

~> **Note:** This resource manages **values** for custom properties that have already been defined at the organization level (e.g. using [`github_organization_custom_properties`](organization_custom_properties.html)). It cannot create new property definitions.

~> **Note:** This resource requires the provider to be configured with an organization owner. Individual user accounts are not supported.

## Example Usage

```hcl
resource "github_repository" "example" {
  name = "example"
}

resource "github_repository_custom_properties" "example" {
  repository = github_repository.example.name

  property {
    name  = "environment"
    value = ["production"]
  }

  property {
    name  = "team"
    value = ["platform"]
  }
}
```

## Example Usage - Multi-Select Property

```hcl
resource "github_repository_custom_properties" "example" {
  repository = "my-repo"

  property {
    name  = "languages"
    value = ["go", "typescript", "python"]
  }

  property {
    name  = "environment"
    value = ["staging"]
  }
}
```

## Argument Reference

The following arguments are supported:

* `repository` - (Required) The name of the repository.

* `property` - (Required) One or more property blocks as defined below. At least one must be specified.

### property

* `name` - (Required) The name of the custom property. Must correspond to a property already defined at the organization level.

* `value` - (Required) The value(s) for the custom property. This is always specified as a set of strings, even for non-multi-select properties. For `string`, `single_select`, `true_false`, and `url` property types, provide a single value. For `multi_select` properties, multiple values can be provided.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - A composite ID in the format `owner:repository`.

* `repository_id` - The ID of the GitHub repository.

## Import

Repository custom properties can be imported using the repository name. When imported, **all** custom property values currently set on the repository will be imported into state. After import, only the properties present in your configuration will continue to be managed; any properties not declared in `property` blocks will be ignored on subsequent plans.

```
terraform import github_repository_custom_properties.example my-repo
```

## Differences from `github_repository_custom_property`

This resource (`github_repository_custom_properties`, plural) manages **multiple** custom property values for a repository in a single resource block, with in-place updates when values change. It is non-authoritative — only the properties declared in the configuration are managed. This is useful when you want to manage multiple properties together as a unit without affecting properties managed elsewhere.

The singular [`github_repository_custom_property`](repository_custom_property.html) resource manages a **single** property value per resource instance. Use it when you need independent lifecycle management for each property.
