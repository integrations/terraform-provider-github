---
page_title: "github_enterprise_actions_runner_group (Resource) - GitHub"
description: |-
  Creates and manages an Actions Runner Group within a GitHub enterprise.
---

# github_enterprise_actions_runner_group (Resource)

This resource allows you to create and manage GitHub Actions runner groups within your GitHub enterprise. You must have admin access to an enterprise to use this resource.

Organization access and repository access are separate controls. After making the group available to an organization, ensure that organization grants its repositories access to the inherited group. The [organization repository-access API](https://docs.github.com/en/rest/actions/self-hosted-runner-groups#set-repository-access-for-a-runner-group-in-an-organization) uses the inherited group's organization-local ID, which can differ from this resource's enterprise group ID.

The [github_actions_hosted_runner](actions_hosted_runner) resource cannot provision runners in enterprise-owned groups. Enterprise-hosted runners must be provisioned separately, for example through the [enterprise hosted-runner API](https://docs.github.com/enterprise-cloud@latest/rest/actions/hosted-runners#create-a-github-hosted-runner-for-an-enterprise).

## Example Usage

```terraform
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
```

## Argument Reference

The following arguments are supported:

- `enterprise_slug` - (Required) The slug of the enterprise.
- `name` - (Required) Name of the runner group
- `visibility` - (Required) Visibility of a runner group to enterprise organizations. Whether the runner group can include `all` or `selected`
- `selected_organization_ids` - (Optional) IDs of the organizations which should be added to the runner group
- `allows_public_repositories` - (Optional) Whether public repositories can be added to the runner group. Defaults to false.
- `restricted_to_workflows` - (Optional) If true, the runner group will be restricted to running only the workflows specified in the selected_workflows array. Defaults to false.
- `selected_workflows` - (Optional) List of workflows the runner group should be allowed to run. This setting will be ignored unless restricted_to_workflows is set to true.
- `network_configuration_id` - (Optional, Computed) The ID of a `github_enterprise_network_configuration` to assign to the runner group, so GitHub-hosted runners in the group run inside your private network. When omitted or set to `null`, the existing assignment is retained. Set this to `""` to clear the assignment. **Clearing currently replaces the runner group** because the GitHub client library cannot encode the explicit null required for in-place removal.

## Attributes Reference

The following additional attributes are exported:

- `id` - The ID of the runner group
- `default` - Whether this is the default runner group
- `etag` - An etag representing the runner group object
- `runners_url` - The GitHub API URL for the runner group's runners
- `network_configuration_id` - The ID of the hosted compute network configuration assigned to the runner group
- `selected_organizations_url` - The GitHub API URL for the runner group's selected organizations

## Import

This resource can be imported using the enterprise slug and the ID of the runner group:

```shell
terraform import github_enterprise_actions_runner_group.test enterprise-slug/42
```
