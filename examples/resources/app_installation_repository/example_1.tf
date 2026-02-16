# Create a repository.
resource "github_repository" "some_repo" {
  name = "some-repo"
}

resource "github_app_installation_repository" "some_app_repo" {
  # The installation id of the app (in the organization).
  installation_id = "1234567"
  repository      = github_repository.some_repo.name
}
