# Create some repositories.
resource "github_repository" "some_repo" {
  name = "some-repo"
}

resource "github_repository" "another_repo" {
  name = "another-repo"
}

resource "github_app_installation_repositories" "some_app_repos" {
  # The installation id of the app (in the organization).
  installation_id       = "1234567"
  selected_repositories = [github_repository.some_repo.name, github_repository.another_repo.name]
}
