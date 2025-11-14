---
layout: "github"
page_title: "GitHub: github_project_item"
description: |-
  Creates and manages project items for GitHub Projects V2
---

# github_project_item

This resource allows you to create and manage items in GitHub Projects V2 (organization-level projects).

~> **Note:** This resource replaces the deprecated `github_project_card` resource, which was used with the now-sunset GitHub Classic Projects.

## Example Usage

```hcl
resource "github_repository" "example" {
  name        = "example"
  has_issues  = true
}

resource "github_issue" "example" {
  repository = github_repository.example.name
  title      = "Example issue"
  body       = "This is an example issue"
}

resource "github_organization_project" "example" {
  name = "Example Project"
  body = "This is an example project"
}

resource "github_project_item" "example" {
  project_number = github_organization_project.example.project_number
  content_id     = github_issue.example.issue_id
  content_type   = "Issue"
}
```

## Example Usage with Archived Item

```hcl
resource "github_project_item" "archived_example" {
  project_number = github_organization_project.example.project_number
  content_id     = github_issue.example.issue_id
  content_type   = "Issue"
  archived       = true
}
```

## Argument Reference

The following arguments are supported:

* `project_number` - (Required) The number of the project (Projects V2).
* `content_id` - (Required) The ID of the issue or pull request to add to the project.
* `content_type` - (Required) Must be either `Issue` or `PullRequest`.
* `archived` - (Optional) Whether the item is archived. Defaults to `false`.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `item_id` - The ID of the project item.
* `node_id` - The node ID of the project item.

## Import

A GitHub Project Item can be imported using the format `org/project_number/item_id`:

```
$ terraform import github_project_item.example myorg/123/456
```

Where:
- `myorg` is the organization name
- `123` is the project number
- `456` is the item ID