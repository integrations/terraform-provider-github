# Add a repository to the team
resource "github_team" "some_team" {
  name        = "SomeTeam"
  description = "Some cool team"
}

resource "github_team_settings" "code_review_settings" {
  team_id = github_team.some_team.id
  review_request_delegation {
    algorithm    = "ROUND_ROBIN"
    member_count = 1
    notify       = true
  }
}
