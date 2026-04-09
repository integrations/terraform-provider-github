terraform {
  required_providers {
    github = {
      source = "integrations/github"
    }
  }
}

provider "github" {
  owner = var.github_owner
}

variable "github_owner" {
  description = "The GitHub organization to manage"
  type        = string
}

# Step 1: Define an org-level custom property
resource "github_organization_custom_properties" "environment" {
  property_name  = "environment"
  value_type     = "single_select"
  required       = false
  allowed_values = ["production", "staging", "development"]
}

# Step 2: Include the custom property in OIDC tokens
resource "github_actions_organization_oidc_custom_property_inclusion" "environment" {
  custom_property_name = "environment"
  depends_on           = [github_organization_custom_properties.environment]
}

# Step 3: Read back the inclusions via data source
data "github_actions_organization_oidc_custom_property_inclusions" "current" {
  depends_on = [github_actions_organization_oidc_custom_property_inclusion.environment]
}

output "included_properties" {
  value = data.github_actions_organization_oidc_custom_property_inclusions.current.custom_property_names
}
