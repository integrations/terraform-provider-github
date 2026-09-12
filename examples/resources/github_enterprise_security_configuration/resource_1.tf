resource "github_enterprise_security_configuration" "default" {
  enterprise_slug                 = "my-enterprise"
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
