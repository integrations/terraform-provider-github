resource "github_organization_private_registry" "my_registry" {
  registry_type = "npm_registry"
  url           = "https://npm.pkg.github.com"
  auth_type     = "username_password"
  username      = "github-actions"
  value         = "super_secret_token_123"
  visibility    = "private"
}
