# Token-authenticated npm registry, available to every repository of the organization.
resource "github_organization_private_registry" "example" {
  registry_type = "npm_registry"
  url           = "https://npm.pkg.github.com"
  auth_type     = "token"
  value         = var.npm_token
  visibility    = "all"
}
