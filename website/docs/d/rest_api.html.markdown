---
layout: "github"
page_title: "GitHub: github_rest_api"
description: |-
  Get information on a GitHub resource with a custom GET request to GitHub REST API.
---

# github_rest_api

Use this data source to retrieve information about a GitHub resource through REST API.

## Example Usage

```hcl
data "github_rest_api" "example" {
  endpoint = "repos/example_repo/git/refs/heads/main"
}
```

## Argument Reference

 * `endpoint` - (Required) REST API endpoint to send the GET request to.

## Attributes Reference

 * `id`       - The GitHub API Request ID
 * `code`     - A response status code.
 * `status`   - A response status string.
 * `headers`  - A JSON string containing response headers.
 * `body`     - A JSON string containing response body.
