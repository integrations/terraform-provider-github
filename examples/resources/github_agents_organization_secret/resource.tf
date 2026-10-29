variable "agent_secret" {
  type        = string
  description = "Value provided by a secret manager or secure input."
  sensitive   = true
}

resource "github_agents_organization_secret" "example" {
  secret_name = "SERVICE_TOKEN"
  visibility  = "selected"
  value       = var.agent_secret
}
