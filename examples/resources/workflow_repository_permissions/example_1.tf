resource "github_repository" "example" {
  name = "my-repository"
}

resource "github_workflow_repository_permissions" "test" {
  default_workflow_permissions     = "read"
  can_approve_pull_request_reviews = true
  repository                       = github_repository.example.name
}
