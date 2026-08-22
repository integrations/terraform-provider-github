data "github_repository" "repo" {
  full_name = "my-org/repo"
}

resource "github_codespaces_organization_secret_repositories" "org_secret_repos" {
  secret_name             = "existing_secret_name"
  selected_repository_ids = [data.github_repository.repo.repo_id]
}
