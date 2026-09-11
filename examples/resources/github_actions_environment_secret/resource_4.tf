# Write-Only Secret Example

variable "example_secret_value" {
  type      = string
  sensitive = true
  ephemeral = true
}

resource "github_actions_environment_secret" "example" {
  repository       = "example-repo"
  environment      = "example-environment"
  secret_name      = "EXAMPLE_SECRET_NAME"
  value_wo         = var.example_secret_value
  value_wo_version = 1
}
