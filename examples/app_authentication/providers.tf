provider "github" {
  owner = var.owner
  // GitHub app credentials can be specified through environment variables
  // An empty app_auth block can be used to avoid picking up GITHUB_TOKEN from the environment
  // app_auth {}
}

terraform {
  required_providers {
    github = {
      source = "integrations/github"
    }
  }
}
