resource "github_team" "some_team" {
  name        = "SomeTeam"
  description = "Some cool team"
}

resource "github_organization_security_manager" "some_team" {
  team_slug = github_team.some_team.slug
}
