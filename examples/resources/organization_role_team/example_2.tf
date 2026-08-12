locals {
  security_manager_id = one([for x in data.github_organization_roles.all_roles.roles : x.role_id if x.name == "security_manager"])
}

data "github_organization_roles" "all_roles" {}

resource "github_organization_role_team" "security_managers" {
  role_id   = local.security_manager_id
  team_slug = "example-team"
}
