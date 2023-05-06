---
layout: "github"
page_title: "GitHub: github_branch_protection_rules"
description: |-
  Get information about a repository branch protection rules.
---

# github\_branch\_protection\_rules

Use this data source to retrieve a list of repository branch protection rules.

## Example Usage

```hcl
data "github_branch_protection_rules" "example" {
  repository = "example"
}
```

## Argument Reference

The following arguments are supported:

* `repository` - (Required) The GitHub repository name.

## Attribute Reference

* `rules` - Collection of Branch Protection Rules. Each of the results conforms to the following scheme:

    * `pattern` - Identifies the protection rule pattern.
