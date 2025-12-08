resource "github_actions_environment_secret" "example_secret" {
  environment       = "example_environment"
  secret_name       = "example_secret_name"
  plaintext_value   = "placeholder"

  lifecycle {
    ignore_changes = [plaintext_value]
  }
}

resource "github_actions_environment_secret" "example_secret" {
  environment       = "example_environment"
  secret_name       = "example_secret_name"
  encrypted_value   = base64sha256("placeholder")

  lifecycle {
    ignore_changes = [encrypted_value]
  }
}
