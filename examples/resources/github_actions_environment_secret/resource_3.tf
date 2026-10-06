# Ignore Drift Secret Example

resource "github_actions_environment_secret" "example" {
  repository  = "example-repo"
  environment = "example-environment"
  secret_name = "EXAMPLE_SECRET_NAME"
  value       = "example-value"

  lifecycle {
    ignore_changes = [updated_at]
  }
}
