resource "github_codespaces_organization_secret" "example_secret" {
  secret_name     = "example_secret_name"
  visibility      = "private"
  plaintext_value = var.some_secret_string
}

resource "github_codespaces_organization_secret" "example_secret" {
  secret_name     = "example_secret_name"
  visibility      = "private"
  encrypted_value = var.some_encrypted_secret_string
}
