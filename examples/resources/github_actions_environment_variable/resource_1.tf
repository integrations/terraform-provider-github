resource "github_actions_environment_variable" "example" {
  repository    = "example-repo"
  environment   = "example-environment"
  variable_name = "EXAMPLE_VARIABLE_NAME"
  value         = "example-value"
}
