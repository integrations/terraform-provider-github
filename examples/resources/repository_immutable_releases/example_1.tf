resource "github_repository" "example" {
  name        = "my-repo"
  description = "GitHub repo managed by Terraform"
  visibility  = "private"
}

resource "github_repository_immutable_releases" "example" {
  repository = github_repository.example.name
  enabled    = true
}
