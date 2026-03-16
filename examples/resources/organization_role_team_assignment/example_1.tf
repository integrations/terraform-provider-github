resource "github_team" "test-team" {
  name     = "test-team"
}

resource "github_organization_role_team_assignment" "test-team-role-assignment" {
  team_slug = github_team.test-team.slug
  role_id   = "8132" # all_repo_read (predefined)
}
