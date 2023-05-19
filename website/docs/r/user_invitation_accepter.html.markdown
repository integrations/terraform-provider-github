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

## Allowing empty invitation IDs

Set `allow_empty_id` when using `for_each` over a list of `github_repository_collaborator.invitation_id`'s.

This allows applying a module again when a new `github_repository_collaborator` resource is added to the `for_each` loop.
This is needed as the `github_repository_collaborator.invitation_id` will be empty after a state refresh when the invitation has been accepted.

Note that when an invitation is accepted manually or by another tool between a state refresh and a `terraform apply` using that refreshed state,
the plan will contain the invitation ID, but the apply will receive an HTTP 404 from the API since the invitation has already been accepted.

This is tracked in [#1157](https://github.com/integrations/terraform-provider-github/issues/1157).

## Argument Reference

The following arguments are supported:

* `invitation_id` - (Optional) ID of the invitation to accept. Must be set when `allow_empty_id` is `false`.
* `allow_empty_id` - (Optional) Allow the ID to be unset. This will result in the resource being skipped when the ID is not set instead of returning an error.
