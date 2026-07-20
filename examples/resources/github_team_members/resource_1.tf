resource "github_team" "example" {
  name = "my-team"
}

resource "github_team_members" "example" {
  team_slug = github_team.example.slug

  members {
    username = "SomeUser"
    role     = "maintainer"
  }

  members {
    username = "AnotherUser"
    role     = "member"
  }
}
