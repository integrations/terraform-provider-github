---
page_title: "github_dependabot_organization_public_key Data Source - terraform-provider-github
description: |-
  Get information on a GitHub Dependabot Organization Public Key.
---

# github_dependabot_organization_public_key (Data Source)

Use this data source to retrieve information about a GitHub Dependabot Organization public key. This data source is required to be used with other GitHub secrets interactions. Note that the provider `token` must have admin rights to an organization to retrieve it's Dependabot public key.

## Example Usage

```terraform
data "github_dependabot_organization_public_key" "example" {}
```

## Attributes Reference

- `key_id` - ID of the key that has been retrieved.
- `key` - Actual key retrieved.
