variable "agent_ciphertext" {
  type        = string
  description = "Base64-encoded sealed-box ciphertext encrypted outside Terraform using the organization's agent public key."
  sensitive   = true
}

variable "agent_key_id" {
  type        = string
  description = "ID of the agent public key used to produce agent_ciphertext."
}

resource "github_agents_organization_secret" "encrypted" {
  secret_name     = "ENCRYPTED_SERVICE_TOKEN"
  visibility      = "private"
  key_id          = var.agent_key_id
  value_encrypted = var.agent_ciphertext
}
