---
layout: "github"
page_title: "GitHub: repository_autolink_references"
description: |-
  Get autolink references for a Github repository.
---

# github_repository_autolink_references

Use this data source to retrieve autolink references for a repository.

## Example Usage

```hcl
data "github_repository_autolink_references" "example" {
    repository = "example-repository"
}
```

## Argument Reference

* `repository` - (Required) Name of the repository to retrieve the autolink references from.

## Attributes Reference

* `autolink_references` - The list of this repository's autolink references. Each element of `autolink_references` has the following attributes:
    * `key_prefix` - Key prefix.
    * `target_url_template` - Target url template.
    * `is_alphanumeric` - True if alphanumeric.
