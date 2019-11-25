---
layout: "github"
page_title: "GitHub: github_organization_project"
description: |-
  Creates and manages projects for GitHub organizations
---

# github_organization_project

This resource allows you to create and manage projects for GitHub organization.

## Example Usage

```hcl
resource "github_organization_project" "project" {
  name = "A Organization Project"
  body = "This is a organization project."
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the project.

* `body` - (Optional) The body of the project.

## Attributes Reference

The following additional attributes are exported:

* `url` - URL of the project
