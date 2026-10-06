# Ignore Drift Secret Example

resource "github_actions_secret" "example" {
  repository  = "example-repo"
  secret_name = "EXAMPLE_SECRET_NAME"
  value       = "example-value"

  lifecycle {
    ignore_changes = [updated_at]
  }
}
