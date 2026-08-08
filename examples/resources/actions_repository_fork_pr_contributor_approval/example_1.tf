resource "github_repository" "example" {
  name       = "my-repository"
  visibility = "public"
}

resource "github_actions_repository_fork_pr_contributor_approval" "test" {
  approval_policy = "all_external_contributors"
  repository      = github_repository.example.name
}
