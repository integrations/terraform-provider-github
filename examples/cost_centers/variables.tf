variable "github_token" {
  description = "GitHub classic personal access token (PAT) for an enterprise admin"
  type        = string
  sensitive   = true
}

variable "github_owner" {
  description = "GitHub organization or user used to configure the provider"
  type        = string
}

variable "enterprise_slug" {
  description = "GitHub Enterprise slug"
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
  description = "Repositories to assign to the cost center, in owner/name format"
  type        = list(string)
  default     = []
}
