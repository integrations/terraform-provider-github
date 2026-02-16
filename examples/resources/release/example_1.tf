resource "github_repository" "repo" {
  name        = "repo"
  description = "GitHub repo managed by Terraform"

  private = false
}

resource "github_release" "example" {
  repository = github_repository.repo.name
  tag_name   = "v1.0.0"
}
