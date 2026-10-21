# Write-Only Secret Example

variable "example_secret_value" {
  type      = string
  sensitive = true
  ephemeral = true
}

resource "github_dependabot_organization_secret" "example" {
  secret_name      = "EXAMPLE_SECRET_NAME"
  value_wo         = var.example_secret_value
  value_wo_version = 1
  visibility       = "all"
}
