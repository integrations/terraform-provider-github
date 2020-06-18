---
layout: "github"
page_title: "GitHub: github_repository_webhook"
description: |-
  Creates and manages repository webhooks within GitHub organizations or personal accounts
---

# github_repository_webhook

This resource allows you to create and manage webhooks for repositories within your
GitHub organization or personal account.

## Example Usage

```hcl
resource "github_repository" "repo" {
  name         = "foo"
  description  = "Terraform acceptance tests"
  homepage_url = "http://example.com/"

  private = false
}

resource "github_repository_webhook" "foo" {
  repository = "${github_repository.repo.name}"

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

* `repository` - (Required) The repository of the webhook.

* `events` - (Required) A list of events which should trigger the webhook. See a list of [available events](https://developer.github.com/v3/activity/events/types/).

* `configuration` - (Required) key/value pair of configuration for this webhook. Available keys are `url`, `content_type`, `secret` and `insecure_ssl`. `secret` is [the shared secret, see API documentation](https://developer.github.com/v3/repos/hooks/#create-a-hook).

* `active` - (Optional) Indicate of the webhook should receive events. Defaults to `true`.

* `name` - (Optional) The type of the webhook. `web` is the default and the only option.

## Attributes Reference

The following additional attributes are exported:

* `url` - URL of the webhook

## Import

Repository webhooks can be imported using the `name` of the repository, combined with the `id` of the webhook, separated by a `/` character.
The `id` of the webhook can be found in the URL of the webhook. For example: `"https://github.com/foo-org/foo-repo/settings/hooks/14711452"`.

Importing uses the name of the repository, as well as the ID of the webhook, e.g.

```
$ terraform import github_repository_webhook.terraform terraform/11235813
```

If secret is populated in the webhook's configuration, the value will be imported as "********".