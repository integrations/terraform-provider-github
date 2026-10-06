resource "github_organization_repository_role" "example" {
  name      = "example"
  base_role = "read"

  permissions = [
    "add_assignee",
    "add_label"
  ]
}
