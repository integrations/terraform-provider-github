---
layout: "github"
page_title: "GitHub: github_organization_webhooks"
description: |-
  Get information on all Github webhooks of the organization.
---

# github\_organization\_webhooks

Use this data source to retrieve all webhooks of the organization.

## Example Usage

To retrieve *all* webhooks of the organization:

```hcl
data "github_organization_webhooks" "all" {}
```

## Attributes Reference

* `webhooks` - An Array of GitHub Webhooks.  Each `webhook` block consists of the fields documented below.
___

The `webhook` block consists of:

 * `id` - the ID of the webhook.
 * `type` - the type of the webhook.
 * `name` - the name of the webhook.
 * `url` - the url of the webhook.
 * `active` - `true` if the webhook is active.
