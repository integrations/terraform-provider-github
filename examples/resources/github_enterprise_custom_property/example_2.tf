resource "github_enterprise_custom_property" "owner" {
  enterprise_slug = "my-enterprise"
  property_name   = "owningTeam"
  value_type      = "string"
  required        = true
  description     = "The team responsible for this repository"
}
