# Un-encrypted Variable Example

resource "github_actions_organization_variable" "example" {
  variable_name = "EXAMPLE_VARIABLE_NAME"
  value         = "example-value"
  visibility    = "all"
}
