resource "github_repository" "collaboration" {
  name        = "collaboration"
  visibility  = "private"
  description = "A collaborative repository"
}

resource "github_repository_collaborator" "collaborator" {
  repository = github_repository.collaboration.name
  username   = var.username
  permission = var.permission
}
