---
layout: "github"
page_title: "GitHub: github_repository_milestone"
description: |-
  Get information on a GitHub Repository Milestone.
---

# github_repository_milestone

Use this data source to retrieve information about a specific GitHub milestone in a repository.

## Example Usage

```hcl
data "github_repository_milestone" "example" {
    owner       = "example-owner"
    repository  = "example-repository"
    number      = 1
}
```

## Argument Reference

 *  `owner`  -  (Required) Owner of the repository.
 
 *  `repository`  -  (Required) Name of the repository to retrieve the milestone from.

 *  `number`  -  (Required) The number of the milestone.

## Attributes Reference

 * `description` - Description of the milestone.
 * `due_date` - The milestone due date (in ISO-8601 `yyyy-mm-dd` format). 
 * `state` - State of the milestone.
 * `title` - Title of the milestone.
