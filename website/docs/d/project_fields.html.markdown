---
layout: "github"
page_title: "GitHub: github_project_fields"
description: |-
  Get information about fields in a GitHub Projects V2 project.
---

# github_project_fields

Use this data source to retrieve information about all fields in a specified GitHub Projects V2 project.

~> **Note**: This data source is only available when using GitHub Projects V2 (beta). Classic Projects are not supported.

## Example Usage

```hcl
# Get fields from an organization project
data "github_project_fields" "org_project" {
  project_number = 1
  organization   = "my-organization"
}

# Get fields from a user project
data "github_project_fields" "user_project" {
  project_number = 2
  username       = "octocat"
}

# Output field names and types
output "field_info" {
  value = [
    for field in data.github_project_fields.org_project.fields : {
      name = field.name
      type = field.data_type
    }
  ]
}

# Find a specific field by name
locals {
  status_field = [
    for field in data.github_project_fields.org_project.fields :
    field if field.name == "Status"
  ][0]
}

# Output the options for a select field
output "status_options" {
  value = local.status_field.options[*].name
}
```

## Argument Reference

The following arguments are supported:

* `project_number` - (Required) The number of the project.
* `organization` - (Optional) The name of the organization that owns the project. Cannot be used with `username`.
* `username` - (Optional) The username that owns the project. Cannot be used with `organization`.

~> **Note**: Either `organization` or `username` must be specified, but not both.

## Attributes Reference

* `fields` - A list of fields in the project. Each field has the following attributes:
  * `id` - The ID of the field.
  * `node_id` - The GraphQL node ID of the field.
  * `name` - The name of the field.
  * `data_type` - The data type of the field (e.g., "text", "number", "date", "single_select", "iteration").
  * `created_at` - The timestamp when the field was created.
  * `updated_at` - The timestamp when the field was last updated.
  * `options` - A list of options for single_select fields. Each option has the following attributes:
    * `id` - The ID of the option.
    * `node_id` - The GraphQL node ID of the option.
    * `name` - The name of the option.
    * `color` - The color of the option.
    * `description` - The description of the option.

## Field Types

The following field types are available in Projects V2:

* `text` - Single line text
* `number` - Numeric values
* `date` - Date values
* `single_select` - Single selection from predefined options
* `iteration` - Iteration/sprint planning field

For `single_select` fields, the `options` attribute will contain the available choices. For other field types, the `options` attribute will be empty.