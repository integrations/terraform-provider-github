---
layout: "github"
page_title: "GitHub: github_issue"
description: |-
  Provides a GitHub issue resource.
---

# github_issue

Provides a GitHub issue resource.

This resource allows you to create and manage issue within your
GitHub repository.

## Example Usage

```hcl
# Create a simple issue
resource "github_repository" "test" {
  name = "tf-acc-test-%s"
  auto_init  = true
  has_issues = true
}

resource "github_issue" "test" {
  repository       = github_repository.test.name
  title            = "My issue title"
  body             = "The body of my issue"
}
```

## Example Usage with milestone and project assignment

```hcl
# Create an issue with milestone and project assignment
resource "github_repository" "test" {
  name = "tf-acc-test-%s"
  auto_init  = true
  has_issues = true
}

resource "github_repository_milestone" "test" {
  owner = split("/", "${github_repository.test.full_name}")[0]
  repository = github_repository.test.name
  title = "v1.0.0"
  description = "General Availability"
  due_date = "2022-11-22"
  state = "open"
}

resource "github_issue" "test" {
  repository       = github_repository.test.name
  title            = "My issue"
  body             = "My issue body"
  labels           = ["bug", "documentation"]
  assignees        = ["bob-github"]
  milestone_number = github_repository_milestone.test.number
}
```

## Argument Reference

The following arguments are supported:

* `owner` - (Required) Owner of the repository the issue belongs to.

* `repository` - (Required) The GitHub repository name

* `title` - (Required) Title of the issue.

* `body` - (Optional) Title of the issue.

* `labels` - (Optional) List of labels to attach to the issue. 

* `assignees` - (Optional) List of Logins to assign the to the issue.

* `milestone_number` - (Optional) Milestone number to assign to the issue

## Attributes Reference

* `number` - (Computed) - The issue number

* `issue_id` - (Computed) - The issue id

## Import

GitHub Issues can be imported using an ID made up of `repository:number`, e.g.

```
$ terraform import github_issue.issue_15 myrepo:15
```
