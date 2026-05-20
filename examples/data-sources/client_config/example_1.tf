data "github_client_config" "current" {}

output "owner" {
  value = data.github_client_config.current.owner
}

output "username" {
  value = data.github_client_config.current.username
}
