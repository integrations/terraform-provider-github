---
layout: "github"
page_title: "GitHub: github_organization_app_installations"
description: |-
  Get information on all GitHub App installations of the organization.
---

# github\_organization\_app_installations

Use this data source to retrieve all GitHub App installations of the organization.

## Example Usage

To retrieve *all* GitHub App installations of the organization:

```hcl
data "github_organization_app_installations" "all" {}
```

## Attributes Reference

* `installations` - List of GitHub App installations in the organization. Each `installation` block consists of the fields documented below.

---

The `installation` block consists of:

 * `id` - The ID of the GitHub App installation.
 * `slug` - The URL-friendly name of the GitHub App.
 * `app_id` - The ID of the GitHub App.
 * `repository_selection` - Whether the installation has access to all repositories or only selected ones. Possible values are `all` or `selected`.
 * `permissions` - A map of the permissions granted to the GitHub App installation.
 * `events` - The list of events the GitHub App installation subscribes to.
 * `html_url` - The URL to the GitHub App installation's settings page.
 * `client_id` - The OAuth client ID of the GitHub App.
 * `target_id` - The ID of the account the GitHub App is installed on.
 * `target_type` - The type of account the GitHub App is installed on. Possible values are `Organization` or `User`.
 * `suspended` - Whether the GitHub App installation is currently suspended.
