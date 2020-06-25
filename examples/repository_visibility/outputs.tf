output "public_repository" {
  description = "Example repository JSON blob"
  value       = github_repository.public
}

output "private_repository" {
  description = "Example repository JSON blob"
  value       = github_repository.private
}

output "internal_repository" {
  description = "Example repository JSON blob"
  value       = github_repository.internal
}
