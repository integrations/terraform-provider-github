data "github_codespaces_public_key" "example_public_key" {
  repository = "example_repository"
}

resource "github_codespaces_secret" "example_secret" {
  repository      = "example_repository"
  secret_name     = "example_secret_name"
  plaintext_value = var.some_secret_string
}

resource "github_codespaces_secret" "example_secret" {
  repository      = "example_repository"
  secret_name     = "example_secret_name"
  encrypted_value = var.some_encrypted_secret_string
}
