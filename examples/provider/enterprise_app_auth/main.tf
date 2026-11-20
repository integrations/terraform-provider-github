terraform {
  required_providers {
    github = {
      source = "integrations/github"
    }
  }
}

provider "github" {
  auth_mode = "app"

  app_auth {
    id              = var.creation_app_id
    installation_id = var.creation_installation_id
    pem_file        = var.creation_private_key
  }

  enterprise_app_auth {
    enterprise_slug           = var.enterprise_slug
    installer_app_id          = var.installer_app_id
    installer_installation_id = var.installer_installation_id
    installer_pem_file        = var.installer_private_key
    automation_app_id         = var.automation_app_id
    automation_client_id      = var.automation_client_id
    automation_pem_file       = var.automation_private_key
  }
}

data "github_enterprise" "selected" {
  slug = var.enterprise_slug
}

resource "github_enterprise_organization" "managed" {
  for_each = var.organizations

  enterprise_id = data.github_enterprise.selected.id
  name          = each.key
  billing_email = each.value.billing_email
  admin_logins  = each.value.admin_logins

  lifecycle {
    prevent_destroy = true
  }
}

variable "enterprise_slug" {
  type        = string
  description = "Enterprise managed by the creation and installer Apps."
}

variable "creation_app_id" {
  type        = string
  description = "Enterprise creation App ID or client ID."
}

variable "creation_installation_id" {
  type        = string
  description = "Enterprise creation App installation ID."
}

variable "creation_private_key" {
  type        = string
  sensitive   = true
  description = "Creation App private key PEM content."
}

variable "installer_app_id" {
  type        = string
  description = "Installer App ID or client ID."
}

variable "installer_installation_id" {
  type        = string
  description = "Installer enterprise installation ID."
}

variable "installer_private_key" {
  type        = string
  sensitive   = true
  description = "Installer App private key PEM content."
}

variable "automation_app_id" {
  type        = string
  description = "Numeric automation App ID."
}

variable "automation_client_id" {
  type        = string
  description = "Automation App client ID matching the numeric ID and private key."
}

variable "automation_private_key" {
  type        = string
  sensitive   = true
  description = "Automation App private key PEM content."
}

variable "organizations" {
  type = map(object({
    billing_email = string
    admin_logins  = set(string)
  }))
  description = "Organizations to create or import into the configured enterprise."
}