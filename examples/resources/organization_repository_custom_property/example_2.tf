resource "github_organization_repository_custom_property" "team_contact" {
  property_name      = "team_contact"
  value_type         = "string"
  description        = "Contact information for the team managing this repository"
  values_editable_by = "org_and_repo_actors"
}
