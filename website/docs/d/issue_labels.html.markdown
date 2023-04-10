---
layout: "github"
page_title: "GitHub: github_issue_labels"
description: |-
  Get the labels for a given repository.
---

# github_labels

Use this data source to retrieve the labels for a given repository.

## Example Usage

```hcl
data "github_labels" "test" {
  repository = "example_repository"
}
```

## Arguments Reference

* `repository` - (Required) The name of the repository.

## Attributes Reference

* `labels` - The list of this repository's labels. Each element of `labels` has the following attributes:
  * `name` - The name of the label.
  * `color` - The hexadecimal color code for the label, without the leading #.
  * `description` - A short description of the label.
  * `url` - The URL of the label.
