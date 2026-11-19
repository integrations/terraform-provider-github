variable "agent_secret" {
  type        = string
  description = "Value provided by a secret manager or secure input."
  sensitive   = true
}

resource "github_repository" "example" {
  name = "agent-secret-example"
}

resource "github_agents_secret" "example" {
  repository  = github_repository.example.name
  secret_name = "SERVICE_TOKEN"
  value       = var.agent_secret
}
