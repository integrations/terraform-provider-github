---
page_title: "github_enterprise_organization (Resource) - GitHub"
description: |-
  Create and manages a GitHub enterprise organization.
---

# github_enterprise_organization (Resource)

This resource allows you to create and manage a GitHub enterprise organization.

## Example Usage

```hcl
resource "github_enterprise_organization" "org" {
  enterprise_id = data.github_enterprise.enterprise.id
  name          = "some-awesome-org"
  display_name  = "Some Awesome Org"
  description   = "Organization created with terraform"
  billing_email = "jon@winteriscoming.com"
  admin_logins  = [
    "jon-snow"
  ]
}
```

## Argument Reference

- `enterprise_id` - (Required) The ID of the enterprise.
- `name` - (Required) The name of the organization.
- `description` - (Optional) The description of the organization.
- `display_name` - (Optional) The display name of the organization.
- `billing_email` - (Required) The billing email address.
- `admin_logins` - (Required) List of organization owner usernames.

## Attributes Reference

The following additional attributes are exported:

- `id` - The node ID of the organization for use with the v4 API.
- `database_id` - The ID of the organization.
- `automation_installation_id` - The automation App installation ID discovered
  when the provider opts into `enterprise_app_auth`. Unset in legacy mode.

## Opt-In App Lifecycle

With the provider's `enterprise_app_auth` block, this resource creates an
organization with the existing enterprise creation App and bootstraps the
automation App with the enterprise installer. The enterprise ID must match the
configured enterprise slug. Node IDs and numeric organization IDs are checked
against enterprise inventory and the organization APIs before accepting state.
Propagation retries are bounded by the create timeout and stop on permission,
suspension, and identity errors. A failed installation POST is not retried.

Refresh and import use the installer's paginated enterprise inventory plus
organization-scoped automation clients. Billing comes from REST `billing_email`,
not GraphQL `organizationBillingEmail`. Missing/null billing is an error; existing
billing state is not replaced with an empty value. Owners are paginated and only
`ADMIN` memberships are retained. Denied, ambiguous, missing-installation, or
identity-mismatch reads return errors and retain the resource ID, rather than
silently planning a replacement. Organizations renamed outside Terraform require
explicit reconciliation/import; refresh does not guess an identity from a new login.

Updates are not supported in this opt-in path, including changes to owners,
billing, display name, and description. Restore the previous configuration or
perform approved external changes and reconcile configuration before refreshing.
Deletion is blocked before any API request until an approved permission and
deletion contract is available. `lifecycle { prevent_destroy = true }` is
recommended. Neither this mode nor the installation ID output dynamically routes
ordinary organization resources such as repositories or teams.

If creation succeeds but installation/profile bootstrap fails, the error reports
the enterprise/org import ID and retains the returned node/database IDs, including
in the SDK apply response. Terraform can mark a failed create as tainted; do not
blindly retry or replace it. Inspect the organization and App installation,
correct the missing permissions/installation through approved administration,
then untaint the existing object if state is complete or explicitly reconcile
state and import it. Import and refresh will not finish a missing installation.
Deletion remains blocked even for a tainted object. If the creation API response
itself is lost, check enterprise inventory and import any created organization
before retrying, because a trustworthy node ID may not have been returned.

## Import

GitHub Enterprise Organization can be imported using the `slug` of the enterprise, combined with the `orgname` of the organization, separated by a `/` character.

```shell
terraform import github_enterprise_organization.org enterp/some-awesome-org
```

In `enterprise_app_auth` mode the enterprise slug in the import ID must match
the configured slug. The automation App must already be installed, active, and
authorized for organization administration and member reads. Import makes no
installation writes and fails on missing permissions instead of discarding state.
