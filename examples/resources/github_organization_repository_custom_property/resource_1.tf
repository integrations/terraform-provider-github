# single_select property with a default value
resource "github_organization_repository_custom_property" "environment" {
  property_name = "environment"
  value_type    = "single_select"
  required      = true
  description   = "The deployment environment for this repository"
  default_value = ["development"]
  allowed_values = [
    "development",
    "staging",
    "production",
  ]
}

# string property that repository actors (not just org owners) can edit
resource "github_organization_repository_custom_property" "team_contact" {
  property_name      = "team_contact"
  value_type         = "string"
  description        = "Contact information for the team managing this repository"
  values_editable_by = "org_and_repo_actors"
}

# true_false property
resource "github_organization_repository_custom_property" "archived" {
  property_name = "archived"
  value_type    = "true_false"
  description   = "Whether this repository is archived"
  default_value = ["false"]
}

# multi_select property; only this type accepts more than one default value
resource "github_organization_repository_custom_property" "compliance" {
  property_name  = "compliance"
  value_type     = "multi_select"
  description    = "Compliance regimes this repository is in scope for"
  allowed_values = ["pci", "sox", "hipaa"]
  default_value  = ["pci", "sox"]
}
