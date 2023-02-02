provider "github" {
  owner = var.owner
  token = var.github_token
}


terraform {
  required_providers {
    github = {
      source  = "integrations/github"
      version = "~> 5.0"
    }
  }
}