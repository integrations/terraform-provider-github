resource "github_organization_custom_role" "example" {
  name        = "example"
  description = "Example custom role that uses the read role as its base"
  base_role   = "read"
  permissions = [
    "add_assignee",
    "add_label",
    "bypass_branch_protection",
    "close_issue",
    "close_pull_request",
    "mark_as_duplicate",
    "create_tag",
    "delete_issue",
    "delete_tag",
    "manage_deploy_keys",
    "push_protected_branch",
    "read_code_scanning",
    "reopen_issue",
    "reopen_pull_request",
    "request_pr_review",
    "resolve_dependabot_alerts",
    "resolve_secret_scanning_alerts",
    "view_secret_scanning_alerts",
    "write_code_scanning"
  ]
}
