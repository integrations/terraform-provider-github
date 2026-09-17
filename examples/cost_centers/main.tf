terraform {
  required_providers {
    github = {
      source  = "integrations/github"
      version = ">= 6.11"
    }
  }
}

provider "github" {
  token = var.github_token
  owner = var.github_owner
}

# The cost center resource manages only the cost center entity itself.
resource "github_enterprise_cost_center" "example" {
  enterprise_slug = var.enterprise_slug
  name            = var.cost_center_name
}

# Use separate authoritative resources for assignments.
# These are optional - only create them if you have items to assign.

resource "github_enterprise_cost_center_users" "example" {
  count = length(var.users) > 0 ? 1 : 0

  enterprise_slug = var.enterprise_slug
  cost_center_id  = github_enterprise_cost_center.example.id
  usernames       = var.users
}

resource "github_enterprise_cost_center_organizations" "example" {
  count = length(var.organizations) > 0 ? 1 : 0

  enterprise_slug     = var.enterprise_slug
  cost_center_id      = github_enterprise_cost_center.example.id
  organization_logins = var.organizations
}

resource "github_enterprise_cost_center_repositories" "example" {
  count = length(var.repositories) > 0 ? 1 : 0

  enterprise_slug  = var.enterprise_slug
  cost_center_id   = github_enterprise_cost_center.example.id
  repository_names = var.repositories
}

# Data sources for reading cost center information
data "github_enterprise_cost_center" "by_id" {
  enterprise_slug = var.enterprise_slug
  cost_center_id  = github_enterprise_cost_center.example.id
}

data "github_enterprise_cost_centers" "active" {
  enterprise_slug = var.enterprise_slug
  state           = "active"

  depends_on = [github_enterprise_cost_center.example]
}
