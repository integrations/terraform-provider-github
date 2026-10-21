# Only find teams without a parent team (root teams only)

data "github_organization_teams" "example" {
  root_teams_only = true
}
