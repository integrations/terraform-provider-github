---
layout: "github"
page_title: "GitHub: github_user_invitation_accepter"
description: |-
  Provides a resource to manage GitHub repository collaborator invitations.
---

# github_user_invitation_accepter

Provides a resource to manage GitHub repository collaborator invitations.

## Example Usage

```hcl
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
```

## Argument Reference

The following arguments are supported:

* `invitation_id` - (Required) ID of the invitation to accept
