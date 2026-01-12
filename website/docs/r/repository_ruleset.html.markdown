---
layout: "github"
page_title: "github_repository_ruleset Resource - terraform-provider-github"
description: |-
  Creates a GitHub repository ruleset.
---

# github_repository_ruleset

Creates a GitHub repository ruleset.

This resource allows you to create and manage rulesets on the repository level. When applied, a new ruleset will be created. When destroyed, that ruleset will be removed.

## Example Usage

```hcl
resource "github_repository" "example" {
  name        = "example"
  description = "Example repository"
}

resource "github_repository_ruleset" "example" {
  name        = "example"
  repository  = github_repository.example.name
  target      = "branch"
  enforcement = "active"

  conditions {
    ref_name {
      include = ["~ALL"]
      exclude = []
    }
  }

  bypass_actors {
    actor_id    = 13473
    actor_type  = "Integration"
    bypass_mode = "always"
  }

  rules {
    creation                = true
    update                  = true
    deletion                = true
    required_linear_history = true
    required_signatures     = true

    required_deployments {
      required_deployment_environments = ["test"]
    }

    required_code_scanning {
      required_code_scanning_tool {
        alerts_threshold          = "errors"
        security_alerts_threshold = "high_or_higher"
        tool                      = "CodeQL"
      }
    }
  }
}

# Example with push ruleset
resource "github_repository_ruleset" "example_push" {
  name        = "example_push"
  repository  = github_repository.example.name
  target      = "push"
  enforcement = "active"

  rules {
    file_path_restriction {
      restricted_file_paths = [".github/workflows/*", "*.env"]
    }

    max_file_size {
      max_file_size = 100  # 100 MB
    }

    max_file_path_length {
      max_file_path_length = 255
    }

    file_extension_restriction {
      restricted_file_extensions = ["*.exe", "*.dll", "*.so"]
    }
  }
}
```

## Argument Reference

- `enforcement` - (Required) (String) Possible values for Enforcement are `disabled`, `active`, `evaluate`. Note: `evaluate` is currently only supported for owners of type `organization`.

- `name` - (Required) (String) The name of the ruleset.

- `rules` - (Required) (Block List, Min: 1, Max: 1) Rules within the ruleset. (see [below for nested schema](#rules))

- `target` - (Required) (String) Possible values are `branch`, `tag` and `push`.

- `bypass_actors` - (Optional) (Block List) The actors that can bypass the rules in this ruleset. (see [below for nested schema](#bypass_actors))

- `conditions` - (Optional) (Block List, Max: 1) Parameters for a repository ruleset condition. For `branch` and `tag` targets, `ref_name` is required. For `push` targets, `ref_name` must NOT be set - conditions are optional for push targets. (see [below for nested schema](#conditions))

- `repository` - (Required) (String) Name of the repository to apply ruleset to.

#### Rules ####

The `rules` block supports the following:

~> **Note:** Rules are target-specific. `branch` and `tag` targets support rules like `creation`, `deletion`, `pull_request`, `required_status_checks`, etc. `push` targets only support `file_path_restriction`, `max_file_size`, `max_file_path_length`, and `file_extension_restriction`. Using the wrong rules for a target will result in a validation error.

- `branch_name_pattern` - (Optional) (Block List, Max: 1) Parameters to be used for the branch_name_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. Conflicts with `tag_name_pattern` as it only applied to rulesets with target `branch`. (see [below for nested schema](#rulesbranch_name_pattern))

- `commit_author_email_pattern` - (Optional) (Block List, Max: 1) Parameters to be used for the commit_author_email_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. (see [below for nested schema](#rulescommit_author_email_pattern))

- `commit_message_pattern` - (Optional) (Block List, Max: 1) Parameters to be used for the commit_message_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. (see [below for nested schema](#rulescommit_message_pattern))

- `committer_email_pattern` - (Optional) (Block List, Max: 1) Parameters to be used for the committer_email_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. (see [below for nested schema](#rulescommitter_email_pattern))

- `creation` - (Optional) (Boolean) Only allow users with bypass permission to create matching refs.

- `deletion` - (Optional) (Boolean) Only allow users with bypass permissions to delete matching refs.

- `non_fast_forward` - (Optional) (Boolean) Prevent users with push access from force pushing to branches.

- `merge_queue` - (Optional) (Block List, Max: 1) Merges must be performed via a merge queue. (see [below for nested schema](#rulesmerge_queue))

- `pull_request` - (Optional) (Block List, Max: 1) Require all commits be made to a non-target branch and submitted via a pull request before they can be merged. (see [below for nested schema](#rulespull_request))

- `copilot_code_review` - (Optional) (Block List, Max: 1) Automatically request Copilot code review for new pull requests if the author has access to Copilot code review and their premium requests quota has not reached the limit. (see [below for nested schema](#rulescopilot_code_review))

- `required_deployments` - (Optional) (Block List, Max: 1) Choose which environments must be successfully deployed to before branches can be merged into a branch that matches this rule. (see [below for nested schema](#rulesrequired_deployments))

- `required_linear_history` - (Optional) (Boolean) Prevent merge commits from being pushed to matching branches.

- `required_signatures` - (Optional) (Boolean) Commits pushed to matching branches must have verified signatures.

- `required_status_checks` - (Optional) (Block List, Max: 1) Choose which status checks must pass before branches can be merged into a branch that matches this rule. When enabled, commits must first be pushed to another branch, then merged or pushed directly to a branch that matches this rule after status checks have passed. (see [below for nested schema](#rulesrequired_status_checks))

- `tag_name_pattern` - (Optional) (Block List, Max: 1) Parameters to be used for the tag_name_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. Conflicts with `branch_name_pattern` as it only applied to rulesets with target `tag`. (see [below for nested schema](#rulestag_name_pattern))

- `required_code_scanning` - (Optional) (Block List, Max: 1) Define which tools must provide code scanning results before the reference is updated. When configured, code scanning must be enabled and have results for both the commit and the reference being updated. Multiple code scanning tools can be specified. (see [below for nested schema](#rulesrequired_code_scanning))

- `file_path_restriction` - (Optional) (Block List, Max 1) Parameters to be used for the file_path_restriction rule. When enabled restricts access to files within the repository. (See [below for nested schema](#rulesfile_path_restriction))

- `max_file_size` - (Optional) (Block List, Max 1) Parameters to be used for the max_file_size rule. When enabled restricts the maximum size of a file that can be pushed to the repository. (See [below for nested schema](#rulesmax_file_size))

- `max_file_path_length` - (Optional) (Block List, Max: 1) Prevent commits that include file paths that exceed a specified character limit from being pushed to the commit graph. This rule only applies to rulesets with target `push`. (see [below for nested schema](#rulesmax_file_path_length))

- `file_extension_restriction` - (Optional) (Block List, Max: 1) Prevent commits that include files with specified file extensions from being pushed to the commit graph. This rule only applies to rulesets with target `push`. (see [below for nested schema](#rulesfile_extension_restriction))
- `update` - (Optional) (Boolean) Only allow users with bypass permission to update matching refs.

- `update_allows_fetch_and_merge` - (Optional) (Boolean) Branch can pull changes from its upstream repository. This is only applicable to forked repositories. Requires `update` to be set to `true`. Note: behaviour is affected by a known bug on the GitHub side which may cause issues when using this parameter.

#### rules.branch_name_pattern ####

- `operator` - (Required) (String) The operator to use for matching. Can be one of: `starts_with`, `ends_with`, `contains`, `regex`.

- `pattern` - (Required) (String) The pattern to match with.

- `name` - (Optional) (String) How this rule will appear to users.

- `negate` - (Optional) (Boolean) If true, the rule will fail if the pattern matches.

#### rules.commit_author_email_pattern ####

- `operator` - (Required) (String) The operator to use for matching. Can be one of: `starts_with`, `ends_with`, `contains`, `regex`.

- `pattern` - (Required) (String) The pattern to match with.

- `name` - (Optional) (String) How this rule will appear to users.

- `negate` - (Optional) (Boolean) If true, the rule will fail if the pattern matches.

#### rules.commit_message_pattern ####

- `operator` - (Required) (String) The operator to use for matching. Can be one of: `starts_with`, `ends_with`, `contains`, `regex`.

- `pattern` - (Required) (String) The pattern to match with.

- `name` - (Optional) (String) How this rule will appear to users.

- `negate` - (Optional) (Boolean) If true, the rule will fail if the pattern matches.

#### rules.committer_email_pattern ####

- `operator` - (Required) (String) The operator to use for matching. Can be one of: `starts_with`, `ends_with`, `contains`, `regex`.

- `pattern` - (Required) (String) The pattern to match with.

- `name` - (Optional) (String) How this rule will appear to users.

- `negate` - (Optional) (Boolean) If true, the rule will fail if the pattern matches.

#### rules.merge_queue ####

- `check_response_timeout_minutes` - (Optional) (Number) Maximum time for a required status check to report a conclusion. After this much time has elapsed, checks that have not reported a conclusion will be assumed to have failed. Defaults to `60`.

- `grouping_strategy` - (Optional) (String) When set to `ALLGREEN`, the merge commit created by merge queue for each PR in the group must pass all required checks to merge. When set to `HEADGREEN`, only the commit at the head of the merge group, i.e. the commit containing changes from all of the PRs in the group, must pass its required checks to merge. Can be one of: `ALLGREEN`, `HEADGREEN`. Defaults to `ALLGREEN`.

- `max_entries_to_build` - (Optional) (Number) Limit the number of queued pull requests requesting checks and workflow runs at the same time. Defaults to `5`.

- `max_entries_to_merge` - (Optional) (Number) Limit the number of queued pull requests that will be merged together in a group. Defaults to `5`.

- `merge_method` - (Optional) (String) Method to use when merging changes from queued pull requests. Can be one of: `MERGE`, `SQUASH`, `REBASE`. Defaults to `MERGE`.

- `min_entries_to_merge` - (Optional) (Number) The minimum number of PRs that will be merged together in a group. Defaults to `1`.

- `min_entries_to_merge_wait_minutes` - (Optional) (Number) The time merge queue should wait after the first PR is added to the queue for the minimum group size to be met. After this time has elapsed, the minimum group size will be ignored and a smaller group will be merged. Defaults to `5`.

#### rules.pull_request ####

- `allowed_merge_methods` - (Optional) (List of String, Min: 1) Array of merge methods to be allowed. Allowed values include `merge`, `squash`, and `rebase`. At least one must be enabled.
- `dismiss_stale_reviews_on_push` - (Optional) (Boolean) New, reviewable commits pushed will dismiss previous pull request review approvals. Defaults to `false`.
- `require_code_owner_review` - (Optional) (Boolean) Require an approving review in pull requests that modify files that have a designated code owner. Defaults to `false`.
- `require_last_push_approval` - (Optional) (Boolean) Whether the most recent reviewable push must be approved by someone other than the person who pushed it. Defaults to `false`.
- `required_approving_review_count` - (Optional) (Number) The number of approving reviews that are required before a pull request can be merged. Defaults to `0`.
- `required_review_thread_resolution` - (Optional) (Boolean) All conversations on code must be resolved before a pull request can be merged. Defaults to `false`.

#### rules.copilot_code_review ####

- `review_on_push` - (Optional) (Boolean) Copilot automatically reviews each new push to the pull request. Defaults to `false`.

- `review_draft_pull_requests` - (Optional) (Boolean) Copilot automatically reviews draft pull requests before they are marked as ready for review. Defaults to `false`.

- `allowed_merge_methods` - (Required) (List of String, Min: 1) Array of merge methods to be allowed. Allowed values include `merge`, `squash`, and `rebase`. At least one must be enabled.

- `required_reviewers` - (Optional) (Block List) Require specific reviewers to approve pull requests. Note: This feature is in beta. (see [below for nested schema](#rulespull_requestrequired_reviewers))

#### rules.pull_request.required_reviewers ####

- `reviewer` - (Required) (Block List, Max: 1) The reviewer that must review matching files. (see [below for nested schema](#rulespull_requestrequired_reviewersreviewer))

- `file_patterns` - (Required) (List of String) File patterns (fnmatch syntax) that this reviewer must approve.

- `minimum_approvals` - (Required) (Number) Minimum number of approvals required from this reviewer. Set to 0 to make approval optional.

#### rules.pull_request.required_reviewers.reviewer ####

- `id` - (Required) (Number) The ID of the reviewer (Team ID).

- `type` - (Required) (String) The type of reviewer. Currently only `Team` is supported.

#### rules.required_deployments ####

- `required_deployment_environments` - (Required) (List of String) The environments that must be successfully deployed to before branches can be merged.

#### rules.required_status_checks ####

- `required_check` - (Required) (Block Set, Min: 1) Status checks that are required. Several can be defined. (see [below for nested schema](#rulesrequired_status_checksrequired_check))

- `strict_required_status_checks_policy` - (Optional) (Boolean) Whether pull requests targeting a matching branch must be tested with the latest code. This setting will not take effect unless at least one status check is enabled. Defaults to `false`.

- `do_not_enforce_on_create` - (Optional) (Boolean) Allow repositories and branches to be created if a check would otherwise prohibit it. Defaults to `false`.

#### rules.required_status_checks.required_check ####

- `context` - (Required) (String) The status check context name that must be present on the commit.

- `integration_id` - (Optional) (Number) The optional integration ID that this status check must originate from. It's a GitHub App ID, which can be obtained by following instructions from the [Get an App API docs](https://docs.github.com/en/rest/apps/apps?apiVersion=2022-11-28#get-an-app).

#### rules.tag_name_pattern ####

- `operator` - (Required) (String) The operator to use for matching. Can be one of: `starts_with`, `ends_with`, `contains`, `regex`.

- `pattern` - (Required) (String) The pattern to match with.

- `name` - (Optional) (String) How this rule will appear to users.

- `negate` - (Optional) (Boolean) If true, the rule will fail if the pattern matches.

#### rules.required_code_scanning ####

- `required_code_scanning_tool` - (Required) (Block Set, Min: 1) Actions code scanning tools that are required. Multiple can be defined. (see [below for nested schema](#rulesrequired_code_scanningrequired_code_scanning_tool))

#### rules.required_code_scanning.required_code_scanning_tool ####

- `alerts_threshold` - (Required) (String) The severity level at which code scanning results that raise alerts block a reference update. Can be one of: `none`, `errors`, `errors_and_warnings`, `all`.

- `security_alerts_threshold` - (Required) (String) The severity level at which code scanning results that raise security alerts block a reference update. Can be one of: `none`, `critical`, `high_or_higher`, `medium_or_higher`, `all`.

- `tool` - (Required) (String) The name of a code scanning tool.

#### rules.file_path_restriction ####

- `restricted_file_paths` - (Required) (Block Set, Min: 1) The file paths that are restricted from being pushed to the commit graph.

#### rules.max_file_size ####

- `max_file_size` - (Required) (Integer) The maximum allowed size, in megabytes (MB), of a file. Valid range is 1-100 MB.

#### rules.max_file_path_length ####

- `max_file_path_length` - (Required) (Integer) The maximum number of characters allowed in file paths.

#### rules.file_extension_restriction ####

- `restricted_file_extensions` - (Required) (Block Set, Min: 1) The file extensions that are restricted from being pushed to the commit graph.

#### bypass_actors ####

- `actor_id` - (Optional) (Number) The ID of the actor that can bypass a ruleset. If `actor_type` is `Integration`, `actor_id` is a GitHub App ID. App ID can be obtained by following instructions from the [Get an App API docs](https://docs.github.com/en/rest/apps/apps?apiVersion=2022-11-28#get-an-app). Some actor types such as `DeployKey` do not have an ID.

- `actor_type` (String) The type of actor that can bypass a ruleset. Can be one of: `RepositoryRole`, `Team`, `Integration`, `OrganizationAdmin`, `DeployKey`.

- `bypass_mode` - (Optional) (String) When the specified actor can bypass the ruleset. pull_request means that an actor can only bypass rules on pull requests. Can be one of: `always`, `pull_request`, `exempt`.

~> Note: at the time of writing this, the following actor types correspond to the following actor IDs:

- `OrganizationAdmin` -> `1`
- `RepositoryRole` (This is the actor type, the following are the base repository roles and their associated IDs.)
    - `maintain` -> `2`
    - `write` -> `4`
    - `admin` -> `5`

#### conditions ####

- `ref_name` - (Optional) (Block List, Max: 1) Required for `branch` and `tag` targets. Must NOT be set for `push` targets. (see [below for nested schema](#conditionsref_name))

~> **Note:** For `push` targets, do not include `ref_name` in conditions. Push rulesets operate on file content, not on refs. The `conditions` block is optional for push targets.

#### conditions.ref_name ####

- `exclude` - (Required) (List of String) Array of ref names or patterns to exclude. The condition will not pass if any of these patterns match.

- `include` - (Required) (List of String) Array of ref names or patterns to include. One of these patterns must match for the condition to pass. Also accepts `~DEFAULT_BRANCH` to include the default branch or `~ALL` to include all branches.

## Attributes Reference

The following additional attributes are exported:

- `etag` (String)

- `node_id` (String) GraphQL global node id for use with v4 API.

- `ruleset_id` (Number) GitHub ID for the ruleset.

## Import

GitHub Repository Rulesets can be imported using the GitHub repository name and ruleset ID e.g.

`$ terraform import github_repository_ruleset.example example:12345`
