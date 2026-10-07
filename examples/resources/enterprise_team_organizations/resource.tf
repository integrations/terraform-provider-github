data "github_enterprise" "enterprise" {
  slug = "my-enterprise"
}

resource "github_enterprise_team" "team" {
  enterprise_slug             = data.github_enterprise.enterprise.slug
  name                        = "Platform"
  organization_selection_type = "selected"
}

resource "github_enterprise_team_organizations" "assignments" {
  enterprise_slug = data.github_enterprise.enterprise.slug
  team_id         = github_enterprise_team.team.team_id

  organization_slugs = [
    "my-org",
    "another-org",
  ]
}
