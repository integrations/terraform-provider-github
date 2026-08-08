resource "github_actions_organization_fork_pr_contributor_approval" "test" {
  approval_policy = "all_external_contributors"
}
