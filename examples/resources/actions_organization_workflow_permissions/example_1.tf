# Basic workflow permissions configuration
resource "github_actions_organization_workflow_permissions" "example" {
  organization_slug = "my-organization"

  default_workflow_permissions     = "read"
  can_approve_pull_request_reviews = false
}

# Allow write permissions and PR approvals
resource "github_actions_organization_workflow_permissions" "permissive" {
  organization_slug = "my-organization"

  default_workflow_permissions     = "write"
  can_approve_pull_request_reviews = true
}
