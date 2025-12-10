data "github_repository" "repo" {
  full_name = "my-org/repo"
}

resource "github_repository_environment" "repo_environment" {
  repository       = data.github_repository.repo.name
  environment      = "example_environment"
}

resource "github_actions_environment_variable" "example_variable" {
  repository       = data.github_repository.repo.name
  environment      = github_repository_environment.repo_environment.environment
  variable_name    = "example_variable_name"
  value            = "example_variable_value"
}
