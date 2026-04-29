resource "github_repository" "example" {
  name        = "example"
  description = "My awesome codebase"
  auto_init   = true
}

resource "github_branch" "development" {
  repository = github_repository.example.name
  branch     = "development"
}

resource "github_branch_default" "default" {
  repository = github_repository.example.name
  branch     = github_branch.development.branch
}
