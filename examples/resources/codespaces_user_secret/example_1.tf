data "github_repository" "repo" {
  full_name = "my-org/repo"
}

resource "github_codespaces_user_secret" "example_secret" {
  secret_name             = "example_secret_name"
  plaintext_value         = var.some_secret_string
  selected_repository_ids = [data.github_repository.repo.repo_id]
}

resource "github_codespaces_user_secret" "example_secret" {
  secret_name             = "example_secret_name"
  encrypted_value         = var.some_encrypted_secret_string
  selected_repository_ids = [data.github_repository.repo.repo_id]
}
