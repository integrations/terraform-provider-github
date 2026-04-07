data "github_enterprise_custom_property" "security_tier" {
  enterprise_slug = "my-enterprise"
  property_name   = "securityTier"
}

resource "github_repository" "example" {
  name       = "example"
  visibility = "private"

  custom_properties = {
    (data.github_enterprise_custom_property.security_tier.property_name) = "tier1"
  }
}
