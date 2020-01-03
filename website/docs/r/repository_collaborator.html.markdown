---
layout: "github"
page_title: "GitHub: github_repository_collaborator"
description: |-
  Provides a GitHub repository collaborator resource.
---

# github_repository_collaborator

Provides a GitHub repository collaborator resource.

This resource allows you to add/remove collaborators from repositories in your
organization. Collaborators can have explicit (and differing levels of) read,
write, or administrator access to specific repositories in your organization,
without giving the user full organization membership.

When applied, an invitation will be sent to the user to become a collaborator
on a repository. When destroyed, either the invitation will be cancelled or the
collaborator will be removed from the repository.

Further documentation on GitHub collaborators:

- [Adding outside collaborators to repositories in your organization](https://help.github.com/articles/adding-outside-collaborators-to-repositories-in-your-organization/)
- [Converting an organization member to an outside collaborator](https://help.github.com/articles/converting-an-organization-member-to-an-outside-collaborator/)

## Example Usage

```hcl
# Add a data source to obtain user details for the user
data "github_user" "some_user" {
  username = "SomeUser"
}

# Add a collaborator to a repository
resource "github_repository_collaborator" "a_repo_collaborator" {
  repository = "our-cool-repo"
  user_id    = data.github_user.some_user.id
  permission = "admin"
}
```

## Argument Reference

The following arguments are supported:

* `repository` - (Required) The GitHub repository
* `user_id` - (Required) The GitHub user ID to add to as a collaborator.
* `permission` - (Optional) The permission of the outside collaborator for the repository.
            Must be one of `pull`, `push`, or `admin`. Defaults to `push`.

## Attribute Reference

* `invitation_id` - ID of the invitation to be used in [`github_user_invitation_accepter`](./user_invitation_accepter.html)
* `username` - The username (login) of the user specified by user_id

## Import

GitHub Repository Collaborators can be imported using an ID made up of two parts; a
repository name and a user identifier. The user can be specified using its
name (login) or numeric ID.

```
$ terraform import github_repository_collaborator.collaborator repo:1

$ terraform import github_repository_collaborator.collaborator org:octocat
```
