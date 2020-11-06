---
layout: "github"
page_title: "GitHub: github_repository_page"
description: |-
  Creates and manages a GitHub Pages for a GitHub repository
---

# github_repository_page

This resource allows you to create and manage a GitHub Page for a GitHub repository.

## Example Usage

```hcl

resource "github_repository" "example" {
  name         = "example"
  description  = "My awesome codebase"
}

resource "github_repository_page" "example" {
  repository = github_repository.example.name
  source {
    path = "docs"
    branch = "main"
  }
}

```

## Argument Reference

* `repository` - (Required) The repository of the page.

* `source` - (Required) The source branch and directory for the rendered Pages site. See [Github Pages Source](#github-pages-source) below for details.

* `cname` - (Optional) The custom domain for the repository.

#### Github Pages Source ####

The `source` block supports the following:

* `branch` - (Required) The repository branch used to publish the site's source files. Can be `master` or `gh-pages`.

* `path` - (Optional) The repository directory from which the site publishes. Can be `/` or `/docs`. (Default: `/`).


## Attributes Reference

The following additional attributes are exported:

* `url` - URL of the page
