---
layout: "github"
page_title: "GitHub: github_repository"
description: |-
  Get details about GitHub repository
---

# github_repository

Use this data source to retrieve information about a GitHub repository.

## Example Usage

```hcl
data "github_repository" "example" {
  full_name = "hashicorp/terraform"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) The name of the repository.

* `full_name` - (Optional) Full name of the repository (in `org/name` format).

## Attributes Reference

* `description` - A description of the repository.

* `homepage_url` - URL of a page describing the project.

* `private` - Whether the repository is private.

* `visibility` - Whether the repository is public, private or internal.

* `has_issues` - Whether the repository has GitHub Issues enabled.

* `has_projects` - Whether the repository has the GitHub Projects enabled.

* `has_wiki` - Whether the repository has the GitHub Wiki enabled.

* `allow_merge_commit` - Whether the repository allows merge commits.

* `allow_squash_merge` - Whether the repository allows squash merges.

* `allow_rebase_merge` - Whether the repository allows rebase merges.

* `has_downloads` - Whether the repository has Downloads feature enabled.

* `default_branch` - The name of the default branch of the repository.

* `archived` - Whether the repository is archived.

* `topics` - The list of topics of the repository.

* `html_url` - URL to the repository on the web.

* `ssh_clone_url` - URL that can be provided to `git clone` to clone the repository via SSH.

* `http_clone_url` - URL that can be provided to `git clone` to clone the repository via HTTPS.

* `git_clone_url` - URL that can be provided to `git clone` to clone the repository anonymously via the git protocol.

* `svn_url` - URL that can be provided to `svn checkout` to check out the repository via GitHub's Subversion protocol emulation.
