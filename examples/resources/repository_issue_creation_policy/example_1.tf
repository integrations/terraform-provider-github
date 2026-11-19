resource "github_repository" "example" {
  name       = "example-repo"
  visibility = "private"
}

resource "github_repository_issue_creation_policy" "example" {
  repository = github_repository.example.name
  policy     = "collaborators_only"
}
