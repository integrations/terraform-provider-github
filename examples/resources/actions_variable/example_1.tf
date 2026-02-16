resource "github_actions_variable" "example_variable" {
  repository    = "example_repository"
  variable_name = "example_variable_name"
  value         = "example_variable_value"
}
