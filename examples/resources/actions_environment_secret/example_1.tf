resource "github_actions_environment_secret" "example_plaintext" {
  repository      = "example-repo"
  environment     = "example-environment"
  secret_name     = "example_secret_name"
  plaintext_value = "example-value
}

resource "github_actions_environment_secret" "example_encrypted" {
  repository      = "example-repo"
  environment     = "example-environment"
  secret_name     = "example_secret_name"
  key_id          = var.key_id
  encrypted_value = var.encrypted_secret_string
}
