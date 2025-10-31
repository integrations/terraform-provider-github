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
  description = "GitHub personal access token with enterprise admin permissions"
  type        = string
  sensitive   = true
}

variable "enterprise_slug" {
  description = "The GitHub Enterprise slug"
  type        = string
}

# Basic Enterprise Settings with minimal configuration
resource "github_enterprise_settings" "basic" {
  enterprise_slug = var.enterprise_slug

  # Allow all actions for all organizations
  actions_enabled_organizations = "all"
  actions_allowed_actions       = "all"

  # Use restrictive workflow permissions
  default_workflow_permissions     = "read"
  can_approve_pull_request_reviews = false
}

# Advanced Enterprise Settings with selective permissions  
resource "github_enterprise_settings" "advanced" {
  enterprise_slug = var.enterprise_slug

  # Enable actions for selected organizations only
  actions_enabled_organizations = "selected"
  
  # Allow only selected actions
  actions_allowed_actions = "selected"
  
  # Only allow GitHub-owned and verified actions
  actions_github_owned_allowed = true
  actions_verified_allowed     = true
  
  # Allow specific action patterns
  actions_patterns_allowed = [
    "actions/cache@*",
    "actions/checkout@*", 
    "actions/setup-node@*",
    "actions/setup-python@*",
    "actions/upload-artifact@*",
    "actions/download-artifact@*",
    "my-org/custom-action@v1"
  ]

  # Grant write permissions to workflows
  default_workflow_permissions     = "write"
  can_approve_pull_request_reviews = true
}

output "basic_enterprise_settings" {
  description = "Basic enterprise settings configuration"
  value = {
    enterprise_slug                  = github_enterprise_settings.basic.enterprise_slug
    actions_enabled_organizations    = github_enterprise_settings.basic.actions_enabled_organizations
    actions_allowed_actions         = github_enterprise_settings.basic.actions_allowed_actions
    default_workflow_permissions    = github_enterprise_settings.basic.default_workflow_permissions
    can_approve_pull_request_reviews = github_enterprise_settings.basic.can_approve_pull_request_reviews
  }
}

output "advanced_enterprise_settings" {
  description = "Advanced enterprise settings configuration"
  value = {
    enterprise_slug                  = github_enterprise_settings.advanced.enterprise_slug
    actions_enabled_organizations    = github_enterprise_settings.advanced.actions_enabled_organizations
    actions_allowed_actions         = github_enterprise_settings.advanced.actions_allowed_actions
    actions_github_owned_allowed    = github_enterprise_settings.advanced.actions_github_owned_allowed
    actions_verified_allowed        = github_enterprise_settings.advanced.actions_verified_allowed
    actions_patterns_allowed        = github_enterprise_settings.advanced.actions_patterns_allowed
    default_workflow_permissions    = github_enterprise_settings.advanced.default_workflow_permissions
    can_approve_pull_request_reviews = github_enterprise_settings.advanced.can_approve_pull_request_reviews
  }
}