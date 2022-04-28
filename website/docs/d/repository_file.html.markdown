---
layout: "github"
page_title: "GitHub: github_repository_file"
description: |-
  Reads files within a GitHub repository
---

# github_repository_file

This data source allows you to read files within a
GitHub repository.


## Example Usage

```hcl
data "github_repository_file" "foo" {
  repository          = github_repository.foo.name
  branch              = "main"
  file                = ".gitignore"
}

```


## Argument Reference

The following arguments are supported:

* `repository` - (Required) The repository to create the file in.

* `file` - (Required) The path of the file to manage.

* `branch` - (Optional) Git branch (defaults to the repository's default branch).

## Attributes Reference

The following additional attributes are exported:

* `content` - The file content.

* `commit_sha` - The SHA of the commit that modified the file.

* `sha` - The SHA blob of the file.

* `commit_author` - Committer author name.

* `commit_email` - Committer email address.

* `commit_message` - Commit message when file was last updated.

* `ref` - The name of the commit/branch/tag.
