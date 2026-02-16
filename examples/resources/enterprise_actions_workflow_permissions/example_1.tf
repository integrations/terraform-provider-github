# Basic workflow permissions configuration
resource "github_enterprise_actions_workflow_permissions" "example" {
  enterprise_slug = "my-enterprise"

  default_workflow_permissions     = "read"
  can_approve_pull_request_reviews = false
}

# Allow write permissions and PR approvals
resource "github_enterprise_actions_workflow_permissions" "permissive" {
  enterprise_slug = "my-enterprise"

  default_workflow_permissions     = "write"
  can_approve_pull_request_reviews = true
}
