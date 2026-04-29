data "github_repository" "repo" {
  full_name = "my-org/repo"
}

resource "github_actions_organization_variable" "example_variable" {
  variable_name           = "example_variable_name"
  visibility              = "selected"
  value                   = "example_variable_value"
  selected_repository_ids = [data.github_repository.repo.repo_id]
}
