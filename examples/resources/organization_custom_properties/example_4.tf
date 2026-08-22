resource "github_organization_custom_properties" "archived" {
  property_name = "archived"
  value_type    = "true_false"
  required      = false
  description   = "Whether this repository is archived"
  default_value = "false"
}
