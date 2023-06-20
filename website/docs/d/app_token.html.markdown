---
layout: "github"
page_title: "GitHub: github_app_token"
description: |-
  Generate a GitHub APP JWT.
---

# github\_app\_token

Use this data source to generate a [GitHub App JWT](https://docs.github.com/en/apps/creating-github-apps/authenticating-with-a-github-app/generating-a-json-web-token-jwt-for-a-github-app).

## Example Usage

```hcl
data "github_app_token" "this" {
  app_id          = "123456"
  installation_id = "78910"
  pem_file        = file("foo/bar.pem")
}
```

## Argument Reference

The following arguments are supported:

* `app_id` - (Required) This is the ID of the GitHub App.

* `installation_id` - (Required) This is the ID of the GitHub App installation.

* `pem_file` - (Required) This is the contents of the GitHub App private key PEM file.

## Attribute Reference

The following additional attributes are exported:

* `token` - The generated GitHub APP JWT.
