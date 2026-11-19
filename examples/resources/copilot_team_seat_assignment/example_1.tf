resource "github_team" "example" {
  name = "my-copilot-team"
}

resource "github_copilot_team_seat_assignment" "example" {
  team = github_team.example.slug
}
