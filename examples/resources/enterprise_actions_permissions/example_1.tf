data "github_organization" "example-org" {
  name = "my-org"
}

resource "github_enterprise_actions_permissions" "test" {
  enterprise_slug       = "my-enterprise"
  allowed_actions       = "selected"
  enabled_organizations = "selected"
  allowed_actions_config {
    github_owned_allowed = true
    patterns_allowed     = ["actions/cache@*", "actions/checkout@*"]
    verified_allowed     = true
  }
  enabled_organizations_config {
    organization_ids = [data.github_organization.example-org.id]
  }
}
