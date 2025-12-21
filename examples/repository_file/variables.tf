variable "organization" {
  description = "GitHub organization used to configure the provider"
  type        = string
}

variable "github_token" {
  description = "GitHub access token used to configure the provider"
  type        = string
}

variable "repository" {
  description = "The name of the repository"
  type        = string
}

variable "file" {
  description = "The name of the file to create"
  type        = string
}
variable "content" {
  description = "The content of the file to create"
  type        = string
}
variable "branch" {
  description = "The branch to create the file in"
  type        = string
  default     = "main"
}
variable "commit_author" {
  description = "The name of the author of the commit"
  type        = string
  default     = ""
}
variable "commit_message" {
  description = "The message of the commit"
  type        = string
  default     = ""
}
variable "commit_email" {
  description = "The email of the author of the commit"
  type        = string
  default     = ""
}
