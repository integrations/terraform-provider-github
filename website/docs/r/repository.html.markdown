---
layout: "github"
page_title: "GitHub: github_repository"
description: |-
  Creates and manages repositories within GitHub organizations or personal accounts
---

# github_repository

This resource allows you to create and manage repositories within your
GitHub organization or personal account.

## Example Usage

```hcl
resource "github_repository" "example" {
  name        = "example"
  description = "My awesome codebase"

  visibility = "public"

  template {
    owner      = "github"
    repository = "terraform-module-template"
  }
}
```

## Example Usage with GitHub Pages Enabled

```hcl
resource "github_repository" "example" {
  name        = "example"
  description = "My awesome web page"

  private = false

  pages {
    source {
      branch = "master"
      path   = "/docs"
    }
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

* `visibility` - (Optional) Can be `public` or `private`. If your organization is associated with an enterprise account using GitHub Enterprise Cloud or GitHub Enterprise Server 2.20+, visibility can also be `internal`. The `visibility` parameter overrides the `private` parameter.

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

* `default_branch` - (Optional) (Deprecated: Use `github_branch_default` resource instead) The name of the default branch of the repository. **NOTE:** This can only be set after a repository has already been created,
and after a correct reference has been created for the target branch inside the repository. This means a user will have to omit this parameter from the
initial repository creation and create the target branch inside of the repository prior to setting this attribute.

* `archived` - (Optional) Specifies if the repository should be archived. Defaults to `false`. **NOTE** Currently, the API does not support unarchiving.

* `archive_on_destroy` - (Optional) Set to `true` to archive the repository instead of deleting on destroy.

* `pages` - (Optional) The repository's GitHub Pages configuration. See [GitHub Pages Configuration](#github-pages-configuration) below for details.

* `topics` - (Optional) The list of topics of the repository.

* `template` - (Optional) Use a template repository to create this resource. See [Template Repositories](#template-repositories) below for details.

* `vulnerability_alerts` (Optional) - Set to `true` to enable security alerts for vulnerable dependencies. Enabling requires alerts to be enabled on the owner level. (Note for importing: GitHub enables the alerts on public repos but disables them on private repos by default.) See [GitHub Documentation](https://help.github.com/en/github/managing-security-vulnerabilities/about-security-alerts-for-vulnerable-dependencies) for details. Note that vulnerability alerts have not been successfully tested on any GitHub Enterprise instance and may be unavailable in those settings.

### GitHub Pages Configuration

The `pages` block supports the following:

* `source` - (Required) The source branch and directory for the rendered Pages site. See [GitHub Pages Source](#github-pages-source) below for details.

* `cname` - (Optional) The custom domain for the repository. This can only be set after the repository has been created.

#### GitHub Pages Source ####

The `source` block supports the following:

* `branch` - (Required) The repository branch used to publish the site's source files. (i.e. `main` or `gh-pages`.

* `path` - (Optional) The repository directory from which the site publishes (Default: `/`).

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

* `node_id` - GraphQL global node id for use with v4 API

* `repo_id` - GitHub ID for the repository

* `pages` - The block consisting of the repository's GitHub Pages configuration with the following additional attributes:
 * `custom_404` - Whether the rendered GitHub Pages site has a custom 404 page.
 * `html_url` - The absolute URL (including scheme) of the rendered GitHub Pages site e.g. `https://username.github.io`.
 * `status` - The GitHub Pages site's build status e.g. `building` or `built`.

* `branches` - The list of this repository's branches. Each element of `branches` has the following attributes:
 * `name` - Name of the branch.
 * `protected` - Whether the branch is protected.


## Import

Repositories can be imported using the `name`, e.g.

```
$ terraform import github_repository.terraform terraform
```
