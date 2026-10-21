data "github_enterprise" "enterprise" {
  slug = "my-enterprise"
}

resource "github_enterprise_team" "example" {
  enterprise_slug             = data.github_enterprise.enterprise.slug
  name                        = "Platform"
  description                 = "Platform Engineering"
  organization_selection_type = "selected"
}
