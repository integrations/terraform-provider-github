resource "github_repository" "example" {
  name = "example-repository"
}

resource "github_actions_repository_oidc_subject_claim_customization_template" "example_template" {
  repository         = github_repository.example.name
  use_default        = false
  include_claim_keys = ["actor", "context", "repository_owner"]

  # Opt this repository into immutable subject claims (owner and repository IDs)
  # ahead of an org-wide rollout.
  use_immutable_subject = true
  sub_claim_prefix      = "custom-prefix"
}
