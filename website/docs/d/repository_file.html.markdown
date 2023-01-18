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

* `repository` - (Required) The repository to read the file from.  
  if `full_name` is passed then its being split into `org` and `repo` (e.g. `bar/foo` split to org `bar` and repo `foo` )  
  if `name` is passed then `owner` is user

* `file` - (Required) The path of the file to manage.

* `branch` - (Optional) Git branch (defaults to `main`).
  The branch must already exist, if its not specified then default branch for repo is being used.

## Attributes Reference

The following additional attributes are exported:

* `content` - The file content.

* `commit_sha` - The SHA of the commit that modified the file.

* `sha` - The SHA blob of the file.

* `commit_author` - Committer author name.

* `commit_email` - Committer email address.

* `commit_message` - Commit message when file was last updated.
