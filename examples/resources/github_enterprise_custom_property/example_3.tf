resource "github_enterprise_custom_property" "contains_pii" {
  enterprise_slug = "my-enterprise"
  property_name   = "containsPII"
  value_type      = "true_false"
  required        = false
  description     = "Whether this repository contains personally identifiable information"
  default_values  = ["false"]
}
