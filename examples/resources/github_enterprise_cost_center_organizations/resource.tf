resource "github_enterprise_cost_center_organizations" "example" {
  enterprise_slug     = "example-enterprise"
  cost_center_id      = github_enterprise_cost_center.example.id
  organization_logins = ["octo-org", "acme-corp"]
}
