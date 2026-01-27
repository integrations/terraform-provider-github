terraform {
  required_providers {
    github = {
      source  = "integrations/github"
      version = "~> 6.11"
    }
  }
}

provider "github" {
  token = var.github_token
  owner = var.enterprise_slug
}

variable "github_token" {
  description = "GitHub classic personal access token (PAT) for an enterprise admin"
  type        = string
  sensitive   = true
}

variable "enterprise_slug" {
  description = "The GitHub Enterprise slug"
  type        = string
}

variable "cost_center_name" {
  description = "Name for the cost center"
  type        = string
}

variable "users" {
  description = "Usernames to assign to the cost center"
  type        = list(string)
  default     = []
}

variable "organizations" {
  description = "Organization logins to assign to the cost center"
  type        = list(string)
  default     = []
}

variable "repositories" {
  description = "Repositories (full name, e.g. org/repo) to assign to the cost center"
  type        = list(string)
  default     = []
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

output "cost_center" {
  description = "Created cost center"
  value = {
    id                 = github_enterprise_cost_center.example.id
    name               = github_enterprise_cost_center.example.name
    state              = github_enterprise_cost_center.example.state
    azure_subscription = github_enterprise_cost_center.example.azure_subscription
  }
}

output "cost_center_from_data_source" {
  description = "Cost center fetched by data source (includes all assignments)"
  value = {
    id            = data.github_enterprise_cost_center.by_id.cost_center_id
    name          = data.github_enterprise_cost_center.by_id.name
    state         = data.github_enterprise_cost_center.by_id.state
    users         = sort(tolist(data.github_enterprise_cost_center.by_id.users))
    organizations = sort(tolist(data.github_enterprise_cost_center.by_id.organizations))
    repositories  = sort(tolist(data.github_enterprise_cost_center.by_id.repositories))
  }
}
