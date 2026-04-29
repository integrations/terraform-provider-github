data "github_repository" "example" {
  full_name = "my-org/repo"
}

resource "github_repository_environment" "example_plaintext" {
  repository  = data.github_repository.example.name
  environment = "example-environment"
}

resource "github_actions_environment_secret" "example_encrypted" {
  repository      = data.github_repository.example.name
  environment     = github_repository_environment.example.environment
  secret_name     = "test_secret_name"
  plaintext_value = "example-value"
}
