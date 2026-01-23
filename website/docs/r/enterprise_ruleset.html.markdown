---
layout: "github"
page_title: "github_enterprise_ruleset Resource - terraform-provider-github"
description: |-
  Creates a GitHub enterprise ruleset.
---

# github_enterprise_ruleset (Resource)

Creates a GitHub enterprise ruleset.

This resource allows you to create and manage rulesets on the enterprise level. When applied, a new ruleset will be created. When destroyed, that ruleset will be removed.

Enterprise rulesets allow you to manage rulesets across multiple organizations within your enterprise, providing centralized control over repository rules and policies.

## Example Usage

### Basic Branch Ruleset

```hcl
resource "github_enterprise_ruleset" "example" {
  enterprise_slug = "my-enterprise"
  name            = "example-branch-ruleset"
  target          = "branch"
  enforcement     = "active"

  conditions {
    organization_name {
      include = ["~ALL"]
      exclude = []
    }

    repository_name {
      include = ["~ALL"]
      exclude = []
    }

    ref_name {
      include = ["~DEFAULT_BRANCH"]
      exclude = []
    }
  }

  bypass_actors {
    actor_id    = 1
    actor_type  = "OrganizationAdmin"
    bypass_mode = "always"
  }

  rules {
    creation                = true
    update                  = true
    deletion                = true
    required_linear_history = true
    required_signatures     = true

    pull_request {
      required_approving_review_count = 2
      require_code_owner_review       = true
      require_last_push_approval      = true
      required_review_thread_resolution = true
    }

    required_status_checks {
      strict_required_status_checks_policy = true

      required_check {
        context = "ci/test"
      }

      required_check {
        context        = "ci/deploy"
        integration_id = 12345
      }
    }
  }
}
```

### Push Ruleset with File Restrictions

```hcl
resource "github_enterprise_ruleset" "push_restrictions" {
  enterprise_slug = "my-enterprise"
  name            = "push-restrictions"
  target          = "push"
  enforcement     = "active"

  conditions {
    organization_id {
      organization_ids = [123456, 789012]
    }

    repository_name {
      include = ["~ALL"]
      exclude = ["legacy-*"]
    }
  }

  rules {
    file_path_restriction {
      restricted_file_paths = [".github/workflows/*", "*.env", "secrets/*"]
    }

    max_file_size {
      max_file_size = 100
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

### Tag Ruleset with Pattern Matching

```hcl
resource "github_enterprise_ruleset" "tag_ruleset" {
  enterprise_slug = "my-enterprise"
  name            = "tag-naming-convention"
  target          = "tag"
  enforcement     = "active"

  conditions {
    organization_name {
      include = ["production-*"]
      exclude = []
    }

    repository_property {
      include = ["repository_tier:production"]
      exclude = []
    }

    ref_name {
      include = ["~ALL"]
      exclude = []
    }
  }

  rules {
    creation = false
    deletion = true

    tag_name_pattern {
      name     = "Semantic versioning"
      operator = "regex"
      pattern  = "^v[0-9]+\\.[0-9]+\\.[0-9]+$"
      negate   = false
    }
  }
}
```

### Enterprise Ruleset with Code Scanning Requirements

```hcl
resource "github_enterprise_ruleset" "security_requirements" {
  enterprise_slug = "my-enterprise"
  name            = "security-requirements"
  target          = "branch"
  enforcement     = "active"

  conditions {
    organization_name {
      include = ["~ALL"]
      exclude = []
    }

    repository_name {
      include = ["~ALL"]
      exclude = []
      protected = true
    }

    ref_name {
      include = ["main", "master"]
      exclude = []
    }
  }

  rules {
    required_code_scanning {
      required_code_scanning_tool {
        tool                      = "CodeQL"
        alerts_threshold          = "errors"
        security_alerts_threshold = "high_or_higher"
      }

      required_code_scanning_tool {
        tool                      = "Semgrep"
        alerts_threshold          = "all"
        security_alerts_threshold = "medium_or_higher"
      }
    }

    required_workflows {
      required_workflow {
        repository_id = 1234567
        path          = ".github/workflows/security-scan.yml"
        ref           = "main"
      }
    }
  }
}
```

### Enterprise Ruleset with Commit Pattern Enforcement

```hcl
resource "github_enterprise_ruleset" "commit_patterns" {
  enterprise_slug = "my-enterprise"
  name            = "commit-conventions"
  target          = "branch"
  enforcement     = "active"

  conditions {
    organization_id {
      organization_ids = [123456]
    }

    repository_name {
      include = ["~ALL"]
      exclude = []
    }

    ref_name {
      include = ["main", "develop"]
      exclude = []
    }
  }

  bypass_actors {
    actor_id    = 2
    actor_type  = "RepositoryRole"
    bypass_mode = "pull_request"
  }

  rules {
    commit_message_pattern {
      name     = "Conventional Commits"
      operator = "regex"
      pattern  = "^(feat|fix|docs|style|refactor|test|chore)(\\(.+\\))?: .+"
      negate   = false
    }

    commit_author_email_pattern {
      name     = "Corporate email required"
      operator = "ends_with"
      pattern  = "@example.com"
      negate   = false
    }

    committer_email_pattern {
      name     = "Corporate email required"
      operator = "ends_with"
      pattern  = "@example.com"
      negate   = false
    }
  }
}
```

### Repository Target Ruleset

```hcl
resource "github_enterprise_ruleset" "repository_management" {
  enterprise_slug = "my-enterprise"
  name            = "repository-management"
  target          = "repository"
  enforcement     = "active"

  bypass_actors {
    actor_id    = 1
    actor_type  = "OrganizationAdmin"
    bypass_mode = "always"
  }

  conditions {
    organization_name {
      include = ["~ALL"]
      exclude = []
    }

    repository_name {
      include = ["~ALL"]
      exclude = []
    }
  }

  rules {
    repository_creation = true
    repository_deletion = true
    repository_transfer = true

    repository_name {
      pattern = "^[a-z][a-z0-9-]*$"
      negate  = false
    }

    repository_visibility {
      internal = true
      private  = true
    }
  }
}
```

## Argument Reference

- `enterprise_slug` - (Required) (String) The slug of the enterprise.

- `name` - (Required) (String) The name of the ruleset.

- `target` - (Required) (String) Possible values are `branch`, `tag`, `push`, and `repository`. Note: The `push` and `repository` targets are in beta and are subject to change.

- `enforcement` - (Required) (String) Possible values for Enforcement are `disabled`, `active`, `evaluate`. Note: `evaluate` is currently only supported for owners of type `organization`.

- `rules` - (Required) (Block List, Min: 1, Max: 1) Rules within the ruleset. (see [below for nested schema](#rules))

- `bypass_actors` - (Optional) (Block List) The actors that can bypass the rules in this ruleset. (see [below for nested schema](#bypass_actors))

- `conditions` - (Optional) (Block List, Max: 1) Parameters for an enterprise ruleset condition. Enterprise rulesets must include organization targeting (organization_name or organization_id) and repository targeting (repository_name or repository_property). For branch and tag targets, ref_name is also required. (see [below for nested schema](#conditions))

### Rules

The `rules` block supports the following:

- `creation` - (Optional) (Boolean) Only allow users with bypass permission to create matching refs.

- `update` - (Optional) (Boolean) Only allow users with bypass permission to update matching refs.

- `deletion` - (Optional) (Boolean) Only allow users with bypass permissions to delete matching refs.

- `required_linear_history` - (Optional) (Boolean) Prevent merge commits from being pushed to matching branches.

- `required_signatures` - (Optional) (Boolean) Commits pushed to matching branches must have verified signatures.

- `non_fast_forward` - (Optional) (Boolean) Prevent users with push access from force pushing to branches.

- `pull_request` - (Optional) (Block List, Max: 1) Require all commits be made to a non-target branch and submitted via a pull request before they can be merged. (see [below for nested schema](#rulespull_request))

- `copilot_code_review` - (Optional) (Block List, Max: 1) Automatically request Copilot code review for new pull requests if the author has access to Copilot code review and their premium requests quota has not reached the limit. (see [below for nested schema](#rulescopilot_code_review))

- `required_status_checks` - (Optional) (Block List, Max: 1) Choose which status checks must pass before branches can be merged into a branch that matches this rule. (see [below for nested schema](#rulesrequired_status_checks))

- `required_workflows` - (Optional) (Block List, Max: 1) Define which Actions workflows must pass before changes can be merged into a branch matching the rule. (see [below for nested schema](#rulesrequired_workflows))

- `required_code_scanning` - (Optional) (Block List, Max: 1) Define which tools must provide code scanning results before the reference is updated. (see [below for nested schema](#rulesrequired_code_scanning))

- `branch_name_pattern` - (Optional) (Block List, Max: 1) Parameters to be used for the branch_name_pattern rule. Conflicts with `tag_name_pattern` as it only applies to rulesets with target `branch`. (see [below for nested schema](#rulesbranch_name_pattern))

- `tag_name_pattern` - (Optional) (Block List, Max: 1) Parameters to be used for the tag_name_pattern rule. Conflicts with `branch_name_pattern` as it only applies to rulesets with target `tag`. (see [below for nested schema](#rulestag_name_pattern))

- `commit_author_email_pattern` - (Optional) (Block List, Max: 1) Parameters to be used for the commit_author_email_pattern rule. (see [below for nested schema](#rulescommit_author_email_pattern))

- `commit_message_pattern` - (Optional) (Block List, Max: 1) Parameters to be used for the commit_message_pattern rule. (see [below for nested schema](#rulescommit_message_pattern))

- `committer_email_pattern` - (Optional) (Block List, Max: 1) Parameters to be used for the committer_email_pattern rule. (see [below for nested schema](#rulescommitter_email_pattern))

- `file_path_restriction` - (Optional) (Block List, Max: 1) Prevent commits that include changes to specified file paths from being pushed to the commit graph. This rule only applies to rulesets with target `push`. (see [below for nested schema](#rulesfile_path_restriction))

- `max_file_size` - (Optional) (Block List, Max: 1) Prevent commits that include files with a specified file size from being pushed to the commit graph. This rule only applies to rulesets with target `push`. (see [below for nested schema](#rulesmax_file_size))

- `max_file_path_length` - (Optional) (Block List, Max: 1) Prevent commits that include file paths that exceed a specified character limit from being pushed to the commit graph. This rule only applies to rulesets with target `push`. (see [below for nested schema](#rulesmax_file_path_length))

- `file_extension_restriction` - (Optional) (Block List, Max: 1) Prevent commits that include files with specified file extensions from being pushed to the commit graph. This rule only applies to rulesets with target `push`. (see [below for nested schema](#rulesfile_extension_restriction))

- `repository_creation` - (Optional) (Boolean) Only allow users with bypass permission to create repositories. Only valid for `repository` target.

- `repository_deletion` - (Optional) (Boolean) Only allow users with bypass permission to delete repositories. Only valid for `repository` target.

- `repository_transfer` - (Optional) (Boolean) Only allow users with bypass permission to transfer repositories. Only valid for `repository` target.

- `repository_name` - (Optional) (Block List, Max: 1) Restrict repository names to match specified patterns. Only valid for `repository` target. (see [below for nested schema](#rulesrepository_name))

- `repository_visibility` - (Optional) (Block List, Max: 1) Restrict repository visibility changes. Only valid for `repository` target. (see [below for nested schema](#rulesrepository_visibility))

#### rules.pull_request

- `dismiss_stale_reviews_on_push` - (Optional) (Boolean) New, reviewable commits pushed will dismiss previous pull request review approvals. Defaults to `false`.

- `require_code_owner_review` - (Optional) (Boolean) Require an approving review in pull requests that modify files that have a designated code owner. Defaults to `false`.

- `require_last_push_approval` - (Optional) (Boolean) Whether the most recent reviewable push must be approved by someone other than the person who pushed it. Defaults to `false`.

- `required_approving_review_count` - (Optional) (Number) The number of approving reviews that are required before a pull request can be merged. Defaults to `0`.

- `required_review_thread_resolution` - (Optional) (Boolean) All conversations on code must be resolved before a pull request can be merged. Defaults to `false`.

- `allowed_merge_methods` - (Optional) (List of String, Min: 1) The merge methods allowed for pull requests. Possible values are `merge`, `squash`, and `rebase`.

#### rules.copilot_code_review

- `review_on_push` - (Optional) (Boolean) Copilot automatically reviews each new push to the pull request. Defaults to `false`.

- `review_draft_pull_requests` - (Optional) (Boolean) Copilot automatically reviews draft pull requests before they are marked as ready for review. Defaults to `false`.

#### rules.required_status_checks

- `required_check` - (Required) (Block Set, Min: 1) Status checks that are required. Several can be defined. (see [below for nested schema](#rulesrequired_status_checksrequired_check))

- `strict_required_status_checks_policy` - (Optional) (Boolean) Whether pull requests targeting a matching branch must be tested with the latest code. This setting will not take effect unless at least one status check is enabled. Defaults to `false`.

- `do_not_enforce_on_create` - (Optional) (Boolean) Allow repositories and branches to be created if a check would otherwise prohibit it. Defaults to `false`.

#### rules.required_status_checks.required_check

- `context` - (Required) (String) The status check context name that must be present on the commit.

- `integration_id` - (Optional) (Number) The optional integration ID that this status check must originate from.

- `do_not_enforce_on_create` - (Optional) (Boolean) Allow repositories and branches to be created if a check would otherwise prohibit it. Defaults to `false`.

#### rules.required_workflows

- `do_not_enforce_on_create` - (Optional) (Boolean) Allow repositories and branches to be created if a check would otherwise prohibit it. Defaults to `false`.

- `required_workflow` - (Required) (Block Set, Min: 1) Actions workflows that are required. Multiple can be defined. (see [below for nested schema](#rulesrequired_workflowsrequired_workflow))

#### rules.required_workflows.required_workflow

- `repository_id` - (Required) (Number) The ID of the repository. Names, full names and repository URLs are not supported.

- `path` - (Required) (String) The path to the YAML definition file of the workflow.

- `ref` - (Optional) (String) The optional ref from which to fetch the workflow. Defaults to `master`.

#### rules.required_code_scanning

- `required_code_scanning_tool` - (Required) (Block Set, Min: 1) Code scanning tools that are required. Multiple can be defined. (see [below for nested schema](#rulesrequired_code_scanningrequired_code_scanning_tool))

#### rules.required_code_scanning.required_code_scanning_tool

- `alerts_threshold` - (Required) (String) The severity level at which code scanning results that raise alerts block a reference update. Can be one of: `none`, `errors`, `errors_and_warnings`, `all`.

- `security_alerts_threshold` - (Required) (String) The severity level at which code scanning results that raise security alerts block a reference update. Can be one of: `none`, `critical`, `high_or_higher`, `medium_or_higher`, `all`.

- `tool` - (Required) (String) The name of a code scanning tool.

#### rules.branch_name_pattern

- `operator` - (Required) (String) The operator to use for matching. Can be one of: `starts_with`, `ends_with`, `contains`, `regex`.

- `pattern` - (Required) (String) The pattern to match with.

- `name` - (Optional) (String) How this rule will appear to users.

- `negate` - (Optional) (Boolean) If true, the rule will fail if the pattern matches.

#### rules.tag_name_pattern

- `operator` - (Required) (String) The operator to use for matching. Can be one of: `starts_with`, `ends_with`, `contains`, `regex`.

- `pattern` - (Required) (String) The pattern to match with.

- `name` - (Optional) (String) How this rule will appear to users.

- `negate` - (Optional) (Boolean) If true, the rule will fail if the pattern matches.

#### rules.commit_author_email_pattern

- `operator` - (Required) (String) The operator to use for matching. Can be one of: `starts_with`, `ends_with`, `contains`, `regex`.

- `pattern` - (Required) (String) The pattern to match with.

- `name` - (Optional) (String) How this rule will appear to users.

- `negate` - (Optional) (Boolean) If true, the rule will fail if the pattern matches.

#### rules.commit_message_pattern

- `operator` - (Required) (String) The operator to use for matching. Can be one of: `starts_with`, `ends_with`, `contains`, `regex`.

- `pattern` - (Required) (String) The pattern to match with.

- `name` - (Optional) (String) How this rule will appear to users.

- `negate` - (Optional) (Boolean) If true, the rule will fail if the pattern matches.

#### rules.committer_email_pattern

- `operator` - (Required) (String) The operator to use for matching. Can be one of: `starts_with`, `ends_with`, `contains`, `regex`.

- `pattern` - (Required) (String) The pattern to match with.

- `name` - (Optional) (String) How this rule will appear to users.

- `negate` - (Optional) (Boolean) If true, the rule will fail if the pattern matches.

#### rules.file_path_restriction

- `restricted_file_paths` - (Required) (List of String, Min: 1) The file paths that are restricted from being pushed to the commit graph.

#### rules.max_file_size

- `max_file_size` - (Required) (Number) The maximum allowed size, in megabytes (MB), of a file. Valid range is 1-100 MB.

#### rules.max_file_path_length

- `max_file_path_length` - (Required) (Number) The maximum number of characters allowed in file paths.

#### rules.file_extension_restriction

- `restricted_file_extensions` - (Required) (List of String, Min: 1) The file extensions that are restricted from being pushed to the commit graph.

#### rules.repository_name

- `pattern` - (Required) (String) The pattern to match repository names against.

- `negate` - (Optional) (Boolean) If true, the rule will fail if the pattern matches. Defaults to `false`.

#### rules.repository_visibility

- `internal` - (Optional) (Boolean) Allow internal visibility for repositories. Defaults to `false`.

- `private` - (Optional) (Boolean) Allow private visibility for repositories. Defaults to `false`.

### bypass_actors

- `actor_id` - (Optional) (Number) The ID of the actor that can bypass a ruleset. When `actor_type` is `OrganizationAdmin`, this should be set to `1`. Some resources such as DeployKey do not have an ID and this should be omitted.

- `actor_type` - (Required) (String) The type of actor that can bypass a ruleset. Can be one of: `Integration`, `OrganizationAdmin`, `RepositoryRole`, `Team`, `DeployKey`.

- `bypass_mode` - (Required) (String) When the specified actor can bypass the ruleset. pull_request means that an actor can only bypass rules on pull requests. Can be one of: `always`, `pull_request`, `exempt`.

~>Note: at the time of writing this, the following actor types correspond to the following actor IDs:

- `OrganizationAdmin` -> `1`
- `RepositoryRole` (This is the actor type, the following are the base repository roles and their associated IDs.)
  - `maintain` -> `2`
  - `write` -> `4`
  - `admin` -> `5`

### conditions

Enterprise rulesets require targeting both organizations and repositories. At least one organization targeting condition (`organization_name` or `organization_id`) and one repository targeting condition (`repository_name`, `repository_id`, or `repository_property`) must be specified. For `branch` and `tag` targets, `ref_name` is also required.

- `organization_name` - (Optional) (Block List, Max: 1) Conditions for organization names that the ruleset targets. Conflicts with `organization_id`. (see [below for nested schema](#conditionsorganization_name))

- `organization_id` - (Optional) (Block List, Max: 1) Conditions for organization IDs that the ruleset targets. Conflicts with `organization_name`. (see [below for nested schema](#conditionsorganization_id))

- `repository_name` - (Optional) (Block List, Max: 1) Conditions for repository names that the ruleset targets. (see [below for nested schema](#conditionsrepository_name))

- `repository_id` - (Optional) (Block List, Max: 1) Conditions for repository IDs that the ruleset targets. (see [below for nested schema](#conditionsrepository_id))

- `repository_property` - (Optional) (Block List, Max: 1) Conditions for repository properties that the ruleset targets. (see [below for nested schema](#conditionsrepository_property))

- `ref_name` - (Optional) (Block List, Max: 1) Conditions for ref names that the ruleset targets. Required for `branch` and `tag` targets. (see [below for nested schema](#conditionsref_name))

#### conditions.organization_name

- `include` - (Required) (List of String) Array of organization name patterns to include. One of these patterns must match for the condition to pass. Also accepts `~ALL` to include all organizations.

- `exclude` - (Required) (List of String) Array of organization name patterns to exclude. The condition will not pass if any of these patterns match.

#### conditions.organization_id

- `organization_ids` - (Required) (List of Number) Array of organization IDs to target. One of these IDs must match for the condition to pass.

#### conditions.repository_name

- `include` - (Required) (List of String) Array of repository name patterns to include. One of these patterns must match for the condition to pass. Also accepts `~ALL` to include all repositories.

- `exclude` - (Required) (List of String) Array of repository name patterns to exclude. The condition will not pass if any of these patterns match.

- `protected` - (Optional) (Boolean) Whether to target only protected repositories. Defaults to `false`.

#### conditions.repository_id

- `repository_ids` - (Required) (List of Number) Array of repository IDs to target. One of these IDs must match for the condition to pass.

#### conditions.repository_property

- `include` - (Required) (List of String) The repository properties to include. All properties must match for the condition to pass. Repository properties are in the format `property_name:property_value`.

- `exclude` - (Required) (List of String) The repository properties to exclude. Repository properties are in the format `property_name:property_value`.

#### conditions.ref_name

- `include` - (Required) (List of String) Array of ref names or patterns to include. One of these patterns must match for the condition to pass. Also accepts `~DEFAULT_BRANCH` to include the default branch or `~ALL` to include all branches.

- `exclude` - (Required) (List of String) Array of ref names or patterns to exclude. The condition will not pass if any of these patterns match.

## Attributes Reference

The following additional attributes are exported:

- `etag` - (String) The etag of the ruleset.

- `node_id` - (String) GraphQL global node id for use with v4 API.

- `ruleset_id` - (Number) GitHub ID for the ruleset.

## Import

GitHub Enterprise Rulesets can be imported using the enterprise slug and ruleset ID in the format `{enterprise_slug}/{ruleset_id}`, e.g.

```sh
terraform import github_enterprise_ruleset.example my-enterprise/12345
```
