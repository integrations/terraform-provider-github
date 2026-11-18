data "github_enterprise_cost_centers" "active" {
  enterprise_slug = "example-enterprise"
  state           = "active"
}
