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
