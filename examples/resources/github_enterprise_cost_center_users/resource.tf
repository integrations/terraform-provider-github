resource "github_enterprise_cost_center_users" "example" {
  enterprise_slug = "example-enterprise"
  cost_center_id  = github_enterprise_cost_center.example.id
  usernames       = ["alice", "bob"]
}
