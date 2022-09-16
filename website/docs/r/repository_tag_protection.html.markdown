---
layout: "github"
page_title: "GitHub: repository_tag_protection"
description: |-
  Creates and manages repository tag protection within GitHub organizations or personal accounts
---

# github_repository_tag_protection

This resource allows you to create and manage a repository tag protection for repositories within your GitHub organization or personal account.

## Example Usage

```hcl
resource "github_repository_tag_protection" "example" {
    repository      = "example-repository"
    pattern         = "v*"
}
```

## Argument Reference

* `repository` - (Required) Name of the repository to add the tag protection to.

* `pattern` - (Required) The pattern of the tag to protect.

## Attributes Reference

The following additional attributes are exported:

* `id` - The ID of the tag protection.

## Import

Repository tag protections can be imported using the `name` of the repository, combined with the `id` of the tag protection, separated by a `/` character.
The `id` of the tag protection can be found using the [GitHub API](https://docs.github.com/en/rest/repos/tags#list-tag-protection-states-for-a-repository).

Importing uses the name of the repository, as well as the ID of the webhook, e.g.

```
$ terraform import github_repository_tag_protection.terraform my-repo/31077
```
