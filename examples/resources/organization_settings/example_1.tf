resource "github_organization_settings" "test" {
  billing_email                                                = "test@example.com"
  company                                                      = "Test Company"
  blog                                                         = "https://example.com"
  email                                                        = "test@example.com"
  twitter_username                                             = "Test"
  location                                                     = "Test Location"
  name                                                         = "Test Name"
  description                                                  = "Test Description"
  has_organization_projects                                    = true
  has_repository_projects                                      = true
  default_repository_permission                                = "read"
  members_can_create_repositories                              = true
  members_can_create_public_repositories                       = true
  members_can_create_private_repositories                      = true
  members_can_create_internal_repositories                     = true
  members_can_create_pages                                     = true
  members_can_create_public_pages                              = true
  members_can_create_private_pages                             = true
  members_can_fork_private_repositories                        = true
  web_commit_signoff_required                                  = true
  advanced_security_enabled_for_new_repositories               = false
  dependabot_alerts_enabled_for_new_repositories               = false
  dependabot_security_updates_enabled_for_new_repositories     = false
  dependency_graph_enabled_for_new_repositories                = false
  secret_scanning_enabled_for_new_repositories                 = false
  secret_scanning_push_protection_enabled_for_new_repositories = false
}
