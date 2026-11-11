---
page_title: "github_actions_organization_fork_pr_contributor_approval (Resource) - GitHub"
description: |-
  Manages the organization-wide fork PR contributor approval policy
---

# github_actions_organization_fork_pr_contributor_approval (Resource)

This resource allows you to set the organization-wide fork pull request contributor approval policy. This controls which fork PR contributors need maintainer approval before their workflows can run on any public repository in the organization. You must be an organization owner to use this resource.

Repositories may override this policy at the repository level (see [`github_actions_repository_fork_pr_contributor_approval`](actions_repository_fork_pr_contributor_approval.md)). Setting the policy at the organization level only establishes the default for repositories that do not have a repository-level override.

The GitHub API for this setting does not expose an "off" state — the policy is always set to one of the three strictness values. If you remove this resource, the policy is reset to GitHub's documented default (`first_time_contributors`).

## Example Usage

```terraform
resource "github_actions_organization_fork_pr_contributor_approval" "test" {
  approval_policy = "all_external_contributors"
}
```

## Argument Reference

The following arguments are supported:

- `approval_policy` - (Required) The organization-wide policy controlling which fork PR contributors need maintainer approval. Possible values are `first_time_contributors_new_to_github`, `first_time_contributors`, or `all_external_contributors`.

## Import

This resource can be imported using the name of the organization:

```shell
terraform import github_actions_organization_fork_pr_contributor_approval.test my-organization
```
