---
page_title: "github_enterprise Data Source - terraform-provider-github
description: |-
  Get an enterprise.
---

# github_enterprise (Data Source)

Use this data source to retrieve basic information about a GitHub enterprise.

## Example Usage

```
data "github_enterprise" "example" {
  slug = "example-co"
}
```

## Attributes Reference

- `id` - The ID of the enterprise.
- `database_id` - The database ID of the enterprise.
- `slug` - The URL slug identifying the enterprise.
- `name` - The name of the enterprise.
- `description` - The description of the enterprise.
- `created_at` - The time the enterprise was created.
- `url` - The url for the enterprise.
