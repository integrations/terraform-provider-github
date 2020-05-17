---
layout: "github"
page_title: "GitHub: github_repository"
description: |-
  Creates and manages repositories within GitHub organizations
---

# github_repository

This resource allows you to create and manage repositories within your
GitHub organization.

This resource cannot currently be used to manage *personal* repositories,
outside of organizations.

## Example Usage

```hcl
resource "github_repository" "example" {
  name        = "example"
  description = "My awesome codebase"

  private = true

  template {
    owner = "github"
    repository = "terraform-module-template"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the repository.

* `description` - (Optional) A description of the repository.

* `homepage_url` - (Optional) URL of a page describing the project.

* `private` - (Optional) Set to `true` to create a private repository.
  Repositories are created as public (e.g. open source) by default.

* `has_issues` - (Optional) Set to `true` to enable the GitHub Issues features
  on the repository.

* `has_projects` - (Optional) Set to `true` to enable the GitHub Projects features on the repository. Per the GitHub [documentation](https://developer.github.com/v3/repos/#create) when in an organization that has disabled repository projects it will default to `false` and will otherwise default to `true`. If you specify `true` when it has been disabled it will return an error.

* `has_wiki` - (Optional) Set to `true` to enable the GitHub Wiki features on
  the repository.
  
* `is_template` - (Optional) Set to `true` to tell GitHub that this is a template repository.

* `allow_merge_commit` - (Optional) Set to `false` to disable merge commits on the repository.

* `allow_squash_merge` - (Optional) Set to `false` to disable squash merges on the repository.

* `allow_rebase_merge` - (Optional) Set to `false` to disable rebase merges on the repository.

* `delete_branch_on_merge` - (Optional) Automatically delete head branch after a pull request is merged. Defaults to `false`.

* `has_downloads` - (Optional) Set to `true` to enable the (deprecated) downloads features on the repository.

* `auto_init` - (Optional) Set to `true` to produce an initial commit in the repository.

* `gitignore_template` - (Optional) Use the [name of the template](https://github.com/github/gitignore) without the extension. For example, "Haskell".

* `license_template` - (Optional) Use the [name of the template](https://github.com/github/choosealicense.com/tree/gh-pages/_licenses) without the extension. For example, "mit" or "mpl-2.0".

* `default_branch` - (Optional) The name of the default branch of the repository. **NOTE:** This can only be set after a repository has already been created,
and after a correct reference has been created for the target branch inside the repository. This means a user will have to omit this parameter from the
initial repository creation and create the target branch inside of the repository prior to setting this attribute.

* `archived` - (Optional) Specifies if the repository should be archived. Defaults to `false`. **NOTE** Currently, the API does not support unarchiving.

* `topics` - (Optional) The list of topics of the repository.

* `template` - (Optional) Use a template repository to create this resource. See [Template Repositories](#template-repositories) below for details.

* `fork_from_repository` - (Optional) The repository to fork from in the format `OWNER/REPOSITORY` e.g. "terraform-providers/terraform-provider-github"

* `fork_into_organization` - (Optional) The Github organization to set as the owner if forked from another repository. By default, the repository will live within the authenticated user's account.  

### Template Repositories

`template` supports the following arguments:

* `owner`: The GitHub organization or user the template repository is owned by.
* `repository`: The name of the template repository.

## Attributes Reference

The following additional attributes are exported:

* `full_name` - A string of the form "orgname/reponame".

* `html_url` - URL to the repository on the web.

* `ssh_clone_url` - URL that can be provided to `git clone` to clone the repository via SSH.

* `http_clone_url` - URL that can be provided to `git clone` to clone the repository via HTTPS.

* `git_clone_url` - URL that can be provided to `git clone` to clone the repository anonymously via the git protocol.

* `svn_url` - URL that can be provided to `svn checkout` to check out the repository via GitHub's Subversion protocol emulation.


## Import

Repositories can be imported using the `name`, e.g.

```
$ terraform import github_repository.terraform terraform
```
