---
page_title: "github_actions_organization_fork_pr_contributor_approval (Data Source) - GitHub"
description: |-
  Read the organization-wide fork PR contributor approval policy
---

# github_actions_organization_fork_pr_contributor_approval (Data Source)

Use this data source to retrieve the current organization-wide fork pull request contributor approval policy.

## Example Usage

```terraform
data "github_actions_organization_fork_pr_contributor_approval" "example" {}
```

## Argument Reference

This data source takes no arguments. The organization is determined by the provider configuration.

## Attributes Reference

- `approval_policy` - The organization-wide fork PR contributor approval policy currently configured. One of `first_time_contributors_new_to_github`, `first_time_contributors`, or `all_external_contributors`.
