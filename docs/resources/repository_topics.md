---
layout: "github"
page_title: "GitHub: github_repository_topics"
description: |-
  Creates and manages the topics on a repository
---

# github_repository_topics

This resource allows you to create and manage topics for repositories within your GitHub organization or personal account.

~> Note: This resource is not compatible with the `topic` attribute of the `github_repository` Use either ``github_repository_topics``
or ``topic`` in ``github_repository``. `github_repository_topics` is only meant to be used if the repository itself is not handled via terraform, for example if it's only read as a datasource (see [issue #1845](https://github.com/integrations/terraform-provider-github/issues/1845)).

## Example Usage

```hcl
data "github_repository" "test" {
    name = "test"
}

resource "github_repository_topics" "test" {
    repository    = github_repository.test.name
    topics        = ["topic-1", "topic-2"]
}
```

## Argument Reference

The following arguments are supported:

* `repository` - (Required) The repository name.

* `topics` - (Required) A list of topics to add to the repository.

## Import

Repository topics can be imported using the `name` of the repository.

```
$ terraform import github_repository_topics.terraform terraform
```
