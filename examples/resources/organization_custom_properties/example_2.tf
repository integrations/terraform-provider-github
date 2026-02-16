resource "github_organization_custom_properties" "team_contact" {
  property_name      = "team_contact"
  value_type         = "string"
  required           = false
  description        = "Contact information for the team managing this repository"
  values_editable_by = "org_and_repo_actors"
}
