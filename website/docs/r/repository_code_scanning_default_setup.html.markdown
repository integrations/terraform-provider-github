---
layout: "github"
page_title: "GitHub: github_repository_code_scanning_default_setup"
description: |-
  Manages code scanning default setup for a repository
---

# github_repository_code_scanning_default_setup

This resource allows you to manage the code scanning default setup configuration for a GitHub repository.
When enabled, GitHub automatically configures CodeQL analysis for supported languages in the repository.

See the [documentation](https://docs.github.com/en/code-security/code-scanning/enabling-code-scanning/configuring-default-setup-for-code-scanning)
for details of usage and how this will impact your repository.

## Example Usage

### Basic usage

```hcl
resource "github_repository" "example" {
  name       = "my-repo"
  visibility = "public"
  auto_init  = true
}

resource "github_repository_code_scanning_default_setup" "example" {
  repository = github_repository.example.name
  state      = "configured"
}
```

### With query suite and languages

```hcl
resource "github_repository" "example" {
  name       = "my-repo"
  visibility = "public"
  auto_init  = true
}

resource "github_repository_code_scanning_default_setup" "example" {
  repository  = github_repository.example.name
  state       = "configured"
  query_suite = "extended"
  languages   = ["javascript-typescript", "python"]
}
```

## Argument Reference

The following arguments are supported:

* `repository` - (Required) The name of the GitHub repository.

* `state` - (Required) The desired state of code scanning default setup. Must be `configured` or `not-configured`. See the [REST API docs](https://docs.github.com/en/rest/code-scanning/code-scanning#update-a-code-scanning-default-setup-configuration).

* `query_suite` - (Optional) The [query suite](https://docs.github.com/en/code-security/code-scanning/managing-your-code-scanning-configuration/codeql-query-suites) to use. Must be `default` or `extended`.

* `languages` - (Optional) The CodeQL languages to be analyzed. If not specified, default setup [automatically includes all supported languages](https://github.blog/changelog/2023-10-23-code-scanning-default-setup-automatically-includes-all-codeql-supported-languages/). Supported values: `actions`, `c-cpp`, `csharp`, `go`, `java-kotlin`, `javascript-typescript`, `python`, `ruby`, `swift`. See the [REST API docs](https://docs.github.com/en/rest/code-scanning/code-scanning#update-a-code-scanning-default-setup-configuration).

## Import

Code scanning default setup can be imported using the repository name:

```sh
terraform import github_repository_code_scanning_default_setup.example my-repo
```
