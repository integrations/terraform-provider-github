data "github_enterprise" "enterprise" {
  slug = "my-enterprise"
}

data "github_organization" "org" {
  name = "my-organization"
}

data "github_repository" "repo1" {
  full_name = "${data.github_organization.org.name}/repo-1"
}

data "github_repository" "repo2" {
  full_name = "${data.github_organization.org.name}/repo-2"
}

# Create an enterprise runner group and share it with the organization
resource "github_enterprise_actions_runner_group" "example" {
  name                      = "my-runner-group"
  enterprise_slug           = data.github_enterprise.enterprise.slug
  visibility                = "selected"
  selected_organization_ids = [data.github_organization.org.id]
}

# Configure repository access for the runner group in the organization
resource "github_organization_inherited_runner_group_settings" "example" {
  organization                 = data.github_organization.org.name
  enterprise_runner_group_name = github_enterprise_actions_runner_group.example.name
  selected_repository_ids      = [data.github_repository.repo1.id, data.github_repository.repo2.id]
  allows_public_repositories   = true
  restricted_to_workflows      = true
  selected_workflows           = ["${data.github_repository.repo1.full_name}/.github/workflows/ci.yml@refs/heads/main"]
}
