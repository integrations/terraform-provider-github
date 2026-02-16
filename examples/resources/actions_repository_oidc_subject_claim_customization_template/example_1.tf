resource "github_repository" "example" {
  name = "example-repository"
}

resource "github_actions_repository_oidc_subject_claim_customization_template" "example_template" {
  repository         = github_repository.example.name
  use_default        = false
  include_claim_keys = ["actor", "context", "repository_owner"]
}
