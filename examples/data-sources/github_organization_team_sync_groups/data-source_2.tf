# Retrieve IdP groups whose names begin with a given prefix.
data "github_organization_team_sync_groups" "filtered" {
  prefix_filter = "myprefix_"
}
