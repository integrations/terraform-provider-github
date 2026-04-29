resource "github_repository" "example" {
  name        = "example"
  description = "My awesome codebase"
  auto_init   = true
}

resource "github_branch_default" "default" {
  repository = github_repository.example.name
  branch     = "development"
  rename     = true
}
