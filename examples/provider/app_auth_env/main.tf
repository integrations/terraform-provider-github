provider "github" {
  owner = var.github_organization
  # Credentials come from the `GITHUB_APP_XXX` environment variables.
}
