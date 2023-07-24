---
layout: "github"
page_title: "GitHub: github_organization_external_identities"
description: |-
  Get a list of organization members and their SAML linked external identity NameID
---

# github_organization_external_identities

Use this data source to retrieve each organization member's SAML or SCIM user
attributes.

## Example Usage

```hcl
data "github_organization_external_identities" "all" {}
```

## Attributes Reference

- `identities` - An Array of identities returned from GitHub

---

Each element in the `identities` block consists of:

- `login` - The username of the GitHub user
- `saml_identity` - An Object containing the user's SAML data. This object will
  be empty if the user is not managed by SAML.
- `scim_identity` - An Object contining the user's SCIM data. This object will
  be empty if the user is not managed by SCIM.

---

If a user is managed by SAML, the `saml_identity` object will contain:

- `name_id` - The member's SAML NameID
- `username` - The member's SAML Username
- `family_name` - The member's SAML Family Name
- `given_name` - The member's SAML Given Name

---

If a user is managed by SCIM, the `scim_identity` object will contain:

- `scim_username` - The member's SCIM Username. (will be empty string if user is
  not managed by SCIM)
- `scim_groups` - The member's SCIM Groups
- `scim_family_name` - The member's SCIM Family Name
- `scim_given_name` - The member's SCIM Given Name
