---
layout: "github"
page_title: "GitHub: github_repository_autolink_reference"
description: |-
  Creates and manages autolink references for a single repository
---

# github_repository_autolink_reference

This resource allows you to create and manage an autolink reference for a single repository.

## Example Usage

```hcl
resource "github_repository" "repo" {
  name         = "oof"
  description  = "GitHub repo managed by Terraform"

  private = false
}

resource "github_repository_autolink_reference" "auto" {
  repository = github_repository.repo.name

  key_prefix = "TICKET-"

  target_url_template = "https://hello.there/TICKET?query=<num>"
}
```

## Argument Reference

The following arguments are supported:

* `repository` - (Required) The repository of the autolink reference.

* `key_prefix` - (Required) This prefix appended by a number will generate a link any time it is found in an issue, pull request, or commit.

* `target_url_template` - (Required) The template of the target URL used for the links; must be a valid URL and contain `<num>` for the reference number

## Attributes Reference

The following additional attributes are exported:

* `etag` - An etag representing the autolink reference object.

## Import

Autolink references can be imported using the `name` of the repository, combined with the `id` of the autolink reference and a `/` character for separating components, e.g.

```sh
terraform import github_repository_autolink_reference.auto oof/123
```
