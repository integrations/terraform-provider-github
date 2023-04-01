---
layout: "github"
page_title: "GitHub: github_actions_organization_variables"
description: |-
  Get Actions variables of the organization
---

# github\_actions\_organization\_variables

Use this data source to retrieve the list of variables of the organization.

## Example Usage

```hcl
data "github_actions_organization_variables" "example" {
}
```

## Argument Reference

## Attributes Reference

 * `variables` - list of variables for the repository
   * `name`         - Name of the variable
   * `value`        - Value of the variable
   * `visibility`   - Visibility of the variable
   * `created_at`   - Timestamp of the variable creation
   * `updated_at`   - Timestamp of the variable last update