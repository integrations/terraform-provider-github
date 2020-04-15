---
layout: "github"
page_title: "GitHub: github_repository_project"
description: |-
  Creates and manages projects for GitHub repositories
---

# github_repository_project

This resource allows you to create and manage projects for GitHub repository.

## Example Usage

```hcl
resource "github_repository" "example" {
  name         = "example"
  description  = "My awesome codebase"
  has_projects = true
}

resource "github_repository_project" "project" {
  name       = "A Repository Project"
  repository = "${github_repository.example.name}"
  body       = "This is a repository project."
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the project.

* `repository` - (Required) The repository of the project.

* `body` - (Optional) The body of the project.

## Attributes Reference

The following additional attributes are exported:

* `url` - URL of the project
