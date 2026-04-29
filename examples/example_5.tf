provider "github" {
  owner = var.github_organization
  app_auth {} # When using `GITHUB_APP_XXX` environment variables
}
