resource "github_team" "security_managers" {
  name        = var.team_name
  description = "A team of organization security managers"
}

resource "github_organization_security_manager" "security_managers" {
  team_slug = github_team.security_managers.slug
}
