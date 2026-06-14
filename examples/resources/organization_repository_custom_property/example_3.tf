resource "github_organization_repository_custom_property" "archived" {
  property_name = "archived"
  value_type    = "true_false"
  description   = "Whether this repository is archived"
  default_value = "false"
}
