resource "github_organization_security_configuration" "default" {
  name                            = "default-config"
  description                     = "Default security configuration"
  advanced_security               = "enabled"
  dependency_graph                = "enabled"
  dependabot_alerts               = "enabled"
  dependabot_security_updates     = "enabled"
  code_scanning_default_setup     = "enabled"
  secret_scanning                 = "enabled"
  secret_scanning_push_protection = "enabled"
  private_vulnerability_reporting = "enabled"
  enforcement                     = "enforced"
}

# Delegated bypass lets specific teams approve secret-scanning push-protection
# bypass requests. Reference the team's numeric ID via the github_team resource
# rather than hardcoding it.
resource "github_team" "security_reviewers" {
  name = "security-reviewers"
}

resource "github_organization_security_configuration" "with_delegated_bypass" {
  name                             = "delegated-bypass-config"
  description                      = "Configuration with delegated bypass reviewers"
  advanced_security                = "enabled"
  secret_scanning                  = "enabled"
  secret_scanning_push_protection  = "enabled"
  secret_scanning_delegated_bypass = "enabled"

  secret_scanning_delegated_bypass_options {
    reviewers {
      reviewer_id   = github_team.security_reviewers.id
      reviewer_type = "TEAM"
    }
  }
}
