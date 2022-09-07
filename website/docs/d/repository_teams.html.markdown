---
layout: "github"
page_title: "GitHub: github_repository_teams"
description: |-
  Get teams which have permission on the given repo.
---

# github\_repository\_teams

Use this data source to retrieve the list of teams which have access to a GitHub repository.

## Example Usage

```hcl
data "github_repository_teams" "example" {
  name = "example"
}
```

## Argument Reference

 * `name` - (Optional) The name of the repository.
 * `full_name` - (Optional) Full name of the repository (in `org/name` format).

## Attributes Reference

 * `teams` - List of teams which have access to the repository
   * `name` - Team name
   * `slug` - Team slug
   * `permission` - Team permission

