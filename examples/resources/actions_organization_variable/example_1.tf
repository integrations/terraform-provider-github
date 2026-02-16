resource "github_actions_organization_variable" "example_variable" {
  variable_name = "example_variable_name"
  visibility    = "private"
  value         = "example_variable_value"
}
