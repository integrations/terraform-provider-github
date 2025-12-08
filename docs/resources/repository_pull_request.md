---
layout: "github"
page_title: "GitHub: repository_pull_request"
description: |-
  Get information on a single GitHub Pull Request.
---

# github_repository_pull_request

This resource allows you to create and manage PullRequests for repositories within your GitHub organization or personal account.

## Example Usage

```hcl
resource "github_repository_pull_request" "example" {
    base_repository = "example-repository"
    base_ref        = "main"
    head_ref        = "feature-branch"
    title           = "My newest feature"
    body            = "This will change everything"
}
```

## Argument Reference

* `base_repository` - (Required) Name of the base repository to retrieve the Pull Requests from.

* `base_ref` - (Required) Name of the branch serving as the base of the Pull Request.

* `head_ref` - (Required) Name of the branch serving as the head of the Pull Request.

* `owner`  - (Optional) Owner of the repository. If not provided, the provider's default owner is used.

* `title` - (Optional) The title of the Pull Request.

* `body` - (Optional) Body of the Pull Request.

* `maintainer_can_modify` - Controls whether the base repository maintainers can modify the Pull Request. Default: false.

## Attributes Reference

* `base_sha` - Head commit SHA of the Pull Request base.

* `draft` - Indicates Whether this Pull Request is a draft.

* `head_sha` - Head commit SHA of the Pull Request head.

* `labels` - List of label names set on the Pull Request.

* `number` - The number of the Pull Request within the repository.

* `opened_at` - Unix timestamp indicating the Pull Request creation time.

* `opened_by` - GitHub login of the user who opened the Pull Request.

* `state` - the current Pull Request state - can be "open", "closed" or "merged".

* `updated_at` - The timestamp of the last Pull Request update.
