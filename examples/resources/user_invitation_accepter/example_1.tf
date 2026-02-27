resource "github_repository" "example" {
  name = "example-repo"
}

resource "github_repository_collaborator" "example" {
  repository = github_repository.example.name
  username   = "example-username"
  permission = "push"
}

provider "github" {
  alias = "invitee"
  token = var.invitee_token
}

resource "github_user_invitation_accepter" "example" {
  provider      = "github.invitee"
  invitation_id = github_repository_collaborator.example.invitation_id
}
