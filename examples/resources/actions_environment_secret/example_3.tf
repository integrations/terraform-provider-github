resource "github_actions_environment_secret" "example_allow_drift" {
  repository      = "example-repo"
  environment     = "example-environment"
  secret_name     = "example_secret_name"
  plaintext_value = "placeholder"

  lifecycle {
    ignore_changes = [updated_at]
  }
}
