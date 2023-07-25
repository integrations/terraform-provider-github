---
layout: "github"
page_title: "GitHub: github_codespaces_organization_public_key"
description: |-
  Get information on a GitHub Codespaces Organization Public Key.
---

# github_codespaces_organization_public_key

Use this data source to retrieve information about a GitHub Codespaces Organization public key. This data source is required to be used with other GitHub secrets interactions.
Note that the provider `token` must have admin rights to an organization to retrieve it's Codespaces public key.

## Example Usage

```hcl
data "github_codespaces_organization_public_key" "example" {}
```

## Attributes Reference

* `key_id` - ID of the key that has been retrieved.
* `key`    - Actual key retrieved.
