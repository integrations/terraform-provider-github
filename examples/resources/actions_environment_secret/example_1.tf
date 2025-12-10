resource "github_actions_environment_secret" "example_secret" {
  environment       = "example_environment"
  secret_name       = "example_secret_name"
  plaintext_value   = var.some_secret_string
}

resource "github_actions_environment_secret" "example_secret" {
  environment       = "example_environment"
  secret_name       = "example_secret_name"
  encrypted_value   = var.some_encrypted_secret_string
}
