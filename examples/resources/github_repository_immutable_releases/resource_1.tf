resource "github_repository" "example" {
  name       = "example-repo"
  visibility = "public"
}

resource "github_repository_immutable_releases" "example" {
  repository = github_repository.example.name
}
