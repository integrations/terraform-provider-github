terraform {
  required_providers {
    github = {
      source  = "integrations/github"
      version = "~> 6.0"
    }
  }
}

provider "github" {
  owner = "integrations"
}

data "github_repository" "example" {
  name = "terraform-provider-github"
}
