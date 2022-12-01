---
layout: "github"
page_title: "GitHub: github_repositories"
description: |-
  Search for GitHub repositories
---

# github_repositories

-> **Note:** The data source will return a maximum of `1000` repositories
	[as documented in official API docs](https://developer.github.com/v3/search/#about-the-search-api).

Use this data source to retrieve a list of GitHub repositories using a search query.

## Example Usage

```hcl
data "github_repositories" "example" {
  query = "org:hashicorp language:Go"
  include_repo_id = true
}
```

## Argument Reference

The following arguments are supported:

* `query` - (Required) Search query. See [documentation for the search syntax](https://help.github.com/articles/understanding-the-search-syntax/).
* `sort` - (Optional) Sorts the repositories returned by the specified attribute. Valid values include `stars`, `fork`, and `updated`. Defaults to `updated`.
* `include_repo_id` - (Optional) Returns a list of found repository IDs
* `results_per_page` - (Optional) Set the number of repositories requested per API call. Can be useful to decrease if requests are timing out or to increase to reduce the number of API calls. Defaults to 100.

## Attributes Reference

* `full_names` - A list of full names of found repositories (e.g. `hashicorp/terraform`)
* `names` - A list of found repository names (e.g. `terraform`)
* `repo_ids` - (Optional) A list of found repository IDs (e.g. `449898861`)
