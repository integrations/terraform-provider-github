resource "github_organization_repository_custom_property" "environment" {
  property_name = "environment"
  value_type    = "single_select"
  required      = true
  description   = "The deployment environment for this repository"
  default_value = "development"
  allowed_values = [
    "development",
    "staging",
    "production",
  ]
}
