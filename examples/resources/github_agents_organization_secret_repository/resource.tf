variable "agent_secret" {
  type        = string
  description = "Value provided by a secret manager or secure input."
  sensitive   = true
}

resource "github_repository" "example" {
  name = "agent-secret-example"
}

resource "github_agents_organization_secret" "example" {
  secret_name = "SERVICE_TOKEN"
  visibility  = "selected"
  value       = var.agent_secret
}

resource "github_agents_organization_secret_repository" "example" {
  secret_name   = github_agents_organization_secret.example.secret_name
  repository_id = github_repository.example.repo_id
}
