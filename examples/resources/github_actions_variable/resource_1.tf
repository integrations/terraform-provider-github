resource "github_actions_variable" "example" {
  repository    = "example-repo"
  variable_name = "EXAMPLE_VARIABLE_NAME"
  value         = "example-value"
}
