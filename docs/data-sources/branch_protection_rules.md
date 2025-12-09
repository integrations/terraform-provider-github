---
page_title: "github_branch_protection_rules Data Source - terraform-provider-github
description: |-
  Get information about a repository branch protection rules.
---

# github_branch_protection_rules (Data Source)

Use this data source to retrieve a list of repository branch protection rules.

## Example Usage

```terraform
data "github_branch_protection_rules" "example" {
  repository = "example"
}
```

## Argument Reference

The following arguments are supported:

- `repository` - (Required) The GitHub repository name.

## Attribute Reference

- `rules` - Collection of Branch Protection Rules. Each of the results conforms to the following scheme:

    - `pattern` - Identifies the protection rule pattern.
