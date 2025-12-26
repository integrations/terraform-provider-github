terraform {
  required_providers {
    github = {
      source  = "integrations/github"
      version = "~> 6.0"
    }
  }
}

provider "github" {
  token = var.github_token
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

resource "github_enterprise_cost_center" "example" {
  enterprise_slug = var.enterprise_slug
  name            = var.cost_center_name

  # Authoritative assignments: Terraform will add/remove to match these lists.
  users         = var.users
  organizations = var.organizations
  repositories  = var.repositories
}

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

output "cost_center_resources" {
  description = "Effective assignments (read from API)"
  value = {
    users         = sort(tolist(github_enterprise_cost_center.example.users))
    organizations = sort(tolist(github_enterprise_cost_center.example.organizations))
    repositories  = sort(tolist(github_enterprise_cost_center.example.repositories))
  }
}

output "cost_center_from_data_source" {
  description = "Cost center fetched by data source"
  value = {
    id            = data.github_enterprise_cost_center.by_id.cost_center_id
    name          = data.github_enterprise_cost_center.by_id.name
    state         = data.github_enterprise_cost_center.by_id.state
    users         = sort(tolist(data.github_enterprise_cost_center.by_id.users))
    organizations = sort(tolist(data.github_enterprise_cost_center.by_id.organizations))
    repositories  = sort(tolist(data.github_enterprise_cost_center.by_id.repositories))
  }
}
