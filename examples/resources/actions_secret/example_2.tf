resource "github_actions_secret" "example_allow_drift" {
  repository      = "example_repository"
  secret_name     = "example_secret_name"
  plaintext_value = "placeholder"

  lifecycle {
    ignore_changes = [remote_updated_at]
  }
}
