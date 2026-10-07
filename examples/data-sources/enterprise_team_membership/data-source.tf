data "github_enterprise_team_membership" "example" {
  enterprise_slug = "my-enterprise"
  team_slug       = "ent:platform"
  username        = "octocat"
}
