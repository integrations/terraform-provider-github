resource "github_repository" "example" {
  name = "example"
}

resource "github_repository_environment" "example" {
  repository  = github_repository.example.name
  environment = "production"
}

resource "github_repository_environment_deployment_protection_rule" "example" {
  repository     = github_repository.example.name
  environment    = github_repository_environment.example.environment
  integration_id = 123456 # The ID of the GitHub App
}
