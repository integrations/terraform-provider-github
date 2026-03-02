provider "github" {
  owner     = var.owner
  auth_mode = "app"
  # Credentials are specified through environment variables:
  # GITHUB_APP_ID, GITHUB_APP_INSTALLATION_ID, GITHUB_APP_PRIVATE_KEY
}

terraform {
  required_providers {
    github = {
      source = "integrations/github"
    }
  }
}
