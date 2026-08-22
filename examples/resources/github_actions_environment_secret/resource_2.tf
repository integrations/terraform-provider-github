# Encrypted Secret Example

data "github_actions_environment_public_key" "example" {
  repository  = "example-repo"
  environment = "example-environment"
}

resource "github_actions_environment_secret" "example" {
  repository      = "example-repo"
  environment     = "example-environment"
  secret_name     = "EXAMPLE_SECRET_NAME"
  key_id          = data.github_actions_environment_public_key.example.key_id
  encrypted_value = var.encrypted_value
}
