---
layout: "github"
page_title: "GitHub: github_organization_webhook"
description: |-
  Creates and manages webhooks for GitHub organizations
---

# github_organization_webhook

This resource allows you to create and manage webhooks for GitHub organization.

## Example Usage

```hcl
resource "github_organization_webhook" "foo" {
  name = "web"

  configuration {
    url          = "https://google.de/"
    content_type = "form"
    insecure_ssl = false
  }

  active = false

  events = ["issues"]
}
```

## Argument Reference

The following arguments are supported:

* `events` - (Required) A list of events which should trigger the webhook. See a list of [available events](https://developer.github.com/v3/activity/events/types/)

* `configuration` - (Required) key/value pair of configuration for this webhook. Available keys are `url`, `content_type`, `secret` and `insecure_ssl`.

* `active` - (Optional) Indicate of the webhook should receive events. Defaults to `true`.

* `name` - (Optional) The type of the webhook. `web` is the default and the only option.

## Attributes Reference

The following additional attributes are exported:

* `url` - URL of the webhook

## Import

Organization webhooks can be imported using the `id` of the webhook.
The `id` of the webhook can be found in the URL of the webhook. For example, `"https://github.com/organizations/foo-org/settings/hooks/123456789"`.

```
$ terraform import github_organization_webhook.terraform 123456789
```

If secret is populated in the webhook's configuration, the value will be imported as "********".
