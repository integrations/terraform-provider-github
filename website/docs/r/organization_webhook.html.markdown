---
layout: "github"
page_title: "GitHub: github_organization_webhook"
sidebar_current: "docs-github-resource-organization-webhook"
description: |-
  Creates and manages webhooks for GitHub organizations
---

# github_organization_webhook

This resource allows you to create and manage webhooks for GitHub organization.

## Example Usage

```hcl
resource "github_organization_webhook" "foo" {
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

## Attributes Reference

The following additional attributes are exported:

* `url` - URL of the webhook
