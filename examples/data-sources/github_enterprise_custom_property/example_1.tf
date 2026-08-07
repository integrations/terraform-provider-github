data "github_enterprise_custom_property" "security_tier" {
  enterprise_slug = "my-enterprise"
  property_name   = "securityTier"
}
