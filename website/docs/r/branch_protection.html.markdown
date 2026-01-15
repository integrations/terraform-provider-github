---
layout: "github"
page_title: "GitHub: github_branch_protection"
description: |-
  Protects a GitHub branch.
---

# github\_branch\_protection

Protects a GitHub branch.

This resource allows you to configure branch protection for repositories in your organization. When applied, the branch will be protected from forced pushes and deletion. Additional constraints, such as required status checks or restrictions on users, teams, and apps, can also be configured.

Note: for the `push_allowances` a given user or team must have specific write access to the repository. If specific write access not provided, github will reject the given actor, which will be the cause of terraform drift.

## Example Usage

```hcl
# Protect the main branch of the foo repository. Additionally, require that
# the "ci/travis" context to be passing and only allow the engineers team merge
# to the branch.

resource "github_branch_protection" "example" {
  repository_id = github_repository.example.node_id
  # also accepts repository name
  # repository_id  = github_repository.example.name

  pattern          = "main"
  enforce_admins   = true
  allows_deletions = true

  required_status_checks {
    strict   = false
    contexts = ["ci/travis"]
  }

  required_pull_request_reviews {
    dismiss_stale_reviews  = true
    restrict_dismissals    = true
    dismissal_restrictions = [
      data.github_user.example.node_id,
      github_team.example.node_id,
      "/exampleuser",
      "exampleorganization/exampleteam",
    ]
  }

  restrict_pushes {
    push_allowances = [
      data.github_user.example.node_id,
      "/exampleuser",
      "exampleorganization/exampleteam",
      # you can have more than one type of restriction (teams + users). If you use
      # more than one type, you must use node_ids of each user and each team.
      # github_team.example.node_id
      # github_user.example-2.node_id
    ]
  }

  force_push_bypassers = [
    data.github_user.example.node_id,
    "/exampleuser",
    "exampleorganization/exampleteam",
    # you can have more than one type of restriction (teams + users)
    # github_team.example.node_id
    # github_team.example-2.node_id
  ]

}

resource "github_repository" "example" {
  name = "test"
}

data "github_user" "example" {
  username = "example"
}

resource "github_team" "example" {
  name = "Example Name"
}

resource "github_team_repository" "example" {
  team_id    = github_team.example.id
  repository = github_repository.example.name
  permission = "pull"
}
```

## Argument Reference

The following arguments are supported:

* `repository_id` - (Required) The name or node ID of the repository associated with this branch protection rule.
* `pattern` - (Required) Identifies the protection rule pattern.
* `enforce_admins` - (Optional) Boolean, setting this to `true` enforces status checks for repository administrators.
* `require_signed_commits` - (Optional) Boolean, setting this to `true` requires all commits to be signed with GPG.
* `required_linear_history` - (Optional) Boolean, setting this to `true` enforces a linear commit Git history, which prevents anyone from pushing merge commits to a branch
* `require_conversation_resolution` - (Optional) Boolean, setting this to `true` requires all conversations on code must be resolved before a pull request can be merged.
* `required_status_checks` - (Optional) Enforce restrictions for required status checks. See [Required Status Checks](#required-status-checks) below for details.
* `required_pull_request_reviews` - (Optional) Enforce restrictions for pull request reviews. See [Required Pull Request Reviews](#required-pull-request-reviews) below for details.
* `restrict_pushes` - (Optional) Restrict pushes to matching branches. See [Restrict Pushes](#restrict-pushes) below for details.
* `force_push_bypassers` - (Optional) The list of actor Names/IDs that are allowed to bypass force push restrictions. Actor names must either begin with a "/" for users or the organization name followed by a "/" for teams. If the list is not empty, `allows_force_pushes` should be set to `false`.
* `allows_deletions` - (Optional) Boolean, setting this to `true` to allow the branch to be deleted.
* `allows_force_pushes` - (Optional) Boolean, setting this to `true` to allow force pushes on the branch to everyone. Set it to `false` if you specify `force_push_bypassers`.
* `lock_branch` - (Optional) Boolean, Setting this to `true` will make the branch read-only and preventing any pushes to it. Defaults to `false`

### Required Status Checks

`required_status_checks` supports the following arguments:

* `strict`: (Optional) Require branches to be up to date before merging. Defaults to `false`.
* `contexts`: (Optional) The list of status checks to require in order to merge into this branch. No status checks are required by default.

~> Note: This attribute can contain multiple string patterns.
If specified, usual value is the [job name](https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions#jobsjob_idname). Otherwise, the [job id](https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions#jobsjob_idname) is defaulted to.
For workflows that use matrixes, append the matrix name to the value using the following pattern `(<matrix_value>[, <matrix_value>])`. Matrixes should be specified based on the order of matrix properties in the workflow file. See [GitHub Documentation]("https://docs.github.com/en/actions/using-jobs/using-a-matrix-for-your-jobs#using-a-matrix-strategy") for more information.
For workflows that use reusable workflows, the pattern is `<initial_workflow.jobs.job.[name/id]> / <reused-workflow.jobs.job.[name/id]>`. This can extend multiple levels.

### Required Pull Request Reviews

`required_pull_request_reviews` supports the following arguments:

* `dismiss_stale_reviews`: (Optional) Dismiss approved reviews automatically when a new commit is pushed. Defaults to `false`.
* `restrict_dismissals`: (Optional) Restrict pull request review dismissals.
* `dismissal_restrictions`: (Optional) The list of actor Names/IDs with dismissal access. If not empty, `restrict_dismissals` is ignored. Actor names must either begin with a "/" for users or the organization name followed by a "/" for teams.
* `pull_request_bypassers`: (Optional) The list of actor Names/IDs that are allowed to bypass pull request requirements. Actor names must either begin with a "/" for users or the organization name followed by a "/" for teams.
* `require_code_owner_reviews`: (Optional) Require an approved review in pull requests including files with a designated code owner. Defaults to `false`.
* `required_approving_review_count`: (Optional) Require x number of approvals to satisfy branch protection requirements. If this is specified it must be a number between 0-6. This requirement matches GitHub's API, see the upstream [documentation](https://developer.github.com/v3/repos/branches/#parameters-1) for more information.
  (https://developer.github.com/v3/repos/branches/#parameters-1) for more information.
* `require_last_push_approval`: (Optional) Require that The most recent push must be approved by someone other than the last pusher.  Defaults to `false`

### Restrict Pushes

`restrict_pushes` supports the following arguments:

* `blocks_creations` - (Optional) Boolean, setting this to `false` allows people, teams, or apps to create new branches matching this rule. Defaults to `true`.
* `push_allowances` - (Optional) A list of actor Names/IDs that may push to the branch. Actor names must either begin with a "/" for users or the organization name followed by a "/" for teams. Organization administrators, repository administrators, and users with the Maintain role on the repository can always push when all other requirements have passed.

## Import

GitHub Branch Protection can be imported using an ID made up of `repository:pattern`, e.g.

```
$ terraform import github_branch_protection.terraform terraform:main
```