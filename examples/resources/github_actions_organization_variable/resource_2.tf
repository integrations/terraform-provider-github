# Variable With Selected Repositories Example

resource "github_actions_organization_variable" "example" {
  variable_name = "EXAMPLE_VARIABLE_NAME"
  value         = "example-value"
  visibility    = "selected"
}

data "github_repository" "example" {
  name = "example-repo"
}

resource "github_actions_organization_variable_repositories" "example" {
  variable_name           = github_actions_organization_variable.example.variable_name
  selected_repository_ids = [data.github_repository.example.repo_id]
}
