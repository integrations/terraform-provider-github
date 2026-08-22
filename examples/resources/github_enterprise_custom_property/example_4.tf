resource "github_enterprise_custom_property" "team_contact" {
  enterprise_slug    = "my-enterprise"
  property_name      = "teamContact"
  value_type         = "string"
  required           = false
  description        = "Contact information for the team managing this repository"
  values_editable_by = "org_and_repo_actors"
}
