data "github_enterprise_app_accessible_organization_repositories" "example" {
  enterprise_slug = "my-enterprise"
  organization    = "my-org"
}
