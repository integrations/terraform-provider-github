data "github_repository" "repo" {
  full_name = "my-org/repo"
}

resource "github_repository_environment" "repo_environment" {
  repository       = data.github_repository.repo.name
  environment      = "example_environment"
}

resource "github_actions_environment_secret" "test_secret" {
  repository       = data.github_repository.repo.name
  environment      = github_repository_environment.repo_environment.environment
  secret_name      = "test_secret_name"
  plaintext_value  = "%s"
}
