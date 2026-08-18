provider "github" {
  owner = var.github_organization

  auth_mode = "app" # Credentials required to come from the `GITHUB_APP_XXX` environment variables.
}
