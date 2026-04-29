resource "github_repository" "repo" {
  name        = "my-repo"
  description = "GitHub repo managed by Terraform"

  private = false

  vulnerability_alerts = true
}


resource "github_repository_dependabot_security_updates" "example" {
  repository = github_repository.test.name
  enabled    = true
}
