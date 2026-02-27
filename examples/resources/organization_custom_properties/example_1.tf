resource "github_organization_custom_properties" "environment" {
  property_name = "environment"
  value_type    = "single_select"
  required      = true
  description   = "The deployment environment for this repository"
  default_value = "development"
  allowed_values = [
    "development",
    "staging", 
    "production"
  ]
}
