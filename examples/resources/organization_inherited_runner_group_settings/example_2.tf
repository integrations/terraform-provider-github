data "github_enterprise" "enterprise" {
  slug = "my-enterprise"
}

data "github_organization" "org" {
  name = "my-organization"
}

resource "github_enterprise_actions_runner_group" "example" {
  name                      = "my-runner-group"
  enterprise_slug           = data.github_enterprise.enterprise.slug
  visibility                = "selected"
  selected_organization_ids = [data.github_organization.org.id]
}

# Make the runner group available to all repositories in the organization
resource "github_organization_inherited_runner_group_settings" "example" {
  organization                 = data.github_organization.org.name
  enterprise_runner_group_name = github_enterprise_actions_runner_group.example.name
  visibility                   = "all"
}
