# Encrypted Secret Example

data "github_actions_public_key" "example" {
  repository = "example-repo"
}

resource "github_actions_secret" "example" {
  repository      = "example-repo"
  secret_name     = "EXAMPLE_SECRET_NAME"
  key_id          = data.github_actions_public_key.example.key_id
  encrypted_value = var.encrypted_value
}
