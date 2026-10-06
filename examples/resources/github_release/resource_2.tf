# Live Example

resource "github_repository" "repo" {
  name      = "repo"
  auto_init = true
}

resource "github_release" "example" {
  repository = github_repository.repo.name
  tag_name   = "v1.0.0"
  draft      = false
  prerelease = false
}
