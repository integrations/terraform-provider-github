---
layout: "github"
page_title: "GitHub: github_enterprise_actions_runner_group"
description: |-
  Creates and manages an Actions Runner Group within a GitHub enterprise.
---

# github_enterprise_actions_runner_group

This resource allows you to create and manage GitHub Actions runner groups within your GitHub enterprise.
You must have admin access to an enterprise to use this resource.

## Example Usage

```hcl
data "github_enterprise" "enterprise" {
  slug = "my-enterprise"
}

resource "github_enterprise_organization" "enterprise_organization" {
  enterprise_id = data.github_enterprise.enterprise.id
  name          = "my-organization"
  billing_email = "octocat@octo.cat"
  admin_logins  = ["octocat"]
}

resource "github_enterprise_actions_runner_group" "example" {
  name                       = "my-awesome-runner-group"
  enterprise_slug            = data.github_enterprise.enterprise.slug
  allows_public_repositories = true
  visibility                 = "selected"
  selected_organization_ids  = [github_enterprise_organization.enterprise_organization.database_id]
  restricted_to_workflows    = true
  selected_workflows         = ["my-organization/my-repo/.github/workflows/cool-workflow.yaml@refs/tags/v1"]
}

resource "github_enterprise_network_configuration" "example" {
  enterprise_slug      = data.github_enterprise.enterprise.slug
  name                 = "my-network-config"
  compute_service      = "actions"
  network_settings_ids = ["23456789ABCDEF1"]
}

resource "github_enterprise_actions_runner_group" "private_networked" {
  enterprise_slug          = data.github_enterprise.enterprise.slug
  name                     = "private-networked-runners"
  visibility               = "all"
  network_configuration_id = github_enterprise_network_configuration.example.id
}
```

## Argument Reference

The following arguments are supported:
* `enterprise_slug`            - (Required) The slug of the enterprise.
* `name`                       - (Required) Name of the runner group
* `visibility`                 - (Required) Visibility of a runner group to enterprise organizations. Whether the runner group can include `all` or `selected`
* `selected_organization_ids`  - (Optional) IDs of the organizations which should be added to the runner group
* `allows_public_repositories` - (Optional) Whether public repositories can be added to the runner group. Defaults to false.
* `restricted_to_workflows`    - (Optional) If true, the runner group will be restricted to running only the workflows specified in the selected_workflows array. Defaults to false.
* `selected_workflows`         - (Optional) List of workflows the runner group should be allowed to run. This setting will be ignored unless restricted_to_workflows is set to true.
* `network_configuration_id`   - (Optional) The ID of a hosted compute network configuration to associate with this runner group. This is the GitHub-side linkage used for GitHub-hosted private networking.

## Attributes Reference

The following additional attributes are exported:

* `id`                         - The ID of the runner group
* `default`                    - Whether this is the default runner group
* `etag`                       - An etag representing the runner group object
* `runners_url`                - The GitHub API URL for the runner group's runners
* `selected_organizations_url` - The GitHub API URL for the runner group's selected organizations
* `network_configuration_id`   - The ID of the hosted compute network configuration associated with this runner group

## Notes

Use `github_enterprise_network_configuration` to manage the hosted compute network configuration, then set `network_configuration_id` on `github_enterprise_actions_runner_group` so GitHub-hosted runners assigned to that enterprise runner group use the private networking association.

## Import

This resource can be imported using the enterprise slug and the ID of the runner group:

```
$ terraform import github_enterprise_actions_runner_group.test enterprise-slug/42
```
