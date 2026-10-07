data "github_enterprise" "enterprise" {
  slug = "my-enterprise"
}

resource "github_enterprise_team" "team" {
  enterprise_slug = data.github_enterprise.enterprise.slug
  name            = "Platform"
}

resource "github_enterprise_team_membership" "member" {
  enterprise_slug = data.github_enterprise.enterprise.slug
  team_id         = github_enterprise_team.team.team_id
  username        = "octocat"
}
