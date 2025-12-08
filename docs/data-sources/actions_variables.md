---
layout: "github"
page_title: "GitHub: github_actions_variables"
description: |-
  Get Actions variables for a repository
---

# github\_actions\_variables

Use this data source to retrieve the list of variables for a GitHub repository.

## Example Usage

```hcl
data "github_actions_variables" "example" {
  name = "example"
}
```

## Argument Reference

 * `name`       - (Optional) The name of the repository.
 * `full_name`  - (Optional) Full name of the repository (in `org/name` format).

## Attributes Reference

 * `variables` - list of variables for the repository
   * `name`         - Name of the variable
   * `value`        - Value of the variable
   * `created_at`   - Timestamp of the variable creation
   * `updated_at`   - Timestamp of the variable last update
 
