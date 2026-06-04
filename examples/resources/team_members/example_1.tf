# Add a user to the organization
resource "github_membership" "membership_for_some_user" {
  username = "SomeUser"
  role     = "member"
}

resource "github_membership" "membership_for_another_user" {
  username = "AnotherUser"
  role     = "member"
}

resource "github_team" "some_team" {
  name        = "SomeTeam"
  description = "Some cool team"
}

resource "github_team_members" "some_team_members" {
  team_id = github_team.some_team.id

  members {
    username = "SomeUser"
    role     = "maintainer"
  }

  members {
    username = "AnotherUser"
    role     = "member"
  }
}
