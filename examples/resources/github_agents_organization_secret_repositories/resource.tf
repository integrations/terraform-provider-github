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

resource "github_agents_organization_secret_repositories" "example" {
  secret_name             = github_agents_organization_secret.example.secret_name
  selected_repository_ids = [github_repository.example.repo_id]
}
