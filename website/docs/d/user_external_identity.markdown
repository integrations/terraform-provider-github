---
layout: "github"
page_title: "GitHub: github_user_external_identity"
description: |-
  Get a specific organization member's SAML/SCIM linked external identity 
---

# github_user_external_identity

Use this data source to retrieve a specific organization member's SAML or SCIM user
attributes.

## Example Usage

```hcl
data "github_user_external_identity" "example_user" {
  username = "example-user"
}
```

## Argument Reference

The following arguments are supported:

- `username` - (Required) The username of the member to fetch external identity for.

## Attributes Reference

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
- `scim_family_name` - The member's SCIM Family Name
- `scim_given_name` - The member's SCIM Given Name
