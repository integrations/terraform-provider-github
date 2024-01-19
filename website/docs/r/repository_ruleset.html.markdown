---
layout: "github"
page_title: "github_repository_ruleset Resource - terraform-provider-github"
description: |-
  Creates a GitHub repository ruleset.
---

# github_repository_ruleset (Resource)

Creates a GitHub repository ruleset.

This resource allows you to create and manage rulesets on the repository level. When applied, a new ruleset will be created. When destroyed, that ruleset will be removed.

## Example Usage

```
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

    
  }
}
```

## Argument Reference

* `enforcement` - (Required) (String) Possible values for Enforcement are `disabled`, `active`, `evaluate`. Note: `evaluate` is currently only supported for owners of type `organization`.

* `name` - (Required) (String) The name of the ruleset.

* `rules` - (Required) (Block List, Min: 1, Max: 1) Rules within the ruleset. (see [below for nested schema](#rules))

* `target` - (Required) (String) Possible values are `branch` and `tag`.

* `bypass_actors` - (Optional) (Block List) The actors that can bypass the rules in this ruleset. (see [below for nested schema](#bypass_actors))

* `conditions` - (Optional) (Block List, Max: 1) Parameters for a repository ruleset ref name condition. (see [below for nested schema](#conditions))

* `repository` - (Optional) (String) Name of the repository to apply rulset to.

#### Rules ####

The `rules` block supports the following:


* `branch_name_pattern` - (Optional) (Block List, Max: 1) Parameters to be used for the branch_name_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. Conflicts with `tag_name_pattern` as it only applied to rulesets with target `branch`. (see [below for nested schema](#rules.branch_name_pattern))

* `commit_author_email_pattern` - (Optional) (Block List, Max: 1) Parameters to be used for the commit_author_email_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. (see [below for nested schema](#rules.commit_author_email_pattern))

* `commit_message_pattern` - (Optional) (Block List, Max: 1) Parameters to be used for the commit_message_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. (see [below for nested schema](#rules.commit_message_pattern))

* `committer_email_pattern` - (Optional) (Block List, Max: 1) Parameters to be used for the committer_email_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. (see [below for nested schema](#rules.committer_email_pattern))

* `creation` - (Optional) (Boolean) Only allow users with bypass permission to create matching refs.

* `deletion` - (Optional) (Boolean) Only allow users with bypass permissions to delete matching refs.

* `non_fast_forward` - (Optional) (Boolean) Prevent users with push access from force pushing to branches.

* `pull_request` - (Optional) (Block List, Max: 1) Require all commits be made to a non-target branch and submitted via a pull request before they can be merged. (see [below for nested schema](#rules.pull_request))

* `required_deployments` - (Optional) (Block List, Max: 1) Choose which environments must be successfully deployed to before branches can be merged into a branch that matches this rule. (see [below for nested schema](#rules.required_deployments))

* `required_linear_history` - (Optional) (Boolean) Prevent merge commits from being pushed to matching branches.

* `required_signatures` - (Optional) (Boolean) Commits pushed to matching branches must have verified signatures.

* `required_status_checks` - (Optional) (Block List, Max: 1) Choose which status checks must pass before branches can be merged into a branch that matches this rule. When enabled, commits must first be pushed to another branch, then merged or pushed directly to a branch that matches this rule after status checks have passed. (see [below for nested schema](#rules.required_status_checks))

* `tag_name_pattern` - (Optional) (Block List, Max: 1) Parameters to be used for the tag_name_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. Conflicts with `branch_name_pattern` as it only applied to rulesets with target `tag`. (see [below for nested schema](#rules.tag_name_pattern))

* `update` - (Optional) (Boolean) Only allow users with bypass permission to update matching refs.
  
* `update_allows_fetch_and_merge` - (Optional) (Boolean) Branch can pull changes from its upstream repository. This is only applicable to forked repositories. Requires `update` to be set to `true`. Note: behaviour is affected by a known bug on the GitHub side which may cause issues when using this parameter.

#### rules.branch_name_pattern ####

* `operator` - (Required) (String) The operator to use for matching. Can be one of: `starts_with`, `ends_with`, `contains`, `regex`.

* `pattern` - (Required) (String) The pattern to match with.

* `name` - (Optional) (String) How this rule will appear to users.

* `negate` - (Optional) (Boolean) If true, the rule will fail if the pattern matches.


#### rules.commit_author_email_pattern ####

* `operator` - (Required) (String) The operator to use for matching. Can be one of: `starts_with`, `ends_with`, `contains`, `regex`.

* `pattern` - (Required) (String) The pattern to match with.

* `name` - (Optional) (String) How this rule will appear to users.

* `negate` - (Optional) (Boolean) If true, the rule will fail if the pattern matches.


#### rules.commit_message_pattern ####

* `operator` - (Required) (String) The operator to use for matching. Can be one of: `starts_with`, `ends_with`, `contains`, `regex`.

* `pattern` - (Required) (String) The pattern to match with.

* `name` - (Optional) (String) How this rule will appear to users.

* `negate` - (Optional) (Boolean) If true, the rule will fail if the pattern matches.


#### rules.committer_email_pattern ####

* `operator` - (Required) (String) The operator to use for matching. Can be one of: `starts_with`, `ends_with`, `contains`, `regex`.

* `pattern` - (Required) (String) The pattern to match with.

* `name` - (Optional) (String) How this rule will appear to users.

* `negate` - (Optional) (Boolean) If true, the rule will fail if the pattern matches.


#### rules.pull_request ####

* `dismiss_stale_reviews_on_push` - (Optional) (Boolean) New, reviewable commits pushed will dismiss previous pull request review approvals. Defaults to `false`.

* `require_code_owner_review` - (Optional) (Boolean) Require an approving review in pull requests that modify files that have a designated code owner. Defaults to `false`.

* `require_last_push_approval` - (Optional) (Boolean) Whether the most recent reviewable push must be approved by someone other than the person who pushed it. Defaults to `false`.

* `required_approving_review_count` - (Optional) (Number) The number of approving reviews that are required before a pull request can be merged. Defaults to `0`.

* `required_review_thread_resolution` - (Optional) (Boolean) All conversations on code must be resolved before a pull request can be merged. Defaults to `false`.


#### rules.required_deployments ####

* `required_deployment_environments` - (Required) (List of String) The environments that must be successfully deployed to before branches can be merged.


#### rules.required_status_checks ####

* `required_check` - (Required) (Block Set, Min: 1) Status checks that are required. Several can be defined. (see [below for nested schema](#rules.required_status_checks.required_check))

* `strict_required_status_checks_policy` - (Optional) (Boolean) Whether pull requests targeting a matching branch must be tested with the latest code. This setting will not take effect unless at least one status check is enabled. Defaults to `false`.

#### rules.required_status_checks.required_check ####

* `context` - (Required) (String) The status check context name that must be present on the commit.

* `integration_id` - (Optional) (Number) The optional integration ID that this status check must originate from.



#### rules.tag_name_pattern ####

* `operator` - (Required) (String) The operator to use for matching. Can be one of: `starts_with`, `ends_with`, `contains`, `regex`.

* `pattern` - (Required) (String) The pattern to match with.

* `name` - (Optional) (String) How this rule will appear to users.

* `negate` - (Optional) (Boolean) If true, the rule will fail if the pattern matches.



#### bypass_actors ####

* `actor_id` - (Required) (Number) The ID of the actor that can bypass a ruleset.

* `actor_type` (String) The type of actor that can bypass a ruleset. Can be one of: `RepositoryRole`, `Team`, `Integration`, `OrganizationAdmin`.

* `bypass_mode` - (Optional) (String) When the specified actor can bypass the ruleset. pull_request means that an actor can only bypass rules on pull requests. Can be one of: `always`, `pull_request`.

~> Note: at the time of writing this, the following actor types correspond to the following actor IDs:
* `OrganizationAdmin` -> `1`
* `RepositoryRole` (This is the actor type, the following are the base repository roles and their associated IDs.)
  * `maintain` -> `2`
  * `write` -> `4`
  * `admin` -> `5`


#### conditions ####

* `ref_name` - (Required) (Block List, Min: 1, Max: 1) (see [below for nested schema](#conditions.ref_name))

#### conditions.ref_name ####

* `exclude` - (Required) (List of String) Array of ref names or patterns to exclude. The condition will not pass if any of these patterns match.

* `include` - (Required) (List of String) Array of ref names or patterns to include. One of these patterns must match for the condition to pass. Also accepts `~DEFAULT_BRANCH` to include the default branch or `~ALL` to include all branches.

## Attributes Reference

The following additional attributes are exported:


* `etag` (String)

* `node_id` (String) GraphQL global node id for use with v4 API.

* `ruleset_id` (Number) GitHub ID for the ruleset.


## Import

GitHub Repository Rulesets can be imported using the GitHub repository name and ruleset ID e.g.

`$ terraform import github_repository_ruleset.example example:12345`
