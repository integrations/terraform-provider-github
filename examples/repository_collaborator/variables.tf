variable "username" {
  description = "Username of a collaborator"
  type        = string
}

variable "permission" {
  description = "Permission level for a collaborator"
  type        = string
}

variable "organization" {
  description = "GitHub organization used to configure the provider"
  type        = string
}

variable "github_token" {
  description = "GitHub access token used to configure the provider"
  type        = string
}
