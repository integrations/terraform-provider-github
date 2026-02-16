resource "github_dependabot_organization_secret" "example_allow_drift" {
  secret_name     = "example_secret_name"
  visibility      = "all"
  plaintext_value = "placeholder"

  lifecycle {
    ignore_changes = [remote_updated_at]
  }
}
