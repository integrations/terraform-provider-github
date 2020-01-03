---
layout: "github"
page_title: "GitHub: github_user"
description: |-
  Get information on a GitHub user.
---

# github\_user

Use this data source to retrieve information about a GitHub user.

## Example Usage

```hcl
data "github_user" "example" {
  username = "example"
}
```

## Argument Reference

* `username` - (Optional) The username.
* `user_id` - (Optional) The user ID.

One (and only one) of `username` and `user_id` must be supplied.

* `minimal` - (Optional) Boolean, setting this to true will stop
the data source from issuing the additional API calls needed to
obtain `gpg_keys` and `ssh_keys`.

## Attributes Reference

* `login` - the user's login.
* `avatar_url` - the user's avatar URL.
* `gravatar_id` - the user's gravatar ID.
* `site_admin` - whether the user is a GitHub admin.
* `name` - the user's full name.
* `company` - the user's company name.
* `blog` - the user's blog location.
* `location` - the user's location.
* `email` - the user's email.
* `gpg_keys` - list of user's GPG keys
* `ssh_keys` - list of user's SSH keys
* `bio` - the user's bio.
* `public_repos` - the number of public repositories.
* `public_gists` - the number of public gists.
* `followers` - the number of followers.
* `following` - the number of following users.
* `created_at` - the creation date.
* `updated_at` - the update date.
