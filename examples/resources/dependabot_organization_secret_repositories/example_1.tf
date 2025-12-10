data "github_repository" "repo" {
  full_name = "my-org/repo"
}

resource "github_dependabot_organization_secret" "example_secret" {
  secret_name     = "example_secret_name"
  visibility      = "private"
  plaintext_value = var.some_secret_string
}

resource "github_dependabot_organization_secret_repositories" "org_secret_repos" {
  secret_name = github_dependabot_organization_secret.example_secret.secret_name
  selected_repository_ids = [data.github_repository.repo.repo_id]
}
