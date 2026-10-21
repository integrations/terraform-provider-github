resource "github_enterprise_cost_center_repositories" "example" {
  enterprise_slug  = "example-enterprise"
  cost_center_id   = github_enterprise_cost_center.example.id
  repository_names = ["octo-org/my-app", "acme-corp/backend-service"]
}
