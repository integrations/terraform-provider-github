---
layout: "github"
page_title: "GitHub: github_project_column"
sidebar_current: "docs-github-resource-project-column"
description: |-
  Creates and manages project columns for GitHub projects
---

# github_repository_project

This resource allows you to create and manage columns for GitHub projects.

## Example Usage

```hcl
resource "github_organization_project" "project" {
  name = "A Organization Project"
  body = "This is a organization project."
}

resource "github_project_column" "column" {
  project_id = "${github_organization_project.project.id}"
  name       = "a column"
}
```

## Argument Reference

The following arguments are supported:

* `project_id` - (Required) The id of an existing project that the column will be created in.

* `name` - (Required) The name of the column.
