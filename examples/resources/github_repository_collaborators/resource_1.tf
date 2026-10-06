# Add collaborators to a repository
resource "github_team" "some_team" {
  name        = "SomeTeam"
  description = "Some cool team"
}

resource "github_repository" "some_repo" {
  name = "some-repo"
}

resource "github_repository_collaborators" "some_repo_collaborators" {
  repository = github_repository.some_repo.name

  user {
    permission = "admin"
    username   = "SomeUser"
  }

  team {
    permission = "pull"
    team_id    = github_team.some_team.slug
  }
}
