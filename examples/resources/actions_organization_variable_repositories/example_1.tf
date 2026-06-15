resource "github_actions_organization_variable" "example" {
  variable_name   = "myvariable"
  plaintext_value = "foo"
  visibility      = "selected"
}

resource "github_repository" "example" {
  name       = "myrepo"
  visibility = "public"
}

resource "github_actions_organization_variable_repositories" "example" {
  variable_name           = github_actions_organization_variable.example.name
  selected_repository_ids = [github_repository.example.repo_id]
}
