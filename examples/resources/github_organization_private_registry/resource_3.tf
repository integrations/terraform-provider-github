# AWS CodeArtifact registry authenticated through OIDC, no stored secret.
resource "github_organization_private_registry" "example" {
  registry_type         = "maven_repository"
  url                   = "https://example-111122223333.d.codeartifact.eu-west-3.amazonaws.com/maven/releases/"
  auth_type             = "oidc_aws"
  oidc_aws_region       = "eu-west-3"
  oidc_aws_account_id   = "111122223333"
  oidc_aws_role_name    = "dependabot-codeartifact"
  oidc_aws_domain       = "example"
  oidc_aws_domain_owner = "111122223333"
  visibility            = "private"
}
