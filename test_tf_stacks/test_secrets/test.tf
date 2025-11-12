terraform {
  required_providers {
    github = {
      source = "integrations/github"
    }
  }
}

provider "github" {
  token = "fake_token_for_validation"
}

# Test both resources with different configurations
resource "github_actions_secret" "test" {
  repository       = "test_repo"
  secret_name      = "test_secret"
  plaintext_value  = "test_value"
  destroy_on_drift = true
}

resource "github_actions_organization_secret" "test" {
  secret_name      = "org_secret"
  encrypted_value  = "dGVzdA=="
  visibility       = "private"
  destroy_on_drift = false
}
 
