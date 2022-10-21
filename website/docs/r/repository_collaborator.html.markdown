---
layout: "github"
page_title: "GitHub: github_repository_collaborator"
description: |-
  Provides a GitHub repository collaborator resource.
---

# github_repository_collaborator

Provides a GitHub repository collaborator resource.

This resource allows you to add/remove collaborators from repositories in your
organization or personal account. For organization repositories, collaborators can
have explicit (and differing levels of) read, write, or administrator access to 
specific repositories, without giving the user full organization membership. 
For personal repositories, collaborators can only be granted write
(implicitly includes read) permission. 

When applied, an invitation will be sent to the user to become a collaborator
on a repository. When destroyed, either the invitation will be cancelled or the
collaborator will be removed from the repository.

Further documentation on GitHub collaborators:

- [Adding outside collaborators to your personal repositories](https://help.github.com/en/github/setting-up-and-managing-your-github-user-account/managing-access-to-your-personal-repositories)
- [Adding outside collaborators to repositories in your organization](https://help.github.com/articles/adding-outside-collaborators-to-repositories-in-your-organization/)
- [Converting an organization member to an outside collaborator](https://help.github.com/articles/converting-an-organization-member-to-an-outside-collaborator/)

## Example Usage

```hcl
# Add a collaborator to a repository
resource "github_repository_collaborator" "a_repo_collaborator" {
  repository = "our-cool-repo"
  username   = "SomeUser"
  permission = "admin"
}
```

## Argument Reference

The following arguments are supported:

* `repository` - (Required) The GitHub repository
* `username` - (Required) The user to add to the repository as a collaborator.
* `permission` - (Optional) The permission of the outside collaborator for the repository.
            Must be one of `pull`, `push`, `maintain`, `triage` or `admin` for organization-owned repositories.
            Must be `push` for personal repositories. Defaults to `push`.
* `permission_diff_suppression` - (Optional) Suppress plan diffs for `triage` and `maintain`.  Defaults to `false`.

## Attribute Reference

In addition to the above arguments, the following attributes are exported:

* `invitation_id` - ID of the invitation to be used in [`github_user_invitation_accepter`](./user_invitation_accepter.html)

## Import

GitHub Repository Collaborators can be imported using an ID made up of `repository:username`, e.g.

```
$ terraform import github_repository_collaborator.collaborator terraform:someuser
```