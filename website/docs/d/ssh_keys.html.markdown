---
layout: "github"
page_title: "GitHub: github_ssh_keys"
description: |-
  Get information on GitHub's SSH keys.
---

# github_ssh_keys

Use this data source to retrieve information about GitHub's SSH keys.

## Example Usage

```hcl
data "github_ssh_keys" "test" {}
```

## Attributes Reference

 * `keys` - An array of GitHub's SSH public keys.
