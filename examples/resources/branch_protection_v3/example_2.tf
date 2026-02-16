# Protect the main branch of the foo repository. Additionally, require that
# the "ci/check" check ran by the Github Actions app is passing and only allow
# the engineers team merge to the branch.

resource "github_branch_protection_v3" "example" {
  repository     = github_repository.example.name
  branch         = "main"
  enforce_admins = true

  required_status_checks {
    strict = false
    checks = [
      "ci/check:824642007264"
    ]
  }

  required_pull_request_reviews {
    dismiss_stale_reviews = true
    dismissal_users       = ["foo-user"]
    dismissal_teams       = [github_team.example.slug]
    dismissal_app         = ["foo-app"]

    bypass_pull_request_allowances {
      users = ["foo-user"]
      teams = [github_team.example.slug]
      apps  = ["foo-app"]
    }
  }

  restrictions {
    users = ["foo-user"]
    teams = [github_team.example.slug]
    apps  = ["foo-app"]
  }
}

resource "github_repository" "example" {
  name = "example"
}

resource "github_team" "example" {
  name = "Example Name"
}

resource "github_team_repository" "example" {
  team_id    = github_team.example.id
  repository = github_repository.example.name
  permission = "pull"
}
