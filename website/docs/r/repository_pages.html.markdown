---
layout: "github"
page_title: "GitHub: github_repository_pages"
description: |-
  Manages GitHub Pages for a repository
---

# github_repository_pages

This resource allows you to manage GitHub Pages for a repository. See the
[documentation](https://docs.github.com/en/pages/getting-started-with-github-pages/about-github-pages)
for details on GitHub Pages.

The authenticated user must be a repository administrator, maintainer, or have the 'manage GitHub Pages settings' permission. OAuth app tokens and personal access tokens (classic) need the repo scope to use this resource.

## Example Usage

### Legacy Build Type

```hcl
resource "github_repository" "example" {
  name       = "my-repo"
  visibility = "public"
  auto_init  = true

  lifecycle {
    ignore_changes = [
      pages,
    ]
  }
}

resource "github_repository_pages" "example" {
  repository = github_repository.example.name
  build_type = "legacy"

  source {
    branch = "main"
    path   = "/"
  }
}
```

### Workflow Build Type (GitHub Actions)

```hcl
resource "github_repository" "example" {
  name       = "my-repo"
  visibility = "public"
  auto_init  = true

  lifecycle {
    ignore_changes = [
      pages,
    ]
  }
}

resource "github_repository_pages" "example" {
  repository = github_repository.example.name
  build_type = "workflow"
}
```

### With Custom Domain

```hcl
resource "github_repository" "example" {
  name       = "my-repo"
  visibility = "public"
  auto_init  = true

  lifecycle {
    ignore_changes = [
      pages,
    ]
  }
}

resource "github_repository_pages" "example" {
  repository = github_repository.example.name
  build_type = "legacy"
  cname          = "example.com"
  https_enforced = true

  source {
    branch = "main"
    path   = "/docs"
  }
}
```

## Argument Reference

The following arguments are supported:

- `repository` - (Required) The repository name to configure GitHub Pages for.

- `build_type` - (Optional) The type of GitHub Pages site to build. Can be `legacy` or `workflow`. Defaults to `legacy`.

- `source` - (Optional) The source branch and directory for the rendered Pages site. Required when `build_type` is `legacy`. See [Source](#source) below for details.

- `cname` - (Optional) The custom domain for the repository.

- `public` - (Optional) Whether the GitHub Pages site is public.

- `https_enforced` - (Optional) Whether HTTPS is enforced for the GitHub Pages site. GitHub Pages sites serve over HTTPS by default; this setting only applies when a custom domain (`cname`) is configured. Requires `cname` to be set.

### Source

The `source` block supports the following:

- `branch` - (Required) The repository branch used to publish the site's source files (e.g., `main` or `gh-pages`).

- `path` - (Optional) The repository directory from which the site publishes. Defaults to `/`. Can be `/` or `/docs`.

## Attribute Reference

In addition to the above arguments, the following attributes are exported:

- `repository_id` - The ID of the repository.

- `custom_404` - Whether the rendered GitHub Pages site has a custom 404 page.

- `html_url` - The absolute URL (with scheme) to the rendered GitHub Pages site.

- `build_status` - The GitHub Pages site's build status (e.g., `building` or `built`).

- `api_url` - The API URL of the GitHub Pages resource.

## Import

GitHub repository pages can be imported using the `repository-slug`, e.g.

```sh
terraform import github_repository_pages.example my-repo
```
