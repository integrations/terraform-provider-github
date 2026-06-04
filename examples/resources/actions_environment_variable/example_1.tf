resource "github_actions_environment_variable" "example" {
  repository    = "example-repo"
  environment   = "example-environment"
  variable_name = "example_variable_name"
  value         = "example-value"
}
