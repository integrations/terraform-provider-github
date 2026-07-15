# Organization-level configuration, set as default for new private/internal
# repos and attached to all repositories without an existing configuration
resource "github_code_security_configuration" "org_baseline" {
  name        = "org-security-baseline"
  description = "Baseline security configuration for all repositories"

  advanced_security                     = "enabled"
  dependency_graph                      = "enabled"
  dependabot_alerts                     = "enabled"
  dependabot_security_updates           = "enabled"
  code_scanning_default_setup           = "enabled"
  secret_scanning                       = "enabled"
  secret_scanning_push_protection       = "enabled"
  secret_scanning_validity_checks       = "enabled"
  secret_scanning_non_provider_patterns = "disabled"
  private_vulnerability_reporting       = "enabled"
  enforcement                           = "enforced"

  default_for_new_repos = "private_and_internal"
  attach_scope          = "all_without_configurations"
}

# Enterprise-level configuration
resource "github_code_security_configuration" "enterprise_baseline" {
  enterprise_slug = "my-enterprise"
  name            = "enterprise-security-baseline"
  description     = "Enterprise-wide security baseline"

  dependabot_alerts               = "enabled"
  secret_scanning                 = "enabled"
  secret_scanning_push_protection = "enabled"

  default_for_new_repos = "all"
}
