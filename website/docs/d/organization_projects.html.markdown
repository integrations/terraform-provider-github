---
layout: "github"
page_title: "GitHub: github_organization_projects"
description: |-
  Get information about Projects V2 in an organization.
---

# github_organization_projects

Use this data source to retrieve information about all Projects V2 in a specified GitHub organization.

~> **Note**: This data source is only available when using GitHub Projects V2 (beta). Classic Projects are not supported. To use Projects V2, you need the proper organization permissions.

## Example Usage

```hcl
data "github_organization_projects" "example" {
  organization = "my-organization"
}

output "project_urls" {
  value = [for project in data.github_organization_projects.example.projects : project.url]
}

# Reference a specific project by title
locals {
  my_project = [for project in data.github_organization_projects.example.projects : project if project.title == "My Project"][0]
}

output "my_project_id" {
  value = local.my_project.id
}
```

## Argument Reference

The following arguments are supported:

* `organization` - (Required) The name of the organization.

## Attributes Reference

* `projects` - A list of Projects V2 in the organization. Each project has the following attributes:
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