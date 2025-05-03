variable "github_token" {
  description = "GitHub access token used to configure the provider"
  type        = string
}

variable "owner" {
  description = "GitHub owner used to configure the provider"
  type        = string
}

variable "team_name" {
  description = "The name to use for the GitHub team"
  type        = string
}
