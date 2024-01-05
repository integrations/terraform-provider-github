provider "github" {
  owner = var.organization
  token = var.github_token
}

terraform {
  required_providers {
    github = {
      source = "integrations/github"
    }
  }
}