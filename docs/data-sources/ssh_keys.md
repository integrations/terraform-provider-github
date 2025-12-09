---
page_title: "github_ssh_keys Data Source - terraform-provider-github
description: |-
  Get information on GitHub's SSH keys.
---

# github_ssh_keys (Data Source)

Use this data source to retrieve information about GitHub's SSH keys.

## Example Usage

```terraform
data "github_ssh_keys" "test" {}
```

## Attributes Reference

- `keys` - An array of GitHub's SSH public keys.
