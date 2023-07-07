---
layout: "github"
page_title: "GitHub: github_organization_external_identities"
description: |-
  Get a list of organization members and their SAML linked external identity NameID
---

# github_organization_external_identities

Use this data source to retrieve each organization member's SAML linked external
identity's NameID.

## Example Usage

```hcl
data "github_organization_external_identities" "all" {}
```

## Attributes Reference

- `identities` - An Array of identities returned from GitHub

---

Each element in the `identities` block consists of:

- `login` - The username of the GitHub user
- `samlIdentityNameID` - The external identity NameID attached to the GitHub
  user
