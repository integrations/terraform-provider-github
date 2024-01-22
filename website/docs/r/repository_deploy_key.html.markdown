---
layout: "github"
page_title: "GitHub: github_repository_deploy_key"
description: |-
  Provides a GitHub repository deploy key resource.
---

# github_repository_deploy_key

Provides a GitHub repository deploy key resource.

A deploy key is an SSH key that is stored on your server and grants
access to a single GitHub repository. This key is attached directly to the repository instead of to a personal user
account.

This resource allows you to add/remove repository deploy keys.

Further documentation on GitHub repository deploy keys:
- [About deploy keys](https://developer.github.com/guides/managing-deploy-keys/#deploy-keys)

## Example Usage

```hcl
# Generate an ssh key using provider "hashicorp/tls"
resource "tls_private_key" "example_repository_deploy_key" {
  algorithm = "ED25519"
}

# Add the ssh key as a deploy key
resource "github_repository_deploy_key" "example_repository_deploy_key" {
  title      = "Repository test key"
  repository = "test-repo"
  key        = tls_private_key.example_repository_deploy_key.public_key_openssh
  read_only  = true
}
```

## Argument Reference

The following arguments are supported:

* `key` - (Required) A SSH key.
* `read_only` - (Required) A boolean qualifying the key to be either read only or read/write.
* `repository` - (Required) Name of the GitHub repository.
* `title` - (Required) A title.

Changing any of the fields forces re-creating the resource.

## Import

Repository deploy keys can be imported using a colon-separated pair of repository name
and GitHub's key id. The latter can be obtained by GitHub's SDKs and API.

```
$ terraform import github_repository_deploy_key.foo test-repo:23824728
```
