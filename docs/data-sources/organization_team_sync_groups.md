---
page_title: "github_organization_team_sync_groups Data Source - terraform-provider-github
description: |-
  Get the external identity provider (IdP) groups for an organization.
---

# github_organization_team_sync_groups (Data Source)

Use this data source to retrieve the identity provider (IdP) groups for an organization.

## Example Usage

```terraform
data "github_organization_team_sync_groups" "test" {}
```

## Attributes Reference

- `groups` - An Array of GitHub Identity Provider Groups. Each `group` block consists of the fields documented below.

---

The `group` block consists of:

- `group_id` - The ID of the IdP group.

- `group_name` - The name of the IdP group.

- `group_description` - The description of the IdP group.
