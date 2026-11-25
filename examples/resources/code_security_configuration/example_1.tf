# Organization-level configuration, set as the default for new private and
# internal repositories
resource "github_code_security_configuration" "org_baseline" {
  name        = "org-security-baseline"
  description = "Baseline security configuration for all repositories"

  advanced_security                     = "enabled"
  dependency_graph                      = "enabled"
  dependency_graph_autosubmit_action    = "enabled"
  dependabot_alerts                     = "enabled"
  dependabot_security_updates           = "enabled"
  code_scanning_default_setup           = "enabled"
  secret_scanning                       = "enabled"
  secret_scanning_push_protection       = "enabled"
  secret_scanning_validity_checks       = "enabled"
  secret_scanning_non_provider_patterns = "disabled"
  private_vulnerability_reporting       = "enabled"
  enforcement                           = "enforced"

  dependency_graph_autosubmit_action_options {
    labeled_runners = false
  }

  code_scanning_default_setup_options {
    runner_type  = "labeled"
    runner_label = "codeql-runners"
  }

  code_scanning_options {
    allow_advanced = true
  }

  default_for_new_repos = "private_and_internal"
}

# Organization-level configuration allowing a team to bypass push protection
resource "github_team" "security" {
  name = "security"
}

resource "github_code_security_configuration" "with_bypass" {
  name = "push-protection-with-bypass"

  advanced_security                = "enabled"
  secret_scanning                  = "enabled"
  secret_scanning_push_protection  = "enabled"
  secret_scanning_delegated_bypass = "enabled"

  secret_scanning_delegated_bypass_options {
    reviewers {
      reviewer_id   = github_team.security.id
      reviewer_type = "TEAM"
    }
  }
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
