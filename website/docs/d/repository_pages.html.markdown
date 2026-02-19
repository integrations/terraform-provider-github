---
layout: "github"
page_title: "GitHub: github_repository_pages"
description: |-
  Get information on GitHub Pages for a repository
---

# github_repository_pages

Use this data source to retrieve GitHub Pages configuration for a repository.

## Example Usage

```hcl
data "github_repository_pages" "example" {
  repository = "my-repo"
}
```

## Argument Reference

The following arguments are supported:

- `repository` - (Required) The repository name to get GitHub Pages information for.

## Attribute Reference

The following attributes are exported:

- `build_type` - The type of GitHub Pages site. Can be `legacy` or `workflow`.

- `cname` - The custom domain for the repository.

- `custom_404` - Whether the rendered GitHub Pages site has a custom 404 page.

- `html_url` - The absolute URL (with scheme) to the rendered GitHub Pages site.

- `source` - The source branch and directory for the rendered Pages site. See [Source](#source) below for details.

- `build_status` - The GitHub Pages site's build status (e.g., `building` or `built`).

- `api_url` - The API URL of the GitHub Pages resource.

- `public` - Whether the GitHub Pages site is public.

### Source

The `source` block contains:

- `branch` - The repository branch used to publish the site's source files.

- `path` - The repository directory from which the site publishes.
