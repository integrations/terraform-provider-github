---
layout: "github"
page_title: "GitHub: github_user"
description: |-
  Provides an import-only GitHub user resource.
---

# github\_user

Provides an import-only GitHub user resource.

This resource allows you to refer to a GitHub user in your configuration,
automatically obtaining its current username (login) when needed, and
providing its unique ID to be used in other resources which refer
to the user.

Note: because this resource is import-only, it can neither be
created nor destroyed (and since it does not have any arguments,
no attempts to update it will be made). If you add this resource to your
configuration but fail to import it before executing a `plan` or `apply`
operation, an error message will be produced. If you have this resource
in your configuration and execute a `destroy` operation, no attempt will
be made to destroy this resource.

## Example Usage

```hcl
resource "github_user" "example" {
}
```

## Attributes Reference

 * `id` - The unique ID (numeric).
 * `username` - The username (login).

## Import

GitHub users can be imported using either their username or unique (numeric) ID, e.g.

```
$ terraform import github_user.example octocat

$ terraform import github_user.example 1234567
```
