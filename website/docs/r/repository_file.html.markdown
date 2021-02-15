---
layout: "github"
page_title: "GitHub: github_repository_file"
description: |-
  Creates and manages files within a GitHub repository
---

# github_repository_file

This resource allows you to create and manage files within a
GitHub repository.


## Example Usage

```hcl

resource "github_repository" "foo" {
  name      = "tf-acc-test-%s"
  auto_init = true
}

resource "github_repository_file" "foo" {
  repository          = github_repository.foo.name
  branch              = "main"
  file                = ".gitignore"
  content             = "**/*.tfstate"
  commit_message      = "Managed by Terraform"
  commit_author       = "Terraform User"
  commit_email        = "terraform@example.com"
  overwrite_on_create = true
}

```


## Argument Reference

The following arguments are supported:

* `repository` - (Required) The repository to create the file in.

* `file` - (Required) The path of the file to manage.

* `content` - (Required) The file content.

* `branch` - (Optional) Git branch (defaults to `main`).
  The branch must already exist, it will not be created if it does not already exist.

* `commit_author` - (Optional) Committer author name to use.

* `commit_email` - (Optional) Committer email address to use.

* `commit_message` - (Optional) Commit message when adding or updating the managed file.

* `overwrite_on_create` - (Optional) Enable overwriting existing files

## Attributes Reference

The following additional attributes are exported:

* `commit_sha` - The SHA of the commit that modified the file.

* `sha` - The SHA blob of the file.


## Import

Repository files can be imported using a combination of the `repo` and `file`, e.g.

```
$ terraform import github_repository_file.gitignore example/.gitignore
```

To import a file from a branch other than main, append `:` and the branch name, e.g.

```
$ terraform import github_repository_file.gitignore example/.gitignore:dev
```
