resource "github_enterprise_custom_property" "security_tier" {
  enterprise_slug = "my-enterprise"
  property_name   = "securityTier"
  value_type      = "single_select"
  required        = true
  description     = "Security classification tier for the repository"
  allowed_values  = ["tier1", "tier2", "tier3"]
}
