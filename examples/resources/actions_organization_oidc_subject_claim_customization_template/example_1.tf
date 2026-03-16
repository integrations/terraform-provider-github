resource "github_actions_organization_oidc_subject_claim_customization_template" "example_template" {
  include_claim_keys = ["actor", "context", "repository_owner"]
}
