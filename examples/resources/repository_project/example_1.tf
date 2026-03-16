resource "github_repository" "example" {
  name         = "example"
  description  = "My awesome codebase"
  has_projects = true
}

resource "github_repository_project" "project" {
  name       = "A Repository Project"
  repository = "${github_repository.example.name}"
  body       = "This is a repository project."
}
