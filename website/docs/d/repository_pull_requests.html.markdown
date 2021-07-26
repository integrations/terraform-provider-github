---
layout: "github"
page_title: "GitHub: repository_pull_requests"
description: |-
  Get information on multiple GitHub Pull Requests.
---

# github_repository_pull_requests

Use this data source to retrieve information about multiple GitHub Pull Requests in a repository.

## Example Usage

```hcl
data "github_repository_pull_requests" "example" {
    base_repository = "example-repository"
    base_ref        = "main"
    sort_by         = "updated"
    sort_direction  = "desc"
    state           = "open"
}
```

## Argument Reference

* `base_repository` - (Required) Name of the base repository to retrieve the Pull Requests from.

* `owner`  - (Optional) Owner of the repository. If not provided, the provider's default owner is used.

* `base_ref` - (Optional) If set, filters Pull Requests by base branch name.

* `head_ref` - (Optional) If set, filters Pull Requests by head user or head organization and branch name in the format of "user:ref-name" or "organization:ref-name". For example: "github:new-script-format" or "octocat:test-branch".

* `sort_by` - (Optional) If set, indicates what to sort results by. Can be either "created", "updated", "popularity" (comment count) or "long-running" (age, filtering by pulls updated in the last month). Default: "created".

* `sort_direction` - (Optional) If set, controls the direction of the sort. Can be either "asc" or "desc". Default: "asc".

* `state` - (Optional) If set, filters Pull Requests by state. Can be "open", "closed", or "all". Default: "open".

## Attributes Reference

* `results` - Collection of Pull Requests matching the filters. Each of the results conforms to the following scheme:

    * `base_ref` - Name of the ref (branch) of the Pull Request base.

    * `base_sha` - Head commit SHA of the Pull Request base.

    * `body` - Body of the Pull Request.

    * `draft` - Indicates Whether this Pull Request is a draft.

    * `head_owner` - Owner of the Pull Request head repository.

    * `head_ref` - Value of the Pull Request `HEAD` reference.

    * `head_repository` - Name of the Pull Request head repository.

    * `head_sha` - Head commit SHA of the Pull Request head.

    * `labels` - List of label names set on the Pull Request.

    * `maintainer_can_modify` - Indicates whether the base repository maintainers can modify the Pull Request.

    * `number` - The number of the Pull Request within the repository.

    * `opened_at` - Unix timestamp indicating the Pull Request creation time.

    * `opened_by` - GitHub login of the user who opened the Pull Request.

    * `state` - the current Pull Request state - can be "open", "closed" or "merged".

    * `title` - The title of the Pull Request.

    * `updated_at` - The timestamp of the last Pull Request update.
