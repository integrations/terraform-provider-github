# Add a repository to the team
resource "github_team" "some_team" {
  name        = "SomeTeam"
  description = "Some cool team"
}

resource "github_repository" "some_repo" {
  name = "some-repo"
}

resource "github_team_repository" "some_team_repo" {
  team_id    = github_team.some_team.id
  repository = github_repository.some_repo.name
  permission = "pull"
}
