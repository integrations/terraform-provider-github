# Secret With Selected Repositories Example

resource "github_actions_organization_secret" "example" {
  secret_name = "EXAMPLE_SECRET_NAME"
  value       = "example-value"
  visibility  = "selected"
}

data "github_repository" "example" {
  name = "example-repo"
}

resource "github_actions_organization_secret_repositories" "example" {
  secret_name             = github_actions_organization_secret.example.secret_name
  selected_repository_ids = [data.github_repository.example.repo_id]
}
