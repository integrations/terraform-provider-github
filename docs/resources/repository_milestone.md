---
layout: "github"
page_title: "GitHub: github_repository_milestone"
description: |-
  Provides a GitHub repository milestone resource.
---

# github_repository_milestone

Provides a GitHub repository milestone resource.

This resource allows you to create and manage milestones for a GitHub Repository within an organization or user account.

## Example Usage

```hcl
# Create a milestone for a repository
resource "github_repository_milestone" "example" {
  owner      = "example-owner"
  repository = "example-repository"
  title      = "v1.1.0"
}
```

## Argument Reference

The following arguments are supported:

* `owner` - (Required) The owner of the GitHub Repository.

* `repository` - (Required) The name of the GitHub Repository.

* `title` - (Required) The title of the milestone.

* `description` - (Optional) A description of the milestone.

* `due_date` - (Optional) The milestone due date. In `yyyy-mm-dd` format.

* `state` - (Optional) The state of the milestone. Either `open` or `closed`. Default: `open`


## Attributes Reference

The following additional attributes are exported:

* `number` - The number of the milestone.

## Import

A GitHub Repository Milestone can be imported using an ID made up of `owner/repository/number`, e.g.

```
$ terraform import github_repository_milestone.example example-owner/example-repository/1
```
