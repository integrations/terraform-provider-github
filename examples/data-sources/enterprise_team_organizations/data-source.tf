data "github_enterprise_team_organizations" "example" {
  enterprise_slug = "my-enterprise"
  team_slug       = "ent:platform"
}

output "assigned_orgs" {
  value = data.github_enterprise_team_organizations.example.organization_slugs
}
