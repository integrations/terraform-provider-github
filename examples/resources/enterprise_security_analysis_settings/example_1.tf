# Basic security settings - enable secret scanning only
resource "github_enterprise_security_analysis_settings" "basic" {
  enterprise_slug = "my-enterprise"
  
  secret_scanning_enabled_for_new_repositories = true
}

# Full security configuration with all features enabled
resource "github_enterprise_security_analysis_settings" "comprehensive" {
  enterprise_slug = "my-enterprise"
  
  advanced_security_enabled_for_new_repositories             = true
  secret_scanning_enabled_for_new_repositories               = true
  secret_scanning_push_protection_enabled_for_new_repositories = true
  secret_scanning_validity_checks_enabled                   = true
  secret_scanning_push_protection_custom_link               = "https://octokit.com/security-guidelines"
}
