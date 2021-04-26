provider "github" {
  owner = var.owner
  app_auth {
    // Empty block to allow the provider configurations to be specified through
    // environment variables.
    // See: https://github.com/hashicorp/terraform-plugin-sdk/issues/142
  }
}

terraform {
  required_providers {
    github = {
      source  = "integrations/github"
    }
  }
}
