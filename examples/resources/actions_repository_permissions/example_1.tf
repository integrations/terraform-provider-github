resource "github_repository" "example" {
  name = "my-repository"
}

resource "github_actions_repository_permissions" "test" {
  allowed_actions = "selected"
  allowed_actions_config {
    github_owned_allowed = true
    patterns_allowed     = ["actions/cache@*", "actions/checkout@*"]
    verified_allowed     = true
  }
  repository = github_repository.example.name
}
