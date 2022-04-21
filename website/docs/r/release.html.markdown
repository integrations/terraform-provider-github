---
layout: "github"
page_title: "GitHub: github_release"
description: |-
  Creates and manages releases within a single GitHub repository
---

# github_release

This resource allows you to create and manage a release in a specific
GitHub repository.

## Example Usage

```hcl
resource "github_repository" "repo" {
  name         = "repo"
  description  = "GitHub repo managed by Terraform"

  private = false
}

resource "github_release" "example" {
  repository = github_repository.repo.name
  tag_name   = "v1.0.0"
}
```

## Example Usage on Non-Default Branch

```hcl
resource "github_repository" "example" {
  name      = "repo"
  auto_init = true
}

resource "github_branch" "example" {
  repository    = github_repository.example.name
  branch        = "branch_name"
  source_branch = github_repository.example.default_branch
}

resource "github_release" "example" {
  repository       = github_repository.example.name
  tag_name         = "v1.0.0"
  target_commitish = github_branch.example.branch
  draft	           = false
  prerelease       = false
}
```

## Argument Reference

The following arguments are supported:

* `repository` - (Required) The name of the repository.

* `tag_name` - (Required) The name of the tag.

* `target_commitish` - (Optional) The branch name or commit SHA the tag is created from. Defaults to the default branch of the repository.

* `name` - (Optional) The name of the release.

* `body` - (Optional) Text describing the contents of the tag.

* `draft` - (Optional) Set to `false` to create a published release.

* `prerelease` - (Optional) Set to `false` to identify the release as a full release.

* `generate_release_notes` - (Optional) Set to `true` to automatically generate the name and body for this release. If `name` is specified, the specified `name` will be used; otherwise, a name will be automatically generated. If `body` is specified, the `body` will be pre-pended to the automatically generated notes.

* `discussion_category_name` - (Optional) If specified, a discussion of the specified category is created and linked to the release. The value must be a category that already exists in the repository. For more information, see [Managing categories for discussions in your repository](https://docs.github.com/discussions/managing-discussions-for-your-community/managing-categories-for-discussions-in-your-repository).

## Attributes Reference

The following additional attributes are exported:

* `release_id` - The ID of the release.

* `created_at` - This is the date of the commit used for the release, and not the date when the release was drafted or published.

* `published_at` - This is the date when the release was published. This will be empty if the release is a draft.

* `html_url` - URL of the release in GitHub. 

* `url` - URL that can be provided to API calls that reference this release.

* `assets_url` - URL that can be provided to API calls displaying the attached assets to this release.

* `upload_url` - URL that can be provided to API calls to upload assets.

* `zipball_url` - URL that can be provided to API calls to fetch the release ZIP archive.

* `tarball_url` - URL that can be provided to API calls to fetch the release TAR archive.

* `node_id` - GraphQL global node id for use with v4 API

## Import

This resource can be imported using the `name` of the repository, combined with the `id` of the release, and a `:` character for separating components, e.g.

```sh
$ terraform import github_release.example repo:12345678
```
