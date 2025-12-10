resource "github_actions_environment_variable" "example_variable" {
  environment       = "example_environment"
  repository        = "example_repository"
  value             = "example_variable_value"
  variable_name     = "example_variable_name"
}
