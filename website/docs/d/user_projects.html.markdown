---
layout: "github"
page_title: "GitHub: github_user_projects"
description: |-
  Get information about Projects V2 for a specific user.
---

# github_user_projects

Use this data source to retrieve information about all Projects V2 for a specified GitHub user.

~> **Note**: This data source is only available when using GitHub Projects V2 (beta). Classic Projects are not supported.

## Example Usage

```hcl
data "github_user_projects" "example" {
  username = "octocat"
}

output "project_titles" {
  value = [for project in data.github_user_projects.example.projects : project.title]
}

# Reference a specific project by title
locals {
  work_project = [for project in data.github_user_projects.example.projects : project if project.title == "Work Tasks"][0]
}

output "work_project_url" {
  value = local.work_project.url
}
```

## Argument Reference

The following arguments are supported:

* `username` - (Required) The username to retrieve projects for.

## Attributes Reference

* `projects` - A list of Projects V2 for the user. Each project has the following attributes:
  * `id` - The ID of the project.
  * `node_id` - The GraphQL node ID of the project.
  * `number` - The project number.
  * `title` - The title of the project.
  * `body` - The body/description of the project.
  * `shortDescription` - The short description of the project.
  * `public` - Whether the project is public.
  * `closed` - Whether the project is closed.
  * `creator` - The username of the user who created the project.
  * `url` - The URL of the project.
  * `created_at` - The timestamp when the project was created.
  * `updated_at` - The timestamp when the project was last updated.
  * `closed_at` - The timestamp when the project was closed (if applicable).
  * `deleted_at` - The timestamp when the project was deleted (if applicable).
  * `delete_by` - The username of the user who deleted the project (if applicable).
  * `owner` - Details about the project owner:
    * `login` - The login name of the owner.
    * `id` - The ID of the owner.
    * `node_id` - The GraphQL node ID of the owner.
    * `avatar_url` - The avatar URL of the owner.
    * `gravatar_id` - The Gravatar ID of the owner.
    * `url` - The URL of the owner.
    * `html_url` - The HTML URL of the owner.
    * `type` - The type of the owner (User or Organization).
    * `site_admin` - Whether the owner is a site administrator.