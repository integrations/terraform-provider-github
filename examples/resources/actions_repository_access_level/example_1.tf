resource "github_repository" "example" {
  name       = "my-repository"
  visibility = "private"
}

resource "github_actions_repository_access_level" "test" {
  access_level = "user"
  repository   = github_repository.example.name
}
