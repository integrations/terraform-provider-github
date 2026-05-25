---
page_title: "github_actions_repository_fork_pr_contributor_approval (Data Source) - GitHub"
description: |-
  Read the fork PR contributor approval policy for a GitHub repository
---

# github_actions_repository_fork_pr_contributor_approval (Data Source)

Use this data source to retrieve the current fork pull request contributor approval policy configured on a GitHub repository.

## Example Usage

```terraform
data "github_actions_repository_fork_pr_contributor_approval" "example" {
  repository = "my-repository"
}
```

## Argument Reference

- `repository` - (Required) The GitHub repository.

## Attributes Reference

- `approval_policy` - The fork PR contributor approval policy currently configured on the repository. One of `first_time_contributors_new_to_github`, `first_time_contributors`, or `all_external_contributors`.
