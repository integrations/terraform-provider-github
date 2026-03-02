provider "github" {
  owner     = var.owner
  auth_mode = "app"
  # Credentials are specified through environment variables:
  # GITHUB_APP_ID, GITHUB_APP_INSTALLATION_ID, GITHUB_APP_PEM_FILE
}

terraform {
  required_providers {
    github = {
      source = "integrations/github"
    }
  }
}
