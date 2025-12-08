---
layout: "github"
page_title: "GitHub: github_repository_webhook"
description: |-
  Creates and manages repository webhooks within GitHub organizations or personal accounts
---

# github_repository_webhook

This resource allows you to create and manage webhooks for repositories within your
GitHub organization or personal account.

~> **Note on Archived Repositories**: When a repository is archived, GitHub makes it read-only, preventing webhook modifications. If you attempt to destroy resources associated with archived repositories, the provider will gracefully handle the operation by logging an informational message and removing the resource from Terraform state without attempting to modify the archived repository.

## Example Usage

```hcl
resource "github_repository" "repo" {
  name         = "foo"
  description  = "Terraform acceptance tests"
  homepage_url = "http://example.com/"

  visibility   = "public"
}

resource "github_repository_webhook" "foo" {
  repository = github_repository.repo.name

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

* `configuration` - (Required) Configuration block for the webhook. [Detailed below.](#configuration)

* `active` - (Optional) Indicate if the webhook should receive events. Defaults to `true`.

### configuration

* `url` - (Required) The URL of the webhook.

* `content_type` - (Required) The content type for the payload. Valid values are either `form` or `json`.

* `secret` - (Optional) The shared secret for the webhook. [See API documentation](https://developer.github.com/v3/repos/hooks/#create-a-hook).

* `insecure_ssl` - (Optional) Insecure SSL boolean toggle. Defaults to `false`.

## Attributes Reference

The following additional attributes are exported:

* `url` - URL of the webhook.  This is a sensitive attribute because it may include basic auth credentials.

## Import

Repository webhooks can be imported using the `name` of the repository, combined with the `id` of the webhook, separated by a `/` character.
The `id` of the webhook can be found in the URL of the webhook. For example: `"https://github.com/foo-org/foo-repo/settings/hooks/14711452"`.

Importing uses the name of the repository, as well as the ID of the webhook, e.g.

```
$ terraform import github_repository_webhook.terraform terraform/11235813
```

If secret is populated in the webhook's configuration, the value will be imported as "********".
