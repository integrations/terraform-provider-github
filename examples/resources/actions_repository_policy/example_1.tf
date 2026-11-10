resource "github_repository" "example" {
  name       = "my-repository"
  visibility = "public"
}

resource "github_actions_repository_policy" "example" {
  repository  = github_repository.example.name
  name        = "allow-pull-request-target"
  enforcement = "active"

  workflow_path {
    include = [".github/workflows/triage.yml"]
  }

  allowed_events = ["pull_request", "pull_request_target"]
}
