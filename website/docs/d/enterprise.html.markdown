---
layout: "github"
page_title: "Github: github_enterprise"
description: |-
  Get an enterprise.
---

# github_enterprise

Use this data source to retrieve basic information about a GitHub enterprise.

## Example Usage

```
data "github_enteprise" "example" {
  slug = "example-co"
}
```

## Attributes Reference

* `id` - The ID of the enterprise.
* `slug` - The URL slug identifying the enterprise.
* `name` - The name of the enteprise.
* `description` - The description of the enterpise.
* `created_at` - The time the enterprise was created.
* `url` - The url for the enterprise.
