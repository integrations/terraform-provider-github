---
layout: "github"
page_title: "GitHub: github_actions_organization_public_key"
description: |-
  Get information on a GitHub Actions Organization Public Key.
---

# github_actions_organization_public_key

Use this data source to retrieve information about a GitHub Actions Organization public key. This data source is required to be used with other GitHub secrets interactions.
Note that the provider `token` must have admin rights to an organization to retrieve it's action public key.

## Example Usage

```hcl
data "github_actions_organization_public_key" "example" {}
```

## Attributes Reference

 * `key_id` - ID of the key that has been retrieved.
 * `key`    - Actual key retrieved.
