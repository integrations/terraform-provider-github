resource "github_actions_organization_secret" "example_plaintext" {
  secret_name     = "example_secret_name"
  visibility      = "all"
  plaintext_value = var.some_secret_string
}

resource "github_actions_organization_secret" "example_encrypted" {
  secret_name     = "example_secret_name"
  visibility      = "all"
  encrypted_value = var.some_encrypted_secret_string
}
