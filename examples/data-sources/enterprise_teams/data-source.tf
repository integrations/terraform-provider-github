data "github_enterprise_teams" "all" {
  enterprise_slug = "my-enterprise"
}

output "enterprise_team_slugs" {
  value = [for t in data.github_enterprise_teams.all.teams : t.slug]
}
