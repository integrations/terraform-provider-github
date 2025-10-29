---
layout: "github"
page_title: "GitHub: github_issue_labels"
description: |-
  Provides GitHub issue labels resource.
---

# github_issue_labels

Provides GitHub issue labels resource.

This resource allows you to create and manage issue labels within your
GitHub organization.

~> Note: github_issue_labels cannot be used in conjunction with github_issue_label or they will fight over what your policy should be.

This resource is authoritative. For adding a label to a repo in a non-authoritative manner, use github_issue_label instead.

If you change the case of a label's name, its' color, or description, this resource will edit the existing label to match the new values. However, if you change the name of a label, this resource will create a new label with the new name and delete the old label. Beware that this will remove the label from any issues it was previously attached to.

~> **Note:** When a repository is archived, Terraform will skip deletion of issue labels to avoid API errors, as archived repositories are read-only. The labels will be removed from Terraform state without attempting to delete them from GitHub.

## Example Usage

```hcl
# Create a new, red colored label
resource "github_issue_labels" "test_repo" {
  repository = "test-repo"

  label {
    name  = "Urgent"
    color = "FF0000"
  }

  label {
    name  = "Critical"
    color = "FF0000"
  }
}
```

## Argument Reference

The following arguments are supported:

* `repository` - (Required) The GitHub repository

* `name` - (Required) The name of the label.

* `color` - (Required) A 6 character hex code, **without the leading #**, identifying the color of the label.

* `description` - (Optional) A short description of the label.

* `url` - (Computed) The URL to the issue label

## Import

GitHub Issue Labels can be imported using the repository `name`, e.g.

```
$ terraform import github_issue_labels.test_repo test_repo
```
