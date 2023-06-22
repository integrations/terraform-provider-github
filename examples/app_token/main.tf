terraform {
  required_providers {
    github = {
      source = "integrations/github"
    }
  }
}

provider "github" {}

data "github_app_token" "this" {
  app_id          = var.app_id
  installation_id = var.installation_id
  pem_file        = file(var.pem_file_path)
}
