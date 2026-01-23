---
layout: "github"
page_title: "github_enterprise_ruleset Data Source - terraform-provider-github"
description: |-
  Use this data source to retrieve information about a GitHub enterprise ruleset.
---

# github_enterprise_ruleset (Data Source)

Use this data source to retrieve information about a GitHub enterprise ruleset.

## Example Usage

```hcl
data "github_enterprise_ruleset" "example" {
  enterprise_slug = "my-enterprise"
  ruleset_id      = 12345
}
```

## Argument Reference

- `enterprise_slug` - (Required) (String) The slug of the enterprise.

- `ruleset_id` - (Required) (Number) The ID of the ruleset to retrieve.

## Attributes Reference

- `name` - (String) The name of the ruleset.

- `target` - (String) The target of the ruleset. Possible values are `branch`, `tag`, `push`, and `repository`.

- `enforcement` - (String) The enforcement level of the ruleset. Possible values are `disabled`, `active`, and `evaluate`.

- `node_id` - (String) GraphQL global node id for use with v4 API.

- `bypass_actors` - (List) The actors that can bypass the rules in this ruleset. (see [below for nested schema](#bypass_actors))

- `conditions` - (List) Parameters for an enterprise ruleset condition. (see [below for nested schema](#conditions))

- `rules` - (List) Rules within the ruleset. (see [below for nested schema](#rules))

### bypass_actors

- `actor_id` - (Number) The ID of the actor that can bypass a ruleset.

- `actor_type` - (String) The type of actor that can bypass a ruleset.

- `bypass_mode` - (String) When the specified actor can bypass the ruleset.

### conditions

- `organization_name` - (List) Conditions for organization names that the ruleset targets. (see [below for nested schema](#conditionsorganization_name))

- `organization_id` - (List) Conditions for organization IDs that the ruleset targets. (see [below for nested schema](#conditionsorganization_id))

- `repository_name` - (List) Conditions for repository names that the ruleset targets. (see [below for nested schema](#conditionsrepository_name))

- `repository_id` - (List) Conditions for repository IDs that the ruleset targets. (see [below for nested schema](#conditionsrepository_id))

- `repository_property` - (List) Conditions for repository properties that the ruleset targets. (see [below for nested schema](#conditionsrepository_property))

- `ref_name` - (List) Conditions for ref names that the ruleset targets. (see [below for nested schema](#conditionsref_name))

### conditions.organization_name

- `include` - (List of String) Array of organization name patterns to include.

- `exclude` - (List of String) Array of organization name patterns to exclude.

### conditions.organization_id

- `organization_ids` - (List of Number) Array of organization IDs to target.

### conditions.repository_name

- `include` - (List of String) Array of repository name patterns to include.

- `exclude` - (List of String) Array of repository name patterns to exclude.

- `protected` - (Boolean) Whether to target only protected repositories.

### conditions.repository_id

- `repository_ids` - (List of Number) Array of repository IDs to target.

### conditions.repository_property

- `include` - (List of String) The repository properties to include. All properties must match for the condition to pass.

- `exclude` - (List of String) The repository properties to exclude.

### conditions.ref_name

- `include` - (List of String) Array of ref names or patterns to include.

- `exclude` - (List of String) Array of ref names or patterns to exclude.

### rules

- `creation` - (Boolean) Only allow users with bypass permission to create matching refs.

- `update` - (Boolean) Only allow users with bypass permission to update matching refs.

- `deletion` - (Boolean) Only allow users with bypass permissions to delete matching refs.

- `required_linear_history` - (Boolean) Prevent merge commits from being pushed to matching branches.

- `required_signatures` - (Boolean) Commits pushed to matching branches must have verified signatures.

- `non_fast_forward` - (Boolean) Prevent users with push access from force pushing to branches.

- `pull_request` - (List) Require all commits be made to a non-target branch and submitted via a pull request. (see [below for nested schema](#rulespull_request))

- `copilot_code_review` - (List) Automatically request Copilot code review for new pull requests. (see [below for nested schema](#rulescopilot_code_review))

- `required_status_checks` - (List) Status checks that are required. (see [below for nested schema](#rulesrequired_status_checks))

- `required_workflows` - (List) Actions workflows that are required. (see [below for nested schema](#rulesrequired_workflows))

- `required_code_scanning` - (List) Code scanning tools that are required. (see [below for nested schema](#rulesrequired_code_scanning))

- `branch_name_pattern` - (List) Parameters for the branch_name_pattern rule. (see [below for nested schema](#rulesbranch_name_pattern))

- `tag_name_pattern` - (List) Parameters for the tag_name_pattern rule. (see [below for nested schema](#rulestag_name_pattern))

- `commit_author_email_pattern` - (List) Parameters for the commit_author_email_pattern rule. (see [below for nested schema](#rulescommit_author_email_pattern))

- `commit_message_pattern` - (List) Parameters for the commit_message_pattern rule. (see [below for nested schema](#rulescommit_message_pattern))

- `committer_email_pattern` - (List) Parameters for the committer_email_pattern rule. (see [below for nested schema](#rulescommitter_email_pattern))

- `file_path_restriction` - (List) File path restrictions for push rulesets. (see [below for nested schema](#rulesfile_path_restriction))

- `max_file_size` - (List) Maximum file size restrictions for push rulesets. (see [below for nested schema](#rulesmax_file_size))

- `max_file_path_length` - (List) Maximum file path length restrictions for push rulesets. (see [below for nested schema](#rulesmax_file_path_length))

- `file_extension_restriction` - (List) File extension restrictions for push rulesets. (see [below for nested schema](#rulesfile_extension_restriction))

- `repository_creation` - (Boolean) Only allow users with bypass permission to create repositories. Only valid for `repository` target.

- `repository_deletion` - (Boolean) Only allow users with bypass permission to delete repositories. Only valid for `repository` target.

- `repository_transfer` - (Boolean) Only allow users with bypass permission to transfer repositories. Only valid for `repository` target.

- `repository_name` - (List) Restrict repository names to match specified patterns. Only valid for `repository` target. (see [below for nested schema](#rulesrepository_name))

- `repository_visibility` - (List) Restrict repository visibility changes. Only valid for `repository` target. (see [below for nested schema](#rulesrepository_visibility))

### rules.pull_request

- `dismiss_stale_reviews_on_push` - (Boolean) New, reviewable commits pushed will dismiss previous pull request review approvals.

- `require_code_owner_review` - (Boolean) Require an approving review in pull requests that modify files that have a designated code owner.

- `require_last_push_approval` - (Boolean) Whether the most recent reviewable push must be approved by someone other than the person who pushed it.

- `required_approving_review_count` - (Number) The number of approving reviews that are required before a pull request can be merged.

- `required_review_thread_resolution` - (Boolean) All conversations on code must be resolved before a pull request can be merged.

- `allowed_merge_methods` - (List of String) The merge methods allowed for pull requests. Possible values are `merge`, `squash`, and `rebase`.

### rules.copilot_code_review

- `review_on_push` - (Boolean) Copilot automatically reviews each new push to the pull request.

- `review_draft_pull_requests` - (Boolean) Copilot automatically reviews draft pull requests before they are marked as ready for review.

### rules.required_status_checks

- `required_check` - (List) Status checks that are required. (see [below for nested schema](#rulesrequired_status_checksrequired_check))

- `strict_required_status_checks_policy` - (Boolean) Whether pull requests targeting a matching branch must be tested with the latest code.

- `do_not_enforce_on_create` - (Boolean) Allow repositories and branches to be created if a check would otherwise prohibit it.

### rules.required_status_checks.required_check

- `context` - (String) The status check context name that must be present on the commit.

- `integration_id` - (Number) The optional integration ID that this status check must originate from.

- `do_not_enforce_on_create` - (Boolean) Allow repositories and branches to be created if a check would otherwise prohibit it.

### rules.required_workflows

- `required_workflow` - (List) Actions workflows that are required. (see [below for nested schema](#rulesrequired_workflowsrequired_workflow))

- `do_not_enforce_on_create` - (Boolean) Allow repositories and branches to be created if a check would otherwise prohibit it.

### rules.required_workflows.required_workflow

- `repository_id` - (Number) The ID of the repository.

- `path` - (String) The path to the YAML definition file of the workflow.

- `ref` - (String) The ref from which to fetch the workflow.

### rules.required_code_scanning

- `required_code_scanning_tool` - (List) Code scanning tools that are required. (see [below for nested schema](#rulesrequired_code_scanningrequired_code_scanning_tool))

### rules.required_code_scanning.required_code_scanning_tool

- `alerts_threshold` - (String) The severity level at which code scanning results that raise alerts block a reference update.

- `security_alerts_threshold` - (String) The severity level at which code scanning results that raise security alerts block a reference update.

- `tool` - (String) The name of a code scanning tool.

### rules.branch_name_pattern

- `operator` - (String) The operator to use for matching.

- `pattern` - (String) The pattern to match with.

- `name` - (String) How this rule will appear to users.

- `negate` - (Boolean) If true, the rule will fail if the pattern matches.

### rules.tag_name_pattern

- `operator` - (String) The operator to use for matching.

- `pattern` - (String) The pattern to match with.

- `name` - (String) How this rule will appear to users.

- `negate` - (Boolean) If true, the rule will fail if the pattern matches.

### rules.commit_author_email_pattern

- `operator` - (String) The operator to use for matching.

- `pattern` - (String) The pattern to match with.

- `name` - (String) How this rule will appear to users.

- `negate` - (Boolean) If true, the rule will fail if the pattern matches.

### rules.commit_message_pattern

- `operator` - (String) The operator to use for matching.

- `pattern` - (String) The pattern to match with.

- `name` - (String) How this rule will appear to users.

- `negate` - (Boolean) If true, the rule will fail if the pattern matches.

### rules.committer_email_pattern

- `operator` - (String) The operator to use for matching.

- `pattern` - (String) The pattern to match with.

- `name` - (String) How this rule will appear to users.

- `negate` - (Boolean) If true, the rule will fail if the pattern matches.

### rules.file_path_restriction

- `restricted_file_paths` - (List of String) The file paths that are restricted from being pushed to the commit graph.

### rules.max_file_size

- `max_file_size` - (Number) The maximum allowed size, in megabytes (MB), of a file.

### rules.max_file_path_length

- `max_file_path_length` - (Number) The maximum number of characters allowed in file paths.

### rules.file_extension_restriction

- `restricted_file_extensions` - (List of String) The file extensions that are restricted from being pushed to the commit graph.

### rules.repository_name

- `pattern` - (String) The pattern to match repository names against.

- `negate` - (Boolean) If true, the rule will fail if the pattern matches.

### rules.repository_visibility

- `internal` - (Boolean) Allow internal visibility for repositories.

- `private` - (Boolean) Allow private visibility for repositories.
