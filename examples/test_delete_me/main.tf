terraform {
  required_providers {
    github = {
      source  = "integrations/github"
    }
  }
}

provider "github" {}

resource "github_repository" "test" {
  name = "test-repo-0"
}

# output "repo_full_name" {
#   value = github_repository.test.full_name
# }

# output "public_key" {
#   value = data.github_codespaces_public_key.example.key
# }

# resource "github_actions_secret" "example_secret" {
#   repository       = github_repository.test.name
#   secret_name      = "example_secret_2"
#   plaintext_value  = "test"
# }

# output "secret_repo_name" {
#   value = github_actions_secret.example_secret.repository
# }

# resource "github_codespaces_secret" "example_secret" {
#   repository       = github_repository.test.name
#   secret_name      = "example_secret_2"
#   plaintext_value  = "test"
# }

# data "github_codespaces_secrets" "secrets" {
#   full_name = "kens-test-org/${github_repository.test.name}"
# }

# output "repo_secrets" {
#   value = data.github_codespaces_secrets.secrets.secrets
# }

# data "github_codespaces_user_public_key" "example" {
# }

# output "user_public_key" {
#   value = data.github_codespaces_user_public_key.example.key
# }

# resource "github_codespaces_user_secret" "example_secret" {
#   secret_name             = "example_user_secret_2"
#   visibility              = "selected"
#   plaintext_value         = "test"
#   selected_repository_ids = [github_repository.test.repo_id]
# }

# data "github_codespaces_user_secrets" "secrets" {
# }

# output "user_secrets" {
#   value = data.github_codespaces_user_secrets.secrets.secrets
# }

data "github_dependabot_organization_public_key" "example" {
}

output "organization_public_key" {
  value = data.github_dependabot_organization_public_key.example.key
}

resource "github_dependabot_organization_secret" "example_secret" {
  secret_name             = "examplp_organization_secret"
  visibility              = "selected"
  plaintext_value         = "test"
  selected_repository_ids = [github_repository.test.repo_id]
}

# data "github_codespaces_organization_secrets" "secrets" {
# }

# output "organization_secrets" {
#   value = data.github_codespaces_organization_secrets.secrets.secrets
# }