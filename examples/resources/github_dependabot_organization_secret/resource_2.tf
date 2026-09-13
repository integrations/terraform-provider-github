# Encrypted Secret Example

data "github_dependabot_organization_public_key" "example" {}

resource "github_dependabot_organization_secret" "example" {
  secret_name     = "EXAMPLE_SECRET_NAME"
  key_id          = data.github_dependabot_organization_public_key.example.key_id
  encrypted_value = var.encrypted_value
  visibility      = "all"
}
