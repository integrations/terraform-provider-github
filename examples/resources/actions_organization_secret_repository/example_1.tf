data "github_repository" "repo" {
  full_name = "my-org/repo"
}

resource "github_actions_organization_secret_repository" "org_secret_repos" {
  secret_name = "EXAMPLE_SECRET_NAME"
  repository_id = github_repository.repo.repo_id
}
