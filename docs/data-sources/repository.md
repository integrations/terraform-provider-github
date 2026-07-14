---
page_title: "github_repository (Data Source) - GitHub"
subcategory: ""
description: |-
  Use this data source to retrieve information about a GitHub repository.
---

# github_repository (Data Source)

Use this data source to retrieve information about a GitHub repository.

## Example Usage

```terraform
data "github_repository" "example" {
  full_name = "hashicorp/terraform"
}
```

<!--
## Schema

### Optional

- `description` (String) A description of the repository.
- `full_name` (String) The full name of the repository, in the format `owner/repo_name`.
- `homepage_url` (String) URL of a page describing the project.
- `name` (String) The name of the repository.

### Read-Only

- `allow_auto_merge` (Boolean) Whether the repository allows auto-merging pull requests.
- `allow_forking` (Boolean) Whether the repository allows private forking; this is only relevant if the repository is owned by an organization and is private or internal.
- `allow_merge_commit` (Boolean) Whether the repository allows merge commits.
- `allow_rebase_merge` (Boolean) Whether the repository allows rebase merges.
- `allow_squash_merge` (Boolean) Whether the repository allows squash merges.
- `allow_update_branch` (Boolean) Whether the repository allows users with push access to update the base branch of pull requests.
- `archived` (Boolean) Whether the repository is archived.
- `default_branch` (String) The default branch of the repository.
- `delete_branch_on_merge` (Boolean) Whether the repository has the option to automatically delete head branches when pull requests are merged enabled.
- `fork` (Boolean) Whether the repository is a fork of another repository.
- `git_clone_url` (String) URL that can be provided to `git clone` to clone the repository anonymously via the git protocol.
- `has_discussions` (Boolean) Whether the repository has discussions enabled.
- `has_downloads` (Boolean, Deprecated) Whether the repository has Downloads feature enabled. This attribute is no longer in use, but it hasn't been removed yet. It will be removed in a future version.
- `has_issues` (Boolean) Whether the repository has issues enabled.
- `has_projects` (Boolean) Whether the repository has projects enabled.
- `has_wiki` (Boolean) Whether the repository has a wiki enabled.
- `html_url` (String) URL to the repository on the web.
- `http_clone_url` (String) URL that can be provided to `git clone` to clone the repository via HTTPS.
- `id` (String) The ID of this resource.
- `is_template` (Boolean) Whether the repository is a template that can be used to generate new repositories.
- `merge_commit_message` (String) The default value for a merge commit message.
- `merge_commit_title` (String) The default value for a merge commit title.
- `node_id` (String) GraphQL global node id for use with v4 API.
- `pages` (List of Object, Deprecated) Use the `github_repository_pages` data source instead. This field will be removed in a future version. (see [below for nested schema](#nestedatt--pages))
- `primary_language` (String) The primary language used in the repository.
- `private` (Boolean) Whether the repository is private.
- `repo_id` (Number) GitHub ID for the repository.
- `repository_license` (List of Object) An Array of GitHub repository licenses. Each `repository_license` block consists of the fields documented below. (see [below for nested schema](#nestedatt--repository_license))
- `squash_merge_commit_message` (String) The default value for a squash merge commit message.
- `squash_merge_commit_title` (String) The default value for a squash merge commit title.
- `ssh_clone_url` (String) URL that can be provided to `git clone` to clone the repository via SSH.
- `svn_url` (String) URL that can be provided to `svn checkout` to check out the repository via GitHub's Subversion protocol emulation.
- `template` (List of Object) The repository source template configuration. (see [below for nested schema](#nestedatt--template))
- `topics` (List of String) The list of topics of the repository.
- `visibility` (String) Whether the repository is public, private or internal.
- `web_commit_signoff_required` (Boolean) Require contributors to sign off on web-based commits.

<a id="nestedatt--pages"></a>
### Nested Schema for `pages`

Read-Only:

- `build_type` (String)
- `cname` (String)
- `custom_404` (Boolean)
- `html_url` (String)
- `source` (List of Object) (see [below for nested schema](#nestedobjatt--pages--source))
- `status` (String)
- `url` (String)

<a id="nestedobjatt--pages--source"></a>
### Nested Schema for `pages.source`

Read-Only:

- `branch` (String)
- `path` (String)



<a id="nestedatt--repository_license"></a>
### Nested Schema for `repository_license`

Read-Only:

- `content` (String)
- `download_url` (String)
- `encoding` (String)
- `git_url` (String)
- `html_url` (String)
- `license` (List of Object) (see [below for nested schema](#nestedobjatt--repository_license--license))
- `name` (String)
- `path` (String)
- `sha` (String)
- `size` (Number)
- `type` (String)
- `url` (String)

<a id="nestedobjatt--repository_license--license"></a>
### Nested Schema for `repository_license.license`

Read-Only:

- `body` (String)
- `conditions` (Set of String)
- `description` (String)
- `featured` (Boolean)
- `html_url` (String)
- `implementation` (String)
- `key` (String)
- `limitations` (Set of String)
- `name` (String)
- `permissions` (Set of String)
- `spdx_id` (String)
- `url` (String)



<a id="nestedatt--template"></a>
### Nested Schema for `template`

Read-Only:

- `owner` (String)
- `repository` (String)
-->

## Schema

### Optional

- `description` (String) A description of the repository.
- `full_name` (String) The full name of the repository, in the format `owner/repo_name`.
- `homepage_url` (String) URL of a page describing the project.
- `name` (String) The name of the repository.

### Read-Only

- `allow_auto_merge` (Boolean) Whether the repository allows auto-merging pull requests.
- `allow_forking` (Boolean) Whether the repository allows private forking; this is only relevant if the repository is owned by an organization and is private or internal.
- `allow_merge_commit` (Boolean) Whether the repository allows merge commits.
- `allow_rebase_merge` (Boolean) Whether the repository allows rebase merges.
- `allow_squash_merge` (Boolean) Whether the repository allows squash merges.
- `allow_update_branch` (Boolean) Whether the repository allows users with push access to update the base branch of pull requests.
- `archived` (Boolean) Whether the repository is archived.
- `default_branch` (String) The default branch of the repository.
- `delete_branch_on_merge` (Boolean) Whether the repository has the option to automatically delete head branches when pull requests are merged enabled.
- `fork` (Boolean) Whether the repository is a fork of another repository.
- `git_clone_url` (String) URL that can be provided to `git clone` to clone the repository anonymously via the git protocol.
- `has_discussions` (Boolean) Whether the repository has discussions enabled.
- `has_downloads` (Boolean, Deprecated) Whether the repository has Downloads feature enabled. This attribute is no longer in use, but it hasn't been removed yet. It will be removed in a future version.
- `has_issues` (Boolean) Whether the repository has issues enabled.
- `has_projects` (Boolean) Whether the repository has projects enabled.
- `has_wiki` (Boolean) Whether the repository has a wiki enabled.
- `html_url` (String) URL to the repository on the web.
- `http_clone_url` (String) URL that can be provided to `git clone` to clone the repository via HTTPS.
- `id` (String) The ID of this resource.
- `is_template` (Boolean) Whether the repository is a template that can be used to generate new repositories.
- `merge_commit_message` (String) The default value for a merge commit message.
- `merge_commit_title` (String) The default value for a merge commit title.
- `node_id` (String) GraphQL global node id for use with v4 API.
- `pages` (List of Object, Deprecated) Use the `github_repository_pages` data source instead. This field will be removed in a future version. (see [below for nested schema](#nestedatt--pages))
- `primary_language` (String) The primary language used in the repository.
- `private` (Boolean) Whether the repository is private.
- `repo_id` (Number) GitHub ID for the repository.
- `repository_license` (List of Object) An Array of GitHub repository licenses. Each `repository_license` block consists of the fields documented below. (see [below for nested schema](#nestedatt--repository_license))
- `squash_merge_commit_message` (String) The default value for a squash merge commit message.
- `squash_merge_commit_title` (String) The default value for a squash merge commit title.
- `ssh_clone_url` (String) URL that can be provided to `git clone` to clone the repository via SSH.
- `svn_url` (String) URL that can be provided to `svn checkout` to check out the repository via GitHub's Subversion protocol emulation.
- `template` (List of Object) The repository source template configuration. (see [below for nested schema](#nestedatt--template))
- `topics` (List of String) The list of topics of the repository.
- `visibility` (String) Whether the repository is public, private or internal.
- `web_commit_signoff_required` (Boolean) Require contributors to sign off on web-based commits.

<a id="nestedatt--pages"></a>
### Nested Schema for `pages`

Read-Only:

- `build_type` (String)
- `cname` (String)
- `custom_404` (Boolean)
- `html_url` (String)
- `source` (List of Object) (see [below for nested schema](#nestedobjatt--pages--source))
- `status` (String)
- `url` (String)

<a id="nestedobjatt--pages--source"></a>
### Nested Schema for `pages.source`

Read-Only:

- `branch` (String)
- `path` (String)



<a id="nestedatt--repository_license"></a>
### Nested Schema for `repository_license`

Read-Only:

- `content` (String) Content of the license file, encoded by encoding scheme mentioned below.
- `download_url` (String) The URL to download the raw content of the license file.
- `encoding` (String) The encoding used for the content (e.g., \"base64\").
- `git_url` (String) The URL to access information about the license file as a Git blob.
- `html_url` (String) The URL to view the license file on GitHub.
- `license` (List of Object) The license information for the license file in the repository. (see [below for nested schema](#nestedobjatt--repository_license--license))
- `name` (String) The name of the license file in the repository.
- `path` (String) The path of the license file in the repository.
- `sha` (String) The SHA hash of the license file.
- `size` (Number) The size of the license file in bytes.
- `type` (String) The type of the license file (e.g., \"file\").
- `url` (String) The URL to access information about the license file on GitHub.

<a id="nestedobjatt--repository_license--license"></a>
### Nested Schema for `repository_license.license`

Read-Only:

- `body` (String) The text of the license.
- `conditions` (Set of String) Conditions associated with the license.
- `description` (String) A description of the license.
- `featured` (Boolean) Indicates if the license is featured.
- `html_url` (String) The URL to view the license details on GitHub.
- `implementation` (String) Details about the implementation of the license.
- `key` (String) A key representing the license type (e.g., "apache-2.0").
- `limitations` (Set of String) "Limitations associated with the license.
- `name` (String) The name of the license (e.g., "Apache License 2.0").
- `permissions` (Set of String) Permissions associated with the license.
- `spdx_id` (String) The SPDX identifier for the license (e.g., \"Apache-2.0\").
- `url` (String) The URL to access information about the license on GitHub.



<a id="nestedatt--template"></a>
### Nested Schema for `template`

Read-Only:

- `owner` (String) Owner of the template repository.
- `repository` (String) Name of the template repository.
