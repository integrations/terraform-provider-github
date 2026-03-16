resource "github_dependabot_secret" "example_plaintext" {
  repository       = "example_repository"
  secret_name      = "example_secret_name"
  plaintext_value  = var.some_secret_string
}

resource "github_dependabot_secret" "example_encrypted" {
  repository       = "example_repository"
  secret_name      = "example_secret_name"
  encrypted_value  = var.some_encrypted_secret_string
}
