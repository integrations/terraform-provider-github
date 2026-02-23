---
layout: "github"
page_title: "GitHub: github_team_settings"
description: |-
  Manages the team settings (in particular the request review delegation settings)
---

# github_team_settings

This resource manages the team settings (in particular the request review delegation settings) within the organization

Creating this resource will alter the team Code Review settings.

The team must both belong to the same organization configured in the provider on GitHub.

~> **Note**: This resource relies on the v4 GraphQl GitHub API. If this API is not available, or the Stone Crop schema preview is not available, then this resource will not work as intended.

## Example Usage

### Notify without delegation

```hcl
resource "github_team" "some_team" {
  name        = "SomeTeam"
  description = "Some cool team"
}

resource "github_team_settings" "code_review_settings" {
  team_id = github_team.some_team.id
  notify  = true
}
```

### Notify with delegation

```hcl
resource "github_team" "some_team" {
  name        = "SomeTeam"
  description = "Some cool team"
}

resource "github_team_settings" "code_review_settings" {
  team_id = github_team.some_team.id
  notify  = true
  review_request_delegation {
    algorithm    = "ROUND_ROBIN"
    member_count = 1
  }
}
```

## Argument Reference

The following arguments are supported:

- `team_id` - (Required) The GitHub team id or the GitHub team slug
- `notify` - (Optional) Whether to notify the entire team when at least one member is also assigned to the pull request. Can be set independently of `review_request_delegation`. Default value is `false`.
- `review_request_delegation` - (Optional) The settings for delegating code reviews to individuals on behalf of the team. If this block is present, even without any fields, then review request delegation will be enabled for the team. See [GitHub Review Request Delegation](#github-review-request-delegation-configuration) below for details. See [GitHub's documentation](https://docs.github.com/en/organizations/organizing-members-into-teams/managing-code-review-settings-for-your-team#configuring-team-notifications) for more configuration details.

### GitHub Review Request Delegation Configuration

The following arguments are supported:

- `algorithm` - (Optional) The algorithm to use when assigning pull requests to team members. Supported values are `ROUND_ROBIN` and `LOAD_BALANCE`. Default value is `ROUND_ROBIN`
- `member_count` - (Optional) The number of team members to assign to a pull request
- `notify` - (Optional, **Deprecated**: Use the top-level `notify` attribute instead.) Whether to notify the entire team when at least one member is also assigned to the pull request. Conflicts with the top-level `notify` attribute.

## Attributes Reference

The following additional attributes are exported:

- `team_slug` - The slug of the Team.
- `team_uid` - The unique node ID of the Team on GitHub. Corresponds to the ID of the `github_team_settings` resource.

## Import

GitHub Teams can be imported using the GitHub team ID, or the team slug e.g.

```text
terraform import github_team_settings.code_review_settings 1234567
```

or,

```text
terraform import github_team_settings.code_review_settings SomeTeam
```
