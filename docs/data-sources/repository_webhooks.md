---
page_title: "github_repository_webhooks Data Source - terraform-provider-github
description: |-
  Get information on all GitHub webhooks of the organization.
---

# github_repository_webhooks (Data Source)

Use this data source to retrieve webhooks for a given repository.

## Example Usage

To retrieve webhooks of a repository:

```terraform
data "github_repository_webhooks" "repo" {
  repository = "foo"
}
```

## Attributes Reference

- `webhooks` - An Array of GitHub Webhooks. Each `webhook` block consists of the fields documented below.

---

The `webhook` block consists of:

- `id` - the ID of the webhook.
- `type` - the type of the webhook.
- `name` - the name of the webhook.
- `url` - the url of the webhook.
- `active` - `true` if the webhook is active.
