---
page_title: "github_client_config (Data Source) - GitHub"
description: |-
  Get information about the configured GitHub provider client, including the resolved owner and authenticated user.
---

# github_client_config (Data Source)

Use this data source to access information about the GitHub provider's client
configuration. This is useful when the `owner` argument is set via the
`GITHUB_OWNER` environment variable (and therefore not available in
configuration), or when you need to know which user the provider is
authenticating as.

## Example Usage

```terraform
data "github_client_config" "current" {}

output "owner" {
  value = data.github_client_config.current.owner
}

output "username" {
  value = data.github_client_config.current.username
}
```

## Argument Reference

This data source has no arguments.

## Attributes Reference

- `id` - The resolved owner name, or the authenticated user's login when no
  owner is configured. Falls back to the API base URL in anonymous mode.
- `owner` - The owner the provider is configured to manage. This reflects the
  value of the `owner` provider argument or the `GITHUB_OWNER` environment
  variable. When neither is set and the provider is authenticated, this is the
  login of the authenticated user.
- `is_organization` - Whether the resolved `owner` is a GitHub organization
  (`true`) or a user (`false`).
- `username` - The login of the user the provider is authenticated as. This
  may differ from `owner` when `owner` is an organization or another user.
  Empty when the provider is configured in anonymous mode.
- `base_url` - The GitHub API base URL the provider is configured to use.
