---
layout: "github"
page_title: "GitHub: github_user_external_identity_by_saml"
description: |-
  Look up a GitHub user by their SAML NameID.
---

# github\_user\_external\_identity\_by\_saml

Use this data source to retrieve a GitHub user's login by their SAML NameID
(typically an email address). This is a reverse lookup — given a SAML identity,
it returns the linked GitHub username.

This complements `github_user_external_identity`, which performs the opposite
lookup (GitHub username to SAML/SCIM identity).

## Example Usage

```hcl
data "github_user_external_identity_by_saml" "example" {
  saml_name_id = "user@example.com"
}

resource "github_team_membership" "example" {
  team_id  = github_team.some_team.id
  username = data.github_user_external_identity_by_saml.example.login
}
```

## Argument Reference

* `saml_name_id` - (Required) The SAML NameID to look up. This is typically
  the user's email address as configured in your identity provider.

## Attribute Reference

* `login` - The GitHub username linked to the SAML identity.
* `username` - Same as `login`.
* `saml_identity` - A map of SAML identity attributes:
  * `name_id` - The SAML NameID value.
  * `username` - The SAML username.
  * `given_name` - The user's given name.
  * `family_name` - The user's family name.
* `scim_identity` - A map of SCIM identity attributes:
  * `username` - The SCIM username.
  * `given_name` - The user's given name.
  * `family_name` - The user's family name.
