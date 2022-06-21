---
layout: "github"
page_title: "GitHub: github_ref"
description: |-
  Get information about a repository ref.
---

# github_ref

Use this data source to retrieve information about a repository ref.

## Example Usage

```hcl
data "github_ref" "development" {
  repository = "example"
  ref     = "heads/development"
}
```

## Argument Reference

The following arguments are supported:

* `repository` - (Required) The GitHub repository name.

* `ref` - (Required) The repository ref to look up. Must be formatted `heads/<ref>` for branches, and `tags/<ref>` for tags.

## Attribute Reference

The following additional attributes are exported:

* `etag` - An etag representing the ref.

* `sha` - A string storing the reference's `HEAD` commit's SHA1.
