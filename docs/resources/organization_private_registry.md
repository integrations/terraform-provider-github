---
page_title: "github_organization_private_registry (Resource) - GitHub"
description: |-
  Creates and manages an organization private registry.
---

# github_organization_private_registry (Resource)

This resource allows you to create and manage an organization's private registry. Centralized private registry configuration for Dependabot allows you to configure rules and credentials for registries at the organization level, replacing the need for repository-level configurations.

## Example Usage

```terraform
resource "github_organization_private_registry" "my_registry" {
  registry_type = "npm_registry"
  url           = "https://npm.pkg.github.com"
  auth_type     = "username_password"
  username      = "github-actions"
  secret        = "super_secret_token_123"
  visibility    = "private"
}
```

## Argument Reference

The following arguments are supported:

- `name` - (Required) The name of the private registry.

- `registry_type` - (Required) The registry type. Can be `maven_repository`, `nuget_feed`, `goproxy_server`, `npm_registry`, `rubygems_server`, `cargo_registry`, `composer_repository`, `docker_registry`, `git_source`, `helm_registry`, `pub_repository`, `python_index`, or `terraform_registry`.
- `visibility` - (Required) Configures the access that repositories have to the organization private registry. Must be one of `all`, `private`, or `selected`.
- `url` - (Required) The URL of the private registry.
- `auth_type` - (Optional) The authentication type for the private registry. Can be `token`, `username_password`, `oidc_azure`, `oidc_aws`, or `oidc_jfrog`. Defaults to `token`.
- `username` - (Optional) The username to use when authenticating with the private registry.
- `secret` - (Optional) The secret/password to use when authenticating with the private registry. This will be encrypted locally before sending to GitHub.
- `key_id` - (Optional) ID of the public key used to encrypt the secret. Required if `encrypted_value` is set directly.
- `encrypted_value` - (Optional) The encrypted value of the secret using the GitHub public key in Base64 format.
- `replaces_base` - (Optional) Whether this registry should replace a base registry configuration.
- `selected_repository_ids` - (Optional) An array of repository IDs that can access the organization private registry. Required when `visibility` is set to `selected`.
- `oidc_audience` - (Optional) The OIDC audience.
- `oidc_azure_tenant_id` - (Optional) The Azure tenant ID. Required when `auth_type` is `oidc_azure`.
- `oidc_azure_client_id` - (Optional) The Azure client ID. Required when `auth_type` is `oidc_azure`.
- `oidc_aws_account_id` - (Optional) The AWS account ID. Required when `auth_type` is `oidc_aws`.
- `oidc_aws_region` - (Optional) The AWS region. Required when `auth_type` is `oidc_aws`.
- `oidc_aws_role_name` - (Optional) The AWS role name. Required when `auth_type` is `oidc_aws`.
- `oidc_aws_domain` - (Optional) The AWS domain. Required when `auth_type` is `oidc_aws`.
- `oidc_aws_domain_owner` - (Optional) The AWS domain owner. Required when `auth_type` is `oidc_aws`.
- `oidc_jfrog_provider_name` - (Optional) The JFrog provider name. Required when `auth_type` is `oidc_jfrog`.
- `oidc_jfrog_identity_mapping_name` - (Optional) The JFrog identity mapping name. Required when `auth_type` is `oidc_jfrog`.

## Attributes Reference

- `id` - The ID of the private registry.
- `created_at` - Timestamp of when the private registry was created.
- `updated_at` - Timestamp of when the private registry was last updated.
