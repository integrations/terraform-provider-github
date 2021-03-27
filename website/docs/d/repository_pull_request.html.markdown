---
layout: "github"
page_title: "GitHub: repository_pull_request"
description: |-
  Get information on a single GitHub Pull Request.
---

# github_repository_pull_request

Use this data source to retrieve information about a specific GitHub Pull Request in a repository.

## Example Usage

```hcl
data "github_repository_pull_request" "example" {
    base_repository = "example_repository"
    number          = 1
}
```

## Argument Reference

*  `base_repository` - (Required) Name of the base repository to retrieve the Pull Request from.

* `number` - (Required) The number of the Pull Request within the repository.

* `owner`  - (Optional) Owner of the repository. If not provided, the provider's default owner is used.

## Attributes Reference

* `base_ref` - Name of the ref (branch) of the Pull Request base.

* `base_sha` - Head commit SHA of the Pull Request base.

* `body` - Body of the Pull Request.

* `draft` - Indicates Whether this Pull Request is a draft.

* `head_owner` - Owner of the Pull Request head repository.

* `head_repository` - Name of the Pull Request head repository.

* `head_sha` - Head commit SHA of the Pull Request head.

* `labels` - List of label names set on the Pull Request.

* `maintainer_can_modify` - Indicates whether the base repository maintainers can modify the Pull Request.

* `opened_at` - Unix timestamp indicating the Pull Request creation time.

* `opened_by` - GitHub login of the user who opened the Pull Request.

* `state` - the current Pull Request state - can be "open", "closed" or "merged".

* `title` - The title of the Pull Request.

* `updated_at` - The timestamp of the last Pull Request update.
 