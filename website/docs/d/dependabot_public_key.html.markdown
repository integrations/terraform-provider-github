---
layout: "github"
page_title: "GitHub: github_dependabot_public_key"
description: |-
  Get information on a GitHub Dependabot Public Key.
---

# github_dependabot_public_key

Use this data source to retrieve information about a GitHub Dependabot public key. This data source is required to be used with other GitHub secrets interactions.
Note that the provider `token` must have admin rights to a repository to retrieve it's Dependabot public key.

## Example Usage

```hcl
data "github_dependabot_public_key" "example" {
  repository = "example_repo"
}
```

## Argument Reference

* `repository` - (Required) Name of the repository to get public key from.

## Attributes Reference

* `key_id` - ID of the key that has been retrieved.
* `key`    - Actual key retrieved.
