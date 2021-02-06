variable "installation_id" {
  description = "ID of an app installation in an organization"
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
