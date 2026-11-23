resource "github_repository" "example" {
  name = "my-repository"
}

resource "github_actions_organization_permissions" "example" {
  enabled_repositories = "selected"

  enabled_repositories_config {
    repository_ids = [github_repository.example.repo_id]
  }

  allowed_actions = "selected"

  allowed_actions_config {
    github_owned_allowed = true
    patterns_allowed     = ["actions/cache@*", "actions/checkout@*"]
    verified_allowed     = true
  }

  sha_pinning_required = true
}
