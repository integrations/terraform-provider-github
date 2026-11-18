# Write-Only Secret Example

variable "example_secret_value" {
  type      = string
  sensitive = true
  ephemeral = true
}

resource "github_dependabot_secret" "example" {
  repository       = "example-repo"
  secret_name      = "EXAMPLE_SECRET_NAME"
  value_wo         = var.example_secret_value
  value_wo_version = 1
}
