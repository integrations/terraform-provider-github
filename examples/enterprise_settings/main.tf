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

# Basic Enterprise Actions Permissions - Allow all actions for all organizations
resource "github_enterprise_actions_permissions" "basic" {
  enterprise_slug = var.enterprise_slug
  
  enabled_organizations = "all"
  allowed_actions      = "all"
}

# Basic Enterprise Workflow Permissions - Restrictive settings
resource "github_enterprise_actions_workflow_permissions" "basic" {
  enterprise_slug = var.enterprise_slug

  default_workflow_permissions     = "read"
  can_approve_pull_request_reviews = false
}

# Advanced Enterprise Actions Permissions - Selective configuration
resource "github_enterprise_actions_permissions" "advanced" {
  enterprise_slug = var.enterprise_slug
  
  enabled_organizations = "selected"
  allowed_actions      = "selected"
  
  # Configure allowed actions when "selected" policy is used
  allowed_actions_config {
    github_owned_allowed = true
    verified_allowed     = true
    patterns_allowed = [
      "actions/cache@*",
      "actions/checkout@*", 
      "actions/setup-node@*",
      "actions/setup-python@*",
      "actions/upload-artifact@*",
      "actions/download-artifact@*",
      "my-org/custom-action@v1"
    ]
  }
  
  # Configure enabled organizations when "selected" policy is used
  enabled_organizations_config {
    organization_ids = [123456, 789012] # Replace with actual org IDs
  }
}

# Advanced Enterprise Workflow Permissions - Permissive settings
resource "github_enterprise_actions_workflow_permissions" "advanced" {
  enterprise_slug = var.enterprise_slug

  default_workflow_permissions     = "write"
  can_approve_pull_request_reviews = true
}

# Security Analysis Settings - Enable security features for new repositories
resource "github_enterprise_security_analysis_settings" "example" {
  enterprise_slug = var.enterprise_slug
  
  advanced_security_enabled_for_new_repositories             = true
  secret_scanning_enabled_for_new_repositories               = true
  secret_scanning_push_protection_enabled_for_new_repositories = true
  secret_scanning_validity_checks_enabled                   = true
  secret_scanning_push_protection_custom_link               = "https://octokit.com/security-help"
}

output "basic_enterprise_actions" {
  description = "Basic enterprise actions permissions configuration"
  value = {
    enterprise_slug       = github_enterprise_actions_permissions.basic.enterprise_slug
    enabled_organizations = github_enterprise_actions_permissions.basic.enabled_organizations
    allowed_actions      = github_enterprise_actions_permissions.basic.allowed_actions
  }
}

output "basic_enterprise_workflow" {
  description = "Basic enterprise workflow permissions configuration"
  value = {
    enterprise_slug                  = github_enterprise_actions_workflow_permissions.basic.enterprise_slug
    default_workflow_permissions     = github_enterprise_actions_workflow_permissions.basic.default_workflow_permissions
    can_approve_pull_request_reviews = github_enterprise_actions_workflow_permissions.basic.can_approve_pull_request_reviews
  }
}

output "advanced_enterprise_actions" {
  description = "Advanced enterprise actions permissions configuration"
  value = {
    enterprise_slug       = github_enterprise_actions_permissions.advanced.enterprise_slug
    enabled_organizations = github_enterprise_actions_permissions.advanced.enabled_organizations
    allowed_actions      = github_enterprise_actions_permissions.advanced.allowed_actions
  }
}

output "advanced_enterprise_workflow" {
  description = "Advanced enterprise workflow permissions configuration"
  value = {
    enterprise_slug                  = github_enterprise_actions_workflow_permissions.advanced.enterprise_slug
    default_workflow_permissions     = github_enterprise_actions_workflow_permissions.advanced.default_workflow_permissions
    can_approve_pull_request_reviews = github_enterprise_actions_workflow_permissions.advanced.can_approve_pull_request_reviews
  }
}
