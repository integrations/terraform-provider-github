---
layout: "github"
page_title: "GitHub: github_repository_autolink_reference"
description: |-
  Creates and manages autolink references for a single repository
---

# github_repository_autolink_reference

This resource allows you to create and manage an autolink reference for a single repository.

When a `github_repository_autolink_reference` resource is created, the provider first checks whether an autolink with the same `key_prefix` already exists in the repository:

* If one exists **with identical settings**, it is imported into Terraform state rather than creating a duplicate (the GitHub API would otherwise return an error).
* If one exists **with different settings** (`target_url_template` or `is_alphanumeric`), the existing autolink is deleted and recreated with the desired configuration.
* If none exists, the autolink is created normally.

## Example Usage

```hcl
resource "github_repository" "repo" {
  name         = "my-repo"
  description  = "GitHub repo managed by Terraform"

  private = false
}

resource "github_repository_autolink_reference" "autolink" {
  repository = github_repository.repo.name

  key_prefix = "TICKET-"

  target_url_template = "https://example.com/TICKET?query=<num>"
}
```

## Argument Reference

The following arguments are supported:

* `repository` - (Required) The repository of the autolink reference.

* `key_prefix` - (Required) This prefix appended by a number will generate a link any time it is found in an issue, pull request, or commit. If an autolink reference with this key prefix already exists, the provider imports it when the settings match, or replaces it when the settings differ.

* `target_url_template` - (Required) The template of the target URL used for the links; must be a valid URL and contain `<num>` for the reference number

* `is_alphanumeric` - (Optional) Whether this autolink reference matches alphanumeric characters. If false, this autolink reference only matches numeric characters. Default is true.

## Attributes Reference

The following additional attributes are exported:

* `etag` - An etag representing the autolink reference object.

## Import

Autolink references can be imported using the `name` of the repository, combined with the `id` or `key prefix` of the autolink reference and a `/` character for separating components, e.g.

### Import by ID

```sh
terraform import github_repository_autolink_reference.auto my-repo/123
```

See the GitHub documentation for how to [list all autolinks of a repository](https://docs.github.com/en/rest/repos/autolinks#list-all-autolinks-of-a-repository) to learn the autolink ids to use with the import command.

### Import by key prefix

```sh
terraform import github_repository_autolink_reference.auto oof/OOF-
```
