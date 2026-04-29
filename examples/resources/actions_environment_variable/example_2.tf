data "github_repository" "example" {
  full_name = "my-org/repo"
}

resource "github_repository_environment" "example" {
  repository  = data.github_repository.example.name
  environment = "example_environment"
}

resource "github_actions_environment_variable" "example" {
  repository    = data.github_repository.example.name
  environment   = github_repository_environment.example.environment
  variable_name = "example_variable_name"
  value         = "example-value"
}
