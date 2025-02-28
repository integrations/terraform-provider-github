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

* `installations` - An Array of GitHub App installations.  Each `installation` block consists of the fields documented below.
___

The `installation` block consists of:

 * `id` - The GitHub app installation id.
 * `slug` - The URL-friendly name of your GitHub App.
 * `app_id` - This is the ID of the GitHub App.
