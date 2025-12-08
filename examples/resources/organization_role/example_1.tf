resource "github_organization_role" "example" {
  name      = "example"
  base_role = "read"

  permissions = [
    "read_organization_custom_org_role",
    "read_organization_custom_repo_role"
  ]
}
