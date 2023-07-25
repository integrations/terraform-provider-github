---
layout: "github"
page_title: "GitHub: actions_repository_oidc_subject_claim_customization_template"
description: |-
  Get a GitHub Actions repository's OpenID Connect customization template
---

# actions_repository_oidc_subject_claim_customization_template

Use this data source to retrieve the OpenID Connect subject claim customization template for a repository

## Example Usage

```hcl
data "github_actions_repository_oidc_subject_claim_customization_template" "example" {
  name = "example_repository"
}
```

## Argument Reference

* `name` - (Required) Name of the repository to get the OpenID Connect subject claim customization template for.

## Attributes Reference

 * `use_default`        - Whether the repository uses the default template.
 * `include_claim_keys` - The list of OpenID Connect claim keys.
