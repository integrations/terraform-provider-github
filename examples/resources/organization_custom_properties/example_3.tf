resource "github_organization_custom_properties" "owner" {
  property_name = "owner"
  value_type    = "string"
  required      = true
  description   = "The team or individual responsible for this repository"
}
