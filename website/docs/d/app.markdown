---
layout: "github"
page_title: "GitHub: github_app"
description: |-
  Get information about an app.
---

# github\_app

Use this data source to retrieve information about an app.

## Example Usage

```hcl
data "github_app" "foobar" {
  slug = "foobar"
}
```

## Argument Reference

The following arguments are supported:

* `slug` - (Required) The URL-friendly name of your GitHub App.


## Attribute Reference

The following additional attributes are exported:

* `description` - (string) The app's description.

* `name` - (string) The app's full name.

* `node_id` - (string) The Node ID of the app.

* `id` - (number) The installation ID.
